# 100EX Golang SDK

A comprehensive Golang SDK for interacting with 100EX exchange APIs, supporting both spot trading and futures trading.

## Features

- **Spot Trading API**: Complete spot trading functionality including market data, order management, and account operations
- **Futures Trading API**: Full futures trading support with position management and advanced order types
- **Type Safety**: All API responses are strongly typed with proper data structures
- **Error Handling**: Comprehensive error handling with custom error types
- **Authentication**: Built-in signature generation for both exchange and futures APIs
- **Decimal Precision**: Uses `shopspring/decimal` for precise financial calculations

## Installation

```bash
go get github.com/yanun0323/byex
```

## Quick Start

### Exchange (Spot Trading)

```go
package main

import (
    "fmt"
    "log"

    "github.com/shopspring/decimal"
    "github.com/yanun0323/byex"
)

func main() {
    // Create client
    client := byex.NewClient("your-api-key", "your-secret-key")
    exchangeAPI := client.Exchange()

    // Get all tickers
    tickers, err := exchangeAPI.GetAllTicker()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d trading pairs\n", len(tickers.Ticker))

    // Create a limit buy order
    order := byex.CreateOrderRequest{
        Symbol: "BTCUSDT",
        Side:   byex.OrderSideBuy,
        Type:   byex.OrderTypeLimit,
        Amount: decimal.NewFromFloat(0.01),
        Price:  decimal.NewFromFloat(45000),
    }

    orderResp, err := exchangeAPI.CreateOrder(order)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Order created with ID: %s\n", orderResp.OrderID)
}
```

### Futures Trading

```go
package main

import (
    "fmt"
    "log"

    "github.com/shopspring/decimal"
    "github.com/yanun0323/byex"
)

func main() {
    // Create client
    client := byex.NewClient("your-api-key", "your-secret-key")
    futuresAPI := client.Futures()

    // Get account information
    account, err := futuresAPI.GetAccount()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Account Balance: %s\n", account.AccountBalance.String())

    // Create a futures order
    order := byex.FuturesCreateOrderRequest{
        FuturesName:  "E-BTC-USDT",
        Type:         byex.OrderTypeLimit,
        Side:         byex.OrderSideBuy,
        Open:         byex.FuturesTradeTypeOpen,
        PositionType: byex.FuturesPositionTypeCross,
        Price:        decimal.NewFromFloat(45000),
        Volume:       decimal.NewFromFloat(10), // 10 multiplier
    }

    orderResp, err := futuresAPI.CreateOrder(order)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Futures order created with ID: %s\n", orderResp.OrderID)
}
```

## API Reference

### Exchange API

#### Market Data

- `GetAllTicker()` - Get all ticker information
- `GetTicker(symbol)` - Get ticker for specific symbol
- `GetDepth(symbol, depth)` - Get order book depth
- `GetKlines(symbol, period, size)` - Get candlestick data
- `GetMarketPrices()` - Get latest price of all trading pairs

#### Trading

- `CreateOrder(request)` - Create new order
- `CancelOrder(symbol, orderID)` - Cancel specific order
- `CancelAllOrders(symbol)` - Cancel all orders for symbol
- `BatchCreateOrders(request)` - Create multiple orders (deprecated)
- `BatchPlaceOrders(symbol, orderList)` - Place multiple orders in batch
- `BatchCancelOrders(symbol, orderIds)` - Cancel multiple orders in batch
- `ReplaceOrder(request)` - Replace an existing order

#### Order Management

- `GetCurrentOrders(symbol, pageSize, page)` - Get current orders
- `GetOrderHistory(symbol, pageSize, page)` - Get order history
- `GetOrderInfo(symbol, orderID)` - Get specific order details
- `GetOrderDetail(symbol, orderID)` - Get detailed order information
- `GetTrades(symbol, pageSize, page)` - Get trade history
- `GetAllTradingRecords(symbol, pageSize, page, id, startDate, endDate, sort)` - Get all trading records with filtering

#### Account

- `GetAccount()` - Get account information
- `GetBalance(coins)` - Get balance for specific coins
- `GetLeverageFinanceBalance(symbol)` - Get leverage finance balance

#### Public

- `GetSymbolsCharge()` - Get symbols with charge information

### Futures API

#### Market Data

- `GetTicker(symbol)` - Get futures ticker
- `GetAllTicker()` - Get all futures tickers
- `GetDepth(symbol, limit)` - Get futures order book
- `GetAllFuturesDepth()` - Get all futures depth information
- `GetKlines(symbol, interval, limit)` - Get futures candlestick data
- `GetIndexPrice(symbol)` - Get index price for specific symbol
- `GetAllIndexPrice()` - Get all index prices
- `GetAllTagIndexPrice()` - Get all tag index prices

#### Trading

