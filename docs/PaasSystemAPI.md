# \PaasSystemAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaasListTemplates**](PaasSystemAPI.md#PaasListTemplates) | **Get** /v1/paas/system/templates | List available templates



## PaasListTemplates

> []PaasListTemplates200ResponseInner PaasListTemplates(ctx).Execute()

List available templates

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
	resp, r, err := apiClient.PaasSystemAPI.PaasListTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaasSystemAPI.PaasListTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaasListTemplates`: []PaasListTemplates200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PaasSystemAPI.PaasListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaasListTemplatesRequest struct via the builder pattern


### Return type

[**[]PaasListTemplates200ResponseInner**](PaasListTemplates200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

