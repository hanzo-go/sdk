# \ChatTagsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteTagsBytag**](ChatTagsAPI.md#ChatDeleteTagsBytag) | **Delete** /v1/chat/tags/{tag} | Delete a conversation tag
[**ChatGetTags**](ChatTagsAPI.md#ChatGetTags) | **Get** /v1/chat/tags | Get all conversation tags
[**ChatPostTags**](ChatTagsAPI.md#ChatPostTags) | **Post** /v1/chat/tags | Create a conversation tag
[**ChatPutTagsBytag**](ChatTagsAPI.md#ChatPutTagsBytag) | **Put** /v1/chat/tags/{tag} | Update a conversation tag
[**ChatPutTagsConvoByconversationid**](ChatTagsAPI.md#ChatPutTagsConvoByconversationid) | **Put** /v1/chat/tags/convo/{conversationId} | Update tags for a conversation



## ChatDeleteTagsBytag

> map[string]interface{} ChatDeleteTagsBytag(ctx, tag).Execute()

Delete a conversation tag

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
	tag := "tag_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatTagsAPI.ChatDeleteTagsBytag(context.Background(), tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatTagsAPI.ChatDeleteTagsBytag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteTagsBytag`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatTagsAPI.ChatDeleteTagsBytag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteTagsBytagRequest struct via the builder pattern


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


## ChatGetTags

> []ChatConversationTag ChatGetTags(ctx).Execute()

Get all conversation tags

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
	resp, r, err := apiClient.ChatTagsAPI.ChatGetTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatTagsAPI.ChatGetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetTags`: []ChatConversationTag
	fmt.Fprintf(os.Stdout, "Response from `ChatTagsAPI.ChatGetTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetTagsRequest struct via the builder pattern


### Return type

[**[]ChatConversationTag**](ChatConversationTag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPostTags

> map[string]interface{} ChatPostTags(ctx).ChatConversationTag(chatConversationTag).Execute()

Create a conversation tag

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
	chatConversationTag := *openapiclient.NewChatConversationTag() // ChatConversationTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatTagsAPI.ChatPostTags(context.Background()).ChatConversationTag(chatConversationTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatTagsAPI.ChatPostTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatTagsAPI.ChatPostTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatConversationTag** | [**ChatConversationTag**](ChatConversationTag.md) |  | 

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


## ChatPutTagsBytag

> map[string]interface{} ChatPutTagsBytag(ctx, tag).ChatConversationTag(chatConversationTag).Execute()

Update a conversation tag

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
	tag := "tag_example" // string | 
	chatConversationTag := *openapiclient.NewChatConversationTag() // ChatConversationTag | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatTagsAPI.ChatPutTagsBytag(context.Background(), tag).ChatConversationTag(chatConversationTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatTagsAPI.ChatPutTagsBytag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutTagsBytag`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatTagsAPI.ChatPutTagsBytag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutTagsBytagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatConversationTag** | [**ChatConversationTag**](ChatConversationTag.md) |  | 

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


## ChatPutTagsConvoByconversationid

> map[string]interface{} ChatPutTagsConvoByconversationid(ctx, conversationId).ChatPutTagsConvoByconversationidRequest(chatPutTagsConvoByconversationidRequest).Execute()

Update tags for a conversation

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
	chatPutTagsConvoByconversationidRequest := *openapiclient.NewChatPutTagsConvoByconversationidRequest() // ChatPutTagsConvoByconversationidRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatTagsAPI.ChatPutTagsConvoByconversationid(context.Background(), conversationId).ChatPutTagsConvoByconversationidRequest(chatPutTagsConvoByconversationidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatTagsAPI.ChatPutTagsConvoByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPutTagsConvoByconversationid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatTagsAPI.ChatPutTagsConvoByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPutTagsConvoByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPutTagsConvoByconversationidRequest** | [**ChatPutTagsConvoByconversationidRequest**](ChatPutTagsConvoByconversationidRequest.md) |  | 

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

