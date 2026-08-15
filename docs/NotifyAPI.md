# \NotifyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotifyHealth**](NotifyAPI.md#GetNotifyHealth) | **Get** /v1/notify/health | Reports that the notify send surface is mounted.
[**PostNotifySend**](NotifyAPI.md#PostNotifySend) | **Post** /v1/notify/send | Delivers one transactional message by email or SMS through the caller org&#39;s own provider credential.
[**PostNotifySendEmail**](NotifyAPI.md#PostNotifySendEmail) | **Post** /v1/notify/send/email | Delivers one transactional email through the caller org&#39;s own provider credential.
[**PostNotifySendSms**](NotifyAPI.md#PostNotifySendSms) | **Post** /v1/notify/send/sms | Delivers one transactional SMS through the caller org&#39;s own provider credential.



## GetNotifyHealth

> NotifyHealth GetNotifyHealth(ctx).Execute()

Reports that the notify send surface is mounted.



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
	resp, r, err := apiClient.NotifyAPI.GetNotifyHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotifyAPI.GetNotifyHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifyHealth`: NotifyHealth
	fmt.Fprintf(os.Stdout, "Response from `NotifyAPI.GetNotifyHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotifyHealthRequest struct via the builder pattern


### Return type

[**NotifyHealth**](NotifyHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySend

> interface{} PostNotifySend(ctx).NotifySend(notifySend).Execute()

Delivers one transactional message by email or SMS through the caller org's own provider credential.



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
	notifySend := *openapiclient.NewNotifySend() // NotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySend(context.Background()).NotifySend(notifySend).Execute()
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
 **notifySend** | [**NotifySend**](NotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySendEmail

> interface{} PostNotifySendEmail(ctx).NotifySend(notifySend).Execute()

Delivers one transactional email through the caller org's own provider credential.



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
	notifySend := *openapiclient.NewNotifySend() // NotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySendEmail(context.Background()).NotifySend(notifySend).Execute()
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
 **notifySend** | [**NotifySend**](NotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotifySendSms

> interface{} PostNotifySendSms(ctx).NotifySend(notifySend).Execute()

Delivers one transactional SMS through the caller org's own provider credential.



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
	notifySend := *openapiclient.NewNotifySend() // NotifySend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotifyAPI.PostNotifySendSms(context.Background()).NotifySend(notifySend).Execute()
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
 **notifySend** | [**NotifySend**](NotifySend.md) |  | 

### Return type

**interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

