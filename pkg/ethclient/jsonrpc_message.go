package ethclient

import "encoding/json"

type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func newJSONRPCMessage(id json.RawMessage, method string, params ...interface{}) (*jsonrpcMessage, error) {
	msg := &jsonrpcMessage{
		Version: "2.0",
		ID:      id,
		Method:  method,
	}
	if params != nil { // prevent sending "params":null
		var err error
		if msg.Params, err = json.Marshal(params); err != nil {
			return nil, err
		}
	}
	return msg, nil
}
