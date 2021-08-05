package futures

import (
	"github.com/adshao/go-binance"
	"context"
	"encoding/json"
)

// GetPositionMarginHistoryService get position margin history service
type GetPositionMarginHistoryService struct {
	C         *binance.Client
	symbol    string
	_type     *int
	startTime *int64
	endTime   *int64
	limit     *int64
}

// Symbol set symbol
func (s *GetPositionMarginHistoryService) Symbol(symbol string) *GetPositionMarginHistoryService {
	s.symbol = symbol
	return s
}

// Type set type
func (s *GetPositionMarginHistoryService) Type(_type int) *GetPositionMarginHistoryService {
	s._type = &_type
	return s
}

// StartTime set startTime
func (s *GetPositionMarginHistoryService) StartTime(startTime int64) *GetPositionMarginHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *GetPositionMarginHistoryService) EndTime(endTime int64) *GetPositionMarginHistoryService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetPositionMarginHistoryService) Limit(limit int64) *GetPositionMarginHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetPositionMarginHistoryService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*PositionMarginHistory, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/positionMargin/history",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s._type != nil {
		r.SetParam("type", *s._type)
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
	res = make([]*PositionMarginHistory, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// PositionMarginHistory define position margin history info
type PositionMarginHistory struct {
	Amount       string `json:"amount"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
	Type         int    `json:"type"`
	PositionSide string `json:"positionSide"`
}
