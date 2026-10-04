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

> KmsKmsConfig GetKmsConfig(ctx).Execute()

Returns the runtime configuration for the KMS console.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.GetKmsConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsConfig`: KmsKmsConfig
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsConfigRequest struct via the builder pattern


### Return type

[**KmsKmsConfig**](KmsKmsConfig.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsHealth

> KmsKmsHealth GetKmsHealth(ctx).Execute()

Reports whether this broker can actually serve secrets.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.GetKmsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsHealth`: KmsKmsHealth
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsHealthRequest struct via the builder pattern


### Return type

[**KmsKmsHealth**](KmsKmsHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsSecrets

> KmsKmsSecrets GetKmsSecrets(ctx).Env(env).Environment(environment).Path(path).SecretPath(secretPath).Execute()

Lists the secrets your org holds, without their values.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	env := "env_example" // string | Env selects the environment, which is part of a secret's storage key. OMITTED means EVERY environment — this is the enumeration surface, so it must be able to answer \"what is in here\" without being told where to look. (optional)
	environment := "environment_example" // string | Environment is the KMS operator's spelling of Env, accepted so one caller need not learn the other's vocabulary. Env wins when both are sent. (optional)
	path := "path_example" // string | Path narrows the listing to one subtree beneath the caller's org root, as a `/`-separated path such as `/ci`. OMITTED means the whole org. (optional)
	secretPath := "secretPath_example" // string | SecretPath is the KMS operator's spelling of Path. Path wins when both are sent. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.GetKmsSecrets(context.Background()).Env(env).Environment(environment).Path(path).SecretPath(secretPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsSecrets`: KmsKmsSecrets
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **env** | **string** | Env selects the environment, which is part of a secret&#39;s storage key. OMITTED means EVERY environment — this is the enumeration surface, so it must be able to answer \&quot;what is in here\&quot; without being told where to look. | 
 **environment** | **string** | Environment is the KMS operator&#39;s spelling of Env, accepted so one caller need not learn the other&#39;s vocabulary. Env wins when both are sent. | 
 **path** | **string** | Path narrows the listing to one subtree beneath the caller&#39;s org root, as a &#x60;/&#x60;-separated path such as &#x60;/ci&#x60;. OMITTED means the whole org. | 
 **secretPath** | **string** | SecretPath is the KMS operator&#39;s spelling of Path. Path wins when both are sent. | 

### Return type

[**KmsKmsSecrets**](KmsKmsSecrets.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKmsAuthLogin

> KmsKmsToken PostKmsAuthLogin(ctx).KmsKmsLogin(kmsKmsLogin).Execute()

Exchanges a machine credential for an IAM bearer token.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	kmsKmsLogin := *openapiclient.NewKmsKmsLogin() // KmsKmsLogin | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.PostKmsAuthLogin(context.Background()).KmsKmsLogin(kmsKmsLogin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.PostKmsAuthLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKmsAuthLogin`: KmsKmsToken
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.PostKmsAuthLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKmsAuthLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsKmsLogin** | [**KmsKmsLogin**](KmsKmsLogin.md) |  | 

### Return type

[**KmsKmsToken**](KmsKmsToken.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKmsSecrets

> KmsKmsStored PostKmsSecrets(ctx).KmsKmsPut(kmsKmsPut).Execute()

Stores or replaces one secret in your org.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	kmsKmsPut := *openapiclient.NewKmsKmsPut() // KmsKmsPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.PostKmsSecrets(context.Background()).KmsKmsPut(kmsKmsPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.PostKmsSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKmsSecrets`: KmsKmsStored
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.PostKmsSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKmsSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsKmsPut** | [**KmsKmsPut**](KmsKmsPut.md) |  | 

### Return type

[**KmsKmsStored**](KmsKmsStored.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

