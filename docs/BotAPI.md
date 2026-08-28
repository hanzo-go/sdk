# \BotAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBotRuns**](BotAPI.md#GetBotRuns) | **Get** /v1/bot/runs | List returns the caller org&#39;s live bot runs, read from the bot runtime and projected into the console contract with each run&#39;s live session URL derived here.
[**PostBotRuns**](BotAPI.md#PostBotRuns) | **Post** /v1/bot/runs | Answers 501 to every call: launching a bot run is not implemented.
[**PostBotRunsByRunidStop**](BotAPI.md#PostBotRunsByRunidStop) | **Post** /v1/bot/runs/{runId}/stop | Stop terminates one of the caller org&#39;s own bot runs and reports its terminal state.



## GetBotRuns

> BotRuns GetBotRuns(ctx).Execute()

List returns the caller org's live bot runs, read from the bot runtime and projected into the console contract with each run's live session URL derived here.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.GetBotRuns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotRuns`: BotRuns
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotRuns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotRunsRequest struct via the builder pattern


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


## PostBotRuns

> PostBotRuns(ctx).Execute()

Answers 501 to every call: launching a bot run is not implemented.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotAPI.PostBotRuns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotRunsRequest struct via the builder pattern


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


## PostBotRunsByRunidStop

> BotStopped PostBotRunsByRunidStop(ctx, runId).Execute()

Stop terminates one of the caller org's own bot runs and reports its terminal state.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	runId := "runId_example" // string | RunID is the run to stop, as the bot runtime named it. It is read from the URL — the `{runId}` segment the router matched on — and a body carrying a different id cannot redirect the stop.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotAPI.PostBotRunsByRunidStop(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotRunsByRunidStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBotRunsByRunidStop`: BotStopped
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.PostBotRunsByRunidStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | RunID is the run to stop, as the bot runtime named it. It is read from the URL — the &#x60;{runId}&#x60; segment the router matched on — and a body carrying a different id cannot redirect the stop. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotRunsByRunidStopRequest struct via the builder pattern


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

