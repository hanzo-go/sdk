# \AffiliateAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AffiliatesApplyAffiliate**](AffiliateAPI.md#AffiliatesApplyAffiliate) | **Post** /v1/affiliates/apply | Apply to the affiliate program
[**AffiliatesAttributeAffiliate**](AffiliateAPI.md#AffiliatesAttributeAffiliate) | **Post** /v1/affiliates/attribute | Record affiliate attribution
[**AffiliatesGetMyAffiliate**](AffiliateAPI.md#AffiliatesGetMyAffiliate) | **Get** /v1/affiliates | Get my affiliate status



## AffiliatesApplyAffiliate

> AffiliatesApplyResponse AffiliatesApplyAffiliate(ctx).AffiliatesApplyRequest(affiliatesApplyRequest).Execute()

Apply to the affiliate program



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
	affiliatesApplyRequest := *openapiclient.NewAffiliatesApplyRequest() // AffiliatesApplyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.AffiliatesApplyAffiliate(context.Background()).AffiliatesApplyRequest(affiliatesApplyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.AffiliatesApplyAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesApplyAffiliate`: AffiliatesApplyResponse
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.AffiliatesApplyAffiliate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesApplyAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliatesApplyRequest** | [**AffiliatesApplyRequest**](AffiliatesApplyRequest.md) |  | 

### Return type

[**AffiliatesApplyResponse**](AffiliatesApplyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesAttributeAffiliate

> AffiliatesAttributeResponse AffiliatesAttributeAffiliate(ctx).AffiliatesAttributeRequest(affiliatesAttributeRequest).Execute()

Record affiliate attribution



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
	affiliatesAttributeRequest := *openapiclient.NewAffiliatesAttributeRequest("Code_example") // AffiliatesAttributeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliateAPI.AffiliatesAttributeAffiliate(context.Background()).AffiliatesAttributeRequest(affiliatesAttributeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.AffiliatesAttributeAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAttributeAffiliate`: AffiliatesAttributeResponse
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.AffiliatesAttributeAffiliate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAttributeAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **affiliatesAttributeRequest** | [**AffiliatesAttributeRequest**](AffiliatesAttributeRequest.md) |  | 

### Return type

[**AffiliatesAttributeResponse**](AffiliatesAttributeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesGetMyAffiliate

> AffiliatesGetMyAffiliate200Response AffiliatesGetMyAffiliate(ctx).Execute()

Get my affiliate status



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
	resp, r, err := apiClient.AffiliateAPI.AffiliatesGetMyAffiliate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliateAPI.AffiliatesGetMyAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesGetMyAffiliate`: AffiliatesGetMyAffiliate200Response
	fmt.Fprintf(os.Stdout, "Response from `AffiliateAPI.AffiliatesGetMyAffiliate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesGetMyAffiliateRequest struct via the builder pattern


### Return type

[**AffiliatesGetMyAffiliate200Response**](AffiliatesGetMyAffiliate200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

