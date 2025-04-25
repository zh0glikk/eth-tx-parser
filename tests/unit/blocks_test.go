package unit

import (
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	memory2 "github.com/zh0glikk/eth-tx-parser/internal/repos/memory"
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"testing"
)

var (
	blocksUse usecases.BlocksUseCase
)

func initBlocksTest() {
	blocksUse = usecases.NewBlocksUseCase(memory2.NewBlockRepo())
}

func TestBlocks_Create(t *testing.T) {
	initBlocksTest()

	err := blocksUse.Create(models.CreateBlockRequest{
		Number: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBlocks_Create_GetLastBlock(t *testing.T) {
	initBlocksTest()

	err := blocksUse.Create(models.CreateBlockRequest{
		Number: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	block, err := blocksUse.GetLastBlock()
	if err != nil {
		t.Fatal(err)
	}
	if block == nil {
		t.Fatal("blocks length should be 1")
	}
	if block.Number != 1 {
		t.Fatal("block number should be 1")
	}

}

func TestBlocks_Create_Multiple_Select_Last(t *testing.T) {
	initBlocksTest()

	for i := 0; i < 100; i++ {
		err := blocksUse.Create(models.CreateBlockRequest{
			Number: uint64(i),
		})
		if err != nil {
			t.Fatal(err)
		}

	}

	block, err := blocksUse.GetLastBlock()
	if err != nil {
		t.Fatal(err)
	}
	if block == nil {
		t.Fatal("block shouldn't be nil")
	}
	if block.Number != 99 {
		t.Fatal("result number should be 99")
	}
}
