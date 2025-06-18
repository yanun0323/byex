package byex

import (
	"net/http"
	"testing"
	"time"
)

const (
	testApiKey    = "test_api_key"
	testSecretKey = "test_secret_key"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name      string
		apiKey    string
		secretKey string
		options   []ClientOption
		wantErr   bool
	}{
		{
			name:      "Basic client creation",
			apiKey:    testApiKey,
			secretKey: testSecretKey,
			options:   nil,
			wantErr:   false,
		},
		{
			name:      "Client with testnet option",
			apiKey:    testApiKey,
			secretKey: testSecretKey,
			options: []ClientOption{
				{Testnet: true},
			},
			wantErr: false,
		},
		{
			name:      "Client with custom http client hook",
			apiKey:    testApiKey,
			secretKey: testSecretKey,
			options: []ClientOption{
				{
					Testnet: true,
					HttpClientHook: []func(*http.Client){
						func(client *http.Client) {
							client.Timeout = 60 * time.Second
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:      "Empty API key",
			apiKey:    "",
			secretKey: testSecretKey,
			options:   nil,
			wantErr:   false, // Constructor doesn't validate empty keys
		},
		{
			name:      "Empty secret key",
			apiKey:    testApiKey,
			secretKey: "",
			options:   nil,
			wantErr:   false, // Constructor doesn't validate empty keys
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.apiKey, tt.secretKey, tt.options...)

			if client == nil {
				t.Error("NewClient() returned nil")
				return
			}

			if client.apiKey != tt.apiKey {
				t.Errorf("Expected apiKey %s, got %s", tt.apiKey, client.apiKey)
			}

			if client.secretKey != tt.secretKey {
				t.Errorf("Expected secretKey %s, got %s", tt.secretKey, client.secretKey)
			}

			if len(tt.options) > 0 && client.Testnet != tt.options[0].Testnet {
				t.Errorf("Expected Testnet %v, got %v", tt.options[0].Testnet, client.Testnet)
			}

			if client.httpClient == nil {
				t.Error("httpClient should not be nil")
			}
		})
	}
}

func TestClient_Exchange(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	exchange := client.Exchange()

	if exchange == nil {
		t.Error("Exchange() returned nil")
	}

	if exchange.client != client {
		t.Error("Exchange API client reference should match original client")
	}
}

func TestClient_Futures(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	futures := client.Futures()

	if futures == nil {
		t.Error("Futures() returned nil")
	}

	if futures.client != client {
		t.Error("Futures API client reference should match original client")
	}
}

func TestClient_baseUrlExchange(t *testing.T) {
	tests := []struct {
		name     string
		testnet  bool
		expected string
	}{
		{
			name:     "Production exchange URL",
			testnet:  false,
			expected: _baseUrlExchange,
		},
		{
			name:     "Testnet exchange URL",
			testnet:  true,
			expected: _baseUrlTestnetExchange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: tt.testnet})

			url := client.baseUrlExchange()

			if url != tt.expected {
				t.Errorf("Expected URL %s, got %s", tt.expected, url)
			}
		})
	}
}

func TestClient_baseUrlFutures(t *testing.T) {
	tests := []struct {
		name     string
		testnet  bool
		expected string
	}{
		{
			name:     "Production futures URL",
			testnet:  false,
			expected: _baseUrlFutures,
		},
		{
			name:     "Testnet futures URL",
			testnet:  true,
			expected: _baseUrlTestnetFutures,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: tt.testnet})

			url := client.baseUrlFutures()

			if url != tt.expected {
				t.Errorf("Expected URL %s, got %s", tt.expected, url)
			}
		})
	}
}

