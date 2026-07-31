# \CaptableAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1CaptableConvertiblesId**](CaptableAPI.md#CloudDeleteV1CaptableConvertiblesId) | **Delete** /v1/captable/convertibles/{id} | DeleteConvertible removes one of the caller org&#39;s convertible notes, taking its principal out of the cap table&#39;s unconverted-instrument totals.
[**CloudDeleteV1CaptableOptionsId**](CaptableAPI.md#CloudDeleteV1CaptableOptionsId) | **Delete** /v1/captable/options/{id} | DeleteOption removes one of the caller org&#39;s option grants, taking its shares out of the cap table&#39;s granted-options and fully-diluted counts.
[**CloudDeleteV1CaptableSafesId**](CaptableAPI.md#CloudDeleteV1CaptableSafesId) | **Delete** /v1/captable/safes/{id} | DeleteSafe removes one of the caller org&#39;s SAFEs, taking its capital out of the cap table&#39;s unconverted-instrument totals.
[**CloudDeleteV1CaptableSharesId**](CaptableAPI.md#CloudDeleteV1CaptableSharesId) | **Delete** /v1/captable/shares/{id} | DeleteShare removes one of the caller org&#39;s share certificates, taking its shares out of the cap table&#39;s outstanding and fully-diluted counts.
[**CloudDeleteV1CaptableStakeholdersId**](CaptableAPI.md#CloudDeleteV1CaptableStakeholdersId) | **Delete** /v1/captable/stakeholders/{id} | DeleteStakeholder removes one of the caller org&#39;s stakeholders.
[**CloudGetV1CaptableCompany**](CaptableAPI.md#CloudGetV1CaptableCompany) | **Get** /v1/captable/company | GetCompany returns the caller org&#39;s cap-table company record.
[**CloudGetV1CaptableConvertibles**](CaptableAPI.md#CloudGetV1CaptableConvertibles) | **Get** /v1/captable/convertibles | ListConvertibles returns the caller org&#39;s convertible notes, newest first.
[**CloudGetV1CaptableEquityPlans**](CaptableAPI.md#CloudGetV1CaptableEquityPlans) | **Get** /v1/captable/equity-plans | ListEquityPlans returns the caller org&#39;s equity plans, newest first.
[**CloudGetV1CaptableInvestments**](CaptableAPI.md#CloudGetV1CaptableInvestments) | **Get** /v1/captable/investments | ListInvestments returns the caller org&#39;s investments, newest first.
[**CloudGetV1CaptableOptions**](CaptableAPI.md#CloudGetV1CaptableOptions) | **Get** /v1/captable/options | ListOptions returns the caller org&#39;s option grants, newest first.
[**CloudGetV1CaptableRounds**](CaptableAPI.md#CloudGetV1CaptableRounds) | **Get** /v1/captable/rounds | ListRounds returns the caller org&#39;s fundraising rounds, newest first.
[**CloudGetV1CaptableRoundsId**](CaptableAPI.md#CloudGetV1CaptableRoundsId) | **Get** /v1/captable/rounds/{id} | GetRound returns one of the caller org&#39;s fundraising rounds together with every investment written into it, oldest first.
[**CloudGetV1CaptableSafes**](CaptableAPI.md#CloudGetV1CaptableSafes) | **Get** /v1/captable/safes | ListSafes returns the caller org&#39;s SAFEs, newest first.
[**CloudGetV1CaptableShareClasses**](CaptableAPI.md#CloudGetV1CaptableShareClasses) | **Get** /v1/captable/share-classes | ListShareClasses returns the caller org&#39;s share classes, in creation order.
[**CloudGetV1CaptableShares**](CaptableAPI.md#CloudGetV1CaptableShares) | **Get** /v1/captable/shares | ListShares returns the caller org&#39;s share certificates, newest first.
[**CloudGetV1CaptableStakeholders**](CaptableAPI.md#CloudGetV1CaptableStakeholders) | **Get** /v1/captable/stakeholders | ListStakeholders returns the caller org&#39;s stakeholders, newest first.
[**CloudGetV1CaptableSummary**](CaptableAPI.md#CloudGetV1CaptableSummary) | **Get** /v1/captable/summary | GetSummary computes the caller org&#39;s cap table.
[**CloudPatchV1CaptableShareClassesById**](CaptableAPI.md#CloudPatchV1CaptableShareClassesById) | **Patch** /v1/captable/share-classes/{id} | 
[**CloudPatchV1CaptableStakeholdersId**](CaptableAPI.md#CloudPatchV1CaptableStakeholdersId) | **Patch** /v1/captable/stakeholders/{id} | UpdateStakeholder changes one of the caller org&#39;s stakeholders.
[**CloudPostV1CaptableConvertibles**](CaptableAPI.md#CloudPostV1CaptableConvertibles) | **Post** /v1/captable/convertibles | 
[**CloudPostV1CaptableEquityPlans**](CaptableAPI.md#CloudPostV1CaptableEquityPlans) | **Post** /v1/captable/equity-plans | 
[**CloudPostV1CaptableOptions**](CaptableAPI.md#CloudPostV1CaptableOptions) | **Post** /v1/captable/options | 
[**CloudPostV1CaptableRounds**](CaptableAPI.md#CloudPostV1CaptableRounds) | **Post** /v1/captable/rounds | 
[**CloudPostV1CaptableRoundsByIdInvestments**](CaptableAPI.md#CloudPostV1CaptableRoundsByIdInvestments) | **Post** /v1/captable/rounds/{id}/investments | 
[**CloudPostV1CaptableRoundsIdClose**](CaptableAPI.md#CloudPostV1CaptableRoundsIdClose) | **Post** /v1/captable/rounds/{id}/close | CloseRound closes one of the caller org&#39;s fundraising rounds, recording the close date and moving its status to CLOSED.
[**CloudPostV1CaptableSafes**](CaptableAPI.md#CloudPostV1CaptableSafes) | **Post** /v1/captable/safes | 
[**CloudPostV1CaptableShareClasses**](CaptableAPI.md#CloudPostV1CaptableShareClasses) | **Post** /v1/captable/share-classes | 
[**CloudPostV1CaptableShares**](CaptableAPI.md#CloudPostV1CaptableShares) | **Post** /v1/captable/shares | 
[**CloudPostV1CaptableSharesTransfer**](CaptableAPI.md#CloudPostV1CaptableSharesTransfer) | **Post** /v1/captable/shares/transfer | 
[**CloudPostV1CaptableStakeholders**](CaptableAPI.md#CloudPostV1CaptableStakeholders) | **Post** /v1/captable/stakeholders | 
[**CloudPutV1CaptableCompany**](CaptableAPI.md#CloudPutV1CaptableCompany) | **Put** /v1/captable/company | UpdateCompany sets the caller org&#39;s company name and incorporation details.



## CloudDeleteV1CaptableConvertiblesId

> CloudCaptableDeleted CloudDeleteV1CaptableConvertiblesId(ctx, id).Execute()

DeleteConvertible removes one of the caller org's convertible notes, taking its principal out of the cap table's unconverted-instrument totals.



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
	id := "id_example" // string | ID is the convertible note to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudDeleteV1CaptableConvertiblesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudDeleteV1CaptableConvertiblesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CaptableConvertiblesId`: CloudCaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudDeleteV1CaptableConvertiblesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the convertible note to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CaptableConvertiblesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableDeleted**](CloudCaptableDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CaptableOptionsId

> CloudCaptableDeleted CloudDeleteV1CaptableOptionsId(ctx, id).Execute()

DeleteOption removes one of the caller org's option grants, taking its shares out of the cap table's granted-options and fully-diluted counts.



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
	id := "id_example" // string | ID is the option grant to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudDeleteV1CaptableOptionsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudDeleteV1CaptableOptionsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CaptableOptionsId`: CloudCaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudDeleteV1CaptableOptionsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the option grant to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CaptableOptionsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableDeleted**](CloudCaptableDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CaptableSafesId

> CloudCaptableDeleted CloudDeleteV1CaptableSafesId(ctx, id).Execute()

DeleteSafe removes one of the caller org's SAFEs, taking its capital out of the cap table's unconverted-instrument totals.



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
	id := "id_example" // string | ID is the SAFE to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudDeleteV1CaptableSafesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudDeleteV1CaptableSafesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CaptableSafesId`: CloudCaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudDeleteV1CaptableSafesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the SAFE to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CaptableSafesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableDeleted**](CloudCaptableDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CaptableSharesId

> CloudCaptableDeleted CloudDeleteV1CaptableSharesId(ctx, id).Execute()

DeleteShare removes one of the caller org's share certificates, taking its shares out of the cap table's outstanding and fully-diluted counts.



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
	id := "id_example" // string | ID is the share certificate to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudDeleteV1CaptableSharesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudDeleteV1CaptableSharesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CaptableSharesId`: CloudCaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudDeleteV1CaptableSharesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the share certificate to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CaptableSharesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableDeleted**](CloudCaptableDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CaptableStakeholdersId

> CloudCaptableDeleted CloudDeleteV1CaptableStakeholdersId(ctx, id).Execute()

DeleteStakeholder removes one of the caller org's stakeholders.



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
	id := "id_example" // string | ID is the stakeholder to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudDeleteV1CaptableStakeholdersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudDeleteV1CaptableStakeholdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CaptableStakeholdersId`: CloudCaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudDeleteV1CaptableStakeholdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the stakeholder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CaptableStakeholdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableDeleted**](CloudCaptableDeleted.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableCompany

> CloudCaptableCompany CloudGetV1CaptableCompany(ctx).Execute()

GetCompany returns the caller org's cap-table company record.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableCompany(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableCompany`: CloudCaptableCompany
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableCompanyRequest struct via the builder pattern


### Return type

[**CloudCaptableCompany**](CloudCaptableCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableConvertibles

> CloudCaptableNotes CloudGetV1CaptableConvertibles(ctx).Execute()

ListConvertibles returns the caller org's convertible notes, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableConvertibles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableConvertibles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableConvertibles`: CloudCaptableNotes
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableConvertibles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableConvertiblesRequest struct via the builder pattern


### Return type

[**CloudCaptableNotes**](CloudCaptableNotes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableEquityPlans

> CloudCaptableEquityPlans CloudGetV1CaptableEquityPlans(ctx).Execute()

ListEquityPlans returns the caller org's equity plans, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableEquityPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableEquityPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableEquityPlans`: CloudCaptableEquityPlans
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableEquityPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableEquityPlansRequest struct via the builder pattern


### Return type

[**CloudCaptableEquityPlans**](CloudCaptableEquityPlans.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableInvestments

> CloudCaptableInvestments CloudGetV1CaptableInvestments(ctx).Execute()

ListInvestments returns the caller org's investments, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableInvestments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableInvestments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableInvestments`: CloudCaptableInvestments
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableInvestments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableInvestmentsRequest struct via the builder pattern


### Return type

[**CloudCaptableInvestments**](CloudCaptableInvestments.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableOptions

> CloudCaptableOptions CloudGetV1CaptableOptions(ctx).Execute()

ListOptions returns the caller org's option grants, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableOptions`: CloudCaptableOptions
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableOptionsRequest struct via the builder pattern


### Return type

[**CloudCaptableOptions**](CloudCaptableOptions.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableRounds

> CloudCaptableRounds CloudGetV1CaptableRounds(ctx).Execute()

ListRounds returns the caller org's fundraising rounds, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableRounds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableRounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableRounds`: CloudCaptableRounds
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableRounds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableRoundsRequest struct via the builder pattern


### Return type

[**CloudCaptableRounds**](CloudCaptableRounds.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableRoundsId

> CloudCaptableRoundDetail CloudGetV1CaptableRoundsId(ctx, id).Execute()

GetRound returns one of the caller org's fundraising rounds together with every investment written into it, oldest first.



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
	id := "id_example" // string | ID is the round to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableRoundsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableRoundsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableRoundsId`: CloudCaptableRoundDetail
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableRoundsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the round to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableRoundsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCaptableRoundDetail**](CloudCaptableRoundDetail.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableSafes

> CloudCaptableSafes CloudGetV1CaptableSafes(ctx).Execute()

ListSafes returns the caller org's SAFEs, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableSafes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableSafes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableSafes`: CloudCaptableSafes
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableSafes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableSafesRequest struct via the builder pattern


### Return type

[**CloudCaptableSafes**](CloudCaptableSafes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableShareClasses

> []CloudCaptableShareClass CloudGetV1CaptableShareClasses(ctx).Execute()

ListShareClasses returns the caller org's share classes, in creation order.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableShareClasses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableShareClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableShareClasses`: []CloudCaptableShareClass
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableShareClasses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableShareClassesRequest struct via the builder pattern


### Return type

[**[]CloudCaptableShareClass**](CloudCaptableShareClass.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableShares

> CloudCaptableShares CloudGetV1CaptableShares(ctx).Execute()

ListShares returns the caller org's share certificates, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableShares`: CloudCaptableShares
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableShares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableSharesRequest struct via the builder pattern


### Return type

[**CloudCaptableShares**](CloudCaptableShares.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableStakeholders

> []CloudCaptableStakeholder CloudGetV1CaptableStakeholders(ctx).Execute()

ListStakeholders returns the caller org's stakeholders, newest first.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableStakeholders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableStakeholders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableStakeholders`: []CloudCaptableStakeholder
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableStakeholders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableStakeholdersRequest struct via the builder pattern


### Return type

[**[]CloudCaptableStakeholder**](CloudCaptableStakeholder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CaptableSummary

> CloudCaptableSummary CloudGetV1CaptableSummary(ctx).Execute()

GetSummary computes the caller org's cap table.



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
	resp, r, err := apiClient.CaptableAPI.CloudGetV1CaptableSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudGetV1CaptableSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CaptableSummary`: CloudCaptableSummary
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudGetV1CaptableSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CaptableSummaryRequest struct via the builder pattern


### Return type

[**CloudCaptableSummary**](CloudCaptableSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1CaptableShareClassesById

> CloudPatchV1CaptableShareClassesById(ctx, id).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPatchV1CaptableShareClassesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPatchV1CaptableShareClassesById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPatchV1CaptableShareClassesByIdRequest struct via the builder pattern


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


## CloudPatchV1CaptableStakeholdersId

> CloudCaptableUpdated CloudPatchV1CaptableStakeholdersId(ctx, id).CloudCaptableStakeholderPatch(cloudCaptableStakeholderPatch).Execute()

UpdateStakeholder changes one of the caller org's stakeholders.



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
	id := "id_example" // string | ID is the stakeholder to update. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.
	cloudCaptableStakeholderPatch := *openapiclient.NewCloudCaptableStakeholderPatch() // CloudCaptableStakeholderPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudPatchV1CaptableStakeholdersId(context.Background(), id).CloudCaptableStakeholderPatch(cloudCaptableStakeholderPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPatchV1CaptableStakeholdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1CaptableStakeholdersId`: CloudCaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudPatchV1CaptableStakeholdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the stakeholder to update. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1CaptableStakeholdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCaptableStakeholderPatch** | [**CloudCaptableStakeholderPatch**](CloudCaptableStakeholderPatch.md) |  | 

### Return type

[**CloudCaptableUpdated**](CloudCaptableUpdated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CaptableConvertibles

> CloudPostV1CaptableConvertibles(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableConvertibles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableConvertibles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableConvertiblesRequest struct via the builder pattern


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


## CloudPostV1CaptableEquityPlans

> CloudPostV1CaptableEquityPlans(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableEquityPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableEquityPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableEquityPlansRequest struct via the builder pattern


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


## CloudPostV1CaptableOptions

> CloudPostV1CaptableOptions(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableOptionsRequest struct via the builder pattern


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


## CloudPostV1CaptableRounds

> CloudPostV1CaptableRounds(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableRounds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableRounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableRoundsRequest struct via the builder pattern


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


## CloudPostV1CaptableRoundsByIdInvestments

> CloudPostV1CaptableRoundsByIdInvestments(ctx, id).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableRoundsByIdInvestments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableRoundsByIdInvestments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1CaptableRoundsByIdInvestmentsRequest struct via the builder pattern


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


## CloudPostV1CaptableRoundsIdClose

> CloudCaptableUpdated CloudPostV1CaptableRoundsIdClose(ctx, id).CloudCaptableRoundCloseRequest(cloudCaptableRoundCloseRequest).Execute()

CloseRound closes one of the caller org's fundraising rounds, recording the close date and moving its status to CLOSED.



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
	id := "id_example" // string | ID is the round to close. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.
	cloudCaptableRoundCloseRequest := *openapiclient.NewCloudCaptableRoundCloseRequest() // CloudCaptableRoundCloseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudPostV1CaptableRoundsIdClose(context.Background(), id).CloudCaptableRoundCloseRequest(cloudCaptableRoundCloseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableRoundsIdClose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CaptableRoundsIdClose`: CloudCaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudPostV1CaptableRoundsIdClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the round to close. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableRoundsIdCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCaptableRoundCloseRequest** | [**CloudCaptableRoundCloseRequest**](CloudCaptableRoundCloseRequest.md) |  | 

### Return type

[**CloudCaptableUpdated**](CloudCaptableUpdated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CaptableSafes

> CloudPostV1CaptableSafes(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableSafes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableSafes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableSafesRequest struct via the builder pattern


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


## CloudPostV1CaptableShareClasses

> CloudPostV1CaptableShareClasses(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableShareClasses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableShareClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableShareClassesRequest struct via the builder pattern


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


## CloudPostV1CaptableShares

> CloudPostV1CaptableShares(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableSharesRequest struct via the builder pattern


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


## CloudPostV1CaptableSharesTransfer

> CloudPostV1CaptableSharesTransfer(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableSharesTransfer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableSharesTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableSharesTransferRequest struct via the builder pattern


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


## CloudPostV1CaptableStakeholders

> CloudPostV1CaptableStakeholders(ctx).Execute()



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
	r, err := apiClient.CaptableAPI.CloudPostV1CaptableStakeholders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPostV1CaptableStakeholders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CaptableStakeholdersRequest struct via the builder pattern


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


## CloudPutV1CaptableCompany

> CloudCaptableUpdated CloudPutV1CaptableCompany(ctx).CloudCaptableCompanyUpdate(cloudCaptableCompanyUpdate).Execute()

UpdateCompany sets the caller org's company name and incorporation details.



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
	cloudCaptableCompanyUpdate := *openapiclient.NewCloudCaptableCompanyUpdate() // CloudCaptableCompanyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.CloudPutV1CaptableCompany(context.Background()).CloudCaptableCompanyUpdate(cloudCaptableCompanyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.CloudPutV1CaptableCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CaptableCompany`: CloudCaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.CloudPutV1CaptableCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CaptableCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCaptableCompanyUpdate** | [**CloudCaptableCompanyUpdate**](CloudCaptableCompanyUpdate.md) |  | 

### Return type

[**CloudCaptableUpdated**](CloudCaptableUpdated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

