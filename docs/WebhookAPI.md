# \WebhookAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhookById**](WebhookAPI.md#DeleteWebhookById) | **Delete** /v1/webhook/{id} | Removes one of the caller org&#39;s webhook endpoints and answers 204 with no body.
[**GetWebhook**](WebhookAPI.md#GetWebhook) | **Get** /v1/webhook | Returns every webhook endpoint the caller&#39;s org has registered, newest first, each with its 7-day delivery and failure counts.
[**GetWebhookById**](WebhookAPI.md#GetWebhookById) | **Get** /v1/webhook/{id} | Returns one of the caller org&#39;s webhook endpoints with its 7-day delivery and failure counts, signing secret redacted.
[**GetWebhookByIdDeliveries**](WebhookAPI.md#GetWebhookByIdDeliveries) | **Get** /v1/webhook/{id}/deliveries | Returns one endpoint&#39;s per-attempt delivery log, newest first — the record of what was sent, what the subscriber answered, and how long it took.
[**PostWebhook**](WebhookAPI.md#PostWebhook) | **Post** /v1/webhook | Registers a new webhook subscription for the caller&#39;s org and answers 201 with the endpoint INCLUDING its freshly minted signing secret.
[**PostWebhookByIdSecret**](WebhookAPI.md#PostWebhookByIdSecret) | **Post** /v1/webhook/{id}/secret | Mints a NEW HMAC signing secret for the endpoint and answers the endpoint WITH it — the only other response besides create that ever carries a secret.
[**PostWebhookByIdTest**](WebhookAPI.md#PostWebhookByIdTest) | **Post** /v1/webhook/{id}/test | Sends ONE signed test event to the endpoint right now and answers the outcome inline, so the console can show whether the subscriber is reachable without waiting for real traffic.
[**PutWebhookById**](WebhookAPI.md#PutWebhookById) | **Put** /v1/webhook/{id} | Replaces the editable fields of one of the caller org&#39;s endpoints — url, events, status and description — and answers the stored row with its secret redacted.



## DeleteWebhookById

> DeleteWebhookById(ctx, id).Execute()

Removes one of the caller org's webhook endpoints and answers 204 with no body.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhookAPI.DeleteWebhookById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeleteWebhookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> EndpointList GetWebhook(ctx).Execute()

Returns every webhook endpoint the caller's org has registered, newest first, each with its 7-day delivery and failure counts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.GetWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: EndpointList
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


### Return type

[**EndpointList**](EndpointList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookById

> Endpoint GetWebhookById(ctx, id).Execute()

Returns one of the caller org's webhook endpoints with its 7-day delivery and failure counts, signing secret redacted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.GetWebhookById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetWebhookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetWebhookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhookByIdDeliveries

> DeliveryList GetWebhookByIdDeliveries(ctx, id).Limit(limit).Status(status).Execute()

Returns one endpoint's per-attempt delivery log, newest first — the record of what was sent, what the subscriber answered, and how long it took.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 
	limit := int64(100) // int64 | Limit caps how many attempts come back: default 50, maximum 200. A value that is not a positive integer reads as the default. (optional)
	status := "failed" // string | Status narrows the log to one outcome: \"ok\", \"retrying\" or \"failed\". Empty returns every attempt. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.GetWebhookByIdDeliveries(context.Background(), id).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.GetWebhookByIdDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhookByIdDeliveries`: DeliveryList
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.GetWebhookByIdDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookByIdDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit caps how many attempts come back: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **status** | **string** | Status narrows the log to one outcome: \&quot;ok\&quot;, \&quot;retrying\&quot; or \&quot;failed\&quot;. Empty returns every attempt. | 

### Return type

[**DeliveryList**](DeliveryList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhook

> Endpoint PostWebhook(ctx).CreateEndpointIn(createEndpointIn).Execute()

Registers a new webhook subscription for the caller's org and answers 201 with the endpoint INCLUDING its freshly minted signing secret.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	createEndpointIn := *openapiclient.NewCreateEndpointIn() // CreateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PostWebhook(context.Background()).CreateEndpointIn(createEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PostWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhook`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PostWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointIn** | [**CreateEndpointIn**](CreateEndpointIn.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhookByIdSecret

> Endpoint PostWebhookByIdSecret(ctx, id).Execute()

Mints a NEW HMAC signing secret for the endpoint and answers the endpoint WITH it — the only other response besides create that ever carries a secret.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PostWebhookByIdSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PostWebhookByIdSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhookByIdSecret`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PostWebhookByIdSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookByIdSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhookByIdTest

> TestResult PostWebhookByIdTest(ctx, id).Execute()

Sends ONE signed test event to the endpoint right now and answers the outcome inline, so the console can show whether the subscriber is reachable without waiting for real traffic.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PostWebhookByIdTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PostWebhookByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhookByIdTest`: TestResult
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PostWebhookByIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhookByIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResult**](TestResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhookById

> Endpoint PutWebhookById(ctx, id).UpdateEndpointIn(updateEndpointIn).Execute()

Replaces the editable fields of one of the caller org's endpoints — url, events, status and description — and answers the stored row with its secret redacted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "wh_9f8c1d2e" // string | 
	updateEndpointIn := *openapiclient.NewUpdateEndpointIn() // UpdateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PutWebhookById(context.Background(), id).UpdateEndpointIn(updateEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PutWebhookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWebhookById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PutWebhookById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEndpointIn** | [**UpdateEndpointIn**](UpdateEndpointIn.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

