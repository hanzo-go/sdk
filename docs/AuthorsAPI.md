# \AuthorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Authors**](AuthorsAPI.md#CloudGetV1Authors) | **Get** /v1/authors | MyAuthorProgram returns the caller&#39;s author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.
[**CloudGetV1AuthorsBasis**](AuthorsAPI.md#CloudGetV1AuthorsBasis) | **Get** /v1/authors/basis | MyRoyaltyBasis returns the AUDIT TRAIL behind the caller&#39;s own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform&#39;s matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.
[**CloudPostV1AuthorsConnect**](AuthorsAPI.md#CloudPostV1AuthorsConnect) | **Post** /v1/authors/connect | ConnectAuthor enrols the caller&#39;s org in the author program at status \&quot;connected\&quot; and returns its enrolment, including the verify code the file method needs.
[**CloudPostV1AuthorsDeploysRecord**](AuthorsAPI.md#CloudPostV1AuthorsDeploysRecord) | **Post** /v1/authors/deploys/record | RecordAuthorDeploy records that the caller&#39;s org deployed a project built from a source repository, which is the edge that makes an author&#39;s work earn royalty.
[**CloudPostV1AuthorsReposVerify**](AuthorsAPI.md#CloudPostV1AuthorsReposVerify) | **Post** /v1/authors/repos/verify | VerifyAuthorRepo proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



## CloudGetV1Authors

> map[string]map[string]interface{} CloudGetV1Authors(ctx).Execute()

MyAuthorProgram returns the caller's author-program dashboard: enrolment status, linked forge login, verified repositories and owner-wide claims, recorded deploys, accrued / pending / paid royalty, and the payout history.



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
	resp, r, err := apiClient.AuthorsAPI.CloudGetV1Authors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.CloudGetV1Authors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Authors`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.CloudGetV1Authors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AuthorsRequest struct via the builder pattern


### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1AuthorsBasis

> map[string]map[string]interface{} CloudGetV1AuthorsBasis(ctx).Period(period).Execute()

MyRoyaltyBasis returns the AUDIT TRAIL behind the caller's own royalty: every ledger row with the spend it was computed from, the share applied at the time, the platform's matching half, whether each row satisfies the formula, and the attribution edges that already existed when the row was written.



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
	resp, r, err := apiClient.AuthorsAPI.CloudGetV1AuthorsBasis(context.Background()).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.CloudGetV1AuthorsBasis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1AuthorsBasis`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.CloudGetV1AuthorsBasis`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1AuthorsBasisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Period is the UTC accrual month, YYYY-MM. Empty means every period; any other shape is refused with 400, because the period is echoed back and used as a SQL filter and is only ever accepted in the one form the accrual latch mints. | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AuthorsConnect

> CloudEnrolment CloudPostV1AuthorsConnect(ctx).CloudConnectRequest(cloudConnectRequest).Execute()

ConnectAuthor enrols the caller's org in the author program at status \"connected\" and returns its enrolment, including the verify code the file method needs.



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
	cloudConnectRequest := *openapiclient.NewCloudConnectRequest() // CloudConnectRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.CloudPostV1AuthorsConnect(context.Background()).CloudConnectRequest(cloudConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.CloudPostV1AuthorsConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AuthorsConnect`: CloudEnrolment
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.CloudPostV1AuthorsConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AuthorsConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudConnectRequest** | [**CloudConnectRequest**](CloudConnectRequest.md) |  | 

### Return type

[**CloudEnrolment**](CloudEnrolment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AuthorsDeploysRecord

> CloudDeployRecord CloudPostV1AuthorsDeploysRecord(ctx).CloudDeployRequest(cloudDeployRequest).Execute()

RecordAuthorDeploy records that the caller's org deployed a project built from a source repository, which is the edge that makes an author's work earn royalty.



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
	cloudDeployRequest := *openapiclient.NewCloudDeployRequest() // CloudDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.CloudPostV1AuthorsDeploysRecord(context.Background()).CloudDeployRequest(cloudDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.CloudPostV1AuthorsDeploysRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AuthorsDeploysRecord`: CloudDeployRecord
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.CloudPostV1AuthorsDeploysRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AuthorsDeploysRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDeployRequest** | [**CloudDeployRequest**](CloudDeployRequest.md) |  | 

### Return type

[**CloudDeployRecord**](CloudDeployRecord.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1AuthorsReposVerify

> CloudClaim CloudPostV1AuthorsReposVerify(ctx).CloudVerifyRequest(cloudVerifyRequest).Execute()

VerifyAuthorRepo proves that the caller owns a repository — or a whole OWNER — and records the claim, which is what makes deploys of that code earn royalty.



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
	cloudVerifyRequest := *openapiclient.NewCloudVerifyRequest() // CloudVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.CloudPostV1AuthorsReposVerify(context.Background()).CloudVerifyRequest(cloudVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.CloudPostV1AuthorsReposVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1AuthorsReposVerify`: CloudClaim
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.CloudPostV1AuthorsReposVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1AuthorsReposVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudVerifyRequest** | [**CloudVerifyRequest**](CloudVerifyRequest.md) |  | 

### Return type

[**CloudClaim**](CloudClaim.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

