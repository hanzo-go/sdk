# \FlowTriggersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowHandleAppEvent**](FlowTriggersAPI.md#FlowHandleAppEvent) | **Post** /v1/flow/app-events/{pieceUrl} | Handle incoming app event for trigger routing
[**FlowListTriggerEvents**](FlowTriggersAPI.md#FlowListTriggerEvents) | **Get** /v1/flow/trigger-events | List trigger events for a flow
[**FlowListTriggerRuns**](FlowTriggersAPI.md#FlowListTriggerRuns) | **Get** /v1/flow/trigger-runs | List trigger run history
[**FlowTestTrigger**](FlowTriggersAPI.md#FlowTestTrigger) | **Post** /v1/flow/test-trigger | Test a trigger and get sample data



## FlowHandleAppEvent

> map[string]interface{} FlowHandleAppEvent(ctx, pieceUrl).Body(body).Execute()

Handle incoming app event for trigger routing

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
	pieceUrl := "pieceUrl_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowTriggersAPI.FlowHandleAppEvent(context.Background(), pieceUrl).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowTriggersAPI.FlowHandleAppEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowHandleAppEvent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowTriggersAPI.FlowHandleAppEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pieceUrl** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowHandleAppEventRequest struct via the builder pattern


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


## FlowListTriggerEvents

> map[string]interface{} FlowListTriggerEvents(ctx).FlowId(flowId).Execute()

List trigger events for a flow

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
	resp, r, err := apiClient.FlowTriggersAPI.FlowListTriggerEvents(context.Background()).FlowId(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowTriggersAPI.FlowListTriggerEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListTriggerEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowTriggersAPI.FlowListTriggerEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListTriggerEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 

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


## FlowListTriggerRuns

> map[string]interface{} FlowListTriggerRuns(ctx).FlowId(flowId).Cursor(cursor).Limit(limit).Execute()

List trigger run history

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
	flowId := "flowId_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowTriggersAPI.FlowListTriggerRuns(context.Background()).FlowId(flowId).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowTriggersAPI.FlowListTriggerRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListTriggerRuns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowTriggersAPI.FlowListTriggerRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListTriggerRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

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


## FlowTestTrigger

> map[string]interface{} FlowTestTrigger(ctx).AutoTestTriggerRequest(autoTestTriggerRequest).Execute()

Test a trigger and get sample data

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
	autoTestTriggerRequest := *openapiclient.NewAutoTestTriggerRequest("FlowId_example", "FlowVersionId_example") // AutoTestTriggerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowTriggersAPI.FlowTestTrigger(context.Background()).AutoTestTriggerRequest(autoTestTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowTriggersAPI.FlowTestTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowTestTrigger`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowTriggersAPI.FlowTestTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowTestTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTestTriggerRequest** | [**AutoTestTriggerRequest**](AutoTestTriggerRequest.md) |  | 

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

