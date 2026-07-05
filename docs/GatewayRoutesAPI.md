# \GatewayRoutesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreateRoute**](GatewayRoutesAPI.md#GatewayCreateRoute) | **Post** /v1/gateway/routes | Create routing rule
[**GatewayDeleteRoute**](GatewayRoutesAPI.md#GatewayDeleteRoute) | **Delete** /v1/gateway/routes/{id} | Delete routing rule
[**GatewayGetRoute**](GatewayRoutesAPI.md#GatewayGetRoute) | **Get** /v1/gateway/routes/{id} | Get routing rule
[**GatewayListRoutes**](GatewayRoutesAPI.md#GatewayListRoutes) | **Get** /v1/gateway/routes | List custom routing rules
[**GatewayUpdateRoute**](GatewayRoutesAPI.md#GatewayUpdateRoute) | **Put** /v1/gateway/routes/{id} | Update routing rule



## GatewayCreateRoute

> GatewayRoutingRule GatewayCreateRoute(ctx).GatewayRoutingRule(gatewayRoutingRule).Execute()

Create routing rule

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
	gatewayRoutingRule := *openapiclient.NewGatewayRoutingRule("Id_example", *openapiclient.NewGatewayRoutingRuleMatch(), *openapiclient.NewGatewayRoutingRuleAction()) // GatewayRoutingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayRoutesAPI.GatewayCreateRoute(context.Background()).GatewayRoutingRule(gatewayRoutingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoutesAPI.GatewayCreateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateRoute`: GatewayRoutingRule
	fmt.Fprintf(os.Stdout, "Response from `GatewayRoutesAPI.GatewayCreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayRoutingRule** | [**GatewayRoutingRule**](GatewayRoutingRule.md) |  | 

### Return type

[**GatewayRoutingRule**](GatewayRoutingRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDeleteRoute

> GatewayDeleteRoute(ctx, id).Execute()

Delete routing rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayRoutesAPI.GatewayDeleteRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoutesAPI.GatewayDeleteRoute``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGatewayDeleteRouteRequest struct via the builder pattern


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


## GatewayGetRoute

> GatewayRoutingRule GatewayGetRoute(ctx, id).Execute()

Get routing rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayRoutesAPI.GatewayGetRoute(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoutesAPI.GatewayGetRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetRoute`: GatewayRoutingRule
	fmt.Fprintf(os.Stdout, "Response from `GatewayRoutesAPI.GatewayGetRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayRoutingRule**](GatewayRoutingRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListRoutes

> []GatewayRoutingRule GatewayListRoutes(ctx).Execute()

List custom routing rules

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
	resp, r, err := apiClient.GatewayRoutesAPI.GatewayListRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoutesAPI.GatewayListRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListRoutes`: []GatewayRoutingRule
	fmt.Fprintf(os.Stdout, "Response from `GatewayRoutesAPI.GatewayListRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListRoutesRequest struct via the builder pattern


### Return type

[**[]GatewayRoutingRule**](GatewayRoutingRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateRoute

> GatewayRoutingRule GatewayUpdateRoute(ctx, id).GatewayRoutingRule(gatewayRoutingRule).Execute()

Update routing rule

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	gatewayRoutingRule := *openapiclient.NewGatewayRoutingRule("Id_example", *openapiclient.NewGatewayRoutingRuleMatch(), *openapiclient.NewGatewayRoutingRuleAction()) // GatewayRoutingRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayRoutesAPI.GatewayUpdateRoute(context.Background(), id).GatewayRoutingRule(gatewayRoutingRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoutesAPI.GatewayUpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayUpdateRoute`: GatewayRoutingRule
	fmt.Fprintf(os.Stdout, "Response from `GatewayRoutesAPI.GatewayUpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gatewayRoutingRule** | [**GatewayRoutingRule**](GatewayRoutingRule.md) |  | 

### Return type

[**GatewayRoutingRule**](GatewayRoutingRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

