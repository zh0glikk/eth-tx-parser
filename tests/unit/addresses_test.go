package unit

import (
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos/memory"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"testing"
)

var (
	addressesUse usecases.AddressesUseCase
)

func initAddressesTest() {
	addressesUse = usecases.NewAddressesUseCase(memory.NewAddressRepo())
}

func TestAddresses_Create(t *testing.T) {
	initAddressesTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: sender,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddresses_Create_Ensure_Exists(t *testing.T) {
	initAddressesTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: sender,
	})
	if err != nil {
		t.Fatal(err)
	}

	ok, _ := addressesUse.IsSubscribed(sender)
	if !ok {
		t.Fatal("address should exist")
	}

	ok, _ = addressesUse.IsSubscribed(sender2)
	if ok {
		t.Fatal("address should not exist")
	}

}

func TestAddresses_Create_Multiple_Select_Page(t *testing.T) {
	initAddressesTest()

	for i := 0; i < 100; i++ {
		err := addressesUse.Subscribe(models.CreateAddressRequest{
			Address: fmt.Sprintf("%s%d", sender, i),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < 100; i++ {
		ok, err := addressesUse.IsSubscribed(fmt.Sprintf("%s%d", sender, i))
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			t.Fatal("address should exist")
		}
	}

}
