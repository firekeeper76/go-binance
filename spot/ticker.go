package spot

import (
	"github.com/firekeeper76/go-binance"
	"github.com/firekeeper76/go-binance/common"
	"context"
	"encoding/json"
)

// ListBookTickersService list best price/qty on the order book for a symbol or symbols
type ListBookTickersService struct {
	C      *binance.Client
	symbol *string
}

// Symbol set symbol
func (s *ListBookTickersService) Symbol(symbol string) *ListBookTickersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListBookTickersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*BookTicker, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	data = common.ToJSONList(data)
	res = make([]*BookTicker, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// BookTicker define book ticker info
type BookTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

// ListPricesService list latest price for a symbol or symbols
type ListPricesService struct {
	C      *binance.Client
	symbol *string
}

// Symbol set symbol
func (s *ListPricesService) Symbol(symbol string) *ListPricesService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*SymbolPrice, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/ticker/price",
	}
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	data = common.ToJSONList(data)
	res = make([]*SymbolPrice, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// SymbolPrice define symbol and price pair
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// ListPriceChangeStatsService show stats of price change in last 24 hours for all symbols
type ListPriceChangeStatsService struct {
	C      *binance.Client
	symbol *string
}

// Symbol set symbol
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*PriceChangeStats, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/ticker/24hr",
	}
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	data = common.ToJSONList(data)
	res = make([]*PriceChangeStats, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// PriceChangeStats define price change stats
type PriceChangeStats struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FristID            int64  `json:"firstId"`
	LastID             int64  `json:"lastId"`
	Count              int64  `json:"count"`
}

// AveragePriceService show current average price for a symbol
type AveragePriceService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *AveragePriceService) Symbol(symbol string) *AveragePriceService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *AveragePriceService) Do(ctx context.Context, opts ...binance.RequestOption) (res *AvgPrice, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/avgPrice",
	}
	r.SetParam("symbol", s.symbol)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(AvgPrice)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// AvgPrice define average price
type AvgPrice struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}