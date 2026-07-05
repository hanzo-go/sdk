# \IamMFAAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerDeleteMfa**](IamMFAAPI.md#IamApiControllerDeleteMfa) | **Delete** /v1/iam/mfa/s/{id} | Api Controller Delete Mfa
[**IamApiControllerGetUserVerifications**](IamMFAAPI.md#IamApiControllerGetUserVerifications) | **Get** /v1/iam/user-payments | Api Controller Get User Verifications
[**IamApiControllerGetVerification**](IamMFAAPI.md#IamApiControllerGetVerification) | **Get** /v1/iam/payments/{id} | Api Controller Get Verification
[**IamApiControllerGetVerifications**](IamMFAAPI.md#IamApiControllerGetVerifications) | **Get** /v1/iam/payments | Api Controller Get Verifications
[**IamApiControllerMfaSetupEnable**](IamMFAAPI.md#IamApiControllerMfaSetupEnable) | **Post** /v1/iam/mfa/setup/enable | Api Controller Mfa Setup Enable
[**IamApiControllerMfaSetupInitiate**](IamMFAAPI.md#IamApiControllerMfaSetupInitiate) | **Post** /v1/iam/mfa/setup/initiate | Api Controller Mfa Setup Initiate
[**IamApiControllerMfaSetupVerify**](IamMFAAPI.md#IamApiControllerMfaSetupVerify) | **Post** /v1/iam/mfa/setup/verify | Api Controller Mfa Setup Verify
[**IamApiControllerSendVerificationCode**](IamMFAAPI.md#IamApiControllerSendVerificationCode) | **Post** /v1/iam/auth/verification-code/send | Api Controller Send Verification Code
[**IamApiControllerSetPreferredMfa**](IamMFAAPI.md#IamApiControllerSetPreferredMfa) | **Post** /v1/iam/mfa/preferred | Api Controller Set Preferred Mfa
[**IamApiControllerVerifyCaptcha**](IamMFAAPI.md#IamApiControllerVerifyCaptcha) | **Post** /v1/iam/captcha/verify | Api Controller Verify Captcha
[**IamApiControllerVerifyCode**](IamMFAAPI.md#IamApiControllerVerifyCode) | **Post** /v1/iam/auth/verification-code/verify | Api Controller Verify Code



## IamApiControllerDeleteMfa

> IamControllersResponse IamApiControllerDeleteMfa(ctx, id).Execute()

Api Controller Delete Mfa



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
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerDeleteMfa(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerDeleteMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteMfa`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerDeleteMfa`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetUserVerifications

> []map[string]interface{} IamApiControllerGetUserVerifications(ctx).Owner(owner).Organization(organization).User(user).Execute()

Api Controller Get User Verifications



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
	owner := "owner_example" // string | The owner of payments
	organization := "organization_example" // string | The organization of the user
	user := "user_example" // string | The username of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerGetUserVerifications(context.Background()).Owner(owner).Organization(organization).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerGetUserVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserVerifications`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerGetUserVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetUserVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of payments | 
 **organization** | **string** | The organization of the user | 
 **user** | **string** | The username of the user | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetVerification

> map[string]interface{} IamApiControllerGetVerification(ctx, id).Execute()

Api Controller Get Verification



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
	id := "id_example" // string | The id ( owner/name ) of the payment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerGetVerification(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerGetVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVerification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerGetVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the payment | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetVerifications

> []map[string]interface{} IamApiControllerGetVerifications(ctx).Owner(owner).Execute()

Api Controller Get Verifications



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
	owner := "owner_example" // string | The owner of payments

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerGetVerifications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerGetVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVerifications`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerGetVerifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetVerificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of payments | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerMfaSetupEnable

> IamControllersResponse IamApiControllerMfaSetupEnable(ctx).Execute()

Api Controller Mfa Setup Enable



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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerMfaSetupEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerMfaSetupEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupEnable`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerMfaSetupEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupEnableRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerMfaSetupInitiate

> IamControllersResponse IamApiControllerMfaSetupInitiate(ctx).Execute()

Api Controller Mfa Setup Initiate



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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerMfaSetupInitiate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerMfaSetupInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupInitiate`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerMfaSetupInitiate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupInitiateRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerMfaSetupVerify

> IamControllersResponse IamApiControllerMfaSetupVerify(ctx).Execute()

Api Controller Mfa Setup Verify



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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerMfaSetupVerify(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerMfaSetupVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupVerify`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerMfaSetupVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupVerifyRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSendVerificationCode

> IamObjectUserinfo IamApiControllerSendVerificationCode(ctx).Execute()

Api Controller Send Verification Code

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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerSendVerificationCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerSendVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendVerificationCode`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerSendVerificationCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSendVerificationCodeRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSetPreferredMfa

> IamControllersResponse IamApiControllerSetPreferredMfa(ctx).Execute()

Api Controller Set Preferred Mfa



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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerSetPreferredMfa(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerSetPreferredMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSetPreferredMfa`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerSetPreferredMfa`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSetPreferredMfaRequest struct via the builder pattern


### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyCaptcha

> IamObjectUserinfo IamApiControllerVerifyCaptcha(ctx).Execute()

Api Controller Verify Captcha

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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerVerifyCaptcha(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerVerifyCaptcha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyCaptcha`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerVerifyCaptcha`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyCaptchaRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyCode

> IamObjectUserinfo IamApiControllerVerifyCode(ctx).Execute()

Api Controller Verify Code

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
	resp, r, err := apiClient.IamMFAAPI.IamApiControllerVerifyCode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamMFAAPI.IamApiControllerVerifyCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyCode`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `IamMFAAPI.IamApiControllerVerifyCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyCodeRequest struct via the builder pattern


### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

