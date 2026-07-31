# \UsageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetRangeUsages**](UsageAPIAPI.md#CloudApiControllerGetRangeUsages) | **Get** /v1/cloud/get-range-usages | Api Controller Get Range Usages
[**CloudApiControllerGetUserTableInfos**](UsageAPIAPI.md#CloudApiControllerGetUserTableInfos) | **Get** /v1/cloud/get-usages | Api Controller Get User Table Infos
[**CloudApiControllerGetUsers**](UsageAPIAPI.md#CloudApiControllerGetUsers) | **Get** /v1/cloud/get-users | Api Controller Get Users



## CloudApiControllerGetRangeUsages

> []CloudObjectUsage CloudApiControllerGetRangeUsages(ctx).Count(count).Execute()

Api Controller Get Range Usages



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
	count := "count_example" // string | count of range usages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPIAPI.CloudApiControllerGetRangeUsages(context.Background()).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPIAPI.CloudApiControllerGetRangeUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetRangeUsages`: []CloudObjectUsage
	fmt.Fprintf(os.Stdout, "Response from `UsageAPIAPI.CloudApiControllerGetRangeUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetRangeUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **string** | count of range usages | 

### Return type

[**[]CloudObjectUsage**](CloudObjectUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetUserTableInfos

> []CloudObjectUsage CloudApiControllerGetUserTableInfos(ctx).Execute()

Api Controller Get User Table Infos



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
	resp, r, err := apiClient.UsageAPIAPI.CloudApiControllerGetUserTableInfos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPIAPI.CloudApiControllerGetUserTableInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetUserTableInfos`: []CloudObjectUsage
	fmt.Fprintf(os.Stdout, "Response from `UsageAPIAPI.CloudApiControllerGetUserTableInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetUserTableInfosRequest struct via the builder pattern


### Return type

[**[]CloudObjectUsage**](CloudObjectUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetUsers

> []string CloudApiControllerGetUsers(ctx).Execute()

Api Controller Get Users



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
	resp, r, err := apiClient.UsageAPIAPI.CloudApiControllerGetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPIAPI.CloudApiControllerGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetUsers`: []string
	fmt.Fprintf(os.Stdout, "Response from `UsageAPIAPI.CloudApiControllerGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetUsersRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

