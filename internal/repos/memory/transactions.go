package memory

import (
	"errors"
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
	"sync"
)

type transactionRepo struct {
	mu   sync.RWMutex
	data []entities.Transaction

	//for unique tx hashes
	unique map[string]struct{}

	//for faster filtering
	incomingTransactions    map[string][]entities.Transaction
	outgoingTransactions    map[string][]entities.Transaction
	related2eoaTransactions map[string][]entities.Transaction
}

func NewTransactionRepo() repos.TransactionsRepo {
	return &transactionRepo{
		mu:                      sync.RWMutex{},
		data:                    nil,
		unique:                  make(map[string]struct{}),
		incomingTransactions:    make(map[string][]entities.Transaction),
		outgoingTransactions:    make(map[string][]entities.Transaction),
		related2eoaTransactions: make(map[string][]entities.Transaction),
	}
}

func (t *transactionRepo) Create(transaction entities.Transaction) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.unique[transaction.Hash]; ok {
		return errors.New(fmt.Sprintf("row with such hash already exists: %s", transaction.Hash))
	}

	transaction.ID = uint64(len(t.data) + 1)
	t.data = append(t.data, transaction)

	t.unique[transaction.Hash] = struct{}{}
	t.incomingTransactions[transaction.To] = append(t.incomingTransactions[transaction.To], transaction)
	t.outgoingTransactions[transaction.From] = append(t.outgoingTransactions[transaction.From], transaction)

	t.related2eoaTransactions[transaction.From] = append(t.related2eoaTransactions[transaction.From], transaction)
	t.related2eoaTransactions[transaction.To] = append(t.related2eoaTransactions[transaction.To], transaction)

	return nil
}

func (t *transactionRepo) Select(req models.SearchTransactionsRequest) ([]entities.Transaction, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if req.From != nil {
		return getPage(req.PageMetadata, t.outgoingTransactions[*req.From]), nil
	}
	if req.To != nil {
		return getPage(req.PageMetadata, t.incomingTransactions[*req.To]), nil
	}
	if req.Address != nil {
		return getPage(req.PageMetadata, t.related2eoaTransactions[*req.Address]), nil
	}

	return getPage(req.PageMetadata, t.data), nil
}
