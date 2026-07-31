# \OraclesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Oracles**](OraclesAPI.md#CloudGetV1Oracles) | **Get** /v1/oracles | ListOracles reports the on-chain price/data oracles from the graph&#39;s O-Chain PriceFeed registry.



## CloudGetV1Oracles

> CloudOraclesOut CloudGetV1Oracles(ctx).Execute()

ListOracles reports the on-chain price/data oracles from the graph's O-Chain PriceFeed registry.



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
	resp, r, err := apiClient.OraclesAPI.CloudGetV1Oracles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OraclesAPI.CloudGetV1Oracles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Oracles`: CloudOraclesOut
	fmt.Fprintf(os.Stdout, "Response from `OraclesAPI.CloudGetV1Oracles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1OraclesRequest struct via the builder pattern


### Return type

[**CloudOraclesOut**](CloudOraclesOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

