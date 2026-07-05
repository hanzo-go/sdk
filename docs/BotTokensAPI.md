# \BotTokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotCreateToken**](BotTokensAPI.md#BotCreateToken) | **Post** /v1/bot/tokens | Create a new API token
[**BotListTokens**](BotTokensAPI.md#BotListTokens) | **Get** /v1/bot/tokens | List current user&#39;s API tokens
[**BotRevokeToken**](BotTokensAPI.md#BotRevokeToken) | **Delete** /v1/bot/tokens/{id} | Revoke an API token



## BotCreateToken

> BotCreateToken200Response BotCreateToken(ctx).BotCreateTokenRequest(botCreateTokenRequest).Execute()

Create a new API token

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
	botCreateTokenRequest := *openapiclient.NewBotCreateTokenRequest("Label_example") // BotCreateTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotTokensAPI.BotCreateToken(context.Background()).BotCreateTokenRequest(botCreateTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotTokensAPI.BotCreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotCreateToken`: BotCreateToken200Response
	fmt.Fprintf(os.Stdout, "Response from `BotTokensAPI.BotCreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **botCreateTokenRequest** | [**BotCreateTokenRequest**](BotCreateTokenRequest.md) |  | 

### Return type

[**BotCreateToken200Response**](BotCreateToken200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListTokens

> BotListTokens200Response BotListTokens(ctx).Execute()

List current user's API tokens

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
	resp, r, err := apiClient.BotTokensAPI.BotListTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotTokensAPI.BotListTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListTokens`: BotListTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `BotTokensAPI.BotListTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBotListTokensRequest struct via the builder pattern


### Return type

[**BotListTokens200Response**](BotListTokens200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotRevokeToken

> AnalyticsHeartbeat200Response BotRevokeToken(ctx, id).Execute()

Revoke an API token

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotTokensAPI.BotRevokeToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotTokensAPI.BotRevokeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotRevokeToken`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `BotTokensAPI.BotRevokeToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

