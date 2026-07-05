# \VisorMachinesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VisorDeleteMachine**](VisorMachinesAPI.md#VisorDeleteMachine) | **Delete** /v1/machines/{id} | Terminate a machine
[**VisorGetMachine**](VisorMachinesAPI.md#VisorGetMachine) | **Get** /v1/machines/{id} | Get one machine by org-scoped name
[**VisorLaunchMachine**](VisorMachinesAPI.md#VisorLaunchMachine) | **Post** /v1/machines | Launch a machine (or dryRun for a price quote)
[**VisorListMachines**](VisorMachinesAPI.md#VisorListMachines) | **Get** /v1/machines | List the org&#39;s machines



## VisorDeleteMachine

> VisorDeleteMachine(ctx, id).Execute()

Terminate a machine

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
	r, err := apiClient.VisorMachinesAPI.VisorDeleteMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorMachinesAPI.VisorDeleteMachine``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVisorDeleteMachineRequest struct via the builder pattern


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


## VisorGetMachine

> VisorMachineView VisorGetMachine(ctx, id).Execute()

Get one machine by org-scoped name

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
	id := "id_example" // string | Org-scoped machine name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VisorMachinesAPI.VisorGetMachine(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorMachinesAPI.VisorGetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorGetMachine`: VisorMachineView
	fmt.Fprintf(os.Stdout, "Response from `VisorMachinesAPI.VisorGetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Org-scoped machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisorMachineView**](VisorMachineView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorLaunchMachine

> map[string]interface{} VisorLaunchMachine(ctx).VisorLaunchRequest(visorLaunchRequest).Execute()

Launch a machine (or dryRun for a price quote)

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
	visorLaunchRequest := *openapiclient.NewVisorLaunchRequest() // VisorLaunchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VisorMachinesAPI.VisorLaunchMachine(context.Background()).VisorLaunchRequest(visorLaunchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorMachinesAPI.VisorLaunchMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorLaunchMachine`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VisorMachinesAPI.VisorLaunchMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVisorLaunchMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visorLaunchRequest** | [**VisorLaunchRequest**](VisorLaunchRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorListMachines

> VisorListMachines200Response VisorListMachines(ctx).Execute()

List the org's machines

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
	resp, r, err := apiClient.VisorMachinesAPI.VisorListMachines(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorMachinesAPI.VisorListMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorListMachines`: VisorListMachines200Response
	fmt.Fprintf(os.Stdout, "Response from `VisorMachinesAPI.VisorListMachines`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVisorListMachinesRequest struct via the builder pattern


### Return type

[**VisorListMachines200Response**](VisorListMachines200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

