# \MachinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindMachineAgent**](MachinesAPI.md#BindMachineAgent) | **Put** /v1/machines/{id}/agent | Binds a cloud Agent to one of the caller org&#39;s machines: the machine is recorded as running that Agent&#39;s @hanzo/bot runtime.
[**DeleteMachine**](MachinesAPI.md#DeleteMachine) | **Delete** /v1/machines/{id} | Terminates one of the caller org&#39;s machines.
[**GetMachine**](MachinesAPI.md#GetMachine) | **Get** /v1/machines/{id} | Returns one of the caller org&#39;s machines by its org-scoped name.
[**GetMachineAgent**](MachinesAPI.md#GetMachineAgent) | **Get** /v1/machines/{id}/agent | Returns the agent binding of one of the caller org&#39;s machines, or 404 when the machine runs no bot runtime.
[**ListMachineAgents**](MachinesAPI.md#ListMachineAgents) | **Get** /v1/machines/agents | Returns every agent↔machine binding in the caller&#39;s org — which machines are running which cloud Agent, with vm&#39;s own reconciled status.
[**ListMachines**](MachinesAPI.md#ListMachines) | **Get** /v1/machines | Returns every machine the caller&#39;s org has — Visor&#39;s registry, the live DigitalOcean droplets and the DOKS worker nodes (deduped into one union), plus the BYO machines that dialed in via &#x60;hanzo link&#x60; (provider \&quot;byo\&quot;).
[**PostMachines**](MachinesAPI.md#PostMachines) | **Post** /v1/machines | Launch a metered machine for your org, or price one first with dryRun
[**UnbindMachineAgent**](MachinesAPI.md#UnbindMachineAgent) | **Delete** /v1/machines/{id}/agent | Detaches the agent runtime from one of the caller org&#39;s machines.



## BindMachineAgent

> AgentBinding BindMachineAgent(ctx, id).BindAgentReq(bindAgentReq).Execute()

Binds a cloud Agent to one of the caller org's machines: the machine is recorded as running that Agent's @hanzo/bot runtime.



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
	id := "id_example" // string | ID is the machine to bind, from the URL path.
	bindAgentReq := *openapiclient.NewBindAgentReq() // BindAgentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.BindMachineAgent(context.Background(), id).BindAgentReq(bindAgentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.BindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindMachineAgent`: AgentBinding
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.BindMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine to bind, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindAgentReq** | [**BindAgentReq**](BindAgentReq.md) |  | 

### Return type

[**AgentBinding**](AgentBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMachine

> DeleteMachine(ctx, id).Execute()

Terminates one of the caller org's machines.



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
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.DeleteMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.DeleteMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachine

> MachineView GetMachine(ctx, id).Execute()

Returns one of the caller org's machines by its org-scoped name.



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
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.GetMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.GetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachine`: MachineView
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.GetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MachineView**](MachineView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineAgent

> AgentBinding GetMachineAgent(ctx, id).Execute()

Returns the agent binding of one of the caller org's machines, or 404 when the machine runs no bot runtime.



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
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.GetMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.GetMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachineAgent`: AgentBinding
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.GetMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AgentBinding**](AgentBinding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachineAgents

> BindingList ListMachineAgents(ctx).Execute()

Returns every agent↔machine binding in the caller's org — which machines are running which cloud Agent, with vm's own reconciled status.



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
	resp, r, err := apiClient.MachinesAPI.ListMachineAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.ListMachineAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMachineAgents`: BindingList
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.ListMachineAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMachineAgentsRequest struct via the builder pattern


### Return type

[**BindingList**](BindingList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachines

> MachineList ListMachines(ctx).Execute()

Returns every machine the caller's org has — Visor's registry, the live DigitalOcean droplets and the DOKS worker nodes (deduped into one union), plus the BYO machines that dialed in via `hanzo link` (provider \"byo\").



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
	resp, r, err := apiClient.MachinesAPI.ListMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.ListMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMachines`: MachineList
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.ListMachines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMachinesRequest struct via the builder pattern


### Return type

[**MachineList**](MachineList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMachines

> PostMachines(ctx).Execute()

Launch a metered machine for your org, or price one first with dryRun



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
	r, err := apiClient.MachinesAPI.PostMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.PostMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostMachinesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindMachineAgent

> UnbindMachineAgent(ctx, id).Execute()

Detaches the agent runtime from one of the caller org's machines.



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
	id := "id_example" // string | ID is the machine's org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachinesAPI.UnbindMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.UnbindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnbindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

