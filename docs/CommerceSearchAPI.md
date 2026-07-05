# \CommerceSearchAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommerceSearchNotes**](CommerceSearchAPI.md#CommerceSearchNotes) | **Post** /v1/commerce/search/note | Search notes
[**CommerceSearchOrders**](CommerceSearchAPI.md#CommerceSearchOrders) | **Get** /v1/commerce/search/order | Search orders
[**CommerceSearchUsers**](CommerceSearchAPI.md#CommerceSearchUsers) | **Get** /v1/commerce/search/user | Search users



## CommerceSearchNotes

> []CommerceNote CommerceSearchNotes(ctx).CommerceSearchNotesRequest(commerceSearchNotesRequest).Execute()

Search notes

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
	commerceSearchNotesRequest := *openapiclient.NewCommerceSearchNotesRequest() // CommerceSearchNotesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceSearchAPI.CommerceSearchNotes(context.Background()).CommerceSearchNotesRequest(commerceSearchNotesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSearchAPI.CommerceSearchNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchNotes`: []CommerceNote
	fmt.Fprintf(os.Stdout, "Response from `CommerceSearchAPI.CommerceSearchNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commerceSearchNotesRequest** | [**CommerceSearchNotesRequest**](CommerceSearchNotesRequest.md) |  | 

### Return type

[**[]CommerceNote**](CommerceNote.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSearchOrders

> []CommerceOrder CommerceSearchOrders(ctx).Q(q).Execute()

Search orders

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
	q := "q_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceSearchAPI.CommerceSearchOrders(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSearchAPI.CommerceSearchOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchOrders`: []CommerceOrder
	fmt.Fprintf(os.Stdout, "Response from `CommerceSearchAPI.CommerceSearchOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 

### Return type

[**[]CommerceOrder**](CommerceOrder.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommerceSearchUsers

> []CommerceUser CommerceSearchUsers(ctx).Q(q).Execute()

Search users

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
	q := "q_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommerceSearchAPI.CommerceSearchUsers(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommerceSearchAPI.CommerceSearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommerceSearchUsers`: []CommerceUser
	fmt.Fprintf(os.Stdout, "Response from `CommerceSearchAPI.CommerceSearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommerceSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 

### Return type

[**[]CommerceUser**](CommerceUser.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

