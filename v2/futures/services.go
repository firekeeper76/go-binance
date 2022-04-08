package futures

import "github.com/adshao/go-binance/v2"

// SideType define side type of order
type SideType string

// PositionSideType define position side type of order
type PositionSideType string

// OrderType define order type
type OrderType string

// TimeInForceType define time in force type of order
type TimeInForceType string

// NewOrderRespType define response JSON verbosity
type NewOrderRespType string

// OrderExecutionType define order execution type
type OrderExecutionType string

// OrderStatusType define order status type
type OrderStatusType string

// SymbolType define symbol type
type SymbolType string

// SymbolStatusType define symbol status type
type SymbolStatusType string

// SymbolFilterType define symbol filter type
type SymbolFilterType string

// SideEffectType define side effect type for orders
type SideEffectType string

// WorkingType define working type
type WorkingType string

// MarginType define margin type
type MarginType string

// ContractType define contract type
type ContractType string

// UserDataEventType define user data event type
type UserDataEventType string

// UserDataEventReasonType define reason type for user data event
type UserDataEventReasonType string

// ForceOrderCloseType define reason type for force order
type ForceOrderCloseType string

// Endpoints
const (
	baseApiMainUrl    = "https://fapi.binance.com"
	baseApiTestnetUrl = "https://testnet.binancefuture.com"
)

