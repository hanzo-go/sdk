# \MachineAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddMachine**](MachineAPIAPI.md#CloudApiControllerAddMachine) | **Post** /v1/cloud/add-machine | Api Controller Add Machine
[**CloudApiControllerDeleteMachine**](MachineAPIAPI.md#CloudApiControllerDeleteMachine) | **Post** /v1/cloud/delete-machine | Api Controller Delete Machine
[**CloudApiControllerGetMachine**](MachineAPIAPI.md#CloudApiControllerGetMachine) | **Get** /v1/cloud/get-machine | Api Controller Get Machine
[**CloudApiControllerGetMachines**](MachineAPIAPI.md#CloudApiControllerGetMachines) | **Get** /v1/cloud/get-machines | Api Controller Get Machines
[**CloudApiControllerUpdateMachine**](MachineAPIAPI.md#CloudApiControllerUpdateMachine) | **Post** /v1/cloud/update-machine | Api Controller Update Machine



## CloudApiControllerAddMachine

> CloudControllersResponse CloudApiControllerAddMachine(ctx).CloudObjectMachine(cloudObjectMachine).Execute()

Api Controller Add Machine



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
	cloudObjectMachine := *openapiclient.NewCloudObjectMachine() // CloudObjectMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAPIAPI.CloudApiControllerAddMachine(context.Background()).CloudObjectMachine(cloudObjectMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAPIAPI.CloudApiControllerAddMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddMachine`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MachineAPIAPI.CloudApiControllerAddMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectMachine** | [**CloudObjectMachine**](CloudObjectMachine.md) | The details of the machine | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteMachine

> CloudControllersResponse CloudApiControllerDeleteMachine(ctx).CloudObjectMachine(cloudObjectMachine).Execute()

Api Controller Delete Machine



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
	cloudObjectMachine := *openapiclient.NewCloudObjectMachine() // CloudObjectMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAPIAPI.CloudApiControllerDeleteMachine(context.Background()).CloudObjectMachine(cloudObjectMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAPIAPI.CloudApiControllerDeleteMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteMachine`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MachineAPIAPI.CloudApiControllerDeleteMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectMachine** | [**CloudObjectMachine**](CloudObjectMachine.md) | The details of the machine | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetMachine

> CloudObjectMachine CloudApiControllerGetMachine(ctx).Id(id).Execute()

Api Controller Get Machine



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
	id := "id_example" // string | The id ( owner/name ) of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAPIAPI.CloudApiControllerGetMachine(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAPIAPI.CloudApiControllerGetMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetMachine`: CloudObjectMachine
	fmt.Fprintf(os.Stdout, "Response from `MachineAPIAPI.CloudApiControllerGetMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the machine | 

### Return type

[**CloudObjectMachine**](CloudObjectMachine.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetMachines

> CloudObjectMachine CloudApiControllerGetMachines(ctx).PageSize(pageSize).P(p).Execute()

Api Controller Get Machines



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
	p := "p_example" // string | The number of the page

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAPIAPI.CloudApiControllerGetMachines(context.Background()).PageSize(pageSize).P(p).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAPIAPI.CloudApiControllerGetMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetMachines`: CloudObjectMachine
	fmt.Fprintf(os.Stdout, "Response from `MachineAPIAPI.CloudApiControllerGetMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | The size of each page | 
 **p** | **string** | The number of the page | 

### Return type

[**CloudObjectMachine**](CloudObjectMachine.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateMachine

> CloudControllersResponse CloudApiControllerUpdateMachine(ctx).Id(id).CloudObjectMachine(cloudObjectMachine).Execute()

Api Controller Update Machine



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
	id := "id_example" // string | The id ( owner/name ) of the machine
	cloudObjectMachine := *openapiclient.NewCloudObjectMachine() // CloudObjectMachine | The details of the machine

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAPIAPI.CloudApiControllerUpdateMachine(context.Background()).Id(id).CloudObjectMachine(cloudObjectMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAPIAPI.CloudApiControllerUpdateMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateMachine`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MachineAPIAPI.CloudApiControllerUpdateMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the machine | 
 **cloudObjectMachine** | [**CloudObjectMachine**](CloudObjectMachine.md) | The details of the machine | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

