package futures

import (
	"context"
	"github.com/adshao/go-binance/v2"
)

// BlvtKlinesService list klines
type BlvtKlinesService struct {
	C         *binance.Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol set symbol
func (s *BlvtKlinesService) Symbol(symbol string) *BlvtKlinesService {
	s.symbol = symbol
	return s
}

// Interval set interval
func (s *BlvtKlinesService) Interval(interval string) *BlvtKlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *BlvtKlinesService) Limit(limit int) *BlvtKlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BlvtKlinesService) StartTime(startTime int64) *BlvtKlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BlvtKlinesService) EndTime(endTime int64) *BlvtKlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *BlvtKlinesService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*Kline, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/lvtKlines",
	}
	r.SetParam("symbol", s.symbol)
	r.SetParam("interval", s.interval)
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	j, jErr := binance.NewJSON(data)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	num := len(j.MustArray())
	res = make([]*Kline, num)
	for i := 0; i < num; i++ {
		item := j.GetIndex(i)
		if len(item.MustArray()) < 11 {
			return nil, binance.NewApiErr("invalid kline response")
		}
		res[i] = &Kline{
			OpenTime:                 item.GetIndex(0).MustInt64(),
			Open:                     item.GetIndex(1).MustString(),
			High:                     item.GetIndex(2).MustString(),
			Low:                      item.GetIndex(3).MustString(),
			Close:                    item.GetIndex(4).MustString(),
			Volume:                   item.GetIndex(5).MustString(),
			CloseTime:                item.GetIndex(6).MustInt64(),
			QuoteAssetVolume:         item.GetIndex(7).MustString(),
			TradeNum:                 item.GetIndex(8).MustInt64(),
			TakerBuyBaseAssetVolume:  item.GetIndex(9).MustString(),
			TakerBuyQuoteAssetVolume: item.GetIndex(10).MustString(),
		}
	}
	return res, nil
}
