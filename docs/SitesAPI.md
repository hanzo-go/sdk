# \SitesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSites**](SitesAPI.md#GetSites) | **Get** /v1/sites | Returns the org&#39;s deployed sites at the pretty URLs they serve at.
[**GetSitesBySlug**](SitesAPI.md#GetSitesBySlug) | **Get** /v1/sites/{slug} | Returns one site — the same row ListSites carries, for one slug.
[**GetSitesBySlugDeployments**](SitesAPI.md#GetSitesBySlugDeployments) | **Get** /v1/sites/{slug}/deployments | Returns a project&#39;s deploy history, newest version first.
[**GetSitesBySlugDeploymentsById**](SitesAPI.md#GetSitesBySlugDeploymentsById) | **Get** /v1/sites/{slug}/deployments/{id} | Returns one deployment of a project by id.
[**GetSitesBySlugReleases**](SitesAPI.md#GetSitesBySlugReleases) | **Get** /v1/sites/{slug}/releases | Returns a site&#39;s releases newest-first, marking the active one — the rollback menu.
[**PostSites**](SitesAPI.md#PostSites) | **Post** /v1/sites | Generates a self-contained, mobile-responsive static site from a natural-language brief and deploys it live in one call.
[**PostSitesBySlugDeployments**](SitesAPI.md#PostSitesBySlugDeployments) | **Post** /v1/sites/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostSitesBySlugDeploymentsByIdComplete**](SitesAPI.md#PostSitesBySlugDeploymentsByIdComplete) | **Post** /v1/sites/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostSitesBySlugPublish**](SitesAPI.md#PostSitesBySlugPublish) | **Post** /v1/sites/{slug}/publish | Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.
[**PostSitesBySlugReleases**](SitesAPI.md#PostSitesBySlugReleases) | **Post** /v1/sites/{slug}/releases | Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.
[**PostSitesBySlugReleasesByReleaseActivate**](SitesAPI.md#PostSitesBySlugReleasesByReleaseActivate) | **Post** /v1/sites/{slug}/releases/{release}/activate | Points the site at an existing release — the go-live, and equally the ROLLBACK.
[**PostSitesDeploy**](SitesAPI.md#PostSitesDeploy) | **Post** /v1/sites/deploy | Deploys a caller-supplied file manifest — the deploy_site capability an agent calls — and answers with where it went live.



## GetSites

> []ProjectsSite GetSites(ctx).Execute()

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
	resp, r, err := apiClient.SitesAPI.GetSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSites`: []ProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesRequest struct via the builder pattern


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


## GetSitesBySlug

> ProjectsSite GetSitesBySlug(ctx, slug).Execute()

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
	resp, r, err := apiClient.SitesAPI.GetSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesBySlug`: ProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSitesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesBySlugRequest struct via the builder pattern


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


## GetSitesBySlugDeployments

> []ProjectsDeployment GetSitesBySlugDeployments(ctx, slug).Execute()

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
	resp, r, err := apiClient.SitesAPI.GetSitesBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSitesBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesBySlugDeployments`: []ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSitesBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesBySlugDeploymentsRequest struct via the builder pattern


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


## GetSitesBySlugDeploymentsById

> ProjectsDeployment GetSitesBySlugDeploymentsById(ctx, slug, id).Execute()

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
	resp, r, err := apiClient.SitesAPI.GetSitesBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSitesBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesBySlugDeploymentsById`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSitesBySlugDeploymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the deployment id, from the path. A deployment of another project — or of another tenant&#39;s project — is not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesBySlugDeploymentsByIdRequest struct via the builder pattern


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


## GetSitesBySlugReleases

> []ProjectsRelease GetSitesBySlugReleases(ctx, slug).Execute()

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
	resp, r, err := apiClient.SitesAPI.GetSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.GetSitesBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSitesBySlugReleases`: []ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.GetSitesBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSitesBySlugReleasesRequest struct via the builder pattern


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


## PostSites

> ProjectsSiteDeploy PostSites(ctx).ProjectsBuildSite(projectsBuildSite).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSites(context.Background()).ProjectsBuildSite(projectsBuildSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSites`: ProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesRequest struct via the builder pattern


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


## PostSitesBySlugDeployments

> ProjectsDeployment PostSitesBySlugDeployments(ctx, slug).ProjectsDeployStart(projectsDeployStart).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesBySlugDeployments(context.Background(), slug).ProjectsDeployStart(projectsDeployStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesBySlugDeployments`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to deploy, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesBySlugDeploymentsRequest struct via the builder pattern


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


## PostSitesBySlugDeploymentsByIdComplete

> ProjectsDeployment PostSitesBySlugDeploymentsByIdComplete(ctx, slug, id).ProjectsComplete(projectsComplete).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesBySlugDeploymentsByIdComplete(context.Background(), slug, id).ProjectsComplete(projectsComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesBySlugDeploymentsByIdComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesBySlugDeploymentsByIdComplete`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesBySlugDeploymentsByIdComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the queued deployment to complete, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesBySlugDeploymentsByIdCompleteRequest struct via the builder pattern


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


## PostSitesBySlugPublish

> ProjectsRelease PostSitesBySlugPublish(ctx, slug).ProjectsPublish(projectsPublish).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesBySlugPublish(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesBySlugPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesBySlugPublish`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesBySlugPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesBySlugPublishRequest struct via the builder pattern


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


## PostSitesBySlugReleases

> ProjectsRelease PostSitesBySlugReleases(ctx, slug).ProjectsPublish(projectsPublish).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesBySlugReleases(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesBySlugReleases`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesBySlugReleasesRequest struct via the builder pattern


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


## PostSitesBySlugReleasesByReleaseActivate

> ProjectsRelease PostSitesBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesBySlugReleasesByReleaseActivate`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesBySlugReleasesByReleaseActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the release belongs to, from the path. | 
**release** | **string** | Release is the content-addressed release id (\&quot;rel_\&quot; + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesBySlugReleasesByReleaseActivateRequest struct via the builder pattern


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


## PostSitesDeploy

> ProjectsSiteDeploy PostSitesDeploy(ctx).ProjectsDeploySite(projectsDeploySite).Execute()

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
	resp, r, err := apiClient.SitesAPI.PostSitesDeploy(context.Background()).ProjectsDeploySite(projectsDeploySite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.PostSitesDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSitesDeploy`: ProjectsSiteDeploy
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.PostSitesDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSitesDeployRequest struct via the builder pattern


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

