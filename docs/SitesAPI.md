# \SitesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsBuildSite**](SitesAPI.md#ProjectsBuildSite) | **Post** /v1/sites | Build a site from a brief and deploy it
[**ProjectsDeploySite**](SitesAPI.md#ProjectsDeploySite) | **Post** /v1/sites/deploy | Deploy a raw file manifest
[**ProjectsListSites**](SitesAPI.md#ProjectsListSites) | **Get** /v1/sites | List the org&#39;s live sites



## ProjectsBuildSite

> ProjectsSiteDeployResult ProjectsBuildSite(ctx).ProjectsBuildSiteRequest(projectsBuildSiteRequest).Execute()

Build a site from a brief and deploy it



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
	projectsBuildSiteRequest := *openapiclient.NewProjectsBuildSiteRequest("Brief_example") // ProjectsBuildSiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.ProjectsBuildSite(context.Background()).ProjectsBuildSiteRequest(projectsBuildSiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.ProjectsBuildSite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsBuildSite`: ProjectsSiteDeployResult
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.ProjectsBuildSite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsBuildSiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsBuildSiteRequest** | [**ProjectsBuildSiteRequest**](ProjectsBuildSiteRequest.md) |  | 

### Return type

[**ProjectsSiteDeployResult**](ProjectsSiteDeployResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDeploySite

> ProjectsSiteDeployResult ProjectsDeploySite(ctx).ProjectsDeploySiteRequest(projectsDeploySiteRequest).Execute()

Deploy a raw file manifest



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
	projectsDeploySiteRequest := *openapiclient.NewProjectsDeploySiteRequest([]openapiclient.ProjectsSiteFile{*openapiclient.NewProjectsSiteFile("Path_example", "Content_example")}) // ProjectsDeploySiteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SitesAPI.ProjectsDeploySite(context.Background()).ProjectsDeploySiteRequest(projectsDeploySiteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.ProjectsDeploySite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsDeploySite`: ProjectsSiteDeployResult
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.ProjectsDeploySite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDeploySiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectsDeploySiteRequest** | [**ProjectsDeploySiteRequest**](ProjectsDeploySiteRequest.md) |  | 

### Return type

[**ProjectsSiteDeployResult**](ProjectsSiteDeployResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsListSites

> []ProjectsSite ProjectsListSites(ctx).Execute()

List the org's live sites



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
	resp, r, err := apiClient.SitesAPI.ProjectsListSites(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SitesAPI.ProjectsListSites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsListSites`: []ProjectsSite
	fmt.Fprintf(os.Stdout, "Response from `SitesAPI.ProjectsListSites`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsListSitesRequest struct via the builder pattern


### Return type

[**[]ProjectsSite**](ProjectsSite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

