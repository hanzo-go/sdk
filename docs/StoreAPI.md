# \StoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1StoreByStoreid**](StoreAPI.md#CloudDeleteV1StoreByStoreid) | **Delete** /v1/store/{storeid} | 
[**CloudDeleteV1StoreByStoreidListingByKey**](StoreAPI.md#CloudDeleteV1StoreByStoreidListingByKey) | **Delete** /v1/store/{storeid}/listing/{key} | 
[**CloudGetV1Store**](StoreAPI.md#CloudGetV1Store) | **Get** /v1/store/ | 
[**CloudGetV1StoreAccess**](StoreAPI.md#CloudGetV1StoreAccess) | **Get** /v1/store/access | 
[**CloudGetV1StoreByStoreid**](StoreAPI.md#CloudGetV1StoreByStoreid) | **Get** /v1/store/{storeid} | 
[**CloudGetV1StoreByStoreidBundleByKey**](StoreAPI.md#CloudGetV1StoreByStoreidBundleByKey) | **Get** /v1/store/{storeid}/bundle/{key} | 
[**CloudGetV1StoreByStoreidListing**](StoreAPI.md#CloudGetV1StoreByStoreidListing) | **Get** /v1/store/{storeid}/listing | 
[**CloudGetV1StoreByStoreidListingByKey**](StoreAPI.md#CloudGetV1StoreByStoreidListingByKey) | **Get** /v1/store/{storeid}/listing/{key} | 
[**CloudGetV1StoreByStoreidProductByKey**](StoreAPI.md#CloudGetV1StoreByStoreidProductByKey) | **Get** /v1/store/{storeid}/product/{key} | 
[**CloudGetV1StoreByStoreidVariantByKey**](StoreAPI.md#CloudGetV1StoreByStoreidVariantByKey) | **Get** /v1/store/{storeid}/variant/{key} | 
[**CloudGetV1StoreCurrent**](StoreAPI.md#CloudGetV1StoreCurrent) | **Get** /v1/store/current | 
[**CloudPatchV1StoreByStoreid**](StoreAPI.md#CloudPatchV1StoreByStoreid) | **Patch** /v1/store/{storeid} | 
[**CloudPatchV1StoreByStoreidListingByKey**](StoreAPI.md#CloudPatchV1StoreByStoreidListingByKey) | **Patch** /v1/store/{storeid}/listing/{key} | 
[**CloudPostV1Store**](StoreAPI.md#CloudPostV1Store) | **Post** /v1/store/ | 
[**CloudPostV1StoreByStoreid**](StoreAPI.md#CloudPostV1StoreByStoreid) | **Post** /v1/store/{storeid} | 
[**CloudPostV1StoreByStoreidAuthorize**](StoreAPI.md#CloudPostV1StoreByStoreidAuthorize) | **Post** /v1/store/{storeid}/authorize | 
[**CloudPostV1StoreByStoreidAuthorizeByOrderid**](StoreAPI.md#CloudPostV1StoreByStoreidAuthorizeByOrderid) | **Post** /v1/store/{storeid}/authorize/{orderid} | 
[**CloudPostV1StoreByStoreidCaptureByOrderid**](StoreAPI.md#CloudPostV1StoreByStoreidCaptureByOrderid) | **Post** /v1/store/{storeid}/capture/{orderid} | 
[**CloudPostV1StoreByStoreidCharge**](StoreAPI.md#CloudPostV1StoreByStoreidCharge) | **Post** /v1/store/{storeid}/charge | 
[**CloudPostV1StoreByStoreidCheckoutAuthorize**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutAuthorize) | **Post** /v1/store/{storeid}/checkout/authorize | 
[**CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid) | **Post** /v1/store/{storeid}/checkout/authorize/{orderid} | 
[**CloudPostV1StoreByStoreidCheckoutCaptureByOrderid**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutCaptureByOrderid) | **Post** /v1/store/{storeid}/checkout/capture/{orderid} | 
[**CloudPostV1StoreByStoreidCheckoutCharge**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutCharge) | **Post** /v1/store/{storeid}/checkout/charge | 
[**CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey) | **Post** /v1/store/{storeid}/checkout/paypal/cancel/{payKey} | 
[**CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey) | **Post** /v1/store/{storeid}/checkout/paypal/confirm/{payKey} | 
[**CloudPostV1StoreByStoreidCheckoutPaypalPay**](StoreAPI.md#CloudPostV1StoreByStoreidCheckoutPaypalPay) | **Post** /v1/store/{storeid}/checkout/paypal/pay | 
[**CloudPostV1StoreByStoreidListingByKey**](StoreAPI.md#CloudPostV1StoreByStoreidListingByKey) | **Post** /v1/store/{storeid}/listing/{key} | 
[**CloudPostV1StoreByStoreidPaypalCancelByPaykey**](StoreAPI.md#CloudPostV1StoreByStoreidPaypalCancelByPaykey) | **Post** /v1/store/{storeid}/paypal/cancel/{payKey} | 
[**CloudPostV1StoreByStoreidPaypalConfirmByPaykey**](StoreAPI.md#CloudPostV1StoreByStoreidPaypalConfirmByPaykey) | **Post** /v1/store/{storeid}/paypal/confirm/{payKey} | 
[**CloudPostV1StoreByStoreidPaypalPay**](StoreAPI.md#CloudPostV1StoreByStoreidPaypalPay) | **Post** /v1/store/{storeid}/paypal/pay | 
[**CloudPostV1StoreByStoreidTrial**](StoreAPI.md#CloudPostV1StoreByStoreidTrial) | **Post** /v1/store/{storeid}/trial | 
[**CloudPostV1StoreStorefrontToken**](StoreAPI.md#CloudPostV1StoreStorefrontToken) | **Post** /v1/store/storefront-token | 
[**CloudPutV1StoreByStoreid**](StoreAPI.md#CloudPutV1StoreByStoreid) | **Put** /v1/store/{storeid} | 
[**CloudPutV1StoreByStoreidListingByKey**](StoreAPI.md#CloudPutV1StoreByStoreidListingByKey) | **Put** /v1/store/{storeid}/listing/{key} | 
[**CommerceCreateStore**](StoreAPI.md#CommerceCreateStore) | **Post** /v1/commerce/store | Create store
[**CommerceCreateStoreListing**](StoreAPI.md#CommerceCreateStoreListing) | **Post** /v1/commerce/store/{storeid}/listing/{key} | Create store listing
[**CommerceDeleteStoreListing**](StoreAPI.md#CommerceDeleteStoreListing) | **Delete** /v1/commerce/store/{storeid}/listing/{key} | Delete store listing
[**CommerceGetStore**](StoreAPI.md#CommerceGetStore) | **Get** /v1/commerce/store/{storeid} | Get store
[**CommerceGetStoreListing**](StoreAPI.md#CommerceGetStoreListing) | **Get** /v1/commerce/store/{storeid}/listing/{key} | Get store listing
[**CommerceGetStoreProduct**](StoreAPI.md#CommerceGetStoreProduct) | **Get** /v1/commerce/store/{storeid}/product/{key} | Get store product
[**CommerceGetStoreVariant**](StoreAPI.md#CommerceGetStoreVariant) | **Get** /v1/commerce/store/{storeid}/variant/{key} | Get store variant
[**CommerceListStoreListings**](StoreAPI.md#CommerceListStoreListings) | **Get** /v1/commerce/store/{storeid}/listing | List store listings
[**CommerceListStores**](StoreAPI.md#CommerceListStores) | **Get** /v1/commerce/store | List stores
[**CommercePatchStoreListing**](StoreAPI.md#CommercePatchStoreListing) | **Patch** /v1/commerce/store/{storeid}/listing/{key} | Partially update store listing
[**CommerceStoreAuthorize**](StoreAPI.md#CommerceStoreAuthorize) | **Post** /v1/commerce/store/{storeid}/checkout/authorize | Authorize payment via store
[**CommerceStoreCharge**](StoreAPI.md#CommerceStoreCharge) | **Post** /v1/commerce/store/{storeid}/checkout/charge | Charge payment via store
[**CommerceUpdateStoreListing**](StoreAPI.md#CommerceUpdateStoreListing) | **Put** /v1/commerce/store/{storeid}/listing/{key} | Update store listing



## CloudDeleteV1StoreByStoreid

> CloudDeleteV1StoreByStoreid(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudDeleteV1StoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudDeleteV1StoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1StoreByStoreidRequest struct via the builder pattern


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


## CloudDeleteV1StoreByStoreidListingByKey

> CloudDeleteV1StoreByStoreidListingByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudDeleteV1StoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudDeleteV1StoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1StoreByStoreidListingByKeyRequest struct via the builder pattern


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


## CloudGetV1Store

> CloudGetV1Store(ctx).Execute()



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
	r, err := apiClient.StoreAPI.CloudGetV1Store(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1Store``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreRequest struct via the builder pattern


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


## CloudGetV1StoreAccess

> CloudGetV1StoreAccess(ctx).Execute()



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
	r, err := apiClient.StoreAPI.CloudGetV1StoreAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreAccessRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreid

> CloudGetV1StoreByStoreid(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreidBundleByKey

> CloudGetV1StoreByStoreidBundleByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreidBundleByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreidBundleByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidBundleByKeyRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreidListing

> CloudGetV1StoreByStoreidListing(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreidListing(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreidListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidListingRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreidListingByKey

> CloudGetV1StoreByStoreidListingByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidListingByKeyRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreidProductByKey

> CloudGetV1StoreByStoreidProductByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreidProductByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreidProductByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidProductByKeyRequest struct via the builder pattern


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


## CloudGetV1StoreByStoreidVariantByKey

> CloudGetV1StoreByStoreidVariantByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudGetV1StoreByStoreidVariantByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreByStoreidVariantByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreByStoreidVariantByKeyRequest struct via the builder pattern


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


## CloudGetV1StoreCurrent

> CloudGetV1StoreCurrent(ctx).Execute()



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
	r, err := apiClient.StoreAPI.CloudGetV1StoreCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudGetV1StoreCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1StoreCurrentRequest struct via the builder pattern


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


## CloudPatchV1StoreByStoreid

> CloudPatchV1StoreByStoreid(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPatchV1StoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPatchV1StoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1StoreByStoreidRequest struct via the builder pattern


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


## CloudPatchV1StoreByStoreidListingByKey

> CloudPatchV1StoreByStoreidListingByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPatchV1StoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPatchV1StoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1StoreByStoreidListingByKeyRequest struct via the builder pattern


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


## CloudPostV1Store

> CloudPostV1Store(ctx).Execute()



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
	r, err := apiClient.StoreAPI.CloudPostV1Store(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1Store``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreid

> CloudPostV1StoreByStoreid(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidAuthorize

> CloudPostV1StoreByStoreidAuthorize(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidAuthorizeRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidAuthorizeByOrderid

> CloudPostV1StoreByStoreidAuthorizeByOrderid(ctx, storeid, orderid).Execute()



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidAuthorizeByOrderidRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCaptureByOrderid

> CloudPostV1StoreByStoreidCaptureByOrderid(ctx, storeid, orderid).Execute()



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCaptureByOrderidRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCharge

> CloudPostV1StoreByStoreidCharge(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidChargeRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutAuthorize

> CloudPostV1StoreByStoreidCheckoutAuthorize(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutAuthorize(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutAuthorizeRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid

> CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid(ctx, storeid, orderid).Execute()



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutAuthorizeByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutAuthorizeByOrderidRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutCaptureByOrderid

> CloudPostV1StoreByStoreidCheckoutCaptureByOrderid(ctx, storeid, orderid).Execute()



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
	storeid := "storeid_example" // string | 
	orderid := "orderid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutCaptureByOrderid(context.Background(), storeid, orderid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutCaptureByOrderid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**orderid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutCaptureByOrderidRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutCharge

> CloudPostV1StoreByStoreidCheckoutCharge(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutCharge(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutChargeRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey

> CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey(ctx, storeid, payKey).Execute()



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutPaypalCancelByPaykeyRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey

> CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutPaypalConfirmByPaykeyRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidCheckoutPaypalPay

> CloudPostV1StoreByStoreidCheckoutPaypalPay(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidCheckoutPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidCheckoutPaypalPayRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidListingByKey

> CloudPostV1StoreByStoreidListingByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidListingByKeyRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidPaypalCancelByPaykey

> CloudPostV1StoreByStoreidPaypalCancelByPaykey(ctx, storeid, payKey).Execute()



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidPaypalCancelByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidPaypalCancelByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidPaypalCancelByPaykeyRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidPaypalConfirmByPaykey

> CloudPostV1StoreByStoreidPaypalConfirmByPaykey(ctx, storeid, payKey).Execute()



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
	storeid := "storeid_example" // string | 
	payKey := "payKey_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidPaypalConfirmByPaykey(context.Background(), storeid, payKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidPaypalConfirmByPaykey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**payKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidPaypalConfirmByPaykeyRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidPaypalPay

> CloudPostV1StoreByStoreidPaypalPay(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidPaypalPay(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidPaypalPay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidPaypalPayRequest struct via the builder pattern


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


## CloudPostV1StoreByStoreidTrial

> CloudPostV1StoreByStoreidTrial(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPostV1StoreByStoreidTrial(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreByStoreidTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreByStoreidTrialRequest struct via the builder pattern


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


## CloudPostV1StoreStorefrontToken

> CloudPostV1StoreStorefrontToken(ctx).Execute()



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
	r, err := apiClient.StoreAPI.CloudPostV1StoreStorefrontToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPostV1StoreStorefrontToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1StoreStorefrontTokenRequest struct via the builder pattern


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


## CloudPutV1StoreByStoreid

> CloudPutV1StoreByStoreid(ctx, storeid).Execute()



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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPutV1StoreByStoreid(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPutV1StoreByStoreid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1StoreByStoreidRequest struct via the builder pattern


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


## CloudPutV1StoreByStoreidListingByKey

> CloudPutV1StoreByStoreidListingByKey(ctx, storeid, key).Execute()



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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CloudPutV1StoreByStoreidListingByKey(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CloudPutV1StoreByStoreidListingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1StoreByStoreidListingByKeyRequest struct via the builder pattern


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


## CommerceCreateStore

> CommerceStore CommerceCreateStore(ctx).CommerceStore(commerceStore).Execute()

Create store

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
	commerceStore := *openapiclient.NewCommerceStore() // CommerceStore | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceCreateStore(context.Background()).CommerceStore(commerceStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceCreateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateStore`: CommerceStore
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceCreateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceStore** | [**CommerceStore**](CommerceStore.md) |  | 

### Return type

[**CommerceStore**](CommerceStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateStoreListing

> CommerceListing CommerceCreateStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Create store listing

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceCreateStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceCreateStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceCreateStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceDeleteStoreListing

> CommerceDeleteStoreListing(ctx, storeid, key).Execute()

Delete store listing

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StoreAPI.CommerceDeleteStoreListing(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceDeleteStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceDeleteStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStore

> CommerceStore CommerceGetStore(ctx, storeid).Execute()

Get store

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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceGetStore(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStore`: CommerceStore
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceStore**](CommerceStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreListing

> CommerceListing CommerceGetStoreListing(ctx, storeid, key).Execute()

Get store listing

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreListing(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreProduct

> CommerceProduct CommerceGetStoreProduct(ctx, storeid, key).Execute()

Get store product

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreProduct(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreProduct`: CommerceProduct
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceProduct**](CommerceProduct.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetStoreVariant

> CommerceVariant CommerceGetStoreVariant(ctx, storeid, key).Execute()

Get store variant

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceGetStoreVariant(context.Background(), storeid, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceGetStoreVariant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetStoreVariant`: CommerceVariant
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceGetStoreVariant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetStoreVariantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CommerceVariant**](CommerceVariant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListStoreListings

> []CommerceListing CommerceListStoreListings(ctx, storeid).Execute()

List store listings

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
	storeid := "storeid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceListStoreListings(context.Background(), storeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceListStoreListings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListStoreListings`: []CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceListStoreListings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListStoreListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListStores

> CommercePaginatedStores CommerceListStores(ctx).Page(page).Display(display).Execute()

List stores

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
	resp, r, err := apiClient.StoreAPI.CommerceListStores(context.Background()).Page(page).Display(display).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceListStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListStores`: CommercePaginatedStores
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceListStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]

### Return type

[**CommercePaginatedStores**](CommercePaginatedStores.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommercePatchStoreListing

> CommerceListing CommercePatchStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Partially update store listing

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommercePatchStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommercePatchStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommercePatchStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommercePatchStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommercePatchStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceStoreAuthorize

> CommerceOrder CommerceStoreAuthorize(ctx, storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Authorize payment via store

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
	storeid := "storeid_example" // string | 
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceStoreAuthorize(context.Background(), storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceStoreAuthorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceStoreAuthorize`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceStoreAuthorize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceStoreAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCheckoutRequest** | [**CommerceCheckoutRequest**](CommerceCheckoutRequest.md) |  | 

### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceStoreCharge

> CommerceOrder CommerceStoreCharge(ctx, storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()

Charge payment via store

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
	storeid := "storeid_example" // string | 
	commerceCheckoutRequest := *openapiclient.NewCommerceCheckoutRequest() // CommerceCheckoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceStoreCharge(context.Background(), storeid).CommerceCheckoutRequest(commerceCheckoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceStoreCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceStoreCharge`: CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceStoreCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceStoreChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commerceCheckoutRequest** | [**CommerceCheckoutRequest**](CommerceCheckoutRequest.md) |  | 

### Return type

[**CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceUpdateStoreListing

> CommerceListing CommerceUpdateStoreListing(ctx, storeid, key).CommerceListing(commerceListing).Execute()

Update store listing

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
	storeid := "storeid_example" // string | 
	key := "key_example" // string | 
	commerceListing := *openapiclient.NewCommerceListing() // CommerceListing | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreAPI.CommerceUpdateStoreListing(context.Background(), storeid, key).CommerceListing(commerceListing).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreAPI.CommerceUpdateStoreListing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceUpdateStoreListing`: CommerceListing
	fmt.Fprintf(os.Stdout, "Response from `StoreAPI.CommerceUpdateStoreListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeid** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceUpdateStoreListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **commerceListing** | [**CommerceListing**](CommerceListing.md) |  | 

### Return type

[**CommerceListing**](CommerceListing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

