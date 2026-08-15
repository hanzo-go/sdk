# \GuideAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGuideCurriculum**](GuideAPI.md#DeleteGuideCurriculum) | **Delete** /v1/guide/curriculum | Clears the caller org&#39;s curriculum override and returns the journey it falls back to — the brand blueprint, else the embedded fixture.
[**GetGuide**](GuideAPI.md#GetGuide) | **Get** /v1/guide | Overview returns the caller org&#39;s launch journey: the active curriculum&#39;s version and title, every step with its state, whether it is available, what blocks it and whether the Business AI can run it, the done/total/percent progress with the next step to take, and the org&#39;s analytics funnel folded in.
[**GetGuideActions**](GuideAPI.md#GetGuideActions) | **Get** /v1/guide/actions | Returns the caller org&#39;s Business AI action ledger, most recent first: every \&quot;do it for me\&quot; tool call, the arguments it ran with, its result and whether it succeeded.
[**GetGuideAnalytics**](GuideAPI.md#GetGuideAnalytics) | **Get** /v1/guide/analytics | Analytics returns the caller org&#39;s funnel from the analytics lens plus the GTM recommendations derived from it.
[**GetGuideBlueprint**](GuideAPI.md#GetGuideBlueprint) | **Get** /v1/guide/blueprint | Returns the FULL authored brand blueprint — every principle, section, step, strategy and template WITH its enabled flag made explicit, including the disabled items the org-facing reads never see — plus the active version number, the brand key it is stored under and the item counts.
[**GetGuideBlueprintVersions**](GuideAPI.md#GetGuideBlueprintVersions) | **Get** /v1/guide/blueprint/versions | Returns the brand blueprint&#39;s version history — every stored version&#39;s number and edit time, newest first — which is the point-in-time-recovery and audit trail behind the authoring plane.
[**GetGuideCurriculum**](GuideAPI.md#GetGuideCurriculum) | **Get** /v1/guide/curriculum | Returns the journey the caller&#39;s org is actually running, and whether it comes from the org&#39;s OWN override (custom) or from the platform default — the brand blueprint, else the embedded fixture.
[**GetGuideProfile**](GuideAPI.md#GetGuideProfile) | **Get** /v1/guide/profile | Profile returns the caller org&#39;s OBSERVED growth profile — the signal set, the classified growth stage, and the org&#39;s own key metrics.
[**GetGuideStrategies**](GuideAPI.md#GetGuideStrategies) | **Get** /v1/guide/strategies | Strategies returns the ENABLED tactics corpus for the caller&#39;s org: the tactics library narrowed by the explicit category/workload filters AND by the org&#39;s OBSERVED growth stage and capability signals (a tactic&#39;s tags are preconditions, so it surfaces only once the org can act on it).
[**GetGuideSuggest**](GuideAPI.md#GetGuideSuggest) | **Get** /v1/guide/suggest | Suggest returns the caller org&#39;s next-best quests: the available, non-terminal steps of its journey ranked by how much downstream work each unblocks, each with the grounded reason it is a good next move and whether the Business AI can run it, plus the org&#39;s funnel and the GTM recommendations derived from it.
[**PatchGuideBlueprintByCollectionById**](GuideAPI.md#PatchGuideBlueprintByCollectionById) | **Patch** /v1/guide/blueprint/{collection}/{id} | Edit — or retire — one item of the brand blueprint
[**PostGuideChat**](GuideAPI.md#PostGuideChat) | **Post** /v1/guide/chat | Chat answers a founder&#39;s question about their launch journey as the Business AI coach: it grounds the reply in the org&#39;s REAL progress, its ranked available quests and its analytics funnel, and returns those candidate quests alongside so the caller can act on one.
[**PostGuideStepsByIdDo**](GuideAPI.md#PostGuideStepsByIdDo) | **Post** /v1/guide/steps/{id}/do | Have the Business AI actually do the step for you
[**PostGuideStepsByIdDone**](GuideAPI.md#PostGuideStepsByIdDone) | **Post** /v1/guide/steps/{id}/done | Mark a step of your org&#39;s journey finished
[**PostGuideStepsByIdReset**](GuideAPI.md#PostGuideStepsByIdReset) | **Post** /v1/guide/steps/{id}/reset | Returns one step of the caller org&#39;s journey to todo — clearing a manual mark or a skip — and returns the refreshed journey.
[**PostGuideStepsByIdSkip**](GuideAPI.md#PostGuideStepsByIdSkip) | **Post** /v1/guide/steps/{id}/skip | Marks one step of the caller org&#39;s journey skipped and returns the refreshed journey.
[**PostGuideStepsByIdStart**](GuideAPI.md#PostGuideStepsByIdStart) | **Post** /v1/guide/steps/{id}/start | Mark a step of your org&#39;s journey started
[**PutGuideBlueprint**](GuideAPI.md#PutGuideBlueprint) | **Put** /v1/guide/blueprint | Publish a new version of the brand blueprint
[**PutGuideCurriculum**](GuideAPI.md#PutGuideCurriculum) | **Put** /v1/guide/curriculum | Replace your org&#39;s journey with a curriculum you author



## DeleteGuideCurriculum

> CurriculumView DeleteGuideCurriculum(ctx).Execute()

Clears the caller org's curriculum override and returns the journey it falls back to — the brand blueprint, else the embedded fixture.



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
	resp, r, err := apiClient.GuideAPI.DeleteGuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.DeleteGuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGuideCurriculum`: CurriculumView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.DeleteGuideCurriculum`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGuideCurriculumRequest struct via the builder pattern


### Return type

[**CurriculumView**](CurriculumView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuide

> OverviewView GetGuide(ctx).Execute()

Overview returns the caller org's launch journey: the active curriculum's version and title, every step with its state, whether it is available, what blocks it and whether the Business AI can run it, the done/total/percent progress with the next step to take, and the org's analytics funnel folded in.



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
	resp, r, err := apiClient.GuideAPI.GetGuide(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuide`: OverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuide`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideRequest struct via the builder pattern


### Return type

[**OverviewView**](OverviewView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideActions

> ActionsView GetGuideActions(ctx).Execute()

Returns the caller org's Business AI action ledger, most recent first: every \"do it for me\" tool call, the arguments it ran with, its result and whether it succeeded.



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
	resp, r, err := apiClient.GuideAPI.GetGuideActions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideActions`: ActionsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideActionsRequest struct via the builder pattern


### Return type

[**ActionsView**](ActionsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideAnalytics

> AnalyticsView GetGuideAnalytics(ctx).Execute()

Analytics returns the caller org's funnel from the analytics lens plus the GTM recommendations derived from it.



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
	resp, r, err := apiClient.GuideAPI.GetGuideAnalytics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideAnalytics`: AnalyticsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideAnalytics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideAnalyticsRequest struct via the builder pattern


### Return type

[**AnalyticsView**](AnalyticsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideBlueprint

> BlueprintView GetGuideBlueprint(ctx).Execute()

Returns the FULL authored brand blueprint — every principle, section, step, strategy and template WITH its enabled flag made explicit, including the disabled items the org-facing reads never see — plus the active version number, the brand key it is stored under and the item counts.



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
	resp, r, err := apiClient.GuideAPI.GetGuideBlueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideBlueprint`: BlueprintView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideBlueprint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideBlueprintRequest struct via the builder pattern


### Return type

[**BlueprintView**](BlueprintView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideBlueprintVersions

> BlueprintVersionsView GetGuideBlueprintVersions(ctx).Execute()

Returns the brand blueprint's version history — every stored version's number and edit time, newest first — which is the point-in-time-recovery and audit trail behind the authoring plane.



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
	resp, r, err := apiClient.GuideAPI.GetGuideBlueprintVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideBlueprintVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideBlueprintVersions`: BlueprintVersionsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideBlueprintVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideBlueprintVersionsRequest struct via the builder pattern


### Return type

[**BlueprintVersionsView**](BlueprintVersionsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideCurriculum

> CurriculumView GetGuideCurriculum(ctx).Execute()

Returns the journey the caller's org is actually running, and whether it comes from the org's OWN override (custom) or from the platform default — the brand blueprint, else the embedded fixture.



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
	resp, r, err := apiClient.GuideAPI.GetGuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideCurriculum`: CurriculumView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideCurriculum`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideCurriculumRequest struct via the builder pattern


### Return type

[**CurriculumView**](CurriculumView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideProfile

> ProfileResponse GetGuideProfile(ctx).Execute()

Profile returns the caller org's OBSERVED growth profile — the signal set, the classified growth stage, and the org's own key metrics.



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
	resp, r, err := apiClient.GuideAPI.GetGuideProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideProfile`: ProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideProfileRequest struct via the builder pattern


### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideStrategies

> CorpusView GetGuideStrategies(ctx).Category(category).Stage(stage).Workload(workload).Execute()

Strategies returns the ENABLED tactics corpus for the caller's org: the tactics library narrowed by the explicit category/workload filters AND by the org's OBSERVED growth stage and capability signals (a tactic's tags are preconditions, so it surfaces only once the org can act on it).



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
	category := "viral-coefficient" // string | Category filters to tactics in exactly this category. (optional)
	stage := "scaling" // string | Stage previews the corpus at a chosen growth stage (research|formed|launched|activated|scaling), overriding the org's observed one. An unknown value is ignored and the observed stage stands. (optional)
	workload := "workload_example" // string | Workload filters to tactics with exactly this workload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuideAPI.GetGuideStrategies(context.Background()).Category(category).Stage(stage).Workload(workload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideStrategies`: CorpusView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideStrategies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Category filters to tactics in exactly this category. | 
 **stage** | **string** | Stage previews the corpus at a chosen growth stage (research|formed|launched|activated|scaling), overriding the org&#39;s observed one. An unknown value is ignored and the observed stage stands. | 
 **workload** | **string** | Workload filters to tactics with exactly this workload. | 

### Return type

[**CorpusView**](CorpusView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGuideSuggest

> SuggestResponse GetGuideSuggest(ctx).Execute()

Suggest returns the caller org's next-best quests: the available, non-terminal steps of its journey ranked by how much downstream work each unblocks, each with the grounded reason it is a good next move and whether the Business AI can run it, plus the org's funnel and the GTM recommendations derived from it.



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
	resp, r, err := apiClient.GuideAPI.GetGuideSuggest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.GetGuideSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGuideSuggest`: SuggestResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.GetGuideSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuideSuggestRequest struct via the builder pattern


### Return type

[**SuggestResponse**](SuggestResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGuideBlueprintByCollectionById

> PatchGuideBlueprintByCollectionById(ctx, collection, id).Execute()

Edit — or retire — one item of the brand blueprint



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
	collection := "collection_example" // string | 
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GuideAPI.PatchGuideBlueprintByCollectionById(context.Background(), collection, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PatchGuideBlueprintByCollectionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collection** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGuideBlueprintByCollectionByIdRequest struct via the builder pattern


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


## PostGuideChat

> ChatResponse PostGuideChat(ctx).ChatRequest(chatRequest).Execute()

Chat answers a founder's question about their launch journey as the Business AI coach: it grounds the reply in the org's REAL progress, its ranked available quests and its analytics funnel, and returns those candidate quests alongside so the caller can act on one.



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
	chatRequest := *openapiclient.NewChatRequest() // ChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuideAPI.PostGuideChat(context.Background()).ChatRequest(chatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGuideChat`: ChatResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.PostGuideChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGuideChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatRequest** | [**ChatRequest**](ChatRequest.md) |  | 

### Return type

[**ChatResponse**](ChatResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGuideStepsByIdDo

> PostGuideStepsByIdDo(ctx, id).Execute()

Have the Business AI actually do the step for you



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
	r, err := apiClient.GuideAPI.PostGuideStepsByIdDo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideStepsByIdDo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGuideStepsByIdDoRequest struct via the builder pattern


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


## PostGuideStepsByIdDone

> PostGuideStepsByIdDone(ctx, id).Execute()

Mark a step of your org's journey finished



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
	r, err := apiClient.GuideAPI.PostGuideStepsByIdDone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideStepsByIdDone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGuideStepsByIdDoneRequest struct via the builder pattern


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


## PostGuideStepsByIdReset

> OverviewView PostGuideStepsByIdReset(ctx, id).Execute()

Returns one step of the caller org's journey to todo — clearing a manual mark or a skip — and returns the refreshed journey.



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
	id := "id_example" // string | ID is the step's id, as it appears in the journey (e.g. \"gsuite\").

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuideAPI.PostGuideStepsByIdReset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideStepsByIdReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGuideStepsByIdReset`: OverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.PostGuideStepsByIdReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the step&#39;s id, as it appears in the journey (e.g. \&quot;gsuite\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGuideStepsByIdResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OverviewView**](OverviewView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGuideStepsByIdSkip

> OverviewView PostGuideStepsByIdSkip(ctx, id).Execute()

Marks one step of the caller org's journey skipped and returns the refreshed journey.



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
	id := "id_example" // string | ID is the step's id, as it appears in the journey (e.g. \"gsuite\").

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuideAPI.PostGuideStepsByIdSkip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideStepsByIdSkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGuideStepsByIdSkip`: OverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.PostGuideStepsByIdSkip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the step&#39;s id, as it appears in the journey (e.g. \&quot;gsuite\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGuideStepsByIdSkipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OverviewView**](OverviewView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGuideStepsByIdStart

> PostGuideStepsByIdStart(ctx, id).Execute()

Mark a step of your org's journey started



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
	r, err := apiClient.GuideAPI.PostGuideStepsByIdStart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PostGuideStepsByIdStart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostGuideStepsByIdStartRequest struct via the builder pattern


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


## PutGuideBlueprint

> PutGuideBlueprint(ctx).Execute()

Publish a new version of the brand blueprint



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
	r, err := apiClient.GuideAPI.PutGuideBlueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PutGuideBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutGuideBlueprintRequest struct via the builder pattern


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


## PutGuideCurriculum

> PutGuideCurriculum(ctx).Execute()

Replace your org's journey with a curriculum you author



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
	r, err := apiClient.GuideAPI.PutGuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.PutGuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutGuideCurriculumRequest struct via the builder pattern


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

