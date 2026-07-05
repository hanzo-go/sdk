# \PlatformGitProviderAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformGitProviderGetAll**](PlatformGitProviderAPI.md#PlatformGitProviderGetAll) | **Get** /v1/platform/gitProvider/getAll | List all git providers



## PlatformGitProviderGetAll

> PlatformTRPCResult PlatformGitProviderGetAll(ctx).Execute()

List all git providers

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
	resp, r, err := apiClient.PlatformGitProviderAPI.PlatformGitProviderGetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformGitProviderAPI.PlatformGitProviderGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformGitProviderGetAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformGitProviderAPI.PlatformGitProviderGetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformGitProviderGetAllRequest struct via the builder pattern


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

