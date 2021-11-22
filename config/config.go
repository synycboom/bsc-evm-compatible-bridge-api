package config

import (
	"encoding/json"
	"io/ioutil"

	svc "github.com/synycboom/bsc-evm-compatible-bridge-api/services"
	ulog "github.com/synycboom/bsc-evm-compatible-bridge-api/utils/log"
)

// Config represents configuration from config.json.
type Config struct {
	Logs       ulog.LogsConfig  `json:"logs"`
	CacheTTLs  map[string]int64 `json:"cache_ttls"`
	DB         DBConfig         `json:"db"`
	SwapConfig SwapConfig       `json:"swap_config"`
	CorsConfig CorsConfig       `json:"cors"`
}

type Consumer struct {
	Username string
	Key      string
}

type SwapConfig struct {
	ETHChainID              int64  `json:"eth_chain_id"`
	BSCChainID              int64  `json:"bsc_chain_id"`
	PolygonChainID          int64  `json:"polygon_chain_id"`
	FantomChainID           int64  `json:"fantom_chain_id"`
	EthErc721SwapAgent      string `json:"eth_erc_721_swap_agent"`
	BSCErc721SwapAgent      string `json:"bsc_erc_721_swap_agent"`
	PolygonErc721SwapAgent  string `json:"polygon_erc_721_swap_agent"`
	FantomErc721SwapAgent   string `json:"fantom_erc_721_swap_agent"`
	EthErc1155SwapAgent     string `json:"eth_erc_1155_swap_agent"`
	BSCErc1155SwapAgent     string `json:"bsc_erc_1155_swap_agent"`
	PolygonErc1155SwapAgent string `json:"polygon_erc_1155_swap_agent"`
	FantomErc1155SwapAgent  string `json:"fantom_erc_1155_swap_agent"`
}

type DBConfig struct {
	DSN      string `json:"dsn"`
	LogLevel string `json:"log_level"`
}

type CorsConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
}

// InitConfigFromFile initializes a new Env from configuration file.
func InitConfigFromFile(configFileName string) *Config {
	bz, err := ioutil.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	var configOpts Config
	if err := json.Unmarshal(bz, &configOpts); err != nil {
		panic(err)
	}

	return &configOpts
}

func InitConfigFromSecret(secretName, secretRegion string) *Config {
	bzStr, err := svc.GetSecret(secretName, secretRegion)
	if err != nil {
		panic(err)
	}

	var configOpts Config
	if err := json.Unmarshal([]byte(bzStr), &configOpts); err != nil {
		panic(err)
	}

	return &configOpts
}
