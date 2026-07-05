# \CommerceSubscriptionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceCancelSubscription**](CommerceSubscriptionsAPI.md#CommerceCancelSubscription) | **Delete** /v1/commerce/subscribe/{subscriptionid} | Cancel subscription
[**CommerceCreateSubscription**](CommerceSubscriptionsAPI.md#CommerceCreateSubscription) | **Post** /v1/commerce/subscribe | Create subscription
[**CommerceGetSubscription**](CommerceSubscriptionsAPI.md#CommerceGetSubscription) | **Get** /v1/commerce/subscribe/{subscriptionid} | Get subscription
[**CommerceUpdateSubscription**](CommerceSubscriptionsAPI.md#CommerceUpdateSubscription) | **Patch** /v1/commerce/subscribe/{subscriptionid} | Update subscription



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
	resp, r, err := apiClient.CommerceSubscriptionsAPI.CommerceCancelSubscription(context.Background(), subscriptionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSubscriptionsAPI.CommerceCancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCancelSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `CommerceSubscriptionsAPI.CommerceCancelSubscription`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceSubscriptionsAPI.CommerceCreateSubscription(context.Background()).CommerceSubscriptionRequest(commerceSubscriptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSubscriptionsAPI.CommerceCreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `CommerceSubscriptionsAPI.CommerceCreateSubscription`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceSubscriptionsAPI.CommerceGetSubscription(context.Background(), subscriptionid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSubscriptionsAPI.CommerceGetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `CommerceSubscriptionsAPI.CommerceGetSubscription`: %v\n", resp)
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
	resp, r, err := apiClient.CommerceSubscriptionsAPI.CommerceUpdateSubscription(context.Background(), subscriptionid).CommerceSubscription(commerceSubscription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSubscriptionsAPI.CommerceUpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateSubscription`: CommerceSubscription
	fmt.Fprintf(os.Stdout, "Response from `CommerceSubscriptionsAPI.CommerceUpdateSubscription`: %v\n", resp)
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

