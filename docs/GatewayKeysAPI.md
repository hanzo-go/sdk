# \GatewayKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayDeleteKey**](GatewayKeysAPI.md#GatewayDeleteKey) | **Post** /v1/gateway/key/delete | Delete key
[**GatewayGenerateKey**](GatewayKeysAPI.md#GatewayGenerateKey) | **Post** /v1/gateway/key/generate | Generate API key
[**GatewayGetKeyInfo**](GatewayKeysAPI.md#GatewayGetKeyInfo) | **Get** /v1/gateway/key/info | Get key info
[**GatewayUpdateKey**](GatewayKeysAPI.md#GatewayUpdateKey) | **Post** /v1/gateway/key/update | Update key



## GatewayDeleteKey

> map[string]interface{} GatewayDeleteKey(ctx).GatewayDeleteKeyRequest(gatewayDeleteKeyRequest).Execute()

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
	gatewayDeleteKeyRequest := *openapiclient.NewGatewayDeleteKeyRequest([]string{"Keys_example"}) // GatewayDeleteKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayKeysAPI.GatewayDeleteKey(context.Background()).GatewayDeleteKeyRequest(gatewayDeleteKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayKeysAPI.GatewayDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayDeleteKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GatewayKeysAPI.GatewayDeleteKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayDeleteKeyRequest** | [**GatewayDeleteKeyRequest**](GatewayDeleteKeyRequest.md) |  | 

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


## GatewayGenerateKey

> GatewayKey GatewayGenerateKey(ctx).GatewayGenerateKeyRequest(gatewayGenerateKeyRequest).Execute()

Generate API key

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
	gatewayGenerateKeyRequest := *openapiclient.NewGatewayGenerateKeyRequest() // GatewayGenerateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayKeysAPI.GatewayGenerateKey(context.Background()).GatewayGenerateKeyRequest(gatewayGenerateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayKeysAPI.GatewayGenerateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGenerateKey`: GatewayKey
	fmt.Fprintf(os.Stdout, "Response from `GatewayKeysAPI.GatewayGenerateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGenerateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayGenerateKeyRequest** | [**GatewayGenerateKeyRequest**](GatewayGenerateKeyRequest.md) |  | 

### Return type

[**GatewayKey**](GatewayKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetKeyInfo

> GatewayKey GatewayGetKeyInfo(ctx).Key(key).Execute()

Get key info

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
	key := "key_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayKeysAPI.GatewayGetKeyInfo(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayKeysAPI.GatewayGetKeyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetKeyInfo`: GatewayKey
	fmt.Fprintf(os.Stdout, "Response from `GatewayKeysAPI.GatewayGetKeyInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetKeyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**GatewayKey**](GatewayKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateKey

> map[string]interface{} GatewayUpdateKey(ctx).GatewayUpdateKeyRequest(gatewayUpdateKeyRequest).Execute()

Update key

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
	gatewayUpdateKeyRequest := *openapiclient.NewGatewayUpdateKeyRequest("Key_example") // GatewayUpdateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayKeysAPI.GatewayUpdateKey(context.Background()).GatewayUpdateKeyRequest(gatewayUpdateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayKeysAPI.GatewayUpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayUpdateKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GatewayKeysAPI.GatewayUpdateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayUpdateKeyRequest** | [**GatewayUpdateKeyRequest**](GatewayUpdateKeyRequest.md) |  | 

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

