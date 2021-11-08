package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
)

type GetERC1155SwapPairsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_1155_swap_pairs.GetErc1155SwapPairsParams) middleware.Responder
}

func (h GetERC1155SwapPairsHandler) Serve(params erc_1155_swap_pairs.GetErc1155SwapPairsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetERC1155SwapPairsHandler creates a new GetERC1155SwapPairsHandler instance.
func NewGetERC1155SwapPairsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC1155SwapPairsHandler {
	return GetERC1155SwapPairsHandler{
		Env: e,
		H: func(e *env.Env, params erc_1155_swap_pairs.GetErc1155SwapPairsParams) middleware.Responder {
			total, pairList, err := e.ERC1155SwapPairDao.GetSwapPairs(params)
			if err != nil {
				return erc_1155_swap_pairs.NewGetErc1155SwapPairsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc1155SwapPairs{Total: total, Pairs: make([]*models.Erc1155SwapPair, 0, len(pairList))}
			for _, t := range pairList {
				res.Pairs = append(res.Pairs, &models.Erc1155SwapPair{
					Available:      t.Available,
					URI:            t.URI,
					CreateTxHash:   t.CreateTxHash,
					CreatedAt:      t.CreatedAt.String(),
					DstChainID:     t.DstChainID,
					DstTokenAddr:   t.DstTokenAddr,
					RegisterTxHash: t.RegisterTxHash,
					Sponsor:        t.Sponsor,
					SrcChainID:     t.SrcChainID,
					SrcTokenAddr:   t.SrcTokenAddr,
					State:          string(t.State),
					UpdatedAt:      t.UpdatedAt.String(),
				})
			}

			return erc_1155_swap_pairs.NewGetErc1155SwapPairsOK().WithPayload(&res)
		},
	}
}
