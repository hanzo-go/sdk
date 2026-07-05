# \CrmContactsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmCreateContact**](CrmContactsAPI.md#CrmCreateContact) | **Post** /v1/crm/contacts | Create a contact
[**CrmDeleteContact**](CrmContactsAPI.md#CrmDeleteContact) | **Delete** /v1/crm/contacts/{id} | Delete a contact (clears opportunity point-of-contact refs)
[**CrmGetContact**](CrmContactsAPI.md#CrmGetContact) | **Get** /v1/crm/contacts/{id} | Contact detail
[**CrmListContacts**](CrmContactsAPI.md#CrmListContacts) | **Get** /v1/crm/contacts | List contacts
[**CrmUpdateContact**](CrmContactsAPI.md#CrmUpdateContact) | **Put** /v1/crm/contacts/{id} | Update a contact



## CrmCreateContact

> CrmContact CrmCreateContact(ctx).CrmContactInput(crmContactInput).Execute()

Create a contact

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
	crmContactInput := *openapiclient.NewCrmContactInput() // CrmContactInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmContactsAPI.CrmCreateContact(context.Background()).CrmContactInput(crmContactInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmContactsAPI.CrmCreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmCreateContact`: CrmContact
	fmt.Fprintf(os.Stdout, "Response from `CrmContactsAPI.CrmCreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crmContactInput** | [**CrmContactInput**](CrmContactInput.md) |  | 

### Return type

[**CrmContact**](CrmContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmDeleteContact

> CrmDeleteContact(ctx, id).Execute()

Delete a contact (clears opportunity point-of-contact refs)

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
	r, err := apiClient.CrmContactsAPI.CrmDeleteContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmContactsAPI.CrmDeleteContact``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCrmDeleteContactRequest struct via the builder pattern


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


## CrmGetContact

> CrmContact CrmGetContact(ctx, id).Execute()

Contact detail

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
	resp, r, err := apiClient.CrmContactsAPI.CrmGetContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmContactsAPI.CrmGetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmGetContact`: CrmContact
	fmt.Fprintf(os.Stdout, "Response from `CrmContactsAPI.CrmGetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CrmContact**](CrmContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmListContacts

> CrmListContacts200Response CrmListContacts(ctx).CompanyId(companyId).Limit(limit).Execute()

List contacts

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
	companyId := "companyId_example" // string | Filter to one company (optional)
	limit := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmContactsAPI.CrmListContacts(context.Background()).CompanyId(companyId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmContactsAPI.CrmListContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmListContacts`: CrmListContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `CrmContactsAPI.CrmListContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCrmListContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **string** | Filter to one company | 
 **limit** | **int32** |  | [default to 200]

### Return type

[**CrmListContacts200Response**](CrmListContacts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CrmUpdateContact

> CrmContact CrmUpdateContact(ctx, id).CrmContactInput(crmContactInput).Execute()

Update a contact

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
	crmContactInput := *openapiclient.NewCrmContactInput() // CrmContactInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CrmContactsAPI.CrmUpdateContact(context.Background(), id).CrmContactInput(crmContactInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmContactsAPI.CrmUpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmUpdateContact`: CrmContact
	fmt.Fprintf(os.Stdout, "Response from `CrmContactsAPI.CrmUpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCrmUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **crmContactInput** | [**CrmContactInput**](CrmContactInput.md) |  | 

### Return type

[**CrmContact**](CrmContact.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

