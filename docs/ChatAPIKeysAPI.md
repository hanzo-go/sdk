# \ChatAPIKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDeleteApiKeysByid**](ChatAPIKeysAPI.md#ChatDeleteApiKeysByid) | **Delete** /v1/chat/api-keys/{id} | Delete an API key
[**ChatGetApiKeys**](ChatAPIKeysAPI.md#ChatGetApiKeys) | **Get** /v1/chat/api-keys | List agent API keys
[**ChatGetApiKeysByid**](ChatAPIKeysAPI.md#ChatGetApiKeysByid) | **Get** /v1/chat/api-keys/{id} | Get an API key by ID
[**ChatPostApiKeys**](ChatAPIKeysAPI.md#ChatPostApiKeys) | **Post** /v1/chat/api-keys | Create an agent API key



## ChatDeleteApiKeysByid

> map[string]interface{} ChatDeleteApiKeysByid(ctx, id).Execute()

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIKeysAPI.ChatDeleteApiKeysByid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIKeysAPI.ChatDeleteApiKeysByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatDeleteApiKeysByid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIKeysAPI.ChatDeleteApiKeysByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatDeleteApiKeysByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ChatGetApiKeys

> []ChatAgentApiKey ChatGetApiKeys(ctx).Execute()

List agent API keys

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
	resp, r, err := apiClient.ChatAPIKeysAPI.ChatGetApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIKeysAPI.ChatGetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetApiKeys`: []ChatAgentApiKey
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIKeysAPI.ChatGetApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetApiKeysRequest struct via the builder pattern


### Return type

[**[]ChatAgentApiKey**](ChatAgentApiKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChatGetApiKeysByid

> map[string]interface{} ChatGetApiKeysByid(ctx, id).Execute()

Get an API key by ID

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
	resp, r, err := apiClient.ChatAPIKeysAPI.ChatGetApiKeysByid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIKeysAPI.ChatGetApiKeysByid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatGetApiKeysByid`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIKeysAPI.ChatGetApiKeysByid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChatGetApiKeysByidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ChatPostApiKeys

> ChatAgentApiKey ChatPostApiKeys(ctx).AutoCreateTableRequest(autoCreateTableRequest).Execute()

Create an agent API key

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
	autoCreateTableRequest := *openapiclient.NewAutoCreateTableRequest("Name_example") // AutoCreateTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChatAPIKeysAPI.ChatPostApiKeys(context.Background()).AutoCreateTableRequest(autoCreateTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChatAPIKeysAPI.ChatPostApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChatPostApiKeys`: ChatAgentApiKey
	fmt.Fprintf(os.Stdout, "Response from `ChatAPIKeysAPI.ChatPostApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChatPostApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateTableRequest** | [**AutoCreateTableRequest**](AutoCreateTableRequest.md) |  | 

### Return type

[**ChatAgentApiKey**](ChatAgentApiKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

