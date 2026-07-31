# \BillingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingBillingBalance**](BillingAPI.md#BillingBillingBalance) | **Get** /v1/billing/balance | Prepaid credit balance
[**BillingBillingGpuCharge**](BillingAPI.md#BillingBillingGpuCharge) | **Post** /v1/billing/gpu-charge | Prepay-only GPU charge
[**BillingBillingGpuEligibility**](BillingAPI.md#BillingBillingGpuEligibility) | **Get** /v1/billing/gpu-eligibility | GPU launch eligibility gate
[**BillingBillingPaymentMethods**](BillingAPI.md#BillingBillingPaymentMethods) | **Get** /v1/billing/payment-methods | Saved payment methods (masked)
[**BillingBillingUsage**](BillingAPI.md#BillingBillingUsage) | **Get** /v1/billing/usage | Per-request usage ledger



## BillingBillingBalance

> BillingBalance BillingBillingBalance(ctx).Currency(currency).Execute()

Prepaid credit balance



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
	currency := "currency_example" // string | Optional currency filter (default usd) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingBillingBalance(context.Background()).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingBillingBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingBalance`: BillingBalance
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingBillingBalance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Optional currency filter (default usd) | 

### Return type

[**BillingBalance**](BillingBalance.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingGpuCharge

> map[string]interface{} BillingBillingGpuCharge(ctx).BillingGpuChargeRequest(billingGpuChargeRequest).Execute()

Prepay-only GPU charge



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
	billingGpuChargeRequest := *openapiclient.NewBillingGpuChargeRequest() // BillingGpuChargeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingBillingGpuCharge(context.Background()).BillingGpuChargeRequest(billingGpuChargeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingBillingGpuCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingGpuCharge`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingBillingGpuCharge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingGpuChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingGpuChargeRequest** | [**BillingGpuChargeRequest**](BillingGpuChargeRequest.md) |  | 

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


## BillingBillingGpuEligibility

> BillingGpuEligibility BillingBillingGpuEligibility(ctx).AmountCents(amountCents).MinPrepaidCents(minPrepaidCents).Currency(currency).Execute()

GPU launch eligibility gate



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
	amountCents := int64(789) // int64 | Immediate charge to test against (optional)
	minPrepaidCents := int64(789) // int64 | 24h-minimum prepaid floor (optional)
	currency := "currency_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingBillingGpuEligibility(context.Background()).AmountCents(amountCents).MinPrepaidCents(minPrepaidCents).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingBillingGpuEligibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingGpuEligibility`: BillingGpuEligibility
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingBillingGpuEligibility`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingGpuEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **amountCents** | **int64** | Immediate charge to test against | 
 **minPrepaidCents** | **int64** | 24h-minimum prepaid floor | 
 **currency** | **string** |  | 

### Return type

[**BillingGpuEligibility**](BillingGpuEligibility.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingPaymentMethods

> BillingPaymentMethods BillingBillingPaymentMethods(ctx).Execute()

Saved payment methods (masked)



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
	resp, r, err := apiClient.BillingAPI.BillingBillingPaymentMethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingBillingPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingPaymentMethods`: BillingPaymentMethods
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingBillingPaymentMethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingPaymentMethodsRequest struct via the builder pattern


### Return type

[**BillingPaymentMethods**](BillingPaymentMethods.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingUsage

> BillingUsageLedger BillingBillingUsage(ctx).Start(start).End(end).Execute()

Per-request usage ledger



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
	start := "start_example" // string | Optional server-side window start (optional)
	end := "end_example" // string | Optional server-side window end (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.BillingBillingUsage(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.BillingBillingUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingUsage`: BillingUsageLedger
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.BillingBillingUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **string** | Optional server-side window start | 
 **end** | **string** | Optional server-side window end | 

### Return type

[**BillingUsageLedger**](BillingUsageLedger.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

