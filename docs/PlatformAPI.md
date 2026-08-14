# \PlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlatformProjectsByProjectAppsByApp**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByApp) | **Delete** /v1/platform/projects/{project}/apps/{app} | Deletes an application and tears down what it runs.
[**DeletePlatformProjectsByProjectAppsByAppDomainsByHost**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByAppDomainsByHost) | **Delete** /v1/platform/projects/{project}/apps/{app}/domains/{host} | Detaches a hostname and releases the claim.
[**DeletePlatformSitesBySlug**](PlatformAPI.md#DeletePlatformSitesBySlug) | **Delete** /v1/platform/sites/{slug} | Deletes a project and takes its site off the internet.
[**DeletePlatformSitesBySlugDomainsByHost**](PlatformAPI.md#DeletePlatformSitesBySlugDomainsByHost) | **Delete** /v1/platform/sites/{slug}/domains/{host} | Gives a custom hostname back, so the name is free to reuse.
[**GetPlatformApps**](PlatformAPI.md#GetPlatformApps) | **Get** /v1/platform/apps | What this organization has declared, and what CD did with it
[**GetPlatformAppsByApp**](PlatformAPI.md#GetPlatformAppsByApp) | **Get** /v1/platform/apps/{app} | One declaration
[**GetPlatformAppsByAppCd**](PlatformAPI.md#GetPlatformAppsByAppCd) | **Get** /v1/platform/apps/{app}/cd | One app&#39;s reconciliation
[**GetPlatformCd**](PlatformAPI.md#GetPlatformCd) | **Get** /v1/platform/cd | The delivery plane
[**GetPlatformCi**](PlatformAPI.md#GetPlatformCi) | **Get** /v1/platform/ci | Continuous integration (not wired)
[**GetPlatformFleet**](PlatformAPI.md#GetPlatformFleet) | **Get** /v1/platform/fleet | Returns the platform&#39;s own service tier, and where it has drifted.
[**GetPlatformFleetByApp**](PlatformAPI.md#GetPlatformFleetByApp) | **Get** /v1/platform/fleet/{app} | Returns one platform service, resolved to production by default.
[**GetPlatformHealth**](PlatformAPI.md#GetPlatformHealth) | **Get** /v1/platform/health | Reports whether this control plane can actually deploy anything.
[**GetPlatformProjects**](PlatformAPI.md#GetPlatformProjects) | **Get** /v1/platform/projects | Returns your org&#39;s projects, each with how many apps live under it.
[**GetPlatformProjectsByProject**](PlatformAPI.md#GetPlatformProjectsByProject) | **Get** /v1/platform/projects/{project} | Returns one project and its app count.
[**GetPlatformProjectsByProjectApps**](PlatformAPI.md#GetPlatformProjectsByProjectApps) | **Get** /v1/platform/projects/{project}/apps | Returns the applications in one project, with what the cluster says about them.
[**GetPlatformProjectsByProjectAppsByApp**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByApp) | **Get** /v1/platform/projects/{project}/apps/{app} | Returns one application, with its live phase, health and secret sync.
[**GetPlatformProjectsByProjectAppsByAppDeployments**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeployments) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments | Returns an app&#39;s deployment history.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsById**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsById) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id} | Returns one deployment of one app.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id}/logs | Returns real logs for a deployment — the build&#39;s, then the app&#39;s.
[**GetPlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDomains) | **Get** /v1/platform/projects/{project}/apps/{app}/domains | Returns every hostname this app answers on.
[**GetPlatformSites**](PlatformAPI.md#GetPlatformSites) | **Get** /v1/platform/sites | Returns every project your org owns.
[**GetPlatformSitesBySlug**](PlatformAPI.md#GetPlatformSitesBySlug) | **Get** /v1/platform/sites/{slug} | Returns one project of yours by slug — its settings, its live URL and the deployment currently serving it.
[**GetPlatformSitesBySlugDeployments**](PlatformAPI.md#GetPlatformSitesBySlugDeployments) | **Get** /v1/platform/sites/{slug}/deployments | Returns a project&#39;s deploy history, newest version first.
[**GetPlatformSitesBySlugDeploymentsById**](PlatformAPI.md#GetPlatformSitesBySlugDeploymentsById) | **Get** /v1/platform/sites/{slug}/deployments/{id} | Returns one deployment of a project by id.
[**GetPlatformSitesBySlugDomains**](PlatformAPI.md#GetPlatformSitesBySlugDomains) | **Get** /v1/platform/sites/{slug}/domains | Returns every custom hostname this site holds: the live ones, plus any pending claim with the DNS records it still owes.
[**GetPlatformSitesBySlugReleases**](PlatformAPI.md#GetPlatformSitesBySlugReleases) | **Get** /v1/platform/sites/{slug}/releases | Returns a site&#39;s releases newest-first, marking the active one — the rollback menu.
[**PatchPlatformSitesBySlug**](PlatformAPI.md#PatchPlatformSitesBySlug) | **Patch** /v1/platform/sites/{slug} | Changes a project&#39;s settings, and only the settings you send.
[**PostPlatformApps**](PlatformAPI.md#PostPlatformApps) | **Post** /v1/platform/apps | Deploy an app through cd.hanzo.ai
[**PostPlatformFleetByAppDeploy**](PlatformAPI.md#PostPlatformFleetByAppDeploy) | **Post** /v1/platform/fleet/{app}/deploy | Rolls a platform service&#39;s pods, in a named environment.
[**PostPlatformProjectsByProjectApps**](PlatformAPI.md#PostPlatformProjectsByProjectApps) | **Post** /v1/platform/projects/{project}/apps | Creates an application from a git repo or a container image.
[**PostPlatformProjectsByProjectAppsByAppDeploy**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppDeploy) | **Post** /v1/platform/projects/{project}/apps/{app}/deploy | Deploys the app — building it first if it comes from git.
[**PostPlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppDomains) | **Post** /v1/platform/projects/{project}/apps/{app}/domains | Attaches a hostname — instantly if you already own it, otherwise with a DNS challenge.
[**PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify) | **Post** /v1/platform/projects/{project}/apps/{app}/domains/{host}/verify | Checks a custom domain&#39;s DNS and turns it on if it passes.
[**PostPlatformProjectsByProjectAppsByAppPreview**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppPreview) | **Post** /v1/platform/projects/{project}/apps/{app}/preview | Puts a branch on its own URL.
[**PostPlatformProjectsByProjectAppsByAppPromote**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppPromote) | **Post** /v1/platform/projects/{project}/apps/{app}/promote | Promotes an already-built release to the app.
[**PostPlatformProjectsByProjectAppsByAppRollback**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppRollback) | **Post** /v1/platform/projects/{project}/apps/{app}/rollback | Goes back to the previous release.
[**PostPlatformProjectsByProjectAppsByAppStart**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppStart) | **Post** /v1/platform/projects/{project}/apps/{app}/start | Starts a stopped app back up.
[**PostPlatformProjectsByProjectAppsByAppStop**](PlatformAPI.md#PostPlatformProjectsByProjectAppsByAppStop) | **Post** /v1/platform/projects/{project}/apps/{app}/stop | Stops an app without deleting it.
[**PostPlatformSites**](PlatformAPI.md#PostPlatformSites) | **Post** /v1/platform/sites | Creates a project — the handle a site is deployed and served under — and answers 201 with it in &#x60;draft&#x60;.
[**PostPlatformSitesBySlugDeploy**](PlatformAPI.md#PostPlatformSitesBySlugDeploy) | **Post** /v1/platform/sites/{slug}/deploy | Upload a built site as one archive and serve it
[**PostPlatformSitesBySlugDeployments**](PlatformAPI.md#PostPlatformSitesBySlugDeployments) | **Post** /v1/platform/sites/{slug}/deployments | Opens a deployment and hands back a short-lived, prefix-scoped grant to write its bytes straight to object storage.
[**PostPlatformSitesBySlugDeploymentsByIdComplete**](PlatformAPI.md#PostPlatformSitesBySlugDeploymentsByIdComplete) | **Post** /v1/platform/sites/{slug}/deployments/{id}/complete | CompleteDeployment is the CI completion hook that flips a queued git deployment to live (or error) once CI has synced the built site to S3.
[**PostPlatformSitesBySlugDomains**](PlatformAPI.md#PostPlatformSitesBySlugDomains) | **Post** /v1/platform/sites/{slug}/domains | Attaches one or more CUSTOM public hostnames to this org&#39;s site.
[**PostPlatformSitesBySlugDomainsByHostVerify**](PlatformAPI.md#PostPlatformSitesBySlugDomainsByHostVerify) | **Post** /v1/platform/sites/{slug}/domains/{host}/verify | Checks the DNS challenge for a pending custom hostname and, when it passes, promotes the host so it begins routing at the edge.
[**PostPlatformSitesBySlugPublish**](PlatformAPI.md#PostPlatformSitesBySlugPublish) | **Post** /v1/platform/sites/{slug}/publish | Promotes a build output into a new release AND goes live with it — create+activate in one call, which is the 99% path.
[**PostPlatformSitesBySlugPurge**](PlatformAPI.md#PostPlatformSitesBySlugPurge) | **Post** /v1/platform/sites/{slug}/purge | Flushes the site&#39;s edge cache without redeploying anything.
[**PostPlatformSitesBySlugReleases**](PlatformAPI.md#PostPlatformSitesBySlugReleases) | **Post** /v1/platform/sites/{slug}/releases | Promotes a build output into a new immutable release WITHOUT serving it — the staged half of publishing, for when you want to check a release before it goes live.
[**PostPlatformSitesBySlugReleasesByReleaseActivate**](PlatformAPI.md#PostPlatformSitesBySlugReleasesByReleaseActivate) | **Post** /v1/platform/sites/{slug}/releases/{release}/activate | Points the site at an existing release — the go-live, and equally the ROLLBACK.
[**PutPlatformProjectsByProjectAppsByAppEnv**](PlatformAPI.md#PutPlatformProjectsByProjectAppsByAppEnv) | **Put** /v1/platform/projects/{project}/apps/{app}/env | Replaces an app&#39;s environment variables.



## DeletePlatformProjectsByProjectAppsByApp

> DeletePlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()

Deletes an application and tears down what it runs.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.DeletePlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeletePlatformProjectsByProjectAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


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


## DeletePlatformProjectsByProjectAppsByAppDomainsByHost

> DeletePlatformProjectsByProjectAppsByAppDomainsByHost(ctx, project, app, host).Execute()

Detaches a hostname and releases the claim.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	host := "host_example" // string | Host is the hostname, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.DeletePlatformProjectsByProjectAppsByAppDomainsByHost(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeletePlatformProjectsByProjectAppsByAppDomainsByHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 
**host** | **string** | Host is the hostname, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlatformProjectsByProjectAppsByAppDomainsByHostRequest struct via the builder pattern


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


## DeletePlatformSitesBySlug

> DeletePlatformSitesBySlug(ctx, slug).Execute()

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
	r, err := apiClient.PlatformAPI.DeletePlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeletePlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePlatformSitesBySlugRequest struct via the builder pattern


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


## DeletePlatformSitesBySlugDomainsByHost

> DeletePlatformSitesBySlugDomainsByHost(ctx, slug, host).Execute()

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
	r, err := apiClient.PlatformAPI.DeletePlatformSitesBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeletePlatformSitesBySlugDomainsByHost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePlatformSitesBySlugDomainsByHostRequest struct via the builder pattern


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


## GetPlatformApps

> GetPlatformApps(ctx).Execute()

What this organization has declared, and what CD did with it



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
	r, err := apiClient.PlatformAPI.GetPlatformApps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsRequest struct via the builder pattern


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


## GetPlatformAppsByApp

> GetPlatformAppsByApp(ctx, app).Execute()

One declaration



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
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.GetPlatformAppsByApp(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsByAppRequest struct via the builder pattern


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


## GetPlatformAppsByAppCd

> GetPlatformAppsByAppCd(ctx, app).Execute()

One app's reconciliation



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
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.GetPlatformAppsByAppCd(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformAppsByAppCd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsByAppCdRequest struct via the builder pattern


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


## GetPlatformCd

> GetPlatformCd(ctx).Execute()

The delivery plane



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
	r, err := apiClient.PlatformAPI.GetPlatformCd(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformCd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformCdRequest struct via the builder pattern


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


## GetPlatformCi

> GetPlatformCi(ctx).Execute()

Continuous integration (not wired)



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
	r, err := apiClient.PlatformAPI.GetPlatformCi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformCi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformCiRequest struct via the builder pattern


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


## GetPlatformFleet

> DriftBoard GetPlatformFleet(ctx).Env(env).Health(health).Org(org).Drift(drift).Execute()

Returns the platform's own service tier, and where it has drifted.



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
	env := "env_example" // string | Env narrows to one lifecycle env: main, test or dev. (optional)
	health := "health_example" // string | Health narrows to one health colour: green, yellow or red. (optional)
	org := "org_example" // string | Org narrows to one image namespace. (optional)
	drift := "drift_example" // string | Drift is `1` or `true` to show only rows that have actually drifted. It is a STRING and not a bool because those two spellings are exactly what the board has always accepted, and a bool would silently widen that to `?drift` alone and to `TRUE` — a behaviour change wearing a type change's clothes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformFleet(context.Background()).Env(env).Health(health).Org(org).Drift(drift).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformFleet`: DriftBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformFleet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** | Env narrows to one lifecycle env: main, test or dev. | 
 **health** | **string** | Health narrows to one health colour: green, yellow or red. | 
 **org** | **string** | Org narrows to one image namespace. | 
 **drift** | **string** | Drift is &#x60;1&#x60; or &#x60;true&#x60; to show only rows that have actually drifted. It is a STRING and not a bool because those two spellings are exactly what the board has always accepted, and a bool would silently widen that to &#x60;?drift&#x60; alone and to &#x60;TRUE&#x60; — a behaviour change wearing a type change&#39;s clothes. | 

### Return type

[**DriftBoard**](DriftBoard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformFleetByApp

> AppView GetPlatformFleetByApp(ctx, app).Env(env).Execute()

Returns one platform service, resolved to production by default.



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
	app := "app_example" // string | App is the service's CR name, from the path. It must be a DNS-1123 label.
	env := "env_example" // string | Env narrows the scan to one lifecycle env: main, test or dev. Omitted, the namespaces are scanned in lifecycle order and the first match wins, so a bare name resolves to PRODUCTION. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformFleetByApp(context.Background(), app).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformFleetByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformFleetByApp`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformFleetByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App is the service&#39;s CR name, from the path. It must be a DNS-1123 label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformFleetByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Env narrows the scan to one lifecycle env: main, test or dev. Omitted, the namespaces are scanned in lifecycle order and the first match wins, so a bare name resolves to PRODUCTION. | 

### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformHealth

> Readiness GetPlatformHealth(ctx).Execute()

Reports whether this control plane can actually deploy anything.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformHealth`: Readiness
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformHealthRequest struct via the builder pattern


### Return type

[**Readiness**](Readiness.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjects

> []ProjectView GetPlatformProjects(ctx).Execute()

Returns your org's projects, each with how many apps live under it.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjects`: []ProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsRequest struct via the builder pattern


### Return type

[**[]ProjectView**](ProjectView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProject

> ProjectView GetPlatformProjectsByProject(ctx, project).Execute()

Returns one project and its app count.



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
	project := "project_example" // string | Project is the project's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProject`: ProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectView**](ProjectView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectApps

> []AppView GetPlatformProjectsByProjectApps(ctx, project).Execute()

Returns the applications in one project, with what the cluster says about them.



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
	project := "project_example" // string | Project is the project's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectApps(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectApps`: []AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByApp

> AppView GetPlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()

Returns one application, with its live phase, health and secret sync.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectAppsByApp`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectAppsByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeployments

> []DeploymentView GetPlatformProjectsByProjectAppsByAppDeployments(ctx, project, app).Execute()

Returns an app's deployment history.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeployments(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectAppsByAppDeployments`: []DeploymentView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsByAppDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeploymentsById

> DeploymentView GetPlatformProjectsByProjectAppsByAppDeploymentsById(ctx, project, app, id).Execute()

Returns one deployment of one app.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	id := "id_example" // string | ID is the deployment's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsById(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectAppsByAppDeploymentsById`: DeploymentView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 
**id** | **string** | ID is the deployment&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsByAppDeploymentsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs

> DeployLogs GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(ctx, project, app, id).Execute()

Returns real logs for a deployment — the build's, then the app's.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	id := "id_example" // string | ID is the deployment's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs`: DeployLogs
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 
**id** | **string** | ID is the deployment&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DeployLogs**](DeployLogs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDomains

> []DomainView GetPlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()

Returns every hostname this app answers on.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProjectAppsByAppDomains`: []DomainView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjectsByProjectAppsByAppDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]DomainView**](DomainView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformSites

> []ProjectsProject GetPlatformSites(ctx).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSites`: []ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesRequest struct via the builder pattern


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


## GetPlatformSitesBySlug

> ProjectsProject GetPlatformSitesBySlug(ctx, slug).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSitesBySlug`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSitesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesBySlugRequest struct via the builder pattern


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


## GetPlatformSitesBySlugDeployments

> []ProjectsDeployment GetPlatformSitesBySlugDeployments(ctx, slug).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSitesBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSitesBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSitesBySlugDeployments`: []ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSitesBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesBySlugDeploymentsRequest struct via the builder pattern


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


## GetPlatformSitesBySlugDeploymentsById

> ProjectsDeployment GetPlatformSitesBySlugDeploymentsById(ctx, slug, id).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSitesBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSitesBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSitesBySlugDeploymentsById`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSitesBySlugDeploymentsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the deployment id, from the path. A deployment of another project — or of another tenant&#39;s project — is not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesBySlugDeploymentsByIdRequest struct via the builder pattern


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


## GetPlatformSitesBySlugDomains

> ProjectsDomains GetPlatformSitesBySlugDomains(ctx, slug).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSitesBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSitesBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSitesBySlugDomains`: ProjectsDomains
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSitesBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## GetPlatformSitesBySlugReleases

> []ProjectsRelease GetPlatformSitesBySlugReleases(ctx, slug).Execute()

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
	resp, r, err := apiClient.PlatformAPI.GetPlatformSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformSitesBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformSitesBySlugReleases`: []ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformSitesBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformSitesBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ProjectsRelease**](ProjectsRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPlatformSitesBySlug

> ProjectsProject PatchPlatformSitesBySlug(ctx, slug).ProjectsUpdate(projectsUpdate).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PatchPlatformSitesBySlug(context.Background(), slug).ProjectsUpdate(projectsUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PatchPlatformSitesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPlatformSitesBySlug`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PatchPlatformSitesBySlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to update, from the path. The URL is the addressing authority — a &#x60;slug&#x60; in the body cannot move the write to another project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPlatformSitesBySlugRequest struct via the builder pattern


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


## PostPlatformApps

> PostPlatformApps(ctx).Execute()

Deploy an app through cd.hanzo.ai



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
	r, err := apiClient.PlatformAPI.PostPlatformApps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformAppsRequest struct via the builder pattern


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


## PostPlatformFleetByAppDeploy

> Restarted PostPlatformFleetByAppDeploy(ctx, app).RestartRef(restartRef).Execute()

Rolls a platform service's pods, in a named environment.



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
	app := "app_example" // string | App is the service's CR name, from the path. It must be a DNS-1123 label.
	restartRef := *openapiclient.NewRestartRef() // RestartRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformFleetByAppDeploy(context.Background(), app).RestartRef(restartRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformFleetByAppDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformFleetByAppDeploy`: Restarted
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformFleetByAppDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App is the service&#39;s CR name, from the path. It must be a DNS-1123 label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformFleetByAppDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restartRef** | [**RestartRef**](RestartRef.md) |  | 

### Return type

[**Restarted**](Restarted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectApps

> AppView PostPlatformProjectsByProjectApps(ctx, project).CreateAppReq(createAppReq).Execute()

Creates an application from a git repo or a container image.



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
	project := "project_example" // string | Project is the project to create the application under, from the path.
	createAppReq := *openapiclient.NewCreateAppReq() // CreateAppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectApps(context.Background(), project).CreateAppReq(createAppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectApps`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project to create the application under, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAppReq** | [**CreateAppReq**](CreateAppReq.md) |  | 

### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDeploy

> DeploymentView PostPlatformProjectsByProjectAppsByAppDeploy(ctx, project, app).DeployReq(deployReq).Execute()

Deploys the app — building it first if it comes from git.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	deployReq := *openapiclient.NewDeployReq() // DeployReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppDeploy(context.Background(), project, app).DeployReq(deployReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppDeploy`: DeploymentView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deployReq** | [**DeployReq**](DeployReq.md) |  | 

### Return type

[**DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDomains

> DomainView PostPlatformProjectsByProjectAppsByAppDomains(ctx, project, app).AddDomainReq(addDomainReq).Execute()

Attaches a hostname — instantly if you already own it, otherwise with a DNS challenge.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	addDomainReq := *openapiclient.NewAddDomainReq() // AddDomainReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).AddDomainReq(addDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppDomains`: DomainView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addDomainReq** | [**AddDomainReq**](AddDomainReq.md) |  | 

### Return type

[**DomainView**](DomainView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify

> DomainView PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify(ctx, project, app, host).Execute()

Checks a custom domain's DNS and turns it on if it passes.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	host := "host_example" // string | Host is the hostname, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify`: DomainView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 
**host** | **string** | Host is the hostname, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppDomainsByHostVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DomainView**](DomainView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppPreview

> PreviewView PostPlatformProjectsByProjectAppsByAppPreview(ctx, project, app).PreviewReq(previewReq).Execute()

Puts a branch on its own URL.



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
	project := "project_example" // string | Project is the project the parent application lives under, from the path.
	app := "app_example" // string | App is the parent application's slug, from the path.
	previewReq := *openapiclient.NewPreviewReq() // PreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppPreview(context.Background(), project, app).PreviewReq(previewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppPreview`: PreviewView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the parent application lives under, from the path. | 
**app** | **string** | App is the parent application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **previewReq** | [**PreviewReq**](PreviewReq.md) |  | 

### Return type

[**PreviewView**](PreviewView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppPromote

> DeploymentView PostPlatformProjectsByProjectAppsByAppPromote(ctx, project, app).PromoteReq(promoteReq).Execute()

Promotes an already-built release to the app.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	promoteReq := *openapiclient.NewPromoteReq() // PromoteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppPromote(context.Background(), project, app).PromoteReq(promoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPromote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppPromote`: DeploymentView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPromote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppPromoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **promoteReq** | [**PromoteReq**](PromoteReq.md) |  | 

### Return type

[**DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppRollback

> DeploymentView PostPlatformProjectsByProjectAppsByAppRollback(ctx, project, app).RollbackReq(rollbackReq).Execute()

Goes back to the previous release.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	rollbackReq := *openapiclient.NewRollbackReq() // RollbackReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppRollback(context.Background(), project, app).RollbackReq(rollbackReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppRollback`: DeploymentView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rollbackReq** | [**RollbackReq**](RollbackReq.md) |  | 

### Return type

[**DeploymentView**](DeploymentView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppStart

> AppView PostPlatformProjectsByProjectAppsByAppStart(ctx, project, app).Execute()

Starts a stopped app back up.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppStart(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppStart`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppStop

> AppView PostPlatformProjectsByProjectAppsByAppStop(ctx, project, app).Execute()

Stops an app without deleting it.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppStop(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppStop`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjectsByProjectAppsByAppStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsByProjectAppsByAppStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformSites

> ProjectsProject PostPlatformSites(ctx).ProjectsCreate(projectsCreate).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSites(context.Background()).ProjectsCreate(projectsCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSites`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesRequest struct via the builder pattern


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


## PostPlatformSitesBySlugDeploy

> ProjectsDeployment PostPlatformSitesBySlugDeploy(ctx, slug).Body(body).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugDeploy(context.Background(), slug).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugDeploy`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugDeploy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugDeployRequest struct via the builder pattern


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


## PostPlatformSitesBySlugDeployments

> ProjectsDeployment PostPlatformSitesBySlugDeployments(ctx, slug).ProjectsDeployStart(projectsDeployStart).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugDeployments(context.Background(), slug).ProjectsDeployStart(projectsDeployStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugDeployments`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to deploy, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugDeploymentsRequest struct via the builder pattern


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


## PostPlatformSitesBySlugDeploymentsByIdComplete

> ProjectsDeployment PostPlatformSitesBySlugDeploymentsByIdComplete(ctx, slug, id).ProjectsComplete(projectsComplete).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugDeploymentsByIdComplete(context.Background(), slug, id).ProjectsComplete(projectsComplete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugDeploymentsByIdComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugDeploymentsByIdComplete`: ProjectsDeployment
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugDeploymentsByIdComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the deployment belongs to, from the path. | 
**id** | **string** | ID is the queued deployment to complete, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugDeploymentsByIdCompleteRequest struct via the builder pattern


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


## PostPlatformSitesBySlugDomains

> ProjectsBoundDomains PostPlatformSitesBySlugDomains(ctx, slug).ProjectsDomainsBind(projectsDomainsBind).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugDomains(context.Background(), slug).ProjectsDomainsBind(projectsDomainsBind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugDomains`: ProjectsBoundDomains
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the hosts attach to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## PostPlatformSitesBySlugDomainsByHostVerify

> ProjectsDomain PostPlatformSitesBySlugDomainsByHostVerify(ctx, slug, host).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugDomainsByHostVerify`: ProjectsDomain
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugDomainsByHostVerify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project the host is attached to, from the path. | 
**host** | **string** | Host is the custom hostname, from the path. It is cleaned to its canonical form (lowercased, trailing dot dropped) before anything is looked up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugDomainsByHostVerifyRequest struct via the builder pattern


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


## PostPlatformSitesBySlugPublish

> ProjectsRelease PostPlatformSitesBySlugPublish(ctx, slug).ProjectsPublish(projectsPublish).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugPublish(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugPublish`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsPublish** | [**ProjectsPublish**](ProjectsPublish.md) |  | 

### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformSitesBySlugPurge

> ProjectsProject PostPlatformSitesBySlugPurge(ctx, slug).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugPurge`: ProjectsProject
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugPurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the project to act on, from the path. It is unique within the caller&#39;s org and nowhere else, so another tenant&#39;s slug is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugPurgeRequest struct via the builder pattern


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


## PostPlatformSitesBySlugReleases

> ProjectsRelease PostPlatformSitesBySlugReleases(ctx, slug).ProjectsPublish(projectsPublish).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugReleases(context.Background(), slug).ProjectsPublish(projectsPublish).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugReleases`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site to publish, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsPublish** | [**ProjectsPublish**](ProjectsPublish.md) |  | 

### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformSitesBySlugReleasesByReleaseActivate

> ProjectsRelease PostPlatformSitesBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()

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
	resp, r, err := apiClient.PlatformAPI.PostPlatformSitesBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformSitesBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformSitesBySlugReleasesByReleaseActivate`: ProjectsRelease
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformSitesBySlugReleasesByReleaseActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug is the site the release belongs to, from the path. | 
**release** | **string** | Release is the content-addressed release id (\&quot;rel_\&quot; + 32 hex chars), from the path. Anything that is not that shape is not found, rather than being interpolated into a storage prefix. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformSitesBySlugReleasesByReleaseActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProjectsRelease**](ProjectsRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlatformProjectsByProjectAppsByAppEnv

> AppView PutPlatformProjectsByProjectAppsByAppEnv(ctx, project, app).SetEnvReq(setEnvReq).Execute()

Replaces an app's environment variables.



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
	project := "project_example" // string | Project is the project the application lives under, from the path.
	app := "app_example" // string | App is the application's slug, from the path.
	setEnvReq := *openapiclient.NewSetEnvReq() // SetEnvReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PutPlatformProjectsByProjectAppsByAppEnv(context.Background(), project, app).SetEnvReq(setEnvReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PutPlatformProjectsByProjectAppsByAppEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlatformProjectsByProjectAppsByAppEnv`: AppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PutPlatformProjectsByProjectAppsByAppEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project the application lives under, from the path. | 
**app** | **string** | App is the application&#39;s slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlatformProjectsByProjectAppsByAppEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setEnvReq** | [**SetEnvReq**](SetEnvReq.md) |  | 

### Return type

[**AppView**](AppView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

