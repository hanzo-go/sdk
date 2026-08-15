# \PricingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPricing**](PricingAPI.md#GetPricing) | **Get** /v1/pricing | Returns the whole pricing catalog in one document: Zen and third-party models, providers, model families, the free-model list, plan and infrastructure pricing.
[**GetPricingBase**](PricingAPI.md#GetPricingBase) | **Get** /v1/pricing/base | Returns the Hanzo Base plans — the managed-instance tiers, each with its monthly and annual price, storage and request allowances and feature list.
[**GetPricingBlockchain**](PricingAPI.md#GetPricingBlockchain) | **Get** /v1/pricing/blockchain | Returns the blockchain access plans — the RPC and node tiers, each with its monthly price, compute-unit allowance and feature list.
[**GetPricingCloud**](PricingAPI.md#GetPricingCloud) | **Get** /v1/pricing/cloud | Returns the public cloud section of the catalog in one document: its instance plans, its regions and its block-storage prices.
[**GetPricingCloudPlans**](PricingAPI.md#GetPricingCloudPlans) | **Get** /v1/pricing/cloud/plans | Returns just the cloud instance plans — each with its vCPU, memory, disk, CPU type, VM allowance, feature list and monthly and hourly price.
[**GetPricingCloudRegions**](PricingAPI.md#GetPricingCloudRegions) | **Get** /v1/pricing/cloud/regions | Returns the regions a cloud instance can be placed in, each with its id, display name and physical location.
[**GetPricingCloudStorage**](PricingAPI.md#GetPricingCloudStorage) | **Get** /v1/pricing/cloud/storage | Returns the block-storage prices of the cloud section: the per-GB monthly rate and the volume size bounds a caller may ask for.
[**GetPricingCompute**](PricingAPI.md#GetPricingCompute) | **Get** /v1/pricing/compute | Returns the compute section of the catalog: the cloud provider and region the prices are quoted for, the monthly markup applied to them, the full instance-size tier list and the named presets.
[**GetPricingComputePresets**](PricingAPI.md#GetPricingComputePresets) | **Get** /v1/pricing/compute/presets | Returns just the named compute sizes — the short, human-labelled list (\&quot;Starter\&quot;, \&quot;Pro\&quot;) a size picker renders, each carrying its provider slug, vCPU, memory, disk and price.
[**GetPricingDatastore**](PricingAPI.md#GetPricingDatastore) | **Get** /v1/pricing/datastore | Returns the Hanzo Datastore rate card: the tier list, the per-GB storage and egress usage rates, the annual discount and the trial.
[**GetPricingFeatured**](PricingAPI.md#GetPricingFeatured) | **Get** /v1/pricing/featured | Returns the models the catalog highlights, filtered to what the caller&#39;s org may see.
[**GetPricingFree**](PricingAPI.md#GetPricingFree) | **Get** /v1/pricing/free | Returns the models that cost nothing to call, filtered to what the caller&#39;s org may see.
[**GetPricingGpu**](PricingAPI.md#GetPricingGpu) | **Get** /v1/pricing/gpu | ListGPUTiers returns the rentable GPU configurations, each with its accelerator count and model, VRAM, vCPU, host memory and hourly price.
[**GetPricingHealth**](PricingAPI.md#GetPricingHealth) | **Get** /v1/pricing/health | Health reports that the pricing subsystem is mounted and serving.
[**GetPricingIam**](PricingAPI.md#GetPricingIam) | **Get** /v1/pricing/iam | ListIAMPlans returns the identity plans — the Hanzo IAM tiers, each with its monthly and annual price, monthly-active-user allowance and feature list.
[**GetPricingModelByName**](PricingAPI.md#GetPricingModelByName) | **Get** /v1/pricing/model/{name} | Returns one model&#39;s catalog entry — its pricing, context window and capabilities as the pricing source records them.
[**GetPricingModels**](PricingAPI.md#GetPricingModels) | **Get** /v1/pricing/models | Returns the whole model catalog — every model the gateway serves, Zen and third-party alike — filtered to what the caller&#39;s org may see.
[**GetPricingPaas**](PricingAPI.md#GetPricingPaas) | **Get** /v1/pricing/paas | ListPaaSPlans returns the application-hosting plans — the deploy-and-host tiers, each with its monthly and annual price, app and memory allowances and feature list.
[**GetPricingPolicy**](PricingAPI.md#GetPricingPolicy) | **Get** /v1/pricing/policy | Returns the pricing policy document: the revenue-sharing terms (the idle-resale share and the open-source share, each with its percentage and who is eligible) and the commitments Hanzo makes about how it bills — no hidden fees, no egress charges, no surprise bills.
[**GetPricingProviders**](PricingAPI.md#GetPricingProviders) | **Get** /v1/pricing/providers | Returns the model providers the catalog knows, each with its info object, filtered to what the caller&#39;s org may see.
[**GetPricingServices**](PricingAPI.md#GetPricingServices) | **Get** /v1/pricing/services | Returns the managed-service rate cards — Search, Crawl, Vector, Console and Managed Services — each with its own tiers, and some with usage rates or a comparison table.
[**GetPricingSubscriptions**](PricingAPI.md#GetPricingSubscriptions) | **Get** /v1/pricing/subscriptions | Returns the API subscription plans — the account-level tiers a customer subscribes to, each with its monthly and annual price, included credit, rate limits and feature list.
[**GetPricingSummary**](PricingAPI.md#GetPricingSummary) | **Get** /v1/pricing/summary | Returns the catalog&#39;s headline statistics — model counts by family and the provider directory.
[**GetPricingTools**](PricingAPI.md#GetPricingTools) | **Get** /v1/pricing/tools | Returns the per-use tool prices — web search, code interpreter, file storage, image generation, speech-to-text and text-to-speech — each with the unit it is billed by and its price in that unit.
[**PostPricingSync**](PricingAPI.md#PostPricingSync) | **Post** /v1/pricing/sync | Refreshes the third-party section of the catalog from its upstream listings and returns the time the refreshed catalog was stamped with.



## GetPricing

> map[string]map[string]interface{} GetPricing(ctx).Execute()

Returns the whole pricing catalog in one document: Zen and third-party models, providers, model families, the free-model list, plan and infrastructure pricing.



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
	resp, r, err := apiClient.PricingAPI.GetPricing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricing`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingBase

> PricingPlanList GetPricingBase(ctx).Execute()

Returns the Hanzo Base plans — the managed-instance tiers, each with its monthly and annual price, storage and request allowances and feature list.



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
	resp, r, err := apiClient.PricingAPI.GetPricingBase(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingBase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingBase`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingBase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingBaseRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingBlockchain

> PricingPlanList GetPricingBlockchain(ctx).Execute()

Returns the blockchain access plans — the RPC and node tiers, each with its monthly price, compute-unit allowance and feature list.



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
	resp, r, err := apiClient.PricingAPI.GetPricingBlockchain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingBlockchain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingBlockchain`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingBlockchain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingBlockchainRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingCloud

> map[string]map[string]interface{} GetPricingCloud(ctx).Execute()

Returns the public cloud section of the catalog in one document: its instance plans, its regions and its block-storage prices.



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
	resp, r, err := apiClient.PricingAPI.GetPricingCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingCloud`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingCloudRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingCloudPlans

> PricingPlanList GetPricingCloudPlans(ctx).Execute()

Returns just the cloud instance plans — each with its vCPU, memory, disk, CPU type, VM allowance, feature list and monthly and hourly price.



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
	resp, r, err := apiClient.PricingAPI.GetPricingCloudPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingCloudPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingCloudPlans`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingCloudPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingCloudPlansRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingCloudRegions

> PricingRegionList GetPricingCloudRegions(ctx).Execute()

Returns the regions a cloud instance can be placed in, each with its id, display name and physical location.



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
	resp, r, err := apiClient.PricingAPI.GetPricingCloudRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingCloudRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingCloudRegions`: PricingRegionList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingCloudRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingCloudRegionsRequest struct via the builder pattern


### Return type

[**PricingRegionList**](PricingRegionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingCloudStorage

> map[string]map[string]interface{} GetPricingCloudStorage(ctx).Execute()

Returns the block-storage prices of the cloud section: the per-GB monthly rate and the volume size bounds a caller may ask for.



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
	resp, r, err := apiClient.PricingAPI.GetPricingCloudStorage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingCloudStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingCloudStorage`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingCloudStorage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingCloudStorageRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingCompute

> map[string]map[string]interface{} GetPricingCompute(ctx).Execute()

Returns the compute section of the catalog: the cloud provider and region the prices are quoted for, the monthly markup applied to them, the full instance-size tier list and the named presets.



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
	resp, r, err := apiClient.PricingAPI.GetPricingCompute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingCompute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingCompute`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingCompute`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingComputeRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingComputePresets

> PricingPresetList GetPricingComputePresets(ctx).Execute()

Returns just the named compute sizes — the short, human-labelled list (\"Starter\", \"Pro\") a size picker renders, each carrying its provider slug, vCPU, memory, disk and price.



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
	resp, r, err := apiClient.PricingAPI.GetPricingComputePresets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingComputePresets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingComputePresets`: PricingPresetList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingComputePresets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingComputePresetsRequest struct via the builder pattern


### Return type

[**PricingPresetList**](PricingPresetList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingDatastore

> map[string]map[string]interface{} GetPricingDatastore(ctx).Execute()

Returns the Hanzo Datastore rate card: the tier list, the per-GB storage and egress usage rates, the annual discount and the trial.



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
	resp, r, err := apiClient.PricingAPI.GetPricingDatastore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingDatastore`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingDatastore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingDatastoreRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingFeatured

> PricingModelList GetPricingFeatured(ctx).Execute()

Returns the models the catalog highlights, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.GetPricingFeatured(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingFeatured``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingFeatured`: PricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingFeatured`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingFeaturedRequest struct via the builder pattern


### Return type

[**PricingModelList**](PricingModelList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingFree

> PricingModelList GetPricingFree(ctx).Execute()

Returns the models that cost nothing to call, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.GetPricingFree(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingFree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingFree`: PricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingFree`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingFreeRequest struct via the builder pattern


### Return type

[**PricingModelList**](PricingModelList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingGpu

> PricingTierList GetPricingGpu(ctx).Execute()

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
	resp, r, err := apiClient.PricingAPI.GetPricingGpu(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingGpu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingGpu`: PricingTierList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingGpu`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingGpuRequest struct via the builder pattern


### Return type

[**PricingTierList**](PricingTierList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingHealth

> PricingHealth GetPricingHealth(ctx).Execute()

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
	resp, r, err := apiClient.PricingAPI.GetPricingHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingHealth`: PricingHealth
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingHealthRequest struct via the builder pattern


### Return type

[**PricingHealth**](PricingHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingIam

> PricingPlanList GetPricingIam(ctx).Execute()

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
	resp, r, err := apiClient.PricingAPI.GetPricingIam(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingIam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingIam`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingIam`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingIamRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingModelByName

> map[string]map[string]interface{} GetPricingModelByName(ctx, name).Execute()

Returns one model's catalog entry — its pricing, context window and capabilities as the pricing source records them.



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
	name := "zen4" // string | Name is the model's name or its slugged id (\"zen4\", \"acme/some-model-1\"), matched case-insensitively. It comes from the path: the URL is the addressing authority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PricingAPI.GetPricingModelByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingModelByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingModelByName`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingModelByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the model&#39;s name or its slugged id (\&quot;zen4\&quot;, \&quot;acme/some-model-1\&quot;), matched case-insensitively. It comes from the path: the URL is the addressing authority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingModelByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingModels

> PricingModelList GetPricingModels(ctx).Execute()

Returns the whole model catalog — every model the gateway serves, Zen and third-party alike — filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.GetPricingModels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingModels`: PricingModelList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingModelsRequest struct via the builder pattern


### Return type

[**PricingModelList**](PricingModelList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPaas

> PricingPlanList GetPricingPaas(ctx).Execute()

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
	resp, r, err := apiClient.PricingAPI.GetPricingPaas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingPaas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPaas`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingPaas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPaasRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingPolicy

> map[string]map[string]interface{} GetPricingPolicy(ctx).Execute()

Returns the pricing policy document: the revenue-sharing terms (the idle-resale share and the open-source share, each with its percentage and who is eligible) and the commitments Hanzo makes about how it bills — no hidden fees, no egress charges, no surprise bills.



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
	resp, r, err := apiClient.PricingAPI.GetPricingPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingPolicy`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingPolicyRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingProviders

> PricingProviderList GetPricingProviders(ctx).Execute()

Returns the model providers the catalog knows, each with its info object, filtered to what the caller's org may see.



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
	resp, r, err := apiClient.PricingAPI.GetPricingProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingProviders`: PricingProviderList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingProvidersRequest struct via the builder pattern


### Return type

[**PricingProviderList**](PricingProviderList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingServices

> map[string]map[string]interface{} GetPricingServices(ctx).Execute()

Returns the managed-service rate cards — Search, Crawl, Vector, Console and Managed Services — each with its own tiers, and some with usage rates or a comparison table.



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
	resp, r, err := apiClient.PricingAPI.GetPricingServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingServices`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingServicesRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingSubscriptions

> PricingPlanList GetPricingSubscriptions(ctx).Execute()

Returns the API subscription plans — the account-level tiers a customer subscribes to, each with its monthly and annual price, included credit, rate limits and feature list.



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
	resp, r, err := apiClient.PricingAPI.GetPricingSubscriptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingSubscriptions`: PricingPlanList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingSubscriptionsRequest struct via the builder pattern


### Return type

[**PricingPlanList**](PricingPlanList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingSummary

> map[string]map[string]interface{} GetPricingSummary(ctx).Execute()

Returns the catalog's headline statistics — model counts by family and the provider directory.



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
	resp, r, err := apiClient.PricingAPI.GetPricingSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingSummary`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingSummaryRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPricingTools

> PricingToolList GetPricingTools(ctx).Execute()

Returns the per-use tool prices — web search, code interpreter, file storage, image generation, speech-to-text and text-to-speech — each with the unit it is billed by and its price in that unit.



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
	resp, r, err := apiClient.PricingAPI.GetPricingTools(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.GetPricingTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPricingTools`: PricingToolList
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.GetPricingTools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingToolsRequest struct via the builder pattern


### Return type

[**PricingToolList**](PricingToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPricingSync

> PricingSyncOut PostPricingSync(ctx).Execute()

Refreshes the third-party section of the catalog from its upstream listings and returns the time the refreshed catalog was stamped with.



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
	resp, r, err := apiClient.PricingAPI.PostPricingSync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PricingAPI.PostPricingSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPricingSync`: PricingSyncOut
	fmt.Fprintf(os.Stdout, "Response from `PricingAPI.PostPricingSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostPricingSyncRequest struct via the builder pattern


### Return type

[**PricingSyncOut**](PricingSyncOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

