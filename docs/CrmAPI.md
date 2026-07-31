# \CrmAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1CrmCompaniesId**](CrmAPI.md#CloudDeleteV1CrmCompaniesId) | **Delete** /v1/crm/companies/{id} | DeleteCompany removes one of the caller org&#39;s companies and answers 204.
[**CloudDeleteV1CrmContactsId**](CrmAPI.md#CloudDeleteV1CrmContactsId) | **Delete** /v1/crm/contacts/{id} | DeleteContact removes one of the caller org&#39;s contacts and answers 204.
[**CloudDeleteV1CrmOpportunitiesId**](CrmAPI.md#CloudDeleteV1CrmOpportunitiesId) | **Delete** /v1/crm/opportunities/{id} | DeleteOpportunity removes one of the caller org&#39;s deals and answers 204.
[**CloudGetV1CrmApplications**](CrmAPI.md#CloudGetV1CrmApplications) | **Get** /v1/crm/applications | ListApplications returns the org&#39;s Startup Program applications, newest first.
[**CloudGetV1CrmApplicationsId**](CrmAPI.md#CloudGetV1CrmApplicationsId) | **Get** /v1/crm/applications/{id} | GetApplication returns one Startup Program application with its AI screen and stage history.
[**CloudGetV1CrmCompanies**](CrmAPI.md#CloudGetV1CrmCompanies) | **Get** /v1/crm/companies | ListCompanies returns the caller org&#39;s companies, most recently updated first.
[**CloudGetV1CrmCompaniesId**](CrmAPI.md#CloudGetV1CrmCompaniesId) | **Get** /v1/crm/companies/{id} | GetCompany returns one of the caller org&#39;s companies.
[**CloudGetV1CrmContacts**](CrmAPI.md#CloudGetV1CrmContacts) | **Get** /v1/crm/contacts | ListContacts returns the caller org&#39;s contacts, most recently updated first.
[**CloudGetV1CrmContactsId**](CrmAPI.md#CloudGetV1CrmContactsId) | **Get** /v1/crm/contacts/{id} | GetContact returns one of the caller org&#39;s contacts.
[**CloudGetV1CrmOpportunities**](CrmAPI.md#CloudGetV1CrmOpportunities) | **Get** /v1/crm/opportunities | ListOpportunities returns the caller org&#39;s deals, most recently updated first.
[**CloudGetV1CrmOpportunitiesId**](CrmAPI.md#CloudGetV1CrmOpportunitiesId) | **Get** /v1/crm/opportunities/{id} | GetOpportunity returns one of the caller org&#39;s deals.
[**CloudGetV1CrmSummary**](CrmAPI.md#CloudGetV1CrmSummary) | **Get** /v1/crm/summary | Summary counts the caller org&#39;s CRM records: companies, contacts, opportunities.
[**CloudPatchV1CrmApplicationsId**](CrmAPI.md#CloudPatchV1CrmApplicationsId) | **Patch** /v1/crm/applications/{id} | PatchApplication moves one Startup Program application through the pipeline.
[**CloudPostV1CrmApplications**](CrmAPI.md#CloudPostV1CrmApplications) | **Post** /v1/crm/applications | 
[**CloudPostV1CrmCompanies**](CrmAPI.md#CloudPostV1CrmCompanies) | **Post** /v1/crm/companies | CreateCompany adds a company to the caller&#39;s org and answers 201 with the stored record.
[**CloudPostV1CrmContacts**](CrmAPI.md#CloudPostV1CrmContacts) | **Post** /v1/crm/contacts | CreateContact adds a person to the caller&#39;s org and answers 201 with the stored record.
[**CloudPostV1CrmOpportunities**](CrmAPI.md#CloudPostV1CrmOpportunities) | **Post** /v1/crm/opportunities | CreateOpportunity adds a deal to the caller&#39;s org and answers 201 with the stored record.
[**CloudPutV1CrmCompaniesId**](CrmAPI.md#CloudPutV1CrmCompaniesId) | **Put** /v1/crm/companies/{id} | UpdateCompany replaces one of the caller org&#39;s companies.
[**CloudPutV1CrmContactsId**](CrmAPI.md#CloudPutV1CrmContactsId) | **Put** /v1/crm/contacts/{id} | UpdateContact replaces one of the caller org&#39;s contacts.
[**CloudPutV1CrmOpportunitiesId**](CrmAPI.md#CloudPutV1CrmOpportunitiesId) | **Put** /v1/crm/opportunities/{id} | UpdateOpportunity replaces one of the caller org&#39;s deals.



## CloudDeleteV1CrmCompaniesId

> CloudDeleteV1CrmCompaniesId(ctx, id).Execute()

DeleteCompany removes one of the caller org's companies and answers 204.



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
	r, err := apiClient.CrmAPI.CloudDeleteV1CrmCompaniesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudDeleteV1CrmCompaniesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1CrmCompaniesIdRequest struct via the builder pattern


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


## CloudDeleteV1CrmContactsId

> CloudDeleteV1CrmContactsId(ctx, id).Execute()

DeleteContact removes one of the caller org's contacts and answers 204.



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
	r, err := apiClient.CrmAPI.CloudDeleteV1CrmContactsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudDeleteV1CrmContactsId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1CrmContactsIdRequest struct via the builder pattern


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


## CloudDeleteV1CrmOpportunitiesId

> CloudDeleteV1CrmOpportunitiesId(ctx, id).Execute()

DeleteOpportunity removes one of the caller org's deals and answers 204.



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
	r, err := apiClient.CrmAPI.CloudDeleteV1CrmOpportunitiesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudDeleteV1CrmOpportunitiesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1CrmOpportunitiesIdRequest struct via the builder pattern


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


## CloudGetV1CrmApplications

> CloudApplicationList CloudGetV1CrmApplications(ctx).Stage(stage).Limit(limit).Execute()

ListApplications returns the org's Startup Program applications, newest first.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmApplications(context.Background()).Stage(stage).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmApplications`: CloudApplicationList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage returns only the applications at that pipeline stage when set: applied, screened, qualified, credits-offered, onboarded or rejected. | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**CloudApplicationList**](CloudApplicationList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmApplicationsId

> CloudApplication CloudGetV1CrmApplicationsId(ctx, id).Execute()

GetApplication returns one Startup Program application with its AI screen and stage history.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmApplicationsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmApplicationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmApplicationsId`: CloudApplication
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmApplicationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmApplicationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudApplication**](CloudApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmCompanies

> CloudCompanyList CloudGetV1CrmCompanies(ctx).Limit(limit).Execute()

ListCompanies returns the caller org's companies, most recently updated first.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmCompanies(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmCompanies`: CloudCompanyList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**CloudCompanyList**](CloudCompanyList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmCompaniesId

> CloudCompany CloudGetV1CrmCompaniesId(ctx, id).Execute()

GetCompany returns one of the caller org's companies.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmCompaniesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmCompaniesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmCompaniesId`: CloudCompany
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmCompaniesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmCompaniesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudCompany**](CloudCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmContacts

> CloudContactList CloudGetV1CrmContacts(ctx).CompanyId(companyId).Limit(limit).Execute()

ListContacts returns the caller org's contacts, most recently updated first.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmContacts(context.Background()).CompanyId(companyId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmContacts`: CloudContactList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | CompanyID returns only the contacts at that company when set. | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**CloudContactList**](CloudContactList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmContactsId

> CloudContact CloudGetV1CrmContactsId(ctx, id).Execute()

GetContact returns one of the caller org's contacts.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmContactsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmContactsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmContactsId`: CloudContact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmContactsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmContactsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudContact**](CloudContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmOpportunities

> CloudOppList CloudGetV1CrmOpportunities(ctx).Stage(stage).Limit(limit).Execute()

ListOpportunities returns the caller org's deals, most recently updated first.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmOpportunities(context.Background()).Stage(stage).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmOpportunities`: CloudOppList
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | **string** | Stage returns only the opportunities at that pipeline stage when set (NEW, SCREENING, MEETING, PROPOSAL or CUSTOMER; case-insensitive). | 
 **limit** | **int32** | Limit caps the rows returned: 200 by default, 1000 at most. | 

### Return type

[**CloudOppList**](CloudOppList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmOpportunitiesId

> CloudOpportunity CloudGetV1CrmOpportunitiesId(ctx, id).Execute()

GetOpportunity returns one of the caller org's deals.



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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmOpportunitiesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmOpportunitiesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmOpportunitiesId`: CloudOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmOpportunitiesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the record to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmOpportunitiesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudOpportunity**](CloudOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CrmSummary

> CloudCrmSummary CloudGetV1CrmSummary(ctx).Execute()

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
	resp, r, err := apiClient.CrmAPI.CloudGetV1CrmSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudGetV1CrmSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CrmSummary`: CloudCrmSummary
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudGetV1CrmSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CrmSummaryRequest struct via the builder pattern


### Return type

[**CloudCrmSummary**](CloudCrmSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1CrmApplicationsId

> CloudApplication CloudPatchV1CrmApplicationsId(ctx, id).CloudPatchApplicationIn(cloudPatchApplicationIn).Execute()

PatchApplication moves one Startup Program application through the pipeline.



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
	cloudPatchApplicationIn := *openapiclient.NewCloudPatchApplicationIn() // CloudPatchApplicationIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPatchV1CrmApplicationsId(context.Background(), id).CloudPatchApplicationIn(cloudPatchApplicationIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPatchV1CrmApplicationsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPatchV1CrmApplicationsId`: CloudApplication
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPatchV1CrmApplicationsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the application to move, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1CrmApplicationsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPatchApplicationIn** | [**CloudPatchApplicationIn**](CloudPatchApplicationIn.md) |  | 

### Return type

[**CloudApplication**](CloudApplication.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CrmApplications

> CloudPostV1CrmApplications(ctx).Execute()



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
	r, err := apiClient.CrmAPI.CloudPostV1CrmApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPostV1CrmApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CrmApplicationsRequest struct via the builder pattern


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


## CloudPostV1CrmCompanies

> CloudCompany CloudPostV1CrmCompanies(ctx).CloudCompanyReq(cloudCompanyReq).Execute()

CreateCompany adds a company to the caller's org and answers 201 with the stored record.



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
	cloudCompanyReq := *openapiclient.NewCloudCompanyReq() // CloudCompanyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPostV1CrmCompanies(context.Background()).CloudCompanyReq(cloudCompanyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPostV1CrmCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CrmCompanies`: CloudCompany
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPostV1CrmCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CrmCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCompanyReq** | [**CloudCompanyReq**](CloudCompanyReq.md) |  | 

### Return type

[**CloudCompany**](CloudCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CrmContacts

> CloudContact CloudPostV1CrmContacts(ctx).CloudContactReq(cloudContactReq).Execute()

CreateContact adds a person to the caller's org and answers 201 with the stored record.



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
	cloudContactReq := *openapiclient.NewCloudContactReq() // CloudContactReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPostV1CrmContacts(context.Background()).CloudContactReq(cloudContactReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPostV1CrmContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CrmContacts`: CloudContact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPostV1CrmContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CrmContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudContactReq** | [**CloudContactReq**](CloudContactReq.md) |  | 

### Return type

[**CloudContact**](CloudContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CrmOpportunities

> CloudOpportunity CloudPostV1CrmOpportunities(ctx).CloudOppReq(cloudOppReq).Execute()

CreateOpportunity adds a deal to the caller's org and answers 201 with the stored record.



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
	cloudOppReq := *openapiclient.NewCloudOppReq() // CloudOppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPostV1CrmOpportunities(context.Background()).CloudOppReq(cloudOppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPostV1CrmOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CrmOpportunities`: CloudOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPostV1CrmOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CrmOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudOppReq** | [**CloudOppReq**](CloudOppReq.md) |  | 

### Return type

[**CloudOpportunity**](CloudOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CrmCompaniesId

> CloudCompany CloudPutV1CrmCompaniesId(ctx, id).CloudCompanyReq(cloudCompanyReq).Execute()

UpdateCompany replaces one of the caller org's companies.



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
	cloudCompanyReq := *openapiclient.NewCloudCompanyReq() // CloudCompanyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPutV1CrmCompaniesId(context.Background(), id).CloudCompanyReq(cloudCompanyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPutV1CrmCompaniesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CrmCompaniesId`: CloudCompany
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPutV1CrmCompaniesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the company to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CrmCompaniesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCompanyReq** | [**CloudCompanyReq**](CloudCompanyReq.md) |  | 

### Return type

[**CloudCompany**](CloudCompany.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CrmContactsId

> CloudContact CloudPutV1CrmContactsId(ctx, id).CloudContactReq(cloudContactReq).Execute()

UpdateContact replaces one of the caller org's contacts.



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
	cloudContactReq := *openapiclient.NewCloudContactReq() // CloudContactReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPutV1CrmContactsId(context.Background(), id).CloudContactReq(cloudContactReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPutV1CrmContactsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CrmContactsId`: CloudContact
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPutV1CrmContactsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the contact to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CrmContactsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudContactReq** | [**CloudContactReq**](CloudContactReq.md) |  | 

### Return type

[**CloudContact**](CloudContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CrmOpportunitiesId

> CloudOpportunity CloudPutV1CrmOpportunitiesId(ctx, id).CloudOppReq(cloudOppReq).Execute()

UpdateOpportunity replaces one of the caller org's deals.



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
	cloudOppReq := *openapiclient.NewCloudOppReq() // CloudOppReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmAPI.CloudPutV1CrmOpportunitiesId(context.Background(), id).CloudOppReq(cloudOppReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmAPI.CloudPutV1CrmOpportunitiesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CrmOpportunitiesId`: CloudOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmAPI.CloudPutV1CrmOpportunitiesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID names the opportunity to update and comes from the path. A create ignores it: the server mints the id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CrmOpportunitiesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudOppReq** | [**CloudOppReq**](CloudOppReq.md) |  | 

### Return type

[**CloudOpportunity**](CloudOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

