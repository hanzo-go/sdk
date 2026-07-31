# \WebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1WebhooksId**](WebhooksAPI.md#CloudDeleteV1WebhooksId) | **Delete** /v1/webhooks/{id} | Removes one of the caller org&#39;s webhook endpoints and answers 204 with no body.
[**CloudGetV1Webhooks**](WebhooksAPI.md#CloudGetV1Webhooks) | **Get** /v1/webhooks | Returns every webhook endpoint the caller&#39;s org has registered, newest first, each with its 7-day delivery and failure counts.
[**CloudGetV1WebhooksId**](WebhooksAPI.md#CloudGetV1WebhooksId) | **Get** /v1/webhooks/{id} | Returns one of the caller org&#39;s webhook endpoints with its 7-day delivery and failure counts, signing secret redacted.
[**CloudGetV1WebhooksIdDeliveries**](WebhooksAPI.md#CloudGetV1WebhooksIdDeliveries) | **Get** /v1/webhooks/{id}/deliveries | Returns one endpoint&#39;s per-attempt delivery log, newest first — the record of what was sent, what the subscriber answered, and how long it took.
[**CloudPostV1Webhooks**](WebhooksAPI.md#CloudPostV1Webhooks) | **Post** /v1/webhooks | Registers a new webhook subscription for the caller&#39;s org and answers 201 with the endpoint INCLUDING its freshly minted signing secret.
[**CloudPostV1WebhooksIdRotateSecret**](WebhooksAPI.md#CloudPostV1WebhooksIdRotateSecret) | **Post** /v1/webhooks/{id}/rotate-secret | Mints a NEW HMAC signing secret for the endpoint and answers the endpoint WITH it — the only other response besides create that ever carries a secret.
[**CloudPostV1WebhooksIdTest**](WebhooksAPI.md#CloudPostV1WebhooksIdTest) | **Post** /v1/webhooks/{id}/test | Sends ONE signed test event to the endpoint right now and answers the outcome inline, so the console can show whether the subscriber is reachable without waiting for real traffic.
[**CloudPutV1WebhooksId**](WebhooksAPI.md#CloudPutV1WebhooksId) | **Put** /v1/webhooks/{id} | Replaces the editable fields of one of the caller org&#39;s endpoints — url, events, status and description — and answers the stored row with its secret redacted.
[**IamApiControllerAddWebhook**](WebhooksAPI.md#IamApiControllerAddWebhook) | **Post** /v1/iam/webhooks | Api Controller Add Webhook
[**IamApiControllerDeleteWebhook**](WebhooksAPI.md#IamApiControllerDeleteWebhook) | **Delete** /v1/iam/webhooks/{id} | Api Controller Delete Webhook
[**IamApiControllerGetWebhook**](WebhooksAPI.md#IamApiControllerGetWebhook) | **Get** /v1/iam/webhooks/{id} | Api Controller Get Webhook
[**IamApiControllerGetWebhooks**](WebhooksAPI.md#IamApiControllerGetWebhooks) | **Get** /v1/iam/webhooks | Api Controller Get Webhooks
[**IamApiControllerUpdateWebhook**](WebhooksAPI.md#IamApiControllerUpdateWebhook) | **Put** /v1/iam/webhooks/{id} | Api Controller Update Webhook
[**SearchDeleteWebhooks**](WebhooksAPI.md#SearchDeleteWebhooks) | **Delete** /v1/search/webhooks | Delete all webhooks
[**SearchGetWebhooks**](WebhooksAPI.md#SearchGetWebhooks) | **Get** /v1/search/webhooks | Get webhook configuration
[**SearchUpdateWebhooks**](WebhooksAPI.md#SearchUpdateWebhooks) | **Patch** /v1/search/webhooks | Update webhook configuration



## CloudDeleteV1WebhooksId

> CloudDeleteV1WebhooksId(ctx, id).Execute()

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
	r, err := apiClient.WebhooksAPI.CloudDeleteV1WebhooksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudDeleteV1WebhooksId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1WebhooksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Webhooks

> CloudEndpointList CloudGetV1Webhooks(ctx).Execute()

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
	resp, r, err := apiClient.WebhooksAPI.CloudGetV1Webhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudGetV1Webhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Webhooks`: CloudEndpointList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudGetV1Webhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WebhooksRequest struct via the builder pattern


### Return type

[**CloudEndpointList**](CloudEndpointList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WebhooksId

> CloudEndpoint CloudGetV1WebhooksId(ctx, id).Execute()

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
	resp, r, err := apiClient.WebhooksAPI.CloudGetV1WebhooksId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudGetV1WebhooksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WebhooksId`: CloudEndpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudGetV1WebhooksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WebhooksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudEndpoint**](CloudEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WebhooksIdDeliveries

> CloudDeliveryList CloudGetV1WebhooksIdDeliveries(ctx, id).Limit(limit).Status(status).Execute()

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
	resp, r, err := apiClient.WebhooksAPI.CloudGetV1WebhooksIdDeliveries(context.Background(), id).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudGetV1WebhooksIdDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WebhooksIdDeliveries`: CloudDeliveryList
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudGetV1WebhooksIdDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WebhooksIdDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps how many attempts come back: default 50, maximum 200. A value that is not a positive integer reads as the default. | 
 **status** | **string** | Status narrows the log to one outcome: \&quot;ok\&quot;, \&quot;retrying\&quot; or \&quot;failed\&quot;. Empty returns every attempt. | 

### Return type

[**CloudDeliveryList**](CloudDeliveryList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Webhooks

> CloudEndpoint CloudPostV1Webhooks(ctx).CloudCreateEndpointIn(cloudCreateEndpointIn).Execute()

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
	cloudCreateEndpointIn := *openapiclient.NewCloudCreateEndpointIn() // CloudCreateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CloudPostV1Webhooks(context.Background()).CloudCreateEndpointIn(cloudCreateEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudPostV1Webhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Webhooks`: CloudEndpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudPostV1Webhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateEndpointIn** | [**CloudCreateEndpointIn**](CloudCreateEndpointIn.md) |  | 

### Return type

[**CloudEndpoint**](CloudEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WebhooksIdRotateSecret

> CloudEndpoint CloudPostV1WebhooksIdRotateSecret(ctx, id).Execute()

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
	resp, r, err := apiClient.WebhooksAPI.CloudPostV1WebhooksIdRotateSecret(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudPostV1WebhooksIdRotateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WebhooksIdRotateSecret`: CloudEndpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudPostV1WebhooksIdRotateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WebhooksIdRotateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudEndpoint**](CloudEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WebhooksIdTest

> CloudTestResult CloudPostV1WebhooksIdTest(ctx, id).Execute()

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
	resp, r, err := apiClient.WebhooksAPI.CloudPostV1WebhooksIdTest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudPostV1WebhooksIdTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WebhooksIdTest`: CloudTestResult
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudPostV1WebhooksIdTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WebhooksIdTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTestResult**](CloudTestResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1WebhooksId

> CloudEndpoint CloudPutV1WebhooksId(ctx, id).CloudUpdateEndpointIn(cloudUpdateEndpointIn).Execute()

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
	cloudUpdateEndpointIn := *openapiclient.NewCloudUpdateEndpointIn() // CloudUpdateEndpointIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.CloudPutV1WebhooksId(context.Background(), id).CloudUpdateEndpointIn(cloudUpdateEndpointIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.CloudPutV1WebhooksId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1WebhooksId`: CloudEndpoint
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.CloudPutV1WebhooksId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1WebhooksIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudUpdateEndpointIn** | [**CloudUpdateEndpointIn**](CloudUpdateEndpointIn.md) |  | 

### Return type

[**CloudEndpoint**](CloudEndpoint.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddWebhook

> IamControllersResponse IamApiControllerAddWebhook(ctx).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Add Webhook



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
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerAddWebhook(context.Background()).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerAddWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerAddWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteWebhook

> IamControllersResponse IamApiControllerDeleteWebhook(ctx, id).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Delete Webhook



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerDeleteWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerDeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWebhook

> IamObjectWebhook IamApiControllerGetWebhook(ctx, id).Execute()

Api Controller Get Webhook



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
	id := "id_example" // string | The id ( owner/name ) of the webhook (default to "built-in/admin")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerGetWebhook(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerGetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhook`: IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerGetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the webhook | [default to &quot;built-in/admin&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectWebhook**](IamObjectWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetWebhooks

> []IamObjectWebhook IamApiControllerGetWebhooks(ctx).Owner(owner).Execute()

Api Controller Get Webhooks



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
	owner := "owner_example" // string | The owner of webhooks (default to "built-in/admin")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerGetWebhooks(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetWebhooks`: []IamObjectWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerGetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of webhooks | [default to &quot;built-in/admin&quot;]

### Return type

[**[]IamObjectWebhook**](IamObjectWebhook.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateWebhook

> IamControllersResponse IamApiControllerUpdateWebhook(ctx, id).IamObjectWebhook(iamObjectWebhook).Execute()

Api Controller Update Webhook



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
	id := "id_example" // string | The id ( owner/name ) of the webhook (default to "built-in/admin")
	iamObjectWebhook := *openapiclient.NewIamObjectWebhook() // IamObjectWebhook | The details of the webhook

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.IamApiControllerUpdateWebhook(context.Background(), id).IamObjectWebhook(iamObjectWebhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.IamApiControllerUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateWebhook`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.IamApiControllerUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the webhook | [default to &quot;built-in/admin&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectWebhook** | [**IamObjectWebhook**](IamObjectWebhook.md) | The details of the webhook | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteWebhooks

> map[string]SearchWebhookResultsValue SearchDeleteWebhooks(ctx).Execute()

Delete all webhooks

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
	resp, r, err := apiClient.WebhooksAPI.SearchDeleteWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchDeleteWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDeleteWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchDeleteWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteWebhooksRequest struct via the builder pattern


### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetWebhooks

> map[string]SearchWebhookResultsValue SearchGetWebhooks(ctx).Execute()

Get webhook configuration

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
	resp, r, err := apiClient.WebhooksAPI.SearchGetWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchGetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchGetWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetWebhooksRequest struct via the builder pattern


### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateWebhooks

> map[string]SearchWebhookResultsValue SearchUpdateWebhooks(ctx).RequestBody(requestBody).Execute()

Update webhook configuration

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
	requestBody := map[string]SearchWebhookSettingsValue{"key": *openapiclient.NewSearchWebhookSettingsValue()} // map[string]SearchWebhookSettingsValue | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.SearchUpdateWebhooks(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.SearchUpdateWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateWebhooks`: map[string]SearchWebhookResultsValue
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.SearchUpdateWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]SearchWebhookSettingsValue**](SearchWebhookSettingsValue.md) |  | 

### Return type

[**map[string]SearchWebhookResultsValue**](SearchWebhookResultsValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

