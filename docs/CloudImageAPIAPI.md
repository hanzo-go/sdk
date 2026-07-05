# \CloudImageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddImage**](CloudImageAPIAPI.md#CloudApiControllerAddImage) | **Post** /v1/cloud/add-image | Api Controller Add Image
[**CloudApiControllerDeleteImage**](CloudImageAPIAPI.md#CloudApiControllerDeleteImage) | **Post** /v1/cloud/delete-image | Api Controller Delete Image
[**CloudApiControllerGetImage**](CloudImageAPIAPI.md#CloudApiControllerGetImage) | **Get** /v1/cloud/get-image | Api Controller Get Image
[**CloudApiControllerGetImages**](CloudImageAPIAPI.md#CloudApiControllerGetImages) | **Get** /v1/cloud/get-images | Api Controller Get Images
[**CloudApiControllerUpdateImage**](CloudImageAPIAPI.md#CloudApiControllerUpdateImage) | **Post** /v1/cloud/update-image | Api Controller Update Image



## CloudApiControllerAddImage

> CloudControllersResponse CloudApiControllerAddImage(ctx).CloudObjectImage(cloudObjectImage).Execute()

Api Controller Add Image



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
	cloudObjectImage := *openapiclient.NewCloudObjectImage() // CloudObjectImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudImageAPIAPI.CloudApiControllerAddImage(context.Background()).CloudObjectImage(cloudObjectImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudImageAPIAPI.CloudApiControllerAddImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddImage`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudImageAPIAPI.CloudApiControllerAddImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectImage** | [**CloudObjectImage**](CloudObjectImage.md) | The details of the image | 

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


## CloudApiControllerDeleteImage

> CloudControllersResponse CloudApiControllerDeleteImage(ctx).CloudObjectImage(cloudObjectImage).Execute()

Api Controller Delete Image



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
	cloudObjectImage := *openapiclient.NewCloudObjectImage() // CloudObjectImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudImageAPIAPI.CloudApiControllerDeleteImage(context.Background()).CloudObjectImage(cloudObjectImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudImageAPIAPI.CloudApiControllerDeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteImage`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudImageAPIAPI.CloudApiControllerDeleteImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectImage** | [**CloudObjectImage**](CloudObjectImage.md) | The details of the image | 

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


## CloudApiControllerGetImage

> CloudObjectImage CloudApiControllerGetImage(ctx).Id(id).Execute()

Api Controller Get Image



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
	id := "id_example" // string | The id ( owner/name ) of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudImageAPIAPI.CloudApiControllerGetImage(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudImageAPIAPI.CloudApiControllerGetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetImage`: CloudObjectImage
	fmt.Fprintf(os.Stdout, "Response from `CloudImageAPIAPI.CloudApiControllerGetImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the image | 

### Return type

[**CloudObjectImage**](CloudObjectImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetImages

> CloudObjectImage CloudApiControllerGetImages(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Images



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
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudImageAPIAPI.CloudApiControllerGetImages(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudImageAPIAPI.CloudApiControllerGetImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetImages`: CloudObjectImage
	fmt.Fprintf(os.Stdout, "Response from `CloudImageAPIAPI.CloudApiControllerGetImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectImage**](CloudObjectImage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateImage

> CloudControllersResponse CloudApiControllerUpdateImage(ctx).Id(id).CloudObjectImage(cloudObjectImage).Execute()

Api Controller Update Image



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
	id := "id_example" // string | The id ( owner/name ) of the image
	cloudObjectImage := *openapiclient.NewCloudObjectImage() // CloudObjectImage | The details of the image

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudImageAPIAPI.CloudApiControllerUpdateImage(context.Background()).Id(id).CloudObjectImage(cloudObjectImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudImageAPIAPI.CloudApiControllerUpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateImage`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudImageAPIAPI.CloudApiControllerUpdateImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the image | 
 **cloudObjectImage** | [**CloudObjectImage**](CloudObjectImage.md) | The details of the image | 

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

