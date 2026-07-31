# \MFAAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerDeleteMfa**](MFAAPI.md#IamApiControllerDeleteMfa) | **Delete** /v1/iam/mfa/s/{id} | Api Controller Delete Mfa
[**IamApiControllerGetUserVerifications**](MFAAPI.md#IamApiControllerGetUserVerifications) | **Get** /v1/iam/user-payments | Api Controller Get User Verifications
[**IamApiControllerGetVerification**](MFAAPI.md#IamApiControllerGetVerification) | **Get** /v1/iam/payments/{id} | Api Controller Get Verification
[**IamApiControllerGetVerifications**](MFAAPI.md#IamApiControllerGetVerifications) | **Get** /v1/iam/payments | Api Controller Get Verifications
[**IamApiControllerMfaSetupEnable**](MFAAPI.md#IamApiControllerMfaSetupEnable) | **Post** /v1/iam/mfa/setup/enable | Api Controller Mfa Setup Enable
[**IamApiControllerMfaSetupInitiate**](MFAAPI.md#IamApiControllerMfaSetupInitiate) | **Post** /v1/iam/mfa/setup/initiate | Api Controller Mfa Setup Initiate
[**IamApiControllerMfaSetupVerify**](MFAAPI.md#IamApiControllerMfaSetupVerify) | **Post** /v1/iam/mfa/setup/verify | Api Controller Mfa Setup Verify
[**IamApiControllerSendVerificationCode**](MFAAPI.md#IamApiControllerSendVerificationCode) | **Post** /v1/iam/auth/verification-code/send | Api Controller Send Verification Code
[**IamApiControllerSetPreferredMfa**](MFAAPI.md#IamApiControllerSetPreferredMfa) | **Post** /v1/iam/mfa/preferred | Api Controller Set Preferred Mfa
[**IamApiControllerVerifyCaptcha**](MFAAPI.md#IamApiControllerVerifyCaptcha) | **Post** /v1/iam/captcha/verify | Api Controller Verify Captcha
[**IamApiControllerVerifyCode**](MFAAPI.md#IamApiControllerVerifyCode) | **Post** /v1/iam/auth/verification-code/verify | Api Controller Verify Code



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
	resp, r, err := apiClient.MFAAPI.IamApiControllerDeleteMfa(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerDeleteMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteMfa`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerDeleteMfa`: %v\n", resp)
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
	resp, r, err := apiClient.MFAAPI.IamApiControllerGetUserVerifications(context.Background()).Owner(owner).Organization(organization).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerGetUserVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetUserVerifications`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerGetUserVerifications`: %v\n", resp)
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
	resp, r, err := apiClient.MFAAPI.IamApiControllerGetVerification(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerGetVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVerification`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerGetVerification`: %v\n", resp)
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
	resp, r, err := apiClient.MFAAPI.IamApiControllerGetVerifications(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerGetVerifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetVerifications`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerGetVerifications`: %v\n", resp)
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

> IamControllersResponse IamApiControllerMfaSetupEnable(ctx).Owner(owner).Name(name).MfaType(mfaType).RecoveryCodes(recoveryCodes).Secret(secret).Dest(dest).Execute()

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
	owner := "owner_example" // string | 
	name := "name_example" // string | 
	mfaType := "mfaType_example" // string | totp | sms | email | radius | push
	recoveryCodes := "recoveryCodes_example" // string | 
	secret := "secret_example" // string | Required for totp; radius/push also require it (optional)
	dest := "dest_example" // string | Destination for sms/email/radius/push (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerMfaSetupEnable(context.Background()).Owner(owner).Name(name).MfaType(mfaType).RecoveryCodes(recoveryCodes).Secret(secret).Dest(dest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerMfaSetupEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupEnable`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerMfaSetupEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** |  | 
 **name** | **string** |  | 
 **mfaType** | **string** | totp | sms | email | radius | push | 
 **recoveryCodes** | **string** |  | 
 **secret** | **string** | Required for totp; radius/push also require it | 
 **dest** | **string** | Destination for sms/email/radius/push | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerMfaSetupInitiate

> IamControllersResponse IamApiControllerMfaSetupInitiate(ctx).MfaType(mfaType).Owner(owner).Name(name).Execute()

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
	mfaType := "mfaType_example" // string | 
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerMfaSetupInitiate(context.Background()).MfaType(mfaType).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerMfaSetupInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupInitiate`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerMfaSetupInitiate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaType** | **string** |  | 
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerMfaSetupVerify

> IamControllersResponse IamApiControllerMfaSetupVerify(ctx).MfaType(mfaType).Passcode(passcode).Secret(secret).Dest(dest).CountryCode(countryCode).Execute()

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
	mfaType := "mfaType_example" // string | 
	passcode := "passcode_example" // string | 
	secret := "secret_example" // string | Required for totp and radius/push (optional)
	dest := "dest_example" // string | Required for sms/email/radius/push (optional)
	countryCode := "countryCode_example" // string | Required for sms (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerMfaSetupVerify(context.Background()).MfaType(mfaType).Passcode(passcode).Secret(secret).Dest(dest).CountryCode(countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerMfaSetupVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerMfaSetupVerify`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerMfaSetupVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerMfaSetupVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaType** | **string** |  | 
 **passcode** | **string** |  | 
 **secret** | **string** | Required for totp and radius/push | 
 **dest** | **string** | Required for sms/email/radius/push | 
 **countryCode** | **string** | Required for sms | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSendVerificationCode

> IamObjectUserinfo IamApiControllerSendVerificationCode(ctx).Dest(dest).Type_(type_).ApplicationId(applicationId).CaptchaType(captchaType).CountryCode(countryCode).Method(method).CheckUser(checkUser).ClientSecret(clientSecret).CaptchaToken(captchaToken).Execute()

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
	dest := "dest_example" // string | 
	type_ := "type__example" // string | 
	applicationId := "applicationId_example" // string | Must be of the form <owner>/<name>
	captchaType := "captchaType_example" // string | 
	countryCode := "countryCode_example" // string |  (optional)
	method := "method_example" // string |  (optional)
	checkUser := "checkUser_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string | Required when captchaType != none (optional)
	captchaToken := "captchaToken_example" // string | Required when captchaType != none (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerSendVerificationCode(context.Background()).Dest(dest).Type_(type_).ApplicationId(applicationId).CaptchaType(captchaType).CountryCode(countryCode).Method(method).CheckUser(checkUser).ClientSecret(clientSecret).CaptchaToken(captchaToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerSendVerificationCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSendVerificationCode`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerSendVerificationCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSendVerificationCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dest** | **string** |  | 
 **type_** | **string** |  | 
 **applicationId** | **string** | Must be of the form &lt;owner&gt;/&lt;name&gt; | 
 **captchaType** | **string** |  | 
 **countryCode** | **string** |  | 
 **method** | **string** |  | 
 **checkUser** | **string** |  | 
 **clientSecret** | **string** | Required when captchaType !&#x3D; none | 
 **captchaToken** | **string** | Required when captchaType !&#x3D; none | 

### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerSetPreferredMfa

> IamControllersResponse IamApiControllerSetPreferredMfa(ctx).MfaType(mfaType).Owner(owner).Name(name).Execute()

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
	mfaType := "mfaType_example" // string | 
	owner := "owner_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerSetPreferredMfa(context.Background()).MfaType(mfaType).Owner(owner).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerSetPreferredMfa``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerSetPreferredMfa`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerSetPreferredMfa`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerSetPreferredMfaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaType** | **string** |  | 
 **owner** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyCaptcha

> IamObjectUserinfo IamApiControllerVerifyCaptcha(ctx).ApplicationId(applicationId).CaptchaType(captchaType).CaptchaToken(captchaToken).ClientSecret(clientSecret).Execute()

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
	applicationId := "applicationId_example" // string | 
	captchaType := "captchaType_example" // string | 
	captchaToken := "captchaToken_example" // string | 
	clientSecret := "clientSecret_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerVerifyCaptcha(context.Background()).ApplicationId(applicationId).CaptchaType(captchaType).CaptchaToken(captchaToken).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerVerifyCaptcha``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyCaptcha`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerVerifyCaptcha`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyCaptchaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **string** |  | 
 **captchaType** | **string** |  | 
 **captchaToken** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyCode

> IamObjectUserinfo IamApiControllerVerifyCode(ctx).IamControllersVerifyCodeForm(iamControllersVerifyCodeForm).Execute()

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
	iamControllersVerifyCodeForm := *openapiclient.NewIamControllersVerifyCodeForm("Username_example", "Code_example") // IamControllersVerifyCodeForm | Handler binds form.AuthForm; this endpoint consumes organization, username, name, code, countryCode. Extra AuthForm fields are accepted but ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MFAAPI.IamApiControllerVerifyCode(context.Background()).IamControllersVerifyCodeForm(iamControllersVerifyCodeForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MFAAPI.IamApiControllerVerifyCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyCode`: IamObjectUserinfo
	fmt.Fprintf(os.Stdout, "Response from `MFAAPI.IamApiControllerVerifyCode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamControllersVerifyCodeForm** | [**IamControllersVerifyCodeForm**](IamControllersVerifyCodeForm.md) | Handler binds form.AuthForm; this endpoint consumes organization, username, name, code, countryCode. Extra AuthForm fields are accepted but ignored. | 

### Return type

[**IamObjectUserinfo**](IamObjectUserinfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

