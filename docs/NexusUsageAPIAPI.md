# \NexusUsageAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusGetRangeUsages**](NexusUsageAPIAPI.md#NexusGetRangeUsages) | **Get** /v1/nexus/get-range-usages | get Range Usages
[**NexusGetUsages**](NexusUsageAPIAPI.md#NexusGetUsages) | **Get** /v1/nexus/get-usages | get Usages
[**NexusGetUsers**](NexusUsageAPIAPI.md#NexusGetUsers) | **Get** /v1/nexus/get-users | get Users



## NexusGetRangeUsages

> []NexusUsage NexusGetRangeUsages(ctx).Count(count).Execute()

get Range Usages



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
	count := "count_example" // string | Count of range usages

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusUsageAPIAPI.NexusGetRangeUsages(context.Background()).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusUsageAPIAPI.NexusGetRangeUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetRangeUsages`: []NexusUsage
	fmt.Fprintf(os.Stdout, "Response from `NexusUsageAPIAPI.NexusGetRangeUsages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetRangeUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **string** | Count of range usages | 

### Return type

[**[]NexusUsage**](NexusUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetUsages

> []NexusUsage NexusGetUsages(ctx).Execute()

get Usages



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
	resp, r, err := apiClient.NexusUsageAPIAPI.NexusGetUsages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusUsageAPIAPI.NexusGetUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetUsages`: []NexusUsage
	fmt.Fprintf(os.Stdout, "Response from `NexusUsageAPIAPI.NexusGetUsages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetUsagesRequest struct via the builder pattern


### Return type

[**[]NexusUsage**](NexusUsage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetUsers

> []string NexusGetUsers(ctx).Execute()

get Users



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
	resp, r, err := apiClient.NexusUsageAPIAPI.NexusGetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusUsageAPIAPI.NexusGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetUsers`: []string
	fmt.Fprintf(os.Stdout, "Response from `NexusUsageAPIAPI.NexusGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetUsersRequest struct via the builder pattern


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

