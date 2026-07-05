# \VisorBindingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VisorBindMachineAgent**](VisorBindingsAPI.md#VisorBindMachineAgent) | **Post** /v1/machines/{id}/bind-agent | Bind a cloud Agent to a machine
[**VisorGetMachineAgentBinding**](VisorBindingsAPI.md#VisorGetMachineAgentBinding) | **Get** /v1/machines/{id}/agent-binding | Get a machine&#39;s agent binding
[**VisorListAgentBindings**](VisorBindingsAPI.md#VisorListAgentBindings) | **Get** /v1/agent-bindings | List the org&#39;s agent bindings
[**VisorUnbindMachineAgent**](VisorBindingsAPI.md#VisorUnbindMachineAgent) | **Delete** /v1/machines/{id}/agent-binding | Unbind the agent from a machine



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
	resp, r, err := apiClient.VisorBindingsAPI.VisorBindMachineAgent(context.Background(), id).VisorBindAgentRequest(visorBindAgentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorBindingsAPI.VisorBindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorBindMachineAgent`: VisorAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `VisorBindingsAPI.VisorBindMachineAgent`: %v\n", resp)
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
	resp, r, err := apiClient.VisorBindingsAPI.VisorGetMachineAgentBinding(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorBindingsAPI.VisorGetMachineAgentBinding``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorGetMachineAgentBinding`: VisorAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `VisorBindingsAPI.VisorGetMachineAgentBinding`: %v\n", resp)
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


## VisorListAgentBindings

> VisorListAgentBindings200Response VisorListAgentBindings(ctx).Execute()

List the org's agent bindings

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
	resp, r, err := apiClient.VisorBindingsAPI.VisorListAgentBindings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorBindingsAPI.VisorListAgentBindings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorListAgentBindings`: VisorListAgentBindings200Response
	fmt.Fprintf(os.Stdout, "Response from `VisorBindingsAPI.VisorListAgentBindings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVisorListAgentBindingsRequest struct via the builder pattern


### Return type

[**VisorListAgentBindings200Response**](VisorListAgentBindings200Response.md)

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
	r, err := apiClient.VisorBindingsAPI.VisorUnbindMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorBindingsAPI.VisorUnbindMachineAgent``: %v\n", err)
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

