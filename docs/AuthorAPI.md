# \AuthorAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthor**](AuthorAPI.md#GetAuthor) | **Get** /v1/author | Returns the caller&#39;s author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.
[**GetAuthorBasis**](AuthorAPI.md#GetAuthorBasis) | **Get** /v1/author/basis | Returns the AUDIT TRAIL behind the caller&#39;s own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform&#39;s matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.
[**PostAuthorConnect**](AuthorAPI.md#PostAuthorConnect) | **Post** /v1/author/connect | Enrols the caller&#39;s org in the author program at status \&quot;connected\&quot; and returns its enrolment, including the verify code the file method needs.
[**PostAuthorDeploysRecord**](AuthorAPI.md#PostAuthorDeploysRecord) | **Post** /v1/author/deploys/record | Records that the caller&#39;s org deployed a project built from a source repository, which is the edge that makes an author&#39;s work earn royalty.
[**PostAuthorReposVerify**](AuthorAPI.md#PostAuthorReposVerify) | **Post** /v1/author/repos/verify | Proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



## GetAuthor

> map[string]interface{} GetAuthor(ctx).Execute()

Returns the caller's author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.



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
	resp, r, err := apiClient.AuthorAPI.GetAuthor(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorAPI.GetAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthor`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorAPI.GetAuthor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorBasis

> map[string]interface{} GetAuthorBasis(ctx).Period(period).Execute()

Returns the AUDIT TRAIL behind the caller's own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform's matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.



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
	period := "2026-07" // string | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400, because the period is echoed back and used as a SQL filter and is only ever accepted in the one form the accrual latch mints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorAPI.GetAuthorBasis(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorAPI.GetAuthorBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorBasis`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorAPI.GetAuthorBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400, because the period is echoed back and used as a SQL filter and is only ever accepted in the one form the accrual latch mints. | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorConnect

> AuthorEnrolment PostAuthorConnect(ctx).AuthorConnectRequest(authorConnectRequest).Execute()

Enrols the caller's org in the author program at status \"connected\" and returns its enrolment, including the verify code the file method needs.



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
	authorConnectRequest := *openapiclient.NewAuthorConnectRequest() // AuthorConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorAPI.PostAuthorConnect(context.Background()).AuthorConnectRequest(authorConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorAPI.PostAuthorConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorConnect`: AuthorEnrolment
	fmt.Fprintf(os.Stdout, "Response from `AuthorAPI.PostAuthorConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorConnectRequest** | [**AuthorConnectRequest**](AuthorConnectRequest.md) |  | 

### Return type

[**AuthorEnrolment**](AuthorEnrolment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorDeploysRecord

> AuthorDeployRecord PostAuthorDeploysRecord(ctx).AuthorDeployRequest(authorDeployRequest).Execute()

Records that the caller's org deployed a project built from a source repository, which is the edge that makes an author's work earn royalty.



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
	authorDeployRequest := *openapiclient.NewAuthorDeployRequest() // AuthorDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorAPI.PostAuthorDeploysRecord(context.Background()).AuthorDeployRequest(authorDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorAPI.PostAuthorDeploysRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorDeploysRecord`: AuthorDeployRecord
	fmt.Fprintf(os.Stdout, "Response from `AuthorAPI.PostAuthorDeploysRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorDeploysRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorDeployRequest** | [**AuthorDeployRequest**](AuthorDeployRequest.md) |  | 

### Return type

[**AuthorDeployRecord**](AuthorDeployRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorReposVerify

> AuthorClaim PostAuthorReposVerify(ctx).AuthorVerifyRequest(authorVerifyRequest).Execute()

Proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



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
	authorVerifyRequest := *openapiclient.NewAuthorVerifyRequest() // AuthorVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorAPI.PostAuthorReposVerify(context.Background()).AuthorVerifyRequest(authorVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorAPI.PostAuthorReposVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorReposVerify`: AuthorClaim
	fmt.Fprintf(os.Stdout, "Response from `AuthorAPI.PostAuthorReposVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorReposVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorVerifyRequest** | [**AuthorVerifyRequest**](AuthorVerifyRequest.md) |  | 

### Return type

[**AuthorClaim**](AuthorClaim.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

