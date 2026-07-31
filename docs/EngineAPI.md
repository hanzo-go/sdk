# \EngineAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1EngineModel**](EngineAPI.md#CloudGetV1EngineModel) | **Get** /v1/engine/model | Model reads one model&#39;s load state — loaded, unloading, or not_found, as the engine itself reports it.
[**CloudGetV1EngineModels**](EngineAPI.md#CloudGetV1EngineModels) | **Get** /v1/engine/models | Models lists the models the engine serves, each with its load state — the server&#39;s own model table (its OpenAI-style list envelope, load status included), relayed verbatim.
[**CloudGetV1EngineStatus**](EngineAPI.md#CloudGetV1EngineStatus) | **Get** /v1/engine/status | Status reports whether the engine deployment is reachable and which build revision it runs — an honest lens for \&quot;is the serving runtime up\&quot;, never a fabricated ok.
[**CloudGetV1EngineSystem**](EngineAPI.md#CloudGetV1EngineSystem) | **Get** /v1/engine/system | System reads the engine host&#39;s inventory: OS, CPU, memory, every accelerator device with its VRAM and compute capability, and the build&#39;s capabilities (CUDA/Metal/flash-attention) — the real hardware under the serving runtime, relayed verbatim.



## CloudGetV1EngineModel

> interface{} CloudGetV1EngineModel(ctx).Model(model).Execute()

Model reads one model's load state — loaded, unloading, or not_found, as the engine itself reports it.



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
	model := "Qwen/Qwen3-4B" // string | Model is the model id to inspect, exactly as the model list reports it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EngineAPI.CloudGetV1EngineModel(context.Background()).Model(model).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.CloudGetV1EngineModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1EngineModel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.CloudGetV1EngineModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EngineModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model** | **string** | Model is the model id to inspect, exactly as the model list reports it. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1EngineModels

> interface{} CloudGetV1EngineModels(ctx).Execute()

Models lists the models the engine serves, each with its load state — the server's own model table (its OpenAI-style list envelope, load status included), relayed verbatim.



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
	resp, r, err := apiClient.EngineAPI.CloudGetV1EngineModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.CloudGetV1EngineModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1EngineModels`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.CloudGetV1EngineModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EngineModelsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1EngineStatus

> CloudEngineStatus CloudGetV1EngineStatus(ctx).Execute()

Status reports whether the engine deployment is reachable and which build revision it runs — an honest lens for \"is the serving runtime up\", never a fabricated ok.



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
	resp, r, err := apiClient.EngineAPI.CloudGetV1EngineStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.CloudGetV1EngineStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1EngineStatus`: CloudEngineStatus
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.CloudGetV1EngineStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EngineStatusRequest struct via the builder pattern


### Return type

[**CloudEngineStatus**](CloudEngineStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1EngineSystem

> interface{} CloudGetV1EngineSystem(ctx).Execute()

System reads the engine host's inventory: OS, CPU, memory, every accelerator device with its VRAM and compute capability, and the build's capabilities (CUDA/Metal/flash-attention) — the real hardware under the serving runtime, relayed verbatim.



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
	resp, r, err := apiClient.EngineAPI.CloudGetV1EngineSystem(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EngineAPI.CloudGetV1EngineSystem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1EngineSystem`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `EngineAPI.CloudGetV1EngineSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EngineSystemRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

