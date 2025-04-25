package main

import (
	"context"
	"github.com/zh0glikk/eth-tx-parser/internal/delivery/api"
	"github.com/zh0glikk/eth-tx-parser/internal/delivery/api/rest"
	"github.com/zh0glikk/eth-tx-parser/internal/repos/memory"
	"github.com/zh0glikk/eth-tx-parser/internal/runners"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"github.com/zh0glikk/eth-tx-parser/pkg/ethclient"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	appWg := new(sync.WaitGroup)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	ethCli := ethclient.NewClient(os.Getenv("RPC_URL"))

	blocksRepo := memory.NewBlockRepo()
	addressRepo := memory.NewAddressRepo()
	transactionRepo := memory.NewTransactionRepo()

	blocksUse := usecases.NewBlocksUseCase(blocksRepo)
	addressesUse := usecases.NewAddressesUseCase(addressRepo)
	transactionsUse := usecases.NewTransactionsUseCase(transactionRepo)
	blockParserUse := usecases.NewBlockParserUseCase(ethCli, addressesUse, blocksUse)

	runners.Start(
		appCtx,
		appWg,
		logger,
		blocksUse,
		transactionsUse,
		blockParserUse,
	)

	router := rest.NewRouter(
		addressesUse,
		blocksUse,
		transactionsUse,
	)
	rs := api.NewServer(router.InitRoutes(), os.Getenv("REST_PORT"))

	go func() {
		if err := rs.Run(); err != nil {
			logger.Error("failed to start http server", slog.String("error", err.Error()))
		}
	}()

	gracefulShutdown(
		logger,
		appCtx,
		cancel,
		appWg,
		rs,
	)
}

func gracefulShutdown(
	logger *slog.Logger,
	appCtx context.Context,
	cancel context.CancelFunc,
	appWg *sync.WaitGroup,
	rs *api.Server,
) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	cancel()

	// wait for all goroutines finished
	appWg.Wait()
	logger.Info("all workers stopped")

	err := rs.Stop(appCtx)
	if err != nil {
		logger.Error("REST server stopped", slog.String("error", err.Error()))
	}

	logger.Info("REST server stopped")

}
