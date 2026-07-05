# \PlatformBillingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformBillingGetBalance**](PlatformBillingAPI.md#PlatformBillingGetBalance) | **Get** /v1/platform/billing/getBalance | Get account balance
[**PlatformBillingGetPlans**](PlatformBillingAPI.md#PlatformBillingGetPlans) | **Get** /v1/platform/billing/getPlans | List subscription plans



## PlatformBillingGetBalance

> PlatformTRPCResult PlatformBillingGetBalance(ctx).Execute()

Get account balance

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
	resp, r, err := apiClient.PlatformBillingAPI.PlatformBillingGetBalance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformBillingAPI.PlatformBillingGetBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformBillingGetBalance`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformBillingAPI.PlatformBillingGetBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformBillingGetBalanceRequest struct via the builder pattern


### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformBillingGetPlans

> PlatformTRPCResult PlatformBillingGetPlans(ctx).Execute()

List subscription plans

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
	resp, r, err := apiClient.PlatformBillingAPI.PlatformBillingGetPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformBillingAPI.PlatformBillingGetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformBillingGetPlans`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformBillingAPI.PlatformBillingGetPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformBillingGetPlansRequest struct via the builder pattern


### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

