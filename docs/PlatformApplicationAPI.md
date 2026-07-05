# \PlatformApplicationAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformApplicationCancelDeployment**](PlatformApplicationAPI.md#PlatformApplicationCancelDeployment) | **Post** /v1/platform/application/cancelDeployment | Cancel an in-progress deployment
[**PlatformApplicationCleanQueues**](PlatformApplicationAPI.md#PlatformApplicationCleanQueues) | **Post** /v1/platform/application/cleanQueues | Clear pending deployment jobs
[**PlatformApplicationCreate**](PlatformApplicationAPI.md#PlatformApplicationCreate) | **Post** /v1/platform/application/create | Create a new application
[**PlatformApplicationDelete**](PlatformApplicationAPI.md#PlatformApplicationDelete) | **Post** /v1/platform/application/delete | Delete an application and all associated resources
[**PlatformApplicationDeploy**](PlatformApplicationAPI.md#PlatformApplicationDeploy) | **Post** /v1/platform/application/deploy | Trigger a new deployment
[**PlatformApplicationDisconnectGitProvider**](PlatformApplicationAPI.md#PlatformApplicationDisconnectGitProvider) | **Post** /v1/platform/application/disconnectGitProvider | Disconnect all git providers from application
[**PlatformApplicationMarkRunning**](PlatformApplicationAPI.md#PlatformApplicationMarkRunning) | **Post** /v1/platform/application/markRunning | Mark application status as running
[**PlatformApplicationMove**](PlatformApplicationAPI.md#PlatformApplicationMove) | **Post** /v1/platform/application/move | Move application to a different environment
[**PlatformApplicationOne**](PlatformApplicationAPI.md#PlatformApplicationOne) | **Get** /v1/platform/application/one | Get application details
[**PlatformApplicationReadAppMonitoring**](PlatformApplicationAPI.md#PlatformApplicationReadAppMonitoring) | **Get** /v1/platform/application/readAppMonitoring | Read container monitoring stats
[**PlatformApplicationReadTraefikConfig**](PlatformApplicationAPI.md#PlatformApplicationReadTraefikConfig) | **Get** /v1/platform/application/readTraefikConfig | Read Traefik routing config for an application
[**PlatformApplicationRedeploy**](PlatformApplicationAPI.md#PlatformApplicationRedeploy) | **Post** /v1/platform/application/redeploy | Redeploy (rebuild) an application
[**PlatformApplicationRefreshToken**](PlatformApplicationAPI.md#PlatformApplicationRefreshToken) | **Post** /v1/platform/application/refreshToken | Regenerate webhook refresh token
[**PlatformApplicationReload**](PlatformApplicationAPI.md#PlatformApplicationReload) | **Post** /v1/platform/application/reload | Reload application container (re-mechanize)
[**PlatformApplicationSaveBitbucketProvider**](PlatformApplicationAPI.md#PlatformApplicationSaveBitbucketProvider) | **Post** /v1/platform/application/saveBitbucketProvider | Configure Bitbucket as source provider
[**PlatformApplicationSaveBuildType**](PlatformApplicationAPI.md#PlatformApplicationSaveBuildType) | **Post** /v1/platform/application/saveBuildType | Configure the build type
[**PlatformApplicationSaveDockerProvider**](PlatformApplicationAPI.md#PlatformApplicationSaveDockerProvider) | **Post** /v1/platform/application/saveDockerProvider | Configure Docker image as source
[**PlatformApplicationSaveEnvironment**](PlatformApplicationAPI.md#PlatformApplicationSaveEnvironment) | **Post** /v1/platform/application/saveEnvironment | Save environment variables and build args
[**PlatformApplicationSaveGitProdiver**](PlatformApplicationAPI.md#PlatformApplicationSaveGitProdiver) | **Post** /v1/platform/application/saveGitProdiver | Configure custom Git URL as source
[**PlatformApplicationSaveGiteaProvider**](PlatformApplicationAPI.md#PlatformApplicationSaveGiteaProvider) | **Post** /v1/platform/application/saveGiteaProvider | Configure Gitea as source provider
[**PlatformApplicationSaveGithubProvider**](PlatformApplicationAPI.md#PlatformApplicationSaveGithubProvider) | **Post** /v1/platform/application/saveGithubProvider | Configure GitHub as source provider
[**PlatformApplicationSaveGitlabProvider**](PlatformApplicationAPI.md#PlatformApplicationSaveGitlabProvider) | **Post** /v1/platform/application/saveGitlabProvider | Configure GitLab as source provider
[**PlatformApplicationStart**](PlatformApplicationAPI.md#PlatformApplicationStart) | **Post** /v1/platform/application/start | Start a stopped application
[**PlatformApplicationStop**](PlatformApplicationAPI.md#PlatformApplicationStop) | **Post** /v1/platform/application/stop | Stop a running application
[**PlatformApplicationUpdate**](PlatformApplicationAPI.md#PlatformApplicationUpdate) | **Post** /v1/platform/application/update | Update application configuration
[**PlatformApplicationUpdateTraefikConfig**](PlatformApplicationAPI.md#PlatformApplicationUpdateTraefikConfig) | **Post** /v1/platform/application/updateTraefikConfig | Write Traefik routing config



## PlatformApplicationCancelDeployment

> PlatformTRPCResult PlatformApplicationCancelDeployment(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Cancel an in-progress deployment

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationCancelDeployment(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationCancelDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationCancelDeployment`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationCancelDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationCancelDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationCleanQueues

> PlatformTRPCResult PlatformApplicationCleanQueues(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Clear pending deployment jobs

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationCleanQueues(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationCleanQueues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationCleanQueues`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationCleanQueues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationCleanQueuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationCreate

> PlatformTRPCResult PlatformApplicationCreate(ctx).PlatformApplicationCreateRequest(platformApplicationCreateRequest).Execute()

Create a new application

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
	platformApplicationCreateRequest := *openapiclient.NewPlatformApplicationCreateRequest() // PlatformApplicationCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationCreate(context.Background()).PlatformApplicationCreateRequest(platformApplicationCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCreateRequest** | [**PlatformApplicationCreateRequest**](PlatformApplicationCreateRequest.md) |  | 

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


## PlatformApplicationDelete

> PlatformTRPCResult PlatformApplicationDelete(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Delete an application and all associated resources

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationDelete(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationDelete`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationDeploy

> PlatformTRPCResult PlatformApplicationDeploy(ctx).PlatformApplicationDeployRequest(platformApplicationDeployRequest).Execute()

Trigger a new deployment

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
	platformApplicationDeployRequest := *openapiclient.NewPlatformApplicationDeployRequest() // PlatformApplicationDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationDeploy(context.Background()).PlatformApplicationDeployRequest(platformApplicationDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationDeploy`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationDeployRequest** | [**PlatformApplicationDeployRequest**](PlatformApplicationDeployRequest.md) |  | 

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


## PlatformApplicationDisconnectGitProvider

> PlatformTRPCResult PlatformApplicationDisconnectGitProvider(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Disconnect all git providers from application

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationDisconnectGitProvider(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationDisconnectGitProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationDisconnectGitProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationDisconnectGitProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationDisconnectGitProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationMarkRunning

> PlatformTRPCResult PlatformApplicationMarkRunning(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Mark application status as running

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationMarkRunning(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationMarkRunning``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationMarkRunning`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationMarkRunning`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationMarkRunningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationMove

> PlatformTRPCResult PlatformApplicationMove(ctx).PlatformApplicationMoveRequest(platformApplicationMoveRequest).Execute()

Move application to a different environment

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
	platformApplicationMoveRequest := *openapiclient.NewPlatformApplicationMoveRequest() // PlatformApplicationMoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationMove(context.Background()).PlatformApplicationMoveRequest(platformApplicationMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationMove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationMove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationMoveRequest** | [**PlatformApplicationMoveRequest**](PlatformApplicationMoveRequest.md) |  | 

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


## PlatformApplicationOne

> PlatformTRPCResult PlatformApplicationOne(ctx).Input(input).Execute()

Get application details

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
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationOneRequest struct via the builder pattern


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


## PlatformApplicationReadAppMonitoring

> PlatformTRPCResult PlatformApplicationReadAppMonitoring(ctx).Input(input).Execute()

Read container monitoring stats

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
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationReadAppMonitoring(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationReadAppMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationReadAppMonitoring`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationReadAppMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationReadAppMonitoringRequest struct via the builder pattern


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


## PlatformApplicationReadTraefikConfig

> PlatformTRPCResult PlatformApplicationReadTraefikConfig(ctx).Input(input).Execute()

Read Traefik routing config for an application

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
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationReadTraefikConfig(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationReadTraefikConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationReadTraefikConfig`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationReadTraefikConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationReadTraefikConfigRequest struct via the builder pattern


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


## PlatformApplicationRedeploy

> PlatformTRPCResult PlatformApplicationRedeploy(ctx).PlatformApplicationDeployRequest(platformApplicationDeployRequest).Execute()

Redeploy (rebuild) an application

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
	platformApplicationDeployRequest := *openapiclient.NewPlatformApplicationDeployRequest() // PlatformApplicationDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationRedeploy(context.Background()).PlatformApplicationDeployRequest(platformApplicationDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationRedeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationRedeploy`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationRedeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationRedeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationDeployRequest** | [**PlatformApplicationDeployRequest**](PlatformApplicationDeployRequest.md) |  | 

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


## PlatformApplicationRefreshToken

> PlatformTRPCResult PlatformApplicationRefreshToken(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Regenerate webhook refresh token

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationRefreshToken(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationRefreshToken`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationReload

> PlatformTRPCResult PlatformApplicationReload(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Reload application container (re-mechanize)

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationReload(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationReload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationReload`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationReload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationReloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationSaveBitbucketProvider

> PlatformTRPCResult PlatformApplicationSaveBitbucketProvider(ctx).PlatformApplicationSaveBitbucketProviderRequest(platformApplicationSaveBitbucketProviderRequest).Execute()

Configure Bitbucket as source provider

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
	platformApplicationSaveBitbucketProviderRequest := *openapiclient.NewPlatformApplicationSaveBitbucketProviderRequest() // PlatformApplicationSaveBitbucketProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveBitbucketProvider(context.Background()).PlatformApplicationSaveBitbucketProviderRequest(platformApplicationSaveBitbucketProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveBitbucketProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveBitbucketProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveBitbucketProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveBitbucketProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveBitbucketProviderRequest** | [**PlatformApplicationSaveBitbucketProviderRequest**](PlatformApplicationSaveBitbucketProviderRequest.md) |  | 

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


## PlatformApplicationSaveBuildType

> PlatformTRPCResult PlatformApplicationSaveBuildType(ctx).PlatformApplicationSaveBuildTypeRequest(platformApplicationSaveBuildTypeRequest).Execute()

Configure the build type

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
	platformApplicationSaveBuildTypeRequest := *openapiclient.NewPlatformApplicationSaveBuildTypeRequest() // PlatformApplicationSaveBuildTypeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveBuildType(context.Background()).PlatformApplicationSaveBuildTypeRequest(platformApplicationSaveBuildTypeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveBuildType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveBuildType`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveBuildType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveBuildTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveBuildTypeRequest** | [**PlatformApplicationSaveBuildTypeRequest**](PlatformApplicationSaveBuildTypeRequest.md) |  | 

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


## PlatformApplicationSaveDockerProvider

> PlatformTRPCResult PlatformApplicationSaveDockerProvider(ctx).PlatformApplicationSaveDockerProviderRequest(platformApplicationSaveDockerProviderRequest).Execute()

Configure Docker image as source

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
	platformApplicationSaveDockerProviderRequest := *openapiclient.NewPlatformApplicationSaveDockerProviderRequest() // PlatformApplicationSaveDockerProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveDockerProvider(context.Background()).PlatformApplicationSaveDockerProviderRequest(platformApplicationSaveDockerProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveDockerProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveDockerProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveDockerProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveDockerProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveDockerProviderRequest** | [**PlatformApplicationSaveDockerProviderRequest**](PlatformApplicationSaveDockerProviderRequest.md) |  | 

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


## PlatformApplicationSaveEnvironment

> PlatformTRPCResult PlatformApplicationSaveEnvironment(ctx).PlatformApplicationSaveEnvironmentRequest(platformApplicationSaveEnvironmentRequest).Execute()

Save environment variables and build args

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
	platformApplicationSaveEnvironmentRequest := *openapiclient.NewPlatformApplicationSaveEnvironmentRequest() // PlatformApplicationSaveEnvironmentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveEnvironment(context.Background()).PlatformApplicationSaveEnvironmentRequest(platformApplicationSaveEnvironmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveEnvironment`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveEnvironmentRequest** | [**PlatformApplicationSaveEnvironmentRequest**](PlatformApplicationSaveEnvironmentRequest.md) |  | 

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


## PlatformApplicationSaveGitProdiver

> PlatformTRPCResult PlatformApplicationSaveGitProdiver(ctx).PlatformApplicationSaveGitProdiverRequest(platformApplicationSaveGitProdiverRequest).Execute()

Configure custom Git URL as source

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
	platformApplicationSaveGitProdiverRequest := *openapiclient.NewPlatformApplicationSaveGitProdiverRequest() // PlatformApplicationSaveGitProdiverRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveGitProdiver(context.Background()).PlatformApplicationSaveGitProdiverRequest(platformApplicationSaveGitProdiverRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveGitProdiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveGitProdiver`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveGitProdiver`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveGitProdiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveGitProdiverRequest** | [**PlatformApplicationSaveGitProdiverRequest**](PlatformApplicationSaveGitProdiverRequest.md) |  | 

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


## PlatformApplicationSaveGiteaProvider

> PlatformTRPCResult PlatformApplicationSaveGiteaProvider(ctx).PlatformApplicationSaveGiteaProviderRequest(platformApplicationSaveGiteaProviderRequest).Execute()

Configure Gitea as source provider

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
	platformApplicationSaveGiteaProviderRequest := *openapiclient.NewPlatformApplicationSaveGiteaProviderRequest() // PlatformApplicationSaveGiteaProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveGiteaProvider(context.Background()).PlatformApplicationSaveGiteaProviderRequest(platformApplicationSaveGiteaProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveGiteaProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveGiteaProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveGiteaProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveGiteaProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveGiteaProviderRequest** | [**PlatformApplicationSaveGiteaProviderRequest**](PlatformApplicationSaveGiteaProviderRequest.md) |  | 

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


## PlatformApplicationSaveGithubProvider

> PlatformTRPCResult PlatformApplicationSaveGithubProvider(ctx).PlatformApplicationSaveGithubProviderRequest(platformApplicationSaveGithubProviderRequest).Execute()

Configure GitHub as source provider

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
	platformApplicationSaveGithubProviderRequest := *openapiclient.NewPlatformApplicationSaveGithubProviderRequest() // PlatformApplicationSaveGithubProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveGithubProvider(context.Background()).PlatformApplicationSaveGithubProviderRequest(platformApplicationSaveGithubProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveGithubProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveGithubProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveGithubProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveGithubProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveGithubProviderRequest** | [**PlatformApplicationSaveGithubProviderRequest**](PlatformApplicationSaveGithubProviderRequest.md) |  | 

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


## PlatformApplicationSaveGitlabProvider

> PlatformTRPCResult PlatformApplicationSaveGitlabProvider(ctx).PlatformApplicationSaveGitlabProviderRequest(platformApplicationSaveGitlabProviderRequest).Execute()

Configure GitLab as source provider

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
	platformApplicationSaveGitlabProviderRequest := *openapiclient.NewPlatformApplicationSaveGitlabProviderRequest() // PlatformApplicationSaveGitlabProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationSaveGitlabProvider(context.Background()).PlatformApplicationSaveGitlabProviderRequest(platformApplicationSaveGitlabProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationSaveGitlabProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationSaveGitlabProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationSaveGitlabProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationSaveGitlabProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationSaveGitlabProviderRequest** | [**PlatformApplicationSaveGitlabProviderRequest**](PlatformApplicationSaveGitlabProviderRequest.md) |  | 

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


## PlatformApplicationStart

> PlatformTRPCResult PlatformApplicationStart(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Start a stopped application

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationStart(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationStart`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationStop

> PlatformTRPCResult PlatformApplicationStop(ctx).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()

Stop a running application

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
	platformApplicationCancelDeploymentRequest := *openapiclient.NewPlatformApplicationCancelDeploymentRequest() // PlatformApplicationCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationStop(context.Background()).PlatformApplicationCancelDeploymentRequest(platformApplicationCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationStop`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationStop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationCancelDeploymentRequest** | [**PlatformApplicationCancelDeploymentRequest**](PlatformApplicationCancelDeploymentRequest.md) |  | 

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


## PlatformApplicationUpdate

> PlatformTRPCResult PlatformApplicationUpdate(ctx).PlatformApplicationUpdateRequest(platformApplicationUpdateRequest).Execute()

Update application configuration

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
	platformApplicationUpdateRequest := *openapiclient.NewPlatformApplicationUpdateRequest() // PlatformApplicationUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationUpdate(context.Background()).PlatformApplicationUpdateRequest(platformApplicationUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationUpdateRequest** | [**PlatformApplicationUpdateRequest**](PlatformApplicationUpdateRequest.md) |  | 

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


## PlatformApplicationUpdateTraefikConfig

> PlatformTRPCResult PlatformApplicationUpdateTraefikConfig(ctx).PlatformApplicationUpdateTraefikConfigRequest(platformApplicationUpdateTraefikConfigRequest).Execute()

Write Traefik routing config

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
	platformApplicationUpdateTraefikConfigRequest := *openapiclient.NewPlatformApplicationUpdateTraefikConfigRequest() // PlatformApplicationUpdateTraefikConfigRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformApplicationAPI.PlatformApplicationUpdateTraefikConfig(context.Background()).PlatformApplicationUpdateTraefikConfigRequest(platformApplicationUpdateTraefikConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformApplicationAPI.PlatformApplicationUpdateTraefikConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformApplicationUpdateTraefikConfig`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformApplicationAPI.PlatformApplicationUpdateTraefikConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformApplicationUpdateTraefikConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformApplicationUpdateTraefikConfigRequest** | [**PlatformApplicationUpdateTraefikConfigRequest**](PlatformApplicationUpdateTraefikConfigRequest.md) |  | 

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

