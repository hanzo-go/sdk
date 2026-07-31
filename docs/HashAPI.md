# \HashAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvHashDeleteField**](HashAPI.md#KvHashDeleteField) | **Delete** /v1/kv/hash/{key}/{field} | Delete hash field
[**KvHashGetAll**](HashAPI.md#KvHashGetAll) | **Get** /v1/kv/hash/{key} | Get all hash fields
[**KvHashGetField**](HashAPI.md#KvHashGetField) | **Get** /v1/kv/hash/{key}/{field} | Get hash field
[**KvHashSet**](HashAPI.md#KvHashSet) | **Put** /v1/kv/hash/{key} | Set hash fields



## KvHashDeleteField

> map[string]interface{} KvHashDeleteField(ctx, key, field).Namespace(namespace).Execute()

Delete hash field

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
	field := "field_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HashAPI.KvHashDeleteField(context.Background(), key, field).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HashAPI.KvHashDeleteField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvHashDeleteField`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HashAPI.KvHashDeleteField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**field** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvHashDeleteFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **string** |  | 

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


## KvHashGetAll

> map[string]string KvHashGetAll(ctx, key).Namespace(namespace).Execute()

Get all hash fields

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
	resp, r, err := apiClient.HashAPI.KvHashGetAll(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HashAPI.KvHashGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvHashGetAll`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `HashAPI.KvHashGetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvHashGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

**map[string]string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvHashGetField

> KmsGetV1KmsSecretsRest200ResponseSecret KvHashGetField(ctx, key, field).Namespace(namespace).Execute()

Get hash field

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
	field := "field_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HashAPI.KvHashGetField(context.Background(), key, field).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HashAPI.KvHashGetField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvHashGetField`: KmsGetV1KmsSecretsRest200ResponseSecret
	fmt.Fprintf(os.Stdout, "Response from `HashAPI.KvHashGetField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 
**field** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvHashGetFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **string** |  | 

### Return type

[**KmsGetV1KmsSecretsRest200ResponseSecret**](KmsGetV1KmsSecretsRest200ResponseSecret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvHashSet

> map[string]interface{} KvHashSet(ctx, key).RequestBody(requestBody).Namespace(namespace).Execute()

Set hash fields

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
	requestBody := map[string]string{"key": "Inner_example"} // map[string]string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HashAPI.KvHashSet(context.Background(), key).RequestBody(requestBody).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HashAPI.KvHashSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvHashSet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `HashAPI.KvHashSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvHashSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]string** |  | 
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

