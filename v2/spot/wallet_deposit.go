package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance"
)

type DepositsAddressService struct {
	C       *binance.Client
	coin    string
	network *string
}

// Coin sets the asset parameter (MANDATORY).
func (s *DepositsAddressService) Coin(v string) *DepositsAddressService {
	s.coin = v
	return s
}

// Network sets the status parameter.
func (s *DepositsAddressService) Network(v string) *DepositsAddressService {
	s.network = &v
	return s
}

// Do sends the request.
func (s *DepositsAddressService) Do(ctx context.Context, opts ...binance.RequestOption) (*DepositAddressResponse, *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/capital/deposit/address",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("coin", s.coin)
	if v := s.network; v != nil {
		r.SetParam("network", *v)
	}

	data, err := s.C.Request(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(DepositAddressResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// DepositAddressResponse represents a response from GetDepositsAddressService.
type DepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"Tag"`
	URL     string `json:"url"`
}

type DepositsHistoryService struct {
	C         *binance.Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int64
	limit     *int64
}

// Coin sets the asset parameter.
func (s *DepositsHistoryService) Coin(asset string) *DepositsHistoryService {
	s.coin = &asset
	return s
}

// Status sets the status parameter.0(0:pending,6: credited but cannot withdraw, 1:success)
func (s *DepositsHistoryService) Status(status int) *DepositsHistoryService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *DepositsHistoryService) StartTime(startTime int64) *DepositsHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *DepositsHistoryService) EndTime(endTime int64) *DepositsHistoryService {
	s.endTime = &endTime
	return s
}
func (s *DepositsHistoryService) Offset(offset int64) *DepositsHistoryService {
	s.offset = &offset
	return s
}
func (s *DepositsHistoryService) Limit(limit int64) *DepositsHistoryService {
	s.limit = &limit
	return s
}

// Do sends the request.
func (s *DepositsHistoryService) Do(ctx context.Context, opts ...binance.RequestOption) (deposits []*Deposit, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/capital/deposit/hisrec",
		SecType:  binance.SecTypeSigned,
	}
	if s.coin != nil {
		r.SetParam("coin", *s.coin)
	}
	if s.status != nil {
		r.SetParam("status", *s.status)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.SetParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}

	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res := make([]*Deposit, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Deposit represents a single deposit entry.
type Deposit struct {
	Amount       string `json:"amount"`
	Coin         string `json:"coin"`
	Network      string `json:"network"`
	Status       int    `json:"status"`
	Address      string `json:"address"`
	AddressTag   string `json:"addressTag"`
	TxID         string `json:"txId"`
	InsertTime   int64  `json:"insertTime"`
	TransferType int    `json:"transferType"`
	ConfirmTimes string `json:"confirmTimes"`
}
