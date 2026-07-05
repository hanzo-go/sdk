# \FlowGlobalConnectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateGlobalConnection**](FlowGlobalConnectionsAPI.md#FlowCreateGlobalConnection) | **Post** /v1/flow/global-connections | Create a global connection (EE)
[**FlowDeleteGlobalConnection**](FlowGlobalConnectionsAPI.md#FlowDeleteGlobalConnection) | **Delete** /v1/flow/global-connections/{id} | Delete a global connection (EE)
[**FlowListGlobalConnections**](FlowGlobalConnectionsAPI.md#FlowListGlobalConnections) | **Get** /v1/flow/global-connections | List platform-wide connections (EE)
[**FlowUpdateGlobalConnection**](FlowGlobalConnectionsAPI.md#FlowUpdateGlobalConnection) | **Post** /v1/flow/global-connections/{id} | Update a global connection (EE)



## FlowCreateGlobalConnection

> map[string]interface{} FlowCreateGlobalConnection(ctx).Body(body).Execute()

Create a global connection (EE)

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowGlobalConnectionsAPI.FlowCreateGlobalConnection(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGlobalConnectionsAPI.FlowCreateGlobalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateGlobalConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGlobalConnectionsAPI.FlowCreateGlobalConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateGlobalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## FlowDeleteGlobalConnection

> FlowDeleteGlobalConnection(ctx, id).Execute()

Delete a global connection (EE)

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
	r, err := apiClient.FlowGlobalConnectionsAPI.FlowDeleteGlobalConnection(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGlobalConnectionsAPI.FlowDeleteGlobalConnection``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowDeleteGlobalConnectionRequest struct via the builder pattern


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


## FlowListGlobalConnections

> map[string]interface{} FlowListGlobalConnections(ctx).Execute()

List platform-wide connections (EE)

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
	resp, r, err := apiClient.FlowGlobalConnectionsAPI.FlowListGlobalConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGlobalConnectionsAPI.FlowListGlobalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListGlobalConnections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGlobalConnectionsAPI.FlowListGlobalConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListGlobalConnectionsRequest struct via the builder pattern


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


## FlowUpdateGlobalConnection

> map[string]interface{} FlowUpdateGlobalConnection(ctx, id).Body(body).Execute()

Update a global connection (EE)

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowGlobalConnectionsAPI.FlowUpdateGlobalConnection(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowGlobalConnectionsAPI.FlowUpdateGlobalConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpdateGlobalConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowGlobalConnectionsAPI.FlowUpdateGlobalConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpdateGlobalConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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

