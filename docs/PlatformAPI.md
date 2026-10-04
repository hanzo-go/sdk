# \PlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlatformProjectsByProject**](PlatformAPI.md#DeletePlatformProjectsByProject) | **Delete** /v1/platform/projects/{project} | Folds a project into another.
[**DeletePlatformProjectsByProjectAppsByApp**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByApp) | **Delete** /v1/platform/projects/{project}/apps/{app} | Deletes an application and tears down what it runs.
[**DeletePlatformProjectsByProjectAppsByAppDomainsByHost**](PlatformAPI.md#DeletePlatformProjectsByProjectAppsByAppDomainsByHost) | **Delete** /v1/platform/projects/{project}/apps/{app}/domains/{host} | Detaches a hostname and releases the claim.
[**GetBuildById**](PlatformAPI.md#GetBuildById) | **Get** /v1/build/{id} | Answers one build: its status, and for a failed build the reason.
[**GetPlatformApps**](PlatformAPI.md#GetPlatformApps) | **Get** /v1/platform/apps | Answers what this organisation has declared, joined with what the delivery plane has done about it.
[**GetPlatformAppsByApp**](PlatformAPI.md#GetPlatformAppsByApp) | **Get** /v1/platform/apps/{app} | Answers ONE declaration — what git says this app is, before the delivery plane has had any say in it.
[**GetPlatformAppsByAppCd**](PlatformAPI.md#GetPlatformAppsByAppCd) | **Get** /v1/platform/apps/{app}/cd | Answers ONE app&#39;s reconciliation alone — the poll a deploy console makes while it waits, without re-reading the whole inventory each time.
[**GetPlatformBuilds**](PlatformAPI.md#GetPlatformBuilds) | **Get** /v1/platform/builds | Returns real build records for your org.
[**GetPlatformCd**](PlatformAPI.md#GetPlatformCd) | **Get** /v1/platform/cd | Answers every Application the delivery plane holds.
[**GetPlatformEnvironments**](PlatformAPI.md#GetPlatformEnvironments) | **Get** /v1/platform/environments | Returns your deploy targets, and what is running on each.
[**GetPlatformFleet**](PlatformAPI.md#GetPlatformFleet) | **Get** /v1/platform/fleet | Returns the platform&#39;s own service tier, and where it has drifted.
[**GetPlatformFleetByApp**](PlatformAPI.md#GetPlatformFleetByApp) | **Get** /v1/platform/fleet/{app} | Returns one platform service, resolved to production by default.
[**GetPlatformHealth**](PlatformAPI.md#GetPlatformHealth) | **Get** /v1/platform/health | Reports whether this control plane can actually deploy anything.
[**GetPlatformPipelines**](PlatformAPI.md#GetPlatformPipelines) | **Get** /v1/platform/pipelines | Returns one build-and-deploy pipeline per app, with its latest run.
[**GetPlatformProjects**](PlatformAPI.md#GetPlatformProjects) | **Get** /v1/platform/projects | Answers every project the caller may see, with how its apps stand.
[**GetPlatformProjectsByProject**](PlatformAPI.md#GetPlatformProjectsByProject) | **Get** /v1/platform/projects/{project} | Answers one project and every app in it.
[**GetPlatformProjectsByProjectApps**](PlatformAPI.md#GetPlatformProjectsByProjectApps) | **Get** /v1/platform/projects/{project}/apps | Returns the applications in one project, with what the cluster says about them.
[**GetPlatformProjectsByProjectAppsByApp**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByApp) | **Get** /v1/platform/projects/{project}/apps/{app} | Returns one application, with its live phase, health and secret sync.
[**GetPlatformProjectsByProjectAppsByAppDeployments**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeployments) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments | Returns an app&#39;s deployment history.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsById**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsById) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id} | Returns one deployment of one app.
[**GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id}/logs | Returns real logs for a deployment — the build&#39;s, then the app&#39;s.
[**GetPlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#GetPlatformProjectsByProjectAppsByAppDomains) | **Get** /v1/platform/projects/{project}/apps/{app}/domains | Returns every hostname this app answers on.
[**GetPlatformReleases**](PlatformAPI.md#GetPlatformReleases) | **Get** /v1/platform/releases | Returns the versions that actually reached the cluster.
[**PostBuild**](PlatformAPI.md#PostBuild) | **Post** /v1/build | Triggers a native build — an image, or the binaries a repo declares.
[**PostPlatformApps**](PlatformAPI.md#PostPlatformApps) | **Post** /v1/platform/apps | Deploy an app through cd.hanzo.ai
[**PostPlatformProjects**](PlatformAPI.md#PostPlatformProjects) | **Post** /v1/platform/projects | Creates a project from the apps it starts with.
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
[**PutPlatformAppsByAppProject**](PlatformAPI.md#PutPlatformAppsByAppProject) | **Put** /v1/platform/apps/{app}/project | Moves an app to a project.
[**PutPlatformProjectsByProject**](PlatformAPI.md#PutPlatformProjectsByProject) | **Put** /v1/platform/projects/{project} | Renames a project.
[**PutPlatformProjectsByProjectAppsByAppEnv**](PlatformAPI.md#PutPlatformProjectsByProjectAppsByAppEnv) | **Put** /v1/platform/projects/{project}/apps/{app}/env | Replaces an app&#39;s environment variables.



## DeletePlatformProjectsByProject

> PlatformProjectWrite DeletePlatformProjectsByProject(ctx, project).Into(into).Org(org).Mode(mode).Execute()

Folds a project into another.



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
	project := "project_example" // string | Project is the project to delete, from the path.
	into := "into_example" // string | Into is the existing project its apps move into. Required: an app always belongs to exactly one project. (optional)
	org := "org_example" // string | Org names the projects' owner, defaulting to the caller's own scope. (optional)
	mode := "mode_example" // string | Mode is `branch` (the default) or `commit`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.DeletePlatformProjectsByProject(context.Background(), project).Into(into).Org(org).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.DeletePlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlatformProjectsByProject`: PlatformProjectWrite
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.DeletePlatformProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project to delete, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlatformProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **into** | **string** | Into is the existing project its apps move into. Required: an app always belongs to exactly one project. | 
 **org** | **string** | Org names the projects&#39; owner, defaulting to the caller&#39;s own scope. | 
 **mode** | **string** | Mode is &#x60;branch&#x60; (the default) or &#x60;commit&#x60;. | 

### Return type

[**PlatformProjectWrite**](PlatformProjectWrite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
- **Accept**: application/problem+json

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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildById

> PlatformRunnerBuildResp GetBuildById(ctx, id).Execute()

Answers one build: its status, and for a failed build the reason.



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
	id := "id_example" // string | ID is the build's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetBuildById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetBuildById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildById`: PlatformRunnerBuildResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetBuildById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the build&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformRunnerBuildResp**](PlatformRunnerBuildResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformApps

> PlatformDeclaredResp GetPlatformApps(ctx).Org(org).Execute()

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
	// response from `GetPlatformApps`: PlatformDeclaredResp
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

[**PlatformDeclaredResp**](PlatformDeclaredResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformAppsByApp

> PlatformDeclaration GetPlatformAppsByApp(ctx, app).Org(org).Execute()

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
	// response from `GetPlatformAppsByApp`: PlatformDeclaration
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

[**PlatformDeclaration**](PlatformDeclaration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformAppsByAppCd

> PlatformCDApp GetPlatformAppsByAppCd(ctx, app).Org(org).Execute()

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
	// response from `GetPlatformAppsByAppCd`: PlatformCDApp
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

[**PlatformCDApp**](PlatformCDApp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformBuilds

> PlatformBuildBoard GetPlatformBuilds(ctx).Execute()

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
	// response from `GetPlatformBuilds`: PlatformBuildBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformBuilds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformBuildsRequest struct via the builder pattern


### Return type

[**PlatformBuildBoard**](PlatformBuildBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformCd

> PlatformCdResp GetPlatformCd(ctx).Execute()

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
	// response from `GetPlatformCd`: PlatformCdResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformCd`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformCdRequest struct via the builder pattern


### Return type

[**PlatformCdResp**](PlatformCdResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformEnvironments

> PlatformEnvironmentBoard GetPlatformEnvironments(ctx).Execute()

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
	// response from `GetPlatformEnvironments`: PlatformEnvironmentBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformEnvironments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformEnvironmentsRequest struct via the builder pattern


### Return type

[**PlatformEnvironmentBoard**](PlatformEnvironmentBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformFleet

> PlatformDriftBoard GetPlatformFleet(ctx).Env(env).Health(health).Org(org).Drift(drift).Execute()

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
	// response from `GetPlatformFleet`: PlatformDriftBoard
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

[**PlatformDriftBoard**](PlatformDriftBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformFleetByApp

> PlatformAppView GetPlatformFleetByApp(ctx, app).Env(env).Execute()

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
	// response from `GetPlatformFleetByApp`: PlatformAppView
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

[**PlatformAppView**](PlatformAppView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformHealth

> PlatformReadiness GetPlatformHealth(ctx).Execute()

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
	// response from `GetPlatformHealth`: PlatformReadiness
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformHealthRequest struct via the builder pattern


### Return type

[**PlatformReadiness**](PlatformReadiness.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformPipelines

> PlatformPipelineBoard GetPlatformPipelines(ctx).Execute()

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
	// response from `GetPlatformPipelines`: PlatformPipelineBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformPipelines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformPipelinesRequest struct via the builder pattern


### Return type

[**PlatformPipelineBoard**](PlatformPipelineBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjects

> PlatformProjectBoard GetPlatformProjects(ctx).Org(org).Execute()

Answers every project the caller may see, with how its apps stand.



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
	org := "org_example" // string | Org names whose projects to read. Omitted, a SuperAdmin reads every owner's and anyone else reads their own; naming another org is an act-as only a SuperAdmin has, and naming a platform directory reads the platform's own. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjects(context.Background()).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjects`: PlatformProjectBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org names whose projects to read. Omitted, a SuperAdmin reads every owner&#39;s and anyone else reads their own; naming another org is an act-as only a SuperAdmin has, and naming a platform directory reads the platform&#39;s own. | 

### Return type

[**PlatformProjectBoard**](PlatformProjectBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProject

> PlatformProjectView GetPlatformProjectsByProject(ctx, project).Org(org).Execute()

Answers one project and every app in it.



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
	org := "org_example" // string | Org names the project's owner, defaulting to the caller's own scope — which for a SuperAdmin whose home is a brand org is the platform's. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.GetPlatformProjectsByProject(context.Background(), project).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.GetPlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlatformProjectsByProject`: PlatformProjectView
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

 **org** | **string** | Org names the project&#39;s owner, defaulting to the caller&#39;s own scope — which for a SuperAdmin whose home is a brand org is the platform&#39;s. | 

### Return type

[**PlatformProjectView**](PlatformProjectView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectApps

> []PlatformAppOut GetPlatformProjectsByProjectApps(ctx, project).Execute()

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
	// response from `GetPlatformProjectsByProjectApps`: []PlatformAppOut
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

[**[]PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByApp

> PlatformAppOut GetPlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()

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
	// response from `GetPlatformProjectsByProjectAppsByApp`: PlatformAppOut
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

[**PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeployments

> []PlatformDeploymentView GetPlatformProjectsByProjectAppsByAppDeployments(ctx, project, app).Execute()

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
	// response from `GetPlatformProjectsByProjectAppsByAppDeployments`: []PlatformDeploymentView
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

[**[]PlatformDeploymentView**](PlatformDeploymentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeploymentsById

> PlatformDeploymentView GetPlatformProjectsByProjectAppsByAppDeploymentsById(ctx, project, app, id).Execute()

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
	// response from `GetPlatformProjectsByProjectAppsByAppDeploymentsById`: PlatformDeploymentView
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

[**PlatformDeploymentView**](PlatformDeploymentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs

> PlatformDeployLogs GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(ctx, project, app, id).Execute()

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
	// response from `GetPlatformProjectsByProjectAppsByAppDeploymentsByIdLogs`: PlatformDeployLogs
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

[**PlatformDeployLogs**](PlatformDeployLogs.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformProjectsByProjectAppsByAppDomains

> []PlatformDomainView GetPlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()

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
	// response from `GetPlatformProjectsByProjectAppsByAppDomains`: []PlatformDomainView
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

[**[]PlatformDomainView**](PlatformDomainView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlatformReleases

> PlatformReleaseBoard GetPlatformReleases(ctx).Execute()

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
	// response from `GetPlatformReleases`: PlatformReleaseBoard
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.GetPlatformReleases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlatformReleasesRequest struct via the builder pattern


### Return type

[**PlatformReleaseBoard**](PlatformReleaseBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBuild

> PlatformRunnerBuildResp PostBuild(ctx).PlatformRunnerBuildReq(platformRunnerBuildReq).Execute()

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
	platformRunnerBuildReq := *openapiclient.NewPlatformRunnerBuildReq() // PlatformRunnerBuildReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostBuild(context.Background()).PlatformRunnerBuildReq(platformRunnerBuildReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostBuild``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBuild`: PlatformRunnerBuildResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostBuild`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformRunnerBuildReq** | [**PlatformRunnerBuildReq**](PlatformRunnerBuildReq.md) |  | 

### Return type

[**PlatformRunnerBuildResp**](PlatformRunnerBuildResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformApps

> DeclareResp PostPlatformApps(ctx).DeclareReq(declareReq).Execute()

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
	declareReq := *openapiclient.NewDeclareReq() // DeclareReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformApps(context.Background()).DeclareReq(declareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformApps`: DeclareResp
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **declareReq** | [**DeclareReq**](DeclareReq.md) |  | 

### Return type

[**DeclareResp**](DeclareResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjects

> PlatformProjectWrite PostPlatformProjects(ctx).PlatformProjectCreate(platformProjectCreate).Execute()

Creates a project from the apps it starts with.



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
	platformProjectCreate := *openapiclient.NewPlatformProjectCreate() // PlatformProjectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjects(context.Background()).PlatformProjectCreate(platformProjectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjects`: PlatformProjectWrite
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformProjectCreate** | [**PlatformProjectCreate**](PlatformProjectCreate.md) |  | 

### Return type

[**PlatformProjectWrite**](PlatformProjectWrite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectApps

> PlatformAppOut PostPlatformProjectsByProjectApps(ctx, project).PlatformCreateAppReq(platformCreateAppReq).Execute()

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
	platformCreateAppReq := *openapiclient.NewPlatformCreateAppReq() // PlatformCreateAppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectApps(context.Background(), project).PlatformCreateAppReq(platformCreateAppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectApps`: PlatformAppOut
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

 **platformCreateAppReq** | [**PlatformCreateAppReq**](PlatformCreateAppReq.md) |  | 

### Return type

[**PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDeploy

> PlatformDeploymentView PostPlatformProjectsByProjectAppsByAppDeploy(ctx, project, app).PlatformDeployReq(platformDeployReq).Execute()

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
	platformDeployReq := *openapiclient.NewPlatformDeployReq() // PlatformDeployReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppDeploy(context.Background(), project, app).PlatformDeployReq(platformDeployReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppDeploy`: PlatformDeploymentView
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


 **platformDeployReq** | [**PlatformDeployReq**](PlatformDeployReq.md) |  | 

### Return type

[**PlatformDeploymentView**](PlatformDeploymentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDomains

> PlatformDomainView PostPlatformProjectsByProjectAppsByAppDomains(ctx, project, app).PlatformAddDomainReq(platformAddDomainReq).Execute()

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
	platformAddDomainReq := *openapiclient.NewPlatformAddDomainReq() // PlatformAddDomainReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).PlatformAddDomainReq(platformAddDomainReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppDomains`: PlatformDomainView
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


 **platformAddDomainReq** | [**PlatformAddDomainReq**](PlatformAddDomainReq.md) |  | 

### Return type

[**PlatformDomainView**](PlatformDomainView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify

> PlatformDomainView PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify(ctx, project, app, host).Execute()

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
	// response from `PostPlatformProjectsByProjectAppsByAppDomainsByHostVerify`: PlatformDomainView
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

[**PlatformDomainView**](PlatformDomainView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppPreview

> PlatformPreviewView PostPlatformProjectsByProjectAppsByAppPreview(ctx, project, app).PlatformPreviewReq(platformPreviewReq).Execute()

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
	platformPreviewReq := *openapiclient.NewPlatformPreviewReq() // PlatformPreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppPreview(context.Background(), project, app).PlatformPreviewReq(platformPreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppPreview`: PlatformPreviewView
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


 **platformPreviewReq** | [**PlatformPreviewReq**](PlatformPreviewReq.md) |  | 

### Return type

[**PlatformPreviewView**](PlatformPreviewView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppPromote

> PlatformDeploymentView PostPlatformProjectsByProjectAppsByAppPromote(ctx, project, app).PlatformPromoteReq(platformPromoteReq).Execute()

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
	platformPromoteReq := *openapiclient.NewPlatformPromoteReq() // PlatformPromoteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppPromote(context.Background(), project, app).PlatformPromoteReq(platformPromoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppPromote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppPromote`: PlatformDeploymentView
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


 **platformPromoteReq** | [**PlatformPromoteReq**](PlatformPromoteReq.md) |  | 

### Return type

[**PlatformDeploymentView**](PlatformDeploymentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppRollback

> PlatformDeploymentView PostPlatformProjectsByProjectAppsByAppRollback(ctx, project, app).PlatformRollbackReq(platformRollbackReq).Execute()

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
	platformRollbackReq := *openapiclient.NewPlatformRollbackReq() // PlatformRollbackReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformProjectsByProjectAppsByAppRollback(context.Background(), project, app).PlatformRollbackReq(platformRollbackReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformProjectsByProjectAppsByAppRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformProjectsByProjectAppsByAppRollback`: PlatformDeploymentView
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


 **platformRollbackReq** | [**PlatformRollbackReq**](PlatformRollbackReq.md) |  | 

### Return type

[**PlatformDeploymentView**](PlatformDeploymentView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppStart

> PlatformAppOut PostPlatformProjectsByProjectAppsByAppStart(ctx, project, app).Execute()

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
	// response from `PostPlatformProjectsByProjectAppsByAppStart`: PlatformAppOut
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

[**PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformProjectsByProjectAppsByAppStop

> PlatformAppOut PostPlatformProjectsByProjectAppsByAppStop(ctx, project, app).Execute()

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
	// response from `PostPlatformProjectsByProjectAppsByAppStop`: PlatformAppOut
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

[**PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlatformRun

> PlatformRunView PostPlatformRun(ctx).PlatformRunReq(platformRunReq).Execute()

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
	platformRunReq := *openapiclient.NewPlatformRunReq() // PlatformRunReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PostPlatformRun(context.Background()).PlatformRunReq(platformRunReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PostPlatformRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlatformRun`: PlatformRunView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PostPlatformRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlatformRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformRunReq** | [**PlatformRunReq**](PlatformRunReq.md) |  | 

### Return type

[**PlatformRunView**](PlatformRunView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlatformAppsByAppProject

> PlatformProjectWrite PutPlatformAppsByAppProject(ctx, app).PlatformAppMove(platformAppMove).Execute()

Moves an app to a project.



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
	app := "app_example" // string | App is the declaration's name, from the path.
	platformAppMove := *openapiclient.NewPlatformAppMove() // PlatformAppMove | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PutPlatformAppsByAppProject(context.Background(), app).PlatformAppMove(platformAppMove).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PutPlatformAppsByAppProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlatformAppsByAppProject`: PlatformProjectWrite
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PutPlatformAppsByAppProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | App is the declaration&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlatformAppsByAppProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformAppMove** | [**PlatformAppMove**](PlatformAppMove.md) |  | 

### Return type

[**PlatformProjectWrite**](PlatformProjectWrite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlatformProjectsByProject

> PlatformProjectWrite PutPlatformProjectsByProject(ctx, project).PlatformProjectRename(platformProjectRename).Execute()

Renames a project.



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
	project := "project_example" // string | Project is the project to rename, from the path.
	platformProjectRename := *openapiclient.NewPlatformProjectRename() // PlatformProjectRename | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PutPlatformProjectsByProject(context.Background(), project).PlatformProjectRename(platformProjectRename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PutPlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlatformProjectsByProject`: PlatformProjectWrite
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PutPlatformProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the project to rename, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlatformProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformProjectRename** | [**PlatformProjectRename**](PlatformProjectRename.md) |  | 

### Return type

[**PlatformProjectWrite**](PlatformProjectWrite.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlatformProjectsByProjectAppsByAppEnv

> PlatformAppOut PutPlatformProjectsByProjectAppsByAppEnv(ctx, project, app).PlatformSetEnvReq(platformSetEnvReq).Execute()

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
	platformSetEnvReq := *openapiclient.NewPlatformSetEnvReq() // PlatformSetEnvReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PutPlatformProjectsByProjectAppsByAppEnv(context.Background(), project, app).PlatformSetEnvReq(platformSetEnvReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PutPlatformProjectsByProjectAppsByAppEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlatformProjectsByProjectAppsByAppEnv`: PlatformAppOut
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


 **platformSetEnvReq** | [**PlatformSetEnvReq**](PlatformSetEnvReq.md) |  | 

### Return type

[**PlatformAppOut**](PlatformAppOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

