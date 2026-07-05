# \EvalsEvaluatorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EvalsEvaluatorsPost**](EvalsEvaluatorsAPI.md#V1EvalsEvaluatorsPost) | **Post** /v1/evals/evaluators | Register an evaluator (pre-built metric or custom rubric)



## V1EvalsEvaluatorsPost

> V1EvalsEvaluatorsPost(ctx).EvalsEvaluator(evalsEvaluator).Execute()

Register an evaluator (pre-built metric or custom rubric)

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
	evalsEvaluator := *openapiclient.NewEvalsEvaluator("Name_example") // EvalsEvaluator | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EvalsEvaluatorsAPI.V1EvalsEvaluatorsPost(context.Background()).EvalsEvaluator(evalsEvaluator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsEvaluatorsAPI.V1EvalsEvaluatorsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsEvaluatorsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalsEvaluator** | [**EvalsEvaluator**](EvalsEvaluator.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

