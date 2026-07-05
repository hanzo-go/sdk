# \EvalsRunsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EvalsHealthGet**](EvalsRunsAPI.md#V1EvalsHealthGet) | **Get** /v1/evals/health | Health check
[**V1EvalsRunsPost**](EvalsRunsAPI.md#V1EvalsRunsPost) | **Post** /v1/evals/runs | Run a dataset against a model with an LLM-as-a-Judge



## V1EvalsHealthGet

> V1EvalsHealthGet200Response V1EvalsHealthGet(ctx).Execute()

Health check

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
	resp, r, err := apiClient.EvalsRunsAPI.V1EvalsHealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsRunsAPI.V1EvalsHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EvalsHealthGet`: V1EvalsHealthGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EvalsRunsAPI.V1EvalsHealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsHealthGetRequest struct via the builder pattern


### Return type

[**V1EvalsHealthGet200Response**](V1EvalsHealthGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EvalsRunsPost

> EvalsRunSummary V1EvalsRunsPost(ctx).EvalsRunRequest(evalsRunRequest).Execute()

Run a dataset against a model with an LLM-as-a-Judge



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
	evalsRunRequest := *openapiclient.NewEvalsRunRequest("Dataset_example", "Model_example") // EvalsRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EvalsRunsAPI.V1EvalsRunsPost(context.Background()).EvalsRunRequest(evalsRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EvalsRunsAPI.V1EvalsRunsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EvalsRunsPost`: EvalsRunSummary
	fmt.Fprintf(os.Stdout, "Response from `EvalsRunsAPI.V1EvalsRunsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EvalsRunsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalsRunRequest** | [**EvalsRunRequest**](EvalsRunRequest.md) |  | 

### Return type

[**EvalsRunSummary**](EvalsRunSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

