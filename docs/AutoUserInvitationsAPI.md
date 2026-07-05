# \AutoUserInvitationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoCreateUserInvitation**](AutoUserInvitationsAPI.md#AutoCreateUserInvitation) | **Post** /v1/auto/user-invitations | Invite a user to a project
[**AutoListUserInvitations**](AutoUserInvitationsAPI.md#AutoListUserInvitations) | **Get** /v1/auto/user-invitations | List pending invitations



## AutoCreateUserInvitation

> map[string]interface{} AutoCreateUserInvitation(ctx).AutoCreateUserInvitationRequest(autoCreateUserInvitationRequest).Execute()

Invite a user to a project

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
	autoCreateUserInvitationRequest := *openapiclient.NewAutoCreateUserInvitationRequest("Email_example", "ProjectRoleId_example") // AutoCreateUserInvitationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoUserInvitationsAPI.AutoCreateUserInvitation(context.Background()).AutoCreateUserInvitationRequest(autoCreateUserInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoUserInvitationsAPI.AutoCreateUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoCreateUserInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoUserInvitationsAPI.AutoCreateUserInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoCreateUserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoCreateUserInvitationRequest** | [**AutoCreateUserInvitationRequest**](AutoCreateUserInvitationRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutoListUserInvitations

> map[string]interface{} AutoListUserInvitations(ctx).Execute()

List pending invitations

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
	resp, r, err := apiClient.AutoUserInvitationsAPI.AutoListUserInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoUserInvitationsAPI.AutoListUserInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListUserInvitations`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoUserInvitationsAPI.AutoListUserInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListUserInvitationsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

