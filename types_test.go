package byex

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
)

// Test ExchangeOrder type
func TestExchangeOrder(t *testing.T) {
	order := ExchangeOrder{
		ID:               "12345",
		Symbol:           "BTCUSDT",
		Type:             "limit",
		Side:             "buy",
		Amount:           decimal.NewFromFloat(0.01),
		Price:            decimal.NewFromFloat(45000),
		Status:           "filled",
		CreatedAt:        1640995200000,
		UpdatedAt:        1640995300000,
		FinishedAt:       1640995400000,
		CancelledAt:      0,
		AvgPrice:         decimal.NewFromFloat(45100),
		Source:           "api",
		Fee:              decimal.NewFromFloat(0.1),
		FeeCurrency:      "USDT",
		FilledAmount:     decimal.NewFromFloat(0.01),
		FilledCashAmount: decimal.NewFromFloat(451),
		FilledFees:       decimal.NewFromFloat(0.1),
	}

	// Test JSON marshaling/unmarshaling
	jsonData, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeOrder: %v", err)
	}

	var unmarshaledOrder ExchangeOrder
	err = json.Unmarshal(jsonData, &unmarshaledOrder)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeOrder: %v", err)
	}

	if unmarshaledOrder.ID != order.ID {
		t.Errorf("Expected ID %s, got %s", order.ID, unmarshaledOrder.ID)
	}

	if unmarshaledOrder.Symbol != order.Symbol {
		t.Errorf("Expected Symbol %s, got %s", order.Symbol, unmarshaledOrder.Symbol)
	}
}

// Test ExchangeTrade type
func TestExchangeTrade(t *testing.T) {
	trade := ExchangeTrade{
		ID:          "trade123",
		OrderID:     "order123",
		Symbol:      "BTCUSDT",
		Side:        "buy",
		Amount:      decimal.NewFromFloat(0.01),
		Price:       decimal.NewFromFloat(45000),
		Fee:         decimal.NewFromFloat(0.05),
		FeeCurrency: "USDT",
		Role:        "taker",
		CreatedAt:   1640995200000,
	}

	jsonData, err := json.Marshal(trade)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeTrade: %v", err)
	}

	var unmarshaledTrade ExchangeTrade
	err = json.Unmarshal(jsonData, &unmarshaledTrade)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeTrade: %v", err)
	}

	if unmarshaledTrade.ID != trade.ID {
		t.Errorf("Expected ID %s, got %s", trade.ID, unmarshaledTrade.ID)
	}
}

// Test ExchangeTicker type
func TestExchangeTicker(t *testing.T) {
	ticker := ExchangeTicker{
		Symbol:      "BTCUSDT",
		High:        decimal.NewFromFloat(50000),
		Low:         decimal.NewFromFloat(45000),
		Last:        decimal.NewFromFloat(48000),
		Vol:         decimal.NewFromFloat(1000),
		Amount:      decimal.NewFromFloat(48000000),
		BuyPrice:    decimal.NewFromFloat(47950),
		SellPrice:   decimal.NewFromFloat(48050),
		NewCoinFlag: 0,
		Change:      decimal.NewFromFloat(1000),
		Rose:        decimal.NewFromFloat(0.02),
	}

	jsonData, err := json.Marshal(ticker)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeTicker: %v", err)
	}

	var unmarshaledTicker ExchangeTicker
	err = json.Unmarshal(jsonData, &unmarshaledTicker)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeTicker: %v", err)
	}

	if unmarshaledTicker.Symbol != ticker.Symbol {
		t.Errorf("Expected Symbol %s, got %s", ticker.Symbol, unmarshaledTicker.Symbol)
	}
}

// Test ExchangeDepth type
func TestExchangeDepth(t *testing.T) {
	depth := ExchangeDepth{
		Asks: [][]decimal.Decimal{
			{decimal.NewFromFloat(48000), decimal.NewFromFloat(0.1)},
			{decimal.NewFromFloat(48100), decimal.NewFromFloat(0.2)},
		},
		Bids: [][]decimal.Decimal{
			{decimal.NewFromFloat(47900), decimal.NewFromFloat(0.15)},
			{decimal.NewFromFloat(47800), decimal.NewFromFloat(0.25)},
		},
	}

	jsonData, err := json.Marshal(depth)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeDepth: %v", err)
	}

	var unmarshaledDepth ExchangeDepth
	err = json.Unmarshal(jsonData, &unmarshaledDepth)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeDepth: %v", err)
	}

	if len(unmarshaledDepth.Asks) != len(depth.Asks) {
		t.Errorf("Expected %d asks, got %d", len(depth.Asks), len(unmarshaledDepth.Asks))
	}

	if len(unmarshaledDepth.Bids) != len(depth.Bids) {
		t.Errorf("Expected %d bids, got %d", len(depth.Bids), len(unmarshaledDepth.Bids))
	}
}

