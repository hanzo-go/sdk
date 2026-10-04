# \BotAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBotMembers**](BotAPI.md#GetBotMembers) | **Get** /v1/bot/members | Returns the caller org&#39;s bots as space members — each with the member account uuid and the Person reference the roster addresses it by.
[**GetBotRuns**](BotAPI.md#GetBotRuns) | **Get** /v1/bot/runs | Returns the caller org&#39;s live bot runs, read from the bot runtime and projected into the console contract with each run&#39;s live session URL derived here.
[**PostBotMembersSync**](BotAPI.md#PostBotMembersSync) | **Post** /v1/bot/members/sync | Re-projects the caller org&#39;s bots as members into every space of the org and removes the ones whose agent is gone.
[**PostBotRunsByRunidStop**](BotAPI.md#PostBotRunsByRunidStop) | **Post** /v1/bot/runs/{runId}/stop | Terminates one of the caller org&#39;s own bot runs and reports its terminal state.



## GetBotMembers

> BotBotRoster GetBotMembers(ctx).Execute()

Returns the caller org's bots as space members — each with the member account uuid and the Person reference the roster addresses it by.



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
	resp, r, err := apiClient.BotAPI.GetBotMembers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.GetBotMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBotMembers`: BotBotRoster
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotMembersRequest struct via the builder pattern


### Return type

[**BotBotRoster**](BotBotRoster.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBotRuns

> BotBotRuns GetBotRuns(ctx).Execute()

Returns the caller org's live bot runs, read from the bot runtime and projected into the console contract with each run's live session URL derived here.



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
	// response from `GetBotRuns`: BotBotRuns
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.GetBotRuns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotRunsRequest struct via the builder pattern


### Return type

[**BotBotRuns**](BotBotRuns.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotMembersSync

> BotBotSync PostBotMembersSync(ctx).Execute()

Re-projects the caller org's bots as members into every space of the org and removes the ones whose agent is gone.



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
	resp, r, err := apiClient.BotAPI.PostBotMembersSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotAPI.PostBotMembersSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBotMembersSync`: BotBotSync
	fmt.Fprintf(os.Stdout, "Response from `BotAPI.PostBotMembersSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostBotMembersSyncRequest struct via the builder pattern


### Return type

[**BotBotSync**](BotBotSync.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBotRunsByRunidStop

> BotBotStopped PostBotRunsByRunidStop(ctx, runId).Execute()

Terminates one of the caller org's own bot runs and reports its terminal state.



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
	// response from `PostBotRunsByRunidStop`: BotBotStopped
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

[**BotBotStopped**](BotBotStopped.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

