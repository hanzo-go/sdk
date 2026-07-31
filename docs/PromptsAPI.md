# \PromptsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromptsCreatePrompt**](PromptsAPI.md#PromptsCreatePrompt) | **Post** /v1/prompts | Create a prompt or append a new version
[**PromptsDeletePrompt**](PromptsAPI.md#PromptsDeletePrompt) | **Delete** /v1/prompts/{name} | Delete a prompt and its version history
[**PromptsGetPrompt**](PromptsAPI.md#PromptsGetPrompt) | **Get** /v1/prompts/{name} | Prompt detail + version history
[**PromptsListPrompts**](PromptsAPI.md#PromptsListPrompts) | **Get** /v1/prompts | List current prompts for the org
[**PromptsPromptMetrics**](PromptsAPI.md#PromptsPromptMetrics) | **Get** /v1/prompts/metrics | Real per-prompt statistics



## PromptsCreatePrompt

> PromptsPromptDetail PromptsCreatePrompt(ctx).PromptsCreatePrompt(promptsCreatePrompt).Execute()

Create a prompt or append a new version

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
	promptsCreatePrompt := *openapiclient.NewPromptsCreatePrompt("Name_example") // PromptsCreatePrompt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.PromptsCreatePrompt(context.Background()).PromptsCreatePrompt(promptsCreatePrompt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PromptsCreatePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsCreatePrompt`: PromptsPromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.PromptsCreatePrompt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromptsCreatePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promptsCreatePrompt** | [**PromptsCreatePrompt**](PromptsCreatePrompt.md) |  | 

### Return type

[**PromptsPromptDetail**](PromptsPromptDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsDeletePrompt

> PromptsDeletePrompt(ctx, name).Execute()

Delete a prompt and its version history

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PromptsAPI.PromptsDeletePrompt(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PromptsDeletePrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsDeletePromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsGetPrompt

> PromptsPromptDetail PromptsGetPrompt(ctx, name).Execute()

Prompt detail + version history

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromptsAPI.PromptsGetPrompt(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PromptsGetPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsGetPrompt`: PromptsPromptDetail
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.PromptsGetPrompt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsGetPromptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PromptsPromptDetail**](PromptsPromptDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsListPrompts

> PromptsListPrompts200Response PromptsListPrompts(ctx).Execute()

List current prompts for the org

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
	resp, r, err := apiClient.PromptsAPI.PromptsListPrompts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PromptsListPrompts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsListPrompts`: PromptsListPrompts200Response
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.PromptsListPrompts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsListPromptsRequest struct via the builder pattern


### Return type

[**PromptsListPrompts200Response**](PromptsListPrompts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromptsPromptMetrics

> PromptsPromptMetrics200Response PromptsPromptMetrics(ctx).Execute()

Real per-prompt statistics

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
	resp, r, err := apiClient.PromptsAPI.PromptsPromptMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromptsAPI.PromptsPromptMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromptsPromptMetrics`: PromptsPromptMetrics200Response
	fmt.Fprintf(os.Stdout, "Response from `PromptsAPI.PromptsPromptMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPromptsPromptMetricsRequest struct via the builder pattern


### Return type

[**PromptsPromptMetrics200Response**](PromptsPromptMetrics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

