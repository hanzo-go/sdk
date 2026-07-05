# \KmsServiceTokensAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetServiceToken**](KmsServiceTokensAPI.md#KmsGetServiceToken) | **Get** /v1/kms/service-token | Get the service token associated with the current request



## KmsGetServiceToken

> KmsGetServiceToken200Response KmsGetServiceToken(ctx).Execute()

Get the service token associated with the current request

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
	resp, r, err := apiClient.KmsServiceTokensAPI.KmsGetServiceToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsServiceTokensAPI.KmsGetServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetServiceToken`: KmsGetServiceToken200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsServiceTokensAPI.KmsGetServiceToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetServiceTokenRequest struct via the builder pattern


### Return type

[**KmsGetServiceToken200Response**](KmsGetServiceToken200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

