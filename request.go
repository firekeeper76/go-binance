package binance

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

const (
	SecTypeNone secType = iota
	SecTypeAPIKey
	SecTypeSigned // if the 'timestamp' parameter is required
)

type Params map[string]interface{}

// Request define an API request
type Request struct {
	Method     string
	Endpoint   string
	query      url.Values
	form       url.Values
	recvWindow int64
	SecType    secType
	header     http.Header
	body       io.Reader
	fullURL    string
}

// AddParam add param with key/value to query string
func (r *Request) AddParam(key string, value interface{}) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// SetParam set param with key/value to query string
func (r *Request) SetParam(key string, value interface{}) *Request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// SetParams set Params with key/values to query string
func (r *Request) SetParams(m Params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

// SetFormParam set param with key/value to request form body
func (r *Request) SetFormParam(key string, value interface{}) *Request {
	if r.form == nil {
		r.form = url.Values{}
	}
	r.form.Set(key, fmt.Sprintf("%v", value))
	return r
}

// SetFormParams set Params with key/values to request form body
func (r *Request) SetFormParams(m Params) *Request {
	for k, v := range m {
		r.SetFormParam(k, v)
	}
	return r
}

func (r *Request) validate() {
	if r.query == nil {
		r.query = url.Values{}
	}
	if r.form == nil {
		r.form = url.Values{}
	}
	return
}


// RequestOption define option type for request
type RequestOption func(*Request)

// WithRecvWindow set recvWindow param for the request
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *Request) {
		r.recvWindow = recvWindow
	}
}

// WithHeader set or add a header value to the request
func WithHeader(key, value string, replace bool) RequestOption {
	return func(r *Request) {
		if r.header == nil {
			r.header = http.Header{}
		}
		if replace {
			r.header.Set(key, value)
		} else {
			r.header.Add(key, value)
		}
	}
}

// WithHeaders set or replace the headers of the request
func WithHeaders(header http.Header) RequestOption {
	return func(r *Request) {
		r.header = header.Clone()
	}
}
