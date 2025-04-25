package rest

import (
	"encoding/json"
	"fmt"
	"github.com/zh0glikk/eth-tx-parser/pkg/utils"
	"net/http"
)

type Response struct {
	Status   string      `json:"status"`
	ErrorMsg *string     `json:"error_msg,omitempty"`
	Result   interface{} `json:"result,omitempty"`
}

func RenderSuccess(w http.ResponseWriter, result interface{}) {
	bb, _ := json.Marshal(Response{
		Status: "success",
		Result: result,
	})
	w.Write(bb)
	w.WriteHeader(http.StatusOK)
}

func RenderBadRequest(w http.ResponseWriter, err error) {
	bb, _ := json.Marshal(Response{
		Status:   "error",
		ErrorMsg: utils.Ptr(err.Error()),
	})
	w.Write(bb)
	w.WriteHeader(http.StatusBadRequest)
}

func RenderInternalServerError(w http.ResponseWriter, err error) {
	bb, _ := json.Marshal(Response{
		Status:   "error",
		ErrorMsg: utils.Ptr(err.Error()),
	})
	w.Write(bb)
	w.WriteHeader(http.StatusInternalServerError)
}

func RenderMethodNotAllowed(w http.ResponseWriter, method string) {
	bb, _ := json.Marshal(Response{
		Status:   "error",
		ErrorMsg: utils.Ptr(fmt.Sprintf("%s method not allowed", method)),
	})
	w.Write(bb)
	w.WriteHeader(http.StatusMethodNotAllowed)
}
