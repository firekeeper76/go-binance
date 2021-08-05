package futures

import (
	"github.com/adshao/go-binance/v2"
	"context"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	C *binance.Client
}

// Do send request
func (s *StartUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (listenKey string, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/fapi/v1/listenKey",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return "", rErr
	}
	j, jErr := binance.NewJSON(data)
	if jErr != nil {
		return "", binance.NewApiErr(jErr.Error())
	}
	listenKey = j.Get("listenKey").MustString()
	return listenKey, nil
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	C         *binance.Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "PUT",
		Endpoint: "/fapi/v1/listenKey",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParam("listenKey", s.listenKey)
	_, err = s.C.Request(ctx, r, opts...)
	return err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	C         *binance.Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "DELETE",
		Endpoint: "/fapi/v1/listenKey",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParam("listenKey", s.listenKey)
	_, err = s.C.Request(ctx, r, opts...)
	return err
}
