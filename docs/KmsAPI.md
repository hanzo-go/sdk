# \KmsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKmsConfig**](KmsAPI.md#GetKmsConfig) | **Get** /v1/kms/config | Returns the runtime configuration for the KMS console.
[**GetKmsHealth**](KmsAPI.md#GetKmsHealth) | **Get** /v1/kms/health | Reports whether this broker can actually serve secrets.
[**GetKmsSecrets**](KmsAPI.md#GetKmsSecrets) | **Get** /v1/kms/secrets | Lists the secrets your org holds, without their values.
[**PostKmsAuthLogin**](KmsAPI.md#PostKmsAuthLogin) | **Post** /v1/kms/auth/login | Exchanges a machine credential for an IAM bearer token.
[**PostKmsSecrets**](KmsAPI.md#PostKmsSecrets) | **Post** /v1/kms/secrets | Stores or replaces one secret in your org.



## GetKmsConfig

> KmsConfig GetKmsConfig(ctx).Execute()

Returns the runtime configuration for the KMS console.



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
	resp, r, err := apiClient.KmsAPI.GetKmsConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsConfig`: KmsConfig
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsConfigRequest struct via the builder pattern


### Return type

[**KmsConfig**](KmsConfig.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsHealth

> KmsHealth GetKmsHealth(ctx).Execute()

Reports whether this broker can actually serve secrets.



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
	resp, r, err := apiClient.KmsAPI.GetKmsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsHealth`: KmsHealth
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsHealthRequest struct via the builder pattern


### Return type

[**KmsHealth**](KmsHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsSecrets

> KmsSecrets GetKmsSecrets(ctx).Env(env).Environment(environment).Path(path).SecretPath(secretPath).Execute()

Lists the secrets your org holds, without their values.



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
	env := "env_example" // string |  (optional)
	environment := "environment_example" // string |  (optional)
	path := "path_example" // string |  (optional)
	secretPath := "secretPath_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.GetKmsSecrets(context.Background()).Env(env).Environment(environment).Path(path).SecretPath(secretPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsSecrets`: KmsSecrets
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** |  | 
 **environment** | **string** |  | 
 **path** | **string** |  | 
 **secretPath** | **string** |  | 

### Return type

[**KmsSecrets**](KmsSecrets.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKmsAuthLogin

> KmsToken PostKmsAuthLogin(ctx).KmsLogin(kmsLogin).Execute()

Exchanges a machine credential for an IAM bearer token.



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
	kmsLogin := *openapiclient.NewKmsLogin() // KmsLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.PostKmsAuthLogin(context.Background()).KmsLogin(kmsLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.PostKmsAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKmsAuthLogin`: KmsToken
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.PostKmsAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKmsAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsLogin** | [**KmsLogin**](KmsLogin.md) |  | 

### Return type

[**KmsToken**](KmsToken.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKmsSecrets

> KmsStored PostKmsSecrets(ctx).KmsPut(kmsPut).Execute()

Stores or replaces one secret in your org.



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
	kmsPut := *openapiclient.NewKmsPut() // KmsPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.PostKmsSecrets(context.Background()).KmsPut(kmsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.PostKmsSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKmsSecrets`: KmsStored
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.PostKmsSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKmsSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsPut** | [**KmsPut**](KmsPut.md) |  | 

### Return type

[**KmsStored**](KmsStored.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

