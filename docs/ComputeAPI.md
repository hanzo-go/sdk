# \ComputeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBot**](ComputeAPI.md#DeleteBot) | **Delete** /v1/compute/bots/{id} | Tears down both halves of a bot: it unbinds the agent (best-effort — a bot with no binding still deletes), then terminates the machine.
[**GetBot**](ComputeAPI.md#GetBot) | **Get** /v1/compute/bots/{id} | Returns one of the caller org&#39;s bot machines with its agent binding.
[**GetComputeRegions**](ComputeAPI.md#GetComputeRegions) | **Get** /v1/compute/regions | The regions a machine or GPU can be launched into
[**GetComputeSizes**](ComputeAPI.md#GetComputeSizes) | **Get** /v1/compute/sizes | The machine and GPU sizes that can be launched
[**ListBots**](ComputeAPI.md#ListBots) | **Get** /v1/compute/bots | Returns the caller org&#39;s bot machines — the kind&#x3D;bot machines — each joined with the agent binding that says which cloud Agent it runs.
[**PostComputeBotsByIdByAction**](ComputeAPI.md#PostComputeBotsByIdByAction) | **Post** /v1/compute/bots/{id}/{action} | Message a bot, or stop it, by naming the action in the path
[**PostComputeBotsLaunch**](ComputeAPI.md#PostComputeBotsLaunch) | **Post** /v1/compute/bots/launch | Launch a bot machine — an agent plus the machine that runs it — or price one



## DeleteBot

> DeleteBot(ctx, id).Execute()

Tears down both halves of a bot: it unbinds the agent (best-effort — a bot with no binding still deletes), then terminates the machine.



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
	id := "id_example" // string | ID is the bot machine's id — the same id the machines surface addresses it by. Scoped to the caller's org upstream, so another tenant's id is 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.DeleteBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.DeleteBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the bot machine&#39;s id — the same id the machines surface addresses it by. Scoped to the caller&#39;s org upstream, so another tenant&#39;s id is 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBot

> BotView GetBot(ctx, id).Execute()

Returns one of the caller org's bot machines with its agent binding.



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
	id := "id_example" // string | ID is the bot machine's id — the same id the machines surface addresses it by. Scoped to the caller's org upstream, so another tenant's id is 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComputeAPI.GetBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBot`: BotView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.GetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the bot machine&#39;s id — the same id the machines surface addresses it by. Scoped to the caller&#39;s org upstream, so another tenant&#39;s id is 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BotView**](BotView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputeRegions

> GetComputeRegions(ctx).Execute()

The regions a machine or GPU can be launched into



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
	r, err := apiClient.ComputeAPI.GetComputeRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetComputeRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputeRegionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComputeSizes

> GetComputeSizes(ctx).Execute()

The machine and GPU sizes that can be launched



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
	r, err := apiClient.ComputeAPI.GetComputeSizes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.GetComputeSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComputeSizesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBots

> BotList ListBots(ctx).Execute()

Returns the caller org's bot machines — the kind=bot machines — each joined with the agent binding that says which cloud Agent it runs.



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
	resp, r, err := apiClient.ComputeAPI.ListBots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.ListBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBots`: BotList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.ListBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBotsRequest struct via the builder pattern


### Return type

[**BotList**](BotList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComputeBotsByIdByAction

> PostComputeBotsByIdByAction(ctx, id, action).Execute()

Message a bot, or stop it, by naming the action in the path



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
	id := "id_example" // string | 
	action := "action_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComputeAPI.PostComputeBotsByIdByAction(context.Background(), id, action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.PostComputeBotsByIdByAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostComputeBotsByIdByActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostComputeBotsLaunch

> PostComputeBotsLaunch(ctx).Execute()

Launch a bot machine — an agent plus the machine that runs it — or price one



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
	r, err := apiClient.ComputeAPI.PostComputeBotsLaunch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.PostComputeBotsLaunch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostComputeBotsLaunchRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

