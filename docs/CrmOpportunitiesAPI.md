# \CrmOpportunitiesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmCreateOpportunity**](CrmOpportunitiesAPI.md#CrmCreateOpportunity) | **Post** /v1/crm/opportunities | Create an opportunity
[**CrmDeleteOpportunity**](CrmOpportunitiesAPI.md#CrmDeleteOpportunity) | **Delete** /v1/crm/opportunities/{id} | Delete an opportunity
[**CrmGetOpportunity**](CrmOpportunitiesAPI.md#CrmGetOpportunity) | **Get** /v1/crm/opportunities/{id} | Opportunity detail
[**CrmListOpportunities**](CrmOpportunitiesAPI.md#CrmListOpportunities) | **Get** /v1/crm/opportunities | List opportunities
[**CrmUpdateOpportunity**](CrmOpportunitiesAPI.md#CrmUpdateOpportunity) | **Put** /v1/crm/opportunities/{id} | Update an opportunity



## CrmCreateOpportunity

> CrmOpportunity CrmCreateOpportunity(ctx).CrmOpportunityInput(crmOpportunityInput).Execute()

Create an opportunity

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
	crmOpportunityInput := *openapiclient.NewCrmOpportunityInput("Name_example") // CrmOpportunityInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmOpportunitiesAPI.CrmCreateOpportunity(context.Background()).CrmOpportunityInput(crmOpportunityInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmOpportunitiesAPI.CrmCreateOpportunity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmCreateOpportunity`: CrmOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmOpportunitiesAPI.CrmCreateOpportunity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmCreateOpportunityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crmOpportunityInput** | [**CrmOpportunityInput**](CrmOpportunityInput.md) |  | 

### Return type

[**CrmOpportunity**](CrmOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmDeleteOpportunity

> CrmDeleteOpportunity(ctx, id).Execute()

Delete an opportunity

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
	r, err := apiClient.CrmOpportunitiesAPI.CrmDeleteOpportunity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmOpportunitiesAPI.CrmDeleteOpportunity``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCrmDeleteOpportunityRequest struct via the builder pattern


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


## CrmGetOpportunity

> CrmOpportunity CrmGetOpportunity(ctx, id).Execute()

Opportunity detail

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
	resp, r, err := apiClient.CrmOpportunitiesAPI.CrmGetOpportunity(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmOpportunitiesAPI.CrmGetOpportunity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmGetOpportunity`: CrmOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmOpportunitiesAPI.CrmGetOpportunity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmGetOpportunityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrmOpportunity**](CrmOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmListOpportunities

> CrmListOpportunities200Response CrmListOpportunities(ctx).Stage(stage).Limit(limit).Execute()

List opportunities

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
	stage := openapiclient.crm_Stage("NEW") // CrmStage | Filter to one pipeline stage (optional)
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmOpportunitiesAPI.CrmListOpportunities(context.Background()).Stage(stage).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmOpportunitiesAPI.CrmListOpportunities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmListOpportunities`: CrmListOpportunities200Response
	fmt.Fprintf(os.Stdout, "Response from `CrmOpportunitiesAPI.CrmListOpportunities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmListOpportunitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stage** | [**CrmStage**](CrmStage.md) | Filter to one pipeline stage | 
 **limit** | **int32** |  | [default to 200]

### Return type

[**CrmListOpportunities200Response**](CrmListOpportunities200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmUpdateOpportunity

> CrmOpportunity CrmUpdateOpportunity(ctx, id).CrmOpportunityInput(crmOpportunityInput).Execute()

Update an opportunity

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
	crmOpportunityInput := *openapiclient.NewCrmOpportunityInput("Name_example") // CrmOpportunityInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmOpportunitiesAPI.CrmUpdateOpportunity(context.Background(), id).CrmOpportunityInput(crmOpportunityInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmOpportunitiesAPI.CrmUpdateOpportunity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmUpdateOpportunity`: CrmOpportunity
	fmt.Fprintf(os.Stdout, "Response from `CrmOpportunitiesAPI.CrmUpdateOpportunity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmUpdateOpportunityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crmOpportunityInput** | [**CrmOpportunityInput**](CrmOpportunityInput.md) |  | 

### Return type

[**CrmOpportunity**](CrmOpportunity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

