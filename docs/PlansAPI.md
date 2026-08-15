# \PlansAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlansEntriesBySlug**](PlansAPI.md#DeletePlansEntriesBySlug) | **Delete** /v1/plans/entries/{slug} | Remove a plan from the authority
[**GetPlans**](PlansAPI.md#GetPlans) | **Get** /v1/plans | Returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller&#39;s catalog.
[**GetPlansBlockchain**](PlansAPI.md#GetPlansBlockchain) | **Get** /v1/plans/blockchain | Returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.
[**GetPlansCloud**](PlansAPI.md#GetPlansCloud) | **Get** /v1/plans/cloud | Returns the cloud plan catalog.
[**GetPlansDns**](PlansAPI.md#GetPlansDns) | **Get** /v1/plans/dns | ListDNSPlans returns the DNS plan catalog: the tiers priced on zones, records per zone and queries per day.
[**GetPlansEntitlementsById**](PlansAPI.md#GetPlansEntitlementsById) | **Get** /v1/plans/entitlements/{id} | Returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.
[**GetPlansEntries**](PlansAPI.md#GetPlansEntries) | **Get** /v1/plans/entries | The raw plan authority rows
[**GetPlansGpu**](PlansAPI.md#GetPlansGpu) | **Get** /v1/plans/gpu | ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPUs, host memory and hourly price.
[**GetPlansHealth**](PlansAPI.md#GetPlansHealth) | **Get** /v1/plans/health | Health reports that the plans subsystem is mounted and serving.
[**GetPlansPolicy**](PlansAPI.md#GetPlansPolicy) | **Get** /v1/plans/policy | Returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.
[**GetPlansRegions**](PlansAPI.md#GetPlansRegions) | **Get** /v1/plans/regions | Returns the regions cloud capacity is offered in, each with its display name and physical location.
[**GetPlansResolveById**](PlansAPI.md#GetPlansResolveById) | **Get** /v1/plans/resolve/{id} | Resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.
[**GetPlansSchema**](PlansAPI.md#GetPlansSchema) | **Get** /v1/plans/schema | Returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.
[**GetPlansStorage**](PlansAPI.md#GetPlansStorage) | **Get** /v1/plans/storage | Returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.
[**GetPlansSubscriptions**](PlansAPI.md#GetPlansSubscriptions) | **Get** /v1/plans/subscriptions | Returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.
[**GetPlansTools**](PlansAPI.md#GetPlansTools) | **Get** /v1/plans/tools | Returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.
[**GetPlansVocab**](PlansAPI.md#GetPlansVocab) | **Get** /v1/plans/vocab | Returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.
[**PostPlansEntries**](PlansAPI.md#PostPlansEntries) | **Post** /v1/plans/entries | Add a subscription plan
[**PostPlansSeed**](PlansAPI.md#PostPlansSeed) | **Post** /v1/plans/seed | Seed the embedded plan catalog, without overwriting administrative edits
[**PutPlansEntriesBySlug**](PlansAPI.md#PutPlansEntriesBySlug) | **Put** /v1/plans/entries/{slug} | Edit a plan, leaving the fields you omit alone



## DeletePlansEntriesBySlug

> DeletePlansEntriesBySlug(ctx, slug).Execute()

Remove a plan from the authority



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlansAPI.DeletePlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.DeletePlansEntriesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlans

> PlanList GetPlans(ctx).Execute()

Returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller's catalog.



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
	resp, r, err := apiClient.PlansAPI.GetPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlans`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


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


## GetPlansBlockchain

> PlanList GetPlansBlockchain(ctx).Execute()

Returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.



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
	resp, r, err := apiClient.PlansAPI.GetPlansBlockchain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansBlockchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansBlockchain`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansBlockchain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansBlockchainRequest struct via the builder pattern


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


## GetPlansCloud

> PlanList GetPlansCloud(ctx).Execute()

Returns the cloud plan catalog.



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
	resp, r, err := apiClient.PlansAPI.GetPlansCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansCloud`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansCloudRequest struct via the builder pattern


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


## GetPlansDns

> PlanList GetPlansDns(ctx).Execute()

ListDNSPlans returns the DNS plan catalog: the tiers priced on zones, records per zone and queries per day.



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
	resp, r, err := apiClient.PlansAPI.GetPlansDns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansDns`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansDns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansDnsRequest struct via the builder pattern


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


## GetPlansEntitlementsById

> PlanEntitlements GetPlansEntitlementsById(ctx, id).Execute()

Returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.



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
	id := "team" // string | ID is the plan's catalog id or slug — \"pro\", \"team\", \"world-enterprise\", \"rpc-growth\". Both are matched, so a slug resolves the plan it names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlansEntitlementsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansEntitlementsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansEntitlementsById`: PlanEntitlements
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansEntitlementsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;pro\&quot;, \&quot;team\&quot;, \&quot;world-enterprise\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansEntitlementsByIdRequest struct via the builder pattern


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


## GetPlansEntries

> GetPlansEntries(ctx).Execute()

The raw plan authority rows



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
	r, err := apiClient.PlansAPI.GetPlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlansGpu

> PlanTierList GetPlansGpu(ctx).Execute()

ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPUs, host memory and hourly price.



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
	resp, r, err := apiClient.PlansAPI.GetPlansGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansGpu`: PlanTierList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansGpuRequest struct via the builder pattern


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


## GetPlansHealth

> PlanHealth GetPlansHealth(ctx).Execute()

Health reports that the plans subsystem is mounted and serving.



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
	resp, r, err := apiClient.PlansAPI.GetPlansHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansHealth`: PlanHealth
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansHealthRequest struct via the builder pattern


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


## GetPlansPolicy

> map[string]interface{} GetPlansPolicy(ctx).Execute()

Returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.



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
	resp, r, err := apiClient.PlansAPI.GetPlansPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansPolicyRequest struct via the builder pattern


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


## GetPlansRegions

> PlanRegionList GetPlansRegions(ctx).Execute()

Returns the regions cloud capacity is offered in, each with its display name and physical location.



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
	resp, r, err := apiClient.PlansAPI.GetPlansRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansRegions`: PlanRegionList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRegionsRequest struct via the builder pattern


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


## GetPlansResolveById

> PlanResolution GetPlansResolveById(ctx, id).Execute()

Resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.



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
	id := "pro" // string | ID is the plan's catalog id or slug — \"pro\", \"team\", \"world-enterprise\", \"rpc-growth\". Both are matched, so a slug resolves the plan it names.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlansResolveById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansResolveById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansResolveById`: PlanResolution
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansResolveById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;pro\&quot;, \&quot;team\&quot;, \&quot;world-enterprise\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansResolveByIdRequest struct via the builder pattern


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


## GetPlansSchema

> PlanSchemas GetPlansSchema(ctx).Execute()

Returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.



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
	resp, r, err := apiClient.PlansAPI.GetPlansSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansSchema`: PlanSchemas
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansSchemaRequest struct via the builder pattern


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


## GetPlansStorage

> map[string]interface{} GetPlansStorage(ctx).Execute()

Returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.



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
	resp, r, err := apiClient.PlansAPI.GetPlansStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansStorage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansStorageRequest struct via the builder pattern


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


## GetPlansSubscriptions

> PlanList GetPlansSubscriptions(ctx).Execute()

Returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.



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
	resp, r, err := apiClient.PlansAPI.GetPlansSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansSubscriptions`: PlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansSubscriptionsRequest struct via the builder pattern


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


## GetPlansTools

> PlanToolList GetPlansTools(ctx).Execute()

Returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.



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
	resp, r, err := apiClient.PlansAPI.GetPlansTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansTools`: PlanToolList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansToolsRequest struct via the builder pattern


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


## GetPlansVocab

> PlanVocab GetPlansVocab(ctx).Execute()

Returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.



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
	resp, r, err := apiClient.PlansAPI.GetPlansVocab(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlansVocab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlansVocab`: PlanVocab
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlansVocab`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansVocabRequest struct via the builder pattern


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


## PostPlansEntries

> PostPlansEntries(ctx).Execute()

Add a subscription plan



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
	r, err := apiClient.PlansAPI.PostPlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PostPlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlansSeed

> PostPlansSeed(ctx).Execute()

Seed the embedded plan catalog, without overwriting administrative edits



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
	r, err := apiClient.PlansAPI.PostPlansSeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PostPlansSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlansSeedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlansEntriesBySlug

> PutPlansEntriesBySlug(ctx, slug).Execute()

Edit a plan, leaving the fields you omit alone



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
	slug := "slug_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlansAPI.PutPlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.PutPlansEntriesBySlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

