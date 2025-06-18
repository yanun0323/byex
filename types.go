package byex

import (
	"github.com/shopspring/decimal"
)

// Exchange API Types

// ExchangeOrder represents an order in the exchange
type ExchangeOrder struct {
	ID               string          `json:"id"`
	Symbol           string          `json:"symbol"`
	Type             string          `json:"type"`
	Side             string          `json:"side"`
	Amount           decimal.Decimal `json:"amount"`
	Price            decimal.Decimal `json:"price"`
	Status           string          `json:"status"`
	CreatedAt        int64           `json:"created_at"`
	UpdatedAt        int64           `json:"updated_at"`
	FinishedAt       int64           `json:"finished_at"`
	CancelledAt      int64           `json:"cancelled_at"`
	AvgPrice         decimal.Decimal `json:"avg_price"`
	Source           string          `json:"source"`
	Fee              decimal.Decimal `json:"fee"`
	FeeCurrency      string          `json:"fee_currency"`
	FilledAmount     decimal.Decimal `json:"filled_amount"`
	FilledCashAmount decimal.Decimal `json:"filled_cash_amount"`
	FilledFees       decimal.Decimal `json:"filled_fees"`
}

// ExchangeTrade represents a trade in the exchange
type ExchangeTrade struct {
	ID          string          `json:"id"`
	OrderID     string          `json:"order_id"`
	Symbol      string          `json:"symbol"`
	Side        string          `json:"side"`
	Amount      decimal.Decimal `json:"amount"`
	Price       decimal.Decimal `json:"price"`
	Fee         decimal.Decimal `json:"fee"`
	FeeCurrency string          `json:"fee_currency"`
	Role        string          `json:"role"`
	CreatedAt   int64           `json:"created_at"`
}

// ExchangeTicker represents ticker information
type ExchangeTicker struct {
	Symbol      string          `json:"symbol"`
	High        decimal.Decimal `json:"high"`
	Low         decimal.Decimal `json:"low"`
	Last        decimal.Decimal `json:"last"`
	Vol         decimal.Decimal `json:"vol"`
	Amount      decimal.Decimal `json:"amount"`
	BuyPrice    decimal.Decimal `json:"buy"`
	SellPrice   decimal.Decimal `json:"sell"`
	NewCoinFlag int             `json:"newCoinFlag"`
	Change      decimal.Decimal `json:"change"`
	Rose        decimal.Decimal `json:"rose"`
}

// ExchangeDepth represents order book depth
type ExchangeDepth struct {
	Asks [][]decimal.Decimal `json:"asks"`
	Bids [][]decimal.Decimal `json:"bids"`
}

