# \MarketsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldCoingecko**](MarketsAPI.md#WorldWorldCoingecko) | **Get** /v1/world/coingecko | Crypto spot prices (CoinGecko)
[**WorldWorldEia**](MarketsAPI.md#WorldWorldEia) | **Get** /v1/world/eia | US EIA energy data (requires EIA_API_KEY)
[**WorldWorldEtfFlows**](MarketsAPI.md#WorldWorldEtfFlows) | **Get** /v1/world/etf-flows | BTC/crypto ETF flows
[**WorldWorldFinnhub**](MarketsAPI.md#WorldWorldFinnhub) | **Get** /v1/world/finnhub | Market data (Finnhub; requires FINNHUB_KEY)
[**WorldWorldFredData**](MarketsAPI.md#WorldWorldFredData) | **Get** /v1/world/fred-data | FRED economic series (requires FRED_API_KEY)
[**WorldWorldIndicators**](MarketsAPI.md#WorldWorldIndicators) | **Get** /v1/world/indicators | Trader indicator suite (VIX/VVIX/MOVE, yield curve + 2s10s, crypto/equity fear-greed, momentum, sector breadth, BTC dominance + perp funding, DXY/metals, risk-on/off composite)
[**WorldWorldMacroSignals**](MarketsAPI.md#WorldWorldMacroSignals) | **Get** /v1/world/macro-signals | Macro market-radar signals
[**WorldWorldPolymarket**](MarketsAPI.md#WorldWorldPolymarket) | **Get** /v1/world/polymarket | Prediction markets (Polymarket gamma)
[**WorldWorldSentiment**](MarketsAPI.md#WorldWorldSentiment) | **Get** /v1/world/sentiment | Realtime news-sentiment index (GDELT tone: global + per-topic + per-region, 24h sparkline + velocity)
[**WorldWorldStablecoinMarkets**](MarketsAPI.md#WorldWorldStablecoinMarkets) | **Get** /v1/world/stablecoin-markets | Stablecoin market health
[**WorldWorldStockIndex**](MarketsAPI.md#WorldWorldStockIndex) | **Get** /v1/world/stock-index | Stock index snapshot
[**WorldWorldWorldbank**](MarketsAPI.md#WorldWorldWorldbank) | **Get** /v1/world/worldbank | World Bank indicators
[**WorldWorldYahooFinance**](MarketsAPI.md#WorldWorldYahooFinance) | **Get** /v1/world/yahoo-finance | Equity/index quotes (Yahoo)



## WorldWorldCoingecko

> map[string]interface{} WorldWorldCoingecko(ctx).Execute()

Crypto spot prices (CoinGecko)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldCoingecko(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldCoingecko``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCoingecko`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldCoingecko`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCoingeckoRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldEia

> map[string]interface{} WorldWorldEia(ctx).Execute()

US EIA energy data (requires EIA_API_KEY)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldEia(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldEia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldEia`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldEia`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldEiaRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldEtfFlows

> map[string]interface{} WorldWorldEtfFlows(ctx).Execute()

BTC/crypto ETF flows

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldEtfFlows(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldEtfFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldEtfFlows`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldEtfFlows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldEtfFlowsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldFinnhub

> map[string]interface{} WorldWorldFinnhub(ctx).Symbol(symbol).Execute()

Market data (Finnhub; requires FINNHUB_KEY)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	symbol := "symbol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldFinnhub(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldFinnhub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldFinnhub`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldFinnhub`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldFinnhubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldFredData

> map[string]interface{} WorldWorldFredData(ctx).SeriesId(seriesId).Execute()

FRED economic series (requires FRED_API_KEY)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	seriesId := "seriesId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldFredData(context.Background()).SeriesId(seriesId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldFredData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldFredData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldFredData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldFredDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldIndicators

> map[string]interface{} WorldWorldIndicators(ctx).Execute()

Trader indicator suite (VIX/VVIX/MOVE, yield curve + 2s10s, crypto/equity fear-greed, momentum, sector breadth, BTC dominance + perp funding, DXY/metals, risk-on/off composite)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldIndicators(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldIndicators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldIndicators`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldIndicators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldIndicatorsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldMacroSignals

> map[string]interface{} WorldWorldMacroSignals(ctx).Execute()

Macro market-radar signals

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldMacroSignals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldMacroSignals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldMacroSignals`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldMacroSignals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldMacroSignalsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldPolymarket

> map[string]interface{} WorldWorldPolymarket(ctx).Execute()

Prediction markets (Polymarket gamma)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldPolymarket(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldPolymarket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldPolymarket`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldPolymarket`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldPolymarketRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldSentiment

> map[string]interface{} WorldWorldSentiment(ctx).Execute()

Realtime news-sentiment index (GDELT tone: global + per-topic + per-region, 24h sparkline + velocity)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldSentiment(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldSentiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldSentiment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldSentiment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldSentimentRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldStablecoinMarkets

> map[string]interface{} WorldWorldStablecoinMarkets(ctx).Execute()

Stablecoin market health

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldStablecoinMarkets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldStablecoinMarkets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldStablecoinMarkets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldStablecoinMarkets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldStablecoinMarketsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldStockIndex

> map[string]interface{} WorldWorldStockIndex(ctx).Execute()

Stock index snapshot

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldStockIndex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldStockIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldStockIndex`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldStockIndex`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldStockIndexRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldWorldbank

> map[string]interface{} WorldWorldWorldbank(ctx).Indicator(indicator).Country(country).Execute()

World Bank indicators

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	indicator := "indicator_example" // string |  (optional)
	country := "country_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldWorldbank(context.Background()).Indicator(indicator).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldWorldbank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldWorldbank`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldWorldbank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldWorldbankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indicator** | **string** |  | 
 **country** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldYahooFinance

> map[string]interface{} WorldWorldYahooFinance(ctx).Symbol(symbol).Execute()

Equity/index quotes (Yahoo)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	symbol := "symbol_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketsAPI.WorldWorldYahooFinance(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketsAPI.WorldWorldYahooFinance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldYahooFinance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketsAPI.WorldWorldYahooFinance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldYahooFinanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

