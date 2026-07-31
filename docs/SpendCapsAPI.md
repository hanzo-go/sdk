# \SpendCapsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillingBillingAuthorizeSpendCap**](SpendCapsAPI.md#BillingBillingAuthorizeSpendCap) | **Get** /v1/billing/spend-alerts/authorize | Per-request spend-cap verdict
[**BillingBillingCreateSpendAlert**](SpendCapsAPI.md#BillingBillingCreateSpendAlert) | **Post** /v1/billing/spend-alerts | Create a spend alert / cap
[**BillingBillingDeleteSpendAlert**](SpendCapsAPI.md#BillingBillingDeleteSpendAlert) | **Delete** /v1/billing/spend-alerts/{id} | Delete a spend alert / cap
[**BillingBillingListSpendAlerts**](SpendCapsAPI.md#BillingBillingListSpendAlerts) | **Get** /v1/billing/spend-alerts | List spend alerts / caps
[**BillingBillingUpdateSpendAlert**](SpendCapsAPI.md#BillingBillingUpdateSpendAlert) | **Patch** /v1/billing/spend-alerts/{id} | Update a spend alert / cap



## BillingBillingAuthorizeSpendCap

> BillingCapVerdict BillingBillingAuthorizeSpendCap(ctx).User(user).Project(project).Service(service).Amount(amount).Pv(pv).Currency(currency).Execute()

Per-request spend-cap verdict



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
	user := "user_example" // string | Org billing key (equals the org slug). (optional)
	project := "project_example" // string | Scope project axis; \"\" or \"default\" = the org-wide default scope. (optional)
	service := "service_example" // string | Scope service axis (server-derived route/provider). (optional)
	amount := int64(789) // int64 | Proposed charge in USD cents to test against the cap (0 = a pure \"already over?\" check). (optional)
	pv := "pv_example" // string | 1 = the project axis is bound to a VALIDATED claim (hardens project-scoped caps); otherwise a project-scoped cap degrades to a soft warn (anti-spoof). (optional)
	currency := "currency_example" // string |  (optional) (default to "usd")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendCapsAPI.BillingBillingAuthorizeSpendCap(context.Background()).User(user).Project(project).Service(service).Amount(amount).Pv(pv).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendCapsAPI.BillingBillingAuthorizeSpendCap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingAuthorizeSpendCap`: BillingCapVerdict
	fmt.Fprintf(os.Stdout, "Response from `SpendCapsAPI.BillingBillingAuthorizeSpendCap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingAuthorizeSpendCapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** | Org billing key (equals the org slug). | 
 **project** | **string** | Scope project axis; \&quot;\&quot; or \&quot;default\&quot; &#x3D; the org-wide default scope. | 
 **service** | **string** | Scope service axis (server-derived route/provider). | 
 **amount** | **int64** | Proposed charge in USD cents to test against the cap (0 &#x3D; a pure \&quot;already over?\&quot; check). | 
 **pv** | **string** | 1 &#x3D; the project axis is bound to a VALIDATED claim (hardens project-scoped caps); otherwise a project-scoped cap degrades to a soft warn (anti-spoof). | 
 **currency** | **string** |  | [default to &quot;usd&quot;]

### Return type

[**BillingCapVerdict**](BillingCapVerdict.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingCreateSpendAlert

> BillingSpendAlert BillingBillingCreateSpendAlert(ctx).BillingSpendAlertCreate(billingSpendAlertCreate).Execute()

Create a spend alert / cap



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
	billingSpendAlertCreate := *openapiclient.NewBillingSpendAlertCreate() // BillingSpendAlertCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendCapsAPI.BillingBillingCreateSpendAlert(context.Background()).BillingSpendAlertCreate(billingSpendAlertCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendCapsAPI.BillingBillingCreateSpendAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingCreateSpendAlert`: BillingSpendAlert
	fmt.Fprintf(os.Stdout, "Response from `SpendCapsAPI.BillingBillingCreateSpendAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingCreateSpendAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billingSpendAlertCreate** | [**BillingSpendAlertCreate**](BillingSpendAlertCreate.md) |  | 

### Return type

[**BillingSpendAlert**](BillingSpendAlert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingDeleteSpendAlert

> BillingBillingDeleteSpendAlert(ctx, id).Execute()

Delete a spend alert / cap



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
	r, err := apiClient.SpendCapsAPI.BillingBillingDeleteSpendAlert(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendCapsAPI.BillingBillingDeleteSpendAlert``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBillingBillingDeleteSpendAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingListSpendAlerts

> []BillingSpendAlert BillingBillingListSpendAlerts(ctx).Execute()

List spend alerts / caps



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
	resp, r, err := apiClient.SpendCapsAPI.BillingBillingListSpendAlerts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendCapsAPI.BillingBillingListSpendAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingListSpendAlerts`: []BillingSpendAlert
	fmt.Fprintf(os.Stdout, "Response from `SpendCapsAPI.BillingBillingListSpendAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingListSpendAlertsRequest struct via the builder pattern


### Return type

[**[]BillingSpendAlert**](BillingSpendAlert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillingBillingUpdateSpendAlert

> BillingSpendAlert BillingBillingUpdateSpendAlert(ctx, id).BillingSpendAlertUpdate(billingSpendAlertUpdate).Execute()

Update a spend alert / cap



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
	billingSpendAlertUpdate := *openapiclient.NewBillingSpendAlertUpdate() // BillingSpendAlertUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpendCapsAPI.BillingBillingUpdateSpendAlert(context.Background(), id).BillingSpendAlertUpdate(billingSpendAlertUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpendCapsAPI.BillingBillingUpdateSpendAlert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillingBillingUpdateSpendAlert`: BillingSpendAlert
	fmt.Fprintf(os.Stdout, "Response from `SpendCapsAPI.BillingBillingUpdateSpendAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillingBillingUpdateSpendAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billingSpendAlertUpdate** | [**BillingSpendAlertUpdate**](BillingSpendAlertUpdate.md) |  | 

### Return type

[**BillingSpendAlert**](BillingSpendAlert.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

