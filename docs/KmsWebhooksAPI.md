# \KmsWebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateWebhook**](KmsWebhooksAPI.md#KmsCreateWebhook) | **Post** /v1/kms/webhooks | Create a webhook
[**KmsDeleteWebhook**](KmsWebhooksAPI.md#KmsDeleteWebhook) | **Delete** /v1/kms/webhooks/{webhookId} | Delete a webhook
[**KmsListWebhooks**](KmsWebhooksAPI.md#KmsListWebhooks) | **Get** /v1/kms/webhooks | List webhooks for a project
[**KmsTestWebhook**](KmsWebhooksAPI.md#KmsTestWebhook) | **Post** /v1/kms/webhooks/{webhookId}/test | Test a webhook
[**KmsUpdateWebhook**](KmsWebhooksAPI.md#KmsUpdateWebhook) | **Patch** /v1/kms/webhooks/{webhookId} | Update a webhook



## KmsCreateWebhook

> KmsCreateWebhook200Response KmsCreateWebhook(ctx).KmsCreateWebhookRequest(kmsCreateWebhookRequest).Execute()

Create a webhook

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
	kmsCreateWebhookRequest := *openapiclient.NewKmsCreateWebhookRequest("WorkspaceId_example", "Environment_example", "WebhookUrl_example") // KmsCreateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsWebhooksAPI.KmsCreateWebhook(context.Background()).KmsCreateWebhookRequest(kmsCreateWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsWebhooksAPI.KmsCreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateWebhook`: KmsCreateWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsWebhooksAPI.KmsCreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateWebhookRequest** | [**KmsCreateWebhookRequest**](KmsCreateWebhookRequest.md) |  | 

### Return type

[**KmsCreateWebhook200Response**](KmsCreateWebhook200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteWebhook

> map[string]interface{} KmsDeleteWebhook(ctx, webhookId).Execute()

Delete a webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsWebhooksAPI.KmsDeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsWebhooksAPI.KmsDeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsWebhooksAPI.KmsDeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListWebhooks

> KmsListWebhooks200Response KmsListWebhooks(ctx).WorkspaceId(workspaceId).Execute()

List webhooks for a project

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsWebhooksAPI.KmsListWebhooks(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsWebhooksAPI.KmsListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListWebhooks`: KmsListWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsWebhooksAPI.KmsListWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

[**KmsListWebhooks200Response**](KmsListWebhooks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsTestWebhook

> map[string]interface{} KmsTestWebhook(ctx, webhookId).Execute()

Test a webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsWebhooksAPI.KmsTestWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsWebhooksAPI.KmsTestWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsTestWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsWebhooksAPI.KmsTestWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsTestWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateWebhook

> KmsCreateWebhook200Response KmsUpdateWebhook(ctx, webhookId).KmsUpdateWebhookRequest(kmsUpdateWebhookRequest).Execute()

Update a webhook

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
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateWebhookRequest := *openapiclient.NewKmsUpdateWebhookRequest() // KmsUpdateWebhookRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsWebhooksAPI.KmsUpdateWebhook(context.Background(), webhookId).KmsUpdateWebhookRequest(kmsUpdateWebhookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsWebhooksAPI.KmsUpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateWebhook`: KmsCreateWebhook200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsWebhooksAPI.KmsUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateWebhookRequest** | [**KmsUpdateWebhookRequest**](KmsUpdateWebhookRequest.md) |  | 

### Return type

[**KmsCreateWebhook200Response**](KmsCreateWebhook200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

