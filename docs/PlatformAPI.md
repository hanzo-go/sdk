# \PlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDeleteV1PlatformProjectsByProjectAppsByApp**](PlatformAPI.md#PlatformDeleteV1PlatformProjectsByProjectAppsByApp) | **Delete** /v1/platform/projects/{project}/apps/{app} | 
[**PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost**](PlatformAPI.md#PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost) | **Delete** /v1/platform/projects/{project}/apps/{app}/domains/{host} | 
[**PlatformDeleteV1PlatformSitesBySlug**](PlatformAPI.md#PlatformDeleteV1PlatformSitesBySlug) | **Delete** /v1/platform/sites/{slug} | 
[**PlatformDeleteV1PlatformSitesBySlugDomainsByHost**](PlatformAPI.md#PlatformDeleteV1PlatformSitesBySlugDomainsByHost) | **Delete** /v1/platform/sites/{slug}/domains/{host} | 
[**PlatformGetV1PlatformFleet**](PlatformAPI.md#PlatformGetV1PlatformFleet) | **Get** /v1/platform/fleet | 
[**PlatformGetV1PlatformFleetByApp**](PlatformAPI.md#PlatformGetV1PlatformFleetByApp) | **Get** /v1/platform/fleet/{app} | 
[**PlatformGetV1PlatformHealth**](PlatformAPI.md#PlatformGetV1PlatformHealth) | **Get** /v1/platform/health | 
[**PlatformGetV1PlatformProjects**](PlatformAPI.md#PlatformGetV1PlatformProjects) | **Get** /v1/platform/projects | 
[**PlatformGetV1PlatformProjectsByProject**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProject) | **Get** /v1/platform/projects/{project} | 
[**PlatformGetV1PlatformProjectsByProjectApps**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectApps) | **Get** /v1/platform/projects/{project}/apps | 
[**PlatformGetV1PlatformProjectsByProjectAppsByApp**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectAppsByApp) | **Get** /v1/platform/projects/{project}/apps/{app} | 
[**PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments | 
[**PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id} | 
[**PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id}/logs | 
[**PlatformGetV1PlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#PlatformGetV1PlatformProjectsByProjectAppsByAppDomains) | **Get** /v1/platform/projects/{project}/apps/{app}/domains | 
[**PlatformGetV1PlatformSites**](PlatformAPI.md#PlatformGetV1PlatformSites) | **Get** /v1/platform/sites | 
[**PlatformGetV1PlatformSitesBySlug**](PlatformAPI.md#PlatformGetV1PlatformSitesBySlug) | **Get** /v1/platform/sites/{slug} | 
[**PlatformGetV1PlatformSitesBySlugDeployments**](PlatformAPI.md#PlatformGetV1PlatformSitesBySlugDeployments) | **Get** /v1/platform/sites/{slug}/deployments | 
[**PlatformGetV1PlatformSitesBySlugDeploymentsById**](PlatformAPI.md#PlatformGetV1PlatformSitesBySlugDeploymentsById) | **Get** /v1/platform/sites/{slug}/deployments/{id} | 
[**PlatformGetV1PlatformSitesBySlugDomains**](PlatformAPI.md#PlatformGetV1PlatformSitesBySlugDomains) | **Get** /v1/platform/sites/{slug}/domains | 
[**PlatformGetV1PlatformSitesBySlugReleases**](PlatformAPI.md#PlatformGetV1PlatformSitesBySlugReleases) | **Get** /v1/platform/sites/{slug}/releases | 
[**PlatformPatchV1PlatformSitesBySlug**](PlatformAPI.md#PlatformPatchV1PlatformSitesBySlug) | **Patch** /v1/platform/sites/{slug} | 
[**PlatformPostV1PlatformFleetByAppDeploy**](PlatformAPI.md#PlatformPostV1PlatformFleetByAppDeploy) | **Post** /v1/platform/fleet/{app}/deploy | 
[**PlatformPostV1PlatformProjectsByProjectApps**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectApps) | **Post** /v1/platform/projects/{project}/apps | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy) | **Post** /v1/platform/projects/{project}/apps/{app}/deploy | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppDomains) | **Post** /v1/platform/projects/{project}/apps/{app}/domains | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify) | **Post** /v1/platform/projects/{project}/apps/{app}/domains/{host}/verify | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppPreview**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppPreview) | **Post** /v1/platform/projects/{project}/apps/{app}/preview | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppPromote**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppPromote) | **Post** /v1/platform/projects/{project}/apps/{app}/promote | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppRollback**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppRollback) | **Post** /v1/platform/projects/{project}/apps/{app}/rollback | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppStart**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppStart) | **Post** /v1/platform/projects/{project}/apps/{app}/start | 
[**PlatformPostV1PlatformProjectsByProjectAppsByAppStop**](PlatformAPI.md#PlatformPostV1PlatformProjectsByProjectAppsByAppStop) | **Post** /v1/platform/projects/{project}/apps/{app}/stop | 
[**PlatformPostV1PlatformSites**](PlatformAPI.md#PlatformPostV1PlatformSites) | **Post** /v1/platform/sites | 
[**PlatformPostV1PlatformSitesBySlugDeploy**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugDeploy) | **Post** /v1/platform/sites/{slug}/deploy | 
[**PlatformPostV1PlatformSitesBySlugDomains**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugDomains) | **Post** /v1/platform/sites/{slug}/domains | 
[**PlatformPostV1PlatformSitesBySlugDomainsByHostVerify**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugDomainsByHostVerify) | **Post** /v1/platform/sites/{slug}/domains/{host}/verify | 
[**PlatformPostV1PlatformSitesBySlugPublish**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugPublish) | **Post** /v1/platform/sites/{slug}/publish | 
[**PlatformPostV1PlatformSitesBySlugPurge**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugPurge) | **Post** /v1/platform/sites/{slug}/purge | 
[**PlatformPostV1PlatformSitesBySlugReleases**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugReleases) | **Post** /v1/platform/sites/{slug}/releases | 
[**PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate**](PlatformAPI.md#PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate) | **Post** /v1/platform/sites/{slug}/releases/{release}/activate | 
[**PlatformPutV1PlatformProjectsByProjectAppsByAppEnv**](PlatformAPI.md#PlatformPutV1PlatformProjectsByProjectAppsByAppEnv) | **Put** /v1/platform/projects/{project}/apps/{app}/env | 



## PlatformDeleteV1PlatformProjectsByProjectAppsByApp

> PlatformDeleteV1PlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformDeleteV1PlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformDeleteV1PlatformProjectsByProjectAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeleteV1PlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


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


## PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost

> PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost(ctx, project, app, host).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 
	host := "host_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHostRequest struct via the builder pattern


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


## PlatformDeleteV1PlatformSitesBySlug

> PlatformDeleteV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformDeleteV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformDeleteV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformDeleteV1PlatformSitesBySlugRequest struct via the builder pattern


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


## PlatformDeleteV1PlatformSitesBySlugDomainsByHost

> PlatformDeleteV1PlatformSitesBySlugDomainsByHost(ctx, slug, host).Execute()



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
	host := "host_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformDeleteV1PlatformSitesBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformDeleteV1PlatformSitesBySlugDomainsByHost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeleteV1PlatformSitesBySlugDomainsByHostRequest struct via the builder pattern


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


## PlatformGetV1PlatformFleet

> PlatformGetV1PlatformFleet(ctx).Env(env).Health(health).Drift(drift).Execute()



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
	env := "env_example" // string | Only rows in this environment (namespace). (optional)
	health := "health_example" // string | Only rows at this health: green | yellow | red. (optional)
	drift := "drift_example" // string | 1 or true — only rows whose live state differs from declared. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformFleet(context.Background()).Env(env).Health(health).Drift(drift).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformFleetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** | Only rows in this environment (namespace). | 
 **health** | **string** | Only rows at this health: green | yellow | red. | 
 **drift** | **string** | 1 or true — only rows whose live state differs from declared. | 

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


## PlatformGetV1PlatformFleetByApp

> PlatformGetV1PlatformFleetByApp(ctx, app).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformFleetByApp(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformFleetByApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformFleetByAppRequest struct via the builder pattern


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


## PlatformGetV1PlatformHealth

> PlatformGetV1PlatformHealth(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformHealthRequest struct via the builder pattern


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


## PlatformGetV1PlatformProjects

> []PlatformProjectView PlatformGetV1PlatformProjects(ctx).Execute()



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
	resp, r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGetV1PlatformProjects`: []PlatformProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformGetV1PlatformProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsRequest struct via the builder pattern


### Return type

[**[]PlatformProjectView**](PlatformProjectView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformGetV1PlatformProjectsByProject

> PlatformProjectView PlatformGetV1PlatformProjectsByProject(ctx, project).Execute()



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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGetV1PlatformProjectsByProject`: PlatformProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformGetV1PlatformProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlatformProjectView**](PlatformProjectView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformGetV1PlatformProjectsByProjectApps

> []PlatformAppView PlatformGetV1PlatformProjectsByProjectApps(ctx, project).Execute()



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
	project := "project_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectApps(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGetV1PlatformProjectsByProjectApps`: []PlatformAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformGetV1PlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PlatformAppView**](PlatformAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformGetV1PlatformProjectsByProjectAppsByApp

> PlatformAppView PlatformGetV1PlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGetV1PlatformProjectsByProjectAppsByApp`: PlatformAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlatformAppView**](PlatformAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments

> PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsRequest struct via the builder pattern


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


## PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById

> PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById(ctx, project, app, id).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdRequest struct via the builder pattern


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


## PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs

> PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(ctx, project, app, id).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogsRequest struct via the builder pattern


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


## PlatformGetV1PlatformProjectsByProjectAppsByAppDomains

> PlatformGetV1PlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


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


## PlatformGetV1PlatformSites

> PlatformGetV1PlatformSites(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesRequest struct via the builder pattern


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


## PlatformGetV1PlatformSitesBySlug

> PlatformGetV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesBySlugRequest struct via the builder pattern


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


## PlatformGetV1PlatformSitesBySlugDeployments

> PlatformGetV1PlatformSitesBySlugDeployments(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSitesBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSitesBySlugDeployments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesBySlugDeploymentsRequest struct via the builder pattern


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


## PlatformGetV1PlatformSitesBySlugDeploymentsById

> PlatformGetV1PlatformSitesBySlugDeploymentsById(ctx, slug, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSitesBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSitesBySlugDeploymentsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesBySlugDeploymentsByIdRequest struct via the builder pattern


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


## PlatformGetV1PlatformSitesBySlugDomains

> PlatformGetV1PlatformSitesBySlugDomains(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSitesBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSitesBySlugDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## PlatformGetV1PlatformSitesBySlugReleases

> PlatformGetV1PlatformSitesBySlugReleases(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformGetV1PlatformSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformGetV1PlatformSitesBySlugReleases``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformGetV1PlatformSitesBySlugReleasesRequest struct via the builder pattern


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


## PlatformPatchV1PlatformSitesBySlug

> PlatformPatchV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPatchV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPatchV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPatchV1PlatformSitesBySlugRequest struct via the builder pattern


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


## PlatformPostV1PlatformFleetByAppDeploy

> PlatformPostV1PlatformFleetByAppDeploy(ctx, app).Env(env).Execute()



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
	env := "env_example" // string | Environment (namespace) holding the Deployment to restart. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformFleetByAppDeploy(context.Background(), app).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformFleetByAppDeploy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformFleetByAppDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Environment (namespace) holding the Deployment to restart. | 

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


## PlatformPostV1PlatformProjectsByProjectApps

> PlatformAppView PlatformPostV1PlatformProjectsByProjectApps(ctx, project).PlatformCreateAppReq(platformCreateAppReq).Execute()



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
	project := "project_example" // string | 
	platformCreateAppReq := *openapiclient.NewPlatformCreateAppReq() // PlatformCreateAppReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectApps(context.Background(), project).PlatformCreateAppReq(platformCreateAppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformPostV1PlatformProjectsByProjectApps`: PlatformAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformPostV1PlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platformCreateAppReq** | [**PlatformCreateAppReq**](PlatformCreateAppReq.md) |  | 

### Return type

[**PlatformAppView**](PlatformAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy

> PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppDeployRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppDomains

> PlatformPostV1PlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify

> PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify(ctx, project, app, host).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 
	host := "host_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerifyRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppPreview

> PlatformPostV1PlatformProjectsByProjectAppsByAppPreview(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppPreview(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppPreviewRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppPromote

> PlatformPostV1PlatformProjectsByProjectAppsByAppPromote(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppPromote(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppPromote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppPromoteRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppRollback

> PlatformPostV1PlatformProjectsByProjectAppsByAppRollback(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppRollback(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppRollbackRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppStart

> PlatformPostV1PlatformProjectsByProjectAppsByAppStart(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppStart(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppStartRequest struct via the builder pattern


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


## PlatformPostV1PlatformProjectsByProjectAppsByAppStop

> PlatformPostV1PlatformProjectsByProjectAppsByAppStop(ctx, project, app).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppStop(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformProjectsByProjectAppsByAppStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformProjectsByProjectAppsByAppStopRequest struct via the builder pattern


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


## PlatformPostV1PlatformSites

> PlatformPostV1PlatformSites(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugDeploy

> PlatformPostV1PlatformSitesBySlugDeploy(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugDeploy(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugDeploy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugDeployRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugDomains

> PlatformPostV1PlatformSitesBySlugDomains(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugDomainsByHostVerify

> PlatformPostV1PlatformSitesBySlugDomainsByHostVerify(ctx, slug, host).Execute()



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
	host := "host_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugDomainsByHostVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugDomainsByHostVerifyRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugPublish

> PlatformPostV1PlatformSitesBySlugPublish(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugPublish(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugPublish``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugPublishRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugPurge

> PlatformPostV1PlatformSitesBySlugPurge(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugPurge``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugPurgeRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugReleases

> PlatformPostV1PlatformSitesBySlugReleases(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugReleases``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugReleasesRequest struct via the builder pattern


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


## PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate

> PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()



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
	release := "release_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformAPI.PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPostV1PlatformSitesBySlugReleasesByReleaseActivate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**release** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPostV1PlatformSitesBySlugReleasesByReleaseActivateRequest struct via the builder pattern


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


## PlatformPutV1PlatformProjectsByProjectAppsByAppEnv

> PlatformAppView PlatformPutV1PlatformProjectsByProjectAppsByAppEnv(ctx, project, app).PlatformSetEnvReq(platformSetEnvReq).Execute()



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
	project := "project_example" // string | 
	app := "app_example" // string | 
	platformSetEnvReq := *openapiclient.NewPlatformSetEnvReq() // PlatformSetEnvReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.PlatformPutV1PlatformProjectsByProjectAppsByAppEnv(context.Background(), project, app).PlatformSetEnvReq(platformSetEnvReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.PlatformPutV1PlatformProjectsByProjectAppsByAppEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformPutV1PlatformProjectsByProjectAppsByAppEnv`: PlatformAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.PlatformPutV1PlatformProjectsByProjectAppsByAppEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformPutV1PlatformProjectsByProjectAppsByAppEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **platformSetEnvReq** | [**PlatformSetEnvReq**](PlatformSetEnvReq.md) |  | 

### Return type

[**PlatformAppView**](PlatformAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

