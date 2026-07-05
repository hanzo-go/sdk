# \PricingModelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingGetFullPricing**](PricingModelsAPI.md#PricingGetFullPricing) | **Get** /v1/pricing | Full pricing data
[**PricingGetModel**](PricingModelsAPI.md#PricingGetModel) | **Get** /v1/pricing/model/{name} | Single model lookup
[**PricingGetPricingSummary**](PricingModelsAPI.md#PricingGetPricingSummary) | **Get** /v1/pricing/summary | Model counts and provider breakdown
[**PricingListFeaturedModels**](PricingModelsAPI.md#PricingListFeaturedModels) | **Get** /v1/pricing/featured | Featured third-party models
[**PricingListFreeModels**](PricingModelsAPI.md#PricingListFreeModels) | **Get** /v1/pricing/free | Free models only
[**PricingListModels**](PricingModelsAPI.md#PricingListModels) | **Get** /v1/pricing/models | List all AI models (OpenAI-compatible)
[**PricingListProviders**](PricingModelsAPI.md#PricingListProviders) | **Get** /v1/pricing/providers | Provider breakdown



## PricingGetFullPricing

> PricingFullPricingResponse PricingGetFullPricing(ctx).Execute()

Full pricing data



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
	resp, r, err := apiClient.PricingModelsAPI.PricingGetFullPricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingGetFullPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetFullPricing`: PricingFullPricingResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingGetFullPricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetFullPricingRequest struct via the builder pattern


### Return type

[**PricingFullPricingResponse**](PricingFullPricingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetModel

> PricingModel PricingGetModel(ctx, name).Execute()

Single model lookup



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
	name := "zen4" // string | Model name or ID (case-insensitive)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingModelsAPI.PricingGetModel(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetModel`: PricingModel
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Model name or ID (case-insensitive) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PricingModel**](PricingModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetPricingSummary

> PricingSummaryResponse PricingGetPricingSummary(ctx).Execute()

Model counts and provider breakdown



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
	resp, r, err := apiClient.PricingModelsAPI.PricingGetPricingSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingGetPricingSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetPricingSummary`: PricingSummaryResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingGetPricingSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetPricingSummaryRequest struct via the builder pattern


### Return type

[**PricingSummaryResponse**](PricingSummaryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListFeaturedModels

> PricingPricingModelsResponse PricingListFeaturedModels(ctx).Execute()

Featured third-party models



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
	resp, r, err := apiClient.PricingModelsAPI.PricingListFeaturedModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingListFeaturedModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListFeaturedModels`: PricingPricingModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingListFeaturedModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListFeaturedModelsRequest struct via the builder pattern


### Return type

[**PricingPricingModelsResponse**](PricingPricingModelsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListFreeModels

> PricingPricingModelsResponse PricingListFreeModels(ctx).Execute()

Free models only



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
	resp, r, err := apiClient.PricingModelsAPI.PricingListFreeModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingListFreeModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListFreeModels`: PricingPricingModelsResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingListFreeModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListFreeModelsRequest struct via the builder pattern


### Return type

[**PricingPricingModelsResponse**](PricingPricingModelsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListModels

> PricingModelListResponse PricingListModels(ctx).Execute()

List all AI models (OpenAI-compatible)



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
	resp, r, err := apiClient.PricingModelsAPI.PricingListModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListModels`: PricingModelListResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingListModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListModelsRequest struct via the builder pattern


### Return type

[**PricingModelListResponse**](PricingModelListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListProviders

> PricingProvidersResponse PricingListProviders(ctx).Execute()

Provider breakdown



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
	resp, r, err := apiClient.PricingModelsAPI.PricingListProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingModelsAPI.PricingListProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListProviders`: PricingProvidersResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingModelsAPI.PricingListProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListProvidersRequest struct via the builder pattern


### Return type

[**PricingProvidersResponse**](PricingProvidersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

