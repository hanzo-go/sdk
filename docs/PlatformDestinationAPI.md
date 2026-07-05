# \PlatformDestinationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDestinationAll**](PlatformDestinationAPI.md#PlatformDestinationAll) | **Get** /v1/platform/destination/all | List backup destinations



## PlatformDestinationAll

> PlatformTRPCResult PlatformDestinationAll(ctx).Execute()

List backup destinations

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
	resp, r, err := apiClient.PlatformDestinationAPI.PlatformDestinationAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDestinationAPI.PlatformDestinationAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDestinationAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDestinationAPI.PlatformDestinationAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDestinationAllRequest struct via the builder pattern


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

