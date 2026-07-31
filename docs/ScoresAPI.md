# \ScoresAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvalsGetV1EvalsScores**](ScoresAPI.md#EvalsGetV1EvalsScores) | **Get** /v1/evals/scores | List scores



## EvalsGetV1EvalsScores

> EvalsGetV1EvalsScores200Response EvalsGetV1EvalsScores(ctx).RunName(runName).Limit(limit).Execute()

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
	resp, r, err := apiClient.ScoresAPI.EvalsGetV1EvalsScores(context.Background()).RunName(runName).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScoresAPI.EvalsGetV1EvalsScores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvalsGetV1EvalsScores`: EvalsGetV1EvalsScores200Response
	fmt.Fprintf(os.Stdout, "Response from `ScoresAPI.EvalsGetV1EvalsScores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvalsGetV1EvalsScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runName** | **string** |  | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**EvalsGetV1EvalsScores200Response**](EvalsGetV1EvalsScores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

