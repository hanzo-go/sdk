# \ZtMeshAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZtListMeshServices**](ZtMeshAPI.md#ZtListMeshServices) | **Get** /v1/mesh/services | List the org&#39;s ZT edge services



## ZtListMeshServices

> ZtListMeshServices200Response ZtListMeshServices(ctx).Execute()

List the org's ZT edge services

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
	resp, r, err := apiClient.ZtMeshAPI.ZtListMeshServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZtMeshAPI.ZtListMeshServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZtListMeshServices`: ZtListMeshServices200Response
	fmt.Fprintf(os.Stdout, "Response from `ZtMeshAPI.ZtListMeshServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZtListMeshServicesRequest struct via the builder pattern


### Return type

[**ZtListMeshServices200Response**](ZtListMeshServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

