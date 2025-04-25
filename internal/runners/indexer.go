package runners

import (
	"context"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"github.com/zh0glikk/eth-tx-parser/pkg/interfaces"
	"log/slog"
	"sync"
	"time"
)

type indexer struct {
	blocksUse       usecases.BlocksUseCase
	transactionsUse usecases.TransactionsUseCase
	blockParserUse  usecases.BlockParserUseCase

	logger *slog.Logger
}

func newIndexer(
	logger *slog.Logger,
	blocksUse usecases.BlocksUseCase,
	transactionsUse usecases.TransactionsUseCase,
	blockParserUse usecases.BlockParserUseCase,
) interfaces.Runner {
	return &indexer{
		blocksUse:       blocksUse,
		transactionsUse: transactionsUse,
		blockParserUse:  blockParserUse,
		logger:          logger,
	}
}

func (i *indexer) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	i.logger.Info("indexer started")

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			i.logger.Info("indexer stopped")
			return
		case <-ticker.C:
			number, err := i.blockParserUse.GetBlock2Process()
			if err != nil {
				i.logger.Error("failed to get block to process", slog.String("err", err.Error()))
				continue
			}
			i.logger.Info("got block to process", slog.Uint64("number", number))

			err = i.processBlock(number)
			if err != nil {
				if err.Error() == "not found" {
					i.logger.Info("block not found, skipping", slog.Uint64("number", number))
					time.Sleep(time.Second * 3)
				}
				continue
			}
		}
	}
}

func (i *indexer) processBlock(number uint64) error {
	txs2store, err := i.blockParserUse.ParseBlockTransactions(number)
	if err != nil {
		return err
	}

	for _, tx := range txs2store {

		i.logger.Info("found tx",
			slog.Uint64("block", tx.Block),
			slog.String("from", tx.From),
			slog.String("to", tx.To),
			slog.String("value", tx.Amount.String()),
			slog.String("token", tx.Token),
			slog.String("hash", tx.Hash),
		)

		err = i.transactionsUse.Create(tx)
		if err != nil {
			return err
		}
	}

	err = i.blocksUse.Create(models.CreateBlockRequest{Number: number})
	if err != nil {
		return err
	}

	i.logger.Info("block processed", slog.Uint64("number", number))

	return nil
}
