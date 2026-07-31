# \MachinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudBindMachineAgent**](MachinesAPI.md#CloudBindMachineAgent) | **Put** /v1/machines/{id}/agent | Binds a cloud Agent to one of the caller org&#39;s machines: the machine is recorded as running that Agent&#39;s @hanzo/bot runtime.
[**CloudDeleteMachine**](MachinesAPI.md#CloudDeleteMachine) | **Delete** /v1/machines/{id} | Terminates one of the caller org&#39;s machines.
[**CloudGetMachine**](MachinesAPI.md#CloudGetMachine) | **Get** /v1/machines/{id} | Returns one of the caller org&#39;s machines by its org-scoped name.
[**CloudGetMachineAgent**](MachinesAPI.md#CloudGetMachineAgent) | **Get** /v1/machines/{id}/agent | Returns the agent binding of one of the caller org&#39;s machines, or 404 when the machine runs no bot runtime.
[**CloudListMachineAgents**](MachinesAPI.md#CloudListMachineAgents) | **Get** /v1/machines/agents | Returns every agent↔machine binding in the caller&#39;s org — which machines are running which cloud Agent, with vm&#39;s own reconciled status.
[**CloudListMachines**](MachinesAPI.md#CloudListMachines) | **Get** /v1/machines | Returns every machine the caller&#39;s org has — Visor&#39;s registry, the live DigitalOcean droplets and the DOKS worker nodes (deduped into one union), plus the BYO machines that dialed in via &#x60;hanzo link&#x60; (provider \&quot;byo\&quot;).
[**CloudPostV1Machines**](MachinesAPI.md#CloudPostV1Machines) | **Post** /v1/machines | 
[**CloudUnbindMachineAgent**](MachinesAPI.md#CloudUnbindMachineAgent) | **Delete** /v1/machines/{id}/agent | Detaches the agent runtime from one of the caller org&#39;s machines.



## CloudBindMachineAgent

> CloudAgentBinding CloudBindMachineAgent(ctx, id).CloudBindAgentReq(cloudBindAgentReq).Execute()

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
	cloudBindAgentReq := *openapiclient.NewCloudBindAgentReq() // CloudBindAgentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachinesAPI.CloudBindMachineAgent(context.Background(), id).CloudBindAgentReq(cloudBindAgentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudBindMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudBindMachineAgent`: CloudAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.CloudBindMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine to bind, from the URL path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudBindMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudBindAgentReq** | [**CloudBindAgentReq**](CloudBindAgentReq.md) |  | 

### Return type

[**CloudAgentBinding**](CloudAgentBinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteMachine

> CloudDeleteMachine(ctx, id).Execute()

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
	r, err := apiClient.MachinesAPI.CloudDeleteMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudDeleteMachine``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteMachineRequest struct via the builder pattern


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


## CloudGetMachine

> CloudMachineView CloudGetMachine(ctx, id).Execute()

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
	resp, r, err := apiClient.MachinesAPI.CloudGetMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudGetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetMachine`: CloudMachineView
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.CloudGetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMachineView**](CloudMachineView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetMachineAgent

> CloudAgentBinding CloudGetMachineAgent(ctx, id).Execute()

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
	resp, r, err := apiClient.MachinesAPI.CloudGetMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudGetMachineAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetMachineAgent`: CloudAgentBinding
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.CloudGetMachineAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the machine&#39;s org-scoped NAME — the stable key Visor addresses a machine by (owner/name), not the ephemeral provider id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetMachineAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAgentBinding**](CloudAgentBinding.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListMachineAgents

> CloudBindingList CloudListMachineAgents(ctx).Execute()

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
	resp, r, err := apiClient.MachinesAPI.CloudListMachineAgents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudListMachineAgents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListMachineAgents`: CloudBindingList
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.CloudListMachineAgents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListMachineAgentsRequest struct via the builder pattern


### Return type

[**CloudBindingList**](CloudBindingList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListMachines

> CloudMachineList CloudListMachines(ctx).Execute()

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
	resp, r, err := apiClient.MachinesAPI.CloudListMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudListMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListMachines`: CloudMachineList
	fmt.Fprintf(os.Stdout, "Response from `MachinesAPI.CloudListMachines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListMachinesRequest struct via the builder pattern


### Return type

[**CloudMachineList**](CloudMachineList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Machines

> CloudPostV1Machines(ctx).Execute()



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
	r, err := apiClient.MachinesAPI.CloudPostV1Machines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudPostV1Machines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MachinesRequest struct via the builder pattern


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


## CloudUnbindMachineAgent

> CloudUnbindMachineAgent(ctx, id).Execute()

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
	r, err := apiClient.MachinesAPI.CloudUnbindMachineAgent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachinesAPI.CloudUnbindMachineAgent``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudUnbindMachineAgentRequest struct via the builder pattern


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

