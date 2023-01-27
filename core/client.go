package core

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/bububa/dataoke-go/util"
)

// Client sdk client
type Client struct {
	http      *http.Client
	appKey    string
	appSecret string
	version   string
	debug     bool
}

// NewClient returns a sdk Client instance
func NewClient(appKey string, appSecret string) *Client {
	return &Client{
		http:      http.DefaultClient,
		appKey:    appKey,
		appSecret: appSecret,
		version:   VERSION,
	}
}

// SetHttpClient set custom http.Client for Client
func (c *Client) SetHttpClient(clt *http.Client) {
	c.http = clt
}

// SetDebug set debug mode for Client
func (c *Client) SetDebug(debug bool) {
	c.debug = true
}

// SetVersion change Client version
func (c *Client) SetVersion(version string) {
	c.version = version
}

// Get http get
func (c *Client) Get(req Request, resp interface{}) error {
	values := util.GetUrlValues()
	req.Values(values)
	c.sign(values)
	gw := util.StringsJoin(GATEWAY, req.Url(), "?", values.Encode())
	util.PutUrlValues(values)
	if c.debug {
		log.Println("[DATAOKE] [GET]: ", gw)
	}
	httpReq, err := http.NewRequest(http.MethodGet, gw, nil)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	return c.fetch(httpReq, resp)
}

// Post http post
func (c *Client) Post(req Request, resp interface{}) error {
	values := util.GetUrlValues()
	req.Values(values)
	c.sign(values)
	gw := util.StringsJoin(GATEWAY, req.Url())
	httpReq, err := http.NewRequest(http.MethodPost, gw, strings.NewReader(values.Encode()))
	if c.debug {
		log.Println("[DATAOKE] [POST]: ", gw)
		log.Println("[DATAOKE] [PARAMS]: ", gw)
		for k := range values {
			log.Printf("[DATAOKE] [PARAMS] %s=%s\n", k, values.Get(k))
		}
	}
	util.PutUrlValues(values)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	return c.fetch(httpReq, resp)
}

func (c *Client) fetch(httpReq *http.Request, resp interface{}) error {
	httpResp, err := c.http.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	ret := GetResponse()
	defer PutResponse(ret)
	var decoder *json.Decoder
	if c.debug {
		body, err := ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
		if err := json.Indent(buf, body, "", "    "); err == nil {
			body = buf.Bytes()
		}
		log.Println("[DATAOKE] [RESP]: ", buf)
		decoder = json.NewDecoder(bytes.NewReader(body))
	} else {
		decoder = json.NewDecoder(httpResp.Body)
	}
	if err := decoder.Decode(ret); err != nil {
		return err
	}
	return ret.Decode(resp)
}

func (c *Client) sign(values url.Values) {
	values.Set("appKey", c.appKey)
	values.Set("version", c.version)
	keys := make([]string, 0, len(values)+2)
	var length int
	for k := range values {
		keys = append(keys, k)
		length += len(k) + 2 + len(values.Get(k))
	}
	sort.Strings(keys)
	length += 4 + len(c.appSecret)
	b := util.GetBytesBuffer()
	b.Grow(length)
	for _, k := range keys {
		b.WriteString(k)
		b.WriteByte('=')
		b.WriteString(values.Get(k))
		b.WriteByte('&')
	}
	b.WriteString("key=")
	b.WriteString(c.appSecret)
	h := md5.New()
	h.Write(b.Bytes())
	values.Set("sign", strings.ToUpper(hex.EncodeToString(h.Sum(nil))))
	util.PutBytesBuffer(b)
}
