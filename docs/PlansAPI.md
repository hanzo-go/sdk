# \PlansAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1PlansEntriesBySlug**](PlansAPI.md#CloudDeleteV1PlansEntriesBySlug) | **Delete** /v1/plans/entries/{slug} | 
[**CloudGetV1Plans**](PlansAPI.md#CloudGetV1Plans) | **Get** /v1/plans | ListCloudPlans returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller&#39;s catalog.
[**CloudGetV1PlansBlockchain**](PlansAPI.md#CloudGetV1PlansBlockchain) | **Get** /v1/plans/blockchain | ListBlockchainPlans returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.
[**CloudGetV1PlansCloud**](PlansAPI.md#CloudGetV1PlansCloud) | **Get** /v1/plans/cloud | ListCloudCapacityPlans returns the cloud plan catalog.
[**CloudGetV1PlansDns**](PlansAPI.md#CloudGetV1PlansDns) | **Get** /v1/plans/dns | ListDNSPlans returns the DNS plan catalog: the tiers priced on zones, records per zone and queries per day.
[**CloudGetV1PlansEntitlementsId**](PlansAPI.md#CloudGetV1PlansEntitlementsId) | **Get** /v1/plans/entitlements/{id} | GetPlanEntitlements returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.
[**CloudGetV1PlansEntries**](PlansAPI.md#CloudGetV1PlansEntries) | **Get** /v1/plans/entries | 
[**CloudGetV1PlansGpu**](PlansAPI.md#CloudGetV1PlansGpu) | **Get** /v1/plans/gpu | ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPUs, host memory and hourly price.
[**CloudGetV1PlansHealth**](PlansAPI.md#CloudGetV1PlansHealth) | **Get** /v1/plans/health | Health reports that the plans subsystem is mounted and serving.
[**CloudGetV1PlansPolicy**](PlansAPI.md#CloudGetV1PlansPolicy) | **Get** /v1/plans/policy | GetPricingPolicy returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.
[**CloudGetV1PlansRegions**](PlansAPI.md#CloudGetV1PlansRegions) | **Get** /v1/plans/regions | ListRegions returns the regions cloud capacity is offered in, each with its display name and physical location.
[**CloudGetV1PlansResolveId**](PlansAPI.md#CloudGetV1PlansResolveId) | **Get** /v1/plans/resolve/{id} | ResolvePlan resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.
[**CloudGetV1PlansSchema**](PlansAPI.md#CloudGetV1PlansSchema) | **Get** /v1/plans/schema | GetPlanSchemas returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.
[**CloudGetV1PlansStorage**](PlansAPI.md#CloudGetV1PlansStorage) | **Get** /v1/plans/storage | GetStoragePricing returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.
[**CloudGetV1PlansSubscriptions**](PlansAPI.md#CloudGetV1PlansSubscriptions) | **Get** /v1/plans/subscriptions | ListSubscriptionPlans returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.
[**CloudGetV1PlansTools**](PlansAPI.md#CloudGetV1PlansTools) | **Get** /v1/plans/tools | ListToolPrices returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.
[**CloudGetV1PlansVocab**](PlansAPI.md#CloudGetV1PlansVocab) | **Get** /v1/plans/vocab | GetEntitlementVocabulary returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.
[**CloudPostV1PlansEntries**](PlansAPI.md#CloudPostV1PlansEntries) | **Post** /v1/plans/entries | 
[**CloudPostV1PlansSeed**](PlansAPI.md#CloudPostV1PlansSeed) | **Post** /v1/plans/seed | 
[**CloudPutV1PlansEntriesBySlug**](PlansAPI.md#CloudPutV1PlansEntriesBySlug) | **Put** /v1/plans/entries/{slug} | 



## CloudDeleteV1PlansEntriesBySlug

> CloudDeleteV1PlansEntriesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlansAPI.CloudDeleteV1PlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudDeleteV1PlansEntriesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1PlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Plans

> CloudPlanList CloudGetV1Plans(ctx).Execute()

ListCloudPlans returns the Hanzo cloud plan catalog: every cloud tier with its price, included capacity, limits and feature list, scoped to the caller's catalog.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1Plans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1Plans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Plans`: CloudPlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1Plans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansRequest struct via the builder pattern


### Return type

[**CloudPlanList**](CloudPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansBlockchain

> CloudPlanList CloudGetV1PlansBlockchain(ctx).Execute()

ListBlockchainPlans returns the blockchain RPC plan catalog: the tiers metered in monthly compute units, with their prices, limits and overage terms.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansBlockchain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansBlockchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansBlockchain`: CloudPlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansBlockchain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansBlockchainRequest struct via the builder pattern


### Return type

[**CloudPlanList**](CloudPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansCloud

> CloudPlanList CloudGetV1PlansCloud(ctx).Execute()

ListCloudCapacityPlans returns the cloud plan catalog.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansCloud`: CloudPlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansCloudRequest struct via the builder pattern


### Return type

[**CloudPlanList**](CloudPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansDns

> CloudPlanList CloudGetV1PlansDns(ctx).Execute()

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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansDns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansDns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansDns`: CloudPlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansDns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansDnsRequest struct via the builder pattern


### Return type

[**CloudPlanList**](CloudPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansEntitlementsId

> CloudPlanEntitlements CloudGetV1PlansEntitlementsId(ctx, id).Execute()

GetPlanEntitlements returns what one plan GRANTS and not what it costs: the canonical namespaced entitlement block and the flat license-feature list derived from it.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansEntitlementsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansEntitlementsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansEntitlementsId`: CloudPlanEntitlements
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansEntitlementsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;pro\&quot;, \&quot;team\&quot;, \&quot;world-enterprise\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansEntitlementsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPlanEntitlements**](CloudPlanEntitlements.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansEntries

> CloudGetV1PlansEntries(ctx).Execute()



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
	r, err := apiClient.PlansAPI.CloudGetV1PlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansGpu

> CloudPlanTierList CloudGetV1PlansGpu(ctx).Execute()

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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansGpu`: CloudPlanTierList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansGpuRequest struct via the builder pattern


### Return type

[**CloudPlanTierList**](CloudPlanTierList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansHealth

> CloudPlanHealth CloudGetV1PlansHealth(ctx).Execute()

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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansHealth`: CloudPlanHealth
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansHealthRequest struct via the builder pattern


### Return type

[**CloudPlanHealth**](CloudPlanHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansPolicy

> map[string]interface{} CloudGetV1PlansPolicy(ctx).Execute()

GetPricingPolicy returns the published pricing policy: whether pricing is transparent, the revenue-sharing terms (idle compute resale and the open-source share) and the principles the catalog is priced by.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansPolicyRequest struct via the builder pattern


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


## CloudGetV1PlansRegions

> CloudPlanRegionList CloudGetV1PlansRegions(ctx).Execute()

ListRegions returns the regions cloud capacity is offered in, each with its display name and physical location.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansRegions`: CloudPlanRegionList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansRegionsRequest struct via the builder pattern


### Return type

[**CloudPlanRegionList**](CloudPlanRegionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansResolveId

> CloudPlanResolution CloudGetV1PlansResolveId(ctx, id).Execute()

ResolvePlan resolves one plan to everything a consumer of the catalog needs at once: its canonical entitlement block, the flat license-feature list a signed license carries, its billing reference, and the catalog it came from.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansResolveId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansResolveId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansResolveId`: CloudPlanResolution
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansResolveId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the plan&#39;s catalog id or slug — \&quot;pro\&quot;, \&quot;team\&quot;, \&quot;world-enterprise\&quot;, \&quot;rpc-growth\&quot;. Both are matched, so a slug resolves the plan it names. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansResolveIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudPlanResolution**](CloudPlanResolution.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansSchema

> CloudPlanSchemas CloudGetV1PlansSchema(ctx).Execute()

GetPlanSchemas returns the two JSON Schema documents this surface speaks: entitlements.schema.json, which declares every entitlement key with its type, unit and enum, and plan.schema.json, which a catalog plan record conforms to.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansSchema(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansSchema`: CloudPlanSchemas
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansSchema`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansSchemaRequest struct via the builder pattern


### Return type

[**CloudPlanSchemas**](CloudPlanSchemas.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansStorage

> map[string]interface{} CloudGetV1PlansStorage(ctx).Execute()

GetStoragePricing returns the block-storage price block: the price per GB per month and the volume size bounds a cloud plan may attach.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansStorage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansStorageRequest struct via the builder pattern


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


## CloudGetV1PlansSubscriptions

> CloudPlanList CloudGetV1PlansSubscriptions(ctx).Execute()

ListSubscriptionPlans returns the subscription ladder — the personal and team tiers a customer buys to use the cloud, each with its monthly and annual price, seat rules, limits and billing reference.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansSubscriptions`: CloudPlanList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansSubscriptionsRequest struct via the builder pattern


### Return type

[**CloudPlanList**](CloudPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansTools

> CloudPlanToolList CloudGetV1PlansTools(ctx).Execute()

ListToolPrices returns the per-use price of every metered tool — web search, code interpreter, image generation, speech — each with the unit it is billed in.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansTools`: CloudPlanToolList
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansToolsRequest struct via the builder pattern


### Return type

[**CloudPlanToolList**](CloudPlanToolList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PlansVocab

> CloudPlanVocab CloudGetV1PlansVocab(ctx).Execute()

GetEntitlementVocabulary returns the entitlement key vocabulary: every key with its namespace, JSON type, nullability, unit, enum and title, the list of namespaces, and the engine features a license can grant.



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
	resp, r, err := apiClient.PlansAPI.CloudGetV1PlansVocab(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudGetV1PlansVocab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PlansVocab`: CloudPlanVocab
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CloudGetV1PlansVocab`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PlansVocabRequest struct via the builder pattern


### Return type

[**CloudPlanVocab**](CloudPlanVocab.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PlansEntries

> CloudPostV1PlansEntries(ctx).Execute()



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
	r, err := apiClient.PlansAPI.CloudPostV1PlansEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudPostV1PlansEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PlansEntriesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PlansSeed

> CloudPostV1PlansSeed(ctx).Execute()



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
	r, err := apiClient.PlansAPI.CloudPostV1PlansSeed(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudPostV1PlansSeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PlansSeedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1PlansEntriesBySlug

> CloudPutV1PlansEntriesBySlug(ctx, slug).Execute()



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
	r, err := apiClient.PlansAPI.CloudPutV1PlansEntriesBySlug(context.Background(), slug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CloudPutV1PlansEntriesBySlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPutV1PlansEntriesBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

