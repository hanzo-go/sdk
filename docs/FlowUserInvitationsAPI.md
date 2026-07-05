# \FlowUserInvitationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowAcceptUserInvitation**](FlowUserInvitationsAPI.md#FlowAcceptUserInvitation) | **Post** /v1/flow/user-invitations/{id}/accept | Accept an invitation
[**FlowCreateUserInvitation**](FlowUserInvitationsAPI.md#FlowCreateUserInvitation) | **Post** /v1/flow/user-invitations | Invite a user to a project
[**FlowListUserInvitations**](FlowUserInvitationsAPI.md#FlowListUserInvitations) | **Get** /v1/flow/user-invitations | List pending invitations



## FlowAcceptUserInvitation

> map[string]interface{} FlowAcceptUserInvitation(ctx, id).Execute()

Accept an invitation

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
	resp, r, err := apiClient.FlowUserInvitationsAPI.FlowAcceptUserInvitation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowUserInvitationsAPI.FlowAcceptUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowAcceptUserInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowUserInvitationsAPI.FlowAcceptUserInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowAcceptUserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## FlowCreateUserInvitation

> map[string]interface{} FlowCreateUserInvitation(ctx).AutoCreateUserInvitationRequest(autoCreateUserInvitationRequest).Execute()

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
	resp, r, err := apiClient.FlowUserInvitationsAPI.FlowCreateUserInvitation(context.Background()).AutoCreateUserInvitationRequest(autoCreateUserInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowUserInvitationsAPI.FlowCreateUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowCreateUserInvitation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowUserInvitationsAPI.FlowCreateUserInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowCreateUserInvitationRequest struct via the builder pattern


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


## FlowListUserInvitations

> map[string]interface{} FlowListUserInvitations(ctx).Execute()

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
	resp, r, err := apiClient.FlowUserInvitationsAPI.FlowListUserInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowUserInvitationsAPI.FlowListUserInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListUserInvitations`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowUserInvitationsAPI.FlowListUserInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListUserInvitationsRequest struct via the builder pattern


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

