# \MarketAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketChains**](MarketAPI.md#GetMarketChains) | **Get** /v1/market/chains | Answers every chain this deployment can read, what is deployed on each, and what its automated market maker amounts to.
[**GetMarketPools**](MarketAPI.md#GetMarketPools) | **Get** /v1/market/pools | Answers the automated market makers on one chain: their two tokens, their fee tier, and what has moved through each.
[**GetMarketSurvey**](MarketAPI.md#GetMarketSurvey) | **Get** /v1/market/survey | Answers which of the four settlement precompiles carry code on one chain.
[**GetMarketToken**](MarketAPI.md#GetMarketToken) | **Get** /v1/market/token | Answers one token&#39;s daily history — open, high, low, close, price and volume per UTC day, oldest first.
[**GetMarketTokens**](MarketAPI.md#GetMarketTokens) | **Get** /v1/market/tokens | Answers the tokens one chain&#39;s indexer has seen, with the decimals a caller needs to read any amount of one correctly.



## GetMarketChains

> Roster GetMarketChains(ctx).Execute()

Answers every chain this deployment can read, what is deployed on each, and what its automated market maker amounts to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketChains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketChains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketChains`: Roster
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketChainsRequest struct via the builder pattern


### Return type

[**Roster**](Roster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketPools

> Pools GetMarketPools(ctx).Chain(chain).Execute()

Answers the automated market makers on one chain: their two tokens, their fee tier, and what has moved through each.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	chain := "chain_example" // string | Chain is the chain's slug — `cchain`, `zoo` — as `chains` reports it. It is the indexer's word for the chain and NOT the chain id: `96369`, `C` and `c-chain` all name nothing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketPools(context.Background()).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketPools`: Pools
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | **string** | Chain is the chain&#39;s slug — &#x60;cchain&#x60;, &#x60;zoo&#x60; — as &#x60;chains&#x60; reports it. It is the indexer&#39;s word for the chain and NOT the chain id: &#x60;96369&#x60;, &#x60;C&#x60; and &#x60;c-chain&#x60; all name nothing. | 

### Return type

[**Pools**](Pools.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketSurvey

> Survey GetMarketSurvey(ctx).Chain(chain).Execute()

Answers which of the four settlement precompiles carry code on one chain.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	chain := "chain_example" // string | Chain is the chain's slug — `cchain`, `zoo` — as `chains` reports it. It is the indexer's word for the chain and NOT the chain id: `96369`, `C` and `c-chain` all name nothing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketSurvey(context.Background()).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketSurvey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketSurvey`: Survey
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketSurvey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketSurveyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | **string** | Chain is the chain&#39;s slug — &#x60;cchain&#x60;, &#x60;zoo&#x60; — as &#x60;chains&#x60; reports it. It is the indexer&#39;s word for the chain and NOT the chain id: &#x60;96369&#x60;, &#x60;C&#x60; and &#x60;c-chain&#x60; all name nothing. | 

### Return type

[**Survey**](Survey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketToken

> History GetMarketToken(ctx).Chain(chain).At(at).Execute()

Answers one token's daily history — open, high, low, close, price and volume per UTC day, oldest first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	chain := "chain_example" // string |  (optional)
	at := "at_example" // string | At is the token's contract address. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketToken(context.Background()).Chain(chain).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketToken`: History
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | **string** |  | 
 **at** | **string** | At is the token&#39;s contract address. | 

### Return type

[**History**](History.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketTokens

> Tokens GetMarketTokens(ctx).Chain(chain).Execute()

Answers the tokens one chain's indexer has seen, with the decimals a caller needs to read any amount of one correctly.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	chain := "chain_example" // string | Chain is the chain's slug — `cchain`, `zoo` — as `chains` reports it. It is the indexer's word for the chain and NOT the chain id: `96369`, `C` and `c-chain` all name nothing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.GetMarketTokens(context.Background()).Chain(chain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.GetMarketTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketTokens`: Tokens
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.GetMarketTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chain** | **string** | Chain is the chain&#39;s slug — &#x60;cchain&#x60;, &#x60;zoo&#x60; — as &#x60;chains&#x60; reports it. It is the indexer&#39;s word for the chain and NOT the chain id: &#x60;96369&#x60;, &#x60;C&#x60; and &#x60;c-chain&#x60; all name nothing. | 

### Return type

[**Tokens**](Tokens.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

