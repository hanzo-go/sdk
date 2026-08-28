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

> SeoAuditOut SeoAudit(ctx).SeoAuditIn(seoAuditIn).Execute()

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
	seoAuditIn := *openapiclient.NewSeoAuditIn() // SeoAuditIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoAudit(context.Background()).SeoAuditIn(seoAuditIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoAudit`: SeoAuditOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoAudit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoAuditIn** | [**SeoAuditIn**](SeoAuditIn.md) |  | 

### Return type

[**SeoAuditOut**](SeoAuditOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoBacklink

> SeoBacklinkOut SeoBacklink(ctx).SeoBacklinkIn(seoBacklinkIn).Execute()

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
	seoBacklinkIn := *openapiclient.NewSeoBacklinkIn() // SeoBacklinkIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoBacklink(context.Background()).SeoBacklinkIn(seoBacklinkIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoBacklink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoBacklink`: SeoBacklinkOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoBacklink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoBacklinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoBacklinkIn** | [**SeoBacklinkIn**](SeoBacklinkIn.md) |  | 

### Return type

[**SeoBacklinkOut**](SeoBacklinkOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoCompetitor

> SeoCompetitorOut SeoCompetitor(ctx).SeoCompetitorIn(seoCompetitorIn).Execute()

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
	seoCompetitorIn := *openapiclient.NewSeoCompetitorIn() // SeoCompetitorIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoCompetitor(context.Background()).SeoCompetitorIn(seoCompetitorIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoCompetitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoCompetitor`: SeoCompetitorOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoCompetitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoCompetitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoCompetitorIn** | [**SeoCompetitorIn**](SeoCompetitorIn.md) |  | 

### Return type

[**SeoCompetitorOut**](SeoCompetitorOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoIdea

> SeoIdeaOut SeoIdea(ctx).SeoIdeaIn(seoIdeaIn).Execute()

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
	seoIdeaIn := *openapiclient.NewSeoIdeaIn() // SeoIdeaIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoIdea(context.Background()).SeoIdeaIn(seoIdeaIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoIdea``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoIdea`: SeoIdeaOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoIdea`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoIdeaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoIdeaIn** | [**SeoIdeaIn**](SeoIdeaIn.md) |  | 

### Return type

[**SeoIdeaOut**](SeoIdeaOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoKeyword

> SeoKeywordOut SeoKeyword(ctx).SeoKeywordIn(seoKeywordIn).Execute()

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
	seoKeywordIn := *openapiclient.NewSeoKeywordIn() // SeoKeywordIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoKeyword(context.Background()).SeoKeywordIn(seoKeywordIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoKeyword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoKeyword`: SeoKeywordOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoKeyword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoKeywordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoKeywordIn** | [**SeoKeywordIn**](SeoKeywordIn.md) |  | 

### Return type

[**SeoKeywordOut**](SeoKeywordOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoRank

> SeoRankOut SeoRank(ctx).SeoRankIn(seoRankIn).Execute()

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
	seoRankIn := *openapiclient.NewSeoRankIn() // SeoRankIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SeoAPI.SeoRank(context.Background()).SeoRankIn(seoRankIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SeoAPI.SeoRank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SeoRank`: SeoRankOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoRank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSeoRankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seoRankIn** | [**SeoRankIn**](SeoRankIn.md) |  | 

### Return type

[**SeoRankOut**](SeoRankOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SeoRate

> SeoRateOut SeoRate(ctx).Execute()

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
	// response from `SeoRate`: SeoRateOut
	fmt.Fprintf(os.Stdout, "Response from `SeoAPI.SeoRate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSeoRateRequest struct via the builder pattern


### Return type

[**SeoRateOut**](SeoRateOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

