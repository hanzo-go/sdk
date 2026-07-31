# \SkillsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BotCreateSkillComment**](SkillsAPI.md#BotCreateSkillComment) | **Post** /v1/bot/skills/{slug}/comments | Add a comment to a skill
[**BotDeleteSkill**](SkillsAPI.md#BotDeleteSkill) | **Delete** /v1/bot/skills/{slug} | Soft-delete a skill (owner or admin only)
[**BotDeleteSkillComment**](SkillsAPI.md#BotDeleteSkillComment) | **Delete** /v1/bot/skills/{slug}/comments/{commentId} | Delete a comment (author or admin only)
[**BotGetSkill**](SkillsAPI.md#BotGetSkill) | **Get** /v1/bot/skills/{slug} | Get skill details by slug
[**BotGetSkillStarStatus**](SkillsAPI.md#BotGetSkillStarStatus) | **Get** /v1/bot/skills/{slug}/stars/me | Check if current user has starred this skill
[**BotGetSkillVersionFiles**](SkillsAPI.md#BotGetSkillVersionFiles) | **Get** /v1/bot/skills/{slug}/versions/{version}/files | Get file listing for a specific version
[**BotListSkillComments**](SkillsAPI.md#BotListSkillComments) | **Get** /v1/bot/skills/{slug}/comments | List comments on a skill
[**BotListSkillVersions**](SkillsAPI.md#BotListSkillVersions) | **Get** /v1/bot/skills/{slug}/versions | List versions of a skill
[**BotListSkills**](SkillsAPI.md#BotListSkills) | **Get** /v1/bot/skills | List published skills (paginated)
[**BotPublishSkillVersion**](SkillsAPI.md#BotPublishSkillVersion) | **Post** /v1/bot/skills/{slug}/publish | Publish a new version of a skill (creates skill if new)
[**BotToggleSkillStar**](SkillsAPI.md#BotToggleSkillStar) | **Post** /v1/bot/skills/{slug}/stars | Star or unstar a skill (toggle)
[**BotUndeleteSkill**](SkillsAPI.md#BotUndeleteSkill) | **Post** /v1/bot/skills/{slug}/undelete | Restore a soft-deleted skill
[**CloudDeleteV1SkillsId**](SkillsAPI.md#CloudDeleteV1SkillsId) | **Delete** /v1/skills/{id} | DeleteSkill removes one of the caller org&#39;s authored skills.
[**CloudGetV1Skills**](SkillsAPI.md#CloudGetV1Skills) | **Get** /v1/skills | ListSkills lists the skills the caller&#39;s org can reach — the brand&#39;s embedded catalogue plus the org&#39;s own authored ones — with each one&#39;s activation flag.
[**CloudGetV1SkillsAuthored**](SkillsAPI.md#CloudGetV1SkillsAuthored) | **Get** /v1/skills/authored | ListAuthoredSkills lists the caller org&#39;s OWN skills with their SKILL.md bodies.
[**CloudPostV1Skills**](SkillsAPI.md#CloudPostV1Skills) | **Post** /v1/skills | PutSkill adds or revises one of the caller org&#39;s own skills, and answers 201 with the stored record.



## BotCreateSkillComment

> BotComment BotCreateSkillComment(ctx, slug).BotCreateSkillCommentRequest(botCreateSkillCommentRequest).Execute()

Add a comment to a skill

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
	botCreateSkillCommentRequest := *openapiclient.NewBotCreateSkillCommentRequest("Body_example") // BotCreateSkillCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotCreateSkillComment(context.Background(), slug).BotCreateSkillCommentRequest(botCreateSkillCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotCreateSkillComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotCreateSkillComment`: BotComment
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotCreateSkillComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotCreateSkillCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **botCreateSkillCommentRequest** | [**BotCreateSkillCommentRequest**](BotCreateSkillCommentRequest.md) |  | 

### Return type

[**BotComment**](BotComment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotDeleteSkill

> AnalyticsHeartbeat200Response BotDeleteSkill(ctx, slug).Execute()

Soft-delete a skill (owner or admin only)

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
	resp, r, err := apiClient.SkillsAPI.BotDeleteSkill(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotDeleteSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotDeleteSkill`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotDeleteSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotDeleteSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotDeleteSkillComment

> AnalyticsHeartbeat200Response BotDeleteSkillComment(ctx, slug, commentId).Execute()

Delete a comment (author or admin only)

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
	commentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotDeleteSkillComment(context.Background(), slug, commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotDeleteSkillComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotDeleteSkillComment`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotDeleteSkillComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**commentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotDeleteSkillCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGetSkill

> BotSkill BotGetSkill(ctx, slug).Execute()

Get skill details by slug

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
	resp, r, err := apiClient.SkillsAPI.BotGetSkill(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotGetSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetSkill`: BotSkill
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotGetSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotSkill**](BotSkill.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGetSkillStarStatus

> BotToggleSkillStar200Response BotGetSkillStarStatus(ctx, slug).Execute()

Check if current user has starred this skill

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
	resp, r, err := apiClient.SkillsAPI.BotGetSkillStarStatus(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotGetSkillStarStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetSkillStarStatus`: BotToggleSkillStar200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotGetSkillStarStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetSkillStarStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotToggleSkillStar200Response**](BotToggleSkillStar200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotGetSkillVersionFiles

> BotGetSkillVersionFiles200Response BotGetSkillVersionFiles(ctx, slug, version).Execute()

Get file listing for a specific version

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
	version := "version_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotGetSkillVersionFiles(context.Background(), slug, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotGetSkillVersionFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotGetSkillVersionFiles`: BotGetSkillVersionFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotGetSkillVersionFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotGetSkillVersionFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BotGetSkillVersionFiles200Response**](BotGetSkillVersionFiles200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListSkillComments

> BotListSkillComments200Response BotListSkillComments(ctx, slug).Execute()

List comments on a skill

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
	resp, r, err := apiClient.SkillsAPI.BotListSkillComments(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotListSkillComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListSkillComments`: BotListSkillComments200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotListSkillComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListSkillCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotListSkillComments200Response**](BotListSkillComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListSkillVersions

> BotListSkillVersions200Response BotListSkillVersions(ctx, slug).Limit(limit).Execute()

List versions of a skill

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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotListSkillVersions(context.Background(), slug).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotListSkillVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListSkillVersions`: BotListSkillVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotListSkillVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotListSkillVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]

### Return type

[**BotListSkillVersions200Response**](BotListSkillVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotListSkills

> BotListSkills200Response BotListSkills(ctx).Sort(sort).Limit(limit).Cursor(cursor).Batch(batch).Execute()

List published skills (paginated)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	sort := "sort_example" // string |  (optional) (default to "updated")
	limit := int32(56) // int32 |  (optional) (default to 50)
	cursor := time.Now() // time.Time | Cursor for pagination (updatedAt ISO timestamp) (optional)
	batch := "batch_example" // string | Filter by batch grouping key (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotListSkills(context.Background()).Sort(sort).Limit(limit).Cursor(cursor).Batch(batch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotListSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotListSkills`: BotListSkills200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotListSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBotListSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** |  | [default to &quot;updated&quot;]
 **limit** | **int32** |  | [default to 50]
 **cursor** | **time.Time** | Cursor for pagination (updatedAt ISO timestamp) | 
 **batch** | **string** | Filter by batch grouping key | 

### Return type

[**BotListSkills200Response**](BotListSkills200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotPublishSkillVersion

> BotPublishSkillVersion200Response BotPublishSkillVersion(ctx, slug).BotPublishSkillVersionRequest(botPublishSkillVersionRequest).Execute()

Publish a new version of a skill (creates skill if new)

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
	botPublishSkillVersionRequest := *openapiclient.NewBotPublishSkillVersionRequest("DisplayName_example", "Version_example", "Changelog_example", []openapiclient.BotPublishSkillVersionRequestFilesInner{*openapiclient.NewBotPublishSkillVersionRequestFilesInner("Path_example", int32(123), "StorageKey_example", "Sha256_example")}) // BotPublishSkillVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.BotPublishSkillVersion(context.Background(), slug).BotPublishSkillVersionRequest(botPublishSkillVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotPublishSkillVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotPublishSkillVersion`: BotPublishSkillVersion200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotPublishSkillVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotPublishSkillVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **botPublishSkillVersionRequest** | [**BotPublishSkillVersionRequest**](BotPublishSkillVersionRequest.md) |  | 

### Return type

[**BotPublishSkillVersion200Response**](BotPublishSkillVersion200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotToggleSkillStar

> BotToggleSkillStar200Response BotToggleSkillStar(ctx, slug).Execute()

Star or unstar a skill (toggle)

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
	resp, r, err := apiClient.SkillsAPI.BotToggleSkillStar(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotToggleSkillStar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotToggleSkillStar`: BotToggleSkillStar200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotToggleSkillStar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotToggleSkillStarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotToggleSkillStar200Response**](BotToggleSkillStar200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BotUndeleteSkill

> AnalyticsHeartbeat200Response BotUndeleteSkill(ctx, slug).Execute()

Restore a soft-deleted skill

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
	resp, r, err := apiClient.SkillsAPI.BotUndeleteSkill(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.BotUndeleteSkill``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BotUndeleteSkill`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.BotUndeleteSkill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBotUndeleteSkillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1SkillsId

> CloudSkillDeleted CloudDeleteV1SkillsId(ctx, id).Execute()

DeleteSkill removes one of the caller org's authored skills.



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
	id := "id_example" // string | ID is the skill to remove, from the path. It is the skill's name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.CloudDeleteV1SkillsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.CloudDeleteV1SkillsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1SkillsId`: CloudSkillDeleted
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.CloudDeleteV1SkillsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the skill to remove, from the path. It is the skill&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1SkillsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSkillDeleted**](CloudSkillDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Skills

> CloudSourceToolList CloudGetV1Skills(ctx).Activated(activated).Execute()

ListSkills lists the skills the caller's org can reach — the brand's embedded catalogue plus the org's own authored ones — with each one's activation flag.



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
	activated := "activated_example" // string | Activated keeps only the tools activated for the caller's org and project, and only when it is exactly the string \"true\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.CloudGetV1Skills(context.Background()).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.CloudGetV1Skills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Skills`: CloudSourceToolList
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.CloudGetV1Skills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**CloudSourceToolList**](CloudSourceToolList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1SkillsAuthored

> CloudAuthoredSkillList CloudGetV1SkillsAuthored(ctx).Execute()

ListAuthoredSkills lists the caller org's OWN skills with their SKILL.md bodies.



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
	resp, r, err := apiClient.SkillsAPI.CloudGetV1SkillsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.CloudGetV1SkillsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SkillsAuthored`: CloudAuthoredSkillList
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.CloudGetV1SkillsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SkillsAuthoredRequest struct via the builder pattern


### Return type

[**CloudAuthoredSkillList**](CloudAuthoredSkillList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Skills

> CloudSkillWritten CloudPostV1Skills(ctx).CloudSkillIn(cloudSkillIn).Execute()

PutSkill adds or revises one of the caller org's own skills, and answers 201 with the stored record.



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
	cloudSkillIn := *openapiclient.NewCloudSkillIn() // CloudSkillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.CloudPostV1Skills(context.Background()).CloudSkillIn(cloudSkillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.CloudPostV1Skills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Skills`: CloudSkillWritten
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.CloudPostV1Skills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1SkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSkillIn** | [**CloudSkillIn**](CloudSkillIn.md) |  | 

### Return type

[**CloudSkillWritten**](CloudSkillWritten.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

