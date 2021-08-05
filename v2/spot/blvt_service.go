package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
)

type BLVTSubscribeService struct {
	C         *binance.Client
	tokenName string
	cost      int
}

type BLVTSubscribeResponse struct {
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	TokenName string `json:"tokenName"`
	Amount    string `json:"amount"`
	Cost      string `json:"cost"`
	Timstamp  int    `json:"timstamp"`
}

func (s *BLVTSubscribeService) Symbol(symbol string) *BLVTSubscribeService {
	s.tokenName = symbol
	return s
}
func (s *BLVTSubscribeService) Cost(cost int) *BLVTSubscribeService {
	s.cost = cost
	return s
}

// Do send request
func (s *BLVTSubscribeService) Do(ctx context.Context, opts ...binance.RequestOption) (res *BLVTSubscribeResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/blvt/subscribe",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("tokenName", s.tokenName)
	r.SetParam("cost", s.cost)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(BLVTSubscribeResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type BLVTRedeemService struct {
	C         *binance.Client
	tokenName string
	amount    decimal.Decimal
}

type RedeemResponse struct {
	Id           int64  `json:"id"`
	Status       string `json:"status"`
	TokenName    string `json:"tokenName"`
	RedeemAmount string `json:"redeemAmount"`
	Amount       string `json:"amount"`
	Timstamp     int    `json:"timstamp"`
}

func (s *BLVTRedeemService) TokenName(symbol string) *BLVTRedeemService {
	s.tokenName = symbol
	return s
}
func (s *BLVTRedeemService) Amount(amount decimal.Decimal) *BLVTRedeemService {
	s.amount = amount
	return s
}

// Do send request
func (s *BLVTRedeemService) Do(ctx context.Context, opts ...binance.RequestOption) (res *RedeemResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/blvt/redeem",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("tokenName", s.tokenName)
	r.SetParam("amount", s.amount)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(RedeemResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type BLVTLimitService struct {
	C         *binance.Client
	tokenName string
}

type BLVTLimitResponse struct {
	TokenName                   string `json:"tokenName"`
	UserDailyTotalPurchaseLimit string `json:"userDailyTotalPurchaseLimit"`
	UserDailyTotalRedeemLimit   string `json:"userDailyTotalRedeemLimit"`
}

func (s *BLVTLimitService) TokenName(symbol string) *BLVTLimitService {
	s.tokenName = symbol
	return s
}

// Do send request
func (s *BLVTLimitService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*BLVTLimitResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/blvt/userLimit",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("tokenName", s.tokenName)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*BLVTLimitResponse, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type BLVTRedeemRecordService struct {
	C         *binance.Client
	tokenName string
	id        *int64
	limit     *int
	startTime *int64
	endTime   *int64
}

func (s *BLVTRedeemRecordService) TokenName(symbol string) *BLVTRedeemRecordService {
	s.tokenName = symbol
	return s
}

func (s *BLVTRedeemRecordService) Id(id int64) *BLVTRedeemRecordService {
	s.id = &id
	return s
}

// Limit set limit
func (s *BLVTRedeemRecordService) Limit(limit int) *BLVTRedeemRecordService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BLVTRedeemRecordService) StartTime(startTime int64) *BLVTRedeemRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BLVTRedeemRecordService) EndTime(endTime int64) *BLVTRedeemRecordService {
	s.endTime = &endTime
	return s
}

type RedeemRecordResponse struct {
	ID         int    `json:"id"`
	TokenName  string `json:"tokenName"`
	Amount     string `json:"amount"`
	Nav        string `json:"nav"`
	Fee        string `json:"fee"`
	NetProceed string `json:"netProceed"`
	Timstamp   int64  `json:"timstamp"`
}

// Do send request
func (s *BLVTRedeemRecordService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*RedeemRecordResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/blvt/redeem/record",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("tokenName", s.tokenName)
	if s.id != nil {
		r.SetParam("id", *s.id)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
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
	res = make([]*RedeemRecordResponse, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type BLVTSubscribeRecordService struct {
	C         *binance.Client
	tokenName string
	id        *int64
	limit     *int
	startTime *int64
	endTime   *int64
}

func (s *BLVTSubscribeRecordService) TokenName(symbol string) *BLVTSubscribeRecordService {
	s.tokenName = symbol
	return s
}

func (s *BLVTSubscribeRecordService) Id(id int64) *BLVTSubscribeRecordService {
	s.id = &id
	return s
}

// Limit set limit
func (s *BLVTSubscribeRecordService) Limit(limit int) *BLVTSubscribeRecordService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *BLVTSubscribeRecordService) StartTime(startTime int64) *BLVTSubscribeRecordService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *BLVTSubscribeRecordService) EndTime(endTime int64) *BLVTSubscribeRecordService {
	s.endTime = &endTime
	return s
}

type SubscribeRecordResponse struct {
	ID          int    `json:"id"`
	TokenName   string `json:"tokenName"`
	Amount      string `json:"amount"`
	Nav         string `json:"nav"`
	Fee         string `json:"fee"`
	TotalCharge string `json:"totalCharge"`
	Timstamp    int64  `json:"timstamp"`
}

// Do send request
func (s *BLVTSubscribeRecordService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*SubscribeRecordResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/blvt/subscribe/record",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("tokenName", s.tokenName)
	if s.id != nil {
		r.SetParam("id", *s.id)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
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
	res = make([]*SubscribeRecordResponse, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type BLVTInfoService struct {
	C         *binance.Client
	tokenName *string
}

func (s *BLVTInfoService) TokenName(symbol string) *BLVTInfoService {
	s.tokenName = &symbol
	return s
}

type BLVTInfoResponse struct {
	TokenName          string           `json:"tokenName"`
	Description        string           `json:"description"`
	Underlying         string           `json:"underlying"`
	TokenIssued        string           `json:"tokenIssued"`
	Basket             string           `json:"basket"`
	CurrentBaskets     []CurrentBaskets `json:"currentBaskets"`
	Nav                string           `json:"nav"`
	RealLeverage       string           `json:"realLeverage"`
	FundingRate        string           `json:"fundingRate"`
	DailyManagementFee string           `json:"dailyManagementFee"`
	PurchaseFeePct     string           `json:"purchaseFeePct"`
	DailyPurchaseLimit string           `json:"dailyPurchaseLimit"`
	RedeemFeePct       string           `json:"redeemFeePct"`
	DailyRedeemLimit   string           `json:"dailyRedeemLimit"`
	Timstamp           int64            `json:"timstamp"`
}
type CurrentBaskets struct {
	Symbol        string `json:"symbol"`
	Amount        string `json:"amount"`
	NotionalValue string `json:"notionalValue"`
}

// Do send request
func (s *BLVTInfoService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*BLVTInfoResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/blvt/tokenInfo",
		SecType:  binance.SecTypeSigned,
	}
	if s.tokenName != nil {
		r.SetParam("tokenName", *s.tokenName)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*BLVTInfoResponse, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}
