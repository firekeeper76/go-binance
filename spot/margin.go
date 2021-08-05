package spot

import (
	"github.com/adshao/go-binance/v2"
	"context"
	"encoding/json"
	"strings"
)

// MarginTransferService transfer between spot account and margin account
type MarginTransferService struct {
	C            *binance.Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginTransferService) Asset(asset string) *MarginTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginTransferService) Amount(amount string) *MarginTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from main account to margin account 2: transfer from margin account to main account
func (s *MarginTransferService) Type(transferType MarginTransferType) *MarginTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *MarginTransferService) Do(ctx context.Context, opts ...binance.RequestOption) (res *TransactionResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/margin/transfer",
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.SetFormParams(m)
	res = new(TransactionResponse)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// TransactionResponse define transaction response
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

// MarginLoanService apply for a loan
type MarginLoanService struct {
	C      *binance.Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...binance.RequestOption) (res *TransactionResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/margin/loan",
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.SetFormParams(m)
	res = new(TransactionResponse)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginRepayService repay loan for margin account
type MarginRepayService struct {
	C      *binance.Client
	asset  string
	amount string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...binance.RequestOption) (res *TransactionResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/margin/repay",
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.SetFormParams(m)
	res = new(TransactionResponse)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// ListMarginLoansService list loan record
type ListMarginLoansService struct {
	C         *binance.Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginLoansService) Asset(asset string) *ListMarginLoansService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginLoansService) TxID(txID int64) *ListMarginLoansService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginLoansService) StartTime(startTime int64) *ListMarginLoansService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginLoansService) EndTime(endTime int64) *ListMarginLoansService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginLoansService) Current(current int64) *ListMarginLoansService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginLoansService) Size(size int64) *ListMarginLoansService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginLoansService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginLoanResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/loan",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("asset", s.asset)
	if s.txID != nil {
		r.SetParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginLoanResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginLoanResponse define margin loan response
type MarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginLoan define margin loan
type MarginLoan struct {
	Asset     string               `json:"asset"`
	Principal string               `json:"principal"`
	Timestamp int64                `json:"timestamp"`
	Status    MarginLoanStatusType `json:"status"`
}

// ListMarginRepaysService list repay record
type ListMarginRepaysService struct {
	C         *binance.Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginRepaysService) Asset(asset string) *ListMarginRepaysService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginRepaysService) TxID(txID int64) *ListMarginRepaysService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginRepaysService) StartTime(startTime int64) *ListMarginRepaysService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginRepaysService) EndTime(endTime int64) *ListMarginRepaysService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginRepaysService) Current(current int64) *ListMarginRepaysService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginRepaysService) Size(size int64) *ListMarginRepaysService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginRepaysService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginRepayResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/repay",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("asset", s.asset)
	if s.txID != nil {
		r.SetParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.SetParam("current", *s.current)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginRepayResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginRepayResponse define margin repay response
type MarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarginRepay define margin repay
type MarginRepay struct {
	Asset     string                `json:"asset"`
	Amount    string                `json:"amount"`
	Interest  string                `json:"interest"`
	Principal string                `json:"principal"`
	Timestamp int64                 `json:"timestamp"`
	Status    MarginRepayStatusType `json:"status"`
	TxID      int64                 `json:"txId"`
}

// GetIsolatedMarginAccountService gets isolated margin account info
type GetIsolatedMarginAccountService struct {
	C *binance.Client

	symbols []string
}

// Symbols set symbols to the isolated margin account
func (s *GetIsolatedMarginAccountService) Symbols(symbols ...string) *GetIsolatedMarginAccountService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *GetIsolatedMarginAccountService) Do(ctx context.Context, opts ...binance.RequestOption) (res *IsolatedMarginAccount, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/isolated/account",
		SecType:  binance.SecTypeSigned,
	}
	if len(s.symbols) > 0 {
		r.SetParam("symbols", strings.Join(s.symbols, ","))
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(IsolatedMarginAccount)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// IsolatedMarginAccount defines isolated user assets of margin account
type IsolatedMarginAccount struct {
	TotalAssetOfBTC     string                `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string                `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string                `json:"totalNetAssetOfBtc"`
	Assets              []IsolatedMarginAsset `json:"assets"`
}

// IsolatedMarginAsset defines isolated margin asset information, like margin level, liquidation price... etc
type IsolatedMarginAsset struct {
	Symbol     string            `json:"symbol"`
	QuoteAsset IsolatedUserAsset `json:"quoteAsset"`
	BaseAsset  IsolatedUserAsset `json:"baseAsset"`

	IsolatedCreated   bool   `json:"isolatedCreated"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// IsolatedUserAsset defines isolated user assets of the margin account
type IsolatedUserAsset struct {
	Asset         string `json:"asset"`
	Borrowed      string `json:"borrowed"`
	Free          string `json:"free"`
	Interest      string `json:"interest"`
	Locked        string `json:"locked"`
	NetAsset      string `json:"netAsset"`
	NetAssetOfBtc string `json:"netAssetOfBtc"`

	BorrowEnabled bool   `json:"borrowEnabled"`
	RepayEnabled  bool   `json:"repayEnabled"`
	TotalAsset    string `json:"totalAsset"`
}

// GetMarginAccountService get margin account info
type GetMarginAccountService struct {
	C *binance.Client
}

// Do send request
func (s *GetMarginAccountService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginAccount, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/account",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginAccount)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginAccount define margin account info
type MarginAccount struct {
	BorrowEnabled       bool        `json:"borrowEnabled"`
	MarginLevel         string      `json:"marginLevel"`
	TotalAssetOfBTC     string      `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string      `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string      `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool        `json:"tradeEnabled"`
	TransferEnabled     bool        `json:"transferEnabled"`
	UserAssets          []UserAsset `json:"userAssets"`
}

// UserAsset define user assets of margin account
type UserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// GetMarginAssetService get margin asset info
type GetMarginAssetService struct {
	C     *binance.Client
	asset string
}

// Asset set asset
func (s *GetMarginAssetService) Asset(asset string) *GetMarginAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginAssetService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginAsset, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/asset",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetParam("asset", s.asset)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginAsset)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginAsset define margin asset info
type MarginAsset struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// GetMarginPairService get margin pair info
type GetMarginPairService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPairService) Symbol(symbol string) *GetMarginPairService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPairService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginPair, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/pair",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetParam("symbol", s.symbol)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginPair)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginPair define margin pair info
type MarginPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginAllPairsService get margin pair info
type GetMarginAllPairsService struct {
	C *binance.Client
}

