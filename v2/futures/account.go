package futures

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
)

// GetBalanceService get account balance
type GetBalanceService struct {
	C *binance.Client
}

// Do send request
func (s *GetBalanceService) Do(ctx context.Context, opts ...binance.RequestOption) (res []*Balance, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v2/balance",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = make([]*Balance, 0)
	jErr := json.Unmarshal(data, &res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Balance define user balance of your account
type Balance struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

// GetAccountService get account info
type GetAccountService struct {
	C *binance.Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...binance.RequestOption) (res *Account, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v1/account",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(Account)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Account define account info
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanTrade                    bool               `json:"canTrade"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	FeeTier                     int                `json:"feeTier"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	UpdateTime                  int64              `json:"updateTime"`
}

// AccountAsset define account asset
type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

// AccountPosition define account position
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               string           `json:"leverage"`
	InitialMargin          string           `json:"initialMargin"`
	MaintMargin            string           `json:"maintMargin"`
	OpenOrderInitialMargin string           `json:"openOrderInitialMargin"`
	PositionInitialMargin  string           `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       string           `json:"unrealizedProfit"`
	EntryPrice             string           `json:"entryPrice"`
	MaxNotional            string           `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            string           `json:"positionAmt"`
	Notional               string           `json:"notional"`
	IsolatedWallet         string           `json:"isolatedWallet"`
	UpdateTime             int64            `json:"updateTime"`
}

type GetAccountV2Service struct {
	C *binance.Client
}

func (s *GetAccountV2Service) Do(ctx context.Context) (*AccountV2, *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v2/account",
		SecType:  binance.SecTypeSigned,
	}
	res, err := s.C.Request(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := new(AccountV2)
	jErr := json.Unmarshal(res, resp)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return resp, nil
}

type AccountV2 struct {
	FeeTier                     int         `json:"feeTier"`
	CanTrade                    bool        `json:"canTrade"`
	CanDeposit                  bool        `json:"canDeposit"`
	CanWithdraw                 bool        `json:"canWithdraw"`
	UpdateTime                  int         `json:"updateTime"`
	TotalInitialMargin          string      `json:"totalInitialMargin"`
	TotalMaintMargin            string      `json:"totalMaintMargin"`
	TotalWalletBalance          string      `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string      `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string      `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string      `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string      `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string      `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string      `json:"totalCrossUnPnl"`
	AvailableBalance            string      `json:"availableBalance"`
	MaxWithdrawAmount           string      `json:"maxWithdrawAmount"`
	Assets                      []Assets    `json:"assets"`
	Positions                   []Positions `json:"positions"`
}
type Assets struct {
	Asset                  string `json:"asset"`
	WalletBalance          string `json:"walletBalance"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	MarginBalance          string `json:"marginBalance"`
	MaintMargin            string `json:"maintMargin"`
	InitialMargin          string `json:"initialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}
type Positions struct {
	Symbol                 string `json:"symbol"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	Leverage               string `json:"leverage"`
	Isolated               bool   `json:"isolated"`
	EntryPrice             string `json:"entryPrice"`
	MaxNotional            string `json:"maxNotional"`
	BidNotional            string `json:"bidNotional"`
	AskNotional            string `json:"askNotional"`
	PositionSide           string `json:"positionSide"`
	PositionAmt            string `json:"positionAmt"`
	UpdateTime             int    `json:"updateTime"`
}


type GetBalanceV2Service struct {
	C *binance.Client
}

func (s *GetBalanceV2Service) Do(ctx context.Context) ([]*BalanceV2, *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v2/balance",
		SecType:  binance.SecTypeSigned,
	}
	res, err := s.C.Request(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := make([]*BalanceV2, 0)
	jErr := json.Unmarshal(res, &resp)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return resp, nil
}

type BalanceV2 struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	AvailableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
	MarginAvailable    bool   `json:"marginAvailable"`
	UpdateTime         int64  `json:"updateTime"`
}