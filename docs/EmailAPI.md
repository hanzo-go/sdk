# \EmailAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmail**](EmailAPI.md#GetEmail) | **Get** /v1/email | Reads a page of the org&#39;s ledger, newest first: every message sent through POST /v1/email, and every message SES or SendGrid told of by its events, sent from anywhere on the org&#39;s account.
[**GetEmailById**](EmailAPI.md#GetEmailById) | **Get** /v1/email/{id} | Reads one message of the org&#39;s ledger, sent through this API or told of by its provider, with what became of it.
[**GetEmailSendgridStatus**](EmailAPI.md#GetEmailSendgridStatus) | **Get** /v1/email/sendgrid/status | Reports the org&#39;s SendGrid configuration: whether it holds the credentials a send needs, the account&#39;s credits as SendGrid answers them, what the org&#39;s ledger counted sent by SendGrid in the last 24 hours, and the DNS records of the sending domain.
[**GetEmailSesStatus**](EmailAPI.md#GetEmailSesStatus) | **Get** /v1/email/ses/status | Reports the org&#39;s Amazon SES configuration: whether it holds the credentials a send needs, the account&#39;s limits as SES answers them (sandbox, sending enabled, 24-hour quota, rate, sent in the last 24 hours), what the org&#39;s ledger counted sent by SES in the last 24 hours, and the DNS records of the sending domain.
[**GetEmailStats**](EmailAPI.md#GetEmailStats) | **Get** /v1/email/stats | Counts the org&#39;s mail over a range: sent, delivered, opened, clicked, bounced, complained and unsubscribed, each a count of messages, in all, by provider, by campaign and by day or hour (UTC).
[**PostEmail**](EmailAPI.md#PostEmail) | **Post** /v1/email | Sends commercial email from the org through SES or SendGrid: one message per recipient, each carrying its own signed unsubscribe link in the footer and in List-Unsubscribe and List-Unsubscribe-Post (RFC 8058), and every http(s) link tagged with utm_source&#x3D;email, utm_medium&#x3D;&lt;provider&gt;, utm_campaign&#x3D;&lt;campaign&gt;, utm_content&#x3D;link&lt;n&gt; and from&#x3D;&lt;message id&gt;, each only where the link does not carry it already.
[**PostEmailVerify**](EmailAPI.md#PostEmailVerify) | **Post** /v1/email/verify | Evaluates an email address using explicitly selected, composable checks.
[**PostEmailWebhooksSendgrid**](EmailAPI.md#PostEmailWebhooksSendgrid) | **Post** /v1/email/webhooks/sendgrid | Receives SendGrid&#39;s Signed Event Webhook: the delivery is an array of events, signed (ECDSA P-256, SHA-256) over its timestamp header followed by the body, under the public key the organization holds (the webhook-public-key credential).
[**PostEmailWebhooksSes**](EmailAPI.md#PostEmailWebhooksSes) | **Post** /v1/email/webhooks/ses | Receives Amazon SES events as an SNS HTTPS subscription delivers them.
[**PutEmailProvider**](EmailAPI.md#PutEmailProvider) | **Put** /v1/email/provider | Sets the provider the org&#39;s mail leaves by when a send pins none: ses or sendgrid.



## GetEmail

> EmailListMessagesResult GetEmail(ctx).Limit(limit).Status(status).Provider(provider).Since(since).Until(until).Recipient(recipient).Cursor(cursor).Execute()

Reads a page of the org's ledger, newest first: every message sent through POST /v1/email, and every message SES or SendGrid told of by its events, sent from anywhere on the org's account.



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
	limit := int64(789) // int64 | Limit is the most messages on the page (default 50, max 200). (optional)
	status := "status_example" // string | Status is one status to list: queued, sent, failed, delivered, opened, clicked, bounced or complained. (optional)
	provider := "provider_example" // string | Provider is ses or sendgrid. (optional)
	since := "since_example" // string | Since is the earliest send time listed, RFC 3339. (optional)
	until := "until_example" // string | Until is the send time listing stops before, RFC 3339. (optional)
	recipient := "recipient_example" // string | Recipient lists the messages to one address. (optional)
	cursor := "cursor_example" // string | Cursor is the next of the page before; empty is the newest page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetEmail(context.Background()).Limit(limit).Status(status).Provider(provider).Since(since).Until(until).Recipient(recipient).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmail`: EmailListMessagesResult
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit is the most messages on the page (default 50, max 200). | 
 **status** | **string** | Status is one status to list: queued, sent, failed, delivered, opened, clicked, bounced or complained. | 
 **provider** | **string** | Provider is ses or sendgrid. | 
 **since** | **string** | Since is the earliest send time listed, RFC 3339. | 
 **until** | **string** | Until is the send time listing stops before, RFC 3339. | 
 **recipient** | **string** | Recipient lists the messages to one address. | 
 **cursor** | **string** | Cursor is the next of the page before; empty is the newest page. | 

### Return type

[**EmailListMessagesResult**](EmailListMessagesResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailById

> EmailMessageRecord GetEmailById(ctx, id).Execute()

Reads one message of the org's ledger, sent through this API or told of by its provider, with what became of it.



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
	id := "id_example" // string | ID is our id for the message (msg_…).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetEmailById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetEmailById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailById`: EmailMessageRecord
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetEmailById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is our id for the message (msg_…). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EmailMessageRecord**](EmailMessageRecord.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSendgridStatus

> EmailSendGridStatusResult GetEmailSendgridStatus(ctx).Domain(domain).Execute()

Reports the org's SendGrid configuration: whether it holds the credentials a send needs, the account's credits as SendGrid answers them, what the org's ledger counted sent by SendGrid in the last 24 hours, and the DNS records of the sending domain.



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
	domain := "domain_example" // string | Domain is an optional sending domain to generate SendGrid whitelabel DKIM, Return-Path, and DMARC records for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetEmailSendgridStatus(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetEmailSendgridStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSendgridStatus`: EmailSendGridStatusResult
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetEmailSendgridStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSendgridStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain is an optional sending domain to generate SendGrid whitelabel DKIM, Return-Path, and DMARC records for. | 

### Return type

[**EmailSendGridStatusResult**](EmailSendGridStatusResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSesStatus

> EmailSESStatusResult GetEmailSesStatus(ctx).Domain(domain).Execute()

Reports the org's Amazon SES configuration: whether it holds the credentials a send needs, the account's limits as SES answers them (sandbox, sending enabled, 24-hour quota, rate, sent in the last 24 hours), what the org's ledger counted sent by SES in the last 24 hours, and the DNS records of the sending domain.



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
	domain := "domain_example" // string | Domain optionally specifies the sending domain to generate Easy DKIM and MAIL FROM records for. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetEmailSesStatus(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetEmailSesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSesStatus`: EmailSESStatusResult
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetEmailSesStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain optionally specifies the sending domain to generate Easy DKIM and MAIL FROM records for. | 

### Return type

[**EmailSESStatusResult**](EmailSESStatusResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailStats

> EmailStats GetEmailStats(ctx).Since(since).Until(until).Interval(interval).Provider(provider).Campaign(campaign).Execute()

Counts the org's mail over a range: sent, delivered, opened, clicked, bounced, complained and unsubscribed, each a count of messages, in all, by provider, by campaign and by day or hour (UTC).



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
	since := "since_example" // string | Since is where the range starts, RFC 3339, rounded down to its day or hour. Empty is 30 days before Until by day, 48 hours by hour. (optional)
	until := "until_example" // string | Until is where the range ends, RFC 3339, rounded up to its day or hour. Empty is now. (optional)
	interval := "interval_example" // string | Interval is day (default, at most 400 days) or hour (at most 31 days), in UTC. (optional)
	provider := "provider_example" // string | Provider counts one rail only: ses or sendgrid. (optional)
	campaign := "campaign_example" // string | Campaign counts one campaign only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.GetEmailStats(context.Background()).Since(since).Until(until).Interval(interval).Provider(provider).Campaign(campaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.GetEmailStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailStats`: EmailStats
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.GetEmailStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **string** | Since is where the range starts, RFC 3339, rounded down to its day or hour. Empty is 30 days before Until by day, 48 hours by hour. | 
 **until** | **string** | Until is where the range ends, RFC 3339, rounded up to its day or hour. Empty is now. | 
 **interval** | **string** | Interval is day (default, at most 400 days) or hour (at most 31 days), in UTC. | 
 **provider** | **string** | Provider counts one rail only: ses or sendgrid. | 
 **campaign** | **string** | Campaign counts one campaign only. | 

### Return type

[**EmailStats**](EmailStats.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmail

> EmailDispatch PostEmail(ctx).EmailSendMessageInput(emailSendMessageInput).IdempotencyKey(idempotencyKey).Execute()

Sends commercial email from the org through SES or SendGrid: one message per recipient, each carrying its own signed unsubscribe link in the footer and in List-Unsubscribe and List-Unsubscribe-Post (RFC 8058), and every http(s) link tagged with utm_source=email, utm_medium=<provider>, utm_campaign=<campaign>, utm_content=link<n> and from=<message id>, each only where the link does not carry it already.



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
	emailSendMessageInput := *openapiclient.NewEmailSendMessageInput() // EmailSendMessageInput | 
	idempotencyKey := "idempotencyKey_example" // string | IdempotencyKey names the send, so a retried request sends each recipient's message once: the same key and recipient name the same message. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.PostEmail(context.Background()).EmailSendMessageInput(emailSendMessageInput).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PostEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEmail`: EmailDispatch
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.PostEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailSendMessageInput** | [**EmailSendMessageInput**](EmailSendMessageInput.md) |  | 
 **idempotencyKey** | **string** | IdempotencyKey names the send, so a retried request sends each recipient&#39;s message once: the same key and recipient name the same message. | 

### Return type

[**EmailDispatch**](EmailDispatch.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmailVerify

> EmailVerificationResult PostEmailVerify(ctx).EmailVerificationInput(emailVerificationInput).Execute()

Evaluates an email address using explicitly selected, composable checks.



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
	emailVerificationInput := *openapiclient.NewEmailVerificationInput() // EmailVerificationInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.PostEmailVerify(context.Background()).EmailVerificationInput(emailVerificationInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PostEmailVerify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEmailVerify`: EmailVerificationResult
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.PostEmailVerify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEmailVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailVerificationInput** | [**EmailVerificationInput**](EmailVerificationInput.md) |  | 

### Return type

[**EmailVerificationResult**](EmailVerificationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEmailWebhooksSendgrid

> PostEmailWebhooksSendgrid(ctx).Execute()

Receives SendGrid's Signed Event Webhook: the delivery is an array of events, signed (ECDSA P-256, SHA-256) over its timestamp header followed by the body, under the public key the organization holds (the webhook-public-key credential).



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
	r, err := apiClient.EmailAPI.PostEmailWebhooksSendgrid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PostEmailWebhooksSendgrid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostEmailWebhooksSendgridRequest struct via the builder pattern


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


## PostEmailWebhooksSes

> PostEmailWebhooksSes(ctx).Execute()

Receives Amazon SES events as an SNS HTTPS subscription delivers them.



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
	r, err := apiClient.EmailAPI.PostEmailWebhooksSes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PostEmailWebhooksSes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostEmailWebhooksSesRequest struct via the builder pattern


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


## PutEmailProvider

> EmailChoice PutEmailProvider(ctx).EmailChoice(emailChoice).Execute()

Sets the provider the org's mail leaves by when a send pins none: ses or sendgrid.



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
	emailChoice := *openapiclient.NewEmailChoice() // EmailChoice | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailAPI.PutEmailProvider(context.Background()).EmailChoice(emailChoice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailAPI.PutEmailProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutEmailProvider`: EmailChoice
	fmt.Fprintf(os.Stdout, "Response from `EmailAPI.PutEmailProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutEmailProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailChoice** | [**EmailChoice**](EmailChoice.md) |  | 

### Return type

[**EmailChoice**](EmailChoice.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

