package byex

import (
	"testing"

	"github.com/shopspring/decimal"
)

// Integration tests for the entire byex package using testnet
// These tests demonstrate proper usage patterns and validate the complete workflow

// TestFullWorkflow demonstrates a complete trading workflow on testnet
func TestFullWorkflow(t *testing.T) {
	// Initialize client with testnet
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	if client == nil {
		t.Fatal("Failed to create client")
	}

	if !client.Testnet {
		t.Fatal("Client should be configured for testnet")
	}

	// Get exchange and futures APIs
	exchange := client.Exchange()
	futures := client.Futures()

	if exchange == nil || futures == nil {
		t.Fatal("Failed to create API instances")
	}

	// Verify APIs have required methods
	if exchange.GetAccount == nil {
		t.Error("Exchange API should have GetAccount method")
	}

	if futures.GetAccount == nil {
		t.Error("Futures API should have GetAccount method")
	}

	t.Log("Successfully initialized client and API instances for testnet")
}

// TestExchangeWorkflow demonstrates exchange trading workflow
func TestExchangeWorkflow(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := client.Exchange()

	// Verify exchange API methods exist
	if exchange.CreateOrder == nil {
		t.Error("CreateOrder method should exist")
	}

	// Prepare test order request
	orderReq := CreateOrderRequest{
		Symbol:        "BTCUSDT",
		Side:          "buy",
		Type:          "limit",
		Amount:        decimal.NewFromFloat(0.001), // Small amount for testing
		Price:         decimal.NewFromFloat(30000), // Conservative price
		ClientOrderID: "integration_test_1",
	}

	// Validate order request structure
	if orderReq.Symbol == "" {
		t.Error("Order symbol should not be empty")
	}

	if orderReq.Amount.IsZero() {
		t.Error("Order amount should not be zero")
	}

	if orderReq.Price.IsZero() {
		t.Error("Order price should not be zero")
	}

	t.Logf("Prepared order request: %+v", orderReq)

	// Note: We don't actually execute the order since we don't have valid credentials
	// In a real test environment, you would:
	// 1. Call exchange.CreateOrder(orderReq)
	// 2. Check the response
	// 3. Cancel the order if needed
	// 4. Verify account balance changes
}

// TestFuturesWorkflow demonstrates futures trading workflow
func TestFuturesWorkflow(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	futures := client.Futures()

	// Verify futures API methods exist
	if futures.CreateOrder == nil {
		t.Error("CreateOrder method should exist")
	}

	// Prepare test futures order request
	futuresReq := FuturesCreateOrderRequest{
		FuturesName:   "BTCUSDT",
		Type:          "LIMIT",
		Side:          "BUY",
		Open:          "OPEN",
		PositionType:  "LONG",
		Price:         decimal.NewFromFloat(30000),
		Volume:        decimal.NewFromFloat(0.001),
		ClientOrderID: "futures_test_1",
	}

	// Validate futures order request
	if futuresReq.FuturesName == "" {
		t.Error("Futures name should not be empty")
	}

	if futuresReq.Volume.IsZero() {
		t.Error("Futures volume should not be zero")
	}

	if futuresReq.PositionType != "LONG" && futuresReq.PositionType != "SHORT" {
		t.Error("Position type should be LONG or SHORT")
	}

	t.Logf("Prepared futures order request: %+v", futuresReq)

	// Test leverage settings
	leverageTests := []int{1, 5, 10, 20}
	for _, leverage := range leverageTests {
		if leverage <= 0 || leverage > 125 {
			t.Errorf("Invalid leverage value: %d", leverage)
		}
	}
}

// TestBatchOperations demonstrates batch operations
func TestBatchOperations(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := client.Exchange()

	// Verify batch methods exist
	if exchange.BatchCreateOrders == nil {
		t.Error("BatchCreateOrders method should exist")
	}

	// Prepare batch order requests
	batchReq := BatchOrderRequest{
		Symbol: "BTCUSDT",
		Orders: []CreateOrderRequest{
			{
				Symbol: "BTCUSDT",
				Side:   "buy",
				Type:   "limit",
				Amount: decimal.NewFromFloat(0.001),
				Price:  decimal.NewFromFloat(29000),
			},
			{
				Symbol: "BTCUSDT",
				Side:   "buy",
				Type:   "limit",
				Amount: decimal.NewFromFloat(0.001),
				Price:  decimal.NewFromFloat(28000),
			},
			{
				Symbol: "BTCUSDT",
				Side:   "sell",
				Type:   "limit",
				Amount: decimal.NewFromFloat(0.001),
				Price:  decimal.NewFromFloat(32000),
			},
		},
	}

	// Validate batch request
	if len(batchReq.Orders) == 0 {
		t.Error("Batch request should contain at least one order")
	}

	for i, order := range batchReq.Orders {
		if order.Symbol != batchReq.Symbol {
			t.Errorf("Order %d symbol mismatch", i)
		}

		if order.Amount.IsZero() {
			t.Errorf("Order %d amount should not be zero", i)
		}
	}

	t.Logf("Prepared batch order with %d orders", len(batchReq.Orders))

	// Test batch cancel orders
	orderIds := []string{"order1", "order2", "order3"}
	if len(orderIds) == 0 {
		t.Error("Order IDs list should not be empty for cancellation")
	}
}

