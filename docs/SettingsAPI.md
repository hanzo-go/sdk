# \SettingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettingsByProduct**](SettingsAPI.md#GetSettingsByProduct) | **Get** /v1/settings/{product} | Reads the caller org&#39;s configuration for one product, with every secret field MASKED — only the names of the set secrets come back, never their values, which live in KMS.
[**PutSettingsByProduct**](SettingsAPI.md#PutSettingsByProduct) | **Put** /v1/settings/{product} | Writes the caller org&#39;s configuration for one product and answers the stored result, secrets masked.



## GetSettingsByProduct

> SettingsView GetSettingsByProduct(ctx, product).Execute()

Reads the caller org's configuration for one product, with every secret field MASKED — only the names of the set secrets come back, never their values, which live in KMS.



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
	product := "product_example" // string | Product is the catalog slug, from the path. Must match ^[a-z0-9][a-z0-9._-]{0,62}$.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetSettingsByProduct(context.Background(), product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetSettingsByProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettingsByProduct`: SettingsView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetSettingsByProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** | Product is the catalog slug, from the path. Must match ^[a-z0-9][a-z0-9._-]{0,62}$. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsView**](SettingsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsByProduct

> SettingsView PutSettingsByProduct(ctx, product).SettingsReq(settingsReq).Execute()

Writes the caller org's configuration for one product and answers the stored result, secrets masked.



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
	product := "product_example" // string | Product is the catalog slug, from the PATH. zip binds the path last, so the URL names the product being written whatever a body field claims.
	settingsReq := *openapiclient.NewSettingsReq() // SettingsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PutSettingsByProduct(context.Background(), product).SettingsReq(settingsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PutSettingsByProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutSettingsByProduct`: SettingsView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PutSettingsByProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** | Product is the catalog slug, from the PATH. zip binds the path last, so the URL names the product being written whatever a body field claims. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsByProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsReq** | [**SettingsReq**](SettingsReq.md) |  | 

### Return type

[**SettingsView**](SettingsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

