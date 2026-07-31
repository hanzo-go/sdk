# \GpusAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudListGpuAlerts**](GpusAPI.md#CloudListGpuAlerts) | **Get** /v1/gpus/alerts | Is an HONEST empty surface: Visor exposes no GPU alert inventory, so this returns [] rather than fabricating alerts.
[**CloudListGpus**](GpusAPI.md#CloudListGpus) | **Get** /v1/gpus | Returns one row per physical accelerator the caller&#39;s org has, derived from its real GPU machines (the size slug says how many cards a node holds) and from the accelerators BYO workers report through nvidia-smi.



## CloudListGpuAlerts

> CloudGpuAlertList CloudListGpuAlerts(ctx).Execute()

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
	resp, r, err := apiClient.GpusAPI.CloudListGpuAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpusAPI.CloudListGpuAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListGpuAlerts`: CloudGpuAlertList
	fmt.Fprintf(os.Stdout, "Response from `GpusAPI.CloudListGpuAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListGpuAlertsRequest struct via the builder pattern


### Return type

[**CloudGpuAlertList**](CloudGpuAlertList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListGpus

> CloudGpuList CloudListGpus(ctx).Execute()

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
	resp, r, err := apiClient.GpusAPI.CloudListGpus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GpusAPI.CloudListGpus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListGpus`: CloudGpuList
	fmt.Fprintf(os.Stdout, "Response from `GpusAPI.CloudListGpus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListGpusRequest struct via the builder pattern


### Return type

[**CloudGpuList**](CloudGpuList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

