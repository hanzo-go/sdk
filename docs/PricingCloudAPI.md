# \PricingCloudAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingGetCloud**](PricingCloudAPI.md#PricingGetCloud) | **Get** /v1/pricing/cloud | Cloud VM plans, regions, and storage
[**PricingGetStoragePricing**](PricingCloudAPI.md#PricingGetStoragePricing) | **Get** /v1/pricing/cloud/storage | Block storage pricing
[**PricingListCloudPlans**](PricingCloudAPI.md#PricingListCloudPlans) | **Get** /v1/pricing/cloud/plans | Cloud VM plans
[**PricingListCloudRegions**](PricingCloudAPI.md#PricingListCloudRegions) | **Get** /v1/pricing/cloud/regions | Available cloud regions



## PricingGetCloud

> PricingCloudResponse PricingGetCloud(ctx).Execute()

Cloud VM plans, regions, and storage



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
	resp, r, err := apiClient.PricingCloudAPI.PricingGetCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingCloudAPI.PricingGetCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetCloud`: PricingCloudResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingCloudAPI.PricingGetCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetCloudRequest struct via the builder pattern


### Return type

[**PricingCloudResponse**](PricingCloudResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetStoragePricing

> PricingBlockStoragePricing PricingGetStoragePricing(ctx).Execute()

Block storage pricing



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
	resp, r, err := apiClient.PricingCloudAPI.PricingGetStoragePricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingCloudAPI.PricingGetStoragePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetStoragePricing`: PricingBlockStoragePricing
	fmt.Fprintf(os.Stdout, "Response from `PricingCloudAPI.PricingGetStoragePricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetStoragePricingRequest struct via the builder pattern


### Return type

[**PricingBlockStoragePricing**](PricingBlockStoragePricing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListCloudPlans

> PricingCloudPlansResponse PricingListCloudPlans(ctx).Execute()

Cloud VM plans



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
	resp, r, err := apiClient.PricingCloudAPI.PricingListCloudPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingCloudAPI.PricingListCloudPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListCloudPlans`: PricingCloudPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingCloudAPI.PricingListCloudPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListCloudPlansRequest struct via the builder pattern


### Return type

[**PricingCloudPlansResponse**](PricingCloudPlansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListCloudRegions

> PricingCloudRegionsResponse PricingListCloudRegions(ctx).Execute()

Available cloud regions



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
	resp, r, err := apiClient.PricingCloudAPI.PricingListCloudRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingCloudAPI.PricingListCloudRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListCloudRegions`: PricingCloudRegionsResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingCloudAPI.PricingListCloudRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListCloudRegionsRequest struct via the builder pattern


### Return type

[**PricingCloudRegionsResponse**](PricingCloudRegionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

