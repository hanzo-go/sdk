# \KmsApprovalPoliciesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateChangeApprovalPolicy**](KmsApprovalPoliciesAPI.md#KmsCreateChangeApprovalPolicy) | **Post** /v1/kms/approval-policies/change | Create a secret change approval policy
[**KmsListChangeApprovalPolicies**](KmsApprovalPoliciesAPI.md#KmsListChangeApprovalPolicies) | **Get** /v1/kms/approval-policies/change | List secret change approval policies



## KmsCreateChangeApprovalPolicy

> KmsCreateChangeApprovalPolicy200Response KmsCreateChangeApprovalPolicy(ctx).KmsCreateChangeApprovalPolicyRequest(kmsCreateChangeApprovalPolicyRequest).Execute()

Create a secret change approval policy

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
	kmsCreateChangeApprovalPolicyRequest := *openapiclient.NewKmsCreateChangeApprovalPolicyRequest("WorkspaceId_example", "Name_example", int32(123), "Environment_example") // KmsCreateChangeApprovalPolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsApprovalPoliciesAPI.KmsCreateChangeApprovalPolicy(context.Background()).KmsCreateChangeApprovalPolicyRequest(kmsCreateChangeApprovalPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsApprovalPoliciesAPI.KmsCreateChangeApprovalPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateChangeApprovalPolicy`: KmsCreateChangeApprovalPolicy200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsApprovalPoliciesAPI.KmsCreateChangeApprovalPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateChangeApprovalPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateChangeApprovalPolicyRequest** | [**KmsCreateChangeApprovalPolicyRequest**](KmsCreateChangeApprovalPolicyRequest.md) |  | 

### Return type

[**KmsCreateChangeApprovalPolicy200Response**](KmsCreateChangeApprovalPolicy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListChangeApprovalPolicies

> KmsListChangeApprovalPolicies200Response KmsListChangeApprovalPolicies(ctx).WorkspaceId(workspaceId).Execute()

List secret change approval policies

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsApprovalPoliciesAPI.KmsListChangeApprovalPolicies(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsApprovalPoliciesAPI.KmsListChangeApprovalPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListChangeApprovalPolicies`: KmsListChangeApprovalPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsApprovalPoliciesAPI.KmsListChangeApprovalPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListChangeApprovalPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 

### Return type

[**KmsListChangeApprovalPolicies200Response**](KmsListChangeApprovalPolicies200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

