# \PricingInfrastructureAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingGetComputePresets**](PricingInfrastructureAPI.md#PricingGetComputePresets) | **Get** /v1/pricing/compute/presets | Curated compute presets
[**PricingGetComputePricing**](PricingInfrastructureAPI.md#PricingGetComputePricing) | **Get** /v1/pricing/compute | Compute tiers
[**PricingListGpuTiers**](PricingInfrastructureAPI.md#PricingListGpuTiers) | **Get** /v1/pricing/gpu | GPU tier pricing
[**PricingListTools**](PricingInfrastructureAPI.md#PricingListTools) | **Get** /v1/pricing/tools | Tool pricing



## PricingGetComputePresets

> PricingGetComputePresets200Response PricingGetComputePresets(ctx).Execute()

Curated compute presets



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
	resp, r, err := apiClient.PricingInfrastructureAPI.PricingGetComputePresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingInfrastructureAPI.PricingGetComputePresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetComputePresets`: PricingGetComputePresets200Response
	fmt.Fprintf(os.Stdout, "Response from `PricingInfrastructureAPI.PricingGetComputePresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetComputePresetsRequest struct via the builder pattern


### Return type

[**PricingGetComputePresets200Response**](PricingGetComputePresets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetComputePricing

> map[string]interface{} PricingGetComputePricing(ctx).Execute()

Compute tiers



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
	resp, r, err := apiClient.PricingInfrastructureAPI.PricingGetComputePricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingInfrastructureAPI.PricingGetComputePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetComputePricing`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingInfrastructureAPI.PricingGetComputePricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetComputePricingRequest struct via the builder pattern


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


## PricingListGpuTiers

> PricingGpuTiersResponse PricingListGpuTiers(ctx).Execute()

GPU tier pricing



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
	resp, r, err := apiClient.PricingInfrastructureAPI.PricingListGpuTiers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingInfrastructureAPI.PricingListGpuTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListGpuTiers`: PricingGpuTiersResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingInfrastructureAPI.PricingListGpuTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListGpuTiersRequest struct via the builder pattern


### Return type

[**PricingGpuTiersResponse**](PricingGpuTiersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListTools

> PricingToolsResponse PricingListTools(ctx).Execute()

Tool pricing



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
	resp, r, err := apiClient.PricingInfrastructureAPI.PricingListTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingInfrastructureAPI.PricingListTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListTools`: PricingToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingInfrastructureAPI.PricingListTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListToolsRequest struct via the builder pattern


### Return type

[**PricingToolsResponse**](PricingToolsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

