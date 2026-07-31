# \CommerceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1CommerceAdminCatalog**](CommerceAPI.md#CloudGetV1CommerceAdminCatalog) | **Get** /v1/commerce/admin/catalog | 
[**CloudGetV1CommerceCatalog**](CommerceAPI.md#CloudGetV1CommerceCatalog) | **Get** /v1/commerce/catalog | 
[**CloudGetV1CommerceCurrencies**](CommerceAPI.md#CloudGetV1CommerceCurrencies) | **Get** /v1/commerce/currencies | 
[**CloudGetV1CommerceDepositsByIdStatus**](CommerceAPI.md#CloudGetV1CommerceDepositsByIdStatus) | **Get** /v1/commerce/deposits/{id}/status | 
[**CloudGetV1CommerceTenant**](CommerceAPI.md#CloudGetV1CommerceTenant) | **Get** /v1/commerce/tenant | 
[**CloudGetV1CommerceTopupRails**](CommerceAPI.md#CloudGetV1CommerceTopupRails) | **Get** /v1/commerce/topup/rails | TopupRails lists the accepted (chain, token, treasury) triples, so a browser can render \&quot;send USDC here\&quot; without the addresses being baked into its bundle.
[**CloudPostV1CommerceDeposits**](CommerceAPI.md#CloudPostV1CommerceDeposits) | **Post** /v1/commerce/deposits | 
[**CloudPostV1CommerceDepositsByIdConfirm**](CommerceAPI.md#CloudPostV1CommerceDepositsByIdConfirm) | **Post** /v1/commerce/deposits/{id}/confirm | 
[**CloudPostV1CommerceTopupWallet**](CommerceAPI.md#CloudPostV1CommerceTopupWallet) | **Post** /v1/commerce/topup/wallet | WalletTopup credits the caller&#39;s org for a stablecoin transfer they already sent to the treasury.
[**CloudPostV1CommerceWebhooksByProvider**](CommerceAPI.md#CloudPostV1CommerceWebhooksByProvider) | **Post** /v1/commerce/webhooks/{provider} | 



## CloudGetV1CommerceAdminCatalog

> CloudGetV1CommerceAdminCatalog(ctx).Execute()



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
	r, err := apiClient.CommerceAPI.CloudGetV1CommerceAdminCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceAdminCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceAdminCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CommerceCatalog

> CloudGetV1CommerceCatalog(ctx).Execute()



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
	r, err := apiClient.CommerceAPI.CloudGetV1CommerceCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceCatalogRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CommerceCurrencies

> CloudGetV1CommerceCurrencies(ctx).Execute()



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
	r, err := apiClient.CommerceAPI.CloudGetV1CommerceCurrencies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceCurrencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceCurrenciesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CommerceDepositsByIdStatus

> CloudGetV1CommerceDepositsByIdStatus(ctx, id).Execute()



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
	r, err := apiClient.CommerceAPI.CloudGetV1CommerceDepositsByIdStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceDepositsByIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceDepositsByIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CommerceTenant

> CloudGetV1CommerceTenant(ctx).Execute()



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
	r, err := apiClient.CommerceAPI.CloudGetV1CommerceTenant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceTenantRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CommerceTopupRails

> CloudRailList CloudGetV1CommerceTopupRails(ctx).Execute()

TopupRails lists the accepted (chain, token, treasury) triples, so a browser can render \"send USDC here\" without the addresses being baked into its bundle.



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
	resp, r, err := apiClient.CommerceAPI.CloudGetV1CommerceTopupRails(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudGetV1CommerceTopupRails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CommerceTopupRails`: CloudRailList
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.CloudGetV1CommerceTopupRails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CommerceTopupRailsRequest struct via the builder pattern


### Return type

[**CloudRailList**](CloudRailList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CommerceDeposits

> CloudPostV1CommerceDeposits(ctx).Execute()



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
	r, err := apiClient.CommerceAPI.CloudPostV1CommerceDeposits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudPostV1CommerceDeposits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CommerceDepositsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CommerceDepositsByIdConfirm

> CloudPostV1CommerceDepositsByIdConfirm(ctx, id).Execute()



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
	r, err := apiClient.CommerceAPI.CloudPostV1CommerceDepositsByIdConfirm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudPostV1CommerceDepositsByIdConfirm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CommerceDepositsByIdConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CommerceTopupWallet

> CloudWalletTopupResp CloudPostV1CommerceTopupWallet(ctx).CloudWalletTopupReq(cloudWalletTopupReq).Execute()

WalletTopup credits the caller's org for a stablecoin transfer they already sent to the treasury.



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
	cloudWalletTopupReq := *openapiclient.NewCloudWalletTopupReq() // CloudWalletTopupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceAPI.CloudPostV1CommerceTopupWallet(context.Background()).CloudWalletTopupReq(cloudWalletTopupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudPostV1CommerceTopupWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CommerceTopupWallet`: CloudWalletTopupResp
	fmt.Fprintf(os.Stdout, "Response from `CommerceAPI.CloudPostV1CommerceTopupWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CommerceTopupWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudWalletTopupReq** | [**CloudWalletTopupReq**](CloudWalletTopupReq.md) |  | 

### Return type

[**CloudWalletTopupResp**](CloudWalletTopupResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CommerceWebhooksByProvider

> CloudPostV1CommerceWebhooksByProvider(ctx, provider).Execute()



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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommerceAPI.CloudPostV1CommerceWebhooksByProvider(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceAPI.CloudPostV1CommerceWebhooksByProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CommerceWebhooksByProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

