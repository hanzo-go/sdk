# \LicensingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicensingDownloadByRelease**](LicensingAPI.md#GetLicensingDownloadByRelease) | **Get** /v1/licensing/download/{release} | Download resolves a release to its artifact, gated on a valid license.
[**GetLicensingHealthz**](LicensingAPI.md#GetLicensingHealthz) | **Get** /v1/licensing/healthz | Health reports which signer this deployment mints with, and in which env.
[**GetLicensingJwks**](LicensingAPI.md#GetLicensingJwks) | **Get** /v1/licensing/jwks | Pubkey publishes the Ed25519 PUBLIC verification key, at both /pubkey and /jwks.
[**GetLicensingPubkey**](LicensingAPI.md#GetLicensingPubkey) | **Get** /v1/licensing/pubkey | Pubkey publishes the Ed25519 PUBLIC verification key, at both /pubkey and /jwks.
[**GetLicensingReleases**](LicensingAPI.md#GetLicensingReleases) | **Get** /v1/licensing/releases | Lists the signed binary releases this deployment can serve.
[**GetLicensingReleasesByRelease**](LicensingAPI.md#GetLicensingReleasesByRelease) | **Get** /v1/licensing/releases/{release} | Reads one release&#39;s metadata: its product, version, platform and the cosign material a client verifies the binary against.
[**PostLicensingFingerprint**](LicensingAPI.md#PostLicensingFingerprint) | **Post** /v1/licensing/fingerprint | Fingerprint turns raw device signals into the opaque value that binds a license to one machine.
[**PostLicensingIssue**](LicensingAPI.md#PostLicensingIssue) | **Post** /v1/licensing/issue | Issue mints a signed license token for a product the caller&#39;s org already pays for.
[**PostLicensingReleases**](LicensingAPI.md#PostLicensingReleases) | **Post** /v1/licensing/releases | Publishes a signed binary release, answering 201 Created.
[**PostLicensingRevoke**](LicensingAPI.md#PostLicensingRevoke) | **Post** /v1/licensing/revoke | Revoke turns off tokens that have already been issued.
[**PostLicensingVerify**](LicensingAPI.md#PostLicensingVerify) | **Post** /v1/licensing/verify | Verify checks a license token online: signature, schema, expiry, app_id and the revocation list.



## GetLicensingDownloadByRelease

> LicensingReleaseAsset GetLicensingDownloadByRelease(ctx, release).XLicenseToken(xLicenseToken).Token(token).Execute()

Download resolves a release to its artifact, gated on a valid license.



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
	release := "release_example" // string | 
	xLicenseToken := "xLicenseToken_example" // string |  (optional)
	token := "token_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.GetLicensingDownloadByRelease(context.Background(), release).XLicenseToken(xLicenseToken).Token(token).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingDownloadByRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingDownloadByRelease`: LicensingReleaseAsset
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingDownloadByRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**release** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingDownloadByReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xLicenseToken** | **string** |  | 
 **token** | **string** |  | 

### Return type

[**LicensingReleaseAsset**](LicensingReleaseAsset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensingHealthz

> LicensingHealthView GetLicensingHealthz(ctx).Execute()

Health reports which signer this deployment mints with, and in which env.



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
	resp, r, err := apiClient.LicensingAPI.GetLicensingHealthz(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingHealthz``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingHealthz`: LicensingHealthView
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingHealthz`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingHealthzRequest struct via the builder pattern


### Return type

[**LicensingHealthView**](LicensingHealthView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensingJwks

> LicensingPubkeyView GetLicensingJwks(ctx).Execute()

Pubkey publishes the Ed25519 PUBLIC verification key, at both /pubkey and /jwks.



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
	resp, r, err := apiClient.LicensingAPI.GetLicensingJwks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingJwks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingJwks`: LicensingPubkeyView
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingJwks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingJwksRequest struct via the builder pattern


### Return type

[**LicensingPubkeyView**](LicensingPubkeyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensingPubkey

> LicensingPubkeyView GetLicensingPubkey(ctx).Execute()

Pubkey publishes the Ed25519 PUBLIC verification key, at both /pubkey and /jwks.



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
	resp, r, err := apiClient.LicensingAPI.GetLicensingPubkey(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingPubkey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingPubkey`: LicensingPubkeyView
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingPubkey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingPubkeyRequest struct via the builder pattern


### Return type

[**LicensingPubkeyView**](LicensingPubkeyView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensingReleases

> LicensingReleaseList GetLicensingReleases(ctx).Execute()

Lists the signed binary releases this deployment can serve.



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
	resp, r, err := apiClient.LicensingAPI.GetLicensingReleases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingReleases`: LicensingReleaseList
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingReleases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingReleasesRequest struct via the builder pattern


### Return type

[**LicensingReleaseList**](LicensingReleaseList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicensingReleasesByRelease

> LicensingRelease GetLicensingReleasesByRelease(ctx, release).Execute()

Reads one release's metadata: its product, version, platform and the cosign material a client verifies the binary against.



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
	release := "release_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.GetLicensingReleasesByRelease(context.Background(), release).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.GetLicensingReleasesByRelease``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensingReleasesByRelease`: LicensingRelease
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.GetLicensingReleasesByRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**release** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensingReleasesByReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicensingRelease**](LicensingRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLicensingFingerprint

> LicensingFingerprintResponse PostLicensingFingerprint(ctx).LicensingFingerprintRequest(licensingFingerprintRequest).Execute()

Fingerprint turns raw device signals into the opaque value that binds a license to one machine.



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
	licensingFingerprintRequest := *openapiclient.NewLicensingFingerprintRequest() // LicensingFingerprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.PostLicensingFingerprint(context.Background()).LicensingFingerprintRequest(licensingFingerprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.PostLicensingFingerprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLicensingFingerprint`: LicensingFingerprintResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.PostLicensingFingerprint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLicensingFingerprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensingFingerprintRequest** | [**LicensingFingerprintRequest**](LicensingFingerprintRequest.md) |  | 

### Return type

[**LicensingFingerprintResponse**](LicensingFingerprintResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLicensingIssue

> LicensingIssueResponse PostLicensingIssue(ctx).LicensingIssueRequest(licensingIssueRequest).Execute()

Issue mints a signed license token for a product the caller's org already pays for.



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
	licensingIssueRequest := *openapiclient.NewLicensingIssueRequest("Product_example") // LicensingIssueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.PostLicensingIssue(context.Background()).LicensingIssueRequest(licensingIssueRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.PostLicensingIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLicensingIssue`: LicensingIssueResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.PostLicensingIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLicensingIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensingIssueRequest** | [**LicensingIssueRequest**](LicensingIssueRequest.md) |  | 

### Return type

[**LicensingIssueResponse**](LicensingIssueResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLicensingReleases

> LicensingRelease PostLicensingReleases(ctx).LicensingRelease(licensingRelease).Execute()

Publishes a signed binary release, answering 201 Created.



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
	licensingRelease := *openapiclient.NewLicensingRelease() // LicensingRelease | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.PostLicensingReleases(context.Background()).LicensingRelease(licensingRelease).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.PostLicensingReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLicensingReleases`: LicensingRelease
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.PostLicensingReleases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLicensingReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensingRelease** | [**LicensingRelease**](LicensingRelease.md) |  | 

### Return type

[**LicensingRelease**](LicensingRelease.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLicensingRevoke

> LicensingRevokeResponse PostLicensingRevoke(ctx).LicensingRevokeRequest(licensingRevokeRequest).Execute()

Revoke turns off tokens that have already been issued.



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
	licensingRevokeRequest := *openapiclient.NewLicensingRevokeRequest("Scope_example", "Value_example") // LicensingRevokeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.PostLicensingRevoke(context.Background()).LicensingRevokeRequest(licensingRevokeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.PostLicensingRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLicensingRevoke`: LicensingRevokeResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.PostLicensingRevoke`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLicensingRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensingRevokeRequest** | [**LicensingRevokeRequest**](LicensingRevokeRequest.md) |  | 

### Return type

[**LicensingRevokeResponse**](LicensingRevokeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLicensingVerify

> LicensingVerifyResponse PostLicensingVerify(ctx).LicensingVerifyRequest(licensingVerifyRequest).Execute()

Verify checks a license token online: signature, schema, expiry, app_id and the revocation list.



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
	licensingVerifyRequest := *openapiclient.NewLicensingVerifyRequest("Token_example") // LicensingVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicensingAPI.PostLicensingVerify(context.Background()).LicensingVerifyRequest(licensingVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicensingAPI.PostLicensingVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLicensingVerify`: LicensingVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `LicensingAPI.PostLicensingVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLicensingVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **licensingVerifyRequest** | [**LicensingVerifyRequest**](LicensingVerifyRequest.md) |  | 

### Return type

[**LicensingVerifyResponse**](LicensingVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

