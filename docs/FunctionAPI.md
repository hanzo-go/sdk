# \FunctionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFunctionByName**](FunctionAPI.md#DeleteFunctionByName) | **Delete** /v1/function/{name} | Removes one of the caller org&#39;s functions and answers 204.
[**GetFunction**](FunctionAPI.md#GetFunction) | **Get** /v1/function | Is every serverless function the caller&#39;s org has published, each with its real 7-day rollup.
[**GetFunctionByName**](FunctionAPI.md#GetFunctionByName) | **Get** /v1/function/{name} | Is one function with everything a detail page needs in one round-trip: its definition, its 7-day rollup, its trigger, its twenty most recent invocations and the NAMES of the secrets it mounts.
[**GetFunctionByNameInvocations**](FunctionAPI.md#GetFunctionByNameInvocations) | **Get** /v1/function/{name}/invocations | Is one function&#39;s past runs, newest first — each with its status, HTTP code, method, time and duration.
[**GetFunctionByNameLogs**](FunctionAPI.md#GetFunctionByNameLogs) | **Get** /v1/function/{name}/logs | Is the output of a function&#39;s most recent run — its error text when that run failed, else what it printed.
[**GetFunctionDeployments**](FunctionAPI.md#GetFunctionDeployments) | **Get** /v1/function/deployments | Is what is live right now — each function&#39;s current record IS its live deployment, so this is the deployment inventory.
[**GetFunctionMetrics**](FunctionAPI.md#GetFunctionMetrics) | **Get** /v1/function/metrics | Is the org&#39;s serverless dashboard over a window: a per-function invocation costLine and how those invocations ended.
[**GetFunctionSecrets**](FunctionAPI.md#GetFunctionSecrets) | **Get** /v1/function/secrets | Is the NAMES of the secrets the caller org&#39;s functions mount.
[**GetFunctionTriggers**](FunctionAPI.md#GetFunctionTriggers) | **Get** /v1/function/triggers | Is what calls the caller org&#39;s functions — one row per function.
[**PostFunction**](FunctionAPI.md#PostFunction) | **Post** /v1/function | Publishes a serverless function under the caller&#39;s org and answers 201 with it.
[**PostFunctionByNameInvoke**](FunctionAPI.md#PostFunctionByNameInvoke) | **Post** /v1/function/{name}/invoke | Runs a function and records a REAL invocation.



## DeleteFunctionByName

> DeleteFunctionByName(ctx, name).Execute()

Removes one of the caller org's functions and answers 204.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the function the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FunctionAPI.DeleteFunctionByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.DeleteFunctionByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunction

> FunctionFnList GetFunction(ctx).Execute()

Is every serverless function the caller's org has published, each with its real 7-day rollup.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunction(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunction`: FunctionFnList
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionRequest struct via the builder pattern


### Return type

[**FunctionFnList**](FunctionFnList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionByName

> FunctionFunctionDetail GetFunctionByName(ctx, name).Execute()

Is one function with everything a detail page needs in one round-trip: its definition, its 7-day rollup, its trigger, its twenty most recent invocations and the NAMES of the secrets it mounts.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the function the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionByName`: FunctionFunctionDetail
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionFunctionDetail**](FunctionFunctionDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionByNameInvocations

> FunctionInvocationList GetFunctionByNameInvocations(ctx, name).Limit(limit).Execute()

Is one function's past runs, newest first — each with its status, HTTP code, method, time and duration.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the function the URL names.
	limit := int64(789) // int64 | Limit caps the page, defaulting to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionByNameInvocations(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionByNameInvocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionByNameInvocations`: FunctionInvocationList
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionByNameInvocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionByNameInvocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit caps the page, defaulting to 100. | 

### Return type

[**FunctionInvocationList**](FunctionInvocationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionByNameLogs

> FunctionLogLines GetFunctionByNameLogs(ctx, name).Execute()

Is the output of a function's most recent run — its error text when that run failed, else what it printed.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | Name is the function the URL names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionByNameLogs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionByNameLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionByNameLogs`: FunctionLogLines
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionByNameLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionByNameLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionLogLines**](FunctionLogLines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionDeployments

> FunctionFnList GetFunctionDeployments(ctx).Execute()

Is what is live right now — each function's current record IS its live deployment, so this is the deployment inventory.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionDeployments`: FunctionFnList
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionDeploymentsRequest struct via the builder pattern


### Return type

[**FunctionFnList**](FunctionFnList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionMetrics

> FunctionUsage GetFunctionMetrics(ctx).Range_(range_).Execute()

Is the org's serverless dashboard over a window: a per-function invocation costLine and how those invocations ended.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	range_ := "range__example" // string | Range is 1H, 6H, 24H (the default), 7D or 30D. Anything else falls back to 24H rather than failing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionMetrics`: FunctionUsage
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is 1H, 6H, 24H (the default), 7D or 30D. Anything else falls back to 24H rather than failing. | 

### Return type

[**FunctionUsage**](FunctionUsage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionSecrets

> FunctionSecretList GetFunctionSecrets(ctx).Execute()

Is the NAMES of the secrets the caller org's functions mount.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionSecrets`: FunctionSecretList
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionSecretsRequest struct via the builder pattern


### Return type

[**FunctionSecretList**](FunctionSecretList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionTriggers

> FunctionTriggerList GetFunctionTriggers(ctx).Execute()

Is what calls the caller org's functions — one row per function.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.GetFunctionTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.GetFunctionTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionTriggers`: FunctionTriggerList
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.GetFunctionTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionTriggersRequest struct via the builder pattern


### Return type

[**FunctionTriggerList**](FunctionTriggerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFunction

> FunctionFunctionView PostFunction(ctx).FunctionDefinition(functionDefinition).Execute()

Publishes a serverless function under the caller's org and answers 201 with it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	functionDefinition := *openapiclient.NewFunctionDefinition("Name_example") // FunctionDefinition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.PostFunction(context.Background()).FunctionDefinition(functionDefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.PostFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFunction`: FunctionFunctionView
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.PostFunction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFunctionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionDefinition** | [**FunctionDefinition**](FunctionDefinition.md) |  | 

### Return type

[**FunctionFunctionView**](FunctionFunctionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFunctionByNameInvoke

> FunctionInvocationView PostFunctionByNameInvoke(ctx, name).FunctionInvokeReq(functionInvokeReq).Execute()

Runs a function and records a REAL invocation.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	name := "name_example" // string | 
	functionInvokeReq := *openapiclient.NewFunctionInvokeReq() // FunctionInvokeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionAPI.PostFunctionByNameInvoke(context.Background(), name).FunctionInvokeReq(functionInvokeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionAPI.PostFunctionByNameInvoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFunctionByNameInvoke`: FunctionInvocationView
	fmt.Fprintf(os.Stdout, "Response from `FunctionAPI.PostFunctionByNameInvoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFunctionByNameInvokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **functionInvokeReq** | [**FunctionInvokeReq**](FunctionInvokeReq.md) |  | 

### Return type

[**FunctionInvocationView**](FunctionInvocationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

