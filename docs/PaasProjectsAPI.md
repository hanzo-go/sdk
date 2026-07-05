# \PaasProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasCreateProject**](PaasProjectsAPI.md#PaasCreateProject) | **Post** /v1/paas/org/{orgId}/project | Create project
[**PaasDeleteProject**](PaasProjectsAPI.md#PaasDeleteProject) | **Delete** /v1/paas/org/{orgId}/project/{projectId} | Delete project
[**PaasGetProject**](PaasProjectsAPI.md#PaasGetProject) | **Get** /v1/paas/org/{orgId}/project/{projectId} | Get project
[**PaasListProjects**](PaasProjectsAPI.md#PaasListProjects) | **Get** /v1/paas/org/{orgId}/project | List projects
[**PaasUpdateProject**](PaasProjectsAPI.md#PaasUpdateProject) | **Put** /v1/paas/org/{orgId}/project/{projectId} | Update project



## PaasCreateProject

> PaasProject PaasCreateProject(ctx, orgId).PaasCreateOrganizationRequest(paasCreateOrganizationRequest).Execute()

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
	orgId := "orgId_example" // string | 
	paasCreateOrganizationRequest := *openapiclient.NewPaasCreateOrganizationRequest("Name_example") // PaasCreateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasProjectsAPI.PaasCreateProject(context.Background(), orgId).PaasCreateOrganizationRequest(paasCreateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasProjectsAPI.PaasCreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasCreateProject`: PaasProject
	fmt.Fprintf(os.Stdout, "Response from `PaasProjectsAPI.PaasCreateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paasCreateOrganizationRequest** | [**PaasCreateOrganizationRequest**](PaasCreateOrganizationRequest.md) |  | 

### Return type

[**PaasProject**](PaasProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasDeleteProject

> map[string]interface{} PaasDeleteProject(ctx, orgId, projectId).Execute()

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasProjectsAPI.PaasDeleteProject(context.Background(), orgId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasProjectsAPI.PaasDeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeleteProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasProjectsAPI.PaasDeleteProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeleteProjectRequest struct via the builder pattern


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


## PaasGetProject

> PaasProject PaasGetProject(ctx, orgId, projectId).Execute()

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasProjectsAPI.PaasGetProject(context.Background(), orgId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasProjectsAPI.PaasGetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetProject`: PaasProject
	fmt.Fprintf(os.Stdout, "Response from `PaasProjectsAPI.PaasGetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaasProject**](PaasProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListProjects

> []PaasProject PaasListProjects(ctx, orgId).Execute()

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
	orgId := "orgId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasProjectsAPI.PaasListProjects(context.Background(), orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasProjectsAPI.PaasListProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListProjects`: []PaasProject
	fmt.Fprintf(os.Stdout, "Response from `PaasProjectsAPI.PaasListProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PaasProject**](PaasProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasUpdateProject

> map[string]interface{} PaasUpdateProject(ctx, orgId, projectId).PaasUpdateOrganizationRequest(paasUpdateOrganizationRequest).Execute()

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	paasUpdateOrganizationRequest := *openapiclient.NewPaasUpdateOrganizationRequest() // PaasUpdateOrganizationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasProjectsAPI.PaasUpdateProject(context.Background(), orgId, projectId).PaasUpdateOrganizationRequest(paasUpdateOrganizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasProjectsAPI.PaasUpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasUpdateProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasProjectsAPI.PaasUpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **paasUpdateOrganizationRequest** | [**PaasUpdateOrganizationRequest**](PaasUpdateOrganizationRequest.md) |  | 

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

