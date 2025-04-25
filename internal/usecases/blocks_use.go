package usecases

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
)

type BlocksUseCase interface {
	Create(data models.CreateBlockRequest) error

	GetLastBlock() (*models.BlockResponse, error)
}

type blocksUse struct {
	repo repos.BlocksRepo
}

func NewBlocksUseCase(
	repo repos.BlocksRepo,
) BlocksUseCase {
	return &blocksUse{
		repo: repo,
	}
}

func (b *blocksUse) Create(data models.CreateBlockRequest) error {
	return b.repo.Create(entities.Block{
		Number: data.Number,
	})
}

func (b *blocksUse) GetLastBlock() (*models.BlockResponse, error) {
	result, err := b.repo.Select(models.SearchBlocksRequest{
		PageMetadata: models.PageMetadata{
			Size:     1,
			OrderDir: "desc",
		},
	})
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}

	return &models.BlockResponse{
		ID:     result[0].ID,
		Number: result[0].Number,
	}, nil
}
