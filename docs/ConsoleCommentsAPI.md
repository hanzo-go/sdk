# \ConsoleCommentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateComment**](ConsoleCommentsAPI.md#ConsoleCreateComment) | **Post** /v1/console/comments | Create a comment
[**ConsoleGetComment**](ConsoleCommentsAPI.md#ConsoleGetComment) | **Get** /v1/console/comments/{commentId} | Get a comment by ID
[**ConsoleListComments**](ConsoleCommentsAPI.md#ConsoleListComments) | **Get** /v1/console/comments | Get all comments



## ConsoleCreateComment

> ConsoleCreateComment200Response ConsoleCreateComment(ctx).ConsoleCreateCommentRequest(consoleCreateCommentRequest).Execute()

Create a comment

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
	consoleCreateCommentRequest := *openapiclient.NewConsoleCreateCommentRequest("ProjectId_example", "ObjectType_example", "ObjectId_example", "Content_example") // ConsoleCreateCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleCommentsAPI.ConsoleCreateComment(context.Background()).ConsoleCreateCommentRequest(consoleCreateCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleCommentsAPI.ConsoleCreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateComment`: ConsoleCreateComment200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleCommentsAPI.ConsoleCreateComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateCommentRequest** | [**ConsoleCreateCommentRequest**](ConsoleCreateCommentRequest.md) |  | 

### Return type

[**ConsoleCreateComment200Response**](ConsoleCreateComment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetComment

> ConsoleComment ConsoleGetComment(ctx, commentId).Execute()

Get a comment by ID

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
	commentId := "commentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleCommentsAPI.ConsoleGetComment(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleCommentsAPI.ConsoleGetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetComment`: ConsoleComment
	fmt.Fprintf(os.Stdout, "Response from `ConsoleCommentsAPI.ConsoleGetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleComment**](ConsoleComment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListComments

> ConsoleListComments200Response ConsoleListComments(ctx).Page(page).Limit(limit).ObjectType(objectType).ObjectId(objectId).AuthorUserId(authorUserId).Execute()

Get all comments

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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)
	objectType := "objectType_example" // string | Filter by object type (trace, observation, session, prompt) (optional)
	objectId := "objectId_example" // string |  (optional)
	authorUserId := "authorUserId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleCommentsAPI.ConsoleListComments(context.Background()).Page(page).Limit(limit).ObjectType(objectType).ObjectId(objectId).AuthorUserId(authorUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleCommentsAPI.ConsoleListComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListComments`: ConsoleListComments200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleCommentsAPI.ConsoleListComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 
 **objectType** | **string** | Filter by object type (trace, observation, session, prompt) | 
 **objectId** | **string** |  | 
 **authorUserId** | **string** |  | 

### Return type

[**ConsoleListComments200Response**](ConsoleListComments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

