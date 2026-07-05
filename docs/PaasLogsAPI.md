# \PaasLogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasGetContainerLogs**](PaasLogsAPI.md#PaasGetContainerLogs) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env/{envId}/containers/{containerId}/logs | Get container logs



## PaasGetContainerLogs

> string PaasGetContainerLogs(ctx, orgId, projectId, envId, containerId).Tail(tail).Follow(follow).Execute()

Get container logs

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 
	containerId := "containerId_example" // string | 
	tail := int32(56) // int32 |  (optional) (default to 100)
	follow := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasLogsAPI.PaasGetContainerLogs(context.Background(), orgId, projectId, envId, containerId).Tail(tail).Follow(follow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasLogsAPI.PaasGetContainerLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetContainerLogs`: string
	fmt.Fprintf(os.Stdout, "Response from `PaasLogsAPI.PaasGetContainerLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 
**containerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetContainerLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tail** | **int32** |  | [default to 100]
 **follow** | **bool** |  | [default to false]

### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

