# \AutoSampleDataAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoGetSampleData**](AutoSampleDataAPI.md#AutoGetSampleData) | **Get** /v1/auto/sample-data | Get sample data for a flow step



## AutoGetSampleData

> map[string]interface{} AutoGetSampleData(ctx).FlowId(flowId).FlowVersionId(flowVersionId).StepName(stepName).Execute()

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
	resp, r, err := apiClient.AutoSampleDataAPI.AutoGetSampleData(context.Background()).FlowId(flowId).FlowVersionId(flowVersionId).StepName(stepName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoSampleDataAPI.AutoGetSampleData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetSampleData`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoSampleDataAPI.AutoGetSampleData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetSampleDataRequest struct via the builder pattern


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

