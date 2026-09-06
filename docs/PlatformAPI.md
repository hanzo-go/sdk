# \PlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlatformProjectsByProjectAppsByApp**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByApp) | **Delete** /v1/platform/projects/{project}/apps/{app} | Deletes an application and tears down what it runs.
[**DeletePlatformProjectsByProjectAppsByAppDomainsByHost**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByAppDomainsByHost) | **Delete** /v1/platform/projects/{project}/apps/{app}/domains/{host} | Detaches a hostname and releases the claim.
[**GetPlatformApps**](PlatformAPI.md#GetPlatformApps) | **Get** /v1/platform/apps | Answers what this organisation has declared, joined with what the delivery plane has done about it.
[**GetPlatformAppsByApp**](PlatformAPI.md#GetPlatformAppsByApp) | **Get** /v1/platform/apps/{app} | Answers ONE declaration — what git says this app is, before the delivery plane has had any say in it.
[**GetPlatformAppsByAppCd**](PlatformAPI.md#GetPlatformAppsByAppCd) | **Get** /v1/platform/apps/{app}/cd | Answers ONE app&#39;s reconciliation alone — the poll a deploy console makes while it waits, without re-reading the whole inventory each time.
[**GetPlatformBuilds**](PlatformAPI.md#GetPlatformBuilds) | **Get** /v1/platform/builds | Returns real build records for your org.
[**GetPlatformCd**](PlatformAPI.md#GetPlatformCd) | **Get** /v1/platform/cd | Answers every Application the delivery plane holds.
[**GetPlatformCi**](PlatformAPI.md#GetPlatformCi) | **Get** /v1/platform/ci | Continuous integration (not wired)
[**GetPlatformEnvironments**](PlatformAPI.md#GetPlatformEnvironments) | **Get** /v1/platform/environments | Returns your deploy targets, and what is running on each.
[**GetPlatformFleet**](PlatformAPI.md#GetPlatformFleet) | **Get** /v1/platform/fleet | Returns the platform&#39;s own service tier, and where it has drifted.
[**GetPlatformFleetByApp**](PlatformAPI.md#GetPlatformFleetByApp) | **Get** /v1/platform/fleet/{app} | Returns one platform service, resolved to production by default.
[**GetPlatformHealth**](PlatformAPI.md#GetPlatformHealth) | **Get** /v1/platform/health | Reports whether this control plane can actually deploy anything.
[**GetPlatformPipelines**](PlatformAPI.md#GetPlatformPipelines) | **Get** /v1/platform/pipelines | Returns one build-and-deploy pipeline per app, with its latest run.
[**GetPlatformProjects**](PlatformAPI.md#GetPlatformProjects) | **Get** /v1/platform/projects | Returns your org&#39;s projects, each with how many apps live under it.
[**GetPlatformProjectsByProject**](PlatformAPI.md#GetPlatformProjectsByProject) | **Get** /v1/platform/projects/{project} | Returns one project and its app count.
[**GetPlatformProjectsByProjectApps**](PlatformAPI.md#GetPlatformProjectsByProjectApps) | **Get** /v1/platform/projects/{project}/apps | Returns the applications in one project, with what the cluster says about them.
[**GetPlatformProjectsByProjectAppsByApp**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByApp) | **Get** /v1/platform/projects/{project}/apps/{app} | Returns one application, with its live phase, health and secret sync.
[**GetPlatformProjectsByProjectAppsByAppDeployments**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeployments) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments | Returns an app&#39;s deployment history.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsById**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsById) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id} | Returns one deployment of one app.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id}/logs | Returns real logs for a deployment — the build&#39;s, then the app&#39;s.
[**GetPlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDomains) | **Get** /v1/platform/projects/{project}/apps/{app}/domains | Returns every hostname this app answers on.
[**GetPlatformReleases**](PlatformAPI.md#GetPlatformReleases) | **Get** /v1/platform/releases | Returns the versions that actually reached the cluster.
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
[**PostPlatformRun**](PlatformAPI.md#PostPlatformRun) | **Post** /v1/platform/run | Runs a container image and gives back a URL.
[**PostPlatformRunner**](PlatformAPI.md#PostPlatformRunner) | **Post** /v1/platform/runner | Triggers a native build — an image, or the binaries a repo declares.
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformApps

> DeclaredResp GetPlatformApps(ctx).Org(org).Execute()

Answers what this organisation has declared, joined with what the delivery plane has done about it.



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
	org := "org_example" // string | Org names the organisation whose declarations to read, defaulting to the caller's own. Only a SuperAdmin may name one that is not theirs; anyone else naming a foreign org is refused, so this widens nothing by itself. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformApps(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformApps`: DeclaredResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org names the organisation whose declarations to read, defaulting to the caller&#39;s own. Only a SuperAdmin may name one that is not theirs; anyone else naming a foreign org is refused, so this widens nothing by itself. | 

### Return type

[**DeclaredResp**](DeclaredResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformAppsByApp

> Declaration GetPlatformAppsByApp(ctx, app).Org(org).Execute()

Answers ONE declaration — what git says this app is, before the delivery plane has had any say in it.



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
	app := "app_example" // string | App is the DNS-1123 label of the declaration. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which app is read whatever else is sent.
	org := "org_example" // string | Org names the organisation the declaration lives in, defaulting to the caller's own and subject to the same SuperAdmin rule as the listing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformAppsByApp(context.Background(), app).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformAppsByApp`: Declaration
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformAppsByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App is the DNS-1123 label of the declaration. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which app is read whatever else is sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | **string** | Org names the organisation the declaration lives in, defaulting to the caller&#39;s own and subject to the same SuperAdmin rule as the listing. | 

### Return type

[**Declaration**](Declaration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformAppsByAppCd

> CDApp GetPlatformAppsByAppCd(ctx, app).Org(org).Execute()

Answers ONE app's reconciliation alone — the poll a deploy console makes while it waits, without re-reading the whole inventory each time.



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
	app := "app_example" // string | App is the DNS-1123 label of the declaration. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which app is read whatever else is sent.
	org := "org_example" // string | Org names the organisation the declaration lives in, defaulting to the caller's own and subject to the same SuperAdmin rule as the listing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformAppsByAppCd(context.Background(), app).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformAppsByAppCd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformAppsByAppCd`: CDApp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformAppsByAppCd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App is the DNS-1123 label of the declaration. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which app is read whatever else is sent. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformAppsByAppCdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **org** | **string** | Org names the organisation the declaration lives in, defaulting to the caller&#39;s own and subject to the same SuperAdmin rule as the listing. | 

### Return type

[**CDApp**](CDApp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformBuilds

> BuildBoard GetPlatformBuilds(ctx).Execute()

Returns real build records for your org.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformBuilds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformBuilds`: BuildBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformBuilds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformBuildsRequest struct via the builder pattern


### Return type

[**BuildBoard**](BuildBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformCd

> CdResp GetPlatformCd(ctx).Execute()

Answers every Application the delivery plane holds.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformCd(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformCd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformCd`: CdResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformCd`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformCdRequest struct via the builder pattern


### Return type

[**CdResp**](CdResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformEnvironments

> EnvironmentBoard GetPlatformEnvironments(ctx).Execute()

Returns your deploy targets, and what is running on each.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformEnvironments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformEnvironments`: EnvironmentBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformEnvironments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformEnvironmentsRequest struct via the builder pattern


### Return type

[**EnvironmentBoard**](EnvironmentBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformPipelines

> PipelineBoard GetPlatformPipelines(ctx).Execute()

Returns one build-and-deploy pipeline per app, with its latest run.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformPipelines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformPipelines`: PipelineBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformPipelines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformPipelinesRequest struct via the builder pattern


### Return type

[**PipelineBoard**](PipelineBoard.md)

### Authorization

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformReleases

> ReleaseBoard GetPlatformReleases(ctx).Execute()

Returns the versions that actually reached the cluster.



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
	resp, r, err := apiClient.PlatformAPI.GetPlatformReleases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformReleases`: ReleaseBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformReleases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformReleasesRequest struct via the builder pattern


### Return type

[**ReleaseBoard**](ReleaseBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformRun

> RunView PostPlatformRun(ctx).RunReq(runReq).Execute()

Runs a container image and gives back a URL.



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
	runReq := *openapiclient.NewRunReq() // RunReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformRun(context.Background()).RunReq(runReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformRun`: RunView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runReq** | [**RunReq**](RunReq.md) |  | 

### Return type

[**RunView**](RunView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformRunner

> RunnerBuildResp PostPlatformRunner(ctx).RunnerBuildReq(runnerBuildReq).Execute()

Triggers a native build — an image, or the binaries a repo declares.



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
	runnerBuildReq := *openapiclient.NewRunnerBuildReq() // RunnerBuildReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformRunner(context.Background()).RunnerBuildReq(runnerBuildReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformRunner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformRunner`: RunnerBuildResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformRunner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformRunnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerBuildReq** | [**RunnerBuildReq**](RunnerBuildReq.md) |  | 

### Return type

[**RunnerBuildResp**](RunnerBuildResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

