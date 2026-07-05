# \NexusActivityAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusGetActivities**](NexusActivityAPIAPI.md#NexusGetActivities) | **Get** /v1/nexus/get-activities | get Activities



## NexusGetActivities

> []NexusActivity NexusGetActivities(ctx).Days(days).Execute()

get Activities



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
	days := "days_example" // string | Number of days

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusActivityAPIAPI.NexusGetActivities(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusActivityAPIAPI.NexusGetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetActivities`: []NexusActivity
	fmt.Fprintf(os.Stdout, "Response from `NexusActivityAPIAPI.NexusGetActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **string** | Number of days | 

### Return type

[**[]NexusActivity**](NexusActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

