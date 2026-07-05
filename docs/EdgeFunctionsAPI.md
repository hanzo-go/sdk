# \EdgeFunctionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeCreateFunction**](EdgeFunctionsAPI.md#EdgeCreateFunction) | **Post** /v1/edge/functions | Create function
[**EdgeDeleteFunction**](EdgeFunctionsAPI.md#EdgeDeleteFunction) | **Delete** /v1/edge/functions/{slug} | Delete function
[**EdgeDeployFunction**](EdgeFunctionsAPI.md#EdgeDeployFunction) | **Post** /v1/edge/functions/{slug}/deploy | Deploy function
[**EdgeGetFunction**](EdgeFunctionsAPI.md#EdgeGetFunction) | **Get** /v1/edge/functions/{slug} | Get function
[**EdgeGetFunctionMetrics**](EdgeFunctionsAPI.md#EdgeGetFunctionMetrics) | **Get** /v1/edge/functions/{slug}/metrics | Get function metrics
[**EdgeInvokeFunction**](EdgeFunctionsAPI.md#EdgeInvokeFunction) | **Post** /v1/edge/functions/{slug}/invoke | Invoke function
[**EdgeListFunctions**](EdgeFunctionsAPI.md#EdgeListFunctions) | **Get** /v1/edge/functions | List functions
[**EdgeUpdateFunction**](EdgeFunctionsAPI.md#EdgeUpdateFunction) | **Put** /v1/edge/functions/{slug} | Update function



## EdgeCreateFunction

> EdgeFunction EdgeCreateFunction(ctx).EdgeFunctionCreate(edgeFunctionCreate).Execute()

Create function

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
	edgeFunctionCreate := *openapiclient.NewEdgeFunctionCreate("Slug_example", "Name_example") // EdgeFunctionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeCreateFunction(context.Background()).EdgeFunctionCreate(edgeFunctionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeCreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeCreateFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeCreateFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeFunctionCreate** | [**EdgeFunctionCreate**](EdgeFunctionCreate.md) |  | 

### Return type

[**EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeDeleteFunction

> map[string]interface{} EdgeDeleteFunction(ctx, slug).Execute()

Delete function

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
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeDeleteFunction(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeDeleteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeleteFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeDeleteFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeDeleteFunctionRequest struct via the builder pattern


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


## EdgeDeployFunction

> EdgeFunction EdgeDeployFunction(ctx, slug).Body(body).Execute()

Deploy function



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
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeDeployFunction(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeDeployFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeployFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeDeployFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeDeployFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[**EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetFunction

> EdgeFunction EdgeGetFunction(ctx, slug).Execute()

Get function

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
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeGetFunction(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeGetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeGetFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeGetFunctionMetrics

> []EdgeFunctionMetrics EdgeGetFunctionMetrics(ctx, slug).From(from).To(to).Granularity(granularity).Execute()

Get function metrics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	slug := "slug_example" // string | 
	from := time.Now() // time.Time |  (optional)
	to := time.Now() // time.Time |  (optional)
	granularity := "granularity_example" // string |  (optional) (default to "hour")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeGetFunctionMetrics(context.Background(), slug).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeGetFunctionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunctionMetrics`: []EdgeFunctionMetrics
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeGetFunctionMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeGetFunctionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **granularity** | **string** |  | [default to &quot;hour&quot;]

### Return type

[**[]EdgeFunctionMetrics**](EdgeFunctionMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeInvokeFunction

> map[string]interface{} EdgeInvokeFunction(ctx, slug).RequestBody(requestBody).Execute()

Invoke function



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
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeInvokeFunction(context.Background(), slug).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeInvokeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeInvokeFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeInvokeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeListFunctions

> []EdgeFunction EdgeListFunctions(ctx).Page(page).PageSize(pageSize).Status(status).Execute()

List functions

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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeListFunctions(context.Background()).Page(page).PageSize(pageSize).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeListFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeListFunctions`: []EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeListFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeListFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]
 **status** | **string** |  | 

### Return type

[**[]EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeUpdateFunction

> EdgeFunction EdgeUpdateFunction(ctx, slug).EdgeFunctionUpdate(edgeFunctionUpdate).Execute()

Update function

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
	edgeFunctionUpdate := *openapiclient.NewEdgeFunctionUpdate() // EdgeFunctionUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EdgeFunctionsAPI.EdgeUpdateFunction(context.Background(), slug).EdgeFunctionUpdate(edgeFunctionUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeFunctionsAPI.EdgeUpdateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeUpdateFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `EdgeFunctionsAPI.EdgeUpdateFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeUpdateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeFunctionUpdate** | [**EdgeFunctionUpdate**](EdgeFunctionUpdate.md) |  | 

### Return type

[**EdgeFunction**](EdgeFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

