# \OraclesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOracles**](OraclesAPI.md#GetOracles) | **Get** /v1/oracles | Reports the on-chain price/data oracles from the graph&#39;s O-Chain PriceFeed registry.



## GetOracles

> OraclesOut GetOracles(ctx).Execute()

Reports the on-chain price/data oracles from the graph's O-Chain PriceFeed registry.



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
	resp, r, err := apiClient.OraclesAPI.GetOracles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OraclesAPI.GetOracles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOracles`: OraclesOut
	fmt.Fprintf(os.Stdout, "Response from `OraclesAPI.GetOracles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOraclesRequest struct via the builder pattern


### Return type

[**OraclesOut**](OraclesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