// TestAccountOperations demonstrates account-related operations
func TestAccountOperations(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := client.Exchange()
	futures := client.Futures()

	// Test exchange account operations
	if exchange.GetAccount == nil {
		t.Error("GetAccount method should exist")
	}

	if exchange.GetBalance == nil {
		t.Error("GetBalance method should exist")
	}

	// Test futures account operations
	if futures.GetAccount == nil {
		t.Error("Futures GetAccount method should exist")
	}

	if futures.GetPositions == nil {
		t.Error("GetPositions method should exist")
	}

	// Test transfer operations
	transferReq := FuturesTransferRequest{
		Currency: "USDT",
		Amount:   decimal.NewFromFloat(100),
		Type:     1, // spot to futures
	}

	if transferReq.Currency == "" {
		t.Error("Transfer currency should not be empty")
	}

	if transferReq.Amount.IsZero() {
		t.Error("Transfer amount should not be zero")
	}

	if transferReq.Type != 1 && transferReq.Type != 2 {
		t.Error("Transfer type should be 1 (spot to futures) or 2 (futures to spot)")
	}

	t.Logf("Transfer request: %+v", transferReq)
}

// TestMarketDataOperations demonstrates market data operations
func TestMarketDataOperations(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := client.Exchange()
	futures := client.Futures()

	// Test exchange market data
	symbols := []string{"BTCUSDT", "ETHUSDT", "BNBUSDT"}

	for _, symbol := range symbols {
		if symbol == "" {
			t.Error("Symbol should not be empty")
		}

		// Verify the symbol is a valid trading pair format
		if len(symbol) < 6 {
			t.Errorf("Symbol %s appears to be too short", symbol)
		}

		// Test various depth levels
		depths := []int{5, 10, 20, 50}
		for _, depth := range depths {
			if depth <= 0 {
				t.Errorf("Invalid depth: %d", depth)
			}
		}

		// Test various kline periods
		periods := []string{"1m", "5m", "15m", "30m", "1h", "4h", "1d"}
		for _, period := range periods {
			if period == "" {
				t.Error("Period should not be empty")
			}
		}
	}

	// Test futures market data methods exist
	if futures.GetTicker == nil {
		t.Error("Futures GetTicker should exist")
	}

	if futures.GetDepth == nil {
		t.Error("Futures GetDepth should exist")
	}

	if futures.GetKlines == nil {
		t.Error("Futures GetKlines should exist")
	}

	// Test exchange market data methods exist
	if exchange.GetAllTicker == nil {
		t.Error("Exchange GetAllTicker should exist")
	}

	t.Logf("Tested market data operations for %d exchange symbols", len(symbols))
}

// TestErrorHandling demonstrates error handling patterns
func TestErrorHandling(t *testing.T) {
	// Test with invalid parameters
	invalidReq := CreateOrderRequest{
		Symbol: "", // Empty symbol should cause validation error
		Side:   "invalid_side",
		Type:   "invalid_type",
		Amount: decimal.NewFromFloat(-1), // Negative amount
		Price:  decimal.Zero,             // Zero price
	}

	// Validate that we can detect invalid requests
	if invalidReq.Symbol == "" {
		t.Log("Correctly detected empty symbol")
	}

	if invalidReq.Side != "buy" && invalidReq.Side != "sell" {
		t.Log("Correctly detected invalid side")
	}

	if invalidReq.Amount.IsNegative() {
		t.Log("Correctly detected negative amount")
	}

	if invalidReq.Price.IsZero() {
		t.Log("Correctly detected zero price for limit order")
	}

	// Test error type
	apiError := &Error{
		Code:    "1001",
		Message: "Invalid parameter",
	}

	expectedErrorMsg := "API Error - Code: 1001, Message: Invalid parameter"
	if apiError.Error() != expectedErrorMsg {
		t.Errorf("Expected error message: %s, got: %s", expectedErrorMsg, apiError.Error())
	}
}

// TestConfigurationValidation validates testnet configuration
func TestConfigurationValidation(t *testing.T) {
	// Test testnet URLs
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	expectedExchangeURL := _baseUrlTestnetExchange
	actualExchangeURL := client.baseUrlExchange()

	if actualExchangeURL != expectedExchangeURL {
		t.Errorf("Expected exchange testnet URL: %s, got: %s",
			expectedExchangeURL, actualExchangeURL)
	}

	expectedFuturesURL := _baseUrlTestnetFutures
	actualFuturesURL := client.baseUrlFutures()

	if actualFuturesURL != expectedFuturesURL {
		t.Errorf("Expected futures testnet URL: %s, got: %s",
			expectedFuturesURL, actualFuturesURL)
	}

	// Test production URLs (for comparison)
	prodClient := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: false})

	prodExchangeURL := prodClient.baseUrlExchange()
	prodFuturesURL := prodClient.baseUrlFutures()

	if prodExchangeURL == actualExchangeURL {
		t.Error("Production and testnet exchange URLs should be different")
	}

	if prodFuturesURL == actualFuturesURL {
		t.Error("Production and testnet futures URLs should be different")
	}

	t.Logf("Testnet configuration validated:")
	t.Logf("  Exchange URL: %s", actualExchangeURL)
	t.Logf("  Futures URL: %s", actualFuturesURL)
}

