package mocks

import (
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/pkg/ethclient"
	"math/big"
)

type ethClientMock struct{}

func (e *ethClientMock) ChainID() (*big.Int, error) {
	return big.NewInt(1), nil
}

func (e *ethClientMock) BlockNumber() (uint64, error) {
	return 100, nil
}

func (e *ethClientMock) GetBlockByNumber(number uint64) (*ethclient.Block, error) {
	return &ethclient.Block{
		Number:    "0x64",
		Timestamp: "0x64",
		Transactions: []ethclient.Transaction{
			{
				From:  "0x0000000000000000000000000000000000000001",
				To:    "0x0000000000000000000000000000000000000002",
				Input: "0x0",
				Value: "0xF",
			},
			{
				From:  "0x0000000000000000000000000000000000000003",
				To:    models.UsdtAddress,
				Input: "0xa9059cbb000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000F4240",
				Value: "0x0",
			},
			{
				From:  "0x0000000000000000000000000000000000000099",
				To:    models.UsdcAddress,
				Input: "0x23b872dd0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000F4240",
				Value: "0x0",
			},
		},
	}, nil
}

func NewEthClientMock() ethclient.Client {
	return &ethClientMock{}
}
