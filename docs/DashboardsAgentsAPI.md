# \DashboardsAgentsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetAgentsDashboardUrl**](DashboardsAgentsAPI.md#AiGetAgentsDashboardUrl) | **Get** /v1/ai/dashboards/agents | Dashboards Agents



## AiGetAgentsDashboardUrl

> AiEnvelope AiGetAgentsDashboardUrl(ctx).Execute()

Dashboards Agents

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
	resp, r, err := apiClient.DashboardsAgentsAPI.AiGetAgentsDashboardUrl(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAgentsAPI.AiGetAgentsDashboardUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetAgentsDashboardUrl`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAgentsAPI.AiGetAgentsDashboardUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetAgentsDashboardUrlRequest struct via the builder pattern


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

