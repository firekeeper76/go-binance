package spot

import (
	"github.com/adshao/go-binance/v2"
	"context"
	"encoding/json"
)

type ParentSpotSummaryService struct {
	C     *binance.Client
	email *string
	page  *int
	size  *int
}
type ParentSpotSummaryResponse struct {
	TotalCount                int    `json:"totalCount"`
	MasterAccountTotalAsset   string `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []struct {
		Email      string `json:"email"`
		TotalAsset string `json:"totalAsset"`
	} `json:"spotSubUserAssetBtcVoList"`
}

func (s *ParentSpotSummaryService) Email(v string) *ParentSpotSummaryService {
	s.email = &v
	return s
}
func (s *ParentSpotSummaryService) Page(p int) *ParentSpotSummaryService {
	s.page = &p
	return s
}
func (s *ParentSpotSummaryService) Size(sz int) *ParentSpotSummaryService {
	s.size = &sz
	return s
}

// Do send request
func (s *ParentSpotSummaryService) Do(ctx context.Context, opts ...binance.RequestOption) (res *ParentSpotSummaryResponse, err *binance.APIError) {
	r := &binance.Request{
		Method:   "GET",
		Endpoint: "/sapi/v1/sub-account/spotSummary",
		SecType:  binance.SecTypeSigned,
	}
	if s.email != nil {
		r.SetParam("email", *s.email)
	}
	if s.page != nil {
		r.SetParam("page", *s.page)
	}
	if s.size != nil {
		r.SetParam("size", *s.size)
	}
	data, rErr := s.C.Request(ctx, r, opts...)
	if rErr != nil {
		return nil, rErr
	}
	res = new(ParentSpotSummaryResponse)
	jErr := json.Unmarshal(data, res)
	if jErr != nil {
		return nil, binance.NewApiErr(jErr.Error())
	}
	return res, nil
}
