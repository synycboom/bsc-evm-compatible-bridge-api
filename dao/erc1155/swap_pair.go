package dao

import (
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swap_pairs"
)

type SwapPairState string

const (
	SwapPairStateRegistrationOngoing    SwapPairState = "registration_ongoing"
	SwapPairStateRegistrationConfirmed  SwapPairState = "registration_confirmed"
	SwapPairStateCreationTxDryRunFailed SwapPairState = "creation_tx_dry_run_failed"
	SwapPairStateCreationTxCreated      SwapPairState = "creation_tx_created"
	SwapPairStateCreationTxSent         SwapPairState = "creation_tx_sent"
	SwapPairStateCreationTxConfirmed    SwapPairState = "creation_tx_confirmed"
	SwapPairStateCreationTxFailed       SwapPairState = "creation_tx_failed"
	SwapPairStateCreationTxMissing      SwapPairState = "creation_tx_missing"
)

type SwapPair struct {
	ID string

	// Basic Token Information
	SrcChainID   string
	DstChainID   string
	SrcTokenAddr string
	DstTokenAddr string
	Sponsor      string
	Available    bool
	Signature    string
	URI          string

	// Pair State
	State SwapPairState

	// Registration Transaction Information
	RegisterTxHash     string
	RegisterHeight     int64
	RegisterBlockHash  string
	RegisterBlockLogID *string

	// Creation Transaction Information
	CreateConsumedFeeAmount string
	CreateGasPrice          string
	CreateGasUsed           int64
	CreateHeight            int64
	CreateTxHash            string
	CreateTrackRetry        int64
	CreateBlockHash         string
	CreateBlockLogID        *string

	MessageLog string

	// Timestamp
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SwapPair) TableName() string {
	return "erc1155_swap_pairs"
}

type SwapPairDaoInterface interface {
	GetSwapPairs(params erc_1155_swap_pairs.GetErc1155SwapPairsParams) (int64, []SwapPair, error)
}

type SwapPairDao struct {
	db *gorm.DB
}

func NewSwapPairDao(db *gorm.DB) SwapPairDaoInterface {
	return &SwapPairDao{db}
}

func (s *SwapPairDao) GetSwapPairs(params erc_1155_swap_pairs.GetErc1155SwapPairsParams) (int64, []SwapPair, error) {
	db := s.db.Model(&SwapPair{}).Where("available = ?", true)

	if params.SrcChainID != nil && *params.SrcChainID != "" {
		db = db.Where("src_chain_id = ?", *params.SrcChainID)
	}
	if params.DstChainID != nil && *params.DstChainID != "" {
		db = db.Where("dst_chain_id = ?", *params.DstChainID)
	}
	if params.SrcTokenAddr != nil && *params.SrcTokenAddr != "" {
		db = db.Where("src_token_addr = ?", strings.ToLower(*params.SrcTokenAddr))
	}
	if params.DstTokenAddr != nil && *params.DstTokenAddr != "" {
		db = db.Where("dst_token_addr = ?", strings.ToLower(*params.DstTokenAddr))
	}
	if params.State != nil && *params.State != "" {
		db = db.Where("state = ?", *params.State)
	}

	var count int64
	dbIns := db.Count(&count)
	if dbIns.Error != nil {
		return 0, nil, dbIns.Error
	}

	offset := 0
	if params.Offset != nil && int(*params.Offset) != 0 {
		offset = int(*params.Offset)
		if int64(offset) > count {
			return count, []SwapPair{}, nil
		}
	}

	limit := 50
	if params.Limit != nil && int(*params.Limit) != 0 {
		limit = int(*params.Limit)
	}
	pairs := make([]SwapPair, 0, limit)
	dbIns = db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&pairs)

	return count, pairs, dbIns.Error
}
