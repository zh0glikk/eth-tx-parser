package repos

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
)

type AddressesRepo interface {
	Create(address entities.Address) error

	IsExist(address string) (bool, error)

	Select(req models.SearchAddressesRequest) ([]entities.Address, error)
}
