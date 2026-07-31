# \AgentSessionsAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudAgentSessionsControllerAppendEvent**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerAppendEvent) | **Post** /v1/agents/sessions/{id}/events | 
[**CloudAgentSessionsControllerGet**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerGet) | **Get** /v1/agents/sessions/{id} | 
[**CloudAgentSessionsControllerList**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerList) | **Get** /v1/agents/sessions | 
[**CloudAgentSessionsControllerMessage**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerMessage) | **Post** /v1/agents/sessions/{id}/message | 
[**CloudAgentSessionsControllerPatch**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerPatch) | **Patch** /v1/agents/sessions/{id} | 
[**CloudAgentSessionsControllerPause**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerPause) | **Post** /v1/agents/sessions/{id}/pause | 
[**CloudAgentSessionsControllerRegister**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerRegister) | **Post** /v1/agents/sessions | 
[**CloudAgentSessionsControllerResume**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerResume) | **Post** /v1/agents/sessions/{id}/resume | 
[**CloudAgentSessionsControllerStop**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerStop) | **Post** /v1/agents/sessions/{id}/stop | 
[**CloudAgentSessionsControllerStream**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerStream) | **Get** /v1/agents/sessions/stream | 
[**CloudAgentSessionsControllerTree**](AgentSessionsAPIAPI.md#CloudAgentSessionsControllerTree) | **Get** /v1/agents/sessions/{id}/tree | 



## CloudAgentSessionsControllerAppendEvent

> CloudAgentsEvent CloudAgentSessionsControllerAppendEvent(ctx, id).CloudAgentsEventRequest(cloudAgentsEventRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsEventRequest := *openapiclient.NewCloudAgentsEventRequest("Kind_example") // CloudAgentsEventRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerAppendEvent(context.Background(), id).CloudAgentsEventRequest(cloudAgentsEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerAppendEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerAppendEvent`: CloudAgentsEvent
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerAppendEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerAppendEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsEventRequest** | [**CloudAgentsEventRequest**](CloudAgentsEventRequest.md) |  | 

### Return type

[**CloudAgentsEvent**](CloudAgentsEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerGet

> CloudAgentsSessionDetail CloudAgentSessionsControllerGet(ctx, id).Execute()





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
	id := "id_example" // string | The session id (sess_...).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerGet`: CloudAgentsSessionDetail
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAgentsSessionDetail**](CloudAgentsSessionDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerList

> CloudAgentSessionsControllerList200Response CloudAgentSessionsControllerList(ctx).Root(root).Parent(parent).Status(status).Limit(limit).Execute()





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
	root := "root_example" // string | Return every session in this tree (rootSessionId == root). (optional)
	parent := "parent_example" // string | Return the direct children of this session. (optional)
	status := "status_example" // string | Filter by status. (optional)
	limit := int32(56) // int32 | Max sessions to return (default 100, max 500). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerList(context.Background()).Root(root).Parent(parent).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerList`: CloudAgentSessionsControllerList200Response
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **root** | **string** | Return every session in this tree (rootSessionId &#x3D;&#x3D; root). | 
 **parent** | **string** | Return the direct children of this session. | 
 **status** | **string** | Filter by status. | 
 **limit** | **int32** | Max sessions to return (default 100, max 500). | 

### Return type

[**CloudAgentSessionsControllerList200Response**](CloudAgentSessionsControllerList200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerMessage

> CloudAgentsControlResult CloudAgentSessionsControllerMessage(ctx, id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsControlRequest := *openapiclient.NewCloudAgentsControlRequest() // CloudAgentsControlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerMessage(context.Background(), id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerMessage`: CloudAgentsControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsControlRequest** | [**CloudAgentsControlRequest**](CloudAgentsControlRequest.md) |  | 

### Return type

[**CloudAgentsControlResult**](CloudAgentsControlResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerPatch

> CloudAgentsSession CloudAgentSessionsControllerPatch(ctx, id).CloudAgentsPatchSessionRequest(cloudAgentsPatchSessionRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsPatchSessionRequest := *openapiclient.NewCloudAgentsPatchSessionRequest() // CloudAgentsPatchSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerPatch(context.Background(), id).CloudAgentsPatchSessionRequest(cloudAgentsPatchSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerPatch`: CloudAgentsSession
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsPatchSessionRequest** | [**CloudAgentsPatchSessionRequest**](CloudAgentsPatchSessionRequest.md) |  | 

### Return type

[**CloudAgentsSession**](CloudAgentsSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerPause

> CloudAgentsControlResult CloudAgentSessionsControllerPause(ctx, id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsControlRequest := *openapiclient.NewCloudAgentsControlRequest() // CloudAgentsControlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerPause(context.Background(), id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerPause`: CloudAgentsControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsControlRequest** | [**CloudAgentsControlRequest**](CloudAgentsControlRequest.md) |  | 

### Return type

[**CloudAgentsControlResult**](CloudAgentsControlResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerRegister

> CloudAgentsSession CloudAgentSessionsControllerRegister(ctx).CloudAgentsRegisterSessionRequest(cloudAgentsRegisterSessionRequest).Execute()





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
	cloudAgentsRegisterSessionRequest := *openapiclient.NewCloudAgentsRegisterSessionRequest("Agent_example") // CloudAgentsRegisterSessionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerRegister(context.Background()).CloudAgentsRegisterSessionRequest(cloudAgentsRegisterSessionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerRegister`: CloudAgentsSession
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAgentsRegisterSessionRequest** | [**CloudAgentsRegisterSessionRequest**](CloudAgentsRegisterSessionRequest.md) |  | 

### Return type

[**CloudAgentsSession**](CloudAgentsSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerResume

> CloudAgentsControlResult CloudAgentSessionsControllerResume(ctx, id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsControlRequest := *openapiclient.NewCloudAgentsControlRequest() // CloudAgentsControlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerResume(context.Background(), id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerResume`: CloudAgentsControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsControlRequest** | [**CloudAgentsControlRequest**](CloudAgentsControlRequest.md) |  | 

### Return type

[**CloudAgentsControlResult**](CloudAgentsControlResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerStop

> CloudAgentsControlResult CloudAgentSessionsControllerStop(ctx, id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()





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
	id := "id_example" // string | The session id (sess_...).
	cloudAgentsControlRequest := *openapiclient.NewCloudAgentsControlRequest() // CloudAgentsControlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerStop(context.Background(), id).CloudAgentsControlRequest(cloudAgentsControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerStop`: CloudAgentsControlResult
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudAgentsControlRequest** | [**CloudAgentsControlRequest**](CloudAgentsControlRequest.md) |  | 

### Return type

[**CloudAgentsControlResult**](CloudAgentsControlResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerStream

> string CloudAgentSessionsControllerStream(ctx).Root(root).Execute()





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
	root := "root_example" // string | Scope the feed to one subagent tree (rootSessionId). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerStream(context.Background()).Root(root).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerStream`: string
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **root** | **string** | Scope the feed to one subagent tree (rootSessionId). | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudAgentSessionsControllerTree

> CloudAgentsTreeNode CloudAgentSessionsControllerTree(ctx, id).Execute()





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
	id := "id_example" // string | The session id (sess_...).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AgentSessionsAPIAPI.CloudAgentSessionsControllerTree(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AgentSessionsAPIAPI.CloudAgentSessionsControllerTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudAgentSessionsControllerTree`: CloudAgentsTreeNode
	fmt.Fprintf(os.Stdout, "Response from `AgentSessionsAPIAPI.CloudAgentSessionsControllerTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The session id (sess_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudAgentSessionsControllerTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAgentsTreeNode**](CloudAgentsTreeNode.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

