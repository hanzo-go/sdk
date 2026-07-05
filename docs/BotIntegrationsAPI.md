# \BotIntegrationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotGetIntegration**](BotIntegrationsAPI.md#BotGetIntegration) | **Get** /v1/bot/integrations/{slug} | Get integration detail with latest version
[**BotListIntegrations**](BotIntegrationsAPI.md#BotListIntegrations) | **Get** /v1/bot/integrations | List integrations (paginated)



## BotGetIntegration

> BotGetIntegration200Response BotGetIntegration(ctx, slug).Execute()

Get integration detail with latest version

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotIntegrationsAPI.BotGetIntegration(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotIntegrationsAPI.BotGetIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetIntegration`: BotGetIntegration200Response
	fmt.Fprintf(os.Stdout, "Response from `BotIntegrationsAPI.BotGetIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotGetIntegration200Response**](BotGetIntegration200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListIntegrations

> BotListIntegrations200Response BotListIntegrations(ctx).Sort(sort).Limit(limit).Cursor(cursor).Execute()

List integrations (paginated)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sort := "sort_example" // string |  (optional) (default to "updated")
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := time.Now() // time.Time |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotIntegrationsAPI.BotListIntegrations(context.Background()).Sort(sort).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotIntegrationsAPI.BotListIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListIntegrations`: BotListIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `BotIntegrationsAPI.BotListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** |  | [default to &quot;updated&quot;]
 **limit** | **int32** |  | [default to 50]
 **cursor** | **time.Time** |  | 

### Return type

[**BotListIntegrations200Response**](BotListIntegrations200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

