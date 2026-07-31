# \AdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AffiliatesAdminApproveAffiliate**](AdminAPI.md#AffiliatesAdminApproveAffiliate) | **Post** /v1/admin/affiliates/{id}/approve | Approve an affiliate and mint its code
[**AffiliatesAdminListAffiliates**](AdminAPI.md#AffiliatesAdminListAffiliates) | **Get** /v1/admin/affiliates | List all affiliates
[**AffiliatesAdminPayoutAffiliate**](AdminAPI.md#AffiliatesAdminPayoutAffiliate) | **Post** /v1/admin/affiliates/{id}/payout | Record a payout
[**AffiliatesAdminSuspendAffiliate**](AdminAPI.md#AffiliatesAdminSuspendAffiliate) | **Post** /v1/admin/affiliates/{id}/suspend | Suspend an affiliate
[**AffiliatesAdminSweepAffiliates**](AdminAPI.md#AffiliatesAdminSweepAffiliates) | **Post** /v1/admin/affiliates/sweep | Run the accrual sweep
[**AnalyticsAdminListUsers**](AdminAPI.md#AnalyticsAdminListUsers) | **Get** /v1/analytics/admin/users | List all users (admin only)
[**AnalyticsAdminListWebsites**](AdminAPI.md#AnalyticsAdminListWebsites) | **Get** /v1/analytics/admin/websites | List all websites for a user (admin only)
[**AuthorsAdminApproveAuthor**](AdminAPI.md#AuthorsAdminApproveAuthor) | **Post** /v1/admin/authors/{id}/approve | Approve an author
[**AuthorsAdminListAuthors**](AdminAPI.md#AuthorsAdminListAuthors) | **Get** /v1/admin/authors | List all authors
[**AuthorsAdminPayoutAuthor**](AdminAPI.md#AuthorsAdminPayoutAuthor) | **Post** /v1/admin/authors/{id}/payout | Record a payout
[**AuthorsAdminSuspendAuthor**](AdminAPI.md#AuthorsAdminSuspendAuthor) | **Post** /v1/admin/authors/{id}/suspend | Suspend an author
[**AuthorsAdminSweepAuthors**](AdminAPI.md#AuthorsAdminSweepAuthors) | **Post** /v1/admin/authors/sweep | Run accrual sweep
[**PluginAdminDisablePlugin**](AdminAPI.md#PluginAdminDisablePlugin) | **Post** /v1/admin/plugins/{name}/disable | Disable a plugin — its routes answer 503, not 404
[**PluginAdminEnablePlugin**](AdminAPI.md#PluginAdminEnablePlugin) | **Post** /v1/admin/plugins/{name}/enable | Enable a stopped or disabled plugin
[**PluginAdminPlugins**](AdminAPI.md#PluginAdminPlugins) | **Get** /v1/admin/plugins | List what each host is running
[**PluginAdminReloadPlugin**](AdminAPI.md#PluginAdminReloadPlugin) | **Post** /v1/admin/plugins/{name}/reload | Reload a plugin onto another build
[**ReferralsAdminListReferrals**](AdminAPI.md#ReferralsAdminListReferrals) | **Get** /v1/admin/referrals | List every referral with a fleet summary (global-admin)
[**ReferralsAdminSweepReferrals**](AdminAPI.md#ReferralsAdminSweepReferrals) | **Post** /v1/admin/referrals/sweep | Qualify-check every pending referral (global-admin)
[**S3AdminInfo**](AdminAPI.md#S3AdminInfo) | **Get** /v1/s3/admin/info | Server information
[**S3AdminUsage**](AdminAPI.md#S3AdminUsage) | **Get** /v1/s3/admin/usage | Storage usage
[**S3CreateServiceAccount**](AdminAPI.md#S3CreateServiceAccount) | **Post** /v1/s3/admin/service-accounts | Create a service account
[**S3ListServiceAccounts**](AdminAPI.md#S3ListServiceAccounts) | **Get** /v1/s3/admin/service-accounts | List service accounts



## AffiliatesAdminApproveAffiliate

> AffiliatesAdminAffiliateEnvelope AffiliatesAdminApproveAffiliate(ctx, id).AffiliatesApproveRequest(affiliatesApproveRequest).Execute()

Approve an affiliate and mint its code



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
	id := "id_example" // string | The affiliate id (e.g. `aff_<hex>`).
	affiliatesApproveRequest := *openapiclient.NewAffiliatesApproveRequest() // AffiliatesApproveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AffiliatesAdminApproveAffiliate(context.Background(), id).AffiliatesApproveRequest(affiliatesApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AffiliatesAdminApproveAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAdminApproveAffiliate`: AffiliatesAdminAffiliateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AffiliatesAdminApproveAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The affiliate id (e.g. &#x60;aff_&lt;hex&gt;&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAdminApproveAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **affiliatesApproveRequest** | [**AffiliatesApproveRequest**](AffiliatesApproveRequest.md) |  | 

### Return type

[**AffiliatesAdminAffiliateEnvelope**](AffiliatesAdminAffiliateEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesAdminListAffiliates

> AffiliatesAdminListEnvelope AffiliatesAdminListAffiliates(ctx).Limit(limit).Execute()

List all affiliates



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
	limit := int32(56) // int32 | Max rows to return (default 500, max 1000). (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AffiliatesAdminListAffiliates(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AffiliatesAdminListAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAdminListAffiliates`: AffiliatesAdminListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AffiliatesAdminListAffiliates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAdminListAffiliatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max rows to return (default 500, max 1000). | [default to 500]

### Return type

[**AffiliatesAdminListEnvelope**](AffiliatesAdminListEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesAdminPayoutAffiliate

> AffiliatesAdminPayoutEnvelope AffiliatesAdminPayoutAffiliate(ctx, id).AffiliatesPayoutRequest(affiliatesPayoutRequest).Execute()

Record a payout



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
	id := "id_example" // string | The affiliate id (e.g. `aff_<hex>`).
	affiliatesPayoutRequest := *openapiclient.NewAffiliatesPayoutRequest(int64(123), "credits") // AffiliatesPayoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AffiliatesAdminPayoutAffiliate(context.Background(), id).AffiliatesPayoutRequest(affiliatesPayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AffiliatesAdminPayoutAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAdminPayoutAffiliate`: AffiliatesAdminPayoutEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AffiliatesAdminPayoutAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The affiliate id (e.g. &#x60;aff_&lt;hex&gt;&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAdminPayoutAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **affiliatesPayoutRequest** | [**AffiliatesPayoutRequest**](AffiliatesPayoutRequest.md) |  | 

### Return type

[**AffiliatesAdminPayoutEnvelope**](AffiliatesAdminPayoutEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesAdminSuspendAffiliate

> AffiliatesAdminAffiliateEnvelope AffiliatesAdminSuspendAffiliate(ctx, id).Execute()

Suspend an affiliate



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
	id := "id_example" // string | The affiliate id (e.g. `aff_<hex>`).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AffiliatesAdminSuspendAffiliate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AffiliatesAdminSuspendAffiliate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAdminSuspendAffiliate`: AffiliatesAdminAffiliateEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AffiliatesAdminSuspendAffiliate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The affiliate id (e.g. &#x60;aff_&lt;hex&gt;&#x60;). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAdminSuspendAffiliateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AffiliatesAdminAffiliateEnvelope**](AffiliatesAdminAffiliateEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AffiliatesAdminSweepAffiliates

> AffiliatesAdminSweepEnvelope AffiliatesAdminSweepAffiliates(ctx).Execute()

Run the accrual sweep



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
	resp, r, err := apiClient.AdminAPI.AffiliatesAdminSweepAffiliates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AffiliatesAdminSweepAffiliates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AffiliatesAdminSweepAffiliates`: AffiliatesAdminSweepEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AffiliatesAdminSweepAffiliates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAffiliatesAdminSweepAffiliatesRequest struct via the builder pattern


### Return type

[**AffiliatesAdminSweepEnvelope**](AffiliatesAdminSweepEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsAdminListUsers

> []AnalyticsAdminListUsers200ResponseInner AnalyticsAdminListUsers(ctx).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all users (admin only)

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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AnalyticsAdminListUsers(context.Background()).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AnalyticsAdminListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListUsers`: []AnalyticsAdminListUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AnalyticsAdminListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsAdminListUsers200ResponseInner**](AnalyticsAdminListUsers200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalyticsAdminListWebsites

> []AnalyticsWebsite AnalyticsAdminListWebsites(ctx).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()

List all websites for a user (admin only)

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	includeOwnedTeams := "includeOwnedTeams_example" // string |  (optional)
	includeAllTeams := "includeAllTeams_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AnalyticsAdminListWebsites(context.Background()).UserId(userId).IncludeOwnedTeams(includeOwnedTeams).IncludeAllTeams(includeAllTeams).Page(page).PageSize(pageSize).OrderBy(orderBy).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AnalyticsAdminListWebsites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyticsAdminListWebsites`: []AnalyticsWebsite
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AnalyticsAdminListWebsites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsAdminListWebsitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **includeOwnedTeams** | **string** |  | 
 **includeAllTeams** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **orderBy** | **string** |  | 
 **search** | **string** |  | 

### Return type

[**[]AnalyticsWebsite**](AnalyticsWebsite.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsAdminApproveAuthor

> AuthorsAdminAuthorEnvelope AuthorsAdminApproveAuthor(ctx, id).AuthorsApproveRequest(authorsApproveRequest).Execute()

Approve an author



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
	id := "id_example" // string | Author id (e.g. aut_...).
	authorsApproveRequest := *openapiclient.NewAuthorsApproveRequest() // AuthorsApproveRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AuthorsAdminApproveAuthor(context.Background(), id).AuthorsApproveRequest(authorsApproveRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AuthorsAdminApproveAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsAdminApproveAuthor`: AuthorsAdminAuthorEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AuthorsAdminApproveAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Author id (e.g. aut_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsAdminApproveAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorsApproveRequest** | [**AuthorsApproveRequest**](AuthorsApproveRequest.md) |  | 

### Return type

[**AuthorsAdminAuthorEnvelope**](AuthorsAdminAuthorEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsAdminListAuthors

> AuthorsAdminListEnvelope AuthorsAdminListAuthors(ctx).Limit(limit).Execute()

List all authors



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
	limit := int32(56) // int32 | Max rows (default 500, capped at 1000). (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AuthorsAdminListAuthors(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AuthorsAdminListAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsAdminListAuthors`: AuthorsAdminListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AuthorsAdminListAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsAdminListAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max rows (default 500, capped at 1000). | [default to 500]

### Return type

[**AuthorsAdminListEnvelope**](AuthorsAdminListEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsAdminPayoutAuthor

> AuthorsAdminPayoutEnvelope AuthorsAdminPayoutAuthor(ctx, id).AuthorsPayoutRequest(authorsPayoutRequest).Execute()

Record a payout



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
	id := "id_example" // string | Author id (e.g. aut_...).
	authorsPayoutRequest := *openapiclient.NewAuthorsPayoutRequest(int64(123), "Method_example") // AuthorsPayoutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AuthorsAdminPayoutAuthor(context.Background(), id).AuthorsPayoutRequest(authorsPayoutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AuthorsAdminPayoutAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsAdminPayoutAuthor`: AuthorsAdminPayoutEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AuthorsAdminPayoutAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Author id (e.g. aut_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsAdminPayoutAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorsPayoutRequest** | [**AuthorsPayoutRequest**](AuthorsPayoutRequest.md) |  | 

### Return type

[**AuthorsAdminPayoutEnvelope**](AuthorsAdminPayoutEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsAdminSuspendAuthor

> AuthorsAdminAuthorEnvelope AuthorsAdminSuspendAuthor(ctx, id).Execute()

Suspend an author



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
	id := "id_example" // string | Author id (e.g. aut_...).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AuthorsAdminSuspendAuthor(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AuthorsAdminSuspendAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsAdminSuspendAuthor`: AuthorsAdminAuthorEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AuthorsAdminSuspendAuthor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Author id (e.g. aut_...). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsAdminSuspendAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorsAdminAuthorEnvelope**](AuthorsAdminAuthorEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsAdminSweepAuthors

> AuthorsAdminSweepEnvelope AuthorsAdminSweepAuthors(ctx).Execute()

Run accrual sweep



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
	resp, r, err := apiClient.AdminAPI.AuthorsAdminSweepAuthors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AuthorsAdminSweepAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsAdminSweepAuthors`: AuthorsAdminSweepEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AuthorsAdminSweepAuthors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsAdminSweepAuthorsRequest struct via the builder pattern


### Return type

[**AuthorsAdminSweepEnvelope**](AuthorsAdminSweepEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginAdminDisablePlugin

> PluginActionOut PluginAdminDisablePlugin(ctx, name).PluginNameIn(pluginNameIn).Execute()

Disable a plugin — its routes answer 503, not 404



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
	name := "name_example" // string | Name is the app, from the path.
	pluginNameIn := *openapiclient.NewPluginNameIn() // PluginNameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PluginAdminDisablePlugin(context.Background(), name).PluginNameIn(pluginNameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PluginAdminDisablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginAdminDisablePlugin`: PluginActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PluginAdminDisablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginAdminDisablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pluginNameIn** | [**PluginNameIn**](PluginNameIn.md) |  | 

### Return type

[**PluginActionOut**](PluginActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginAdminEnablePlugin

> PluginActionOut PluginAdminEnablePlugin(ctx, name).PluginNameIn(pluginNameIn).Execute()

Enable a stopped or disabled plugin



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
	name := "name_example" // string | Name is the app, from the path.
	pluginNameIn := *openapiclient.NewPluginNameIn() // PluginNameIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PluginAdminEnablePlugin(context.Background(), name).PluginNameIn(pluginNameIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PluginAdminEnablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginAdminEnablePlugin`: PluginActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PluginAdminEnablePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginAdminEnablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pluginNameIn** | [**PluginNameIn**](PluginNameIn.md) |  | 

### Return type

[**PluginActionOut**](PluginActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginAdminPlugins

> PluginListOut PluginAdminPlugins(ctx).Scope(scope).Execute()

List what each host is running



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
	scope := "scope_example" // string | Scope \"host\" answers for THIS host only. Default \"fleet\" fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PluginAdminPlugins(context.Background()).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PluginAdminPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginAdminPlugins`: PluginListOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PluginAdminPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPluginAdminPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope \&quot;host\&quot; answers for THIS host only. Default \&quot;fleet\&quot; fans out to every live peer. A peer answers a host-scoped read, which is what stops the fan-out recursing.  | 

### Return type

[**PluginListOut**](PluginListOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PluginAdminReloadPlugin

> PluginActionOut PluginAdminReloadPlugin(ctx, name).PluginReloadIn(pluginReloadIn).Execute()

Reload a plugin onto another build



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
	name := "name_example" // string | Name is the app, from the path. It must be one the manifest declares.
	pluginReloadIn := *openapiclient.NewPluginReloadIn() // PluginReloadIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.PluginAdminReloadPlugin(context.Background(), name).PluginReloadIn(pluginReloadIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.PluginAdminReloadPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PluginAdminReloadPlugin`: PluginActionOut
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.PluginAdminReloadPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the app, from the path. It must be one the manifest declares. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPluginAdminReloadPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pluginReloadIn** | [**PluginReloadIn**](PluginReloadIn.md) |  | 

### Return type

[**PluginActionOut**](PluginActionOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferralsAdminListReferrals

> ReferralsAdminListEnvelope ReferralsAdminListReferrals(ctx).Limit(limit).Execute()

List every referral with a fleet summary (global-admin)



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
	limit := int32(56) // int32 | Max rows to return. Defaults to 500 when absent/invalid/<=0; capped at 1000.  (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.ReferralsAdminListReferrals(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReferralsAdminListReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferralsAdminListReferrals`: ReferralsAdminListEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReferralsAdminListReferrals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReferralsAdminListReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max rows to return. Defaults to 500 when absent/invalid/&lt;&#x3D;0; capped at 1000.  | [default to 500]

### Return type

[**ReferralsAdminListEnvelope**](ReferralsAdminListEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferralsAdminSweepReferrals

> ReferralsAdminSweepEnvelope ReferralsAdminSweepReferrals(ctx).Execute()

Qualify-check every pending referral (global-admin)



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
	resp, r, err := apiClient.AdminAPI.ReferralsAdminSweepReferrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.ReferralsAdminSweepReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferralsAdminSweepReferrals`: ReferralsAdminSweepEnvelope
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.ReferralsAdminSweepReferrals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferralsAdminSweepReferralsRequest struct via the builder pattern


### Return type

[**ReferralsAdminSweepEnvelope**](ReferralsAdminSweepEnvelope.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3AdminInfo

> S3AdminInfo200Response S3AdminInfo(ctx).Execute()

Server information



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
	resp, r, err := apiClient.AdminAPI.S3AdminInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3AdminInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminInfo`: S3AdminInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3AdminInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminInfoRequest struct via the builder pattern


### Return type

[**S3AdminInfo200Response**](S3AdminInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3AdminUsage

> S3UsageInfo S3AdminUsage(ctx).Execute()

Storage usage



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
	resp, r, err := apiClient.AdminAPI.S3AdminUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3AdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminUsage`: S3UsageInfo
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3AdminUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminUsageRequest struct via the builder pattern


### Return type

[**S3UsageInfo**](S3UsageInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3CreateServiceAccount

> S3ServiceAccount S3CreateServiceAccount(ctx).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()

Create a service account



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
	s3CreateServiceAccountRequest := *openapiclient.NewS3CreateServiceAccountRequest() // S3CreateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.S3CreateServiceAccount(context.Background()).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3CreateServiceAccount`: S3ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3CreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3CreateServiceAccountRequest** | [**S3CreateServiceAccountRequest**](S3CreateServiceAccountRequest.md) |  | 

### Return type

[**S3ServiceAccount**](S3ServiceAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ListServiceAccounts

> S3ListServiceAccounts200Response S3ListServiceAccounts(ctx).Execute()

List service accounts

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
	resp, r, err := apiClient.AdminAPI.S3ListServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.S3ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3ListServiceAccounts`: S3ListServiceAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.S3ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ListServiceAccountsRequest struct via the builder pattern


### Return type

[**S3ListServiceAccounts200Response**](S3ListServiceAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

