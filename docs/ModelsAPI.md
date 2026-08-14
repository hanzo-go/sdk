# \ModelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetModels**](ModelsAPI.md#GetModels) | **Get** /v1/models | Returns the list of available models from the routing table.
[**GetModelsByModelAccess**](ModelsAPI.md#GetModelsByModelAccess) | **Get** /v1/models/{model}/access | Returns the caller&#39;s own standing for a gated model: \&quot;granted\&quot;, \&quot;requested\&quot;, or empty when they have never asked.
[**GetModelsProviders**](ModelsAPI.md#GetModelsProviders) | **Get** /v1/models/providers | Public, secret-free list of the providers serving the models that GET /v1/models lists — the same source, projected.
[**PostModelsByModelAccess**](ModelsAPI.md#PostModelsByModelAccess) | **Post** /v1/models/{model}/access | Records the caller&#39;s waitlist request for a gated model and answers their new standing.



## GetModels

> GetModels(ctx).Execute()

Returns the list of available models from the routing table.



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
	r, err := apiClient.ModelsAPI.GetModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsByModelAccess

> GetModelsByModelAccess(ctx, model).Execute()

Returns the caller's own standing for a gated model: \"granted\", \"requested\", or empty when they have never asked.



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
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModelsAPI.GetModelsByModelAccess(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelsByModelAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsByModelAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelsProviders

> GetModelsProviders(ctx).Execute()

Public, secret-free list of the providers serving the models that GET /v1/models lists — the same source, projected.



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
	r, err := apiClient.ModelsAPI.GetModelsProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GetModelsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelsProvidersRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostModelsByModelAccess

> PostModelsByModelAccess(ctx, model).Execute()

Records the caller's waitlist request for a gated model and answers their new standing.



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
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModelsAPI.PostModelsByModelAccess(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.PostModelsByModelAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostModelsByModelAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

