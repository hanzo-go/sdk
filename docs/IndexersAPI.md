# \IndexersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Indexers**](IndexersAPI.md#CloudGetV1Indexers) | **Get** /v1/indexers | ListIndexers reports the deployment&#39;s chain indexer(s) and how far each has indexed.



## CloudGetV1Indexers

> CloudIndexersOut CloudGetV1Indexers(ctx).Execute()

ListIndexers reports the deployment's chain indexer(s) and how far each has indexed.



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
	resp, r, err := apiClient.IndexersAPI.CloudGetV1Indexers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexersAPI.CloudGetV1Indexers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Indexers`: CloudIndexersOut
	fmt.Fprintf(os.Stdout, "Response from `IndexersAPI.CloudGetV1Indexers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IndexersRequest struct via the builder pattern


### Return type

[**CloudIndexersOut**](CloudIndexersOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

