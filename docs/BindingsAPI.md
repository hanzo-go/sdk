# \BindingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VisorBindMachineAgent**](BindingsAPI.md#VisorBindMachineAgent) | **Post** /v1/machines/{id}/bind-agent | Bind a cloud Agent to a machine
[**VisorGetMachineAgentBinding**](BindingsAPI.md#VisorGetMachineAgentBinding) | **Get** /v1/machines/{id}/agent-binding | Get a machine&#39;s agent binding
[**VisorUnbindMachineAgent**](BindingsAPI.md#VisorUnbindMachineAgent) | **Delete** /v1/machines/{id}/agent-binding | Unbind the agent from a machine



## VisorBindMachineAgent

> VisorAgentBinding VisorBindMachineAgent(ctx, id).VisorBindAgentRequest(visorBindAgentRequest).Execute()

Bind a cloud Agent to a machine

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
	id := "id_example" // string | 
	visorBindAgentRequest := *openapiclient.NewVisorBindAgentRequest("AgentName_example") // VisorBindAgentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BindingsAPI.VisorBindMachineAgent(context.Background(), id).VisorBindAgentRequest(visorBindAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BindingsAPI.VisorBindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorBindMachineAgent`: VisorAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `BindingsAPI.VisorBindMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorBindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **visorBindAgentRequest** | [**VisorBindAgentRequest**](VisorBindAgentRequest.md) |  | 

### Return type

[**VisorAgentBinding**](VisorAgentBinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorGetMachineAgentBinding

> VisorAgentBinding VisorGetMachineAgentBinding(ctx, id).Execute()

Get a machine's agent binding

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BindingsAPI.VisorGetMachineAgentBinding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BindingsAPI.VisorGetMachineAgentBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorGetMachineAgentBinding`: VisorAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `BindingsAPI.VisorGetMachineAgentBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorGetMachineAgentBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisorAgentBinding**](VisorAgentBinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorUnbindMachineAgent

> VisorUnbindMachineAgent(ctx, id).Execute()

Unbind the agent from a machine

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BindingsAPI.VisorUnbindMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BindingsAPI.VisorUnbindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorUnbindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

