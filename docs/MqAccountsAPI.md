# \MqAccountsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqGetAccount**](MqAccountsAPI.md#MqGetAccount) | **Get** /v1/mq/accounts/{id} | Get account info
[**MqListAccountConnections**](MqAccountsAPI.md#MqListAccountConnections) | **Get** /v1/mq/accounts/{id}/connections | List account connections
[**MqListAccounts**](MqAccountsAPI.md#MqListAccounts) | **Get** /v1/mq/accounts | List accounts



## MqGetAccount

> MqAccount MqGetAccount(ctx, id).Execute()

Get account info



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
	id := "id_example" // string | Account ID (maps to IAM org_id).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAccountsAPI.MqGetAccount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAccountsAPI.MqGetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetAccount`: MqAccount
	fmt.Fprintf(os.Stdout, "Response from `MqAccountsAPI.MqGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (maps to IAM org_id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqAccount**](MqAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListAccountConnections

> MqListAccountConnections200Response MqListAccountConnections(ctx, id).Limit(limit).Offset(offset).Execute()

List account connections



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
	id := "id_example" // string | Account ID (maps to IAM org_id).
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAccountsAPI.MqListAccountConnections(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAccountsAPI.MqListAccountConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListAccountConnections`: MqListAccountConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `MqAccountsAPI.MqListAccountConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID (maps to IAM org_id). | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqListAccountConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListAccountConnections200Response**](MqListAccountConnections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListAccounts

> MqListAccounts200Response MqListAccounts(ctx).Limit(limit).Offset(offset).Execute()

List accounts



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
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqAccountsAPI.MqListAccounts(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqAccountsAPI.MqListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListAccounts`: MqListAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `MqAccountsAPI.MqListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListAccounts200Response**](MqListAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

