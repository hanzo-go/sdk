# \DbConsumptionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbGetConsumption**](DbConsumptionAPI.md#DbGetConsumption) | **Get** /v1/db/consumption | Get usage metrics



## DbGetConsumption

> DbGetConsumption200Response DbGetConsumption(ctx).From(from).To(to).ProjectId(projectId).Granularity(granularity).Execute()

Get usage metrics



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
	from := time.Now() // time.Time | 
	to := time.Now() // time.Time | 
	projectId := "projectId_example" // string | Filter by project (omit for all projects) (optional)
	granularity := "granularity_example" // string |  (optional) (default to "hourly")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbConsumptionAPI.DbGetConsumption(context.Background()).From(from).To(to).ProjectId(projectId).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbConsumptionAPI.DbGetConsumption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetConsumption`: DbGetConsumption200Response
	fmt.Fprintf(os.Stdout, "Response from `DbConsumptionAPI.DbGetConsumption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDbGetConsumptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **projectId** | **string** | Filter by project (omit for all projects) | 
 **granularity** | **string** |  | [default to &quot;hourly&quot;]

### Return type

[**DbGetConsumption200Response**](DbGetConsumption200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

