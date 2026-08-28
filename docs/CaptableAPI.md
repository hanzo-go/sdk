# \CaptableAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCaptableConvertiblesById**](CaptableAPI.md#DeleteCaptableConvertiblesById) | **Delete** /v1/captable/convertibles/{id} | Removes one of the caller org&#39;s convertible notes, taking its principal out of the cap table&#39;s unconverted-instrument totals.
[**DeleteCaptableOptionsById**](CaptableAPI.md#DeleteCaptableOptionsById) | **Delete** /v1/captable/options/{id} | Removes one of the caller org&#39;s option grants, taking its shares out of the cap table&#39;s granted-options and fully-diluted counts.
[**DeleteCaptableSafesById**](CaptableAPI.md#DeleteCaptableSafesById) | **Delete** /v1/captable/safes/{id} | Removes one of the caller org&#39;s SAFEs, taking its capital out of the cap table&#39;s unconverted-instrument totals.
[**DeleteCaptableSharesById**](CaptableAPI.md#DeleteCaptableSharesById) | **Delete** /v1/captable/shares/{id} | Removes one of the caller org&#39;s share certificates, taking its shares out of the cap table&#39;s outstanding and fully-diluted counts.
[**DeleteCaptableStakeholdersById**](CaptableAPI.md#DeleteCaptableStakeholdersById) | **Delete** /v1/captable/stakeholders/{id} | Removes one of the caller org&#39;s stakeholders.
[**GetCaptableClasses**](CaptableAPI.md#GetCaptableClasses) | **Get** /v1/captable/classes | Returns the caller org&#39;s share classes, in creation order.
[**GetCaptableCompany**](CaptableAPI.md#GetCaptableCompany) | **Get** /v1/captable/company | Returns the caller org&#39;s cap-table company record.
[**GetCaptableConvertibles**](CaptableAPI.md#GetCaptableConvertibles) | **Get** /v1/captable/convertibles | Returns the caller org&#39;s convertible notes, newest first.
[**GetCaptableInvestments**](CaptableAPI.md#GetCaptableInvestments) | **Get** /v1/captable/investments | Returns the caller org&#39;s investments, newest first.
[**GetCaptableOptions**](CaptableAPI.md#GetCaptableOptions) | **Get** /v1/captable/options | Returns the caller org&#39;s option grants, newest first.
[**GetCaptablePlans**](CaptableAPI.md#GetCaptablePlans) | **Get** /v1/captable/plans | Returns the caller org&#39;s equity plans, newest first.
[**GetCaptableRounds**](CaptableAPI.md#GetCaptableRounds) | **Get** /v1/captable/rounds | Returns the caller org&#39;s fundraising rounds, newest first.
[**GetCaptableRoundsById**](CaptableAPI.md#GetCaptableRoundsById) | **Get** /v1/captable/rounds/{id} | Returns one of the caller org&#39;s fundraising rounds together with every investment written into it, oldest first.
[**GetCaptableSafes**](CaptableAPI.md#GetCaptableSafes) | **Get** /v1/captable/safes | Returns the caller org&#39;s SAFEs, newest first.
[**GetCaptableShares**](CaptableAPI.md#GetCaptableShares) | **Get** /v1/captable/shares | Returns the caller org&#39;s share certificates, newest first.
[**GetCaptableStakeholders**](CaptableAPI.md#GetCaptableStakeholders) | **Get** /v1/captable/stakeholders | Returns the caller org&#39;s stakeholders, newest first.
[**GetCaptableSummary**](CaptableAPI.md#GetCaptableSummary) | **Get** /v1/captable/summary | Computes the caller org&#39;s cap table.
[**PatchCaptableClassesById**](CaptableAPI.md#PatchCaptableClassesById) | **Patch** /v1/captable/classes/{id} | Replaces one share class&#39;s terms.
[**PatchCaptableStakeholdersById**](CaptableAPI.md#PatchCaptableStakeholdersById) | **Patch** /v1/captable/stakeholders/{id} | Changes one of the caller org&#39;s stakeholders.
[**PostCaptableClasses**](CaptableAPI.md#PostCaptableClasses) | **Post** /v1/captable/classes | Defines a new class of shares.
[**PostCaptableConvertibles**](CaptableAPI.md#PostCaptableConvertibles) | **Post** /v1/captable/convertibles | Records a convertible note.
[**PostCaptableOptions**](CaptableAPI.md#PostCaptableOptions) | **Post** /v1/captable/options | Grants options to a stakeholder from an equity plan.
[**PostCaptablePlans**](CaptableAPI.md#PostCaptablePlans) | **Post** /v1/captable/plans | Opens an equity plan that options are granted from.
[**PostCaptableRounds**](CaptableAPI.md#PostCaptableRounds) | **Post** /v1/captable/rounds | Opens a priced round that investments can be added to.
[**PostCaptableRoundsByIdClose**](CaptableAPI.md#PostCaptableRoundsByIdClose) | **Post** /v1/captable/rounds/{id}/close | Closes one of the caller org&#39;s fundraising rounds, recording the close date and moving its status to CLOSED.
[**PostCaptableRoundsByIdInvestments**](CaptableAPI.md#PostCaptableRoundsByIdInvestments) | **Post** /v1/captable/rounds/{id}/investments | Records one investor&#39;s money into an open round.
[**PostCaptableSafes**](CaptableAPI.md#PostCaptableSafes) | **Post** /v1/captable/safes | Records a SAFE — a simple agreement for future equity.
[**PostCaptableShares**](CaptableAPI.md#PostCaptableShares) | **Post** /v1/captable/shares | Issues a share certificate to a stakeholder.
[**PostCaptableSharesTransfer**](CaptableAPI.md#PostCaptableSharesTransfer) | **Post** /v1/captable/shares/transfer | Moves shares from one stakeholder to another.
[**PostCaptableStakeholders**](CaptableAPI.md#PostCaptableStakeholders) | **Post** /v1/captable/stakeholders | Add stakeholders to the cap table
[**PutCaptableCompany**](CaptableAPI.md#PutCaptableCompany) | **Put** /v1/captable/company | Sets the caller org&#39;s company name and incorporation details.



## DeleteCaptableConvertiblesById

> CaptableDeleted DeleteCaptableConvertiblesById(ctx, id).Execute()

Removes one of the caller org's convertible notes, taking its principal out of the cap table's unconverted-instrument totals.



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
	id := "id_example" // string | ID is the convertible note to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.DeleteCaptableConvertiblesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.DeleteCaptableConvertiblesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCaptableConvertiblesById`: CaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.DeleteCaptableConvertiblesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the convertible note to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptableConvertiblesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableDeleted**](CaptableDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptableOptionsById

> CaptableDeleted DeleteCaptableOptionsById(ctx, id).Execute()

Removes one of the caller org's option grants, taking its shares out of the cap table's granted-options and fully-diluted counts.



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
	id := "id_example" // string | ID is the option grant to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.DeleteCaptableOptionsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.DeleteCaptableOptionsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCaptableOptionsById`: CaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.DeleteCaptableOptionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the option grant to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptableOptionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableDeleted**](CaptableDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptableSafesById

> CaptableDeleted DeleteCaptableSafesById(ctx, id).Execute()

Removes one of the caller org's SAFEs, taking its capital out of the cap table's unconverted-instrument totals.



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
	id := "id_example" // string | ID is the SAFE to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.DeleteCaptableSafesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.DeleteCaptableSafesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCaptableSafesById`: CaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.DeleteCaptableSafesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the SAFE to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptableSafesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableDeleted**](CaptableDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptableSharesById

> CaptableDeleted DeleteCaptableSharesById(ctx, id).Execute()

Removes one of the caller org's share certificates, taking its shares out of the cap table's outstanding and fully-diluted counts.



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
	id := "id_example" // string | ID is the share certificate to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.DeleteCaptableSharesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.DeleteCaptableSharesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCaptableSharesById`: CaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.DeleteCaptableSharesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the share certificate to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptableSharesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableDeleted**](CaptableDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCaptableStakeholdersById

> CaptableDeleted DeleteCaptableStakeholdersById(ctx, id).Execute()

Removes one of the caller org's stakeholders.



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
	id := "id_example" // string | ID is the stakeholder to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.DeleteCaptableStakeholdersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.DeleteCaptableStakeholdersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCaptableStakeholdersById`: CaptableDeleted
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.DeleteCaptableStakeholdersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the stakeholder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCaptableStakeholdersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableDeleted**](CaptableDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableClasses

> []CaptableShareClass GetCaptableClasses(ctx).Execute()

Returns the caller org's share classes, in creation order.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableClasses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableClasses`: []CaptableShareClass
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableClasses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableClassesRequest struct via the builder pattern


### Return type

[**[]CaptableShareClass**](CaptableShareClass.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableCompany

> CaptableCompany GetCaptableCompany(ctx).Execute()

Returns the caller org's cap-table company record.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableCompany(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableCompany`: CaptableCompany
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableCompany`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableCompanyRequest struct via the builder pattern


### Return type

[**CaptableCompany**](CaptableCompany.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableConvertibles

> CaptableNotes GetCaptableConvertibles(ctx).Execute()

Returns the caller org's convertible notes, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableConvertibles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableConvertibles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableConvertibles`: CaptableNotes
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableConvertibles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableConvertiblesRequest struct via the builder pattern


### Return type

[**CaptableNotes**](CaptableNotes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableInvestments

> CaptableInvestments GetCaptableInvestments(ctx).Execute()

Returns the caller org's investments, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableInvestments(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableInvestments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableInvestments`: CaptableInvestments
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableInvestments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableInvestmentsRequest struct via the builder pattern


### Return type

[**CaptableInvestments**](CaptableInvestments.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableOptions

> CaptableOptions GetCaptableOptions(ctx).Execute()

Returns the caller org's option grants, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableOptions`: CaptableOptions
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableOptionsRequest struct via the builder pattern


### Return type

[**CaptableOptions**](CaptableOptions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptablePlans

> CaptableEquityPlans GetCaptablePlans(ctx).Execute()

Returns the caller org's equity plans, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptablePlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptablePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptablePlans`: CaptableEquityPlans
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptablePlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptablePlansRequest struct via the builder pattern


### Return type

[**CaptableEquityPlans**](CaptableEquityPlans.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableRounds

> CaptableRounds GetCaptableRounds(ctx).Execute()

Returns the caller org's fundraising rounds, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableRounds(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableRounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableRounds`: CaptableRounds
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableRounds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableRoundsRequest struct via the builder pattern


### Return type

[**CaptableRounds**](CaptableRounds.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableRoundsById

> CaptableRoundDetail GetCaptableRoundsById(ctx, id).Execute()

Returns one of the caller org's fundraising rounds together with every investment written into it, oldest first.



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
	id := "id_example" // string | ID is the round to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.GetCaptableRoundsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableRoundsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableRoundsById`: CaptableRoundDetail
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableRoundsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the round to read. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableRoundsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CaptableRoundDetail**](CaptableRoundDetail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableSafes

> CaptableSafes GetCaptableSafes(ctx).Execute()

Returns the caller org's SAFEs, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableSafes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableSafes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableSafes`: CaptableSafes
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableSafes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableSafesRequest struct via the builder pattern


### Return type

[**CaptableSafes**](CaptableSafes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableShares

> CaptableShares GetCaptableShares(ctx).Execute()

Returns the caller org's share certificates, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableShares`: CaptableShares
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableShares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableSharesRequest struct via the builder pattern


### Return type

[**CaptableShares**](CaptableShares.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableStakeholders

> []CaptableStakeholder GetCaptableStakeholders(ctx).Execute()

Returns the caller org's stakeholders, newest first.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableStakeholders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableStakeholders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableStakeholders`: []CaptableStakeholder
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableStakeholders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableStakeholdersRequest struct via the builder pattern


### Return type

[**[]CaptableStakeholder**](CaptableStakeholder.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCaptableSummary

> CaptableSummary GetCaptableSummary(ctx).Execute()

Computes the caller org's cap table.



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
	resp, r, err := apiClient.CaptableAPI.GetCaptableSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.GetCaptableSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCaptableSummary`: CaptableSummary
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.GetCaptableSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCaptableSummaryRequest struct via the builder pattern


### Return type

[**CaptableSummary**](CaptableSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCaptableClassesById

> CaptableUpdated PatchCaptableClassesById(ctx, id).CaptableShareClassAmend(captableShareClassAmend).Execute()

Replaces one share class's terms.



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
	id := "id_example" // string | ID addresses the resource. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which row is written whatever a body claims.
	captableShareClassAmend := *openapiclient.NewCaptableShareClassAmend() // CaptableShareClassAmend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PatchCaptableClassesById(context.Background(), id).CaptableShareClassAmend(captableShareClassAmend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PatchCaptableClassesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCaptableClassesById`: CaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PatchCaptableClassesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID addresses the resource. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which row is written whatever a body claims. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCaptableClassesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captableShareClassAmend** | [**CaptableShareClassAmend**](CaptableShareClassAmend.md) |  | 

### Return type

[**CaptableUpdated**](CaptableUpdated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCaptableStakeholdersById

> CaptableUpdated PatchCaptableStakeholdersById(ctx, id).CaptableStakeholderPatch(captableStakeholderPatch).Execute()

Changes one of the caller org's stakeholders.



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
	id := "id_example" // string | ID is the stakeholder to update. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.
	captableStakeholderPatch := *openapiclient.NewCaptableStakeholderPatch() // CaptableStakeholderPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PatchCaptableStakeholdersById(context.Background(), id).CaptableStakeholderPatch(captableStakeholderPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PatchCaptableStakeholdersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCaptableStakeholdersById`: CaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PatchCaptableStakeholdersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the stakeholder to update. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCaptableStakeholdersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captableStakeholderPatch** | [**CaptableStakeholderPatch**](CaptableStakeholderPatch.md) |  | 

### Return type

[**CaptableUpdated**](CaptableUpdated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableClasses

> CaptableCreated PostCaptableClasses(ctx).CaptableShareClassIn(captableShareClassIn).Execute()

Defines a new class of shares.



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
	captableShareClassIn := *openapiclient.NewCaptableShareClassIn() // CaptableShareClassIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableClasses(context.Background()).CaptableShareClassIn(captableShareClassIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableClasses`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableClasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableShareClassIn** | [**CaptableShareClassIn**](CaptableShareClassIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableConvertibles

> CaptableCreated PostCaptableConvertibles(ctx).CaptableConvertibleIn(captableConvertibleIn).Execute()

Records a convertible note.



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
	captableConvertibleIn := *openapiclient.NewCaptableConvertibleIn() // CaptableConvertibleIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableConvertibles(context.Background()).CaptableConvertibleIn(captableConvertibleIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableConvertibles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableConvertibles`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableConvertibles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableConvertiblesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableConvertibleIn** | [**CaptableConvertibleIn**](CaptableConvertibleIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableOptions

> CaptableCreated PostCaptableOptions(ctx).CaptableOptionIn(captableOptionIn).Execute()

Grants options to a stakeholder from an equity plan.



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
	captableOptionIn := *openapiclient.NewCaptableOptionIn() // CaptableOptionIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableOptions(context.Background()).CaptableOptionIn(captableOptionIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableOptions`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableOptionIn** | [**CaptableOptionIn**](CaptableOptionIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptablePlans

> CaptableCreated PostCaptablePlans(ctx).CaptableEquityPlanIn(captableEquityPlanIn).Execute()

Opens an equity plan that options are granted from.



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
	captableEquityPlanIn := *openapiclient.NewCaptableEquityPlanIn() // CaptableEquityPlanIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptablePlans(context.Background()).CaptableEquityPlanIn(captableEquityPlanIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptablePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptablePlans`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptablePlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptablePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableEquityPlanIn** | [**CaptableEquityPlanIn**](CaptableEquityPlanIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableRounds

> CaptableCreated PostCaptableRounds(ctx).CaptableRoundIn(captableRoundIn).Execute()

Opens a priced round that investments can be added to.



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
	captableRoundIn := *openapiclient.NewCaptableRoundIn() // CaptableRoundIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableRounds(context.Background()).CaptableRoundIn(captableRoundIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableRounds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableRounds`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableRounds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableRoundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableRoundIn** | [**CaptableRoundIn**](CaptableRoundIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableRoundsByIdClose

> CaptableUpdated PostCaptableRoundsByIdClose(ctx, id).CaptableRoundCloseRequest(captableRoundCloseRequest).Execute()

Closes one of the caller org's fundraising rounds, recording the close date and moving its status to CLOSED.



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
	id := "id_example" // string | ID is the round to close. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller's principal, so an id from another tenant is simply not found.
	captableRoundCloseRequest := *openapiclient.NewCaptableRoundCloseRequest() // CaptableRoundCloseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableRoundsByIdClose(context.Background(), id).CaptableRoundCloseRequest(captableRoundCloseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableRoundsByIdClose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableRoundsByIdClose`: CaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableRoundsByIdClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the round to close. It is the path segment: the URL is the addressing authority, and the org it is resolved in comes from the caller&#39;s principal, so an id from another tenant is simply not found. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableRoundsByIdCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captableRoundCloseRequest** | [**CaptableRoundCloseRequest**](CaptableRoundCloseRequest.md) |  | 

### Return type

[**CaptableUpdated**](CaptableUpdated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableRoundsByIdInvestments

> CaptableInvested PostCaptableRoundsByIdInvestments(ctx, id).CaptableInvestmentIn(captableInvestmentIn).Execute()

Records one investor's money into an open round.



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
	id := "id_example" // string | ID is the round to invest in. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which round is written whatever a body claims.
	captableInvestmentIn := *openapiclient.NewCaptableInvestmentIn() // CaptableInvestmentIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableRoundsByIdInvestments(context.Background(), id).CaptableInvestmentIn(captableInvestmentIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableRoundsByIdInvestments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableRoundsByIdInvestments`: CaptableInvested
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableRoundsByIdInvestments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the round to invest in. The URL is the addressing authority — a path segment binds after the body and after the query — so the address decides which round is written whatever a body claims. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableRoundsByIdInvestmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **captableInvestmentIn** | [**CaptableInvestmentIn**](CaptableInvestmentIn.md) |  | 

### Return type

[**CaptableInvested**](CaptableInvested.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableSafes

> CaptableCreated PostCaptableSafes(ctx).CaptableSafeIn(captableSafeIn).Execute()

Records a SAFE — a simple agreement for future equity.



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
	captableSafeIn := *openapiclient.NewCaptableSafeIn() // CaptableSafeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableSafes(context.Background()).CaptableSafeIn(captableSafeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableSafes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableSafes`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableSafes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableSafesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableSafeIn** | [**CaptableSafeIn**](CaptableSafeIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableShares

> CaptableCreated PostCaptableShares(ctx).CaptableShareIn(captableShareIn).Execute()

Issues a share certificate to a stakeholder.



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
	captableShareIn := *openapiclient.NewCaptableShareIn() // CaptableShareIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableShares(context.Background()).CaptableShareIn(captableShareIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableShares`: CaptableCreated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableShares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableSharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableShareIn** | [**CaptableShareIn**](CaptableShareIn.md) |  | 

### Return type

[**CaptableCreated**](CaptableCreated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableSharesTransfer

> CaptableTransferred PostCaptableSharesTransfer(ctx).CaptableShareTransfer(captableShareTransfer).Execute()

Moves shares from one stakeholder to another.



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
	captableShareTransfer := *openapiclient.NewCaptableShareTransfer() // CaptableShareTransfer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PostCaptableSharesTransfer(context.Background()).CaptableShareTransfer(captableShareTransfer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableSharesTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCaptableSharesTransfer`: CaptableTransferred
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PostCaptableSharesTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableSharesTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableShareTransfer** | [**CaptableShareTransfer**](CaptableShareTransfer.md) |  | 

### Return type

[**CaptableTransferred**](CaptableTransferred.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCaptableStakeholders

> PostCaptableStakeholders(ctx).Execute()

Add stakeholders to the cap table



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
	r, err := apiClient.CaptableAPI.PostCaptableStakeholders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PostCaptableStakeholders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCaptableStakeholdersRequest struct via the builder pattern


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


## PutCaptableCompany

> CaptableUpdated PutCaptableCompany(ctx).CaptableCompanyUpdate(captableCompanyUpdate).Execute()

Sets the caller org's company name and incorporation details.



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
	captableCompanyUpdate := *openapiclient.NewCaptableCompanyUpdate() // CaptableCompanyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CaptableAPI.PutCaptableCompany(context.Background()).CaptableCompanyUpdate(captableCompanyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptableAPI.PutCaptableCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCaptableCompany`: CaptableUpdated
	fmt.Fprintf(os.Stdout, "Response from `CaptableAPI.PutCaptableCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCaptableCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **captableCompanyUpdate** | [**CaptableCompanyUpdate**](CaptableCompanyUpdate.md) |  | 

### Return type

[**CaptableUpdated**](CaptableUpdated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

