package handler

import (
	"net/http"

	"github.com/go-openapi/swag"

	"github.com/go-openapi/runtime/middleware"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/models"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_721_swap_pairs"
	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/env"
)

type GetERC721SwapPairsHandler struct {
	*env.Env
	H func(e *env.Env, params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder
}

func (h GetERC721SwapPairsHandler) Serve(params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder {
	responder := h.H(h.Env, params)
	return responder
}

// NewGetERC721SwapPairsHandler creates a new GetERC721SwapPairsHandler instance.
func NewGetERC721SwapPairsHandler(e *env.Env, api *operations.BscEvmCompatibleBridgeAPIAPI) GetERC721SwapPairsHandler {
	return GetERC721SwapPairsHandler{
		Env: e,
		H: func(e *env.Env, params erc_721_swap_pairs.GetErc721SwapPairsParams) middleware.Responder {
			total, pairList, err := e.ERC721SwapPairDao.GetSwapPairs(params)
			if err != nil {
				return erc_721_swap_pairs.NewGetErc721SwapPairsBadRequest().WithPayload(&models.Error{Code: http.StatusBadRequest, Message: swag.String(err.Error())})
			}
			res := models.Erc721SwapPairs{Total: total, Pairs: make([]*models.Erc721SwapPair, 0, len(pairList))}
			for _, t := range pairList {
				res.Pairs = append(res.Pairs, &models.Erc721SwapPair{
					Available:      t.Available,
					BaseURI:        t.BaseURI,
					CreateTxHash:   t.CreateTxHash,
					CreatedAt:      t.CreatedAt.String(),
					DstChainID:     t.DstChainID,
					DstTokenAddr:   t.DstTokenAddr,
					DstTokenName:   t.DstTokenName,
					RegisterTxHash: t.RegisterTxHash,
					Sponsor:        t.Sponsor,
					SrcChainID:     t.SrcChainID,
					SrcTokenAddr:   t.SrcTokenAddr,
					SrcTokenName:   t.SrcTokenName,
					State:          string(t.State),
					Symbol:         t.Symbol,
					UpdatedAt:      t.UpdatedAt.String(),
				})
			}

			return erc_721_swap_pairs.NewGetErc721SwapPairsOK().WithPayload(&res)
		},
	}
}
