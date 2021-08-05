package futures

import (
	"github.com/firekeeper76/go-binance"
	"context"
	"encoding/json"
)

// GetIncomeHistoryService get position margin history service
type GetIncomeHistoryService struct {
	C          *binance.Client
	symbol     string
	incomeType string
	startTime  *int64
	endTime    *int64
	limit      *int64
}

// Symbol set symbol
func (s *GetIncomeHistoryService) Symbol(symbol string) *GetIncomeHistoryService {
	s.symbol = symbol
	return s
}

// IncomeType set income type
func (s *GetIncomeHistoryService) IncomeType(incomeType string) *GetIncomeHistoryService {
	s.incomeType = incomeType
	return s
}

// StartTime set startTime
func (s *GetIncomeHistoryService) StartTime(startTime int64) *GetIncomeHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetIncomeHistoryService) EndTime(endTime int64) *GetIncomeHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetIncomeHistoryService) Limit(limit int64) *GetIncomeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetIncomeHistoryService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*IncomeHistory, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/income",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.incomeType != "" {
		r.SetParam("incomeType", s.incomeType)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*IncomeHistory, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// IncomeHistory define position margin history info
type IncomeHistory struct {
	Asset      string `json:"asset"`
	Income     string `json:"income"`
	IncomeType string `json:"incomeType"`
	Info       string `json:"info"`
	Symbol     string `json:"symbol"`
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`
	TradeID    string `json:"tradeId"`
}
