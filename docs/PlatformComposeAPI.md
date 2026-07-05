# \PlatformComposeAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformComposeCancelDeployment**](PlatformComposeAPI.md#PlatformComposeCancelDeployment) | **Post** /v1/platform/compose/cancelDeployment | Cancel in-progress compose deployment
[**PlatformComposeCreate**](PlatformComposeAPI.md#PlatformComposeCreate) | **Post** /v1/platform/compose/create | Create a Docker Compose service
[**PlatformComposeDelete**](PlatformComposeAPI.md#PlatformComposeDelete) | **Post** /v1/platform/compose/delete | Delete a compose service
[**PlatformComposeDeploy**](PlatformComposeAPI.md#PlatformComposeDeploy) | **Post** /v1/platform/compose/deploy | Deploy a compose service
[**PlatformComposeDeployTemplate**](PlatformComposeAPI.md#PlatformComposeDeployTemplate) | **Post** /v1/platform/compose/deployTemplate | Deploy a one-click template
[**PlatformComposeDisconnectGitProvider**](PlatformComposeAPI.md#PlatformComposeDisconnectGitProvider) | **Post** /v1/platform/compose/disconnectGitProvider | Disconnect git providers from compose
[**PlatformComposeGetConvertedCompose**](PlatformComposeAPI.md#PlatformComposeGetConvertedCompose) | **Get** /v1/platform/compose/getConvertedCompose | Get final compose YAML with domains injected
[**PlatformComposeGetDefaultCommand**](PlatformComposeAPI.md#PlatformComposeGetDefaultCommand) | **Get** /v1/platform/compose/getDefaultCommand | Get the docker compose command
[**PlatformComposeGetTags**](PlatformComposeAPI.md#PlatformComposeGetTags) | **Get** /v1/platform/compose/getTags | Get unique tags from all templates
[**PlatformComposeLoadServices**](PlatformComposeAPI.md#PlatformComposeLoadServices) | **Get** /v1/platform/compose/loadServices | List running services within a compose stack
[**PlatformComposeMove**](PlatformComposeAPI.md#PlatformComposeMove) | **Post** /v1/platform/compose/move | Move compose to a different environment
[**PlatformComposeOne**](PlatformComposeAPI.md#PlatformComposeOne) | **Get** /v1/platform/compose/one | Get compose service details
[**PlatformComposeRedeploy**](PlatformComposeAPI.md#PlatformComposeRedeploy) | **Post** /v1/platform/compose/redeploy | Redeploy (rebuild) a compose service
[**PlatformComposeRefreshToken**](PlatformComposeAPI.md#PlatformComposeRefreshToken) | **Post** /v1/platform/compose/refreshToken | Regenerate webhook token
[**PlatformComposeStart**](PlatformComposeAPI.md#PlatformComposeStart) | **Post** /v1/platform/compose/start | Start a compose service
[**PlatformComposeStop**](PlatformComposeAPI.md#PlatformComposeStop) | **Post** /v1/platform/compose/stop | Stop a compose service
[**PlatformComposeTemplates**](PlatformComposeAPI.md#PlatformComposeTemplates) | **Get** /v1/platform/compose/templates | List available one-click templates
[**PlatformComposeUpdate**](PlatformComposeAPI.md#PlatformComposeUpdate) | **Post** /v1/platform/compose/update | Update compose configuration



## PlatformComposeCancelDeployment

> PlatformTRPCResult PlatformComposeCancelDeployment(ctx).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()

Cancel in-progress compose deployment

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
	platformComposeCancelDeploymentRequest := *openapiclient.NewPlatformComposeCancelDeploymentRequest() // PlatformComposeCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeCancelDeployment(context.Background()).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeCancelDeployment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeCancelDeployment`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeCancelDeployment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeCancelDeploymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCancelDeploymentRequest** | [**PlatformComposeCancelDeploymentRequest**](PlatformComposeCancelDeploymentRequest.md) |  | 

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


## PlatformComposeCreate

> PlatformTRPCResult PlatformComposeCreate(ctx).PlatformComposeCreateRequest(platformComposeCreateRequest).Execute()

Create a Docker Compose service

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
	platformComposeCreateRequest := *openapiclient.NewPlatformComposeCreateRequest() // PlatformComposeCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeCreate(context.Background()).PlatformComposeCreateRequest(platformComposeCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCreateRequest** | [**PlatformComposeCreateRequest**](PlatformComposeCreateRequest.md) |  | 

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


## PlatformComposeDelete

> PlatformTRPCResult PlatformComposeDelete(ctx).PlatformComposeDeleteRequest(platformComposeDeleteRequest).Execute()

Delete a compose service

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
	platformComposeDeleteRequest := *openapiclient.NewPlatformComposeDeleteRequest() // PlatformComposeDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeDelete(context.Background()).PlatformComposeDeleteRequest(platformComposeDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeDelete`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeDeleteRequest** | [**PlatformComposeDeleteRequest**](PlatformComposeDeleteRequest.md) |  | 

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


## PlatformComposeDeploy

> PlatformTRPCResult PlatformComposeDeploy(ctx).PlatformComposeDeployRequest(platformComposeDeployRequest).Execute()

Deploy a compose service

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
	platformComposeDeployRequest := *openapiclient.NewPlatformComposeDeployRequest() // PlatformComposeDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeDeploy(context.Background()).PlatformComposeDeployRequest(platformComposeDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeDeploy`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeDeployRequest** | [**PlatformComposeDeployRequest**](PlatformComposeDeployRequest.md) |  | 

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


## PlatformComposeDeployTemplate

> PlatformTRPCResult PlatformComposeDeployTemplate(ctx).PlatformComposeDeployTemplateRequest(platformComposeDeployTemplateRequest).Execute()

Deploy a one-click template

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
	platformComposeDeployTemplateRequest := *openapiclient.NewPlatformComposeDeployTemplateRequest() // PlatformComposeDeployTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeDeployTemplate(context.Background()).PlatformComposeDeployTemplateRequest(platformComposeDeployTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeDeployTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeDeployTemplate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeDeployTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeDeployTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeDeployTemplateRequest** | [**PlatformComposeDeployTemplateRequest**](PlatformComposeDeployTemplateRequest.md) |  | 

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


## PlatformComposeDisconnectGitProvider

> PlatformTRPCResult PlatformComposeDisconnectGitProvider(ctx).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()

Disconnect git providers from compose

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
	platformComposeCancelDeploymentRequest := *openapiclient.NewPlatformComposeCancelDeploymentRequest() // PlatformComposeCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeDisconnectGitProvider(context.Background()).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeDisconnectGitProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeDisconnectGitProvider`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeDisconnectGitProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeDisconnectGitProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCancelDeploymentRequest** | [**PlatformComposeCancelDeploymentRequest**](PlatformComposeCancelDeploymentRequest.md) |  | 

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


## PlatformComposeGetConvertedCompose

> PlatformTRPCResult PlatformComposeGetConvertedCompose(ctx).Input(input).Execute()

Get final compose YAML with domains injected

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeGetConvertedCompose(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeGetConvertedCompose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeGetConvertedCompose`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeGetConvertedCompose`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeGetConvertedComposeRequest struct via the builder pattern


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


## PlatformComposeGetDefaultCommand

> PlatformTRPCResult PlatformComposeGetDefaultCommand(ctx).Input(input).Execute()

Get the docker compose command

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeGetDefaultCommand(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeGetDefaultCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeGetDefaultCommand`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeGetDefaultCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeGetDefaultCommandRequest struct via the builder pattern


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


## PlatformComposeGetTags

> PlatformTRPCResult PlatformComposeGetTags(ctx).Input(input).Execute()

Get unique tags from all templates

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeGetTags(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeGetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeGetTags`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeGetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeGetTagsRequest struct via the builder pattern


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


## PlatformComposeLoadServices

> PlatformTRPCResult PlatformComposeLoadServices(ctx).Input(input).Execute()

List running services within a compose stack

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeLoadServices(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeLoadServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeLoadServices`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeLoadServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeLoadServicesRequest struct via the builder pattern


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


## PlatformComposeMove

> PlatformTRPCResult PlatformComposeMove(ctx).PlatformComposeMoveRequest(platformComposeMoveRequest).Execute()

Move compose to a different environment

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
	platformComposeMoveRequest := *openapiclient.NewPlatformComposeMoveRequest() // PlatformComposeMoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeMove(context.Background()).PlatformComposeMoveRequest(platformComposeMoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeMove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeMove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeMoveRequest** | [**PlatformComposeMoveRequest**](PlatformComposeMoveRequest.md) |  | 

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


## PlatformComposeOne

> PlatformTRPCResult PlatformComposeOne(ctx).Input(input).Execute()

Get compose service details

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeOneRequest struct via the builder pattern


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


## PlatformComposeRedeploy

> PlatformTRPCResult PlatformComposeRedeploy(ctx).PlatformComposeDeployRequest(platformComposeDeployRequest).Execute()

Redeploy (rebuild) a compose service

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
	platformComposeDeployRequest := *openapiclient.NewPlatformComposeDeployRequest() // PlatformComposeDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeRedeploy(context.Background()).PlatformComposeDeployRequest(platformComposeDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeRedeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeRedeploy`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeRedeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeRedeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeDeployRequest** | [**PlatformComposeDeployRequest**](PlatformComposeDeployRequest.md) |  | 

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


## PlatformComposeRefreshToken

> PlatformTRPCResult PlatformComposeRefreshToken(ctx).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()

Regenerate webhook token

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
	platformComposeCancelDeploymentRequest := *openapiclient.NewPlatformComposeCancelDeploymentRequest() // PlatformComposeCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeRefreshToken(context.Background()).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeRefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeRefreshToken`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeRefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCancelDeploymentRequest** | [**PlatformComposeCancelDeploymentRequest**](PlatformComposeCancelDeploymentRequest.md) |  | 

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


## PlatformComposeStart

> PlatformTRPCResult PlatformComposeStart(ctx).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()

Start a compose service

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
	platformComposeCancelDeploymentRequest := *openapiclient.NewPlatformComposeCancelDeploymentRequest() // PlatformComposeCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeStart(context.Background()).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeStart`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCancelDeploymentRequest** | [**PlatformComposeCancelDeploymentRequest**](PlatformComposeCancelDeploymentRequest.md) |  | 

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


## PlatformComposeStop

> PlatformTRPCResult PlatformComposeStop(ctx).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()

Stop a compose service

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
	platformComposeCancelDeploymentRequest := *openapiclient.NewPlatformComposeCancelDeploymentRequest() // PlatformComposeCancelDeploymentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeStop(context.Background()).PlatformComposeCancelDeploymentRequest(platformComposeCancelDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeStop`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeStop`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeCancelDeploymentRequest** | [**PlatformComposeCancelDeploymentRequest**](PlatformComposeCancelDeploymentRequest.md) |  | 

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


## PlatformComposeTemplates

> PlatformTRPCResult PlatformComposeTemplates(ctx).Input(input).Execute()

List available one-click templates

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
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeTemplates(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeTemplates`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **input** | **string** | URL-encoded JSON input for tRPC queries | 

### Return type

[**PlatformTRPCResult**](PlatformTRPCResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformComposeUpdate

> PlatformTRPCResult PlatformComposeUpdate(ctx).PlatformComposeUpdateRequest(platformComposeUpdateRequest).Execute()

Update compose configuration

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
	platformComposeUpdateRequest := *openapiclient.NewPlatformComposeUpdateRequest() // PlatformComposeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformComposeAPI.PlatformComposeUpdate(context.Background()).PlatformComposeUpdateRequest(platformComposeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformComposeAPI.PlatformComposeUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformComposeUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformComposeAPI.PlatformComposeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformComposeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformComposeUpdateRequest** | [**PlatformComposeUpdateRequest**](PlatformComposeUpdateRequest.md) |  | 

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

