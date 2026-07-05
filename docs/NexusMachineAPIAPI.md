# \NexusMachineAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddMachine**](NexusMachineAPIAPI.md#NexusAddMachine) | **Post** /v1/nexus/add-machine | add Machine
[**NexusDeleteMachine**](NexusMachineAPIAPI.md#NexusDeleteMachine) | **Post** /v1/nexus/delete-machine | delete Machine
[**NexusGetMachine**](NexusMachineAPIAPI.md#NexusGetMachine) | **Get** /v1/nexus/get-machine | get Machine
[**NexusGetMachines**](NexusMachineAPIAPI.md#NexusGetMachines) | **Get** /v1/nexus/get-machines | get Machines
[**NexusUpdateMachine**](NexusMachineAPIAPI.md#NexusUpdateMachine) | **Post** /v1/nexus/update-machine | update Machine



## NexusAddMachine

> NexusResponse NexusAddMachine(ctx).NexusMachine(nexusMachine).Execute()

add Machine



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
	nexusMachine := *openapiclient.NewNexusMachine() // NexusMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMachineAPIAPI.NexusAddMachine(context.Background()).NexusMachine(nexusMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMachineAPIAPI.NexusAddMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddMachine`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusMachineAPIAPI.NexusAddMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusMachine** | [**NexusMachine**](NexusMachine.md) | The details of the machine | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteMachine

> NexusResponse NexusDeleteMachine(ctx).NexusMachine(nexusMachine).Execute()

delete Machine



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
	nexusMachine := *openapiclient.NewNexusMachine() // NexusMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMachineAPIAPI.NexusDeleteMachine(context.Background()).NexusMachine(nexusMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMachineAPIAPI.NexusDeleteMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteMachine`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusMachineAPIAPI.NexusDeleteMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusMachine** | [**NexusMachine**](NexusMachine.md) | The details of the machine | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetMachine

> NexusMachine NexusGetMachine(ctx).Id(id).Execute()

get Machine



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
	id := "id_example" // string | The id (owner/name) of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMachineAPIAPI.NexusGetMachine(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMachineAPIAPI.NexusGetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetMachine`: NexusMachine
	fmt.Fprintf(os.Stdout, "Response from `NexusMachineAPIAPI.NexusGetMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the machine | 

### Return type

[**NexusMachine**](NexusMachine.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetMachines

> NexusMachine NexusGetMachines(ctx).PageSize(pageSize).P(p).Execute()

get Machines



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
	pageSize := "pageSize_example" // string | The size of each page
	p := "p_example" // string | The page number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMachineAPIAPI.NexusGetMachines(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMachineAPIAPI.NexusGetMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetMachines`: NexusMachine
	fmt.Fprintf(os.Stdout, "Response from `NexusMachineAPIAPI.NexusGetMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The page number | 

### Return type

[**NexusMachine**](NexusMachine.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateMachine

> NexusResponse NexusUpdateMachine(ctx).Id(id).NexusMachine(nexusMachine).Execute()

update Machine



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
	id := "id_example" // string | The id (owner/name) of the machine
	nexusMachine := *openapiclient.NewNexusMachine() // NexusMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusMachineAPIAPI.NexusUpdateMachine(context.Background()).Id(id).NexusMachine(nexusMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusMachineAPIAPI.NexusUpdateMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateMachine`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusMachineAPIAPI.NexusUpdateMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the machine | 
 **nexusMachine** | [**NexusMachine**](NexusMachine.md) | The details of the machine | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

