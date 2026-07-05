# \PlatformDeploymentAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDeploymentAll**](PlatformDeploymentAPI.md#PlatformDeploymentAll) | **Get** /v1/platform/deployment/all | List deployments for an application
[**PlatformDeploymentAllByCompose**](PlatformDeploymentAPI.md#PlatformDeploymentAllByCompose) | **Get** /v1/platform/deployment/allByCompose | List deployments for a compose service
[**PlatformDeploymentAllByServer**](PlatformDeploymentAPI.md#PlatformDeploymentAllByServer) | **Get** /v1/platform/deployment/allByServer | List deployments on a server
[**PlatformDeploymentAllByType**](PlatformDeploymentAPI.md#PlatformDeploymentAllByType) | **Get** /v1/platform/deployment/allByType | List deployments by resource type and ID
[**PlatformDeploymentKillProcess**](PlatformDeploymentAPI.md#PlatformDeploymentKillProcess) | **Post** /v1/platform/deployment/killProcess | Kill a running deployment process
[**PlatformWebhookDeployApplication**](PlatformDeploymentAPI.md#PlatformWebhookDeployApplication) | **Post** /v1/platform/deploy/{refreshToken} | Webhook to trigger application deployment
[**PlatformWebhookDeployCompose**](PlatformDeploymentAPI.md#PlatformWebhookDeployCompose) | **Post** /v1/platform/deploy/compose/{refreshToken} | Webhook to trigger compose deployment
[**PlatformWebhookGithub**](PlatformDeploymentAPI.md#PlatformWebhookGithub) | **Post** /v1/platform/deploy/github | GitHub webhook receiver for auto-deploy



## PlatformDeploymentAll

> PlatformTRPCResult PlatformDeploymentAll(ctx).Input(input).Execute()

List deployments for an application

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeploymentAPI.PlatformDeploymentAll(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformDeploymentAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeploymentAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeploymentAPI.PlatformDeploymentAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeploymentAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformDeploymentAllByCompose

> PlatformTRPCResult PlatformDeploymentAllByCompose(ctx).Input(input).Execute()

List deployments for a compose service

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeploymentAPI.PlatformDeploymentAllByCompose(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformDeploymentAllByCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeploymentAllByCompose`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeploymentAPI.PlatformDeploymentAllByCompose`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeploymentAllByComposeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformDeploymentAllByServer

> PlatformTRPCResult PlatformDeploymentAllByServer(ctx).Input(input).Execute()

List deployments on a server

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeploymentAPI.PlatformDeploymentAllByServer(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformDeploymentAllByServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeploymentAllByServer`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeploymentAPI.PlatformDeploymentAllByServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeploymentAllByServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformDeploymentAllByType

> PlatformTRPCResult PlatformDeploymentAllByType(ctx).Input(input).Execute()

List deployments by resource type and ID

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
	input := "input_example" // string | URL-encoded JSON input for tRPC queries (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeploymentAPI.PlatformDeploymentAllByType(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformDeploymentAllByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeploymentAllByType`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeploymentAPI.PlatformDeploymentAllByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeploymentAllByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformDeploymentKillProcess

> PlatformTRPCResult PlatformDeploymentKillProcess(ctx).PlatformDeploymentKillProcessRequest(platformDeploymentKillProcessRequest).Execute()

Kill a running deployment process

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
	platformDeploymentKillProcessRequest := *openapiclient.NewPlatformDeploymentKillProcessRequest() // PlatformDeploymentKillProcessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDeploymentAPI.PlatformDeploymentKillProcess(context.Background()).PlatformDeploymentKillProcessRequest(platformDeploymentKillProcessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformDeploymentKillProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDeploymentKillProcess`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDeploymentAPI.PlatformDeploymentKillProcess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDeploymentKillProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDeploymentKillProcessRequest** | [**PlatformDeploymentKillProcessRequest**](PlatformDeploymentKillProcessRequest.md) |  | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformWebhookDeployApplication

> PlatformWebhookDeployApplication(ctx, refreshToken).Execute()

Webhook to trigger application deployment

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
	refreshToken := "refreshToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformDeploymentAPI.PlatformWebhookDeployApplication(context.Background(), refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformWebhookDeployApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refreshToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformWebhookDeployApplicationRequest struct via the builder pattern


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


## PlatformWebhookDeployCompose

> PlatformWebhookDeployCompose(ctx, refreshToken).Execute()

Webhook to trigger compose deployment

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
	refreshToken := "refreshToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformDeploymentAPI.PlatformWebhookDeployCompose(context.Background(), refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformWebhookDeployCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refreshToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformWebhookDeployComposeRequest struct via the builder pattern


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


## PlatformWebhookGithub

> PlatformWebhookGithub(ctx).Body(body).Execute()

GitHub webhook receiver for auto-deploy

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
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlatformDeploymentAPI.PlatformWebhookGithub(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDeploymentAPI.PlatformWebhookGithub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformWebhookGithubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

