# \MlDeploymentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MlDeployModel**](MlDeploymentsAPI.md#MlDeployModel) | **Post** /v1/ml/deploy | Deploy a model
[**MlGetDeployment**](MlDeploymentsAPI.md#MlGetDeployment) | **Get** /v1/ml/deployments/{deployment_id} | Get deployment details
[**MlListDeployments**](MlDeploymentsAPI.md#MlListDeployments) | **Get** /v1/ml/deployments | List deployments
[**MlStopDeployment**](MlDeploymentsAPI.md#MlStopDeployment) | **Delete** /v1/ml/deployments/{deployment_id} | Stop a deployment



## MlDeployModel

> MlDeployment MlDeployModel(ctx).MlDeployModelRequest(mlDeployModelRequest).Execute()

Deploy a model



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
	mlDeployModelRequest := *openapiclient.NewMlDeployModelRequest("ModelId_example") // MlDeployModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlDeploymentsAPI.MlDeployModel(context.Background()).MlDeployModelRequest(mlDeployModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlDeploymentsAPI.MlDeployModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlDeployModel`: MlDeployment
	fmt.Fprintf(os.Stdout, "Response from `MlDeploymentsAPI.MlDeployModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlDeployModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mlDeployModelRequest** | [**MlDeployModelRequest**](MlDeployModelRequest.md) |  | 

### Return type

[**MlDeployment**](MlDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlGetDeployment

> MlDeployment MlGetDeployment(ctx, deploymentId).Execute()

Get deployment details

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
	deploymentId := "deploymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlDeploymentsAPI.MlGetDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlDeploymentsAPI.MlGetDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetDeployment`: MlDeployment
	fmt.Fprintf(os.Stdout, "Response from `MlDeploymentsAPI.MlGetDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlGetDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlDeployment**](MlDeployment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlListDeployments

> MlListDeployments200Response MlListDeployments(ctx).Environment(environment).Status(status).Execute()

List deployments

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
	environment := "environment_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlDeploymentsAPI.MlListDeployments(context.Background()).Environment(environment).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlDeploymentsAPI.MlListDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListDeployments`: MlListDeployments200Response
	fmt.Fprintf(os.Stdout, "Response from `MlDeploymentsAPI.MlListDeployments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlListDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**MlListDeployments200Response**](MlListDeployments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlStopDeployment

> MlStopDeployment(ctx, deploymentId).Execute()

Stop a deployment

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
	deploymentId := "deploymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MlDeploymentsAPI.MlStopDeployment(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlDeploymentsAPI.MlStopDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlStopDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

