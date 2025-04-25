package unit

import (
	"errors"
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	memory2 "github.com/zh0glikk/eth-tx-parser/internal/repos/memory"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"github.com/zh0glikk/eth-tx-parser/pkg/utils"
	"math/big"
	"sync"
	"testing"
	"time"
)

var (
	sender    string = "eoa_1"
	receiver  string = "eoa_2"
	sender2   string = "eoa_3"
	receiver2 string = "eoa_4"

	uniqueHash string = "unique_hash"
)

var (
	transactionsUse usecases.TransactionsUseCase
)

func initTransactionsTest() {
	transactionsUse = usecases.NewTransactionsUseCase(memory2.NewTransactionRepo())
}

func TestTransactions_Create(t *testing.T) {
	initTransactionsTest()

	err := transactionsUse.Create(models.CreateTransactionRequest{
		Block:     1,
		BlockTime: time.Now().UTC().Unix(),
		Hash:      uniqueHash,
		From:      sender,
		To:        receiver,
		Amount:    big.NewInt(100),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactions_Create_Duplication(t *testing.T) {
	initTransactionsTest()

	err := transactionsUse.Create(models.CreateTransactionRequest{
		Block:     1,
		BlockTime: time.Now().UTC().Unix(),
		Hash:      uniqueHash,
		From:      sender,
		To:        receiver,
		Amount:    big.NewInt(100),
	})
	if err != nil {
		t.Fatal(err)
	}

	err = transactionsUse.Create(models.CreateTransactionRequest{
		Block:     1,
		BlockTime: time.Now().UTC().Unix(),
		Hash:      uniqueHash,
		From:      sender,
		To:        receiver,
		Amount:    big.NewInt(100),
	})
	if err == nil {
		t.Fatal(errors.New("should fail"))
	}
}

func TestTransactions_Create_Select_By_From(t *testing.T) {
	initTransactionsTest()

	err := transactionsUse.Create(models.CreateTransactionRequest{
		Block:     1,
		BlockTime: time.Now().UTC().Unix(),
		Hash:      uniqueHash,
		From:      sender,
		To:        receiver,
		Amount:    big.NewInt(100),
	})
	if err != nil {
		t.Fatal(err)
	}

	result, err := transactionsUse.Select(models.SearchTransactionsRequest{
		From: utils.Ptr(sender),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		t.Fatal("should return one transaction")
	}
	if result[0].Hash != uniqueHash {
		t.Fatal("hashes are not equal")
	}
}

func TestTransactions_Create_Select_By_To(t *testing.T) {
	initTransactionsTest()

	err := transactionsUse.Create(models.CreateTransactionRequest{
		Block:     1,
		BlockTime: time.Now().UTC().Unix(),
		Hash:      uniqueHash,
		From:      sender,
		To:        receiver,
		Amount:    big.NewInt(100),
	})
	if err != nil {
		t.Fatal(err)
	}

	result, err := transactionsUse.Select(models.SearchTransactionsRequest{
		To: utils.Ptr(receiver),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 1 {
		t.Fatal("should return one transaction")
	}
	if result[0].Hash != uniqueHash {
		t.Fatal("hashes are not equal")
	}
}

func TestTransactions_Create_Multiple_Select_By_Address(t *testing.T) {
	initTransactionsTest()

	for i := 0; i < 100; i++ {
		err := transactionsUse.Create(models.CreateTransactionRequest{
			Block:     1,
			BlockTime: time.Now().UTC().Unix(),
			Hash:      fmt.Sprintf("%s%d_1", uniqueHash, i),
			From:      sender,
			To:        receiver,
			Amount:    big.NewInt(100),
		})
		if err != nil {
			t.Fatal(err)
		}
		err = transactionsUse.Create(models.CreateTransactionRequest{
			Block:     1,
			BlockTime: time.Now().UTC().Unix(),
			Hash:      fmt.Sprintf("%s%d_2", uniqueHash, i),
			From:      sender2,
			To:        receiver2,
			Amount:    big.NewInt(100),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	result, err := transactionsUse.Select(models.SearchTransactionsRequest{
		Address: utils.Ptr(receiver),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 100 {
		t.Fatal("should return 100 transactions")
	}
}

func TestTransactions_Create_Multiple_Select_Page(t *testing.T) {
	initTransactionsTest()

	for i := 0; i < 100; i++ {
		err := transactionsUse.Create(models.CreateTransactionRequest{
			Block:     1,
			BlockTime: time.Now().UTC().Unix(),
			Hash:      fmt.Sprintf("%s%d_1", uniqueHash, i),
			From:      sender,
			To:        receiver,
			Amount:    big.NewInt(100),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	page := 0
	size := 12
	for {
		result, err := transactionsUse.Select(models.SearchTransactionsRequest{
			PageMetadata: models.PageMetadata{
				Size: size,
				Page: page,
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		if page == 8 && len(result) == 4 {
			//success case
			break
		}

		if len(result) != size {
			t.Fatal("should return 10 transactions")
		}
		page += 1
	}

}

func TestTransactions_Create_Concurrent_Multiple_Select_By_Address(t *testing.T) {
	initTransactionsTest()

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := transactionsUse.Create(models.CreateTransactionRequest{
				Block:     1,
				BlockTime: time.Now().UTC().Unix(),
				Hash:      fmt.Sprintf("%s%d_1", uniqueHash, i),
				From:      sender,
				To:        receiver,
				Amount:    big.NewInt(100),
			})
			if err != nil {
				t.Fatal(err)
			}
			err = transactionsUse.Create(models.CreateTransactionRequest{
				Block:     1,
				BlockTime: time.Now().UTC().Unix(),
				Hash:      fmt.Sprintf("%s%d_2", uniqueHash, i),
				From:      sender2,
				To:        receiver2,
				Amount:    big.NewInt(100),
			})
			if err != nil {
				t.Fatal(err)
			}
		}()
	}

	wg.Wait()
	result, err := transactionsUse.Select(models.SearchTransactionsRequest{
		PageMetadata: models.PageMetadata{
			OrderDir: "desc",
		},
		Address: utils.Ptr(receiver),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 100 {
		t.Fatal("should return 100 transactions")
	}

}
