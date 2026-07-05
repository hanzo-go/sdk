# \MlModelsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MlGetModel**](MlModelsAPI.md#MlGetModel) | **Get** /v1/ml/models/{model_id} | Get model details
[**MlListModels**](MlModelsAPI.md#MlListModels) | **Get** /v1/ml/models | List models
[**MlPromoteModel**](MlModelsAPI.md#MlPromoteModel) | **Post** /v1/ml/models/{model_id}/promote | Promote a model
[**MlRegisterModel**](MlModelsAPI.md#MlRegisterModel) | **Post** /v1/ml/models | Register a model
[**MlRollbackModel**](MlModelsAPI.md#MlRollbackModel) | **Post** /v1/ml/models/{model_id}/rollback | Rollback a model



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
	resp, r, err := apiClient.MlModelsAPI.MlGetModel(context.Background(), modelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlModelsAPI.MlGetModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlGetModel`: MlModel
	fmt.Fprintf(os.Stdout, "Response from `MlModelsAPI.MlGetModel`: %v\n", resp)
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


## MlListModels

> MlListModels200Response MlListModels(ctx).Stage(stage).Search(search).Execute()

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
	stage := "stage_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlModelsAPI.MlListModels(context.Background()).Stage(stage).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlModelsAPI.MlListModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlListModels`: MlListModels200Response
	fmt.Fprintf(os.Stdout, "Response from `MlModelsAPI.MlListModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlListModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**MlListModels200Response**](MlListModels200Response.md)

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
	resp, r, err := apiClient.MlModelsAPI.MlPromoteModel(context.Background(), modelId).MlPromoteModelRequest(mlPromoteModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlModelsAPI.MlPromoteModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlPromoteModel`: MlModel
	fmt.Fprintf(os.Stdout, "Response from `MlModelsAPI.MlPromoteModel`: %v\n", resp)
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


## MlRegisterModel

> MlModel MlRegisterModel(ctx).MlRegisterModelRequest(mlRegisterModelRequest).Execute()

Register a model

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
	mlRegisterModelRequest := *openapiclient.NewMlRegisterModelRequest("Name_example", "Version_example") // MlRegisterModelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MlModelsAPI.MlRegisterModel(context.Background()).MlRegisterModelRequest(mlRegisterModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlModelsAPI.MlRegisterModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MlRegisterModel`: MlModel
	fmt.Fprintf(os.Stdout, "Response from `MlModelsAPI.MlRegisterModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMlRegisterModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mlRegisterModelRequest** | [**MlRegisterModelRequest**](MlRegisterModelRequest.md) |  | 

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
	r, err := apiClient.MlModelsAPI.MlRollbackModel(context.Background(), modelId).MlRollbackModelRequest(mlRollbackModelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MlModelsAPI.MlRollbackModel``: %v\n", err)
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

