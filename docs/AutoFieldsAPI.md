# \AutoFieldsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateField**](AutoFieldsAPI.md#AutoCreateField) | **Post** /v1/auto/fields | Create a field in a table
[**AutoListFields**](AutoFieldsAPI.md#AutoListFields) | **Get** /v1/auto/fields | List fields for a table



## AutoCreateField

> map[string]interface{} AutoCreateField(ctx).AutoCreateFieldRequest(autoCreateFieldRequest).Execute()

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
	resp, r, err := apiClient.AutoFieldsAPI.AutoCreateField(context.Background()).AutoCreateFieldRequest(autoCreateFieldRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFieldsAPI.AutoCreateField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateField`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFieldsAPI.AutoCreateField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateFieldRequest struct via the builder pattern


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


## AutoListFields

> map[string]interface{} AutoListFields(ctx).TableId(tableId).Execute()

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
	resp, r, err := apiClient.AutoFieldsAPI.AutoListFields(context.Background()).TableId(tableId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoFieldsAPI.AutoListFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListFields`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoFieldsAPI.AutoListFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListFieldsRequest struct via the builder pattern


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

