# \PlatformScheduleAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformScheduleCreate**](PlatformScheduleAPI.md#PlatformScheduleCreate) | **Post** /v1/platform/schedule/create | Create a scheduled task
[**PlatformScheduleList**](PlatformScheduleAPI.md#PlatformScheduleList) | **Get** /v1/platform/schedule/list | List scheduled tasks



## PlatformScheduleCreate

> PlatformTRPCResult PlatformScheduleCreate(ctx).PlatformScheduleCreateRequest(platformScheduleCreateRequest).Execute()

Create a scheduled task

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
	platformScheduleCreateRequest := *openapiclient.NewPlatformScheduleCreateRequest() // PlatformScheduleCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformScheduleAPI.PlatformScheduleCreate(context.Background()).PlatformScheduleCreateRequest(platformScheduleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformScheduleAPI.PlatformScheduleCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformScheduleCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformScheduleAPI.PlatformScheduleCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformScheduleCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformScheduleCreateRequest** | [**PlatformScheduleCreateRequest**](PlatformScheduleCreateRequest.md) |  | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformScheduleList

> PlatformTRPCResult PlatformScheduleList(ctx).Input(input).Execute()

List scheduled tasks

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformScheduleAPI.PlatformScheduleList(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformScheduleAPI.PlatformScheduleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformScheduleList`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformScheduleAPI.PlatformScheduleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformScheduleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

