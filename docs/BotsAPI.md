# \BotsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBots**](BotsAPI.md#GetBots) | **Get** /v1/bots | List returns the caller org&#39;s live bot runs, read from the bot runtime and projected into the console contract with each run&#39;s live session URL derived here.
[**PostBotsByRunidStop**](BotsAPI.md#PostBotsByRunidStop) | **Post** /v1/bots/{runId}/stop | Stop terminates one of the caller org&#39;s own bot runs and reports its terminal state.
[**PostBotsRun**](BotsAPI.md#PostBotsRun) | **Post** /v1/bots/run | Reserved address for launching a bot run — not implemented, always 501



## GetBots

> BotRuns GetBots(ctx).Execute()

List returns the caller org's live bot runs, read from the bot runtime and projected into the console contract with each run's live session URL derived here.



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
	resp, r, err := apiClient.BotsAPI.GetBots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.GetBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBots`: BotRuns
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.GetBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotsRequest struct via the builder pattern


### Return type

[**BotRuns**](BotRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotsByRunidStop

> BotStopped PostBotsByRunidStop(ctx, runId).Execute()

Stop terminates one of the caller org's own bot runs and reports its terminal state.



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
	runId := "runId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotsAPI.PostBotsByRunidStop(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.PostBotsByRunidStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBotsByRunidStop`: BotStopped
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.PostBotsByRunidStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotsByRunidStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotStopped**](BotStopped.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotsRun

> PostBotsRun(ctx).Execute()

Reserved address for launching a bot run — not implemented, always 501



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
	r, err := apiClient.BotsAPI.PostBotsRun(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.PostBotsRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotsRunRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

