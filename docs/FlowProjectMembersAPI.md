# \FlowProjectMembersAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowAddProjectMember**](FlowProjectMembersAPI.md#FlowAddProjectMember) | **Post** /v1/flow/project-members | Add a project member (EE)
[**FlowListProjectMembers**](FlowProjectMembersAPI.md#FlowListProjectMembers) | **Get** /v1/flow/project-members | List project members (EE)
[**FlowRemoveProjectMember**](FlowProjectMembersAPI.md#FlowRemoveProjectMember) | **Delete** /v1/flow/project-members/{id} | Remove a project member (EE)



## FlowAddProjectMember

> map[string]interface{} FlowAddProjectMember(ctx).FlowAddProjectMemberRequest(flowAddProjectMemberRequest).Execute()

Add a project member (EE)

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
	flowAddProjectMemberRequest := *openapiclient.NewFlowAddProjectMemberRequest("UserId_example", "ProjectRoleId_example") // FlowAddProjectMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowProjectMembersAPI.FlowAddProjectMember(context.Background()).FlowAddProjectMemberRequest(flowAddProjectMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowProjectMembersAPI.FlowAddProjectMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowAddProjectMember`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowProjectMembersAPI.FlowAddProjectMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowAddProjectMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flowAddProjectMemberRequest** | [**FlowAddProjectMemberRequest**](FlowAddProjectMemberRequest.md) |  | 

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


## FlowListProjectMembers

> map[string]interface{} FlowListProjectMembers(ctx).Execute()

List project members (EE)

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
	resp, r, err := apiClient.FlowProjectMembersAPI.FlowListProjectMembers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowProjectMembersAPI.FlowListProjectMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListProjectMembers`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowProjectMembersAPI.FlowListProjectMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFlowListProjectMembersRequest struct via the builder pattern


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


## FlowRemoveProjectMember

> FlowRemoveProjectMember(ctx, id).Execute()

Remove a project member (EE)

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
	r, err := apiClient.FlowProjectMembersAPI.FlowRemoveProjectMember(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowProjectMembersAPI.FlowRemoveProjectMember``: %v\n", err)
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

Other parameters are passed through a pointer to a apiFlowRemoveProjectMemberRequest struct via the builder pattern


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