// Test ExchangeKline type
func TestExchangeKline(t *testing.T) {
	kline := ExchangeKline{
		Time:   1640995200000,
		Open:   decimal.NewFromFloat(47000),
		High:   decimal.NewFromFloat(48000),
		Low:    decimal.NewFromFloat(46500),
		Close:  decimal.NewFromFloat(47800),
		Volume: decimal.NewFromFloat(100),
	}

	jsonData, err := json.Marshal(kline)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeKline: %v", err)
	}

	var unmarshaledKline ExchangeKline
	err = json.Unmarshal(jsonData, &unmarshaledKline)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeKline: %v", err)
	}

	if unmarshaledKline.Time != kline.Time {
		t.Errorf("Expected Time %d, got %d", kline.Time, unmarshaledKline.Time)
	}
}

// Test ExchangeAccount type
func TestExchangeAccount(t *testing.T) {
	account := ExchangeAccount{
		TotalAsset: decimal.NewFromFloat(10000),
		CoinList: []CoinBalance{
			{
				Coin:     "BTC",
				Normal:   decimal.NewFromFloat(0.1),
				Locked:   decimal.NewFromFloat(0.05),
				BtcValue: decimal.NewFromFloat(0.15),
				RmbValue: decimal.NewFromFloat(7500),
			},
			{
				Coin:     "USDT",
				Normal:   decimal.NewFromFloat(2500),
				Locked:   decimal.NewFromFloat(500),
				BtcValue: decimal.NewFromFloat(0.05),
				RmbValue: decimal.NewFromFloat(3000),
			},
		},
		NormalCount:   decimal.NewFromFloat(5000),
		LockedCount:   decimal.NewFromFloat(1000),
		FreezingCount: decimal.NewFromFloat(0),
		BtcValuation:  decimal.NewFromFloat(0.2),
		RmbValuation:  decimal.NewFromFloat(10000),
	}

	jsonData, err := json.Marshal(account)
	if err != nil {
		t.Errorf("Failed to marshal ExchangeAccount: %v", err)
	}

	var unmarshaledAccount ExchangeAccount
	err = json.Unmarshal(jsonData, &unmarshaledAccount)
	if err != nil {
		t.Errorf("Failed to unmarshal ExchangeAccount: %v", err)
	}

	if len(unmarshaledAccount.CoinList) != len(account.CoinList) {
		t.Errorf("Expected %d coins, got %d", len(account.CoinList), len(unmarshaledAccount.CoinList))
	}
}

// Test CoinBalance type
func TestCoinBalance(t *testing.T) {
	balance := CoinBalance{
		Coin:     "BTC",
		Normal:   decimal.NewFromFloat(0.1),
		Locked:   decimal.NewFromFloat(0.05),
		BtcValue: decimal.NewFromFloat(0.15),
		RmbValue: decimal.NewFromFloat(7500),
	}

	jsonData, err := json.Marshal(balance)
	if err != nil {
		t.Errorf("Failed to marshal CoinBalance: %v", err)
	}

	var unmarshaledBalance CoinBalance
	err = json.Unmarshal(jsonData, &unmarshaledBalance)
	if err != nil {
		t.Errorf("Failed to unmarshal CoinBalance: %v", err)
	}

	if unmarshaledBalance.Coin != balance.Coin {
		t.Errorf("Expected Coin %s, got %s", balance.Coin, unmarshaledBalance.Coin)
	}
}

