# \NexusSystemAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusGetPrometheusInfo**](NexusSystemAPIAPI.md#NexusGetPrometheusInfo) | **Get** /v1/nexus/get-prometheus-info | get Prometheus Info
[**NexusGetSystemInfo**](NexusSystemAPIAPI.md#NexusGetSystemInfo) | **Get** /v1/nexus/get-system-info | get System Info
[**NexusGetVersionInfo**](NexusSystemAPIAPI.md#NexusGetVersionInfo) | **Get** /v1/nexus/get-version-info | get Version Info



## NexusGetPrometheusInfo

> NexusPrometheusInfo NexusGetPrometheusInfo(ctx).Execute()

get Prometheus Info



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
	resp, r, err := apiClient.NexusSystemAPIAPI.NexusGetPrometheusInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusSystemAPIAPI.NexusGetPrometheusInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetPrometheusInfo`: NexusPrometheusInfo
	fmt.Fprintf(os.Stdout, "Response from `NexusSystemAPIAPI.NexusGetPrometheusInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetPrometheusInfoRequest struct via the builder pattern


### Return type

[**NexusPrometheusInfo**](NexusPrometheusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetSystemInfo

> NexusSystemInfo NexusGetSystemInfo(ctx).Execute()

get System Info



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
	resp, r, err := apiClient.NexusSystemAPIAPI.NexusGetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusSystemAPIAPI.NexusGetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetSystemInfo`: NexusSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `NexusSystemAPIAPI.NexusGetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetSystemInfoRequest struct via the builder pattern


### Return type

[**NexusSystemInfo**](NexusSystemInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetVersionInfo

> NexusVersionInfo NexusGetVersionInfo(ctx).Execute()

get Version Info



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
	resp, r, err := apiClient.NexusSystemAPIAPI.NexusGetVersionInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusSystemAPIAPI.NexusGetVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetVersionInfo`: NexusVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `NexusSystemAPIAPI.NexusGetVersionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetVersionInfoRequest struct via the builder pattern


### Return type

[**NexusVersionInfo**](NexusVersionInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

