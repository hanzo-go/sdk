# \SeoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SeoAudit**](SeoAPI.md#SeoAudit) | **Post** /v1/seo/audit | Fetch one page and report what it gets wrong
[**SeoBacklink**](SeoAPI.md#SeoBacklink) | **Post** /v1/seo/backlinks | Who links to a target, and how much of it is broken or spam
[**SeoCompetitor**](SeoAPI.md#SeoCompetitor) | **Post** /v1/seo/competitors | The domains that place for the same phrases
[**SeoIdea**](SeoAPI.md#SeoIdea) | **Post** /v1/seo/ideas | Grow a seed phrase into the phrases nobody named yet
[**SeoKeyword**](SeoAPI.md#SeoKeyword) | **Post** /v1/seo/keywords | How often named phrases are searched, and what a click costs
[**SeoRank**](SeoAPI.md#SeoRank) | **Post** /v1/seo/rankings | Every phrase a domain already places for, with its position
[**SeoRate**](SeoAPI.md#SeoRate) | **Get** /v1/seo/rates | What each call on this surface costs, from the vendor&#39;s own list



## SeoAudit

> SeoSeoAuditOut SeoAudit(ctx).SeoSeoAuditIn(seoSeoAuditIn).Execute()

Fetch one page and report what it gets wrong



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoAuditIn := *openapiclient.NewSeoSeoAuditIn() // SeoSeoAuditIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoAudit(context.Background()).SeoSeoAuditIn(seoSeoAuditIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoAudit`: SeoSeoAuditOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoAuditIn** | [**SeoSeoAuditIn**](SeoSeoAuditIn.md) |  | 

### Return type

[**SeoSeoAuditOut**](SeoSeoAuditOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoBacklink

> SeoSeoBacklinkOut SeoBacklink(ctx).SeoSeoBacklinkIn(seoSeoBacklinkIn).Execute()

Who links to a target, and how much of it is broken or spam



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoBacklinkIn := *openapiclient.NewSeoSeoBacklinkIn() // SeoSeoBacklinkIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoBacklink(context.Background()).SeoSeoBacklinkIn(seoSeoBacklinkIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoBacklink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoBacklink`: SeoSeoBacklinkOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoBacklink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoBacklinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoBacklinkIn** | [**SeoSeoBacklinkIn**](SeoSeoBacklinkIn.md) |  | 

### Return type

[**SeoSeoBacklinkOut**](SeoSeoBacklinkOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoCompetitor

> SeoSeoCompetitorOut SeoCompetitor(ctx).SeoSeoCompetitorIn(seoSeoCompetitorIn).Execute()

The domains that place for the same phrases



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoCompetitorIn := *openapiclient.NewSeoSeoCompetitorIn() // SeoSeoCompetitorIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoCompetitor(context.Background()).SeoSeoCompetitorIn(seoSeoCompetitorIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoCompetitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoCompetitor`: SeoSeoCompetitorOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoCompetitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoCompetitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoCompetitorIn** | [**SeoSeoCompetitorIn**](SeoSeoCompetitorIn.md) |  | 

### Return type

[**SeoSeoCompetitorOut**](SeoSeoCompetitorOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoIdea

> SeoSeoIdeaOut SeoIdea(ctx).SeoSeoIdeaIn(seoSeoIdeaIn).Execute()

Grow a seed phrase into the phrases nobody named yet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoIdeaIn := *openapiclient.NewSeoSeoIdeaIn() // SeoSeoIdeaIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoIdea(context.Background()).SeoSeoIdeaIn(seoSeoIdeaIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoIdea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoIdea`: SeoSeoIdeaOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoIdea`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoIdeaIn** | [**SeoSeoIdeaIn**](SeoSeoIdeaIn.md) |  | 

### Return type

[**SeoSeoIdeaOut**](SeoSeoIdeaOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoKeyword

> SeoSeoKeywordOut SeoKeyword(ctx).SeoSeoKeywordIn(seoSeoKeywordIn).Execute()

How often named phrases are searched, and what a click costs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoKeywordIn := *openapiclient.NewSeoSeoKeywordIn() // SeoSeoKeywordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoKeyword(context.Background()).SeoSeoKeywordIn(seoSeoKeywordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoKeyword`: SeoSeoKeywordOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoKeyword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoKeywordIn** | [**SeoSeoKeywordIn**](SeoSeoKeywordIn.md) |  | 

### Return type

[**SeoSeoKeywordOut**](SeoSeoKeywordOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoRank

> SeoSeoRankOut SeoRank(ctx).SeoSeoRankIn(seoSeoRankIn).Execute()

Every phrase a domain already places for, with its position



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	seoSeoRankIn := *openapiclient.NewSeoSeoRankIn() // SeoSeoRankIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoRank(context.Background()).SeoSeoRankIn(seoSeoRankIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoRank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoRank`: SeoSeoRankOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoRank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoRankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoSeoRankIn** | [**SeoSeoRankIn**](SeoSeoRankIn.md) |  | 

### Return type

[**SeoSeoRankOut**](SeoSeoRankOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoRate

> SeoSeoRateOut SeoRate(ctx).Execute()

What each call on this surface costs, from the vendor's own list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoRate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoRate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoRate`: SeoSeoRateOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoRate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSeoRateRequest struct via the builder pattern


### Return type

[**SeoSeoRateOut**](SeoSeoRateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

