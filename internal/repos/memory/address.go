package memory

import (
	"errors"
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
	"sync"
)

type addressRepo struct {
	mu   sync.RWMutex
	data []entities.Address

	//for unique addresses
	unique map[string]struct{}
}

func NewAddressRepo() repos.AddressesRepo {
	return &addressRepo{
		mu:     sync.RWMutex{},
		data:   nil,
		unique: make(map[string]struct{}),
	}
}

func (t *addressRepo) Create(address entities.Address) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.unique[address.Address]; ok {
		return errors.New(fmt.Sprintf("row with such address already exists: %s", address.Address))
	}

	address.ID = uint64(len(t.data) + 1)
	t.data = append(t.data, address)
	t.unique[address.Address] = struct{}{}

	return nil
}

func (t *addressRepo) IsExist(address string) (bool, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, ok := t.unique[address]
	return ok, nil
}

func (t *addressRepo) Select(req models.SearchAddressesRequest) ([]entities.Address, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return getPage(req.PageMetadata, t.data), nil
}
