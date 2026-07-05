# \NexusImageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddImage**](NexusImageAPIAPI.md#NexusAddImage) | **Post** /v1/nexus/add-image | add Image
[**NexusDeleteImage**](NexusImageAPIAPI.md#NexusDeleteImage) | **Post** /v1/nexus/delete-image | delete Image
[**NexusGetImage**](NexusImageAPIAPI.md#NexusGetImage) | **Get** /v1/nexus/get-image | get Image
[**NexusGetImages**](NexusImageAPIAPI.md#NexusGetImages) | **Get** /v1/nexus/get-images | get Images
[**NexusUpdateImage**](NexusImageAPIAPI.md#NexusUpdateImage) | **Post** /v1/nexus/update-image | update Image



## NexusAddImage

> NexusResponse NexusAddImage(ctx).NexusImage(nexusImage).Execute()

add Image



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
	nexusImage := *openapiclient.NewNexusImage() // NexusImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusImageAPIAPI.NexusAddImage(context.Background()).NexusImage(nexusImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusImageAPIAPI.NexusAddImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddImage`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusImageAPIAPI.NexusAddImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusImage** | [**NexusImage**](NexusImage.md) | The details of the image | 

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


## NexusDeleteImage

> NexusResponse NexusDeleteImage(ctx).NexusImage(nexusImage).Execute()

delete Image



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
	nexusImage := *openapiclient.NewNexusImage() // NexusImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusImageAPIAPI.NexusDeleteImage(context.Background()).NexusImage(nexusImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusImageAPIAPI.NexusDeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteImage`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusImageAPIAPI.NexusDeleteImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusImage** | [**NexusImage**](NexusImage.md) | The details of the image | 

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


## NexusGetImage

> NexusImage NexusGetImage(ctx).Id(id).Execute()

get Image



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
	id := "id_example" // string | The id (owner/name) of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusImageAPIAPI.NexusGetImage(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusImageAPIAPI.NexusGetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetImage`: NexusImage
	fmt.Fprintf(os.Stdout, "Response from `NexusImageAPIAPI.NexusGetImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the image | 

### Return type

[**NexusImage**](NexusImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetImages

> NexusImage NexusGetImages(ctx).PageSize(pageSize).P(p).Execute()

get Images



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusImageAPIAPI.NexusGetImages(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusImageAPIAPI.NexusGetImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetImages`: NexusImage
	fmt.Fprintf(os.Stdout, "Response from `NexusImageAPIAPI.NexusGetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusImage**](NexusImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateImage

> NexusResponse NexusUpdateImage(ctx).Id(id).NexusImage(nexusImage).Execute()

update Image



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
	id := "id_example" // string | The id (owner/name) of the image
	nexusImage := *openapiclient.NewNexusImage() // NexusImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusImageAPIAPI.NexusUpdateImage(context.Background()).Id(id).NexusImage(nexusImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusImageAPIAPI.NexusUpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateImage`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusImageAPIAPI.NexusUpdateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the image | 
 **nexusImage** | [**NexusImage**](NexusImage.md) | The details of the image | 

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

