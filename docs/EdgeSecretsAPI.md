# \EdgeSecretsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeCreateSecret**](EdgeSecretsAPI.md#EdgeCreateSecret) | **Post** /v1/edge/secrets | Create secret
[**EdgeDeleteSecret**](EdgeSecretsAPI.md#EdgeDeleteSecret) | **Delete** /v1/edge/secrets/{name} | Delete secret
[**EdgeListSecrets**](EdgeSecretsAPI.md#EdgeListSecrets) | **Get** /v1/edge/secrets | List secrets
[**EdgeUpdateSecret**](EdgeSecretsAPI.md#EdgeUpdateSecret) | **Put** /v1/edge/secrets/{name} | Update secret



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
	resp, r, err := apiClient.EdgeSecretsAPI.EdgeCreateSecret(context.Background()).EdgeSecretCreate(edgeSecretCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSecretsAPI.EdgeCreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeCreateSecret`: EdgeSecret
	fmt.Fprintf(os.Stdout, "Response from `EdgeSecretsAPI.EdgeCreateSecret`: %v\n", resp)
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
	resp, r, err := apiClient.EdgeSecretsAPI.EdgeDeleteSecret(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSecretsAPI.EdgeDeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeDeleteSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeSecretsAPI.EdgeDeleteSecret`: %v\n", resp)
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
	resp, r, err := apiClient.EdgeSecretsAPI.EdgeListSecrets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSecretsAPI.EdgeListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeListSecrets`: []EdgeListSecrets200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `EdgeSecretsAPI.EdgeListSecrets`: %v\n", resp)
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
	resp, r, err := apiClient.EdgeSecretsAPI.EdgeUpdateSecret(context.Background(), name).EdgeUpdateSecretRequest(edgeUpdateSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EdgeSecretsAPI.EdgeUpdateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EdgeUpdateSecret`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EdgeSecretsAPI.EdgeUpdateSecret`: %v\n", resp)
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

