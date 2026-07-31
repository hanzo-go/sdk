# \PoliciesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthzAuthzAddPolicy**](PoliciesAPI.md#AuthzAuthzAddPolicy) | **Post** /v1/authz/policies | Add a policy
[**AuthzAuthzListPolicies**](PoliciesAPI.md#AuthzAuthzListPolicies) | **Get** /v1/authz/policies | List policies
[**AuthzAuthzRemovePolicy**](PoliciesAPI.md#AuthzAuthzRemovePolicy) | **Delete** /v1/authz/policies | Remove a policy
[**S3GetBucketPolicy**](PoliciesAPI.md#S3GetBucketPolicy) | **Get** /v1/s3/{bucket}?policy | Get bucket policy
[**S3PutBucketPolicy**](PoliciesAPI.md#S3PutBucketPolicy) | **Put** /v1/s3/{bucket}?policy | Set bucket policy



## AuthzAuthzAddPolicy

> AuthzAddPolicyResponse AuthzAuthzAddPolicy(ctx).AuthzEnforceRequest(authzEnforceRequest).Execute()

Add a policy



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
	authzEnforceRequest := *openapiclient.NewAuthzEnforceRequest("Sub_example", "Obj_example", "Act_example") // AuthzEnforceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AuthzAuthzAddPolicy(context.Background()).AuthzEnforceRequest(authzEnforceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AuthzAuthzAddPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzAddPolicy`: AuthzAddPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AuthzAuthzAddPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzAddPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authzEnforceRequest** | [**AuthzEnforceRequest**](AuthzEnforceRequest.md) |  | 

### Return type

[**AuthzAddPolicyResponse**](AuthzAddPolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthzAuthzListPolicies

> AuthzPolicyListResponse AuthzAuthzListPolicies(ctx).Execute()

List policies



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
	resp, r, err := apiClient.PoliciesAPI.AuthzAuthzListPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AuthzAuthzListPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzListPolicies`: AuthzPolicyListResponse
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AuthzAuthzListPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzListPoliciesRequest struct via the builder pattern


### Return type

[**AuthzPolicyListResponse**](AuthzPolicyListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthzAuthzRemovePolicy

> AuthzRemovePolicyResponse AuthzAuthzRemovePolicy(ctx).AuthzEnforceRequest(authzEnforceRequest).Execute()

Remove a policy



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
	authzEnforceRequest := *openapiclient.NewAuthzEnforceRequest("Sub_example", "Obj_example", "Act_example") // AuthzEnforceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.AuthzAuthzRemovePolicy(context.Background()).AuthzEnforceRequest(authzEnforceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.AuthzAuthzRemovePolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthzAuthzRemovePolicy`: AuthzRemovePolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.AuthzAuthzRemovePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthzAuthzRemovePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authzEnforceRequest** | [**AuthzEnforceRequest**](AuthzEnforceRequest.md) |  | 

### Return type

[**AuthzRemovePolicyResponse**](AuthzRemovePolicyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3GetBucketPolicy

> S3BucketPolicy S3GetBucketPolicy(ctx, bucket).Execute()

Get bucket policy

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
	bucket := "bucket_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoliciesAPI.S3GetBucketPolicy(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.S3GetBucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3GetBucketPolicy`: S3BucketPolicy
	fmt.Fprintf(os.Stdout, "Response from `PoliciesAPI.S3GetBucketPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3GetBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3BucketPolicy**](S3BucketPolicy.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3PutBucketPolicy

> S3PutBucketPolicy(ctx, bucket).S3BucketPolicy(s3BucketPolicy).Execute()

Set bucket policy

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
	bucket := "bucket_example" // string | 
	s3BucketPolicy := *openapiclient.NewS3BucketPolicy() // S3BucketPolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoliciesAPI.S3PutBucketPolicy(context.Background(), bucket).S3BucketPolicy(s3BucketPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoliciesAPI.S3PutBucketPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3PutBucketPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3BucketPolicy** | [**S3BucketPolicy**](S3BucketPolicy.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

