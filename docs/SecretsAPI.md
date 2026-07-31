# \SecretsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeCreateSecret**](SecretsAPI.md#EdgeCreateSecret) | **Post** /v1/edge/secrets | Create secret
[**EdgeDeleteSecret**](SecretsAPI.md#EdgeDeleteSecret) | **Delete** /v1/edge/secrets/{name} | Delete secret
[**EdgeListSecrets**](SecretsAPI.md#EdgeListSecrets) | **Get** /v1/edge/secrets | List secrets
[**EdgeUpdateSecret**](SecretsAPI.md#EdgeUpdateSecret) | **Put** /v1/edge/secrets/{name} | Update secret
[**KmsDeleteV1KmsSecretsRest**](SecretsAPI.md#KmsDeleteV1KmsSecretsRest) | **Delete** /v1/kms/secrets/{rest} | Delete a secret
[**KmsGetV1KmsSecretsRest**](SecretsAPI.md#KmsGetV1KmsSecretsRest) | **Get** /v1/kms/secrets/{rest} | Read one secret value



## EdgeCreateSecret

> EdgeSecret EdgeCreateSecret(ctx).EdgeSecretCreate(edgeSecretCreate).Execute()

Create secret

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
	edgeSecretCreate := *openapiclient.NewEdgeSecretCreate("Name_example", "Value_example") // EdgeSecretCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.EdgeCreateSecret(context.Background()).EdgeSecretCreate(edgeSecretCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.EdgeCreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeCreateSecret`: EdgeSecret
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.EdgeCreateSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEdgeCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgeSecretCreate** | [**EdgeSecretCreate**](EdgeSecretCreate.md) |  | 

### Return type

[**EdgeSecret**](EdgeSecret.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeDeleteSecret

> map[string]interface{} EdgeDeleteSecret(ctx, name).Execute()

Delete secret

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.EdgeDeleteSecret(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.EdgeDeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeleteSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.EdgeDeleteSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeDeleteSecretRequest struct via the builder pattern


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


## EdgeListSecrets

> []EdgeListSecrets200ResponseInner EdgeListSecrets(ctx).Execute()

List secrets



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
	resp, r, err := apiClient.SecretsAPI.EdgeListSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.EdgeListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeListSecrets`: []EdgeListSecrets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.EdgeListSecrets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeListSecretsRequest struct via the builder pattern


### Return type

[**[]EdgeListSecrets200ResponseInner**](EdgeListSecrets200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EdgeUpdateSecret

> map[string]interface{} EdgeUpdateSecret(ctx, name).EdgeUpdateSecretRequest(edgeUpdateSecretRequest).Execute()

Update secret

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
	name := "name_example" // string | 
	edgeUpdateSecretRequest := *openapiclient.NewEdgeUpdateSecretRequest("Value_example") // EdgeUpdateSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.EdgeUpdateSecret(context.Background(), name).EdgeUpdateSecretRequest(edgeUpdateSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.EdgeUpdateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeUpdateSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.EdgeUpdateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEdgeUpdateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **edgeUpdateSecretRequest** | [**EdgeUpdateSecretRequest**](EdgeUpdateSecretRequest.md) |  | 

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


## KmsDeleteV1KmsSecretsRest

> KmsDeleteV1KmsSecretsRest(ctx, rest).Env(env).Execute()

Delete a secret

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
	rest := "rest_example" // string | Path and name joined. Split at the LAST slash: `a/b/C` is path `a/b`, name `C`. Escape each segment individually.
	env := "env_example" // string | Environment slug; defaults to `default` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretsAPI.KmsDeleteV1KmsSecretsRest(context.Background(), rest).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.KmsDeleteV1KmsSecretsRest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rest** | **string** | Path and name joined. Split at the LAST slash: &#x60;a/b/C&#x60; is path &#x60;a/b&#x60;, name &#x60;C&#x60;. Escape each segment individually. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteV1KmsSecretsRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Environment slug; defaults to &#x60;default&#x60; | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsSecretsRest

> KmsGetV1KmsSecretsRest200Response KmsGetV1KmsSecretsRest(ctx, rest).Env(env).Execute()

Read one secret value

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
	rest := "rest_example" // string | Path and name joined. Split at the LAST slash: `a/b/C` is path `a/b`, name `C`. Escape each segment individually.
	env := "env_example" // string | Environment slug; defaults to `default` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsAPI.KmsGetV1KmsSecretsRest(context.Background(), rest).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsAPI.KmsGetV1KmsSecretsRest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetV1KmsSecretsRest`: KmsGetV1KmsSecretsRest200Response
	fmt.Fprintf(os.Stdout, "Response from `SecretsAPI.KmsGetV1KmsSecretsRest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rest** | **string** | Path and name joined. Split at the LAST slash: &#x60;a/b/C&#x60; is path &#x60;a/b&#x60;, name &#x60;C&#x60;. Escape each segment individually. | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsSecretsRestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Environment slug; defaults to &#x60;default&#x60; | 

### Return type

[**KmsGetV1KmsSecretsRest200Response**](KmsGetV1KmsSecretsRest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

