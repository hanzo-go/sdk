# \SubscriptionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCancelSubscription**](SubscriptionsAPI.md#CommerceCancelSubscription) | **Delete** /v1/commerce/subscribe/{subscriptionid} | Cancel subscription
[**CommerceCreateSubscription**](SubscriptionsAPI.md#CommerceCreateSubscription) | **Post** /v1/commerce/subscribe | Create subscription
[**CommerceGetSubscription**](SubscriptionsAPI.md#CommerceGetSubscription) | **Get** /v1/commerce/subscribe/{subscriptionid} | Get subscription
[**CommerceUpdateSubscription**](SubscriptionsAPI.md#CommerceUpdateSubscription) | **Patch** /v1/commerce/subscribe/{subscriptionid} | Update subscription
[**PricingListBlockchainPlans**](SubscriptionsAPI.md#PricingListBlockchainPlans) | **Get** /v1/pricing/blockchain | Blockchain / RPC plans
[**PricingListPlans**](SubscriptionsAPI.md#PricingListPlans) | **Get** /v1/pricing/plans | Subscription plans
[**PricingListSubscriptions**](SubscriptionsAPI.md#PricingListSubscriptions) | **Get** /v1/pricing/subscriptions | Subscription plans



## CommerceCancelSubscription

> CommerceSubscription CommerceCancelSubscription(ctx, subscriptionid).Execute()

Cancel subscription

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
	subscriptionid := "subscriptionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CommerceCancelSubscription(context.Background(), subscriptionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CommerceCancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCancelSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CommerceCancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceSubscription**](CommerceSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateSubscription

> CommerceSubscription CommerceCreateSubscription(ctx).CommerceSubscriptionRequest(commerceSubscriptionRequest).Execute()

Create subscription

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
	commerceSubscriptionRequest := *openapiclient.NewCommerceSubscriptionRequest("PlanId_example") // CommerceSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CommerceCreateSubscription(context.Background()).CommerceSubscriptionRequest(commerceSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CommerceCreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CommerceCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceSubscriptionRequest** | [**CommerceSubscriptionRequest**](CommerceSubscriptionRequest.md) |  | 

### Return type

[**CommerceSubscription**](CommerceSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetSubscription

> CommerceSubscription CommerceGetSubscription(ctx, subscriptionid).Execute()

Get subscription

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
	subscriptionid := "subscriptionid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CommerceGetSubscription(context.Background(), subscriptionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CommerceGetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CommerceGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceSubscription**](CommerceSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateSubscription

> CommerceSubscription CommerceUpdateSubscription(ctx, subscriptionid).CommerceSubscription(commerceSubscription).Execute()

Update subscription

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
	subscriptionid := "subscriptionid_example" // string | 
	commerceSubscription := *openapiclient.NewCommerceSubscription() // CommerceSubscription | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CommerceUpdateSubscription(context.Background(), subscriptionid).CommerceSubscription(commerceSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CommerceUpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CommerceUpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceSubscription** | [**CommerceSubscription**](CommerceSubscription.md) |  | 

### Return type

[**CommerceSubscription**](CommerceSubscription.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.SubscriptionsAPI.PricingListBlockchainPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PricingListBlockchainPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListBlockchainPlans`: PricingListBlockchainPlans200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PricingListBlockchainPlans`: %v\n", resp)
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
	resp, r, err := apiClient.SubscriptionsAPI.PricingListPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PricingListPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListPlans`: PricingSubscriptionPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PricingListPlans`: %v\n", resp)
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
	resp, r, err := apiClient.SubscriptionsAPI.PricingListSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PricingListSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListSubscriptions`: PricingSubscriptionPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PricingListSubscriptions`: %v\n", resp)
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

