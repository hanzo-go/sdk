# \IndexersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIndexers**](IndexersAPI.md#GetIndexers) | **Get** /v1/indexers | Reports the deployment&#39;s chain indexer(s) and how far each has indexed.



## GetIndexers

> IndexersOut GetIndexers(ctx).Execute()

Reports the deployment's chain indexer(s) and how far each has indexed.



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
	resp, r, err := apiClient.IndexersAPI.GetIndexers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IndexersAPI.GetIndexers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIndexers`: IndexersOut
	fmt.Fprintf(os.Stdout, "Response from `IndexersAPI.GetIndexers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndexersRequest struct via the builder pattern


### Return type

[**IndexersOut**](IndexersOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

