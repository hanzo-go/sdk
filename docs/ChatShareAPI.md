# \ChatShareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteShareByshareid**](ChatShareAPI.md#ChatDeleteShareByshareid) | **Delete** /v1/chat/share/{shareId} | Delete a shared link
[**ChatGetShare**](ChatShareAPI.md#ChatGetShare) | **Get** /v1/chat/share | List shared links
[**ChatGetShareByshareid**](ChatShareAPI.md#ChatGetShareByshareid) | **Get** /v1/chat/share/{shareId} | Get shared conversation messages
[**ChatGetShareLinkByconversationid**](ChatShareAPI.md#ChatGetShareLinkByconversationid) | **Get** /v1/chat/share/link/{conversationId} | Get shared link for a conversation
[**ChatPatchShareByshareid**](ChatShareAPI.md#ChatPatchShareByshareid) | **Patch** /v1/chat/share/{shareId} | Update a shared link
[**ChatPostShareByconversationid**](ChatShareAPI.md#ChatPostShareByconversationid) | **Post** /v1/chat/share/{conversationId} | Create a shared link



## ChatDeleteShareByshareid

> map[string]interface{} ChatDeleteShareByshareid(ctx, shareId).Execute()

Delete a shared link

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatShareAPI.ChatDeleteShareByshareid(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatDeleteShareByshareid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteShareByshareid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatDeleteShareByshareid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteShareByshareidRequest struct via the builder pattern


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


## ChatGetShare

> ChatGetShare200Response ChatGetShare(ctx).Cursor(cursor).PageSize(pageSize).IsPublic(isPublic).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

List shared links

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
	cursor := "cursor_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	isPublic := "isPublic_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional) (default to "createdAt")
	sortDirection := "sortDirection_example" // string |  (optional) (default to "desc")
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatShareAPI.ChatGetShare(context.Background()).Cursor(cursor).PageSize(pageSize).IsPublic(isPublic).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatGetShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetShare`: ChatGetShare200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatGetShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatGetShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **pageSize** | **int32** |  | [default to 10]
 **isPublic** | **string** |  | 
 **sortBy** | **string** |  | [default to &quot;createdAt&quot;]
 **sortDirection** | **string** |  | [default to &quot;desc&quot;]
 **search** | **string** |  | 

### Return type

[**ChatGetShare200Response**](ChatGetShare200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetShareByshareid

> map[string]interface{} ChatGetShareByshareid(ctx, shareId).Execute()

Get shared conversation messages

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatShareAPI.ChatGetShareByshareid(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatGetShareByshareid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetShareByshareid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatGetShareByshareid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetShareByshareidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetShareLinkByconversationid

> ChatGetShareLinkByconversationid200Response ChatGetShareLinkByconversationid(ctx, conversationId).Execute()

Get shared link for a conversation

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
	resp, r, err := apiClient.ChatShareAPI.ChatGetShareLinkByconversationid(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatGetShareLinkByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetShareLinkByconversationid`: ChatGetShareLinkByconversationid200Response
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatGetShareLinkByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetShareLinkByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChatGetShareLinkByconversationid200Response**](ChatGetShareLinkByconversationid200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatPatchShareByshareid

> map[string]interface{} ChatPatchShareByshareid(ctx, shareId).Execute()

Update a shared link

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
	shareId := "shareId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatShareAPI.ChatPatchShareByshareid(context.Background(), shareId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatPatchShareByshareid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPatchShareByshareid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatPatchShareByshareid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPatchShareByshareidRequest struct via the builder pattern


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


## ChatPostShareByconversationid

> map[string]interface{} ChatPostShareByconversationid(ctx, conversationId).ChatPostShareByconversationidRequest(chatPostShareByconversationidRequest).Execute()

Create a shared link

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
	chatPostShareByconversationidRequest := *openapiclient.NewChatPostShareByconversationidRequest() // ChatPostShareByconversationidRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatShareAPI.ChatPostShareByconversationid(context.Background(), conversationId).ChatPostShareByconversationidRequest(chatPostShareByconversationidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatShareAPI.ChatPostShareByconversationid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostShareByconversationid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatShareAPI.ChatPostShareByconversationid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatPostShareByconversationidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chatPostShareByconversationidRequest** | [**ChatPostShareByconversationidRequest**](ChatPostShareByconversationidRequest.md) |  | 

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

