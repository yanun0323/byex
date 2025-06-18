package byex

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// FuturesAPI represents the futures API methods
type FuturesAPI struct {
	client *Client
}

// NewFuturesAPI creates a new futures API instance
func NewFuturesAPI(client *Client) *FuturesAPI {
	return &FuturesAPI{client: client}
}

// Market Data APIs

// GetTicker gets futures ticker information for a specific symbol
func (f *FuturesAPI) GetTicker(symbol string) (*FuturesTicker, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/ticker", params)
	if err != nil {
		return nil, err
	}

	var result FuturesTicker
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse ticker response: %w", err)
	}

	return &result, nil
}

// GetDepth gets order book depth for a specific futures symbol
func (f *FuturesAPI) GetDepth(symbol string, limit int) (*ExchangeDepth, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/depth", params)
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

// GetKlines gets futures candlestick data
func (f *FuturesAPI) GetKlines(symbol, interval string, limit int) ([]ExchangeKline, error) {
	params := map[string]string{
		"symbol":   symbol,
		"interval": interval,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/klines", params)
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

// CreateOrder creates a new futures order
func (f *FuturesAPI) CreateOrder(req FuturesCreateOrderRequest) (*OrderResponse, error) {
	resp, err := f.client.doFuturesRequest("POST", "/fapi/v1/trade/order", req)
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

// CancelOrder cancels a futures order
func (f *FuturesAPI) CancelOrder(futuresName, orderID string) error {
	req := map[string]string{
		"futuresName": futuresName,
		"orderId":     orderID,
	}

	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/trade/cancel", req)
	return err
}

// CancelAllOrders cancels all futures orders for a symbol
func (f *FuturesAPI) CancelAllOrders(futuresName string) error {
	req := map[string]string{
		"futuresName": futuresName,
	}

	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/trade/cancelAll", req)
	return err
}

// GetCurrentOrders gets current futures orders
func (f *FuturesAPI) GetCurrentOrders(futuresName string) ([]FuturesOrder, error) {
	params := map[string]string{
		"futuresName": futuresName,
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/trade/openOrders", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesOrder
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse orders response: %w", err)
	}

	return result, nil
}

// GetOrderHistory gets futures order history
func (f *FuturesAPI) GetOrderHistory(futuresName string, limit int) ([]FuturesOrder, error) {
	params := map[string]string{
		"futuresName": futuresName,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/trade/allOrders", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesOrder
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse orders response: %w", err)
	}

	return result, nil
}

// GetOrderInfo gets specific futures order information
func (f *FuturesAPI) GetOrderInfo(futuresName, orderID string) (*FuturesOrder, error) {
	params := map[string]string{
		"futuresName": futuresName,
		"orderId":     orderID,
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/trade/order", params)
	if err != nil {
		return nil, err
	}

	var result FuturesOrder
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse order response: %w", err)
	}

	return &result, nil
}

// GetTrades gets futures trade history
func (f *FuturesAPI) GetTrades(futuresName string, limit int) ([]FuturesTrade, error) {
	params := map[string]string{
		"futuresName": futuresName,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/trade/userTrades", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesTrade
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse trades response: %w", err)
	}

	return result, nil
}

// Position and Account APIs

// GetPositions gets futures positions
func (f *FuturesAPI) GetPositions(futuresName string) ([]FuturesPosition, error) {
	params := map[string]string{}
	if futuresName != "" {
		params["futuresName"] = futuresName
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/position/positions", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesPosition
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse positions response: %w", err)
	}

	return result, nil
}

// GetAccount gets futures account information
func (f *FuturesAPI) GetAccount() (*FuturesAccount, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/account/balance", nil)
	if err != nil {
		return nil, err
	}

	var result FuturesAccount
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse account response: %w", err)
	}

	return &result, nil
}

// SetLeverage sets leverage for a futures symbol
func (f *FuturesAPI) SetLeverage(futuresName string, leverage int) error {
	req := map[string]interface{}{
		"futuresName": futuresName,
		"leverage":    leverage,
	}

	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/position/leverage", req)
	return err
}

// SetMarginType sets margin type for a futures symbol
func (f *FuturesAPI) SetMarginType(futuresName, marginType string) error {
	req := map[string]string{
		"futuresName": futuresName,
		"marginType":  marginType,
	}

	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/position/marginType", req)
	return err
}

// ModifyMargin modifies position margin
func (f *FuturesAPI) ModifyMargin(futuresName string, amount float64, marginType int) error {
	req := map[string]interface{}{
		"futuresName": futuresName,
		"amount":      amount,
		"type":        marginType, // 1: Add margin, 2: Reduce margin
	}

	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/position/positionMargin", req)
	return err
}

// GetAllTicker gets all futures ticker information
func (f *FuturesAPI) GetAllTicker() ([]FuturesTicker, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/ticker/24hr", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesTicker
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse all ticker response: %w", err)
	}

	return result, nil
}

// GetIndexPrice gets index price for a specific symbol
func (f *FuturesAPI) GetIndexPrice(symbol string) (*FuturesIndexPrice, error) {
	params := map[string]string{
		"symbol": symbol,
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/premiumIndex", params)
	if err != nil {
		return nil, err
	}

	var result FuturesIndexPrice
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse index price response: %w", err)
	}

	return &result, nil
}

// GetAllIndexPrice gets all index prices
func (f *FuturesAPI) GetAllIndexPrice() ([]FuturesIndexPrice, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/premiumIndex", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesIndexPrice
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse all index prices response: %w", err)
	}

	return result, nil
}

// GetAllTagIndexPrice gets all tag index prices
func (f *FuturesAPI) GetAllTagIndexPrice() ([]FuturesIndexPrice, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/indexPrice", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesIndexPrice
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse tag index prices response: %w", err)
	}

	return result, nil
}

