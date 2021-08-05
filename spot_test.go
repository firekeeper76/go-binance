package binance_test

import (
	"go-binance"
	"go-binance/futures"
	"go-binance/spot"
	"context"
	"fmt"
	"testing"
	"time"
)

var (
	key    = ""
	secret = ""
)

func TestSpotWs(t *testing.T) {
	spot.WebsocketKeepalive = true
	wsHandler := func(event *spot.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println("BTCUSDT SubDepth:", err)
	}
	_, _, err := spot.WsDepthServe("BTCUSDT", wsHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second)
		return
	}
	select {}
}
func TestGetServerTime(t *testing.T) {
	c := binance.NewFuturesClient("", "")
	c.Debug = true
	st, err := futures.NewServerTimeService(c).Do(context.Background())
	fmt.Println(err)
	fmt.Println(st)
}
func TestGetAccount(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	st, err := spot.NewGetAccountService(c).Do(context.Background())
	fmt.Println(err)
	fmt.Println(st)
}
func TestGetApi(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	st, err := spot.NewWalletApiService(c).Do(context.Background())
	fmt.Println(err)
	fmt.Println(st)
}
func TestGetWalletGetAll(t *testing.T) {
	c := binance.NewClient(key, secret)
	st, err := spot.NewWalletGetAllService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
	}
	fmt.Println(st)
}
func TestGetDeposit(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	st, err := spot.NewDepositsAddressService(c).Coin("USDT").Network("bsc").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", st)
}
func TestGetAssetDetail(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	st, err := spot.NewAssetDetailService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", st)
}
func TestGetLimit(t *testing.T) {
	c := binance.NewClient(key, secret)
	st, err := spot.NewBLVTLimitService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", st)
}
func TestGetDepositHis(t *testing.T) {
	c := binance.NewClient(key, secret)
	st, err := spot.NewDepositsHistoryService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", st)
}

func TestWithdrawHis(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewWithdrawHistoryService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v", v)
	}
}

func TestAccountStatus(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewAccountStatusService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
}
func TestTransHistory(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	res, err := spot.NewTransHistoryService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res.UserAssetDribblets {
		fmt.Printf("%+v", v)
	}
}
func TestParentSpotSummary(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewParentSpotSummaryService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res.SpotSubUserAssetBtcVoList {
		fmt.Printf("%+v", v)
	}
}
func TestExchangeInfo(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewExchangeInfoService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res.RateLimits {
		fmt.Printf("%+v \n", v)
	}
}
func TestDepth(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewDepthService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res.Bids {
		fmt.Printf("%+v", v)
	}
}
func TestKline(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewKlinesService(c).Symbol("BTCUSDT").Interval("1351").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res {
		fmt.Printf("%+v", v.Close)
	}
}

func TestMyTrades(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListTradesService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res {
		fmt.Printf("%+v", v)
	}
}
func TestHistoricalTrades(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewHistoricalTradesService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res {
		fmt.Printf("%+v", v)
	}
}
func TestAggTrades(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewAggTradesService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res {
		fmt.Printf("%+v", v)
	}
}
func TestRecentTrades(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewRecentTradesService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
	for _, v := range res {
		fmt.Printf("%+v", v)
	}
}

func TestListBookTickers(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListBookTickersService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
func TestListPrices(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListPricesService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestListPriceChangeStats(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListPriceChangeStatsService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestAveragePrice(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewAveragePriceService(c).Symbol("SUSHIUPUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestListFuturesTransfer(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListFuturesTransferService(c).Asset("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestListMarginLoans(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListMarginLoansService(c).Asset("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestListMarginRepays(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListMarginRepaysService(c).Asset("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}

func TestGetIsolatedMarginAccountService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetIsolatedMarginAccountService(c).Symbols("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}
func TestGetMarginAccountService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginAccountService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	//for _, v := range res {
	//	fmt.Printf("%+v \n", v)
	//}
}
func TestGetMarginAssetService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginAssetService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestNewGetMarginPairService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginPairService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestGetMarginAllPairsService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginAllPairsService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}
func TestNewGetMarginPriceIndexService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginPriceIndexService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)

}

func TestNewListMarginTradesService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListMarginTradesService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)

}
func TestGetMaxBorrowableService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMaxBorrowableService(c).Asset("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestGetMaxTransferableService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMaxTransferableService(c).Asset("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}

func TestNewStartIsolatedMarginUserStreamService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewStartIsolatedMarginUserStreamService(c).Symbol("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}
func TestNewGetMarginOrderService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewGetMarginOrderService(c).Symbol("BTC").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}

func TestNewListMarginOpenOrdersService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListMarginOpenOrdersService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}

func TestNewListMarginOrdersService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewListMarginOrdersService(c).Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
}

func TestEtfLimit(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewBLVTLimitService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestEtfRedeemRecord(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewBLVTRedeemRecordService(c).TokenName("BTCDOWN").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewBLVTSubscribeRecordService(t *testing.T) {
	c := binance.NewClient(key, secret)
	res, err := spot.NewBLVTSubscribeRecordService(c).TokenName("SUSHIUP").Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	fmt.Printf("%+v \n", res)
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}

func TestNewBLVTInfoService(t *testing.T) {
	c := binance.NewClient(key, secret)
	c.Debug = true
	res, err := spot.NewBLVTInfoService(c).Do(context.Background())
	if err != nil {
		fmt.Println(err.Header)
		fmt.Println(err)
	}
	for _, v := range res {
		fmt.Printf("%+v \n", v)
	}
}