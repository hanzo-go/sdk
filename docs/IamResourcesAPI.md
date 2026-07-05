# \IamResourcesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddResource**](IamResourcesAPI.md#IamApiControllerAddResource) | **Post** /v1/iam/resources | Api Controller Add Resource
[**IamApiControllerDeleteResource**](IamResourcesAPI.md#IamApiControllerDeleteResource) | **Delete** /v1/iam/resources/{id} | Api Controller Delete Resource
[**IamApiControllerGetResource**](IamResourcesAPI.md#IamApiControllerGetResource) | **Get** /v1/iam/resources/{id} | Api Controller Get Resource
[**IamApiControllerGetResources**](IamResourcesAPI.md#IamApiControllerGetResources) | **Get** /v1/iam/resources | Api Controller Get Resources
[**IamApiControllerUpdateResource**](IamResourcesAPI.md#IamApiControllerUpdateResource) | **Put** /v1/iam/resources/{id} | Api Controller Update Resource
[**IamApiControllerUploadResource**](IamResourcesAPI.md#IamApiControllerUploadResource) | **Post** /v1/iam/resources/upload | Api Controller Upload Resource



## IamApiControllerAddResource

> IamControllersResponse IamApiControllerAddResource(ctx).IamObjectResource(iamObjectResource).Execute()

Api Controller Add Resource

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
	iamObjectResource := *openapiclient.NewIamObjectResource() // IamObjectResource | Resource object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerAddResource(context.Background()).IamObjectResource(iamObjectResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerAddResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddResource`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerAddResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectResource** | [**IamObjectResource**](IamObjectResource.md) | Resource object | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteResource

> IamControllersResponse IamApiControllerDeleteResource(ctx, id).IamObjectResource(iamObjectResource).Execute()

Api Controller Delete Resource

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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectResource := *openapiclient.NewIamObjectResource() // IamObjectResource | Resource object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerDeleteResource(context.Background(), id).IamObjectResource(iamObjectResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerDeleteResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteResource`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerDeleteResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectResource** | [**IamObjectResource**](IamObjectResource.md) | Resource object | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetResource

> IamObjectResource IamApiControllerGetResource(ctx, id).Execute()

Api Controller Get Resource



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
	id := "id_example" // string | The id ( owner/name ) of resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerGetResource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerGetResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetResource`: IamObjectResource
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerGetResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectResource**](IamObjectResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetResources

> []IamObjectResource IamApiControllerGetResources(ctx).Owner(owner).User(user).PageSize(pageSize).P(p).Field(field).Value(value).SortField(sortField).SortOrder(sortOrder).Execute()

Api Controller Get Resources



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
	owner := "owner_example" // string | Owner
	user := "user_example" // string | User
	pageSize := int32(56) // int32 | Page Size (optional)
	p := int32(56) // int32 | Page Number (optional)
	field := "field_example" // string | Field (optional)
	value := "value_example" // string | Value (optional)
	sortField := "sortField_example" // string | Sort Field (optional)
	sortOrder := "sortOrder_example" // string | Sort Order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerGetResources(context.Background()).Owner(owner).User(user).PageSize(pageSize).P(p).Field(field).Value(value).SortField(sortField).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerGetResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetResources`: []IamObjectResource
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerGetResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner | 
 **user** | **string** | User | 
 **pageSize** | **int32** | Page Size | 
 **p** | **int32** | Page Number | 
 **field** | **string** | Field | 
 **value** | **string** | Value | 
 **sortField** | **string** | Sort Field | 
 **sortOrder** | **string** | Sort Order | 

### Return type

[**[]IamObjectResource**](IamObjectResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateResource

> IamControllersResponse IamApiControllerUpdateResource(ctx, id).IamObjectResource(iamObjectResource).Execute()

Api Controller Update Resource



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
	id := "id_example" // string | The id ( owner/name ) of resource
	iamObjectResource := *openapiclient.NewIamObjectResource() // IamObjectResource | The resource object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerUpdateResource(context.Background(), id).IamObjectResource(iamObjectResource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerUpdateResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateResource`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerUpdateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectResource** | [**IamObjectResource**](IamObjectResource.md) | The resource object | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUploadResource

> IamObjectResource IamApiControllerUploadResource(ctx).Owner(owner).User(user).Application(application).FullFilePath(fullFilePath).File(file).Tag(tag).Parent(parent).CreatedTime(createdTime).Description(description).Execute()

Api Controller Upload Resource

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
	owner := "owner_example" // string | Owner
	user := "user_example" // string | User
	application := "application_example" // string | Application
	fullFilePath := "fullFilePath_example" // string | Full File Path
	file := os.NewFile(1234, "some_file") // *os.File | Resource file
	tag := "tag_example" // string | Tag (optional)
	parent := "parent_example" // string | Parent (optional)
	createdTime := "createdTime_example" // string | Created Time (optional)
	description := "description_example" // string | Description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamResourcesAPI.IamApiControllerUploadResource(context.Background()).Owner(owner).User(user).Application(application).FullFilePath(fullFilePath).File(file).Tag(tag).Parent(parent).CreatedTime(createdTime).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamResourcesAPI.IamApiControllerUploadResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUploadResource`: IamObjectResource
	fmt.Fprintf(os.Stdout, "Response from `IamResourcesAPI.IamApiControllerUploadResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUploadResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | Owner | 
 **user** | **string** | User | 
 **application** | **string** | Application | 
 **fullFilePath** | **string** | Full File Path | 
 **file** | ***os.File** | Resource file | 
 **tag** | **string** | Tag | 
 **parent** | **string** | Parent | 
 **createdTime** | **string** | Created Time | 
 **description** | **string** | Description | 

### Return type

[**IamObjectResource**](IamObjectResource.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

