### github.com/adshao/go-binance/v2

A Golang SDK for [binance](https://www.binance.com) API.

[![Build Status](https://travis-ci.org/adshao/github.com/adshao/go-binance/v2.svg?branch=master)](https://travis-ci.org/adshao/github.com/adshao/go-binance/v2)
[![GoDoc](https://godoc.org/github.com/adshao/github.com/adshao/go-binance/v2?status.svg)](https://godoc.org/github.com/adshao/github.com/adshao/go-binance/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/adshao/github.com/adshao/go-binance/v2)](https://goreportcard.com/report/github.com/adshao/github.com/adshao/go-binance/v2)
[![codecov](https://codecov.io/gh/adshao/github.com/adshao/go-binance/v2/branch/master/graph/badge.svg)](https://codecov.io/gh/adshao/github.com/adshao/go-binance/v2)

All the REST APIs listed in [binance API document](https://github.com/binance-exchange/binance-official-api-docs) are implemented, as well as the websocket APIs.

For best compatibility, please use Go >= 1.8.

Make sure you have read binance API document before continuing.

### API List

Name | Description | Status
------------ | ------------ | ------------
[rest-api.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/rest-api.md) | Details on the Rest API (/api) | <input type="checkbox" checked> Implemented
[web-socket-streams.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/web-socket-streams.md) | Details on available streams and payloads | <input type="checkbox" checked>  Implemented
[user-data-stream.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/user-data-stream.md) | Details on the dedicated account stream | <input type="checkbox" checked>  Implemented
[margin-api.md](https://github.com/binance-exchange/binance-official-api-docs/blob/master/margin-api.md) | Details on the Margin API (/sapi) | <input type="checkbox" checked>  Implemented
[futures-api.md](https://binance-docs.github.io/apidocs/futures/en/#general-info) | Details on the Futures API (/fapi) | <input type="checkbox" checked>  Partially Implemented
[delivery-api.md](https://binance-docs.github.io/apidocs/delivery/en/#general-info) | Details on the Coin-M Futures API (/dapi) | <input type="checkbox" checked>  Partially Implemented

### Installation

```shell
go get github.com/adshao/github.com/adshao/go-binance/v2/v2
```

### Importing

```golang
import (
    "github.com/adshao/github.com/adshao/go-binance/v2/v2"
)
```

### Documentation

[![GoDoc](https://godoc.org/github.com/adshao/github.com/adshao/go-binance/v2?status.svg)](https://godoc.org/github.com/adshao/github.com/adshao/go-binance/v2)

### REST API

#### Setup

Init client for API services. Get APIKey/SecretKey from your binance account.

```golang
var (
    apiKey = "your api key"
    secretKey = "your secret key"
)
spotclient := binance.NewClient(apiKey, secretKey)
futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
deliveryClient := binance.NewDeliveryClient(apiKey, secretKey)  // Coin-M Futures
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

Following are some simple examples, please refer to [godoc](https://godoc.org/github.com/adshao/github.com/adshao/go-binance/v2) for full references.

#### Create Order

```golang
order, err := spot.NewCreateOrderService(spotclient).Symbol("BNBETH").
        Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
        TimeInForce(binance.TimeInForceTypeGTC).Quantity("5").
        Price("0.0030000").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(order)

// Use Test() instead of Do() for testing.
```

#### Get Order

```golang
order, err := spot.NewGetOrderService(spotclient).Symbol("BNBETH").
    OrderID(4432844).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(order)
```

#### Cancel Order

```golang
_, err := spot.NewCancelOrderService(spotclient).Symbol("BNBETH").
    OrderID(4432844).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
```

#### List Open Orders

```golang
openOrders, err := spot.NewListOpenOrdersService(spotClient).Symbol("BNBETH").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, o := range openOrders {
    fmt.Println(o)
}
```

#### List Orders

```golang
orders, err := spot.NewListOrdersService(spotClient).Symbol("BNBETH").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, o := range orders {
    fmt.Println(o)
}
```

#### List Ticker Prices

```golang
prices, err := spot.NewListPricesService(spotClient).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, p := range prices {
    fmt.Println(p)
}
```

#### Show Depth

```golang
res, err := spot..NewDepthService(spotClient).Symbol("LTCBTC").
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

#### List Klines

```golang
klines, err := spot.NewKlinesService(spotClient).Symbol("LTCBTC").
    Interval("15m").Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, k := range klines {
    fmt.Println(k)
}
```

#### List Aggregate Trades

```golang
trades, err := spot.NewAggTradesService(spotClient).
    Symbol("LTCBTC").StartTime(1508673256594).EndTime(1508673256595).
    Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
for _, t := range trades {
    fmt.Println(t)
}
```

#### Get Account

```golang
res, err := spot.NewGetAccountService(spotClient).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

#### Start User Stream

```golang
res, err := spot.NewStartUserStreamService(spotClient).Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)
```

### Websocket

You don't need Client in websocket API. Just call binance.WsXxxServe(args, handler, errHandler).

> For delivery API you can use `delivery.WsXxxServe(args, handler, errHandler)`.

#### Depth

```golang
wsDepthHandler := func(event *binance.WsDepthEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, stopC, err := binance.WsDepthServe("LTCBTC", wsDepthHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
// use stopC to exit
go func() {
    time.Sleep(5 * time.Second)
    stopC <- struct{}{}
}()
// remove this if you do not want to be blocked here
<-doneC
```

#### Kline

```golang
wsKlineHandler := func(event *binance.WsKlineEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsKlineServe("LTCBTC", "1m", wsKlineHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### Aggregate

```golang
wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
    fmt.Println(event)
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsAggTradeServe("LTCBTC", wsAggTradeHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### User Data

```golang
wsHandler := func(message []byte) {
    fmt.Println(string(message))
}
errHandler := func(err error) {
    fmt.Println(err)
}
doneC, _, err := binance.WsUserDataServe(listenKey, wsHandler, errHandler)
if err != nil {
    fmt.Println(err)
    return
}
<-doneC
```

#### Setting Server Time

Your system time may be incorrect and you may use following function to set the time offset based off Binance Server Time:

```golang
// use the client future for Futures
spot.NewSetServerTimeService(futuresClient).Do(context.Background())
```

Or you can also overwrite the `TimeOffset` yourself:

```golang
client.TimeOffset = 123
```

### Testnet

You can use the testnet by enabling the corresponding flag.

> Note that you can't use your regular API and Secret keys for the testnet. You have to create an account on
> the testnet websites : [https://testnet.binancefuture.com/](https://testnet.binancefuture.com/) for futures and delivery
> or [https://testnet.binance.vision/](https://testnet.binance.vision/) for the Spot Test Network.

#### Spot

You can create a new test client Or
Use the `binance.UseTestnet` flag before calling the client creation and the websockets methods.

```go
import (
    "github.com/adshao/github.com/adshao/go-binance/v2/v2"
)

spot.UseTestnet = true

client := binance.NewTestClient(apiKey, secretKey)
```

#### Futures (usd(s)-m futures)

You can create a new test client Or
Use the `futures.UseTestnet` flag before calling the websockets methods

```go
import (
    "github.com/adshao/github.com/adshao/go-binance/v2/v2/futures"
)
futures.UseWsTestnet = true
futuresClient = futures.NewFuturesTestClient(ApiKey, SecretKey)
```

#### Delivery (coin-m futures)

You can create a new test client Or
Use the `delivery.UseWsTestnet` flag before calling the websockets methods

```go
import (
    "github.com/adshao/github.com/adshao/go-binance/v2/v2/delivery"
)

delivery.UseWsTestnet = true
delivery.NewDeliveryTestClient(ApiKey, SecretKey)
```

