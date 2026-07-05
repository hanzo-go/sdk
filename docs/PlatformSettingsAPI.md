# \PlatformSettingsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlatformHealth**](PlatformSettingsAPI.md#PlatformHealth) | **Get** /v1/platform/health | Platform health (REST)
[**PlatformHealthcheck**](PlatformSettingsAPI.md#PlatformHealthcheck) | **Get** /v1/platform/healthcheck | Platform healthcheck (REST)
[**PlatformSettingsCleanAll**](PlatformSettingsAPI.md#PlatformSettingsCleanAll) | **Post** /v1/platform/settings/cleanAll | Full Docker cleanup (admin)
[**PlatformSettingsGetHanzoVersion**](PlatformSettingsAPI.md#PlatformSettingsGetHanzoVersion) | **Get** /v1/platform/settings/getHanzoVersion | Get platform version
[**PlatformSettingsHealth**](PlatformSettingsAPI.md#PlatformSettingsHealth) | **Get** /v1/platform/settings/health | Platform health check (tRPC)
[**PlatformSettingsIsCloud**](PlatformSettingsAPI.md#PlatformSettingsIsCloud) | **Get** /v1/platform/settings/isCloud | Check if running in cloud mode
[**PlatformSettingsReloadTraefik**](PlatformSettingsAPI.md#PlatformSettingsReloadTraefik) | **Post** /v1/platform/settings/reloadTraefik | Reload Traefik configuration (admin)



## PlatformHealth

> PlatformHealth(ctx).Execute()

Platform health (REST)

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
	r, err := apiClient.PlatformSettingsAPI.PlatformHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformHealthRequest struct via the builder pattern


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


## PlatformHealthcheck

> PlatformHealthcheck(ctx).Execute()

Platform healthcheck (REST)

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
	r, err := apiClient.PlatformSettingsAPI.PlatformHealthcheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformHealthcheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformHealthcheckRequest struct via the builder pattern


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


## PlatformSettingsCleanAll

> PlatformTRPCResult PlatformSettingsCleanAll(ctx).Execute()

Full Docker cleanup (admin)

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
	resp, r, err := apiClient.PlatformSettingsAPI.PlatformSettingsCleanAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformSettingsCleanAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSettingsCleanAll`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSettingsAPI.PlatformSettingsCleanAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSettingsCleanAllRequest struct via the builder pattern


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


## PlatformSettingsGetHanzoVersion

> PlatformTRPCResult PlatformSettingsGetHanzoVersion(ctx).Execute()

Get platform version

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
	resp, r, err := apiClient.PlatformSettingsAPI.PlatformSettingsGetHanzoVersion(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformSettingsGetHanzoVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSettingsGetHanzoVersion`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSettingsAPI.PlatformSettingsGetHanzoVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSettingsGetHanzoVersionRequest struct via the builder pattern


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


## PlatformSettingsHealth

> PlatformTRPCResult PlatformSettingsHealth(ctx).Execute()

Platform health check (tRPC)

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
	resp, r, err := apiClient.PlatformSettingsAPI.PlatformSettingsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformSettingsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSettingsHealth`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSettingsAPI.PlatformSettingsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSettingsHealthRequest struct via the builder pattern


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


## PlatformSettingsIsCloud

> PlatformTRPCResult PlatformSettingsIsCloud(ctx).Execute()

Check if running in cloud mode

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
	resp, r, err := apiClient.PlatformSettingsAPI.PlatformSettingsIsCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformSettingsIsCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSettingsIsCloud`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSettingsAPI.PlatformSettingsIsCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSettingsIsCloudRequest struct via the builder pattern


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


## PlatformSettingsReloadTraefik

> PlatformTRPCResult PlatformSettingsReloadTraefik(ctx).PlatformSettingsReloadTraefikRequest(platformSettingsReloadTraefikRequest).Execute()

Reload Traefik configuration (admin)

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
	platformSettingsReloadTraefikRequest := *openapiclient.NewPlatformSettingsReloadTraefikRequest() // PlatformSettingsReloadTraefikRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlatformSettingsAPI.PlatformSettingsReloadTraefik(context.Background()).PlatformSettingsReloadTraefikRequest(platformSettingsReloadTraefikRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlatformSettingsAPI.PlatformSettingsReloadTraefik``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlatformSettingsReloadTraefik`: PlatformTRPCResult
	fmt.Fprintf(os.Stdout, "Response from `PlatformSettingsAPI.PlatformSettingsReloadTraefik`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlatformSettingsReloadTraefikRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platformSettingsReloadTraefikRequest** | [**PlatformSettingsReloadTraefikRequest**](PlatformSettingsReloadTraefikRequest.md) |  | 

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

