package byex

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// ExchangeAPI represents the exchange API methods
type ExchangeAPI struct {
	client *Client
}

// NewExchangeAPI creates a new exchange API instance
func NewExchangeAPI(client *Client) *ExchangeAPI {
	return &ExchangeAPI{client: client}
}

// Market Data APIs

// GetAllTicker gets all ticker information for trading pairs
func (e *ExchangeAPI) GetAllTicker() (*TickerListResponse, error) {
	resp, err := e.client.doExchangeRequest("GET", "/open/api/get_allticker", nil)
	if err != nil {
		return nil, err
	}

	var result TickerListResponse
	if err := json.Unmarshal(resp.Data.([]byte), &result); err != nil {
		return nil, fmt.Errorf("failed to parse ticker response: %w", err)
	}

	return &result, nil
}

// GetTicker gets ticker information for a specific symbol
func (e *ExchangeAPI) GetTicker(symbol string) (*ExchangeTicker, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/get_ticker", params)
	if err != nil {
		return nil, err
	}

	var result ExchangeTicker
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse ticker response: %w", err)
	}

	return &result, nil
}

// GetDepth gets order book depth for a specific symbol
func (e *ExchangeAPI) GetDepth(symbol string, depth int) (*ExchangeDepth, error) {
	params := map[string]string{
		"symbol": symbol,
		"type":   strconv.Itoa(depth),
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/market_dept", params)
	if err != nil {
		return nil, err
	}

	var result ExchangeDepth
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse depth response: %w", err)
	}

	return &result, nil
}

// GetKlines gets candlestick data for a specific symbol
func (e *ExchangeAPI) GetKlines(symbol, period string, size int) ([]ExchangeKline, error) {
	params := map[string]string{
		"symbol": symbol,
		"period": period,
		"size":   strconv.Itoa(size),
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/get_records", params)
	if err != nil {
		return nil, err
	}

	var result []ExchangeKline
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse klines response: %w", err)
	}

	return result, nil
}

// Trading APIs

// CreateOrder creates a new order
func (e *ExchangeAPI) CreateOrder(req CreateOrderRequest) (*OrderResponse, error) {
	params := map[string]string{
		"symbol": req.Symbol,
		"side":   req.Side,
		"type":   req.Type,
	}

	if !req.Amount.IsZero() {
		params["volume"] = req.Amount.String()
	}
	if !req.Price.IsZero() {
		params["price"] = req.Price.String()
	}
	if req.ClientOrderID != "" {
		params["client_order_id"] = req.ClientOrderID
	}

	resp, err := e.client.doExchangeRequest("POST", "/open/api/create_order", params)
	if err != nil {
		return nil, err
	}

	var result OrderResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse order response: %w", err)
	}

	return &result, nil
}

// CancelOrder cancels an existing order
func (e *ExchangeAPI) CancelOrder(symbol, orderID string) error {
	params := map[string]string{
		"symbol":   symbol,
		"order_id": orderID,
	}

	_, err := e.client.doExchangeRequest("POST", "/open/api/cancel_order", params)
	return err
}

// CancelAllOrders cancels all orders for a symbol
func (e *ExchangeAPI) CancelAllOrders(symbol string) error {
	params := map[string]string{
		"symbol": symbol,
	}

	_, err := e.client.doExchangeRequest("POST", "/open/api/cancel_order_all", params)
	return err
}

// BatchCreateOrders creates multiple orders at once
func (e *ExchangeAPI) BatchCreateOrders(req BatchOrderRequest) error {
	params := map[string]string{
		"symbol": req.Symbol,
	}

	// Convert orders to JSON string
	ordersJSON, err := json.Marshal(req.Orders)
	if err != nil {
		return fmt.Errorf("failed to marshal orders: %w", err)
	}
	params["orders_data"] = string(ordersJSON)

	_, err = e.client.doExchangeRequest("POST", "/open/api/mass_replace", params)
	return err
}

// GetCurrentOrders gets current orders (executing or unexecuted)
func (e *ExchangeAPI) GetCurrentOrders(symbol string, pageSize, page int) (*OrderListResponse, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if page > 0 {
		params["page"] = strconv.Itoa(page)
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/v2/new_order", params)
	if err != nil {
		return nil, err
	}

	var result OrderListResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse orders response: %w", err)
	}

	return &result, nil
}

// GetOrderHistory gets order history
func (e *ExchangeAPI) GetOrderHistory(symbol string, pageSize, page int) (*OrderListResponse, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if page > 0 {
		params["page"] = strconv.Itoa(page)
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/v2/all_order", params)
	if err != nil {
		return nil, err
	}

	var result OrderListResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse orders response: %w", err)
	}

	return &result, nil
}

// GetOrderInfo gets specific order information
func (e *ExchangeAPI) GetOrderInfo(symbol, orderID string) (*ExchangeOrder, error) {
	params := map[string]string{
		"symbol":   symbol,
		"order_id": orderID,
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/order_info", params)
	if err != nil {
		return nil, err
	}

	var result ExchangeOrder
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse order response: %w", err)
	}

	return &result, nil
}

// GetTrades gets trade history
func (e *ExchangeAPI) GetTrades(symbol string, pageSize, page int) (*TradeListResponse, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if page > 0 {
		params["page"] = strconv.Itoa(page)
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/v2/my_trades", params)
	if err != nil {
		return nil, err
	}

	var result TradeListResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse trades response: %w", err)
	}

	return &result, nil
}

