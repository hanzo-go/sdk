# \WebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhooksById**](WebhooksAPI.md#DeleteWebhooksById) | **Delete** /v1/webhooks/{id} | Removes one of the caller org&#39;s webhook endpoints and answers 204 with no body.
[**GetWebhooks**](WebhooksAPI.md#GetWebhooks) | **Get** /v1/webhooks | Returns every webhook endpoint the caller&#39;s org has registered, newest first, each with its 7-day delivery and failure counts.
[**GetWebhooksById**](WebhooksAPI.md#GetWebhooksById) | **Get** /v1/webhooks/{id} | Returns one of the caller org&#39;s webhook endpoints with its 7-day delivery and failure counts, signing secret redacted.
[**GetWebhooksByIdDeliveries**](WebhooksAPI.md#GetWebhooksByIdDeliveries) | **Get** /v1/webhooks/{id}/deliveries | Returns one endpoint&#39;s per-attempt delivery log, newest first — the record of what was sent, what the subscriber answered, and how long it took.
[**PostWebhooks**](WebhooksAPI.md#PostWebhooks) | **Post** /v1/webhooks | Registers a new webhook subscription for the caller&#39;s org and answers 201 with the endpoint INCLUDING its freshly minted signing secret.
[**PostWebhooksByIdSecret**](WebhooksAPI.md#PostWebhooksByIdSecret) | **Post** /v1/webhooks/{id}/secret | Mints a NEW HMAC signing secret for the endpoint and answers the endpoint WITH it — the only other response besides create that ever carries a secret.
[**PostWebhooksByIdTest**](WebhooksAPI.md#PostWebhooksByIdTest) | **Post** /v1/webhooks/{id}/test | Sends ONE signed test event to the endpoint right now and answers the outcome inline, so the console can show whether the subscriber is reachable without waiting for real traffic.
[**PutWebhooksById**](WebhooksAPI.md#PutWebhooksById) | **Put** /v1/webhooks/{id} | Replaces the editable fields of one of the caller org&#39;s endpoints — url, events, status and description — and answers the stored row with its secret redacted.



## DeleteWebhooksById

> DeleteWebhooksById(ctx, id).Execute()

Removes one of the caller org's webhook endpoints and answers 204 with no body.



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
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.DeleteWebhooksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhooksById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteWebhooksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> EndpointList GetWebhooks(ctx).Execute()

Returns every webhook endpoint the caller's org has registered, newest first, each with its 7-day delivery and failure counts.



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
	resp, r, err := apiClient.WebhooksAPI.GetWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooks`: EndpointList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


### Return type

[**EndpointList**](EndpointList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksById

> Endpoint GetWebhooksById(ctx, id).Execute()

Returns one of the caller org's webhook endpoints with its 7-day delivery and failure counts, signing secret redacted.



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
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhooksById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooksByIdDeliveries

> DeliveryList GetWebhooksByIdDeliveries(ctx, id).Limit(limit).Status(status).Execute()

Returns one endpoint's per-attempt delivery log, newest first — the record of what was sent, what the subscriber answered, and how long it took.



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
	id := "wh_9f8c1d2e" // string | 
	limit := int32(100) // int32 | Limit caps how many attempts come back: default 50, maximum 200. A value that is not a positive integer reads as the default. (optional)
	status := "failed" // string | Status narrows the log to one outcome: \"ok\", \"retrying\" or \"failed\". Empty returns every attempt. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetWebhooksByIdDeliveries(context.Background(), id).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetWebhooksByIdDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooksByIdDeliveries`: DeliveryList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetWebhooksByIdDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksByIdDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps how many attempts come back: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **status** | **string** | Status narrows the log to one outcome: \&quot;ok\&quot;, \&quot;retrying\&quot; or \&quot;failed\&quot;. Empty returns every attempt. | 

### Return type

[**DeliveryList**](DeliveryList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooks

> Endpoint PostWebhooks(ctx).CreateEndpointIn(createEndpointIn).Execute()

Registers a new webhook subscription for the caller's org and answers 201 with the endpoint INCLUDING its freshly minted signing secret.



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
	createEndpointIn := *openapiclient.NewCreateEndpointIn() // CreateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.PostWebhooks(context.Background()).CreateEndpointIn(createEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhooks`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PostWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEndpointIn** | [**CreateEndpointIn**](CreateEndpointIn.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksByIdSecret

> Endpoint PostWebhooksByIdSecret(ctx, id).Execute()

Mints a NEW HMAC signing secret for the endpoint and answers the endpoint WITH it — the only other response besides create that ever carries a secret.



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
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.PostWebhooksByIdSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooksByIdSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhooksByIdSecret`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PostWebhooksByIdSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksByIdSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWebhooksByIdTest

> TestResult PostWebhooksByIdTest(ctx, id).Execute()

Sends ONE signed test event to the endpoint right now and answers the outcome inline, so the console can show whether the subscriber is reachable without waiting for real traffic.



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
	id := "wh_9f8c1d2e" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.PostWebhooksByIdTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PostWebhooksByIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWebhooksByIdTest`: TestResult
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PostWebhooksByIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWebhooksByIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TestResult**](TestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutWebhooksById

> Endpoint PutWebhooksById(ctx, id).UpdateEndpointIn(updateEndpointIn).Execute()

Replaces the editable fields of one of the caller org's endpoints — url, events, status and description — and answers the stored row with its secret redacted.



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
	id := "wh_9f8c1d2e" // string | 
	updateEndpointIn := *openapiclient.NewUpdateEndpointIn() // UpdateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.PutWebhooksById(context.Background(), id).UpdateEndpointIn(updateEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.PutWebhooksById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutWebhooksById`: Endpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.PutWebhooksById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutWebhooksByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEndpointIn** | [**UpdateEndpointIn**](UpdateEndpointIn.md) |  | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

