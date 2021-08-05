package spot

import "github.com/adshao/go-binance/v2"

// SideType define side type of order
type SideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// MarginTransferType define margin transfer type
type MarginTransferType int

// MarginLoanStatusType define margin loan status type
type MarginLoanStatusType string

// MarginRepayStatusType define margin repay status type
type MarginRepayStatusType string

// FuturesTransferStatusType define futures transfer status type
type FuturesTransferStatusType string

// SideEffectType define side effect type for orders
type SideEffectType string

// FuturesTransferType define futures transfer type
type FuturesTransferType int

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	OrderTypeLimit           OrderType = "LIMIT"
	OrderTypeMarket          OrderType = "MARKET"
	OrderTypeLimitMaker      OrderType = "LIMIT_MAKER"
	OrderTypeStopLoss        OrderType = "STOP_LOSS"
	OrderTypeStopLossLimit   OrderType = "STOP_LOSS_LIMIT"
	OrderTypeTakeProfit      OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitLimit OrderType = "TAKE_PROFIT_LIMIT"

	TimeInForceTypeGTC TimeInForceType = "GTC"
	TimeInForceTypeIOC TimeInForceType = "IOC"
	TimeInForceTypeFOK TimeInForceType = "FOK"

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"
	NewOrderRespTypeFULL   NewOrderRespType = "FULL"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypePendingCancel   OrderStatusType = "PENDING_CANCEL"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"

	SymbolTypeSpot SymbolType = "SPOT"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize          SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePriceFilter      SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice     SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMinNotional      SymbolFilterType = "MIN_NOTIONAL"
	SymbolFilterTypeIcebergParts     SymbolFilterType = "ICEBERG_PARTS"
	SymbolFilterTypeMarketLotSize    SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumAlgoOrders SymbolFilterType = "MAX_NUM_ALGO_ORDERS"

	MarginTransferTypeToMargin MarginTransferType = 1
	MarginTransferTypeToMain   MarginTransferType = 2

	FuturesTransferTypeToFutures FuturesTransferType = 1
	FuturesTransferTypeToMain    FuturesTransferType = 2

	MarginLoanStatusTypePending   MarginLoanStatusType = "PENDING"
	MarginLoanStatusTypeConfirmed MarginLoanStatusType = "CONFIRMED"
	MarginLoanStatusTypeFailed    MarginLoanStatusType = "FAILED"

	MarginRepayStatusTypePending   MarginRepayStatusType = "PENDING"
	MarginRepayStatusTypeConfirmed MarginRepayStatusType = "CONFIRMED"
	MarginRepayStatusTypeFailed    MarginRepayStatusType = "FAILED"

	FuturesTransferStatusTypePending   FuturesTransferStatusType = "PENDING"
	FuturesTransferStatusTypeConfirmed FuturesTransferStatusType = "CONFIRMED"
	FuturesTransferStatusTypeFailed    FuturesTransferStatusType = "FAILED"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func NewPingService(c *binance.Client) *PingService {
	return &PingService{C: c}
}

func NewServerTimeService(c *binance.Client) *ServerTimeService {
	return &ServerTimeService{C: c}
}

func NewBLVTSubscribeService(c *binance.Client) *BLVTSubscribeService {
	return &BLVTSubscribeService{C: c}
}

func NewBLVTRedeemService(c *binance.Client) *BLVTRedeemService {
	return &BLVTRedeemService{C: c}
}

func NewBLVTLimitService(c *binance.Client) *BLVTLimitService {
	return &BLVTLimitService{C: c}
}

func NewBLVTSubscribeRecordService(c *binance.Client) *BLVTSubscribeRecordService {
	return &BLVTSubscribeRecordService{C: c}
}

func NewBLVTInfoService(c *binance.Client) *BLVTInfoService {
	return &BLVTInfoService{C: c}
}