// Do send request
func (s *GetMarginAllPairsService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*MarginAllPair, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/allPairs",
		SecType:  binance.SecTypeAPIKey,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*MarginAllPair, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginAllPair define margin pair info
type MarginAllPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// GetMarginPriceIndexService get margin price index
type GetMarginPriceIndexService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPriceIndexService) Symbol(symbol string) *GetMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPriceIndexService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MarginPriceIndex, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/priceIndex",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetParam("symbol", s.symbol)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MarginPriceIndex)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MarginPriceIndex define margin price index
type MarginPriceIndex struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// ListMarginTradesService list trades
type ListMarginTradesService struct {
	C          *binance.Client
	symbol     string
	startTime  *int64
	endTime    *int64
	limit      *int
	fromID     *int64
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginTradesService) Symbol(symbol string) *ListMarginTradesService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginTradesService) IsIsolated(isIsolated bool) *ListMarginTradesService {
	s.isIsolated = isIsolated
	return s
}

// StartTime set starttime
func (s *ListMarginTradesService) StartTime(startTime int64) *ListMarginTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginTradesService) EndTime(endTime int64) *ListMarginTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginTradesService) Limit(limit int) *ListMarginTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListMarginTradesService) FromID(fromID int64) *ListMarginTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *ListMarginTradesService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*TradeV3, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/myTrades",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.SetParam("fromId", *s.fromID)
	}
	if s.isIsolated {
		r.SetParam("isIsolated", "TRUE")
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*TradeV3, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// GetMaxBorrowableService get max borrowable of asset
type GetMaxBorrowableService struct {
	C     *binance.Client
	asset string
}

// Asset set asset
func (s *GetMaxBorrowableService) Asset(asset string) *GetMaxBorrowableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxBorrowableService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MaxBorrowable, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/maxBorrowable",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("asset", s.asset)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MaxBorrowable)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MaxBorrowable define max borrowable response
type MaxBorrowable struct {
	Amount string `json:"amount"`
}

// GetMaxTransferableService get max transferable of asset
type GetMaxTransferableService struct {
	C     *binance.Client
	asset string
}

// Asset set asset
func (s *GetMaxTransferableService) Asset(asset string) *GetMaxTransferableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxTransferableService) Do(ctx context.Context, opts ...binance.RequestOption) (res *MaxTransferable, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/margin/maxTransferable",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("asset", s.asset)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(MaxTransferable)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// MaxTransferable define max transferable response
type MaxTransferable struct {
	Amount string `json:"amount"`
}

// StartIsolatedMarginUserStreamService create listen key for margin user stream service
type StartIsolatedMarginUserStreamService struct {
	C      *binance.Client
	symbol string
}

// Symbol sets the user stream to isolated margin user stream
func (s *StartIsolatedMarginUserStreamService) Symbol(symbol string) *StartIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *StartIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (listenKey string, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/userDataStream/isolated",
		SecType:  binance.SecTypeAPIKey,
	}

	r.SetFormParam("symbol", s.symbol)

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

// KeepaliveIsolatedMarginUserStreamService updates listen key for isolated margin user data stream
type KeepaliveIsolatedMarginUserStreamService struct {
	C         *binance.Client
	listenKey string
	symbol    string
}

// Symbol set symbol to the isolated margin keepalive request
func (s *KeepaliveIsolatedMarginUserStreamService) Symbol(symbol string) *KeepaliveIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// ListenKey set listen key
func (s *KeepaliveIsolatedMarginUserStreamService) ListenKey(listenKey string) *KeepaliveIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "PUT",
		Endpoint: "/sapi/v1/userDataStream/isolated",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetFormParam("listenKey", s.listenKey)
	r.SetFormParam("symbol", s.symbol)
	_, rErr := s.C.Request(ctx, r, opts...)
	return rErr
}

