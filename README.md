# eth-tx-parser


## Description

Service parses block by block and stores transactions for subscribed addressees.
Service exposes public interface for external usage via http Rest Api.
- Supported transaction types: `native ethereum transfer`, `erc20 transfer`, `erc20 transferFrom`
- Supported erc20 tokens: `USDT`, `USDC` - can be easily extendable in future

## How to run
1. export env params (example in `.example.env`)
2. `go run main.go`


## Project structure

- `./main.go` - entrypoint
- `./internal/delivery/api/` - http server with REST-api endpoints
- `./internal/entities/` - entities layer (models for storing in any storage)
- `./internal/models/` - models layer for communication via api and use cases
- `./internal/repos/` - interfaces for repository layer
- `./internal/repos/memory` - implementation of in-memory storage
- `./internal/runners/indexer.go` - blockchain indexer runner, iteratively parses block by block
- `./internal/usecases/` - use cases layer with main logic
- `./pkg/ethclient` - JsonRpc client for communication with EVM nodes
- `./pkg/utils` - some additional helper functions
- `./tests` - tests for use cases layer

## Env params
- `RPC_URL` - http rpc node url
- `REST_PORT` - port for rest api server

Check `.example.env` file for details


## Api docs

### Subscribe Address

#### Endpoint: POST: `/subcribe/`

#### Body:

```
{
  "address": "0xb5d85CBf7cB3EE0D56b3bB207D5Fc4B82f43F511"
}
```

#### Response:
```
{
    "status": "success",
    "result": {
        "ok": true
    }
}
```

### Get Current block

- returns last parsed block by indexer

#### Endpoint: Get: `/get_current_block/`

#### Response:
```
{
    "status": "success",
    "result": {
        "number": 22345080
    }
}
```

### Get transactions

- returns last parsed block by indexer

#### Endpoint: Get: `/get_transactions/`

#### Filters:
- `address` string
- `size` int (default 10)
- `order_dir` string (asc || desc, default desc)

#### Example:
`http://localhost:8005/get_transactions?address=0xb5d85CBf7cB3EE0D56b3bB207D5Fc4B82f43F511&size=3&order_dir=desc`

#### Response:
```
{
    "status": "success",
    "result": [
        {
            "id": 17,
            "block": 22345451,
            "block_time": 1745577035,
            "hash": "0xf6b0639bda5d5362d5e0449d56510bf0888fbf3f6e90470a3b63ec8b827c9dc2",
            "from": "0xb5d85cbf7cb3ee0d56b3bb207d5fc4b82f43f511",
            "to": "0x70bf9d6f5d81c67f293adcc6ab4c13eea1f6ae38",
            "amount": 2769940000000000,
            "fee": 0,
            "type": "eth-transfer"
        },
        {
            "id": 16,
            "block": 22345451,
            "block_time": 1745577035,
            "hash": "0xd713326a518a5d89a9f90d72d5d29dd745147d0a17b012098403f4b9f6b88896",
            "from": "0xb5d85cbf7cb3ee0d56b3bb207d5fc4b82f43f511",
            "to": "0xbb9b5f83cc62523d77dbddaa48fbd426b03227a7",
            "amount": 12302950000000000,
            "fee": 0,
            "type": "eth-transfer"
        },
        {
            "id": 15,
            "block": 22345451,
            "block_time": 1745577035,
            "hash": "0x2ae1d451067decbed53a476022957364887951567dd312478e61c111453ea815",
            "from": "0xb5d85cbf7cb3ee0d56b3bb207d5fc4b82f43f511",
            "to": "0x225585994063b5a5d47aa35fc28c12bcc19d1e15",
            "amount": 9937440000000000,
            "fee": 0,
            "type": "eth-transfer"
        }
    ]
}
```
