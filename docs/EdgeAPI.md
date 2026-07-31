# \EdgeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayInvokeEdgeFunction**](EdgeAPI.md#GatewayInvokeEdgeFunction) | **Post** /v1/gateway/edge/{slug} | Invoke edge function via gateway
[**GatewayInvokeEdgeFunctionGet**](EdgeAPI.md#GatewayInvokeEdgeFunctionGet) | **Get** /v1/gateway/edge/{slug} | Invoke edge function (GET)
[**ZtListEdgeNodes**](EdgeAPI.md#ZtListEdgeNodes) | **Get** /v1/edge/nodes | List the org&#39;s ZT edge-routers



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


## ZtListEdgeNodes

> ZtListEdgeNodes200Response ZtListEdgeNodes(ctx).Execute()

List the org's ZT edge-routers

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
	resp, r, err := apiClient.EdgeAPI.ZtListEdgeNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeAPI.ZtListEdgeNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZtListEdgeNodes`: ZtListEdgeNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `EdgeAPI.ZtListEdgeNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZtListEdgeNodesRequest struct via the builder pattern


### Return type

[**ZtListEdgeNodes200Response**](ZtListEdgeNodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

