# \ProjectAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectBySlug**](ProjectAPI.md#DeleteProjectBySlug) | **Delete** /v1/project/{slug} | Deletes a project and takes its site off the internet.
[**DeleteProjectBySlugDomainsByHost**](ProjectAPI.md#DeleteProjectBySlugDomainsByHost) | **Delete** /v1/project/{slug}/domains/{host} | Gives a custom hostname back, so the name is free to reuse.
[**DeleteProjectBySlugStar**](ProjectAPI.md#DeleteProjectBySlugStar) | **Delete** /v1/project/{slug}/star | Removes the caller&#39;s own bookmark from a project, and answers whether it is starred afterwards.
[**DeleteProjectsBySlug**](ProjectAPI.md#DeleteProjectsBySlug) | **Delete** /v1/projects/{slug} | Deletes a project and takes its site off the internet.
[**DeleteProjectsBySlugDomainsByHost**](ProjectAPI.md#DeleteProjectsBySlugDomainsByHost) | **Delete** /v1/projects/{slug}/domains/{host} | Gives a custom hostname back, so the name is free to reuse.
[**DeleteProjectsBySlugStar**](ProjectAPI.md#DeleteProjectsBySlugStar) | **Delete** /v1/projects/{slug}/star | Removes the caller&#39;s own bookmark from a project, and answers whether it is starred afterwards.
[**GetProject**](ProjectAPI.md#GetProject) | **Get** /v1/project | Returns every project your org owns.
[**GetProjectBySlug**](ProjectAPI.md#GetProjectBySlug) | **Get** /v1/project/{slug} | Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.
[**GetProjectBySlugDeployments**](ProjectAPI.md#GetProjectBySlugDeployments) | **Get** /v1/project/{slug}/deployments | Returns a project&#39;s deploy history, newest version first.
[**GetProjectBySlugDeploymentsById**](ProjectAPI.md#GetProjectBySlugDeploymentsById) | **Get** /v1/project/{slug}/deployments/{id} | Returns one deployment of a project by id.
[**GetProjectBySlugDomains**](ProjectAPI.md#GetProjectBySlugDomains) | **Get** /v1/project/{slug}/domains | Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.
[**GetProjectBySlugReleases**](ProjectAPI.md#GetProjectBySlugReleases) | **Get** /v1/project/{slug}/releases | Returns a site&#39;s releases newest-first, marking the active one — the rollback menu.
[**GetProjectBySlugShot**](ProjectAPI.md#GetProjectBySlugShot) | **Get** /v1/project/{slug}/shot | Get a PNG of the project&#39;s live site
[**GetProjectEdge**](ProjectAPI.md#GetProjectEdge) | **Get** /v1/project/edge | health reports whether a publish reaches readers, rather than whether it was accepted.
[**GetProjectSites**](ProjectAPI.md#GetProjectSites) | **Get** /v1/project/sites | Returns the org&#39;s deployed sites at the pretty URLs they serve at.
[**GetProjectSitesBySlug**](ProjectAPI.md#GetProjectSitesBySlug) | **Get** /v1/project/sites/{slug} | Returns one site — the same row ListSites carries, for one slug.
[**GetProjectTags**](ProjectAPI.md#GetProjectTags) | **Get** /v1/project/tags | The site&#39;s browser tag set for the hosted tag — which pixels to inject, by publishable key
[**GetProjects**](ProjectAPI.md#GetProjects) | **Get** /v1/projects | Returns every project your org owns.
[**GetProjectsBySlug**](ProjectAPI.md#GetProjectsBySlug) | **Get** /v1/projects/{slug} | Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.
[**GetProjectsBySlugDeployments**](ProjectAPI.md#GetProjectsBySlugDeployments) | **Get** /v1/projects/{slug}/deployments | Returns a project&#39;s deploy history, newest version first.
[**GetProjectsBySlugDeploymentsById**](ProjectAPI.md#GetProjectsBySlugDeploymentsById) | **Get** /v1/projects/{slug}/deployments/{id} | Returns one deployment of a project by id.
[**GetProjectsBySlugDomains**](ProjectAPI.md#GetProjectsBySlugDomains) | **Get** /v1/projects/{slug}/domains | Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.
[**GetProjectsBySlugReleases**](ProjectAPI.md#GetProjectsBySlugReleases) | **Get** /v1/projects/{slug}/releases | Returns a site&#39;s releases newest-first, marking the active one — the rollback menu.
[**GetProjectsBySlugShot**](ProjectAPI.md#GetProjectsBySlugShot) | **Get** /v1/projects/{slug}/shot | Get a PNG of the project&#39;s live site
[**GetProjectsEdge**](ProjectAPI.md#GetProjectsEdge) | **Get** /v1/projects/edge | health reports whether a publish reaches readers, rather than whether it was accepted.
[**GetProjectsSites**](ProjectAPI.md#GetProjectsSites) | **Get** /v1/projects/sites | Returns the org&#39;s deployed sites at the pretty URLs they serve at.
[**GetProjectsSitesBySlug**](ProjectAPI.md#GetProjectsSitesBySlug) | **Get** /v1/projects/sites/{slug} | Returns one site — the same row ListSites carries, for one slug.
[**PatchProjectBySlug**](ProjectAPI.md#PatchProjectBySlug) | **Patch** /v1/project/{slug} | Changes a project&#39;s settings, and only the settings you send.
[**PatchProjectsBySlug**](ProjectAPI.md#PatchProjectsBySlug) | **Patch** /v1/projects/{slug} | Changes a project&#39;s settings, and only the settings you send.
[**PostProject**](ProjectAPI.md#PostProject) | **Post** /v1/project | Creates a project — the handle a site is deployed and served under — and answers 201 with it in &#x60;draft&#x60;.
[**PostProjectBySlugDeploy**](ProjectAPI.md#PostProjectBySlugDeploy) | **Post** /v1/project/{slug}/deploy | Upload a built site as one archive and serve it
[**PostProjectBySlugDeployments**](ProjectAPI.md#PostProjectBySlugDeployments) | **Post** /v1/project/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostProjectBySlugDeploymentsByIdComplete**](ProjectAPI.md#PostProjectBySlugDeploymentsByIdComplete) | **Post** /v1/project/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostProjectBySlugDomains**](ProjectAPI.md#PostProjectBySlugDomains) | **Post** /v1/project/{slug}/domains | Attaches one or more CUSTOM public hostnames to this org&#39;s site.
[**PostProjectBySlugDomainsByHostVerify**](ProjectAPI.md#PostProjectBySlugDomainsByHostVerify) | **Post** /v1/project/{slug}/domains/{host}/verify | Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.
[**PostProjectBySlugPublish**](ProjectAPI.md#PostProjectBySlugPublish) | **Post** /v1/project/{slug}/publish | Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.
[**PostProjectBySlugPurge**](ProjectAPI.md#PostProjectBySlugPurge) | **Post** /v1/project/{slug}/purge | Flushes the site&#39;s edge cache without redeploying anything.
[**PostProjectBySlugReleases**](ProjectAPI.md#PostProjectBySlugReleases) | **Post** /v1/project/{slug}/releases | Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.
[**PostProjectBySlugReleasesByReleaseActivate**](ProjectAPI.md#PostProjectBySlugReleasesByReleaseActivate) | **Post** /v1/project/{slug}/releases/{release}/activate | Points the site at an existing release — the go-live, and equally the ROLLBACK.
[**PostProjectFork**](ProjectAPI.md#PostProjectFork) | **Post** /v1/project/fork | Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org&#39;s app serving at &lt;slug&gt;.hanzo.app).
[**PostProjectSites**](ProjectAPI.md#PostProjectSites) | **Post** /v1/project/sites | Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.
[**PostProjectSitesDeploy**](ProjectAPI.md#PostProjectSitesDeploy) | **Post** /v1/project/sites/deploy | Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.
[**PostProjects**](ProjectAPI.md#PostProjects) | **Post** /v1/projects | Creates a project — the handle a site is deployed and served under — and answers 201 with it in &#x60;draft&#x60;.
[**PostProjectsBySlugDeploy**](ProjectAPI.md#PostProjectsBySlugDeploy) | **Post** /v1/projects/{slug}/deploy | Upload a built site as one archive and serve it
[**PostProjectsBySlugDeployments**](ProjectAPI.md#PostProjectsBySlugDeployments) | **Post** /v1/projects/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostProjectsBySlugDeploymentsByIdComplete**](ProjectAPI.md#PostProjectsBySlugDeploymentsByIdComplete) | **Post** /v1/projects/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostProjectsBySlugDomains**](ProjectAPI.md#PostProjectsBySlugDomains) | **Post** /v1/projects/{slug}/domains | Attaches one or more CUSTOM public hostnames to this org&#39;s site.
[**PostProjectsBySlugDomainsByHostVerify**](ProjectAPI.md#PostProjectsBySlugDomainsByHostVerify) | **Post** /v1/projects/{slug}/domains/{host}/verify | Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.
[**PostProjectsBySlugPublish**](ProjectAPI.md#PostProjectsBySlugPublish) | **Post** /v1/projects/{slug}/publish | Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.
[**PostProjectsBySlugPurge**](ProjectAPI.md#PostProjectsBySlugPurge) | **Post** /v1/projects/{slug}/purge | Flushes the site&#39;s edge cache without redeploying anything.
[**PostProjectsBySlugReleases**](ProjectAPI.md#PostProjectsBySlugReleases) | **Post** /v1/projects/{slug}/releases | Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.
[**PostProjectsBySlugReleasesByReleaseActivate**](ProjectAPI.md#PostProjectsBySlugReleasesByReleaseActivate) | **Post** /v1/projects/{slug}/releases/{release}/activate | Points the site at an existing release — the go-live, and equally the ROLLBACK.
[**PostProjectsFork**](ProjectAPI.md#PostProjectsFork) | **Post** /v1/projects/fork | Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org&#39;s app serving at &lt;slug&gt;.hanzo.app).
[**PostProjectsSites**](ProjectAPI.md#PostProjectsSites) | **Post** /v1/projects/sites | Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.
[**PostProjectsSitesDeploy**](ProjectAPI.md#PostProjectsSitesDeploy) | **Post** /v1/projects/sites/deploy | Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.
[**PutProjectBySlugStar**](ProjectAPI.md#PutProjectBySlugStar) | **Put** /v1/project/{slug}/star | Bookmarks a project for the person calling, and answers whether it is starred afterwards.
[**PutProjectsBySlugStar**](ProjectAPI.md#PutProjectsBySlugStar) | **Put** /v1/projects/{slug}/star | Bookmarks a project for the person calling, and answers whether it is starred afterwards.



## DeleteProjectBySlug

> DeleteProjectBySlug(ctx, slug).Execute()

Deletes a project and takes its site off the internet.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteProjectBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProjectBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectBySlugDomainsByHost

> DeleteProjectBySlugDomainsByHost(ctx, slug, host).Execute()

Gives a custom hostname back, so the name is free to reuse.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteProjectBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectBySlugDomainsByHost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProjectBySlugDomainsByHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectBySlugStar

> ProjectProjectsStar DeleteProjectBySlugStar(ctx, slug).Execute()

Removes the caller's own bookmark from a project, and answers whether it is starred afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.DeleteProjectBySlugStar(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectBySlugStar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectBySlugStar`: ProjectProjectsStar
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DeleteProjectBySlugStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectBySlugStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsStar**](ProjectProjectsStar.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteProjectsBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectsBySlug``: %v\n", err)
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
- **Accept**: application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.DeleteProjectsBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectsBySlugDomainsByHost``: %v\n", err)
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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectsBySlugStar

> ProjectProjectsStar DeleteProjectsBySlugStar(ctx, slug).Execute()

Removes the caller's own bookmark from a project, and answers whether it is starred afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.DeleteProjectsBySlugStar(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.DeleteProjectsBySlugStar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectsBySlugStar`: ProjectProjectsStar
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.DeleteProjectsBySlugStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsBySlugStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsStar**](ProjectProjectsStar.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> []ProjectProjectsProject GetProject(ctx).Execute()

Returns every project your org owns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProject(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProject`: []ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


### Return type

[**[]ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlug

> ProjectProjectsProject GetProjectBySlug(ctx, slug).Execute()

Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBySlug`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlugDeployments

> []ProjectProjectsDeployment GetProjectBySlugDeployments(ctx, slug).Execute()

Returns a project's deploy history, newest version first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBySlugDeployments`: []ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlugDeploymentsById

> ProjectProjectsDeployment GetProjectBySlugDeploymentsById(ctx, slug, id).Execute()

Returns one deployment of a project by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the deployment id, from the path. A deployment of another project — or of another tenant's project — is not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBySlugDeploymentsById`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectBySlugDeploymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the deployment id, from the path. A deployment of another project — or of another tenant&#39;s project — is not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugDeploymentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlugDomains

> ProjectProjectsDomains GetProjectBySlugDomains(ctx, slug).Execute()

Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBySlugDomains`: ProjectProjectsDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsDomains**](ProjectProjectsDomains.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlugReleases

> []ProjectProjectsRelease GetProjectBySlugReleases(ctx, slug).Execute()

Returns a site's releases newest-first, marking the active one — the rollback menu.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectBySlugReleases`: []ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectBySlugShot

> GetProjectBySlugShot(ctx, slug).Execute()

Get a PNG of the project's live site



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.GetProjectBySlugShot(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectBySlugShot``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetProjectBySlugShotRequest struct via the builder pattern


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


## GetProjectEdge

> ProjectEdgeState GetProjectEdge(ctx).Execute()

health reports whether a publish reaches readers, rather than whether it was accepted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectEdge(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEdge`: ProjectEdgeState
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectEdge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEdgeRequest struct via the builder pattern


### Return type

[**ProjectEdgeState**](ProjectEdgeState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSites

> []ProjectProjectsSite GetProjectSites(ctx).Execute()

Returns the org's deployed sites at the pretty URLs they serve at.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSites`: []ProjectProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSitesRequest struct via the builder pattern


### Return type

[**[]ProjectProjectsSite**](ProjectProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectSitesBySlug

> ProjectProjectsSite GetProjectSitesBySlug(ctx, slug).Execute()

Returns one site — the same row ListSites carries, for one slug.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectSitesBySlug`: ProjectProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectSitesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectSitesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsSite**](ProjectProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTags

> TagConfig GetProjectTags(ctx).Execute()

The site's browser tag set for the hosted tag — which pixels to inject, by publishable key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTags`: TagConfig
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTagsRequest struct via the builder pattern


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


## GetProjects

> []ProjectProjectsProject GetProjects(ctx).Execute()

Returns every project your org owns.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: []ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


### Return type

[**[]ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlug

> ProjectProjectsProject GetProjectsBySlug(ctx, slug).Execute()

Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlug`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsBySlug`: %v\n", resp)
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

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDeployments

> []ProjectProjectsDeployment GetProjectsBySlugDeployments(ctx, slug).Execute()

Returns a project's deploy history, newest version first.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDeployments`: []ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsBySlugDeployments`: %v\n", resp)
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

[**[]ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDeploymentsById

> ProjectProjectsDeployment GetProjectsBySlugDeploymentsById(ctx, slug, id).Execute()

Returns one deployment of a project by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the deployment id, from the path. A deployment of another project — or of another tenant's project — is not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDeploymentsById`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsBySlugDeploymentsById`: %v\n", resp)
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

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugDomains

> ProjectProjectsDomains GetProjectsBySlugDomains(ctx, slug).Execute()

Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugDomains`: ProjectProjectsDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsBySlugDomains`: %v\n", resp)
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

[**ProjectProjectsDomains**](ProjectProjectsDomains.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsBySlugReleases

> []ProjectProjectsRelease GetProjectsBySlugReleases(ctx, slug).Execute()

Returns a site's releases newest-first, marking the active one — the rollback menu.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsBySlugReleases`: []ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsBySlugReleases`: %v\n", resp)
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

[**[]ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAPI.GetProjectsBySlugShot(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsBySlugShot``: %v\n", err)
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

> ProjectEdgeState GetProjectsEdge(ctx).Execute()

health reports whether a publish reaches readers, rather than whether it was accepted.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsEdge(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsEdge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsEdge`: ProjectEdgeState
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsEdge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsEdgeRequest struct via the builder pattern


### Return type

[**ProjectEdgeState**](ProjectEdgeState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsSites

> []ProjectProjectsSite GetProjectsSites(ctx).Execute()

Returns the org's deployed sites at the pretty URLs they serve at.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsSites`: []ProjectProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsSitesRequest struct via the builder pattern


### Return type

[**[]ProjectProjectsSite**](ProjectProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsSitesBySlug

> ProjectProjectsSite GetProjectsSitesBySlug(ctx, slug).Execute()

Returns one site — the same row ListSites carries, for one slug.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectsSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectsSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsSitesBySlug`: ProjectProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectsSitesBySlug`: %v\n", resp)
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

[**ProjectProjectsSite**](ProjectProjectsSite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectBySlug

> ProjectProjectsProject PatchProjectBySlug(ctx, slug).ProjectProjectsUpdate(projectProjectsUpdate).Execute()

Changes a project's settings, and only the settings you send.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to update, from the path. The URL is the addressing authority — a `slug` in the body cannot move the write to another project.
	projectProjectsUpdate := *openapiclient.NewProjectProjectsUpdate() // ProjectProjectsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PatchProjectBySlug(context.Background(), slug).ProjectProjectsUpdate(projectProjectsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PatchProjectBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectBySlug`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PatchProjectBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to update, from the path. The URL is the addressing authority — a &#x60;slug&#x60; in the body cannot move the write to another project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProjectBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectsUpdate** | [**ProjectProjectsUpdate**](ProjectProjectsUpdate.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProjectsBySlug

> ProjectProjectsProject PatchProjectsBySlug(ctx, slug).ProjectProjectsUpdate(projectProjectsUpdate).Execute()

Changes a project's settings, and only the settings you send.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to update, from the path. The URL is the addressing authority — a `slug` in the body cannot move the write to another project.
	projectProjectsUpdate := *openapiclient.NewProjectProjectsUpdate() // ProjectProjectsUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PatchProjectsBySlug(context.Background(), slug).ProjectProjectsUpdate(projectProjectsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PatchProjectsBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchProjectsBySlug`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PatchProjectsBySlug`: %v\n", resp)
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

 **projectProjectsUpdate** | [**ProjectProjectsUpdate**](ProjectProjectsUpdate.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProject

> ProjectProjectsProject PostProject(ctx).ProjectProjectsCreate(projectProjectsCreate).Execute()

Creates a project — the handle a site is deployed and served under — and answers 201 with it in `draft`.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsCreate := *openapiclient.NewProjectProjectsCreate() // ProjectProjectsCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProject(context.Background()).ProjectProjectsCreate(projectProjectsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProject`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsCreate** | [**ProjectProjectsCreate**](ProjectProjectsCreate.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugDeploy

> ProjectsDeployment PostProjectBySlugDeploy(ctx, slug).Body(body).Execute()

Upload a built site as one archive and serve it



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugDeploy(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugDeploy`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugDeployRequest struct via the builder pattern


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


## PostProjectBySlugDeployments

> ProjectProjectsDeployment PostProjectBySlugDeployments(ctx, slug).ProjectProjectsDeployStart(projectProjectsDeployStart).Execute()

Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to deploy, from the path.
	projectProjectsDeployStart := *openapiclient.NewProjectProjectsDeployStart() // ProjectProjectsDeployStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugDeployments(context.Background(), slug).ProjectProjectsDeployStart(projectProjectsDeployStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugDeployments`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to deploy, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectsDeployStart** | [**ProjectProjectsDeployStart**](ProjectProjectsDeployStart.md) |  | 

### Return type

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugDeploymentsByIdComplete

> ProjectProjectsDeployment PostProjectBySlugDeploymentsByIdComplete(ctx, slug, id).ProjectProjectsComplete(projectProjectsComplete).Execute()

CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the queued deployment to complete, from the path.
	projectProjectsComplete := *openapiclient.NewProjectProjectsComplete() // ProjectProjectsComplete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugDeploymentsByIdComplete(context.Background(), slug, id).ProjectProjectsComplete(projectProjectsComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugDeploymentsByIdComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugDeploymentsByIdComplete`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugDeploymentsByIdComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the queued deployment to complete, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugDeploymentsByIdCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectProjectsComplete** | [**ProjectProjectsComplete**](ProjectProjectsComplete.md) |  | 

### Return type

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugDomains

> ProjectProjectsBoundDomains PostProjectBySlugDomains(ctx, slug).ProjectProjectsDomainsBind(projectProjectsDomainsBind).Execute()

Attaches one or more CUSTOM public hostnames to this org's site.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site the hosts attach to, from the path.
	projectProjectsDomainsBind := *openapiclient.NewProjectProjectsDomainsBind() // ProjectProjectsDomainsBind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugDomains(context.Background(), slug).ProjectProjectsDomainsBind(projectProjectsDomainsBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugDomains`: ProjectProjectsBoundDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the hosts attach to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectsDomainsBind** | [**ProjectProjectsDomainsBind**](ProjectProjectsDomainsBind.md) |  | 

### Return type

[**ProjectProjectsBoundDomains**](ProjectProjectsBoundDomains.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugDomainsByHostVerify

> ProjectProjectsDomain PostProjectBySlugDomainsByHostVerify(ctx, slug, host).Execute()

Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugDomainsByHostVerify`: ProjectProjectsDomain
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugDomainsByHostVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the host is attached to, from the path. | 
**host** | **string** | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugDomainsByHostVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectProjectsDomain**](ProjectProjectsDomain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugPublish

> ProjectProjectsRelease PostProjectBySlugPublish(ctx, slug).ProjectProjectsPublish(projectProjectsPublish).Execute()

Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectProjectsPublish := *openapiclient.NewProjectProjectsPublish() // ProjectProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugPublish(context.Background(), slug).ProjectProjectsPublish(projectProjectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugPublish`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectsPublish** | [**ProjectProjectsPublish**](ProjectProjectsPublish.md) |  | 

### Return type

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugPurge

> ProjectProjectsProject PostProjectBySlugPurge(ctx, slug).Execute()

Flushes the site's edge cache without redeploying anything.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugPurge`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugReleases

> ProjectProjectsRelease PostProjectBySlugReleases(ctx, slug).ProjectProjectsPublish(projectProjectsPublish).Execute()

Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectProjectsPublish := *openapiclient.NewProjectProjectsPublish() // ProjectProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugReleases(context.Background(), slug).ProjectProjectsPublish(projectProjectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugReleases`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectProjectsPublish** | [**ProjectProjectsPublish**](ProjectProjectsPublish.md) |  | 

### Return type

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectBySlugReleasesByReleaseActivate

> ProjectProjectsRelease PostProjectBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()

Points the site at an existing release — the go-live, and equally the ROLLBACK.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site the release belongs to, from the path.
	release := "release_example" // string | Release is the content-addressed release id (\"rel_\" + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectBySlugReleasesByReleaseActivate`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectBySlugReleasesByReleaseActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the release belongs to, from the path. | 
**release** | **string** | Release is the content-addressed release id (\&quot;rel_\&quot; + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectBySlugReleasesByReleaseActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectFork

> ProjectProjectsProject PostProjectFork(ctx).ProjectProjectsFork(projectProjectsFork).Execute()

Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org's app serving at <slug>.hanzo.app).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsFork := *openapiclient.NewProjectProjectsFork() // ProjectProjectsFork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectFork(context.Background()).ProjectProjectsFork(projectProjectsFork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectFork`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsFork** | [**ProjectProjectsFork**](ProjectProjectsFork.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectSites

> ProjectProjectsSiteDeploy PostProjectSites(ctx).ProjectProjectsBuildSite(projectProjectsBuildSite).Execute()

Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsBuildSite := *openapiclient.NewProjectProjectsBuildSite() // ProjectProjectsBuildSite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectSites(context.Background()).ProjectProjectsBuildSite(projectProjectsBuildSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectSites`: ProjectProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsBuildSite** | [**ProjectProjectsBuildSite**](ProjectProjectsBuildSite.md) |  | 

### Return type

[**ProjectProjectsSiteDeploy**](ProjectProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectSitesDeploy

> ProjectProjectsSiteDeploy PostProjectSitesDeploy(ctx).ProjectProjectsDeploySite(projectProjectsDeploySite).Execute()

Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsDeploySite := *openapiclient.NewProjectProjectsDeploySite() // ProjectProjectsDeploySite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectSitesDeploy(context.Background()).ProjectProjectsDeploySite(projectProjectsDeploySite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectSitesDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectSitesDeploy`: ProjectProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectSitesDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectSitesDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsDeploySite** | [**ProjectProjectsDeploySite**](ProjectProjectsDeploySite.md) |  | 

### Return type

[**ProjectProjectsSiteDeploy**](ProjectProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjects

> ProjectProjectsProject PostProjects(ctx).ProjectProjectsCreate(projectProjectsCreate).Execute()

Creates a project — the handle a site is deployed and served under — and answers 201 with it in `draft`.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsCreate := *openapiclient.NewProjectProjectsCreate() // ProjectProjectsCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjects(context.Background()).ProjectProjectsCreate(projectProjectsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjects`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsCreate** | [**ProjectProjectsCreate**](ProjectProjectsCreate.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugDeploy(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeploy`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugDeploy`: %v\n", resp)
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

> ProjectProjectsDeployment PostProjectsBySlugDeployments(ctx, slug).ProjectProjectsDeployStart(projectProjectsDeployStart).Execute()

Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to deploy, from the path.
	projectProjectsDeployStart := *openapiclient.NewProjectProjectsDeployStart() // ProjectProjectsDeployStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugDeployments(context.Background(), slug).ProjectProjectsDeployStart(projectProjectsDeployStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeployments`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugDeployments`: %v\n", resp)
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

 **projectProjectsDeployStart** | [**ProjectProjectsDeployStart**](ProjectProjectsDeployStart.md) |  | 

### Return type

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDeploymentsByIdComplete

> ProjectProjectsDeployment PostProjectsBySlugDeploymentsByIdComplete(ctx, slug, id).ProjectProjectsComplete(projectProjectsComplete).Execute()

CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the deployment belongs to, from the path.
	id := "id_example" // string | ID is the queued deployment to complete, from the path.
	projectProjectsComplete := *openapiclient.NewProjectProjectsComplete() // ProjectProjectsComplete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugDeploymentsByIdComplete(context.Background(), slug, id).ProjectProjectsComplete(projectProjectsComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugDeploymentsByIdComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDeploymentsByIdComplete`: ProjectProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugDeploymentsByIdComplete`: %v\n", resp)
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


 **projectProjectsComplete** | [**ProjectProjectsComplete**](ProjectProjectsComplete.md) |  | 

### Return type

[**ProjectProjectsDeployment**](ProjectProjectsDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDomains

> ProjectProjectsBoundDomains PostProjectsBySlugDomains(ctx, slug).ProjectProjectsDomainsBind(projectProjectsDomainsBind).Execute()

Attaches one or more CUSTOM public hostnames to this org's site.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site the hosts attach to, from the path.
	projectProjectsDomainsBind := *openapiclient.NewProjectProjectsDomainsBind() // ProjectProjectsDomainsBind | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugDomains(context.Background(), slug).ProjectProjectsDomainsBind(projectProjectsDomainsBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDomains`: ProjectProjectsBoundDomains
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugDomains`: %v\n", resp)
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

 **projectProjectsDomainsBind** | [**ProjectProjectsDomainsBind**](ProjectProjectsDomainsBind.md) |  | 

### Return type

[**ProjectProjectsBoundDomains**](ProjectProjectsBoundDomains.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugDomainsByHostVerify

> ProjectProjectsDomain PostProjectsBySlugDomainsByHostVerify(ctx, slug, host).Execute()

Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project the host is attached to, from the path.
	host := "host_example" // string | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugDomainsByHostVerify`: ProjectProjectsDomain
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugDomainsByHostVerify`: %v\n", resp)
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

[**ProjectProjectsDomain**](ProjectProjectsDomain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugPublish

> ProjectProjectsRelease PostProjectsBySlugPublish(ctx, slug).ProjectProjectsPublish(projectProjectsPublish).Execute()

Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectProjectsPublish := *openapiclient.NewProjectProjectsPublish() // ProjectProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugPublish(context.Background(), slug).ProjectProjectsPublish(projectProjectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugPublish`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugPublish`: %v\n", resp)
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

 **projectProjectsPublish** | [**ProjectProjectsPublish**](ProjectProjectsPublish.md) |  | 

### Return type

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugPurge

> ProjectProjectsProject PostProjectsBySlugPurge(ctx, slug).Execute()

Flushes the site's edge cache without redeploying anything.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugPurge`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugPurge`: %v\n", resp)
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

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugReleases

> ProjectProjectsRelease PostProjectsBySlugReleases(ctx, slug).ProjectProjectsPublish(projectProjectsPublish).Execute()

Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site to publish, from the path.
	projectProjectsPublish := *openapiclient.NewProjectProjectsPublish() // ProjectProjectsPublish | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugReleases(context.Background(), slug).ProjectProjectsPublish(projectProjectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugReleases`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugReleases`: %v\n", resp)
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

 **projectProjectsPublish** | [**ProjectProjectsPublish**](ProjectProjectsPublish.md) |  | 

### Return type

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsBySlugReleasesByReleaseActivate

> ProjectProjectsRelease PostProjectsBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()

Points the site at an existing release — the go-live, and equally the ROLLBACK.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the site the release belongs to, from the path.
	release := "release_example" // string | Release is the content-addressed release id (\"rel_\" + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsBySlugReleasesByReleaseActivate`: ProjectProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsBySlugReleasesByReleaseActivate`: %v\n", resp)
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

[**ProjectProjectsRelease**](ProjectProjectsRelease.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsFork

> ProjectProjectsProject PostProjectsFork(ctx).ProjectProjectsFork(projectProjectsFork).Execute()

Creates a project seeded from a PUBLISHED EXAMPLE — either a starter-kit template from the ONE embedded gallery catalog, or any live project on the platform (an example a seeded creator published, or another org's app serving at <slug>.hanzo.app).



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsFork := *openapiclient.NewProjectProjectsFork() // ProjectProjectsFork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsFork(context.Background()).ProjectProjectsFork(projectProjectsFork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsFork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsFork`: ProjectProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsFork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsForkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsFork** | [**ProjectProjectsFork**](ProjectProjectsFork.md) |  | 

### Return type

[**ProjectProjectsProject**](ProjectProjectsProject.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsSites

> ProjectProjectsSiteDeploy PostProjectsSites(ctx).ProjectProjectsBuildSite(projectProjectsBuildSite).Execute()

Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsBuildSite := *openapiclient.NewProjectProjectsBuildSite() // ProjectProjectsBuildSite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsSites(context.Background()).ProjectProjectsBuildSite(projectProjectsBuildSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsSites`: ProjectProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsSitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsBuildSite** | [**ProjectProjectsBuildSite**](ProjectProjectsBuildSite.md) |  | 

### Return type

[**ProjectProjectsSiteDeploy**](ProjectProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProjectsSitesDeploy

> ProjectProjectsSiteDeploy PostProjectsSitesDeploy(ctx).ProjectProjectsDeploySite(projectProjectsDeploySite).Execute()

Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	projectProjectsDeploySite := *openapiclient.NewProjectProjectsDeploySite() // ProjectProjectsDeploySite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PostProjectsSitesDeploy(context.Background()).ProjectProjectsDeploySite(projectProjectsDeploySite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PostProjectsSitesDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProjectsSitesDeploy`: ProjectProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PostProjectsSitesDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProjectsSitesDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectProjectsDeploySite** | [**ProjectProjectsDeploySite**](ProjectProjectsDeploySite.md) |  | 

### Return type

[**ProjectProjectsSiteDeploy**](ProjectProjectsSiteDeploy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectBySlugStar

> ProjectProjectsStar PutProjectBySlugStar(ctx, slug).Execute()

Bookmarks a project for the person calling, and answers whether it is starred afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PutProjectBySlugStar(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PutProjectBySlugStar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectBySlugStar`: ProjectProjectsStar
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PutProjectBySlugStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectBySlugStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsStar**](ProjectProjectsStar.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProjectsBySlugStar

> ProjectProjectsStar PutProjectsBySlugStar(ctx, slug).Execute()

Bookmarks a project for the person calling, and answers whether it is starred afterwards.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	slug := "slug_example" // string | Slug is the project to act on, from the path. It is unique within the caller's org and nowhere else, so another tenant's slug is a 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.PutProjectsBySlugStar(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.PutProjectsBySlugStar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutProjectsBySlugStar`: ProjectProjectsStar
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.PutProjectsBySlugStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProjectsBySlugStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectProjectsStar**](ProjectProjectsStar.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

