# \KmsKMSAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateKmsKey**](KmsKMSAPI.md#KmsCreateKmsKey) | **Post** /v1/kms/kms/keys | Create a KMS encryption key
[**KmsDecryptData**](KmsKMSAPI.md#KmsDecryptData) | **Post** /v1/kms/kms/keys/{keyId}/decrypt | Decrypt data with a KMS key
[**KmsDeleteKmsKey**](KmsKMSAPI.md#KmsDeleteKmsKey) | **Delete** /v1/kms/kms/keys/{keyId} | Delete a KMS key
[**KmsEncryptData**](KmsKMSAPI.md#KmsEncryptData) | **Post** /v1/kms/kms/keys/{keyId}/encrypt | Encrypt data with a KMS key
[**KmsListKmsKeys**](KmsKMSAPI.md#KmsListKmsKeys) | **Get** /v1/kms/kms/keys | List KMS encryption keys
[**KmsUpdateKmsKey**](KmsKMSAPI.md#KmsUpdateKmsKey) | **Patch** /v1/kms/kms/keys/{keyId} | Update a KMS key



## KmsCreateKmsKey

> KmsCreateKmsKey200Response KmsCreateKmsKey(ctx).KmsCreateKmsKeyRequest(kmsCreateKmsKeyRequest).Execute()

Create a KMS encryption key

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
	kmsCreateKmsKeyRequest := *openapiclient.NewKmsCreateKmsKeyRequest("Name_example", "ProjectId_example") // KmsCreateKmsKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsCreateKmsKey(context.Background()).KmsCreateKmsKeyRequest(kmsCreateKmsKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsCreateKmsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateKmsKey`: KmsCreateKmsKey200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsCreateKmsKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateKmsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateKmsKeyRequest** | [**KmsCreateKmsKeyRequest**](KmsCreateKmsKeyRequest.md) |  | 

### Return type

[**KmsCreateKmsKey200Response**](KmsCreateKmsKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDecryptData

> KmsDecryptDataResponse KmsDecryptData(ctx, keyId).KmsDecryptDataRequest(kmsDecryptDataRequest).Execute()

Decrypt data with a KMS key

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsDecryptDataRequest := *openapiclient.NewKmsDecryptDataRequest("Ciphertext_example") // KmsDecryptDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsDecryptData(context.Background(), keyId).KmsDecryptDataRequest(kmsDecryptDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsDecryptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDecryptData`: KmsDecryptDataResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsDecryptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDecryptDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsDecryptDataRequest** | [**KmsDecryptDataRequest**](KmsDecryptDataRequest.md) |  | 

### Return type

[**KmsDecryptDataResponse**](KmsDecryptDataResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteKmsKey

> map[string]interface{} KmsDeleteKmsKey(ctx, keyId).Execute()

Delete a KMS key

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsDeleteKmsKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsDeleteKmsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteKmsKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsDeleteKmsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteKmsKeyRequest struct via the builder pattern


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


## KmsEncryptData

> KmsEncryptDataResponse KmsEncryptData(ctx, keyId).KmsEncryptDataRequest(kmsEncryptDataRequest).Execute()

Encrypt data with a KMS key

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsEncryptDataRequest := *openapiclient.NewKmsEncryptDataRequest("Plaintext_example") // KmsEncryptDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsEncryptData(context.Background(), keyId).KmsEncryptDataRequest(kmsEncryptDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsEncryptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsEncryptData`: KmsEncryptDataResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsEncryptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsEncryptDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsEncryptDataRequest** | [**KmsEncryptDataRequest**](KmsEncryptDataRequest.md) |  | 

### Return type

[**KmsEncryptDataResponse**](KmsEncryptDataResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListKmsKeys

> KmsListKmsKeys200Response KmsListKmsKeys(ctx).ProjectId(projectId).Execute()

List KMS encryption keys

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsListKmsKeys(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsListKmsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListKmsKeys`: KmsListKmsKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsListKmsKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListKmsKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 

### Return type

[**KmsListKmsKeys200Response**](KmsListKmsKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateKmsKey

> KmsCreateKmsKey200Response KmsUpdateKmsKey(ctx, keyId).KmsUpdateKmsKeyRequest(kmsUpdateKmsKeyRequest).Execute()

Update a KMS key

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
	keyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateKmsKeyRequest := *openapiclient.NewKmsUpdateKmsKeyRequest() // KmsUpdateKmsKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsKMSAPI.KmsUpdateKmsKey(context.Background(), keyId).KmsUpdateKmsKeyRequest(kmsUpdateKmsKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsKMSAPI.KmsUpdateKmsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateKmsKey`: KmsCreateKmsKey200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsKMSAPI.KmsUpdateKmsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateKmsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateKmsKeyRequest** | [**KmsUpdateKmsKeyRequest**](KmsUpdateKmsKeyRequest.md) |  | 

### Return type

[**KmsCreateKmsKey200Response**](KmsCreateKmsKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

