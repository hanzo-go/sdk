# \ConsoleSessionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleGetSession**](ConsoleSessionsAPI.md#ConsoleGetSession) | **Get** /v1/console/sessions/{sessionId} | Get a session with its traces
[**ConsoleListSessions**](ConsoleSessionsAPI.md#ConsoleListSessions) | **Get** /v1/console/sessions | Get sessions



## ConsoleGetSession

> ConsoleSessionWithTraces ConsoleGetSession(ctx, sessionId).Execute()

Get a session with its traces

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
	sessionId := "sessionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleSessionsAPI.ConsoleGetSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleSessionsAPI.ConsoleGetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetSession`: ConsoleSessionWithTraces
	fmt.Fprintf(os.Stdout, "Response from `ConsoleSessionsAPI.ConsoleGetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleSessionWithTraces**](ConsoleSessionWithTraces.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListSessions

> ConsoleListSessions200Response ConsoleListSessions(ctx).Page(page).Limit(limit).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Environment(environment).Execute()

Get sessions

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
	fromTimestamp := time.Now() // time.Time |  (optional)
	toTimestamp := time.Now() // time.Time |  (optional)
	environment := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleSessionsAPI.ConsoleListSessions(context.Background()).Page(page).Limit(limit).FromTimestamp(fromTimestamp).ToTimestamp(toTimestamp).Environment(environment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleSessionsAPI.ConsoleListSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListSessions`: ConsoleListSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleSessionsAPI.ConsoleListSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **fromTimestamp** | **time.Time** |  | 
 **toTimestamp** | **time.Time** |  | 
 **environment** | **[]string** |  | 

### Return type

[**ConsoleListSessions200Response**](ConsoleListSessions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

