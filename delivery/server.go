package delivery

import (
	"go-binance"
	"context"
)

// PingService ping server
type PingService struct {
	C *binance.Client
}

// Do send request
func (s *PingService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/ping",
	}
	_, rErr := s.C.Request(ctx, r, opts...)
	return rErr
}

// ServerTimeService get server time
type ServerTimeService struct {
	C *binance.Client
}

// Do send request
func (s *ServerTimeService) Do(ctx context.Context, opts ...binance.RequestOption) (serverTime int64, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/time",
	}
	data, rErr := s.C.Request(ctx, r, opts...)
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

// SetServerTimeService set server time
type SetServerTimeService struct {
	C *binance.Client
}

// Do send request
func (s *SetServerTimeService) Do(ctx context.Context, opts ...binance.RequestOption) (timeOffset int64, err *binance.APIError) {
	serverTime, rErr := NewServerTimeService(s.C).Do(ctx)
	if rErr != nil {
		return 0, rErr
	}
	timeOffset = binance.CurrentTimestamp() - serverTime
	s.C.TimeOffset = timeOffset
	return timeOffset, nil
}
