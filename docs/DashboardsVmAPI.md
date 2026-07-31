# \DashboardsVmAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetVmDashboardUrl**](DashboardsVmAPI.md#AiGetVmDashboardUrl) | **Get** /v1/ai/dashboards/vm | Dashboards Vm



## AiGetVmDashboardUrl

> AiEnvelope AiGetVmDashboardUrl(ctx).Execute()

Dashboards Vm

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
	resp, r, err := apiClient.DashboardsVmAPI.AiGetVmDashboardUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsVmAPI.AiGetVmDashboardUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetVmDashboardUrl`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DashboardsVmAPI.AiGetVmDashboardUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetVmDashboardUrlRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

