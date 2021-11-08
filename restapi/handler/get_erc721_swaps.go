package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
)

type GetERC721SwapsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_721_swaps.GetErc721SwapsParams) middleware.Responder
}

func (h GetERC721SwapsHandler) Serve(params erc_721_swaps.GetErc721SwapsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetswapsHandler creates a new GetERC721SwapsHandler instance.
func NewGetERC721SwapsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC721SwapsHandler {
	return GetERC721SwapsHandler{
		Env: e,
		H: func(e *env.Env, params erc_721_swaps.GetErc721SwapsParams) middleware.Responder {
			total, SwapsList, err := e.ERC721SwapDao.GetSwaps(params)
			if err != nil {
				return erc_721_swaps.NewGetErc721SwapsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc721Swaps{Total: total, Erc721Swaps: make([]*models.Erc721Swap, 0, len(SwapsList))}
			for _, s := range SwapsList {
				res.Erc721Swaps = append(res.Erc721Swaps, &models.Erc721Swap{
					BaseURI:       s.BaseURI,
					CreatedAt:     s.CreatedAt.String(),
					DstChainID:    s.DstChainID,
					DstTokenAddr:  s.DstTokenAddr,
					DstTokenName:  s.DstTokenName,
					FillTxHash:    s.FillTxHash,
					Recipient:     s.Recipient,
					RequestTxHash: s.RequestTxHash,
					Sender:        s.Sender,
					SrcChainID:    s.SrcChainID,
					SrcTokenAddr:  s.SrcTokenAddr,
					SrcTokenName:  s.SrcTokenName,
					State:         string(s.State),
					SwapDirection: string(s.SwapDirection),
					TokenID:       s.TokenID,
					TokenURI:      s.TokenURI,
					UpdatedAt:     s.UpdatedAt.String(),
				})
			}

			return erc_721_swaps.NewGetErc721SwapsOK().WithPayload(&res)
		},
	}
}
