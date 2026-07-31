# \GitAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1GitKeysId**](GitAPI.md#CloudDeleteV1GitKeysId) | **Delete** /v1/git/keys/{id} | Removes a registered SSH key, scoped to the caller&#39;s org: an org can only delete its own, and a key id it does not own is not found.
[**CloudDeleteV1GitReposName**](GitAPI.md#CloudDeleteV1GitReposName) | **Delete** /v1/git/repos/{name} | Removes a repo&#39;s metadata and purges its storage.
[**CloudDeleteV1GitReposNameMirrorsId**](GitAPI.md#CloudDeleteV1GitReposNameMirrorsId) | **Delete** /v1/git/repos/{name}/mirrors/{id} | Removes one outbound mirror target; later pushes stop being forwarded to it.
[**CloudDeleteV1GitReposNameSubscriptionsId**](GitAPI.md#CloudDeleteV1GitReposNameSubscriptionsId) | **Delete** /v1/git/repos/{name}/subscriptions/{id} | Removes one Slack subscription from a repo; the notifier stops posting that repo&#39;s events to that channel.
[**CloudGetV1GitByOrgByProjectByRepoInfoRefs**](GitAPI.md#CloudGetV1GitByOrgByProjectByRepoInfoRefs) | **Get** /v1/git/{org}/{project}/{repo}/info/refs | 
[**CloudGetV1GitByOrgByRepoInfoRefs**](GitAPI.md#CloudGetV1GitByOrgByRepoInfoRefs) | **Get** /v1/git/{org}/{repo}/info/refs | 
[**CloudGetV1GitKeys**](GitAPI.md#CloudGetV1GitKeys) | **Get** /v1/git/keys | Returns the SSH public keys registered to the caller&#39;s org — the keys that authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60;.
[**CloudGetV1GitRepos**](GitAPI.md#CloudGetV1GitRepos) | **Get** /v1/git/repos | Returns the repos in the caller&#39;s scope, most recently updated first.
[**CloudGetV1GitReposName**](GitAPI.md#CloudGetV1GitReposName) | **Get** /v1/git/repos/{name} | Returns one repo with its live ref state: every branch name and the resolved HEAD commit.
[**CloudGetV1GitReposNameBlob**](GitAPI.md#CloudGetV1GitReposNameBlob) | **Get** /v1/git/repos/{name}/blob | Returns one file&#39;s bytes at one revision.
[**CloudGetV1GitReposNameCommits**](GitAPI.md#CloudGetV1GitReposNameCommits) | **Get** /v1/git/repos/{name}/commits | Walks a ref&#39;s history newest first, or one path&#39;s history when a path is given.
[**CloudGetV1GitReposNameFiles**](GitAPI.md#CloudGetV1GitReposNameFiles) | **Get** /v1/git/repos/{name}/files | Returns every file a glob selects at one revision, WITH its bytes and the revision they came from.
[**CloudGetV1GitReposNameMirrors**](GitAPI.md#CloudGetV1GitReposNameMirrors) | **Get** /v1/git/repos/{name}/mirrors | Returns a repo&#39;s outbound mirror targets — the downstream remotes the mirror reactor pushes to whenever a push lands here.
[**CloudGetV1GitReposNameReadme**](GitAPI.md#CloudGetV1GitReposNameReadme) | **Get** /v1/git/repos/{name}/readme | Returns the README at the tree root as plain text — unrendered, so the caller decides how to present it.
[**CloudGetV1GitReposNameRefs**](GitAPI.md#CloudGetV1GitReposNameRefs) | **Get** /v1/git/repos/{name}/refs | Lists a repo&#39;s branches, tags and default branch — what a branch picker needs in one call.
[**CloudGetV1GitReposNameSubscriptions**](GitAPI.md#CloudGetV1GitReposNameSubscriptions) | **Get** /v1/git/repos/{name}/subscriptions | Returns a repo&#39;s Slack subscriptions — which channels the lifecycle notifier posts this repo&#39;s push and deploy events to.
[**CloudGetV1GitReposNameTree**](GitAPI.md#CloudGetV1GitReposNameTree) | **Get** /v1/git/repos/{name}/tree | Lists the immediate children of one directory at one revision, directories before files.
[**CloudGetV1GitUsage**](GitAPI.md#CloudGetV1GitUsage) | **Get** /v1/git/usage | Returns per-repo and total storage bytes for the caller&#39;s org — the queryable, per-tenant number commerce and o11y meter on.
[**CloudPatchV1GitReposName**](GitAPI.md#CloudPatchV1GitReposName) | **Patch** /v1/git/repos/{name} | Flips a repo&#39;s public bit, the one mutable repo setting today.
[**CloudPostV1GitByOrgByProjectByRepoGitReceivePack**](GitAPI.md#CloudPostV1GitByOrgByProjectByRepoGitReceivePack) | **Post** /v1/git/{org}/{project}/{repo}/git-receive-pack | 
[**CloudPostV1GitByOrgByProjectByRepoGitUploadPack**](GitAPI.md#CloudPostV1GitByOrgByProjectByRepoGitUploadPack) | **Post** /v1/git/{org}/{project}/{repo}/git-upload-pack | 
[**CloudPostV1GitByOrgByRepoGitReceivePack**](GitAPI.md#CloudPostV1GitByOrgByRepoGitReceivePack) | **Post** /v1/git/{org}/{repo}/git-receive-pack | 
[**CloudPostV1GitByOrgByRepoGitUploadPack**](GitAPI.md#CloudPostV1GitByOrgByRepoGitUploadPack) | **Post** /v1/git/{org}/{repo}/git-upload-pack | 
[**CloudPostV1GitKeys**](GitAPI.md#CloudPostV1GitKeys) | **Post** /v1/git/keys | Registers an SSH public key so it can authenticate &#x60;git clone git@&lt;host&gt;:&lt;org&gt;/&lt;repo&gt;.git&#x60; for the caller&#39;s org.
[**CloudPostV1GitRepos**](GitAPI.md#CloudPostV1GitRepos) | **Post** /v1/git/repos | Provisions an empty bare repository in the caller&#39;s scope and returns it with its clone URLs.
[**CloudPostV1GitReposNameGc**](GitAPI.md#CloudPostV1GitReposNameGc) | **Post** /v1/git/repos/{name}/gc | Repacks a repo into one bitmapped pack and rewrites its commit-graph, so the next clone reuses the bitmap instead of walking the whole object graph.
[**CloudPostV1GitReposNameMirror**](GitAPI.md#CloudPostV1GitReposNameMirror) | **Post** /v1/git/repos/{name}/mirror | Imports an external git repository into the caller&#39;s repo, provisioning it on first use.
[**CloudPostV1GitReposNameMirrors**](GitAPI.md#CloudPostV1GitReposNameMirrors) | **Post** /v1/git/repos/{name}/mirrors | Registers a downstream remote the repo&#39;s advanced refs are pushed to whenever a push lands here.
[**CloudPostV1GitReposNamePush**](GitAPI.md#CloudPostV1GitReposNamePush) | **Post** /v1/git/repos/{name}/push | Lands a set of files as one commit without a git client — the hanzo.app builder&#39;s push.
[**CloudPostV1GitReposNameSubscriptions**](GitAPI.md#CloudPostV1GitReposNameSubscriptions) | **Post** /v1/git/repos/{name}/subscriptions | Binds a Slack channel to a repo, so the lifecycle notifier posts that repo&#39;s push and deploy events there.
[**CloudPostV1GitWebhook**](GitAPI.md#CloudPostV1GitWebhook) | **Post** /v1/git/webhook | 
[**CloudPostV1GitZapCreaterepo**](GitAPI.md#CloudPostV1GitZapCreaterepo) | **Post** /v1/git/zap/createRepo | 
[**CloudPostV1GitZapDeleterepo**](GitAPI.md#CloudPostV1GitZapDeleterepo) | **Post** /v1/git/zap/deleteRepo | 
[**CloudPostV1GitZapGetrepo**](GitAPI.md#CloudPostV1GitZapGetrepo) | **Post** /v1/git/zap/getRepo | 
[**CloudPostV1GitZapListrepos**](GitAPI.md#CloudPostV1GitZapListrepos) | **Post** /v1/git/zap/listRepos | 
[**CloudPostV1GitZapUsage**](GitAPI.md#CloudPostV1GitZapUsage) | **Post** /v1/git/zap/usage | 



## CloudDeleteV1GitKeysId

> map[string]interface{} CloudDeleteV1GitKeysId(ctx, id).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudDeleteV1GitKeysId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudDeleteV1GitKeysId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1GitKeysId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudDeleteV1GitKeysId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the key&#39;s identifier (\&quot;gitkey_…\&quot;), from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1GitKeysIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1GitReposName

> map[string]interface{} CloudDeleteV1GitReposName(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudDeleteV1GitReposName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudDeleteV1GitReposName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1GitReposName`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudDeleteV1GitReposName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1GitReposNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1GitReposNameMirrorsId

> map[string]interface{} CloudDeleteV1GitReposNameMirrorsId(ctx, name, id).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudDeleteV1GitReposNameMirrorsId(context.Background(), name, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudDeleteV1GitReposNameMirrorsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1GitReposNameMirrorsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudDeleteV1GitReposNameMirrorsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**id** | **string** | ID is the row to remove, from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1GitReposNameMirrorsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1GitReposNameSubscriptionsId

> map[string]interface{} CloudDeleteV1GitReposNameSubscriptionsId(ctx, name, id).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudDeleteV1GitReposNameSubscriptionsId(context.Background(), name, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudDeleteV1GitReposNameSubscriptionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1GitReposNameSubscriptionsId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudDeleteV1GitReposNameSubscriptionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo, from the :name path segment. | 
**id** | **string** | ID is the row to remove, from the :id path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1GitReposNameSubscriptionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitByOrgByProjectByRepoInfoRefs

> CloudGetV1GitByOrgByProjectByRepoInfoRefs(ctx, org, project, repo).Execute()



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
	r, err := apiClient.GitAPI.CloudGetV1GitByOrgByProjectByRepoInfoRefs(context.Background(), org, project, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitByOrgByProjectByRepoInfoRefs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1GitByOrgByProjectByRepoInfoRefsRequest struct via the builder pattern


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


## CloudGetV1GitByOrgByRepoInfoRefs

> CloudGetV1GitByOrgByRepoInfoRefs(ctx, org, repo).Execute()



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
	r, err := apiClient.GitAPI.CloudGetV1GitByOrgByRepoInfoRefs(context.Background(), org, repo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitByOrgByRepoInfoRefs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudGetV1GitByOrgByRepoInfoRefsRequest struct via the builder pattern


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


## CloudGetV1GitKeys

> CloudKeyList CloudGetV1GitKeys(ctx).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitKeys`: CloudKeyList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitKeysRequest struct via the builder pattern


### Return type

[**CloudKeyList**](CloudKeyList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitRepos

> CloudRepoList CloudGetV1GitRepos(ctx).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitRepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitRepos`: CloudRepoList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitRepos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposRequest struct via the builder pattern


### Return type

[**CloudRepoList**](CloudRepoList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposName

> CloudRepoView CloudGetV1GitReposName(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposName`: CloudRepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRepoView**](CloudRepoView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameBlob

> CloudBlobJSON CloudGetV1GitReposNameBlob(ctx, name).Ref(ref).Path(path).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameBlob(context.Background(), name).Ref(ref).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameBlob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameBlob`: CloudBlobJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameBlob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameBlobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **path** | **string** | Path is repo-relative; empty is the tree root. Traversal is stripped. | 

### Return type

[**CloudBlobJSON**](CloudBlobJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameCommits

> CloudCommitsJSON CloudGetV1GitReposNameCommits(ctx, name).Ref(ref).Path(path).Limit(limit).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameCommits(context.Background(), name).Ref(ref).Path(path).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameCommits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameCommits`: CloudCommitsJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is the branch, tag or commit to walk back from; empty means HEAD. | 
 **path** | **string** | Path narrows the history to commits touching it; empty walks the whole ref. | 
 **limit** | **int32** | Limit caps the page. Anything not positive means 50; the cap is 100. | 

### Return type

[**CloudCommitsJSON**](CloudCommitsJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameFiles

> CloudFilesJSON CloudGetV1GitReposNameFiles(ctx, name).Ref(ref).Glob(glob).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameFiles(context.Background(), name).Ref(ref).Glob(glob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameFiles`: CloudFilesJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **glob** | **string** | Glob selects files, matched segment by segment so &#x60;*&#x60; never crosses a &#x60;/&#x60;. &#x60;**&#x60; matches zero or more whole segments. | 

### Return type

[**CloudFilesJSON**](CloudFilesJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameMirrors

> CloudMirrorList CloudGetV1GitReposNameMirrors(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameMirrors(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameMirrors`: CloudMirrorList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMirrorList**](CloudMirrorList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameReadme

> CloudReadmeJSON CloudGetV1GitReposNameReadme(ctx, name).Ref(ref).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameReadme(context.Background(), name).Ref(ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameReadme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameReadme`: CloudReadmeJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameReadme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameReadmeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 

### Return type

[**CloudReadmeJSON**](CloudReadmeJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameRefs

> CloudRefsJSON CloudGetV1GitReposNameRefs(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameRefs(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameRefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameRefs`: CloudRefsJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameRefs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameRefsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRefsJSON**](CloudRefsJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameSubscriptions

> CloudSubscriptionList CloudGetV1GitReposNameSubscriptions(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameSubscriptions(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameSubscriptions`: CloudSubscriptionList
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSubscriptionList**](CloudSubscriptionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitReposNameTree

> CloudTreeJSON CloudGetV1GitReposNameTree(ctx, name).Ref(ref).Path(path).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitReposNameTree(context.Background(), name).Ref(ref).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitReposNameTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitReposNameTree`: CloudTreeJSON
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitReposNameTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to read, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitReposNameTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | Ref is a branch, tag or commit; empty means the repo&#39;s HEAD. | 
 **path** | **string** | Path is repo-relative; empty is the tree root. Traversal is stripped. | 

### Return type

[**CloudTreeJSON**](CloudTreeJSON.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GitUsage

> CloudUsageView CloudGetV1GitUsage(ctx).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudGetV1GitUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudGetV1GitUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GitUsage`: CloudUsageView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudGetV1GitUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GitUsageRequest struct via the builder pattern


### Return type

[**CloudUsageView**](CloudUsageView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1GitReposName

> CloudRepoView CloudPatchV1GitReposName(ctx, name).CloudPatchIn(cloudPatchIn).Execute()

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
	cloudPatchIn := *openapiclient.NewCloudPatchIn() // CloudPatchIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPatchV1GitReposName(context.Background(), name).CloudPatchIn(cloudPatchIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPatchV1GitReposName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1GitReposName`: CloudRepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPatchV1GitReposName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to update, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1GitReposNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchIn** | [**CloudPatchIn**](CloudPatchIn.md) |  | 

### Return type

[**CloudRepoView**](CloudRepoView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitByOrgByProjectByRepoGitReceivePack

> CloudPostV1GitByOrgByProjectByRepoGitReceivePack(ctx, org, project, repo).Body(body).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitByOrgByProjectByRepoGitReceivePack(context.Background(), org, project, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitByOrgByProjectByRepoGitReceivePack``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GitByOrgByProjectByRepoGitReceivePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitByOrgByProjectByRepoGitUploadPack

> CloudPostV1GitByOrgByProjectByRepoGitUploadPack(ctx, org, project, repo).Body(body).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitByOrgByProjectByRepoGitUploadPack(context.Background(), org, project, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitByOrgByProjectByRepoGitUploadPack``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GitByOrgByProjectByRepoGitUploadPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitByOrgByRepoGitReceivePack

> CloudPostV1GitByOrgByRepoGitReceivePack(ctx, org, repo).Body(body).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitByOrgByRepoGitReceivePack(context.Background(), org, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitByOrgByRepoGitReceivePack``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GitByOrgByRepoGitReceivePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitByOrgByRepoGitUploadPack

> CloudPostV1GitByOrgByRepoGitUploadPack(ctx, org, repo).Body(body).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitByOrgByRepoGitUploadPack(context.Background(), org, repo).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitByOrgByRepoGitUploadPack``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GitByOrgByRepoGitUploadPackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitKeys

> CloudKeyView CloudPostV1GitKeys(ctx).CloudRegisterKeyReq(cloudRegisterKeyReq).Execute()

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
	cloudRegisterKeyReq := *openapiclient.NewCloudRegisterKeyReq() // CloudRegisterKeyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitKeys(context.Background()).CloudRegisterKeyReq(cloudRegisterKeyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitKeys`: CloudKeyView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRegisterKeyReq** | [**CloudRegisterKeyReq**](CloudRegisterKeyReq.md) |  | 

### Return type

[**CloudKeyView**](CloudKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitRepos

> CloudRepoView CloudPostV1GitRepos(ctx).CloudCreateReq(cloudCreateReq).Execute()

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
	cloudCreateReq := *openapiclient.NewCloudCreateReq() // CloudCreateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitRepos(context.Background()).CloudCreateReq(cloudCreateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitRepos`: CloudRepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateReq** | [**CloudCreateReq**](CloudCreateReq.md) |  | 

### Return type

[**CloudRepoView**](CloudRepoView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitReposNameGc

> CloudGcOut CloudPostV1GitReposNameGc(ctx, name).Execute()

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
	resp, r, err := apiClient.GitAPI.CloudPostV1GitReposNameGc(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitReposNameGc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitReposNameGc`: CloudGcOut
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitReposNameGc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo&#39;s org-unique handle, from the :name path segment. A trailing \&quot;.git\&quot; is stripped. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposNameGcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudGcOut**](CloudGcOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitReposNameMirror

> CloudRepoView CloudPostV1GitReposNameMirror(ctx, name).CloudMirrorReq(cloudMirrorReq).Execute()

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
	cloudMirrorReq := *openapiclient.NewCloudMirrorReq() // CloudMirrorReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitReposNameMirror(context.Background(), name).CloudMirrorReq(cloudMirrorReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitReposNameMirror``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitReposNameMirror`: CloudRepoView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitReposNameMirror`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the local repo to mirror into, from the :name path segment. It is CREATED on first use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposNameMirrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudMirrorReq** | [**CloudMirrorReq**](CloudMirrorReq.md) |  | 

### Return type

[**CloudRepoView**](CloudRepoView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitReposNameMirrors

> CloudMirrorTargetView CloudPostV1GitReposNameMirrors(ctx, name).CloudMirrorTargetReq(cloudMirrorTargetReq).Execute()

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
	cloudMirrorTargetReq := *openapiclient.NewCloudMirrorTargetReq() // CloudMirrorTargetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitReposNameMirrors(context.Background(), name).CloudMirrorTargetReq(cloudMirrorTargetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitReposNameMirrors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitReposNameMirrors`: CloudMirrorTargetView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitReposNameMirrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo whose advanced refs are pushed downstream, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposNameMirrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudMirrorTargetReq** | [**CloudMirrorTargetReq**](CloudMirrorTargetReq.md) |  | 

### Return type

[**CloudMirrorTargetView**](CloudMirrorTargetView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitReposNamePush

> CloudPushResp CloudPostV1GitReposNamePush(ctx, name).CloudPushReq(cloudPushReq).Execute()

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
	cloudPushReq := *openapiclient.NewCloudPushReq() // CloudPushReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitReposNamePush(context.Background(), name).CloudPushReq(cloudPushReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitReposNamePush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitReposNamePush`: CloudPushResp
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitReposNamePush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to push into, from the :name path segment. It is CREATED on first push if it does not exist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposNamePushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPushReq** | [**CloudPushReq**](CloudPushReq.md) |  | 

### Return type

[**CloudPushResp**](CloudPushResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitReposNameSubscriptions

> CloudSubscriptionView CloudPostV1GitReposNameSubscriptions(ctx, name).CloudSubscribeReq(cloudSubscribeReq).Execute()

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
	cloudSubscribeReq := *openapiclient.NewCloudSubscribeReq() // CloudSubscribeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitAPI.CloudPostV1GitReposNameSubscriptions(context.Background(), name).CloudSubscribeReq(cloudSubscribeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitReposNameSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GitReposNameSubscriptions`: CloudSubscriptionView
	fmt.Fprintf(os.Stdout, "Response from `GitAPI.CloudPostV1GitReposNameSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the repo to subscribe, from the :name path segment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitReposNameSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSubscribeReq** | [**CloudSubscribeReq**](CloudSubscribeReq.md) |  | 

### Return type

[**CloudSubscriptionView**](CloudSubscriptionView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitWebhook

> CloudPostV1GitWebhook(ctx).CloudPushEvent(cloudPushEvent).Execute()



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
	cloudPushEvent := *openapiclient.NewCloudPushEvent() // CloudPushEvent |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.CloudPostV1GitWebhook(context.Background()).CloudPushEvent(cloudPushEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPushEvent** | [**CloudPushEvent**](CloudPushEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitZapCreaterepo

> CloudPostV1GitZapCreaterepo(ctx).CloudZapProcReq(cloudZapProcReq).Execute()



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
	cloudZapProcReq := *openapiclient.NewCloudZapProcReq() // CloudZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.CloudPostV1GitZapCreaterepo(context.Background()).CloudZapProcReq(cloudZapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitZapCreaterepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitZapCreaterepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudZapProcReq** | [**CloudZapProcReq**](CloudZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitZapDeleterepo

> CloudPostV1GitZapDeleterepo(ctx).CloudZapProcReq(cloudZapProcReq).Execute()



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
	cloudZapProcReq := *openapiclient.NewCloudZapProcReq() // CloudZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.CloudPostV1GitZapDeleterepo(context.Background()).CloudZapProcReq(cloudZapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitZapDeleterepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitZapDeleterepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudZapProcReq** | [**CloudZapProcReq**](CloudZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitZapGetrepo

> CloudPostV1GitZapGetrepo(ctx).CloudZapProcReq(cloudZapProcReq).Execute()



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
	cloudZapProcReq := *openapiclient.NewCloudZapProcReq() // CloudZapProcReq |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitAPI.CloudPostV1GitZapGetrepo(context.Background()).CloudZapProcReq(cloudZapProcReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitZapGetrepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitZapGetrepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudZapProcReq** | [**CloudZapProcReq**](CloudZapProcReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GitZapListrepos

> CloudPostV1GitZapListrepos(ctx).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitZapListrepos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitZapListrepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitZapListreposRequest struct via the builder pattern


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


## CloudPostV1GitZapUsage

> CloudPostV1GitZapUsage(ctx).Execute()



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
	r, err := apiClient.GitAPI.CloudPostV1GitZapUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitAPI.CloudPostV1GitZapUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GitZapUsageRequest struct via the builder pattern


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

