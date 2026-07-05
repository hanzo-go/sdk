# \ConsoleTracesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleDeleteTrace**](ConsoleTracesAPI.md#ConsoleDeleteTrace) | **Delete** /v1/console/traces/{traceId} | Delete a specific trace
[**ConsoleDeleteTraces**](ConsoleTracesAPI.md#ConsoleDeleteTraces) | **Delete** /v1/console/traces | Delete multiple traces
[**ConsoleGetTrace**](ConsoleTracesAPI.md#ConsoleGetTrace) | **Get** /v1/console/traces/{traceId} | Get a specific trace with full details
[**ConsoleListTraces**](ConsoleTracesAPI.md#ConsoleListTraces) | **Get** /v1/console/traces | Get list of traces



## ConsoleDeleteTrace

> ConsoleDeleteDatasetItem200Response ConsoleDeleteTrace(ctx, traceId).Execute()

Delete a specific trace

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
	resp, r, err := apiClient.ConsoleTracesAPI.ConsoleDeleteTrace(context.Background(), traceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleTracesAPI.ConsoleDeleteTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteTrace`: ConsoleDeleteDatasetItem200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleTracesAPI.ConsoleDeleteTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleDeleteDatasetItem200Response**](ConsoleDeleteDatasetItem200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleDeleteTraces

> ConsoleDeleteDatasetItem200Response ConsoleDeleteTraces(ctx).ConsoleDeleteTracesRequest(consoleDeleteTracesRequest).Execute()

Delete multiple traces

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
	consoleDeleteTracesRequest := *openapiclient.NewConsoleDeleteTracesRequest([]string{"TraceIds_example"}) // ConsoleDeleteTracesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleTracesAPI.ConsoleDeleteTraces(context.Background()).ConsoleDeleteTracesRequest(consoleDeleteTracesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleTracesAPI.ConsoleDeleteTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleDeleteTraces`: ConsoleDeleteDatasetItem200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleTracesAPI.ConsoleDeleteTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleDeleteTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleDeleteTracesRequest** | [**ConsoleDeleteTracesRequest**](ConsoleDeleteTracesRequest.md) |  | 

### Return type

[**ConsoleDeleteDatasetItem200Response**](ConsoleDeleteDatasetItem200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetTrace

> ConsoleTraceWithFullDetails ConsoleGetTrace(ctx, traceId).Execute()

Get a specific trace with full details

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
	resp, r, err := apiClient.ConsoleTracesAPI.ConsoleGetTrace(context.Background(), traceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleTracesAPI.ConsoleGetTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetTrace`: ConsoleTraceWithFullDetails
	fmt.Fprintf(os.Stdout, "Response from `ConsoleTracesAPI.ConsoleGetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleTraceWithFullDetails**](ConsoleTraceWithFullDetails.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListTraces

> ConsoleListTraces200Response ConsoleListTraces(ctx).Page(page).Limit(limit).UserId(userId).Name(name).SessionId(sessionId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).OrderBy(orderBy).Tags(tags).Version(version).Release(release).Environment(environment).Filter(filter).Execute()

Get list of traces

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
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 |  (optional) (default to 50)
	userId := "userId_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	sessionId := "sessionId_example" // string |  (optional)
	fromTimestamp := time.Now() // time.Time |  (optional)
	toTimestamp := time.Now() // time.Time |  (optional)
	orderBy := "orderBy_example" // string | Format: field.asc|desc (e.g., timestamp.desc) (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	version := "version_example" // string |  (optional)
	release := "release_example" // string |  (optional)
	environment := []string{"Inner_example"} // []string |  (optional)
	filter := "filter_example" // string | JSON array of filter conditions (overrides other filters when provided) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleTracesAPI.ConsoleListTraces(context.Background()).Page(page).Limit(limit).UserId(userId).Name(name).SessionId(sessionId).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).OrderBy(orderBy).Tags(tags).Version(version).Release(release).Environment(environment).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleTracesAPI.ConsoleListTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListTraces`: ConsoleListTraces200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleTracesAPI.ConsoleListTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **userId** | **string** |  | 
 **name** | **string** |  | 
 **sessionId** | **string** |  | 
 **fromTimestamp** | **time.Time** |  | 
 **toTimestamp** | **time.Time** |  | 
 **orderBy** | **string** | Format: field.asc|desc (e.g., timestamp.desc) | 
 **tags** | **[]string** |  | 
 **version** | **string** |  | 
 **release** | **string** |  | 
 **environment** | **[]string** |  | 
 **filter** | **string** | JSON array of filter conditions (overrides other filters when provided) | 

### Return type

[**ConsoleListTraces200Response**](ConsoleListTraces200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

