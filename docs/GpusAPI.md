# \GpusAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListGpuAlerts**](GpusAPI.md#ListGpuAlerts) | **Get** /v1/gpus/alerts | Is an HONEST empty surface: Visor exposes no GPU alert inventory, so this returns [] rather than fabricating alerts.
[**ListGpus**](GpusAPI.md#ListGpus) | **Get** /v1/gpus | Returns one row per physical accelerator the caller&#39;s org has, derived from its real GPU machines (the size slug says how many cards a node holds) and from the accelerators BYO workers report through nvidia-smi.



## ListGpuAlerts

> GpuAlertList ListGpuAlerts(ctx).Execute()

Is an HONEST empty surface: Visor exposes no GPU alert inventory, so this returns [] rather than fabricating alerts.



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
	resp, r, err := apiClient.GpusAPI.ListGpuAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpusAPI.ListGpuAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGpuAlerts`: GpuAlertList
	fmt.Fprintf(os.Stdout, "Response from `GpusAPI.ListGpuAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGpuAlertsRequest struct via the builder pattern


### Return type

[**GpuAlertList**](GpuAlertList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGpus

> GpuList ListGpus(ctx).Execute()

Returns one row per physical accelerator the caller's org has, derived from its real GPU machines (the size slug says how many cards a node holds) and from the accelerators BYO workers report through nvidia-smi.



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
	resp, r, err := apiClient.GpusAPI.ListGpus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpusAPI.ListGpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGpus`: GpuList
	fmt.Fprintf(os.Stdout, "Response from `GpusAPI.ListGpus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGpusRequest struct via the builder pattern


### Return type

[**GpuList**](GpuList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

