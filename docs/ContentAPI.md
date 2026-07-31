# \ContentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1ContentBoard**](ContentAPI.md#CloudGetV1ContentBoard) | **Get** /v1/content/board | GetBoard aggregates the caller org&#39;s marketing content across every publishable content type into ONE queue board — the cross-type read the framework&#39;s per-DocType list cannot give.
[**CloudGetV1ContentChannels**](ContentAPI.md#CloudGetV1ContentChannels) | **Get** /v1/content/channels | GetChannels lists the distribution channels the caller&#39;s org has connected — the social integrations a publish can target.
[**CloudGetV1ContentLifecycle**](ContentAPI.md#CloudGetV1ContentLifecycle) | **Get** /v1/content/lifecycle | GetLifecycle returns the ONE marketing-content state machine: the ordered lifecycle states, which state a fresh document starts in, which one is publicly live, and the legal successors of every state.
[**CloudPostV1ContentDoctypeNameTransition**](ContentAPI.md#CloudPostV1ContentDoctypeNameTransition) | **Post** /v1/content/{doctype}/{name}/transition | PostTransition moves one content item to a new lifecycle state and, on the move to published, fans it out to the item&#39;s channels.
[**CloudPostV1ContentGenerate**](ContentAPI.md#CloudPostV1ContentGenerate) | **Post** /v1/content/generate | 
[**CloudPostV1ContentPublish**](ContentAPI.md#CloudPostV1ContentPublish) | **Post** /v1/content/publish | Publish distributes one CMS content item to the channels recorded on it and returns the honest per-channel outcome.



## CloudGetV1ContentBoard

> CloudBoardPage CloudGetV1ContentBoard(ctx).Status(status).Project(project).Doctype(doctype).Limit(limit).Execute()

GetBoard aggregates the caller org's marketing content across every publishable content type into ONE queue board — the cross-type read the framework's per-DocType list cannot give.



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
	status := "queued" // string | Status keeps only items in one lifecycle state (draft, in_review, approved, queued, published, archived). An undefined state is refused. (optional)
	project := "project_example" // string | Project keeps only items in one brand/site sub-scope. (optional)
	doctype := "doctype_example" // string | DocType keeps only one content type; omitted, the board spans every publishable type. An unknown type is refused. (optional)
	limit := int32(50) // int32 | Limit caps the rows returned, clamped to 1000. Defaults to 200, which is also what a non-positive or unparseable value takes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.CloudGetV1ContentBoard(context.Background()).Status(status).Project(project).Doctype(doctype).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudGetV1ContentBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ContentBoard`: CloudBoardPage
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CloudGetV1ContentBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ContentBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only items in one lifecycle state (draft, in_review, approved, queued, published, archived). An undefined state is refused. | 
 **project** | **string** | Project keeps only items in one brand/site sub-scope. | 
 **doctype** | **string** | DocType keeps only one content type; omitted, the board spans every publishable type. An unknown type is refused. | 
 **limit** | **int32** | Limit caps the rows returned, clamped to 1000. Defaults to 200, which is also what a non-positive or unparseable value takes. | 

### Return type

[**CloudBoardPage**](CloudBoardPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ContentChannels

> CloudChannelList CloudGetV1ContentChannels(ctx).Execute()

GetChannels lists the distribution channels the caller's org has connected — the social integrations a publish can target.



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
	resp, r, err := apiClient.ContentAPI.CloudGetV1ContentChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudGetV1ContentChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ContentChannels`: CloudChannelList
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CloudGetV1ContentChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ContentChannelsRequest struct via the builder pattern


### Return type

[**CloudChannelList**](CloudChannelList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ContentLifecycle

> CloudStateGraph CloudGetV1ContentLifecycle(ctx).Execute()

GetLifecycle returns the ONE marketing-content state machine: the ordered lifecycle states, which state a fresh document starts in, which one is publicly live, and the legal successors of every state.



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
	resp, r, err := apiClient.ContentAPI.CloudGetV1ContentLifecycle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudGetV1ContentLifecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ContentLifecycle`: CloudStateGraph
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CloudGetV1ContentLifecycle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ContentLifecycleRequest struct via the builder pattern


### Return type

[**CloudStateGraph**](CloudStateGraph.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ContentDoctypeNameTransition

> CloudTransitionResult CloudPostV1ContentDoctypeNameTransition(ctx, doctype, name).CloudTransitionIn(cloudTransitionIn).Execute()

PostTransition moves one content item to a new lifecycle state and, on the move to published, fans it out to the item's channels.



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
	doctype := "SocialPost" // string | DocType is the content type to act on, from the path.
	name := "spring-teaser" // string | Name is the document to act on, from the path.
	cloudTransitionIn := *openapiclient.NewCloudTransitionIn() // CloudTransitionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.CloudPostV1ContentDoctypeNameTransition(context.Background(), doctype, name).CloudTransitionIn(cloudTransitionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudPostV1ContentDoctypeNameTransition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ContentDoctypeNameTransition`: CloudTransitionResult
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CloudPostV1ContentDoctypeNameTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the content type to act on, from the path. | 
**name** | **string** | Name is the document to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ContentDoctypeNameTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudTransitionIn** | [**CloudTransitionIn**](CloudTransitionIn.md) |  | 

### Return type

[**CloudTransitionResult**](CloudTransitionResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ContentGenerate

> CloudPostV1ContentGenerate(ctx).Execute()



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
	r, err := apiClient.ContentAPI.CloudPostV1ContentGenerate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudPostV1ContentGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ContentGenerateRequest struct via the builder pattern


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


## CloudPostV1ContentPublish

> CloudPublishResult CloudPostV1ContentPublish(ctx).CloudPublishInput(cloudPublishInput).Execute()

Publish distributes one CMS content item to the channels recorded on it and returns the honest per-channel outcome.



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
	cloudPublishInput := *openapiclient.NewCloudPublishInput() // CloudPublishInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.CloudPostV1ContentPublish(context.Background()).CloudPublishInput(cloudPublishInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.CloudPostV1ContentPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ContentPublish`: CloudPublishResult
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.CloudPostV1ContentPublish`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ContentPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPublishInput** | [**CloudPublishInput**](CloudPublishInput.md) |  | 

### Return type

[**CloudPublishResult**](CloudPublishResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

