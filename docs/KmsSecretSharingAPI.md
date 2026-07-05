# \KmsSecretSharingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSharedSecret**](KmsSecretSharingAPI.md#KmsCreateSharedSecret) | **Post** /v1/kms/secret-sharing/shared | Create a shared secret
[**KmsDeleteSharedSecret**](KmsSecretSharingAPI.md#KmsDeleteSharedSecret) | **Delete** /v1/kms/secret-sharing/shared/{sharedSecretId} | Delete a shared secret
[**KmsGetSharedSecret**](KmsSecretSharingAPI.md#KmsGetSharedSecret) | **Get** /v1/kms/secret-sharing/shared/{sharedSecretId} | Get a shared secret by ID (consumes a view)
[**KmsListSharedSecrets**](KmsSecretSharingAPI.md#KmsListSharedSecrets) | **Get** /v1/kms/secret-sharing/shared | List shared secrets created by the user



## KmsCreateSharedSecret

> KmsCreateSharedSecret200Response KmsCreateSharedSecret(ctx).KmsCreateSharedSecretRequest(kmsCreateSharedSecretRequest).Execute()

Create a shared secret

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
	kmsCreateSharedSecretRequest := *openapiclient.NewKmsCreateSharedSecretRequest("SecretValue_example") // KmsCreateSharedSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSharingAPI.KmsCreateSharedSecret(context.Background()).KmsCreateSharedSecretRequest(kmsCreateSharedSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSharingAPI.KmsCreateSharedSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSharedSecret`: KmsCreateSharedSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSharingAPI.KmsCreateSharedSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSharedSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateSharedSecretRequest** | [**KmsCreateSharedSecretRequest**](KmsCreateSharedSecretRequest.md) |  | 

### Return type

[**KmsCreateSharedSecret200Response**](KmsCreateSharedSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSharedSecret

> map[string]interface{} KmsDeleteSharedSecret(ctx, sharedSecretId).Execute()

Delete a shared secret

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
	sharedSecretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSharingAPI.KmsDeleteSharedSecret(context.Background(), sharedSecretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSharingAPI.KmsDeleteSharedSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSharedSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSharingAPI.KmsDeleteSharedSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedSecretId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSharedSecretRequest struct via the builder pattern


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


## KmsGetSharedSecret

> KmsGetSharedSecret200Response KmsGetSharedSecret(ctx, sharedSecretId).Execute()

Get a shared secret by ID (consumes a view)

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
	sharedSecretId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSharingAPI.KmsGetSharedSecret(context.Background(), sharedSecretId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSharingAPI.KmsGetSharedSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetSharedSecret`: KmsGetSharedSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSharingAPI.KmsGetSharedSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sharedSecretId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetSharedSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsGetSharedSecret200Response**](KmsGetSharedSecret200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListSharedSecrets

> KmsListSharedSecrets200Response KmsListSharedSecrets(ctx).Execute()

List shared secrets created by the user

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
	resp, r, err := apiClient.KmsSecretSharingAPI.KmsListSharedSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSharingAPI.KmsListSharedSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSharedSecrets`: KmsListSharedSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSharingAPI.KmsListSharedSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSharedSecretsRequest struct via the builder pattern


### Return type

[**KmsListSharedSecrets200Response**](KmsListSharedSecrets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

