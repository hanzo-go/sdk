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
[**GetProjectsBySlugReleases**](ProjectsAPI.md#GetProjectsBySlugReleases) | **Get** /v1/projects/{slug}/releases | Returns a site&#39;s releases newest-first, marking the active one — the rollback menu.
[**GetProjectsBySlugShot**](ProjectsAPI.md#GetProjectsBySlugShot) | **Get** /v1/projects/{slug}/shot | Get a PNG of the project&#39;s live site
[**GetProjectsEdge**](ProjectsAPI.md#GetProjectsEdge) | **Get** /v1/projects/edge | health reports whether a publish reaches readers, rather than whether it was accepted.
[**GetProjectsSites**](ProjectsAPI.md#GetProjectsSites) | **Get** /v1/projects/sites | Returns the org&#39;s deployed sites at the pretty URLs they serve at.
[**GetProjectsSitesBySlug**](ProjectsAPI.md#GetProjectsSitesBySlug) | **Get** /v1/projects/sites/{slug} | Returns one site — the same row ListSites carries, for one slug.
[**GetProjectsTags**](ProjectsAPI.md#GetProjectsTags) | **Get** /v1/projects/tags | The site&#39;s browser tag set for the hosted tag — which pixels to inject, by publishable key
[**PatchProjectsBySlug**](ProjectsAPI.md#PatchProjectsBySlug) | **Patch** /v1/projects/{slug} | Changes a project&#39;s settings, and only the settings you send.
[**PostProjects**](ProjectsAPI.md#PostProjects) | **Post** /v1/projects | Creates a project — the handle a site is deployed and served under — and answers 201 with it in &#x60;draft&#x60;.
[**PostProjectsBySlugDeploy**](ProjectsAPI.md#PostProjectsBySlugDeploy) | **Post** /v1/projects/{slug}/deploy | Upload a built site as one archive and serve it
[**PostProjectsBySlugDeployments**](ProjectsAPI.md#PostProjectsBySlugDeployments) | **Post** /v1/projects/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostProjectsBySlugDeploymentsByIdComplete**](ProjectsAPI.md#PostProjectsBySlugDeploymentsByIdComplete) | **Post** /v1/projects/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostProjectsBySlugDomains**](ProjectsAPI.md#PostProjectsBySlugDomains) | **Post** /v1/projects/{slug}/domains | Attaches one or more CUSTOM public hostnames to this org&#39;s site.
[**PostProjectsBySlugDomainsByHostVerify**](ProjectsAPI.md#PostProjectsBySlugDomainsByHostVerify) | **Post** /v1/projects/{slug}/domains/{host}/verify | Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.
[**PostProjectsBySlugPublish**](ProjectsAPI.md#PostProjectsBySlugPublish) | **Post** /v1/projects/{slug}/publish | Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.
[**PostProjectsBySlugPurge**](ProjectsAPI.md#PostProjectsBySlugPurge) | **Post** /v1/projects/{slug}/purge | Flushes the site&#39;s edge cache without redeploying anything.
[**PostProjectsBySlugReleases**](ProjectsAPI.md#PostProjectsBySlugReleases) | **Post** /v1/projects/{slug}/releases | Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.
[**PostProjectsBySlugReleasesByReleaseActivate**](ProjectsAPI.md#PostProjectsBySlugReleasesByReleaseActivate) | **Post** /v1/projects/{slug}/releases/{release}/activate | Points the site at an existing release — the go-live, and equally the ROLLBACK.
[**PostProjectsFork**](ProjectsAPI.md#PostProjectsFork) | **Post** /v1/projects/fork | Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org&#39;s app serving at &lt;slug&gt;.hanzo.app).
[**PostProjectsSites**](ProjectsAPI.md#PostProjectsSites) | **Post** /v1/projects/sites | Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.
[**PostProjectsSitesDeploy**](ProjectsAPI.md#PostProjectsSitesDeploy) | **Post** /v1/projects/sites/deploy | Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.



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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugReleases

> []ProjectsRelease GetProjectsBySlugReleases(ctx, slug).Execute()

Returns a site's releases newest-first, marking the active one — the rollback menu.



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugReleases`: []ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectsRelease**](ProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugShot

> GetProjectsBySlugShot(ctx, slug).Execute()

Get a PNG of the project's live site



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.GetProjectsBySlugShot(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsBySlugShot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsBySlugShotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsEdge

> EdgeState GetProjectsEdge(ctx).Execute()

health reports whether a publish reaches readers, rather than whether it was accepted.



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsEdge(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsEdge`: EdgeState
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsEdge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsEdgeRequest struct via the builder pattern


### Return type

[**EdgeState**](EdgeState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsSites

> []ProjectsSite GetProjectsSites(ctx).Execute()

Returns the org's deployed sites at the pretty URLs they serve at.



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsSites`: []ProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsSitesRequest struct via the builder pattern


### Return type

[**[]ProjectsSite**](ProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsSitesBySlug

> ProjectsSite GetProjectsSitesBySlug(ctx, slug).Execute()

Returns one site — the same row ListSites carries, for one slug.



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsSitesBySlug`: ProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsSitesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsSitesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectsSite**](ProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsTags

> TagConfig GetProjectsTags(ctx).Execute()

The site's browser tag set for the hosted tag — which pixels to inject, by publishable key



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
	resp, r, err := apiClient.ProjectsAPI.GetProjectsTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsTags`: TagConfig
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsTagsRequest struct via the builder pattern


### Return type

[**TagConfig**](TagConfig.md)

### Authorization

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugPublish

> ProjectsRelease PostProjectsBySlugPublish(ctx, slug).ProjectsPublish(projectsPublish).Execute()

Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.



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
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectsPublish := *openapiclient.NewProjectsPublish() // ProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugPublish(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugPublish`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsPublish** | [**ProjectsPublish**](ProjectsPublish.md) |  | 

### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugReleases

> ProjectsRelease PostProjectsBySlugReleases(ctx, slug).ProjectsPublish(projectsPublish).Execute()

Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.



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
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectsPublish := *openapiclient.NewProjectsPublish() // ProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugReleases(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugReleases`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsPublish** | [**ProjectsPublish**](ProjectsPublish.md) |  | 

### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugReleasesByReleaseActivate

> ProjectsRelease PostProjectsBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()

Points the site at an existing release — the go-live, and equally the ROLLBACK.



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
	slug := "slug_example" // string | Slug is the site the release belongs to, from the path.
	release := "release_example" // string | Release is the content-addressed release id (\"rel_\" + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugReleasesByReleaseActivate`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsBySlugReleasesByReleaseActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the release belongs to, from the path. | 
**release** | **string** | Release is the content-addressed release id (\&quot;rel_\&quot; + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsBySlugReleasesByReleaseActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsSites

> ProjectsSiteDeploy PostProjectsSites(ctx).ProjectsBuildSite(projectsBuildSite).Execute()

Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.



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
	projectsBuildSite := *openapiclient.NewProjectsBuildSite() // ProjectsBuildSite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsSites(context.Background()).ProjectsBuildSite(projectsBuildSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsSites`: ProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsBuildSite** | [**ProjectsBuildSite**](ProjectsBuildSite.md) |  | 

### Return type

[**ProjectsSiteDeploy**](ProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsSitesDeploy

> ProjectsSiteDeploy PostProjectsSitesDeploy(ctx).ProjectsDeploySite(projectsDeploySite).Execute()

Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.



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
	projectsDeploySite := *openapiclient.NewProjectsDeploySite() // ProjectsDeploySite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.PostProjectsSitesDeploy(context.Background()).ProjectsDeploySite(projectsDeploySite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.PostProjectsSitesDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsSitesDeploy`: ProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.PostProjectsSitesDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsSitesDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsDeploySite** | [**ProjectsDeploySite**](ProjectsDeploySite.md) |  | 

### Return type

[**ProjectsSiteDeploy**](ProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