func NewBLVTRedeemRecordService(c *binance.Client) *BLVTRedeemRecordService {
	return &BLVTRedeemRecordService{C: c}
}

func NewWalletGetAllService(c *binance.Client) *WalletGetAllService {
	return &WalletGetAllService{C: c}
}

func NewCreateOrderService(c *binance.Client) *CreateOrderService {
	return &CreateOrderService{C: c}
}

func NewWalletApiService(c *binance.Client) *WalletApiService {
	return &WalletApiService{C: c}
}

func NewGetAccountService(c *binance.Client) *GetAccountService {
	return &GetAccountService{C: c}
}

func NewGetAccountSnapshotService(c *binance.Client) *GetAccountSnapshotService {
	return &GetAccountSnapshotService{C: c}
}

func NewAssetDetailService(c *binance.Client) *AssetDetailService {
	return &AssetDetailService{C: c}
}

func NewWalletEnableFastService(c *binance.Client) *WalletEnableFastService {
	return &WalletEnableFastService{C: c}
}

func NewWalletDisableFastService(c *binance.Client) *WalletDisableFastService {
	return &WalletDisableFastService{C: c}
}

func NewWalletWithdrawService(c *binance.Client) *WalletDisableFastService {
	return &WalletDisableFastService{C: c}
}

func NewDepositsAddressService(c *binance.Client) *DepositsAddressService {
	return &DepositsAddressService{C: c}
}

func NewDepositsHistoryService(c *binance.Client) *DepositsHistoryService {
	return &DepositsHistoryService{C: c}
}

func NewWithdrawHistoryService(c *binance.Client) *WithdrawHistoryService {
	return &WithdrawHistoryService{C: c}
}

func NewAccountStatusService(c *binance.Client) *AccountStatusService {
	return &AccountStatusService{C: c}
}

func NewTransHistoryService(c *binance.Client) *TransHistoryService {
	return &TransHistoryService{C: c}
}

func NewTransferService(c *binance.Client) *TransferService {
	return &TransferService{C: c}
}

func NewParentSpotSummaryService(c *binance.Client) *ParentSpotSummaryService {
	return &ParentSpotSummaryService{C: c}
}

func NewExchangeInfoService(c *binance.Client) *ExchangeInfoService {
	return &ExchangeInfoService{C: c}
}

func NewDepthService(c *binance.Client) *DepthService {
	return &DepthService{C: c}
}

func NewKlinesService(c *binance.Client) *KlinesService {
	return &KlinesService{C: c}
}

func NewStartUserStreamService(c *binance.Client) *StartUserStreamService {
	return &StartUserStreamService{C: c}
}

func NewKeepaliveUserStreamService(c *binance.Client) *KeepaliveUserStreamService {
	return &KeepaliveUserStreamService{C: c}
}

func NewCloseUserStreamService(c *binance.Client) *CloseUserStreamService {
	return &CloseUserStreamService{C: c}
}

func NewListTradesService(c *binance.Client) *ListTradesService {
	return &ListTradesService{C: c}
}

func NewHistoricalTradesService(c *binance.Client) *HistoricalTradesService {
	return &HistoricalTradesService{C: c}
}

func NewAggTradesService(c *binance.Client) *AggTradesService {
	return &AggTradesService{C: c}
}

func NewRecentTradesService(c *binance.Client) *RecentTradesService {
	return &RecentTradesService{C: c}
}

func NewListBookTickersService(c *binance.Client) *ListBookTickersService {
	return &ListBookTickersService{C: c}
}

func NewListPricesService(c *binance.Client) *ListPricesService {
	return &ListPricesService{C: c}
}

func NewListPriceChangeStatsService(c *binance.Client) *ListPriceChangeStatsService {
	return &ListPriceChangeStatsService{C: c}
}

func NewAveragePriceService(c *binance.Client) *AveragePriceService {
	return &AveragePriceService{C: c}
}

func NewFuturesTransferService(c *binance.Client) *FuturesTransferService {
	return &FuturesTransferService{C: c}
}

