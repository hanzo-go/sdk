# \PlanAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPlan**](PlanAPI.md#GetPlan) | **Get** /v1/plan | Returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller&#39;s catalog.
[**GetPlanBlockchain**](PlanAPI.md#GetPlanBlockchain) | **Get** /v1/plan/blockchain | Returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.
[**GetPlanDns**](PlanAPI.md#GetPlanDns) | **Get** /v1/plan/dns | ListDNSPlans returns the DNS plan catalog: the tiers priced on zones, records per zone and queries per day.
[**GetPlanEntitlementsById**](PlanAPI.md#GetPlanEntitlementsById) | **Get** /v1/plan/entitlements/{id} | Returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.
[**GetPlanGpu**](PlanAPI.md#GetPlanGpu) | **Get** /v1/plan/gpu | ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPUs, host memory and hourly price.
[**GetPlanHealth**](PlanAPI.md#GetPlanHealth) | **Get** /v1/plan/health | Health reports that the plans subsystem is mounted and serving.
[**GetPlanPolicy**](PlanAPI.md#GetPlanPolicy) | **Get** /v1/plan/policy | Returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.
[**GetPlanRegions**](PlanAPI.md#GetPlanRegions) | **Get** /v1/plan/regions | Returns the regions cloud capacity is offered in, each with its display name and physical location.
[**GetPlanResolveById**](PlanAPI.md#GetPlanResolveById) | **Get** /v1/plan/resolve/{id} | Resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.
[**GetPlanSchema**](PlanAPI.md#GetPlanSchema) | **Get** /v1/plan/schema | Returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.
[**GetPlanStorage**](PlanAPI.md#GetPlanStorage) | **Get** /v1/plan/storage | Returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.
[**GetPlanSubscriptions**](PlanAPI.md#GetPlanSubscriptions) | **Get** /v1/plan/subscriptions | Returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.
[**GetPlanTools**](PlanAPI.md#GetPlanTools) | **Get** /v1/plan/tools | Returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.
[**GetPlanVocab**](PlanAPI.md#GetPlanVocab) | **Get** /v1/plan/vocab | Returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.



## GetPlan

> PlanList GetPlan(ctx).Execute()

Returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller's catalog.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlan(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


### Return type

[**PlanList**](PlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanBlockchain

> PlanList GetPlanBlockchain(ctx).Execute()

Returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanBlockchain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanBlockchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanBlockchain`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanBlockchain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanBlockchainRequest struct via the builder pattern


### Return type

[**PlanList**](PlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanDns

> PlanList GetPlanDns(ctx).Execute()

ListDNSPlans returns the DNS plan catalog: the tiers priced on zones, records per zone and queries per day.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanDns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanDns`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanDns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanDnsRequest struct via the builder pattern


### Return type

[**PlanList**](PlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanEntitlementsById

> PlanEntitlements GetPlanEntitlementsById(ctx, id).Execute()

Returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "team" // string | ID is the plan's catalog id or slug — \"dev\", \"max\", \"team\", \"rpc-growth\". Both are matched, so a slug resolves the plan it names. A withdrawn id still resolves for a renewal, which is why this takes an id rather than a ladder position.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanEntitlementsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanEntitlementsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanEntitlementsById`: PlanEntitlements
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanEntitlementsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;dev\&quot;, \&quot;max\&quot;, \&quot;team\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. A withdrawn id still resolves for a renewal, which is why this takes an id rather than a ladder position. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanEntitlementsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlanEntitlements**](PlanEntitlements.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanGpu

> PlanTierList GetPlanGpu(ctx).Execute()

ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPUs, host memory and hourly price.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanGpu`: PlanTierList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanGpuRequest struct via the builder pattern


### Return type

[**PlanTierList**](PlanTierList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanHealth

> PlanHealth GetPlanHealth(ctx).Execute()

Health reports that the plans subsystem is mounted and serving.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanHealth`: PlanHealth
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanHealthRequest struct via the builder pattern


### Return type

[**PlanHealth**](PlanHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanPolicy

> map[string]interface{} GetPlanPolicy(ctx).Execute()

Returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanPolicyRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanRegions

> PlanRegionList GetPlanRegions(ctx).Execute()

Returns the regions cloud capacity is offered in, each with its display name and physical location.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanRegions`: PlanRegionList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRegionsRequest struct via the builder pattern


### Return type

[**PlanRegionList**](PlanRegionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanResolveById

> PlanResolution GetPlanResolveById(ctx, id).Execute()

Resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "dev" // string | ID is the plan's catalog id or slug — \"dev\", \"max\", \"team\", \"rpc-growth\". Both are matched, so a slug resolves the plan it names. A withdrawn id still resolves for a renewal, which is why this takes an id rather than a ladder position.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanResolveById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanResolveById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanResolveById`: PlanResolution
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanResolveById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;dev\&quot;, \&quot;max\&quot;, \&quot;team\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. A withdrawn id still resolves for a renewal, which is why this takes an id rather than a ladder position. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanResolveByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlanResolution**](PlanResolution.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanSchema

> PlanSchemas GetPlanSchema(ctx).Execute()

Returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSchema`: PlanSchemas
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanSchemaRequest struct via the builder pattern


### Return type

[**PlanSchemas**](PlanSchemas.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanStorage

> map[string]interface{} GetPlanStorage(ctx).Execute()

Returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanStorage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanStorageRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanSubscriptions

> PlanList GetPlanSubscriptions(ctx).Execute()

Returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanSubscriptions`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanSubscriptionsRequest struct via the builder pattern


### Return type

[**PlanList**](PlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanTools

> PlanToolList GetPlanTools(ctx).Execute()

Returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanTools`: PlanToolList
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanToolsRequest struct via the builder pattern


### Return type

[**PlanToolList**](PlanToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanVocab

> PlanVocab GetPlanVocab(ctx).Execute()

Returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlanAPI.GetPlanVocab(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlanAPI.GetPlanVocab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanVocab`: PlanVocab
	fmt.Fprintf(os.Stdout, "Response from `PlanAPI.GetPlanVocab`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanVocabRequest struct via the builder pattern


### Return type

[**PlanVocab**](PlanVocab.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

