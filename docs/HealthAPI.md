# \HealthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayHealthCheck**](HealthAPI.md#GatewayHealthCheck) | **Get** /healthz | Health check
[**GatewayLivelinessCheck**](HealthAPI.md#GatewayLivelinessCheck) | **Get** /v1/gateway/health/liveliness | Liveliness check
[**GatewayReadinessCheck**](HealthAPI.md#GatewayReadinessCheck) | **Get** /v1/gateway/health/readiness | Readiness check
[**KmsGetV1KmsHealthz**](HealthAPI.md#KmsGetV1KmsHealthz) | **Get** /v1/kms/healthz | Liveness
[**ReferralsReferralsHealth**](HealthAPI.md#ReferralsReferralsHealth) | **Get** /v1/referrals/health | Liveness probe



## GatewayHealthCheck

> GatewayHealthCheck200Response GatewayHealthCheck(ctx).Execute()

Health check

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
	resp, r, err := apiClient.HealthAPI.GatewayHealthCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GatewayHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayHealthCheck`: GatewayHealthCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GatewayHealthCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayHealthCheckRequest struct via the builder pattern


### Return type

[**GatewayHealthCheck200Response**](GatewayHealthCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayLivelinessCheck

> string GatewayLivelinessCheck(ctx).Execute()

Liveliness check

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
	resp, r, err := apiClient.HealthAPI.GatewayLivelinessCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GatewayLivelinessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayLivelinessCheck`: string
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GatewayLivelinessCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayLivelinessCheckRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayReadinessCheck

> GatewayReadinessCheck200Response GatewayReadinessCheck(ctx).Execute()

Readiness check

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
	resp, r, err := apiClient.HealthAPI.GatewayReadinessCheck(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.GatewayReadinessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayReadinessCheck`: GatewayReadinessCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.GatewayReadinessCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayReadinessCheckRequest struct via the builder pattern


### Return type

[**GatewayReadinessCheck200Response**](GatewayReadinessCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsHealthz

> KmsGetV1KmsHealthz200Response KmsGetV1KmsHealthz(ctx).Execute()

Liveness



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
	resp, r, err := apiClient.HealthAPI.KmsGetV1KmsHealthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.KmsGetV1KmsHealthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetV1KmsHealthz`: KmsGetV1KmsHealthz200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.KmsGetV1KmsHealthz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsHealthzRequest struct via the builder pattern


### Return type

[**KmsGetV1KmsHealthz200Response**](KmsGetV1KmsHealthz200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferralsReferralsHealth

> ReferralsReferralsHealth(ctx).Execute()

Liveness probe



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
	r, err := apiClient.HealthAPI.ReferralsReferralsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.ReferralsReferralsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferralsReferralsHealthRequest struct via the builder pattern


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