// ExchangeKline represents candlestick data
type ExchangeKline struct {
	Time   int64           `json:"time"`
	Open   decimal.Decimal `json:"open"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Close  decimal.Decimal `json:"close"`
	Volume decimal.Decimal `json:"volume"`
}

// ExchangeAccount represents user account information
type ExchangeAccount struct {
	TotalAsset    decimal.Decimal `json:"total_asset"`
	CoinList      []CoinBalance   `json:"coin_list"`
	NormalCount   decimal.Decimal `json:"normal_count"`
	LockedCount   decimal.Decimal `json:"locked_count"`
	FreezingCount decimal.Decimal `json:"freezing_count"`
	BtcValuation  decimal.Decimal `json:"btc_valuation"`
	RmbValuation  decimal.Decimal `json:"rmb_valuation"`
}

// CoinBalance represents balance for a specific coin
type CoinBalance struct {
	Coin     string          `json:"coin"`
	Normal   decimal.Decimal `json:"normal"`
	Locked   decimal.Decimal `json:"locked"`
	BtcValue decimal.Decimal `json:"btcValue"`
	RmbValue decimal.Decimal `json:"rmbValue"`
}

// Futures API Types

// FuturesOrder represents a futures order
type FuturesOrder struct {
	OrderID       string          `json:"orderId"`
	ClientOrderID string          `json:"clientOrderId"`
	Symbol        string          `json:"symbol"`
	Type          string          `json:"type"`
	Side          string          `json:"side"`
	Open          string          `json:"open"`
	PositionType  string          `json:"positionType"`
	Price         decimal.Decimal `json:"price"`
	Volume        decimal.Decimal `json:"volume"`
	Status        string          `json:"status"`
	CreatedAt     int64           `json:"created_at"`
	UpdatedAt     int64           `json:"updated_at"`
}

// FuturesTrade represents a futures trade
type FuturesTrade struct {
	ID        string          `json:"id"`
	OrderID   string          `json:"order_id"`
	Symbol    string          `json:"symbol"`
	Side      string          `json:"side"`
	Volume    decimal.Decimal `json:"volume"`
	Price     decimal.Decimal `json:"price"`
	Fee       decimal.Decimal `json:"fee"`
	Timestamp int64           `json:"timestamp"`
}

// FuturesPosition represents a futures position
type FuturesPosition struct {
	Symbol            string          `json:"symbol"`
	PositionSide      string          `json:"positionSide"`
	PositionAmt       decimal.Decimal `json:"positionAmt"`
	AvgPrice          decimal.Decimal `json:"avgPrice"`
	UnrealizedPnl     decimal.Decimal `json:"unrealizedPnl"`
	RealizedPnl       decimal.Decimal `json:"realizedPnl"`
	MarginType        string          `json:"marginType"`
	InitialMargin     decimal.Decimal `json:"initialMargin"`
	MaintenanceMargin decimal.Decimal `json:"maintenanceMargin"`
	PositionValue     decimal.Decimal `json:"positionValue"`
	Leverage          decimal.Decimal `json:"leverage"`
}

// FuturesAccount represents futures account information
type FuturesAccount struct {
	AccountId       string          `json:"accountId"`
	CollateralCoin  string          `json:"collateralCoin"`
	AccountBalance  decimal.Decimal `json:"accountBalance"`
	TotalMargin     decimal.Decimal `json:"totalMargin"`
	TotalPnl        decimal.Decimal `json:"totalPnl"`
	AvailableMargin decimal.Decimal `json:"availableMargin"`
}

// FuturesTicker represents futures ticker information
type FuturesTicker struct {
	Symbol             string          `json:"symbol"`
	PriceChange        decimal.Decimal `json:"priceChange"`
	PriceChangePercent decimal.Decimal `json:"priceChangePercent"`
	WeightedAvgPrice   decimal.Decimal `json:"weightedAvgPrice"`
	LastPrice          decimal.Decimal `json:"lastPrice"`
	LastQty            decimal.Decimal `json:"lastQty"`
	OpenPrice          decimal.Decimal `json:"openPrice"`
	HighPrice          decimal.Decimal `json:"highPrice"`
	LowPrice           decimal.Decimal `json:"lowPrice"`
	Volume             decimal.Decimal `json:"volume"`
	QuoteVolume        decimal.Decimal `json:"quoteVolume"`
	OpenTime           int64           `json:"openTime"`
	CloseTime          int64           `json:"closeTime"`
	Count              int64           `json:"count"`
}

// Request Types

// CreateOrderRequest represents a create order request
type CreateOrderRequest struct {
	Symbol        string          `json:"symbol"`
	Side          string          `json:"side"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
	ClientOrderID string          `json:"client_order_id,omitempty"`
}

// FuturesCreateOrderRequest represents a futures create order request
type FuturesCreateOrderRequest struct {
	FuturesName   string          `json:"futuresName"`
	Type          string          `json:"type"`
	Side          string          `json:"side"`
	Open          string          `json:"open"`
	PositionType  string          `json:"positionType"`
	Price         decimal.Decimal `json:"price,omitempty"`
	Volume        decimal.Decimal `json:"volume"`
	ClientOrderID string          `json:"clientOrderId,omitempty"`
}

// BatchOrderRequest represents a batch order request
type BatchOrderRequest struct {
	Symbol string               `json:"symbol"`
	Orders []CreateOrderRequest `json:"orders"`
}

// Response Types

// OrderResponse represents order creation response
type OrderResponse struct {
	OrderID string `json:"orderId"`
}

// OrderListResponse represents order list response
type OrderListResponse struct {
	Count      int             `json:"count"`
	ResultList []ExchangeOrder `json:"resultList"`
}

// TradeListResponse represents trade list response
type TradeListResponse struct {
	Count      int             `json:"count"`
	ResultList []ExchangeTrade `json:"resultList"`
}

// TickerListResponse represents ticker list response
type TickerListResponse struct {
	Date   int64            `json:"date"`
	Ticker []ExchangeTicker `json:"ticker"`
}

// Constants for order types, sides, etc.
const (
	// Order Types
	OrderTypeLimit  = "LIMIT"
	OrderTypeMarket = "MARKET"

	// Order Sides
	OrderSideBuy  = "BUY"
	OrderSideSell = "SELL"

	// Order Status
	OrderStatusNew             = "NEW"
	OrderStatusPartiallyFilled = "PARTIALLY_FILLED"
	OrderStatusFilled          = "FILLED"
	OrderStatusCancelled       = "CANCELLED"
	OrderStatusRejected        = "REJECTED"

	// Futures Position Types
	FuturesPositionTypeCross    = "1"
	FuturesPositionTypeIsolated = "2"

	// Futures Trade Types
	FuturesTradeTypeOpen  = "OPEN"
	FuturesTradeTypeClose = "CLOSE"
)

// Additional Exchange Types

// BatchOrder represents a single order in batch operations
type BatchOrder struct {
	Volume        decimal.Decimal `json:"volume"`
	Price         decimal.Decimal `json:"price,omitempty"`
	Side          string          `json:"side"`
	Type          int             `json:"type"`       // 1: limit, 2: market
	VolumeType    int             `json:"volumeType"` // 1: amount value, 2: base coin amount
	ClientOrderID string          `json:"clientOrderId,omitempty"`
}

