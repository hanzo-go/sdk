# \TrackerProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackerCreateProject**](TrackerProjectsAPI.md#TrackerCreateProject) | **Post** /v1/tracker/projects | Create a project
[**TrackerDeleteProject**](TrackerProjectsAPI.md#TrackerDeleteProject) | **Delete** /v1/tracker/projects/{key} | Delete a project and all its issues
[**TrackerGetProject**](TrackerProjectsAPI.md#TrackerGetProject) | **Get** /v1/tracker/projects/{key} | Get a project
[**TrackerListProjects**](TrackerProjectsAPI.md#TrackerListProjects) | **Get** /v1/tracker/projects | List projects
[**TrackerUpdateProject**](TrackerProjectsAPI.md#TrackerUpdateProject) | **Patch** /v1/tracker/projects/{key} | Update a project



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
	resp, r, err := apiClient.TrackerProjectsAPI.TrackerCreateProject(context.Background()).TrackerCreateProjectRequest(trackerCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerProjectsAPI.TrackerCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerCreateProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerProjectsAPI.TrackerCreateProject`: %v\n", resp)
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
	r, err := apiClient.TrackerProjectsAPI.TrackerDeleteProject(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerProjectsAPI.TrackerDeleteProject``: %v\n", err)
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
	resp, r, err := apiClient.TrackerProjectsAPI.TrackerGetProject(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerProjectsAPI.TrackerGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerGetProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerProjectsAPI.TrackerGetProject`: %v\n", resp)
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
	resp, r, err := apiClient.TrackerProjectsAPI.TrackerListProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerProjectsAPI.TrackerListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerListProjects`: []TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerProjectsAPI.TrackerListProjects`: %v\n", resp)
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
	resp, r, err := apiClient.TrackerProjectsAPI.TrackerUpdateProject(context.Background(), key).TrackerUpdateProjectRequest(trackerUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerProjectsAPI.TrackerUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerUpdateProject`: TrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerProjectsAPI.TrackerUpdateProject`: %v\n", resp)
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

