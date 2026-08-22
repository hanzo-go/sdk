# \SocialAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSocialAccountsById**](SocialAPI.md#DeleteSocialAccountsById) | **Delete** /v1/social/accounts/{id} | Removes one connected account from the org and answers 204 with no body; an id that is not there is 404.
[**DeleteSocialPostsById**](SocialAPI.md#DeleteSocialPostsById) | **Delete** /v1/social/posts/{id} | Removes one post from the org and answers 204 with no body; an id that is not there is 404.
[**GetSocialAccounts**](SocialAPI.md#GetSocialAccounts) | **Get** /v1/social/accounts | Returns the org&#39;s connected accounts — each one&#39;s id, network, handle, status and timestamps, most-recently-updated first.
[**GetSocialAccountsById**](SocialAPI.md#GetSocialAccountsById) | **Get** /v1/social/accounts/{id} | Returns one of the org&#39;s connected accounts by id — its network, handle, status and timestamps — or 404.
[**GetSocialPosts**](SocialAPI.md#GetSocialPosts) | **Get** /v1/social/posts | Returns the org&#39;s posts — content, channel, status, scheduled time, media and timestamps — most-recently-updated first.
[**GetSocialPostsById**](SocialAPI.md#GetSocialPostsById) | **Get** /v1/social/posts/{id} | Returns one of the org&#39;s posts by id, with its current status, scheduled time, media and — once it has published — the account and external id it published under.
[**GetSocialProviders**](SocialAPI.md#GetSocialProviders) | **Get** /v1/social/providers | Reports each supported network&#39;s publish-readiness: whether this deployment holds the OAuth application credentials for it and, when it does not, exactly which environment variables are missing.
[**GetSocialSummary**](SocialAPI.md#GetSocialSummary) | **Get** /v1/social/summary | Returns four counts for the caller&#39;s org: total posts, how many are scheduled, how many have published, and how many accounts are connected.
[**PostSocialAccounts**](SocialAPI.md#PostSocialAccounts) | **Post** /v1/social/accounts | Records a social account for the org and answers 201 with the stored row, including the generated id later calls address it by.
[**PostSocialPosts**](SocialAPI.md#PostSocialPosts) | **Post** /v1/social/posts | Stores a post for the org and answers 201 with the stored row.
[**PostSocialPostsByIdPublish**](SocialAPI.md#PostSocialPostsByIdPublish) | **Post** /v1/social/posts/{id}/publish | Publishes the post immediately to the connected accounts on its channel and answers with the updated row, carrying the account and external id it published under.
[**PutSocialAccountsById**](SocialAPI.md#PutSocialAccountsById) | **Put** /v1/social/accounts/{id} | Replaces the account&#39;s network, handle and status with what the body carries, and answers with the stored row.
[**PutSocialPostsById**](SocialAPI.md#PutSocialPostsById) | **Put** /v1/social/posts/{id} | Replaces the post&#39;s content, channel, status, scheduled time and media with what the body carries, and answers with the stored row.



## DeleteSocialAccountsById

> DeleteSocialAccountsById(ctx, id).Execute()

Removes one connected account from the org and answers 204 with no body; an id that is not there is 404.



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
	id := "id_example" // string | ID is the account or post to act on, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SocialAPI.DeleteSocialAccountsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.DeleteSocialAccountsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the account or post to act on, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialAccountsByIdRequest struct via the builder pattern


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


## DeleteSocialPostsById

> DeleteSocialPostsById(ctx, id).Execute()

Removes one post from the org and answers 204 with no body; an id that is not there is 404.



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
	id := "id_example" // string | ID is the account or post to act on, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SocialAPI.DeleteSocialPostsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.DeleteSocialPostsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the account or post to act on, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSocialPostsByIdRequest struct via the builder pattern


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


## GetSocialAccounts

> SocialAccounts GetSocialAccounts(ctx).Provider(provider).Limit(limit).Execute()

Returns the org's connected accounts — each one's id, network, handle, status and timestamps, most-recently-updated first.



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
	provider := "provider_example" // string | Provider keeps only accounts on one network — x, facebook, instagram, linkedin, tiktok, youtube or threads. Omit it for every network. It is lower-cased and trimmed before it is matched, and a value that names no network simply matches nothing rather than being refused. (optional)
	limit := "limit_example" // string | Limit bounds the page, defaulting to 200 and capped at 1000. It is a string rather than an integer on purpose: the route parses it with a leading trim and falls back to the default on anything it cannot read, so `?limit=%2050` is a page of fifty today. An integer field would refuse the space and read an unparseable value as zero, which is a different page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialAccounts(context.Background()).Provider(provider).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialAccounts`: SocialAccounts
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider keeps only accounts on one network — x, facebook, instagram, linkedin, tiktok, youtube or threads. Omit it for every network. It is lower-cased and trimmed before it is matched, and a value that names no network simply matches nothing rather than being refused. | 
 **limit** | **string** | Limit bounds the page, defaulting to 200 and capped at 1000. It is a string rather than an integer on purpose: the route parses it with a leading trim and falls back to the default on anything it cannot read, so &#x60;?limit&#x3D;%2050&#x60; is a page of fifty today. An integer field would refuse the space and read an unparseable value as zero, which is a different page. | 

### Return type

[**SocialAccounts**](SocialAccounts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialAccountsById

> SocialAccount GetSocialAccountsById(ctx, id).Execute()

Returns one of the org's connected accounts by id — its network, handle, status and timestamps — or 404.



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
	id := "id_example" // string | ID is the account or post to act on, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialAccountsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialAccountsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialAccountsById`: SocialAccount
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialAccountsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the account or post to act on, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialAccountsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SocialAccount**](SocialAccount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPosts

> SocialPosts GetSocialPosts(ctx).Status(status).Limit(limit).Execute()

Returns the org's posts — content, channel, status, scheduled time, media and timestamps — most-recently-updated first.



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
	status := "status_example" // string | Status keeps only posts in one state — draft, scheduled, published or failed. Omit it for every state. The transient publishing claim is not a user-visible state and matching it is not useful. (optional)
	limit := "limit_example" // string | Limit bounds the page, defaulting to 200 and capped at 1000. A string for the same reason accountFilter.Limit is. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialPosts(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPosts`: SocialPosts
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only posts in one state — draft, scheduled, published or failed. Omit it for every state. The transient publishing claim is not a user-visible state and matching it is not useful. | 
 **limit** | **string** | Limit bounds the page, defaulting to 200 and capped at 1000. A string for the same reason accountFilter.Limit is. | 

### Return type

[**SocialPosts**](SocialPosts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialPostsById

> SocialPost GetSocialPostsById(ctx, id).Execute()

Returns one of the org's posts by id, with its current status, scheduled time, media and — once it has published — the account and external id it published under.



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
	id := "id_example" // string | ID is the account or post to act on, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.GetSocialPostsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialPostsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialPostsById`: SocialPost
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialPostsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the account or post to act on, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialPostsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SocialPost**](SocialPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialProviders

> SocialProviders GetSocialProviders(ctx).Execute()

Reports each supported network's publish-readiness: whether this deployment holds the OAuth application credentials for it and, when it does not, exactly which environment variables are missing.



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
	resp, r, err := apiClient.SocialAPI.GetSocialProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialProviders`: SocialProviders
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialProvidersRequest struct via the builder pattern


### Return type

[**SocialProviders**](SocialProviders.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSocialSummary

> SocialSummary GetSocialSummary(ctx).Execute()

Returns four counts for the caller's org: total posts, how many are scheduled, how many have published, and how many accounts are connected.



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
	resp, r, err := apiClient.SocialAPI.GetSocialSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.GetSocialSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSocialSummary`: SocialSummary
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.GetSocialSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSocialSummaryRequest struct via the builder pattern


### Return type

[**SocialSummary**](SocialSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSocialAccounts

> SocialAccount PostSocialAccounts(ctx).SocialAccountBody(socialAccountBody).Execute()

Records a social account for the org and answers 201 with the stored row, including the generated id later calls address it by.



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
	socialAccountBody := *openapiclient.NewSocialAccountBody() // SocialAccountBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.PostSocialAccounts(context.Background()).SocialAccountBody(socialAccountBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.PostSocialAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSocialAccounts`: SocialAccount
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.PostSocialAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSocialAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialAccountBody** | [**SocialAccountBody**](SocialAccountBody.md) |  | 

### Return type

[**SocialAccount**](SocialAccount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSocialPosts

> SocialPost PostSocialPosts(ctx).SocialPostBody(socialPostBody).Execute()

Stores a post for the org and answers 201 with the stored row.



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
	socialPostBody := *openapiclient.NewSocialPostBody() // SocialPostBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.PostSocialPosts(context.Background()).SocialPostBody(socialPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.PostSocialPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSocialPosts`: SocialPost
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.PostSocialPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSocialPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **socialPostBody** | [**SocialPostBody**](SocialPostBody.md) |  | 

### Return type

[**SocialPost**](SocialPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSocialPostsByIdPublish

> SocialPost PostSocialPostsByIdPublish(ctx, id).Execute()

Publishes the post immediately to the connected accounts on its channel and answers with the updated row, carrying the account and external id it published under.



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
	id := "id_example" // string | ID is the account or post to act on, taken from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.PostSocialPostsByIdPublish(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.PostSocialPostsByIdPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSocialPostsByIdPublish`: SocialPost
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.PostSocialPostsByIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the account or post to act on, taken from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSocialPostsByIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SocialPost**](SocialPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSocialAccountsById

> SocialAccount PutSocialAccountsById(ctx, id).SocialAccountWrite(socialAccountWrite).Execute()

Replaces the account's network, handle and status with what the body carries, and answers with the stored row.



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
	id := "id_example" // string | 
	socialAccountWrite := *openapiclient.NewSocialAccountWrite() // SocialAccountWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.PutSocialAccountsById(context.Background(), id).SocialAccountWrite(socialAccountWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.PutSocialAccountsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSocialAccountsById`: SocialAccount
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.PutSocialAccountsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSocialAccountsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialAccountWrite** | [**SocialAccountWrite**](SocialAccountWrite.md) |  | 

### Return type

[**SocialAccount**](SocialAccount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSocialPostsById

> SocialPost PutSocialPostsById(ctx, id).SocialPostWrite(socialPostWrite).Execute()

Replaces the post's content, channel, status, scheduled time and media with what the body carries, and answers with the stored row.



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
	id := "id_example" // string | 
	socialPostWrite := *openapiclient.NewSocialPostWrite() // SocialPostWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SocialAPI.PutSocialPostsById(context.Background(), id).SocialPostWrite(socialPostWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SocialAPI.PutSocialPostsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSocialPostsById`: SocialPost
	fmt.Fprintf(os.Stdout, "Response from `SocialAPI.PutSocialPostsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSocialPostsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **socialPostWrite** | [**SocialPostWrite**](SocialPostWrite.md) |  | 

### Return type

[**SocialPost**](SocialPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

