package env

import (
	"github.com/synycboom/bsc-evm-compatible-bridge-api/config"
	erc1155dao "github.com/synycboom/bsc-evm-compatible-bridge-api/dao/erc1155"
	erc721dao "github.com/synycboom/bsc-evm-compatible-bridge-api/dao/erc721"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/cache"
)

type Env struct {
	Config *config.Config

	ERC721SwapPairDao  erc721dao.SwapPairDaoInterface
	ERC721SwapDao      erc721dao.SwapDaoInterface
	ERC1155SwapPairDao erc1155dao.SwapPairDaoInterface
	ERC1155SwapDao     erc1155dao.SwapDaoInterface

	Cache cache.Store
}
