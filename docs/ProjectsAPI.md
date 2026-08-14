# \ProjectsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectsBySlug**](ProjectsAPI.md#DeleteProjectsBySlug) | **Delete** /v1/projects/{slug} | Deletes a project and takes its site off the internet.
[**DeleteProjectsBySlugDomainsByHost**](ProjectsAPI.md#DeleteProjectsBySlugDomainsByHost) | **Delete** /v1/projects/{slug}/domains/{host} | Gives a custom hostname back, so the name is free to reuse.
[**GetProjects**](ProjectsAPI.md#GetProjects) | **Get** /v1/projects | Returns every project your org owns.
[**GetProjectsBySlug**](ProjectsAPI.md#GetProjectsBySlug) | **Get** /v1/projects/{slug} | Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.
[**GetProjectsBySlugDeployments**](ProjectsAPI.md#GetProjectsBySlugDeployments) | **Get** /v1/projects/{slug}/deployments | Returns a project&#39;s deploy history, newest version first.
[**GetProjectsBySlugDeploymentsById**](ProjectsAPI.md#GetProjectsBySlugDeploymentsById) | **Get** /v1/projects/{slug}/deployments/{id} | Returns one deployment of a project by id.
[**GetProjectsBySlugDomains**](ProjectsAPI.md#GetProjectsBySlugDomains) | **Get** /v1/projects/{slug}/domains | Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.
[**PatchProjectsBySlug**](ProjectsAPI.md#PatchProjectsBySlug) | **Patch** /v1/projects/{slug} | Changes a project&#39;s settings, and only the settings you send.
[**PostProjects**](ProjectsAPI.md#PostProjects) | **Post** /v1/projects | Creates a project — the handle a site is deployed and served under — and answers 201 with it in &#x60;draft&#x60;.
[**PostProjectsBySlugDeploy**](ProjectsAPI.md#PostProjectsBySlugDeploy) | **Post** /v1/projects/{slug}/deploy | Upload a built site as one archive and serve it
[**PostProjectsBySlugDeployments**](ProjectsAPI.md#PostProjectsBySlugDeployments) | **Post** /v1/projects/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostProjectsBySlugDeploymentsByIdComplete**](ProjectsAPI.md#PostProjectsBySlugDeploymentsByIdComplete) | **Post** /v1/projects/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostProjectsBySlugDomains**](ProjectsAPI.md#PostProjectsBySlugDomains) | **Post** /v1/projects/{slug}/domains | Attaches one or more CUSTOM public hostnames to this org&#39;s site.
[**PostProjectsBySlugDomainsByHostVerify**](ProjectsAPI.md#PostProjectsBySlugDomainsByHostVerify) | **Post** /v1/projects/{slug}/domains/{host}/verify | Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.
[**PostProjectsBySlugPurge**](ProjectsAPI.md#PostProjectsBySlugPurge) | **Post** /v1/projects/{slug}/purge | Flushes the site&#39;s edge cache without redeploying anything.
[**PostProjectsFork**](ProjectsAPI.md#PostProjectsFork) | **Post** /v1/projects/fork | Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org&#39;s app serving at &lt;slug&gt;.hanzo.app).



## DeleteProjectsBySlug

> DeleteProjectsBySlug(ctx, slug).Execute()

Deletes a project and takes its site off the internet.



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
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProjectsBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectsBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectsBySlugDomainsByHost

> DeleteProjectsBySlugDomainsByHost(ctx, slug, host).Execute()

Gives a custom hostname back, so the name is free to reuse.



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
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProjectsBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectsBySlugDomainsByHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the host is attached to, from the path. | 
**host** | **string** | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsBySlugDomainsByHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> []ProjectsProject GetProjects(ctx).Execute()

Returns every project your org owns.



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
	resp, r, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: []ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


### Return type

[**[]ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlug

> ProjectsProject GetProjectsBySlug(ctx, slug).Execute()

Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.



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
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlug`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDeployments

> []ProjectsDeployment GetProjectsBySlugDeployments(ctx, slug).Execute()

Returns a project's deploy history, newest version first.



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
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDeployments`: []ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDeploymentsById

> ProjectsDeployment GetProjectsBySlugDeploymentsById(ctx, slug, id).Execute()

Returns one deployment of a project by id.



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
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the deployment id, from the path. A deployment of another project — or of another tenant's project — is not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDeploymentsById`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsBySlugDeploymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the deployment id, from the path. A deployment of another project — or of another tenant&#39;s project — is not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugDeploymentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDomains

> ProjectsDomains GetProjectsBySlugDomains(ctx, slug).Execute()

Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.



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
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDomains`: ProjectsDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsDomains**](ProjectsDomains.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectsBySlug

> ProjectsProject PatchProjectsBySlug(ctx, slug).ProjectsUpdate(projectsUpdate).Execute()

Changes a project's settings, and only the settings you send.



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
	slug := "slug_example" // string | Slug is the project to update, from the path. The URL is the addressing authority — a `slug` in the body cannot move the write to another project.
	projectsUpdate := *openapiclient.NewProjectsUpdate() // ProjectsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PatchProjectsBySlug(context.Background(), slug).ProjectsUpdate(projectsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PatchProjectsBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectsBySlug`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PatchProjectsBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to update, from the path. The URL is the addressing authority — a &#x60;slug&#x60; in the body cannot move the write to another project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectsBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsUpdate** | [**ProjectsUpdate**](ProjectsUpdate.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjects

> ProjectsProject PostProjects(ctx).ProjectsCreate(projectsCreate).Execute()

Creates a project — the handle a site is deployed and served under — and answers 201 with it in `draft`.



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
	projectsCreate := *openapiclient.NewProjectsCreate() // ProjectsCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjects(context.Background()).ProjectsCreate(projectsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjects`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsCreate** | [**ProjectsCreate**](ProjectsCreate.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDeploy

> ProjectsDeployment PostProjectsBySlugDeploy(ctx, slug).Body(body).Execute()

Upload a built site as one archive and serve it



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
	slug := "slug_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugDeploy(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeploy`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** |  | 

### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDeployments

> ProjectsDeployment PostProjectsBySlugDeployments(ctx, slug).ProjectsDeployStart(projectsDeployStart).Execute()

Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.



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
	slug := "slug_example" // string | Slug is the site to deploy, from the path.
	projectsDeployStart := *openapiclient.NewProjectsDeployStart() // ProjectsDeployStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugDeployments(context.Background(), slug).ProjectsDeployStart(projectsDeployStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeployments`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to deploy, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsDeployStart** | [**ProjectsDeployStart**](ProjectsDeployStart.md) |  | 

### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDeploymentsByIdComplete

> ProjectsDeployment PostProjectsBySlugDeploymentsByIdComplete(ctx, slug, id).ProjectsComplete(projectsComplete).Execute()

CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.



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
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the queued deployment to complete, from the path.
	projectsComplete := *openapiclient.NewProjectsComplete() // ProjectsComplete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugDeploymentsByIdComplete(context.Background(), slug, id).ProjectsComplete(projectsComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugDeploymentsByIdComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeploymentsByIdComplete`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugDeploymentsByIdComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the queued deployment to complete, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugDeploymentsByIdCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectsComplete** | [**ProjectsComplete**](ProjectsComplete.md) |  | 

### Return type

[**ProjectsDeployment**](ProjectsDeployment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDomains

> ProjectsBoundDomains PostProjectsBySlugDomains(ctx, slug).ProjectsDomainsBind(projectsDomainsBind).Execute()

Attaches one or more CUSTOM public hostnames to this org's site.



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
	slug := "slug_example" // string | Slug is the site the hosts attach to, from the path.
	projectsDomainsBind := *openapiclient.NewProjectsDomainsBind() // ProjectsDomainsBind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugDomains(context.Background(), slug).ProjectsDomainsBind(projectsDomainsBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDomains`: ProjectsBoundDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the hosts attach to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsDomainsBind** | [**ProjectsDomainsBind**](ProjectsDomainsBind.md) |  | 

### Return type

[**ProjectsBoundDomains**](ProjectsBoundDomains.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDomainsByHostVerify

> ProjectsDomain PostProjectsBySlugDomainsByHostVerify(ctx, slug, host).Execute()

Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.



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
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDomainsByHostVerify`: ProjectsDomain
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugDomainsByHostVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the host is attached to, from the path. | 
**host** | **string** | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugDomainsByHostVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsDomain**](ProjectsDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugPurge

> ProjectsProject PostProjectsBySlugPurge(ctx, slug).Execute()

Flushes the site's edge cache without redeploying anything.



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
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugPurge`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsFork

> ProjectsProject PostProjectsFork(ctx).ProjectsFork(projectsFork).Execute()

Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org's app serving at <slug>.hanzo.app).



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
	projectsFork := *openapiclient.NewProjectsFork() // ProjectsFork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsFork(context.Background()).ProjectsFork(projectsFork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsFork`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsFork** | [**ProjectsFork**](ProjectsFork.md) |  | 

### Return type

[**ProjectsProject**](ProjectsProject.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

