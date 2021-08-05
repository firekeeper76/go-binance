package delivery

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
)

// GetPositionRiskService get account balance
type GetPositionRiskService struct {
	C           *binance.Client
	pair        *string
	marginAsset *string
}

// MarginAsset set margin asset
func (s *GetPositionRiskService) MarginAsset(marginAsset string) *GetPositionRiskService {
	s.marginAsset = &marginAsset
	return s
}

// Pair set pair
func (s *GetPositionRiskService) Pair(pair string) *GetPositionRiskService {
	s.pair = &pair
	return s
}

// Do send request
func (s *GetPositionRiskService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*PositionRisk, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/positionRisk",
		SecType:  binance.SecTypeSigned,
	}
	if s.marginAsset != nil {
		r.SetParam("marginAsset", *s.marginAsset)
	}
	if s.pair != nil {
		r.SetParam("pair", *s.pair)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*PositionRisk, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// PositionRisk define position risk info
type PositionRisk struct {
	Symbol           string `json:"symbol"`
	PositionAmt      string `json:"positionAmt"`
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	LiquidationPrice string `json:"liquidationPrice"`
	Leverage         string `json:"leverage"`
	MaxQuantity      string `json:"maxQty"`
	MarginType       string `json:"marginType"`
	IsolatedMargin   string `json:"isolatedMargin"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	PositionSide     string `json:"positionSide"`
}
