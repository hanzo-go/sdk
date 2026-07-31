# \HealthAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthzAuthzHealth**](HealthAPI.md#AuthzAuthzHealth) | **Get** /v1/authz/health | Liveness probe
[**AuthzAuthzReadyz**](HealthAPI.md#AuthzAuthzReadyz) | **Get** /v1/authz/readyz | Readiness probe
[**GatewayLivelinessCheck**](HealthAPI.md#GatewayLivelinessCheck) | **Get** /v1/gateway/health/liveliness | Liveliness check
[**GatewayReadinessCheck**](HealthAPI.md#GatewayReadinessCheck) | **Get** /v1/gateway/health/readiness | Readiness check
[**KmsGetV1KmsHealth**](HealthAPI.md#KmsGetV1KmsHealth) | **Get** /v1/kms/health | Liveness
[**KmsGetV1KmsHealthz**](HealthAPI.md#KmsGetV1KmsHealthz) | **Get** /v1/kms/healthz | Liveness
[**NotifyNotifyHealth**](HealthAPI.md#NotifyNotifyHealth) | **Get** /v1/notify/health | Liveness probe
[**PricingTriggerSync**](HealthAPI.md#PricingTriggerSync) | **Post** /v1/pricing/sync | Trigger manual sync
[**ReferralsReferralsHealth**](HealthAPI.md#ReferralsReferralsHealth) | **Get** /v1/referrals/health | Liveness probe
[**SearchGetHealth**](HealthAPI.md#SearchGetHealth) | **Get** /health | Health check



## AuthzAuthzHealth

> AuthzAuthzHealth200Response AuthzAuthzHealth(ctx).Execute()

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
	resp, r, err := apiClient.HealthAPI.AuthzAuthzHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.AuthzAuthzHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzHealth`: AuthzAuthzHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.AuthzAuthzHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzHealthRequest struct via the builder pattern


### Return type

[**AuthzAuthzHealth200Response**](AuthzAuthzHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthzAuthzReadyz

> AuthzAuthzReadyz200Response AuthzAuthzReadyz(ctx).Execute()

Readiness probe



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
	resp, r, err := apiClient.HealthAPI.AuthzAuthzReadyz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.AuthzAuthzReadyz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzReadyz`: AuthzAuthzReadyz200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.AuthzAuthzReadyz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzReadyzRequest struct via the builder pattern


### Return type

[**AuthzAuthzReadyz200Response**](AuthzAuthzReadyz200Response.md)

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


## KmsGetV1KmsHealth

> KmsGetV1KmsHealth200Response KmsGetV1KmsHealth(ctx).Execute()

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
	resp, r, err := apiClient.HealthAPI.KmsGetV1KmsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.KmsGetV1KmsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetV1KmsHealth`: KmsGetV1KmsHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.KmsGetV1KmsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsHealthRequest struct via the builder pattern


### Return type

[**KmsGetV1KmsHealth200Response**](KmsGetV1KmsHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsHealthz

> KmsGetV1KmsHealth200Response KmsGetV1KmsHealthz(ctx).Execute()

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
	// response from `KmsGetV1KmsHealthz`: KmsGetV1KmsHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.KmsGetV1KmsHealthz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsHealthzRequest struct via the builder pattern


### Return type

[**KmsGetV1KmsHealth200Response**](KmsGetV1KmsHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyNotifyHealth

> NotifyHealthResponse NotifyNotifyHealth(ctx).Execute()

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
	resp, r, err := apiClient.HealthAPI.NotifyNotifyHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.NotifyNotifyHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotifyNotifyHealth`: NotifyHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.NotifyNotifyHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotifyNotifyHealthRequest struct via the builder pattern


### Return type

[**NotifyHealthResponse**](NotifyHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingTriggerSync

> PricingTriggerSync200Response PricingTriggerSync(ctx).Execute()

Trigger manual sync



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
	resp, r, err := apiClient.HealthAPI.PricingTriggerSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.PricingTriggerSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingTriggerSync`: PricingTriggerSync200Response
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.PricingTriggerSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingTriggerSyncRequest struct via the builder pattern


### Return type

[**PricingTriggerSync200Response**](PricingTriggerSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

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


## SearchGetHealth

> SearchHealthResponse SearchGetHealth(ctx).Execute()

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
	resp, r, err := apiClient.HealthAPI.SearchGetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.SearchGetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetHealth`: SearchHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `HealthAPI.SearchGetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetHealthRequest struct via the builder pattern


### Return type

[**SearchHealthResponse**](SearchHealthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

