# \AuditAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminListAudit**](AuditAPI.md#AdminAdminListAudit) | **Get** /v1/admin/audit | Query the tamper-evident audit trail
[**AdminAdminVerifyAudit**](AuditAPI.md#AdminAdminVerifyAudit) | **Get** /v1/admin/audit/verify | Verify audit-chain integrity



## AdminAdminListAudit

> AdminAdminListAudit200Response AdminAdminListAudit(ctx).Org(org).Sub(sub).Action(action).Resource(resource).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()

Query the tamper-evident audit trail

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
	org := "org_example" // string |  (optional)
	sub := "sub_example" // string |  (optional)
	action := "action_example" // string |  (optional)
	resource := "resource_example" // string |  (optional)
	result := "result_example" // string |  (optional)
	since := time.Now() // time.Time |  (optional)
	until := time.Now() // time.Time |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 100)
	p := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.AdminAdminListAudit(context.Background()).Org(org).Sub(sub).Action(action).Resource(resource).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.AdminAdminListAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminListAudit`: AdminAdminListAudit200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.AdminAdminListAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminListAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** |  | 
 **sub** | **string** |  | 
 **action** | **string** |  | 
 **resource** | **string** |  | 
 **result** | **string** |  | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **pageSize** | **int32** |  | [default to 100]
 **p** | **int32** |  | [default to 1]

### Return type

[**AdminAdminListAudit200Response**](AdminAdminListAudit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminVerifyAudit

> AdminAdminVerifyAudit200Response AdminAdminVerifyAudit(ctx).Execute()

Verify audit-chain integrity

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
	resp, r, err := apiClient.AuditAPI.AdminAdminVerifyAudit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.AdminAdminVerifyAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminVerifyAudit`: AdminAdminVerifyAudit200Response
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.AdminAdminVerifyAudit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminVerifyAuditRequest struct via the builder pattern


### Return type

[**AdminAdminVerifyAudit200Response**](AdminAdminVerifyAudit200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

