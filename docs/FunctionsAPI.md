# \FunctionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFunctionsByName**](FunctionsAPI.md#DeleteFunctionsByName) | **Delete** /v1/functions/{name} | Removes one of the caller org&#39;s functions and answers 204.
[**GetFunctions**](FunctionsAPI.md#GetFunctions) | **Get** /v1/functions | Is every serverless function the caller&#39;s org has published, each with its real 7-day rollup.
[**GetFunctionsByName**](FunctionsAPI.md#GetFunctionsByName) | **Get** /v1/functions/{name} | Is one function with everything a detail page needs in one round-trip: its definition, its 7-day rollup, its trigger, its twenty most recent invocations and the NAMES of the secrets it mounts.
[**GetFunctionsByNameInvocations**](FunctionsAPI.md#GetFunctionsByNameInvocations) | **Get** /v1/functions/{name}/invocations | Is one function&#39;s past runs, newest first — each with its status, HTTP code, method, time and duration.
[**GetFunctionsByNameLogs**](FunctionsAPI.md#GetFunctionsByNameLogs) | **Get** /v1/functions/{name}/logs | Is the output of a function&#39;s most recent run — its error text when that run failed, else what it printed.
[**GetFunctionsDeployments**](FunctionsAPI.md#GetFunctionsDeployments) | **Get** /v1/functions/deployments | Is what is live right now — each function&#39;s current record IS its live deployment, so this is the deployment inventory.
[**GetFunctionsMetrics**](FunctionsAPI.md#GetFunctionsMetrics) | **Get** /v1/functions/metrics | Is the org&#39;s serverless dashboard over a window: a per-function invocation costLine and how those invocations ended.
[**GetFunctionsSecrets**](FunctionsAPI.md#GetFunctionsSecrets) | **Get** /v1/functions/secrets | Is the NAMES of the secrets the caller org&#39;s functions mount.
[**GetFunctionsTriggers**](FunctionsAPI.md#GetFunctionsTriggers) | **Get** /v1/functions/triggers | Is what calls the caller org&#39;s functions — one row per function.
[**PostFunctions**](FunctionsAPI.md#PostFunctions) | **Post** /v1/functions | Publishes a serverless function under the caller&#39;s org and answers 201 with it.
[**PostFunctionsByNameInvoke**](FunctionsAPI.md#PostFunctionsByNameInvoke) | **Post** /v1/functions/{name}/invoke | Runs a function and records a REAL invocation.



## DeleteFunctionsByName

> map[string]interface{} DeleteFunctionsByName(ctx, name).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.DeleteFunctionsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.DeleteFunctionsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFunctionsByName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.DeleteFunctionsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFunctionsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctions

> FnList GetFunctions(ctx).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctions`: FnList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsRequest struct via the builder pattern


### Return type

[**FnList**](FnList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsByName

> FunctionDetail GetFunctionsByName(ctx, name).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsByName`: FunctionDetail
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FunctionDetail**](FunctionDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsByNameInvocations

> InvocationList GetFunctionsByNameInvocations(ctx, name).Limit(limit).Execute()

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
	limit := int32(56) // int32 | Limit caps the page, defaulting to 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsByNameInvocations(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsByNameInvocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsByNameInvocations`: InvocationList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsByNameInvocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsByNameInvocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps the page, defaulting to 100. | 

### Return type

[**InvocationList**](InvocationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsByNameLogs

> LogLines GetFunctionsByNameLogs(ctx, name).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsByNameLogs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsByNameLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsByNameLogs`: LogLines
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsByNameLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the function the URL names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsByNameLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogLines**](LogLines.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsDeployments

> FnList GetFunctionsDeployments(ctx).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsDeployments`: FnList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsDeployments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsDeploymentsRequest struct via the builder pattern


### Return type

[**FnList**](FnList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsMetrics

> Usage GetFunctionsMetrics(ctx).Range_(range_).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsMetrics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsMetrics`: Usage
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Range is 1H, 6H, 24H (the default), 7D or 30D. Anything else falls back to 24H rather than failing. | 

### Return type

[**Usage**](Usage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsSecrets

> SecretList GetFunctionsSecrets(ctx).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsSecrets`: SecretList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsSecretsRequest struct via the builder pattern


### Return type

[**SecretList**](SecretList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunctionsTriggers

> TriggerList GetFunctionsTriggers(ctx).Execute()

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
	resp, r, err := apiClient.FunctionsAPI.GetFunctionsTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.GetFunctionsTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFunctionsTriggers`: TriggerList
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.GetFunctionsTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunctionsTriggersRequest struct via the builder pattern


### Return type

[**TriggerList**](TriggerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFunctions

> FunctionView PostFunctions(ctx).Definition(definition).Execute()

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
	definition := *openapiclient.NewDefinition("Name_example") // Definition | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.PostFunctions(context.Background()).Definition(definition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.PostFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFunctions`: FunctionView
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.PostFunctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **definition** | [**Definition**](Definition.md) |  | 

### Return type

[**FunctionView**](FunctionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFunctionsByNameInvoke

> InvocationView PostFunctionsByNameInvoke(ctx, name).InvokeReq(invokeReq).Execute()

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
	invokeReq := *openapiclient.NewInvokeReq() // InvokeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FunctionsAPI.PostFunctionsByNameInvoke(context.Background(), name).InvokeReq(invokeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsAPI.PostFunctionsByNameInvoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFunctionsByNameInvoke`: InvocationView
	fmt.Fprintf(os.Stdout, "Response from `FunctionsAPI.PostFunctionsByNameInvoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFunctionsByNameInvokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invokeReq** | [**InvokeReq**](InvokeReq.md) |  | 

### Return type

[**InvocationView**](InvocationView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

