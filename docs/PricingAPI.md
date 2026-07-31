# \PricingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Pricing**](PricingAPI.md#CloudGetV1Pricing) | **Get** /v1/pricing | GetPricing returns the whole pricing catalog in one document: Zen and third-party models, providers, model families, the free-model list, plan and infrastructure pricing.
[**CloudGetV1PricingBase**](PricingAPI.md#CloudGetV1PricingBase) | **Get** /v1/pricing/base | ListBasePlans returns the Hanzo Base plans — the managed-instance tiers, each with its monthly and annual price, storage and request allowances and feature list.
[**CloudGetV1PricingBlockchain**](PricingAPI.md#CloudGetV1PricingBlockchain) | **Get** /v1/pricing/blockchain | ListBlockchainPlans returns the blockchain access plans — the RPC and node tiers, each with its monthly price, compute-unit allowance and feature list.
[**CloudGetV1PricingCloud**](PricingAPI.md#CloudGetV1PricingCloud) | **Get** /v1/pricing/cloud | GetCloudPricing returns the public cloud section of the catalog in one document: its instance plans, its regions and its block-storage prices.
[**CloudGetV1PricingCloudPlans**](PricingAPI.md#CloudGetV1PricingCloudPlans) | **Get** /v1/pricing/cloud/plans | GetCloudPlans returns just the cloud instance plans — each with its vCPU, memory, disk, CPU type, VM allowance, feature list and monthly and hourly price.
[**CloudGetV1PricingCloudRegions**](PricingAPI.md#CloudGetV1PricingCloudRegions) | **Get** /v1/pricing/cloud/regions | GetCloudRegions returns the regions a cloud instance can be placed in, each with its id, display name and physical location.
[**CloudGetV1PricingCloudStorage**](PricingAPI.md#CloudGetV1PricingCloudStorage) | **Get** /v1/pricing/cloud/storage | GetCloudStoragePricing returns the block-storage prices of the cloud section: the per-GB monthly rate and the volume size bounds a caller may ask for.
[**CloudGetV1PricingCompute**](PricingAPI.md#CloudGetV1PricingCompute) | **Get** /v1/pricing/compute | GetComputePricing returns the compute section of the catalog: the cloud provider and region the prices are quoted for, the monthly markup applied to them, the full instance-size tier list and the named presets.
[**CloudGetV1PricingComputePresets**](PricingAPI.md#CloudGetV1PricingComputePresets) | **Get** /v1/pricing/compute/presets | GetComputePresets returns just the named compute sizes — the short, human-labelled list (\&quot;Starter\&quot;, \&quot;Pro\&quot;) a size picker renders, each carrying its provider slug, vCPU, memory, disk and price.
[**CloudGetV1PricingFeatured**](PricingAPI.md#CloudGetV1PricingFeatured) | **Get** /v1/pricing/featured | ListFeaturedModels returns the models the catalog highlights, filtered to what the caller&#39;s org may see.
[**CloudGetV1PricingFree**](PricingAPI.md#CloudGetV1PricingFree) | **Get** /v1/pricing/free | ListFreeModels returns the models that cost nothing to call, filtered to what the caller&#39;s org may see.
[**CloudGetV1PricingGpu**](PricingAPI.md#CloudGetV1PricingGpu) | **Get** /v1/pricing/gpu | ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPU, host memory and hourly price.
[**CloudGetV1PricingHealth**](PricingAPI.md#CloudGetV1PricingHealth) | **Get** /v1/pricing/health | Health reports that the pricing subsystem is mounted and serving.
[**CloudGetV1PricingIam**](PricingAPI.md#CloudGetV1PricingIam) | **Get** /v1/pricing/iam | ListIAMPlans returns the identity plans — the Hanzo IAM tiers, each with its monthly and annual price, monthly-active-user allowance and feature list.
[**CloudGetV1PricingModelName**](PricingAPI.md#CloudGetV1PricingModelName) | **Get** /v1/pricing/model/{name} | GetModel returns one model&#39;s catalog entry — its pricing, context window and capabilities as the pricing source records them.
[**CloudGetV1PricingModels**](PricingAPI.md#CloudGetV1PricingModels) | **Get** /v1/pricing/models | ListModels returns the whole model catalog — Hanzo&#39;s own Zen models and every third-party model — filtered to what the caller&#39;s org may see.
[**CloudGetV1PricingPaas**](PricingAPI.md#CloudGetV1PricingPaas) | **Get** /v1/pricing/paas | ListPaaSPlans returns the application-hosting plans — the deploy-and-host tiers, each with its monthly and annual price, app and memory allowances and feature list.
[**CloudGetV1PricingPolicy2**](PricingAPI.md#CloudGetV1PricingPolicy2) | **Get** /v1/pricing/policy | GetPricingPolicy returns the pricing policy document: the revenue-sharing terms (the idle-resale share and the open-source share, each with its percentage and who is eligible) and the commitments Hanzo makes about how it bills — no hidden fees, no egress charges, no surprise bills.
[**CloudGetV1PricingProviders**](PricingAPI.md#CloudGetV1PricingProviders) | **Get** /v1/pricing/providers | ListProviders returns the model providers the catalog knows, each with its info object, filtered to what the caller&#39;s org may see.
[**CloudGetV1PricingSubscriptions**](PricingAPI.md#CloudGetV1PricingSubscriptions) | **Get** /v1/pricing/subscriptions | ListSubscriptionPlans returns the API subscription plans — the account-level tiers a customer subscribes to, each with its monthly and annual price, included credit, rate limits and feature list.
[**CloudGetV1PricingSummary**](PricingAPI.md#CloudGetV1PricingSummary) | **Get** /v1/pricing/summary | GetPricingSummary returns the catalog&#39;s headline statistics — model counts by family and the provider directory.
[**CloudGetV1PricingTools**](PricingAPI.md#CloudGetV1PricingTools) | **Get** /v1/pricing/tools | ListToolPrices returns the per-use tool prices — web search, code interpreter, file storage, image generation, speech-to-text and text-to-speech — each with the unit it is billed by and its price in that unit.
[**CloudPostV1PricingSync**](PricingAPI.md#CloudPostV1PricingSync) | **Post** /v1/pricing/sync | SyncPricing refreshes the third-party section of the catalog from its upstream listings and returns the time the refreshed catalog was stamped with.



## CloudGetV1Pricing

> map[string]map[string]interface{} CloudGetV1Pricing(ctx).Execute()

GetPricing returns the whole pricing catalog in one document: Zen and third-party models, providers, model families, the free-model list, plan and infrastructure pricing.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1Pricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1Pricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Pricing`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1Pricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingBase

> CloudPricingPlanList CloudGetV1PricingBase(ctx).Execute()

ListBasePlans returns the Hanzo Base plans — the managed-instance tiers, each with its monthly and annual price, storage and request allowances and feature list.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingBase(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingBase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingBase`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingBase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingBaseRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingBlockchain

> CloudPricingPlanList CloudGetV1PricingBlockchain(ctx).Execute()

ListBlockchainPlans returns the blockchain access plans — the RPC and node tiers, each with its monthly price, compute-unit allowance and feature list.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingBlockchain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingBlockchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingBlockchain`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingBlockchain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingBlockchainRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingCloud

> map[string]map[string]interface{} CloudGetV1PricingCloud(ctx).Execute()

GetCloudPricing returns the public cloud section of the catalog in one document: its instance plans, its regions and its block-storage prices.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingCloud`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingCloudRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingCloudPlans

> CloudPricingPlanList CloudGetV1PricingCloudPlans(ctx).Execute()

GetCloudPlans returns just the cloud instance plans — each with its vCPU, memory, disk, CPU type, VM allowance, feature list and monthly and hourly price.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingCloudPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingCloudPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingCloudPlans`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingCloudPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingCloudPlansRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingCloudRegions

> CloudPricingRegionList CloudGetV1PricingCloudRegions(ctx).Execute()

GetCloudRegions returns the regions a cloud instance can be placed in, each with its id, display name and physical location.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingCloudRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingCloudRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingCloudRegions`: CloudPricingRegionList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingCloudRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingCloudRegionsRequest struct via the builder pattern


### Return type

[**CloudPricingRegionList**](CloudPricingRegionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingCloudStorage

> map[string]map[string]interface{} CloudGetV1PricingCloudStorage(ctx).Execute()

GetCloudStoragePricing returns the block-storage prices of the cloud section: the per-GB monthly rate and the volume size bounds a caller may ask for.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingCloudStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingCloudStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingCloudStorage`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingCloudStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingCloudStorageRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingCompute

> map[string]map[string]interface{} CloudGetV1PricingCompute(ctx).Execute()

GetComputePricing returns the compute section of the catalog: the cloud provider and region the prices are quoted for, the monthly markup applied to them, the full instance-size tier list and the named presets.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingCompute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingCompute`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingCompute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingComputeRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingComputePresets

> CloudPricingPresetList CloudGetV1PricingComputePresets(ctx).Execute()

GetComputePresets returns just the named compute sizes — the short, human-labelled list (\"Starter\", \"Pro\") a size picker renders, each carrying its provider slug, vCPU, memory, disk and price.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingComputePresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingComputePresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingComputePresets`: CloudPricingPresetList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingComputePresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingComputePresetsRequest struct via the builder pattern


### Return type

[**CloudPricingPresetList**](CloudPricingPresetList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingFeatured

> CloudPricingModelList CloudGetV1PricingFeatured(ctx).Execute()

ListFeaturedModels returns the models the catalog highlights, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingFeatured(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingFeatured``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingFeatured`: CloudPricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingFeatured`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingFeaturedRequest struct via the builder pattern


### Return type

[**CloudPricingModelList**](CloudPricingModelList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingFree

> CloudPricingModelList CloudGetV1PricingFree(ctx).Execute()

ListFreeModels returns the models that cost nothing to call, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingFree(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingFree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingFree`: CloudPricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingFree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingFreeRequest struct via the builder pattern


### Return type

[**CloudPricingModelList**](CloudPricingModelList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingGpu

> CloudPricingTierList CloudGetV1PricingGpu(ctx).Execute()

ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPU, host memory and hourly price.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingGpu`: CloudPricingTierList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingGpuRequest struct via the builder pattern


### Return type

[**CloudPricingTierList**](CloudPricingTierList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingHealth

> CloudPricingHealth CloudGetV1PricingHealth(ctx).Execute()

Health reports that the pricing subsystem is mounted and serving.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingHealth`: CloudPricingHealth
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingHealthRequest struct via the builder pattern


### Return type

[**CloudPricingHealth**](CloudPricingHealth.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingIam

> CloudPricingPlanList CloudGetV1PricingIam(ctx).Execute()

ListIAMPlans returns the identity plans — the Hanzo IAM tiers, each with its monthly and annual price, monthly-active-user allowance and feature list.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingIam(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingIam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingIam`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingIam`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingIamRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingModelName

> map[string]map[string]interface{} CloudGetV1PricingModelName(ctx, name).Execute()

GetModel returns one model's catalog entry — its pricing, context window and capabilities as the pricing source records them.



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
	name := "zen4" // string | Name is the model's name or its slugged id (\"zen4\", \"anthropic/claude-opus-4.6\"), matched case-insensitively. It comes from the path: the URL is the addressing authority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingModelName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingModelName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingModelName`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingModelName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the model&#39;s name or its slugged id (\&quot;zen4\&quot;, \&quot;anthropic/claude-opus-4.6\&quot;), matched case-insensitively. It comes from the path: the URL is the addressing authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingModelNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingModels

> CloudPricingModelList CloudGetV1PricingModels(ctx).Execute()

ListModels returns the whole model catalog — Hanzo's own Zen models and every third-party model — filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingModels`: CloudPricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingModelsRequest struct via the builder pattern


### Return type

[**CloudPricingModelList**](CloudPricingModelList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingPaas

> CloudPricingPlanList CloudGetV1PricingPaas(ctx).Execute()

ListPaaSPlans returns the application-hosting plans — the deploy-and-host tiers, each with its monthly and annual price, app and memory allowances and feature list.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingPaas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingPaas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingPaas`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingPaas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingPaasRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingPolicy2

> map[string]map[string]interface{} CloudGetV1PricingPolicy2(ctx).Execute()

GetPricingPolicy returns the pricing policy document: the revenue-sharing terms (the idle-resale share and the open-source share, each with its percentage and who is eligible) and the commitments Hanzo makes about how it bills — no hidden fees, no egress charges, no surprise bills.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingPolicy2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingPolicy2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingPolicy2`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingPolicy2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingPolicy2Request struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingProviders

> CloudPricingProviderList CloudGetV1PricingProviders(ctx).Execute()

ListProviders returns the model providers the catalog knows, each with its info object, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingProviders`: CloudPricingProviderList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingProvidersRequest struct via the builder pattern


### Return type

[**CloudPricingProviderList**](CloudPricingProviderList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingSubscriptions

> CloudPricingPlanList CloudGetV1PricingSubscriptions(ctx).Execute()

ListSubscriptionPlans returns the API subscription plans — the account-level tiers a customer subscribes to, each with its monthly and annual price, included credit, rate limits and feature list.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingSubscriptions`: CloudPricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingSubscriptionsRequest struct via the builder pattern


### Return type

[**CloudPricingPlanList**](CloudPricingPlanList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingSummary

> map[string]map[string]interface{} CloudGetV1PricingSummary(ctx).Execute()

GetPricingSummary returns the catalog's headline statistics — model counts by family and the provider directory.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingSummary`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingSummaryRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1PricingTools

> CloudPricingToolList CloudGetV1PricingTools(ctx).Execute()

ListToolPrices returns the per-use tool prices — web search, code interpreter, file storage, image generation, speech-to-text and text-to-speech — each with the unit it is billed by and its price in that unit.



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
	resp, r, err := apiClient.PricingAPI.CloudGetV1PricingTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudGetV1PricingTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1PricingTools`: CloudPricingToolList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudGetV1PricingTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PricingToolsRequest struct via the builder pattern


### Return type

[**CloudPricingToolList**](CloudPricingToolList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1PricingSync

> CloudPricingSyncOut CloudPostV1PricingSync(ctx).Execute()

SyncPricing refreshes the third-party section of the catalog from its upstream listings and returns the time the refreshed catalog was stamped with.



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
	resp, r, err := apiClient.PricingAPI.CloudPostV1PricingSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.CloudPostV1PricingSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1PricingSync`: CloudPricingSyncOut
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.CloudPostV1PricingSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1PricingSyncRequest struct via the builder pattern


### Return type

[**CloudPricingSyncOut**](CloudPricingSyncOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

