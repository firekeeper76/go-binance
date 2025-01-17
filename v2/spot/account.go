package spot

import (
	"context"
	"encoding/json"
	"github.com/adshao/go-binance/v2"
)

// GetAccountService get account info
type GetAccountService struct {
	C *binance.Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...binance.RequestOption) (res *Account, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/api/v3/account",
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
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	Balances         []Balance `json:"balances"`
}

// Balance define user balance of your account
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// GetAccountSnapshotService all account orders; active, canceled, or filled
type GetAccountSnapshotService struct {
	C           *binance.Client
	accountType string
	startTime   *int64
	endTime     *int64
	limit       *int
}

// Type set account type ("SPOT", "MARGIN", "FUTURES")
func (s *GetAccountSnapshotService) Type(accountType string) *GetAccountSnapshotService {
	s.accountType = accountType
	return s
}

// StartTime set starttime
func (s *GetAccountSnapshotService) StartTime(startTime int64) *GetAccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *GetAccountSnapshotService) EndTime(endTime int64) *GetAccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAccountSnapshotService) Limit(limit int) *GetAccountSnapshotService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAccountSnapshotService) Do(ctx context.Context, opts ...binance.RequestOption) (res *Snapshot, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/accountSnapshot",
		SecType:  binance.SecTypeSigned,
	}
	r.SetParam("type", s.accountType)

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
	res = new(Snapshot)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}

// Snapshot define snapshot
type Snapshot struct {
	Code     int            `json:"code"`
	Msg      string         `json:"msg"`
	Snapshot []*SnapshotVos `json:"snapshotVos"`
}

// SnapshotVos define content of a snapshot
type SnapshotVos struct {
	Data       *SnapshotData `json:"data"`
	Type       string        `json:"type"`
	UpdateTime int64         `json:"updateTime"`
}

// SnapshotData define content of a snapshot
type SnapshotData struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`

	Balances   []*SnapshotBalances   `json:"balances"`
	UserAssets []*SnapshotUserAssets `json:"userAssets"`
	Assets     []*SnapshotAssets     `json:"assets"`
	Positions  []*SnapshotPositions  `json:"position"`
}

// SnapshotBalances define snapshot balances
type SnapshotBalances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// SnapshotUserAssets define snapshot user assets
type SnapshotUserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// SnapshotAssets define snapshot assets
type SnapshotAssets struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

// SnapshotPositions define snapshot positions
type SnapshotPositions struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

type AccountStatusService struct {
	C *binance.Client
}

// Do send request
func (s *AccountStatusService) Do(ctx context.Context, opts ...binance.RequestOption) (status string, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/account/status",
		SecType:  binance.SecTypeSigned,
	}
	data, rErr := s.C.Request(ctx, r)
	if rErr != nil {
		return "", rErr
	}
	m := map[string]string{}
	jErr := json.Unmarshal(data, &m)
	if jErr != nil {
		return "", binance.NewApiErr(jErr.Error())
	}
	status = m["data"]
	return status, nil
}


type GetAccountV2Service struct {
	C *binance.Client
}

func (s *GetAccountV2Service) Do(ctx context.Context) (*AccountV2Response, *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/fapi/v2/account",
		SecType:  binance.SecTypeSigned,
	}
	res, err := s.C.Request(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := new(AccountV2Response)
	jErr := json.Unmarshal(res, resp)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return resp, nil
}

type AccountV2Response struct {
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