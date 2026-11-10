# \AccountAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountKeysById**](AccountAPI.md#DeleteAccountKeysById) | **Delete** /v1/account/keys/{id} | Revokes exactly one API key, by id.
[**GetAccountAppearance**](AccountAPI.md#GetAccountAppearance) | **Get** /v1/account/appearance | Returns the signed-in caller&#39;s own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.
[**GetAccountAvatarByOrgByUserByDigest**](AccountAPI.md#GetAccountAvatarByOrgByUserByDigest) | **Get** /v1/account/avatar/{org}/{user}/{digest} | Fetch a profile photo
[**GetAccountCsrf**](AccountAPI.md#GetAccountCsrf) | **Get** /v1/account/csrf | Mints the anti-forgery token a browser echoes as X-CSRF-Token on every change it asks for.
[**GetAccountEmbed**](AccountAPI.md#GetAccountEmbed) | **Get** /v1/account/embed | Reports whether one of this brand&#39;s shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.
[**GetAccountKeys**](AccountAPI.md#GetAccountKeys) | **Get** /v1/account/keys | Returns the API keys the caller may see, newest first: their own, or every key the org holds when the caller administers it.
[**PatchAccountKeysById**](AccountAPI.md#PatchAccountKeysById) | **Patch** /v1/account/keys/{id} | Changes one key&#39;s name, permissions, budget, rate or expiry.
[**PostAccountAppearance**](AccountAPI.md#PostAccountAppearance) | **Post** /v1/account/appearance | Stores the caller&#39;s appearance preference on their IAM account, preserving every other field of the row.
[**PostAccountAvatar**](AccountAPI.md#PostAccountAvatar) | **Post** /v1/account/avatar | Set your profile photo
[**PostAccountKeys**](AccountAPI.md#PostAccountKeys) | **Post** /v1/account/keys | Creates a NEW API key and returns its secret ONCE.
[**PostAccountOrgs**](AccountAPI.md#PostAccountOrgs) | **Post** /v1/account/orgs | Creates an organization and makes the caller its owner.



## DeleteAccountKeysById

> AccountApiKey DeleteAccountKeysById(ctx, id).Execute()

Revokes exactly one API key, by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the key to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DeleteAccountKeysById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountKeysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountKeysById`: AccountApiKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccountKeysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountKeysByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountApiKey**](AccountApiKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAppearance

> AccountAppearance GetAccountAppearance(ctx).Execute()

Returns the signed-in caller's own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountAppearance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAppearance`: AccountAppearance
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountAppearance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAppearanceRequest struct via the builder pattern


### Return type

[**AccountAppearance**](AccountAppearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAvatarByOrgByUserByDigest

> GetAccountAvatarByOrgByUserByDigest(ctx, org, user, digest).Execute()

Fetch a profile photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	org := "org_example" // string | 
	user := "user_example" // string | 
	digest := "digest_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.GetAccountAvatarByOrgByUserByDigest(context.Background(), org, user, digest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountAvatarByOrgByUserByDigest``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetAccountAvatarByOrgByUserByDigestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCsrf

> AccountCsrfResp GetAccountCsrf(ctx).Execute()

Mints the anti-forgery token a browser echoes as X-CSRF-Token on every change it asks for.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountCsrf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCsrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCsrf`: AccountCsrfResp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCsrf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCsrfRequest struct via the builder pattern


### Return type

[**AccountCsrfResp**](AccountCsrfResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountEmbed

> AccountEmbedStatusResp GetAccountEmbed(ctx).App(app).Execute()

Reports whether one of this brand's shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	app := "cms" // string | App is the embedded app to report on: cms (Content Studio), erp or help. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountEmbed(context.Background()).App(app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountEmbed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEmbed`: AccountEmbedStatusResp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountEmbed`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEmbedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **app** | **string** | App is the embedded app to report on: cms (Content Studio), erp or help. | 

### Return type

[**AccountEmbedStatusResp**](AccountEmbedStatusResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountKeys

> AccountApiKeyList GetAccountKeys(ctx).Execute()

Returns the API keys the caller may see, newest first: their own, or every key the org holds when the caller administers it.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountKeys`: AccountApiKeyList
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountKeysRequest struct via the builder pattern


### Return type

[**AccountApiKeyList**](AccountApiKeyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccountKeysById

> AccountApiKey PatchAccountKeysById(ctx, id).AccountKeyEdit(accountKeyEdit).Execute()

Changes one key's name, permissions, budget, rate or expiry.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	id := "id_example" // string | ID is the key to edit, from the path.
	accountKeyEdit := *openapiclient.NewAccountKeyEdit() // AccountKeyEdit | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PatchAccountKeysById(context.Background(), id).AccountKeyEdit(accountKeyEdit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PatchAccountKeysById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAccountKeysById`: AccountApiKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PatchAccountKeysById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the key to edit, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountKeysByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountKeyEdit** | [**AccountKeyEdit**](AccountKeyEdit.md) |  | 

### Return type

[**AccountApiKey**](AccountApiKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountAppearance

> AccountAppearance PostAccountAppearance(ctx).AccountAppearance(accountAppearance).Execute()

Stores the caller's appearance preference on their IAM account, preserving every other field of the row.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	accountAppearance := *openapiclient.NewAccountAppearance() // AccountAppearance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountAppearance(context.Background()).AccountAppearance(accountAppearance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountAppearance`: AccountAppearance
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountAppearance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountAppearanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountAppearance** | [**AccountAppearance**](AccountAppearance.md) |  | 

### Return type

[**AccountAppearance**](AccountAppearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountAvatar

> PostAccountAvatar(ctx).Execute()

Set your profile photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountAPI.PostAccountAvatar(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountAvatarRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountKeys

> AccountApiKey PostAccountKeys(ctx).AccountKeyIn(accountKeyIn).Execute()

Creates a NEW API key and returns its secret ONCE.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	accountKeyIn := *openapiclient.NewAccountKeyIn() // AccountKeyIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountKeys(context.Background()).AccountKeyIn(accountKeyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountKeys`: AccountApiKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountKeyIn** | [**AccountKeyIn**](AccountKeyIn.md) |  | 

### Return type

[**AccountApiKey**](AccountApiKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountOrgs

> AccountOnboardResp PostAccountOrgs(ctx).AccountOnboardReq(accountOnboardReq).Execute()

Creates an organization and makes the caller its owner.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {
	accountOnboardReq := *openapiclient.NewAccountOnboardReq() // AccountOnboardReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountOrgs(context.Background()).AccountOnboardReq(accountOnboardReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountOrgs`: AccountOnboardResp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountOnboardReq** | [**AccountOnboardReq**](AccountOnboardReq.md) |  | 

### Return type

[**AccountOnboardResp**](AccountOnboardResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

