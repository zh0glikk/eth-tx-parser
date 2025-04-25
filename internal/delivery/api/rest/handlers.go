package rest

import (
	"encoding/json"
	"errors"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"net/http"
)

func (rs *Router) subscribeAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderMethodNotAllowed(w, r.Method)
		return
	}

	var request models.CreateAddressRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		RenderBadRequest(w, err)
		return
	}

	err = rs.addressUse.Subscribe(request)
	if err != nil {
		RenderInternalServerError(w, err)
		return
	}

	RenderSuccess(w, struct {
		OK bool `json:"ok"`
	}{
		OK: true,
	})
}

func (rs *Router) getCurrentBlockNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderMethodNotAllowed(w, r.Method)
		return
	}

	block, err := rs.blocksUse.GetLastBlock()
	if err != nil {
		RenderInternalServerError(w, err)
		return
	}

	RenderSuccess(w, struct {
		Number uint64 `json:"number"`
	}{
		Number: block.Number,
	})
}

func (rs *Router) getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderMethodNotAllowed(w, r.Method)
		return
	}

	var request models.SearchTransactionsRequest

	address := r.URL.Query().Get("address")
	if address == "" {
		RenderBadRequest(w, errors.New("address is required"))
		return
	}

	orderDir := r.URL.Query().Get("order_dir")
	if orderDir == "" {
		orderDir = "desc"
	}

	size, err := parseInt(r.URL.Query(), "size")
	if err != nil {
		RenderBadRequest(w, err)
		return
	}
	if size == 0 {
		size = 10
	}

	page, err := parseInt(r.URL.Query(), "page")
	if err != nil {
		RenderBadRequest(w, err)
		return
	}

	request.Address = &address
	request.OrderDir = orderDir
	request.Size = size
	request.Page = page

	transactions, err := rs.transactionUse.Select(request)
	if err != nil {
		RenderInternalServerError(w, err)
		return
	}

	RenderSuccess(w, transactions)
}
