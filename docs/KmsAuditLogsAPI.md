# \KmsAuditLogsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsListAuditLogs**](KmsAuditLogsAPI.md#KmsListAuditLogs) | **Get** /v1/kms/events | List audit log events



## KmsListAuditLogs

> KmsListAuditLogs200Response KmsListAuditLogs(ctx).WorkspaceId(workspaceId).EventType(eventType).Actor(actor).Offset(offset).Limit(limit).Execute()

List audit log events

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	eventType := "eventType_example" // string |  (optional)
	actor := "actor_example" // string |  (optional)
	offset := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAuditLogsAPI.KmsListAuditLogs(context.Background()).WorkspaceId(workspaceId).EventType(eventType).Actor(actor).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAuditLogsAPI.KmsListAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListAuditLogs`: KmsListAuditLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAuditLogsAPI.KmsListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **eventType** | **string** |  | 
 **actor** | **string** |  | 
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**KmsListAuditLogs200Response**](KmsListAuditLogs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

