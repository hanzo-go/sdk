# \CsrfAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Csrf**](CsrfAPI.md#CloudGetV1Csrf) | **Get** /v1/csrf | IssueCSRFToken mints the anti-CSRF token a browser echoes as X-CSRF-Token on every money write (mint/revoke a key, top up, onboard, and the billing/commerce write verbs).



## CloudGetV1Csrf

> CloudCsrfResp CloudGetV1Csrf(ctx).Execute()

IssueCSRFToken mints the anti-CSRF token a browser echoes as X-CSRF-Token on every money write (mint/revoke a key, top up, onboard, and the billing/commerce write verbs).



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
	resp, r, err := apiClient.CsrfAPI.CloudGetV1Csrf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsrfAPI.CloudGetV1Csrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Csrf`: CloudCsrfResp
	fmt.Fprintf(os.Stdout, "Response from `CsrfAPI.CloudGetV1Csrf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CsrfRequest struct via the builder pattern


### Return type

[**CloudCsrfResp**](CloudCsrfResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

