package usecases

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
	"strings"
)

type AddressesUseCase interface {
	Subscribe(data models.CreateAddressRequest) error
	IsSubscribed(address string) (bool, error)
}

type addressesUse struct {
	repo repos.AddressesRepo
}

func NewAddressesUseCase(repo repos.AddressesRepo) AddressesUseCase {
	return &addressesUse{
		repo: repo,
	}
}

func (uc *addressesUse) Subscribe(data models.CreateAddressRequest) error {
	address := strings.ToLower(data.Address)

	ok, err := uc.repo.IsExist(address)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	err = uc.repo.Create(entities.Address{
		Address: address,
	})

	return err
}

func (uc *addressesUse) IsSubscribed(address string) (bool, error) {
	ok, err := uc.repo.IsExist(strings.ToLower(address))
	if err != nil {
		return false, err
	}
	return ok, nil
}
