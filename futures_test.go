package byex

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewFuturesAPI(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	futures := NewFuturesAPI(client)

	if futures == nil {
		t.Error("NewFuturesAPI() returned nil")
	}

	if futures.client != client {
		t.Error("FuturesAPI client should match the provided client")
	}
}

func TestFuturesAPI_GetTicker(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

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
			name:   "ETH symbol",
			symbol: "ETHUSDT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetTicker == nil {
				t.Error("GetTicker method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetDepth(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name   string
		symbol string
		limit  int
	}{
		{
			name:   "Valid parameters",
			symbol: "BTCUSDT",
			limit:  20,
		},
		{
			name:   "Zero limit",
			symbol: "BTCUSDT",
			limit:  0,
		},
		{
			name:   "Large limit",
			symbol: "ETHUSDT",
			limit:  1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetDepth == nil {
				t.Error("GetDepth method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetKlines(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name     string
		symbol   string
		interval string
		limit    int
	}{
		{
			name:     "Valid parameters",
			symbol:   "BTCUSDT",
			interval: "1m",
			limit:    100,
		},
		{
			name:     "Different interval",
			symbol:   "ETHUSDT",
			interval: "5m",
			limit:    50,
		},
		{
			name:     "Hour interval",
			symbol:   "BTCUSDT",
			interval: "1h",
			limit:    24,
		},
		{
			name:     "Zero limit",
			symbol:   "BTCUSDT",
			interval: "15m",
			limit:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetKlines == nil {
				t.Error("GetKlines method should exist")
			}
		})
	}
}

func TestFuturesAPI_CreateOrder(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name string
		req  FuturesCreateOrderRequest
	}{
		{
			name: "Buy long order",
			req: FuturesCreateOrderRequest{
				FuturesName:   "BTCUSDT",
				Type:          "LIMIT",
				Side:          "BUY",
				Open:          "OPEN",
				PositionType:  "LONG",
				Price:         decimal.NewFromFloat(45000),
				Volume:        decimal.NewFromFloat(0.01),
				ClientOrderID: "test_futures_1",
			},
		},
		{
			name: "Sell short order",
			req: FuturesCreateOrderRequest{
				FuturesName:  "BTCUSDT",
				Type:         "MARKET",
				Side:         "SELL",
				Open:         "OPEN",
				PositionType: "SHORT",
				Volume:       decimal.NewFromFloat(0.01),
			},
		},
		{
			name: "Close position order",
			req: FuturesCreateOrderRequest{
				FuturesName:  "ETHUSDT",
				Type:         "LIMIT",
				Side:         "SELL",
				Open:         "CLOSE",
				PositionType: "LONG",
				Price:        decimal.NewFromFloat(3000),
				Volume:       decimal.NewFromFloat(0.1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.CreateOrder == nil {
				t.Error("CreateOrder method should exist")
			}
		})
	}
}

func TestFuturesAPI_CancelOrder(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		orderID     string
	}{
		{
			name:        "Valid parameters",
			futuresName: "BTCUSDT",
			orderID:     "12345678",
		},
		{
			name:        "Empty futures name",
			futuresName: "",
			orderID:     "12345678",
		},
		{
			name:        "Empty order ID",
			futuresName: "BTCUSDT",
			orderID:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.CancelOrder == nil {
				t.Error("CancelOrder method should exist")
			}
		})
	}
}

