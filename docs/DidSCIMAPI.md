# \DidSCIMAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DidScimCreateUser**](DidSCIMAPI.md#DidScimCreateUser) | **Post** /v1/did/scim/v2/Users | SCIM create user
[**DidScimListUsers**](DidSCIMAPI.md#DidScimListUsers) | **Get** /v1/did/scim/v2/Users | SCIM list users



## DidScimCreateUser

> DidScimCreateUser(ctx).Body(body).Execute()

SCIM create user



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DidSCIMAPI.DidScimCreateUser(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidSCIMAPI.DidScimCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDidScimCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/scim+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidScimListUsers

> map[string]interface{} DidScimListUsers(ctx).Filter(filter).Count(count).StartIndex(startIndex).Execute()

SCIM list users



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
	filter := "filter_example" // string | SCIM filter expression (optional)
	count := int32(56) // int32 |  (optional) (default to 100)
	startIndex := int32(56) // int32 |  (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidSCIMAPI.DidScimListUsers(context.Background()).Filter(filter).Count(count).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidSCIMAPI.DidScimListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidScimListUsers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DidSCIMAPI.DidScimListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDidScimListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | SCIM filter expression | 
 **count** | **int32** |  | [default to 100]
 **startIndex** | **int32** |  | [default to 1]

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/scim+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

