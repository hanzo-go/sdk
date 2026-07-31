# \SendAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotifyNotifySend**](SendAPI.md#NotifyNotifySend) | **Post** /v1/notify/send | Send a notification (generic channel)
[**NotifyNotifySendEmail**](SendAPI.md#NotifyNotifySendEmail) | **Post** /v1/notify/send/email | Send an email notification
[**NotifyNotifySendSms**](SendAPI.md#NotifyNotifySendSms) | **Post** /v1/notify/send/sms | Send an SMS notification



## NotifyNotifySend

> NotifyNotifySend200Response NotifyNotifySend(ctx).Sync(sync).NotifySendRequest(notifySendRequest).Execute()

Send a notification (generic channel)



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
	sync := "sync_example" // string | Must be `true`. Async dispatch is not available in the cloud fold; any other value yields `503`. 
	notifySendRequest := *openapiclient.NewNotifySendRequest([]string{"To_example"}) // NotifySendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.NotifyNotifySend(context.Background()).Sync(sync).NotifySendRequest(notifySendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.NotifyNotifySend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotifyNotifySend`: NotifyNotifySend200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.NotifyNotifySend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyNotifySendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sync** | **string** | Must be &#x60;true&#x60;. Async dispatch is not available in the cloud fold; any other value yields &#x60;503&#x60;.  | 
 **notifySendRequest** | [**NotifySendRequest**](NotifySendRequest.md) |  | 

### Return type

[**NotifyNotifySend200Response**](NotifyNotifySend200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyNotifySendEmail

> NotifyNotifySend200Response NotifyNotifySendEmail(ctx).Sync(sync).NotifySendRequest(notifySendRequest).Execute()

Send an email notification



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
	sync := "sync_example" // string | Must be `true`; otherwise `503`.
	notifySendRequest := *openapiclient.NewNotifySendRequest([]string{"To_example"}) // NotifySendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.NotifyNotifySendEmail(context.Background()).Sync(sync).NotifySendRequest(notifySendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.NotifyNotifySendEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotifyNotifySendEmail`: NotifyNotifySend200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.NotifyNotifySendEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyNotifySendEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sync** | **string** | Must be &#x60;true&#x60;; otherwise &#x60;503&#x60;. | 
 **notifySendRequest** | [**NotifySendRequest**](NotifySendRequest.md) |  | 

### Return type

[**NotifyNotifySend200Response**](NotifyNotifySend200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyNotifySendSms

> NotifyNotifySend200Response NotifyNotifySendSms(ctx).Sync(sync).NotifySendRequest(notifySendRequest).Execute()

Send an SMS notification



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
	sync := "sync_example" // string | Must be `true`; otherwise `503`.
	notifySendRequest := *openapiclient.NewNotifySendRequest([]string{"To_example"}) // NotifySendRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SendAPI.NotifyNotifySendSms(context.Background()).Sync(sync).NotifySendRequest(notifySendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SendAPI.NotifyNotifySendSms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotifyNotifySendSms`: NotifyNotifySend200Response
	fmt.Fprintf(os.Stdout, "Response from `SendAPI.NotifyNotifySendSms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotifyNotifySendSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sync** | **string** | Must be &#x60;true&#x60;; otherwise &#x60;503&#x60;. | 
 **notifySendRequest** | [**NotifySendRequest**](NotifySendRequest.md) |  | 

### Return type

[**NotifyNotifySend200Response**](NotifyNotifySend200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