// GetAllFuturesDepth gets all futures depth information
func (f *FuturesAPI) GetAllFuturesDepth() (map[string]ExchangeDepth, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/depth/all", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]ExchangeDepth
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse all futures depth response: %w", err)
	}

	return result, nil
}

// BatchCreateOrders creates multiple futures orders in batch
func (f *FuturesAPI) BatchCreateOrders(req FuturesBatchOrderRequest) ([]OrderResponse, error) {
	resp, err := f.client.doFuturesRequest("POST", "/fapi/v1/batchOrders", req)
	if err != nil {
		return nil, err
	}

	var result []OrderResponse
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse batch orders response: %w", err)
	}

	return result, nil
}

// BatchCancelOrders cancels multiple futures orders in batch
func (f *FuturesAPI) BatchCancelOrders(futuresName string, orderIds []string) error {
	req := map[string]interface{}{
		"futuresName": futuresName,
		"orderIdList": orderIds,
	}

	_, err := f.client.doFuturesRequest("DELETE", "/fapi/v1/batchOrders", req)
	return err
}

// GetCapital gets futures capital/fund information
func (f *FuturesAPI) GetCapital() ([]FuturesCapital, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/balance", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesCapital
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse capital response: %w", err)
	}

	return result, nil
}

// GetFutureAccounts gets futures account information
func (f *FuturesAPI) GetFutureAccounts() ([]FuturesAccount, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/account", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesAccount
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse futures accounts response: %w", err)
	}

	return result, nil
}

// CreateFutureAccount creates a new futures account
func (f *FuturesAPI) CreateFutureAccount() error {
	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/account", nil)
	return err
}

// FundTransfer transfers funds between spot and futures accounts
func (f *FuturesAPI) FundTransfer(req FuturesTransferRequest) error {
	_, err := f.client.doFuturesRequest("POST", "/fapi/v1/transfer", req)
	return err
}

// GetAllPositions gets all account positions
func (f *FuturesAPI) GetAllPositions() ([]FuturesPosition, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/positionRisk", nil)
	if err != nil {
		return nil, err
	}

	var result []FuturesPosition
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse all positions response: %w", err)
	}

	return result, nil
}

// GetFutures gets futures symbol information
func (f *FuturesAPI) GetFutures() ([]map[string]interface{}, error) {
	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/exchangeInfo", nil)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse futures info response: %w", err)
	}

	return result, nil
}

// GetOpeningOrders gets opening orders (alternative method)
func (f *FuturesAPI) GetOpeningOrders(futuresName string, limit int) ([]FuturesOrder, error) {
	params := map[string]string{
		"futuresName": futuresName,
	}

	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/openOrders", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesOrder
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse opening orders response: %w", err)
	}

	return result, nil
}

// GetMyTrades gets user trades (alternative method name for consistency)
func (f *FuturesAPI) GetMyTrades(futuresName string, fromId string, limit int) ([]FuturesTrade, error) {
	params := map[string]string{
		"futuresName": futuresName,
	}

	if fromId != "" {
		params["fromId"] = fromId
	}
	if limit > 0 {
		params["limit"] = strconv.Itoa(limit)
	}

	resp, err := f.client.doFuturesRequest("GET", "/fapi/v1/userTrades", params)
	if err != nil {
		return nil, err
	}

	var result []FuturesTrade
	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response data: %w", err)
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return nil, fmt.Errorf("failed to parse my trades response: %w", err)
	}

	return result, nil
}
