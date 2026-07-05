# \KmsProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateProject**](KmsProjectsAPI.md#KmsCreateProject) | **Post** /v1/kms/projects | Create a project
[**KmsDeleteProject**](KmsProjectsAPI.md#KmsDeleteProject) | **Delete** /v1/kms/projects/{projectId} | Delete a project
[**KmsGetProject**](KmsProjectsAPI.md#KmsGetProject) | **Get** /v1/kms/projects/{projectId} | Get a project by ID
[**KmsListProjectUsers**](KmsProjectsAPI.md#KmsListProjectUsers) | **Get** /v1/kms/projects/{projectId}/users | List project members
[**KmsUpdateProject**](KmsProjectsAPI.md#KmsUpdateProject) | **Patch** /v1/kms/projects/{projectId} | Update a project



## KmsCreateProject

> KmsCreateProject200Response KmsCreateProject(ctx).KmsCreateProjectRequest(kmsCreateProjectRequest).Execute()

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
	kmsCreateProjectRequest := *openapiclient.NewKmsCreateProjectRequest("ProjectName_example", "OrganizationId_example") // KmsCreateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsProjectsAPI.KmsCreateProject(context.Background()).KmsCreateProjectRequest(kmsCreateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsProjectsAPI.KmsCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateProject`: KmsCreateProject200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsProjectsAPI.KmsCreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateProjectRequest** | [**KmsCreateProjectRequest**](KmsCreateProjectRequest.md) |  | 

### Return type

[**KmsCreateProject200Response**](KmsCreateProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteProject

> map[string]interface{} KmsDeleteProject(ctx, projectId).Execute()

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsProjectsAPI.KmsDeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsProjectsAPI.KmsDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsProjectsAPI.KmsDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteProjectRequest struct via the builder pattern


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


## KmsGetProject

> KmsCreateProject200Response KmsGetProject(ctx, projectId).Execute()

Get a project by ID

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
	resp, r, err := apiClient.KmsProjectsAPI.KmsGetProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsProjectsAPI.KmsGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetProject`: KmsCreateProject200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsProjectsAPI.KmsGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsCreateProject200Response**](KmsCreateProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListProjectUsers

> KmsListProjectUsers200Response KmsListProjectUsers(ctx, projectId).IncludeGroupMembers(includeGroupMembers).Execute()

List project members

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
	includeGroupMembers := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsProjectsAPI.KmsListProjectUsers(context.Background(), projectId).IncludeGroupMembers(includeGroupMembers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsProjectsAPI.KmsListProjectUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListProjectUsers`: KmsListProjectUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsProjectsAPI.KmsListProjectUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsListProjectUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGroupMembers** | **bool** |  | [default to false]

### Return type

[**KmsListProjectUsers200Response**](KmsListProjectUsers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateProject

> KmsCreateProject200Response KmsUpdateProject(ctx, projectId).KmsUpdateProjectRequest(kmsUpdateProjectRequest).Execute()

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateProjectRequest := *openapiclient.NewKmsUpdateProjectRequest() // KmsUpdateProjectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsProjectsAPI.KmsUpdateProject(context.Background(), projectId).KmsUpdateProjectRequest(kmsUpdateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsProjectsAPI.KmsUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateProject`: KmsCreateProject200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsProjectsAPI.KmsUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateProjectRequest** | [**KmsUpdateProjectRequest**](KmsUpdateProjectRequest.md) |  | 

### Return type

[**KmsCreateProject200Response**](KmsCreateProject200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

