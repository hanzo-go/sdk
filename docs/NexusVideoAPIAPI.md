# \NexusVideoAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddVideo**](NexusVideoAPIAPI.md#NexusAddVideo) | **Post** /v1/nexus/add-video | add Video
[**NexusDeleteVideo**](NexusVideoAPIAPI.md#NexusDeleteVideo) | **Post** /v1/nexus/delete-video | delete Video
[**NexusGetGlobalVideos**](NexusVideoAPIAPI.md#NexusGetGlobalVideos) | **Get** /v1/nexus/get-global-videos | get Global Videos
[**NexusGetVideo**](NexusVideoAPIAPI.md#NexusGetVideo) | **Get** /v1/nexus/get-video | get Video
[**NexusGetVideos**](NexusVideoAPIAPI.md#NexusGetVideos) | **Get** /v1/nexus/get-videos | get Videos
[**NexusUpdateVideo**](NexusVideoAPIAPI.md#NexusUpdateVideo) | **Post** /v1/nexus/update-video | update Video
[**NexusUploadVideo**](NexusVideoAPIAPI.md#NexusUploadVideo) | **Post** /v1/nexus/upload-video | upload Video



## NexusAddVideo

> NexusResponse NexusAddVideo(ctx).NexusVideo(nexusVideo).Execute()

add Video



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
	nexusVideo := *openapiclient.NewNexusVideo() // NexusVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusAddVideo(context.Background()).NexusVideo(nexusVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusAddVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddVideo`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusAddVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusVideo** | [**NexusVideo**](NexusVideo.md) | The details of the video | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteVideo

> NexusResponse NexusDeleteVideo(ctx).NexusVideo(nexusVideo).Execute()

delete Video



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
	nexusVideo := *openapiclient.NewNexusVideo() // NexusVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusDeleteVideo(context.Background()).NexusVideo(nexusVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusDeleteVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteVideo`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusDeleteVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusVideo** | [**NexusVideo**](NexusVideo.md) | The details of the video | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalVideos

> []NexusVideo NexusGetGlobalVideos(ctx).Execute()

get Global Videos



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
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusGetGlobalVideos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusGetGlobalVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalVideos`: []NexusVideo
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusGetGlobalVideos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalVideosRequest struct via the builder pattern


### Return type

[**[]NexusVideo**](NexusVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetVideo

> NexusVideo NexusGetVideo(ctx).Id(id).Execute()

get Video



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
	id := "id_example" // string | The id of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusGetVideo(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusGetVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetVideo`: NexusVideo
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusGetVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of the video | 

### Return type

[**NexusVideo**](NexusVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetVideos

> []NexusVideo NexusGetVideos(ctx).Owner(owner).Execute()

get Videos



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
	owner := "owner_example" // string | The owner of the videos

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusGetVideos(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusGetVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetVideos`: []NexusVideo
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusGetVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the videos | 

### Return type

[**[]NexusVideo**](NexusVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateVideo

> NexusResponse NexusUpdateVideo(ctx).Id(id).NexusVideo(nexusVideo).Execute()

update Video



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
	id := "id_example" // string | The id (owner/name) of the video
	nexusVideo := *openapiclient.NewNexusVideo() // NexusVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusUpdateVideo(context.Background()).Id(id).NexusVideo(nexusVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusUpdateVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateVideo`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusUpdateVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the video | 
 **nexusVideo** | [**NexusVideo**](NexusVideo.md) | The details of the video | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUploadVideo

> string NexusUploadVideo(ctx).File(file).Execute()

upload Video



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
	file := os.NewFile(1234, "some_file") // *os.File | The video file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusVideoAPIAPI.NexusUploadVideo(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusVideoAPIAPI.NexusUploadVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUploadVideo`: string
	fmt.Fprintf(os.Stdout, "Response from `NexusVideoAPIAPI.NexusUploadVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUploadVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The video file | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

