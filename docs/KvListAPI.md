# \KvListAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvListPop**](KvListAPI.md#KvListPop) | **Post** /v1/kv/list/{key}/pop | Pop from list
[**KvListPush**](KvListAPI.md#KvListPush) | **Post** /v1/kv/list/{key}/push | Push to list
[**KvListRange**](KvListAPI.md#KvListRange) | **Get** /v1/kv/list/{key} | Get list range



## KvListPop

> KvListPop200Response KvListPop(ctx, key).Namespace(namespace).KvListPopRequest(kvListPopRequest).Execute()

Pop from list

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
	kvListPopRequest := *openapiclient.NewKvListPopRequest() // KvListPopRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvListAPI.KvListPop(context.Background(), key).Namespace(namespace).KvListPopRequest(kvListPopRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvListAPI.KvListPop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListPop`: KvListPop200Response
	fmt.Fprintf(os.Stdout, "Response from `KvListAPI.KvListPop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvListPopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **kvListPopRequest** | [**KvListPopRequest**](KvListPopRequest.md) |  | 

### Return type

[**KvListPop200Response**](KvListPop200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvListPush

> KvListPush200Response KvListPush(ctx, key).KvListPushRequest(kvListPushRequest).Namespace(namespace).Execute()

Push to list

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
	kvListPushRequest := *openapiclient.NewKvListPushRequest([]string{"Values_example"}) // KvListPushRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvListAPI.KvListPush(context.Background(), key).KvListPushRequest(kvListPushRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvListAPI.KvListPush``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListPush`: KvListPush200Response
	fmt.Fprintf(os.Stdout, "Response from `KvListAPI.KvListPush`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvListPushRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvListPushRequest** | [**KvListPushRequest**](KvListPushRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

[**KvListPush200Response**](KvListPush200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvListRange

> KvListRange200Response KvListRange(ctx, key).Start(start).Stop(stop).Namespace(namespace).Execute()

Get list range

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
	start := int32(56) // int32 |  (optional) (default to 0)
	stop := int32(56) // int32 |  (optional) (default to -1)
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KvListAPI.KvListRange(context.Background(), key).Start(start).Stop(stop).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KvListAPI.KvListRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvListRange`: KvListRange200Response
	fmt.Fprintf(os.Stdout, "Response from `KvListAPI.KvListRange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvListRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** |  | [default to 0]
 **stop** | **int32** |  | [default to -1]
 **namespace** | **string** |  | 

### Return type

[**KvListRange200Response**](KvListRange200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