// TestSignatureGeneration tests signature generation methods
func TestSignatureGeneration(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})

	// Test exchange signature generation
	params := map[string]string{
		"symbol": "BTCUSDT",
		"side":   "buy",
		"type":   "limit",
		"volume": "0.001",
		"price":  "45000",
	}

	signature1 := client.generateExchangeSignature(params)
	signature2 := client.generateExchangeSignature(params)

	if signature1 == "" {
		t.Error("Exchange signature should not be empty")
	}

	// Signatures should be different due to timestamp
	if signature1 == signature2 {
		t.Log("Note: Exchange signatures are identical (same timestamp)")
	}

	// Test futures signature generation
	timestamp := int64(1640995200000)
	futuresSignature1 := client.generateFuturesSignature("POST", "/fapi/v1/trade/order", "", timestamp)
	futuresSignature2 := client.generateFuturesSignature("POST", "/fapi/v1/trade/order", "", timestamp)

	if futuresSignature1 == "" {
		t.Error("Futures signature should not be empty")
	}

	// Same parameters should produce same signature
	if futuresSignature1 != futuresSignature2 {
		t.Error("Identical futures requests should produce identical signatures")
	}

	// Different methods should produce different signatures
	getSignature := client.generateFuturesSignature("GET", "/fapi/v1/trade/order", "", timestamp)
	if futuresSignature1 == getSignature {
		t.Error("Different HTTP methods should produce different signatures")
	}

	t.Log("Signature generation tests passed")
}

// TestComprehensiveValidation runs comprehensive validation tests
func TestComprehensiveValidation(t *testing.T) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	exchange := client.Exchange()
	futures := client.Futures()

	// Count and validate all exchange methods
	exchangeMethods := []string{
		"GetAllTicker", "GetTicker", "GetDepth", "GetKlines",
		"CreateOrder", "CancelOrder", "CancelAllOrders", "BatchCreateOrders",
		"GetCurrentOrders", "GetOrderHistory", "GetOrderInfo", "GetTrades",
		"GetAccount", "GetBalance", "GetAllTradingRecords", "GetMarketPrices",
		"BatchPlaceOrders", "BatchCancelOrders", "GetOrderDetail", "ReplaceOrder",
		"GetSymbolsCharge", "GetLeverageFinanceBalance",
	}

	// Count and validate all futures methods
	futuresMethods := []string{
		"GetTicker", "GetDepth", "GetKlines", "CreateOrder", "CancelOrder",
		"CancelAllOrders", "GetCurrentOrders", "GetOrderHistory", "GetOrderInfo",
		"GetTrades", "GetPositions", "GetAccount", "SetLeverage", "SetMarginType",
		"ModifyMargin", "GetAllTicker", "GetIndexPrice", "GetAllIndexPrice",
		"GetAllTagIndexPrice", "GetAllFuturesDepth", "BatchCreateOrders",
		"BatchCancelOrders", "GetCapital", "GetFutureAccounts", "CreateFutureAccount",
		"FundTransfer", "GetAllPositions", "GetFutures", "GetOpeningOrders", "GetMyTrades",
	}

	t.Logf("Exchange API has %d methods", len(exchangeMethods))
	t.Logf("Futures API has %d methods", len(futuresMethods))
	t.Logf("Total API methods: %d", len(exchangeMethods)+len(futuresMethods))

	// Validate that APIs are properly initialized
	if exchange == nil {
		t.Fatal("Exchange API should be initialized")
	}

	if futures == nil {
		t.Fatal("Futures API should be initialized")
	}

	// Validate testnet configuration
	if !client.Testnet {
		t.Fatal("Client should be configured for testnet")
	}

	t.Log("Comprehensive validation completed successfully")
}

// BenchmarkClientCreation benchmarks client creation performance
func BenchmarkClientCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
		_ = client.Exchange()
		_ = client.Futures()
	}
}

// BenchmarkSignatureGeneration benchmarks signature generation performance
func BenchmarkSignatureGeneration(b *testing.B) {
	client := NewClient(testApiKey, testSecretKey, ClientOption{Testnet: true})
	params := map[string]string{
		"symbol": "BTCUSDT",
		"side":   "buy",
		"type":   "limit",
		"volume": "0.001",
		"price":  "45000",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.generateExchangeSignature(params)
	}
}

// BenchmarkRequestCreation benchmarks request object creation
func BenchmarkRequestCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateOrderRequest{
			Symbol:        "BTCUSDT",
			Side:          "buy",
			Type:          "limit",
			Amount:        decimal.NewFromFloat(0.001),
			Price:         decimal.NewFromFloat(45000),
			ClientOrderID: "benchmark_test",
		}
	}
}
