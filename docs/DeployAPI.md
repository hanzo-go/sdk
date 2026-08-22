# \DeployAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeployApplications**](DeployAPI.md#GetDeployApplications) | **Get** /v1/deploy/applications | Returns the fleet as an argocd ApplicationList: one projected Application per operator App CR, carrying the image tag the CR DECLARES, the tag actually RUNNING in the cluster&#39;s Deployment, the reconciled health, and the sync verdict those two produce (declared &#x3D;&#x3D; running ⇒ Synced, both known and different ⇒ OutOfSync, either unknown ⇒ Unknown).
[**GetDeployApplicationsByName**](DeployAPI.md#GetDeployApplicationsByName) | **Get** /v1/deploy/applications/{name} | Returns ONE projected argocd Application by name, with status.resources filled in from its reconciled resource tree — which is what makes it the detail view rather than a row of the list.
[**GetDeployApplicationsByNameResourceTree**](DeployAPI.md#GetDeployApplicationsByNameResourceTree) | **Get** /v1/deploy/applications/{name}/resource-tree | Returns one application&#39;s argocd ApplicationTree: the objects the operator reconciled from its App CR, reached by ownerRef — the Deployment and, under it, the ReplicaSet and Pods, plus the Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget and ConfigMaps it owns — each node carrying its parent edges and its health.
[**GetDeployApplicationsByNameRevisionsByRevisionMetadata**](DeployAPI.md#GetDeployApplicationsByNameRevisionsByRevisionMetadata) | **Get** /v1/deploy/applications/{name}/revisions/{revision}/metadata | Returns the argocd RevisionMetadata for one revision of one application — what the detail view shows beside a revision.
[**GetDeployApplicationsByNameSyncwindows**](DeployAPI.md#GetDeployApplicationsByNameSyncwindows) | **Get** /v1/deploy/applications/{name}/syncwindows | Returns one application&#39;s argocd ApplicationSyncWindowState — the answer to \&quot;is anything blocking a sync of this application right now?\&quot;.
[**GetDeployCallback**](DeployAPI.md#GetDeployCallback) | **Get** /v1/deploy/callback | Finish the sign-in round trip and mint the console session
[**GetDeployClusters**](DeployAPI.md#GetDeployClusters) | **Get** /v1/deploy/clusters | Returns the argocd ClusterList of the destinations the caller&#39;s applications reconcile into: one entry per distinct destination server, carrying the count of applications reconciling into it.
[**GetDeployGitops**](DeployAPI.md#GetDeployGitops) | **Get** /v1/deploy/gitops | Lists every Hanzo CD Application in the cluster: the git source each one polls, the commit it last APPLIED, how its last sync operation ended, and its recent deploy history — newest deploy first, ordered by namespace then name.
[**GetDeployHealth**](DeployAPI.md#GetDeployHealth) | **Get** /v1/deploy/health | Whether this control plane can actually reach the cluster it deploys to
[**GetDeployLogin**](DeployAPI.md#GetDeployLogin) | **Get** /v1/deploy/login | Start the sign-in round trip for this console
[**GetDeployProjects**](DeployAPI.md#GetDeployProjects) | **Get** /v1/deploy/projects | Returns the argocd AppProjectList this console groups and filters applications by.
[**GetDeploySessionUserinfo**](DeployAPI.md#GetDeploySessionUserinfo) | **Get** /v1/deploy/session/userinfo | Answers \&quot;is this browser signed in, and if not where does it sign in?\&quot; — the dashboard SPA&#39;s bootstrap question, and the only route on this plane that answers for an anonymous caller.
[**GetDeploySettings**](DeployAPI.md#GetDeploySettings) | **Get** /v1/deploy/settings | Returns the argocd AuthSettings object the dashboard SPA awaits before its first render.
[**GetDeployStreamApplications**](DeployAPI.md#GetDeployStreamApplications) | **Get** /v1/deploy/stream/applications | Live application fleet updates as Server-Sent Events
[**GetDeployStreamApplicationsByNameResourceTree**](DeployAPI.md#GetDeployStreamApplicationsByNameResourceTree) | **Get** /v1/deploy/stream/applications/{name}/resource-tree | Live resource tree for one application, as Server-Sent Events
[**GetDeployVersion**](DeployAPI.md#GetDeployVersion) | **Get** /v1/deploy/version | Returns the argocd VersionMessage the dashboard SPA reads at bootstrap.
[**PostDeployApplicationsByNameRollback**](DeployAPI.md#PostDeployApplicationsByNameRollback) | **Post** /v1/deploy/applications/{name}/rollback | The console&#39;s rollback control — today it requests a reconcile, nothing more
[**PostDeployApplicationsByNameSync**](DeployAPI.md#PostDeployApplicationsByNameSync) | **Post** /v1/deploy/applications/{name}/sync | Ask the operator to reconcile one application now
[**PostDeployLogout**](DeployAPI.md#PostDeployLogout) | **Post** /v1/deploy/logout | End the console session on this host
[**PostDeployReconcile**](DeployAPI.md#PostDeployReconcile) | **Post** /v1/deploy/reconcile | Render the configured git source and apply it to the cluster, once



## GetDeployApplications

> ArgoAppList GetDeployApplications(ctx).Execute()

Returns the fleet as an argocd ApplicationList: one projected Application per operator App CR, carrying the image tag the CR DECLARES, the tag actually RUNNING in the cluster's Deployment, the reconciled health, and the sync verdict those two produce (declared == running ⇒ Synced, both known and different ⇒ OutOfSync, either unknown ⇒ Unknown).



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
	resp, r, err := apiClient.DeployAPI.GetDeployApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployApplications`: ArgoAppList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployApplicationsRequest struct via the builder pattern


### Return type

[**ArgoAppList**](ArgoAppList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployApplicationsByName

> ArgoApp GetDeployApplicationsByName(ctx, name).Execute()

Returns ONE projected argocd Application by name, with status.resources filled in from its reconciled resource tree — which is what makes it the detail view rather than a row of the list.



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
	resp, r, err := apiClient.DeployAPI.GetDeployApplicationsByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployApplicationsByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployApplicationsByName`: ArgoApp
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployApplicationsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployApplicationsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArgoApp**](ArgoApp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployApplicationsByNameResourceTree

> ArgoTree GetDeployApplicationsByNameResourceTree(ctx, name).Execute()

Returns one application's argocd ApplicationTree: the objects the operator reconciled from its App CR, reached by ownerRef — the Deployment and, under it, the ReplicaSet and Pods, plus the Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget and ConfigMaps it owns — each node carrying its parent edges and its health.



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
	resp, r, err := apiClient.DeployAPI.GetDeployApplicationsByNameResourceTree(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployApplicationsByNameResourceTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployApplicationsByNameResourceTree`: ArgoTree
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployApplicationsByNameResourceTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployApplicationsByNameResourceTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArgoTree**](ArgoTree.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployApplicationsByNameRevisionsByRevisionMetadata

> ArgoRevisionMetadata GetDeployApplicationsByNameRevisionsByRevisionMetadata(ctx, name, revision).Execute()

Returns the argocd RevisionMetadata for one revision of one application — what the detail view shows beside a revision.



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
	resp, r, err := apiClient.DeployAPI.GetDeployApplicationsByNameRevisionsByRevisionMetadata(context.Background(), name, revision).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployApplicationsByNameRevisionsByRevisionMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployApplicationsByNameRevisionsByRevisionMetadata`: ArgoRevisionMetadata
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployApplicationsByNameRevisionsByRevisionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label. | 
**revision** | **string** | Revision is the revision to describe, from the path. The empty revision and \&quot;HEAD\&quot; both mean \&quot;whatever this application currently declares\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployApplicationsByNameRevisionsByRevisionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ArgoRevisionMetadata**](ArgoRevisionMetadata.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployApplicationsByNameSyncwindows

> ArgoSyncWindows GetDeployApplicationsByNameSyncwindows(ctx, name).Execute()

Returns one application's argocd ApplicationSyncWindowState — the answer to \"is anything blocking a sync of this application right now?\".



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
	resp, r, err := apiClient.DeployAPI.GetDeployApplicationsByNameSyncwindows(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployApplicationsByNameSyncwindows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployApplicationsByNameSyncwindows`: ArgoSyncWindows
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployApplicationsByNameSyncwindows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the application to read, from the path. It must be a DNS-1123 label (lowercase alphanumerics and hyphens, starting and ending alphanumeric) — every operator App CR&#39;s metadata.name satisfies that, and anything else is a 400 rather than a lookup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployApplicationsByNameSyncwindowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArgoSyncWindows**](ArgoSyncWindows.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployCallback

> GetDeployCallback(ctx).Execute()

Finish the sign-in round trip and mint the console session



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
	r, err := apiClient.DeployAPI.GetDeployCallback(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployCallbackRequest struct via the builder pattern


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


## GetDeployClusters

> ArgoClusterList GetDeployClusters(ctx).Execute()

Returns the argocd ClusterList of the destinations the caller's applications reconcile into: one entry per distinct destination server, carrying the count of applications reconciling into it.



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
	resp, r, err := apiClient.DeployAPI.GetDeployClusters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployClusters`: ArgoClusterList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployClustersRequest struct via the builder pattern


### Return type

[**ArgoClusterList**](ArgoClusterList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployGitops

> GitOpsPlane GetDeployGitops(ctx).Execute()

Lists every Hanzo CD Application in the cluster: the git source each one polls, the commit it last APPLIED, how its last sync operation ended, and its recent deploy history — newest deploy first, ordered by namespace then name.



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
	resp, r, err := apiClient.DeployAPI.GetDeployGitops(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployGitops``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployGitops`: GitOpsPlane
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployGitops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployGitopsRequest struct via the builder pattern


### Return type

[**GitOpsPlane**](GitOpsPlane.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployHealth

> GetDeployHealth(ctx).Execute()

Whether this control plane can actually reach the cluster it deploys to



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
	r, err := apiClient.DeployAPI.GetDeployHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployHealthRequest struct via the builder pattern


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


## GetDeployLogin

> GetDeployLogin(ctx).Execute()

Start the sign-in round trip for this console



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
	r, err := apiClient.DeployAPI.GetDeployLogin(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployLoginRequest struct via the builder pattern


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


## GetDeployProjects

> ArgoProjectList GetDeployProjects(ctx).Execute()

Returns the argocd AppProjectList this console groups and filters applications by.



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
	resp, r, err := apiClient.DeployAPI.GetDeployProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployProjects`: ArgoProjectList
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployProjectsRequest struct via the builder pattern


### Return type

[**ArgoProjectList**](ArgoProjectList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploySessionUserinfo

> SessionUser GetDeploySessionUserinfo(ctx).Execute()

Answers \"is this browser signed in, and if not where does it sign in?\" — the dashboard SPA's bootstrap question, and the only route on this plane that answers for an anonymous caller.



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
	resp, r, err := apiClient.DeployAPI.GetDeploySessionUserinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeploySessionUserinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploySessionUserinfo`: SessionUser
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeploySessionUserinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploySessionUserinfoRequest struct via the builder pattern


### Return type

[**SessionUser**](SessionUser.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeploySettings

> ConsoleSettings GetDeploySettings(ctx).Execute()

Returns the argocd AuthSettings object the dashboard SPA awaits before its first render.



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
	resp, r, err := apiClient.DeployAPI.GetDeploySettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeploySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeploySettings`: ConsoleSettings
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeploySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeploySettingsRequest struct via the builder pattern


### Return type

[**ConsoleSettings**](ConsoleSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeployStreamApplications

> GetDeployStreamApplications(ctx).Execute()

Live application fleet updates as Server-Sent Events



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
	r, err := apiClient.DeployAPI.GetDeployStreamApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployStreamApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployStreamApplicationsRequest struct via the builder pattern


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


## GetDeployStreamApplicationsByNameResourceTree

> GetDeployStreamApplicationsByNameResourceTree(ctx, name).Execute()

Live resource tree for one application, as Server-Sent Events



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
	r, err := apiClient.DeployAPI.GetDeployStreamApplicationsByNameResourceTree(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployStreamApplicationsByNameResourceTree``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetDeployStreamApplicationsByNameResourceTreeRequest struct via the builder pattern


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


## GetDeployVersion

> VersionMessage GetDeployVersion(ctx).Execute()

Returns the argocd VersionMessage the dashboard SPA reads at bootstrap.



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
	resp, r, err := apiClient.DeployAPI.GetDeployVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.GetDeployVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeployVersion`: VersionMessage
	fmt.Fprintf(os.Stdout, "Response from `DeployAPI.GetDeployVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeployVersionRequest struct via the builder pattern


### Return type

[**VersionMessage**](VersionMessage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeployApplicationsByNameRollback

> PostDeployApplicationsByNameRollback(ctx, name).Execute()

The console's rollback control — today it requests a reconcile, nothing more



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
	r, err := apiClient.DeployAPI.PostDeployApplicationsByNameRollback(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.PostDeployApplicationsByNameRollback``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostDeployApplicationsByNameRollbackRequest struct via the builder pattern


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


## PostDeployApplicationsByNameSync

> PostDeployApplicationsByNameSync(ctx, name).Execute()

Ask the operator to reconcile one application now



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
	r, err := apiClient.DeployAPI.PostDeployApplicationsByNameSync(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.PostDeployApplicationsByNameSync``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostDeployApplicationsByNameSyncRequest struct via the builder pattern


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


## PostDeployLogout

> PostDeployLogout(ctx).Execute()

End the console session on this host



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
	r, err := apiClient.DeployAPI.PostDeployLogout(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.PostDeployLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeployLogoutRequest struct via the builder pattern


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


## PostDeployReconcile

> PostDeployReconcile(ctx).Execute()

Render the configured git source and apply it to the cluster, once



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
	r, err := apiClient.DeployAPI.PostDeployReconcile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeployAPI.PostDeployReconcile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeployReconcileRequest struct via the builder pattern


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

