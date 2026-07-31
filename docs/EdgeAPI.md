# \EdgeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1EdgeNodes**](EdgeAPI.md#CloudGetV1EdgeNodes) | **Get** /v1/edge/nodes | Returns the Zero Trust edge-routers the caller&#39;s org owns.
[**GatewayInvokeEdgeFunction**](EdgeAPI.md#GatewayInvokeEdgeFunction) | **Post** /v1/gateway/edge/{slug} | Invoke edge function via gateway
[**GatewayInvokeEdgeFunctionGet**](EdgeAPI.md#GatewayInvokeEdgeFunctionGet) | **Get** /v1/gateway/edge/{slug} | Invoke edge function (GET)



## CloudGetV1EdgeNodes

> CloudEdgeNodeList CloudGetV1EdgeNodes(ctx).Execute()

Returns the Zero Trust edge-routers the caller's org owns.



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
	resp, r, err := apiClient.EdgeAPI.CloudGetV1EdgeNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.CloudGetV1EdgeNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1EdgeNodes`: CloudEdgeNodeList
	fmt.Fprintf(os.Stdout, "Response from `EdgeAPI.CloudGetV1EdgeNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EdgeNodesRequest struct via the builder pattern


### Return type

[**CloudEdgeNodeList**](CloudEdgeNodeList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayInvokeEdgeFunction

> map[string]interface{} GatewayInvokeEdgeFunction(ctx, slug).Body(body).Execute()

Invoke edge function via gateway

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
	slug := "slug_example" // string | Function slug
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeAPI.GatewayInvokeEdgeFunction(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.GatewayInvokeEdgeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayInvokeEdgeFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeAPI.GatewayInvokeEdgeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Function slug | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayInvokeEdgeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream
- **Accept**: application/json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayInvokeEdgeFunctionGet

> map[string]interface{} GatewayInvokeEdgeFunctionGet(ctx, slug).Execute()

Invoke edge function (GET)

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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeAPI.GatewayInvokeEdgeFunctionGet(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.GatewayInvokeEdgeFunctionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayInvokeEdgeFunctionGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeAPI.GatewayInvokeEdgeFunctionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayInvokeEdgeFunctionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

