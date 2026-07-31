# \AffiliatesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Affiliates**](AffiliatesAPI.md#CloudGetV1Affiliates) | **Get** /v1/affiliates | 
[**CloudGetV1AffiliatesLeaderboard**](AffiliatesAPI.md#CloudGetV1AffiliatesLeaderboard) | **Get** /v1/affiliates/leaderboard | 
[**CloudGetV1AffiliatesMe**](AffiliatesAPI.md#CloudGetV1AffiliatesMe) | **Get** /v1/affiliates/me | 
[**CloudGetV1AffiliatesMeEarnings**](AffiliatesAPI.md#CloudGetV1AffiliatesMeEarnings) | **Get** /v1/affiliates/me/earnings | 
[**CloudGetV1AffiliatesMeLinks**](AffiliatesAPI.md#CloudGetV1AffiliatesMeLinks) | **Get** /v1/affiliates/me/links | 
[**CloudPostV1AffiliatesApply**](AffiliatesAPI.md#CloudPostV1AffiliatesApply) | **Post** /v1/affiliates/apply | 
[**CloudPostV1AffiliatesAttribute**](AffiliatesAPI.md#CloudPostV1AffiliatesAttribute) | **Post** /v1/affiliates/attribute | 
[**CloudPostV1AffiliatesClick**](AffiliatesAPI.md#CloudPostV1AffiliatesClick) | **Post** /v1/affiliates/click | 
[**CloudPostV1AffiliatesMeHandle**](AffiliatesAPI.md#CloudPostV1AffiliatesMeHandle) | **Post** /v1/affiliates/me/handle | 
[**CloudPostV1AffiliatesMeLinks**](AffiliatesAPI.md#CloudPostV1AffiliatesMeLinks) | **Post** /v1/affiliates/me/links | 
[**CommerceConnectAffiliate**](AffiliatesAPI.md#CommerceConnectAffiliate) | **Get** /v1/commerce/affiliate/{affiliateid}/connect | Connect affiliate
[**CommerceCreateAffiliate**](AffiliatesAPI.md#CommerceCreateAffiliate) | **Post** /v1/commerce/affiliate | Create affiliate
[**CommerceGetAffiliate**](AffiliatesAPI.md#CommerceGetAffiliate) | **Get** /v1/commerce/affiliate/{affiliateid} | Get affiliate
[**CommerceListAffiliates**](AffiliatesAPI.md#CommerceListAffiliates) | **Get** /v1/commerce/affiliate | List affiliates



## CloudGetV1Affiliates

> CloudGetV1Affiliates(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudGetV1Affiliates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudGetV1Affiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AffiliatesRequest struct via the builder pattern


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


## CloudGetV1AffiliatesLeaderboard

> CloudGetV1AffiliatesLeaderboard(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudGetV1AffiliatesLeaderboard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudGetV1AffiliatesLeaderboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AffiliatesLeaderboardRequest struct via the builder pattern


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


## CloudGetV1AffiliatesMe

> CloudGetV1AffiliatesMe(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudGetV1AffiliatesMe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudGetV1AffiliatesMe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AffiliatesMeRequest struct via the builder pattern


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


## CloudGetV1AffiliatesMeEarnings

> CloudGetV1AffiliatesMeEarnings(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudGetV1AffiliatesMeEarnings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudGetV1AffiliatesMeEarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AffiliatesMeEarningsRequest struct via the builder pattern


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


## CloudGetV1AffiliatesMeLinks

> CloudGetV1AffiliatesMeLinks(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudGetV1AffiliatesMeLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudGetV1AffiliatesMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AffiliatesMeLinksRequest struct via the builder pattern


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


## CloudPostV1AffiliatesApply

> CloudPostV1AffiliatesApply(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudPostV1AffiliatesApply(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudPostV1AffiliatesApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AffiliatesApplyRequest struct via the builder pattern


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


## CloudPostV1AffiliatesAttribute

> CloudPostV1AffiliatesAttribute(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudPostV1AffiliatesAttribute(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudPostV1AffiliatesAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AffiliatesAttributeRequest struct via the builder pattern


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


## CloudPostV1AffiliatesClick

> CloudPostV1AffiliatesClick(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudPostV1AffiliatesClick(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudPostV1AffiliatesClick``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AffiliatesClickRequest struct via the builder pattern


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


## CloudPostV1AffiliatesMeHandle

> CloudPostV1AffiliatesMeHandle(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudPostV1AffiliatesMeHandle(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudPostV1AffiliatesMeHandle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AffiliatesMeHandleRequest struct via the builder pattern


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


## CloudPostV1AffiliatesMeLinks

> CloudPostV1AffiliatesMeLinks(ctx).Execute()



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
	r, err := apiClient.AffiliatesAPI.CloudPostV1AffiliatesMeLinks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CloudPostV1AffiliatesMeLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AffiliatesMeLinksRequest struct via the builder pattern


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


## CommerceConnectAffiliate

> map[string]interface{} CommerceConnectAffiliate(ctx, affiliateid).Execute()

Connect affiliate

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
	affiliateid := "affiliateid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.CommerceConnectAffiliate(context.Background(), affiliateid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CommerceConnectAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceConnectAffiliate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.CommerceConnectAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**affiliateid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceConnectAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CommerceCreateAffiliate

> CommerceAffiliate CommerceCreateAffiliate(ctx).CommerceAffiliate(commerceAffiliate).Execute()

Create affiliate

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
	commerceAffiliate := *openapiclient.NewCommerceAffiliate() // CommerceAffiliate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.CommerceCreateAffiliate(context.Background()).CommerceAffiliate(commerceAffiliate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CommerceCreateAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceCreateAffiliate`: CommerceAffiliate
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.CommerceCreateAffiliate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceCreateAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceAffiliate** | [**CommerceAffiliate**](CommerceAffiliate.md) |  | 

### Return type

[**CommerceAffiliate**](CommerceAffiliate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceGetAffiliate

> CommerceAffiliate CommerceGetAffiliate(ctx, affiliateid).Execute()

Get affiliate

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
	affiliateid := "affiliateid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.CommerceGetAffiliate(context.Background(), affiliateid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CommerceGetAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceGetAffiliate`: CommerceAffiliate
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.CommerceGetAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**affiliateid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommerceGetAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommerceAffiliate**](CommerceAffiliate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceListAffiliates

> CommercePaginatedAffiliates CommerceListAffiliates(ctx).Page(page).Display(display).Execute()

List affiliates

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
	page := int32(56) // int32 | Page number (1-indexed) (optional) (default to 1)
	display := int32(56) // int32 | Number of items per page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AffiliatesAPI.CommerceListAffiliates(context.Background()).Page(page).Display(display).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AffiliatesAPI.CommerceListAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceListAffiliates`: CommercePaginatedAffiliates
	fmt.Fprintf(os.Stdout, "Response from `AffiliatesAPI.CommerceListAffiliates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceListAffiliatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed) | [default to 1]
 **display** | **int32** | Number of items per page | [default to 20]

### Return type

[**CommercePaginatedAffiliates**](CommercePaginatedAffiliates.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

