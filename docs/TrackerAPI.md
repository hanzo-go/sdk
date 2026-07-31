# \TrackerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1TrackerProjectsKey**](TrackerAPI.md#CloudDeleteV1TrackerProjectsKey) | **Delete** /v1/tracker/projects/{key} | DeleteProject removes one tracker project of the caller&#39;s org AND every issue filed under it, and answers 204 with no body.
[**CloudDeleteV1TrackerProjectsKeyIssuesNum**](TrackerAPI.md#CloudDeleteV1TrackerProjectsKeyIssuesNum) | **Delete** /v1/tracker/projects/{key}/issues/{num} | DeleteIssue removes one issue from a tracker project and answers 204 with no body.
[**CloudGetV1TrackerProjects**](TrackerAPI.md#CloudGetV1TrackerProjects) | **Get** /v1/tracker/projects | ListProjects returns every tracker project in the caller&#39;s org, newest first.
[**CloudGetV1TrackerProjectsKey**](TrackerAPI.md#CloudGetV1TrackerProjectsKey) | **Get** /v1/tracker/projects/{key} | GetProject returns one tracker project of the caller&#39;s org by its key — its name, description and timestamps.
[**CloudGetV1TrackerProjectsKeyIssues**](TrackerAPI.md#CloudGetV1TrackerProjectsKeyIssues) | **Get** /v1/tracker/projects/{key}/issues | ListIssues returns the issues of one tracker project, optionally filtered by status, kind, repo and source.
[**CloudGetV1TrackerProjectsKeyIssuesNum**](TrackerAPI.md#CloudGetV1TrackerProjectsKeyIssuesNum) | **Get** /v1/tracker/projects/{key}/issues/{num} | GetIssue returns one issue of one tracker project by its per-project number — title, description, status, priority, assignee, labels, kind, source and its git bindings.
[**CloudPatchV1TrackerProjectsKey**](TrackerAPI.md#CloudPatchV1TrackerProjectsKey) | **Patch** /v1/tracker/projects/{key} | UpdateProject renames a tracker project or rewrites its description, and returns the updated project.
[**CloudPatchV1TrackerProjectsKeyIssuesNum**](TrackerAPI.md#CloudPatchV1TrackerProjectsKeyIssuesNum) | **Patch** /v1/tracker/projects/{key}/issues/{num} | UpdateIssue edits one issue in place and returns it — retitle it, rewrite its body, move it between board columns, reprioritize, reassign, or replace its labels.
[**CloudPostV1TrackerProjects**](TrackerAPI.md#CloudPostV1TrackerProjects) | **Post** /v1/tracker/projects | 
[**CloudPostV1TrackerProjectsByKeyIssues**](TrackerAPI.md#CloudPostV1TrackerProjectsByKeyIssues) | **Post** /v1/tracker/projects/{key}/issues | 



## CloudDeleteV1TrackerProjectsKey

> CloudDeleteV1TrackerProjectsKey(ctx, key).Execute()

DeleteProject removes one tracker project of the caller's org AND every issue filed under it, and answers 204 with no body.



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
	key := "key_example" // string | Key is the project's org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\"ENG\", \"OPS2\"). Matched case-insensitively.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackerAPI.CloudDeleteV1TrackerProjectsKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudDeleteV1TrackerProjectsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project&#39;s org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\&quot;ENG\&quot;, \&quot;OPS2\&quot;). Matched case-insensitively. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TrackerProjectsKeyRequest struct via the builder pattern


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


## CloudDeleteV1TrackerProjectsKeyIssuesNum

> CloudDeleteV1TrackerProjectsKeyIssuesNum(ctx, key, num).Execute()

DeleteIssue removes one issue from a tracker project and answers 204 with no body.



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
	key := "key_example" // string | Key is the issue's project, from the path.
	num := int32(56) // int32 | Num is the issue's number within that project — the digits of KEY-14. Positive; anything else is refused with 400.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackerAPI.CloudDeleteV1TrackerProjectsKeyIssuesNum(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudDeleteV1TrackerProjectsKeyIssuesNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the issue&#39;s project, from the path. | 
**num** | **int32** | Num is the issue&#39;s number within that project — the digits of KEY-14. Positive; anything else is refused with 400. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1TrackerProjectsKeyIssuesNumRequest struct via the builder pattern


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


## CloudGetV1TrackerProjects

> []CloudTrackerProject CloudGetV1TrackerProjects(ctx).Execute()

ListProjects returns every tracker project in the caller's org, newest first.



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
	resp, r, err := apiClient.TrackerAPI.CloudGetV1TrackerProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudGetV1TrackerProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrackerProjects`: []CloudTrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudGetV1TrackerProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrackerProjectsRequest struct via the builder pattern


### Return type

[**[]CloudTrackerProject**](CloudTrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrackerProjectsKey

> CloudTrackerProject CloudGetV1TrackerProjectsKey(ctx, key).Execute()

GetProject returns one tracker project of the caller's org by its key — its name, description and timestamps.



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
	key := "key_example" // string | Key is the project's org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\"ENG\", \"OPS2\"). Matched case-insensitively.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerAPI.CloudGetV1TrackerProjectsKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudGetV1TrackerProjectsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrackerProjectsKey`: CloudTrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudGetV1TrackerProjectsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project&#39;s org-unique handle: 2-8 uppercase alphanumerics starting with a letter (\&quot;ENG\&quot;, \&quot;OPS2\&quot;). Matched case-insensitively. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrackerProjectsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudTrackerProject**](CloudTrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrackerProjectsKeyIssues

> []CloudIssueView CloudGetV1TrackerProjectsKeyIssues(ctx, key).Status(status).Kind(kind).Repo(repo).Source(source).Execute()

ListIssues returns the issues of one tracker project, optionally filtered by status, kind, repo and source.



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
	key := "ENG" // string | Key is the project whose issues to list, from the path.
	status := "status_example" // string | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. (optional)
	kind := "pr" // string | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. (optional)
	repo := "hanzoai/cloud" // string | Repo keeps only issues bound to that git repository. (optional)
	source := "source_example" // string | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerAPI.CloudGetV1TrackerProjectsKeyIssues(context.Background(), key).Status(status).Kind(kind).Repo(repo).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudGetV1TrackerProjectsKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrackerProjectsKeyIssues`: []CloudIssueView
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudGetV1TrackerProjectsKeyIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project whose issues to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrackerProjectsKeyIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Status keeps only issues in that board column: backlog, todo, in_progress, done or canceled. An unknown value is refused with 400. | 
 **kind** | **string** | Kind keeps only work items of that shape: issue, pr or epic. An unknown value is refused with 400. | 
 **repo** | **string** | Repo keeps only issues bound to that git repository. | 
 **source** | **string** | Source keeps only issues opened from that surface: team, git, crm, helpdesk, cms or agent. An unknown value is refused with 400. | 

### Return type

[**[]CloudIssueView**](CloudIssueView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1TrackerProjectsKeyIssuesNum

> CloudIssueView CloudGetV1TrackerProjectsKeyIssuesNum(ctx, key, num).Execute()

GetIssue returns one issue of one tracker project by its per-project number — title, description, status, priority, assignee, labels, kind, source and its git bindings.



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
	key := "key_example" // string | Key is the issue's project, from the path.
	num := int32(56) // int32 | Num is the issue's number within that project — the digits of KEY-14. Positive; anything else is refused with 400.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerAPI.CloudGetV1TrackerProjectsKeyIssuesNum(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudGetV1TrackerProjectsKeyIssuesNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1TrackerProjectsKeyIssuesNum`: CloudIssueView
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudGetV1TrackerProjectsKeyIssuesNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the issue&#39;s project, from the path. | 
**num** | **int32** | Num is the issue&#39;s number within that project — the digits of KEY-14. Positive; anything else is refused with 400. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1TrackerProjectsKeyIssuesNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudIssueView**](CloudIssueView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1TrackerProjectsKey

> CloudTrackerProject CloudPatchV1TrackerProjectsKey(ctx, key).CloudProjectPatch(cloudProjectPatch).Execute()

UpdateProject renames a tracker project or rewrites its description, and returns the updated project.



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
	key := "key_example" // string | Key is the project to update, from the path.
	cloudProjectPatch := *openapiclient.NewCloudProjectPatch() // CloudProjectPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerAPI.CloudPatchV1TrackerProjectsKey(context.Background(), key).CloudProjectPatch(cloudProjectPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudPatchV1TrackerProjectsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1TrackerProjectsKey`: CloudTrackerProject
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudPatchV1TrackerProjectsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the project to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1TrackerProjectsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudProjectPatch** | [**CloudProjectPatch**](CloudProjectPatch.md) |  | 

### Return type

[**CloudTrackerProject**](CloudTrackerProject.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1TrackerProjectsKeyIssuesNum

> CloudIssueView CloudPatchV1TrackerProjectsKeyIssuesNum(ctx, key, num).CloudIssuePatch(cloudIssuePatch).Execute()

UpdateIssue edits one issue in place and returns it — retitle it, rewrite its body, move it between board columns, reprioritize, reassign, or replace its labels.



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
	key := "ENG" // string | Key is the issue's project, from the path.
	num := int32(14) // int32 | Num is the issue's number within that project, from the path.
	cloudIssuePatch := *openapiclient.NewCloudIssuePatch() // CloudIssuePatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerAPI.CloudPatchV1TrackerProjectsKeyIssuesNum(context.Background(), key, num).CloudIssuePatch(cloudIssuePatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudPatchV1TrackerProjectsKeyIssuesNum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1TrackerProjectsKeyIssuesNum`: CloudIssueView
	fmt.Fprintf(os.Stdout, "Response from `TrackerAPI.CloudPatchV1TrackerProjectsKeyIssuesNum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the issue&#39;s project, from the path. | 
**num** | **int32** | Num is the issue&#39;s number within that project, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1TrackerProjectsKeyIssuesNumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cloudIssuePatch** | [**CloudIssuePatch**](CloudIssuePatch.md) |  | 

### Return type

[**CloudIssueView**](CloudIssueView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1TrackerProjects

> CloudPostV1TrackerProjects(ctx).Execute()



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
	r, err := apiClient.TrackerAPI.CloudPostV1TrackerProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudPostV1TrackerProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TrackerProjectsRequest struct via the builder pattern


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


## CloudPostV1TrackerProjectsByKeyIssues

> CloudPostV1TrackerProjectsByKeyIssues(ctx, key).Execute()



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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackerAPI.CloudPostV1TrackerProjectsByKeyIssues(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerAPI.CloudPostV1TrackerProjectsByKeyIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1TrackerProjectsByKeyIssuesRequest struct via the builder pattern


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

