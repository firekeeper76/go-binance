package futures

import (
	"github.com/adshao/go-binance/v2"
	"context"
	"encoding/json"
)

// ChangeLeverageService change user's initial leverage of specific symbol market
type ChangeLeverageService struct {
	C        *binance.Client
	symbol   string
	leverage int
}

// Symbol set symbol
func (s *ChangeLeverageService) Symbol(symbol string) *ChangeLeverageService {
	s.symbol = symbol
	return s
}

// Leverage set leverage
func (s *ChangeLeverageService) Leverage(leverage int) *ChangeLeverageService {
	s.leverage = leverage
	return s
}

// Do send request
func (s *ChangeLeverageService) Do(ctx context.Context, opts ...binance.RequestOption) (res *SymbolLeverage, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/fapi/v1/leverage",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParams(binance.Params{
		"symbol":   s.symbol,
		"leverage": s.leverage,
	})
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(SymbolLeverage)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// SymbolLeverage define leverage info of symbol
type SymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}

// ChangeMarginTypeService change user's margin type of specific symbol market
type ChangeMarginTypeService struct {
	C          *binance.Client
	symbol     string
	marginType MarginType
}

// Symbol set symbol
func (s *ChangeMarginTypeService) Symbol(symbol string) *ChangeMarginTypeService {
	s.symbol = symbol
	return s
}

// MarginType set margin type
func (s *ChangeMarginTypeService) MarginType(marginType MarginType) *ChangeMarginTypeService {
	s.marginType = marginType
	return s
}

// Do send request
func (s *ChangeMarginTypeService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/fapi/v1/marginType",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParams(binance.Params{
		"symbol":     s.symbol,
		"marginType": s.marginType,
	})
	_, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return rErr
	}
	return nil
}

// UpdatePositionMarginService update isolated position margin
type UpdatePositionMarginService struct {
	C            *binance.Client
	symbol       string
	positionSide *PositionSideType
	amount       string
	actionType   int
}

// Symbol set symbol
func (s *UpdatePositionMarginService) Symbol(symbol string) *UpdatePositionMarginService {
	s.symbol = symbol
	return s
}

// PositionSide Side set side
func (s *UpdatePositionMarginService) PositionSide(positionSide PositionSideType) *UpdatePositionMarginService {
	s.positionSide = &positionSide
	return s
}

// Amount set position margin amount
func (s *UpdatePositionMarginService) Amount(amount string) *UpdatePositionMarginService {
	s.amount = amount
	return s
}

// Type set action type: 1: Add postion margin，2: Reduce postion margin
func (s *UpdatePositionMarginService) Type(actionType int) *UpdatePositionMarginService {
	s.actionType = actionType
	return s
}

// Do send request
func (s *UpdatePositionMarginService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/fapi/v1/positionMargin",
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"symbol": s.symbol,
		"amount": s.amount,
		"type":   s.actionType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	r.SetFormParams(m)

	_, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return rErr
	}
	return nil
}

// ChangePositionModeService change user's position mode
type ChangePositionModeService struct {
	C        *binance.Client
	dualSide string
}

// DualSide Change user's position mode: true - Hedge Mode, false - One-way Mode
func (s *ChangePositionModeService) DualSide(dualSide bool) *ChangePositionModeService {
	if dualSide {
		s.dualSide = "true"
	} else {
		s.dualSide = "false"
	}
	return s
}

// Do send request
func (s *ChangePositionModeService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: "/fapi/v1/positionSide/dual",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParams(binance.Params{
		"dualSidePosition": s.dualSide,
	})
	_, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return rErr
	}
	return nil
}

// GetPositionModeService get user's position mode
type GetPositionModeService struct {
	C *binance.Client
}

// PositionMode Response of user's position mode
type PositionMode struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

// Do send request
func (s *GetPositionModeService) Do(ctx context.Context, opts ...binance.RequestOption) (res *PositionMode, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/positionSide/dual",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParams(binance.Params{})
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(PositionMode)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}
