# \KmsIdentityAuthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsAttachUniversalAuth**](KmsIdentityAuthAPI.md#KmsAttachUniversalAuth) | **Post** /v1/kms/auth/universal-auth/identities/{identityId} | Attach Universal Auth to an identity
[**KmsCreateIdentityToken**](KmsIdentityAuthAPI.md#KmsCreateIdentityToken) | **Post** /v1/kms/auth/token-auth/identities/{identityId}/tokens | Create an identity token
[**KmsCreateUniversalAuthClientSecret**](KmsIdentityAuthAPI.md#KmsCreateUniversalAuthClientSecret) | **Post** /v1/kms/auth/universal-auth/identities/{identityId}/client-secrets | Create a client secret for Universal Auth
[**KmsGetUniversalAuth**](KmsIdentityAuthAPI.md#KmsGetUniversalAuth) | **Get** /v1/kms/auth/universal-auth/identities/{identityId} | Get Universal Auth configuration for an identity
[**KmsUniversalAuthLogin**](KmsIdentityAuthAPI.md#KmsUniversalAuthLogin) | **Post** /v1/kms/auth/universal-auth/login | Login with Universal Auth



## KmsAttachUniversalAuth

> KmsUniversalAuthConfig KmsAttachUniversalAuth(ctx, identityId).KmsAttachUniversalAuthRequest(kmsAttachUniversalAuthRequest).Execute()

Attach Universal Auth to an identity

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsAttachUniversalAuthRequest := *openapiclient.NewKmsAttachUniversalAuthRequest() // KmsAttachUniversalAuthRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentityAuthAPI.KmsAttachUniversalAuth(context.Background(), identityId).KmsAttachUniversalAuthRequest(kmsAttachUniversalAuthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentityAuthAPI.KmsAttachUniversalAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsAttachUniversalAuth`: KmsUniversalAuthConfig
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentityAuthAPI.KmsAttachUniversalAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsAttachUniversalAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsAttachUniversalAuthRequest** | [**KmsAttachUniversalAuthRequest**](KmsAttachUniversalAuthRequest.md) |  | 

### Return type

[**KmsUniversalAuthConfig**](KmsUniversalAuthConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsCreateIdentityToken

> KmsTokenResponse KmsCreateIdentityToken(ctx, identityId).Execute()

Create an identity token

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentityAuthAPI.KmsCreateIdentityToken(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentityAuthAPI.KmsCreateIdentityToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateIdentityToken`: KmsTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentityAuthAPI.KmsCreateIdentityToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateIdentityTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsTokenResponse**](KmsTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsCreateUniversalAuthClientSecret

> KmsCreateClientSecretResponse KmsCreateUniversalAuthClientSecret(ctx, identityId).KmsCreateUniversalAuthClientSecretRequest(kmsCreateUniversalAuthClientSecretRequest).Execute()

Create a client secret for Universal Auth

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsCreateUniversalAuthClientSecretRequest := *openapiclient.NewKmsCreateUniversalAuthClientSecretRequest() // KmsCreateUniversalAuthClientSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentityAuthAPI.KmsCreateUniversalAuthClientSecret(context.Background(), identityId).KmsCreateUniversalAuthClientSecretRequest(kmsCreateUniversalAuthClientSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentityAuthAPI.KmsCreateUniversalAuthClientSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateUniversalAuthClientSecret`: KmsCreateClientSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentityAuthAPI.KmsCreateUniversalAuthClientSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateUniversalAuthClientSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsCreateUniversalAuthClientSecretRequest** | [**KmsCreateUniversalAuthClientSecretRequest**](KmsCreateUniversalAuthClientSecretRequest.md) |  | 

### Return type

[**KmsCreateClientSecretResponse**](KmsCreateClientSecretResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetUniversalAuth

> KmsUniversalAuthConfig KmsGetUniversalAuth(ctx, identityId).Execute()

Get Universal Auth configuration for an identity

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
	identityId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentityAuthAPI.KmsGetUniversalAuth(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentityAuthAPI.KmsGetUniversalAuth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetUniversalAuth`: KmsUniversalAuthConfig
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentityAuthAPI.KmsGetUniversalAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetUniversalAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsUniversalAuthConfig**](KmsUniversalAuthConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUniversalAuthLogin

> KmsTokenResponse KmsUniversalAuthLogin(ctx).KmsUniversalAuthLoginRequest(kmsUniversalAuthLoginRequest).Execute()

Login with Universal Auth



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
	kmsUniversalAuthLoginRequest := *openapiclient.NewKmsUniversalAuthLoginRequest("ClientId_example", "ClientSecret_example") // KmsUniversalAuthLoginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIdentityAuthAPI.KmsUniversalAuthLogin(context.Background()).KmsUniversalAuthLoginRequest(kmsUniversalAuthLoginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIdentityAuthAPI.KmsUniversalAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUniversalAuthLogin`: KmsTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `KmsIdentityAuthAPI.KmsUniversalAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsUniversalAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsUniversalAuthLoginRequest** | [**KmsUniversalAuthLoginRequest**](KmsUniversalAuthLoginRequest.md) |  | 

### Return type

[**KmsTokenResponse**](KmsTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

