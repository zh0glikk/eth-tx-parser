package usecases

import (
	"github.com/zh0glikk/eth-tx-parser/internal/entities"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/internal/repos"
	"github.com/zh0glikk/eth-tx-parser/pkg/utils"
	"strings"
)

type TransactionsUseCase interface {
	Create(data models.CreateTransactionRequest) error

	Select(req models.SearchTransactionsRequest) ([]models.TransactionResponse, error)
}

type transactionsUse struct {
	repo repos.TransactionsRepo
}

func NewTransactionsUseCase(repo repos.TransactionsRepo) TransactionsUseCase {
	return &transactionsUse{
		repo: repo,
	}
}

func (u *transactionsUse) Create(data models.CreateTransactionRequest) error {
	return u.repo.Create(entities.Transaction{
		Block:     data.Block,
		BlockTime: data.BlockTime,
		Hash:      data.Hash,
		From:      strings.ToLower(data.From),
		To:        strings.ToLower(data.To),
		Amount:    data.Amount,
		Type:      string(data.Type),
		Token:     data.Token,
	})
}

func (u *transactionsUse) Select(req models.SearchTransactionsRequest) ([]models.TransactionResponse, error) {
	if req.From != nil {
		req.From = utils.Ptr(strings.ToLower(*req.From))
	}
	if req.To != nil {
		req.To = utils.Ptr(strings.ToLower(*req.To))
	}
	if req.Address != nil {
		req.Address = utils.Ptr(strings.ToLower(*req.Address))
	}

	transactions, err := u.repo.Select(req)
	if err != nil {
		return nil, err
	}

	res := make([]models.TransactionResponse, 0, len(transactions))
	for _, tx := range transactions {
		res = append(res, models.TransactionResponse{
			ID:        tx.ID,
			Block:     tx.Block,
			BlockTime: tx.BlockTime,
			Hash:      tx.Hash,
			From:      tx.From,
			To:        tx.To,
			Amount:    tx.Amount,
			Type:      models.TransactionType(tx.Type),
			Token:     tx.Token,
		})
	}

	return res, nil
}
