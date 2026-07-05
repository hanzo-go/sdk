# \OperativeSessionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperativeCreateSession**](OperativeSessionsAPI.md#OperativeCreateSession) | **Post** /v1/operative/sessions | Create a new operative session
[**OperativeDeleteSession**](OperativeSessionsAPI.md#OperativeDeleteSession) | **Delete** /v1/operative/sessions/{sessionId} | Terminate and delete a session
[**OperativeGetSession**](OperativeSessionsAPI.md#OperativeGetSession) | **Get** /v1/operative/sessions/{sessionId} | Get session details
[**OperativeListSessions**](OperativeSessionsAPI.md#OperativeListSessions) | **Get** /v1/operative/sessions | List all active sessions
[**OperativeResetSession**](OperativeSessionsAPI.md#OperativeResetSession) | **Post** /v1/operative/sessions/{sessionId}/reset | Reset the session environment



## OperativeCreateSession

> OperativeSession OperativeCreateSession(ctx).OperativeCreateSessionRequest(operativeCreateSessionRequest).Execute()

Create a new operative session



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
	operativeCreateSessionRequest := *openapiclient.NewOperativeCreateSessionRequest("Provider_example") // OperativeCreateSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeSessionsAPI.OperativeCreateSession(context.Background()).OperativeCreateSessionRequest(operativeCreateSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeSessionsAPI.OperativeCreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeCreateSession`: OperativeSession
	fmt.Fprintf(os.Stdout, "Response from `OperativeSessionsAPI.OperativeCreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperativeCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operativeCreateSessionRequest** | [**OperativeCreateSessionRequest**](OperativeCreateSessionRequest.md) |  | 

### Return type

[**OperativeSession**](OperativeSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperativeDeleteSession

> OperativeDeleteSession(ctx, sessionId).Execute()

Terminate and delete a session

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique session identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OperativeSessionsAPI.OperativeDeleteSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeSessionsAPI.OperativeDeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeDeleteSessionRequest struct via the builder pattern


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


## OperativeGetSession

> OperativeSession OperativeGetSession(ctx, sessionId).Execute()

Get session details

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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique session identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeSessionsAPI.OperativeGetSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeSessionsAPI.OperativeGetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeGetSession`: OperativeSession
	fmt.Fprintf(os.Stdout, "Response from `OperativeSessionsAPI.OperativeGetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeGetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperativeSession**](OperativeSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperativeListSessions

> OperativeListSessions200Response OperativeListSessions(ctx).Status(status).Limit(limit).Offset(offset).Execute()

List all active sessions

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
	status := "status_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 20)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeSessionsAPI.OperativeListSessions(context.Background()).Status(status).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeSessionsAPI.OperativeListSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeListSessions`: OperativeListSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `OperativeSessionsAPI.OperativeListSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOperativeListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]

### Return type

[**OperativeListSessions200Response**](OperativeListSessions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperativeResetSession

> OperativeSession OperativeResetSession(ctx, sessionId).Execute()

Reset the session environment



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
	sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Unique session identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OperativeSessionsAPI.OperativeResetSession(context.Background(), sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OperativeSessionsAPI.OperativeResetSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperativeResetSession`: OperativeSession
	fmt.Fprintf(os.Stdout, "Response from `OperativeSessionsAPI.OperativeResetSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Unique session identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperativeResetSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OperativeSession**](OperativeSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

