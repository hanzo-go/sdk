# \FlowFieldsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateField**](FlowFieldsAPI.md#FlowCreateField) | **Post** /v1/flow/fields | Create a field in a table
[**FlowListFields**](FlowFieldsAPI.md#FlowListFields) | **Get** /v1/flow/fields | List fields for a table



## FlowCreateField

> map[string]interface{} FlowCreateField(ctx).AutoCreateFieldRequest(autoCreateFieldRequest).Execute()

Create a field in a table

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
	autoCreateFieldRequest := *openapiclient.NewAutoCreateFieldRequest("Name_example", "Type_example", "TableId_example") // AutoCreateFieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFieldsAPI.FlowCreateField(context.Background()).AutoCreateFieldRequest(autoCreateFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFieldsAPI.FlowCreateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateField`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFieldsAPI.FlowCreateField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateFieldRequest** | [**AutoCreateFieldRequest**](AutoCreateFieldRequest.md) |  | 

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


## FlowListFields

> map[string]interface{} FlowListFields(ctx).TableId(tableId).Execute()

List fields for a table

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
	tableId := "tableId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowFieldsAPI.FlowListFields(context.Background()).TableId(tableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowFieldsAPI.FlowListFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListFields`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowFieldsAPI.FlowListFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tableId** | **string** |  | 

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

