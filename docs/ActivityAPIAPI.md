# \ActivityAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetActivities**](ActivityAPIAPI.md#CloudApiControllerGetActivities) | **Get** /v1/cloud/get-activities | Api Controller Get Activities



## CloudApiControllerGetActivities

> []CloudObjectActivity CloudApiControllerGetActivities(ctx).Days(days).Execute()

Api Controller Get Activities



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
	days := "days_example" // string | days count

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityAPIAPI.CloudApiControllerGetActivities(context.Background()).Days(days).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityAPIAPI.CloudApiControllerGetActivities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetActivities`: []CloudObjectActivity
	fmt.Fprintf(os.Stdout, "Response from `ActivityAPIAPI.CloudApiControllerGetActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **string** | days count | 

### Return type

[**[]CloudObjectActivity**](CloudObjectActivity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

