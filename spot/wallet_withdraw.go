package spot

import (
	"github.com/firekeeper76/go-binance"
	"context"
	"encoding/json"
)

type WalletWithdrawService struct {
	C                  *binance.Client
	coin               string
	withdrawOrderID    *string
	network            *string
	address            string
	addressTag         *string
	amount             string
	transactionFeeFlag *bool
	name               *string
}

func (s *WalletWithdrawService) Coin(symbol string) *WalletWithdrawService {
	s.coin = symbol
	return s
}

// WithdrawOrderID sets the withdrawOrderID parameter.
func (s *WalletWithdrawService) WithdrawOrderID(v string) *WalletWithdrawService {
	s.withdrawOrderID = &v
	return s
}

// Network sets the network parameter.
func (s *WalletWithdrawService) Network(v string) *WalletWithdrawService {
	s.network = &v
	return s
}

// Address sets the address parameter (MANDATORY).
func (s *WalletWithdrawService) Address(v string) *WalletWithdrawService {
	s.address = v
	return s
}

// AddressTag sets the addressTag parameter.
func (s *WalletWithdrawService) AddressTag(v string) *WalletWithdrawService {
	s.addressTag = &v
	return s
}

// Amount sets the amount parameter (MANDATORY).
func (s *WalletWithdrawService) Amount(v string) *WalletWithdrawService {
	s.amount = v
	return s
}

// TransactionFeeFlag sets the transactionFeeFlag parameter.
func (s *WalletWithdrawService) TransactionFeeFlag(v bool) *WalletWithdrawService {
	s.transactionFeeFlag = &v
	return s
}

// Name sets the name parameter.
func (s *WalletWithdrawService) Name(v string) *WalletWithdrawService {
	s.name = &v
	return s
}

type WalletWithdrawResponse struct {
	Id string `json:"id"`
}

// Do send request
func (s *WalletWithdrawService) Do(ctx context.Context, opts ...binance.RequestOption) (res *WalletWithdrawResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/capital/withdraw/apply",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("coin", s.coin)
	r.SetParam("address", s.address)
	r.SetParam("amount", s.amount)
	if v := s.withdrawOrderID; v != nil {
		r.SetParam("withdrawOrderId", *v)
	}
	if v := s.network; v != nil {
		r.SetParam("network", *v)
	}
	if v := s.addressTag; v != nil {
		r.SetParam("addressTag", *v)
	}
	if v := s.transactionFeeFlag; v != nil {
		r.SetParam("transactionFeeFlag", *v)
	}
	if v := s.name; v != nil {
		r.SetParam("name", *v)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(WalletWithdrawResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type WithdrawHistoryService struct {
	C         *binance.Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int64
	limit     *int64
}

// Coin sets the asset parameter.
func (s *WithdrawHistoryService) Coin(asset string) *WithdrawHistoryService {
	s.coin = &asset
	return s
}

// Status sets the status parameter.
func (s *WithdrawHistoryService) Status(status int) *WithdrawHistoryService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *WithdrawHistoryService) StartTime(startTime int64) *WithdrawHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *WithdrawHistoryService) EndTime(endTime int64) *WithdrawHistoryService {
	s.endTime = &endTime
	return s
}
func (s *WithdrawHistoryService) Offset(offset int64) *WithdrawHistoryService {
	s.offset = &offset
	return s
}
func (s *WithdrawHistoryService) Limit(limit int64) *WithdrawHistoryService {
	s.limit = &limit
	return s
}

// Do sends the request.
func (s *WithdrawHistoryService) Do(ctx context.Context, opts ...binance.RequestOption) (withdraws []*Withdraw, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/capital/withdraw/history",
		SecType:  binance.SecTypeSigned,
	}
	if s.coin != nil {
		r.SetParam("asset", *s.coin)
	}
	if s.status != nil {
		r.SetParam("status", *s.status)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res := make([]*Withdraw, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Withdraw represents a single withdraw entry.
type Withdraw struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	TxID            string `json:"txId"`
}

type WalletGetAllService struct {
	C *binance.Client
}

type WalletGetAllResponse struct {
	Coin              string        `json:"coin"`
	DepositAllEnable  bool          `json:"depositAllEnable"`
	Free              string        `json:"free"`
	Freeze            string        `json:"freeze"`
	Ipoable           string        `json:"ipoable"`
	Ipoing            string        `json:"ipoing"`
	IsLegalMoney      bool          `json:"isLegalMoney"`
	Locked            string        `json:"locked"`
	Name              string        `json:"name"`
	NetworkList       []NetworkList `json:"networkList"`
	Storage           string        `json:"storage"`
	Trading           bool          `json:"trading"`
	WithdrawAllEnable bool          `json:"withdrawAllEnable"`
	Withdrawing       string        `json:"withdrawing"`
}
type NetworkList struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc,omitempty"`
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"`
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	WithdrawDesc            string `json:"withdrawDesc,omitempty"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawMin             string `json:"withdrawMin"`
	InsertTime              int64  `json:"insertTime,omitempty"`
	UpdateTime              int64  `json:"updateTime,omitempty"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple,omitempty"`
}

// Do send request
func (s *WalletGetAllService) Do(ctx context.Context) (res []*WalletGetAllResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/capital/config/getall",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*WalletGetAllResponse, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type WalletEnableFastService struct {
	C *binance.Client
}

// Do send request
func (s *WalletEnableFastService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/account/enableFastWithdrawSwitch",
		SecType:  binance.SecTypeSigned,
	}
	_, err = s.C.Request(ctx, r, opts...)
	return err
}

type WalletDisableFastService struct {
	C *binance.Client
}

// Do send request
func (s *WalletDisableFastService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/account/disableFastWithdrawSwitch",
		SecType:  binance.SecTypeSigned,
	}
	_, err = s.C.Request(ctx, r, opts...)
	return err
}


type WalletApiService struct {
	C *binance.Client
}

type ApiResponse struct {
	IPRestrict                     bool  `json:"ipRestrict"`
	CreateTime                     int64 `json:"createTime"`
	EnableWithdrawals              bool  `json:"enableWithdrawals"`
	EnableInternalTransfer         bool  `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool  `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool  `json:"enableVanillaOptions"`
	EnableReading                  bool  `json:"enableReading"`
	EnableFutures                  bool  `json:"enableFutures"`
	EnableMargin                   bool  `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool  `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime int64 `json:"tradingAuthorityExpirationTime"`
}

// Do send request
func (s *WalletApiService) Do(ctx context.Context, opts ...binance.RequestOption) (res *ApiResponse, err error) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/account/apiRestrictions",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(ApiResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}