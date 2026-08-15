# \MarketingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingAudiencesById**](MarketingAPI.md#DeleteMarketingAudiencesById) | **Delete** /v1/marketing/audiences/{id} | Removes one of the caller org&#39;s audiences and answers 204.
[**DeleteMarketingCalendarById**](MarketingAPI.md#DeleteMarketingCalendarById) | **Delete** /v1/marketing/calendar/{id} | Removes one of the caller org&#39;s posts and answers 204.
[**DeleteMarketingCampaignsById**](MarketingAPI.md#DeleteMarketingCampaignsById) | **Delete** /v1/marketing/campaigns/{id} | Removes one of the caller org&#39;s campaigns and answers 204.
[**DeleteMarketingSuppressions**](MarketingAPI.md#DeleteMarketingSuppressions) | **Delete** /v1/marketing/suppressions | Re-subscribes an address on one channel and answers 204.
[**GetMarketingAudiences**](MarketingAPI.md#GetMarketingAudiences) | **Get** /v1/marketing/audiences | Returns the org&#39;s saved audiences, most recently updated first.
[**GetMarketingAudiencesById**](MarketingAPI.md#GetMarketingAudiencesById) | **Get** /v1/marketing/audiences/{id} | Returns one of the caller org&#39;s saved audiences.
[**GetMarketingAudiencesByIdPreview**](MarketingAPI.md#GetMarketingAudiencesByIdPreview) | **Get** /v1/marketing/audiences/{id}/preview | Evaluates the cohort LIVE — the same resolution an enrollment would run — and reports how big it is and how many real mailboxes it reaches.
[**GetMarketingCalendar**](MarketingAPI.md#GetMarketingCalendar) | **Get** /v1/marketing/calendar | Returns the org&#39;s calendar, soonest scheduled first, optionally narrowed to one status.
[**GetMarketingCalendarById**](MarketingAPI.md#GetMarketingCalendarById) | **Get** /v1/marketing/calendar/{id} | Returns one of the caller org&#39;s posts, including the exact error behind a failed publish.
[**GetMarketingCampaigns**](MarketingAPI.md#GetMarketingCampaigns) | **Get** /v1/marketing/campaigns | Returns the org&#39;s campaigns, most recently updated first, optionally narrowed to one lifecycle status.
[**GetMarketingCampaignsById**](MarketingAPI.md#GetMarketingCampaignsById) | **Get** /v1/marketing/campaigns/{id} | Returns one of the caller org&#39;s campaigns.
[**GetMarketingPromos**](MarketingAPI.md#GetMarketingPromos) | **Get** /v1/marketing/promos | Returns every promo the deployment offers with its live counters: how many orgs have redeemed it and how many redemptions remain under the cap.
[**GetMarketingPromosByCodeEligibility**](MarketingAPI.md#GetMarketingPromosByCodeEligibility) | **Get** /v1/marketing/promos/{code}/eligibility | Prices a promo against a plan and seat count.
[**GetMarketingPromosByCodeRedemption**](MarketingAPI.md#GetMarketingPromosByCodeRedemption) | **Get** /v1/marketing/promos/{code}/redemption | Returns the caller org&#39;s OWN redemption of a promo — an org-scoped read, so it can never surface another tenant&#39;s.
[**GetMarketingSequences**](MarketingAPI.md#GetMarketingSequences) | **Get** /v1/marketing/sequences | Returns the org&#39;s drip sequences, most recently updated first.
[**GetMarketingSequencesById**](MarketingAPI.md#GetMarketingSequencesById) | **Get** /v1/marketing/sequences/{id} | Returns one of the caller org&#39;s sequences together with its steps in send order.
[**GetMarketingSequencesByIdEnrollments**](MarketingAPI.md#GetMarketingSequencesByIdEnrollments) | **Get** /v1/marketing/sequences/{id}/enrollments | Returns who is walking one sequence, most recently enrolled first, with each walk&#39;s current step and next due time.
[**GetMarketingSequencesByIdSteps**](MarketingAPI.md#GetMarketingSequencesByIdSteps) | **Get** /v1/marketing/sequences/{id}/steps | Returns one sequence&#39;s steps in send order.
[**GetMarketingSummary**](MarketingAPI.md#GetMarketingSummary) | **Get** /v1/marketing/summary | Rolls up the caller org&#39;s campaigns: how many there are, how many are active, and the summed budget and spend in cents.
[**GetMarketingSuppressions**](MarketingAPI.md#GetMarketingSuppressions) | **Get** /v1/marketing/suppressions | Returns the org&#39;s opt-out list, newest first — everyone the send gate will refuse to deliver to.
[**GetMarketingUnsubscribe**](MarketingAPI.md#GetMarketingUnsubscribe) | **Get** /v1/marketing/unsubscribe | Is the PUBLIC one-click endpoint (no principal): a recipient clicks the signed link in an email footer.
[**PostMarketingAudiences**](MarketingAPI.md#PostMarketingAudiences) | **Post** /v1/marketing/audiences | Saves a cohort filter for the caller&#39;s org.
[**PostMarketingCalendar**](MarketingAPI.md#PostMarketingCalendar) | **Post** /v1/marketing/calendar | Adds a post to the content calendar.
[**PostMarketingCalendarByIdPublish**](MarketingAPI.md#PostMarketingCalendarByIdPublish) | **Post** /v1/marketing/calendar/{id}/publish | Publishes a post NOW, synchronously, whatever its schedule.
[**PostMarketingCampaigns**](MarketingAPI.md#PostMarketingCampaigns) | **Post** /v1/marketing/campaigns | Registers a campaign in the caller&#39;s org.
[**PostMarketingCampaignsByIdSchedule**](MarketingAPI.md#PostMarketingCampaignsByIdSchedule) | **Post** /v1/marketing/campaigns/{id}/schedule | Sets a campaign&#39;s send time and moves it to \&quot;scheduled\&quot;.
[**PostMarketingPromosByCodeRedeem**](MarketingAPI.md#PostMarketingPromosByCodeRedeem) | **Post** /v1/marketing/promos/{code}/redeem | Records the caller org&#39;s claim on a promo.
[**PostMarketingSequences**](MarketingAPI.md#PostMarketingSequences) | **Post** /v1/marketing/sequences | Registers a drip sequence in the caller&#39;s org.
[**PostMarketingSequencesByIdEnroll**](MarketingAPI.md#PostMarketingSequencesByIdEnroll) | **Post** /v1/marketing/sequences/{id}/enroll | Adds one contact or a whole audience to a sequence and schedules the first step for each.
[**PostMarketingSequencesByIdEnrollmentsByEidCancel**](MarketingAPI.md#PostMarketingSequencesByIdEnrollmentsByEidCancel) | **Post** /v1/marketing/sequences/{id}/enrollments/{eid}/cancel | Stops one walk mid-sequence and answers 204: no further step is sent, and steps already delivered are not recalled.
[**PostMarketingSequencesByIdStatus**](MarketingAPI.md#PostMarketingSequencesByIdStatus) | **Post** /v1/marketing/sequences/{id}/status | Flips draft/active/archived — the activation gate for sending, since only an active sequence accepts enrollments.
[**PostMarketingSequencesByIdSteps**](MarketingAPI.md#PostMarketingSequencesByIdSteps) | **Post** /v1/marketing/sequences/{id}/steps | Appends a message to the END of a sequence: the new step&#39;s idx is one past the last, so steps arrive in the order they are added.
[**PostMarketingSuppressions**](MarketingAPI.md#PostMarketingSuppressions) | **Post** /v1/marketing/suppressions | Records an opt-out for the org (admin / self-service management).
[**PutMarketingCalendarById**](MarketingAPI.md#PutMarketingCalendarById) | **Put** /v1/marketing/calendar/{id} | Replaces a post&#39;s editable fields.
[**PutMarketingCampaignsById**](MarketingAPI.md#PutMarketingCampaignsById) | **Put** /v1/marketing/campaigns/{id} | Replaces a campaign&#39;s editable fields.



## DeleteMarketingAudiencesById

> DeleteMarketingAudiencesById(ctx, id).Execute()

Removes one of the caller org's audiences and answers 204.



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
	id := "aud_4c1e9b7a2d6f0538e4a7c9b1d3f5027a" // string | ID is the audience id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingAPI.DeleteMarketingAudiencesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.DeleteMarketingAudiencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the audience id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingAudiencesByIdRequest struct via the builder pattern


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


## DeleteMarketingCalendarById

> DeleteMarketingCalendarById(ctx, id).Execute()

Removes one of the caller org's posts and answers 204.



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
	id := "cal_1d7f3b9e5a2c8046f1b3d5a7c9e02468" // string | ID is the post id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingAPI.DeleteMarketingCalendarById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.DeleteMarketingCalendarById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the post id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCalendarByIdRequest struct via the builder pattern


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


## DeleteMarketingCampaignsById

> DeleteMarketingCampaignsById(ctx, id).Execute()

Removes one of the caller org's campaigns and answers 204.



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
	id := "camp_9f2a1c7d4e8b0a6f3d2c5b1e7a9f4c60" // string | ID is the campaign id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingAPI.DeleteMarketingCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.DeleteMarketingCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingCampaignsByIdRequest struct via the builder pattern


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


## DeleteMarketingSuppressions

> DeleteMarketingSuppressions(ctx).Channel(channel).Address(address).Reason(reason).CreatedAt(createdAt).Execute()

Re-subscribes an address on one channel and answers 204.



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
	channel := "email" // string | Channel is the surface opted out of: email, sms, social, meta, google or tiktok. Empty means email. Opting out of one leaves the others reachable. (optional)
	address := "person@example.com" // string | Address is the recipient, normalized (lower-cased, trimmed) so an opt-out cannot be slipped past on a case or whitespace difference. Required. (optional)
	reason := "reason_example" // string | Reason is a free-text note, capped at 1024 bytes. The public one-click endpoint records \"one-click unsubscribe\". (optional)
	createdAt := int32(56) // int32 | CreatedAt is unix seconds, server-assigned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingAPI.DeleteMarketingSuppressions(context.Background()).Channel(channel).Address(address).Reason(reason).CreatedAt(createdAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.DeleteMarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the surface opted out of: email, sms, social, meta, google or tiktok. Empty means email. Opting out of one leaves the others reachable. | 
 **address** | **string** | Address is the recipient, normalized (lower-cased, trimmed) so an opt-out cannot be slipped past on a case or whitespace difference. Required. | 
 **reason** | **string** | Reason is a free-text note, capped at 1024 bytes. The public one-click endpoint records \&quot;one-click unsubscribe\&quot;. | 
 **createdAt** | **int32** | CreatedAt is unix seconds, server-assigned. | 

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


## GetMarketingAudiences

> AudienceList GetMarketingAudiences(ctx).Limit(limit).Execute()

Returns the org's saved audiences, most recently updated first.



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
	limit := int32(50) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingAudiences(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingAudiences`: AudienceList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**AudienceList**](AudienceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingAudiencesById

> Audience GetMarketingAudiencesById(ctx, id).Execute()

Returns one of the caller org's saved audiences.



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
	id := "aud_4c1e9b7a2d6f0538e4a7c9b1d3f5027a" // string | ID is the audience id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingAudiencesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingAudiencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingAudiencesById`: Audience
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingAudiencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the audience id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingAudiencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Audience**](Audience.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingAudiencesByIdPreview

> AudiencePreview GetMarketingAudiencesByIdPreview(ctx, id).Execute()

Evaluates the cohort LIVE — the same resolution an enrollment would run — and reports how big it is and how many real mailboxes it reaches.



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
	id := "aud_4c1e9b7a2d6f0538e4a7c9b1d3f5027a" // string | ID is the audience id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingAudiencesByIdPreview(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingAudiencesByIdPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingAudiencesByIdPreview`: AudiencePreview
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingAudiencesByIdPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the audience id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingAudiencesByIdPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AudiencePreview**](AudiencePreview.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCalendar

> PostList GetMarketingCalendar(ctx).Status(status).Limit(limit).Execute()

Returns the org's calendar, soonest scheduled first, optionally narrowed to one status.



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
	status := "scheduled" // string | Status keeps only posts in that state (draft, scheduled, published, failed, canceled). Empty means every post. (optional)
	limit := int32(50) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingCalendar(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCalendar`: PostList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only posts in that state (draft, scheduled, published, failed, canceled). Empty means every post. | 
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**PostList**](PostList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCalendarById

> CalendarPost GetMarketingCalendarById(ctx, id).Execute()

Returns one of the caller org's posts, including the exact error behind a failed publish.



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
	id := "cal_1d7f3b9e5a2c8046f1b3d5a7c9e02468" // string | ID is the post id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingCalendarById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingCalendarById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCalendarById`: CalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingCalendarById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the post id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCalendarByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalendarPost**](CalendarPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaigns

> CampaignList GetMarketingCampaigns(ctx).Status(status).Limit(limit).Execute()

Returns the org's campaigns, most recently updated first, optionally narrowed to one lifecycle status.



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
	status := "active" // string | Status keeps only campaigns in that lifecycle state (draft, scheduled, active, paused, completed). Empty means every campaign. (optional)
	limit := int32(25) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingCampaigns(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaigns`: CampaignList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only campaigns in that lifecycle state (draft, scheduled, active, paused, completed). Empty means every campaign. | 
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CampaignList**](CampaignList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingCampaignsById

> Campaign GetMarketingCampaignsById(ctx, id).Execute()

Returns one of the caller org's campaigns.



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
	id := "camp_9f2a1c7d4e8b0a6f3d2c5b1e7a9f4c60" // string | ID is the campaign id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingCampaignsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingCampaignsById`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingPromos

> PromoList GetMarketingPromos(ctx).Execute()

Returns every promo the deployment offers with its live counters: how many orgs have redeemed it and how many redemptions remain under the cap.



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
	resp, r, err := apiClient.MarketingAPI.GetMarketingPromos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingPromos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingPromos`: PromoList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingPromos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingPromosRequest struct via the builder pattern


### Return type

[**PromoList**](PromoList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingPromosByCodeEligibility

> Quote GetMarketingPromosByCodeEligibility(ctx, code).Plan(plan).Seats(seats).Execute()

Prices a promo against a plan and seat count.



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
	code := "first1000" // string | Code is the promo code from the path.
	plan := "team" // string | Plan is the plan being priced: pro, max or team. Anything else (including the free Developer plan) has no list price and so nothing to discount. (optional)
	seats := int32(12) // int32 | Seats is the Team seat count; 0 means 1, and it is ignored for the single-seat plans. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingPromosByCodeEligibility(context.Background(), code).Plan(plan).Seats(seats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingPromosByCodeEligibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingPromosByCodeEligibility`: Quote
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingPromosByCodeEligibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingPromosByCodeEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plan** | **string** | Plan is the plan being priced: pro, max or team. Anything else (including the free Developer plan) has no list price and so nothing to discount. | 
 **seats** | **int32** | Seats is the Team seat count; 0 means 1, and it is ignored for the single-seat plans. | 

### Return type

[**Quote**](Quote.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingPromosByCodeRedemption

> Redemption GetMarketingPromosByCodeRedemption(ctx, code).Execute()

Returns the caller org's OWN redemption of a promo — an org-scoped read, so it can never surface another tenant's.



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
	code := "first1000" // string | Code is the promo code from the path, e.g. \"first1000\".

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingPromosByCodeRedemption(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingPromosByCodeRedemption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingPromosByCodeRedemption`: Redemption
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingPromosByCodeRedemption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path, e.g. \&quot;first1000\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingPromosByCodeRedemptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Redemption**](Redemption.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSequences

> SequenceList GetMarketingSequences(ctx).Limit(limit).Execute()

Returns the org's drip sequences, most recently updated first.



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
	limit := int32(50) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingSequences(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSequences`: SequenceList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**SequenceList**](SequenceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSequencesById

> SequenceView GetMarketingSequencesById(ctx, id).Execute()

Returns one of the caller org's sequences together with its steps in send order.



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
	id := "seq_7b3e5a1c9d024f68b0a3e7c5d9f1a248" // string | ID is the sequence id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingSequencesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSequencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSequencesById`: SequenceView
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSequencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSequencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SequenceView**](SequenceView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSequencesByIdEnrollments

> EnrollmentList GetMarketingSequencesByIdEnrollments(ctx, id).Limit(limit).Execute()

Returns who is walking one sequence, most recently enrolled first, with each walk's current step and next due time.



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
	id := "seq_7b3e5a1c9d024f68b0a3e7c5d9f1a248" // string | ID is the sequence id from the path.
	limit := int32(100) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingSequencesByIdEnrollments(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSequencesByIdEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSequencesByIdEnrollments`: EnrollmentList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSequencesByIdEnrollments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSequencesByIdEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**EnrollmentList**](EnrollmentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSequencesByIdSteps

> StepList GetMarketingSequencesByIdSteps(ctx, id).Execute()

Returns one sequence's steps in send order.



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
	id := "seq_7b3e5a1c9d024f68b0a3e7c5d9f1a248" // string | ID is the sequence id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingSequencesByIdSteps(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSequencesByIdSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSequencesByIdSteps`: StepList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSequencesByIdSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSequencesByIdStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StepList**](StepList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSummary

> Summary GetMarketingSummary(ctx).Execute()

Rolls up the caller org's campaigns: how many there are, how many are active, and the summed budget and spend in cents.



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
	resp, r, err := apiClient.MarketingAPI.GetMarketingSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSummary`: Summary
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSummaryRequest struct via the builder pattern


### Return type

[**Summary**](Summary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingSuppressions

> SuppressionList GetMarketingSuppressions(ctx).Limit(limit).Execute()

Returns the org's opt-out list, newest first — everyone the send gate will refuse to deliver to.



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
	limit := int32(100) // int32 | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingSuppressions(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingSuppressions`: SuppressionList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**SuppressionList**](SuppressionList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingUnsubscribe

> Unsubscribed GetMarketingUnsubscribe(ctx).Org(org).Channel(channel).Address(address).Token(token).Execute()

Is the PUBLIC one-click endpoint (no principal): a recipient clicks the signed link in an email footer.



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
	org := "acme" // string | Org is the org the link was minted for. (optional)
	channel := "email" // string | Channel is the surface to opt out of. (optional)
	address := "person@example.com" // string | Address is the recipient to opt out. (optional)
	token := "9f2a…" // string | Token is the HMAC over (org, channel, address). It is the ONLY authority here — there is no principal — so it binds the request to one tuple and nothing else. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.GetMarketingUnsubscribe(context.Background()).Org(org).Channel(channel).Address(address).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.GetMarketingUnsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingUnsubscribe`: Unsubscribed
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.GetMarketingUnsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org is the org the link was minted for. | 
 **channel** | **string** | Channel is the surface to opt out of. | 
 **address** | **string** | Address is the recipient to opt out. | 
 **token** | **string** | Token is the HMAC over (org, channel, address). It is the ONLY authority here — there is no principal — so it binds the request to one tuple and nothing else. | 

### Return type

[**Unsubscribed**](Unsubscribed.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingAudiences

> Audience PostMarketingAudiences(ctx).Audience(audience).Execute()

Saves a cohort filter for the caller's org.



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
	audience := *openapiclient.NewAudience() // Audience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingAudiences(context.Background()).Audience(audience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingAudiences`: Audience
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audience** | [**Audience**](Audience.md) |  | 

### Return type

[**Audience**](Audience.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCalendar

> CalendarPost PostMarketingCalendar(ctx).CalendarPost(calendarPost).Execute()

Adds a post to the content calendar.



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
	calendarPost := *openapiclient.NewCalendarPost() // CalendarPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingCalendar(context.Background()).CalendarPost(calendarPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCalendar`: CalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calendarPost** | [**CalendarPost**](CalendarPost.md) |  | 

### Return type

[**CalendarPost**](CalendarPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCalendarByIdPublish

> CalendarPost PostMarketingCalendarByIdPublish(ctx, id).Execute()

Publishes a post NOW, synchronously, whatever its schedule.



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
	id := "cal_1d7f3b9e5a2c8046f1b3d5a7c9e02468" // string | ID is the post id from the path, as returned by create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingCalendarByIdPublish(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingCalendarByIdPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCalendarByIdPublish`: CalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingCalendarByIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the post id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCalendarByIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalendarPost**](CalendarPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaigns

> Campaign PostMarketingCampaigns(ctx).Campaign(campaign).Execute()

Registers a campaign in the caller's org.



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
	campaign := *openapiclient.NewCampaign() // Campaign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingCampaigns(context.Background()).Campaign(campaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaigns`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaign** | [**Campaign**](Campaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingCampaignsByIdSchedule

> Campaign PostMarketingCampaignsByIdSchedule(ctx, id).ScheduleInput(scheduleInput).Execute()

Sets a campaign's send time and moves it to \"scheduled\".



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
	id := "camp_9f2a1c7d4e8b0a6f3d2c5b1e7a9f4c60" // string | ID is the campaign id from the path.
	scheduleInput := *openapiclient.NewScheduleInput() // ScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingCampaignsByIdSchedule(context.Background(), id).ScheduleInput(scheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingCampaignsByIdSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingCampaignsByIdSchedule`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingCampaignsByIdSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingCampaignsByIdScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scheduleInput** | [**ScheduleInput**](ScheduleInput.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingPromosByCodeRedeem

> RedeemResult PostMarketingPromosByCodeRedeem(ctx, code).RedeemInput(redeemInput).Execute()

Records the caller org's claim on a promo.



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
	code := "first1000" // string | Code is the promo code from the path.
	redeemInput := *openapiclient.NewRedeemInput() // RedeemInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingPromosByCodeRedeem(context.Background(), code).RedeemInput(redeemInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingPromosByCodeRedeem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingPromosByCodeRedeem`: RedeemResult
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingPromosByCodeRedeem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingPromosByCodeRedeemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redeemInput** | [**RedeemInput**](RedeemInput.md) |  | 

### Return type

[**RedeemResult**](RedeemResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingSequences

> Sequence PostMarketingSequences(ctx).Sequence(sequence).Execute()

Registers a drip sequence in the caller's org.



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
	sequence := *openapiclient.NewSequence() // Sequence | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingSequences(context.Background()).Sequence(sequence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingSequences`: Sequence
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sequence** | [**Sequence**](Sequence.md) |  | 

### Return type

[**Sequence**](Sequence.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingSequencesByIdEnroll

> EnrollResult PostMarketingSequencesByIdEnroll(ctx, id).EnrollInput(enrollInput).Execute()

Adds one contact or a whole audience to a sequence and schedules the first step for each.



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
	id := "id_example" // string | ID is the sequence id from the path.
	enrollInput := *openapiclient.NewEnrollInput() // EnrollInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingSequencesByIdEnroll(context.Background(), id).EnrollInput(enrollInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSequencesByIdEnroll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingSequencesByIdEnroll`: EnrollResult
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingSequencesByIdEnroll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSequencesByIdEnrollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enrollInput** | [**EnrollInput**](EnrollInput.md) |  | 

### Return type

[**EnrollResult**](EnrollResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingSequencesByIdEnrollmentsByEidCancel

> PostMarketingSequencesByIdEnrollmentsByEidCancel(ctx, id, eid).Execute()

Stops one walk mid-sequence and answers 204: no further step is sent, and steps already delivered are not recalled.



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
	id := "seq_7b3e5a1c9d024f68b0a3e7c5d9f1a248" // string | ID is the sequence id from the path.
	eid := "enr_2a8d6f0b4c1e9375a0d2f6b8c4e19f73" // string | EID is the enrollment id from the path, as returned by a single-address enroll.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MarketingAPI.PostMarketingSequencesByIdEnrollmentsByEidCancel(context.Background(), id, eid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSequencesByIdEnrollmentsByEidCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 
**eid** | **string** | EID is the enrollment id from the path, as returned by a single-address enroll. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSequencesByIdEnrollmentsByEidCancelRequest struct via the builder pattern


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


## PostMarketingSequencesByIdStatus

> SequenceStatus PostMarketingSequencesByIdStatus(ctx, id).SequenceStatus(sequenceStatus).Execute()

Flips draft/active/archived — the activation gate for sending, since only an active sequence accepts enrollments.



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
	id := "seq_7b3e5a1c9d024f68b0a3e7c5d9f1a248" // string | ID is the sequence id from the path.
	sequenceStatus := *openapiclient.NewSequenceStatus() // SequenceStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingSequencesByIdStatus(context.Background(), id).SequenceStatus(sequenceStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSequencesByIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingSequencesByIdStatus`: SequenceStatus
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingSequencesByIdStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSequencesByIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sequenceStatus** | [**SequenceStatus**](SequenceStatus.md) |  | 

### Return type

[**SequenceStatus**](SequenceStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingSequencesByIdSteps

> Step PostMarketingSequencesByIdSteps(ctx, id).StepInput(stepInput).Execute()

Appends a message to the END of a sequence: the new step's idx is one past the last, so steps arrive in the order they are added.



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
	id := "id_example" // string | SequenceID is the sequence id from the path (the route's :id).
	stepInput := *openapiclient.NewStepInput() // StepInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingSequencesByIdSteps(context.Background(), id).StepInput(stepInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSequencesByIdSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingSequencesByIdSteps`: Step
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingSequencesByIdSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SequenceID is the sequence id from the path (the route&#39;s :id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSequencesByIdStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stepInput** | [**StepInput**](StepInput.md) |  | 

### Return type

[**Step**](Step.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingSuppressions

> Suppression PostMarketingSuppressions(ctx).Suppression(suppression).Execute()

Records an opt-out for the org (admin / self-service management).



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
	suppression := *openapiclient.NewSuppression() // Suppression | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PostMarketingSuppressions(context.Background()).Suppression(suppression).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PostMarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingSuppressions`: Suppression
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PostMarketingSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suppression** | [**Suppression**](Suppression.md) |  | 

### Return type

[**Suppression**](Suppression.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCalendarById

> CalendarPost PutMarketingCalendarById(ctx, id).CalendarPost(calendarPost).Execute()

Replaces a post's editable fields.



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
	id := "id_example" // string | ID is the server-assigned post id (\"cal_\" + 128 random bits).
	calendarPost := *openapiclient.NewCalendarPost() // CalendarPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PutMarketingCalendarById(context.Background(), id).CalendarPost(calendarPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PutMarketingCalendarById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCalendarById`: CalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PutMarketingCalendarById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server-assigned post id (\&quot;cal_\&quot; + 128 random bits). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCalendarByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calendarPost** | [**CalendarPost**](CalendarPost.md) |  | 

### Return type

[**CalendarPost**](CalendarPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingCampaignsById

> Campaign PutMarketingCampaignsById(ctx, id).Campaign(campaign).Execute()

Replaces a campaign's editable fields.



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
	id := "id_example" // string | ID is the server-assigned campaign id (\"camp_\" + 128 random bits).
	campaign := *openapiclient.NewCampaign() // Campaign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.PutMarketingCampaignsById(context.Background(), id).Campaign(campaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.PutMarketingCampaignsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingCampaignsById`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.PutMarketingCampaignsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server-assigned campaign id (\&quot;camp_\&quot; + 128 random bits). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingCampaignsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaign** | [**Campaign**](Campaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

