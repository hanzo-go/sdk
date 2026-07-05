# \StreamConsumerGroupsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamCommitOffsets**](StreamConsumerGroupsAPI.md#StreamCommitOffsets) | **Put** /v1/stream/groups/{group_id}/offsets | Commit offsets
[**StreamDeleteConsumerGroup**](StreamConsumerGroupsAPI.md#StreamDeleteConsumerGroup) | **Delete** /v1/stream/groups/{group_id} | Delete a consumer group
[**StreamGetConsumerGroup**](StreamConsumerGroupsAPI.md#StreamGetConsumerGroup) | **Get** /v1/stream/groups/{group_id} | Get consumer group details
[**StreamGetGroupOffsets**](StreamConsumerGroupsAPI.md#StreamGetGroupOffsets) | **Get** /v1/stream/groups/{group_id}/offsets | Get committed offsets
[**StreamListConsumerGroups**](StreamConsumerGroupsAPI.md#StreamListConsumerGroups) | **Get** /v1/stream/groups | List consumer groups



## StreamCommitOffsets

> StreamCommitOffsets(ctx, groupId).StreamGetGroupOffsets200Response(streamGetGroupOffsets200Response).Execute()

Commit offsets



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
	groupId := "groupId_example" // string | 
	streamGetGroupOffsets200Response := *openapiclient.NewStreamGetGroupOffsets200Response() // StreamGetGroupOffsets200Response | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StreamConsumerGroupsAPI.StreamCommitOffsets(context.Background(), groupId).StreamGetGroupOffsets200Response(streamGetGroupOffsets200Response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamConsumerGroupsAPI.StreamCommitOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamCommitOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamGetGroupOffsets200Response** | [**StreamGetGroupOffsets200Response**](StreamGetGroupOffsets200Response.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamDeleteConsumerGroup

> StreamDeleteConsumerGroup(ctx, groupId).Execute()

Delete a consumer group

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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StreamConsumerGroupsAPI.StreamDeleteConsumerGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamConsumerGroupsAPI.StreamDeleteConsumerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamDeleteConsumerGroupRequest struct via the builder pattern


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


## StreamGetConsumerGroup

> StreamConsumerGroup StreamGetConsumerGroup(ctx, groupId).Execute()

Get consumer group details



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamConsumerGroupsAPI.StreamGetConsumerGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamConsumerGroupsAPI.StreamGetConsumerGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamGetConsumerGroup`: StreamConsumerGroup
	fmt.Fprintf(os.Stdout, "Response from `StreamConsumerGroupsAPI.StreamGetConsumerGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGetConsumerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StreamConsumerGroup**](StreamConsumerGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamGetGroupOffsets

> StreamGetGroupOffsets200Response StreamGetGroupOffsets(ctx, groupId).Topic(topic).Execute()

Get committed offsets

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
	groupId := "groupId_example" // string | 
	topic := "topic_example" // string | Filter by topic (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamConsumerGroupsAPI.StreamGetGroupOffsets(context.Background(), groupId).Topic(topic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamConsumerGroupsAPI.StreamGetGroupOffsets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamGetGroupOffsets`: StreamGetGroupOffsets200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamConsumerGroupsAPI.StreamGetGroupOffsets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamGetGroupOffsetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **topic** | **string** | Filter by topic | 

### Return type

[**StreamGetGroupOffsets200Response**](StreamGetGroupOffsets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamListConsumerGroups

> StreamListConsumerGroups200Response StreamListConsumerGroups(ctx).Execute()

List consumer groups

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
	resp, r, err := apiClient.StreamConsumerGroupsAPI.StreamListConsumerGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamConsumerGroupsAPI.StreamListConsumerGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamListConsumerGroups`: StreamListConsumerGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamConsumerGroupsAPI.StreamListConsumerGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamListConsumerGroupsRequest struct via the builder pattern


### Return type

[**StreamListConsumerGroups200Response**](StreamListConsumerGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

