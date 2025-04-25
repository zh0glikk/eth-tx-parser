package entities

import (
	"github.com/zh0glikk/eth-tx-parser/pkg/interfaces"
	"math/big"
)

type Transaction struct {
	ID        uint64
	Block     uint64
	BlockTime int64
	Hash      string
	From      string
	To        string
	Amount    *big.Int
	Type      string
	Token     string
}

func (a Transaction) Less(b interfaces.Sortable) bool {
	return a.ID < b.(Transaction).ID
}
