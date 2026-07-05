# \AdminObservabilityAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminCompute**](AdminObservabilityAPI.md#AdminAdminCompute) | **Get** /v1/admin/compute | Cross-tenant compute analytics
[**AdminAdminO11y**](AdminObservabilityAPI.md#AdminAdminO11y) | **Get** /v1/admin/o11y | Fleet-wide observability board



## AdminAdminCompute

> AdminAdminCompute200Response AdminAdminCompute(ctx).Kind(kind).Org(org).Range_(range_).Execute()

Cross-tenant compute analytics

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
	kind := "kind_example" // string | bot | machine | cluster | nodepool | container | function | … (optional)
	org := "org_example" // string |  (optional)
	range_ := "range__example" // string |  (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminObservabilityAPI.AdminAdminCompute(context.Background()).Kind(kind).Org(org).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminObservabilityAPI.AdminAdminCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminCompute`: AdminAdminCompute200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminObservabilityAPI.AdminAdminCompute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminComputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | bot | machine | cluster | nodepool | container | function | … | 
 **org** | **string** |  | 
 **range_** | **string** |  | [default to &quot;30d&quot;]

### Return type

[**AdminAdminCompute200Response**](AdminAdminCompute200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminO11y

> AdminAdminO11y200Response AdminAdminO11y(ctx).Range_(range_).Execute()

Fleet-wide observability board

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
	range_ := "range__example" // string |  (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminObservabilityAPI.AdminAdminO11y(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminObservabilityAPI.AdminAdminO11y``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminO11y`: AdminAdminO11y200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminObservabilityAPI.AdminAdminO11y`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminO11yRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  | [default to &quot;30d&quot;]

### Return type

[**AdminAdminO11y200Response**](AdminAdminO11y200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

