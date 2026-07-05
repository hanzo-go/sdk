# \GraphOraclesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphListOracles**](GraphOraclesAPI.md#GraphListOracles) | **Get** /v1/oracles | List on-chain price/data oracles



## GraphListOracles

> GraphListOracles200Response GraphListOracles(ctx).Execute()

List on-chain price/data oracles

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
	resp, r, err := apiClient.GraphOraclesAPI.GraphListOracles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphOraclesAPI.GraphListOracles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphListOracles`: GraphListOracles200Response
	fmt.Fprintf(os.Stdout, "Response from `GraphOraclesAPI.GraphListOracles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGraphListOraclesRequest struct via the builder pattern


### Return type

[**GraphListOracles200Response**](GraphListOracles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

