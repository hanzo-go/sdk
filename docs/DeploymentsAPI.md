# \DeploymentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MlDeployModel**](DeploymentsAPI.md#MlDeployModel) | **Post** /v1/ml/deploy | Deploy a model
[**MlGetDeployment**](DeploymentsAPI.md#MlGetDeployment) | **Get** /v1/ml/deployments/{deployment_id} | Get deployment details
[**MlListDeployments**](DeploymentsAPI.md#MlListDeployments) | **Get** /v1/ml/deployments | List deployments
[**MlStopDeployment**](DeploymentsAPI.md#MlStopDeployment) | **Delete** /v1/ml/deployments/{deployment_id} | Stop a deployment
[**ProjectsCompleteDeployment**](DeploymentsAPI.md#ProjectsCompleteDeployment) | **Post** /v1/projects/{slug}/deployments/{id}/complete | CI completion hook for a git deployment
[**ProjectsDeployProject**](DeploymentsAPI.md#ProjectsDeployProject) | **Post** /v1/projects/{slug}/deploy | Deploy a project to the S3 origin
[**ProjectsGetDeployment**](DeploymentsAPI.md#ProjectsGetDeployment) | **Get** /v1/projects/{slug}/deployments/{id} | Get a deployment
[**ProjectsListDeployments**](DeploymentsAPI.md#ProjectsListDeployments) | **Get** /v1/projects/{slug}/deployments | List deployments for a project
[**ProjectsPurgeProject**](DeploymentsAPI.md#ProjectsPurgeProject) | **Post** /v1/projects/{slug}/purge | Purge the site&#39;s edge cache



## MlDeployModel

> MlDeployment MlDeployModel(ctx).MlDeployModelRequest(mlDeployModelRequest).Execute()

Deploy a model



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
	mlDeployModelRequest := *openapiclient.NewMlDeployModelRequest("ModelId_example") // MlDeployModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.MlDeployModel(context.Background()).MlDeployModelRequest(mlDeployModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.MlDeployModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlDeployModel`: MlDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.MlDeployModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlDeployModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mlDeployModelRequest** | [**MlDeployModelRequest**](MlDeployModelRequest.md) |  | 

### Return type

[**MlDeployment**](MlDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlGetDeployment

> MlDeployment MlGetDeployment(ctx, deploymentId).Execute()

Get deployment details

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
	deploymentId := "deploymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.MlGetDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.MlGetDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetDeployment`: MlDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.MlGetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlDeployment**](MlDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListDeployments

> MlListDeployments200Response MlListDeployments(ctx).Environment(environment).Status(status).Execute()

List deployments

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
	environment := "environment_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.MlListDeployments(context.Background()).Environment(environment).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.MlListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListDeployments`: MlListDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.MlListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**MlListDeployments200Response**](MlListDeployments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlStopDeployment

> MlStopDeployment(ctx, deploymentId).Execute()

Stop a deployment

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
	deploymentId := "deploymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentsAPI.MlStopDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.MlStopDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlStopDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCompleteDeployment

> ProjectsDeployment ProjectsCompleteDeployment(ctx, slug, id).ProjectsCompleteDeploymentRequest(projectsCompleteDeploymentRequest).Execute()

CI completion hook for a git deployment



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
	id := "id_example" // string | Deployment id (e.g. dep_...).
	projectsCompleteDeploymentRequest := *openapiclient.NewProjectsCompleteDeploymentRequest("Status_example") // ProjectsCompleteDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.ProjectsCompleteDeployment(context.Background(), slug, id).ProjectsCompleteDeploymentRequest(projectsCompleteDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.ProjectsCompleteDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsCompleteDeployment`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.ProjectsCompleteDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 
**id** | **string** | Deployment id (e.g. dep_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCompleteDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectsCompleteDeploymentRequest** | [**ProjectsCompleteDeploymentRequest**](ProjectsCompleteDeploymentRequest.md) |  | 

### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDeployProject

> ProjectsDeployment ProjectsDeployProject(ctx, slug).ProjectsGitDeployRequest(projectsGitDeployRequest).Execute()

Deploy a project to the S3 origin



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
	projectsGitDeployRequest := *openapiclient.NewProjectsGitDeployRequest() // ProjectsGitDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.ProjectsDeployProject(context.Background(), slug).ProjectsGitDeployRequest(projectsGitDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.ProjectsDeployProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsDeployProject`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.ProjectsDeployProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeployProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsGitDeployRequest** | [**ProjectsGitDeployRequest**](ProjectsGitDeployRequest.md) |  | 

### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGetDeployment

> ProjectsDeployment ProjectsGetDeployment(ctx, slug, id).Execute()

Get a deployment



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
	id := "id_example" // string | Deployment id (e.g. dep_...).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.ProjectsGetDeployment(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.ProjectsGetDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGetDeployment`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.ProjectsGetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 
**id** | **string** | Deployment id (e.g. dep_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListDeployments

> []ProjectsDeployment ProjectsListDeployments(ctx, slug).Execute()

List deployments for a project



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
	resp, r, err := apiClient.DeploymentsAPI.ProjectsListDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.ProjectsListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsListDeployments`: []ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.ProjectsListDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPurgeProject

> ProjectsProject ProjectsPurgeProject(ctx, slug).Execute()

Purge the site's edge cache



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
	resp, r, err := apiClient.DeploymentsAPI.ProjectsPurgeProject(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.ProjectsPurgeProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsPurgeProject`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.ProjectsPurgeProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Org-unique project handle (lowercased); also the S3-origin key segment and the subdomain label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPurgeProjectRequest struct via the builder pattern


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

