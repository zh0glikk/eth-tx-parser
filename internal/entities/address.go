package entities

import "github.com/zh0glikk/eth-tx-parser/pkg/interfaces"

type Address struct {
	ID      uint64
	Address string
}

func (a Address) Less(b interfaces.Sortable) bool {
	return a.ID < b.(Address).ID
}