// Test Futures types
func TestFuturesOrder(t *testing.T) {
	order := FuturesOrder{
		OrderID:       "futures123",
		ClientOrderID: "client123",
		Symbol:        "BTCUSDT",
		Type:          "LIMIT",
		Side:          "BUY",
		Open:          "OPEN",
		PositionType:  "LONG",
		Price:         decimal.NewFromFloat(45000),
		Volume:        decimal.NewFromFloat(0.01),
		Status:        "NEW",
		CreatedAt:     1640995200000,
		UpdatedAt:     1640995300000,
	}

	jsonData, err := json.Marshal(order)
	if err != nil {
		t.Errorf("Failed to marshal FuturesOrder: %v", err)
	}

	var unmarshaledOrder FuturesOrder
	err = json.Unmarshal(jsonData, &unmarshaledOrder)
	if err != nil {
		t.Errorf("Failed to unmarshal FuturesOrder: %v", err)
	}

	if unmarshaledOrder.OrderID != order.OrderID {
		t.Errorf("Expected OrderID %s, got %s", order.OrderID, unmarshaledOrder.OrderID)
	}
}

// Test FuturesPosition type
func TestFuturesPosition(t *testing.T) {
	position := FuturesPosition{
		Symbol:            "BTCUSDT",
		PositionSide:      "LONG",
		PositionAmt:       decimal.NewFromFloat(0.01),
		AvgPrice:          decimal.NewFromFloat(45000),
		UnrealizedPnl:     decimal.NewFromFloat(100),
		RealizedPnl:       decimal.NewFromFloat(50),
		MarginType:        "ISOLATED",
		InitialMargin:     decimal.NewFromFloat(450),
		MaintenanceMargin: decimal.NewFromFloat(45),
		PositionValue:     decimal.NewFromFloat(450),
		Leverage:          decimal.NewFromInt(10),
	}

	jsonData, err := json.Marshal(position)
	if err != nil {
		t.Errorf("Failed to marshal FuturesPosition: %v", err)
	}

	var unmarshaledPosition FuturesPosition
	err = json.Unmarshal(jsonData, &unmarshaledPosition)
	if err != nil {
		t.Errorf("Failed to unmarshal FuturesPosition: %v", err)
	}

	if unmarshaledPosition.Symbol != position.Symbol {
		t.Errorf("Expected Symbol %s, got %s", position.Symbol, unmarshaledPosition.Symbol)
	}
}

// Test CreateOrderRequest type
func TestCreateOrderRequest(t *testing.T) {
	req := CreateOrderRequest{
		Symbol:        "BTCUSDT",
		Side:          "buy",
		Type:          "limit",
		Amount:        decimal.NewFromFloat(0.01),
		Price:         decimal.NewFromFloat(45000),
		ClientOrderID: "test123",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal CreateOrderRequest: %v", err)
	}

	var unmarshaledReq CreateOrderRequest
	err = json.Unmarshal(jsonData, &unmarshaledReq)
	if err != nil {
		t.Errorf("Failed to unmarshal CreateOrderRequest: %v", err)
	}

	if unmarshaledReq.Symbol != req.Symbol {
		t.Errorf("Expected Symbol %s, got %s", req.Symbol, unmarshaledReq.Symbol)
	}
}

// Test FuturesCreateOrderRequest type
func TestFuturesCreateOrderRequest(t *testing.T) {
	req := FuturesCreateOrderRequest{
		FuturesName:   "BTCUSDT",
		Type:          "LIMIT",
		Side:          "BUY",
		Open:          "OPEN",
		PositionType:  "LONG",
		Price:         decimal.NewFromFloat(45000),
		Volume:        decimal.NewFromFloat(0.01),
		ClientOrderID: "futures_test123",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal FuturesCreateOrderRequest: %v", err)
	}

	var unmarshaledReq FuturesCreateOrderRequest
	err = json.Unmarshal(jsonData, &unmarshaledReq)
	if err != nil {
		t.Errorf("Failed to unmarshal FuturesCreateOrderRequest: %v", err)
	}

	if unmarshaledReq.FuturesName != req.FuturesName {
		t.Errorf("Expected FuturesName %s, got %s", req.FuturesName, unmarshaledReq.FuturesName)
	}
}

// Test BatchOrder type
func TestBatchOrder(t *testing.T) {
	batchOrder := BatchOrder{
		Volume:        decimal.NewFromFloat(0.01),
		Price:         decimal.NewFromFloat(45000),
		Side:          "buy",
		Type:          1,
		VolumeType:    1,
		ClientOrderID: "batch123",
	}

	jsonData, err := json.Marshal(batchOrder)
	if err != nil {
		t.Errorf("Failed to marshal BatchOrder: %v", err)
	}

	var unmarshaledOrder BatchOrder
	err = json.Unmarshal(jsonData, &unmarshaledOrder)
	if err != nil {
		t.Errorf("Failed to unmarshal BatchOrder: %v", err)
	}

	if unmarshaledOrder.Side != batchOrder.Side {
		t.Errorf("Expected Side %s, got %s", batchOrder.Side, unmarshaledOrder.Side)
	}

	if unmarshaledOrder.Type != batchOrder.Type {
		t.Errorf("Expected Type %d, got %d", batchOrder.Type, unmarshaledOrder.Type)
	}
}

