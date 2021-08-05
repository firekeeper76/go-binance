package delivery

import (
	"github.com/adshao/go-binance/v2"
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
	closePosition    *bool
	activationPrice  *string
	callbackRate     *string
	workingType      *WorkingType
	priceProtect     *bool
	newOrderRespType NewOrderRespType
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
	data, err := s.createOrder(ctx, "/dapi/v1/order", opts...)
	if err != nil {
		return nil, err
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
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	AvgPrice         string           `json:"avgPrice"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	ClosePosition    bool             `json:"closePosition"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	OrigType         OrderType        `json:"origType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	PriceProtect     bool             `json:"priceProtect"`
}

// ListOpenOrdersService list opened orders
type ListOpenOrdersService struct {
	C      *binance.Client
	symbol string
	pair   string
}

// Symbol set symbol
func (s *ListOpenOrdersService) Symbol(symbol string) *ListOpenOrdersService {
	s.symbol = symbol
	return s
}

// Pair set pair
func (s *ListOpenOrdersService) Pair(pair string) *ListOpenOrdersService {
	s.pair = pair
	return s
}

// Do send request
func (s *ListOpenOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*Order, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/openOrders",
		SecType:  binance.SecTypeSigned,
	}
	if s.symbol != "" {
		r.SetParam("symbol", s.symbol)
	}
	if s.pair != "" {
		r.SetParam("pair", s.symbol)
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
		Endpoint: "/dapi/v1/order",
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
	AvgPrice         string           `json:"avgPrice"`
	ClientOrderID    string           `json:"clientOrderId"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	OrigType         OrderType        `json:"origType"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	ClosePosition    bool             `json:"closePosition"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	Time             int64            `json:"time"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	PriceProtect     bool             `json:"priceProtect"`
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	C         *binance.Client
	symbol    string
	pair      string
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

// Pair set pair
func (s *ListOrdersService) Pair(pair string) *ListOrdersService {
	s.pair = pair
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
		Endpoint: "/dapi/v1/allOrders",
		SecType:  binance.SecTypeSigned,
	}
	if s.symbol != "" {
		r.SetParam("symbol", s.symbol)
	}
	if s.pair != "" {
		r.SetParam("pair", s.pair)
	}
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
		Endpoint: "/dapi/v1/order",
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
	AvgPrice         string           `json:"avgPrice"`
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumBase          string           `json:"cumBase"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	OrigType         OrderType        `json:"origType"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	ClosePosition    bool             `json:"closePosition"`
	Symbol           string           `json:"symbol"`
	Pair             string           `json:"pair"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
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
		Endpoint: "/dapi/v1/allOpenOrders",
		SecType:  binance.SecTypeSigned,
	}
	r.SetFormParam("symbol", s.symbol)
	_, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return  rErr
	}
	return nil
}

// ListLiquidationOrdersService list liquidation orders
type ListLiquidationOrdersService struct {
	C         *binance.Client
	symbol    *string
	pair      *string
	startTime *int64
	endTime   *int64
	limit     *int
}

// Symbol set symbol
func (s *ListLiquidationOrdersService) Symbol(symbol string) *ListLiquidationOrdersService {
	s.symbol = &symbol
	return s
}

// Pair set pair
func (s *ListLiquidationOrdersService) Pair(pair string) *ListLiquidationOrdersService {
	s.pair = &pair
	return s
}

// StartTime set startTime
func (s *ListLiquidationOrdersService) StartTime(startTime int64) *ListLiquidationOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set startTime
func (s *ListLiquidationOrdersService) EndTime(endTime int64) *ListLiquidationOrdersService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListLiquidationOrdersService) Limit(limit int) *ListLiquidationOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListLiquidationOrdersService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*LiquidationOrder, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/allForceOrders",
		SecType:  binance.SecTypeNone,
	}
	if s.pair != nil {
		r.SetParam("pair", *s.pair)
	}
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
	res = make([]*LiquidationOrder, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
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
