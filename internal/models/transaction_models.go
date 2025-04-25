package models

import "math/big"

type CreateTransactionRequest struct {
	Block     uint64
	BlockTime int64
	Hash      string
	From      string
	To        string
	Amount    *big.Int
	Type      TransactionType
	Token     string
}

type SearchTransactionsRequest struct {
	PageMetadata
	From *string
	To   *string

	Address *string
}

type TransactionResponse struct {
	ID        uint64          `json:"id"`
	Block     uint64          `json:"block"`
	BlockTime int64           `json:"block_time"`
	Hash      string          `json:"hash"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Amount    *big.Int        `json:"amount"`
	Type      TransactionType `json:"type"`
	Token     string          `json:"token,omitempty"`
}

type TransactionType string

const (
	EthTransfer   TransactionType = "eth-transfer"
	Erc20Transfer TransactionType = "erc20-transfer"
)
