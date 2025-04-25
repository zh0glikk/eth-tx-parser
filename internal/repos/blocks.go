package repos

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
)

type BlocksRepo interface {
	Create(block entities.Block) error

	Select(req models.SearchBlocksRequest) ([]entities.Block, error)
}
