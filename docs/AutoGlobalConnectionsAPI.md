# \AutoGlobalConnectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoListGlobalConnections**](AutoGlobalConnectionsAPI.md#AutoListGlobalConnections) | **Get** /v1/auto/global-connections | List platform-wide connections (EE)



## AutoListGlobalConnections

> map[string]interface{} AutoListGlobalConnections(ctx).Execute()

List platform-wide connections (EE)

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
	resp, r, err := apiClient.AutoGlobalConnectionsAPI.AutoListGlobalConnections(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoGlobalConnectionsAPI.AutoListGlobalConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListGlobalConnections`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoGlobalConnectionsAPI.AutoListGlobalConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListGlobalConnectionsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

