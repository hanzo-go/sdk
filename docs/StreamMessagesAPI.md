# \StreamMessagesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamConsumeMessages**](StreamMessagesAPI.md#StreamConsumeMessages) | **Get** /v1/stream/topics/{topic}/messages | Consume messages
[**StreamProduceMessages**](StreamMessagesAPI.md#StreamProduceMessages) | **Post** /v1/stream/topics/{topic}/messages | Produce messages



## StreamConsumeMessages

> StreamConsumeMessages200Response StreamConsumeMessages(ctx, topic).Partition(partition).Offset(offset).Limit(limit).Timeout(timeout).Execute()

Consume messages



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
	topic := "topic_example" // string | 
	partition := int32(56) // int32 | Partition to consume from (optional) (default to 0)
	offset := "offset_example" // string | Starting offset (earliest, latest, or numeric offset) (optional) (default to "latest")
	limit := int32(56) // int32 | Maximum records to return (optional) (default to 100)
	timeout := int32(56) // int32 | Long-poll timeout in milliseconds (optional) (default to 5000)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamMessagesAPI.StreamConsumeMessages(context.Background(), topic).Partition(partition).Offset(offset).Limit(limit).Timeout(timeout).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamMessagesAPI.StreamConsumeMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamConsumeMessages`: StreamConsumeMessages200Response
	fmt.Fprintf(os.Stdout, "Response from `StreamMessagesAPI.StreamConsumeMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamConsumeMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partition** | **int32** | Partition to consume from | [default to 0]
 **offset** | **string** | Starting offset (earliest, latest, or numeric offset) | [default to &quot;latest&quot;]
 **limit** | **int32** | Maximum records to return | [default to 100]
 **timeout** | **int32** | Long-poll timeout in milliseconds | [default to 5000]

### Return type

[**StreamConsumeMessages200Response**](StreamConsumeMessages200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StreamProduceMessages

> StreamProduceResponse StreamProduceMessages(ctx, topic).StreamProduceRequest(streamProduceRequest).Execute()

Produce messages



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
	topic := "topic_example" // string | 
	streamProduceRequest := *openapiclient.NewStreamProduceRequest([]openapiclient.StreamProduceRequestRecordsInner{*openapiclient.NewStreamProduceRequestRecordsInner()}) // StreamProduceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StreamMessagesAPI.StreamProduceMessages(context.Background(), topic).StreamProduceRequest(streamProduceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StreamMessagesAPI.StreamProduceMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StreamProduceMessages`: StreamProduceResponse
	fmt.Fprintf(os.Stdout, "Response from `StreamMessagesAPI.StreamProduceMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamProduceMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamProduceRequest** | [**StreamProduceRequest**](StreamProduceRequest.md) |  | 

### Return type

[**StreamProduceResponse**](StreamProduceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

