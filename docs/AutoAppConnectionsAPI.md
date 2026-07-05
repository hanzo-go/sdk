# \AutoAppConnectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoDeleteAppConnection**](AutoAppConnectionsAPI.md#AutoDeleteAppConnection) | **Delete** /v1/auto/app-connections/{id} | Delete an app connection
[**AutoListAppConnections**](AutoAppConnectionsAPI.md#AutoListAppConnections) | **Get** /v1/auto/app-connections | List app connections
[**AutoUpdateAppConnection**](AutoAppConnectionsAPI.md#AutoUpdateAppConnection) | **Post** /v1/auto/app-connections/{id} | Update an app connection
[**AutoUpsertAppConnection**](AutoAppConnectionsAPI.md#AutoUpsertAppConnection) | **Post** /v1/auto/app-connections | Upsert an app connection



## AutoDeleteAppConnection

> AutoDeleteAppConnection(ctx, id).Execute()

Delete an app connection

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoAppConnectionsAPI.AutoDeleteAppConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAppConnectionsAPI.AutoDeleteAppConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutoDeleteAppConnectionRequest struct via the builder pattern


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


## AutoListAppConnections

> AutoListAppConnections200Response AutoListAppConnections(ctx).PieceName(pieceName).DisplayName(displayName).Status(status).Cursor(cursor).Limit(limit).Execute()

List app connections

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
	pieceName := "pieceName_example" // string |  (optional)
	displayName := "displayName_example" // string |  (optional)
	status := []string{"Status_example"} // []string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAppConnectionsAPI.AutoListAppConnections(context.Background()).PieceName(pieceName).DisplayName(displayName).Status(status).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAppConnectionsAPI.AutoListAppConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListAppConnections`: AutoListAppConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `AutoAppConnectionsAPI.AutoListAppConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoListAppConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pieceName** | **string** |  | 
 **displayName** | **string** |  | 
 **status** | **[]string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]

### Return type

[**AutoListAppConnections200Response**](AutoListAppConnections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoUpdateAppConnection

> map[string]interface{} AutoUpdateAppConnection(ctx, id).AutoUpdateAppConnectionRequest(autoUpdateAppConnectionRequest).Execute()

Update an app connection

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
	id := "id_example" // string | 
	autoUpdateAppConnectionRequest := *openapiclient.NewAutoUpdateAppConnectionRequest() // AutoUpdateAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAppConnectionsAPI.AutoUpdateAppConnection(context.Background(), id).AutoUpdateAppConnectionRequest(autoUpdateAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAppConnectionsAPI.AutoUpdateAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpdateAppConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoAppConnectionsAPI.AutoUpdateAppConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpdateAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoUpdateAppConnectionRequest** | [**AutoUpdateAppConnectionRequest**](AutoUpdateAppConnectionRequest.md) |  | 

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


## AutoUpsertAppConnection

> AutoAppConnection AutoUpsertAppConnection(ctx).AutoUpsertAppConnectionRequest(autoUpsertAppConnectionRequest).Execute()

Upsert an app connection

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
	autoUpsertAppConnectionRequest := *openapiclient.NewAutoUpsertAppConnectionRequest("DisplayName_example", "PieceName_example", "Type_example", map[string]interface{}(123)) // AutoUpsertAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoAppConnectionsAPI.AutoUpsertAppConnection(context.Background()).AutoUpsertAppConnectionRequest(autoUpsertAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoAppConnectionsAPI.AutoUpsertAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpsertAppConnection`: AutoAppConnection
	fmt.Fprintf(os.Stdout, "Response from `AutoAppConnectionsAPI.AutoUpsertAppConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpsertAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoUpsertAppConnectionRequest** | [**AutoUpsertAppConnectionRequest**](AutoUpsertAppConnectionRequest.md) |  | 

### Return type

[**AutoAppConnection**](AutoAppConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

