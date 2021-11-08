package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swaps"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
)

type GetERC1155SwapsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_1155_swaps.GetErc1155SwapsParams) middleware.Responder
}

func (h GetERC1155SwapsHandler) Serve(params erc_1155_swaps.GetErc1155SwapsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetERC1155SwapsHandler creates a new GetERC1155SwapsHandler instance.
func NewGetERC1155SwapsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC1155SwapsHandler {
	return GetERC1155SwapsHandler{
		Env: e,
		H: func(e *env.Env, params erc_1155_swaps.GetErc1155SwapsParams) middleware.Responder {
			total, SwapsList, err := e.ERC1155SwapDao.GetSwaps(params)
			if err != nil {
				return erc_1155_swaps.NewGetErc1155SwapsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc1155Swaps{Total: total, Erc1155Swaps: make([]*models.Erc1155Swap, 0, len(SwapsList))}
			for _, s := range SwapsList {
				var ids []string
				if err := json.Unmarshal(s.IDs, &ids); err != nil {
					return erc_1155_swaps.NewGetErc1155SwapsBadRequest().WithPayload(&models.Error{Code: http.StatusInternalServerError, Message: swag.String(err.Error())})
				}
				var amounts []string
				if err := json.Unmarshal(s.Amounts, &amounts); err != nil {
					return erc_1155_swaps.NewGetErc1155SwapsBadRequest().WithPayload(&models.Error{Code: http.StatusInternalServerError, Message: swag.String(err.Error())})
				}

				res.Erc1155Swaps = append(res.Erc1155Swaps, &models.Erc1155Swap{
					CreatedAt:     s.CreatedAt.String(),
					DstChainID:    s.DstChainID,
					DstTokenAddr:  s.DstTokenAddr,
					FillTxHash:    s.FillTxHash,
					Recipient:     s.Recipient,
					RequestTxHash: s.RequestTxHash,
					Sender:        s.Sender,
					SrcChainID:    s.SrcChainID,
					SrcTokenAddr:  s.SrcTokenAddr,
					State:         string(s.State),
					SwapDirection: string(s.SwapDirection),
					Ids:           ids,
					Amounts:       amounts,
					UpdatedAt:     s.UpdatedAt.String(),
				})
			}

			return erc_1155_swaps.NewGetErc1155SwapsOK().WithPayload(&res)
		},
	}
}
