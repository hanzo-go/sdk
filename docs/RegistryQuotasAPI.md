# \RegistryQuotasAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryListQuotas**](RegistryQuotasAPI.md#RegistryListQuotas) | **Get** /v1/registry/quotas | List storage quotas



## RegistryListQuotas

> []RegistryQuota RegistryListQuotas(ctx).Page(page).PageSize(pageSize).Sort(sort).Execute()

List storage quotas

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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistryQuotasAPI.RegistryListQuotas(context.Background()).Page(page).PageSize(pageSize).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistryQuotasAPI.RegistryListQuotas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegistryListQuotas`: []RegistryQuota
	fmt.Fprintf(os.Stdout, "Response from `RegistryQuotasAPI.RegistryListQuotas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegistryListQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **sort** | **string** |  | 

### Return type

[**[]RegistryQuota**](RegistryQuota.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

