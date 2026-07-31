# \TrainingContributionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetTrainingContribution**](TrainingContributionAPI.md#AiGetTrainingContribution) | **Get** /v1/ai/training-contribution | Training Contribution
[**AiPutUpdateTrainingContribution**](TrainingContributionAPI.md#AiPutUpdateTrainingContribution) | **Put** /v1/ai/training-contribution | Training Contribution
[**AiUpdateTrainingContribution**](TrainingContributionAPI.md#AiUpdateTrainingContribution) | **Patch** /v1/ai/training-contribution | Training Contribution



## AiGetTrainingContribution

> AiEnvelope AiGetTrainingContribution(ctx).Execute()

Training Contribution

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
	resp, r, err := apiClient.TrainingContributionAPI.AiGetTrainingContribution(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingContributionAPI.AiGetTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetTrainingContribution`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingContributionAPI.AiGetTrainingContribution`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetTrainingContributionRequest struct via the builder pattern


### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiPutUpdateTrainingContribution

> AiEnvelope AiPutUpdateTrainingContribution(ctx).Body(body).Execute()

Training Contribution

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingContributionAPI.AiPutUpdateTrainingContribution(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingContributionAPI.AiPutUpdateTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiPutUpdateTrainingContribution`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingContributionAPI.AiPutUpdateTrainingContribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiPutUpdateTrainingContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AiUpdateTrainingContribution

> AiEnvelope AiUpdateTrainingContribution(ctx).Body(body).Execute()

Training Contribution

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrainingContributionAPI.AiUpdateTrainingContribution(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrainingContributionAPI.AiUpdateTrainingContribution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiUpdateTrainingContribution`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `TrainingContributionAPI.AiUpdateTrainingContribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAiUpdateTrainingContributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**AiEnvelope**](AiEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