// Global enums
const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"

	PositionSideTypeBoth  PositionSideType = "BOTH"
	PositionSideTypeLong  PositionSideType = "LONG"
	PositionSideTypeShort PositionSideType = "SHORT"

	OrderTypeLimit              OrderType = "LIMIT"
	OrderTypeMarket             OrderType = "MARKET"
	OrderTypeStop               OrderType = "STOP"
	OrderTypeStopMarket         OrderType = "STOP_MARKET"
	OrderTypeTakeProfit         OrderType = "TAKE_PROFIT"
	OrderTypeTakeProfitMarket   OrderType = "TAKE_PROFIT_MARKET"
	OrderTypeTrailingStopMarket OrderType = "TRAILING_STOP_MARKET"

	TimeInForceTypeGTC TimeInForceType = "GTC" // Good Till Cancel
	TimeInForceTypeIOC TimeInForceType = "IOC" // Immediate or Cancel
	TimeInForceTypeFOK TimeInForceType = "FOK" // Fill or Kill
	TimeInForceTypeGTX TimeInForceType = "GTX" // Good Till Crossing (Post Only)

	NewOrderRespTypeACK    NewOrderRespType = "ACK"
	NewOrderRespTypeRESULT NewOrderRespType = "RESULT"

	OrderExecutionTypeNew         OrderExecutionType = "NEW"
	OrderExecutionTypePartialFill OrderExecutionType = "PARTIAL_FILL"
	OrderExecutionTypeFill        OrderExecutionType = "FILL"
	OrderExecutionTypeCanceled    OrderExecutionType = "CANCELED"
	OrderExecutionTypeCalculated  OrderExecutionType = "CALCULATED"
	OrderExecutionTypeExpired     OrderExecutionType = "EXPIRED"
	OrderExecutionTypeTrade       OrderExecutionType = "TRADE"

	OrderStatusTypeNew             OrderStatusType = "NEW"
	OrderStatusTypePartiallyFilled OrderStatusType = "PARTIALLY_FILLED"
	OrderStatusTypeFilled          OrderStatusType = "FILLED"
	OrderStatusTypeCanceled        OrderStatusType = "CANCELED"
	OrderStatusTypeRejected        OrderStatusType = "REJECTED"
	OrderStatusTypeExpired         OrderStatusType = "EXPIRED"
	OrderStatusTypeNewInsurance    OrderStatusType = "NEW_INSURANCE"
	OrderStatusTypeNewADL          OrderStatusType = "NEW_ADL"

	SymbolTypeFuture SymbolType = "FUTURE"

	WorkingTypeMarkPrice     WorkingType = "MARK_PRICE"
	WorkingTypeContractPrice WorkingType = "CONTRACT_PRICE"

	SymbolStatusTypePreTrading   SymbolStatusType = "PRE_TRADING"
	SymbolStatusTypeTrading      SymbolStatusType = "TRADING"
	SymbolStatusTypePostTrading  SymbolStatusType = "POST_TRADING"
	SymbolStatusTypeEndOfDay     SymbolStatusType = "END_OF_DAY"
	SymbolStatusTypeHalt         SymbolStatusType = "HALT"
	SymbolStatusTypeAuctionMatch SymbolStatusType = "AUCTION_MATCH"
	SymbolStatusTypeBreak        SymbolStatusType = "BREAK"

	SymbolFilterTypeLotSize          SymbolFilterType = "LOT_SIZE"
	SymbolFilterTypePrice            SymbolFilterType = "PRICE_FILTER"
	SymbolFilterTypePercentPrice     SymbolFilterType = "PERCENT_PRICE"
	SymbolFilterTypeMarketLotSize    SymbolFilterType = "MARKET_LOT_SIZE"
	SymbolFilterTypeMaxNumOrders     SymbolFilterType = "MAX_NUM_ORDERS"
	SymbolFilterTypeMaxNumAlgoOrders SymbolFilterType = "MAX_NUM_ALGO_ORDERS"

	SideEffectTypeNoSideEffect SideEffectType = "NO_SIDE_EFFECT"
	SideEffectTypeMarginBuy    SideEffectType = "MARGIN_BUY"
	SideEffectTypeAutoRepay    SideEffectType = "AUTO_REPAY"

	MarginTypeIsolated MarginType = "ISOLATED"
	MarginTypeCrossed  MarginType = "CROSSED"

	ContractTypePerpetual ContractType = "PERPETUAL"

	UserDataEventTypeListenKeyExpired    UserDataEventType = "listenKeyExpired"
	UserDataEventTypeMarginCall          UserDataEventType = "MARGIN_CALL"
	UserDataEventTypeAccountUpdate       UserDataEventType = "ACCOUNT_UPDATE"
	UserDataEventTypeOrderTradeUpdate    UserDataEventType = "ORDER_TRADE_UPDATE"
	UserDataEventTypeAccountConfigUpdate UserDataEventType = "ACCOUNT_CONFIG_UPDATE"

	UserDataEventReasonTypeDeposit             UserDataEventReasonType = "DEPOSIT"
	UserDataEventReasonTypeWithdraw            UserDataEventReasonType = "WITHDRAW"
	UserDataEventReasonTypeOrder               UserDataEventReasonType = "ORDER"
	UserDataEventReasonTypeFundingFee          UserDataEventReasonType = "FUNDING_FEE"
	UserDataEventReasonTypeWithdrawReject      UserDataEventReasonType = "WITHDRAW_REJECT"
	UserDataEventReasonTypeAdjustment          UserDataEventReasonType = "ADJUSTMENT"
	UserDataEventReasonTypeInsuranceClear      UserDataEventReasonType = "INSURANCE_CLEAR"
	UserDataEventReasonTypeAdminDeposit        UserDataEventReasonType = "ADMIN_DEPOSIT"
	UserDataEventReasonTypeAdminWithdraw       UserDataEventReasonType = "ADMIN_WITHDRAW"
	UserDataEventReasonTypeMarginTransfer      UserDataEventReasonType = "MARGIN_TRANSFER"
	UserDataEventReasonTypeMarginTypeChange    UserDataEventReasonType = "MARGIN_TYPE_CHANGE"
	UserDataEventReasonTypeAssetTransfer       UserDataEventReasonType = "ASSET_TRANSFER"
	UserDataEventReasonTypeOptionsPremiumFee   UserDataEventReasonType = "OPTIONS_PREMIUM_FEE"
	UserDataEventReasonTypeOptionsSettleProfit UserDataEventReasonType = "OPTIONS_SETTLE_PROFIT"

	ForceOrderCloseTypeLiquidation ForceOrderCloseType = "LIQUIDATION"
	ForceOrderCloseTypeADL         ForceOrderCloseType = "ADL"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func NewClient(key, secret string) *binance.Client {
	return binance.NewFuturesClient(key, secret)
}

func NewTestClient(key, secret string) *binance.Client {
	return binance.NewFuturesTestClient(key, secret)
}

