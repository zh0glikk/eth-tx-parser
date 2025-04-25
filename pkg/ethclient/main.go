package ethclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

type Client interface {
	ChainID() (*big.Int, error)
	BlockNumber() (uint64, error)

	GetBlockByNumber(number uint64) (*Block, error)
}

type client struct {
	idCounter atomic.Uint32

	rpcUrl string
	cli    *http.Client
}

func NewClient(rpcUrl string) Client {
	return &client{
		idCounter: atomic.Uint32{},
		rpcUrl:    rpcUrl,
		cli: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (c *client) GetBlockByNumber(number uint64) (*Block, error) {
	result, err := c.call("eth_getBlockByNumber", number, true)
	if err != nil {
		return nil, err
	}

	var block Block
	err = json.Unmarshal(result, &block)
	if err != nil {
		return nil, err
	}

	if block.Number == "" {
		return nil, errors.New("not found")
	}

	return &block, nil
}

func (c *client) BlockNumber() (uint64, error) {
	result, err := c.call("eth_blockNumber")
	if err != nil {
		return 0, err
	}
	var v string
	err = json.Unmarshal(result, &v)
	if err != nil {
		return 0, err
	}
	value, err := strconv.ParseUint(v[2:], 16, 64)
	if err != nil {
		return 0, err
	}

	return value, nil
}

func (c *client) ChainID() (*big.Int, error) {
	result, err := c.call("eth_chainId")
	if err != nil {
		return nil, err
	}
	var v string
	err = json.Unmarshal(result, &v)
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseUint(v[2:], 16, 64)
	if err != nil {
		return nil, err
	}

	return big.NewInt(int64(value)), nil
}

func (c *client) nextID() json.RawMessage {
	id := c.idCounter.Add(1)
	return strconv.AppendUint(nil, uint64(id), 10)
}

func (c *client) call(method string, params ...interface{}) (json.RawMessage, error) {
	msg, err := newJSONRPCMessage(c.nextID(), method, params...)
	if err != nil {
		return nil, err
	}

	bb, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest("POST", c.rpcUrl, bytes.NewBuffer(bb))
	if err != nil {
		return nil, err
	}
	response, err := c.cli.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var jsonRpcResult jsonrpcMessage
	err = json.NewDecoder(response.Body).Decode(&jsonRpcResult)
	if err != nil {
		return nil, err
	}

	return jsonRpcResult.Result, nil
}