- `CreateOrder(request)` - Create futures order
- `CancelOrder(futuresName, orderID)` - Cancel futures order
- `CancelAllOrders(futuresName)` - Cancel all futures orders
- `BatchCreateOrders(request)` - Create multiple futures orders in batch
- `BatchCancelOrders(futuresName, orderIds)` - Cancel multiple futures orders in batch

#### Order Management

- `GetCurrentOrders(futuresName)` - Get current futures orders
- `GetOpeningOrders(futuresName, limit)` - Get opening orders (alternative method)
- `GetOrderHistory(futuresName, limit)` - Get futures order history
- `GetOrderInfo(futuresName, orderID)` - Get specific futures order
- `GetTrades(futuresName, limit)` - Get futures trade history
- `GetMyTrades(futuresName, fromId, limit)` - Get user trades (alternative method)

#### Position & Account

- `GetPositions(futuresName)` - Get futures positions
- `GetAllPositions()` - Get all account positions
- `GetAccount()` - Get futures account information
- `GetFutureAccounts()` - Get futures account information (alternative method)
- `GetCapital()` - Get futures capital/fund information
- `CreateFutureAccount()` - Create a new futures account
- `SetLeverage(futuresName, leverage)` - Set leverage
- `SetMarginType(futuresName, marginType)` - Set margin type
- `ModifyMargin(futuresName, amount, type)` - Modify position margin
- `FundTransfer(request)` - Transfer funds between spot and futures accounts

#### Exchange Info

- `GetFutures()` - Get futures symbol information

## Data Types

### Order Types

- `OrderTypeLimit` - Limit order
- `OrderTypeMarket` - Market order

### Order Sides

- `OrderSideBuy` - Buy order
- `OrderSideSell` - Sell order

### Order Status

- `OrderStatusNew` - New order
- `OrderStatusPartiallyFilled` - Partially filled
- `OrderStatusFilled` - Fully filled
- `OrderStatusCancelled` - Cancelled
- `OrderStatusRejected` - Rejected

### Futures Position Types

- `FuturesPositionTypeCross` - Cross margin mode
- `FuturesPositionTypeIsolated` - Isolated margin mode

### Futures Trade Types

- `FuturesTradeTypeOpen` - Open position
- `FuturesTradeTypeClose` - Close position

## Advanced Features

### Batch Operations

```go
// Batch place orders
orders := []byex.BatchOrder{
    {
        Volume:     decimal.NewFromFloat(0.01),
        Price:      decimal.NewFromFloat(45000),
        Side:       byex.OrderSideBuy,
        Type:       1, // limit order
        VolumeType: 2, // base coin amount
    },
    {
        Volume:     decimal.NewFromFloat(0.02),
        Price:      decimal.NewFromFloat(46000),
        Side:       byex.OrderSideBuy,
        Type:       1, // limit order
        VolumeType: 2, // base coin amount
    },
}

result, err := exchangeAPI.BatchPlaceOrders("BTCUSDT", orders)
```

### Order Replacement

```go
// Replace an existing order
replaceReq := byex.ReplaceOrderRequest{
    Symbol:        "BTCUSDT",
    CancelOrderID: "old-order-id",
    Side:          byex.OrderSideBuy,
    Type:          byex.OrderTypeLimit,
    Amount:        decimal.NewFromFloat(0.02),
    Price:         decimal.NewFromFloat(44000),
}

newOrder, err := exchangeAPI.ReplaceOrder(replaceReq)
```

### Fund Transfer

```go
// Transfer funds from spot to futures
transferReq := byex.FuturesTransferRequest{
    Currency: "USDT",
    Amount:   decimal.NewFromFloat(1000),
    Type:     1, // 1: spot to futures, 2: futures to spot
}

err := futuresAPI.FundTransfer(transferReq)
```

## Error Handling

The SDK provides comprehensive error handling:

```go
orderResp, err := exchangeAPI.CreateOrder(order)
if err != nil {
    if apiErr, ok := err.(*byex.Error); ok {
        fmt.Printf("API Error - Code: %s, Message: %s\n", apiErr.Code, apiErr.Message)
    } else {
        fmt.Printf("General error: %s\n", err.Error())
    }
    return
}
```

## Authentication

The SDK automatically handles authentication for both exchange and futures APIs:

- **Exchange API**: Uses MD5 signature with query parameters
- **Futures API**: Uses HMAC-SHA256 signature with headers

Both authentication methods are implemented according to the official 100EX API documentation.

## Rate Limiting

Please be aware of the rate limits imposed by the 100EX API. The SDK does not implement rate limiting internally, so you should handle this in your application logic.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License.

## Disclaimer

This SDK is unofficial and not affiliated with 100EX. Use at your own risk. Always test thoroughly before using in production environments.
