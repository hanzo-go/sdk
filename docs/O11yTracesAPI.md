# \O11yTracesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yGetDependencies**](O11yTracesAPI.md#O11yGetDependencies) | **Get** /v1/o11y/dependencies | Get service dependency graph
[**O11yGetTrace**](O11yTracesAPI.md#O11yGetTrace) | **Get** /v1/o11y/traces/{traceId} | Get trace by ID
[**O11yListServiceOperations**](O11yTracesAPI.md#O11yListServiceOperations) | **Get** /v1/o11y/services/{name}/operations | List service operations
[**O11yListServices**](O11yTracesAPI.md#O11yListServices) | **Get** /v1/o11y/services | List traced services
[**O11ySearchTraces**](O11yTracesAPI.md#O11ySearchTraces) | **Get** /v1/o11y/traces | Search traces



## O11yGetDependencies

> O11yGetDependencies200Response O11yGetDependencies(ctx).Start(start).End(end).Execute()

Get service dependency graph

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
	start := time.Now() // time.Time | 
	end := time.Now() // time.Time | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yTracesAPI.O11yGetDependencies(context.Background()).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yTracesAPI.O11yGetDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yGetDependencies`: O11yGetDependencies200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yTracesAPI.O11yGetDependencies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yGetDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 

### Return type

[**O11yGetDependencies200Response**](O11yGetDependencies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yGetTrace

> O11yTrace O11yGetTrace(ctx, traceId).Execute()

Get trace by ID

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
	traceId := "traceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yTracesAPI.O11yGetTrace(context.Background(), traceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yTracesAPI.O11yGetTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yGetTrace`: O11yTrace
	fmt.Fprintf(os.Stdout, "Response from `O11yTracesAPI.O11yGetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**O11yTrace**](O11yTrace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yListServiceOperations

> O11yListServiceOperations200Response O11yListServiceOperations(ctx, name).SpanKind(spanKind).Execute()

List service operations

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
	spanKind := "spanKind_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yTracesAPI.O11yListServiceOperations(context.Background(), name).SpanKind(spanKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yTracesAPI.O11yListServiceOperations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yListServiceOperations`: O11yListServiceOperations200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yTracesAPI.O11yListServiceOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiO11yListServiceOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **spanKind** | **string** |  | 

### Return type

[**O11yListServiceOperations200Response**](O11yListServiceOperations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yListServices

> O11yListServices200Response O11yListServices(ctx).Execute()

List traced services

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
	resp, r, err := apiClient.O11yTracesAPI.O11yListServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yTracesAPI.O11yListServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yListServices`: O11yListServices200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yTracesAPI.O11yListServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yListServicesRequest struct via the builder pattern


### Return type

[**O11yListServices200Response**](O11yListServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11ySearchTraces

> O11ySearchTraces200Response O11ySearchTraces(ctx).Service(service).Operation(operation).Tags(tags).MinDuration(minDuration).MaxDuration(maxDuration).Start(start).End(end).Limit(limit).Execute()

Search traces

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
	service := "service_example" // string | Filter by service name (optional)
	operation := "operation_example" // string | Filter by operation name (optional)
	tags := "tags_example" // string | Key-value tag filter (e.g. http.status_code=500) (optional)
	minDuration := "minDuration_example" // string | Minimum trace duration (e.g. 100ms, 1s) (optional)
	maxDuration := "maxDuration_example" // string |  (optional)
	start := time.Now() // time.Time |  (optional)
	end := time.Now() // time.Time |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yTracesAPI.O11ySearchTraces(context.Background()).Service(service).Operation(operation).Tags(tags).MinDuration(minDuration).MaxDuration(maxDuration).Start(start).End(end).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yTracesAPI.O11ySearchTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11ySearchTraces`: O11ySearchTraces200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yTracesAPI.O11ySearchTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11ySearchTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | Filter by service name | 
 **operation** | **string** | Filter by operation name | 
 **tags** | **string** | Key-value tag filter (e.g. http.status_code&#x3D;500) | 
 **minDuration** | **string** | Minimum trace duration (e.g. 100ms, 1s) | 
 **maxDuration** | **string** |  | 
 **start** | **time.Time** |  | 
 **end** | **time.Time** |  | 
 **limit** | **int32** |  | [default to 20]

### Return type

[**O11ySearchTraces200Response**](O11ySearchTraces200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

