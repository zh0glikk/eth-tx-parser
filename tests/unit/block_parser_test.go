package unit

import (
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"github.com/zh0glikk/eth-tx-parser/tests/unit/mocks"
	"strings"
	"testing"
)

var (
	blockParserUse usecases.BlockParserUseCase
)

func initBlockParserTest() {
	initBlocksTest()
	initAddressesTest()

	client := mocks.NewEthClientMock()

	blockParserUse = usecases.NewBlockParserUseCase(client, addressesUse, blocksUse)

}

func TestParseTransactions_Eth_transfer_by_from(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000001",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000001" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000002" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.EthTransfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "15" {
		t.Fatal("invalid amount")
	}
}

func TestParseTransactions_Eth_transfer_by_to(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000002",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000001" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000002" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.EthTransfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "15" {
		t.Fatal("invalid amount")
	}
}

func TestParseTransactions_Usdt_transfer_by_from(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000003",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000003" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000004" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.Erc20Transfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "1000000" {
		t.Fatal("invalid amount")
	}
	if transactions[0].Token != strings.ToLower(models.UsdtAddress) {
		t.Fatal("invalid token")
	}
}

func TestParseTransactions_Usdt_transfer_by_to(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000004",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000003" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000004" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.Erc20Transfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "1000000" {
		t.Fatal("invalid amount")
	}
	if transactions[0].Token != strings.ToLower(models.UsdtAddress) {
		t.Fatal("invalid token")
	}
}

func TestParseTransactions_Usdc_transfer_from_by_from(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000005",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000005" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000006" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.Erc20Transfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "1000000" {
		t.Fatal("invalid amount")
	}
	if transactions[0].Token != strings.ToLower(models.UsdcAddress) {
		t.Fatal("invalid token")
	}
}

func TestParseTransactions_Usdc_transfer_from_by_to(t *testing.T) {
	initBlockParserTest()

	err := addressesUse.Subscribe(models.CreateAddressRequest{
		Address: "0x0000000000000000000000000000000000000006",
	})
	if err != nil {
		t.Fatal(err)
	}

	transactions, err := blockParserUse.ParseBlockTransactions(100)
	if err != nil {
		t.Fatal(err)
	}

	if len(transactions) != 1 {
		t.Fatal("invalid amount of transactions")
	}
	if transactions[0].From != "0x0000000000000000000000000000000000000005" {
		t.Fatal("invalid from")
	}
	if transactions[0].To != "0x0000000000000000000000000000000000000006" {
		t.Fatal("invalid to")
	}
	if transactions[0].Type != models.Erc20Transfer {
		t.Fatal("invalid transaction type")
	}
	if transactions[0].Amount.String() != "1000000" {
		t.Fatal("invalid amount")
	}
	if transactions[0].Token != strings.ToLower(models.UsdcAddress) {
		t.Fatal("invalid token")
	}
}