func NewListFuturesTransferService(c *binance.Client) *ListFuturesTransferService {
	return &ListFuturesTransferService{C: c}
}

func NewMarginTransferService(c *binance.Client) *MarginTransferService {
	return &MarginTransferService{C: c}
}

func NewMarginLoanService(c *binance.Client) *MarginLoanService {
	return &MarginLoanService{C: c}
}

func NewMarginRepayService(c *binance.Client) *MarginRepayService {
	return &MarginRepayService{C: c}
}

func NewListMarginLoansService(c *binance.Client) *ListMarginLoansService {
	return &ListMarginLoansService{C: c}
}

func NewListMarginRepaysService(c *binance.Client) *ListMarginRepaysService {
	return &ListMarginRepaysService{C: c}
}

func NewGetIsolatedMarginAccountService(c *binance.Client) *GetIsolatedMarginAccountService {
	return &GetIsolatedMarginAccountService{C: c}
}

func NewGetMarginAccountService(c *binance.Client) *GetMarginAccountService {
	return &GetMarginAccountService{C: c}
}

func NewGetMarginAssetService(c *binance.Client) *GetMarginAssetService {
	return &GetMarginAssetService{C: c}
}

func NewGetMarginPairService(c *binance.Client) *GetMarginPairService {
	return &GetMarginPairService{C: c}
}

func NewGetMarginAllPairsService(c *binance.Client) *GetMarginAllPairsService {
	return &GetMarginAllPairsService{C: c}
}

func NewGetMarginPriceIndexService(c *binance.Client) *GetMarginPriceIndexService {
	return &GetMarginPriceIndexService{C: c}
}

func NewListMarginTradesService(c *binance.Client) *ListMarginTradesService {
	return &ListMarginTradesService{C: c}
}

func NewGetMaxBorrowableService(c *binance.Client) *GetMaxBorrowableService {
	return &GetMaxBorrowableService{C: c}
}

func NewGetMaxTransferableService(c *binance.Client) *GetMaxTransferableService {
	return &GetMaxTransferableService{C: c}
}

func NewStartIsolatedMarginUserStreamService(c *binance.Client) *StartIsolatedMarginUserStreamService {
	return &StartIsolatedMarginUserStreamService{C: c}
}

func NewKeepaliveIsolatedMarginUserStreamService(c *binance.Client) *KeepaliveIsolatedMarginUserStreamService {
	return &KeepaliveIsolatedMarginUserStreamService{C: c}
}

func NewCloseIsolatedMarginUserStreamService(c *binance.Client) *CloseIsolatedMarginUserStreamService {
	return &CloseIsolatedMarginUserStreamService{C: c}
}

func NewStartMarginUserStreamService(c *binance.Client) *StartMarginUserStreamService {
	return &StartMarginUserStreamService{C: c}
}

func NewKeepaliveMarginUserStreamService(c *binance.Client) *KeepaliveMarginUserStreamService {
	return &KeepaliveMarginUserStreamService{C: c}
}

func NewCloseMarginUserStreamService(c *binance.Client) *CloseMarginUserStreamService {
	return &CloseMarginUserStreamService{C: c}
}

func NewCreateMarginOrderService(c *binance.Client) *CreateMarginOrderService {
	return &CreateMarginOrderService{C: c}
}

func NewCancelMarginOrderService(c *binance.Client) *CancelMarginOrderService {
	return &CancelMarginOrderService{C: c}
}

func NewGetMarginOrderService(c *binance.Client) *GetMarginOrderService {
	return &GetMarginOrderService{C: c}
}

func NewListMarginOpenOrdersService(c *binance.Client) *ListMarginOpenOrdersService {
	return &ListMarginOpenOrdersService{C: c}
}

func NewListMarginOrdersService(c *binance.Client) *ListMarginOrdersService {
	return &ListMarginOrdersService{C: c}
}