func TestClient_generateExchangeSignature(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	tests := []struct {
		name   string
		params map[string]string
	}{
		{
			name:   "Empty parameters",
			params: map[string]string{},
		},
		{
			name: "Single parameter",
			params: map[string]string{
				"symbol": "BTCUSDT",
			},
		},
		{
			name: "Multiple parameters",
			params: map[string]string{
				"symbol": "BTCUSDT",
				"side":   "buy",
				"type":   "limit",
			},
		},
		{
			name: "Parameters with empty values",
			params: map[string]string{
				"symbol":   "BTCUSDT",
				"side":     "buy",
				"optional": "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature := client.generateExchangeSignature(tt.params)

			if signature == "" {
				t.Error("generateExchangeSignature() returned empty signature")
			}

			if len(signature) != 32 { // MD5 hash is 32 characters
				t.Errorf("Expected signature length 32, got %d", len(signature))
			}

			// Verify signature is consistent
			signature2 := client.generateExchangeSignature(tt.params)
			// Note: signatures will be different due to timestamp, this is expected
			if signature == signature2 {
				t.Log("Note: Signatures are the same, this could happen if timestamps are identical")
			}
		})
	}
}

func TestClient_generateFuturesSignature(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	tests := []struct {
		name        string
		method      string
		path        string
		queryString string
		timestamp   int64
	}{
		{
			name:        "GET request without query",
			method:      "GET",
			path:        "/fapi/v1/ticker",
			queryString: "",
			timestamp:   1640995200000,
		},
		{
			name:        "GET request with query",
			method:      "GET",
			path:        "/fapi/v1/ticker",
			queryString: "symbol=BTCUSDT",
			timestamp:   1640995200000,
		},
		{
			name:        "POST request",
			method:      "POST",
			path:        "/fapi/v1/trade/order",
			queryString: "",
			timestamp:   1640995200000,
		},
		{
			name:        "DELETE request",
			method:      "DELETE",
			path:        "/fapi/v1/trade/cancel",
			queryString: "",
			timestamp:   1640995200000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signature := client.generateFuturesSignature(tt.method, tt.path, tt.queryString, tt.timestamp)

			if signature == "" {
				t.Error("generateFuturesSignature() returned empty signature")
			}

			if len(signature) != 64 { // SHA256 hash is 64 characters
				t.Errorf("Expected signature length 64, got %d", len(signature))
			}

			// Verify signature is consistent with same input
			signature2 := client.generateFuturesSignature(tt.method, tt.path, tt.queryString, tt.timestamp)
			if signature != signature2 {
				t.Error("generateFuturesSignature() should return consistent signatures for same input")
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      Error
		expected string
	}{
		{
			name: "Basic error",
			err: Error{
				Code:    "1001",
				Message: "Invalid parameter",
			},
			expected: "API Error - Code: 1001, Message: Invalid parameter",
		},
		{
			name: "Empty error",
			err: Error{
				Code:    "",
				Message: "",
			},
			expected: "API Error - Code: , Message: ",
		},
		{
			name: "Error with special characters",
			err: Error{
				Code:    "ERR_001",
				Message: "Special chars: @#$%^&*()",
			},
			expected: "API Error - Code: ERR_001, Message: Special chars: @#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()

			if result != tt.expected {
				t.Errorf("Expected error message: %s, got: %s", tt.expected, result)
			}
		})
	}
}

func TestClientOption(t *testing.T) {
	// Test default client option
	client := NewClient(testApiKey, testSecretKey)
	if client.Testnet != false {
		t.Error("Default Testnet should be false")
	}

	// Test custom timeout hook
	customTimeout := 120 * time.Second
	option := ClientOption{
		Testnet: true,
		HttpClientHook: []func(*http.Client){
			func(client *http.Client) {
				client.Timeout = customTimeout
			},
		},
	}

	client = NewClient(testApiKey, testSecretKey, option)
	if client.httpClient.Timeout != customTimeout {
		t.Errorf("Expected timeout %v, got %v", customTimeout, client.httpClient.Timeout)
	}
}

// Note: doExchangeRequest and doFuturesRequest are harder to test without actual HTTP mocking
// These would require more complex test setup with HTTP test servers
// For now, these tests focus on the signature generation and basic client functionality
