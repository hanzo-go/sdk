# \AutoApiKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateApiKey**](AutoApiKeysAPI.md#AutoCreateApiKey) | **Post** /v1/auto/api-keys | Create an API key (EE)
[**AutoDeleteApiKey**](AutoApiKeysAPI.md#AutoDeleteApiKey) | **Delete** /v1/auto/api-keys/{id} | Delete an API key (EE)
[**AutoListApiKeys**](AutoApiKeysAPI.md#AutoListApiKeys) | **Get** /v1/auto/api-keys | List API keys (EE)



## AutoCreateApiKey

> map[string]interface{} AutoCreateApiKey(ctx).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()

Create an API key (EE)

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
	autoCreateApiKeyRequest := *openapiclient.NewAutoCreateApiKeyRequest("DisplayName_example") // AutoCreateApiKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoApiKeysAPI.AutoCreateApiKey(context.Background()).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoApiKeysAPI.AutoCreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateApiKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoApiKeysAPI.AutoCreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateApiKeyRequest** | [**AutoCreateApiKeyRequest**](AutoCreateApiKeyRequest.md) |  | 

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


## AutoDeleteApiKey

> AutoDeleteApiKey(ctx, id).Execute()

Delete an API key (EE)

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
	r, err := apiClient.AutoApiKeysAPI.AutoDeleteApiKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoApiKeysAPI.AutoDeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoDeleteApiKeyRequest struct via the builder pattern


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


## AutoListApiKeys

> map[string]interface{} AutoListApiKeys(ctx).Execute()

List API keys (EE)

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
	resp, r, err := apiClient.AutoApiKeysAPI.AutoListApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoApiKeysAPI.AutoListApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListApiKeys`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoApiKeysAPI.AutoListApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListApiKeysRequest struct via the builder pattern


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

