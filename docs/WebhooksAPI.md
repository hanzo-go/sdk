# \WebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddWebhook**](WebhooksAPI.md#IamApiControllerAddWebhook) | **Post** /v1/iam/webhooks | Api Controller Add Webhook
[**IamApiControllerDeleteWebhook**](WebhooksAPI.md#IamApiControllerDeleteWebhook) | **Delete** /v1/iam/webhooks/{id} | Api Controller Delete Webhook
[**IamApiControllerGetWebhook**](WebhooksAPI.md#IamApiControllerGetWebhook) | **Get** /v1/iam/webhooks/{id} | Api Controller Get Webhook
[**IamApiControllerGetWebhooks**](WebhooksAPI.md#IamApiControllerGetWebhooks) | **Get** /v1/iam/webhooks | Api Controller Get Webhooks
[**IamApiControllerUpdateWebhook**](WebhooksAPI.md#IamApiControllerUpdateWebhook) | **Put** /v1/iam/webhooks/{id} | Api Controller Update Webhook
[**SearchDeleteWebhooks**](WebhooksAPI.md#SearchDeleteWebhooks) | **Delete** /v1/search/webhooks | Delete all webhooks
[**SearchGetWebhooks**](WebhooksAPI.md#SearchGetWebhooks) | **Get** /v1/search/webhooks | Get webhook configuration
[**SearchUpdateWebhooks**](WebhooksAPI.md#SearchUpdateWebhooks) | **Patch** /v1/search/webhooks | Update webhook configuration



## IamApiControllerAddWebhook

> IamControllersResponse IamApiControllerAddWebhook(ctx).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Add Webhook



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
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerAddWebhook(context.Background()).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerAddWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerAddWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteWebhook

> IamControllersResponse IamApiControllerDeleteWebhook(ctx, id).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Delete Webhook



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerDeleteWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerDeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWebhook

> IamObjectWebhook IamApiControllerGetWebhook(ctx, id).Execute()

Api Controller Get Webhook



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
	id := "id_example" // string | The id ( owner/name ) of the webhook (default to "built-in/admin")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerGetWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerGetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhook`: IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the webhook | [default to &quot;built-in/admin&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectWebhook**](IamObjectWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWebhooks

> []IamObjectWebhook IamApiControllerGetWebhooks(ctx).Owner(owner).Execute()

Api Controller Get Webhooks



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
	owner := "owner_example" // string | The owner of webhooks (default to "built-in/admin")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerGetWebhooks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhooks`: []IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerGetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of webhooks | [default to &quot;built-in/admin&quot;]

### Return type

[**[]IamObjectWebhook**](IamObjectWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateWebhook

> IamControllersResponse IamApiControllerUpdateWebhook(ctx, id).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Update Webhook



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
	id := "id_example" // string | The id ( owner/name ) of the webhook (default to "built-in/admin")
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerUpdateWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the webhook | [default to &quot;built-in/admin&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.WebhooksAPI.SearchDeleteWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchDeleteWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchDeleteWebhooks`: %v\n", resp)
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
	resp, r, err := apiClient.WebhooksAPI.SearchGetWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchGetWebhooks`: %v\n", resp)
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
	resp, r, err := apiClient.WebhooksAPI.SearchUpdateWebhooks(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchUpdateWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchUpdateWebhooks`: %v\n", resp)
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

