# \GatewayEmbeddingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreateEmbedding**](GatewayEmbeddingsAPI.md#GatewayCreateEmbedding) | **Post** /v1/gateway/embeddings | Create embeddings



## GatewayCreateEmbedding

> GatewayEmbeddingResponse GatewayCreateEmbedding(ctx).GatewayEmbeddingRequest(gatewayEmbeddingRequest).Execute()

Create embeddings

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
	gatewayEmbeddingRequest := *openapiclient.NewGatewayEmbeddingRequest("Model_example", openapiclient.gateway_createCompletion_request_prompt{ArrayOfString: new([]string)}) // GatewayEmbeddingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayEmbeddingsAPI.GatewayCreateEmbedding(context.Background()).GatewayEmbeddingRequest(gatewayEmbeddingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayEmbeddingsAPI.GatewayCreateEmbedding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreateEmbedding`: GatewayEmbeddingResponse
	fmt.Fprintf(os.Stdout, "Response from `GatewayEmbeddingsAPI.GatewayCreateEmbedding`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateEmbeddingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayEmbeddingRequest** | [**GatewayEmbeddingRequest**](GatewayEmbeddingRequest.md) |  | 

### Return type

[**GatewayEmbeddingResponse**](GatewayEmbeddingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

