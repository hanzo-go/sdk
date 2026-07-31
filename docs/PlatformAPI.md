# \PlatformAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1PlatformProjectsByProjectAppsByApp**](PlatformAPI.md#CloudDeleteV1PlatformProjectsByProjectAppsByApp) | **Delete** /v1/platform/projects/{project}/apps/{app} | 
[**CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost**](PlatformAPI.md#CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost) | **Delete** /v1/platform/projects/{project}/apps/{app}/domains/{host} | 
[**CloudDeleteV1PlatformSitesBySlug**](PlatformAPI.md#CloudDeleteV1PlatformSitesBySlug) | **Delete** /v1/platform/sites/{slug} | 
[**CloudDeleteV1PlatformSitesBySlugDomainsByHost**](PlatformAPI.md#CloudDeleteV1PlatformSitesBySlugDomainsByHost) | **Delete** /v1/platform/sites/{slug}/domains/{host} | 
[**CloudGetV1PlatformFleet**](PlatformAPI.md#CloudGetV1PlatformFleet) | **Get** /v1/platform/fleet | 
[**CloudGetV1PlatformFleetByApp**](PlatformAPI.md#CloudGetV1PlatformFleetByApp) | **Get** /v1/platform/fleet/{app} | 
[**CloudGetV1PlatformHealth**](PlatformAPI.md#CloudGetV1PlatformHealth) | **Get** /v1/platform/health | 
[**CloudGetV1PlatformProjects**](PlatformAPI.md#CloudGetV1PlatformProjects) | **Get** /v1/platform/projects | 
[**CloudGetV1PlatformProjectsByProject**](PlatformAPI.md#CloudGetV1PlatformProjectsByProject) | **Get** /v1/platform/projects/{project} | 
[**CloudGetV1PlatformProjectsByProjectApps**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectApps) | **Get** /v1/platform/projects/{project}/apps | 
[**CloudGetV1PlatformProjectsByProjectAppsByApp**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectAppsByApp) | **Get** /v1/platform/projects/{project}/apps/{app} | 
[**CloudGetV1PlatformProjectsByProjectAppsByAppDeployments**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectAppsByAppDeployments) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments | 
[**CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id} | 
[**CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs) | **Get** /v1/platform/projects/{project}/apps/{app}/deployments/{id}/logs | 
[**CloudGetV1PlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#CloudGetV1PlatformProjectsByProjectAppsByAppDomains) | **Get** /v1/platform/projects/{project}/apps/{app}/domains | 
[**CloudGetV1PlatformSites**](PlatformAPI.md#CloudGetV1PlatformSites) | **Get** /v1/platform/sites | 
[**CloudGetV1PlatformSitesBySlug**](PlatformAPI.md#CloudGetV1PlatformSitesBySlug) | **Get** /v1/platform/sites/{slug} | 
[**CloudGetV1PlatformSitesBySlugDeployments**](PlatformAPI.md#CloudGetV1PlatformSitesBySlugDeployments) | **Get** /v1/platform/sites/{slug}/deployments | 
[**CloudGetV1PlatformSitesBySlugDeploymentsById**](PlatformAPI.md#CloudGetV1PlatformSitesBySlugDeploymentsById) | **Get** /v1/platform/sites/{slug}/deployments/{id} | 
[**CloudGetV1PlatformSitesBySlugDomains**](PlatformAPI.md#CloudGetV1PlatformSitesBySlugDomains) | **Get** /v1/platform/sites/{slug}/domains | 
[**CloudGetV1PlatformSitesBySlugReleases**](PlatformAPI.md#CloudGetV1PlatformSitesBySlugReleases) | **Get** /v1/platform/sites/{slug}/releases | 
[**CloudPatchV1PlatformSitesBySlug**](PlatformAPI.md#CloudPatchV1PlatformSitesBySlug) | **Patch** /v1/platform/sites/{slug} | 
[**CloudPostV1PlatformFleetByAppDeploy**](PlatformAPI.md#CloudPostV1PlatformFleetByAppDeploy) | **Post** /v1/platform/fleet/{app}/deploy | 
[**CloudPostV1PlatformProjectsByProjectApps**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectApps) | **Post** /v1/platform/projects/{project}/apps | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppDeploy**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppDeploy) | **Post** /v1/platform/projects/{project}/apps/{app}/deploy | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppDomains**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppDomains) | **Post** /v1/platform/projects/{project}/apps/{app}/domains | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify) | **Post** /v1/platform/projects/{project}/apps/{app}/domains/{host}/verify | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppPreview**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppPreview) | **Post** /v1/platform/projects/{project}/apps/{app}/preview | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppPromote**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppPromote) | **Post** /v1/platform/projects/{project}/apps/{app}/promote | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppRollback**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppRollback) | **Post** /v1/platform/projects/{project}/apps/{app}/rollback | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppStart**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppStart) | **Post** /v1/platform/projects/{project}/apps/{app}/start | 
[**CloudPostV1PlatformProjectsByProjectAppsByAppStop**](PlatformAPI.md#CloudPostV1PlatformProjectsByProjectAppsByAppStop) | **Post** /v1/platform/projects/{project}/apps/{app}/stop | 
[**CloudPostV1PlatformSites**](PlatformAPI.md#CloudPostV1PlatformSites) | **Post** /v1/platform/sites | 
[**CloudPostV1PlatformSitesBySlugDeploy**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugDeploy) | **Post** /v1/platform/sites/{slug}/deploy | 
[**CloudPostV1PlatformSitesBySlugDomains**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugDomains) | **Post** /v1/platform/sites/{slug}/domains | 
[**CloudPostV1PlatformSitesBySlugDomainsByHostVerify**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugDomainsByHostVerify) | **Post** /v1/platform/sites/{slug}/domains/{host}/verify | 
[**CloudPostV1PlatformSitesBySlugPublish**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugPublish) | **Post** /v1/platform/sites/{slug}/publish | 
[**CloudPostV1PlatformSitesBySlugPurge**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugPurge) | **Post** /v1/platform/sites/{slug}/purge | 
[**CloudPostV1PlatformSitesBySlugReleases**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugReleases) | **Post** /v1/platform/sites/{slug}/releases | 
[**CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate**](PlatformAPI.md#CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate) | **Post** /v1/platform/sites/{slug}/releases/{release}/activate | 
[**CloudPutV1PlatformProjectsByProjectAppsByAppEnv**](PlatformAPI.md#CloudPutV1PlatformProjectsByProjectAppsByAppEnv) | **Put** /v1/platform/projects/{project}/apps/{app}/env | 



## CloudDeleteV1PlatformProjectsByProjectAppsByApp

> CloudDeleteV1PlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudDeleteV1PlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudDeleteV1PlatformProjectsByProjectAppsByApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


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


## CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost

> CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost(ctx, project, app, host).Execute()



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
	r, err := apiClient.PlatformAPI.CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PlatformProjectsByProjectAppsByAppDomainsByHostRequest struct via the builder pattern


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


## CloudDeleteV1PlatformSitesBySlug

> CloudDeleteV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudDeleteV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudDeleteV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PlatformSitesBySlugRequest struct via the builder pattern


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


## CloudDeleteV1PlatformSitesBySlugDomainsByHost

> CloudDeleteV1PlatformSitesBySlugDomainsByHost(ctx, slug, host).Execute()



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
	r, err := apiClient.PlatformAPI.CloudDeleteV1PlatformSitesBySlugDomainsByHost(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudDeleteV1PlatformSitesBySlugDomainsByHost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PlatformSitesBySlugDomainsByHostRequest struct via the builder pattern


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


## CloudGetV1PlatformFleet

> CloudGetV1PlatformFleet(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformFleetRequest struct via the builder pattern


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


## CloudGetV1PlatformFleetByApp

> CloudGetV1PlatformFleetByApp(ctx, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformFleetByApp(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformFleetByApp``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformFleetByAppRequest struct via the builder pattern


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


## CloudGetV1PlatformHealth

> CloudGetV1PlatformHealth(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformHealthRequest struct via the builder pattern


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


## CloudGetV1PlatformProjects

> []CloudProjectView CloudGetV1PlatformProjects(ctx).Execute()



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
	resp, r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlatformProjects`: []CloudProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudGetV1PlatformProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsRequest struct via the builder pattern


### Return type

[**[]CloudProjectView**](CloudProjectView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlatformProjectsByProject

> CloudProjectView CloudGetV1PlatformProjectsByProject(ctx, project).Execute()



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
	resp, r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlatformProjectsByProject`: CloudProjectView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudGetV1PlatformProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudProjectView**](CloudProjectView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlatformProjectsByProjectApps

> []CloudAppView CloudGetV1PlatformProjectsByProjectApps(ctx, project).Execute()



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
	resp, r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectApps(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlatformProjectsByProjectApps`: []CloudAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudGetV1PlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CloudAppView**](CloudAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlatformProjectsByProjectAppsByApp

> CloudAppView CloudGetV1PlatformProjectsByProjectAppsByApp(ctx, project, app).Execute()



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
	resp, r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByApp(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlatformProjectsByProjectAppsByApp`: CloudAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsByAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudAppView**](CloudAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlatformProjectsByProjectAppsByAppDeployments

> CloudGetV1PlatformProjectsByProjectAppsByAppDeployments(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeployments(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeployments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsRequest struct via the builder pattern


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


## CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById

> CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById(ctx, project, app, id).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdRequest struct via the builder pattern


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


## CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs

> CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(ctx, project, app, id).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs(context.Background(), project, app, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsByAppDeploymentsByIdLogsRequest struct via the builder pattern


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


## CloudGetV1PlatformProjectsByProjectAppsByAppDomains

> CloudGetV1PlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


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


## CloudGetV1PlatformSites

> CloudGetV1PlatformSites(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesRequest struct via the builder pattern


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


## CloudGetV1PlatformSitesBySlug

> CloudGetV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesBySlugRequest struct via the builder pattern


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


## CloudGetV1PlatformSitesBySlugDeployments

> CloudGetV1PlatformSitesBySlugDeployments(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSitesBySlugDeployments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSitesBySlugDeployments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesBySlugDeploymentsRequest struct via the builder pattern


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


## CloudGetV1PlatformSitesBySlugDeploymentsById

> CloudGetV1PlatformSitesBySlugDeploymentsById(ctx, slug, id).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSitesBySlugDeploymentsById(context.Background(), slug, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSitesBySlugDeploymentsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesBySlugDeploymentsByIdRequest struct via the builder pattern


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


## CloudGetV1PlatformSitesBySlugDomains

> CloudGetV1PlatformSitesBySlugDomains(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSitesBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSitesBySlugDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## CloudGetV1PlatformSitesBySlugReleases

> CloudGetV1PlatformSitesBySlugReleases(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudGetV1PlatformSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudGetV1PlatformSitesBySlugReleases``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1PlatformSitesBySlugReleasesRequest struct via the builder pattern


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


## CloudPatchV1PlatformSitesBySlug

> CloudPatchV1PlatformSitesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPatchV1PlatformSitesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPatchV1PlatformSitesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPatchV1PlatformSitesBySlugRequest struct via the builder pattern


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


## CloudPostV1PlatformFleetByAppDeploy

> CloudPostV1PlatformFleetByAppDeploy(ctx, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformFleetByAppDeploy(context.Background(), app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformFleetByAppDeploy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformFleetByAppDeployRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectApps

> CloudAppView CloudPostV1PlatformProjectsByProjectApps(ctx, project).CloudCreateAppReq(cloudCreateAppReq).Execute()



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
	cloudCreateAppReq := *openapiclient.NewCloudCreateAppReq() // CloudCreateAppReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectApps(context.Background(), project).CloudCreateAppReq(cloudCreateAppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PlatformProjectsByProjectApps`: CloudAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudPostV1PlatformProjectsByProjectApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCreateAppReq** | [**CloudCreateAppReq**](CloudCreateAppReq.md) |  | 

### Return type

[**CloudAppView**](CloudAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PlatformProjectsByProjectAppsByAppDeploy

> CloudPostV1PlatformProjectsByProjectAppsByAppDeploy(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDeploy(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDeploy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppDeployRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppDomains

> CloudPostV1PlatformProjectsByProjectAppsByAppDomains(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDomains(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppDomainsRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify

> CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify(ctx, project, app, host).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify(context.Background(), project, app, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerify``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppDomainsByHostVerifyRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppPreview

> CloudPostV1PlatformProjectsByProjectAppsByAppPreview(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppPreview(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppPreview``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppPreviewRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppPromote

> CloudPostV1PlatformProjectsByProjectAppsByAppPromote(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppPromote(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppPromote``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppPromoteRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppRollback

> CloudPostV1PlatformProjectsByProjectAppsByAppRollback(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppRollback(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppRollback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppRollbackRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppStart

> CloudPostV1PlatformProjectsByProjectAppsByAppStart(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppStart(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppStart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppStartRequest struct via the builder pattern


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


## CloudPostV1PlatformProjectsByProjectAppsByAppStop

> CloudPostV1PlatformProjectsByProjectAppsByAppStop(ctx, project, app).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppStop(context.Background(), project, app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformProjectsByProjectAppsByAppStop``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformProjectsByProjectAppsByAppStopRequest struct via the builder pattern


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


## CloudPostV1PlatformSites

> CloudPostV1PlatformSites(ctx).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugDeploy

> CloudPostV1PlatformSitesBySlugDeploy(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugDeploy(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugDeploy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugDeployRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugDomains

> CloudPostV1PlatformSitesBySlugDomains(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugDomains(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugDomains``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugDomainsRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugDomainsByHostVerify

> CloudPostV1PlatformSitesBySlugDomainsByHostVerify(ctx, slug, host).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugDomainsByHostVerify(context.Background(), slug, host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugDomainsByHostVerify``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugDomainsByHostVerifyRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugPublish

> CloudPostV1PlatformSitesBySlugPublish(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugPublish(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugPublish``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugPublishRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugPurge

> CloudPostV1PlatformSitesBySlugPurge(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugPurge(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugPurge``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugPurgeRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugReleases

> CloudPostV1PlatformSitesBySlugReleases(ctx, slug).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugReleases(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugReleases``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugReleasesRequest struct via the builder pattern


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


## CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate

> CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate(ctx, slug, release).Execute()



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
	r, err := apiClient.PlatformAPI.CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate(context.Background(), slug, release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPostV1PlatformSitesBySlugReleasesByReleaseActivate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1PlatformSitesBySlugReleasesByReleaseActivateRequest struct via the builder pattern


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


## CloudPutV1PlatformProjectsByProjectAppsByAppEnv

> CloudAppView CloudPutV1PlatformProjectsByProjectAppsByAppEnv(ctx, project, app).CloudSetEnvReq(cloudSetEnvReq).Execute()



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
	cloudSetEnvReq := *openapiclient.NewCloudSetEnvReq() // CloudSetEnvReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformAPI.CloudPutV1PlatformProjectsByProjectAppsByAppEnv(context.Background(), project, app).CloudSetEnvReq(cloudSetEnvReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformAPI.CloudPutV1PlatformProjectsByProjectAppsByAppEnv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1PlatformProjectsByProjectAppsByAppEnv`: CloudAppView
	fmt.Fprintf(os.Stdout, "Response from `PlatformAPI.CloudPutV1PlatformProjectsByProjectAppsByAppEnv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 
**app** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1PlatformProjectsByProjectAppsByAppEnvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudSetEnvReq** | [**CloudSetEnvReq**](CloudSetEnvReq.md) |  | 

### Return type

[**CloudAppView**](CloudAppView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

