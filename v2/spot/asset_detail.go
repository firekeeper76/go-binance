package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
	"github.com/shopspring/decimal"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type AssetDetailService struct {
	C *binance.Client
}

// Do sends the request.
func (s *AssetDetailService) Do(ctx context.Context, opt ...binance.RequestOption) (assetDetails *AssetDetail, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/asset/assetDetail",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opt...)
	if rErr != nil {
		return
	}
	res := new(AssetDetail)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// AssetDetail represents the detail of an asset
type AssetDetail struct {
	CTR CTR `json:"CTR"`
	SKY SKY `json:"SKY"`
}
type CTR struct {
	MinWithdrawAmount string          `json:"minWithdrawAmount"`
	DepositStatus     bool            `json:"depositStatus"`
	WithdrawFee       decimal.Decimal `json:"withdrawFee"`
	WithdrawStatus    bool            `json:"withdrawStatus"`
	DepositTip        string          `json:"depositTip"`
}
type SKY struct {
	MinWithdrawAmount string          `json:"minWithdrawAmount"`
	DepositStatus     bool            `json:"depositStatus"`
	WithdrawFee       decimal.Decimal `json:"withdrawFee"`
	WithdrawStatus    bool            `json:"withdrawStatus"`
}
