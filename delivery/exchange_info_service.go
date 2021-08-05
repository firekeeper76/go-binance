package delivery

import (
	"go-binance"
	"context"
	"encoding/json"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	C *binance.Client
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...binance.RequestOption) (res *ExchangeInfo, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/dapi/v1/exchangeInfo",
		SecType:  binance.SecTypeNone,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(ExchangeInfo)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}

// Symbol market symbol
type Symbol struct {
	OrderType             []OrderType              `json:"OrderType"`
	TimeInForce           []TimeInForceType        `json:"timeInForce"`
	Filters               []map[string]interface{} `json:"filters"`
	Symbol                string                   `json:"symbol"`
	Pair                  string                   `json:"pair"`
	ContractType          string                   `json:"contractType"`
	DeliveryDate          int64                    `json:"deliveryDate"`
	OnboardDate           int64                    `json:"onboardDate"`
	ContractStatus        string                   `json:"contractStatus"`
	ContractSize          int                      `json:"contractSize"`
	PricePrecision        int                      `json:"pricePrecision"`
	QuantityPrecision     int                      `json:"quantityPrecision"`
	MaintMarginPercent    string                   `json:"maintMarginPercent"`
	RequiredMarginPercent string                   `json:"requiredMarginPercent"`
	QuoteAsset            string                   `json:"quoteAsset"`
	BaseAsset             string                   `json:"baseAsset"`
	MarginAsset           string                   `json:"marginAsset"`
	BaseAssetPrecision    int                      `json:"baseAssetPrecision"`
	QuotePrecision        int                      `json:"quotePrecision"`
	EqualQtyPrecision     int                      `json:"equalQtyPrecision"`
	TriggerProtect        string                   `json:"triggerProtect"`
	UnderlyingType        string                   `json:"underlyingType"`
	UnderlyingSubType     []interface{}            `json:"underlyingSubType"`
}

// LotSizeFilter define lot size filter of symbol
type LotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// PriceFilter define price filter of symbol
type PriceFilter struct {
	MaxPrice string `json:"maxPrice"`
	MinPrice string `json:"minPrice"`
	TickSize string `json:"tickSize"`
}

// PercentPriceFilter define percent price filter of symbol
type PercentPriceFilter struct {
	MultiplierDecimal int    `json:"multiplierDecimal"`
	MultiplierUp      string `json:"multiplierUp"`
	MultiplierDown    string `json:"multiplierDown"`
}

// MarketLotSizeFilter define market lot size filter of symbol
type MarketLotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MaxNumOrdersFilter define max num orders filter of symbol
type MaxNumOrdersFilter struct {
	Limit int64 `json:"limit"`
}

// LotSizeFilter return lot size filter of symbol
func (s *Symbol) LotSizeFilter() *LotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeLotSize) {
			f := &LotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PriceFilter return price filter of symbol
func (s *Symbol) PriceFilter() *PriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePrice) {
			f := &PriceFilter{}
			if i, ok := filter["maxPrice"]; ok {
				f.MaxPrice = i.(string)
			}
			if i, ok := filter["minPrice"]; ok {
				f.MinPrice = i.(string)
			}
			if i, ok := filter["tickSize"]; ok {
				f.TickSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PercentPriceFilter return percent price filter of symbol
func (s *Symbol) PercentPriceFilter() *PercentPriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePercentPrice) {
			f := &PercentPriceFilter{}
			if i, ok := filter["multiplierDecimal"]; ok {
				f.MultiplierDecimal = int(i.(float64))
			}
			if i, ok := filter["multiplierUp"]; ok {
				f.MultiplierUp = i.(string)
			}
			if i, ok := filter["multiplierDown"]; ok {
				f.MultiplierDown = i.(string)
			}
			return f
		}
	}
	return nil
}

// MarketLotSizeFilter return market lot size filter of symbol
func (s *Symbol) MarketLotSizeFilter() *MarketLotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMarketLotSize) {
			f := &MarketLotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// MaxNumOrdersFilter return max num orders filter of symbol
func (s *Symbol) MaxNumOrdersFilter() *MaxNumOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumOrders) {
			f := &MaxNumOrdersFilter{}
			if i, ok := filter["limit"]; ok {
				f.Limit = int64(i.(float64))
			}
			return f
		}
	}
	return nil
}
