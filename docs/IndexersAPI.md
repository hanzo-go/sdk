# \IndexersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GraphListIndexers**](IndexersAPI.md#GraphListIndexers) | **Get** /v1/indexers | List the deployment&#39;s chain indexer(s)



## GraphListIndexers

> GraphListIndexers200Response GraphListIndexers(ctx).Execute()

List the deployment's chain indexer(s)

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
	resp, r, err := apiClient.IndexersAPI.GraphListIndexers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexersAPI.GraphListIndexers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GraphListIndexers`: GraphListIndexers200Response
	fmt.Fprintf(os.Stdout, "Response from `IndexersAPI.GraphListIndexers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGraphListIndexersRequest struct via the builder pattern


### Return type

[**GraphListIndexers200Response**](GraphListIndexers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

