# \VideoAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddVideo**](VideoAPIAPI.md#CloudApiControllerAddVideo) | **Post** /v1/cloud/add-video | Api Controller Add Video
[**CloudApiControllerDeleteVideo**](VideoAPIAPI.md#CloudApiControllerDeleteVideo) | **Post** /v1/cloud/delete-video | Api Controller Delete Video
[**CloudApiControllerGetGlobalVideos**](VideoAPIAPI.md#CloudApiControllerGetGlobalVideos) | **Get** /v1/cloud/get-global-videos | Api Controller Get Global Videos
[**CloudApiControllerGetVideo**](VideoAPIAPI.md#CloudApiControllerGetVideo) | **Get** /v1/cloud/get-video | Api Controller Get Video
[**CloudApiControllerGetVideos**](VideoAPIAPI.md#CloudApiControllerGetVideos) | **Get** /v1/cloud/get-videos | Api Controller Get Videos
[**CloudApiControllerUpdateVideo**](VideoAPIAPI.md#CloudApiControllerUpdateVideo) | **Post** /v1/cloud/update-video | Api Controller Update Video
[**CloudApiControllerUploadVideo**](VideoAPIAPI.md#CloudApiControllerUploadVideo) | **Post** /v1/cloud/upload-video | Api Controller Upload Video



## CloudApiControllerAddVideo

> CloudControllersResponse CloudApiControllerAddVideo(ctx).CloudObjectVideo(cloudObjectVideo).Execute()

Api Controller Add Video



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
	cloudObjectVideo := *openapiclient.NewCloudObjectVideo() // CloudObjectVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerAddVideo(context.Background()).CloudObjectVideo(cloudObjectVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerAddVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddVideo`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerAddVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectVideo** | [**CloudObjectVideo**](CloudObjectVideo.md) | The details of the video | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteVideo

> CloudControllersResponse CloudApiControllerDeleteVideo(ctx).CloudObjectVideo(cloudObjectVideo).Execute()

Api Controller Delete Video



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
	cloudObjectVideo := *openapiclient.NewCloudObjectVideo() // CloudObjectVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerDeleteVideo(context.Background()).CloudObjectVideo(cloudObjectVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerDeleteVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteVideo`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerDeleteVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectVideo** | [**CloudObjectVideo**](CloudObjectVideo.md) | The details of the video | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalVideos

> []CloudObjectVideo CloudApiControllerGetGlobalVideos(ctx).Execute()

Api Controller Get Global Videos



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
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerGetGlobalVideos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerGetGlobalVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalVideos`: []CloudObjectVideo
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerGetGlobalVideos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalVideosRequest struct via the builder pattern


### Return type

[**[]CloudObjectVideo**](CloudObjectVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetVideo

> CloudObjectVideo CloudApiControllerGetVideo(ctx).Id(id).Execute()

Api Controller Get Video



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
	id := "id_example" // string | The id of video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerGetVideo(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerGetVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetVideo`: CloudObjectVideo
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerGetVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id of video | 

### Return type

[**CloudObjectVideo**](CloudObjectVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetVideos

> []CloudObjectVideo CloudApiControllerGetVideos(ctx).Owner(owner).Execute()

Api Controller Get Videos



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
	owner := "owner_example" // string | The owner of videos

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerGetVideos(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerGetVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetVideos`: []CloudObjectVideo
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerGetVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of videos | 

### Return type

[**[]CloudObjectVideo**](CloudObjectVideo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateVideo

> CloudControllersResponse CloudApiControllerUpdateVideo(ctx).Id(id).CloudObjectVideo(cloudObjectVideo).Execute()

Api Controller Update Video



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
	cloudObjectVideo := *openapiclient.NewCloudObjectVideo() // CloudObjectVideo | The details of the video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerUpdateVideo(context.Background()).Id(id).CloudObjectVideo(cloudObjectVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerUpdateVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateVideo`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerUpdateVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the video | 
 **cloudObjectVideo** | [**CloudObjectVideo**](CloudObjectVideo.md) | The details of the video | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUploadVideo

> string CloudApiControllerUploadVideo(ctx).File(file).Execute()

Api Controller Upload Video



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
	file := os.NewFile(1234, "some_file") // *os.File | The video file to upload

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPIAPI.CloudApiControllerUploadVideo(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPIAPI.CloudApiControllerUploadVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUploadVideo`: string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPIAPI.CloudApiControllerUploadVideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUploadVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The video file to upload | 

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

