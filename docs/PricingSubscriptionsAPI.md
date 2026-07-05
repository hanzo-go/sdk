# \PricingSubscriptionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingListBlockchainPlans**](PricingSubscriptionsAPI.md#PricingListBlockchainPlans) | **Get** /v1/pricing/blockchain | Blockchain / RPC plans
[**PricingListPlans**](PricingSubscriptionsAPI.md#PricingListPlans) | **Get** /v1/pricing/plans | Subscription plans
[**PricingListSubscriptions**](PricingSubscriptionsAPI.md#PricingListSubscriptions) | **Get** /v1/pricing/subscriptions | Subscription plans



## PricingListBlockchainPlans

> PricingListBlockchainPlans200Response PricingListBlockchainPlans(ctx).Execute()

Blockchain / RPC plans



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
	resp, r, err := apiClient.PricingSubscriptionsAPI.PricingListBlockchainPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingSubscriptionsAPI.PricingListBlockchainPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListBlockchainPlans`: PricingListBlockchainPlans200Response
	fmt.Fprintf(os.Stdout, "Response from `PricingSubscriptionsAPI.PricingListBlockchainPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListBlockchainPlansRequest struct via the builder pattern


### Return type

[**PricingListBlockchainPlans200Response**](PricingListBlockchainPlans200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListPlans

> PricingSubscriptionPlansResponse PricingListPlans(ctx).Execute()

Subscription plans



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
	resp, r, err := apiClient.PricingSubscriptionsAPI.PricingListPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingSubscriptionsAPI.PricingListPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListPlans`: PricingSubscriptionPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingSubscriptionsAPI.PricingListPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListPlansRequest struct via the builder pattern


### Return type

[**PricingSubscriptionPlansResponse**](PricingSubscriptionPlansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListSubscriptions

> PricingSubscriptionPlansResponse PricingListSubscriptions(ctx).Execute()

Subscription plans



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
	resp, r, err := apiClient.PricingSubscriptionsAPI.PricingListSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingSubscriptionsAPI.PricingListSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListSubscriptions`: PricingSubscriptionPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `PricingSubscriptionsAPI.PricingListSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListSubscriptionsRequest struct via the builder pattern


### Return type

[**PricingSubscriptionPlansResponse**](PricingSubscriptionPlansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

