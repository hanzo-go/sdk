# \PromotionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminDeletePromo**](PromotionsAPI.md#AdminAdminDeletePromo) | **Delete** /v1/admin/promos/{id} | Delete a discount promo
[**AdminAdminGetPromo**](PromotionsAPI.md#AdminAdminGetPromo) | **Get** /v1/admin/promos/{id} | Get one discount promo
[**AdminAdminUpdatePromo**](PromotionsAPI.md#AdminAdminUpdatePromo) | **Patch** /v1/admin/promos/{id} | Update a discount promo



## AdminAdminDeletePromo

> AdminAdminDeletePromo200Response AdminAdminDeletePromo(ctx, id).Execute()

Delete a discount promo

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromotionsAPI.AdminAdminDeletePromo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromotionsAPI.AdminAdminDeletePromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminDeletePromo`: AdminAdminDeletePromo200Response
	fmt.Fprintf(os.Stdout, "Response from `PromotionsAPI.AdminAdminDeletePromo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminDeletePromoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAdminDeletePromo200Response**](AdminAdminDeletePromo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminGetPromo

> AdminAdminGetPromo200Response AdminAdminGetPromo(ctx, id).Execute()

Get one discount promo

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromotionsAPI.AdminAdminGetPromo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromotionsAPI.AdminAdminGetPromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminGetPromo`: AdminAdminGetPromo200Response
	fmt.Fprintf(os.Stdout, "Response from `PromotionsAPI.AdminAdminGetPromo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminGetPromoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminAdminGetPromo200Response**](AdminAdminGetPromo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminUpdatePromo

> AdminAdminGetPromo200Response AdminAdminUpdatePromo(ctx, id).AdminPromoUpdate(adminPromoUpdate).Execute()

Update a discount promo



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
	id := "id_example" // string | 
	adminPromoUpdate := *openapiclient.NewAdminPromoUpdate() // AdminPromoUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromotionsAPI.AdminAdminUpdatePromo(context.Background(), id).AdminPromoUpdate(adminPromoUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromotionsAPI.AdminAdminUpdatePromo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminUpdatePromo`: AdminAdminGetPromo200Response
	fmt.Fprintf(os.Stdout, "Response from `PromotionsAPI.AdminAdminUpdatePromo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminUpdatePromoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminPromoUpdate** | [**AdminPromoUpdate**](AdminPromoUpdate.md) |  | 

### Return type

[**AdminAdminGetPromo200Response**](AdminAdminGetPromo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

