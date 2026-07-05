# \IamWebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddWebhook**](IamWebhooksAPI.md#IamApiControllerAddWebhook) | **Post** /v1/iam/webhooks | Api Controller Add Webhook
[**IamApiControllerDeleteWebhook**](IamWebhooksAPI.md#IamApiControllerDeleteWebhook) | **Delete** /v1/iam/webhooks/{id} | Api Controller Delete Webhook
[**IamApiControllerGetWebhook**](IamWebhooksAPI.md#IamApiControllerGetWebhook) | **Get** /v1/iam/webhooks/{id} | Api Controller Get Webhook
[**IamApiControllerGetWebhooks**](IamWebhooksAPI.md#IamApiControllerGetWebhooks) | **Get** /v1/iam/webhooks | Api Controller Get Webhooks
[**IamApiControllerUpdateWebhook**](IamWebhooksAPI.md#IamApiControllerUpdateWebhook) | **Put** /v1/iam/webhooks/{id} | Api Controller Update Webhook



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
	resp, r, err := apiClient.IamWebhooksAPI.IamApiControllerAddWebhook(context.Background()).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamWebhooksAPI.IamApiControllerAddWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamWebhooksAPI.IamApiControllerAddWebhook`: %v\n", resp)
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
	resp, r, err := apiClient.IamWebhooksAPI.IamApiControllerDeleteWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamWebhooksAPI.IamApiControllerDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamWebhooksAPI.IamApiControllerDeleteWebhook`: %v\n", resp)
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
	resp, r, err := apiClient.IamWebhooksAPI.IamApiControllerGetWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamWebhooksAPI.IamApiControllerGetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhook`: IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `IamWebhooksAPI.IamApiControllerGetWebhook`: %v\n", resp)
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
	resp, r, err := apiClient.IamWebhooksAPI.IamApiControllerGetWebhooks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamWebhooksAPI.IamApiControllerGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhooks`: []IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `IamWebhooksAPI.IamApiControllerGetWebhooks`: %v\n", resp)
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
	resp, r, err := apiClient.IamWebhooksAPI.IamApiControllerUpdateWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamWebhooksAPI.IamApiControllerUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamWebhooksAPI.IamApiControllerUpdateWebhook`: %v\n", resp)
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

