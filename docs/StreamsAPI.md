# \StreamsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KvStreamAdd**](StreamsAPI.md#KvStreamAdd) | **Post** /v1/kv/streams/{key}/add | Add entry to stream
[**KvStreamInfo**](StreamsAPI.md#KvStreamInfo) | **Get** /v1/kv/streams/{key}/info | Get stream info
[**KvStreamRead**](StreamsAPI.md#KvStreamRead) | **Get** /v1/kv/streams/{key} | Read stream entries



## KvStreamAdd

> KvStreamAdd201Response KvStreamAdd(ctx, key).KvStreamAddRequest(kvStreamAddRequest).Namespace(namespace).Execute()

Add entry to stream

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
	kvStreamAddRequest := *openapiclient.NewKvStreamAddRequest(map[string]string{"key": "Inner_example"}) // KvStreamAddRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.KvStreamAdd(context.Background(), key).KvStreamAddRequest(kvStreamAddRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.KvStreamAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvStreamAdd`: KvStreamAdd201Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.KvStreamAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStreamAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvStreamAddRequest** | [**KvStreamAddRequest**](KvStreamAddRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

[**KvStreamAdd201Response**](KvStreamAdd201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvStreamInfo

> KvStreamInfo200Response KvStreamInfo(ctx, key).Namespace(namespace).Execute()

Get stream info

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
	resp, r, err := apiClient.StreamsAPI.KvStreamInfo(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.KvStreamInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvStreamInfo`: KvStreamInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.KvStreamInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStreamInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

[**KvStreamInfo200Response**](KvStreamInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvStreamRead

> KvStreamRead200Response KvStreamRead(ctx, key).Start(start).End(end).Count(count).Namespace(namespace).Execute()

Read stream entries

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
	start := "start_example" // string | Start ID (- for beginning) (optional) (default to "-")
	end := "end_example" // string | End ID (+ for latest) (optional) (default to "+")
	count := int32(56) // int32 |  (optional) (default to 100)
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamsAPI.KvStreamRead(context.Background(), key).Start(start).End(end).Count(count).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamsAPI.KvStreamRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvStreamRead`: KvStreamRead200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamsAPI.KvStreamRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvStreamReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **string** | Start ID (- for beginning) | [default to &quot;-&quot;]
 **end** | **string** | End ID (+ for latest) | [default to &quot;+&quot;]
 **count** | **int32** |  | [default to 100]
 **namespace** | **string** |  | 

### Return type

[**KvStreamRead200Response**](KvStreamRead200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

