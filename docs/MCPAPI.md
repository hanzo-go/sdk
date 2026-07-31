# \MCPAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationsMcp**](MCPAPI.md#AutomationsMcp) | **Post** /v1/automations/mcp | JSON-RPC 2.0 tool surface over connector actions



## AutomationsMcp

> AutomationsMcpResponse AutomationsMcp(ctx).AutomationsMcpRequest(automationsMcpRequest).Execute()

JSON-RPC 2.0 tool surface over connector actions



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
	automationsMcpRequest := *openapiclient.NewAutomationsMcpRequest("2.0", "Method_example") // AutomationsMcpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MCPAPI.AutomationsMcp(context.Background()).AutomationsMcpRequest(automationsMcpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MCPAPI.AutomationsMcp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsMcp`: AutomationsMcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MCPAPI.AutomationsMcp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsMcpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationsMcpRequest** | [**AutomationsMcpRequest**](AutomationsMcpRequest.md) |  | 

### Return type

[**AutomationsMcpResponse**](AutomationsMcpResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

