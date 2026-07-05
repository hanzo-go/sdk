# \PlatformServerAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformServerAll**](PlatformServerAPI.md#PlatformServerAll) | **Get** /v1/platform/server/all | List all servers with service counts
[**PlatformServerCreate**](PlatformServerAPI.md#PlatformServerCreate) | **Post** /v1/platform/server/create | Register a new remote server
[**PlatformServerGetServerMetrics**](PlatformServerAPI.md#PlatformServerGetServerMetrics) | **Get** /v1/platform/server/getServerMetrics | Fetch server metrics (CPU, memory, disk, network)
[**PlatformServerOne**](PlatformServerAPI.md#PlatformServerOne) | **Get** /v1/platform/server/one | Get server details
[**PlatformServerPublicIp**](PlatformServerAPI.md#PlatformServerPublicIp) | **Get** /v1/platform/server/publicIp | Get platform host public IP
[**PlatformServerRemove**](PlatformServerAPI.md#PlatformServerRemove) | **Post** /v1/platform/server/remove | Remove a server (must have no active services)
[**PlatformServerSecurity**](PlatformServerAPI.md#PlatformServerSecurity) | **Get** /v1/platform/server/security | Run security audit (ufw, ssh, fail2ban)
[**PlatformServerSetup**](PlatformServerAPI.md#PlatformServerSetup) | **Post** /v1/platform/server/setup | Run initial setup (install Docker, Traefik, etc.)
[**PlatformServerSetupMonitoring**](PlatformServerAPI.md#PlatformServerSetupMonitoring) | **Post** /v1/platform/server/setupMonitoring | Configure and deploy monitoring agent
[**PlatformServerUpdate**](PlatformServerAPI.md#PlatformServerUpdate) | **Post** /v1/platform/server/update | Update server configuration
[**PlatformServerValidate**](PlatformServerAPI.md#PlatformServerValidate) | **Get** /v1/platform/server/validate | Validate server capabilities



## PlatformServerAll

> PlatformTRPCResult PlatformServerAll(ctx).Execute()

List all servers with service counts

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerAllRequest struct via the builder pattern


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


## PlatformServerCreate

> PlatformTRPCResult PlatformServerCreate(ctx).PlatformServerCreateRequest(platformServerCreateRequest).Execute()

Register a new remote server

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
	platformServerCreateRequest := *openapiclient.NewPlatformServerCreateRequest() // PlatformServerCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerCreate(context.Background()).PlatformServerCreateRequest(platformServerCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerCreate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformServerCreateRequest** | [**PlatformServerCreateRequest**](PlatformServerCreateRequest.md) |  | 

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


## PlatformServerGetServerMetrics

> PlatformTRPCResult PlatformServerGetServerMetrics(ctx).Input(input).Execute()

Fetch server metrics (CPU, memory, disk, network)

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerGetServerMetrics(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerGetServerMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerGetServerMetrics`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerGetServerMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerGetServerMetricsRequest struct via the builder pattern


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


## PlatformServerOne

> PlatformTRPCResult PlatformServerOne(ctx).Input(input).Execute()

Get server details

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerOne(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerOne``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerOne`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerOne`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerOneRequest struct via the builder pattern


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


## PlatformServerPublicIp

> PlatformTRPCResult PlatformServerPublicIp(ctx).Execute()

Get platform host public IP

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerPublicIp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerPublicIp`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerPublicIp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerPublicIpRequest struct via the builder pattern


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


## PlatformServerRemove

> PlatformTRPCResult PlatformServerRemove(ctx).PlatformServerRemoveRequest(platformServerRemoveRequest).Execute()

Remove a server (must have no active services)

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
	platformServerRemoveRequest := *openapiclient.NewPlatformServerRemoveRequest() // PlatformServerRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerRemove(context.Background()).PlatformServerRemoveRequest(platformServerRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerRemove`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformServerRemoveRequest** | [**PlatformServerRemoveRequest**](PlatformServerRemoveRequest.md) |  | 

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


## PlatformServerSecurity

> PlatformTRPCResult PlatformServerSecurity(ctx).Input(input).Execute()

Run security audit (ufw, ssh, fail2ban)

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerSecurity(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerSecurity`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerSecurity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerSecurityRequest struct via the builder pattern


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


## PlatformServerSetup

> PlatformTRPCResult PlatformServerSetup(ctx).PlatformServerRemoveRequest(platformServerRemoveRequest).Execute()

Run initial setup (install Docker, Traefik, etc.)

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
	platformServerRemoveRequest := *openapiclient.NewPlatformServerRemoveRequest() // PlatformServerRemoveRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerSetup(context.Background()).PlatformServerRemoveRequest(platformServerRemoveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerSetup`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerSetup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformServerRemoveRequest** | [**PlatformServerRemoveRequest**](PlatformServerRemoveRequest.md) |  | 

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


## PlatformServerSetupMonitoring

> PlatformTRPCResult PlatformServerSetupMonitoring(ctx).PlatformServerSetupMonitoringRequest(platformServerSetupMonitoringRequest).Execute()

Configure and deploy monitoring agent

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
	platformServerSetupMonitoringRequest := *openapiclient.NewPlatformServerSetupMonitoringRequest() // PlatformServerSetupMonitoringRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerSetupMonitoring(context.Background()).PlatformServerSetupMonitoringRequest(platformServerSetupMonitoringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerSetupMonitoring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerSetupMonitoring`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerSetupMonitoring`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerSetupMonitoringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformServerSetupMonitoringRequest** | [**PlatformServerSetupMonitoringRequest**](PlatformServerSetupMonitoringRequest.md) |  | 

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


## PlatformServerUpdate

> PlatformTRPCResult PlatformServerUpdate(ctx).PlatformServerUpdateRequest(platformServerUpdateRequest).Execute()

Update server configuration

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
	platformServerUpdateRequest := *openapiclient.NewPlatformServerUpdateRequest() // PlatformServerUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerUpdate(context.Background()).PlatformServerUpdateRequest(platformServerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerUpdate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformServerUpdateRequest** | [**PlatformServerUpdateRequest**](PlatformServerUpdateRequest.md) |  | 

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


## PlatformServerValidate

> PlatformTRPCResult PlatformServerValidate(ctx).Input(input).Execute()

Validate server capabilities

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
	resp, r, err := apiClient.PlatformServerAPI.PlatformServerValidate(context.Background()).Input(input).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformServerAPI.PlatformServerValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformServerValidate`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformServerAPI.PlatformServerValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformServerValidateRequest struct via the builder pattern


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

