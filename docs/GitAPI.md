# \GitAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGitKeysById**](GitAPI.md#DeleteGitKeysById) | **Delete** /v1/git/keys/{id} | Removes a registered SSH key, scoped to the caller&#39;s org: an org can only delete its own, and a key id it does not own is not found.
[**DeleteGitReposByName**](GitAPI.md#DeleteGitReposByName) | **Delete** /v1/git/repos/{name} | Removes a repo&#39;s metadata and purges its storage.
[**DeleteGitReposByNameSubscriptionsById**](GitAPI.md#DeleteGitReposByNameSubscriptionsById) | **Delete** /v1/git/repos/{name}/subscriptions/{id} | Removes one Slack subscription from a repo; the notifier stops posting that repo&#39;s events to that channel.
[**DeleteGitReposByNameTargetsById**](GitAPI.md#DeleteGitReposByNameTargetsById) | **Delete** /v1/git/repos/{name}/targets/{id} | Removes one outbound mirror target; later pushes stop being forwarded to it.
[**GetGit**](GitAPI.md#GetGit) | **Get** /v1/git | Browse your org&#39;s repositories
[**GetGitByOrgByProjectByRepoInfoRefs**](GitAPI.md#GetGitByOrgByProjectByRepoInfoRefs) | **Get** /v1/git/{org}/{project}/{repo}/info/refs | Advertise a repository&#39;s refs to a git client
[**GetGitByOrgByRepo**](GitAPI.md#GetGitByOrgByRepo) | **Get** /v1/git/{org}/{repo} | Open a repository&#39;s home page
[**GetGitByOrgByRepoCommits**](GitAPI.md#GetGitByOrgByRepoCommits) | **Get** /v1/git/{org}/{repo}/commits | Read a repository&#39;s commit log
[**GetGitByOrgByRepoInfoRefs**](GitAPI.md#GetGitByOrgByRepoInfoRefs) | **Get** /v1/git/{org}/{repo}/info/refs | Advertise a repository&#39;s refs to a git client
[**GetGitExplore**](GitAPI.md#GetGitExplore) | **Get** /v1/git/explore | Discover public repositories across every org
[**GetGitKeys**](GitAPI.md#GetGitKeys) | **Get** /v1/git/keys | Returns the SSH public keys registered to the caller&#39;s org — the keys that authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60;.
[**GetGitPools**](GitAPI.md#GetGitPools) | **Get** /v1/git/pools | Returns the capacity this org has declared and how many daemons have entered each pool.
[**GetGitRepos**](GitAPI.md#GetGitRepos) | **Get** /v1/git/repos | Returns the repos in the caller&#39;s scope, most recently updated first.
[**GetGitReposByName**](GitAPI.md#GetGitReposByName) | **Get** /v1/git/repos/{name} | Returns one repo with its live ref state: every branch name and the resolved HEAD commit.
[**GetGitReposByNameBlob**](GitAPI.md#GetGitReposByNameBlob) | **Get** /v1/git/repos/{name}/blob | Returns one file&#39;s bytes at one revision.
[**GetGitReposByNameCommits**](GitAPI.md#GetGitReposByNameCommits) | **Get** /v1/git/repos/{name}/commits | Walks a ref&#39;s history newest first, or one path&#39;s history when a path is given.
[**GetGitReposByNameFiles**](GitAPI.md#GetGitReposByNameFiles) | **Get** /v1/git/repos/{name}/files | Returns every file a glob selects at one revision, WITH its bytes and the revision they came from.
[**GetGitReposByNamePulls**](GitAPI.md#GetGitReposByNamePulls) | **Get** /v1/git/repos/{name}/pulls | Returns a repo&#39;s pull requests, newest number first — what is waiting to be reviewed, and what has already landed.
[**GetGitReposByNamePullsByNumber**](GitAPI.md#GetGitReposByNamePullsByNumber) | **Get** /v1/git/repos/{name}/pulls/{number} | Returns one pull request by its per-repo number.
[**GetGitReposByNameReadme**](GitAPI.md#GetGitReposByNameReadme) | **Get** /v1/git/repos/{name}/readme | Returns the README at the tree root as plain text — unrendered, so the caller decides how to present it.
[**GetGitReposByNameRefs**](GitAPI.md#GetGitReposByNameRefs) | **Get** /v1/git/repos/{name}/refs | Lists a repo&#39;s branches, tags and default branch — what a branch picker needs in one call.
[**GetGitReposByNameSubscriptions**](GitAPI.md#GetGitReposByNameSubscriptions) | **Get** /v1/git/repos/{name}/subscriptions | Returns a repo&#39;s Slack subscriptions — which channels the lifecycle notifier posts this repo&#39;s push and deploy events to.
[**GetGitReposByNameTargets**](GitAPI.md#GetGitReposByNameTargets) | **Get** /v1/git/repos/{name}/targets | Returns a repo&#39;s outbound mirror targets — the downstream remotes the mirror reactor pushes to whenever a push lands here.
[**GetGitReposByNameTree**](GitAPI.md#GetGitReposByNameTree) | **Get** /v1/git/repos/{name}/tree | Lists the immediate children of one directory at one revision, directories before files.
[**GetGitRunners**](GitAPI.md#GetGitRunners) | **Get** /v1/git/runners | Returns the daemons registered into this org&#39;s pools, newest first, with when each was last heard from.
[**GetGitRuns**](GitAPI.md#GetGitRuns) | **Get** /v1/git/runs | Returns this org&#39;s runs, newest first.
[**GetGitRunsById**](GitAPI.md#GetGitRunsById) | **Get** /v1/git/runs/{id} | Returns one run.
[**GetGitUsage**](GitAPI.md#GetGitUsage) | **Get** /v1/git/usage | Returns per-repo and total storage bytes for the caller&#39;s org — the queryable, per-tenant number commerce and o11y meter on.
[**GetGitWorkflows**](GitAPI.md#GetGitWorkflows) | **Get** /v1/git/workflows | Reports the workflows a repository declares at a ref and which declared pool would execute each job — the answer to \&quot;would a push here run, and where\&quot;.
[**PatchGitReposByName**](GitAPI.md#PatchGitReposByName) | **Patch** /v1/git/repos/{name} | Flips a repo&#39;s public bit, the one mutable repo setting today.
[**PostGitByOrgByProjectByRepoGitReceivePack**](GitAPI.md#PostGitByOrgByProjectByRepoGitReceivePack) | **Post** /v1/git/{org}/{project}/{repo}/git-receive-pack | Accept a push, and turn it into a build
[**PostGitByOrgByProjectByRepoGitUploadPack**](GitAPI.md#PostGitByOrgByProjectByRepoGitUploadPack) | **Post** /v1/git/{org}/{project}/{repo}/git-upload-pack | Serve a clone or fetch
[**PostGitByOrgByRepoGitReceivePack**](GitAPI.md#PostGitByOrgByRepoGitReceivePack) | **Post** /v1/git/{org}/{repo}/git-receive-pack | Accept a push, and turn it into a build
[**PostGitByOrgByRepoGitUploadPack**](GitAPI.md#PostGitByOrgByRepoGitUploadPack) | **Post** /v1/git/{org}/{repo}/git-upload-pack | Serve a clone or fetch
[**PostGitKeys**](GitAPI.md#PostGitKeys) | **Post** /v1/git/keys | Registers an SSH public key so it can authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60; for the caller&#39;s org.
[**PostGitPools**](GitAPI.md#PostGitPools) | **Post** /v1/git/pools | Records the capacity an org has, and answers with the secret a runner daemon presents to enter it.
[**PostGitRepos**](GitAPI.md#PostGitRepos) | **Post** /v1/git/repos | Provisions an empty bare repository in the caller&#39;s scope and returns it with its clone URLs.
[**PostGitReposByNameGc**](GitAPI.md#PostGitReposByNameGc) | **Post** /v1/git/repos/{name}/gc | Repacks a repo into one bitmapped pack and rewrites its commit-graph, so the next clone reuses the bitmap instead of walking the whole object graph.
[**PostGitReposByNameMirror**](GitAPI.md#PostGitReposByNameMirror) | **Post** /v1/git/repos/{name}/mirror | Imports an external git repository into the caller&#39;s repo, provisioning it on first use.
[**PostGitReposByNamePulls**](GitAPI.md#PostGitReposByNamePulls) | **Post** /v1/git/repos/{name}/pulls | Proposes a branch for merging and returns it with its number.
[**PostGitReposByNamePullsByNumberMerge**](GitAPI.md#PostGitReposByNamePullsByNumberMerge) | **Post** /v1/git/repos/{name}/pulls/{number}/merge | Merges an open pull request by FAST-FORWARDING base to head, and answers the proposal in its merged state with the revision base now points at.
[**PostGitReposByNamePush**](GitAPI.md#PostGitReposByNamePush) | **Post** /v1/git/repos/{name}/push | Lands a set of files as one commit without a git client — the hanzo.app builder&#39;s push.
[**PostGitReposByNameSubscriptions**](GitAPI.md#PostGitReposByNameSubscriptions) | **Post** /v1/git/repos/{name}/subscriptions | Binds a Slack channel to a repo, so the lifecycle notifier posts that repo&#39;s push and deploy events there.
[**PostGitReposByNameTargets**](GitAPI.md#PostGitReposByNameTargets) | **Post** /v1/git/repos/{name}/targets | Registers a downstream remote the repo&#39;s advanced refs are pushed to whenever a push lands here.
[**PostGitRuns**](GitAPI.md#PostGitRuns) | **Post** /v1/git/runs | Runs a repository&#39;s workflows at a ref, on demand.
[**PostGitWebhook**](GitAPI.md#PostGitWebhook) | **Post** /v1/git/webhook | Retired — push-to-deploy has no inbound webhook
[**PostRunnerDeclare**](GitAPI.md#PostRunnerDeclare) | **Post** /v1/runner/declare | Republishes what a registered runner can do, and answers with what this side understands, so the two learn about each other from one exchange.
[**PostRunnerLog**](GitAPI.md#PostRunnerLog) | **Post** /v1/runner/log | Adds console output to a task&#39;s log and answers with how far that log is durable, so the runner knows where to resend from.
[**PostRunnerRegister**](GitAPI.md#PostRunnerRegister) | **Post** /v1/runner/register | Trades a pool&#39;s join secret for a runner identity and the token that authenticates every later call.
[**PostRunnerState**](GitAPI.md#PostRunnerState) | **Post** /v1/runner/state | Records a task&#39;s progress and that of its steps, and answers with the result this side now holds — which is how a runner learns its task was stopped from somewhere else.
[**PostRunnerTask**](GitAPI.md#PostRunnerTask) | **Post** /v1/runner/task | Hands the runner a job to execute, if its pool has one, and answers immediately either way.



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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
- **Accept**: application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
- **Accept**: application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGitReposByNameTargetsById

> DeleteGitReposByNameTargetsById(ctx, name, id).Execute()

Removes one outbound mirror target; later pushes stop being forwarded to it.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	id := "mir_2d90" // string | ID is the row to remove, from the :id path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.DeleteGitReposByNameTargetsById(context.Background(), name, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.DeleteGitReposByNameTargetsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGitReposByNameTargetsByIdRequest struct via the builder pattern


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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> GitKeyList GetGitKeys(ctx).Execute()

Returns the SSH public keys registered to the caller's org — the keys that authenticate `git clone git@<host>:<org>/<repo>.git`.



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
	resp, r, err := apiClient.GitAPI.GetGitKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitKeys`: GitKeyList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitKeysRequest struct via the builder pattern


### Return type

[**GitKeyList**](GitKeyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitPools

> GitPoolList GetGitPools(ctx).Execute()

Returns the capacity this org has declared and how many daemons have entered each pool.



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
	resp, r, err := apiClient.GitAPI.GetGitPools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitPools`: GitPoolList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitPoolsRequest struct via the builder pattern


### Return type

[**GitPoolList**](GitPoolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitRepos

> GitRepoList GetGitRepos(ctx).Execute()

Returns the repos in the caller's scope, most recently updated first.



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
	resp, r, err := apiClient.GitAPI.GetGitRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRepos`: GitRepoList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposRequest struct via the builder pattern


### Return type

[**GitRepoList**](GitRepoList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByName

> GitRepoView GetGitReposByName(ctx, name).Execute()

Returns one repo with its live ref state: every branch name and the resolved HEAD commit.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByName`: GitRepoView
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

[**GitRepoView**](GitRepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameBlob

> GitBlobJSON GetGitReposByNameBlob(ctx, name).Ref(ref).Path(path).Execute()

Returns one file's bytes at one revision.



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
	// response from `GetGitReposByNameBlob`: GitBlobJSON
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

[**GitBlobJSON**](GitBlobJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameCommits

> GitCommitsJSON GetGitReposByNameCommits(ctx, name).Ref(ref).Path(path).Limit(limit).Execute()

Walks a ref's history newest first, or one path's history when a path is given.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is the branch, tag or commit to walk back from; empty means HEAD. (optional)
	path := "path_example" // string | Path narrows the history to commits touching it; empty walks the whole ref. (optional)
	limit := int64(2) // int64 | Limit caps the page. Anything not positive means 50; the cap is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameCommits(context.Background(), name).Ref(ref).Path(path).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameCommits`: GitCommitsJSON
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
 **limit** | **int64** | Limit caps the page. Anything not positive means 50; the cap is 100. | 

### Return type

[**GitCommitsJSON**](GitCommitsJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameFiles

> GitFilesJSON GetGitReposByNameFiles(ctx, name).Ref(ref).Glob(glob).Execute()

Returns every file a glob selects at one revision, WITH its bytes and the revision they came from.



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
	// response from `GetGitReposByNameFiles`: GitFilesJSON
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

[**GitFilesJSON**](GitFilesJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNamePulls

> GitPullList GetGitReposByNamePulls(ctx, name).State(state).Execute()

Returns a repo's pull requests, newest number first — what is waiting to be reviewed, and what has already landed.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	state := "open" // string | State narrows the list to \"open\" or \"merged\". Omit it for every proposal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNamePulls(context.Background(), name).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNamePulls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNamePulls`: GitPullList
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

[**GitPullList**](GitPullList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNamePullsByNumber

> GitPullView GetGitReposByNamePullsByNumber(ctx, name, number).Execute()

Returns one pull request by its per-repo number.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	number := int64(4) // int64 | Number is the proposal's per-repo number, from the :number path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNamePullsByNumber(context.Background(), name, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNamePullsByNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNamePullsByNumber`: GitPullView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNamePullsByNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**number** | **int64** | Number is the proposal&#39;s per-repo number, from the :number path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNamePullsByNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitPullView**](GitPullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameReadme

> GitReadmeJSON GetGitReposByNameReadme(ctx, name).Ref(ref).Execute()

Returns the README at the tree root as plain text — unrendered, so the caller decides how to present it.



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
	name := "widgets" // string | Name is the repo to read, from the :name path segment.
	ref := "main" // string | Ref is a branch, tag or commit; empty means the repo's HEAD. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameReadme(context.Background(), name).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameReadme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameReadme`: GitReadmeJSON
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

[**GitReadmeJSON**](GitReadmeJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameRefs

> GitRefsJSON GetGitReposByNameRefs(ctx, name).Execute()

Lists a repo's branches, tags and default branch — what a branch picker needs in one call.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameRefs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameRefs`: GitRefsJSON
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

[**GitRefsJSON**](GitRefsJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameSubscriptions

> GitSubscriptionList GetGitReposByNameSubscriptions(ctx, name).Execute()

Returns a repo's Slack subscriptions — which channels the lifecycle notifier posts this repo's push and deploy events to.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameSubscriptions(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameSubscriptions`: GitSubscriptionList
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

[**GitSubscriptionList**](GitSubscriptionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameTargets

> GitMirrorList GetGitReposByNameTargets(ctx, name).Execute()

Returns a repo's outbound mirror targets — the downstream remotes the mirror reactor pushes to whenever a push lands here.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitReposByNameTargets(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitReposByNameTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitReposByNameTargets`: GitMirrorList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitReposByNameTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitReposByNameTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitMirrorList**](GitMirrorList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitReposByNameTree

> GitTreeJSON GetGitReposByNameTree(ctx, name).Ref(ref).Path(path).Execute()

Lists the immediate children of one directory at one revision, directories before files.



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
	// response from `GetGitReposByNameTree`: GitTreeJSON
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

[**GitTreeJSON**](GitTreeJSON.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitRunners

> GitRunnerList GetGitRunners(ctx).Execute()

Returns the daemons registered into this org's pools, newest first, with when each was last heard from.



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
	resp, r, err := apiClient.GitAPI.GetGitRunners(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitRunners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRunners`: GitRunnerList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitRunners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitRunnersRequest struct via the builder pattern


### Return type

[**GitRunnerList**](GitRunnerList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitRuns

> GitWorkflowRuns GetGitRuns(ctx).Repo(repo).Limit(limit).Execute()

Returns this org's runs, newest first.



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
	repo := "repo_example" // string | Repo restricts the listing to one repository. Empty lists the whole org. (optional)
	limit := int64(789) // int64 | Limit caps the answer; 0 means the default of 50, and 200 is the ceiling. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitRuns(context.Background()).Repo(repo).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRuns`: GitWorkflowRuns
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGitRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | Repo restricts the listing to one repository. Empty lists the whole org. | 
 **limit** | **int64** | Limit caps the answer; 0 means the default of 50, and 200 is the ceiling. | 

### Return type

[**GitWorkflowRuns**](GitWorkflowRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitRunsById

> GitWorkflowRun GetGitRunsById(ctx, id).Execute()

Returns one run.



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
	id := "id_example" // string | ID is the run to read, from the :id path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitRunsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitRunsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitRunsById`: GitWorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitRunsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the run to read, from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitRunsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitWorkflowRun**](GitWorkflowRun.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitUsage

> GitUsageView GetGitUsage(ctx).Execute()

Returns per-repo and total storage bytes for the caller's org — the queryable, per-tenant number commerce and o11y meter on.



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
	resp, r, err := apiClient.GitAPI.GetGitUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitUsage`: GitUsageView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitUsageRequest struct via the builder pattern


### Return type

[**GitUsageView**](GitUsageView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitWorkflows

> GitWorkflowList GetGitWorkflows(ctx).Repo(repo).Ref(ref).Execute()

Reports the workflows a repository declares at a ref and which declared pool would execute each job — the answer to \"would a push here run, and where\".



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
	repo := "repo_example" // string | Repo is the repository whose workflows to read. (optional)
	ref := "ref_example" // string | Ref is the branch to read them at; empty means the default. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.GetGitWorkflows(context.Background()).Repo(repo).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.GetGitWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitWorkflows`: GitWorkflowList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.GetGitWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGitWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **string** | Repo is the repository whose workflows to read. | 
 **ref** | **string** | Ref is the branch to read them at; empty means the default. | 

### Return type

[**GitWorkflowList**](GitWorkflowList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGitReposByName

> GitRepoView PatchGitReposByName(ctx, name).GitPatchIn(gitPatchIn).Execute()

Flips a repo's public bit, the one mutable repo setting today.



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
	name := "widgets" // string | Name is the repo to update, from the :name path segment.
	gitPatchIn := *openapiclient.NewGitPatchIn() // GitPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PatchGitReposByName(context.Background(), name).GitPatchIn(gitPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PatchGitReposByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchGitReposByName`: GitRepoView
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

 **gitPatchIn** | [**GitPatchIn**](GitPatchIn.md) |  | 

### Return type

[**GitRepoView**](GitRepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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
	openapiclient "github.com/hanzoai/go-sdk/v8"
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

> GitKeyView PostGitKeys(ctx).GitRegisterKeyReq(gitRegisterKeyReq).Execute()

Registers an SSH public key so it can authenticate `git clone git@<host>:<org>/<repo>.git` for the caller's org.



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
	gitRegisterKeyReq := *openapiclient.NewGitRegisterKeyReq() // GitRegisterKeyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitKeys(context.Background()).GitRegisterKeyReq(gitRegisterKeyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitKeys`: GitKeyView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitRegisterKeyReq** | [**GitRegisterKeyReq**](GitRegisterKeyReq.md) |  | 

### Return type

[**GitKeyView**](GitKeyView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitPools

> GitPoolDeclared PostGitPools(ctx).GitPoolDeclare(gitPoolDeclare).Execute()

Records the capacity an org has, and answers with the secret a runner daemon presents to enter it.



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
	gitPoolDeclare := *openapiclient.NewGitPoolDeclare() // GitPoolDeclare | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitPools(context.Background()).GitPoolDeclare(gitPoolDeclare).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitPools`: GitPoolDeclared
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitPoolDeclare** | [**GitPoolDeclare**](GitPoolDeclare.md) |  | 

### Return type

[**GitPoolDeclared**](GitPoolDeclared.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitRepos

> GitRepoView PostGitRepos(ctx).GitCreateReq(gitCreateReq).Execute()

Provisions an empty bare repository in the caller's scope and returns it with its clone URLs.



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
	gitCreateReq := *openapiclient.NewGitCreateReq() // GitCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitRepos(context.Background()).GitCreateReq(gitCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitRepos`: GitRepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitCreateReq** | [**GitCreateReq**](GitCreateReq.md) |  | 

### Return type

[**GitRepoView**](GitRepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameGc

> GitGcOut PostGitReposByNameGc(ctx, name).Execute()

Repacks a repo into one bitmapped pack and rewrites its commit-graph, so the next clone reuses the bitmap instead of walking the whole object graph.



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
	name := "widgets" // string | Name is the repo's org-unique handle, from the :name path segment. A trailing \".git\" is stripped.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameGc(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameGc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameGc`: GitGcOut
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

[**GitGcOut**](GitGcOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameMirror

> GitRepoView PostGitReposByNameMirror(ctx, name).GitMirrorReq(gitMirrorReq).Execute()

Imports an external git repository into the caller's repo, provisioning it on first use.



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
	name := "widgets" // string | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use.
	gitMirrorReq := *openapiclient.NewGitMirrorReq() // GitMirrorReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameMirror(context.Background(), name).GitMirrorReq(gitMirrorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameMirror``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameMirror`: GitRepoView
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

 **gitMirrorReq** | [**GitMirrorReq**](GitMirrorReq.md) |  | 

### Return type

[**GitRepoView**](GitRepoView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePulls

> GitPullView PostGitReposByNamePulls(ctx, name).GitOpenReq(gitOpenReq).Execute()

Proposes a branch for merging and returns it with its number.



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
	name := "widgets" // string | Name is the repo the proposal belongs to, from the :name path segment.
	gitOpenReq := *openapiclient.NewGitOpenReq() // GitOpenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePulls(context.Background(), name).GitOpenReq(gitOpenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePulls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePulls`: GitPullView
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

 **gitOpenReq** | [**GitOpenReq**](GitOpenReq.md) |  | 

### Return type

[**GitPullView**](GitPullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePullsByNumberMerge

> GitPullView PostGitReposByNamePullsByNumberMerge(ctx, name, number).Execute()

Merges an open pull request by FAST-FORWARDING base to head, and answers the proposal in its merged state with the revision base now points at.



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
	name := "widgets" // string | Name is the repo, from the :name path segment.
	number := int64(4) // int64 | Number is the proposal's per-repo number, from the :number path segment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePullsByNumberMerge(context.Background(), name, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePullsByNumberMerge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePullsByNumberMerge`: GitPullView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNamePullsByNumberMerge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**number** | **int64** | Number is the proposal&#39;s per-repo number, from the :number path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNamePullsByNumberMergeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitPullView**](GitPullView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNamePush

> GitPushResp PostGitReposByNamePush(ctx, name).GitPushReq(gitPushReq).Execute()

Lands a set of files as one commit without a git client — the hanzo.app builder's push.



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
	name := "widgets" // string | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist.
	gitPushReq := *openapiclient.NewGitPushReq() // GitPushReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNamePush(context.Background(), name).GitPushReq(gitPushReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNamePush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNamePush`: GitPushResp
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

 **gitPushReq** | [**GitPushReq**](GitPushReq.md) |  | 

### Return type

[**GitPushResp**](GitPushResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameSubscriptions

> GitSubscriptionView PostGitReposByNameSubscriptions(ctx, name).GitSubscribeReq(gitSubscribeReq).Execute()

Binds a Slack channel to a repo, so the lifecycle notifier posts that repo's push and deploy events there.



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
	name := "widgets" // string | Name is the repo to subscribe, from the :name path segment.
	gitSubscribeReq := *openapiclient.NewGitSubscribeReq() // GitSubscribeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameSubscriptions(context.Background(), name).GitSubscribeReq(gitSubscribeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameSubscriptions`: GitSubscriptionView
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

 **gitSubscribeReq** | [**GitSubscribeReq**](GitSubscribeReq.md) |  | 

### Return type

[**GitSubscriptionView**](GitSubscriptionView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitReposByNameTargets

> GitMirrorTargetView PostGitReposByNameTargets(ctx, name).GitMirrorTargetReq(gitMirrorTargetReq).Execute()

Registers a downstream remote the repo's advanced refs are pushed to whenever a push lands here.



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
	name := "widgets" // string | Name is the repo whose advanced refs are pushed downstream, from the :name path segment.
	gitMirrorTargetReq := *openapiclient.NewGitMirrorTargetReq() // GitMirrorTargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitReposByNameTargets(context.Background(), name).GitMirrorTargetReq(gitMirrorTargetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitReposByNameTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitReposByNameTargets`: GitMirrorTargetView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitReposByNameTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGitReposByNameTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gitMirrorTargetReq** | [**GitMirrorTargetReq**](GitMirrorTargetReq.md) |  | 

### Return type

[**GitMirrorTargetView**](GitMirrorTargetView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitRuns

> GitWorkflowRuns PostGitRuns(ctx).GitRunStart(gitRunStart).Execute()

Runs a repository's workflows at a ref, on demand.



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
	gitRunStart := *openapiclient.NewGitRunStart() // GitRunStart | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostGitRuns(context.Background()).GitRunStart(gitRunStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostGitRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGitRuns`: GitWorkflowRuns
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostGitRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGitRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitRunStart** | [**GitRunStart**](GitRunStart.md) |  | 

### Return type

[**GitWorkflowRuns**](GitWorkflowRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGitWebhook

> PostGitWebhook(ctx).Execute()

Retired — push-to-deploy has no inbound webhook



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


## PostRunnerDeclare

> RunnerDeclareOut PostRunnerDeclare(ctx).RunnerDeclareIn(runnerDeclareIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()

Republishes what a registered runner can do, and answers with what this side understands, so the two learn about each other from one exchange.



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
	runnerDeclareIn := *openapiclient.NewRunnerDeclareIn() // RunnerDeclareIn | 
	xRunnerUuid := "xRunnerUuid_example" // string |  (optional)
	xRunnerToken := "xRunnerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostRunnerDeclare(context.Background()).RunnerDeclareIn(runnerDeclareIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostRunnerDeclare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerDeclare`: RunnerDeclareOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostRunnerDeclare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerDeclareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerDeclareIn** | [**RunnerDeclareIn**](RunnerDeclareIn.md) |  | 
 **xRunnerUuid** | **string** |  | 
 **xRunnerToken** | **string** |  | 

### Return type

[**RunnerDeclareOut**](RunnerDeclareOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerLog

> RunnerLogOut PostRunnerLog(ctx).RunnerLogIn(runnerLogIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()

Adds console output to a task's log and answers with how far that log is durable, so the runner knows where to resend from.



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
	runnerLogIn := *openapiclient.NewRunnerLogIn() // RunnerLogIn | 
	xRunnerUuid := "xRunnerUuid_example" // string |  (optional)
	xRunnerToken := "xRunnerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostRunnerLog(context.Background()).RunnerLogIn(runnerLogIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostRunnerLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerLog`: RunnerLogOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostRunnerLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerLogIn** | [**RunnerLogIn**](RunnerLogIn.md) |  | 
 **xRunnerUuid** | **string** |  | 
 **xRunnerToken** | **string** |  | 

### Return type

[**RunnerLogOut**](RunnerLogOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerRegister

> RunnerRegisterOut PostRunnerRegister(ctx).RunnerRegisterIn(runnerRegisterIn).Execute()

Trades a pool's join secret for a runner identity and the token that authenticates every later call.



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
	runnerRegisterIn := *openapiclient.NewRunnerRegisterIn() // RunnerRegisterIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostRunnerRegister(context.Background()).RunnerRegisterIn(runnerRegisterIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostRunnerRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerRegister`: RunnerRegisterOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostRunnerRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerRegisterIn** | [**RunnerRegisterIn**](RunnerRegisterIn.md) |  | 

### Return type

[**RunnerRegisterOut**](RunnerRegisterOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerState

> RunnerStateOut PostRunnerState(ctx).RunnerStateIn(runnerStateIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()

Records a task's progress and that of its steps, and answers with the result this side now holds — which is how a runner learns its task was stopped from somewhere else.



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
	runnerStateIn := *openapiclient.NewRunnerStateIn() // RunnerStateIn | 
	xRunnerUuid := "xRunnerUuid_example" // string |  (optional)
	xRunnerToken := "xRunnerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostRunnerState(context.Background()).RunnerStateIn(runnerStateIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostRunnerState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerState`: RunnerStateOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostRunnerState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerStateIn** | [**RunnerStateIn**](RunnerStateIn.md) |  | 
 **xRunnerUuid** | **string** |  | 
 **xRunnerToken** | **string** |  | 

### Return type

[**RunnerStateOut**](RunnerStateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRunnerTask

> RunnerTaskOut PostRunnerTask(ctx).RunnerTaskIn(runnerTaskIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()

Hands the runner a job to execute, if its pool has one, and answers immediately either way.



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
	runnerTaskIn := *openapiclient.NewRunnerTaskIn() // RunnerTaskIn | 
	xRunnerUuid := "xRunnerUuid_example" // string |  (optional)
	xRunnerToken := "xRunnerToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.PostRunnerTask(context.Background()).RunnerTaskIn(runnerTaskIn).XRunnerUuid(xRunnerUuid).XRunnerToken(xRunnerToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.PostRunnerTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRunnerTask`: RunnerTaskOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.PostRunnerTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRunnerTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runnerTaskIn** | [**RunnerTaskIn**](RunnerTaskIn.md) |  | 
 **xRunnerUuid** | **string** |  | 
 **xRunnerToken** | **string** |  | 

### Return type

[**RunnerTaskOut**](RunnerTaskOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

