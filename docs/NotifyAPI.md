# \NotifyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNotifyCredentialsByProviderByKey**](NotifyAPI.md#DeleteNotifyCredentialsByProviderByKey) | **Delete** /v1/notify/credentials/{provider}/{key} | Removes one of your org&#39;s notify provider credentials.
[**GetNotifyHealth**](NotifyAPI.md#GetNotifyHealth) | **Get** /v1/notify/health | Reports that the notify send surface is mounted.
[**PostNotifySend**](NotifyAPI.md#PostNotifySend) | **Post** /v1/notify/send | Delivers one transactional message by email or SMS through the caller org&#39;s own provider credential.
[**PostNotifySendEmail**](NotifyAPI.md#PostNotifySendEmail) | **Post** /v1/notify/send/email | Delivers one transactional email through the caller org&#39;s own provider credential.
[**PostNotifySendSms**](NotifyAPI.md#PostNotifySendSms) | **Post** /v1/notify/send/sms | Delivers one transactional SMS through the caller org&#39;s own provider credential.
[**PutNotifyCredentialsByProviderByKey**](NotifyAPI.md#PutNotifyCredentialsByProviderByKey) | **Put** /v1/notify/credentials/{provider}/{key} | Sets one of your org&#39;s notify provider credentials.



## DeleteNotifyCredentialsByProviderByKey

> DeleteNotifyCredentialsByProviderByKey(ctx, provider, key).Execute()

Removes one of your org's notify provider credentials.



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
	provider := "provider_example" // string | Provider is the delivery provider the credential is for.
	key := "key_example" // string | Key is the credential's name within that provider.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotifyAPI.DeleteNotifyCredentialsByProviderByKey(context.Background(), provider, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.DeleteNotifyCredentialsByProviderByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the delivery provider the credential is for. | 
**key** | **string** | Key is the credential&#39;s name within that provider. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotifyCredentialsByProviderByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifyHealth

> NotifyNotifyHealth GetNotifyHealth(ctx).Execute()

Reports that the notify send surface is mounted.



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
	resp, r, err := apiClient.NotifyAPI.GetNotifyHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.GetNotifyHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifyHealth`: NotifyNotifyHealth
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.GetNotifyHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotifyHealthRequest struct via the builder pattern


### Return type

[**NotifyNotifyHealth**](NotifyNotifyHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySend

> interface{} PostNotifySend(ctx).NotifyNotifySend(notifyNotifySend).Execute()

Delivers one transactional message by email or SMS through the caller org's own provider credential.



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
	notifyNotifySend := *openapiclient.NewNotifyNotifySend() // NotifyNotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySend(context.Background()).NotifyNotifySend(notifyNotifySend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.PostNotifySend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNotifySend`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.PostNotifySend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotifySendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyNotifySend** | [**NotifyNotifySend**](NotifyNotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySendEmail

> interface{} PostNotifySendEmail(ctx).NotifyNotifySend(notifyNotifySend).Execute()

Delivers one transactional email through the caller org's own provider credential.



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
	notifyNotifySend := *openapiclient.NewNotifyNotifySend() // NotifyNotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySendEmail(context.Background()).NotifyNotifySend(notifyNotifySend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.PostNotifySendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNotifySendEmail`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.PostNotifySendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotifySendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyNotifySend** | [**NotifyNotifySend**](NotifyNotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySendSms

> interface{} PostNotifySendSms(ctx).NotifyNotifySend(notifyNotifySend).Execute()

Delivers one transactional SMS through the caller org's own provider credential.



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
	notifyNotifySend := *openapiclient.NewNotifyNotifySend() // NotifyNotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySendSms(context.Background()).NotifyNotifySend(notifyNotifySend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.PostNotifySendSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNotifySendSms`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.PostNotifySendSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotifySendSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notifyNotifySend** | [**NotifyNotifySend**](NotifyNotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNotifyCredentialsByProviderByKey

> NotifyNotifyStored PutNotifyCredentialsByProviderByKey(ctx, provider, key).NotifyNotifyCredential(notifyNotifyCredential).Execute()

Sets one of your org's notify provider credentials.



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
	provider := "twilio" // string | Provider is the delivery provider the credential is for: twilio, plivo, twilio_email or mail.
	key := "auth-token" // string | Key is one of the credentials that provider reads, such as auth-token or smtp-host. A key the provider does not read is refused.
	notifyNotifyCredential := *openapiclient.NewNotifyNotifyCredential() // NotifyNotifyCredential | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PutNotifyCredentialsByProviderByKey(context.Background(), provider, key).NotifyNotifyCredential(notifyNotifyCredential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.PutNotifyCredentialsByProviderByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNotifyCredentialsByProviderByKey`: NotifyNotifyStored
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.PutNotifyCredentialsByProviderByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the delivery provider the credential is for: twilio, plivo, twilio_email or mail. | 
**key** | **string** | Key is one of the credentials that provider reads, such as auth-token or smtp-host. A key the provider does not read is refused. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNotifyCredentialsByProviderByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notifyNotifyCredential** | [**NotifyNotifyCredential**](NotifyNotifyCredential.md) |  | 

### Return type

[**NotifyNotifyStored**](NotifyNotifyStored.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

