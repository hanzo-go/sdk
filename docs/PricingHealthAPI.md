# \PricingHealthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingTriggerSync**](PricingHealthAPI.md#PricingTriggerSync) | **Post** /v1/pricing/sync | Trigger manual sync



## PricingTriggerSync

> PricingTriggerSync200Response PricingTriggerSync(ctx).Execute()

Trigger manual sync



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
	resp, r, err := apiClient.PricingHealthAPI.PricingTriggerSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingHealthAPI.PricingTriggerSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingTriggerSync`: PricingTriggerSync200Response
	fmt.Fprintf(os.Stdout, "Response from `PricingHealthAPI.PricingTriggerSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingTriggerSyncRequest struct via the builder pattern


### Return type

[**PricingTriggerSync200Response**](PricingTriggerSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

