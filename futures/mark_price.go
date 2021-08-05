package futures

import (
	"context"
	"encoding/json"
	"github.com/firekeeper76/go-binance"
	"github.com/firekeeper76/go-binance/common"
)

// PremiumIndexService get premium index
type PremiumIndexService struct {
	C      *binance.Client
	symbol *string
}

// Symbol set symbol
func (s *PremiumIndexService) Symbol(symbol string) *PremiumIndexService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *PremiumIndexService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*PremiumIndex, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/premiumIndex",
		SecType:  binance.SecTypeNone,
	}
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	data = common.ToJSONList(data)
	res = make([]*PremiumIndex, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// PremiumIndex define premium index of mark price
type PremiumIndex struct {
	Symbol          string `json:"symbol"`
	MarkPrice       string `json:"markPrice"`
	LastFundingRate string `json:"lastFundingRate"`
	NextFundingTime int64  `json:"nextFundingTime"`
	Time            int64  `json:"time"`
}

// FundingRateService get funding rate
type FundingRateService struct {
	C         *binance.Client
	symbol    string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *FundingRateService) Symbol(symbol string) *FundingRateService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *FundingRateService) StartTime(startTime int64) *FundingRateService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *FundingRateService) EndTime(endTime int64) *FundingRateService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *FundingRateService) Limit(limit int) *FundingRateService {
	s.limit = &limit
	return s
}

// Do send request
func (s *FundingRateService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*FundingRate, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/fundingRate",
		SecType:  binance.SecTypeNone,
	}
	r.SetParam("symbol", s.symbol)
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
	res = make([]*FundingRate, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// FundingRate define funding rate of mark price
type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	Time        int64  `json:"time"`
}

// GetLeverageBracketService get funding rate
type GetLeverageBracketService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *GetLeverageBracketService) Symbol(symbol string) *GetLeverageBracketService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetLeverageBracketService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*LeverageBracket, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/leverageBracket",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.symbol != "" {
		r.SetParam("symbol", s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}

	if s.symbol != "" {
		data = common.ToJSONList(data)
	}
	res = make([]*LeverageBracket, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// LeverageBracket define the leverage bracket
type LeverageBracket struct {
	Symbol   string    `json:"symbol"`
	Brackets []Bracket `json:"brackets"`
}

// Bracket define the bracket
type Bracket struct {
	Bracket          int     `json:"bracket"`
	InitialLeverage  int     `json:"initialLeverage"`
	NotionalCap      float64 `json:"notionalCap"`
	NotionalFloor    float64 `json:"notionalFloor"`
	MaintMarginRatio float64 `json:"maintMarginRatio"`
	Cum              float64 `json:"cum"`
}
