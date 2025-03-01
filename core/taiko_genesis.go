package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	taikoGenesis "github.com/ethereum/go-ethereum/core/taiko_genesis"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// TaikoGenesisBlock returns the Taiko network genesis block configs.
func TaikoGenesisBlock(networkID uint64) *Genesis {
	chainConfig := params.TaikoChainConfig

	var allocJSON []byte
	switch networkID {
	case params.TaikoInternalL2ANetworkID.Uint64():
		chainConfig.ChainID = params.TaikoInternalL2ANetworkID
		allocJSON = taikoGenesis.InternalL2AGenesisAllocJSON
	case params.TaikoInternalL2BNetworkID.Uint64():
		chainConfig.ChainID = params.TaikoInternalL2BNetworkID
		allocJSON = taikoGenesis.InternalL2BGenesisAllocJSON
	case params.SnaefellsjokullNetworkID.Uint64():
		chainConfig.ChainID = params.SnaefellsjokullNetworkID
		allocJSON = taikoGenesis.SnaefellsjokullGenesisAllocJSON
	case params.AskjaNetworkID.Uint64():
		chainConfig.ChainID = params.AskjaNetworkID
		allocJSON = taikoGenesis.AskjaGenesisAllocJSON
	case params.GrimsvotnNetworkID.Uint64():
		chainConfig.ChainID = params.GrimsvotnNetworkID
		allocJSON = taikoGenesis.GrimsvotnGenesisAllocJSON
	case params.EldfellNetworkID.Uint64():
		chainConfig.ChainID = params.EldfellNetworkID
		allocJSON = taikoGenesis.EldfellGenesisAllocJSON
	case params.JolnirNetworkID.Uint64():
		chainConfig.ChainID = params.JolnirNetworkID
		allocJSON = taikoGenesis.JolnirGenesisAllocJSON
	case params.KatlaNetworkID.Uint64():
		chainConfig.ChainID = params.KatlaNetworkID
		allocJSON = taikoGenesis.KatlaGenesisAllocJSON
	case params.HeklaNetworkID.Uint64():
		chainConfig.ChainID = params.HeklaNetworkID
		allocJSON = taikoGenesis.HeklaGenesisAllocJSON
	default:
		chainConfig.ChainID = params.TaikoInternalL2ANetworkID
		allocJSON = taikoGenesis.InternalL2AGenesisAllocJSON
	}

	var alloc GenesisAlloc
	if err := alloc.UnmarshalJSON(allocJSON); err != nil {
		log.Crit("unmarshal alloc json error", "error", err)
	}

	return &Genesis{
		Config:     chainConfig,
		ExtraData:  []byte{},
		GasLimit:   uint64(15_000_000),
		Difficulty: common.Big0,
		Alloc:      alloc,
		GasUsed:    0,
		BaseFee:    new(big.Int).SetUint64(10_000_000),
	}
}
