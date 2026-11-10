# \KmsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKmsSecretsByWildcard1**](KmsAPI.md#DeleteKmsSecretsByWildcard1) | **Delete** /v1/kms/secrets/{wildcard1} | Removes one secret from your org.
[**GetKmsConfig**](KmsAPI.md#GetKmsConfig) | **Get** /v1/kms/config | Returns the runtime configuration for the KMS console.
[**GetKmsHealth**](KmsAPI.md#GetKmsHealth) | **Get** /v1/kms/health | Reports whether this broker can actually serve secrets.
[**GetKmsSecrets**](KmsAPI.md#GetKmsSecrets) | **Get** /v1/kms/secrets | Lists the secrets your org holds, without their values.
[**GetKmsSecretsByWildcard1**](KmsAPI.md#GetKmsSecretsByWildcard1) | **Get** /v1/kms/secrets/{wildcard1} | Reads one secret&#39;s value from your org.
[**PostKmsAuthLogin**](KmsAPI.md#PostKmsAuthLogin) | **Post** /v1/kms/auth/login | Exchanges a machine credential for an IAM bearer token.
[**PostKmsSecrets**](KmsAPI.md#PostKmsSecrets) | **Post** /v1/kms/secrets | Stores or replaces one secret in your org.



## DeleteKmsSecretsByWildcard1

> KmsKmsRemoved DeleteKmsSecretsByWildcard1(ctx, wildcard1).Env(env).Execute()

Removes one secret from your org.



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
	wildcard1 := "wildcard1_example" // string | Secret is the coordinate beneath the caller's own org root: an optional `/`-separated subpath and then the name, such as `ci/deploy/token`. Over HTTP it is the trailing path itself, and the trailing path WINS over any other spelling sent with it. There is no org in it — the tenant comes from the validated claim — so another tenant's secret is not merely refused, it is unnameable. OMITTED is refused with a 400: there is no secret named \"everything\", and a blank address must not read as one.
	env := "prod" // string | Env selects the environment to resolve the secret in. It is part of the storage key, so one name in two environments is two secrets. OMITTED means the `default` environment — where a WRITE refuses to default, because a misplaced write strands a value no reader looks for, a read or a delete aimed at the wrong environment simply answers 404 and the caller learns. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.DeleteKmsSecretsByWildcard1(context.Background(), wildcard1).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.DeleteKmsSecretsByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKmsSecretsByWildcard1`: KmsKmsRemoved
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.DeleteKmsSecretsByWildcard1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** | Secret is the coordinate beneath the caller&#39;s own org root: an optional &#x60;/&#x60;-separated subpath and then the name, such as &#x60;ci/deploy/token&#x60;. Over HTTP it is the trailing path itself, and the trailing path WINS over any other spelling sent with it. There is no org in it — the tenant comes from the validated claim — so another tenant&#39;s secret is not merely refused, it is unnameable. OMITTED is refused with a 400: there is no secret named \&quot;everything\&quot;, and a blank address must not read as one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKmsSecretsByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Env selects the environment to resolve the secret in. It is part of the storage key, so one name in two environments is two secrets. OMITTED means the &#x60;default&#x60; environment — where a WRITE refuses to default, because a misplaced write strands a value no reader looks for, a read or a delete aimed at the wrong environment simply answers 404 and the caller learns. | 

### Return type

[**KmsKmsRemoved**](KmsKmsRemoved.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetKmsSecretsByWildcard1

> KmsKmsSecret GetKmsSecretsByWildcard1(ctx, wildcard1).Env(env).Execute()

Reads one secret's value from your org.



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
	wildcard1 := "wildcard1_example" // string | Secret is the coordinate beneath the caller's own org root: an optional `/`-separated subpath and then the name, such as `ci/deploy/token`. Over HTTP it is the trailing path itself, and the trailing path WINS over any other spelling sent with it. There is no org in it — the tenant comes from the validated claim — so another tenant's secret is not merely refused, it is unnameable. OMITTED is refused with a 400: there is no secret named \"everything\", and a blank address must not read as one.
	env := "prod" // string | Env selects the environment to resolve the secret in. It is part of the storage key, so one name in two environments is two secrets. OMITTED means the `default` environment — where a WRITE refuses to default, because a misplaced write strands a value no reader looks for, a read or a delete aimed at the wrong environment simply answers 404 and the caller learns. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.GetKmsSecretsByWildcard1(context.Background(), wildcard1).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.GetKmsSecretsByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsSecretsByWildcard1`: KmsKmsSecret
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.GetKmsSecretsByWildcard1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** | Secret is the coordinate beneath the caller&#39;s own org root: an optional &#x60;/&#x60;-separated subpath and then the name, such as &#x60;ci/deploy/token&#x60;. Over HTTP it is the trailing path itself, and the trailing path WINS over any other spelling sent with it. There is no org in it — the tenant comes from the validated claim — so another tenant&#39;s secret is not merely refused, it is unnameable. OMITTED is refused with a 400: there is no secret named \&quot;everything\&quot;, and a blank address must not read as one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsSecretsByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Env selects the environment to resolve the secret in. It is part of the storage key, so one name in two environments is two secrets. OMITTED means the &#x60;default&#x60; environment — where a WRITE refuses to default, because a misplaced write strands a value no reader looks for, a read or a delete aimed at the wrong environment simply answers 404 and the caller learns. | 

### Return type

[**KmsKmsSecret**](KmsKmsSecret.md)

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