// Test OrderResponse type
func TestOrderResponse(t *testing.T) {
	response := OrderResponse{
		OrderID: "response123",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal OrderResponse: %v", err)
	}

	var unmarshaledResponse OrderResponse
	err = json.Unmarshal(jsonData, &unmarshaledResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal OrderResponse: %v", err)
	}

	if unmarshaledResponse.OrderID != response.OrderID {
		t.Errorf("Expected OrderID %s, got %s", response.OrderID, unmarshaledResponse.OrderID)
	}
}

// Test OrderListResponse type
func TestOrderListResponse(t *testing.T) {
	response := OrderListResponse{
		Count: 2,
		ResultList: []ExchangeOrder{
			{
				ID:     "order1",
				Symbol: "BTCUSDT",
				Type:   "limit",
				Side:   "buy",
				Amount: decimal.NewFromFloat(0.01),
				Price:  decimal.NewFromFloat(45000),
			},
			{
				ID:     "order2",
				Symbol: "ETHUSDT",
				Type:   "market",
				Side:   "sell",
				Amount: decimal.NewFromFloat(0.1),
				Price:  decimal.NewFromFloat(3000),
			},
		},
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		t.Errorf("Failed to marshal OrderListResponse: %v", err)
	}

	var unmarshaledResponse OrderListResponse
	err = json.Unmarshal(jsonData, &unmarshaledResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal OrderListResponse: %v", err)
	}

	if unmarshaledResponse.Count != response.Count {
		t.Errorf("Expected Count %d, got %d", response.Count, unmarshaledResponse.Count)
	}

	if len(unmarshaledResponse.ResultList) != len(response.ResultList) {
		t.Errorf("Expected %d orders, got %d", len(response.ResultList), len(unmarshaledResponse.ResultList))
	}
}

// Test SymbolCharge type
func TestSymbolCharge(t *testing.T) {
	charge := SymbolCharge{
		Symbol:              "BTCUSDT",
		BaseAsset:           "BTC",
		QuoteAsset:          "USDT",
		BaseAssetPrecision:  8,
		QuoteAssetPrecision: 2,
		Status:              "TRADING",
		TakerCommission:     decimal.NewFromFloat(0.001),
		MakerCommission:     decimal.NewFromFloat(0.0008),
		MinPrice:            decimal.NewFromFloat(0.01),
		MaxPrice:            decimal.NewFromFloat(1000000),
		TickSize:            decimal.NewFromFloat(0.01),
		MinQty:              decimal.NewFromFloat(0.00001),
		MaxQty:              decimal.NewFromFloat(9000),
		StepSize:            decimal.NewFromFloat(0.00001),
	}

	jsonData, err := json.Marshal(charge)
	if err != nil {
		t.Errorf("Failed to marshal SymbolCharge: %v", err)
	}

	var unmarshaledCharge SymbolCharge
	err = json.Unmarshal(jsonData, &unmarshaledCharge)
	if err != nil {
		t.Errorf("Failed to unmarshal SymbolCharge: %v", err)
	}

	if unmarshaledCharge.Symbol != charge.Symbol {
		t.Errorf("Expected Symbol %s, got %s", charge.Symbol, unmarshaledCharge.Symbol)
	}

	if unmarshaledCharge.BaseAssetPrecision != charge.BaseAssetPrecision {
		t.Errorf("Expected BaseAssetPrecision %d, got %d", charge.BaseAssetPrecision, unmarshaledCharge.BaseAssetPrecision)
	}
}

