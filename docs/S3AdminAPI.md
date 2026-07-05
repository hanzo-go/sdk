# \S3AdminAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3AdminInfo**](S3AdminAPI.md#S3AdminInfo) | **Get** /v1/s3/admin/info | Server information
[**S3AdminUsage**](S3AdminAPI.md#S3AdminUsage) | **Get** /v1/s3/admin/usage | Storage usage
[**S3CreateServiceAccount**](S3AdminAPI.md#S3CreateServiceAccount) | **Post** /v1/s3/admin/service-accounts | Create a service account
[**S3ListServiceAccounts**](S3AdminAPI.md#S3ListServiceAccounts) | **Get** /v1/s3/admin/service-accounts | List service accounts



## S3AdminInfo

> S3AdminInfo200Response S3AdminInfo(ctx).Execute()

Server information



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
	resp, r, err := apiClient.S3AdminAPI.S3AdminInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3AdminAPI.S3AdminInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminInfo`: S3AdminInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `S3AdminAPI.S3AdminInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminInfoRequest struct via the builder pattern


### Return type

[**S3AdminInfo200Response**](S3AdminInfo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3AdminUsage

> S3UsageInfo S3AdminUsage(ctx).Execute()

Storage usage



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
	resp, r, err := apiClient.S3AdminAPI.S3AdminUsage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3AdminAPI.S3AdminUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3AdminUsage`: S3UsageInfo
	fmt.Fprintf(os.Stdout, "Response from `S3AdminAPI.S3AdminUsage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3AdminUsageRequest struct via the builder pattern


### Return type

[**S3UsageInfo**](S3UsageInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3CreateServiceAccount

> S3ServiceAccount S3CreateServiceAccount(ctx).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()

Create a service account



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
	s3CreateServiceAccountRequest := *openapiclient.NewS3CreateServiceAccountRequest() // S3CreateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.S3AdminAPI.S3CreateServiceAccount(context.Background()).S3CreateServiceAccountRequest(s3CreateServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3AdminAPI.S3CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3CreateServiceAccount`: S3ServiceAccount
	fmt.Fprintf(os.Stdout, "Response from `S3AdminAPI.S3CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiS3CreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3CreateServiceAccountRequest** | [**S3CreateServiceAccountRequest**](S3CreateServiceAccountRequest.md) |  | 

### Return type

[**S3ServiceAccount**](S3ServiceAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ListServiceAccounts

> S3ListServiceAccounts200Response S3ListServiceAccounts(ctx).Execute()

List service accounts

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
	resp, r, err := apiClient.S3AdminAPI.S3ListServiceAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `S3AdminAPI.S3ListServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `S3ListServiceAccounts`: S3ListServiceAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `S3AdminAPI.S3ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ListServiceAccountsRequest struct via the builder pattern


### Return type

[**S3ListServiceAccounts200Response**](S3ListServiceAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

