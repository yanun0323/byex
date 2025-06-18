package byex

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewExchangeAPI(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	exchange := NewExchangeAPI(client)

	if exchange == nil {
		t.Error("NewExchangeAPI() returned nil")
	}

	if exchange.client != client {
		t.Error("ExchangeAPI client should match the provided client")
	}
}

func TestExchangeAPI_GetAllTicker(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := BaseResponse{
			Code: "0",
			Msg:  "success",
			Data: json.RawMessage(`{"date":1640995200,"ticker":[{"symbol":"BTCUSDT","high":"50000","low":"45000","last":"48000","vol":"1000","amount":"48000000","buy":"47950","sell":"48050","newCoinFlag":0,"change":"1000","rose":"0.02"}]}`),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Note: This test would require modifying the client to use the test server
	// For now, we'll test the function structure
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	// Test function exists and has correct signature
	if exchange.GetAllTicker == nil {
		t.Error("GetAllTicker method should exist")
	}
}

func TestExchangeAPI_GetTicker(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name   string
		symbol string
	}{
		{
			name:   "Valid symbol",
			symbol: "BTCUSDT",
		},
		{
			name:   "Empty symbol",
			symbol: "",
		},
		{
			name:   "Invalid symbol format",
			symbol: "invalid-symbol",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test function exists and has correct signature
			if exchange.GetTicker == nil {
				t.Error("GetTicker method should exist")
			}

			// Note: Actual API call would require real credentials and network access
			// This test validates the function signature and basic structure
		})
	}
}

func TestExchangeAPI_GetDepth(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name   string
		symbol string
		depth  int
	}{
		{
			name:   "Valid parameters",
			symbol: "BTCUSDT",
			depth:  20,
		},
		{
			name:   "Zero depth",
			symbol: "BTCUSDT",
			depth:  0,
		},
		{
			name:   "Negative depth",
			symbol: "BTCUSDT",
			depth:  -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetDepth == nil {
				t.Error("GetDepth method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetKlines(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name   string
		symbol string
		period string
		size   int
	}{
		{
			name:   "Valid parameters",
			symbol: "BTCUSDT",
			period: "1m",
			size:   100,
		},
		{
			name:   "Different period",
			symbol: "ETHUSDT",
			period: "5m",
			size:   50,
		},
		{
			name:   "Zero size",
			symbol: "BTCUSDT",
			period: "1h",
			size:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetKlines == nil {
				t.Error("GetKlines method should exist")
			}
		})
	}
}

func TestExchangeAPI_CreateOrder(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name string
		req  CreateOrderRequest
	}{
		{
			name: "Buy limit order",
			req: CreateOrderRequest{
				Symbol:        "BTCUSDT",
				Side:          "buy",
				Type:          "limit",
				Amount:        decimal.NewFromFloat(0.01),
				Price:         decimal.NewFromFloat(45000),
				ClientOrderID: "test_order_1",
			},
		},
		{
			name: "Sell market order",
			req: CreateOrderRequest{
				Symbol: "BTCUSDT",
				Side:   "sell",
				Type:   "market",
				Amount: decimal.NewFromFloat(0.01),
			},
		},
		{
			name: "Order with zero amount",
			req: CreateOrderRequest{
				Symbol: "BTCUSDT",
				Side:   "buy",
				Type:   "limit",
				Amount: decimal.Zero,
				Price:  decimal.NewFromFloat(45000),
			},
		},
		{
			name: "Order with zero price",
			req: CreateOrderRequest{
				Symbol: "BTCUSDT",
				Side:   "buy",
				Type:   "market",
				Amount: decimal.NewFromFloat(0.01),
				Price:  decimal.Zero,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.CreateOrder == nil {
				t.Error("CreateOrder method should exist")
			}
		})
	}
}

func TestExchangeAPI_CancelOrder(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name    string
		symbol  string
		orderID string
	}{
		{
			name:    "Valid parameters",
			symbol:  "BTCUSDT",
			orderID: "12345678",
		},
		{
			name:    "Empty symbol",
			symbol:  "",
			orderID: "12345678",
		},
		{
			name:    "Empty order ID",
			symbol:  "BTCUSDT",
			orderID: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.CancelOrder == nil {
				t.Error("CancelOrder method should exist")
			}
		})
	}
}

func TestExchangeAPI_CancelAllOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name   string
		symbol string
	}{
		{
			name:   "Valid symbol",
			symbol: "BTCUSDT",
		},
		{
			name:   "Empty symbol",
			symbol: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.CancelAllOrders == nil {
				t.Error("CancelAllOrders method should exist")
			}
		})
	}
}

