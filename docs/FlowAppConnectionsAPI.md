# \FlowAppConnectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowDeleteAppConnection**](FlowAppConnectionsAPI.md#FlowDeleteAppConnection) | **Delete** /v1/flow/app-connections/{id} | Delete an app connection
[**FlowListAppConnectionOwners**](FlowAppConnectionsAPI.md#FlowListAppConnectionOwners) | **Get** /v1/flow/app-connections/owners | List app connection owners
[**FlowListAppConnections**](FlowAppConnectionsAPI.md#FlowListAppConnections) | **Get** /v1/flow/app-connections | List app connections
[**FlowReplaceAppConnections**](FlowAppConnectionsAPI.md#FlowReplaceAppConnections) | **Post** /v1/flow/app-connections/replace | Replace one connection with another across all flows
[**FlowUpdateAppConnection**](FlowAppConnectionsAPI.md#FlowUpdateAppConnection) | **Post** /v1/flow/app-connections/{id} | Update an app connection
[**FlowUpsertAppConnection**](FlowAppConnectionsAPI.md#FlowUpsertAppConnection) | **Post** /v1/flow/app-connections | Upsert an app connection



## FlowDeleteAppConnection

> FlowDeleteAppConnection(ctx, id).Execute()

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
	r, err := apiClient.FlowAppConnectionsAPI.FlowDeleteAppConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowDeleteAppConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowDeleteAppConnectionRequest struct via the builder pattern


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


## FlowListAppConnectionOwners

> FlowSeekPage FlowListAppConnectionOwners(ctx).Execute()

List app connection owners

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
	resp, r, err := apiClient.FlowAppConnectionsAPI.FlowListAppConnectionOwners(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowListAppConnectionOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListAppConnectionOwners`: FlowSeekPage
	fmt.Fprintf(os.Stdout, "Response from `FlowAppConnectionsAPI.FlowListAppConnectionOwners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListAppConnectionOwnersRequest struct via the builder pattern


### Return type

[**FlowSeekPage**](FlowSeekPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowListAppConnections

> FlowListAppConnections200Response FlowListAppConnections(ctx).PieceName(pieceName).DisplayName(displayName).Status(status).Cursor(cursor).Limit(limit).Execute()

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
	resp, r, err := apiClient.FlowAppConnectionsAPI.FlowListAppConnections(context.Background()).PieceName(pieceName).DisplayName(displayName).Status(status).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowListAppConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListAppConnections`: FlowListAppConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `FlowAppConnectionsAPI.FlowListAppConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListAppConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pieceName** | **string** |  | 
 **displayName** | **string** |  | 
 **status** | **[]string** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | [default to 10]

### Return type

[**FlowListAppConnections200Response**](FlowListAppConnections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowReplaceAppConnections

> map[string]interface{} FlowReplaceAppConnections(ctx).FlowReplaceAppConnectionsRequest(flowReplaceAppConnectionsRequest).Execute()

Replace one connection with another across all flows

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
	flowReplaceAppConnectionsRequest := *openapiclient.NewFlowReplaceAppConnectionsRequest("SourceAppConnectionId_example", "TargetAppConnectionId_example") // FlowReplaceAppConnectionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAppConnectionsAPI.FlowReplaceAppConnections(context.Background()).FlowReplaceAppConnectionsRequest(flowReplaceAppConnectionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowReplaceAppConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowReplaceAppConnections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowAppConnectionsAPI.FlowReplaceAppConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowReplaceAppConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowReplaceAppConnectionsRequest** | [**FlowReplaceAppConnectionsRequest**](FlowReplaceAppConnectionsRequest.md) |  | 

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


## FlowUpdateAppConnection

> FlowAppConnection FlowUpdateAppConnection(ctx, id).FlowUpdateAppConnectionRequest(flowUpdateAppConnectionRequest).Execute()

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
	flowUpdateAppConnectionRequest := *openapiclient.NewFlowUpdateAppConnectionRequest() // FlowUpdateAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAppConnectionsAPI.FlowUpdateAppConnection(context.Background(), id).FlowUpdateAppConnectionRequest(flowUpdateAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowUpdateAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdateAppConnection`: FlowAppConnection
	fmt.Fprintf(os.Stdout, "Response from `FlowAppConnectionsAPI.FlowUpdateAppConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdateAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flowUpdateAppConnectionRequest** | [**FlowUpdateAppConnectionRequest**](FlowUpdateAppConnectionRequest.md) |  | 

### Return type

[**FlowAppConnection**](FlowAppConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowUpsertAppConnection

> FlowAppConnection FlowUpsertAppConnection(ctx).FlowUpsertAppConnectionRequest(flowUpsertAppConnectionRequest).Execute()

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
	flowUpsertAppConnectionRequest := *openapiclient.NewFlowUpsertAppConnectionRequest("DisplayName_example", "PieceName_example", "Type_example", map[string]interface{}(123)) // FlowUpsertAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowAppConnectionsAPI.FlowUpsertAppConnection(context.Background()).FlowUpsertAppConnectionRequest(flowUpsertAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowAppConnectionsAPI.FlowUpsertAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpsertAppConnection`: FlowAppConnection
	fmt.Fprintf(os.Stdout, "Response from `FlowAppConnectionsAPI.FlowUpsertAppConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpsertAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowUpsertAppConnectionRequest** | [**FlowUpsertAppConnectionRequest**](FlowUpsertAppConnectionRequest.md) |  | 

### Return type

[**FlowAppConnection**](FlowAppConnection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

