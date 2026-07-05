# \PlatformDomainAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformDomainByApplicationId**](PlatformDomainAPI.md#PlatformDomainByApplicationId) | **Get** /v1/platform/domain/byApplicationId | List domains for an application
[**PlatformDomainByComposeId**](PlatformDomainAPI.md#PlatformDomainByComposeId) | **Get** /v1/platform/domain/byComposeId | List domains for a compose service
[**PlatformDomainCreate**](PlatformDomainAPI.md#PlatformDomainCreate) | **Post** /v1/platform/domain/create | Create a domain mapping
[**PlatformDomainDelete**](PlatformDomainAPI.md#PlatformDomainDelete) | **Post** /v1/platform/domain/delete | Delete a domain
[**PlatformDomainGenerateDomain**](PlatformDomainAPI.md#PlatformDomainGenerateDomain) | **Post** /v1/platform/domain/generateDomain | Generate a traefik.me auto-domain
[**PlatformDomainOne**](PlatformDomainAPI.md#PlatformDomainOne) | **Get** /v1/platform/domain/one | Get a domain by ID
[**PlatformDomainUpdate**](PlatformDomainAPI.md#PlatformDomainUpdate) | **Post** /v1/platform/domain/update | Update a domain mapping
[**PlatformDomainValidateDomain**](PlatformDomainAPI.md#PlatformDomainValidateDomain) | **Post** /v1/platform/domain/validateDomain | Validate DNS for a domain



## PlatformDomainByApplicationId

> PlatformTRPCResult PlatformDomainByApplicationId(ctx).Input(input).Execute()

List domains for an application

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
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainByApplicationId(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainByApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainByApplicationId`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainByApplicationId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainByApplicationIdRequest struct via the builder pattern


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


## PlatformDomainByComposeId

> PlatformTRPCResult PlatformDomainByComposeId(ctx).Input(input).Execute()

List domains for a compose service

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
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainByComposeId(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainByComposeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainByComposeId`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainByComposeId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainByComposeIdRequest struct via the builder pattern


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


## PlatformDomainCreate

> PlatformTRPCResult PlatformDomainCreate(ctx).PlatformDomainCreateRequest(platformDomainCreateRequest).Execute()

Create a domain mapping

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
	platformDomainCreateRequest := *openapiclient.NewPlatformDomainCreateRequest() // PlatformDomainCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainCreate(context.Background()).PlatformDomainCreateRequest(platformDomainCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDomainCreateRequest** | [**PlatformDomainCreateRequest**](PlatformDomainCreateRequest.md) |  | 

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


## PlatformDomainDelete

> PlatformTRPCResult PlatformDomainDelete(ctx).PlatformDomainDeleteRequest(platformDomainDeleteRequest).Execute()

Delete a domain

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
	platformDomainDeleteRequest := *openapiclient.NewPlatformDomainDeleteRequest() // PlatformDomainDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainDelete(context.Background()).PlatformDomainDeleteRequest(platformDomainDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainDelete`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDomainDeleteRequest** | [**PlatformDomainDeleteRequest**](PlatformDomainDeleteRequest.md) |  | 

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


## PlatformDomainGenerateDomain

> PlatformTRPCResult PlatformDomainGenerateDomain(ctx).PlatformDomainGenerateDomainRequest(platformDomainGenerateDomainRequest).Execute()

Generate a traefik.me auto-domain

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
	platformDomainGenerateDomainRequest := *openapiclient.NewPlatformDomainGenerateDomainRequest() // PlatformDomainGenerateDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainGenerateDomain(context.Background()).PlatformDomainGenerateDomainRequest(platformDomainGenerateDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainGenerateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainGenerateDomain`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainGenerateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainGenerateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDomainGenerateDomainRequest** | [**PlatformDomainGenerateDomainRequest**](PlatformDomainGenerateDomainRequest.md) |  | 

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


## PlatformDomainOne

> PlatformTRPCResult PlatformDomainOne(ctx).Input(input).Execute()

Get a domain by ID

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
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainOneRequest struct via the builder pattern


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


## PlatformDomainUpdate

> PlatformTRPCResult PlatformDomainUpdate(ctx).PlatformDomainUpdateRequest(platformDomainUpdateRequest).Execute()

Update a domain mapping

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
	platformDomainUpdateRequest := *openapiclient.NewPlatformDomainUpdateRequest() // PlatformDomainUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainUpdate(context.Background()).PlatformDomainUpdateRequest(platformDomainUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDomainUpdateRequest** | [**PlatformDomainUpdateRequest**](PlatformDomainUpdateRequest.md) |  | 

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


## PlatformDomainValidateDomain

> PlatformTRPCResult PlatformDomainValidateDomain(ctx).PlatformDomainValidateDomainRequest(platformDomainValidateDomainRequest).Execute()

Validate DNS for a domain

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
	platformDomainValidateDomainRequest := *openapiclient.NewPlatformDomainValidateDomainRequest() // PlatformDomainValidateDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformDomainAPI.PlatformDomainValidateDomain(context.Background()).PlatformDomainValidateDomainRequest(platformDomainValidateDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformDomainAPI.PlatformDomainValidateDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformDomainValidateDomain`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformDomainAPI.PlatformDomainValidateDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformDomainValidateDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformDomainValidateDomainRequest** | [**PlatformDomainValidateDomainRequest**](PlatformDomainValidateDomainRequest.md) |  | 

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

