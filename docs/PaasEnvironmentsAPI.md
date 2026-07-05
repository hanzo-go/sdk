# \PaasEnvironmentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasCreateEnvironment**](PaasEnvironmentsAPI.md#PaasCreateEnvironment) | **Post** /v1/paas/org/{orgId}/project/{projectId}/env | Create environment
[**PaasDeleteEnvironment**](PaasEnvironmentsAPI.md#PaasDeleteEnvironment) | **Delete** /v1/paas/org/{orgId}/project/{projectId}/env/{envId} | Delete environment
[**PaasGetEnvironment**](PaasEnvironmentsAPI.md#PaasGetEnvironment) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env/{envId} | Get environment
[**PaasListEnvironments**](PaasEnvironmentsAPI.md#PaasListEnvironments) | **Get** /v1/paas/org/{orgId}/project/{projectId}/env | List environments



## PaasCreateEnvironment

> map[string]interface{} PaasCreateEnvironment(ctx, orgId, projectId).AutoCreateTableRequest(autoCreateTableRequest).Execute()

Create environment

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	autoCreateTableRequest := *openapiclient.NewAutoCreateTableRequest("Name_example") // AutoCreateTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasEnvironmentsAPI.PaasCreateEnvironment(context.Background(), orgId, projectId).AutoCreateTableRequest(autoCreateTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasEnvironmentsAPI.PaasCreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasCreateEnvironment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasEnvironmentsAPI.PaasCreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoCreateTableRequest** | [**AutoCreateTableRequest**](AutoCreateTableRequest.md) |  | 

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


## PaasDeleteEnvironment

> map[string]interface{} PaasDeleteEnvironment(ctx, orgId, projectId, envId).Execute()

Delete environment

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasEnvironmentsAPI.PaasDeleteEnvironment(context.Background(), orgId, projectId, envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasEnvironmentsAPI.PaasDeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasDeleteEnvironment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaasEnvironmentsAPI.PaasDeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasDeleteEnvironmentRequest struct via the builder pattern


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


## PaasGetEnvironment

> PaasEnvironment PaasGetEnvironment(ctx, orgId, projectId, envId).Execute()

Get environment

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 
	envId := "envId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasEnvironmentsAPI.PaasGetEnvironment(context.Background(), orgId, projectId, envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasEnvironmentsAPI.PaasGetEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasGetEnvironment`: PaasEnvironment
	fmt.Fprintf(os.Stdout, "Response from `PaasEnvironmentsAPI.PaasGetEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasGetEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PaasEnvironment**](PaasEnvironment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaasListEnvironments

> []PaasEnvironment PaasListEnvironments(ctx, orgId, projectId).Execute()

List environments

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
	orgId := "orgId_example" // string | 
	projectId := "projectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaasEnvironmentsAPI.PaasListEnvironments(context.Background(), orgId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasEnvironmentsAPI.PaasListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListEnvironments`: []PaasEnvironment
	fmt.Fprintf(os.Stdout, "Response from `PaasEnvironmentsAPI.PaasListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PaasEnvironment**](PaasEnvironment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

