# \KmsUsersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsGetCurrentUser**](KmsUsersAPI.md#KmsGetCurrentUser) | **Get** /v1/kms/user | Get the current authenticated user



## KmsGetCurrentUser

> KmsGetCurrentUser200Response KmsGetCurrentUser(ctx).Execute()

Get the current authenticated user

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
	resp, r, err := apiClient.KmsUsersAPI.KmsGetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsUsersAPI.KmsGetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetCurrentUser`: KmsGetCurrentUser200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsUsersAPI.KmsGetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetCurrentUserRequest struct via the builder pattern


### Return type

[**KmsGetCurrentUser200Response**](KmsGetCurrentUser200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

