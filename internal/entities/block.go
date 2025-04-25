package entities

import "github.com/zh0glikk/eth-tx-parser/pkg/interfaces"

type Block struct {
	ID     uint64
	Number uint64
}

func (a Block) Less(b interfaces.Sortable) bool {
	return a.ID < b.(Block).ID
}
