# \CrmAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmCompaniesById**](CrmAPI.md#DeleteCrmCompaniesById) | **Delete** /v1/crm/companies/{id} | Removes one of the caller org&#39;s companies and answers 204.
[**DeleteCrmContactsById**](CrmAPI.md#DeleteCrmContactsById) | **Delete** /v1/crm/contacts/{id} | Removes one of the caller org&#39;s contacts and answers 204.
[**DeleteCrmOpportunitiesById**](CrmAPI.md#DeleteCrmOpportunitiesById) | **Delete** /v1/crm/opportunities/{id} | Removes one of the caller org&#39;s deals and answers 204.
[**GetCrmApplications**](CrmAPI.md#GetCrmApplications) | **Get** /v1/crm/applications | Returns the org&#39;s Startup Program applications, newest first.
[**GetCrmApplicationsById**](CrmAPI.md#GetCrmApplicationsById) | **Get** /v1/crm/applications/{id} | Returns one Startup Program application with its AI screen and stage history.
[**GetCrmCompanies**](CrmAPI.md#GetCrmCompanies) | **Get** /v1/crm/companies | Returns the caller org&#39;s companies, most recently updated first.
[**GetCrmCompaniesById**](CrmAPI.md#GetCrmCompaniesById) | **Get** /v1/crm/companies/{id} | Returns one of the caller org&#39;s companies.
[**GetCrmContacts**](CrmAPI.md#GetCrmContacts) | **Get** /v1/crm/contacts | Returns the caller org&#39;s contacts, most recently updated first.
[**GetCrmContactsById**](CrmAPI.md#GetCrmContactsById) | **Get** /v1/crm/contacts/{id} | Returns one of the caller org&#39;s contacts.
[**GetCrmOpportunities**](CrmAPI.md#GetCrmOpportunities) | **Get** /v1/crm/opportunities | Returns the caller org&#39;s deals, most recently updated first.
[**GetCrmOpportunitiesById**](CrmAPI.md#GetCrmOpportunitiesById) | **Get** /v1/crm/opportunities/{id} | Returns one of the caller org&#39;s deals.
[**GetCrmSummary**](CrmAPI.md#GetCrmSummary) | **Get** /v1/crm/summary | Summary counts the caller org&#39;s CRM records: companies, contacts, opportunities.
[**PatchCrmApplicationsById**](CrmAPI.md#PatchCrmApplicationsById) | **Patch** /v1/crm/applications/{id} | Moves one Startup Program application through the pipeline.
[**PostCrmApplications**](CrmAPI.md#PostCrmApplications) | **Post** /v1/crm/applications | Apply to the Startup Program from the public form
[**PostCrmCompanies**](CrmAPI.md#PostCrmCompanies) | **Post** /v1/crm/companies | Adds a company to the caller&#39;s org and answers 201 with the stored record.
[**PostCrmContacts**](CrmAPI.md#PostCrmContacts) | **Post** /v1/crm/contacts | Adds a person to the caller&#39;s org and answers 201 with the stored record.
[**PostCrmOpportunities**](CrmAPI.md#PostCrmOpportunities) | **Post** /v1/crm/opportunities | Adds a deal to the caller&#39;s org and answers 201 with the stored record.
[**PutCrmCompaniesById**](CrmAPI.md#PutCrmCompaniesById) | **Put** /v1/crm/companies/{id} | Replaces one of the caller org&#39;s companies.
[**PutCrmContactsById**](CrmAPI.md#PutCrmContactsById) | **Put** /v1/crm/contacts/{id} | Replaces one of the caller org&#39;s contacts.
[**PutCrmOpportunitiesById**](CrmAPI.md#PutCrmOpportunitiesById) | **Put** /v1/crm/opportunities/{id} | Replaces one of the caller org&#39;s deals.



## DeleteCrmCompaniesById

> DeleteCrmCompaniesById(ctx, id).Execute()

Removes one of the caller org's companies and answers 204.



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
	id := "comp_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrmAPI.DeleteCrmCompaniesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.DeleteCrmCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmCompaniesByIdRequest struct via the builder pattern


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


## DeleteCrmContactsById

> DeleteCrmContactsById(ctx, id).Execute()

Removes one of the caller org's contacts and answers 204.



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
	id := "cont_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrmAPI.DeleteCrmContactsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.DeleteCrmContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmContactsByIdRequest struct via the builder pattern


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


## DeleteCrmOpportunitiesById

> DeleteCrmOpportunitiesById(ctx, id).Execute()

Removes one of the caller org's deals and answers 204.



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
	id := "oppo_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CrmAPI.DeleteCrmOpportunitiesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.DeleteCrmOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmOpportunitiesByIdRequest struct via the builder pattern


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


## GetCrmApplications

> ApplicationList GetCrmApplications(ctx).Stage(stage).Limit(limit).Execute()

Returns the org's Startup Program applications, newest first.



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
	stage := "stage_example" // string | Stage returns only the applications at that pipeline stage when set: applied, screened, qualified, credits-offered, onboarded or rejected. (optional)
	limit := int32(56) // int32 | Limit caps the rows returned: 200 by default, 1000 at most. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmApplications(context.Background()).Stage(stage).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmApplications`: ApplicationList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage returns only the applications at that pipeline stage when set: applied, screened, qualified, credits-offered, onboarded or rejected. | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**ApplicationList**](ApplicationList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmApplicationsById

> ProgramApplication GetCrmApplicationsById(ctx, id).Execute()

Returns one Startup Program application with its AI screen and stage history.



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
	id := "appl_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmApplicationsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmApplicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmApplicationsById`: ProgramApplication
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmApplicationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmApplicationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProgramApplication**](ProgramApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmCompanies

> CompanyList GetCrmCompanies(ctx).Limit(limit).Execute()

Returns the caller org's companies, most recently updated first.



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
	limit := int32(56) // int32 | Limit caps the rows returned: 200 by default, 1000 at most. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmCompanies(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmCompanies`: CompanyList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**CompanyList**](CompanyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmCompaniesById

> Company GetCrmCompaniesById(ctx, id).Execute()

Returns one of the caller org's companies.



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
	id := "comp_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmCompaniesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmCompaniesById`: Company
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmCompaniesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Company**](Company.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmContacts

> ContactList GetCrmContacts(ctx).CompanyId(companyId).Limit(limit).Execute()

Returns the caller org's contacts, most recently updated first.



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
	companyId := "companyId_example" // string | CompanyID returns only the contacts at that company when set. (optional)
	limit := int32(56) // int32 | Limit caps the rows returned: 200 by default, 1000 at most. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmContacts(context.Background()).CompanyId(companyId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmContacts`: ContactList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | CompanyID returns only the contacts at that company when set. | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**ContactList**](ContactList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmContactsById

> Contact GetCrmContactsById(ctx, id).Execute()

Returns one of the caller org's contacts.



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
	id := "cont_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmContactsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmContactsById`: Contact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmOpportunities

> OppList GetCrmOpportunities(ctx).Stage(stage).Limit(limit).Execute()

Returns the caller org's deals, most recently updated first.



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
	stage := "stage_example" // string | Stage returns only the opportunities at that pipeline stage when set (NEW, SCREENING, MEETING, PROPOSAL or CUSTOMER; case-insensitive). (optional)
	limit := int32(56) // int32 | Limit caps the rows returned: 200 by default, 1000 at most. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmOpportunities(context.Background()).Stage(stage).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmOpportunities`: OppList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage returns only the opportunities at that pipeline stage when set (NEW, SCREENING, MEETING, PROPOSAL or CUSTOMER; case-insensitive). | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**OppList**](OppList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmOpportunitiesById

> Opportunity GetCrmOpportunitiesById(ctx, id).Execute()

Returns one of the caller org's deals.



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
	id := "oppo_1" // string | ID is the record to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.GetCrmOpportunitiesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmOpportunitiesById`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmOpportunitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmOpportunitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Opportunity**](Opportunity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmSummary

> CrmSummary GetCrmSummary(ctx).Execute()

Summary counts the caller org's CRM records: companies, contacts, opportunities.



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
	resp, r, err := apiClient.CrmAPI.GetCrmSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.GetCrmSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmSummary`: CrmSummary
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.GetCrmSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmSummaryRequest struct via the builder pattern


### Return type

[**CrmSummary**](CrmSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmApplicationsById

> ProgramApplication PatchCrmApplicationsById(ctx, id).PatchApplicationIn(patchApplicationIn).Execute()

Moves one Startup Program application through the pipeline.



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
	id := "appl_1" // string | ID is the application to move, from the path.
	patchApplicationIn := *openapiclient.NewPatchApplicationIn() // PatchApplicationIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PatchCrmApplicationsById(context.Background(), id).PatchApplicationIn(patchApplicationIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PatchCrmApplicationsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmApplicationsById`: ProgramApplication
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PatchCrmApplicationsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the application to move, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmApplicationsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApplicationIn** | [**PatchApplicationIn**](PatchApplicationIn.md) |  | 

### Return type

[**ProgramApplication**](ProgramApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmApplications

> PostCrmApplications(ctx).Execute()

Apply to the Startup Program from the public form



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
	r, err := apiClient.CrmAPI.PostCrmApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PostCrmApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmApplicationsRequest struct via the builder pattern


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


## PostCrmCompanies

> Company PostCrmCompanies(ctx).CompanyReq(companyReq).Execute()

Adds a company to the caller's org and answers 201 with the stored record.



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
	companyReq := *openapiclient.NewCompanyReq() // CompanyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PostCrmCompanies(context.Background()).CompanyReq(companyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PostCrmCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmCompanies`: Company
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PostCrmCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyReq** | [**CompanyReq**](CompanyReq.md) |  | 

### Return type

[**Company**](Company.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmContacts

> Contact PostCrmContacts(ctx).ContactReq(contactReq).Execute()

Adds a person to the caller's org and answers 201 with the stored record.



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
	contactReq := *openapiclient.NewContactReq() // ContactReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PostCrmContacts(context.Background()).ContactReq(contactReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PostCrmContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmContacts`: Contact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PostCrmContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactReq** | [**ContactReq**](ContactReq.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmOpportunities

> Opportunity PostCrmOpportunities(ctx).OppReq(oppReq).Execute()

Adds a deal to the caller's org and answers 201 with the stored record.



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
	oppReq := *openapiclient.NewOppReq() // OppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PostCrmOpportunities(context.Background()).OppReq(oppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PostCrmOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmOpportunities`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PostCrmOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oppReq** | [**OppReq**](OppReq.md) |  | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmCompaniesById

> Company PutCrmCompaniesById(ctx, id).CompanyReq(companyReq).Execute()

Replaces one of the caller org's companies.



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
	id := "comp_1" // string | ID names the company to update and comes from the path. A create ignores it: the server mints the id.
	companyReq := *openapiclient.NewCompanyReq() // CompanyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PutCrmCompaniesById(context.Background(), id).CompanyReq(companyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PutCrmCompaniesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmCompaniesById`: Company
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PutCrmCompaniesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the company to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmCompaniesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyReq** | [**CompanyReq**](CompanyReq.md) |  | 

### Return type

[**Company**](Company.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmContactsById

> Contact PutCrmContactsById(ctx, id).ContactReq(contactReq).Execute()

Replaces one of the caller org's contacts.



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
	id := "cont_1" // string | ID names the contact to update and comes from the path. A create ignores it: the server mints the id.
	contactReq := *openapiclient.NewContactReq() // ContactReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PutCrmContactsById(context.Background(), id).ContactReq(contactReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PutCrmContactsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmContactsById`: Contact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PutCrmContactsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the contact to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmContactsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contactReq** | [**ContactReq**](ContactReq.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmOpportunitiesById

> Opportunity PutCrmOpportunitiesById(ctx, id).OppReq(oppReq).Execute()

Replaces one of the caller org's deals.



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
	id := "oppo_1" // string | ID names the opportunity to update and comes from the path. A create ignores it: the server mints the id.
	oppReq := *openapiclient.NewOppReq() // OppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.PutCrmOpportunitiesById(context.Background(), id).OppReq(oppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.PutCrmOpportunitiesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmOpportunitiesById`: Opportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.PutCrmOpportunitiesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the opportunity to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmOpportunitiesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oppReq** | [**OppReq**](OppReq.md) |  | 

### Return type

[**Opportunity**](Opportunity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

