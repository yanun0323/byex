package byex

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// API endpoints
	_baseUrlExchange = "https://openapi.100ex.com"
	_baseUrlFutures  = "https://futuresopenapi.100ex.com"

	// Testnet API endpoints
	_baseUrlTestnetExchange = "https://openapi.100extest.com"
	_baseUrlTestnetFutures  = "https://futuresopenapi.100extest.com"
)

var (
	defaultClientOption = ClientOption{
		Testnet:        false,
		HttpClientHook: nil,
	}
)

// Client represents the 100EX API client
type Client struct {
	apiKey     string
	secretKey  string
	httpClient *http.Client
	Testnet    bool
}

type ClientOption struct {
	Testnet        bool
	HttpClientHook []func(*http.Client)
}

// NewClient creates a new client
func NewClient(apiKey, secretKey string, opt ...ClientOption) *Client {
	o := defaultClientOption
	if len(opt) != 0 {
		o = opt[0]
	}

	c := &Client{
		apiKey:     apiKey,
		secretKey:  secretKey,
		httpClient: &http.Client{Timeout: 30 * time.Second},
		Testnet:    o.Testnet,
	}

	for _, hook := range o.HttpClientHook {
		hook(c.httpClient)
	}

	return c
}

// Exchange returns a new ExchangeAPI instance
func (c *Client) Exchange() *ExchangeAPI {
	return NewExchangeAPI(c)
}

// Futures returns a new FuturesAPI instance
func (c *Client) Futures() *FuturesAPI {
	return NewFuturesAPI(c)
}

// BaseResponse represents the common response structure
type BaseResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Error represents an API error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("API Error - Code: %s, Message: %s", e.Code, e.Message)
}

func (c *Client) baseUrlExchange() string {
	if c.Testnet {
		return _baseUrlTestnetExchange
	}
	return _baseUrlExchange
}

func (c *Client) baseUrlFutures() string {
	if c.Testnet {
		return _baseUrlTestnetFutures
	}
	return _baseUrlFutures
}

// generateExchangeSignature generates signature for exchange APIs
func (c *Client) generateExchangeSignature(params map[string]string) string {
	// Add required parameters
	params["api_key"] = c.apiKey
	params["time"] = strconv.FormatInt(time.Now().UnixMilli(), 10)

	// Sort keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build sign string
	var parts []string
	for _, k := range keys {
		if params[k] != "" {
			parts = append(parts, k+params[k])
		}
	}

	signString := strings.Join(parts, "") + c.secretKey

	// Generate MD5 hash
	hash := md5.New()
	hash.Write([]byte(signString))
	return hex.EncodeToString(hash.Sum(nil))
}

// generateFuturesSignature generates signature for futures APIs
func (c *Client) generateFuturesSignature(method, path, queryString string, timestamp int64) string {
	message := fmt.Sprintf("%d%s%s", timestamp, method, path)
	if queryString != "" {
		message += "?" + queryString
	}

	mac := hmac.New(sha256.New, []byte(c.secretKey))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

// doExchangeRequest performs HTTP request for exchange APIs
func (c *Client) doExchangeRequest(method, path string, params map[string]string) (*BaseResponse, error) {
	if params == nil {
		params = make(map[string]string)
	}

	// Generate signature
	sign := c.generateExchangeSignature(params)
	params["sign"] = sign

	// Build URL
	reqURL := c.baseUrlExchange() + path

	var req *http.Request
	var err error

	if method == "GET" {
		// For GET requests, add params to query string
		u, _ := url.Parse(reqURL)
		q := u.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		u.RawQuery = q.Encode()
		reqURL = u.String()

		req, err = http.NewRequest(method, reqURL, nil)
	} else {
		// For POST requests, send as form data
		form := url.Values{}
		for k, v := range params {
			form.Set(k, v)
		}

		req, err = http.NewRequest(method, reqURL, strings.NewReader(form.Encode()))
		if err == nil {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.executeRequest(req)
}

// doFuturesRequest performs HTTP request for futures APIs
func (c *Client) doFuturesRequest(method, path string, params interface{}) (*BaseResponse, error) {
	timestamp := time.Now().UnixMilli()

	// Build URL
	reqURL := c.baseUrlFutures() + path

	var req *http.Request
	var err error
	var queryString string

	if method == "GET" && params != nil {
		// For GET requests with params, add to query string
		if paramMap, ok := params.(map[string]string); ok {
			u, _ := url.Parse(reqURL)
			q := u.Query()
			for k, v := range paramMap {
				if v != "" {
					q.Set(k, v)
				}
			}
			if len(q) > 0 {
				queryString = q.Encode()
				u.RawQuery = queryString
				reqURL = u.String()
			}
		}
		req, err = http.NewRequest(method, reqURL, nil)
	} else if method == "POST" && params != nil {
		// For POST requests, send as JSON
		jsonData, jsonErr := json.Marshal(params)
		if jsonErr != nil {
			return nil, fmt.Errorf("failed to marshal params: %w", jsonErr)
		}
		req, err = http.NewRequest(method, reqURL, bytes.NewBuffer(jsonData))
		if err == nil {
			req.Header.Set("Content-Type", "application/json")
		}
	} else {
		req, err = http.NewRequest(method, reqURL, nil)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Generate signature
	sign := c.generateFuturesSignature(method, path, queryString, timestamp)

	// Set required headers
	req.Header.Set("X-CH-APIKEY", c.apiKey)
	req.Header.Set("X-CH-TS", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-CH-SIGN", sign)

	return c.executeRequest(req)
}

// executeRequest executes the HTTP request and parses response
func (c *Client) executeRequest(req *http.Request) (*BaseResponse, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var baseResp BaseResponse
	if err := json.Unmarshal(body, &baseResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	// Check for API errors
	if baseResp.Code != "0" && baseResp.Code != "" {
		return nil, &Error{
			Code:    baseResp.Code,
			Message: baseResp.Msg,
		}
	}

	return &baseResp, nil
}
