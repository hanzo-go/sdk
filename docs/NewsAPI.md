# \NewsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldArxiv**](NewsAPI.md#WorldWorldArxiv) | **Get** /v1/world/arxiv | arXiv research feed (backs Robotics/Quantum lenses)
[**WorldWorldGdeltDoc**](NewsAPI.md#WorldWorldGdeltDoc) | **Get** /v1/world/gdelt-doc | GDELT article search
[**WorldWorldGdeltGeo**](NewsAPI.md#WorldWorldGdeltGeo) | **Get** /v1/world/gdelt-geo | GDELT geo-tagged events
[**WorldWorldGithubTrending**](NewsAPI.md#WorldWorldGithubTrending) | **Get** /v1/world/github-trending | GitHub trending repositories
[**WorldWorldHackernews**](NewsAPI.md#WorldWorldHackernews) | **Get** /v1/world/hackernews | Hacker News stories
[**WorldWorldOgStory**](NewsAPI.md#WorldWorldOgStory) | **Get** /v1/world/og-story | Open-graph story card
[**WorldWorldRssProxy**](NewsAPI.md#WorldWorldRssProxy) | **Get** /v1/world/rss-proxy | Allowlisted RSS feed proxy (SSRF-bounded)
[**WorldWorldStory**](NewsAPI.md#WorldWorldStory) | **Get** /v1/world/story | Story detail
[**WorldWorldTechEvents**](NewsAPI.md#WorldWorldTechEvents) | **Get** /v1/world/tech-events | Curated technology events



## WorldWorldArxiv

> map[string]interface{} WorldWorldArxiv(ctx).Q(q).Cat(cat).Execute()

arXiv research feed (backs Robotics/Quantum lenses)

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
	q := "q_example" // string |  (optional)
	cat := "cat_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.WorldWorldArxiv(context.Background()).Q(q).Cat(cat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldArxiv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldArxiv`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldArxiv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldArxivRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **cat** | **string** |  | 

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


## WorldWorldGdeltDoc

> map[string]interface{} WorldWorldGdeltDoc(ctx).Query(query).Maxrecords(maxrecords).Timespan(timespan).Execute()

GDELT article search

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
	query := "query_example" // string | 
	maxrecords := "maxrecords_example" // string |  (optional)
	timespan := "timespan_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.WorldWorldGdeltDoc(context.Background()).Query(query).Maxrecords(maxrecords).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldGdeltDoc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldGdeltDoc`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldGdeltDoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldGdeltDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **maxrecords** | **string** |  | 
 **timespan** | **string** |  | 

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


## WorldWorldGdeltGeo

> map[string]interface{} WorldWorldGdeltGeo(ctx).Query(query).Format(format).Maxrecords(maxrecords).Timespan(timespan).Execute()

GDELT geo-tagged events

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
	query := "query_example" // string | 
	format := "format_example" // string |  (optional)
	maxrecords := "maxrecords_example" // string |  (optional)
	timespan := "timespan_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.WorldWorldGdeltGeo(context.Background()).Query(query).Format(format).Maxrecords(maxrecords).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldGdeltGeo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldGdeltGeo`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldGdeltGeo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldGdeltGeoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **format** | **string** |  | 
 **maxrecords** | **string** |  | 
 **timespan** | **string** |  | 

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


## WorldWorldGithubTrending

> map[string]interface{} WorldWorldGithubTrending(ctx).Execute()

GitHub trending repositories

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
	resp, r, err := apiClient.NewsAPI.WorldWorldGithubTrending(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldGithubTrending``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldGithubTrending`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldGithubTrending`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldGithubTrendingRequest struct via the builder pattern


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


## WorldWorldHackernews

> map[string]interface{} WorldWorldHackernews(ctx).Type_(type_).Execute()

Hacker News stories

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
	type_ := "type__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.WorldWorldHackernews(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldHackernews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldHackernews`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldHackernews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldHackernewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 

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


## WorldWorldOgStory

> map[string]interface{} WorldWorldOgStory(ctx).Execute()

Open-graph story card

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
	resp, r, err := apiClient.NewsAPI.WorldWorldOgStory(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldOgStory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldOgStory`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldOgStory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldOgStoryRequest struct via the builder pattern


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


## WorldWorldRssProxy

> map[string]interface{} WorldWorldRssProxy(ctx).Url(url).Execute()

Allowlisted RSS feed proxy (SSRF-bounded)

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
	url := "url_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsAPI.WorldWorldRssProxy(context.Background()).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldRssProxy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldRssProxy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldRssProxy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldRssProxyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **string** |  | 

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


## WorldWorldStory

> map[string]interface{} WorldWorldStory(ctx).Execute()

Story detail

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
	resp, r, err := apiClient.NewsAPI.WorldWorldStory(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldStory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldStory`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldStory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldStoryRequest struct via the builder pattern


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


## WorldWorldTechEvents

> map[string]interface{} WorldWorldTechEvents(ctx).Execute()

Curated technology events

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
	resp, r, err := apiClient.NewsAPI.WorldWorldTechEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsAPI.WorldWorldTechEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldTechEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NewsAPI.WorldWorldTechEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldTechEventsRequest struct via the builder pattern


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