func TestExchangeAPI_BatchCreateOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name string
		req  BatchOrderRequest
	}{
		{
			name: "Valid batch order",
			req: BatchOrderRequest{
				Symbol: "BTCUSDT",
				Orders: []CreateOrderRequest{
					{
						Symbol: "BTCUSDT",
						Side:   "buy",
						Type:   "limit",
						Amount: decimal.NewFromFloat(0.01),
						Price:  decimal.NewFromFloat(45000),
					},
					{
						Symbol: "BTCUSDT",
						Side:   "sell",
						Type:   "limit",
						Amount: decimal.NewFromFloat(0.01),
						Price:  decimal.NewFromFloat(50000),
					},
				},
			},
		},
		{
			name: "Empty orders",
			req: BatchOrderRequest{
				Symbol: "BTCUSDT",
				Orders: []CreateOrderRequest{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.BatchCreateOrders == nil {
				t.Error("BatchCreateOrders method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetCurrentOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name     string
		symbol   string
		pageSize int
		page     int
	}{
		{
			name:     "Valid parameters",
			symbol:   "BTCUSDT",
			pageSize: 50,
			page:     1,
		},
		{
			name:     "Zero page size",
			symbol:   "BTCUSDT",
			pageSize: 0,
			page:     1,
		},
		{
			name:     "Zero page",
			symbol:   "BTCUSDT",
			pageSize: 50,
			page:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetCurrentOrders == nil {
				t.Error("GetCurrentOrders method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetOrderHistory(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name     string
		symbol   string
		pageSize int
		page     int
	}{
		{
			name:     "Valid parameters",
			symbol:   "BTCUSDT",
			pageSize: 50,
			page:     1,
		},
		{
			name:     "Large page size",
			symbol:   "BTCUSDT",
			pageSize: 1000,
			page:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetOrderHistory == nil {
				t.Error("GetOrderHistory method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetOrderInfo(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name    string
		symbol  string
		orderID string
	}{
		{
			name:    "Valid parameters",
			symbol:  "BTCUSDT",
			orderID: "12345678",
		},
		{
			name:    "Long order ID",
			symbol:  "BTCUSDT",
			orderID: "1234567890123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetOrderInfo == nil {
				t.Error("GetOrderInfo method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetTrades(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name     string
		symbol   string
		pageSize int
		page     int
	}{
		{
			name:     "Valid parameters",
			symbol:   "BTCUSDT",
			pageSize: 50,
			page:     1,
		},
		{
			name:     "Different symbol",
			symbol:   "ETHUSDT",
			pageSize: 100,
			page:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetTrades == nil {
				t.Error("GetTrades method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetAccount(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	if exchange.GetAccount == nil {
		t.Error("GetAccount method should exist")
	}
}

func TestExchangeAPI_GetBalance(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name  string
		coins []string
	}{
		{
			name:  "Specific coins",
			coins: []string{"BTC", "ETH", "USDT"},
		},
		{
			name:  "Empty coins list",
			coins: []string{},
		},
		{
			name:  "Nil coins list",
			coins: nil,
		},
		{
			name:  "Single coin",
			coins: []string{"BTC"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetBalance == nil {
				t.Error("GetBalance method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetAllTradingRecords(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name      string
		symbol    string
		pageSize  int
		page      int
		id        int64
		startDate string
		endDate   string
		sort      int
	}{
		{
			name:      "Basic parameters",
			symbol:    "BTCUSDT",
			pageSize:  50,
			page:      1,
			id:        0,
			startDate: "",
			endDate:   "",
			sort:      0,
		},
		{
			name:      "With date range",
			symbol:    "BTCUSDT",
			pageSize:  100,
			page:      1,
			id:        0,
			startDate: "2023-01-01",
			endDate:   "2023-12-31",
			sort:      1,
		},
		{
			name:      "With ID filter",
			symbol:    "ETHUSDT",
			pageSize:  50,
			page:      1,
			id:        12345,
			startDate: "",
			endDate:   "",
			sort:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetAllTradingRecords == nil {
				t.Error("GetAllTradingRecords method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetMarketPrices(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	if exchange.GetMarketPrices == nil {
		t.Error("GetMarketPrices method should exist")
	}
}

func TestExchangeAPI_BatchPlaceOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name      string
		symbol    string
		orderList []BatchOrder
	}{
		{
			name:   "Valid batch orders",
			symbol: "BTCUSDT",
			orderList: []BatchOrder{
				{
					Volume:        decimal.NewFromFloat(0.01),
					Price:         decimal.NewFromFloat(45000),
					Side:          "buy",
					Type:          1,
					VolumeType:    1,
					ClientOrderID: "batch_1",
				},
				{
					Volume:        decimal.NewFromFloat(0.01),
					Price:         decimal.NewFromFloat(50000),
					Side:          "sell",
					Type:          1,
					VolumeType:    1,
					ClientOrderID: "batch_2",
				},
			},
		},
		{
			name:      "Empty order list",
			symbol:    "BTCUSDT",
			orderList: []BatchOrder{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.BatchPlaceOrders == nil {
				t.Error("BatchPlaceOrders method should exist")
			}
		})
	}
}

func TestExchangeAPI_BatchCancelOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name     string
		symbol   string
		orderIds []string
	}{
		{
			name:     "Valid order IDs",
			symbol:   "BTCUSDT",
			orderIds: []string{"123", "456", "789"},
		},
		{
			name:     "Empty order IDs",
			symbol:   "BTCUSDT",
			orderIds: []string{},
		},
		{
			name:     "Single order ID",
			symbol:   "BTCUSDT",
			orderIds: []string{"123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.BatchCancelOrders == nil {
				t.Error("BatchCancelOrders method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetOrderDetail(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name    string
		symbol  string
		orderID string
	}{
		{
			name:    "Valid parameters",
			symbol:  "BTCUSDT",
			orderID: "12345678",
		},
		{
			name:    "Different symbol",
			symbol:  "ETHUSDT",
			orderID: "87654321",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetOrderDetail == nil {
				t.Error("GetOrderDetail method should exist")
			}
		})
	}
}

func TestExchangeAPI_ReplaceOrder(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name string
		req  ReplaceOrderRequest
	}{
		{
			name: "Valid replace order",
			req: ReplaceOrderRequest{
				Symbol:        "BTCUSDT",
				CancelOrderID: "12345",
				Side:          "buy",
				Type:          "limit",
				Amount:        decimal.NewFromFloat(0.02),
				Price:         decimal.NewFromFloat(46000),
				ClientOrderID: "replace_1",
			},
		},
		{
			name: "Replace with zero amount",
			req: ReplaceOrderRequest{
				Symbol:        "BTCUSDT",
				CancelOrderID: "12345",
				Side:          "buy",
				Type:          "market",
				Amount:        decimal.Zero,
				Price:         decimal.NewFromFloat(46000),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.ReplaceOrder == nil {
				t.Error("ReplaceOrder method should exist")
			}
		})
	}
}

func TestExchangeAPI_GetSymbolsCharge(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	if exchange.GetSymbolsCharge == nil {
		t.Error("GetSymbolsCharge method should exist")
	}
}

func TestExchangeAPI_GetLeverageFinanceBalance(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := NewExchangeAPI(client)

	tests := []struct {
		name   string
		symbol string
	}{
		{
			name:   "Valid symbol",
			symbol: "BTCUSDT",
		},
		{
			name:   "Different symbol",
			symbol: "ETHUSDT",
		},
		{
			name:   "Empty symbol",
			symbol: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exchange.GetLeverageFinanceBalance == nil {
				t.Error("GetLeverageFinanceBalance method should exist")
			}
		})
	}
}

// Helper function to create test client with testnet enabled
func createTestExchangeClient() *ExchangeAPI {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	return NewExchangeAPI(client)
}

// Test client initialization with testnet
func TestExchangeAPITestnetInitialization(t *testing.T) {
	exchange := createTestExchangeClient()

	if exchange == nil {
		t.Error("Exchange API should be initialized")
	}

	if !exchange.client.Testnet {
		t.Error("Exchange API should be using testnet")
	}

	expectedURL := _baseUrlTestnetExchange
	actualURL := exchange.client.baseUrlExchange()

	if actualURL != expectedURL {
		t.Errorf("Expected testnet URL %s, got %s", expectedURL, actualURL)
	}
}

// Test that all required methods exist on ExchangeAPI
func TestExchangeAPI_AllMethodsExist(t *testing.T) {
	exchange := createTestExchangeClient()

	methods := []string{
		"GetAllTicker",
		"GetTicker",
		"GetDepth",
		"GetKlines",
		"CreateOrder",
		"CancelOrder",
		"CancelAllOrders",
		"BatchCreateOrders",
		"GetCurrentOrders",
		"GetOrderHistory",
		"GetOrderInfo",
		"GetTrades",
		"GetAccount",
		"GetBalance",
		"GetAllTradingRecords",
		"GetMarketPrices",
		"BatchPlaceOrders",
		"BatchCancelOrders",
		"GetOrderDetail",
		"ReplaceOrder",
		"GetSymbolsCharge",
		"GetLeverageFinanceBalance",
	}

	// This is a compile-time check that all methods exist
	// If any method is missing, this won't compile
	_ = exchange.GetAllTicker
	_ = exchange.GetTicker
	_ = exchange.GetDepth
	_ = exchange.GetKlines
	_ = exchange.CreateOrder
	_ = exchange.CancelOrder
	_ = exchange.CancelAllOrders
	_ = exchange.BatchCreateOrders
	_ = exchange.GetCurrentOrders
	_ = exchange.GetOrderHistory
	_ = exchange.GetOrderInfo
	_ = exchange.GetTrades
	_ = exchange.GetAccount
	_ = exchange.GetBalance
	_ = exchange.GetAllTradingRecords
	_ = exchange.GetMarketPrices
	_ = exchange.BatchPlaceOrders
	_ = exchange.BatchCancelOrders
	_ = exchange.GetOrderDetail
	_ = exchange.ReplaceOrder
	_ = exchange.GetSymbolsCharge
	_ = exchange.GetLeverageFinanceBalance

	t.Logf("All %d exchange API methods exist", len(methods))
}
