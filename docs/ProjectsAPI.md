# \ProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsCreateProject**](ProjectsAPI.md#ProjectsCreateProject) | **Post** /v1/projects | Create a project
[**ProjectsDeleteProject**](ProjectsAPI.md#ProjectsDeleteProject) | **Delete** /v1/projects/{slug} | Delete a project
[**ProjectsForkProject**](ProjectsAPI.md#ProjectsForkProject) | **Post** /v1/projects/fork | Fork a starter template into a new project
[**ProjectsGetProject**](ProjectsAPI.md#ProjectsGetProject) | **Get** /v1/projects/{slug} | Get a project
[**ProjectsListProjects**](ProjectsAPI.md#ProjectsListProjects) | **Get** /v1/projects | List projects
[**ProjectsUpdateProject**](ProjectsAPI.md#ProjectsUpdateProject) | **Patch** /v1/projects/{slug} | Update a project
[**TrackerCreateProject**](ProjectsAPI.md#TrackerCreateProject) | **Post** /v1/tracker/projects | Create a project
[**TrackerDeleteProject**](ProjectsAPI.md#TrackerDeleteProject) | **Delete** /v1/tracker/projects/{key} | Delete a project and all its issues
[**TrackerGetProject**](ProjectsAPI.md#TrackerGetProject) | **Get** /v1/tracker/projects/{key} | Get a project
[**TrackerListProjects**](ProjectsAPI.md#TrackerListProjects) | **Get** /v1/tracker/projects | List projects
[**TrackerUpdateProject**](ProjectsAPI.md#TrackerUpdateProject) | **Patch** /v1/tracker/projects/{key} | Update a project



## ProjectsCreateProject

> ProjectsProject ProjectsCreateProject(ctx).ProjectsCreateProjectRequest(projectsCreateProjectRequest).Execute()

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
	projectsCreateProjectRequest := *openapiclient.NewProjectsCreateProjectRequest("Name_example") // ProjectsCreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsCreateProject(context.Background()).ProjectsCreateProjectRequest(projectsCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCreateProject`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsCreateProjectRequest** | [**ProjectsCreateProjectRequest**](ProjectsCreateProjectRequest.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDeleteProject

> ProjectsDeleteProject(ctx, slug).Execute()

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
	slug := "slug_example" // string | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsDeleteProject(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeleteProjectRequest struct via the builder pattern


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


## ProjectsForkProject

> ProjectsProject ProjectsForkProject(ctx).ProjectsForkProjectRequest(projectsForkProjectRequest).Execute()

Fork a starter template into a new project



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
	projectsForkProjectRequest := *openapiclient.NewProjectsForkProjectRequest("Slug_example") // ProjectsForkProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsForkProject(context.Background()).ProjectsForkProjectRequest(projectsForkProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsForkProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsForkProject`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsForkProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsForkProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsForkProjectRequest** | [**ProjectsForkProjectRequest**](ProjectsForkProjectRequest.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGetProject

> ProjectsProject ProjectsGetProject(ctx, slug).Execute()

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
	slug := "slug_example" // string | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsGetProject(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGetProject`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListProjects

> []ProjectsProject ProjectsListProjects(ctx).Execute()

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
	resp, r, err := apiClient.ProjectsAPI.ProjectsListProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsListProjects`: []ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListProjectsRequest struct via the builder pattern


### Return type

[**[]ProjectsProject**](ProjectsProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdateProject

> ProjectsProject ProjectsUpdateProject(ctx, slug).ProjectsUpdateProjectRequest(projectsUpdateProjectRequest).Execute()

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
	slug := "slug_example" // string | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label.
	projectsUpdateProjectRequest := *openapiclient.NewProjectsUpdateProjectRequest() // ProjectsUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsUpdateProject(context.Background(), slug).ProjectsUpdateProjectRequest(projectsUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsUpdateProject`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdateProjectRequest** | [**ProjectsUpdateProjectRequest**](ProjectsUpdateProjectRequest.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerCreateProject

> TrackerProject TrackerCreateProject(ctx).TrackerCreateProjectRequest(trackerCreateProjectRequest).Execute()

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
	trackerCreateProjectRequest := *openapiclient.NewTrackerCreateProjectRequest("Name_example") // TrackerCreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.TrackerCreateProject(context.Background()).TrackerCreateProjectRequest(trackerCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TrackerCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerCreateProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.TrackerCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackerCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackerCreateProjectRequest** | [**TrackerCreateProjectRequest**](TrackerCreateProjectRequest.md) |  | 

### Return type

[**TrackerProject**](TrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerDeleteProject

> TrackerDeleteProject(ctx, key).Execute()

Delete a project and all its issues

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
	key := "key_example" // string | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.TrackerDeleteProject(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TrackerDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$) | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerDeleteProjectRequest struct via the builder pattern


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


## TrackerGetProject

> TrackerProject TrackerGetProject(ctx, key).Execute()

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
	key := "key_example" // string | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.TrackerGetProject(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TrackerGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerGetProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.TrackerGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$) | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrackerProject**](TrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerListProjects

> []TrackerProject TrackerListProjects(ctx).Execute()

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
	resp, r, err := apiClient.ProjectsAPI.TrackerListProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TrackerListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerListProjects`: []TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.TrackerListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerListProjectsRequest struct via the builder pattern


### Return type

[**[]TrackerProject**](TrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerUpdateProject

> TrackerProject TrackerUpdateProject(ctx, key).TrackerUpdateProjectRequest(trackerUpdateProjectRequest).Execute()

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
	key := "key_example" // string | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$)
	trackerUpdateProjectRequest := *openapiclient.NewTrackerUpdateProjectRequest() // TrackerUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.TrackerUpdateProject(context.Background(), key).TrackerUpdateProjectRequest(trackerUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.TrackerUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerUpdateProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.TrackerUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key (uppercase, ^[A-Z][A-Z0-9]{1,7}$) | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackerUpdateProjectRequest** | [**TrackerUpdateProjectRequest**](TrackerUpdateProjectRequest.md) |  | 

### Return type

[**TrackerProject**](TrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

