package spot

import (
	"github.com/firekeeper76/go-binance"
	"context"
)

type ServerTimeService struct {
	C *binance.Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context) (serverTime int64, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/time",
		SecType:  binance.SecTypeNone,
	}
	data, rErr := s.C.Request(ctx, r)
	if rErr != nil {
		return 0, rErr
	}
	j, jErr := binance.NewJSON(data)
	if jErr != nil {
		return 0, binance.NewApiErr(jErr.Error())
	}
	serverTime = j.Get("serverTime").MustInt64()
	return serverTime, nil
}

// PingService ping server
type PingService struct {
	C *binance.Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/ping",
	}
	_, err = s.C.Request(ctx, r, opts...)
	return err
}
