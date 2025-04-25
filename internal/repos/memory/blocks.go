package memory

import (
	"errors"
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
	"sync"
)

type blockRepo struct {
	mu   sync.RWMutex
	data []entities.Block

	//for unique blocks
	unique map[uint64]struct{}
}

func NewBlockRepo() repos.BlocksRepo {
	return &blockRepo{
		mu:     sync.RWMutex{},
		data:   nil,
		unique: make(map[uint64]struct{}),
	}
}

func (t *blockRepo) Create(block entities.Block) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.unique[block.Number]; ok {
		return errors.New(fmt.Sprintf("row with such block already exists: %s", block.Number))
	}

	block.ID = uint64(len(t.data) + 1)
	t.data = append(t.data, block)
	t.unique[block.Number] = struct{}{}

	return nil
}

func (t *blockRepo) Select(req models.SearchBlocksRequest) ([]entities.Block, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return getPage(req.PageMetadata, t.data), nil
}
