# \ContentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetContentBoard**](ContentAPI.md#GetContentBoard) | **Get** /v1/content/board | Aggregates the caller org&#39;s marketing content across every publishable content type into ONE queue board — the cross-type read the framework&#39;s per-DocType list cannot give.
[**GetContentChannels**](ContentAPI.md#GetContentChannels) | **Get** /v1/content/channels | Lists the distribution channels the caller&#39;s org has connected — the social integrations a publish can target.
[**GetContentLifecycle**](ContentAPI.md#GetContentLifecycle) | **Get** /v1/content/lifecycle | Returns the ONE marketing-content state machine: the ordered lifecycle states, which state a fresh document starts in, which one is publicly live, and the legal successors of every state.
[**PostContentByDoctypeByNameTransition**](ContentAPI.md#PostContentByDoctypeByNameTransition) | **Post** /v1/content/{doctype}/{name}/transition | Moves one content item to a new lifecycle state and, on the move to published, fans it out to the item&#39;s channels.
[**PostContentGenerate**](ContentAPI.md#PostContentGenerate) | **Post** /v1/content/generate | Draft a piece of marketing content and file it in the CMS as a draft.
[**PostContentPublish**](ContentAPI.md#PostContentPublish) | **Post** /v1/content/publish | Publish distributes one CMS content item to the channels recorded on it and returns the honest per-channel outcome.



## GetContentBoard

> BoardPage GetContentBoard(ctx).Status(status).Project(project).Doctype(doctype).Limit(limit).Execute()

Aggregates the caller org's marketing content across every publishable content type into ONE queue board — the cross-type read the framework's per-DocType list cannot give.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	status := "queued" // string | Status keeps only items in one lifecycle state (draft, in_review, approved, queued, published, archived). An undefined state is refused. (optional)
	project := "project_example" // string | Project keeps only items in one brand/site sub-scope. (optional)
	doctype := "doctype_example" // string | DocType keeps only one content type; omitted, the board spans every publishable type. An unknown type is refused. (optional)
	limit := int64(50) // int64 | Limit caps the rows returned, clamped to 1000. Defaults to 200, which is also what a non-positive or unparseable value takes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentBoard(context.Background()).Status(status).Project(project).Doctype(doctype).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentBoard`: BoardPage
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContentBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only items in one lifecycle state (draft, in_review, approved, queued, published, archived). An undefined state is refused. | 
 **project** | **string** | Project keeps only items in one brand/site sub-scope. | 
 **doctype** | **string** | DocType keeps only one content type; omitted, the board spans every publishable type. An unknown type is refused. | 
 **limit** | **int64** | Limit caps the rows returned, clamped to 1000. Defaults to 200, which is also what a non-positive or unparseable value takes. | 

### Return type

[**BoardPage**](BoardPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentChannels

> ChannelList GetContentChannels(ctx).Execute()

Lists the distribution channels the caller's org has connected — the social integrations a publish can target.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentChannels`: ChannelList
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentChannelsRequest struct via the builder pattern


### Return type

[**ChannelList**](ChannelList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentLifecycle

> StateGraph GetContentLifecycle(ctx).Execute()

Returns the ONE marketing-content state machine: the ordered lifecycle states, which state a fresh document starts in, which one is publicly live, and the legal successors of every state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetContentLifecycle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetContentLifecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentLifecycle`: StateGraph
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetContentLifecycle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentLifecycleRequest struct via the builder pattern


### Return type

[**StateGraph**](StateGraph.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContentByDoctypeByNameTransition

> TransitionResult PostContentByDoctypeByNameTransition(ctx, doctype, name).TransitionIn(transitionIn).Execute()

Moves one content item to a new lifecycle state and, on the move to published, fans it out to the item's channels.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	doctype := "marketing.SocialPost" // string | DocType is the content type to act on, from the path.
	name := "spring-teaser" // string | Name is the document to act on, from the path.
	transitionIn := *openapiclient.NewTransitionIn() // TransitionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PostContentByDoctypeByNameTransition(context.Background(), doctype, name).TransitionIn(transitionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PostContentByDoctypeByNameTransition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContentByDoctypeByNameTransition`: TransitionResult
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PostContentByDoctypeByNameTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the content type to act on, from the path. | 
**name** | **string** | Name is the document to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostContentByDoctypeByNameTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transitionIn** | [**TransitionIn**](TransitionIn.md) |  | 

### Return type

[**TransitionResult**](TransitionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContentGenerate

> GenerateResult PostContentGenerate(ctx).GenerateInput(generateInput).Execute()

Draft a piece of marketing content and file it in the CMS as a draft.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	generateInput := *openapiclient.NewGenerateInput() // GenerateInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PostContentGenerate(context.Background()).GenerateInput(generateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PostContentGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContentGenerate`: GenerateResult
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PostContentGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContentGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **generateInput** | [**GenerateInput**](GenerateInput.md) |  | 

### Return type

[**GenerateResult**](GenerateResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostContentPublish

> PublishResult PostContentPublish(ctx).PublishInput(publishInput).Execute()

Publish distributes one CMS content item to the channels recorded on it and returns the honest per-channel outcome.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	publishInput := *openapiclient.NewPublishInput() // PublishInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PostContentPublish(context.Background()).PublishInput(publishInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PostContentPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostContentPublish`: PublishResult
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PostContentPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostContentPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publishInput** | [**PublishInput**](PublishInput.md) |  | 

### Return type

[**PublishResult**](PublishResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

