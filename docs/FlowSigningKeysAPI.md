# \FlowSigningKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowCreateSigningKey**](FlowSigningKeysAPI.md#FlowCreateSigningKey) | **Post** /v1/flow/signing-keys | Create a signing key (EE)
[**FlowListSigningKeys**](FlowSigningKeysAPI.md#FlowListSigningKeys) | **Get** /v1/flow/signing-keys | List signing keys (EE)



## FlowCreateSigningKey

> map[string]interface{} FlowCreateSigningKey(ctx).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()

Create a signing key (EE)

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
	resp, r, err := apiClient.FlowSigningKeysAPI.FlowCreateSigningKey(context.Background()).AutoCreateApiKeyRequest(autoCreateApiKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSigningKeysAPI.FlowCreateSigningKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateSigningKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSigningKeysAPI.FlowCreateSigningKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateSigningKeyRequest struct via the builder pattern


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


## FlowListSigningKeys

> map[string]interface{} FlowListSigningKeys(ctx).Execute()

List signing keys (EE)

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
	resp, r, err := apiClient.FlowSigningKeysAPI.FlowListSigningKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowSigningKeysAPI.FlowListSigningKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListSigningKeys`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowSigningKeysAPI.FlowListSigningKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListSigningKeysRequest struct via the builder pattern


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