func NewPingService(c *binance.Client) *PingService {
	return &PingService{C: c}
}

func NewSetServerTimeService(c *binance.Client) *SetServerTimeService {
	return &SetServerTimeService{C: c}
}

func NewServerTimeService(c *binance.Client) *ServerTimeService {
	return &ServerTimeService{C: c}
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

func NewDepthService(c *binance.Client) *DepthService {
	return &DepthService{C: c}
}

func NewKlinesService(c *binance.Client) *KlinesService {
	return &KlinesService{C: c}
}

func NewGetAccountService(c *binance.Client) *GetAccountService {
	return &GetAccountService{C: c}
}

func NewExchangeInfoService(c *binance.Client) *ExchangeInfoService {
	return &ExchangeInfoService{C: c}
}

func NewPremiumIndexService(c *binance.Client) *PremiumIndexService {
	return &PremiumIndexService{C: c}
}

func NewFundingRateService(c *binance.Client) *FundingRateService {
	return &FundingRateService{C: c}
}

func NewGetLeverageBracketService(c *binance.Client) *GetLeverageBracketService {
	return &GetLeverageBracketService{C: c}
}

func NewCreateOrderService(c *binance.Client) *CreateOrderService {
	return &CreateOrderService{C: c}
}

func NewListOpenOrdersService(c *binance.Client) *ListOpenOrdersService {
	return &ListOpenOrdersService{C: c}
}

func NewGetOrderService(c *binance.Client) *GetOrderService {
	return &GetOrderService{C: c}
}

func NewListOrdersService(c *binance.Client) *ListOrdersService {
	return &ListOrdersService{C: c}
}

func NewCancelOrderService(c *binance.Client) *CancelOrderService {
	return &CancelOrderService{C: c}
}

func NewCancelAllOpenOrdersService(c *binance.Client) *CancelAllOpenOrdersService {
	return &CancelAllOpenOrdersService{C: c}
}

func NewListUserLiquidationOrdersService(c *binance.Client) *ListUserLiquidationOrdersService {
	return &ListUserLiquidationOrdersService{C: c}
}

func NewGetPositionRiskService(c *binance.Client) *GetPositionRiskService {
	return &GetPositionRiskService{C: c}
}

func NewChangeLeverageService(c *binance.Client) *ChangeLeverageService {
	return &ChangeLeverageService{C: c}
}

func NewChangeMarginTypeService(c *binance.Client) *ChangeMarginTypeService {
	return &ChangeMarginTypeService{C: c}
}

func NewUpdatePositionMarginService(c *binance.Client) *UpdatePositionMarginService {
	return &UpdatePositionMarginService{C: c}
}

func NewChangePositionModeService(c *binance.Client) *ChangePositionModeService {
	return &ChangePositionModeService{C: c}
}

func NewGetPositionModeService(c *binance.Client) *GetPositionModeService {
	return &GetPositionModeService{C: c}
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

func NewGetPositionMarginHistoryService(c *binance.Client) *GetPositionMarginHistoryService {
	return &GetPositionMarginHistoryService{C: c}
}

func NewHistoricalTradesService(c *binance.Client) *HistoricalTradesService {
	return &HistoricalTradesService{C: c}
}

func NewRecentTradesService(c *binance.Client) *RecentTradesService {
	return &RecentTradesService{C: c}
}

func NewAggTradesService(c *binance.Client) *AggTradesService {
	return &AggTradesService{C: c}
}

func NewListAccountTradeService(c *binance.Client) *ListAccountTradeService {
	return &ListAccountTradeService{C: c}
}

func NewGetIncomeHistoryService(c *binance.Client) *GetIncomeHistoryService {
	return &GetIncomeHistoryService{C: c}
}

func NewBlvtKlinesService(c *binance.Client) *BlvtKlinesService {
	return &BlvtKlinesService{C: c}
}

func NewGetAccountV2Service(c *binance.Client) *GetAccountV2Service {
	return &GetAccountV2Service{C: c}
}

func NewGetBalanceV2Service(c *binance.Client) *GetBalanceV2Service {
	return &GetBalanceV2Service{C: c}
}
