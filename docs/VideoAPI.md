# \VideoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldYoutubeEmbed**](VideoAPI.md#WorldWorldYoutubeEmbed) | **Get** /v1/world/youtube/embed | Self-contained IFrame-API player page (text/html)
[**WorldWorldYoutubeLive**](VideoAPI.md#WorldWorldYoutubeLive) | **Get** /v1/world/youtube/live | Resolve a channel handle to its current LIVE video id



## WorldWorldYoutubeEmbed

> string WorldWorldYoutubeEmbed(ctx).VideoId(videoId).Execute()

Self-contained IFrame-API player page (text/html)

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
	videoId := "videoId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.WorldWorldYoutubeEmbed(context.Background()).VideoId(videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.WorldWorldYoutubeEmbed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldYoutubeEmbed`: string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.WorldWorldYoutubeEmbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldYoutubeEmbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoId** | **string** |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldYoutubeLive

> map[string]interface{} WorldWorldYoutubeLive(ctx).Channel(channel).Execute()

Resolve a channel handle to its current LIVE video id

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
	channel := "channel_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.WorldWorldYoutubeLive(context.Background()).Channel(channel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.WorldWorldYoutubeLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldYoutubeLive`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.WorldWorldYoutubeLive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldYoutubeLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** |  | 

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

