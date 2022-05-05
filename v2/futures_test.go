package binance_test

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"testing"
	"time"
)

func TestWs(t *testing.T) {
	binance.WebsocketKeepalive = true
	wsHandler := func(event *futures.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println("BTCUSDT SubDepth:", err)
	}
	_, _, err := futures.WsPartialDepthServeWithRate("BTCUSDT", 5, 100*time.Millisecond, wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
		return
	}
	select {}
}

func TestNewStartUserStreamService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewStartUserStreamService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}

func TestNewKeepaliveUserStreamService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	err := futures.NewKeepaliveUserStreamService(c).ListenKey("3VLu3Ew6AiNrPbjNQDr2sjvfmafYZvVK5tEHDOJpfmKRQI59sRWZ8t7WFZywuMnZ").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
}

func TestNewCloseUserStreamService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	err := futures.NewCloseUserStreamService(c).ListenKey("3VLu3Ew6AiNrPbjNQDr2sjvfmafYZvVK5tEHDOJpfmKRQI59sRWZ8t7WFZywuMnZ").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
}

func TestNewDepthService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewDepthService(c).Symbol("BTCUSDT").Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res.Bids {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewKlinesService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewKlinesService(c).Symbol("BTCUSDT").Interval("1m").Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestGetAccountService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	_, err := futures.NewGetAccountService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	//for _, v := range res.Positions {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestNewExchangeInfoService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewExchangeInfoService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res.RateLimits {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewPremiumIndexService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewPremiumIndexService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewFundingRateService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewFundingRateService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewGetLeverageBracketService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetLeverageBracketService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewCreateOrderService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewCreateOrderService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestNewListOpenOrdersService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListOpenOrdersService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestNewGetOrderService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetOrderService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestNewListOrdersService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListOrdersService(c).Symbol("BAKEUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	//fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewCancelOrderService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewCancelOrderService(c).Symbol("BAKEUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestNewListLiquidationOrdersService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListUserLiquidationOrdersService(c).Symbol("BAKEUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewGetPositionRiskService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetPositionRiskService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
func TestNewChangeLeverageService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewChangeLeverageService(c).Symbol("BAKEUSDT").Leverage(20).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestNewGetPositionModeService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetPositionModeService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestNewListBookTickersService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListBookTickersService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
func TestNewListPricesService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListPricesService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewListPriceChangeStatsService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListPriceChangeStatsService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewGetPositionMarginHistoryService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetPositionMarginHistoryService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewHistoricalTradesService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewHistoricalTradesService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewRecentTradesService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewRecentTradesService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewAggTradesService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewAggTradesService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewListAccountTradeService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewListAccountTradeService(c).Symbol("BAKEUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewGetIncomeHistoryService(t *testing.T) {
	c := binance.NewFuturesClient(key, secret)
	c.Debug = true
	res, err := futures.NewGetIncomeHistoryService(c).Symbol("BAKEUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
