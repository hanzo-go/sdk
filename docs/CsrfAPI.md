# \CsrfAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCsrf**](CsrfAPI.md#GetCsrf) | **Get** /v1/csrf | IssueCSRFToken mints the anti-CSRF token a browser echoes as X-CSRF-Token on every money write (mint/revoke a key, top up, onboard, and the billing/commerce write verbs).



## GetCsrf

> CsrfResp GetCsrf(ctx).Execute()

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
	resp, r, err := apiClient.CsrfAPI.GetCsrf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsrfAPI.GetCsrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCsrf`: CsrfResp
	fmt.Fprintf(os.Stdout, "Response from `CsrfAPI.GetCsrf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsrfRequest struct via the builder pattern


### Return type

[**CsrfResp**](CsrfResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

