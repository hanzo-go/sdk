# \CampaignAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1CampaignId**](CampaignAPI.md#CloudDeleteV1CampaignId) | **Delete** /v1/campaign/{id} | DeleteCampaign removes one campaign of the caller&#39;s org and answers 204 with no body.
[**CloudDeleteV1CampaignIdChannelsKind**](CampaignAPI.md#CloudDeleteV1CampaignIdChannelsKind) | **Delete** /v1/campaign/{id}/channels/{kind} | RemoveCampaignChannel drops one channel from a campaign and returns the updated campaign.
[**CloudGetV1Campaign**](CampaignAPI.md#CloudGetV1Campaign) | **Get** /v1/campaign | ListCampaigns returns the org&#39;s campaigns, newest first, optionally narrowed to one status.
[**CloudGetV1CampaignId**](CampaignAPI.md#CloudGetV1CampaignId) | **Get** /v1/campaign/{id} | GetCampaign returns one campaign of the caller&#39;s org — its name, audience, creatives, channels with their per-channel launch state, schedule, budget and status.
[**CloudGetV1CampaignIdMetrics**](CampaignAPI.md#CloudGetV1CampaignIdMetrics) | **Get** /v1/campaign/{id}/metrics | CampaignMetrics returns a campaign&#39;s results over a window: the analytics funnel (impressions, clicks, conversions, revenue, visitors), the spend each channel&#39;s connector reports, and the derived growth KPIs — CTR, CVR, CAC and ROAS.
[**CloudGetV1CampaignSummary**](CampaignAPI.md#CloudGetV1CampaignSummary) | **Get** /v1/campaign/summary | SummarizeCampaigns returns the org&#39;s go-to-market roll-up: how many campaigns exist, how many are live, their total budget in cents, and which channel executors this deployment can actually reach.
[**CloudPostV1Campaign**](CampaignAPI.md#CloudPostV1Campaign) | **Post** /v1/campaign | CreateCampaign creates a campaign as a DRAFT and returns it.
[**CloudPostV1CampaignByIdLaunch**](CampaignAPI.md#CloudPostV1CampaignByIdLaunch) | **Post** /v1/campaign/{id}/launch | 
[**CloudPostV1CampaignByIdPause**](CampaignAPI.md#CloudPostV1CampaignByIdPause) | **Post** /v1/campaign/{id}/pause | 
[**CloudPostV1CampaignIdChannels**](CampaignAPI.md#CloudPostV1CampaignIdChannels) | **Post** /v1/campaign/{id}/channels | AddCampaignChannel adds a channel to a campaign, or REPLACES the one it already has of that kind, and returns the updated campaign.
[**CloudPutV1CampaignId**](CampaignAPI.md#CloudPutV1CampaignId) | **Put** /v1/campaign/{id} | UpdateCampaign rewrites a campaign&#39;s core fields — name, audience, creatives, schedule and budget — and returns the updated campaign.



## CloudDeleteV1CampaignId

> CloudDeleteV1CampaignId(ctx, id).Execute()

DeleteCampaign removes one campaign of the caller's org and answers 204 with no body.



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
	id := "id_example" // string | ID is the campaign's server-minted handle, \"cmp_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignAPI.CloudDeleteV1CampaignId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudDeleteV1CampaignId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign&#39;s server-minted handle, \&quot;cmp_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CampaignIdRequest struct via the builder pattern


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


## CloudDeleteV1CampaignIdChannelsKind

> CloudCampaignRecord CloudDeleteV1CampaignIdChannelsKind(ctx, id, kind).Execute()

RemoveCampaignChannel drops one channel from a campaign and returns the updated campaign.



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
	id := "id_example" // string | ID is the campaign, from the path.
	kind := "kind_example" // string | Kind is the channel to remove: paid, organic or email.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudDeleteV1CampaignIdChannelsKind(context.Background(), id, kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudDeleteV1CampaignIdChannelsKind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CampaignIdChannelsKind`: CloudCampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudDeleteV1CampaignIdChannelsKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign, from the path. | 
**kind** | **string** | Kind is the channel to remove: paid, organic or email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CampaignIdChannelsKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudCampaignRecord**](CloudCampaignRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Campaign

> CloudCampaignPage CloudGetV1Campaign(ctx).Status(status).Limit(limit).Execute()

ListCampaigns returns the org's campaigns, newest first, optionally narrowed to one status.



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
	status := "live" // string | Status keeps only campaigns in that state: draft, live, paused or failed. Empty means any. (optional)
	limit := int32(50) // int32 | Limit bounds the page. 0 or less means the default of 200; anything above 1000 is clamped to 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudGetV1Campaign(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudGetV1Campaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Campaign`: CloudCampaignPage
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudGetV1Campaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only campaigns in that state: draft, live, paused or failed. Empty means any. | 
 **limit** | **int32** | Limit bounds the page. 0 or less means the default of 200; anything above 1000 is clamped to 1000. | 

### Return type

[**CloudCampaignPage**](CloudCampaignPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CampaignId

> CloudCampaignRecord CloudGetV1CampaignId(ctx, id).Execute()

GetCampaign returns one campaign of the caller's org — its name, audience, creatives, channels with their per-channel launch state, schedule, budget and status.



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
	id := "id_example" // string | ID is the campaign's server-minted handle, \"cmp_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudGetV1CampaignId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudGetV1CampaignId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CampaignId`: CloudCampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudGetV1CampaignId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign&#39;s server-minted handle, \&quot;cmp_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CampaignIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCampaignRecord**](CloudCampaignRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CampaignIdMetrics

> CloudCampaignResults CloudGetV1CampaignIdMetrics(ctx, id).Range_(range_).Start(start).End(end).Execute()

CampaignMetrics returns a campaign's results over a window: the analytics funnel (impressions, clicks, conversions, revenue, visitors), the spend each channel's connector reports, and the derived growth KPIs — CTR, CVR, CAC and ROAS.



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
	id := "cmp_1f…" // string | ID is the campaign to report on, from the path.
	range_ := "7d" // string | Range is the lookback window: 24h, 7d, 30d or 90d. Anything else, including empty, reads as 30d. (optional)
	start := "start_example" // string | Start is an explicit RFC3339 window start. Honored only together with End, and only when End is after it. (optional)
	end := "end_example" // string | End is an explicit RFC3339 window end. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudGetV1CampaignIdMetrics(context.Background(), id).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudGetV1CampaignIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CampaignIdMetrics`: CloudCampaignResults
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudGetV1CampaignIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to report on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CampaignIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** | Range is the lookback window: 24h, 7d, 30d or 90d. Anything else, including empty, reads as 30d. | 
 **start** | **string** | Start is an explicit RFC3339 window start. Honored only together with End, and only when End is after it. | 
 **end** | **string** | End is an explicit RFC3339 window end. | 

### Return type

[**CloudCampaignResults**](CloudCampaignResults.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CampaignSummary

> CloudCampaignSummary CloudGetV1CampaignSummary(ctx).Execute()

SummarizeCampaigns returns the org's go-to-market roll-up: how many campaigns exist, how many are live, their total budget in cents, and which channel executors this deployment can actually reach.



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
	resp, r, err := apiClient.CampaignAPI.CloudGetV1CampaignSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudGetV1CampaignSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CampaignSummary`: CloudCampaignSummary
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudGetV1CampaignSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CampaignSummaryRequest struct via the builder pattern


### Return type

[**CloudCampaignSummary**](CloudCampaignSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Campaign

> CloudCampaignRecord CloudPostV1Campaign(ctx).CloudCampaignWrite(cloudCampaignWrite).Execute()

CreateCampaign creates a campaign as a DRAFT and returns it.



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
	cloudCampaignWrite := *openapiclient.NewCloudCampaignWrite() // CloudCampaignWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudPostV1Campaign(context.Background()).CloudCampaignWrite(cloudCampaignWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudPostV1Campaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Campaign`: CloudCampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudPostV1Campaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCampaignWrite** | [**CloudCampaignWrite**](CloudCampaignWrite.md) |  | 

### Return type

[**CloudCampaignRecord**](CloudCampaignRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CampaignByIdLaunch

> CloudPostV1CampaignByIdLaunch(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignAPI.CloudPostV1CampaignByIdLaunch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudPostV1CampaignByIdLaunch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CampaignByIdLaunchRequest struct via the builder pattern


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


## CloudPostV1CampaignByIdPause

> CloudPostV1CampaignByIdPause(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignAPI.CloudPostV1CampaignByIdPause(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudPostV1CampaignByIdPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CampaignByIdPauseRequest struct via the builder pattern


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


## CloudPostV1CampaignIdChannels

> CloudCampaignRecord CloudPostV1CampaignIdChannels(ctx, id).CloudChannelAdd(cloudChannelAdd).Execute()

AddCampaignChannel adds a channel to a campaign, or REPLACES the one it already has of that kind, and returns the updated campaign.



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
	id := "cmp_1f…" // string | ID is the campaign to add the channel to, from the path.
	cloudChannelAdd := *openapiclient.NewCloudChannelAdd() // CloudChannelAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudPostV1CampaignIdChannels(context.Background(), id).CloudChannelAdd(cloudChannelAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudPostV1CampaignIdChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CampaignIdChannels`: CloudCampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudPostV1CampaignIdChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to add the channel to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CampaignIdChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudChannelAdd** | [**CloudChannelAdd**](CloudChannelAdd.md) |  | 

### Return type

[**CloudCampaignRecord**](CloudCampaignRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CampaignId

> CloudCampaignRecord CloudPutV1CampaignId(ctx, id).CloudCampaignUpdate(cloudCampaignUpdate).Execute()

UpdateCampaign rewrites a campaign's core fields — name, audience, creatives, schedule and budget — and returns the updated campaign.



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
	id := "id_example" // string | ID is the campaign to update, from the path.
	cloudCampaignUpdate := *openapiclient.NewCloudCampaignUpdate() // CloudCampaignUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.CloudPutV1CampaignId(context.Background(), id).CloudCampaignUpdate(cloudCampaignUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CloudPutV1CampaignId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CampaignId`: CloudCampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.CloudPutV1CampaignId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CampaignIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCampaignUpdate** | [**CloudCampaignUpdate**](CloudCampaignUpdate.md) |  | 

### Return type

[**CloudCampaignRecord**](CloudCampaignRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