// Test LeverageFinanceBalance type
func TestLeverageFinanceBalance(t *testing.T) {
	balance := LeverageFinanceBalance{
		Symbol:        "BTCUSDT",
		BaseAsset:     "BTC",
		QuoteAsset:    "USDT",
		BaseBalance:   decimal.NewFromFloat(0.1),
		QuoteBalance:  decimal.NewFromFloat(1000),
		BaseBorrowed:  decimal.NewFromFloat(0.05),
		QuoteBorrowed: decimal.NewFromFloat(500),
		BaseInterest:  decimal.NewFromFloat(0.001),
		QuoteInterest: decimal.NewFromFloat(1),
		BaseNetAsset:  decimal.NewFromFloat(0.049),
		QuoteNetAsset: decimal.NewFromFloat(499),
		MarginLevel:   decimal.NewFromFloat(2.5),
	}

	jsonData, err := json.Marshal(balance)
	if err != nil {
		t.Errorf("Failed to marshal LeverageFinanceBalance: %v", err)
	}

	var unmarshaledBalance LeverageFinanceBalance
	err = json.Unmarshal(jsonData, &unmarshaledBalance)
	if err != nil {
		t.Errorf("Failed to unmarshal LeverageFinanceBalance: %v", err)
	}

	if unmarshaledBalance.Symbol != balance.Symbol {
		t.Errorf("Expected Symbol %s, got %s", balance.Symbol, unmarshaledBalance.Symbol)
	}
}

// Test FuturesTransferRequest type
func TestFuturesTransferRequest(t *testing.T) {
	req := FuturesTransferRequest{
		Currency: "USDT",
		Amount:   decimal.NewFromFloat(100),
		Type:     1, // spot to futures
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Errorf("Failed to marshal FuturesTransferRequest: %v", err)
	}

	var unmarshaledReq FuturesTransferRequest
	err = json.Unmarshal(jsonData, &unmarshaledReq)
	if err != nil {
		t.Errorf("Failed to unmarshal FuturesTransferRequest: %v", err)
	}

	if unmarshaledReq.Currency != req.Currency {
		t.Errorf("Expected Currency %s, got %s", req.Currency, unmarshaledReq.Currency)
	}

	if unmarshaledReq.Type != req.Type {
		t.Errorf("Expected Type %d, got %d", req.Type, unmarshaledReq.Type)
	}
}

// Test decimal operations in types
func TestDecimalOperations(t *testing.T) {
	// Test zero decimal values
	zeroAmount := decimal.Zero
	if !zeroAmount.IsZero() {
		t.Error("decimal.Zero should be zero")
	}

	// Test decimal comparison
	price1 := decimal.NewFromFloat(45000)
	price2 := decimal.NewFromFloat(45000)
	if !price1.Equal(price2) {
		t.Error("Equal decimals should be equal")
	}

	// Test decimal arithmetic
	volume := decimal.NewFromFloat(0.01)
	price := decimal.NewFromFloat(45000)
	total := volume.Mul(price)
	expected := decimal.NewFromFloat(450)
	if !total.Equal(expected) {
		t.Errorf("Expected total %s, got %s", expected, total)
	}
}

// Test request validation patterns
func TestRequestValidation(t *testing.T) {
	// Test CreateOrderRequest with required fields
	req := CreateOrderRequest{
		Symbol: "BTCUSDT",
		Side:   "buy",
		Type:   "limit",
		Amount: decimal.NewFromFloat(0.01),
		Price:  decimal.NewFromFloat(45000),
	}

	if req.Symbol == "" {
		t.Error("Symbol should not be empty")
	}

	if req.Side != "buy" && req.Side != "sell" {
		t.Error("Side should be buy or sell")
	}

	if req.Amount.IsZero() {
		t.Error("Amount should not be zero for valid orders")
	}

	// Test FuturesCreateOrderRequest validation
	futuresReq := FuturesCreateOrderRequest{
		FuturesName:  "BTCUSDT",
		Type:         "LIMIT",
		Side:         "BUY",
		Open:         "OPEN",
		PositionType: "LONG",
		Volume:       decimal.NewFromFloat(0.01),
	}

	if futuresReq.FuturesName == "" {
		t.Error("FuturesName should not be empty")
	}

	if futuresReq.Volume.IsZero() {
		t.Error("Volume should not be zero for valid orders")
	}
}

