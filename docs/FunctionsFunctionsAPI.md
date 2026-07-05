# \FunctionsFunctionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FunctionsCreateFunction**](FunctionsFunctionsAPI.md#FunctionsCreateFunction) | **Post** /v1/functions | Create a function
[**FunctionsDeleteFunction**](FunctionsFunctionsAPI.md#FunctionsDeleteFunction) | **Delete** /v1/functions/{name} | Delete a function
[**FunctionsGetFunction**](FunctionsFunctionsAPI.md#FunctionsGetFunction) | **Get** /v1/functions/{name} | Get a function (with triggers, recent invocations, secrets)
[**FunctionsGetFunctionLogs**](FunctionsFunctionsAPI.md#FunctionsGetFunctionLogs) | **Get** /v1/functions/{name}/logs | Get the latest invocation logs
[**FunctionsInvokeFunction**](FunctionsFunctionsAPI.md#FunctionsInvokeFunction) | **Post** /v1/functions/{name}/invoke | Invoke a function (metered compute)
[**FunctionsListDeployments**](FunctionsFunctionsAPI.md#FunctionsListDeployments) | **Get** /v1/functions/deployments | List deployed functions
[**FunctionsListFunctionSecrets**](FunctionsFunctionsAPI.md#FunctionsListFunctionSecrets) | **Get** /v1/functions/secrets | List function secrets (names only)
[**FunctionsListFunctions**](FunctionsFunctionsAPI.md#FunctionsListFunctions) | **Get** /v1/functions | List functions
[**FunctionsListInvocations**](FunctionsFunctionsAPI.md#FunctionsListInvocations) | **Get** /v1/functions/{name}/invocations | List a function&#39;s invocations
[**FunctionsListTriggers**](FunctionsFunctionsAPI.md#FunctionsListTriggers) | **Get** /v1/functions/triggers | List triggers across functions



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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsCreateFunction(context.Background()).FunctionsCreateFunctionRequest(functionsCreateFunctionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsCreateFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsCreateFunction`: FunctionsFunction
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsCreateFunction`: %v\n", resp)
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
	r, err := apiClient.FunctionsFunctionsAPI.FunctionsDeleteFunction(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsDeleteFunction``: %v\n", err)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsGetFunction(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsGetFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetFunction`: FunctionsFunctionDetail
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsGetFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsGetFunctionLogs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsGetFunctionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsGetFunctionLogs`: FunctionsGetFunctionLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsGetFunctionLogs`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsInvokeFunction(context.Background(), name).FunctionsInvokeRequest(functionsInvokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsInvokeFunction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsInvokeFunction`: FunctionsInvocation
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsInvokeFunction`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsListDeployments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListDeployments`: FunctionsListFunctions200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsListDeployments`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsListFunctionSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsListFunctionSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListFunctionSecrets`: FunctionsListFunctionSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsListFunctionSecrets`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsListFunctions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsListFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListFunctions`: FunctionsListFunctions200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsListFunctions`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsListInvocations(context.Background(), name).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsListInvocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListInvocations`: FunctionsListInvocations200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsListInvocations`: %v\n", resp)
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
	resp, r, err := apiClient.FunctionsFunctionsAPI.FunctionsListTriggers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FunctionsFunctionsAPI.FunctionsListTriggers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FunctionsListTriggers`: FunctionsListTriggers200Response
	fmt.Fprintf(os.Stdout, "Response from `FunctionsFunctionsAPI.FunctionsListTriggers`: %v\n", resp)
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

