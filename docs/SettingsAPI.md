# \SettingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1SettingsProduct**](SettingsAPI.md#CloudGetV1SettingsProduct) | **Get** /v1/settings/{product} | GetSettings reads the caller org&#39;s configuration for one product, with every secret field MASKED — only the names of the set secrets come back, never their values, which live in KMS.
[**CloudPutV1SettingsProduct**](SettingsAPI.md#CloudPutV1SettingsProduct) | **Put** /v1/settings/{product} | PutSettings writes the caller org&#39;s configuration for one product and answers the stored result, secrets masked.
[**SearchGetSettings**](SettingsAPI.md#SearchGetSettings) | **Get** /v1/search/indexes/{indexUid}/settings | Get all index settings
[**SearchResetSettings**](SettingsAPI.md#SearchResetSettings) | **Delete** /v1/search/indexes/{indexUid}/settings | Reset all settings to defaults
[**SearchUpdateSettings**](SettingsAPI.md#SearchUpdateSettings) | **Patch** /v1/search/indexes/{indexUid}/settings | Update index settings



## CloudGetV1SettingsProduct

> CloudSettingsView CloudGetV1SettingsProduct(ctx, product).Execute()

GetSettings reads the caller org's configuration for one product, with every secret field MASKED — only the names of the set secrets come back, never their values, which live in KMS.



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
	resp, r, err := apiClient.SettingsAPI.CloudGetV1SettingsProduct(context.Background(), product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CloudGetV1SettingsProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1SettingsProduct`: CloudSettingsView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CloudGetV1SettingsProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** | Product is the catalog slug, from the path. Must match ^[a-z0-9][a-z0-9._-]{0,62}$. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1SettingsProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSettingsView**](CloudSettingsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1SettingsProduct

> CloudSettingsView CloudPutV1SettingsProduct(ctx, product).CloudSettingsReq(cloudSettingsReq).Execute()

PutSettings writes the caller org's configuration for one product and answers the stored result, secrets masked.



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
	cloudSettingsReq := *openapiclient.NewCloudSettingsReq() // CloudSettingsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.CloudPutV1SettingsProduct(context.Background(), product).CloudSettingsReq(cloudSettingsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.CloudPutV1SettingsProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1SettingsProduct`: CloudSettingsView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.CloudPutV1SettingsProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**product** | **string** | Product is the catalog slug, from the PATH. zip binds the path last, so the URL names the product being written whatever a body field claims. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1SettingsProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSettingsReq** | [**CloudSettingsReq**](CloudSettingsReq.md) |  | 

### Return type

[**CloudSettingsView**](CloudSettingsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetSettings

> SearchSettings SearchGetSettings(ctx, indexUid).Execute()

Get all index settings

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
	indexUid := "indexUid_example" // string | Unique index identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SearchGetSettings(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SearchGetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetSettings`: SearchSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SearchGetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchSettings**](SearchSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchResetSettings

> SearchSummarizedTaskView SearchResetSettings(ctx, indexUid).Execute()

Reset all settings to defaults

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
	indexUid := "indexUid_example" // string | Unique index identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SearchResetSettings(context.Background(), indexUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SearchResetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchResetSettings`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SearchResetSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchResetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateSettings

> SearchSummarizedTaskView SearchUpdateSettings(ctx, indexUid).SearchSettings(searchSettings).Execute()

Update index settings

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
	indexUid := "indexUid_example" // string | Unique index identifier
	searchSettings := *openapiclient.NewSearchSettings() // SearchSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.SearchUpdateSettings(context.Background(), indexUid).SearchSettings(searchSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.SearchUpdateSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateSettings`: SearchSummarizedTaskView
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.SearchUpdateSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indexUid** | **string** | Unique index identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchSettings** | [**SearchSettings**](SearchSettings.md) |  | 

### Return type

[**SearchSummarizedTaskView**](SearchSummarizedTaskView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

