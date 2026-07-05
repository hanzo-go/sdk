# \ConsoleHealthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleGetHealth**](ConsoleHealthAPI.md#ConsoleGetHealth) | **Get** /v1/console/health | Check health of API and database



## ConsoleGetHealth

> ConsoleHealthResponse ConsoleGetHealth(ctx).Execute()

Check health of API and database

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
	resp, r, err := apiClient.ConsoleHealthAPI.ConsoleGetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleHealthAPI.ConsoleGetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetHealth`: ConsoleHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `ConsoleHealthAPI.ConsoleGetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetHealthRequest struct via the builder pattern


### Return type

[**ConsoleHealthResponse**](ConsoleHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

