# \AppProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppCreateProject**](AppProjectsAPI.md#AppCreateProject) | **Post** /v1/projects | Create a project
[**AppDeleteProject**](AppProjectsAPI.md#AppDeleteProject) | **Delete** /v1/projects/{slug} | Delete a project
[**AppGetProject**](AppProjectsAPI.md#AppGetProject) | **Get** /v1/projects/{slug} | Get a project
[**AppListProjects**](AppProjectsAPI.md#AppListProjects) | **Get** /v1/projects | List projects
[**AppUpdateProject**](AppProjectsAPI.md#AppUpdateProject) | **Patch** /v1/projects/{slug} | Update a project



## AppCreateProject

> AppProject AppCreateProject(ctx).AppProjectCreate(appProjectCreate).Execute()

Create a project

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
	appProjectCreate := *openapiclient.NewAppProjectCreate("Slug_example") // AppProjectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppProjectsAPI.AppCreateProject(context.Background()).AppProjectCreate(appProjectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppProjectsAPI.AppCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppCreateProject`: AppProject
	fmt.Fprintf(os.Stdout, "Response from `AppProjectsAPI.AppCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appProjectCreate** | [**AppProjectCreate**](AppProjectCreate.md) |  | 

### Return type

[**AppProject**](AppProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppDeleteProject

> AppDeleteProject(ctx, slug).Execute()

Delete a project

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
	slug := "slug_example" // string | The project's URL-safe slug (unique within the org).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppProjectsAPI.AppDeleteProject(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppProjectsAPI.AppDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The project&#39;s URL-safe slug (unique within the org). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppGetProject

> AppProject AppGetProject(ctx, slug).Execute()

Get a project

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
	slug := "slug_example" // string | The project's URL-safe slug (unique within the org).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppProjectsAPI.AppGetProject(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppProjectsAPI.AppGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppGetProject`: AppProject
	fmt.Fprintf(os.Stdout, "Response from `AppProjectsAPI.AppGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The project&#39;s URL-safe slug (unique within the org). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppProject**](AppProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppListProjects

> AppListProjects200Response AppListProjects(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppProjectsAPI.AppListProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppProjectsAPI.AppListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppListProjects`: AppListProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `AppProjectsAPI.AppListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAppListProjectsRequest struct via the builder pattern


### Return type

[**AppListProjects200Response**](AppListProjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppUpdateProject

> AppProject AppUpdateProject(ctx, slug).AppProjectUpdate(appProjectUpdate).Execute()

Update a project

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
	slug := "slug_example" // string | The project's URL-safe slug (unique within the org).
	appProjectUpdate := *openapiclient.NewAppProjectUpdate() // AppProjectUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppProjectsAPI.AppUpdateProject(context.Background(), slug).AppProjectUpdate(appProjectUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppProjectsAPI.AppUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppUpdateProject`: AppProject
	fmt.Fprintf(os.Stdout, "Response from `AppProjectsAPI.AppUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The project&#39;s URL-safe slug (unique within the org). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appProjectUpdate** | [**AppProjectUpdate**](AppProjectUpdate.md) |  | 

### Return type

[**AppProject**](AppProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

