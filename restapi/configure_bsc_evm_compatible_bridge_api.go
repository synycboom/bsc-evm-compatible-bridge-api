// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/svc_info"
)

//go:generate swagger generate server --target ../../bsc-evm-compatible-bridge-api --name BscEvmCompatibleBridgeAPI --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.BscEvmCompatibleBridgeAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BscEvmCompatibleBridgeAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.Erc721SwapPairsGetErc721SwapPairsHandler == nil {
		api.Erc721SwapPairsGetErc721SwapPairsHandler = erc_721_swap_pairs.GetErc721SwapPairsHandlerFunc(func(params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_721_swap_pairs.GetErc721SwapPairs has not yet been implemented")
		})
	}
	if api.Erc721SwapsGetErc721SwapsHandler == nil {
		api.Erc721SwapsGetErc721SwapsHandler = erc_721_swaps.GetErc721SwapsHandlerFunc(func(params erc_721_swaps.GetErc721SwapsParams) middleware.Responder {
			return middleware.NotImplemented("operation erc_721_swaps.GetErc721Swaps has not yet been implemented")
		})
	}
	if api.SvcInfoGetInfoHandler == nil {
		api.SvcInfoGetInfoHandler = svc_info.GetInfoHandlerFunc(func(params svc_info.GetInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation svc_info.GetInfo has not yet been implemented")
		})
	}

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
