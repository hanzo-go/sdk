# \CommerceAffiliatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceConnectAffiliate**](CommerceAffiliatesAPI.md#CommerceConnectAffiliate) | **Get** /v1/commerce/affiliate/{affiliateid}/connect | Connect affiliate
[**CommerceCreateAffiliate**](CommerceAffiliatesAPI.md#CommerceCreateAffiliate) | **Post** /v1/commerce/affiliate | Create affiliate
[**CommerceGetAffiliate**](CommerceAffiliatesAPI.md#CommerceGetAffiliate) | **Get** /v1/commerce/affiliate/{affiliateid} | Get affiliate
[**CommerceListAffiliates**](CommerceAffiliatesAPI.md#CommerceListAffiliates) | **Get** /v1/commerce/affiliate | List affiliates



## CommerceConnectAffiliate

> map[string]interface{} CommerceConnectAffiliate(ctx, affiliateid).Execute()

Connect affiliate

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
	affiliateid := "affiliateid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAffiliatesAPI.CommerceConnectAffiliate(context.Background(), affiliateid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAffiliatesAPI.CommerceConnectAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceConnectAffiliate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommerceAffiliatesAPI.CommerceConnectAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**affiliateid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceConnectAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateAffiliate

> CommerceAffiliate CommerceCreateAffiliate(ctx).CommerceAffiliate(commerceAffiliate).Execute()

Create affiliate

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
	commerceAffiliate := *openapiclient.NewCommerceAffiliate() // CommerceAffiliate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAffiliatesAPI.CommerceCreateAffiliate(context.Background()).CommerceAffiliate(commerceAffiliate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAffiliatesAPI.CommerceCreateAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateAffiliate`: CommerceAffiliate
	fmt.Fprintf(os.Stdout, "Response from `CommerceAffiliatesAPI.CommerceCreateAffiliate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceAffiliate** | [**CommerceAffiliate**](CommerceAffiliate.md) |  | 

### Return type

[**CommerceAffiliate**](CommerceAffiliate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetAffiliate

> CommerceAffiliate CommerceGetAffiliate(ctx, affiliateid).Execute()

Get affiliate

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
	affiliateid := "affiliateid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAffiliatesAPI.CommerceGetAffiliate(context.Background(), affiliateid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAffiliatesAPI.CommerceGetAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetAffiliate`: CommerceAffiliate
	fmt.Fprintf(os.Stdout, "Response from `CommerceAffiliatesAPI.CommerceGetAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**affiliateid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceAffiliate**](CommerceAffiliate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListAffiliates

> CommercePaginatedAffiliates CommerceListAffiliates(ctx).Page(page).Display(display).Execute()

List affiliates

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
	page := int32(56) // int32 | Page number (1-indexed) (optional) (default to 1)
	display := int32(56) // int32 | Number of items per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAffiliatesAPI.CommerceListAffiliates(context.Background()).Page(page).Display(display).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAffiliatesAPI.CommerceListAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListAffiliates`: CommercePaginatedAffiliates
	fmt.Fprintf(os.Stdout, "Response from `CommerceAffiliatesAPI.CommerceListAffiliates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListAffiliatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]

### Return type

[**CommercePaginatedAffiliates**](CommercePaginatedAffiliates.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

