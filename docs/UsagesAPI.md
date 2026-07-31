# \UsagesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AiGetCloudUsages**](UsagesAPI.md#AiGetCloudUsages) | **Get** /v1/ai/usages/cloud | Cloud (usage)
[**AiGetRangeUsages**](UsagesAPI.md#AiGetRangeUsages) | **Get** /v1/ai/usages/range | Range (usage)
[**AiGetUsages**](UsagesAPI.md#AiGetUsages) | **Get** /v1/ai/usages | List usages



## AiGetCloudUsages

> AiEnvelope AiGetCloudUsages(ctx).Execute()

Cloud (usage)

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
	resp, r, err := apiClient.UsagesAPI.AiGetCloudUsages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsagesAPI.AiGetCloudUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetCloudUsages`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsagesAPI.AiGetCloudUsages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetCloudUsagesRequest struct via the builder pattern


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


## AiGetRangeUsages

> AiEnvelope AiGetRangeUsages(ctx).Execute()

Range (usage)

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
	resp, r, err := apiClient.UsagesAPI.AiGetRangeUsages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsagesAPI.AiGetRangeUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetRangeUsages`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsagesAPI.AiGetRangeUsages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetRangeUsagesRequest struct via the builder pattern


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


## AiGetUsages

> AiEnvelope AiGetUsages(ctx).Execute()

List usages



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
	resp, r, err := apiClient.UsagesAPI.AiGetUsages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsagesAPI.AiGetUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AiGetUsages`: AiEnvelope
	fmt.Fprintf(os.Stdout, "Response from `UsagesAPI.AiGetUsages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAiGetUsagesRequest struct via the builder pattern


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

