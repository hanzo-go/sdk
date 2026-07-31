# \CloudAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PricingGetCloud**](CloudAPI.md#PricingGetCloud) | **Get** /v1/pricing/cloud | Cloud VM plans, regions, and storage
[**PricingGetStoragePricing**](CloudAPI.md#PricingGetStoragePricing) | **Get** /v1/pricing/cloud/storage | Block storage pricing
[**PricingListCloudPlans**](CloudAPI.md#PricingListCloudPlans) | **Get** /v1/pricing/cloud/plans | Cloud VM plans
[**PricingListCloudRegions**](CloudAPI.md#PricingListCloudRegions) | **Get** /v1/pricing/cloud/regions | Available cloud regions
[**WorldWorldCloudAnalytics**](CloudAPI.md#WorldWorldCloudAnalytics) | **Get** /v1/world/cloud/analytics | ADMIN. Web analytics aggregate — top pages/referrers/countries, live visitors (analytics.hanzo.ai). Requires an admin-org IAM bearer.
[**WorldWorldCloudByoGpu**](CloudAPI.md#WorldWorldCloudByoGpu) | **Get** /v1/world/cloud/byo-gpu | Public BYO-GPU map data — connected GPU workers by region (real counts when a service token is wired server-side, else demo-flagged). No auth.
[**WorldWorldCloudChainNodes**](CloudAPI.md#WorldWorldCloudChainNodes) | **Get** /v1/world/cloud/chain-nodes | Public blockchain-network map data — per-network block height, peer count, live flag, and modeled node positions (positionsModeled:true; counts are real, geo is illustrative). No auth.
[**WorldWorldCloudFleet**](CloudAPI.md#WorldWorldCloudFleet) | **Get** /v1/world/cloud/fleet | ADMIN. Machines + GPUs grouped by provider/region (visor). Requires an admin-org IAM bearer; 401 without a token, 403 for non-admin.
[**WorldWorldCloudLlm**](CloudAPI.md#WorldWorldCloudLlm) | **Get** /v1/world/cloud/llm | ADMIN. Platform LLM observability — per-model/per-org usage, tokens, cost, errors, trace latency (cloud /v1/admin/o11y). Requires an admin-org IAM bearer.
[**WorldWorldCloudModels**](CloudAPI.md#WorldWorldCloudModels) | **Get** /v1/world/cloud/models | Public served-model catalog + scale (from the gateway /v1/models). No auth.
[**WorldWorldCloudPulse**](CloudAPI.md#WorldWorldCloudPulse) | **Get** /v1/world/cloud-pulse | Public platform aggregate (SaaS variant). Anonymized counts; demo-flagged unless a service token is wired server-side.
[**WorldWorldCloudServices**](CloudAPI.md#WorldWorldCloudServices) | **Get** /v1/world/cloud/services | ADMIN. Per-subsystem health + RED metrics (o11y). Requires an admin-org IAM bearer.
[**WorldWorldCloudTraffic**](CloudAPI.md#WorldWorldCloudTraffic) | **Get** /v1/world/cloud/traffic | Public request-traffic arcs — country-level origin → nearest region, weight-normalized (real analytics when a service token is wired server-side, else demo-flagged). No auth.



## PricingGetCloud

> PricingCloudResponse PricingGetCloud(ctx).Execute()

Cloud VM plans, regions, and storage



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
	resp, r, err := apiClient.CloudAPI.PricingGetCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PricingGetCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetCloud`: PricingCloudResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PricingGetCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetCloudRequest struct via the builder pattern


### Return type

[**PricingCloudResponse**](PricingCloudResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingGetStoragePricing

> PricingBlockStoragePricing PricingGetStoragePricing(ctx).Execute()

Block storage pricing



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
	resp, r, err := apiClient.CloudAPI.PricingGetStoragePricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PricingGetStoragePricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingGetStoragePricing`: PricingBlockStoragePricing
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PricingGetStoragePricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingGetStoragePricingRequest struct via the builder pattern


### Return type

[**PricingBlockStoragePricing**](PricingBlockStoragePricing.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListCloudPlans

> PricingCloudPlansResponse PricingListCloudPlans(ctx).Execute()

Cloud VM plans



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
	resp, r, err := apiClient.CloudAPI.PricingListCloudPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PricingListCloudPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListCloudPlans`: PricingCloudPlansResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PricingListCloudPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListCloudPlansRequest struct via the builder pattern


### Return type

[**PricingCloudPlansResponse**](PricingCloudPlansResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PricingListCloudRegions

> PricingCloudRegionsResponse PricingListCloudRegions(ctx).Execute()

Available cloud regions



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
	resp, r, err := apiClient.CloudAPI.PricingListCloudRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PricingListCloudRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PricingListCloudRegions`: PricingCloudRegionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PricingListCloudRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPricingListCloudRegionsRequest struct via the builder pattern


### Return type

[**PricingCloudRegionsResponse**](PricingCloudRegionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudAnalytics

> map[string]interface{} WorldWorldCloudAnalytics(ctx).Execute()

ADMIN. Web analytics aggregate — top pages/referrers/countries, live visitors (analytics.hanzo.ai). Requires an admin-org IAM bearer.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudAnalytics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudAnalytics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudAnalytics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudAnalyticsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudByoGpu

> map[string]interface{} WorldWorldCloudByoGpu(ctx).Execute()

Public BYO-GPU map data — connected GPU workers by region (real counts when a service token is wired server-side, else demo-flagged). No auth.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudByoGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudByoGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudByoGpu`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudByoGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudByoGpuRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudChainNodes

> map[string]interface{} WorldWorldCloudChainNodes(ctx).Execute()

Public blockchain-network map data — per-network block height, peer count, live flag, and modeled node positions (positionsModeled:true; counts are real, geo is illustrative). No auth.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudChainNodes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudChainNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudChainNodes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudChainNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudChainNodesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudFleet

> map[string]interface{} WorldWorldCloudFleet(ctx).Execute()

ADMIN. Machines + GPUs grouped by provider/region (visor). Requires an admin-org IAM bearer; 401 without a token, 403 for non-admin.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudFleet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudFleet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudFleet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudFleet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudFleetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudLlm

> map[string]interface{} WorldWorldCloudLlm(ctx).Range_(range_).Execute()

ADMIN. Platform LLM observability — per-model/per-org usage, tokens, cost, errors, trace latency (cloud /v1/admin/o11y). Requires an admin-org IAM bearer.

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
	range_ := "range__example" // string |  (optional) (default to "24h")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudLlm(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudLlm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudLlm`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudLlm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudLlmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  | [default to &quot;24h&quot;]

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudModels

> map[string]interface{} WorldWorldCloudModels(ctx).Execute()

Public served-model catalog + scale (from the gateway /v1/models). No auth.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudModels`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudModelsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudPulse

> map[string]interface{} WorldWorldCloudPulse(ctx).Execute()

Public platform aggregate (SaaS variant). Anonymized counts; demo-flagged unless a service token is wired server-side.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudPulse(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudPulse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudPulse`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudPulse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudPulseRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudServices

> map[string]interface{} WorldWorldCloudServices(ctx).Execute()

ADMIN. Per-subsystem health + RED metrics (o11y). Requires an admin-org IAM bearer.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudServices`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudServicesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCloudTraffic

> map[string]interface{} WorldWorldCloudTraffic(ctx).Execute()

Public request-traffic arcs — country-level origin → nearest region, weight-normalized (real analytics when a service token is wired server-side, else demo-flagged). No auth.

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
	resp, r, err := apiClient.CloudAPI.WorldWorldCloudTraffic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.WorldWorldCloudTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudTraffic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.WorldWorldCloudTraffic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudTrafficRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

