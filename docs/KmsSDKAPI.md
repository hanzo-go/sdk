# \KmsSDKAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsSdkSecretsOp**](KmsSDKAPI.md#KmsSdkSecretsOp) | **Post** /v1/sdk/secrets | Enveloped secret + threshold-key operation



## KmsSdkSecretsOp

> map[string]interface{} KmsSdkSecretsOp(ctx).KmsSdkEnvelope(kmsSdkEnvelope).Execute()

Enveloped secret + threshold-key operation



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
	kmsSdkEnvelope := *openapiclient.NewKmsSdkEnvelope(int32(1), *openapiclient.NewKmsSdkEnvelopeIdentity(int32(66), "Node_example", string(123), "Path_example", string(123)), int64(123), "Nonce_example", int32(64), map[string]interface{}(123), string(123)) // KmsSdkEnvelope | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSDKAPI.KmsSdkSecretsOp(context.Background()).KmsSdkEnvelope(kmsSdkEnvelope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSDKAPI.KmsSdkSecretsOp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsSdkSecretsOp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSDKAPI.KmsSdkSecretsOp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsSdkSecretsOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsSdkEnvelope** | [**KmsSdkEnvelope**](KmsSdkEnvelope.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

