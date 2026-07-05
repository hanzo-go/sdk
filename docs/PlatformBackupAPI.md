# \PlatformBackupAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformBackupCreate**](PlatformBackupAPI.md#PlatformBackupCreate) | **Post** /v1/platform/backup/create | Create a backup schedule



## PlatformBackupCreate

> PlatformTRPCResult PlatformBackupCreate(ctx).PlatformBackupCreateRequest(platformBackupCreateRequest).Execute()

Create a backup schedule

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
	platformBackupCreateRequest := *openapiclient.NewPlatformBackupCreateRequest() // PlatformBackupCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformBackupAPI.PlatformBackupCreate(context.Background()).PlatformBackupCreateRequest(platformBackupCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformBackupAPI.PlatformBackupCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformBackupCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformBackupAPI.PlatformBackupCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformBackupCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformBackupCreateRequest** | [**PlatformBackupCreateRequest**](PlatformBackupCreateRequest.md) |  | 

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

