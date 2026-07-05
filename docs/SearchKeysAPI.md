# \SearchKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchCreateKey**](SearchKeysAPI.md#SearchCreateKey) | **Post** /v1/search/keys | Create an API key
[**SearchDeleteKey**](SearchKeysAPI.md#SearchDeleteKey) | **Delete** /v1/search/keys/{keyOrUid} | Delete an API key
[**SearchGetKey**](SearchKeysAPI.md#SearchGetKey) | **Get** /v1/search/keys/{keyOrUid} | Get an API key
[**SearchListKeys**](SearchKeysAPI.md#SearchListKeys) | **Get** /v1/search/keys | List API keys
[**SearchUpdateKey**](SearchKeysAPI.md#SearchUpdateKey) | **Patch** /v1/search/keys/{keyOrUid} | Update an API key



## SearchCreateKey

> SearchKeyView SearchCreateKey(ctx).SearchCreateApiKey(searchCreateApiKey).Execute()

Create an API key

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
	searchCreateApiKey := *openapiclient.NewSearchCreateApiKey([]string{"Actions_example"}, []string{"Indexes_example"}) // SearchCreateApiKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchKeysAPI.SearchCreateKey(context.Background()).SearchCreateApiKey(searchCreateApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchKeysAPI.SearchCreateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCreateKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `SearchKeysAPI.SearchCreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCreateApiKey** | [**SearchCreateApiKey**](SearchCreateApiKey.md) |  | 

### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteKey

> SearchDeleteKey(ctx, keyOrUid).Execute()

Delete an API key

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
	keyOrUid := "keyOrUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchKeysAPI.SearchDeleteKey(context.Background(), keyOrUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchKeysAPI.SearchDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteKeyRequest struct via the builder pattern


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


## SearchGetKey

> SearchKeyView SearchGetKey(ctx, keyOrUid).Execute()

Get an API key

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
	keyOrUid := "keyOrUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchKeysAPI.SearchGetKey(context.Background(), keyOrUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchKeysAPI.SearchGetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `SearchKeysAPI.SearchGetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListKeys

> SearchPaginatedKeys SearchListKeys(ctx).Offset(offset).Limit(limit).Execute()

List API keys

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
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchKeysAPI.SearchListKeys(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchKeysAPI.SearchListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListKeys`: SearchPaginatedKeys
	fmt.Fprintf(os.Stdout, "Response from `SearchKeysAPI.SearchListKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**SearchPaginatedKeys**](SearchPaginatedKeys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateKey

> SearchKeyView SearchUpdateKey(ctx, keyOrUid).SearchUpdateKeyRequest(searchUpdateKeyRequest).Execute()

Update an API key

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
	keyOrUid := "keyOrUid_example" // string | 
	searchUpdateKeyRequest := *openapiclient.NewSearchUpdateKeyRequest() // SearchUpdateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchKeysAPI.SearchUpdateKey(context.Background(), keyOrUid).SearchUpdateKeyRequest(searchUpdateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchKeysAPI.SearchUpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `SearchKeysAPI.SearchUpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchUpdateKeyRequest** | [**SearchUpdateKeyRequest**](SearchUpdateKeyRequest.md) |  | 

### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

