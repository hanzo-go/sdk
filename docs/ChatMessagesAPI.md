# \ChatMessagesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteMessagesByconversationidBymessageid**](ChatMessagesAPI.md#ChatDeleteMessagesByconversationidBymessageid) | **Delete** /v1/chat/messages/{conversationId}/{messageId} | Delete a message
[**ChatGetMessages**](ChatMessagesAPI.md#ChatGetMessages) | **Get** /v1/chat/messages | Query messages
[**ChatGetMessagesByconversationid**](ChatMessagesAPI.md#ChatGetMessagesByconversationid) | **Get** /v1/chat/messages/{conversationId} | Get all messages in a conversation
[**ChatGetMessagesByconversationidBymessageid**](ChatMessagesAPI.md#ChatGetMessagesByconversationidBymessageid) | **Get** /v1/chat/messages/{conversationId}/{messageId} | Get a specific message
[**ChatPostMessagesArtifactBymessageid**](ChatMessagesAPI.md#ChatPostMessagesArtifactBymessageid) | **Post** /v1/chat/messages/artifact/{messageId} | Edit artifact content in a message
[**ChatPostMessagesBranch**](ChatMessagesAPI.md#ChatPostMessagesBranch) | **Post** /v1/chat/messages/branch | Create a branch message
[**ChatPostMessagesByconversationid**](ChatMessagesAPI.md#ChatPostMessagesByconversationid) | **Post** /v1/chat/messages/{conversationId} | Save a message to a conversation
[**ChatPutMessagesByconversationidBymessageid**](ChatMessagesAPI.md#ChatPutMessagesByconversationidBymessageid) | **Put** /v1/chat/messages/{conversationId}/{messageId} | Update a message
[**ChatPutMessagesByconversationidBymessageidFeedback**](ChatMessagesAPI.md#ChatPutMessagesByconversationidBymessageidFeedback) | **Put** /v1/chat/messages/{conversationId}/{messageId}/feedback | Update message feedback



## ChatDeleteMessagesByconversationidBymessageid

> ChatDeleteMessagesByconversationidBymessageid(ctx, conversationId, messageId).Execute()

Delete a message

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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChatMessagesAPI.ChatDeleteMessagesByconversationidBymessageid(context.Background(), conversationId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatDeleteMessagesByconversationidBymessageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteMessagesByconversationidBymessageidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMessages

> ChatGetMessages200Response ChatGetMessages(ctx).ConversationId(conversationId).MessageId(messageId).Search(search).Cursor(cursor).PageSize(pageSize).SortBy(sortBy).SortDirection(sortDirection).Execute()

Query messages



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
	conversationId := "conversationId_example" // string |  (optional)
	messageId := "messageId_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 25)
	sortBy := "sortBy_example" // string |  (optional) (default to "createdAt")
	sortDirection := "sortDirection_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatGetMessages(context.Background()).ConversationId(conversationId).MessageId(messageId).Search(search).Cursor(cursor).PageSize(pageSize).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatGetMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMessages`: ChatGetMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatGetMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversationId** | **string** |  | 
 **messageId** | **string** |  | 
 **search** | **string** |  | 
 **cursor** | **string** |  | 
 **pageSize** | **int32** |  | [default to 25]
 **sortBy** | **string** |  | [default to &quot;createdAt&quot;]
 **sortDirection** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**ChatGetMessages200Response**](ChatGetMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMessagesByconversationid

> []ChatMessage ChatGetMessagesByconversationid(ctx, conversationId).Execute()

Get all messages in a conversation

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
	resp, r, err := apiClient.ChatMessagesAPI.ChatGetMessagesByconversationid(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatGetMessagesByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMessagesByconversationid`: []ChatMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatGetMessagesByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMessagesByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ChatMessage**](ChatMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetMessagesByconversationidBymessageid

> ChatMessage ChatGetMessagesByconversationidBymessageid(ctx, conversationId, messageId).Execute()

Get a specific message

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
	messageId := "messageId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatGetMessagesByconversationidBymessageid(context.Background(), conversationId, messageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatGetMessagesByconversationidBymessageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetMessagesByconversationidBymessageid`: ChatMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatGetMessagesByconversationidBymessageid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetMessagesByconversationidBymessageidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChatMessage**](ChatMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMessagesArtifactBymessageid

> ChatPostMessagesArtifactBymessageid200Response ChatPostMessagesArtifactBymessageid(ctx, messageId).ChatPostMessagesArtifactBymessageidRequest(chatPostMessagesArtifactBymessageidRequest).Execute()

Edit artifact content in a message

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
	messageId := "messageId_example" // string | 
	chatPostMessagesArtifactBymessageidRequest := *openapiclient.NewChatPostMessagesArtifactBymessageidRequest(int32(123), "Original_example", "Updated_example") // ChatPostMessagesArtifactBymessageidRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatPostMessagesArtifactBymessageid(context.Background(), messageId).ChatPostMessagesArtifactBymessageidRequest(chatPostMessagesArtifactBymessageidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatPostMessagesArtifactBymessageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMessagesArtifactBymessageid`: ChatPostMessagesArtifactBymessageid200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatPostMessagesArtifactBymessageid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMessagesArtifactBymessageidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostMessagesArtifactBymessageidRequest** | [**ChatPostMessagesArtifactBymessageidRequest**](ChatPostMessagesArtifactBymessageidRequest.md) |  | 

### Return type

[**ChatPostMessagesArtifactBymessageid200Response**](ChatPostMessagesArtifactBymessageid200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMessagesBranch

> ChatMessage ChatPostMessagesBranch(ctx).ChatPostMessagesBranchRequest(chatPostMessagesBranchRequest).Execute()

Create a branch message



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
	chatPostMessagesBranchRequest := *openapiclient.NewChatPostMessagesBranchRequest("MessageId_example", "AgentId_example") // ChatPostMessagesBranchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatPostMessagesBranch(context.Background()).ChatPostMessagesBranchRequest(chatPostMessagesBranchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatPostMessagesBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMessagesBranch`: ChatMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatPostMessagesBranch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMessagesBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostMessagesBranchRequest** | [**ChatPostMessagesBranchRequest**](ChatPostMessagesBranchRequest.md) |  | 

### Return type

[**ChatMessage**](ChatMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostMessagesByconversationid

> ChatMessage ChatPostMessagesByconversationid(ctx, conversationId).ChatMessage(chatMessage).Execute()

Save a message to a conversation

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
	chatMessage := *openapiclient.NewChatMessage() // ChatMessage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatPostMessagesByconversationid(context.Background(), conversationId).ChatMessage(chatMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatPostMessagesByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostMessagesByconversationid`: ChatMessage
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatPostMessagesByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostMessagesByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatMessage** | [**ChatMessage**](ChatMessage.md) |  | 

### Return type

[**ChatMessage**](ChatMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPutMessagesByconversationidBymessageid

> map[string]interface{} ChatPutMessagesByconversationidBymessageid(ctx, conversationId, messageId).ChatPutMessagesByconversationidBymessageidRequest(chatPutMessagesByconversationidBymessageidRequest).Execute()

Update a message

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
	messageId := "messageId_example" // string | 
	chatPutMessagesByconversationidBymessageidRequest := *openapiclient.NewChatPutMessagesByconversationidBymessageidRequest() // ChatPutMessagesByconversationidBymessageidRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatPutMessagesByconversationidBymessageid(context.Background(), conversationId, messageId).ChatPutMessagesByconversationidBymessageidRequest(chatPutMessagesByconversationidBymessageidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatPutMessagesByconversationidBymessageid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutMessagesByconversationidBymessageid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatPutMessagesByconversationidBymessageid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutMessagesByconversationidBymessageidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chatPutMessagesByconversationidBymessageidRequest** | [**ChatPutMessagesByconversationidBymessageidRequest**](ChatPutMessagesByconversationidBymessageidRequest.md) |  | 

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


## ChatPutMessagesByconversationidBymessageidFeedback

> ChatPutMessagesByconversationidBymessageidFeedback200Response ChatPutMessagesByconversationidBymessageidFeedback(ctx, conversationId, messageId).ChatPutMessagesByconversationidBymessageidFeedbackRequest(chatPutMessagesByconversationidBymessageidFeedbackRequest).Execute()

Update message feedback

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
	messageId := "messageId_example" // string | 
	chatPutMessagesByconversationidBymessageidFeedbackRequest := *openapiclient.NewChatPutMessagesByconversationidBymessageidFeedbackRequest() // ChatPutMessagesByconversationidBymessageidFeedbackRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatMessagesAPI.ChatPutMessagesByconversationidBymessageidFeedback(context.Background(), conversationId, messageId).ChatPutMessagesByconversationidBymessageidFeedbackRequest(chatPutMessagesByconversationidBymessageidFeedbackRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatMessagesAPI.ChatPutMessagesByconversationidBymessageidFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutMessagesByconversationidBymessageidFeedback`: ChatPutMessagesByconversationidBymessageidFeedback200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatMessagesAPI.ChatPutMessagesByconversationidBymessageidFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutMessagesByconversationidBymessageidFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chatPutMessagesByconversationidBymessageidFeedbackRequest** | [**ChatPutMessagesByconversationidBymessageidFeedbackRequest**](ChatPutMessagesByconversationidBymessageidFeedbackRequest.md) |  | 

### Return type

[**ChatPutMessagesByconversationidBymessageidFeedback200Response**](ChatPutMessagesByconversationidBymessageidFeedback200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

