package binance_test

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance"
	"github.com/adshao/go-binance/delivery"
	"github.com/adshao/go-binance/futures"
	"testing"
	"time"
)

func TestWsDelivery(t *testing.T) {
	futures.WebsocketKeepalive = true
	delivery.WebsocketKeepalive = true
	wsHandler := func(event *delivery.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println("BTCUSDT SubDepth:", err)
	}
	_, _, err := delivery.WsKlineServe("BTCUSD_210924", "1m", wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
		return
	}
	select {}
}

func TestUserWsDelivery(t *testing.T) {
	delivery.WebsocketKeepalive = true
	wsHandler := func(event []byte) {
		fmt.Println(string(event))
	}
	errHandler := func(err error) {
		fmt.Println("BTCUSDT SubDepth:", err)
	}
	_, _, err := delivery.WsUserDataServe("jgI1sdvJb1idpM20P1daQM8b7uDfedX30D2vavufS0CBBhe5mGOS14YfXan4LWlT", wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
		return
	}
	select {}
}
func TestDeliveryServerTime(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewServerTimeService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestNewGetBalanceService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewGetBalanceService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewGetAccountService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewGetAccountService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res.Assets {
		fmt.Printf("%+v \n", v)
	}
}

func TestUserStream(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewStartUserStreamService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestKeepUserStream(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	err := delivery.NewKeepaliveUserStreamService(c).ListenKey("EbCF3ErHJ7FLnBfMsy2xlDyRh4EzuxSmREt7kaLeQmCoOeTP2Kr0ne0lTkUhjQ0w").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
}
func TestCloseUserStream(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	err := delivery.NewCloseUserStreamService(c).ListenKey("EbCF3ErHJ7FLnBfMsy2xlDyRh4EzuxSmREt7kaLeQmCoOeTP2Kr0ne0lTkUhjQ0w").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
}

func TestDeliveryNewExchangeInfoService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewExchangeInfoService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res.Symbols {
		fmt.Printf("%+v \n", v.Symbol)
	}
}

func TestDeliveryNewKlinesService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewKlinesService(c).Symbol("BTCUSD_210924").Interval("1m").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestGetPositionRiskService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewGetPositionRiskService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestDeliveryNewGetPositionModeService(t *testing.T) {
	c := binance.NewDeliveryClient(key, secret)
	c.Debug = true
	res, err := delivery.NewGetPositionModeService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
