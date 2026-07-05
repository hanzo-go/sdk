# \VisorGPUsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VisorListGPUs**](VisorGPUsAPI.md#VisorListGPUs) | **Get** /v1/gpus | List per-accelerator GPU inventory (derived from GPU machines)
[**VisorListGpuAlerts**](VisorGPUsAPI.md#VisorListGpuAlerts) | **Get** /v1/gpus/alerts | List GPU alerts (honest empty — Visor carries no alert inventory)



## VisorListGPUs

> VisorListGPUs200Response VisorListGPUs(ctx).Execute()

List per-accelerator GPU inventory (derived from GPU machines)

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
	resp, r, err := apiClient.VisorGPUsAPI.VisorListGPUs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorGPUsAPI.VisorListGPUs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorListGPUs`: VisorListGPUs200Response
	fmt.Fprintf(os.Stdout, "Response from `VisorGPUsAPI.VisorListGPUs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVisorListGPUsRequest struct via the builder pattern


### Return type

[**VisorListGPUs200Response**](VisorListGPUs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorListGpuAlerts

> VisorListGpuAlerts200Response VisorListGpuAlerts(ctx).Execute()

List GPU alerts (honest empty — Visor carries no alert inventory)

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
	resp, r, err := apiClient.VisorGPUsAPI.VisorListGpuAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VisorGPUsAPI.VisorListGpuAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorListGpuAlerts`: VisorListGpuAlerts200Response
	fmt.Fprintf(os.Stdout, "Response from `VisorGPUsAPI.VisorListGpuAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVisorListGpuAlertsRequest struct via the builder pattern


### Return type

[**VisorListGpuAlerts200Response**](VisorListGpuAlerts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

