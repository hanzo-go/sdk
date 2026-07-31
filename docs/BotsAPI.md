# \BotsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Bots**](BotsAPI.md#CloudGetV1Bots) | **Get** /v1/bots | List returns the caller org&#39;s live bot runs, read from the bot runtime and projected into the console contract with each run&#39;s live session URL derived here.
[**CloudPostV1BotsRun**](BotsAPI.md#CloudPostV1BotsRun) | **Post** /v1/bots/run | 
[**CloudPostV1BotsRunIdStop**](BotsAPI.md#CloudPostV1BotsRunIdStop) | **Post** /v1/bots/{runId}/stop | Stop terminates one of the caller org&#39;s own bot runs and reports its terminal state.
[**VisorBotAction**](BotsAPI.md#VisorBotAction) | **Post** /v1/bots/{id}/{action} | Act on a bot (stop, pause, or message)
[**VisorDeleteBot**](BotsAPI.md#VisorDeleteBot) | **Delete** /v1/bots/{id} | Terminate a bot (unbind agent + delete machine)
[**VisorGetBot**](BotsAPI.md#VisorGetBot) | **Get** /v1/bots/{id} | Get one bot by id
[**VisorLaunchBot**](BotsAPI.md#VisorLaunchBot) | **Post** /v1/bots/launch | Launch a bot (machine + agent binding), or dryRun for a quote



## CloudGetV1Bots

> CloudBotRuns CloudGetV1Bots(ctx).Execute()

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
	resp, r, err := apiClient.BotsAPI.CloudGetV1Bots(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.CloudGetV1Bots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Bots`: CloudBotRuns
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.CloudGetV1Bots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1BotsRequest struct via the builder pattern


### Return type

[**CloudBotRuns**](CloudBotRuns.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1BotsRun

> CloudPostV1BotsRun(ctx).Execute()



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
	r, err := apiClient.BotsAPI.CloudPostV1BotsRun(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.CloudPostV1BotsRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BotsRunRequest struct via the builder pattern


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


## CloudPostV1BotsRunIdStop

> CloudBotStopped CloudPostV1BotsRunIdStop(ctx, runId).Execute()

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
	resp, r, err := apiClient.BotsAPI.CloudPostV1BotsRunIdStop(context.Background(), runId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.CloudPostV1BotsRunIdStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1BotsRunIdStop`: CloudBotStopped
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.CloudPostV1BotsRunIdStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1BotsRunIdStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudBotStopped**](CloudBotStopped.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorBotAction

> map[string]interface{} VisorBotAction(ctx, id, action).Body(body).Execute()

Act on a bot (stop, pause, or message)



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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotsAPI.VisorBotAction(context.Background(), id, action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.VisorBotAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorBotAction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.VisorBotAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**action** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorBotActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorDeleteBot

> VisorDeleteBot(ctx, id).Execute()

Terminate a bot (unbind agent + delete machine)

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BotsAPI.VisorDeleteBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.VisorDeleteBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorDeleteBotRequest struct via the builder pattern


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


## VisorGetBot

> VisorBotView VisorGetBot(ctx, id).Execute()

Get one bot by id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotsAPI.VisorGetBot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.VisorGetBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorGetBot`: VisorBotView
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.VisorGetBot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVisorGetBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VisorBotView**](VisorBotView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VisorLaunchBot

> map[string]interface{} VisorLaunchBot(ctx).VisorBotLaunchRequest(visorBotLaunchRequest).Execute()

Launch a bot (machine + agent binding), or dryRun for a quote

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
	visorBotLaunchRequest := *openapiclient.NewVisorBotLaunchRequest() // VisorBotLaunchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BotsAPI.VisorLaunchBot(context.Background()).VisorBotLaunchRequest(visorBotLaunchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BotsAPI.VisorLaunchBot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VisorLaunchBot`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BotsAPI.VisorLaunchBot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVisorLaunchBotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **visorBotLaunchRequest** | [**VisorBotLaunchRequest**](VisorBotLaunchRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

