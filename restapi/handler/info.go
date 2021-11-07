package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/svc_info"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
)

type GetInfosHandler struct {
	*env.Env
	H func(e *env.Env, params svc_info.GetInfoParams) middleware.Responder
}

func (h GetInfosHandler) Serve(params svc_info.GetInfoParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

func NewGetInfoHandler(e *env.Env, _ *operations.BscEvmCompatibleBridgeAPIAPI) GetInfosHandler {
	return GetInfosHandler{
		Env: e,
		H: func(e *env.Env, _ svc_info.GetInfoParams) middleware.Responder {
			info := models.ServiceInfo{
				BscSwapAgent: e.Config.SwapConfig.BSCSwapAgent,
				EthSwapAgent: e.Config.SwapConfig.EthSwapAgent,
			}
			return svc_info.NewGetInfoOK().WithPayload(&info)
		},
	}
}
