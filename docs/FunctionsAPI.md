# \FunctionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeCreateFunction**](FunctionsAPI.md#EdgeCreateFunction) | **Post** /v1/edge/functions | Create function
[**EdgeDeleteFunction**](FunctionsAPI.md#EdgeDeleteFunction) | **Delete** /v1/edge/functions/{slug} | Delete function
[**EdgeDeployFunction**](FunctionsAPI.md#EdgeDeployFunction) | **Post** /v1/edge/functions/{slug}/deploy | Deploy function
[**EdgeGetFunction**](FunctionsAPI.md#EdgeGetFunction) | **Get** /v1/edge/functions/{slug} | Get function
[**EdgeGetFunctionMetrics**](FunctionsAPI.md#EdgeGetFunctionMetrics) | **Get** /v1/edge/functions/{slug}/metrics | Get function metrics
[**EdgeInvokeFunction**](FunctionsAPI.md#EdgeInvokeFunction) | **Post** /v1/edge/functions/{slug}/invoke | Invoke function
[**EdgeListFunctions**](FunctionsAPI.md#EdgeListFunctions) | **Get** /v1/edge/functions | List functions
[**EdgeUpdateFunction**](FunctionsAPI.md#EdgeUpdateFunction) | **Put** /v1/edge/functions/{slug} | Update function
[**FunctionsCreateFunction**](FunctionsAPI.md#FunctionsCreateFunction) | **Post** /v1/functions | Create a function
[**FunctionsDeleteFunction**](FunctionsAPI.md#FunctionsDeleteFunction) | **Delete** /v1/functions/{name} | Delete a function
[**FunctionsGetFunction**](FunctionsAPI.md#FunctionsGetFunction) | **Get** /v1/functions/{name} | Get a function (with triggers, recent invocations, secrets)
[**FunctionsGetFunctionLogs**](FunctionsAPI.md#FunctionsGetFunctionLogs) | **Get** /v1/functions/{name}/logs | Get the latest invocation logs
[**FunctionsInvokeFunction**](FunctionsAPI.md#FunctionsInvokeFunction) | **Post** /v1/functions/{name}/invoke | Invoke a function (metered compute)
[**FunctionsListDeployments**](FunctionsAPI.md#FunctionsListDeployments) | **Get** /v1/functions/deployments | List deployed functions
[**FunctionsListFunctionSecrets**](FunctionsAPI.md#FunctionsListFunctionSecrets) | **Get** /v1/functions/secrets | List function secrets (names only)
[**FunctionsListFunctions**](FunctionsAPI.md#FunctionsListFunctions) | **Get** /v1/functions | List functions
[**FunctionsListInvocations**](FunctionsAPI.md#FunctionsListInvocations) | **Get** /v1/functions/{name}/invocations | List a function&#39;s invocations
[**FunctionsListTriggers**](FunctionsAPI.md#FunctionsListTriggers) | **Get** /v1/functions/triggers | List triggers across functions



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
	resp, r, err := apiClient.FunctionsAPI.EdgeCreateFunction(context.Background()).EdgeFunctionCreate(edgeFunctionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeCreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeCreateFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeCreateFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeDeleteFunction(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeDeleteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeleteFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeDeleteFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeDeployFunction(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeDeployFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeployFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeDeployFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeGetFunction(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeGetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeGetFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeGetFunctionMetrics(context.Background(), slug).From(from).To(to).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeGetFunctionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeGetFunctionMetrics`: []EdgeFunctionMetrics
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeGetFunctionMetrics`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeInvokeFunction(context.Background(), slug).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeInvokeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeInvokeFunction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeInvokeFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeListFunctions(context.Background()).Page(page).PageSize(pageSize).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeListFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeListFunctions`: []EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeListFunctions`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsAPI.EdgeUpdateFunction(context.Background(), slug).EdgeFunctionUpdate(edgeFunctionUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.EdgeUpdateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeUpdateFunction`: EdgeFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.EdgeUpdateFunction`: %v\n", resp)
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


## FunctionsCreateFunction

> FunctionsFunction FunctionsCreateFunction(ctx).FunctionsCreateFunctionRequest(functionsCreateFunctionRequest).Execute()

Create a function

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
	functionsCreateFunctionRequest := *openapiclient.NewFunctionsCreateFunctionRequest("Name_example") // FunctionsCreateFunctionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsCreateFunction(context.Background()).FunctionsCreateFunctionRequest(functionsCreateFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsCreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsCreateFunction`: FunctionsFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsCreateFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsCreateFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionsCreateFunctionRequest** | [**FunctionsCreateFunctionRequest**](FunctionsCreateFunctionRequest.md) |  | 

### Return type

[**FunctionsFunction**](FunctionsFunction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsDeleteFunction

> FunctionsDeleteFunction(ctx, name).Execute()

Delete a function

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionsAPI.FunctionsDeleteFunction(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsDeleteFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsDeleteFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetFunction

> FunctionsFunctionDetail FunctionsGetFunction(ctx, name).Execute()

Get a function (with triggers, recent invocations, secrets)

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGetFunction(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetFunction`: FunctionsFunctionDetail
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGetFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionsFunctionDetail**](FunctionsFunctionDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsGetFunctionLogs

> FunctionsGetFunctionLogs200Response FunctionsGetFunctionLogs(ctx, name).Execute()

Get the latest invocation logs

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsGetFunctionLogs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsGetFunctionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetFunctionLogs`: FunctionsGetFunctionLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsGetFunctionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsGetFunctionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionsGetFunctionLogs200Response**](FunctionsGetFunctionLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsInvokeFunction

> FunctionsInvocation FunctionsInvokeFunction(ctx, name).FunctionsInvokeRequest(functionsInvokeRequest).Execute()

Invoke a function (metered compute)

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
	name := "name_example" // string | 
	functionsInvokeRequest := *openapiclient.NewFunctionsInvokeRequest() // FunctionsInvokeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsInvokeFunction(context.Background(), name).FunctionsInvokeRequest(functionsInvokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsInvokeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsInvokeFunction`: FunctionsInvocation
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsInvokeFunction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsInvokeFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionsInvokeRequest** | [**FunctionsInvokeRequest**](FunctionsInvokeRequest.md) |  | 

### Return type

[**FunctionsInvocation**](FunctionsInvocation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsListDeployments

> FunctionsListFunctions200Response FunctionsListDeployments(ctx).Execute()

List deployed functions

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
	resp, r, err := apiClient.FunctionsAPI.FunctionsListDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListDeployments`: FunctionsListFunctions200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsListDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsListDeploymentsRequest struct via the builder pattern


### Return type

[**FunctionsListFunctions200Response**](FunctionsListFunctions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsListFunctionSecrets

> FunctionsListFunctionSecrets200Response FunctionsListFunctionSecrets(ctx).Execute()

List function secrets (names only)

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
	resp, r, err := apiClient.FunctionsAPI.FunctionsListFunctionSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsListFunctionSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListFunctionSecrets`: FunctionsListFunctionSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsListFunctionSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsListFunctionSecretsRequest struct via the builder pattern


### Return type

[**FunctionsListFunctionSecrets200Response**](FunctionsListFunctionSecrets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsListFunctions

> FunctionsListFunctions200Response FunctionsListFunctions(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsListFunctions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsListFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListFunctions`: FunctionsListFunctions200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsListFunctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsListFunctionsRequest struct via the builder pattern


### Return type

[**FunctionsListFunctions200Response**](FunctionsListFunctions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsListInvocations

> FunctionsListInvocations200Response FunctionsListInvocations(ctx, name).Limit(limit).Execute()

List a function's invocations

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
	name := "name_example" // string | 
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.FunctionsListInvocations(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsListInvocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListInvocations`: FunctionsListInvocations200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsListInvocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsListInvocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 

### Return type

[**FunctionsListInvocations200Response**](FunctionsListInvocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FunctionsListTriggers

> FunctionsListTriggers200Response FunctionsListTriggers(ctx).Execute()

List triggers across functions

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
	resp, r, err := apiClient.FunctionsAPI.FunctionsListTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.FunctionsListTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListTriggers`: FunctionsListTriggers200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.FunctionsListTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsListTriggersRequest struct via the builder pattern


### Return type

[**FunctionsListTriggers200Response**](FunctionsListTriggers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

