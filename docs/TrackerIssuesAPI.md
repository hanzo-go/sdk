# \TrackerIssuesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackerCreateIssue**](TrackerIssuesAPI.md#TrackerCreateIssue) | **Post** /v1/tracker/projects/{key}/issues | Create an issue
[**TrackerDeleteIssue**](TrackerIssuesAPI.md#TrackerDeleteIssue) | **Delete** /v1/tracker/projects/{key}/issues/{num} | Delete an issue
[**TrackerGetIssue**](TrackerIssuesAPI.md#TrackerGetIssue) | **Get** /v1/tracker/projects/{key}/issues/{num} | Get an issue
[**TrackerListIssues**](TrackerIssuesAPI.md#TrackerListIssues) | **Get** /v1/tracker/projects/{key}/issues | List issues (board/list)
[**TrackerUpdateIssue**](TrackerIssuesAPI.md#TrackerUpdateIssue) | **Patch** /v1/tracker/projects/{key}/issues/{num} | Update an issue



## TrackerCreateIssue

> TrackerIssue TrackerCreateIssue(ctx, key).TrackerCreateIssueRequest(trackerCreateIssueRequest).Execute()

Create an issue

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
	key := "key_example" // string | Project key
	trackerCreateIssueRequest := *openapiclient.NewTrackerCreateIssueRequest("Title_example") // TrackerCreateIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerIssuesAPI.TrackerCreateIssue(context.Background(), key).TrackerCreateIssueRequest(trackerCreateIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerIssuesAPI.TrackerCreateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerCreateIssue`: TrackerIssue
	fmt.Fprintf(os.Stdout, "Response from `TrackerIssuesAPI.TrackerCreateIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerCreateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **trackerCreateIssueRequest** | [**TrackerCreateIssueRequest**](TrackerCreateIssueRequest.md) |  | 

### Return type

[**TrackerIssue**](TrackerIssue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerDeleteIssue

> TrackerDeleteIssue(ctx, key, num).Execute()

Delete an issue

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
	key := "key_example" // string | Project key
	num := int32(56) // int32 | Per-project issue number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TrackerIssuesAPI.TrackerDeleteIssue(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerIssuesAPI.TrackerDeleteIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key | 
**num** | **int32** | Per-project issue number | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerDeleteIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerGetIssue

> TrackerIssue TrackerGetIssue(ctx, key, num).Execute()

Get an issue

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
	key := "key_example" // string | Project key
	num := int32(56) // int32 | Per-project issue number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerIssuesAPI.TrackerGetIssue(context.Background(), key, num).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerIssuesAPI.TrackerGetIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerGetIssue`: TrackerIssue
	fmt.Fprintf(os.Stdout, "Response from `TrackerIssuesAPI.TrackerGetIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key | 
**num** | **int32** | Per-project issue number | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TrackerIssue**](TrackerIssue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerListIssues

> []TrackerIssue TrackerListIssues(ctx, key).Status(status).Execute()

List issues (board/list)

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
	key := "key_example" // string | Project key
	status := openapiclient.tracker_IssueStatus("backlog") // TrackerIssueStatus | Filter to one board column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerIssuesAPI.TrackerListIssues(context.Background(), key).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerIssuesAPI.TrackerListIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerListIssues`: []TrackerIssue
	fmt.Fprintf(os.Stdout, "Response from `TrackerIssuesAPI.TrackerListIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerListIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**TrackerIssueStatus**](TrackerIssueStatus.md) | Filter to one board column | 

### Return type

[**[]TrackerIssue**](TrackerIssue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackerUpdateIssue

> TrackerIssue TrackerUpdateIssue(ctx, key, num).TrackerUpdateIssueRequest(trackerUpdateIssueRequest).Execute()

Update an issue

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
	key := "key_example" // string | Project key
	num := int32(56) // int32 | Per-project issue number
	trackerUpdateIssueRequest := *openapiclient.NewTrackerUpdateIssueRequest() // TrackerUpdateIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrackerIssuesAPI.TrackerUpdateIssue(context.Background(), key, num).TrackerUpdateIssueRequest(trackerUpdateIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrackerIssuesAPI.TrackerUpdateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackerUpdateIssue`: TrackerIssue
	fmt.Fprintf(os.Stdout, "Response from `TrackerIssuesAPI.TrackerUpdateIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Project key | 
**num** | **int32** | Per-project issue number | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrackerUpdateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trackerUpdateIssueRequest** | [**TrackerUpdateIssueRequest**](TrackerUpdateIssueRequest.md) |  | 

### Return type

[**TrackerIssue**](TrackerIssue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

