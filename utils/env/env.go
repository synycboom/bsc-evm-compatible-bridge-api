package env

import (
	"github.com/synycboom/bsc-evm-compatible-bridge-api/config"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/dao"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/cache"
)

type Env struct {
	Config *config.Config

	SwapPairDao dao.SwapPairDaoInterface
	SwapDao     dao.SwapDaoInterface

	Cache cache.Store
}
