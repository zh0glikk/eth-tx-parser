package ethclient

import (
	"math/big"
	"strconv"
)

type Block struct {
	Number       string        `json:"number"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
}

func (b Block) BlockNumber() uint64 {
	value, err := strconv.ParseUint(b.Number[2:], 16, 64)
	if err != nil {
		return 0
	}
	return value
}

func (b Block) BlockTimestamp() int64 {
	value, err := strconv.ParseUint(b.Timestamp[2:], 16, 64)
	if err != nil {
		return 0
	}
	return int64(value)
}

type Transaction struct {
	BlockHash            string        `json:"blockHash"`
	BlockNumber          string        `json:"blockNumber"`
	From                 string        `json:"from"`
	Gas                  string        `json:"gas"`
	GasPrice             string        `json:"gasPrice"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	TransactionIndex     string        `json:"transactionIndex"`
	Value                string        `json:"value"`
	Type                 string        `json:"type"`
	AccessList           []interface{} `json:"accessList"`
	ChainId              string        `json:"chainId"`
	V                    string        `json:"v"`
	YParity              string        `json:"yParity"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
}

func (tx Transaction) GetValue() *big.Int {
	v, _ := new(big.Int).SetString(tx.Value[2:], 16)
	return v
}
