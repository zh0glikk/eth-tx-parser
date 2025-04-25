package runners

import (
	"context"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"github.com/zh0glikk/eth-tx-parser/pkg/interfaces"
	"log/slog"
	"sync"
)

func Start(
	appCtx context.Context,
	appWg *sync.WaitGroup,
	logger *slog.Logger,
	blocksUse usecases.BlocksUseCase,
	transactionsUse usecases.TransactionsUseCase,
	blockParserUse usecases.BlockParserUseCase,
) {

	runners2start := []interfaces.Runner{
		newIndexer(
			logger,
			blocksUse,
			transactionsUse,
			blockParserUse,
		),
	}

	for i := range runners2start {
		appWg.Add(1)
		go runners2start[i].Run(appCtx, appWg)
	}

}
