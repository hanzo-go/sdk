# \AuditAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Audit**](AuditAPI.md#CloudGetV1Audit) | **Get** /v1/audit | List reads the caller&#39;s OWN org audit trail, newest first, with the total the filter matched so a console can page it.



## CloudGetV1Audit

> CloudTrailPage CloudGetV1Audit(ctx).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()

List reads the caller's OWN org audit trail, newest first, with the total the filter matched so a console can page it.



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
	sub := "sub_example" // string | Sub narrows the trail to one actor — the validated subject that made the request. Blank means every actor in the org. (optional)
	action := "machine.create" // string | Action narrows it to one action name, e.g. \"machine.create\". (optional)
	resource := "resource_example" // string | Resource narrows it to one resource TYPE, e.g. \"apikey\". (optional)
	resourceId := "resourceId_example" // string | ResourceID narrows it to one resource instance. (optional)
	result := "success" // string | Result narrows it to one outcome: \"success\", \"deny\" or \"error\". (optional)
	since := "2026-07-01T00:00:00Z" // string | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. (optional)
	until := "until_example" // string | Until is the upper time bound, RFC3339, with the same tolerance. (optional)
	pageSize := "50" // string | PageSize is rows per page, default 100. A value that is not a positive integer falls back to the default. (optional)
	p := "p_example" // string | Page is the 1-based page number, driving the offset. Anything below 2 reads the first page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditAPI.CloudGetV1Audit(context.Background()).Sub(sub).Action(action).Resource(resource).ResourceId(resourceId).Result(result).Since(since).Until(until).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditAPI.CloudGetV1Audit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Audit`: CloudTrailPage
	fmt.Fprintf(os.Stdout, "Response from `AuditAPI.CloudGetV1Audit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sub** | **string** | Sub narrows the trail to one actor — the validated subject that made the request. Blank means every actor in the org. | 
 **action** | **string** | Action narrows it to one action name, e.g. \&quot;machine.create\&quot;. | 
 **resource** | **string** | Resource narrows it to one resource TYPE, e.g. \&quot;apikey\&quot;. | 
 **resourceId** | **string** | ResourceID narrows it to one resource instance. | 
 **result** | **string** | Result narrows it to one outcome: \&quot;success\&quot;, \&quot;deny\&quot; or \&quot;error\&quot;. | 
 **since** | **string** | Since is the inclusive lower time bound, RFC3339. An unparseable value is ignored rather than refused — one malformed filter must not hide the trail. | 
 **until** | **string** | Until is the upper time bound, RFC3339, with the same tolerance. | 
 **pageSize** | **string** | PageSize is rows per page, default 100. A value that is not a positive integer falls back to the default. | 
 **p** | **string** | Page is the 1-based page number, driving the offset. Anything below 2 reads the first page. | 

### Return type

[**CloudTrailPage**](CloudTrailPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

