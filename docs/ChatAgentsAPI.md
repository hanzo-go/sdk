# \ChatAgentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteAgentsByid**](ChatAgentsAPI.md#ChatDeleteAgentsByid) | **Delete** /v1/chat/agents/{id} | Delete an agent
[**ChatGetAgents**](ChatAgentsAPI.md#ChatGetAgents) | **Get** /v1/chat/agents | List agents
[**ChatGetAgentsByid**](ChatAgentsAPI.md#ChatGetAgentsByid) | **Get** /v1/chat/agents/{id} | Get an agent (basic info)
[**ChatGetAgentsByidExpanded**](ChatAgentsAPI.md#ChatGetAgentsByidExpanded) | **Get** /v1/chat/agents/{id}/expanded | Get agent with full configuration details
[**ChatGetAgentsCategories**](ChatAgentsAPI.md#ChatGetAgentsCategories) | **Get** /v1/chat/agents/categories | Get agent categories with counts
[**ChatGetAgentsChatActive**](ChatAgentsAPI.md#ChatGetAgentsChatActive) | **Get** /v1/chat/agents/chat/active | Get active generation job IDs
[**ChatGetAgentsChatStatusByconversationid**](ChatAgentsAPI.md#ChatGetAgentsChatStatusByconversationid) | **Get** /v1/chat/agents/chat/status/{conversationId} | Check generation status for a conversation
[**ChatGetAgentsChatStreamBystreamid**](ChatAgentsAPI.md#ChatGetAgentsChatStreamBystreamid) | **Get** /v1/chat/agents/chat/stream/{streamId} | Subscribe to a generation stream
[**ChatGetAgentsTools**](ChatAgentsAPI.md#ChatGetAgentsTools) | **Get** /v1/chat/agents/tools | List available agent tools
[**ChatGetAgentsToolsBytoolidAuth**](ChatAgentsAPI.md#ChatGetAgentsToolsBytoolidAuth) | **Get** /v1/chat/agents/tools/{toolId}/auth | Verify tool authentication
[**ChatGetAgentsToolsCalls**](ChatAgentsAPI.md#ChatGetAgentsToolsCalls) | **Get** /v1/chat/agents/tools/calls | Get tool call history
[**ChatPatchAgentsByid**](ChatAgentsAPI.md#ChatPatchAgentsByid) | **Patch** /v1/chat/agents/{id} | Update an agent
[**ChatPostAgents**](ChatAgentsAPI.md#ChatPostAgents) | **Post** /v1/chat/agents | Create an agent
[**ChatPostAgentsByidDuplicate**](ChatAgentsAPI.md#ChatPostAgentsByidDuplicate) | **Post** /v1/chat/agents/{id}/duplicate | Duplicate an agent
[**ChatPostAgentsByidRevert**](ChatAgentsAPI.md#ChatPostAgentsByidRevert) | **Post** /v1/chat/agents/{id}/revert | Revert agent to a previous version
[**ChatPostAgentsChat**](ChatAgentsAPI.md#ChatPostAgentsChat) | **Post** /v1/chat/agents/chat | Chat with an agent
[**ChatPostAgentsChatAbort**](ChatAgentsAPI.md#ChatPostAgentsChatAbort) | **Post** /v1/chat/agents/chat/abort | Abort an ongoing agent generation
[**ChatPostAgentsChatByendpoint**](ChatAgentsAPI.md#ChatPostAgentsChatByendpoint) | **Post** /v1/chat/agents/chat/{endpoint} | Chat with an ephemeral agent
[**ChatPostAgentsToolsBytoolidCall**](ChatAgentsAPI.md#ChatPostAgentsToolsBytoolidCall) | **Post** /v1/chat/agents/tools/{toolId}/call | Execute a tool call



## ChatDeleteAgentsByid

> map[string]interface{} ChatDeleteAgentsByid(ctx, id).Execute()

Delete an agent

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatDeleteAgentsByid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatDeleteAgentsByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteAgentsByid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatDeleteAgentsByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteAgentsByidRequest struct via the builder pattern


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


## ChatGetAgents

> ChatAgentListResponse ChatGetAgents(ctx).Limit(limit).After(after).SortBy(sortBy).SortDirection(sortDirection).Execute()

List agents

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
	limit := int32(56) // int32 |  (optional)
	after := "after_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	sortDirection := "sortDirection_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgents(context.Background()).Limit(limit).After(after).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgents`: ChatAgentListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **after** | **string** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 

### Return type

[**ChatAgentListResponse**](ChatAgentListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsByid

> ChatAgent ChatGetAgentsByid(ctx, id).Execute()

Get an agent (basic info)

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsByid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsByid`: ChatAgent
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatAgent**](ChatAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsByidExpanded

> ChatAgent ChatGetAgentsByidExpanded(ctx, id).Execute()

Get agent with full configuration details



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsByidExpanded(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsByidExpanded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsByidExpanded`: ChatAgent
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsByidExpanded`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsByidExpandedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatAgent**](ChatAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsCategories

> map[string]interface{} ChatGetAgentsCategories(ctx).Execute()

Get agent categories with counts

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
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsCategoriesRequest struct via the builder pattern


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


## ChatGetAgentsChatActive

> ChatGetAgentsChatActive200Response ChatGetAgentsChatActive(ctx).Execute()

Get active generation job IDs

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
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsChatActive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsChatActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsChatActive`: ChatGetAgentsChatActive200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsChatActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsChatActiveRequest struct via the builder pattern


### Return type

[**ChatGetAgentsChatActive200Response**](ChatGetAgentsChatActive200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsChatStatusByconversationid

> ChatGetAgentsChatStatusByconversationid200Response ChatGetAgentsChatStatusByconversationid(ctx, conversationId).Execute()

Check generation status for a conversation

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
	conversationId := "conversationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsChatStatusByconversationid(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsChatStatusByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsChatStatusByconversationid`: ChatGetAgentsChatStatusByconversationid200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsChatStatusByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsChatStatusByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetAgentsChatStatusByconversationid200Response**](ChatGetAgentsChatStatusByconversationid200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsChatStreamBystreamid

> string ChatGetAgentsChatStreamBystreamid(ctx, streamId).Resume(resume).Execute()

Subscribe to a generation stream



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
	streamId := "streamId_example" // string | 
	resume := "resume_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsChatStreamBystreamid(context.Background(), streamId).Resume(resume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsChatStreamBystreamid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsChatStreamBystreamid`: string
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsChatStreamBystreamid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsChatStreamBystreamidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resume** | **string** |  | 

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


## ChatGetAgentsTools

> []ChatTool ChatGetAgentsTools(ctx).Execute()

List available agent tools

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
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsTools`: []ChatTool
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsToolsRequest struct via the builder pattern


### Return type

[**[]ChatTool**](ChatTool.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsToolsBytoolidAuth

> ChatGetAgentsToolsBytoolidAuth200Response ChatGetAgentsToolsBytoolidAuth(ctx, toolId).Execute()

Verify tool authentication

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
	toolId := "toolId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsToolsBytoolidAuth(context.Background(), toolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsToolsBytoolidAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsToolsBytoolidAuth`: ChatGetAgentsToolsBytoolidAuth200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsToolsBytoolidAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsToolsBytoolidAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetAgentsToolsBytoolidAuth200Response**](ChatGetAgentsToolsBytoolidAuth200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetAgentsToolsCalls

> map[string]interface{} ChatGetAgentsToolsCalls(ctx).Execute()

Get tool call history

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
	resp, r, err := apiClient.ChatAgentsAPI.ChatGetAgentsToolsCalls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatGetAgentsToolsCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetAgentsToolsCalls`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatGetAgentsToolsCalls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetAgentsToolsCallsRequest struct via the builder pattern


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


## ChatPatchAgentsByid

> ChatAgent ChatPatchAgentsByid(ctx, id).ChatAgentCreateParams(chatAgentCreateParams).Execute()

Update an agent

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
	id := "id_example" // string | 
	chatAgentCreateParams := *openapiclient.NewChatAgentCreateParams("Name_example") // ChatAgentCreateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPatchAgentsByid(context.Background(), id).ChatAgentCreateParams(chatAgentCreateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPatchAgentsByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchAgentsByid`: ChatAgent
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPatchAgentsByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchAgentsByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAgentCreateParams** | [**ChatAgentCreateParams**](ChatAgentCreateParams.md) |  | 

### Return type

[**ChatAgent**](ChatAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgents

> ChatAgent ChatPostAgents(ctx).ChatAgentCreateParams(chatAgentCreateParams).Execute()

Create an agent

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
	chatAgentCreateParams := *openapiclient.NewChatAgentCreateParams("Name_example") // ChatAgentCreateParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgents(context.Background()).ChatAgentCreateParams(chatAgentCreateParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgents`: ChatAgent
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatAgentCreateParams** | [**ChatAgentCreateParams**](ChatAgentCreateParams.md) |  | 

### Return type

[**ChatAgent**](ChatAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsByidDuplicate

> ChatAgent ChatPostAgentsByidDuplicate(ctx, id).Execute()

Duplicate an agent

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsByidDuplicate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsByidDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsByidDuplicate`: ChatAgent
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsByidDuplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsByidDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatAgent**](ChatAgent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsByidRevert

> map[string]interface{} ChatPostAgentsByidRevert(ctx, id).ChatPostAgentsByidRevertRequest(chatPostAgentsByidRevertRequest).Execute()

Revert agent to a previous version

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
	id := "id_example" // string | 
	chatPostAgentsByidRevertRequest := *openapiclient.NewChatPostAgentsByidRevertRequest(int32(123)) // ChatPostAgentsByidRevertRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsByidRevert(context.Background(), id).ChatPostAgentsByidRevertRequest(chatPostAgentsByidRevertRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsByidRevert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsByidRevert`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsByidRevert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsByidRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostAgentsByidRevertRequest** | [**ChatPostAgentsByidRevertRequest**](ChatPostAgentsByidRevertRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsChat

> string ChatPostAgentsChat(ctx).ChatAgentChatRequest(chatAgentChatRequest).Execute()

Chat with an agent



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
	chatAgentChatRequest := *openapiclient.NewChatAgentChatRequest() // ChatAgentChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsChat(context.Background()).ChatAgentChatRequest(chatAgentChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsChat`: string
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatAgentChatRequest** | [**ChatAgentChatRequest**](ChatAgentChatRequest.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsChatAbort

> ChatPostAgentsChatAbort200Response ChatPostAgentsChatAbort(ctx).ChatPostAgentsChatAbortRequest(chatPostAgentsChatAbortRequest).Execute()

Abort an ongoing agent generation

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
	chatPostAgentsChatAbortRequest := *openapiclient.NewChatPostAgentsChatAbortRequest() // ChatPostAgentsChatAbortRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsChatAbort(context.Background()).ChatPostAgentsChatAbortRequest(chatPostAgentsChatAbortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsChatAbort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsChatAbort`: ChatPostAgentsChatAbort200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsChatAbort`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsChatAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostAgentsChatAbortRequest** | [**ChatPostAgentsChatAbortRequest**](ChatPostAgentsChatAbortRequest.md) |  | 

### Return type

[**ChatPostAgentsChatAbort200Response**](ChatPostAgentsChatAbort200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsChatByendpoint

> string ChatPostAgentsChatByendpoint(ctx, endpoint).ChatAgentChatRequest(chatAgentChatRequest).Execute()

Chat with an ephemeral agent

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
	endpoint := "endpoint_example" // string | 
	chatAgentChatRequest := *openapiclient.NewChatAgentChatRequest() // ChatAgentChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsChatByendpoint(context.Background(), endpoint).ChatAgentChatRequest(chatAgentChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsChatByendpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsChatByendpoint`: string
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsChatByendpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpoint** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsChatByendpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatAgentChatRequest** | [**ChatAgentChatRequest**](ChatAgentChatRequest.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostAgentsToolsBytoolidCall

> map[string]interface{} ChatPostAgentsToolsBytoolidCall(ctx, toolId).Body(body).Execute()

Execute a tool call

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
	toolId := "toolId_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAgentsAPI.ChatPostAgentsToolsBytoolidCall(context.Background(), toolId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAgentsAPI.ChatPostAgentsToolsBytoolidCall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostAgentsToolsBytoolidCall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAgentsAPI.ChatPostAgentsToolsBytoolidCall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**toolId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostAgentsToolsBytoolidCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

