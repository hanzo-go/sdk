# \GuardAuditAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuardGetAuditLog**](GuardAuditAPI.md#GuardGetAuditLog) | **Get** /v1/guard/audit | Get audit log



## GuardGetAuditLog

> GuardGetAuditLog200Response GuardGetAuditLog(ctx).UserId(userId).SessionId(sessionId).Result(result).Since(since).Until(until).Limit(limit).Execute()

Get audit log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/hanzoai/go-sdk"
)

func main() {
	userId := "userId_example" // string | Filter by user ID (optional)
	sessionId := "sessionId_example" // string | Filter by session ID (optional)
	result := "result_example" // string | Filter by result type (optional)
	since := time.Now() // time.Time | Entries after this timestamp (optional)
	until := time.Now() // time.Time | Entries before this timestamp (optional)
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuardAuditAPI.GuardGetAuditLog(context.Background()).UserId(userId).SessionId(sessionId).Result(result).Since(since).Until(until).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuardAuditAPI.GuardGetAuditLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GuardGetAuditLog`: GuardGetAuditLog200Response
	fmt.Fprintf(os.Stdout, "Response from `GuardAuditAPI.GuardGetAuditLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Filter by user ID | 
 **sessionId** | **string** | Filter by session ID | 
 **result** | **string** | Filter by result type | 
 **since** | **time.Time** | Entries after this timestamp | 
 **until** | **time.Time** | Entries before this timestamp | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**GuardGetAuditLog200Response**](GuardGetAuditLog200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

