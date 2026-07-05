# \ChatConversationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteConvos**](ChatConversationsAPI.md#ChatDeleteConvos) | **Delete** /v1/chat/convos | Delete a conversation
[**ChatDeleteConvosAll**](ChatConversationsAPI.md#ChatDeleteConvosAll) | **Delete** /v1/chat/convos/all | Delete all conversations
[**ChatGetConvos**](ChatConversationsAPI.md#ChatGetConvos) | **Get** /v1/chat/convos | List conversations
[**ChatGetConvosByconversationid**](ChatConversationsAPI.md#ChatGetConvosByconversationid) | **Get** /v1/chat/convos/{conversationId} | Get a conversation
[**ChatGetConvosGenTitleByconversationid**](ChatConversationsAPI.md#ChatGetConvosGenTitleByconversationid) | **Get** /v1/chat/convos/gen_title/{conversationId} | Get generated title for conversation
[**ChatPostConvosArchive**](ChatConversationsAPI.md#ChatPostConvosArchive) | **Post** /v1/chat/convos/archive | Archive or unarchive a conversation
[**ChatPostConvosDuplicate**](ChatConversationsAPI.md#ChatPostConvosDuplicate) | **Post** /v1/chat/convos/duplicate | Duplicate a conversation
[**ChatPostConvosFork**](ChatConversationsAPI.md#ChatPostConvosFork) | **Post** /v1/chat/convos/fork | Fork a conversation
[**ChatPostConvosImport**](ChatConversationsAPI.md#ChatPostConvosImport) | **Post** /v1/chat/convos/import | Import conversations from JSON file
[**ChatPostConvosUpdate**](ChatConversationsAPI.md#ChatPostConvosUpdate) | **Post** /v1/chat/convos/update | Update a conversation title



## ChatDeleteConvos

> map[string]interface{} ChatDeleteConvos(ctx).ChatDeleteConvosRequest(chatDeleteConvosRequest).Execute()

Delete a conversation

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
	chatDeleteConvosRequest := *openapiclient.NewChatDeleteConvosRequest() // ChatDeleteConvosRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatDeleteConvos(context.Background()).ChatDeleteConvosRequest(chatDeleteConvosRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatDeleteConvos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteConvos`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatDeleteConvos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteConvosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatDeleteConvosRequest** | [**ChatDeleteConvosRequest**](ChatDeleteConvosRequest.md) |  | 

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


## ChatDeleteConvosAll

> map[string]interface{} ChatDeleteConvosAll(ctx).Execute()

Delete all conversations

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
	resp, r, err := apiClient.ChatConversationsAPI.ChatDeleteConvosAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatDeleteConvosAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteConvosAll`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatDeleteConvosAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteConvosAllRequest struct via the builder pattern


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


## ChatGetConvos

> ChatConversationListResponse ChatGetConvos(ctx).Limit(limit).Cursor(cursor).IsArchived(isArchived).Tags(tags).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()

List conversations

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
	limit := int32(56) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string |  (optional)
	isArchived := "isArchived_example" // string |  (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	search := "search_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional) (default to "updatedAt")
	sortDirection := "sortDirection_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatGetConvos(context.Background()).Limit(limit).Cursor(cursor).IsArchived(isArchived).Tags(tags).Search(search).SortBy(sortBy).SortDirection(sortDirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatGetConvos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetConvos`: ChatConversationListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatGetConvos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetConvosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 25]
 **cursor** | **string** |  | 
 **isArchived** | **string** |  | 
 **tags** | **[]string** |  | 
 **search** | **string** |  | 
 **sortBy** | **string** |  | [default to &quot;updatedAt&quot;]
 **sortDirection** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**ChatConversationListResponse**](ChatConversationListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetConvosByconversationid

> ChatConversation ChatGetConvosByconversationid(ctx, conversationId).Execute()

Get a conversation

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
	resp, r, err := apiClient.ChatConversationsAPI.ChatGetConvosByconversationid(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatGetConvosByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetConvosByconversationid`: ChatConversation
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatGetConvosByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetConvosByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatConversation**](ChatConversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetConvosGenTitleByconversationid

> ChatGetConvosGenTitleByconversationid200Response ChatGetConvosGenTitleByconversationid(ctx, conversationId).Execute()

Get generated title for conversation



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
	resp, r, err := apiClient.ChatConversationsAPI.ChatGetConvosGenTitleByconversationid(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatGetConvosGenTitleByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetConvosGenTitleByconversationid`: ChatGetConvosGenTitleByconversationid200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatGetConvosGenTitleByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetConvosGenTitleByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetConvosGenTitleByconversationid200Response**](ChatGetConvosGenTitleByconversationid200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostConvosArchive

> ChatConversation ChatPostConvosArchive(ctx).ChatPostConvosArchiveRequest(chatPostConvosArchiveRequest).Execute()

Archive or unarchive a conversation

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
	chatPostConvosArchiveRequest := *openapiclient.NewChatPostConvosArchiveRequest() // ChatPostConvosArchiveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatPostConvosArchive(context.Background()).ChatPostConvosArchiveRequest(chatPostConvosArchiveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatPostConvosArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostConvosArchive`: ChatConversation
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatPostConvosArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostConvosArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostConvosArchiveRequest** | [**ChatPostConvosArchiveRequest**](ChatPostConvosArchiveRequest.md) |  | 

### Return type

[**ChatConversation**](ChatConversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostConvosDuplicate

> map[string]interface{} ChatPostConvosDuplicate(ctx).ChatPostConvosDuplicateRequest(chatPostConvosDuplicateRequest).Execute()

Duplicate a conversation

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
	chatPostConvosDuplicateRequest := *openapiclient.NewChatPostConvosDuplicateRequest("ConversationId_example") // ChatPostConvosDuplicateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatPostConvosDuplicate(context.Background()).ChatPostConvosDuplicateRequest(chatPostConvosDuplicateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatPostConvosDuplicate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostConvosDuplicate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatPostConvosDuplicate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostConvosDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostConvosDuplicateRequest** | [**ChatPostConvosDuplicateRequest**](ChatPostConvosDuplicateRequest.md) |  | 

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


## ChatPostConvosFork

> ChatPostConvosFork200Response ChatPostConvosFork(ctx).ChatPostConvosForkRequest(chatPostConvosForkRequest).Execute()

Fork a conversation

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
	chatPostConvosForkRequest := *openapiclient.NewChatPostConvosForkRequest("ConversationId_example", "MessageId_example") // ChatPostConvosForkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatPostConvosFork(context.Background()).ChatPostConvosForkRequest(chatPostConvosForkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatPostConvosFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostConvosFork`: ChatPostConvosFork200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatPostConvosFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostConvosForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostConvosForkRequest** | [**ChatPostConvosForkRequest**](ChatPostConvosForkRequest.md) |  | 

### Return type

[**ChatPostConvosFork200Response**](ChatPostConvosFork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostConvosImport

> map[string]interface{} ChatPostConvosImport(ctx).File(file).Execute()

Import conversations from JSON file

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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatPostConvosImport(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatPostConvosImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostConvosImport`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatPostConvosImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostConvosImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostConvosUpdate

> map[string]interface{} ChatPostConvosUpdate(ctx).ChatPostConvosUpdateRequest(chatPostConvosUpdateRequest).Execute()

Update a conversation title

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
	chatPostConvosUpdateRequest := *openapiclient.NewChatPostConvosUpdateRequest() // ChatPostConvosUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatConversationsAPI.ChatPostConvosUpdate(context.Background()).ChatPostConvosUpdateRequest(chatPostConvosUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatConversationsAPI.ChatPostConvosUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostConvosUpdate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatConversationsAPI.ChatPostConvosUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostConvosUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatPostConvosUpdateRequest** | [**ChatPostConvosUpdateRequest**](ChatPostConvosUpdateRequest.md) |  | 

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

