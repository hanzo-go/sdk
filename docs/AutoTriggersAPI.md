# \AutoTriggersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoListTriggerEvents**](AutoTriggersAPI.md#AutoListTriggerEvents) | **Get** /v1/auto/trigger-events | List trigger events for a flow
[**AutoListTriggerRuns**](AutoTriggersAPI.md#AutoListTriggerRuns) | **Get** /v1/auto/trigger-runs | List trigger run history
[**AutoTestTrigger**](AutoTriggersAPI.md#AutoTestTrigger) | **Post** /v1/auto/test-trigger | Test a trigger and get sample data



## AutoListTriggerEvents

> map[string]interface{} AutoListTriggerEvents(ctx).FlowId(flowId).Execute()

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
	resp, r, err := apiClient.AutoTriggersAPI.AutoListTriggerEvents(context.Background()).FlowId(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTriggersAPI.AutoListTriggerEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListTriggerEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTriggersAPI.AutoListTriggerEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListTriggerEventsRequest struct via the builder pattern


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


## AutoListTriggerRuns

> map[string]interface{} AutoListTriggerRuns(ctx).FlowId(flowId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTriggersAPI.AutoListTriggerRuns(context.Background()).FlowId(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTriggersAPI.AutoListTriggerRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListTriggerRuns`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTriggersAPI.AutoListTriggerRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListTriggerRunsRequest struct via the builder pattern


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


## AutoTestTrigger

> map[string]interface{} AutoTestTrigger(ctx).AutoTestTriggerRequest(autoTestTriggerRequest).Execute()

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
	resp, r, err := apiClient.AutoTriggersAPI.AutoTestTrigger(context.Background()).AutoTestTriggerRequest(autoTestTriggerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTriggersAPI.AutoTestTrigger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoTestTrigger`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTriggersAPI.AutoTestTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoTestTriggerRequest struct via the builder pattern


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

