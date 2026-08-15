# \CampaignAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCampaignById**](CampaignAPI.md#DeleteCampaignById) | **Delete** /v1/campaign/{id} | Removes one campaign of the caller&#39;s org and answers 204 with no body.
[**DeleteCampaignByIdChannelsByKind**](CampaignAPI.md#DeleteCampaignByIdChannelsByKind) | **Delete** /v1/campaign/{id}/channels/{kind} | Drops one channel from a campaign and returns the updated campaign.
[**GetCampaign**](CampaignAPI.md#GetCampaign) | **Get** /v1/campaign | Returns the org&#39;s campaigns, newest first, optionally narrowed to one status.
[**GetCampaignById**](CampaignAPI.md#GetCampaignById) | **Get** /v1/campaign/{id} | Returns one campaign of the caller&#39;s org — its name, audience, creatives, channels with their per-channel launch state, schedule, budget and status.
[**GetCampaignByIdMetrics**](CampaignAPI.md#GetCampaignByIdMetrics) | **Get** /v1/campaign/{id}/metrics | Returns a campaign&#39;s results over a window: the analytics funnel (impressions, clicks, conversions, revenue, visitors), the spend each channel&#39;s connector reports, and the derived growth KPIs — CTR, CVR, CAC and ROAS.
[**GetCampaignSummary**](CampaignAPI.md#GetCampaignSummary) | **Get** /v1/campaign/summary | Returns the org&#39;s go-to-market roll-up: how many campaigns exist, how many are live, their total budget in cents, and which channel executors this deployment can actually reach.
[**PostCampaign**](CampaignAPI.md#PostCampaign) | **Post** /v1/campaign | Creates a campaign as a DRAFT and returns it.
[**PostCampaignByIdChannels**](CampaignAPI.md#PostCampaignByIdChannels) | **Post** /v1/campaign/{id}/channels | Adds a channel to a campaign, or REPLACES the one it already has of that kind, and returns the updated campaign.
[**PostCampaignByIdLaunch**](CampaignAPI.md#PostCampaignByIdLaunch) | **Post** /v1/campaign/{id}/launch | Launch a campaign across every channel it declares
[**PostCampaignByIdPause**](CampaignAPI.md#PostCampaignByIdPause) | **Post** /v1/campaign/{id}/pause | Pause every live channel on a campaign at its provider
[**PutCampaignById**](CampaignAPI.md#PutCampaignById) | **Put** /v1/campaign/{id} | Rewrites a campaign&#39;s core fields — name, audience, creatives, schedule and budget — and returns the updated campaign.



## DeleteCampaignById

> DeleteCampaignById(ctx, id).Execute()

Removes one campaign of the caller's org and answers 204 with no body.



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
	r, err := apiClient.CampaignAPI.DeleteCampaignById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.DeleteCampaignById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCampaignByIdRequest struct via the builder pattern


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


## DeleteCampaignByIdChannelsByKind

> CampaignRecord DeleteCampaignByIdChannelsByKind(ctx, id, kind).Execute()

Drops one channel from a campaign and returns the updated campaign.



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
	resp, r, err := apiClient.CampaignAPI.DeleteCampaignByIdChannelsByKind(context.Background(), id, kind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.DeleteCampaignByIdChannelsByKind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCampaignByIdChannelsByKind`: CampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.DeleteCampaignByIdChannelsByKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign, from the path. | 
**kind** | **string** | Kind is the channel to remove: paid, organic or email. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignByIdChannelsByKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CampaignRecord**](CampaignRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> CampaignPage GetCampaign(ctx).Status(status).Limit(limit).Execute()

Returns the org's campaigns, newest first, optionally narrowed to one status.



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
	resp, r, err := apiClient.CampaignAPI.GetCampaign(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: CampaignPage
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only campaigns in that state: draft, live, paused or failed. Empty means any. | 
 **limit** | **int32** | Limit bounds the page. 0 or less means the default of 200; anything above 1000 is clamped to 1000. | 

### Return type

[**CampaignPage**](CampaignPage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignById

> CampaignRecord GetCampaignById(ctx, id).Execute()

Returns one campaign of the caller's org — its name, audience, creatives, channels with their per-channel launch state, schedule, budget and status.



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
	resp, r, err := apiClient.CampaignAPI.GetCampaignById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignById`: CampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign&#39;s server-minted handle, \&quot;cmp_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignRecord**](CampaignRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByIdMetrics

> CampaignResults GetCampaignByIdMetrics(ctx, id).Range_(range_).Start(start).End(end).Execute()

Returns a campaign's results over a window: the analytics funnel (impressions, clicks, conversions, revenue, visitors), the spend each channel's connector reports, and the derived growth KPIs — CTR, CVR, CAC and ROAS.



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
	resp, r, err := apiClient.CampaignAPI.GetCampaignByIdMetrics(context.Background(), id).Range_(range_).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignByIdMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignByIdMetrics`: CampaignResults
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignByIdMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to report on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByIdMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** | Range is the lookback window: 24h, 7d, 30d or 90d. Anything else, including empty, reads as 30d. | 
 **start** | **string** | Start is an explicit RFC3339 window start. Honored only together with End, and only when End is after it. | 
 **end** | **string** | End is an explicit RFC3339 window end. | 

### Return type

[**CampaignResults**](CampaignResults.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSummary

> CampaignSummary GetCampaignSummary(ctx).Execute()

Returns the org's go-to-market roll-up: how many campaigns exist, how many are live, their total budget in cents, and which channel executors this deployment can actually reach.



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
	resp, r, err := apiClient.CampaignAPI.GetCampaignSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignSummary`: CampaignSummary
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSummaryRequest struct via the builder pattern


### Return type

[**CampaignSummary**](CampaignSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaign

> CampaignRecord PostCampaign(ctx).CampaignWrite(campaignWrite).Execute()

Creates a campaign as a DRAFT and returns it.



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
	campaignWrite := *openapiclient.NewCampaignWrite() // CampaignWrite | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.PostCampaign(context.Background()).CampaignWrite(campaignWrite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PostCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaign`: CampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.PostCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignWrite** | [**CampaignWrite**](CampaignWrite.md) |  | 

### Return type

[**CampaignRecord**](CampaignRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignByIdChannels

> CampaignRecord PostCampaignByIdChannels(ctx, id).ChannelAdd(channelAdd).Execute()

Adds a channel to a campaign, or REPLACES the one it already has of that kind, and returns the updated campaign.



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
	channelAdd := *openapiclient.NewChannelAdd() // ChannelAdd | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.PostCampaignByIdChannels(context.Background(), id).ChannelAdd(channelAdd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PostCampaignByIdChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaignByIdChannels`: CampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.PostCampaignByIdChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to add the channel to, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignByIdChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelAdd** | [**ChannelAdd**](ChannelAdd.md) |  | 

### Return type

[**CampaignRecord**](CampaignRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaignByIdLaunch

> PostCampaignByIdLaunch(ctx, id).Execute()

Launch a campaign across every channel it declares



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
	r, err := apiClient.CampaignAPI.PostCampaignByIdLaunch(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PostCampaignByIdLaunch``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCampaignByIdLaunchRequest struct via the builder pattern


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


## PostCampaignByIdPause

> PostCampaignByIdPause(ctx, id).Execute()

Pause every live channel on a campaign at its provider



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
	r, err := apiClient.CampaignAPI.PostCampaignByIdPause(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PostCampaignByIdPause``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCampaignByIdPauseRequest struct via the builder pattern


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


## PutCampaignById

> CampaignRecord PutCampaignById(ctx, id).CampaignUpdate(campaignUpdate).Execute()

Rewrites a campaign's core fields — name, audience, creatives, schedule and budget — and returns the updated campaign.



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
	campaignUpdate := *openapiclient.NewCampaignUpdate() // CampaignUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.PutCampaignById(context.Background(), id).CampaignUpdate(campaignUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PutCampaignById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCampaignById`: CampaignRecord
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.PutCampaignById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign to update, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCampaignByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignUpdate** | [**CampaignUpdate**](CampaignUpdate.md) |  | 

### Return type

[**CampaignRecord**](CampaignRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

