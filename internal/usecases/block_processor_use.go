package usecases

import (
	"errors"
	"github.com/zh0glikk/eth-tx-parser/internal/models"
	"github.com/zh0glikk/eth-tx-parser/pkg/ethclient"
	"math/big"
	"strings"
)

const (
	transferSig     = "0xa9059cbb"
	transferFromSig = "0x23b872dd"
)

type BlockParserUseCase interface {
	GetBlock2Process() (uint64, error)
	ParseBlockTransactions(number uint64) ([]models.CreateTransactionRequest, error)
}

type blockParserUse struct {
	cli ethclient.Client

	addressesUse AddressesUseCase
	blocksUse    BlocksUseCase
}

func NewBlockParserUseCase(
	cli ethclient.Client,
	addressesUse AddressesUseCase,
	blocksUse BlocksUseCase,
) BlockParserUseCase {
	return &blockParserUse{
		cli:          cli,
		addressesUse: addressesUse,
		blocksUse:    blocksUse,
	}
}

func (b *blockParserUse) GetBlock2Process() (uint64, error) {
	block, err := b.blocksUse.GetLastBlock()
	if err != nil {
		return 0, err
	}
	if block == nil {
		number, err := b.cli.BlockNumber()
		if err != nil {
			return 0, err
		}

		return number, nil
	}
	return block.Number + 1, nil
}

func (b *blockParserUse) ParseBlockTransactions(number uint64) ([]models.CreateTransactionRequest, error) {
	block, err := b.cli.GetBlockByNumber(number)
	if err != nil {
		return nil, err
	}

	var transactions []models.CreateTransactionRequest

	for _, tx := range block.Transactions {
		if strings.ToLower(tx.To) == strings.ToLower(models.UsdtAddress) ||
			strings.ToLower(tx.To) == strings.ToLower(models.UsdcAddress) {

			if b.isErc20Transfer(tx.Input) {
				to, amount, err := b.parseErc20Transfer(tx.Input)
				if err != nil {
					return nil, err
				}

				ok, err := b.ensureSubscribed(tx.From, to)
				if err != nil {
					return nil, err
				}
				if !ok {
					continue
				}

				transactions = append(transactions, models.CreateTransactionRequest{
					Block:     number,
					BlockTime: block.BlockTimestamp(),
					Hash:      tx.Hash,
					From:      tx.From,
					To:        to,
					Amount:    amount,
					Type:      models.Erc20Transfer,
					Token:     tx.To,
				})
				continue
			}
			if b.isErc20TransferFrom(tx.Input) {
				from, to, amount, err := b.parseErc20TransferFrom(tx.Input)
				if err != nil {
					return nil, err
				}

				ok, err := b.ensureSubscribed(from, to)
				if err != nil {
					return nil, err
				}
				if !ok {
					continue
				}

				transactions = append(transactions, models.CreateTransactionRequest{
					Block:     number,
					BlockTime: block.BlockTimestamp(),
					Hash:      tx.Hash,
					From:      from,
					To:        to,
					Amount:    amount,
					Type:      models.Erc20Transfer,
					Token:     tx.To,
				})
				continue
			}

			continue
		}

		//process eth transfer
		ok, err := b.ensureSubscribed(tx.From, tx.To)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		transactions = append(transactions, models.CreateTransactionRequest{
			Block:     number,
			BlockTime: block.BlockTimestamp(),
			Hash:      tx.Hash,
			From:      tx.From,
			To:        tx.To,
			Amount:    tx.GetValue(),
			Type:      models.EthTransfer,
		})
	}

	return transactions, nil
}

func (b *blockParserUse) ensureSubscribed(from, to string) (bool, error) {
	ok, err := b.addressesUse.IsSubscribed(from)
	if err != nil {
		return false, err
	}
	if !ok {
		ok, err = b.addressesUse.IsSubscribed(to)
		if err != nil {
			return false, err
		}
	}
	if !ok {
		return false, nil
	}

	return true, nil
}

func (b *blockParserUse) isErc20Transfer(callData string) bool {
	if len(callData) == 138 && callData[:10] == transferSig {
		return true
	}

	return false
}

func (b *blockParserUse) isErc20TransferFrom(callData string) bool {
	if len(callData) == 202 && callData[:10] == transferFromSig {
		return true
	}

	return false
}

func (b *blockParserUse) parseErc20Transfer(callData string) (string, *big.Int, error) {
	to := "0x" + callData[34:74]
	amount, ok := big.NewInt(0).SetString(callData[74:], 16)
	if !ok {
		return "", nil, errors.New("failed to parse amount")
	}

	return to, amount, nil
}

func (b *blockParserUse) parseErc20TransferFrom(callData string) (string, string, *big.Int, error) {
	from := "0x" + callData[34:74]
	to := "0x" + callData[98:138]
	amount, ok := big.NewInt(0).SetString(callData[138:], 16)
	if !ok {
		return "", "", nil, errors.New("failed to parse amount")
	}

	return from, to, amount, nil
}
