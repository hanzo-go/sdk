# \FlowSampleDataAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowGetSampleData**](FlowSampleDataAPI.md#FlowGetSampleData) | **Get** /v1/flow/sample-data | Get sample data for a flow step
[**FlowSaveSampleData**](FlowSampleDataAPI.md#FlowSaveSampleData) | **Post** /v1/flow/sample-data | Save sample data for a flow step



## FlowGetSampleData

> map[string]interface{} FlowGetSampleData(ctx).FlowId(flowId).FlowVersionId(flowVersionId).StepName(stepName).Execute()

Get sample data for a flow step

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
	flowVersionId := "flowVersionId_example" // string | 
	stepName := "stepName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowSampleDataAPI.FlowGetSampleData(context.Background()).FlowId(flowId).FlowVersionId(flowVersionId).StepName(stepName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSampleDataAPI.FlowGetSampleData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetSampleData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSampleDataAPI.FlowGetSampleData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetSampleDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowId** | **string** |  | 
 **flowVersionId** | **string** |  | 
 **stepName** | **string** |  | 

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


## FlowSaveSampleData

> map[string]interface{} FlowSaveSampleData(ctx).FlowSaveSampleDataRequest(flowSaveSampleDataRequest).Execute()

Save sample data for a flow step

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
	flowSaveSampleDataRequest := *openapiclient.NewFlowSaveSampleDataRequest("FlowId_example", "FlowVersionId_example", "StepName_example", map[string]interface{}(123)) // FlowSaveSampleDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowSampleDataAPI.FlowSaveSampleData(context.Background()).FlowSaveSampleDataRequest(flowSaveSampleDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSampleDataAPI.FlowSaveSampleData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowSaveSampleData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSampleDataAPI.FlowSaveSampleData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowSaveSampleDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowSaveSampleDataRequest** | [**FlowSaveSampleDataRequest**](FlowSaveSampleDataRequest.md) |  | 

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

