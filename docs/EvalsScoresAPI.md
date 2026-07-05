# \EvalsScoresAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EvalsScoresGet**](EvalsScoresAPI.md#V1EvalsScoresGet) | **Get** /v1/evals/scores | List scores



## V1EvalsScoresGet

> V1EvalsScoresGet200Response V1EvalsScoresGet(ctx).RunName(runName).Limit(limit).Execute()

List scores

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
	runName := "runName_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsScoresAPI.V1EvalsScoresGet(context.Background()).RunName(runName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsScoresAPI.V1EvalsScoresGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EvalsScoresGet`: V1EvalsScoresGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EvalsScoresAPI.V1EvalsScoresGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsScoresGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runName** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**V1EvalsScoresGet200Response**](V1EvalsScoresGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

