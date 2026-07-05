# \ZtEdgeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ZtListEdgeNodes**](ZtEdgeAPI.md#ZtListEdgeNodes) | **Get** /v1/edge/nodes | List the org&#39;s ZT edge-routers



## ZtListEdgeNodes

> ZtListEdgeNodes200Response ZtListEdgeNodes(ctx).Execute()

List the org's ZT edge-routers

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
	resp, r, err := apiClient.ZtEdgeAPI.ZtListEdgeNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZtEdgeAPI.ZtListEdgeNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ZtListEdgeNodes`: ZtListEdgeNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `ZtEdgeAPI.ZtListEdgeNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiZtListEdgeNodesRequest struct via the builder pattern


### Return type

[**ZtListEdgeNodes200Response**](ZtListEdgeNodes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

