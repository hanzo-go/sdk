# \RegistryProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryCreateProject**](RegistryProjectsAPI.md#RegistryCreateProject) | **Post** /v1/registry/projects | Create project
[**RegistryDeleteProject**](RegistryProjectsAPI.md#RegistryDeleteProject) | **Delete** /v1/registry/projects/{name} | Delete project
[**RegistryGetProject**](RegistryProjectsAPI.md#RegistryGetProject) | **Get** /v1/registry/projects/{name} | Get project
[**RegistryListProjects**](RegistryProjectsAPI.md#RegistryListProjects) | **Get** /v1/registry/projects | List projects
[**RegistryUpdateProject**](RegistryProjectsAPI.md#RegistryUpdateProject) | **Put** /v1/registry/projects/{name} | Update project



## RegistryCreateProject

> map[string]interface{} RegistryCreateProject(ctx).RegistryProjectCreate(registryProjectCreate).Execute()

Create project

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
	registryProjectCreate := *openapiclient.NewRegistryProjectCreate("ProjectName_example") // RegistryProjectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryProjectsAPI.RegistryCreateProject(context.Background()).RegistryProjectCreate(registryProjectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryProjectsAPI.RegistryCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryCreateProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryProjectsAPI.RegistryCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registryProjectCreate** | [**RegistryProjectCreate**](RegistryProjectCreate.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryDeleteProject

> map[string]interface{} RegistryDeleteProject(ctx, name).Execute()

Delete project

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryProjectsAPI.RegistryDeleteProject(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryProjectsAPI.RegistryDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryDeleteProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryProjectsAPI.RegistryDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryDeleteProjectRequest struct via the builder pattern


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


## RegistryGetProject

> RegistryProject RegistryGetProject(ctx, name).Execute()

Get project

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryProjectsAPI.RegistryGetProject(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryProjectsAPI.RegistryGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryGetProject`: RegistryProject
	fmt.Fprintf(os.Stdout, "Response from `RegistryProjectsAPI.RegistryGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistryProject**](RegistryProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryListProjects

> []RegistryProject RegistryListProjects(ctx).Name(name).Public(public).Page(page).PageSize(pageSize).Sort(sort).Execute()

List projects

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
	name := "name_example" // string | Filter by project name (fuzzy match) (optional)
	public := true // bool | Filter by public/private (optional)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sort := "sort_example" // string |  (optional) (default to "creation_time")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryProjectsAPI.RegistryListProjects(context.Background()).Name(name).Public(public).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryProjectsAPI.RegistryListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListProjects`: []RegistryProject
	fmt.Fprintf(os.Stdout, "Response from `RegistryProjectsAPI.RegistryListProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter by project name (fuzzy match) | 
 **public** | **bool** | Filter by public/private | 
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sort** | **string** |  | [default to &quot;creation_time&quot;]

### Return type

[**[]RegistryProject**](RegistryProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistryUpdateProject

> map[string]interface{} RegistryUpdateProject(ctx, name).RegistryUpdateProjectRequest(registryUpdateProjectRequest).Execute()

Update project

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
	name := "name_example" // string | 
	registryUpdateProjectRequest := *openapiclient.NewRegistryUpdateProjectRequest() // RegistryUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryProjectsAPI.RegistryUpdateProject(context.Background(), name).RegistryUpdateProjectRequest(registryUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryProjectsAPI.RegistryUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryUpdateProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RegistryProjectsAPI.RegistryUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistryUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registryUpdateProjectRequest** | [**RegistryUpdateProjectRequest**](RegistryUpdateProjectRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

