# \GitAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGitKeysById**](GitAPI.md#DeleteGitKeysById) | **Delete** /v1/git/keys/{id} | Removes a registered SSH key, scoped to the caller&#39;s org: an org can only delete its own, and a key id it does not own is not found.
[**DeleteGitReposByName**](GitAPI.md#DeleteGitReposByName) | **Delete** /v1/git/repos/{name} | Removes a repo&#39;s metadata and purges its storage.
[**DeleteGitReposByNameMirrorsById**](GitAPI.md#DeleteGitReposByNameMirrorsById) | **Delete** /v1/git/repos/{name}/mirrors/{id} | Removes one outbound mirror target; later pushes stop being forwarded to it.
[**DeleteGitReposByNameSubscriptionsById**](GitAPI.md#DeleteGitReposByNameSubscriptionsById) | **Delete** /v1/git/repos/{name}/subscriptions/{id} | Removes one Slack subscription from a repo; the notifier stops posting that repo&#39;s events to that channel.
[**GetGit**](GitAPI.md#GetGit) | **Get** /v1/git | Browse your org&#39;s repositories
[**GetGitByOrgByProjectByRepoInfoRefs**](GitAPI.md#GetGitByOrgByProjectByRepoInfoRefs) | **Get** /v1/git/{org}/{project}/{repo}/info/refs | Advertise a repository&#39;s refs to a git client
[**GetGitByOrgByRepo**](GitAPI.md#GetGitByOrgByRepo) | **Get** /v1/git/{org}/{repo} | Open a repository&#39;s home page
[**GetGitByOrgByRepoCommits**](GitAPI.md#GetGitByOrgByRepoCommits) | **Get** /v1/git/{org}/{repo}/commits | Read a repository&#39;s commit log
[**GetGitByOrgByRepoInfoRefs**](GitAPI.md#GetGitByOrgByRepoInfoRefs) | **Get** /v1/git/{org}/{repo}/info/refs | Advertise a repository&#39;s refs to a git client
[**GetGitExplore**](GitAPI.md#GetGitExplore) | **Get** /v1/git/explore | Discover public repositories across every org
[**GetGitKeys**](GitAPI.md#GetGitKeys) | **Get** /v1/git/keys | Returns the SSH public keys registered to the caller&#39;s org — the keys that authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60;.
[**GetGitRepos**](GitAPI.md#GetGitRepos) | **Get** /v1/git/repos | Returns the repos in the caller&#39;s scope, most recently updated first.
[**GetGitReposByName**](GitAPI.md#GetGitReposByName) | **Get** /v1/git/repos/{name} | Returns one repo with its live ref state: every branch name and the resolved HEAD commit.
[**GetGitReposByNameBlob**](GitAPI.md#GetGitReposByNameBlob) | **Get** /v1/git/repos/{name}/blob | Returns one file&#39;s bytes at one revision.
[**GetGitReposByNameCommits**](GitAPI.md#GetGitReposByNameCommits) | **Get** /v1/git/repos/{name}/commits | Walks a ref&#39;s history newest first, or one path&#39;s history when a path is given.
[**GetGitReposByNameFiles**](GitAPI.md#GetGitReposByNameFiles) | **Get** /v1/git/repos/{name}/files | Returns every file a glob selects at one revision, WITH its bytes and the revision they came from.
[**GetGitReposByNameMirrors**](GitAPI.md#GetGitReposByNameMirrors) | **Get** /v1/git/repos/{name}/mirrors | Returns a repo&#39;s outbound mirror targets — the downstream remotes the mirror reactor pushes to whenever a push lands here.
[**GetGitReposByNamePulls**](GitAPI.md#GetGitReposByNamePulls) | **Get** /v1/git/repos/{name}/pulls | Returns a repo&#39;s pull requests, newest number first — what is waiting to be reviewed, and what has already landed.
[**GetGitReposByNamePullsByNumber**](GitAPI.md#GetGitReposByNamePullsByNumber) | **Get** /v1/git/repos/{name}/pulls/{number} | Returns one pull request by its per-repo number.
[**GetGitReposByNameReadme**](GitAPI.md#GetGitReposByNameReadme) | **Get** /v1/git/repos/{name}/readme | Returns the README at the tree root as plain text — unrendered, so the caller decides how to present it.
[**GetGitReposByNameRefs**](GitAPI.md#GetGitReposByNameRefs) | **Get** /v1/git/repos/{name}/refs | Lists a repo&#39;s branches, tags and default branch — what a branch picker needs in one call.
[**GetGitReposByNameSubscriptions**](GitAPI.md#GetGitReposByNameSubscriptions) | **Get** /v1/git/repos/{name}/subscriptions | Returns a repo&#39;s Slack subscriptions — which channels the lifecycle notifier posts this repo&#39;s push and deploy events to.
[**GetGitReposByNameTree**](GitAPI.md#GetGitReposByNameTree) | **Get** /v1/git/repos/{name}/tree | Lists the immediate children of one directory at one revision, directories before files.
[**GetGitUsage**](GitAPI.md#GetGitUsage) | **Get** /v1/git/usage | Returns per-repo and total storage bytes for the caller&#39;s org — the queryable, per-tenant number commerce and o11y meter on.
[**PatchGitReposByName**](GitAPI.md#PatchGitReposByName) | **Patch** /v1/git/repos/{name} | Flips a repo&#39;s public bit, the one mutable repo setting today.
[**PostGitByOrgByProjectByRepoGitReceivePack**](GitAPI.md#PostGitByOrgByProjectByRepoGitReceivePack) | **Post** /v1/git/{org}/{project}/{repo}/git-receive-pack | Accept a push, and turn it into a build
[**PostGitByOrgByProjectByRepoGitUploadPack**](GitAPI.md#PostGitByOrgByProjectByRepoGitUploadPack) | **Post** /v1/git/{org}/{project}/{repo}/git-upload-pack | Serve a clone or fetch
[**PostGitByOrgByRepoGitReceivePack**](GitAPI.md#PostGitByOrgByRepoGitReceivePack) | **Post** /v1/git/{org}/{repo}/git-receive-pack | Accept a push, and turn it into a build
[**PostGitByOrgByRepoGitUploadPack**](GitAPI.md#PostGitByOrgByRepoGitUploadPack) | **Post** /v1/git/{org}/{repo}/git-upload-pack | Serve a clone or fetch
[**PostGitKeys**](GitAPI.md#PostGitKeys) | **Post** /v1/git/keys | Registers an SSH public key so it can authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60; for the caller&#39;s org.
[**PostGitRepos**](GitAPI.md#PostGitRepos) | **Post** /v1/git/repos | Provisions an empty bare repository in the caller&#39;s scope and returns it with its clone URLs.
[**PostGitReposByNameGc**](GitAPI.md#PostGitReposByNameGc) | **Post** /v1/git/repos/{name}/gc | Repacks a repo into one bitmapped pack and rewrites its commit-graph, so the next clone reuses the bitmap instead of walking the whole object graph.
[**PostGitReposByNameMirror**](GitAPI.md#PostGitReposByNameMirror) | **Post** /v1/git/repos/{name}/mirror | Imports an external git repository into the caller&#39;s repo, provisioning it on first use.
[**PostGitReposByNameMirrors**](GitAPI.md#PostGitReposByNameMirrors) | **Post** /v1/git/repos/{name}/mirrors | Registers a downstream remote the repo&#39;s advanced refs are pushed to whenever a push lands here.
[**PostGitReposByNamePulls**](GitAPI.md#PostGitReposByNamePulls) | **Post** /v1/git/repos/{name}/pulls | Proposes a branch for merging and returns it with its number.
[**PostGitReposByNamePullsByNumberMerge**](GitAPI.md#PostGitReposByNamePullsByNumberMerge) | **Post** /v1/git/repos/{name}/pulls/{number}/merge | Merges an open pull request by FAST-FORWARDING base to head, and answers the proposal in its merged state with the revision base now points at.
[**PostGitReposByNamePush**](GitAPI.md#PostGitReposByNamePush) | **Post** /v1/git/repos/{name}/push | Lands a set of files as one commit without a git client — the hanzo.app builder&#39;s push.
[**PostGitReposByNameSubscriptions**](GitAPI.md#PostGitReposByNameSubscriptions) | **Post** /v1/git/repos/{name}/subscriptions | Binds a Slack channel to a repo, so the lifecycle notifier posts that repo&#39;s push and deploy events there.
[**PostGitWebhook**](GitAPI.md#PostGitWebhook) | **Post** /v1/git/webhook | Retired — forge pushes build via platform.hanzo.ai
[**PostGitZapCreaterepo**](GitAPI.md#PostGitZapCreaterepo) | **Post** /v1/git/zap/createRepo | Create a repository over the ZAP transport
[**PostGitZapDeleterepo**](GitAPI.md#PostGitZapDeleterepo) | **Post** /v1/git/zap/deleteRepo | Delete a repository over the ZAP transport
[**PostGitZapGetrepo**](GitAPI.md#PostGitZapGetrepo) | **Post** /v1/git/zap/getRepo | Read one repository over the ZAP transport
[**PostGitZapListrepos**](GitAPI.md#PostGitZapListrepos) | **Post** /v1/git/zap/listRepos | List your repositories over the ZAP transport
[**PostGitZapUsage**](GitAPI.md#PostGitZapUsage) | **Post** /v1/git/zap/usage | Report your org&#39;s git storage footprint over the ZAP transport



## DeleteGitKeysById

> DeleteGitKeysById(ctx, id).Execute()

Removes a registered SSH key, scoped to the caller's org: an org can only delete its own, and a key id it does not own is not found.



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
	id := "gitkey_4a1b" // string | ID is the key's identifier (\"gitkey_…\"), from the :id path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.DeleteGitKeysById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.DeleteGitKeysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the key&#39;s identifier (\&quot;gitkey_…\&quot;), from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitKeysByIdRequest struct via the builder pattern


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


## DeleteGitReposByName

> DeleteGitReposByName(ctx, name).Execute()

Removes a repo's metadata and purges its storage.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.DeleteGitReposByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.DeleteGitReposByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitReposByNameRequest struct via the builder pattern


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


## DeleteGitReposByNameMirrorsById

> DeleteGitReposByNameMirrorsById(ctx, name, id).Execute()

Removes one outbound mirror target; later pushes stop being forwarded to it.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	id := "mir_2d90" // string | ID is the row to remove, from the :id path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.DeleteGitReposByNameMirrorsById(context.Background(), name, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.DeleteGitReposByNameMirrorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**id** | **string** | ID is the row to remove, from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitReposByNameMirrorsByIdRequest struct via the builder pattern


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


## DeleteGitReposByNameSubscriptionsById

> DeleteGitReposByNameSubscriptionsById(ctx, name, id).Execute()

Removes one Slack subscription from a repo; the notifier stops posting that repo's events to that channel.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	id := "sub_7c2e" // string | ID is the row to remove, from the :id path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.DeleteGitReposByNameSubscriptionsById(context.Background(), name, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.DeleteGitReposByNameSubscriptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**id** | **string** | ID is the row to remove, from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitReposByNameSubscriptionsByIdRequest struct via the builder pattern


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


## GetGit

> GetGit(ctx).Execute()

Browse your org's repositories



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
	r, err := apiClient.GitAPI.GetGit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitRequest struct via the builder pattern


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


## GetGitByOrgByProjectByRepoInfoRefs

> GetGitByOrgByProjectByRepoInfoRefs(ctx, org, project, repo).Execute()

Advertise a repository's refs to a git client



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
	org := "org_example" // string | 
	project := "project_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.GetGitByOrgByProjectByRepoInfoRefs(context.Background(), org, project, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitByOrgByProjectByRepoInfoRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**project** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitByOrgByProjectByRepoInfoRefsRequest struct via the builder pattern


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


## GetGitByOrgByRepo

> GetGitByOrgByRepo(ctx, org, repo).Execute()

Open a repository's home page



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
	org := "org_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.GetGitByOrgByRepo(context.Background(), org, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitByOrgByRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitByOrgByRepoRequest struct via the builder pattern


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


## GetGitByOrgByRepoCommits

> GetGitByOrgByRepoCommits(ctx, org, repo).Execute()

Read a repository's commit log



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
	org := "org_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.GetGitByOrgByRepoCommits(context.Background(), org, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitByOrgByRepoCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitByOrgByRepoCommitsRequest struct via the builder pattern


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


## GetGitByOrgByRepoInfoRefs

> GetGitByOrgByRepoInfoRefs(ctx, org, repo).Execute()

Advertise a repository's refs to a git client



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
	org := "org_example" // string | 
	repo := "repo_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.GetGitByOrgByRepoInfoRefs(context.Background(), org, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitByOrgByRepoInfoRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitByOrgByRepoInfoRefsRequest struct via the builder pattern


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


## GetGitExplore

> GetGitExplore(ctx).Execute()

Discover public repositories across every org



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
	r, err := apiClient.GitAPI.GetGitExplore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitExplore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitExploreRequest struct via the builder pattern


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


## GetGitKeys

> KeyList GetGitKeys(ctx).Execute()

Returns the SSH public keys registered to the caller's org — the keys that authenticate `git clone git@<host>:<org>/<repo>.git`.



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
	resp, r, err := apiClient.GitAPI.GetGitKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitKeys`: KeyList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitKeysRequest struct via the builder pattern


### Return type

[**KeyList**](KeyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitRepos

> RepoList GetGitRepos(ctx).Execute()

Returns the repos in the caller's scope, most recently updated first.



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
	resp, r, err := apiClient.GitAPI.GetGitRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRepos`: RepoList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposRequest struct via the builder pattern


### Return type

[**RepoList**](RepoList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByName

> RepoView GetGitReposByName(ctx, name).Execute()

Returns one repo with its live ref state: every branch name and the resolved HEAD commit.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByName`: RepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepoView**](RepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameBlob

> BlobJSON GetGitReposByNameBlob(ctx, name).Ref(ref).Path(path).Execute()

Returns one file's bytes at one revision.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is a branch, tag or commit; empty means the repo's HEAD. (optional)
	path := "go.mod" // string | Path is repo-relative; empty is the tree root. Traversal is stripped. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameBlob(context.Background(), name).Ref(ref).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameBlob`: BlobJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **path** | **string** | Path is repo-relative; empty is the tree root. Traversal is stripped. | 

### Return type

[**BlobJSON**](BlobJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameCommits

> CommitsJSON GetGitReposByNameCommits(ctx, name).Ref(ref).Path(path).Limit(limit).Execute()

Walks a ref's history newest first, or one path's history when a path is given.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is the branch, tag or commit to walk back from; empty means HEAD. (optional)
	path := "path_example" // string | Path narrows the history to commits touching it; empty walks the whole ref. (optional)
	limit := int32(2) // int32 | Limit caps the page. Anything not positive means 50; the cap is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameCommits(context.Background(), name).Ref(ref).Path(path).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameCommits`: CommitsJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is the branch, tag or commit to walk back from; empty means HEAD. | 
 **path** | **string** | Path narrows the history to commits touching it; empty walks the whole ref. | 
 **limit** | **int32** | Limit caps the page. Anything not positive means 50; the cap is 100. | 

### Return type

[**CommitsJSON**](CommitsJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameFiles

> FilesJSON GetGitReposByNameFiles(ctx, name).Ref(ref).Glob(glob).Execute()

Returns every file a glob selects at one revision, WITH its bytes and the revision they came from.



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
	name := "universe" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is a branch, tag or commit; empty means the repo's HEAD. (optional)
	glob := "charts/app/values/*/*.yaml" // string | Glob selects files, matched segment by segment so `*` never crosses a `/`. `**` matches zero or more whole segments. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameFiles(context.Background(), name).Ref(ref).Glob(glob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameFiles`: FilesJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **glob** | **string** | Glob selects files, matched segment by segment so &#x60;*&#x60; never crosses a &#x60;/&#x60;. &#x60;**&#x60; matches zero or more whole segments. | 

### Return type

[**FilesJSON**](FilesJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameMirrors

> MirrorList GetGitReposByNameMirrors(ctx, name).Execute()

Returns a repo's outbound mirror targets — the downstream remotes the mirror reactor pushes to whenever a push lands here.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameMirrors(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameMirrors`: MirrorList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MirrorList**](MirrorList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNamePulls

> PullList GetGitReposByNamePulls(ctx, name).State(state).Execute()

Returns a repo's pull requests, newest number first — what is waiting to be reviewed, and what has already landed.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	state := "open" // string | State narrows the list to \"open\" or \"merged\". Omit it for every proposal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNamePulls(context.Background(), name).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNamePulls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNamePulls`: PullList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNamePulls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNamePullsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **string** | State narrows the list to \&quot;open\&quot; or \&quot;merged\&quot;. Omit it for every proposal. | 

### Return type

[**PullList**](PullList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNamePullsByNumber

> PullView GetGitReposByNamePullsByNumber(ctx, name, number).Execute()

Returns one pull request by its per-repo number.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	number := int32(4) // int32 | Number is the proposal's per-repo number, from the :number path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNamePullsByNumber(context.Background(), name, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNamePullsByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNamePullsByNumber`: PullView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNamePullsByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**number** | **int32** | Number is the proposal&#39;s per-repo number, from the :number path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNamePullsByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PullView**](PullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameReadme

> ReadmeJSON GetGitReposByNameReadme(ctx, name).Ref(ref).Execute()

Returns the README at the tree root as plain text — unrendered, so the caller decides how to present it.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is a branch, tag or commit; empty means the repo's HEAD. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameReadme(context.Background(), name).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameReadme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameReadme`: ReadmeJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameReadme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameReadmeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 

### Return type

[**ReadmeJSON**](ReadmeJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameRefs

> RefsJSON GetGitReposByNameRefs(ctx, name).Execute()

Lists a repo's branches, tags and default branch — what a branch picker needs in one call.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameRefs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameRefs`: RefsJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefsJSON**](RefsJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameSubscriptions

> SubscriptionList GetGitReposByNameSubscriptions(ctx, name).Execute()

Returns a repo's Slack subscriptions — which channels the lifecycle notifier posts this repo's push and deploy events to.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameSubscriptions(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameSubscriptions`: SubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionList**](SubscriptionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameTree

> TreeJSON GetGitReposByNameTree(ctx, name).Ref(ref).Path(path).Execute()

Lists the immediate children of one directory at one revision, directories before files.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is a branch, tag or commit; empty means the repo's HEAD. (optional)
	path := "cmd" // string | Path is repo-relative; empty is the tree root. Traversal is stripped. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameTree(context.Background(), name).Ref(ref).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameTree`: TreeJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **path** | **string** | Path is repo-relative; empty is the tree root. Traversal is stripped. | 

### Return type

[**TreeJSON**](TreeJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitUsage

> UsageView GetGitUsage(ctx).Execute()

Returns per-repo and total storage bytes for the caller's org — the queryable, per-tenant number commerce and o11y meter on.



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
	resp, r, err := apiClient.GitAPI.GetGitUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitUsage`: UsageView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitUsageRequest struct via the builder pattern


### Return type

[**UsageView**](UsageView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGitReposByName

> RepoView PatchGitReposByName(ctx, name).PatchIn(patchIn).Execute()

Flips a repo's public bit, the one mutable repo setting today.



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
	name := "widgets" // string | Name is the repo to update, from the :name path segment.
	patchIn := *openapiclient.NewPatchIn() // PatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PatchGitReposByName(context.Background(), name).PatchIn(patchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PatchGitReposByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchGitReposByName`: RepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PatchGitReposByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to update, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGitReposByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchIn** | [**PatchIn**](PatchIn.md) |  | 

### Return type

[**RepoView**](RepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitByOrgByProjectByRepoGitReceivePack

> PostGitByOrgByProjectByRepoGitReceivePack(ctx, org, project, repo).Body(body).Execute()

Accept a push, and turn it into a build



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
	org := "org_example" // string | 
	project := "project_example" // string | 
	repo := "repo_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitByOrgByProjectByRepoGitReceivePack(context.Background(), org, project, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitByOrgByProjectByRepoGitReceivePack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**project** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitByOrgByProjectByRepoGitReceivePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitByOrgByProjectByRepoGitUploadPack

> PostGitByOrgByProjectByRepoGitUploadPack(ctx, org, project, repo).Body(body).Execute()

Serve a clone or fetch



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
	org := "org_example" // string | 
	project := "project_example" // string | 
	repo := "repo_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitByOrgByProjectByRepoGitUploadPack(context.Background(), org, project, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitByOrgByProjectByRepoGitUploadPack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**project** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitByOrgByProjectByRepoGitUploadPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitByOrgByRepoGitReceivePack

> PostGitByOrgByRepoGitReceivePack(ctx, org, repo).Body(body).Execute()

Accept a push, and turn it into a build



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
	org := "org_example" // string | 
	repo := "repo_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitByOrgByRepoGitReceivePack(context.Background(), org, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitByOrgByRepoGitReceivePack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitByOrgByRepoGitReceivePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitByOrgByRepoGitUploadPack

> PostGitByOrgByRepoGitUploadPack(ctx, org, repo).Body(body).Execute()

Serve a clone or fetch



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
	org := "org_example" // string | 
	repo := "repo_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitByOrgByRepoGitUploadPack(context.Background(), org, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitByOrgByRepoGitUploadPack``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**repo** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitByOrgByRepoGitUploadPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitKeys

> KeyView PostGitKeys(ctx).RegisterKeyReq(registerKeyReq).Execute()

Registers an SSH public key so it can authenticate `git clone git@<host>:<org>/<repo>.git` for the caller's org.



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
	registerKeyReq := *openapiclient.NewRegisterKeyReq() // RegisterKeyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitKeys(context.Background()).RegisterKeyReq(registerKeyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitKeys`: KeyView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerKeyReq** | [**RegisterKeyReq**](RegisterKeyReq.md) |  | 

### Return type

[**KeyView**](KeyView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitRepos

> RepoView PostGitRepos(ctx).CreateReq(createReq).Execute()

Provisions an empty bare repository in the caller's scope and returns it with its clone URLs.



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
	createReq := *openapiclient.NewCreateReq() // CreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitRepos(context.Background()).CreateReq(createReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitRepos`: RepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReq** | [**CreateReq**](CreateReq.md) |  | 

### Return type

[**RepoView**](RepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameGc

> GcOut PostGitReposByNameGc(ctx, name).Execute()

Repacks a repo into one bitmapped pack and rewrites its commit-graph, so the next clone reuses the bitmap instead of walking the whole object graph.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameGc(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameGc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameGc`: GcOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNameGc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNameGcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GcOut**](GcOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameMirror

> RepoView PostGitReposByNameMirror(ctx, name).MirrorReq(mirrorReq).Execute()

Imports an external git repository into the caller's repo, provisioning it on first use.



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
	name := "widgets" // string | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use.
	mirrorReq := *openapiclient.NewMirrorReq() // MirrorReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameMirror(context.Background(), name).MirrorReq(mirrorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameMirror``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameMirror`: RepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNameMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNameMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorReq** | [**MirrorReq**](MirrorReq.md) |  | 

### Return type

[**RepoView**](RepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameMirrors

> MirrorTargetView PostGitReposByNameMirrors(ctx, name).MirrorTargetReq(mirrorTargetReq).Execute()

Registers a downstream remote the repo's advanced refs are pushed to whenever a push lands here.



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
	name := "widgets" // string | Name is the repo whose advanced refs are pushed downstream, from the :name path segment.
	mirrorTargetReq := *openapiclient.NewMirrorTargetReq() // MirrorTargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameMirrors(context.Background(), name).MirrorTargetReq(mirrorTargetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameMirrors`: MirrorTargetView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNameMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNameMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mirrorTargetReq** | [**MirrorTargetReq**](MirrorTargetReq.md) |  | 

### Return type

[**MirrorTargetView**](MirrorTargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePulls

> PullView PostGitReposByNamePulls(ctx, name).OpenReq(openReq).Execute()

Proposes a branch for merging and returns it with its number.



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
	name := "widgets" // string | Name is the repo the proposal belongs to, from the :name path segment.
	openReq := *openapiclient.NewOpenReq() // OpenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePulls(context.Background(), name).OpenReq(openReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePulls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePulls`: PullView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNamePulls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo the proposal belongs to, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNamePullsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openReq** | [**OpenReq**](OpenReq.md) |  | 

### Return type

[**PullView**](PullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePullsByNumberMerge

> PullView PostGitReposByNamePullsByNumberMerge(ctx, name, number).Execute()

Merges an open pull request by FAST-FORWARDING base to head, and answers the proposal in its merged state with the revision base now points at.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	number := int32(4) // int32 | Number is the proposal's per-repo number, from the :number path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePullsByNumberMerge(context.Background(), name, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePullsByNumberMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePullsByNumberMerge`: PullView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNamePullsByNumberMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**number** | **int32** | Number is the proposal&#39;s per-repo number, from the :number path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNamePullsByNumberMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PullView**](PullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePush

> PushResp PostGitReposByNamePush(ctx, name).PushReq(pushReq).Execute()

Lands a set of files as one commit without a git client — the hanzo.app builder's push.



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
	name := "widgets" // string | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist.
	pushReq := *openapiclient.NewPushReq() // PushReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePush(context.Background(), name).PushReq(pushReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePush`: PushResp
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNamePush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNamePushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pushReq** | [**PushReq**](PushReq.md) |  | 

### Return type

[**PushResp**](PushResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameSubscriptions

> SubscriptionView PostGitReposByNameSubscriptions(ctx, name).SubscribeReq(subscribeReq).Execute()

Binds a Slack channel to a repo, so the lifecycle notifier posts that repo's push and deploy events there.



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
	name := "widgets" // string | Name is the repo to subscribe, from the :name path segment.
	subscribeReq := *openapiclient.NewSubscribeReq() // SubscribeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameSubscriptions(context.Background(), name).SubscribeReq(subscribeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameSubscriptions`: SubscriptionView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNameSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to subscribe, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNameSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscribeReq** | [**SubscribeReq**](SubscribeReq.md) |  | 

### Return type

[**SubscriptionView**](SubscriptionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitWebhook

> PostGitWebhook(ctx).Execute()

Retired — forge pushes build via platform.hanzo.ai



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
	r, err := apiClient.GitAPI.PostGitWebhook(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitWebhookRequest struct via the builder pattern


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


## PostGitZapCreaterepo

> PostGitZapCreaterepo(ctx).ZapProcReq(zapProcReq).Execute()

Create a repository over the ZAP transport



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
	zapProcReq := *openapiclient.NewZapProcReq() // ZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitZapCreaterepo(context.Background()).ZapProcReq(zapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitZapCreaterepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitZapCreaterepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zapProcReq** | [**ZapProcReq**](ZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitZapDeleterepo

> PostGitZapDeleterepo(ctx).ZapProcReq(zapProcReq).Execute()

Delete a repository over the ZAP transport



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
	zapProcReq := *openapiclient.NewZapProcReq() // ZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitZapDeleterepo(context.Background()).ZapProcReq(zapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitZapDeleterepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitZapDeleterepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zapProcReq** | [**ZapProcReq**](ZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitZapGetrepo

> PostGitZapGetrepo(ctx).ZapProcReq(zapProcReq).Execute()

Read one repository over the ZAP transport



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
	zapProcReq := *openapiclient.NewZapProcReq() // ZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.PostGitZapGetrepo(context.Background()).ZapProcReq(zapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitZapGetrepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitZapGetrepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zapProcReq** | [**ZapProcReq**](ZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitZapListrepos

> PostGitZapListrepos(ctx).Execute()

List your repositories over the ZAP transport



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
	r, err := apiClient.GitAPI.PostGitZapListrepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitZapListrepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitZapListreposRequest struct via the builder pattern


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


## PostGitZapUsage

> PostGitZapUsage(ctx).Execute()

Report your org's git storage footprint over the ZAP transport



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
	r, err := apiClient.GitAPI.PostGitZapUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitZapUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitZapUsageRequest struct via the builder pattern


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

