# \PlatformRollbackAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformRollbackRollback**](PlatformRollbackAPI.md#PlatformRollbackRollback) | **Post** /v1/platform/rollback/rollback | Rollback to a previous deployment



## PlatformRollbackRollback

> PlatformTRPCResult PlatformRollbackRollback(ctx).PlatformDeploymentKillProcessRequest(platformDeploymentKillProcessRequest).Execute()

Rollback to a previous deployment

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
	platformDeploymentKillProcessRequest := *openapiclient.NewPlatformDeploymentKillProcessRequest() // PlatformDeploymentKillProcessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformRollbackAPI.PlatformRollbackRollback(context.Background()).PlatformDeploymentKillProcessRequest(platformDeploymentKillProcessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformRollbackAPI.PlatformRollbackRollback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformRollbackRollback`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformRollbackAPI.PlatformRollbackRollback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformRollbackRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDeploymentKillProcessRequest** | [**PlatformDeploymentKillProcessRequest**](PlatformDeploymentKillProcessRequest.md) |  | 

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

