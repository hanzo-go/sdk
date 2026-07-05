# \KvKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvBatchOperation**](KvKeysAPI.md#KvBatchOperation) | **Post** /v1/kv/batch | Batch get/set/delete
[**KvDeleteKey**](KvKeysAPI.md#KvDeleteKey) | **Delete** /v1/kv/keys/{key} | Delete key
[**KvGetKey**](KvKeysAPI.md#KvGetKey) | **Get** /v1/kv/keys/{key} | Get key value
[**KvIncrKey**](KvKeysAPI.md#KvIncrKey) | **Post** /v1/kv/keys/{key}/incr | Increment numeric key
[**KvScanKeys**](KvKeysAPI.md#KvScanKeys) | **Get** /v1/kv/keys | Scan keys
[**KvSetKey**](KvKeysAPI.md#KvSetKey) | **Put** /v1/kv/keys/{key} | Set key value
[**KvSetKeyTTL**](KvKeysAPI.md#KvSetKeyTTL) | **Put** /v1/kv/keys/{key}/ttl | Set key TTL



## KvBatchOperation

> KvBatchOperation200Response KvBatchOperation(ctx).KvBatchOperationRequest(kvBatchOperationRequest).Execute()

Batch get/set/delete

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
	kvBatchOperationRequest := *openapiclient.NewKvBatchOperationRequest() // KvBatchOperationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvBatchOperation(context.Background()).KvBatchOperationRequest(kvBatchOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvBatchOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvBatchOperation`: KvBatchOperation200Response
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvBatchOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvBatchOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvBatchOperationRequest** | [**KvBatchOperationRequest**](KvBatchOperationRequest.md) |  | 

### Return type

[**KvBatchOperation200Response**](KvBatchOperation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvDeleteKey

> KvDeleteKey200Response KvDeleteKey(ctx, key).Namespace(namespace).Execute()

Delete key

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvDeleteKey(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvDeleteKey`: KvDeleteKey200Response
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvDeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

[**KvDeleteKey200Response**](KvDeleteKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvGetKey

> KvKeyValue KvGetKey(ctx, key).Namespace(namespace).Execute()

Get key value

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvGetKey(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvGetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetKey`: KvKeyValue
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvGetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

[**KvKeyValue**](KvKeyValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvIncrKey

> AnalyticsGetSessionStats200ResponseValue KvIncrKey(ctx, key).Namespace(namespace).KvIncrKeyRequest(kvIncrKeyRequest).Execute()

Increment numeric key

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)
	kvIncrKeyRequest := *openapiclient.NewKvIncrKeyRequest() // KvIncrKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvIncrKey(context.Background(), key).Namespace(namespace).KvIncrKeyRequest(kvIncrKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvIncrKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvIncrKey`: AnalyticsGetSessionStats200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvIncrKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvIncrKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **kvIncrKeyRequest** | [**KvIncrKeyRequest**](KvIncrKeyRequest.md) |  | 

### Return type

[**AnalyticsGetSessionStats200ResponseValue**](AnalyticsGetSessionStats200ResponseValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvScanKeys

> KvScanKeys200Response KvScanKeys(ctx).Pattern(pattern).Type_(type_).Cursor(cursor).Count(count).Namespace(namespace).Execute()

Scan keys

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
	pattern := "pattern_example" // string | Glob-style pattern (e.g. user:*, session:*) (optional) (default to "*")
	type_ := "type__example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional) (default to "0")
	count := int32(56) // int32 |  (optional) (default to 100)
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvScanKeys(context.Background()).Pattern(pattern).Type_(type_).Cursor(cursor).Count(count).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvScanKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvScanKeys`: KvScanKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvScanKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvScanKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pattern** | **string** | Glob-style pattern (e.g. user:*, session:*) | [default to &quot;*&quot;]
 **type_** | **string** |  | 
 **cursor** | **string** |  | [default to &quot;0&quot;]
 **count** | **int32** |  | [default to 100]
 **namespace** | **string** |  | 

### Return type

[**KvScanKeys200Response**](KvScanKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSetKey

> KvKeyValue KvSetKey(ctx, key).KvSetKeyRequest(kvSetKeyRequest).Namespace(namespace).Execute()

Set key value

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
	key := "key_example" // string | 
	kvSetKeyRequest := *openapiclient.NewKvSetKeyRequest("Value_example") // KvSetKeyRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvSetKey(context.Background(), key).KvSetKeyRequest(kvSetKeyRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvSetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSetKey`: KvKeyValue
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvSetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvSetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvSetKeyRequest** | [**KvSetKeyRequest**](KvSetKeyRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

[**KvKeyValue**](KvKeyValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSetKeyTTL

> map[string]interface{} KvSetKeyTTL(ctx, key).KvSetKeyTTLRequest(kvSetKeyTTLRequest).Namespace(namespace).Execute()

Set key TTL

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
	key := "key_example" // string | 
	kvSetKeyTTLRequest := *openapiclient.NewKvSetKeyTTLRequest(int32(123)) // KvSetKeyTTLRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvKeysAPI.KvSetKeyTTL(context.Background(), key).KvSetKeyTTLRequest(kvSetKeyTTLRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvKeysAPI.KvSetKeyTTL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSetKeyTTL`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KvKeysAPI.KvSetKeyTTL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvSetKeyTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvSetKeyTTLRequest** | [**KvSetKeyTTLRequest**](KvSetKeyTTLRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

