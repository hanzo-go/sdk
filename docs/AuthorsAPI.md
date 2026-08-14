# \AuthorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthors**](AuthorsAPI.md#GetAuthors) | **Get** /v1/authors | Returns the caller&#39;s author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.
[**GetAuthorsBasis**](AuthorsAPI.md#GetAuthorsBasis) | **Get** /v1/authors/basis | Returns the AUDIT TRAIL behind the caller&#39;s own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform&#39;s matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.
[**PostAuthorsConnect**](AuthorsAPI.md#PostAuthorsConnect) | **Post** /v1/authors/connect | Enrols the caller&#39;s org in the author program at status \&quot;connected\&quot; and returns its enrolment, including the verify code the file method needs.
[**PostAuthorsDeploysRecord**](AuthorsAPI.md#PostAuthorsDeploysRecord) | **Post** /v1/authors/deploys/record | Records that the caller&#39;s org deployed a project built from a source repository, which is the edge that makes an author&#39;s work earn royalty.
[**PostAuthorsReposVerify**](AuthorsAPI.md#PostAuthorsReposVerify) | **Post** /v1/authors/repos/verify | Proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



## GetAuthors

> map[string]map[string]interface{} GetAuthors(ctx).Execute()

Returns the caller's author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.



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
	resp, r, err := apiClient.AuthorsAPI.GetAuthors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.GetAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthors`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.GetAuthors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorsRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthorsBasis

> map[string]map[string]interface{} GetAuthorsBasis(ctx).Period(period).Execute()

Returns the AUDIT TRAIL behind the caller's own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform's matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.



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
	period := "2026-07" // string | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400, because the period is echoed back and used as a SQL filter and is only ever accepted in the one form the accrual latch mints. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.GetAuthorsBasis(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.GetAuthorsBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthorsBasis`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.GetAuthorsBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorsBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400, because the period is echoed back and used as a SQL filter and is only ever accepted in the one form the accrual latch mints. | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorsConnect

> Enrolment PostAuthorsConnect(ctx).ConnectRequest(connectRequest).Execute()

Enrols the caller's org in the author program at status \"connected\" and returns its enrolment, including the verify code the file method needs.



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
	connectRequest := *openapiclient.NewConnectRequest() // ConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.PostAuthorsConnect(context.Background()).ConnectRequest(connectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.PostAuthorsConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorsConnect`: Enrolment
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.PostAuthorsConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorsConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectRequest** | [**ConnectRequest**](ConnectRequest.md) |  | 

### Return type

[**Enrolment**](Enrolment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorsDeploysRecord

> DeployRecord PostAuthorsDeploysRecord(ctx).DeployRequest(deployRequest).Execute()

Records that the caller's org deployed a project built from a source repository, which is the edge that makes an author's work earn royalty.



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
	deployRequest := *openapiclient.NewDeployRequest() // DeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.PostAuthorsDeploysRecord(context.Background()).DeployRequest(deployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.PostAuthorsDeploysRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorsDeploysRecord`: DeployRecord
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.PostAuthorsDeploysRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorsDeploysRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployRequest** | [**DeployRequest**](DeployRequest.md) |  | 

### Return type

[**DeployRecord**](DeployRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthorsReposVerify

> Claim PostAuthorsReposVerify(ctx).VerifyRequest(verifyRequest).Execute()

Proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



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
	verifyRequest := *openapiclient.NewVerifyRequest() // VerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.PostAuthorsReposVerify(context.Background()).VerifyRequest(verifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.PostAuthorsReposVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthorsReposVerify`: Claim
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.PostAuthorsReposVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthorsReposVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyRequest** | [**VerifyRequest**](VerifyRequest.md) |  | 

### Return type

[**Claim**](Claim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