// CloseIsolatedMarginUserStreamService delete listen key
type CloseIsolatedMarginUserStreamService struct {
	C         *binance.Client
	listenKey string

	symbol string
}

// ListenKey set listen key
func (s *CloseIsolatedMarginUserStreamService) ListenKey(listenKey string) *CloseIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Symbol set symbol to the isolated margin user stream close request
func (s *CloseIsolatedMarginUserStreamService) Symbol(symbol string) *CloseIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CloseIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "DELETE",
		Endpoint: "/sapi/v1/userDataStream/isolated",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetFormParam("listenKey", s.listenKey)
	r.SetFormParam("symbol", s.symbol)
	_, rErr := s.C.Request(ctx, r, opts...)
	return rErr
}

// StartMarginUserStreamService create listen key for margin user stream service
type StartMarginUserStreamService struct {
	C *binance.Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (listenKey string, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/userDataStream",
		SecType:  binance.SecTypeAPIKey,
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

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	C         *binance.Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "PUT",
		Endpoint: "/sapi/v1/userDataStream",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetFormParam("listenKey", s.listenKey)
	_, rErr := s.C.Request(ctx, r, opts...)
	return rErr
}

// CloseMarginUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	C         *binance.Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "DELETE",
		Endpoint: "/sapi/v1/userDataStream",
		SecType:  binance.SecTypeAPIKey,
	}
	r.SetFormParam("listenKey", s.listenKey)
	_, rErr := s.C.Request(ctx, r, opts...)
	return rErr
}
