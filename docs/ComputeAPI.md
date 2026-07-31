# \ComputeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteBot**](ComputeAPI.md#CloudDeleteBot) | **Delete** /v1/compute/bots/{id} | Tears down both halves of a bot: it unbinds the agent (best-effort — a bot with no binding still deletes), then terminates the machine.
[**CloudGetBot**](ComputeAPI.md#CloudGetBot) | **Get** /v1/compute/bots/{id} | Returns one of the caller org&#39;s bot machines with its agent binding.
[**CloudGetV1ComputeRegions**](ComputeAPI.md#CloudGetV1ComputeRegions) | **Get** /v1/compute/regions | 
[**CloudGetV1ComputeSizes**](ComputeAPI.md#CloudGetV1ComputeSizes) | **Get** /v1/compute/sizes | 
[**CloudListBots**](ComputeAPI.md#CloudListBots) | **Get** /v1/compute/bots | Returns the caller org&#39;s bot machines — the kind&#x3D;bot machines — each joined with the agent binding that says which cloud Agent it runs.
[**CloudPostV1ComputeBotsByIdByAction**](ComputeAPI.md#CloudPostV1ComputeBotsByIdByAction) | **Post** /v1/compute/bots/{id}/{action} | 
[**CloudPostV1ComputeBotsLaunch**](ComputeAPI.md#CloudPostV1ComputeBotsLaunch) | **Post** /v1/compute/bots/launch | 



## CloudDeleteBot

> CloudDeleteBot(ctx, id).Execute()

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
	r, err := apiClient.ComputeAPI.CloudDeleteBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudDeleteBot``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetBot

> CloudBotView CloudGetBot(ctx, id).Execute()

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
	resp, r, err := apiClient.ComputeAPI.CloudGetBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudGetBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetBot`: CloudBotView
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.CloudGetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the bot machine&#39;s id — the same id the machines surface addresses it by. Scoped to the caller&#39;s org upstream, so another tenant&#39;s id is 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudBotView**](CloudBotView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComputeRegions

> CloudGetV1ComputeRegions(ctx).Execute()



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
	r, err := apiClient.ComputeAPI.CloudGetV1ComputeRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudGetV1ComputeRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComputeRegionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ComputeSizes

> CloudGetV1ComputeSizes(ctx).Execute()



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
	r, err := apiClient.ComputeAPI.CloudGetV1ComputeSizes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudGetV1ComputeSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ComputeSizesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudListBots

> CloudBotList CloudListBots(ctx).Execute()

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
	resp, r, err := apiClient.ComputeAPI.CloudListBots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudListBots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudListBots`: CloudBotList
	fmt.Fprintf(os.Stdout, "Response from `ComputeAPI.CloudListBots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudListBotsRequest struct via the builder pattern


### Return type

[**CloudBotList**](CloudBotList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComputeBotsByIdByAction

> CloudPostV1ComputeBotsByIdByAction(ctx, id, action).Execute()



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
	r, err := apiClient.ComputeAPI.CloudPostV1ComputeBotsByIdByAction(context.Background(), id, action).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudPostV1ComputeBotsByIdByAction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1ComputeBotsByIdByActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ComputeBotsLaunch

> CloudPostV1ComputeBotsLaunch(ctx).Execute()



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
	r, err := apiClient.ComputeAPI.CloudPostV1ComputeBotsLaunch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComputeAPI.CloudPostV1ComputeBotsLaunch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ComputeBotsLaunchRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

