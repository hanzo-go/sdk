# \SearchWebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchDeleteWebhooks**](SearchWebhooksAPI.md#SearchDeleteWebhooks) | **Delete** /v1/search/webhooks | Delete all webhooks
[**SearchGetWebhooks**](SearchWebhooksAPI.md#SearchGetWebhooks) | **Get** /v1/search/webhooks | Get webhook configuration
[**SearchUpdateWebhooks**](SearchWebhooksAPI.md#SearchUpdateWebhooks) | **Patch** /v1/search/webhooks | Update webhook configuration



## SearchDeleteWebhooks

> map[string]SearchWebhookResultsValue SearchDeleteWebhooks(ctx).Execute()

Delete all webhooks

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
	resp, r, err := apiClient.SearchWebhooksAPI.SearchDeleteWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchWebhooksAPI.SearchDeleteWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `SearchWebhooksAPI.SearchDeleteWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteWebhooksRequest struct via the builder pattern


### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetWebhooks

> map[string]SearchWebhookResultsValue SearchGetWebhooks(ctx).Execute()

Get webhook configuration

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
	resp, r, err := apiClient.SearchWebhooksAPI.SearchGetWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchWebhooksAPI.SearchGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `SearchWebhooksAPI.SearchGetWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetWebhooksRequest struct via the builder pattern


### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateWebhooks

> map[string]SearchWebhookResultsValue SearchUpdateWebhooks(ctx).RequestBody(requestBody).Execute()

Update webhook configuration

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
	requestBody := map[string]SearchWebhookSettingsValue{"key": *openapiclient.NewSearchWebhookSettingsValue()} // map[string]SearchWebhookSettingsValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchWebhooksAPI.SearchUpdateWebhooks(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchWebhooksAPI.SearchUpdateWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `SearchWebhooksAPI.SearchUpdateWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]SearchWebhookSettingsValue**](SearchWebhookSettingsValue.md) |  | 

### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

