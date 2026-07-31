# \DeployAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1DeployAccountCanIByWildcard1**](DeployAPI.md#CloudGetV1DeployAccountCanIByWildcard1) | **Get** /v1/deploy/account/can-i/{wildcard1} | 
[**CloudGetV1DeployApplications**](DeployAPI.md#CloudGetV1DeployApplications) | **Get** /v1/deploy/applications | ListDeployApplications returns the fleet as an argocd ApplicationList: one projected Application per operator App CR, carrying the image tag the CR DECLARES, the tag actually RUNNING in the cluster&#39;s Deployment, the reconciled health, and the sync verdict those two produce (declared &#x3D;&#x3D; running ⇒ Synced, both known and different ⇒ OutOfSync, either unknown ⇒ Unknown).
[**CloudGetV1DeployApplicationsName**](DeployAPI.md#CloudGetV1DeployApplicationsName) | **Get** /v1/deploy/applications/{name} | GetDeployApplication returns ONE projected argocd Application by name, with status.resources filled in from its reconciled resource tree — which is what makes it the detail view rather than a row of the list.
[**CloudGetV1DeployApplicationsNameResourceTree**](DeployAPI.md#CloudGetV1DeployApplicationsNameResourceTree) | **Get** /v1/deploy/applications/{name}/resource-tree | GetDeployResourceTree returns one application&#39;s argocd ApplicationTree: the objects the operator reconciled from its App CR, reached by ownerRef — the Deployment and, under it, the ReplicaSet and Pods, plus the Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget and ConfigMaps it owns — each node carrying its parent edges and its health.
[**CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata**](DeployAPI.md#CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata) | **Get** /v1/deploy/applications/{name}/revisions/{revision}/metadata | GetDeployRevisionMetadata returns the argocd RevisionMetadata for one revision of one application — what the detail view shows beside a revision.
[**CloudGetV1DeployApplicationsNameSyncwindows**](DeployAPI.md#CloudGetV1DeployApplicationsNameSyncwindows) | **Get** /v1/deploy/applications/{name}/syncwindows | GetDeploySyncWindows returns one application&#39;s argocd ApplicationSyncWindowState — the answer to \&quot;is anything blocking a sync of this application right now?\&quot;.
[**CloudGetV1DeployCallback**](DeployAPI.md#CloudGetV1DeployCallback) | **Get** /v1/deploy/callback | 
[**CloudGetV1DeployClusters**](DeployAPI.md#CloudGetV1DeployClusters) | **Get** /v1/deploy/clusters | ListDeployClusters returns the argocd ClusterList of the destinations the caller&#39;s applications reconcile into: one entry per distinct destination server, carrying the count of applications reconciling into it.
[**CloudGetV1DeployGitops**](DeployAPI.md#CloudGetV1DeployGitops) | **Get** /v1/deploy/gitops | GetDeployGitOps lists every Hanzo CD Application in the cluster: the git source each one polls, the commit it last APPLIED, how its last sync operation ended, and its recent deploy history — newest deploy first, ordered by namespace then name.
[**CloudGetV1DeployHealth**](DeployAPI.md#CloudGetV1DeployHealth) | **Get** /v1/deploy/health | 
[**CloudGetV1DeployLogin**](DeployAPI.md#CloudGetV1DeployLogin) | **Get** /v1/deploy/login | 
[**CloudGetV1DeployProjects**](DeployAPI.md#CloudGetV1DeployProjects) | **Get** /v1/deploy/projects | ListDeployProjects returns the argocd AppProjectList this console groups and filters applications by.
[**CloudGetV1DeploySessionUserinfo**](DeployAPI.md#CloudGetV1DeploySessionUserinfo) | **Get** /v1/deploy/session/userinfo | GetDeploySession answers \&quot;is this browser signed in, and if not where does it sign in?\&quot; — the dashboard SPA&#39;s bootstrap question, and the only route on this plane that answers for an anonymous caller.
[**CloudGetV1DeploySettings**](DeployAPI.md#CloudGetV1DeploySettings) | **Get** /v1/deploy/settings | GetDeploySettings returns the argocd AuthSettings object the dashboard SPA awaits before its first render.
[**CloudGetV1DeployStreamApplications**](DeployAPI.md#CloudGetV1DeployStreamApplications) | **Get** /v1/deploy/stream/applications | 
[**CloudGetV1DeployStreamApplicationsByNameResourceTree**](DeployAPI.md#CloudGetV1DeployStreamApplicationsByNameResourceTree) | **Get** /v1/deploy/stream/applications/{name}/resource-tree | 
[**CloudGetV1DeployVersion**](DeployAPI.md#CloudGetV1DeployVersion) | **Get** /v1/deploy/version | GetDeployVersion returns the argocd VersionMessage the dashboard SPA reads at bootstrap.
[**CloudPostV1DeployApplicationsByNameRollback**](DeployAPI.md#CloudPostV1DeployApplicationsByNameRollback) | **Post** /v1/deploy/applications/{name}/rollback | 
[**CloudPostV1DeployApplicationsByNameSync**](DeployAPI.md#CloudPostV1DeployApplicationsByNameSync) | **Post** /v1/deploy/applications/{name}/sync | 
[**CloudPostV1DeployLogout**](DeployAPI.md#CloudPostV1DeployLogout) | **Post** /v1/deploy/logout | 
[**CloudPostV1DeployReconcile**](DeployAPI.md#CloudPostV1DeployReconcile) | **Post** /v1/deploy/reconcile | 



## CloudGetV1DeployAccountCanIByWildcard1

> CloudGetV1DeployAccountCanIByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeployAPI.CloudGetV1DeployAccountCanIByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployAccountCanIByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployAccountCanIByWildcard1Request struct via the builder pattern


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


## CloudGetV1DeployApplications

> CloudArgoAppList CloudGetV1DeployApplications(ctx).Execute()

ListDeployApplications returns the fleet as an argocd ApplicationList: one projected Application per operator App CR, carrying the image tag the CR DECLARES, the tag actually RUNNING in the cluster's Deployment, the reconciled health, and the sync verdict those two produce (declared == running ⇒ Synced, both known and different ⇒ OutOfSync, either unknown ⇒ Unknown).



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployApplications`: CloudArgoAppList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployApplicationsRequest struct via the builder pattern


### Return type

[**CloudArgoAppList**](CloudArgoAppList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployApplicationsName

> CloudArgoApp CloudGetV1DeployApplicationsName(ctx, name).Execute()

GetDeployApplication returns ONE projected argocd Application by name, with status.resources filled in from its reconciled resource tree — which is what makes it the detail view rather than a row of the list.



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
	name := "name_example" // string | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR's metadata.name satisfies that, and anything else is a 400 rather than a lookup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployApplicationsName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployApplicationsName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployApplicationsName`: CloudArgoApp
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployApplicationsName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployApplicationsNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudArgoApp**](CloudArgoApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployApplicationsNameResourceTree

> CloudArgoTree CloudGetV1DeployApplicationsNameResourceTree(ctx, name).Execute()

GetDeployResourceTree returns one application's argocd ApplicationTree: the objects the operator reconciled from its App CR, reached by ownerRef — the Deployment and, under it, the ReplicaSet and Pods, plus the Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget and ConfigMaps it owns — each node carrying its parent edges and its health.



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
	name := "name_example" // string | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR's metadata.name satisfies that, and anything else is a 400 rather than a lookup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployApplicationsNameResourceTree(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployApplicationsNameResourceTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployApplicationsNameResourceTree`: CloudArgoTree
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployApplicationsNameResourceTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployApplicationsNameResourceTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudArgoTree**](CloudArgoTree.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata

> CloudArgoRevisionMetadata CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata(ctx, name, revision).Execute()

GetDeployRevisionMetadata returns the argocd RevisionMetadata for one revision of one application — what the detail view shows beside a revision.



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
	name := "name_example" // string | Name is the application to read, from the path. It must be a DNS-1123 label.
	revision := "revision_example" // string | Revision is the revision to describe, from the path. The empty revision and \"HEAD\" both mean \"whatever this application currently declares\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata(context.Background(), name, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata`: CloudArgoRevisionMetadata
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployApplicationsNameRevisionsRevisionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label. | 
**revision** | **string** | Revision is the revision to describe, from the path. The empty revision and \&quot;HEAD\&quot; both mean \&quot;whatever this application currently declares\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployApplicationsNameRevisionsRevisionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudArgoRevisionMetadata**](CloudArgoRevisionMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployApplicationsNameSyncwindows

> CloudArgoSyncWindows CloudGetV1DeployApplicationsNameSyncwindows(ctx, name).Execute()

GetDeploySyncWindows returns one application's argocd ApplicationSyncWindowState — the answer to \"is anything blocking a sync of this application right now?\".



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
	name := "name_example" // string | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR's metadata.name satisfies that, and anything else is a 400 rather than a lookup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployApplicationsNameSyncwindows(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployApplicationsNameSyncwindows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployApplicationsNameSyncwindows`: CloudArgoSyncWindows
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployApplicationsNameSyncwindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployApplicationsNameSyncwindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudArgoSyncWindows**](CloudArgoSyncWindows.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployCallback

> CloudGetV1DeployCallback(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudGetV1DeployCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployCallbackRequest struct via the builder pattern


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


## CloudGetV1DeployClusters

> CloudArgoClusterList CloudGetV1DeployClusters(ctx).Execute()

ListDeployClusters returns the argocd ClusterList of the destinations the caller's applications reconcile into: one entry per distinct destination server, carrying the count of applications reconciling into it.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployClusters`: CloudArgoClusterList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployClustersRequest struct via the builder pattern


### Return type

[**CloudArgoClusterList**](CloudArgoClusterList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployGitops

> CloudGitOpsPlane CloudGetV1DeployGitops(ctx).Execute()

GetDeployGitOps lists every Hanzo CD Application in the cluster: the git source each one polls, the commit it last APPLIED, how its last sync operation ended, and its recent deploy history — newest deploy first, ordered by namespace then name.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployGitops(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployGitops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployGitops`: CloudGitOpsPlane
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployGitops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployGitopsRequest struct via the builder pattern


### Return type

[**CloudGitOpsPlane**](CloudGitOpsPlane.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployHealth

> CloudGetV1DeployHealth(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudGetV1DeployHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployHealthRequest struct via the builder pattern


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


## CloudGetV1DeployLogin

> CloudGetV1DeployLogin(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudGetV1DeployLogin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployLoginRequest struct via the builder pattern


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


## CloudGetV1DeployProjects

> CloudArgoProjectList CloudGetV1DeployProjects(ctx).Execute()

ListDeployProjects returns the argocd AppProjectList this console groups and filters applications by.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployProjects`: CloudArgoProjectList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployProjectsRequest struct via the builder pattern


### Return type

[**CloudArgoProjectList**](CloudArgoProjectList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeploySessionUserinfo

> CloudSessionUser CloudGetV1DeploySessionUserinfo(ctx).Execute()

GetDeploySession answers \"is this browser signed in, and if not where does it sign in?\" — the dashboard SPA's bootstrap question, and the only route on this plane that answers for an anonymous caller.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeploySessionUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeploySessionUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeploySessionUserinfo`: CloudSessionUser
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeploySessionUserinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeploySessionUserinfoRequest struct via the builder pattern


### Return type

[**CloudSessionUser**](CloudSessionUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeploySettings

> CloudConsoleSettings CloudGetV1DeploySettings(ctx).Execute()

GetDeploySettings returns the argocd AuthSettings object the dashboard SPA awaits before its first render.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeploySettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeploySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeploySettings`: CloudConsoleSettings
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeploySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeploySettingsRequest struct via the builder pattern


### Return type

[**CloudConsoleSettings**](CloudConsoleSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1DeployStreamApplications

> CloudGetV1DeployStreamApplications(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudGetV1DeployStreamApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployStreamApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployStreamApplicationsRequest struct via the builder pattern


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


## CloudGetV1DeployStreamApplicationsByNameResourceTree

> CloudGetV1DeployStreamApplicationsByNameResourceTree(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeployAPI.CloudGetV1DeployStreamApplicationsByNameResourceTree(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployStreamApplicationsByNameResourceTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployStreamApplicationsByNameResourceTreeRequest struct via the builder pattern


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


## CloudGetV1DeployVersion

> CloudVersionMessage CloudGetV1DeployVersion(ctx).Execute()

GetDeployVersion returns the argocd VersionMessage the dashboard SPA reads at bootstrap.



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
	resp, r, err := apiClient.DeployAPI.CloudGetV1DeployVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudGetV1DeployVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1DeployVersion`: CloudVersionMessage
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.CloudGetV1DeployVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1DeployVersionRequest struct via the builder pattern


### Return type

[**CloudVersionMessage**](CloudVersionMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1DeployApplicationsByNameRollback

> CloudPostV1DeployApplicationsByNameRollback(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeployAPI.CloudPostV1DeployApplicationsByNameRollback(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudPostV1DeployApplicationsByNameRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DeployApplicationsByNameRollbackRequest struct via the builder pattern


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


## CloudPostV1DeployApplicationsByNameSync

> CloudPostV1DeployApplicationsByNameSync(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeployAPI.CloudPostV1DeployApplicationsByNameSync(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudPostV1DeployApplicationsByNameSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DeployApplicationsByNameSyncRequest struct via the builder pattern


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


## CloudPostV1DeployLogout

> CloudPostV1DeployLogout(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudPostV1DeployLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudPostV1DeployLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DeployLogoutRequest struct via the builder pattern


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


## CloudPostV1DeployReconcile

> CloudPostV1DeployReconcile(ctx).Execute()



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
	r, err := apiClient.DeployAPI.CloudPostV1DeployReconcile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.CloudPostV1DeployReconcile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1DeployReconcileRequest struct via the builder pattern


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

