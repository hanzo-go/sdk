# \AutoTablesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateTable**](AutoTablesAPI.md#AutoCreateTable) | **Post** /v1/auto/tables | Create a table
[**AutoListTables**](AutoTablesAPI.md#AutoListTables) | **Get** /v1/auto/tables | List tables



## AutoCreateTable

> map[string]interface{} AutoCreateTable(ctx).AutoCreateTableRequest(autoCreateTableRequest).Execute()

Create a table

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
	autoCreateTableRequest := *openapiclient.NewAutoCreateTableRequest("Name_example") // AutoCreateTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoTablesAPI.AutoCreateTable(context.Background()).AutoCreateTableRequest(autoCreateTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTablesAPI.AutoCreateTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateTable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTablesAPI.AutoCreateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateTableRequest** | [**AutoCreateTableRequest**](AutoCreateTableRequest.md) |  | 

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


## AutoListTables

> map[string]interface{} AutoListTables(ctx).Execute()

List tables

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
	resp, r, err := apiClient.AutoTablesAPI.AutoListTables(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoTablesAPI.AutoListTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListTables`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoTablesAPI.AutoListTables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListTablesRequest struct via the builder pattern


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

