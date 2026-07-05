# \ConsoleScoreConfigsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleCreateScoreConfig**](ConsoleScoreConfigsAPI.md#ConsoleCreateScoreConfig) | **Post** /v1/console/score-configs | Create a score config
[**ConsoleGetScoreConfig**](ConsoleScoreConfigsAPI.md#ConsoleGetScoreConfig) | **Get** /v1/console/score-configs/{configId} | Get a score config
[**ConsoleListScoreConfigs**](ConsoleScoreConfigsAPI.md#ConsoleListScoreConfigs) | **Get** /v1/console/score-configs | Get all score configs
[**ConsoleUpdateScoreConfig**](ConsoleScoreConfigsAPI.md#ConsoleUpdateScoreConfig) | **Patch** /v1/console/score-configs/{configId} | Update a score config



## ConsoleCreateScoreConfig

> ConsoleScoreConfig ConsoleCreateScoreConfig(ctx).ConsoleCreateScoreConfigRequest(consoleCreateScoreConfigRequest).Execute()

Create a score config

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
	consoleCreateScoreConfigRequest := *openapiclient.NewConsoleCreateScoreConfigRequest("Name_example", "DataType_example") // ConsoleCreateScoreConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoreConfigsAPI.ConsoleCreateScoreConfig(context.Background()).ConsoleCreateScoreConfigRequest(consoleCreateScoreConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoreConfigsAPI.ConsoleCreateScoreConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleCreateScoreConfig`: ConsoleScoreConfig
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoreConfigsAPI.ConsoleCreateScoreConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleCreateScoreConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consoleCreateScoreConfigRequest** | [**ConsoleCreateScoreConfigRequest**](ConsoleCreateScoreConfigRequest.md) |  | 

### Return type

[**ConsoleScoreConfig**](ConsoleScoreConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleGetScoreConfig

> ConsoleScoreConfig ConsoleGetScoreConfig(ctx, configId).Execute()

Get a score config

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
	configId := "configId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoreConfigsAPI.ConsoleGetScoreConfig(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoreConfigsAPI.ConsoleGetScoreConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleGetScoreConfig`: ConsoleScoreConfig
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoreConfigsAPI.ConsoleGetScoreConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleGetScoreConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConsoleScoreConfig**](ConsoleScoreConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleListScoreConfigs

> ConsoleListScoreConfigs200Response ConsoleListScoreConfigs(ctx).Page(page).Limit(limit).Execute()

Get all score configs

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
	page := int32(56) // int32 |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoreConfigsAPI.ConsoleListScoreConfigs(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoreConfigsAPI.ConsoleListScoreConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleListScoreConfigs`: ConsoleListScoreConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoreConfigsAPI.ConsoleListScoreConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleListScoreConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**ConsoleListScoreConfigs200Response**](ConsoleListScoreConfigs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleUpdateScoreConfig

> ConsoleScoreConfig ConsoleUpdateScoreConfig(ctx, configId).ConsoleUpdateScoreConfigRequest(consoleUpdateScoreConfigRequest).Execute()

Update a score config

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
	configId := "configId_example" // string | 
	consoleUpdateScoreConfigRequest := *openapiclient.NewConsoleUpdateScoreConfigRequest() // ConsoleUpdateScoreConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleScoreConfigsAPI.ConsoleUpdateScoreConfig(context.Background(), configId).ConsoleUpdateScoreConfigRequest(consoleUpdateScoreConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleScoreConfigsAPI.ConsoleUpdateScoreConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleUpdateScoreConfig`: ConsoleScoreConfig
	fmt.Fprintf(os.Stdout, "Response from `ConsoleScoreConfigsAPI.ConsoleUpdateScoreConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsoleUpdateScoreConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **consoleUpdateScoreConfigRequest** | [**ConsoleUpdateScoreConfigRequest**](ConsoleUpdateScoreConfigRequest.md) |  | 

### Return type

[**ConsoleScoreConfig**](ConsoleScoreConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

