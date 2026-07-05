# \KmsSecretTagsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSecretTag**](KmsSecretTagsAPI.md#KmsCreateSecretTag) | **Post** /v1/kms/projects/{projectId}/tags | Create a secret tag
[**KmsDeleteSecretTag**](KmsSecretTagsAPI.md#KmsDeleteSecretTag) | **Delete** /v1/kms/projects/{projectId}/tags/{tagId} | Delete a secret tag
[**KmsGetSecretTag**](KmsSecretTagsAPI.md#KmsGetSecretTag) | **Get** /v1/kms/projects/{projectId}/tags/{tagId} | Get a secret tag by ID
[**KmsListSecretTags**](KmsSecretTagsAPI.md#KmsListSecretTags) | **Get** /v1/kms/projects/{projectId}/tags | List secret tags for a project
[**KmsUpdateSecretTag**](KmsSecretTagsAPI.md#KmsUpdateSecretTag) | **Patch** /v1/kms/projects/{projectId}/tags/{tagId} | Update a secret tag



## KmsCreateSecretTag

> KmsCreateSecretTag200Response KmsCreateSecretTag(ctx, projectId).KmsCreateSecretTagRequest(kmsCreateSecretTagRequest).Execute()

Create a secret tag

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsCreateSecretTagRequest := *openapiclient.NewKmsCreateSecretTagRequest("Name_example", "Slug_example", "Color_example") // KmsCreateSecretTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretTagsAPI.KmsCreateSecretTag(context.Background(), projectId).KmsCreateSecretTagRequest(kmsCreateSecretTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretTagsAPI.KmsCreateSecretTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSecretTag`: KmsCreateSecretTag200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretTagsAPI.KmsCreateSecretTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSecretTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsCreateSecretTagRequest** | [**KmsCreateSecretTagRequest**](KmsCreateSecretTagRequest.md) |  | 

### Return type

[**KmsCreateSecretTag200Response**](KmsCreateSecretTag200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSecretTag

> map[string]interface{} KmsDeleteSecretTag(ctx, projectId, tagId).Execute()

Delete a secret tag

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretTagsAPI.KmsDeleteSecretTag(context.Background(), projectId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretTagsAPI.KmsDeleteSecretTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSecretTag`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretTagsAPI.KmsDeleteSecretTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSecretTagRequest struct via the builder pattern


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


## KmsGetSecretTag

> KmsCreateSecretTag200Response KmsGetSecretTag(ctx, projectId, tagId).Execute()

Get a secret tag by ID

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretTagsAPI.KmsGetSecretTag(context.Background(), projectId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretTagsAPI.KmsGetSecretTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetSecretTag`: KmsCreateSecretTag200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretTagsAPI.KmsGetSecretTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetSecretTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KmsCreateSecretTag200Response**](KmsCreateSecretTag200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListSecretTags

> KmsListSecretTags200Response KmsListSecretTags(ctx, projectId).Execute()

List secret tags for a project

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretTagsAPI.KmsListSecretTags(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretTagsAPI.KmsListSecretTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSecretTags`: KmsListSecretTags200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretTagsAPI.KmsListSecretTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSecretTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsListSecretTags200Response**](KmsListSecretTags200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateSecretTag

> KmsCreateSecretTag200Response KmsUpdateSecretTag(ctx, projectId, tagId).KmsUpdateSecretTagRequest(kmsUpdateSecretTagRequest).Execute()

Update a secret tag

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tagId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateSecretTagRequest := *openapiclient.NewKmsUpdateSecretTagRequest() // KmsUpdateSecretTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretTagsAPI.KmsUpdateSecretTag(context.Background(), projectId, tagId).KmsUpdateSecretTagRequest(kmsUpdateSecretTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretTagsAPI.KmsUpdateSecretTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateSecretTag`: KmsCreateSecretTag200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretTagsAPI.KmsUpdateSecretTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateSecretTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kmsUpdateSecretTagRequest** | [**KmsUpdateSecretTagRequest**](KmsUpdateSecretTagRequest.md) |  | 

### Return type

[**KmsCreateSecretTag200Response**](KmsCreateSecretTag200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

