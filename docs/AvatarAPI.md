# \AvatarAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvatarByOrgByUserByDigest**](AvatarAPI.md#GetAvatarByOrgByUserByDigest) | **Get** /v1/avatar/{org}/{user}/{digest} | Fetch a profile photo
[**PostAvatar**](AvatarAPI.md#PostAvatar) | **Post** /v1/avatar | Set your profile photo



## GetAvatarByOrgByUserByDigest

> GetAvatarByOrgByUserByDigest(ctx, org, user, digest).Execute()

Fetch a profile photo



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
	org := "org_example" // string | 
	user := "user_example" // string | 
	digest := "digest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarAPI.GetAvatarByOrgByUserByDigest(context.Background(), org, user, digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarAPI.GetAvatarByOrgByUserByDigest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**user** | **string** |  | 
**digest** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarByOrgByUserByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAvatar

> PostAvatar(ctx).Execute()

Set your profile photo



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
	r, err := apiClient.AvatarAPI.PostAvatar(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarAPI.PostAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAvatarRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

