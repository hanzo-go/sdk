# \IamSessionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddSession**](IamSessionsAPI.md#IamApiControllerAddSession) | **Post** /v1/iam/sessions | Api Controller Add Session
[**IamApiControllerDeleteSession**](IamSessionsAPI.md#IamApiControllerDeleteSession) | **Delete** /v1/iam/sessions/{id} | Api Controller Delete Session
[**IamApiControllerGetSessions**](IamSessionsAPI.md#IamApiControllerGetSessions) | **Get** /v1/iam/sessions | Api Controller Get Sessions
[**IamApiControllerGetSingleSession**](IamSessionsAPI.md#IamApiControllerGetSingleSession) | **Get** /v1/iam/sessions/{id} | Api Controller Get Single Session
[**IamApiControllerIsSessionDuplicated**](IamSessionsAPI.md#IamApiControllerIsSessionDuplicated) | **Get** /v1/iam/is-session-duplicated | Api Controller Is Session Duplicated
[**IamApiControllerUpdateSession**](IamSessionsAPI.md#IamApiControllerUpdateSession) | **Put** /v1/iam/sessions/{id} | Api Controller Update Session



## IamApiControllerAddSession

> IamControllersResponse IamApiControllerAddSession(ctx).IamObjectSession(iamObjectSession).Execute()

Api Controller Add Session



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
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerAddSession(context.Background()).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerAddSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerAddSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to add | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteSession

> IamControllersResponse IamApiControllerDeleteSession(ctx, id).IamObjectSession(iamObjectSession).Execute()

Api Controller Delete Session



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerDeleteSession(context.Background(), id).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerDeleteSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerDeleteSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to delete | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSessions

> []string IamApiControllerGetSessions(ctx).Owner(owner).Execute()

Api Controller Get Sessions



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
	owner := "owner_example" // string | The organization name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerGetSessions(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerGetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSessions`: []string
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerGetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The organization name | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetSingleSession

> []string IamApiControllerGetSingleSession(ctx, id).SessionPkId(sessionPkId).Execute()

Api Controller Get Single Session



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
	sessionPkId := "sessionPkId_example" // string | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in)
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerGetSingleSession(context.Background(), id).SessionPkId(sessionPkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerGetSingleSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetSingleSession`: []string
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerGetSingleSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetSingleSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionPkId** | **string** | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in) | 


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerIsSessionDuplicated

> []string IamApiControllerIsSessionDuplicated(ctx).SessionPkId(sessionPkId).SessionId(sessionId).Execute()

Api Controller Is Session Duplicated



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
	sessionPkId := "sessionPkId_example" // string | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in)
	sessionId := "sessionId_example" // string | The specific session ID to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerIsSessionDuplicated(context.Background()).SessionPkId(sessionPkId).SessionId(sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerIsSessionDuplicated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerIsSessionDuplicated`: []string
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerIsSessionDuplicated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerIsSessionDuplicatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionPkId** | **string** | The session ID in format: organization/user/application (e.g., built-in/admin/app-built-in) | 
 **sessionId** | **string** | The specific session ID to check | 

### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateSession

> IamControllersResponse IamApiControllerUpdateSession(ctx, id).IamObjectSession(iamObjectSession).Execute()

Api Controller Update Session



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectSession := *openapiclient.NewIamObjectSession() // IamObjectSession | The session object to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamSessionsAPI.IamApiControllerUpdateSession(context.Background(), id).IamObjectSession(iamObjectSession).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamSessionsAPI.IamApiControllerUpdateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateSession`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamSessionsAPI.IamApiControllerUpdateSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectSession** | [**IamObjectSession**](IamObjectSession.md) | The session object to update | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

