# \SearchLogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGetStderrLogs**](SearchLogsAPI.md#SearchGetStderrLogs) | **Get** /v1/search/logs/stderr | Get stderr log configuration
[**SearchResetStderrLogs**](SearchLogsAPI.md#SearchResetStderrLogs) | **Delete** /v1/search/logs/stderr | Reset stderr log level to default
[**SearchStreamLogs**](SearchLogsAPI.md#SearchStreamLogs) | **Post** /v1/search/logs/stream | Stream logs
[**SearchUpdateStderrLogs**](SearchLogsAPI.md#SearchUpdateStderrLogs) | **Put** /v1/search/logs/stderr | Update stderr log level



## SearchGetStderrLogs

> SearchGetStderrLogs200Response SearchGetStderrLogs(ctx).Execute()

Get stderr log configuration

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
	resp, r, err := apiClient.SearchLogsAPI.SearchGetStderrLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchLogsAPI.SearchGetStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchLogsAPI.SearchGetStderrLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetStderrLogsRequest struct via the builder pattern


### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchResetStderrLogs

> SearchGetStderrLogs200Response SearchResetStderrLogs(ctx).Execute()

Reset stderr log level to default

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
	resp, r, err := apiClient.SearchLogsAPI.SearchResetStderrLogs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchLogsAPI.SearchResetStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchResetStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchLogsAPI.SearchResetStderrLogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchResetStderrLogsRequest struct via the builder pattern


### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStreamLogs

> string SearchStreamLogs(ctx).SearchStreamLogsRequest(searchStreamLogsRequest).Execute()

Stream logs

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
	searchStreamLogsRequest := *openapiclient.NewSearchStreamLogsRequest("Target_example", "Mode_example") // SearchStreamLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchLogsAPI.SearchStreamLogs(context.Background()).SearchStreamLogsRequest(searchStreamLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchLogsAPI.SearchStreamLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchStreamLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `SearchLogsAPI.SearchStreamLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStreamLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchStreamLogsRequest** | [**SearchStreamLogsRequest**](SearchStreamLogsRequest.md) |  | 

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateStderrLogs

> SearchGetStderrLogs200Response SearchUpdateStderrLogs(ctx).SearchUpdateStderrLogsRequest(searchUpdateStderrLogsRequest).Execute()

Update stderr log level

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
	searchUpdateStderrLogsRequest := *openapiclient.NewSearchUpdateStderrLogsRequest("Target_example") // SearchUpdateStderrLogsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchLogsAPI.SearchUpdateStderrLogs(context.Background()).SearchUpdateStderrLogsRequest(searchUpdateStderrLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchLogsAPI.SearchUpdateStderrLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateStderrLogs`: SearchGetStderrLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchLogsAPI.SearchUpdateStderrLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateStderrLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchUpdateStderrLogsRequest** | [**SearchUpdateStderrLogsRequest**](SearchUpdateStderrLogsRequest.md) |  | 

### Return type

[**SearchGetStderrLogs200Response**](SearchGetStderrLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

