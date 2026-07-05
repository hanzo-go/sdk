# \KmsEnvironmentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateEnvironment**](KmsEnvironmentsAPI.md#KmsCreateEnvironment) | **Post** /v1/kms/projects/{projectId}/environments | Create an environment
[**KmsDeleteEnvironment**](KmsEnvironmentsAPI.md#KmsDeleteEnvironment) | **Delete** /v1/kms/projects/{projectId}/environments/{envId} | Delete an environment
[**KmsListEnvironments**](KmsEnvironmentsAPI.md#KmsListEnvironments) | **Get** /v1/kms/projects/{projectId}/environments | List project environments
[**KmsUpdateEnvironment**](KmsEnvironmentsAPI.md#KmsUpdateEnvironment) | **Patch** /v1/kms/projects/{projectId}/environments/{envId} | Update an environment



## KmsCreateEnvironment

> KmsCreateEnvironment200Response KmsCreateEnvironment(ctx, projectId).KmsCreateEnvironmentRequest(kmsCreateEnvironmentRequest).Execute()

Create an environment

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
	kmsCreateEnvironmentRequest := *openapiclient.NewKmsCreateEnvironmentRequest("Name_example", "Slug_example") // KmsCreateEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsEnvironmentsAPI.KmsCreateEnvironment(context.Background(), projectId).KmsCreateEnvironmentRequest(kmsCreateEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsEnvironmentsAPI.KmsCreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateEnvironment`: KmsCreateEnvironment200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsEnvironmentsAPI.KmsCreateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsCreateEnvironmentRequest** | [**KmsCreateEnvironmentRequest**](KmsCreateEnvironmentRequest.md) |  | 

### Return type

[**KmsCreateEnvironment200Response**](KmsCreateEnvironment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteEnvironment

> map[string]interface{} KmsDeleteEnvironment(ctx, projectId, envId).Execute()

Delete an environment

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
	envId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsEnvironmentsAPI.KmsDeleteEnvironment(context.Background(), projectId, envId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsEnvironmentsAPI.KmsDeleteEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteEnvironment`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsEnvironmentsAPI.KmsDeleteEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteEnvironmentRequest struct via the builder pattern


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


## KmsListEnvironments

> KmsListEnvironments200Response KmsListEnvironments(ctx, projectId).Execute()

List project environments

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
	resp, r, err := apiClient.KmsEnvironmentsAPI.KmsListEnvironments(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsEnvironmentsAPI.KmsListEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListEnvironments`: KmsListEnvironments200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsEnvironmentsAPI.KmsListEnvironments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsListEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsListEnvironments200Response**](KmsListEnvironments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateEnvironment

> KmsCreateEnvironment200Response KmsUpdateEnvironment(ctx, projectId, envId).KmsUpdateEnvironmentRequest(kmsUpdateEnvironmentRequest).Execute()

Update an environment

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
	envId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateEnvironmentRequest := *openapiclient.NewKmsUpdateEnvironmentRequest() // KmsUpdateEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsEnvironmentsAPI.KmsUpdateEnvironment(context.Background(), projectId, envId).KmsUpdateEnvironmentRequest(kmsUpdateEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsEnvironmentsAPI.KmsUpdateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateEnvironment`: KmsCreateEnvironment200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsEnvironmentsAPI.KmsUpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**envId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kmsUpdateEnvironmentRequest** | [**KmsUpdateEnvironmentRequest**](KmsUpdateEnvironmentRequest.md) |  | 

### Return type

[**KmsCreateEnvironment200Response**](KmsCreateEnvironment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

