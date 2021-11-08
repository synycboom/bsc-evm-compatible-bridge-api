package dao

import (
	"strings"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/restapi/operations/erc_1155_swaps"
)

type SwapState string
type SwapDirection string

const (
	SwapStateRequestOngoing     SwapState = "request_ongoing"
	SwapStateRequestRejected    SwapState = "request_rejected"
	SwapStateRequestConfirmed   SwapState = "request_confirmed"
	SwapStateFillTxDryRunFailed SwapState = "fill_tx_dry_run_failed"
	SwapStateFillTxCreated      SwapState = "fill_tx_created"
	SwapStateFillTxSent         SwapState = "fill_tx_sent"
	SwapStateFillTxConfirmed    SwapState = "fill_tx_confirmed"
	SwapStateFillTxFailed       SwapState = "fill_tx_failed"
	SwapStateFillTxMissing      SwapState = "fill_tx_missing"

	SwapDirectionForward  SwapDirection = "forward"
	SwapDirectionBackward SwapDirection = "backward"
)

type Swap struct {
	ID string

	// Basic Token Information
	SrcChainID   string
	DstChainID   string
	SrcTokenAddr string
	DstTokenAddr string
	Sender       string
	Recipient    string
	IDs          datatypes.JSON `gorm:"not null"`
	Amounts      datatypes.JSON `gorm:"not null"`
	Signature    string

	// Swap State
	State         SwapState
	SwapDirection SwapDirection

	// Request Transaction Information
	RequestTxHash     string
	RequestHeight     int64
	RequestBlockHash  string
	RequestTrackRetry int64

	// Fill Transaction Information
	FillConsumedFeeAmount string
	FillGasPrice          string
	FillGasUsed           int64
	FillHeight            int64
	FillTxHash            string
	FillTrackRetry        int64
	FillBlockHash         string

	MessageLog string

	// Timestamp
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Swap) TableName() string {
	return "erc1155_swaps"
}

type SwapDaoInterface interface {
	GetSwaps(params erc_1155_swaps.GetErc1155SwapsParams) (int64, []Swap, error)
}

type SwapDao struct {
	db *gorm.DB
}

func NewSwapDao(db *gorm.DB) SwapDaoInterface {
	return &SwapDao{db}
}

func (s *SwapDao) GetSwaps(params erc_1155_swaps.GetErc1155SwapsParams) (int64, []Swap, error) {
	db := s.db.Model(&Swap{})

	if params.State != nil && *params.State != "" {
		db = db.Where("state = ?", strings.ToLower(*params.State))
	}
	if params.RequestTxHash != nil && *params.RequestTxHash != "" {
		db = db.Where("request_tx_hash = ?", strings.ToLower(*params.RequestTxHash))
	}

	db = db.Where("sender = ?", strings.ToLower(params.Sender))

	var count int64
	dbIns := db.Count(&count)
	if dbIns.Error != nil {
		return 0, nil, dbIns.Error
	}

	offset := 0
	if params.Offset != nil && int(*params.Offset) != 0 {
		offset = int(*params.Offset)
		if int64(offset) > count {
			return count, []Swap{}, nil
		}
	}

	limit := 50
	if params.Limit != nil && int(*params.Limit) != 0 {
		limit = int(*params.Limit)
	}

	swaps := make([]Swap, 0, limit)
	dbIns = db.Offset(int(offset)).Limit(int(limit)).Order("created_at DESC").Find(&swaps)

	return count, swaps, dbIns.Error
}
