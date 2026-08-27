# \AccountAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccountKeys**](AccountAPI.md#DeleteAccountKeys) | **Delete** /v1/account/keys | Revokes the caller&#39;s own API key of the requested class.
[**GetAccountAppearance**](AccountAPI.md#GetAccountAppearance) | **Get** /v1/account/appearance | Returns the signed-in caller&#39;s own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.
[**GetAccountAvatarByOrgByUserByDigest**](AccountAPI.md#GetAccountAvatarByOrgByUserByDigest) | **Get** /v1/account/avatar/{org}/{user}/{digest} | Fetch a profile photo
[**GetAccountCsrf**](AccountAPI.md#GetAccountCsrf) | **Get** /v1/account/csrf | IssueCSRFToken mints the anti-forgery token a browser echoes as X-CSRF-Token on every change it asks for.
[**GetAccountEmbed**](AccountAPI.md#GetAccountEmbed) | **Get** /v1/account/embed | Reports whether one of this brand&#39;s shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.
[**GetAccountKeys**](AccountAPI.md#GetAccountKeys) | **Get** /v1/account/keys | Returns the caller&#39;s own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.
[**PostAccountAppearance**](AccountAPI.md#PostAccountAppearance) | **Post** /v1/account/appearance | Stores the caller&#39;s appearance preference on their IAM account, preserving every other field of the row.
[**PostAccountAvatar**](AccountAPI.md#PostAccountAvatar) | **Post** /v1/account/avatar | Set your profile photo
[**PostAccountKeys**](AccountAPI.md#PostAccountKeys) | **Post** /v1/account/keys | Creates — or rotates — the caller&#39;s API key of the requested type and returns it ONCE.
[**PostAccountOrgs**](AccountAPI.md#PostAccountOrgs) | **Post** /v1/account/orgs | Onboard creates the caller&#39;s organization.



## DeleteAccountKeys

> RevokedKey DeleteAccountKeys(ctx).Type_(type_).Execute()

Revokes the caller's own API key of the requested class.



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
	type_ := "publishable" // string | Type is the key class to act on: \"secret\" (sk-, session-equivalent, belongs on a server) or \"publishable\" (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.DeleteAccountKeys(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.DeleteAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccountKeys`: RevokedKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.DeleteAccountKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type is the key class to act on: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. | 

### Return type

[**RevokedKey**](RevokedKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAppearance

> Appearance GetAccountAppearance(ctx).Execute()

Returns the signed-in caller's own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.



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
	resp, r, err := apiClient.AccountAPI.GetAccountAppearance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAppearance`: Appearance
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountAppearance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAppearanceRequest struct via the builder pattern


### Return type

[**Appearance**](Appearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk"
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

> CsrfResp GetAccountCsrf(ctx).Execute()

IssueCSRFToken mints the anti-forgery token a browser echoes as X-CSRF-Token on every change it asks for.



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
	resp, r, err := apiClient.AccountAPI.GetAccountCsrf(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountCsrf``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCsrf`: CsrfResp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountCsrf`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCsrfRequest struct via the builder pattern


### Return type

[**CsrfResp**](CsrfResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountEmbed

> EmbedStatusResp GetAccountEmbed(ctx).App(app).Execute()

Reports whether one of this brand's shared embedded apps (cms, erp, help) may be framed by the caller and is actually running, so a console module can choose between the embed and the provision panel.



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
	app := "cms" // string | App is the embedded app to report on: cms (Content Studio), erp or help. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetAccountEmbed(context.Background()).App(app).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountEmbed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEmbed`: EmbedStatusResp
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

[**EmbedStatusResp**](EmbedStatusResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountKeys

> ApiKeyList GetAccountKeys(ctx).Execute()

Returns the caller's own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.



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
	resp, r, err := apiClient.AccountAPI.GetAccountKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountKeys`: ApiKeyList
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetAccountKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountKeysRequest struct via the builder pattern


### Return type

[**ApiKeyList**](ApiKeyList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountAppearance

> Appearance PostAccountAppearance(ctx).Appearance(appearance).Execute()

Stores the caller's appearance preference on their IAM account, preserving every other field of the row.



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
	appearance := *openapiclient.NewAppearance() // Appearance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountAppearance(context.Background()).Appearance(appearance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountAppearance`: Appearance
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountAppearance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountAppearanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appearance** | [**Appearance**](Appearance.md) |  | 

### Return type

[**Appearance**](Appearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	openapiclient "github.com/hanzoai/go-sdk"
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

> MintedKey PostAccountKeys(ctx).KeyTypeIn(keyTypeIn).Execute()

Creates — or rotates — the caller's API key of the requested type and returns it ONCE.



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
	keyTypeIn := *openapiclient.NewKeyTypeIn() // KeyTypeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountKeys(context.Background()).KeyTypeIn(keyTypeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountKeys`: MintedKey
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyTypeIn** | [**KeyTypeIn**](KeyTypeIn.md) |  | 

### Return type

[**MintedKey**](MintedKey.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccountOrgs

> OnboardResp PostAccountOrgs(ctx).OnboardReq(onboardReq).Execute()

Onboard creates the caller's organization.



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
	onboardReq := *openapiclient.NewOnboardReq() // OnboardReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PostAccountOrgs(context.Background()).OnboardReq(onboardReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PostAccountOrgs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccountOrgs`: OnboardResp
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PostAccountOrgs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccountOrgsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onboardReq** | [**OnboardReq**](OnboardReq.md) |  | 

### Return type

[**OnboardResp**](OnboardResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

