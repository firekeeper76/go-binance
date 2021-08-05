package binance_test

import (
	"context"
	"encoding/json"
	"fmt"
	"go-binance"
	"go-binance/futures"
	"go-binance/spot"
	"testing"
)

func TestChanged(t *testing.T) {
	client := binance.NewFuturesClient("key", "secret")
	// change 1. 接口请求service改为外部注入client
	res, err := spot.NewServerTimeService(client).
		Do(context.Background())
	if err != nil {
		// change 2. err *APIError 包含接口返回的header和http状态码
		// 有些信息都在header返回  原sdk没有header信息
		fmt.Printf("%+v", err.Header)
		fmt.Printf("%+v", err.HttpCode)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v", res)
}

func TestChangedFutures(t *testing.T) {
	client :=  binance.NewClient("key", "secret")
	// change 1. 接口请求service改为外部注入client
	res, err := futures.NewServerTimeService(client).
		Do(context.Background())
	if err != nil {
		// change 2. err *APIError 包含接口返回的header和http状态码
		// 有些信息都在header返回  原sdk没有header信息
		fmt.Printf("%+v", err.Header)
		fmt.Printf("%+v", err.HttpCode)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v", res)
}

// change 3. 可在外部自定义service 有些sdk没有实现的接口 可以自己实现
type CustomService struct {
	C  *binance.Client
}

type CustomResponse struct {
}

func (s *CustomService) Do(ctx context.Context) (*CustomResponse, error) {
	r := & binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/接口",
		SecType:   binance.SecTypeNone,
	}
	res, err := s.C.Request(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := new(CustomResponse)
	jErr := json.Unmarshal(res, resp)
	if jErr != nil {
		return nil,  binance.NewApiErr(jErr.Error())
	}
	return resp, nil
}
