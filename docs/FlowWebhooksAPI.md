# \FlowWebhooksAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowHandleWebhook**](FlowWebhooksAPI.md#FlowHandleWebhook) | **Post** /v1/flow/webhooks/{flowId} | Receive incoming webhook for a flow trigger
[**FlowHandleWebhookSync**](FlowWebhooksAPI.md#FlowHandleWebhookSync) | **Post** /v1/flow/webhooks/{flowId}/sync | Receive webhook and wait for flow run result
[**FlowSimulateWebhook**](FlowWebhooksAPI.md#FlowSimulateWebhook) | **Get** /v1/flow/webhooks/{flowId}/simulate | Simulate a webhook to capture sample data



## FlowHandleWebhook

> map[string]interface{} FlowHandleWebhook(ctx, flowId).Body(body).Execute()

Receive incoming webhook for a flow trigger

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
	flowId := "flowId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowWebhooksAPI.FlowHandleWebhook(context.Background(), flowId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowWebhooksAPI.FlowHandleWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowHandleWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowWebhooksAPI.FlowHandleWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowHandleWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowHandleWebhookSync

> map[string]interface{} FlowHandleWebhookSync(ctx, flowId).Body(body).Execute()

Receive webhook and wait for flow run result

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
	flowId := "flowId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowWebhooksAPI.FlowHandleWebhookSync(context.Background(), flowId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowWebhooksAPI.FlowHandleWebhookSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowHandleWebhookSync`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowWebhooksAPI.FlowHandleWebhookSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowHandleWebhookSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowSimulateWebhook

> map[string]interface{} FlowSimulateWebhook(ctx, flowId).Execute()

Simulate a webhook to capture sample data

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
	flowId := "flowId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowWebhooksAPI.FlowSimulateWebhook(context.Background(), flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowWebhooksAPI.FlowSimulateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowSimulateWebhook`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowWebhooksAPI.FlowSimulateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flowId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowSimulateWebhookRequest struct via the builder pattern


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

