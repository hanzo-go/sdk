# \ModelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiListModels**](ModelsAPI.md#AiListModels) | **Get** /v1/models | List available models
[**AiRetrieveModel**](ModelsAPI.md#AiRetrieveModel) | **Get** /v1/models/{model} | Retrieve a model
[**GatewayGetModel**](ModelsAPI.md#GatewayGetModel) | **Get** /v1/gateway/models/{model} | Get model
[**GatewayListModels**](ModelsAPI.md#GatewayListModels) | **Get** /v1/gateway/models | List models
[**MlGetModel**](ModelsAPI.md#MlGetModel) | **Get** /v1/ml/models/{model_id} | Get model details
[**MlPromoteModel**](ModelsAPI.md#MlPromoteModel) | **Post** /v1/ml/models/{model_id}/promote | Promote a model
[**MlRollbackModel**](ModelsAPI.md#MlRollbackModel) | **Post** /v1/ml/models/{model_id}/rollback | Rollback a model



## AiListModels

> AiModelList AiListModels(ctx).Execute()

List available models



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
	resp, r, err := apiClient.ModelsAPI.AiListModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.AiListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiListModels`: AiModelList
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.AiListModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiListModelsRequest struct via the builder pattern


### Return type

[**AiModelList**](AiModelList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiRetrieveModel

> AiModel AiRetrieveModel(ctx, model).Execute()

Retrieve a model

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
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.AiRetrieveModel(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.AiRetrieveModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiRetrieveModel`: AiModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.AiRetrieveModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAiRetrieveModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiModel**](AiModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetModel

> GatewayModel GatewayGetModel(ctx, model).Execute()

Get model

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
	model := "model_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.GatewayGetModel(context.Background(), model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GatewayGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetModel`: GatewayModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GatewayGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**model** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayModel**](GatewayModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayListModels

> GatewayListModels200Response GatewayListModels(ctx).Execute()

List models

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
	resp, r, err := apiClient.ModelsAPI.GatewayListModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.GatewayListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayListModels`: GatewayListModels200Response
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.GatewayListModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayListModelsRequest struct via the builder pattern


### Return type

[**GatewayListModels200Response**](GatewayListModels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlGetModel

> MlModel MlGetModel(ctx, modelId).Execute()

Get model details

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
	modelId := "modelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.MlGetModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.MlGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetModel`: MlModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.MlGetModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlGetModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MlModel**](MlModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlPromoteModel

> MlModel MlPromoteModel(ctx, modelId).MlPromoteModelRequest(mlPromoteModelRequest).Execute()

Promote a model



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
	modelId := "modelId_example" // string | 
	mlPromoteModelRequest := *openapiclient.NewMlPromoteModelRequest("Stage_example") // MlPromoteModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelsAPI.MlPromoteModel(context.Background(), modelId).MlPromoteModelRequest(mlPromoteModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.MlPromoteModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlPromoteModel`: MlModel
	fmt.Fprintf(os.Stdout, "Response from `ModelsAPI.MlPromoteModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlPromoteModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mlPromoteModelRequest** | [**MlPromoteModelRequest**](MlPromoteModelRequest.md) |  | 

### Return type

[**MlModel**](MlModel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MlRollbackModel

> MlRollbackModel(ctx, modelId).MlRollbackModelRequest(mlRollbackModelRequest).Execute()

Rollback a model



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
	modelId := "modelId_example" // string | 
	mlRollbackModelRequest := *openapiclient.NewMlRollbackModelRequest("Version_example") // MlRollbackModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ModelsAPI.MlRollbackModel(context.Background(), modelId).MlRollbackModelRequest(mlRollbackModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelsAPI.MlRollbackModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMlRollbackModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mlRollbackModelRequest** | [**MlRollbackModelRequest**](MlRollbackModelRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

