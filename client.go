package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Endpoints
const (
	baseAPIMainURL    = "https://api1.binance.com"
	baseAPITestnetURL = "https://testnet.binance.vision"

	baseFuturesApiMainUrl    = "https://fapi.binance.com"
	baseFuturesApiTestnetUrl = "https://testnet.binancefuture.com"

	baseDeliveryApiMainUrl    = "https://dapi.binance.com"
	baseDeliveryApiTestnetUrl = "https://testnet.binancefuture.com"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	TimeOffset int64
	Logger     logger
	//do         func(req *http.Request) (*http.Response, error)
}

type logger interface {
	Printf(format string, v ...interface{})
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseAPIMainURL,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

func NewTestClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseAPITestnetURL,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

func NewFuturesClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseFuturesApiMainUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

func NewFuturesTestClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseFuturesApiTestnetUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}
func NewDeliveryClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseDeliveryApiMainUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

func NewDeliveryTestClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseDeliveryApiTestnetUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}
func CurrentTimestamp() int64 {
	now := time.Now()
	return now.UnixNano() / int64(time.Millisecond)
}

func (c *Client) parseRequest(r *Request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	r.validate()
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.Endpoint)
	if r.recvWindow > 0 {
		r.SetParam(recvWindowKey, r.recvWindow)
	}
	if r.SecType == SecTypeSigned {
		r.SetParam(timestampKey, CurrentTimestamp()-c.TimeOffset)
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.SecType == SecTypeAPIKey || r.SecType == SecTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}
	if r.SecType == SecTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) Request(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, apiErr *APIError) {
	err := c.parseRequest(r, opts...)
	if err != nil {
		return nil, NewApiErr(err.Error())
	}
	req, _ := http.NewRequest(r.Method, r.fullURL, r.body)
	req.WithContext(ctx)
	req.Header = r.header
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, NewApiErr(err.Error())
	}
	bodyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, NewApiErr(err.Error())
	}
	defer func() {
		cerr := res.Body.Close()
		if apiErr == nil && cerr != nil {
			apiErr = NewApiErr(cerr.Error())
		}
	}()
	//c.debug("response header:%v", res.Header)
	if c.Debug {
		for k, h := range res.Header {
			if strings.Contains(k, "Weight") {
				c.debug("%s:%v", k, h)
			}
		}
	}
	if res.StatusCode >= 400 {
		apiErr = new(APIError)
		e := json.Unmarshal(bodyData, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
			apiErr.Message = e.Error()
		}
		apiErr.Header = res.Header
		apiErr.HttpCode = res.StatusCode
		return nil, apiErr
	}
	return bodyData, nil
}
func NewJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}
