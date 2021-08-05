package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance"
)

// FuturesTransferService transfer asset between spot account and futures account
type FuturesTransferService struct {
	C            *binance.Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *FuturesTransferService) Asset(asset string) *FuturesTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *FuturesTransferService) Amount(amount string) *FuturesTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from spot account to futures account 2: transfer from futures account to spot account
func (s *FuturesTransferService) Type(transferType FuturesTransferType) *FuturesTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *FuturesTransferService) Do(ctx context.Context, opts ...binance.RequestOption) (res *TransactionResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/futures/transfer",
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.SetFormParams(m)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(TransactionResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// ListFuturesTransferService list futures transfer
type ListFuturesTransferService struct {
	C         *binance.Client
	asset     string
	startTime int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListFuturesTransferService) Asset(asset string) *ListFuturesTransferService {
	s.asset = asset
	return s
}

// StartTime set start time
func (s *ListFuturesTransferService) StartTime(startTime int64) *ListFuturesTransferService {
	s.startTime = startTime
	return s
}

// EndTime set end time
func (s *ListFuturesTransferService) EndTime(endTime int64) *ListFuturesTransferService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListFuturesTransferService) Current(current int64) *ListFuturesTransferService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListFuturesTransferService) Size(size int64) *ListFuturesTransferService {
	s.size = &size
	return s
}

// Do send request
func (s *ListFuturesTransferService) Do(ctx context.Context, opts ...binance.RequestOption) (res *FuturesTransferHistory, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/futures/transfer",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParams(binance.Params{
		"asset":     s.asset,
		"startTime": s.startTime,
	})
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(FuturesTransferHistory)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// FuturesTransferHistory define futures transfer history
type FuturesTransferHistory struct {
	Rows  []FuturesTransfer `json:"rows"`
	Total int64             `json:"total"`
}

// FuturesTransfer define futures transfer history item
type FuturesTransfer struct {
	Asset     string                    `json:"asset"`
	TranID    int64                     `json:"tranId"`
	Amount    string                    `json:"amount"`
	Type      int64                     `json:"type"`
	Timestamp int64                     `json:"timestamp"`
	Status    FuturesTransferStatusType `json:"status"`
}