// BatchOrderResponse represents the response from batch order operations
type BatchOrderResponse struct {
	Success []BatchOrderResult `json:"success"`
	Failed  []BatchOrderResult `json:"failed"`
}

// BatchOrderResult represents a single order result in batch operations
type BatchOrderResult struct {
	Index         int    `json:"index"`
	OrderID       string `json:"orderId,omitempty"`
	ClientOrderID string `json:"clientOrderId,omitempty"`
	Error         string `json:"error,omitempty"`
}

// ExchangeOrderDetail represents detailed order information
type ExchangeOrderDetail struct {
	ExchangeOrder
	Trades []ExchangeTrade `json:"trades"`
}

// ReplaceOrderRequest represents a replace order request
type ReplaceOrderRequest struct {
	Symbol        string          `json:"symbol"`
	CancelOrderID string          `json:"cancel_order"`
	Side          string          `json:"side"`
	Type          string          `json:"type"`
	Amount        decimal.Decimal `json:"amount,omitempty"`
	Price         decimal.Decimal `json:"price,omitempty"`
	ClientOrderID string          `json:"client_order_id,omitempty"`
}

// SymbolCharge represents symbol with charge information
type SymbolCharge struct {
	Symbol              string          `json:"symbol"`
	BaseAsset           string          `json:"baseAsset"`
	QuoteAsset          string          `json:"quoteAsset"`
	BaseAssetPrecision  int             `json:"baseAssetPrecision"`
	QuoteAssetPrecision int             `json:"quoteAssetPrecision"`
	Status              string          `json:"status"`
	TakerCommission     decimal.Decimal `json:"takerCommission"`
	MakerCommission     decimal.Decimal `json:"makerCommission"`
	MinPrice            decimal.Decimal `json:"minPrice"`
	MaxPrice            decimal.Decimal `json:"maxPrice"`
	TickSize            decimal.Decimal `json:"tickSize"`
	MinQty              decimal.Decimal `json:"minQty"`
	MaxQty              decimal.Decimal `json:"maxQty"`
	StepSize            decimal.Decimal `json:"stepSize"`
}

// LeverageFinanceBalance represents leverage finance balance information
type LeverageFinanceBalance struct {
	Symbol        string          `json:"symbol"`
	BaseAsset     string          `json:"baseAsset"`
	QuoteAsset    string          `json:"quoteAsset"`
	BaseBalance   decimal.Decimal `json:"baseBalance"`
	QuoteBalance  decimal.Decimal `json:"quoteBalance"`
	BaseBorrowed  decimal.Decimal `json:"baseBorrowed"`
	QuoteBorrowed decimal.Decimal `json:"quoteBorrowed"`
	BaseInterest  decimal.Decimal `json:"baseInterest"`
	QuoteInterest decimal.Decimal `json:"quoteInterest"`
	BaseNetAsset  decimal.Decimal `json:"baseNetAsset"`
	QuoteNetAsset decimal.Decimal `json:"quoteNetAsset"`
	MarginLevel   decimal.Decimal `json:"marginLevel"`
}

// Additional Futures Types

// FuturesOrderDetail represents detailed futures order information
type FuturesOrderDetail struct {
	FuturesOrder
	Trades []FuturesTrade `json:"trades"`
}

// FuturesBatchOrderRequest represents batch order request for futures
type FuturesBatchOrderRequest struct {
	FuturesName string                      `json:"futuresName"`
	Orders      []FuturesCreateOrderRequest `json:"orders"`
}

// FuturesIndexPrice represents index price information
type FuturesIndexPrice struct {
	Symbol     string          `json:"symbol"`
	IndexPrice decimal.Decimal `json:"indexPrice"`
	Time       int64           `json:"time"`
}

// FuturesCapital represents capital/fund information
type FuturesCapital struct {
	Asset                  string          `json:"asset"`
	WalletBalance          decimal.Decimal `json:"walletBalance"`
	UnrealizedPnl          decimal.Decimal `json:"unrealizedPnl"`
	MarginBalance          decimal.Decimal `json:"marginBalance"`
	MaintMargin            decimal.Decimal `json:"maintMargin"`
	InitialMargin          decimal.Decimal `json:"initialMargin"`
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"`
	CrossWalletBalance     decimal.Decimal `json:"crossWalletBalance"`
	CrossUnPnl             decimal.Decimal `json:"crossUnPnl"`
	AvailableBalance       decimal.Decimal `json:"availableBalance"`
	MaxWithdrawAmount      decimal.Decimal `json:"maxWithdrawAmount"`
}

// FuturesTransferRequest represents fund transfer request
type FuturesTransferRequest struct {
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
	Type     int             `json:"type"` // 1: spot to futures, 2: futures to spot
}
