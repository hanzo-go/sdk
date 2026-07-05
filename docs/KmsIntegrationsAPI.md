# \KmsIntegrationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateIntegration**](KmsIntegrationsAPI.md#KmsCreateIntegration) | **Post** /v1/kms/integration | Create an integration
[**KmsDeleteIntegration**](KmsIntegrationsAPI.md#KmsDeleteIntegration) | **Delete** /v1/kms/integration/{integrationId} | Delete an integration
[**KmsUpdateIntegration**](KmsIntegrationsAPI.md#KmsUpdateIntegration) | **Patch** /v1/kms/integration/{integrationId} | Update an integration



## KmsCreateIntegration

> KmsCreateIntegration200Response KmsCreateIntegration(ctx).KmsCreateIntegrationRequest(kmsCreateIntegrationRequest).Execute()

Create an integration

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
	kmsCreateIntegrationRequest := *openapiclient.NewKmsCreateIntegrationRequest("IntegrationAuthId_example", "Integration_example", "App_example", "SourceEnvironment_example", "SecretPath_example") // KmsCreateIntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIntegrationsAPI.KmsCreateIntegration(context.Background()).KmsCreateIntegrationRequest(kmsCreateIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIntegrationsAPI.KmsCreateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateIntegration`: KmsCreateIntegration200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsIntegrationsAPI.KmsCreateIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateIntegrationRequest** | [**KmsCreateIntegrationRequest**](KmsCreateIntegrationRequest.md) |  | 

### Return type

[**KmsCreateIntegration200Response**](KmsCreateIntegration200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteIntegration

> map[string]interface{} KmsDeleteIntegration(ctx, integrationId).Execute()

Delete an integration

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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIntegrationsAPI.KmsDeleteIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIntegrationsAPI.KmsDeleteIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteIntegration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsIntegrationsAPI.KmsDeleteIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteIntegrationRequest struct via the builder pattern


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


## KmsUpdateIntegration

> KmsCreateIntegration200Response KmsUpdateIntegration(ctx, integrationId).KmsUpdateIntegrationRequest(kmsUpdateIntegrationRequest).Execute()

Update an integration

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
	integrationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateIntegrationRequest := *openapiclient.NewKmsUpdateIntegrationRequest() // KmsUpdateIntegrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsIntegrationsAPI.KmsUpdateIntegration(context.Background(), integrationId).KmsUpdateIntegrationRequest(kmsUpdateIntegrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsIntegrationsAPI.KmsUpdateIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateIntegration`: KmsCreateIntegration200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsIntegrationsAPI.KmsUpdateIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateIntegrationRequest** | [**KmsUpdateIntegrationRequest**](KmsUpdateIntegrationRequest.md) |  | 

### Return type

[**KmsCreateIntegration200Response**](KmsCreateIntegration200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

