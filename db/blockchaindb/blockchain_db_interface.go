package blockchaindb

import (
	"it-chain/service/blockchain"
)

type BlockChainDB interface {
	AddBlock(block *blockchain.Block) error
	GetBlockByNumber(blockNumber uint64) (*blockchain.Block, error)
	GetBlockByHash(hash string) (*blockchain.Block, error)
}