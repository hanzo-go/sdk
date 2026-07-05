# \RegistryWebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryCreateWebhook**](RegistryWebhooksAPI.md#RegistryCreateWebhook) | **Post** /v1/registry/webhooks | Create webhook
[**RegistryDeleteWebhook**](RegistryWebhooksAPI.md#RegistryDeleteWebhook) | **Delete** /v1/registry/webhooks/{id} | Delete webhook
[**RegistryGetWebhook**](RegistryWebhooksAPI.md#RegistryGetWebhook) | **Get** /v1/registry/webhooks/{id} | Get webhook
[**RegistryListWebhooks**](RegistryWebhooksAPI.md#RegistryListWebhooks) | **Get** /v1/registry/webhooks | List webhooks
[**RegistryUpdateWebhook**](RegistryWebhooksAPI.md#RegistryUpdateWebhook) | **Put** /v1/registry/webhooks/{id} | Update webhook



## RegistryCreateWebhook

> map[string]interface{} RegistryCreateWebhook(ctx).RegistryWebhookCreate(registryWebhookCreate).Execute()

Create webhook

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
	registryWebhookCreate := *openapiclient.NewRegistryWebhookCreate("Name_example", []string{"EventTypes_example"}, []openapiclient.RegistryWebhookCreateTargetsInner{*openapiclient.NewRegistryWebhookCreateTargetsInner("Type_example", "Address_example")}) // RegistryWebhookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryWebhooksAPI.RegistryCreateWebhook(context.Background()).RegistryWebhookCreate(registryWebhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryWebhooksAPI.RegistryCreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryCreateWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryWebhooksAPI.RegistryCreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryWebhookCreate** | [**RegistryWebhookCreate**](RegistryWebhookCreate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDeleteWebhook

> map[string]interface{} RegistryDeleteWebhook(ctx, id).Execute()

Delete webhook

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryWebhooksAPI.RegistryDeleteWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryWebhooksAPI.RegistryDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryDeleteWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryWebhooksAPI.RegistryDeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RegistryGetWebhook

> RegistryWebhook RegistryGetWebhook(ctx, id).Execute()

Get webhook

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryWebhooksAPI.RegistryGetWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryWebhooksAPI.RegistryGetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryGetWebhook`: RegistryWebhook
	fmt.Fprintf(os.Stdout, "Response from `RegistryWebhooksAPI.RegistryGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryWebhook**](RegistryWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListWebhooks

> []RegistryWebhook RegistryListWebhooks(ctx).ProjectId(projectId).Page(page).PageSize(pageSize).Execute()

List webhooks

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
	projectId := int32(56) // int32 | Filter by project ID (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryWebhooksAPI.RegistryListWebhooks(context.Background()).ProjectId(projectId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryWebhooksAPI.RegistryListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListWebhooks`: []RegistryWebhook
	fmt.Fprintf(os.Stdout, "Response from `RegistryWebhooksAPI.RegistryListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32** | Filter by project ID | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]

### Return type

[**[]RegistryWebhook**](RegistryWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryUpdateWebhook

> map[string]interface{} RegistryUpdateWebhook(ctx, id).RegistryWebhookCreate(registryWebhookCreate).Execute()

Update webhook

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
	id := int32(56) // int32 | 
	registryWebhookCreate := *openapiclient.NewRegistryWebhookCreate("Name_example", []string{"EventTypes_example"}, []openapiclient.RegistryWebhookCreateTargetsInner{*openapiclient.NewRegistryWebhookCreateTargetsInner("Type_example", "Address_example")}) // RegistryWebhookCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryWebhooksAPI.RegistryUpdateWebhook(context.Background(), id).RegistryWebhookCreate(registryWebhookCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryWebhooksAPI.RegistryUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryUpdateWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryWebhooksAPI.RegistryUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registryWebhookCreate** | [**RegistryWebhookCreate**](RegistryWebhookCreate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

