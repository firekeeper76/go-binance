package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
)

type TransHistoryService struct {
	C         *binance.Client
	startTime *int64
	endTime   *int64
}

func (s *TransHistoryService) StartTime(startTime int64) *TransHistoryService {
	s.startTime = &startTime
	return s
}

func (s *TransHistoryService) EndTime(endTime int64) *TransHistoryService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *TransHistoryService) Do(ctx context.Context, opts ...binance.RequestOption) (res *TransHistoryResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/asset/dribblet",
		SecType:  binance.SecTypeSigned,
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	data, rErr := s.C.Request(ctx, r)
	if rErr != nil {
		return nil, rErr
	}
	res = new(TransHistoryResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

type TransHistoryResponse struct {
	Total              int                  `json:"total"`
	UserAssetDribblets []UserAssetDribblets `json:"userAssetDribblets"`
}
type UserAssetDribblets struct {
	TotalTransferedAmount    string                     `json:"totalTransferedAmount"`
	TotalServiceChargeAmount string                     `json:"totalServiceChargeAmount"`
	TransID                  int64                      `json:"transId"`
	UserAssetDribbletDetails []UserAssetDribbletDetails `json:"userAssetDribbletDetails"`
}
type UserAssetDribbletDetails struct {
	TransID             int    `json:"transId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	Amount              string `json:"amount"`
	OperateTime         int64  `json:"operateTime"`
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}

type TransferService struct {
	C     *binance.Client
	asset []string
}

// Asset set asset.
func (s *TransferService) Asset(asset []string) *TransferService {
	s.asset = asset
	return s
}

// Do sends the request.
func (s *TransferService) Do(ctx context.Context, opts ...binance.RequestOption) (res *DustTransferResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/sapi/v1/asset/dust",
		SecType:  binance.SecTypeSigned,
	}
	for _, a := range s.asset {
		r.AddParam("asset", a)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(DustTransferResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// DustTransferResponse represents the response from DustTransferService.
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
}
