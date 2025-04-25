package rest

import (
	"github.com/zh0glikk/eth-tx-parser/internal/usecases"
	"net/http"
)

type Router struct {
	addressUse     usecases.AddressesUseCase
	blocksUse      usecases.BlocksUseCase
	transactionUse usecases.TransactionsUseCase
}

func NewRouter(
	addressUse usecases.AddressesUseCase,
	blocksUse usecases.BlocksUseCase,
	transactionUse usecases.TransactionsUseCase,
) *Router {
	return &Router{
		addressUse:     addressUse,
		blocksUse:      blocksUse,
		transactionUse: transactionUse,
	}
}

func (rs *Router) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/subscribe", rs.subscribeAddress)
	mux.HandleFunc("/get_current_block", rs.getCurrentBlockNumber)
	mux.HandleFunc("/get_transactions", rs.getTransactions)

	return mux
}
