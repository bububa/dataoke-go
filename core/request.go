package core

import (
	"encoding/json"
	"net/url"
	"strconv"
	"sync"

	"github.com/bububa/dataoke-go/util"
)

// Request api request interface
type Request interface {
	Values(url.Values)
	Url() string
}

// Response api response
type Response struct {
	RequestID string          `json:"requestId,omitempty"`
	Time      int64           `json:"time,omitempty"`
	Code      int             `json:"code,omitempty"`
	Msg       string          `json:"msg,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// IsError check api response is error
func (r Response) IsError() bool {
	return r.Code != 0
}

// Error implement error interface
func (r Response) Error() string {
	return util.StringsJoin("code:", strconv.Itoa(r.Code), ", msg:", r.Msg)
}

// Decode api response result
func (r Response) Decode(resp interface{}) error {
	if r.IsError() {
		return r
	}
	if resp != nil && r.Data != nil {
		return json.Unmarshal(r.Data, resp)
	}
	return nil
}

var responsePool = sync.Pool{
	New: func() any {
		return new(Response)
	},
}

// GetResponse get response from sync.Pool
func GetResponse() *Response {
	return responsePool.Get().(*Response)
}

// PutResponse put response to sync.Pool
func PutResponse(r *Response) {
	r.RequestID = ""
	r.Time = 0
	r.Code = 0
	r.Msg = ""
	r.Data = nil
	responsePool.Put(r)
}