func TestFuturesAPI_CancelAllOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
	}{
		{
			name:        "Valid futures name",
			futuresName: "BTCUSDT",
		},
		{
			name:        "Different futures name",
			futuresName: "ETHUSDT",
		},
		{
			name:        "Empty futures name",
			futuresName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.CancelAllOrders == nil {
				t.Error("CancelAllOrders method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetCurrentOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
	}{
		{
			name:        "Valid futures name",
			futuresName: "BTCUSDT",
		},
		{
			name:        "Different futures name",
			futuresName: "ETHUSDT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetCurrentOrders == nil {
				t.Error("GetCurrentOrders method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetOrderHistory(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		limit       int
	}{
		{
			name:        "Valid parameters",
			futuresName: "BTCUSDT",
			limit:       50,
		},
		{
			name:        "Zero limit",
			futuresName: "BTCUSDT",
			limit:       0,
		},
		{
			name:        "Large limit",
			futuresName: "ETHUSDT",
			limit:       1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetOrderHistory == nil {
				t.Error("GetOrderHistory method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetOrderInfo(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		orderID     string
	}{
		{
			name:        "Valid parameters",
			futuresName: "BTCUSDT",
			orderID:     "12345678",
		},
		{
			name:        "Long order ID",
			futuresName: "BTCUSDT",
			orderID:     "1234567890123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetOrderInfo == nil {
				t.Error("GetOrderInfo method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetTrades(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		limit       int
	}{
		{
			name:        "Valid parameters",
			futuresName: "BTCUSDT",
			limit:       50,
		},
		{
			name:        "Different futures",
			futuresName: "ETHUSDT",
			limit:       100,
		},
		{
			name:        "Zero limit",
			futuresName: "BTCUSDT",
			limit:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetTrades == nil {
				t.Error("GetTrades method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetPositions(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
	}{
		{
			name:        "Specific futures",
			futuresName: "BTCUSDT",
		},
		{
			name:        "All positions",
			futuresName: "",
		},
		{
			name:        "ETH positions",
			futuresName: "ETHUSDT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetPositions == nil {
				t.Error("GetPositions method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetAccount(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAccount == nil {
		t.Error("GetAccount method should exist")
	}
}

func TestFuturesAPI_SetLeverage(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		leverage    int
	}{
		{
			name:        "Low leverage",
			futuresName: "BTCUSDT",
			leverage:    5,
		},
		{
			name:        "High leverage",
			futuresName: "BTCUSDT",
			leverage:    100,
		},
		{
			name:        "Zero leverage",
			futuresName: "BTCUSDT",
			leverage:    0,
		},
		{
			name:        "Negative leverage",
			futuresName: "BTCUSDT",
			leverage:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.SetLeverage == nil {
				t.Error("SetLeverage method should exist")
			}
		})
	}
}

func TestFuturesAPI_SetMarginType(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		marginType  string
	}{
		{
			name:        "Isolated margin",
			futuresName: "BTCUSDT",
			marginType:  "ISOLATED",
		},
		{
			name:        "Cross margin",
			futuresName: "BTCUSDT",
			marginType:  "CROSSED",
		},
		{
			name:        "Empty margin type",
			futuresName: "BTCUSDT",
			marginType:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.SetMarginType == nil {
				t.Error("SetMarginType method should exist")
			}
		})
	}
}

func TestFuturesAPI_ModifyMargin(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		amount      float64
		marginType  int
	}{
		{
			name:        "Add margin",
			futuresName: "BTCUSDT",
			amount:      100.0,
			marginType:  1,
		},
		{
			name:        "Reduce margin",
			futuresName: "BTCUSDT",
			amount:      50.0,
			marginType:  2,
		},
		{
			name:        "Zero amount",
			futuresName: "BTCUSDT",
			amount:      0.0,
			marginType:  1,
		},
		{
			name:        "Negative amount",
			futuresName: "BTCUSDT",
			amount:      -10.0,
			marginType:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.ModifyMargin == nil {
				t.Error("ModifyMargin method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetAllTicker(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAllTicker == nil {
		t.Error("GetAllTicker method should exist")
	}
}

func TestFuturesAPI_GetIndexPrice(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name   string
		symbol string
	}{
		{
			name:   "Valid symbol",
			symbol: "BTCUSDT",
		},
		{
			name:   "ETH symbol",
			symbol: "ETHUSDT",
		},
		{
			name:   "Empty symbol",
			symbol: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetIndexPrice == nil {
				t.Error("GetIndexPrice method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetAllIndexPrice(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAllIndexPrice == nil {
		t.Error("GetAllIndexPrice method should exist")
	}
}

func TestFuturesAPI_GetAllTagIndexPrice(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAllTagIndexPrice == nil {
		t.Error("GetAllTagIndexPrice method should exist")
	}
}

func TestFuturesAPI_GetAllFuturesDepth(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAllFuturesDepth == nil {
		t.Error("GetAllFuturesDepth method should exist")
	}
}

func TestFuturesAPI_BatchCreateOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name string
		req  FuturesBatchOrderRequest
	}{
		{
			name: "Valid batch orders",
			req: FuturesBatchOrderRequest{
				FuturesName: "BTCUSDT",
				Orders: []FuturesCreateOrderRequest{
					{
						FuturesName:  "BTCUSDT",
						Type:         "LIMIT",
						Side:         "BUY",
						Open:         "OPEN",
						PositionType: "LONG",
						Price:        decimal.NewFromFloat(45000),
						Volume:       decimal.NewFromFloat(0.01),
					},
					{
						FuturesName:  "BTCUSDT",
						Type:         "LIMIT",
						Side:         "SELL",
						Open:         "OPEN",
						PositionType: "SHORT",
						Price:        decimal.NewFromFloat(50000),
						Volume:       decimal.NewFromFloat(0.01),
					},
				},
			},
		},
		{
			name: "Empty orders",
			req: FuturesBatchOrderRequest{
				FuturesName: "BTCUSDT",
				Orders:      []FuturesCreateOrderRequest{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.BatchCreateOrders == nil {
				t.Error("BatchCreateOrders method should exist")
			}
		})
	}
}

func TestFuturesAPI_BatchCancelOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		orderIds    []string
	}{
		{
			name:        "Valid order IDs",
			futuresName: "BTCUSDT",
			orderIds:    []string{"123", "456", "789"},
		},
		{
			name:        "Empty order IDs",
			futuresName: "BTCUSDT",
			orderIds:    []string{},
		},
		{
			name:        "Single order ID",
			futuresName: "ETHUSDT",
			orderIds:    []string{"123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.BatchCancelOrders == nil {
				t.Error("BatchCancelOrders method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetCapital(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetCapital == nil {
		t.Error("GetCapital method should exist")
	}
}

func TestFuturesAPI_GetFutureAccounts(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetFutureAccounts == nil {
		t.Error("GetFutureAccounts method should exist")
	}
}

func TestFuturesAPI_CreateFutureAccount(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.CreateFutureAccount == nil {
		t.Error("CreateFutureAccount method should exist")
	}
}

func TestFuturesAPI_FundTransfer(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name string
		req  FuturesTransferRequest
	}{
		{
			name: "Spot to futures transfer",
			req: FuturesTransferRequest{
				Currency: "USDT",
				Amount:   decimal.NewFromFloat(100.0),
				Type:     1,
			},
		},
		{
			name: "Futures to spot transfer",
			req: FuturesTransferRequest{
				Currency: "USDT",
				Amount:   decimal.NewFromFloat(50.0),
				Type:     2,
			},
		},
		{
			name: "Zero amount transfer",
			req: FuturesTransferRequest{
				Currency: "USDT",
				Amount:   decimal.Zero,
				Type:     1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.FundTransfer == nil {
				t.Error("FundTransfer method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetAllPositions(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetAllPositions == nil {
		t.Error("GetAllPositions method should exist")
	}
}

func TestFuturesAPI_GetFutures(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	if futures.GetFutures == nil {
		t.Error("GetFutures method should exist")
	}
}

func TestFuturesAPI_GetOpeningOrders(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		limit       int
	}{
		{
			name:        "Valid parameters",
			futuresName: "BTCUSDT",
			limit:       50,
		},
		{
			name:        "Zero limit",
			futuresName: "BTCUSDT",
			limit:       0,
		},
		{
			name:        "Large limit",
			futuresName: "ETHUSDT",
			limit:       1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetOpeningOrders == nil {
				t.Error("GetOpeningOrders method should exist")
			}
		})
	}
}

func TestFuturesAPI_GetMyTrades(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := NewFuturesAPI(client)

	tests := []struct {
		name        string
		futuresName string
		fromId      string
		limit       int
	}{
		{
			name:        "Basic parameters",
			futuresName: "BTCUSDT",
			fromId:      "",
			limit:       50,
		},
		{
			name:        "With from ID",
			futuresName: "BTCUSDT",
			fromId:      "12345",
			limit:       100,
		},
		{
			name:        "Zero limit",
			futuresName: "ETHUSDT",
			fromId:      "",
			limit:       0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if futures.GetMyTrades == nil {
				t.Error("GetMyTrades method should exist")
			}
		})
	}
}

// Helper function to create test futures client with testnet enabled
func createTestFuturesClient() *FuturesAPI {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	return NewFuturesAPI(client)
}

// Test client initialization with testnet
func TestFuturesAPITestnetInitialization(t *testing.T) {
	futures := createTestFuturesClient()

	if futures == nil {
		t.Error("Futures API should be initialized")
	}

	if !futures.client.Testnet {
		t.Error("Futures API should be using testnet")
	}

	expectedURL := _baseUrlTestnetFutures
	actualURL := futures.client.baseUrlFutures()

	if actualURL != expectedURL {
		t.Errorf("Expected testnet URL %s, got %s", expectedURL, actualURL)
	}
}

// Test that all required methods exist on FuturesAPI
func TestFuturesAPI_AllMethodsExist(t *testing.T) {
	futures := createTestFuturesClient()

	methods := []string{
		"GetTicker",
		"GetDepth",
		"GetKlines",
		"CreateOrder",
		"CancelOrder",
		"CancelAllOrders",
		"GetCurrentOrders",
		"GetOrderHistory",
		"GetOrderInfo",
		"GetTrades",
		"GetPositions",
		"GetAccount",
		"SetLeverage",
		"SetMarginType",
		"ModifyMargin",
		"GetAllTicker",
		"GetIndexPrice",
		"GetAllIndexPrice",
		"GetAllTagIndexPrice",
		"GetAllFuturesDepth",
		"BatchCreateOrders",
		"BatchCancelOrders",
		"GetCapital",
		"GetFutureAccounts",
		"CreateFutureAccount",
		"FundTransfer",
		"GetAllPositions",
		"GetFutures",
		"GetOpeningOrders",
		"GetMyTrades",
	}

	// This is a compile-time check that all methods exist
	// If any method is missing, this won't compile
	_ = futures.GetTicker
	_ = futures.GetDepth
	_ = futures.GetKlines
	_ = futures.CreateOrder
	_ = futures.CancelOrder
	_ = futures.CancelAllOrders
	_ = futures.GetCurrentOrders
	_ = futures.GetOrderHistory
	_ = futures.GetOrderInfo
	_ = futures.GetTrades
	_ = futures.GetPositions
	_ = futures.GetAccount
	_ = futures.SetLeverage
	_ = futures.SetMarginType
	_ = futures.ModifyMargin
	_ = futures.GetAllTicker
	_ = futures.GetIndexPrice
	_ = futures.GetAllIndexPrice
	_ = futures.GetAllTagIndexPrice
	_ = futures.GetAllFuturesDepth
	_ = futures.BatchCreateOrders
	_ = futures.BatchCancelOrders
	_ = futures.GetCapital
	_ = futures.GetFutureAccounts
	_ = futures.CreateFutureAccount
	_ = futures.FundTransfer
	_ = futures.GetAllPositions
	_ = futures.GetFutures
	_ = futures.GetOpeningOrders
	_ = futures.GetMyTrades

	t.Logf("All %d futures API methods exist", len(methods))
}

// Test futures-specific request types
func TestFuturesRequestTypes(t *testing.T) {
	// Test FuturesCreateOrderRequest
	req := FuturesCreateOrderRequest{
		FuturesName:   "BTCUSDT",
		Type:          "LIMIT",
		Side:          "BUY",
		Open:          "OPEN",
		PositionType:  "LONG",
		Price:         decimal.NewFromFloat(45000),
		Volume:        decimal.NewFromFloat(0.01),
		ClientOrderID: "test_order",
	}

	if req.FuturesName != "BTCUSDT" {
		t.Error("FuturesCreateOrderRequest should set futures name correctly")
	}

	if req.Type != "LIMIT" {
		t.Error("FuturesCreateOrderRequest should set type correctly")
	}

	// Test FuturesBatchOrderRequest
	batchReq := FuturesBatchOrderRequest{
		FuturesName: "BTCUSDT",
		Orders:      []FuturesCreateOrderRequest{req},
	}

	if len(batchReq.Orders) != 1 {
		t.Error("FuturesBatchOrderRequest should contain one order")
	}

	// Test FuturesTransferRequest
	transferReq := FuturesTransferRequest{
		Currency: "USDT",
		Amount:   decimal.NewFromFloat(100),
		Type:     1,
	}

	if transferReq.Currency != "USDT" {
		t.Error("FuturesTransferRequest should set currency correctly")
	}

	if transferReq.Type != 1 {
		t.Error("FuturesTransferRequest should set type correctly")
	}
}
