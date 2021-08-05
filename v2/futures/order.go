package futures

import (
	"github.com/adshao/go-binance"
	"context"
	"encoding/json"
)

// CreateOrderService create order
type CreateOrderService struct {
	C                *binance.Client
	symbol           string
	side             SideType
	positionSide     *PositionSideType
	orderType        OrderType
	timeInForce      *TimeInForceType
	quantity         string
	reduceOnly       *bool
	price            *string
	newClientOrderID *string
	stopPrice        *string
	workingType      *WorkingType
	activationPrice  *string
	callbackRate     *string
	priceProtect     *bool
	newOrderRespType NewOrderRespType
	closePosition    *bool
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// PositionSide set side
func (s *CreateOrderService) PositionSide(positionSide PositionSideType) *CreateOrderService {
	s.positionSide = &positionSide
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

// TimeInForce set timeInForce
func (s *CreateOrderService) TimeInForce(timeInForce TimeInForceType) *CreateOrderService {
	s.timeInForce = &timeInForce
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity string) *CreateOrderService {
	s.quantity = quantity
	return s
}

// ReduceOnly set reduceOnly
func (s *CreateOrderService) ReduceOnly(reduceOnly bool) *CreateOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// NewClientOrderID set newClientOrderID
func (s *CreateOrderService) NewClientOrderID(newClientOrderID string) *CreateOrderService {
	s.newClientOrderID = &newClientOrderID
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice string) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// WorkingType set workingType
func (s *CreateOrderService) WorkingType(workingType WorkingType) *CreateOrderService {
	s.workingType = &workingType
	return s
}

// ActivationPrice set activationPrice
func (s *CreateOrderService) ActivationPrice(activationPrice string) *CreateOrderService {
	s.activationPrice = &activationPrice
	return s
}

// CallbackRate set callbackRate
func (s *CreateOrderService) CallbackRate(callbackRate string) *CreateOrderService {
	s.callbackRate = &callbackRate
	return s
}

// PriceProtect set priceProtect
func (s *CreateOrderService) PriceProtect(priceProtect bool) *CreateOrderService {
	s.priceProtect = &priceProtect
	return s
}

// NewOrderResponseType set newOrderResponseType
func (s *CreateOrderService) NewOrderResponseType(newOrderResponseType NewOrderRespType) *CreateOrderService {
	s.newOrderRespType = newOrderResponseType
	return s
}

// ClosePosition set closePosition
func (s *CreateOrderService) ClosePosition(closePosition bool) *CreateOrderService {
	s.closePosition = &closePosition
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, Endpoint string, opts ...binance.RequestOption) (data []byte, err *binance.APIError) {
	r := &binance.Request{
		Method:   "POST",
		Endpoint: Endpoint,
		SecType:  binance.SecTypeSigned,
	}
	m := binance.Params{
		"symbol":           s.symbol,
		"side":             s.side,
		"type":             s.orderType,
		"quantity":         s.quantity,
		"newOrderRespType": s.newOrderRespType,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}
	if s.timeInForce != nil {
		m["timeInForce"] = *s.timeInForce
	}
	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}
	if s.price != nil {
		m["price"] = *s.price
	}
	if s.newClientOrderID != nil {
		m["newClientOrderId"] = *s.newClientOrderID
	}
	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}
	if s.workingType != nil {
		m["workingType"] = *s.workingType
	}
	if s.priceProtect != nil {
		m["priceProtect"] = *s.priceProtect
	}
	if s.activationPrice != nil {
		m["activationPrice"] = *s.activationPrice
	}
	if s.callbackRate != nil {
		m["callbackRate"] = *s.callbackRate
	}
	if s.closePosition != nil {
		m["closePosition"] = *s.closePosition
	}
	r.SetFormParams(m)
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	return data, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...binance.RequestOption) (res *CreateOrderResponse, err *binance.APIError) {
	data, rErr := s.createOrder(ctx, "/fapi/v1/order", opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(CreateOrderResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// CreateOrderResponse define create order response
type CreateOrderResponse struct {
	Symbol           string           `json:"symbol"`
	OrderID          int64            `json:"orderId"`
	ClientOrderID    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuote         string           `json:"cumQuote"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	Side             SideType         `json:"side"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	AvgPrice         string           `json:"avgPrice"`
	PositionSide     PositionSideType `json:"positionSide"`
	ClosePosition    bool             `json:"closePosition"`
	PriceProtect     bool             `json:"priceProtect"`
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*Order, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/openOrders",
		SecType:  binance.SecTypeSigned,
	}
	if s.symbol != "" {
		r.SetParam("symbol", s.symbol)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*Order, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// GetOrderService get an order
type GetOrderService struct {
	C                 *binance.Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *GetOrderService) Symbol(symbol string) *GetOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *GetOrderService) OrigClientOrderID(origClientOrderID string) *GetOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...binance.RequestOption) (res *Order, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/order",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.SetParam("origClientOrderId", *s.origClientOrderID)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(Order)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Order define order info
type Order struct {
	Symbol           string           `json:"symbol"`
	OrderID          int64            `json:"orderId"`
	ClientOrderID    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	Status           OrderStatusType  `json:"status"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	Side             SideType         `json:"side"`
	StopPrice        string           `json:"stopPrice"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	AvgPrice         string           `json:"avgPrice"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	C         *binance.Client
	symbol    string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListOrdersService) Symbol(symbol string) *ListOrdersService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *ListOrdersService) OrderID(orderID int64) *ListOrdersService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListOrdersService) StartTime(startTime int64) *ListOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListOrdersService) EndTime(endTime int64) *ListOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*Order, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/allOrders",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*Order, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	C                 *binance.Client
	symbol            string
	orderID           *int64
	origClientOrderID *string
}

// Symbol set symbol
func (s *CancelOrderService) Symbol(symbol string) *CancelOrderService {
	s.symbol = symbol
	return s
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = &orderID
	return s
}

// OrigClientOrderID set origClientOrderID
func (s *CancelOrderService) OrigClientOrderID(origClientOrderID string) *CancelOrderService {
	s.origClientOrderID = &origClientOrderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...binance.RequestOption) (res *CancelOrderResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "DELETE",
		Endpoint: "/fapi/v1/order",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParam("symbol", s.symbol)
	if s.orderID != nil {
		r.SetFormParam("orderId", *s.orderID)
	}
	if s.origClientOrderID != nil {
		r.SetFormParam("origClientOrderId", *s.origClientOrderID)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(CancelOrderResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// CancelOrderResponse define response of canceling order
type CancelOrderResponse struct {
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	Symbol           string           `json:"symbol"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

// CancelAllOpenOrdersService cancel all open orders
type CancelAllOpenOrdersService struct {
	C      *binance.Client
	symbol string
}

// Symbol set symbol
func (s *CancelAllOpenOrdersService) Symbol(symbol string) *CancelAllOpenOrdersService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CancelAllOpenOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (err *binance.APIError) {
	r := &binance.Request{
		Method:   "DELETE",
		Endpoint: "/fapi/v1/allOpenOrders",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParam("symbol", s.symbol)
	_, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return rErr
	}
	return nil
}


// LiquidationOrder define liquidation order
type LiquidationOrder struct {
	Symbol           string          `json:"symbol"`
	Price            string          `json:"price"`
	OrigQuantity     string          `json:"origQty"`
	ExecutedQuantity string          `json:"executedQty"`
	AveragePrice     string          `json:"avragePrice"`
	Status           OrderStatusType `json:"status"`
	TimeInForce      TimeInForceType `json:"timeInForce"`
	Type             OrderType       `json:"type"`
	Side             SideType        `json:"side"`
	Time             int64           `json:"time"`
}

// ListUserLiquidationOrdersService lists user's liquidation orders
type ListUserLiquidationOrdersService struct {
	C             *binance.Client
	symbol        *string
	autoCloseType ForceOrderCloseType
	startTime     *int64
	endTime       *int64
	limit         *int
}

// Symbol set symbol
func (s *ListUserLiquidationOrdersService) Symbol(symbol string) *ListUserLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// AutoCloseType set symbol
func (s *ListUserLiquidationOrdersService) AutoCloseType(autoCloseType ForceOrderCloseType) *ListUserLiquidationOrdersService {
	s.autoCloseType = autoCloseType
	return s
}

// StartTime set startTime
func (s *ListUserLiquidationOrdersService) StartTime(startTime int64) *ListUserLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListUserLiquidationOrdersService) EndTime(endTime int64) *ListUserLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListUserLiquidationOrdersService) Limit(limit int) *ListUserLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListUserLiquidationOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*UserLiquidationOrder, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/forceOrders",
		SecType:  binance.SecTypeSigned,
	}

	r.SetParam("autoCloseType", s.autoCloseType)
	if s.symbol != nil {
		r.SetParam("symbol", *s.symbol)
	}
	if s.startTime != nil {
		r.SetParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.SetParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.SetParam("limit", *s.limit)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*UserLiquidationOrder, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// UserLiquidationOrder defines user's liquidation order
type UserLiquidationOrder struct {
	OrderId          int64            `json:"orderId"`
	Symbol           string           `json:"symbol"`
	Status           OrderStatusType  `json:"status"`
	ClientOrderId    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	AveragePrice     string           `json:"avgPrice"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuote         string           `json:"cumQuote"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ReduceOnly       bool             `json:"reduceOnly"`
	ClosePosition    bool             `json:"closePosition"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	StopPrice        string           `json:"stopPrice"`
	WorkingType      WorkingType      `json:"workingType"`
	OrigType         string           `json:"origType"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
}