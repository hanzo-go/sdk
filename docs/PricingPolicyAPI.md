# \PricingPolicyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1PricingPolicy**](PricingPolicyAPI.md#CloudGetV1PricingPolicy) | **Get** /v1/pricing-policy | GetPricingPolicyAlias returns the pricing policy document at its top-level address.



## CloudGetV1PricingPolicy

> map[string]map[string]interface{} CloudGetV1PricingPolicy(ctx).Execute()

GetPricingPolicyAlias returns the pricing policy document at its top-level address.



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
	resp, r, err := apiClient.PricingPolicyAPI.CloudGetV1PricingPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingPolicyAPI.CloudGetV1PricingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingPolicy`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingPolicyAPI.CloudGetV1PricingPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingPolicyRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