// Test all response list types
func TestResponseTypes(t *testing.T) {
	// Test TickerListResponse
	tickerList := TickerListResponse{
		Date: 1640995200,
		Ticker: []ExchangeTicker{
			{
				Symbol: "BTCUSDT",
				High:   decimal.NewFromFloat(50000),
				Low:    decimal.NewFromFloat(45000),
				Last:   decimal.NewFromFloat(48000),
			},
		},
	}

	if len(tickerList.Ticker) != 1 {
		t.Error("TickerListResponse should contain one ticker")
	}

	// Test TradeListResponse
	tradeList := TradeListResponse{
		Count: 1,
		ResultList: []ExchangeTrade{
			{
				ID:      "trade1",
				OrderID: "order1",
				Symbol:  "BTCUSDT",
				Side:    "buy",
				Amount:  decimal.NewFromFloat(0.01),
				Price:   decimal.NewFromFloat(45000),
			},
		},
	}

	if tradeList.Count != 1 {
		t.Error("TradeListResponse count should match result list length")
	}

	// Test BatchOrderResponse
	batchResponse := BatchOrderResponse{
		Success: []BatchOrderResult{
			{
				Index:   0,
				OrderID: "success1",
			},
		},
		Failed: []BatchOrderResult{
			{
				Index: 1,
				Error: "Insufficient balance",
			},
		},
	}

	if len(batchResponse.Success) != 1 || len(batchResponse.Failed) != 1 {
		t.Error("BatchOrderResponse should contain success and failed results")
	}
}

// Test futures-specific types
func TestFuturesTypes(t *testing.T) {
	// Test FuturesAccount
	account := FuturesAccount{
		AccountId:       "account123",
		CollateralCoin:  "USDT",
		AccountBalance:  decimal.NewFromFloat(1000),
		TotalMargin:     decimal.NewFromFloat(500),
		TotalPnl:        decimal.NewFromFloat(50),
		AvailableMargin: decimal.NewFromFloat(450),
	}

	if account.AccountId == "" {
		t.Error("FuturesAccount should have AccountId")
	}

	// Test FuturesTicker
	ticker := FuturesTicker{
		Symbol:             "BTCUSDT",
		PriceChange:        decimal.NewFromFloat(1000),
		PriceChangePercent: decimal.NewFromFloat(2.1),
		WeightedAvgPrice:   decimal.NewFromFloat(47500),
		LastPrice:          decimal.NewFromFloat(48000),
		OpenPrice:          decimal.NewFromFloat(47000),
		HighPrice:          decimal.NewFromFloat(49000),
		LowPrice:           decimal.NewFromFloat(46000),
		Volume:             decimal.NewFromFloat(1000000),
		QuoteVolume:        decimal.NewFromFloat(47500000000),
		OpenTime:           1640995200000,
		CloseTime:          1641081600000,
		Count:              50000,
	}

	if ticker.Symbol == "" {
		t.Error("FuturesTicker should have Symbol")
	}

	// Test FuturesIndexPrice
	indexPrice := FuturesIndexPrice{
		Symbol:     "BTCUSDT",
		IndexPrice: decimal.NewFromFloat(47800),
		Time:       1640995200000,
	}

	if indexPrice.Symbol == "" {
		t.Error("FuturesIndexPrice should have Symbol")
	}

	// Test FuturesCapital
	capital := FuturesCapital{
		Asset:                  "USDT",
		WalletBalance:          decimal.NewFromFloat(1000),
		UnrealizedPnl:          decimal.NewFromFloat(50),
		MarginBalance:          decimal.NewFromFloat(1050),
		MaintMargin:            decimal.NewFromFloat(100),
		InitialMargin:          decimal.NewFromFloat(200),
		PositionInitialMargin:  decimal.NewFromFloat(150),
		OpenOrderInitialMargin: decimal.NewFromFloat(50),
		CrossWalletBalance:     decimal.NewFromFloat(1000),
		CrossUnPnl:             decimal.NewFromFloat(50),
		AvailableBalance:       decimal.NewFromFloat(800),
		MaxWithdrawAmount:      decimal.NewFromFloat(800),
	}

	if capital.Asset == "" {
		t.Error("FuturesCapital should have Asset")
	}
}

// Test edge cases and error conditions
func TestEdgeCases(t *testing.T) {
	// Test empty decimal values
	var emptyDecimal decimal.Decimal
	if !emptyDecimal.IsZero() {
		t.Error("Empty decimal should be zero")
	}

	// Test negative decimal values
	negativeAmount := decimal.NewFromFloat(-10)
	if negativeAmount.IsPositive() {
		t.Error("Negative amount should not be positive")
	}

	// Test large decimal values
	largePrice := decimal.NewFromFloat(999999999.99)
	if largePrice.IsZero() {
		t.Error("Large price should not be zero")
	}

	// Test string conversion
	priceStr := "45000.50"
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		t.Errorf("Failed to parse decimal from string: %v", err)
	}

	if price.String() != priceStr {
		t.Errorf("Expected price string %s, got %s", priceStr, price.String())
	}
}
