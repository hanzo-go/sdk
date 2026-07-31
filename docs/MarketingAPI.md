# \MarketingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1MarketingAudiencesId**](MarketingAPI.md#CloudDeleteV1MarketingAudiencesId) | **Delete** /v1/marketing/audiences/{id} | Removes one of the caller org&#39;s audiences and answers 204.
[**CloudDeleteV1MarketingCalendarId**](MarketingAPI.md#CloudDeleteV1MarketingCalendarId) | **Delete** /v1/marketing/calendar/{id} | Removes one of the caller org&#39;s posts and answers 204.
[**CloudDeleteV1MarketingCampaignsId**](MarketingAPI.md#CloudDeleteV1MarketingCampaignsId) | **Delete** /v1/marketing/campaigns/{id} | Removes one of the caller org&#39;s campaigns and answers 204.
[**CloudDeleteV1MarketingSuppressions**](MarketingAPI.md#CloudDeleteV1MarketingSuppressions) | **Delete** /v1/marketing/suppressions | Re-subscribes an address on one channel and answers 204.
[**CloudGetV1MarketingAudiences**](MarketingAPI.md#CloudGetV1MarketingAudiences) | **Get** /v1/marketing/audiences | Returns the org&#39;s saved audiences, most recently updated first.
[**CloudGetV1MarketingAudiencesId**](MarketingAPI.md#CloudGetV1MarketingAudiencesId) | **Get** /v1/marketing/audiences/{id} | Returns one of the caller org&#39;s saved audiences.
[**CloudGetV1MarketingAudiencesIdPreview**](MarketingAPI.md#CloudGetV1MarketingAudiencesIdPreview) | **Get** /v1/marketing/audiences/{id}/preview | Evaluates the cohort LIVE — the same resolution an enrollment would run — and reports how big it is and how many real mailboxes it reaches.
[**CloudGetV1MarketingCalendar**](MarketingAPI.md#CloudGetV1MarketingCalendar) | **Get** /v1/marketing/calendar | Returns the org&#39;s calendar, soonest scheduled first, optionally narrowed to one status.
[**CloudGetV1MarketingCalendarId**](MarketingAPI.md#CloudGetV1MarketingCalendarId) | **Get** /v1/marketing/calendar/{id} | Returns one of the caller org&#39;s posts, including the exact error behind a failed publish.
[**CloudGetV1MarketingCampaigns**](MarketingAPI.md#CloudGetV1MarketingCampaigns) | **Get** /v1/marketing/campaigns | Returns the org&#39;s campaigns, most recently updated first, optionally narrowed to one lifecycle status.
[**CloudGetV1MarketingCampaignsId**](MarketingAPI.md#CloudGetV1MarketingCampaignsId) | **Get** /v1/marketing/campaigns/{id} | Returns one of the caller org&#39;s campaigns.
[**CloudGetV1MarketingPromos**](MarketingAPI.md#CloudGetV1MarketingPromos) | **Get** /v1/marketing/promos | Returns every promo the deployment offers with its live counters: how many orgs have redeemed it and how many redemptions remain under the cap.
[**CloudGetV1MarketingPromosCodeEligibility**](MarketingAPI.md#CloudGetV1MarketingPromosCodeEligibility) | **Get** /v1/marketing/promos/{code}/eligibility | Prices a promo against a plan and seat count.
[**CloudGetV1MarketingPromosCodeRedemption**](MarketingAPI.md#CloudGetV1MarketingPromosCodeRedemption) | **Get** /v1/marketing/promos/{code}/redemption | Returns the caller org&#39;s OWN redemption of a promo — an org-scoped read, so it can never surface another tenant&#39;s.
[**CloudGetV1MarketingSequences**](MarketingAPI.md#CloudGetV1MarketingSequences) | **Get** /v1/marketing/sequences | Returns the org&#39;s drip sequences, most recently updated first.
[**CloudGetV1MarketingSequencesId**](MarketingAPI.md#CloudGetV1MarketingSequencesId) | **Get** /v1/marketing/sequences/{id} | Returns one of the caller org&#39;s sequences together with its steps in send order.
[**CloudGetV1MarketingSequencesIdEnrollments**](MarketingAPI.md#CloudGetV1MarketingSequencesIdEnrollments) | **Get** /v1/marketing/sequences/{id}/enrollments | Returns who is walking one sequence, most recently enrolled first, with each walk&#39;s current step and next due time.
[**CloudGetV1MarketingSequencesIdSteps**](MarketingAPI.md#CloudGetV1MarketingSequencesIdSteps) | **Get** /v1/marketing/sequences/{id}/steps | Returns one sequence&#39;s steps in send order.
[**CloudGetV1MarketingSummary**](MarketingAPI.md#CloudGetV1MarketingSummary) | **Get** /v1/marketing/summary | Rolls up the caller org&#39;s campaigns: how many there are, how many are active, and the summed budget and spend in cents.
[**CloudGetV1MarketingSuppressions**](MarketingAPI.md#CloudGetV1MarketingSuppressions) | **Get** /v1/marketing/suppressions | Returns the org&#39;s opt-out list, newest first — everyone the send gate will refuse to deliver to.
[**CloudGetV1MarketingUnsubscribe**](MarketingAPI.md#CloudGetV1MarketingUnsubscribe) | **Get** /v1/marketing/unsubscribe | Is the PUBLIC one-click endpoint (no principal): a recipient clicks the signed link in an email footer.
[**CloudPostV1MarketingAudiences**](MarketingAPI.md#CloudPostV1MarketingAudiences) | **Post** /v1/marketing/audiences | Saves a cohort filter for the caller&#39;s org.
[**CloudPostV1MarketingCalendar**](MarketingAPI.md#CloudPostV1MarketingCalendar) | **Post** /v1/marketing/calendar | Adds a post to the content calendar.
[**CloudPostV1MarketingCalendarIdPublish**](MarketingAPI.md#CloudPostV1MarketingCalendarIdPublish) | **Post** /v1/marketing/calendar/{id}/publish | Publishes a post NOW, synchronously, whatever its schedule.
[**CloudPostV1MarketingCampaigns**](MarketingAPI.md#CloudPostV1MarketingCampaigns) | **Post** /v1/marketing/campaigns | Registers a campaign in the caller&#39;s org.
[**CloudPostV1MarketingCampaignsIdSchedule**](MarketingAPI.md#CloudPostV1MarketingCampaignsIdSchedule) | **Post** /v1/marketing/campaigns/{id}/schedule | Sets a campaign&#39;s send time and moves it to \&quot;scheduled\&quot;.
[**CloudPostV1MarketingPromosCodeRedeem**](MarketingAPI.md#CloudPostV1MarketingPromosCodeRedeem) | **Post** /v1/marketing/promos/{code}/redeem | Redeems the promo for the caller&#39;s org, crediting the discount value to its wallet through the finance ledger.
[**CloudPostV1MarketingSequences**](MarketingAPI.md#CloudPostV1MarketingSequences) | **Post** /v1/marketing/sequences | Registers a drip sequence in the caller&#39;s org.
[**CloudPostV1MarketingSequencesIdEnroll**](MarketingAPI.md#CloudPostV1MarketingSequencesIdEnroll) | **Post** /v1/marketing/sequences/{id}/enroll | Adds one contact or a whole audience to a sequence and schedules the first step for each.
[**CloudPostV1MarketingSequencesIdEnrollmentsEidCancel**](MarketingAPI.md#CloudPostV1MarketingSequencesIdEnrollmentsEidCancel) | **Post** /v1/marketing/sequences/{id}/enrollments/{eid}/cancel | Stops one walk mid-sequence and answers 204: no further step is sent, and steps already delivered are not recalled.
[**CloudPostV1MarketingSequencesIdStatus**](MarketingAPI.md#CloudPostV1MarketingSequencesIdStatus) | **Post** /v1/marketing/sequences/{id}/status | Flips draft/active/archived — the activation gate for sending, since only an active sequence accepts enrollments.
[**CloudPostV1MarketingSequencesIdSteps**](MarketingAPI.md#CloudPostV1MarketingSequencesIdSteps) | **Post** /v1/marketing/sequences/{id}/steps | Appends a message to the END of a sequence: the new step&#39;s idx is one past the last, so steps arrive in the order they are added.
[**CloudPostV1MarketingSuppressions**](MarketingAPI.md#CloudPostV1MarketingSuppressions) | **Post** /v1/marketing/suppressions | Records an opt-out for the org (admin / self-service management).
[**CloudPutV1MarketingCalendarId**](MarketingAPI.md#CloudPutV1MarketingCalendarId) | **Put** /v1/marketing/calendar/{id} | Replaces a post&#39;s editable fields.
[**CloudPutV1MarketingCampaignsId**](MarketingAPI.md#CloudPutV1MarketingCampaignsId) | **Put** /v1/marketing/campaigns/{id} | Replaces a campaign&#39;s editable fields.
[**CommerceCreateMarketingCampaign**](MarketingAPI.md#CommerceCreateMarketingCampaign) | **Post** /v1/commerce/marketing | Create marketing campaign



## CloudDeleteV1MarketingAudiencesId

> CloudDeleteV1MarketingAudiencesId(ctx, id).Execute()

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
	r, err := apiClient.MarketingAPI.CloudDeleteV1MarketingAudiencesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudDeleteV1MarketingAudiencesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MarketingAudiencesIdRequest struct via the builder pattern


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


## CloudDeleteV1MarketingCalendarId

> CloudDeleteV1MarketingCalendarId(ctx, id).Execute()

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
	r, err := apiClient.MarketingAPI.CloudDeleteV1MarketingCalendarId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudDeleteV1MarketingCalendarId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MarketingCalendarIdRequest struct via the builder pattern


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


## CloudDeleteV1MarketingCampaignsId

> CloudDeleteV1MarketingCampaignsId(ctx, id).Execute()

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
	r, err := apiClient.MarketingAPI.CloudDeleteV1MarketingCampaignsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudDeleteV1MarketingCampaignsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1MarketingCampaignsIdRequest struct via the builder pattern


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


## CloudDeleteV1MarketingSuppressions

> CloudDeleteV1MarketingSuppressions(ctx).Channel(channel).Address(address).Reason(reason).CreatedAt(createdAt).Execute()

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
	r, err := apiClient.MarketingAPI.CloudDeleteV1MarketingSuppressions(context.Background()).Channel(channel).Address(address).Reason(reason).CreatedAt(createdAt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudDeleteV1MarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1MarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channel** | **string** | Channel is the surface opted out of: email, sms, social, meta, google or tiktok. Empty means email. Opting out of one leaves the others reachable. | 
 **address** | **string** | Address is the recipient, normalized (lower-cased, trimmed) so an opt-out cannot be slipped past on a case or whitespace difference. Required. | 
 **reason** | **string** | Reason is a free-text note, capped at 1024 bytes. The public one-click endpoint records \&quot;one-click unsubscribe\&quot;. | 
 **createdAt** | **int32** | CreatedAt is unix seconds, server-assigned. | 

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


## CloudGetV1MarketingAudiences

> CloudAudienceList CloudGetV1MarketingAudiences(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingAudiences(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingAudiences`: CloudAudienceList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudAudienceList**](CloudAudienceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingAudiencesId

> CloudAudience CloudGetV1MarketingAudiencesId(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingAudiencesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingAudiencesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingAudiencesId`: CloudAudience
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingAudiencesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the audience id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingAudiencesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAudience**](CloudAudience.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingAudiencesIdPreview

> CloudAudiencePreview CloudGetV1MarketingAudiencesIdPreview(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingAudiencesIdPreview(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingAudiencesIdPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingAudiencesIdPreview`: CloudAudiencePreview
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingAudiencesIdPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the audience id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingAudiencesIdPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudAudiencePreview**](CloudAudiencePreview.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingCalendar

> CloudPostList CloudGetV1MarketingCalendar(ctx).Status(status).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingCalendar(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingCalendar`: CloudPostList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only posts in that state (draft, scheduled, published, failed, canceled). Empty means every post. | 
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudPostList**](CloudPostList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingCalendarId

> CloudCalendarPost CloudGetV1MarketingCalendarId(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingCalendarId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingCalendarId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingCalendarId`: CloudCalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingCalendarId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the post id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingCalendarIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCalendarPost**](CloudCalendarPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingCampaigns

> CloudCampaignList CloudGetV1MarketingCampaigns(ctx).Status(status).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingCampaigns(context.Background()).Status(status).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingCampaigns`: CloudCampaignList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Status keeps only campaigns in that lifecycle state (draft, scheduled, active, paused, completed). Empty means every campaign. | 
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudCampaignList**](CloudCampaignList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingCampaignsId

> CloudCampaign CloudGetV1MarketingCampaignsId(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingCampaignsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingCampaignsId`: CloudCampaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCampaign**](CloudCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingPromos

> CloudPromoList CloudGetV1MarketingPromos(ctx).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingPromos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingPromos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingPromos`: CloudPromoList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingPromos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingPromosRequest struct via the builder pattern


### Return type

[**CloudPromoList**](CloudPromoList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingPromosCodeEligibility

> CloudQuote CloudGetV1MarketingPromosCodeEligibility(ctx, code).Plan(plan).Seats(seats).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingPromosCodeEligibility(context.Background(), code).Plan(plan).Seats(seats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingPromosCodeEligibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingPromosCodeEligibility`: CloudQuote
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingPromosCodeEligibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingPromosCodeEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plan** | **string** | Plan is the plan being priced: pro, max or team. Anything else (including the free Developer plan) has no list price and so nothing to discount. | 
 **seats** | **int32** | Seats is the Team seat count; 0 means 1, and it is ignored for the single-seat plans. | 

### Return type

[**CloudQuote**](CloudQuote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingPromosCodeRedemption

> CloudRedemption CloudGetV1MarketingPromosCodeRedemption(ctx, code).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingPromosCodeRedemption(context.Background(), code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingPromosCodeRedemption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingPromosCodeRedemption`: CloudRedemption
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingPromosCodeRedemption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path, e.g. \&quot;first1000\&quot;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingPromosCodeRedemptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRedemption**](CloudRedemption.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSequences

> CloudSequenceList CloudGetV1MarketingSequences(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSequences(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSequences`: CloudSequenceList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudSequenceList**](CloudSequenceList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSequencesId

> CloudSequenceView CloudGetV1MarketingSequencesId(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSequencesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSequencesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSequencesId`: CloudSequenceView
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSequencesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSequencesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSequenceView**](CloudSequenceView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSequencesIdEnrollments

> CloudEnrollmentList CloudGetV1MarketingSequencesIdEnrollments(ctx, id).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSequencesIdEnrollments(context.Background(), id).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSequencesIdEnrollments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSequencesIdEnrollments`: CloudEnrollmentList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSequencesIdEnrollments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSequencesIdEnrollmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudEnrollmentList**](CloudEnrollmentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSequencesIdSteps

> CloudStepList CloudGetV1MarketingSequencesIdSteps(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSequencesIdSteps(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSequencesIdSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSequencesIdSteps`: CloudStepList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSequencesIdSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSequencesIdStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudStepList**](CloudStepList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSummary

> CloudSummary CloudGetV1MarketingSummary(ctx).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSummary`: CloudSummary
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSummaryRequest struct via the builder pattern


### Return type

[**CloudSummary**](CloudSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingSuppressions

> CloudSuppressionList CloudGetV1MarketingSuppressions(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingSuppressions(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingSuppressions`: CloudSuppressionList
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned; 0 means 200 and nothing above 1000 is honoured. | 

### Return type

[**CloudSuppressionList**](CloudSuppressionList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1MarketingUnsubscribe

> CloudUnsubscribed CloudGetV1MarketingUnsubscribe(ctx).Org(org).Channel(channel).Address(address).Token(token).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudGetV1MarketingUnsubscribe(context.Background()).Org(org).Channel(channel).Address(address).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudGetV1MarketingUnsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1MarketingUnsubscribe`: CloudUnsubscribed
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudGetV1MarketingUnsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1MarketingUnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | Org is the org the link was minted for. | 
 **channel** | **string** | Channel is the surface to opt out of. | 
 **address** | **string** | Address is the recipient to opt out. | 
 **token** | **string** | Token is the HMAC over (org, channel, address). It is the ONLY authority here — there is no principal — so it binds the request to one tuple and nothing else. | 

### Return type

[**CloudUnsubscribed**](CloudUnsubscribed.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingAudiences

> CloudAudience CloudPostV1MarketingAudiences(ctx).CloudAudience(cloudAudience).Execute()

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
	cloudAudience := *openapiclient.NewCloudAudience() // CloudAudience | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingAudiences(context.Background()).CloudAudience(cloudAudience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingAudiences`: CloudAudience
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudAudience** | [**CloudAudience**](CloudAudience.md) |  | 

### Return type

[**CloudAudience**](CloudAudience.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingCalendar

> CloudCalendarPost CloudPostV1MarketingCalendar(ctx).CloudCalendarPost(cloudCalendarPost).Execute()

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
	cloudCalendarPost := *openapiclient.NewCloudCalendarPost() // CloudCalendarPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingCalendar(context.Background()).CloudCalendarPost(cloudCalendarPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingCalendar`: CloudCalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingCalendar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCalendarPost** | [**CloudCalendarPost**](CloudCalendarPost.md) |  | 

### Return type

[**CloudCalendarPost**](CloudCalendarPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingCalendarIdPublish

> CloudCalendarPost CloudPostV1MarketingCalendarIdPublish(ctx, id).Execute()

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
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingCalendarIdPublish(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingCalendarIdPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingCalendarIdPublish`: CloudCalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingCalendarIdPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the post id from the path, as returned by create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingCalendarIdPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCalendarPost**](CloudCalendarPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingCampaigns

> CloudCampaign CloudPostV1MarketingCampaigns(ctx).CloudCampaign(cloudCampaign).Execute()

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
	cloudCampaign := *openapiclient.NewCloudCampaign() // CloudCampaign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingCampaigns(context.Background()).CloudCampaign(cloudCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingCampaigns`: CloudCampaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCampaign** | [**CloudCampaign**](CloudCampaign.md) |  | 

### Return type

[**CloudCampaign**](CloudCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingCampaignsIdSchedule

> CloudCampaign CloudPostV1MarketingCampaignsIdSchedule(ctx, id).CloudScheduleInput(cloudScheduleInput).Execute()

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
	cloudScheduleInput := *openapiclient.NewCloudScheduleInput() // CloudScheduleInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingCampaignsIdSchedule(context.Background(), id).CloudScheduleInput(cloudScheduleInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingCampaignsIdSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingCampaignsIdSchedule`: CloudCampaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingCampaignsIdSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the campaign id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingCampaignsIdScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudScheduleInput** | [**CloudScheduleInput**](CloudScheduleInput.md) |  | 

### Return type

[**CloudCampaign**](CloudCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingPromosCodeRedeem

> CloudRedeemResult CloudPostV1MarketingPromosCodeRedeem(ctx, code).CloudRedeemInput(cloudRedeemInput).Execute()

Redeems the promo for the caller's org, crediting the discount value to its wallet through the finance ledger.



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
	cloudRedeemInput := *openapiclient.NewCloudRedeemInput() // CloudRedeemInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingPromosCodeRedeem(context.Background(), code).CloudRedeemInput(cloudRedeemInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingPromosCodeRedeem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingPromosCodeRedeem`: CloudRedeemResult
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingPromosCodeRedeem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string** | Code is the promo code from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingPromosCodeRedeemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRedeemInput** | [**CloudRedeemInput**](CloudRedeemInput.md) |  | 

### Return type

[**CloudRedeemResult**](CloudRedeemResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingSequences

> CloudSequence CloudPostV1MarketingSequences(ctx).CloudSequence(cloudSequence).Execute()

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
	cloudSequence := *openapiclient.NewCloudSequence() // CloudSequence | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingSequences(context.Background()).CloudSequence(cloudSequence).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSequences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingSequences`: CloudSequence
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingSequences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSequencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSequence** | [**CloudSequence**](CloudSequence.md) |  | 

### Return type

[**CloudSequence**](CloudSequence.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingSequencesIdEnroll

> CloudEnrollResult CloudPostV1MarketingSequencesIdEnroll(ctx, id).CloudEnrollInput(cloudEnrollInput).Execute()

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
	cloudEnrollInput := *openapiclient.NewCloudEnrollInput() // CloudEnrollInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingSequencesIdEnroll(context.Background(), id).CloudEnrollInput(cloudEnrollInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSequencesIdEnroll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingSequencesIdEnroll`: CloudEnrollResult
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingSequencesIdEnroll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSequencesIdEnrollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudEnrollInput** | [**CloudEnrollInput**](CloudEnrollInput.md) |  | 

### Return type

[**CloudEnrollResult**](CloudEnrollResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingSequencesIdEnrollmentsEidCancel

> CloudPostV1MarketingSequencesIdEnrollmentsEidCancel(ctx, id, eid).Execute()

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
	r, err := apiClient.MarketingAPI.CloudPostV1MarketingSequencesIdEnrollmentsEidCancel(context.Background(), id, eid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSequencesIdEnrollmentsEidCancel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSequencesIdEnrollmentsEidCancelRequest struct via the builder pattern


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


## CloudPostV1MarketingSequencesIdStatus

> CloudSequenceStatus CloudPostV1MarketingSequencesIdStatus(ctx, id).CloudSequenceStatus(cloudSequenceStatus).Execute()

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
	cloudSequenceStatus := *openapiclient.NewCloudSequenceStatus() // CloudSequenceStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingSequencesIdStatus(context.Background(), id).CloudSequenceStatus(cloudSequenceStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSequencesIdStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingSequencesIdStatus`: CloudSequenceStatus
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingSequencesIdStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the sequence id from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSequencesIdStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSequenceStatus** | [**CloudSequenceStatus**](CloudSequenceStatus.md) |  | 

### Return type

[**CloudSequenceStatus**](CloudSequenceStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingSequencesIdSteps

> CloudStep CloudPostV1MarketingSequencesIdSteps(ctx, id).CloudStepInput(cloudStepInput).Execute()

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
	cloudStepInput := *openapiclient.NewCloudStepInput() // CloudStepInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingSequencesIdSteps(context.Background(), id).CloudStepInput(cloudStepInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSequencesIdSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingSequencesIdSteps`: CloudStep
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingSequencesIdSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SequenceID is the sequence id from the path (the route&#39;s :id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSequencesIdStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudStepInput** | [**CloudStepInput**](CloudStepInput.md) |  | 

### Return type

[**CloudStep**](CloudStep.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1MarketingSuppressions

> CloudSuppression CloudPostV1MarketingSuppressions(ctx).CloudSuppression(cloudSuppression).Execute()

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
	cloudSuppression := *openapiclient.NewCloudSuppression() // CloudSuppression | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPostV1MarketingSuppressions(context.Background()).CloudSuppression(cloudSuppression).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPostV1MarketingSuppressions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1MarketingSuppressions`: CloudSuppression
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPostV1MarketingSuppressions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1MarketingSuppressionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudSuppression** | [**CloudSuppression**](CloudSuppression.md) |  | 

### Return type

[**CloudSuppression**](CloudSuppression.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1MarketingCalendarId

> CloudCalendarPost CloudPutV1MarketingCalendarId(ctx, id).CloudCalendarPost(cloudCalendarPost).Execute()

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
	cloudCalendarPost := *openapiclient.NewCloudCalendarPost() // CloudCalendarPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPutV1MarketingCalendarId(context.Background(), id).CloudCalendarPost(cloudCalendarPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPutV1MarketingCalendarId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1MarketingCalendarId`: CloudCalendarPost
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPutV1MarketingCalendarId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server-assigned post id (\&quot;cal_\&quot; + 128 random bits). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1MarketingCalendarIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCalendarPost** | [**CloudCalendarPost**](CloudCalendarPost.md) |  | 

### Return type

[**CloudCalendarPost**](CloudCalendarPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1MarketingCampaignsId

> CloudCampaign CloudPutV1MarketingCampaignsId(ctx, id).CloudCampaign(cloudCampaign).Execute()

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
	cloudCampaign := *openapiclient.NewCloudCampaign() // CloudCampaign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CloudPutV1MarketingCampaignsId(context.Background(), id).CloudCampaign(cloudCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CloudPutV1MarketingCampaignsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1MarketingCampaignsId`: CloudCampaign
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CloudPutV1MarketingCampaignsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the server-assigned campaign id (\&quot;camp_\&quot; + 128 random bits). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1MarketingCampaignsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCampaign** | [**CloudCampaign**](CloudCampaign.md) |  | 

### Return type

[**CloudCampaign**](CloudCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceCreateMarketingCampaign

> map[string]interface{} CommerceCreateMarketingCampaign(ctx).Body(body).Execute()

Create marketing campaign

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingAPI.CommerceCreateMarketingCampaign(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingAPI.CommerceCreateMarketingCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateMarketingCampaign`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MarketingAPI.CommerceCreateMarketingCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateMarketingCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

