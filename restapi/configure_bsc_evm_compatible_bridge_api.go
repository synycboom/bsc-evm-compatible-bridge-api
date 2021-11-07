// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/dao"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/middlewares"
	handler "github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/handler"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/svc_info"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/services"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/cache"

	cfg "github.com/synycboom/bsc-evm-compatible-bridge-api/config"
	uenv "github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/log"
)

//go:generate swagger generate server --target ../../bsc-evm-compatible-bridge-api --name BscEvmCompatibleBridgeAPI --spec ../swagger.yml --principal interface{}

const cacheServiceTickRate = time.Hour

var (
	config       *cfg.Config
	cacheService services.Service

	erc721SwapPairCache,
	infoCache,
	erc721SwapCache *middlewares.MWCacher

	env *uenv.Env
)

var cliOpts = struct {
	ConfigFileName string `short:"c" long:"config-file" description:"Config filename"`
	SecretName     string `short:"s" long:"secret-name" description:"the secret name of the config"`
	SecretRegion   string `short:"r" long:"secret-region" description:"the secret region of the config"`
}{}

func configureFlags(api *operations.BscEvmCompatibleBridgeAPIAPI) {
	param1 := swag.CommandLineOptionsGroup{
		ShortDescription: "config",
		Options:          &cliOpts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{param1}
}

func configureAPI(api *operations.BscEvmCompatibleBridgeAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	api.Logger = func(str string, args ...interface{}) {
		level, str := log.ParsePrefixedLogString(str)
		log.GetLogger(level)(str, args...)
	}

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.Erc721SwapPairsGetErc721SwapPairsHandler = erc_721_swap_pairs.GetErc721SwapPairsHandlerFunc(func(params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder {
		return erc721SwapCache.Serve(params.HTTPRequest, func() middleware.Responder {
			return handler.NewGetSwapPairsHandler(env, api).Serve(params)
		}, api.JSONProducer)
	})

	api.Erc721SwapsGetErc721SwapsHandler = erc_721_swaps.GetErc721SwapsHandlerFunc(func(params erc_721_swaps.GetErc721SwapsParams) middleware.Responder {
		return erc721SwapPairCache.Serve(params.HTTPRequest, func() middleware.Responder {
			return handler.NewGetSwapsHandler(env, api).Serve(params)
		}, api.JSONProducer)
	})

	api.SvcInfoGetInfoHandler = svc_info.GetInfoHandlerFunc(func(params svc_info.GetInfoParams) middleware.Responder {
		return infoCache.Serve(params.HTTPRequest, func() middleware.Responder {
			return handler.NewGetInfoHandler(env, api).Serve(params)
		}, api.JSONProducer)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	configFileName := cliOpts.ConfigFileName
	secretName := cliOpts.SecretName
	secretRegion := cliOpts.SecretRegion
	if configFileName == "" && (secretName == "" || secretRegion == "") {
		panic("missing config file path and secret path")
	}
	if configFileName != "" {
		config = cfg.InitConfigFromFile(configFileName)
	}
	if secretName != "" && secretRegion != "" {
		config = cfg.InitConfigFromSecret(secretName, secretRegion, config)
	}

	// init logger
	log.InitLogger(config.Logs)

	// init cache
	store := cache.NewMemStorage()
	swapPairCacheMS := config.CacheTTLs["swap_pairs"] * time.Millisecond.Nanoseconds()
	erc721SwapPairCache = middlewares.NewMWCacher(store, time.Duration(swapPairCacheMS))
	swapsCacheMs := config.CacheTTLs["swaps"] * time.Millisecond.Nanoseconds()
	erc721SwapCache = middlewares.NewMWCacher(store, time.Duration(swapsCacheMs))
	infoCacheMs := config.CacheTTLs["info"] * time.Millisecond.Nanoseconds()
	infoCache = middlewares.NewMWCacher(store, time.Duration(infoCacheMs))

	cacheService = services.NewCacheService(store, cacheServiceTickRate)
	if err := cacheService.Start(); err != nil {
		panic(err) // fatal
	}
	// init db
	dbConfig := config.DB
	swapPairDao, swapDao, err := dao.NewDaoServices(dbConfig.DSN, dbConfig.LogLevel)
	if err != nil {
		panic(err)
	}

	// init env
	env = &uenv.Env{
		Config:      config,
		SwapPairDao: swapPairDao,
		SwapDao:     swapDao,
		Cache:       store,
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
