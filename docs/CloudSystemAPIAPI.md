# \CloudSystemAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerGetPrometheusInfo**](CloudSystemAPIAPI.md#CloudApiControllerGetPrometheusInfo) | **Get** /v1/cloud/get-prometheus-info | Api Controller Get Prometheus Info
[**CloudApiControllerGetSystemInfo**](CloudSystemAPIAPI.md#CloudApiControllerGetSystemInfo) | **Get** /v1/cloud/get-system-info | Api Controller Get System Info
[**CloudApiControllerGetVersionInfo**](CloudSystemAPIAPI.md#CloudApiControllerGetVersionInfo) | **Get** /v1/cloud/get-version-info | Api Controller Get Version Info



## CloudApiControllerGetPrometheusInfo

> CloudObjectPrometheusInfo CloudApiControllerGetPrometheusInfo(ctx).Execute()

Api Controller Get Prometheus Info



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
	resp, r, err := apiClient.CloudSystemAPIAPI.CloudApiControllerGetPrometheusInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSystemAPIAPI.CloudApiControllerGetPrometheusInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetPrometheusInfo`: CloudObjectPrometheusInfo
	fmt.Fprintf(os.Stdout, "Response from `CloudSystemAPIAPI.CloudApiControllerGetPrometheusInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetPrometheusInfoRequest struct via the builder pattern


### Return type

[**CloudObjectPrometheusInfo**](CloudObjectPrometheusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetSystemInfo

> CloudUtilSystemInfo CloudApiControllerGetSystemInfo(ctx).Execute()

Api Controller Get System Info



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
	resp, r, err := apiClient.CloudSystemAPIAPI.CloudApiControllerGetSystemInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSystemAPIAPI.CloudApiControllerGetSystemInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetSystemInfo`: CloudUtilSystemInfo
	fmt.Fprintf(os.Stdout, "Response from `CloudSystemAPIAPI.CloudApiControllerGetSystemInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetSystemInfoRequest struct via the builder pattern


### Return type

[**CloudUtilSystemInfo**](CloudUtilSystemInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetVersionInfo

> CloudUtilVersionInfo CloudApiControllerGetVersionInfo(ctx).Execute()

Api Controller Get Version Info



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
	resp, r, err := apiClient.CloudSystemAPIAPI.CloudApiControllerGetVersionInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSystemAPIAPI.CloudApiControllerGetVersionInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetVersionInfo`: CloudUtilVersionInfo
	fmt.Fprintf(os.Stdout, "Response from `CloudSystemAPIAPI.CloudApiControllerGetVersionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetVersionInfoRequest struct via the builder pattern


### Return type

[**CloudUtilVersionInfo**](CloudUtilVersionInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

