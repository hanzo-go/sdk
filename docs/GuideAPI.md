# \GuideAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1GuideCurriculum**](GuideAPI.md#CloudDeleteV1GuideCurriculum) | **Delete** /v1/guide/curriculum | DeleteCurriculum clears the caller org&#39;s curriculum override and returns the journey it falls back to — the brand blueprint, else the embedded fixture.
[**CloudGetV1Guide**](GuideAPI.md#CloudGetV1Guide) | **Get** /v1/guide | Overview returns the caller org&#39;s launch journey: the active curriculum&#39;s version and title, every step with its state, whether it is available, what blocks it and whether the Business AI can run it, the done/total/percent progress with the next step to take, and the org&#39;s analytics funnel folded in.
[**CloudGetV1GuideActions**](GuideAPI.md#CloudGetV1GuideActions) | **Get** /v1/guide/actions | ListActions returns the caller org&#39;s Business AI action ledger, most recent first: every \&quot;do it for me\&quot; tool call, the arguments it ran with, its result and whether it succeeded.
[**CloudGetV1GuideAnalytics**](GuideAPI.md#CloudGetV1GuideAnalytics) | **Get** /v1/guide/analytics | Analytics returns the caller org&#39;s funnel from the analytics lens plus the GTM recommendations derived from it.
[**CloudGetV1GuideBlueprint**](GuideAPI.md#CloudGetV1GuideBlueprint) | **Get** /v1/guide/blueprint | GetBlueprint returns the FULL authored brand blueprint — every principle, section, step, strategy and template WITH its enabled flag made explicit, including the disabled items the org-facing reads never see — plus the active version number, the brand key it is stored under and the item counts.
[**CloudGetV1GuideBlueprintVersions**](GuideAPI.md#CloudGetV1GuideBlueprintVersions) | **Get** /v1/guide/blueprint/versions | ListBlueprintVersions returns the brand blueprint&#39;s version history — every stored version&#39;s number and edit time, newest first — which is the point-in-time-recovery and audit trail behind the authoring plane.
[**CloudGetV1GuideCurriculum**](GuideAPI.md#CloudGetV1GuideCurriculum) | **Get** /v1/guide/curriculum | GetCurriculum returns the journey the caller&#39;s org is actually running, and whether it comes from the org&#39;s OWN override (custom) or from the platform default — the brand blueprint, else the embedded fixture.
[**CloudGetV1GuideProfile**](GuideAPI.md#CloudGetV1GuideProfile) | **Get** /v1/guide/profile | Profile returns the caller org&#39;s OBSERVED growth profile — the signal set, the classified growth stage, and the org&#39;s own key metrics.
[**CloudGetV1GuideStrategies**](GuideAPI.md#CloudGetV1GuideStrategies) | **Get** /v1/guide/strategies | Strategies returns the ENABLED tactics corpus for the caller&#39;s org: the tactics library narrowed by the explicit category/workload filters AND by the org&#39;s OBSERVED growth stage and capability signals (a tactic&#39;s tags are preconditions, so it surfaces only once the org can act on it).
[**CloudGetV1GuideSuggest**](GuideAPI.md#CloudGetV1GuideSuggest) | **Get** /v1/guide/suggest | Suggest returns the caller org&#39;s next-best quests: the available, non-terminal steps of its journey ranked by how much downstream work each unblocks, each with the grounded reason it is a good next move and whether the Business AI can run it, plus the org&#39;s funnel and the GTM recommendations derived from it.
[**CloudPatchV1GuideBlueprintByCollectionById**](GuideAPI.md#CloudPatchV1GuideBlueprintByCollectionById) | **Patch** /v1/guide/blueprint/{collection}/{id} | 
[**CloudPostV1GuideChat**](GuideAPI.md#CloudPostV1GuideChat) | **Post** /v1/guide/chat | Chat answers a founder&#39;s question about their launch journey as the Business AI coach: it grounds the reply in the org&#39;s REAL progress, its ranked available quests and its analytics funnel, and returns those candidate quests alongside so the caller can act on one.
[**CloudPostV1GuideStepsByIdDo**](GuideAPI.md#CloudPostV1GuideStepsByIdDo) | **Post** /v1/guide/steps/{id}/do | 
[**CloudPostV1GuideStepsByIdDone**](GuideAPI.md#CloudPostV1GuideStepsByIdDone) | **Post** /v1/guide/steps/{id}/done | 
[**CloudPostV1GuideStepsByIdStart**](GuideAPI.md#CloudPostV1GuideStepsByIdStart) | **Post** /v1/guide/steps/{id}/start | 
[**CloudPostV1GuideStepsIdReset**](GuideAPI.md#CloudPostV1GuideStepsIdReset) | **Post** /v1/guide/steps/{id}/reset | ResetStep returns one step of the caller org&#39;s journey to todo — clearing a manual mark or a skip — and returns the refreshed journey.
[**CloudPostV1GuideStepsIdSkip**](GuideAPI.md#CloudPostV1GuideStepsIdSkip) | **Post** /v1/guide/steps/{id}/skip | SkipStep marks one step of the caller org&#39;s journey skipped and returns the refreshed journey.
[**CloudPutV1GuideBlueprint**](GuideAPI.md#CloudPutV1GuideBlueprint) | **Put** /v1/guide/blueprint | 
[**CloudPutV1GuideCurriculum**](GuideAPI.md#CloudPutV1GuideCurriculum) | **Put** /v1/guide/curriculum | 



## CloudDeleteV1GuideCurriculum

> CloudCurriculumView CloudDeleteV1GuideCurriculum(ctx).Execute()

DeleteCurriculum clears the caller org's curriculum override and returns the journey it falls back to — the brand blueprint, else the embedded fixture.



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
	resp, r, err := apiClient.GuideAPI.CloudDeleteV1GuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudDeleteV1GuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1GuideCurriculum`: CloudCurriculumView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudDeleteV1GuideCurriculum`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1GuideCurriculumRequest struct via the builder pattern


### Return type

[**CloudCurriculumView**](CloudCurriculumView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Guide

> CloudOverviewView CloudGetV1Guide(ctx).Execute()

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
	resp, r, err := apiClient.GuideAPI.CloudGetV1Guide(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1Guide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Guide`: CloudOverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1Guide`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideRequest struct via the builder pattern


### Return type

[**CloudOverviewView**](CloudOverviewView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideActions

> CloudActionsView CloudGetV1GuideActions(ctx).Execute()

ListActions returns the caller org's Business AI action ledger, most recent first: every \"do it for me\" tool call, the arguments it ran with, its result and whether it succeeded.



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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideActions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideActions`: CloudActionsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideActions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideActionsRequest struct via the builder pattern


### Return type

[**CloudActionsView**](CloudActionsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideAnalytics

> CloudAnalyticsView CloudGetV1GuideAnalytics(ctx).Execute()

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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideAnalytics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideAnalytics`: CloudAnalyticsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideAnalytics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideAnalyticsRequest struct via the builder pattern


### Return type

[**CloudAnalyticsView**](CloudAnalyticsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideBlueprint

> CloudBlueprintView CloudGetV1GuideBlueprint(ctx).Execute()

GetBlueprint returns the FULL authored brand blueprint — every principle, section, step, strategy and template WITH its enabled flag made explicit, including the disabled items the org-facing reads never see — plus the active version number, the brand key it is stored under and the item counts.



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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideBlueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideBlueprint`: CloudBlueprintView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideBlueprint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideBlueprintRequest struct via the builder pattern


### Return type

[**CloudBlueprintView**](CloudBlueprintView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideBlueprintVersions

> CloudBlueprintVersionsView CloudGetV1GuideBlueprintVersions(ctx).Execute()

ListBlueprintVersions returns the brand blueprint's version history — every stored version's number and edit time, newest first — which is the point-in-time-recovery and audit trail behind the authoring plane.



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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideBlueprintVersions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideBlueprintVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideBlueprintVersions`: CloudBlueprintVersionsView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideBlueprintVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideBlueprintVersionsRequest struct via the builder pattern


### Return type

[**CloudBlueprintVersionsView**](CloudBlueprintVersionsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideCurriculum

> CloudCurriculumView CloudGetV1GuideCurriculum(ctx).Execute()

GetCurriculum returns the journey the caller's org is actually running, and whether it comes from the org's OWN override (custom) or from the platform default — the brand blueprint, else the embedded fixture.



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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideCurriculum`: CloudCurriculumView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideCurriculum`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideCurriculumRequest struct via the builder pattern


### Return type

[**CloudCurriculumView**](CloudCurriculumView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideProfile

> CloudProfileResponse CloudGetV1GuideProfile(ctx).Execute()

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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideProfile`: CloudProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideProfileRequest struct via the builder pattern


### Return type

[**CloudProfileResponse**](CloudProfileResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideStrategies

> CloudCorpusView CloudGetV1GuideStrategies(ctx).Category(category).Stage(stage).Workload(workload).Execute()

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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideStrategies(context.Background()).Category(category).Stage(stage).Workload(workload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideStrategies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideStrategies`: CloudCorpusView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideStrategies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Category filters to tactics in exactly this category. | 
 **stage** | **string** | Stage previews the corpus at a chosen growth stage (research|formed|launched|activated|scaling), overriding the org&#39;s observed one. An unknown value is ignored and the observed stage stands. | 
 **workload** | **string** | Workload filters to tactics with exactly this workload. | 

### Return type

[**CloudCorpusView**](CloudCorpusView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1GuideSuggest

> CloudSuggestResponse CloudGetV1GuideSuggest(ctx).Execute()

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
	resp, r, err := apiClient.GuideAPI.CloudGetV1GuideSuggest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudGetV1GuideSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1GuideSuggest`: CloudSuggestResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudGetV1GuideSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1GuideSuggestRequest struct via the builder pattern


### Return type

[**CloudSuggestResponse**](CloudSuggestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1GuideBlueprintByCollectionById

> CloudPatchV1GuideBlueprintByCollectionById(ctx, collection, id).Execute()



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
	r, err := apiClient.GuideAPI.CloudPatchV1GuideBlueprintByCollectionById(context.Background(), collection, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPatchV1GuideBlueprintByCollectionById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPatchV1GuideBlueprintByCollectionByIdRequest struct via the builder pattern


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


## CloudPostV1GuideChat

> CloudChatResponse CloudPostV1GuideChat(ctx).CloudChatRequest(cloudChatRequest).Execute()

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
	cloudChatRequest := *openapiclient.NewCloudChatRequest() // CloudChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuideAPI.CloudPostV1GuideChat(context.Background()).CloudChatRequest(cloudChatRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GuideChat`: CloudChatResponse
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudPostV1GuideChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GuideChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudChatRequest** | [**CloudChatRequest**](CloudChatRequest.md) |  | 

### Return type

[**CloudChatResponse**](CloudChatResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GuideStepsByIdDo

> CloudPostV1GuideStepsByIdDo(ctx, id).Execute()



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
	r, err := apiClient.GuideAPI.CloudPostV1GuideStepsByIdDo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideStepsByIdDo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GuideStepsByIdDoRequest struct via the builder pattern


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


## CloudPostV1GuideStepsByIdDone

> CloudPostV1GuideStepsByIdDone(ctx, id).Execute()



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
	r, err := apiClient.GuideAPI.CloudPostV1GuideStepsByIdDone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideStepsByIdDone``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GuideStepsByIdDoneRequest struct via the builder pattern


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


## CloudPostV1GuideStepsByIdStart

> CloudPostV1GuideStepsByIdStart(ctx, id).Execute()



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
	r, err := apiClient.GuideAPI.CloudPostV1GuideStepsByIdStart(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideStepsByIdStart``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1GuideStepsByIdStartRequest struct via the builder pattern


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


## CloudPostV1GuideStepsIdReset

> CloudOverviewView CloudPostV1GuideStepsIdReset(ctx, id).Execute()

ResetStep returns one step of the caller org's journey to todo — clearing a manual mark or a skip — and returns the refreshed journey.



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
	resp, r, err := apiClient.GuideAPI.CloudPostV1GuideStepsIdReset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideStepsIdReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GuideStepsIdReset`: CloudOverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudPostV1GuideStepsIdReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the step&#39;s id, as it appears in the journey (e.g. \&quot;gsuite\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GuideStepsIdResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudOverviewView**](CloudOverviewView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1GuideStepsIdSkip

> CloudOverviewView CloudPostV1GuideStepsIdSkip(ctx, id).Execute()

SkipStep marks one step of the caller org's journey skipped and returns the refreshed journey.



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
	resp, r, err := apiClient.GuideAPI.CloudPostV1GuideStepsIdSkip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPostV1GuideStepsIdSkip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1GuideStepsIdSkip`: CloudOverviewView
	fmt.Fprintf(os.Stdout, "Response from `GuideAPI.CloudPostV1GuideStepsIdSkip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the step&#39;s id, as it appears in the journey (e.g. \&quot;gsuite\&quot;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1GuideStepsIdSkipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudOverviewView**](CloudOverviewView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1GuideBlueprint

> CloudPutV1GuideBlueprint(ctx).Execute()



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
	r, err := apiClient.GuideAPI.CloudPutV1GuideBlueprint(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPutV1GuideBlueprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1GuideBlueprintRequest struct via the builder pattern


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


## CloudPutV1GuideCurriculum

> CloudPutV1GuideCurriculum(ctx).Execute()



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
	r, err := apiClient.GuideAPI.CloudPutV1GuideCurriculum(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuideAPI.CloudPutV1GuideCurriculum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1GuideCurriculumRequest struct via the builder pattern


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

