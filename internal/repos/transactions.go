package repos

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
)

type TransactionsRepo interface {
	Create(transaction entities.Transaction) error

	Select(req models.SearchTransactionsRequest) ([]entities.Transaction, error)
}