// Account APIs

// GetAccount gets account information
func (e *ExchangeAPI) GetAccount() (*ExchangeAccount, error) {
	resp, err := e.client.doExchangeRequest("GET", "/open/api/user/account", nil)
	if err != nil {
		return nil, err
	}

	var result ExchangeAccount
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse account response: %w", err)
	}

	return &result, nil
}

// GetBalance gets balance for specific coins
func (e *ExchangeAPI) GetBalance(coins []string) ([]CoinBalance, error) {
	params := map[string]string{}
	if len(coins) > 0 {
		coinsJSON, err := json.Marshal(coins)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal coins: %w", err)
		}
		params["coins"] = string(coinsJSON)
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/user/account", params)
	if err != nil {
		return nil, err
	}

	var account ExchangeAccount
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &account); err != nil {
		return nil, fmt.Errorf("failed to parse account response: %w", err)
	}

	return account.CoinList, nil
}

// GetAllTradingRecords gets all trading records with advanced filtering
func (e *ExchangeAPI) GetAllTradingRecords(symbol string, pageSize, page int, id int64, startDate, endDate string, sort int) (*TradeListResponse, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if pageSize > 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if page > 0 {
		params["page"] = strconv.Itoa(page)
	}
	if id > 0 {
		params["id"] = strconv.FormatInt(id, 10)
	}
	if startDate != "" {
		params["startDate"] = startDate
	}
	if endDate != "" {
		params["endDate"] = endDate
	}
	if sort > 0 {
		params["sort"] = strconv.Itoa(sort)
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/all_trade", params)
	if err != nil {
		return nil, err
	}

	var result TradeListResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse all trades response: %w", err)
	}

	return &result, nil
}

// GetMarketPrices gets the latest price of all trading pairs
func (e *ExchangeAPI) GetMarketPrices() (map[string]decimal.Decimal, error) {
	resp, err := e.client.doExchangeRequest("GET", "/open/api/market", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]decimal.Decimal
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse market prices response: %w", err)
	}

	return result, nil
}

// BatchPlaceOrders places multiple orders in batch
func (e *ExchangeAPI) BatchPlaceOrders(symbol string, orderList []BatchOrder) (*BatchOrderResponse, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	// Convert orderList to JSON string
	orderListJSON, err := json.Marshal(orderList)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal order list: %w", err)
	}
	params["orderList"] = string(orderListJSON)

	resp, err := e.client.doExchangeRequest("POST", "/open/api/batchOrders", params)
	if err != nil {
		return nil, err
	}

	var result BatchOrderResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse batch orders response: %w", err)
	}

	return &result, nil
}

// BatchCancelOrders cancels multiple orders in batch
func (e *ExchangeAPI) BatchCancelOrders(symbol string, orderIds []string) error {
	params := map[string]string{
		"symbol": symbol,
	}

	// Convert orderIds to JSON string
	orderIdsJSON, err := json.Marshal(orderIds)
	if err != nil {
		return fmt.Errorf("failed to marshal order IDs: %w", err)
	}
	params["orderIds"] = string(orderIdsJSON)

	_, err = e.client.doExchangeRequest("POST", "/open/api/batchCancelOrders", params)
	return err
}

// GetOrderDetail gets detailed order information
func (e *ExchangeAPI) GetOrderDetail(symbol, orderID string) (*ExchangeOrderDetail, error) {
	params := map[string]string{
		"symbol":   symbol,
		"order_id": orderID,
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/order_info", params)
	if err != nil {
		return nil, err
	}

	var result ExchangeOrderDetail
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse order detail response: %w", err)
	}

	return &result, nil
}

// ReplaceOrder replaces an existing order
func (e *ExchangeAPI) ReplaceOrder(req ReplaceOrderRequest) (*OrderResponse, error) {
	params := map[string]string{
		"symbol":       req.Symbol,
		"cancel_order": req.CancelOrderID,
		"side":         req.Side,
		"type":         req.Type,
	}

	if !req.Amount.IsZero() {
		params["volume"] = req.Amount.String()
	}
	if !req.Price.IsZero() {
		params["price"] = req.Price.String()
	}
	if req.ClientOrderID != "" {
		params["client_order_id"] = req.ClientOrderID
	}

	resp, err := e.client.doExchangeRequest("POST", "/open/api/replace_order", params)
	if err != nil {
		return nil, err
	}

	var result OrderResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse replace order response: %w", err)
	}

	return &result, nil
}

// GetSymbolsCharge gets symbols with charge information
func (e *ExchangeAPI) GetSymbolsCharge() ([]SymbolCharge, error) {
	resp, err := e.client.doExchangeRequest("GET", "/open/api/common/symbols", nil)
	if err != nil {
		return nil, err
	}

	var result []SymbolCharge
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse symbols charge response: %w", err)
	}

	return result, nil
}

// GetLeverageFinanceBalance gets leverage finance balance
func (e *ExchangeAPI) GetLeverageFinanceBalance(symbol string) (*LeverageFinanceBalance, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := e.client.doExchangeRequest("GET", "/open/api/leverFinance/account", params)
	if err != nil {
		return nil, err
	}

	var result LeverageFinanceBalance
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse leverage finance balance response: %w", err)
	}

	return &result, nil
}
