# \EngineAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EngineModel**](EngineAPI.md#EngineModel) | **Get** /v1/engine/model | Read one model&#39;s load state on the serving runtime
[**EngineModels**](EngineAPI.md#EngineModels) | **Get** /v1/engine/models | List the models the serving runtime holds, with each one&#39;s load state
[**EngineStatus**](EngineAPI.md#EngineStatus) | **Get** /v1/engine/status | Whether the serving runtime is reachable, and which build it runs
[**EngineSystem**](EngineAPI.md#EngineSystem) | **Get** /v1/engine/system | The serving host&#39;s own inventory: devices, memory and build capabilities



## EngineModel

> interface{} EngineModel(ctx).Model(model).Execute()

Read one model's load state on the serving runtime



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
	model := "Qwen/Qwen3-4B" // string | Model is the model id to inspect, exactly as the model list reports it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.EngineModel(context.Background()).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.EngineModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineModel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.EngineModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEngineModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | Model is the model id to inspect, exactly as the model list reports it. | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineModels

> interface{} EngineModels(ctx).Execute()

List the models the serving runtime holds, with each one's load state



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
	resp, r, err := apiClient.EngineAPI.EngineModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.EngineModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineModels`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.EngineModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEngineModelsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineStatus

> EngineStatus EngineStatus(ctx).Execute()

Whether the serving runtime is reachable, and which build it runs



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
	resp, r, err := apiClient.EngineAPI.EngineStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.EngineStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineStatus`: EngineStatus
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.EngineStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEngineStatusRequest struct via the builder pattern


### Return type

[**EngineStatus**](EngineStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EngineSystem

> interface{} EngineSystem(ctx).Execute()

The serving host's own inventory: devices, memory and build capabilities



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
	resp, r, err := apiClient.EngineAPI.EngineSystem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.EngineSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EngineSystem`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.EngineSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEngineSystemRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

