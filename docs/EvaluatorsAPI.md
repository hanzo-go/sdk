# \EvaluatorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvalsPostV1EvalsEvaluators**](EvaluatorsAPI.md#EvalsPostV1EvalsEvaluators) | **Post** /v1/evals/evaluators | Register an evaluator (pre-built metric or custom rubric)



## EvalsPostV1EvalsEvaluators

> EvalsPostV1EvalsEvaluators(ctx).EvalsEvaluator(evalsEvaluator).Execute()

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
	r, err := apiClient.EvaluatorsAPI.EvalsPostV1EvalsEvaluators(context.Background()).EvalsEvaluator(evalsEvaluator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvaluatorsAPI.EvalsPostV1EvalsEvaluators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvalsPostV1EvalsEvaluatorsRequest struct via the builder pattern


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

