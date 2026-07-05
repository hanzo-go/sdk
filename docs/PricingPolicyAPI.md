# \PricingPolicyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingGetPricingPolicy**](PricingPolicyAPI.md#PricingGetPricingPolicy) | **Get** /v1/pricing/policy | Pricing policy and revenue sharing



## PricingGetPricingPolicy

> PricingPricingPolicy PricingGetPricingPolicy(ctx).Execute()

Pricing policy and revenue sharing



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
	resp, r, err := apiClient.PricingPolicyAPI.PricingGetPricingPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPolicyAPI.PricingGetPricingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetPricingPolicy`: PricingPricingPolicy
	fmt.Fprintf(os.Stdout, "Response from `PricingPolicyAPI.PricingGetPricingPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetPricingPolicyRequest struct via the builder pattern


### Return type

[**PricingPricingPolicy**](PricingPricingPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

