# \FormAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForm**](FormAPI.md#CreateForm) | **Post** /v1/form/forms | Makes a form in the caller&#39;s org and answers it as stored, its fields&#39; ids assigned and its defaults filled: kind form, visibility public, language en, one question at a time, and the platform&#39;s default sinks when it names none.
[**CreateFormTheme**](FormAPI.md#CreateFormTheme) | **Post** /v1/form/themes | Makes a theme in the caller&#39;s org: a name, four colours (#rrggbb), a font family and an https background image, each optional but the name.
[**CreateFormWebhook**](FormAPI.md#CreateFormWebhook) | **Post** /v1/form/forms/{id}/webhooks | Adds a webhook to a form: every completed response is posted to its https address as JSON, signed with Hanzo-Signature: t&#x3D;&lt;unix seconds&gt;,sha256&#x3D;&lt;base64 of HMAC-SHA256(secret, \&quot;&lt;seconds&gt;.\&quot; + body)&gt;; a receiver refuses a timestamp older than five minutes and an event_id it has seen.
[**DeleteForm**](FormAPI.md#DeleteForm) | **Delete** /v1/form/forms/{id} | Deletes a form with every response, webhook and delivery it has, and the webhooks&#39; signing secrets.
[**DeleteFormResponse**](FormAPI.md#DeleteFormResponse) | **Delete** /v1/form/forms/{id}/responses/{response} | Deletes one response of a form, and its webhook deliveries.
[**DeleteFormTheme**](FormAPI.md#DeleteFormTheme) | **Delete** /v1/form/themes/{id} | Deletes a theme.
[**DeleteFormWebhook**](FormAPI.md#DeleteFormWebhook) | **Delete** /v1/form/forms/{id}/webhooks/{webhook} | Deletes a webhook, its deliveries and its secret.
[**ExportFormResponses**](FormAPI.md#ExportFormResponses) | **Get** /v1/form/forms/{id}/export | Exports a form&#39;s responses as JSON or CSV, a page of up to 5000 at a time under the same filter a listing takes; next is the page to ask for after it, 0 at the end.
[**FileFormResponse**](FormAPI.md#FileFormResponse) | **Post** /v1/form/forms/{id}/responses/{response}/file | Files a completed response on the CRM&#39;s timeline now, and answers the response with whether it is filed and, when it is not, why.
[**GetForm**](FormAPI.md#GetForm) | **Get** /v1/form/forms/{id} | Reads one form of the caller&#39;s org, by its id or its slug.
[**GetFormResponse**](FormAPI.md#GetFormResponse) | **Get** /v1/form/forms/{id}/responses/{response} | Reads one response of a form.
[**GetFormTheme**](FormAPI.md#GetFormTheme) | **Get** /v1/form/themes/{id} | Reads one theme of the caller&#39;s org.
[**GetFormWebhook**](FormAPI.md#GetFormWebhook) | **Get** /v1/form/forms/{id}/webhooks/{webhook} | Reads one webhook of a form.
[**GetPublicForm**](FormAPI.md#GetPublicForm) | **Get** /v1/form/public/{id} | Reads a published form as its page draws it: the questions, the screens, the hidden values it takes, how it is answered and its theme.
[**ImportFormResponses**](FormAPI.md#ImportFormResponses) | **Post** /v1/form/forms/{id}/import | Brings responses made elsewhere into a form, keeping when each was made, a batch of at most 500 at a time.
[**ListFormDeliveries**](FormAPI.md#ListFormDeliveries) | **Get** /v1/form/forms/{id}/webhooks/{webhook}/deliveries | Lists a webhook&#39;s deliveries, newest first: each response posted, how many attempts it took, the last status the address answered and why it failed.
[**ListFormResponses**](FormAPI.md#ListFormResponses) | **Get** /v1/form/forms/{id}/responses | Lists a form&#39;s responses, newest first, a page at a time: since and until bound when they were made, completed keeps the finished or the open ones, query keeps the ones whose answers hold it, and included_response_ids and excluded_response_ids keep or drop named ones.
[**ListFormThemes**](FormAPI.md#ListFormThemes) | **Get** /v1/form/themes | Lists the caller org&#39;s themes by name.
[**ListFormWebhooks**](FormAPI.md#ListFormWebhooks) | **Get** /v1/form/forms/{id}/webhooks | Lists a form&#39;s webhooks.
[**ListForms**](FormAPI.md#ListForms) | **Get** /v1/form/forms | Lists the caller org&#39;s forms, newest first, a page at a time; search keeps the forms whose title holds it and tag the ones carrying it.
[**RedeliverFormWebhook**](FormAPI.md#RedeliverFormWebhook) | **Post** /v1/form/forms/{id}/webhooks/{webhook}/deliveries/{delivery}/retry | Posts a delivery&#39;s response to its webhook again, now, and answers how it went.
[**ReplaceForm**](FormAPI.md#ReplaceForm) | **Put** /v1/form/forms/{id} | Replaces a form whole and answers it as stored.
[**ReplaceFormTheme**](FormAPI.md#ReplaceFormTheme) | **Put** /v1/form/themes/{id} | Replaces a theme whole; every form drawn in it is drawn in the new one.
[**ReplaceFormWebhook**](FormAPI.md#ReplaceFormWebhook) | **Put** /v1/form/forms/{id}/webhooks/{webhook} | Replaces a webhook&#39;s tag, address and whether it is enabled.
[**SubmitFormResponse**](FormAPI.md#SubmitFormResponse) | **Post** /v1/form/public/{id}/responses | Answers a published form.
[**SummarizeForm**](FormAPI.md#SummarizeForm) | **Get** /v1/form/forms/{id}/summary | Sums up a form&#39;s responses: how many were begun and completed, and for each field how many answered it, the distribution of what was said (choices, points, yes and no), the mean, least and most of a number, and how many open responses stopped in front of it.
[**UpdateForm**](FormAPI.md#UpdateForm) | **Patch** /v1/form/forms/{id} | Changes the parts of a form the body names and keeps the rest, then answers it as stored.



## CreateForm

> FormForm CreateForm(ctx).FormForm(formForm).Execute()

Makes a form in the caller's org and answers it as stored, its fields' ids assigned and its defaults filled: kind form, visibility public, language en, one question at a time, and the platform's default sinks when it names none.



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
	formForm := *openapiclient.NewFormForm() // FormForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.CreateForm(context.Background()).FormForm(formForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.CreateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForm`: FormForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.CreateForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formForm** | [**FormForm**](FormForm.md) |  | 

### Return type

[**FormForm**](FormForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFormTheme

> FormTheme CreateFormTheme(ctx).FormTheme(formTheme).Execute()

Makes a theme in the caller's org: a name, four colours (#rrggbb), a font family and an https background image, each optional but the name.



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
	formTheme := *openapiclient.NewFormTheme() // FormTheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.CreateFormTheme(context.Background()).FormTheme(formTheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.CreateFormTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormTheme`: FormTheme
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.CreateFormTheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formTheme** | [**FormTheme**](FormTheme.md) |  | 

### Return type

[**FormTheme**](FormTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFormWebhook

> FormWebhook CreateFormWebhook(ctx, id).FormWebhookIn(formWebhookIn).Execute()

Adds a webhook to a form: every completed response is posted to its https address as JSON, signed with Hanzo-Signature: t=<unix seconds>,sha256=<base64 of HMAC-SHA256(secret, \"<seconds>.\" + body)>; a receiver refuses a timestamp older than five minutes and an event_id it has seen.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	formWebhookIn := *openapiclient.NewFormWebhookIn() // FormWebhookIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.CreateFormWebhook(context.Background(), id).FormWebhookIn(formWebhookIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.CreateFormWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormWebhook`: FormWebhook
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.CreateFormWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formWebhookIn** | [**FormWebhookIn**](FormWebhookIn.md) |  | 

### Return type

[**FormWebhook**](FormWebhook.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForm

> DeleteForm(ctx, id).Execute()

Deletes a form with every response, webhook and delivery it has, and the webhooks' signing secrets.



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
	id := "id_example" // string | ID is the form's id or its slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormAPI.DeleteForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.DeleteForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or its slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormRequest struct via the builder pattern


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


## DeleteFormResponse

> DeleteFormResponse(ctx, id, response).Execute()

Deletes one response of a form, and its webhook deliveries.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	response := "response_example" // string | Response is the response's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormAPI.DeleteFormResponse(context.Background(), id, response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.DeleteFormResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**response** | **string** | Response is the response&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormResponseRequest struct via the builder pattern


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


## DeleteFormTheme

> DeleteFormTheme(ctx, id).Execute()

Deletes a theme.



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
	id := "id_example" // string | ID is the theme's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormAPI.DeleteFormTheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.DeleteFormTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the theme&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormThemeRequest struct via the builder pattern


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


## DeleteFormWebhook

> DeleteFormWebhook(ctx, id, webhook).Execute()

Deletes a webhook, its deliveries and its secret.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	webhook := "webhook_example" // string | Webhook is the webhook's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FormAPI.DeleteFormWebhook(context.Background(), id, webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.DeleteFormWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**webhook** | **string** | Webhook is the webhook&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormWebhookRequest struct via the builder pattern


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


## ExportFormResponses

> FormExport ExportFormResponses(ctx, id).Page(page).PageSize(pageSize).Since(since).Until(until).Completed(completed).Query(query).Filed(filed).Sort(sort).IncludedResponseIds(includedResponseIds).ExcludedResponseIds(excludedResponseIds).Format(format).Execute()

Exports a form's responses as JSON or CSV, a page of up to 5000 at a time under the same filter a listing takes; next is the page to ask for after it, 0 at the end.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	page := int64(789) // int64 | Page is the page, from 1. Default 1. (optional)
	pageSize := int64(789) // int64 | PageSize is how many a page holds: at most 1000, default 25. (optional)
	since := "since_example" // string | Since keeps the responses made at or after it, RFC 3339. (optional)
	until := "until_example" // string | Until keeps the responses made before it, RFC 3339. (optional)
	completed := true // bool | Completed keeps the completed responses (true) or the open ones (false); absent keeps both. (optional)
	query := "query_example" // string | Query keeps the responses whose answers hold it. (optional)
	filed := true // bool | Filed keeps the responses on the CRM's timeline (true) or not yet on it (false); absent keeps both. (optional)
	sort := "sort_example" // string | Sort is submitted_at,desc (the default) or submitted_at,asc. (optional)
	includedResponseIds := "includedResponseIds_example" // string | IncludedResponseIDs keeps only these responses, comma separated, at most 90. (optional)
	excludedResponseIds := "excludedResponseIds_example" // string | ExcludedResponseIDs drops these responses, comma separated, at most 90. (optional)
	format := "format_example" // string | Format is json (the default) or csv. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ExportFormResponses(context.Background(), id).Page(page).PageSize(pageSize).Since(since).Until(until).Completed(completed).Query(query).Filed(filed).Sort(sort).IncludedResponseIds(includedResponseIds).ExcludedResponseIds(excludedResponseIds).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ExportFormResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFormResponses`: FormExport
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ExportFormResponses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFormResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page is the page, from 1. Default 1. | 
 **pageSize** | **int64** | PageSize is how many a page holds: at most 1000, default 25. | 
 **since** | **string** | Since keeps the responses made at or after it, RFC 3339. | 
 **until** | **string** | Until keeps the responses made before it, RFC 3339. | 
 **completed** | **bool** | Completed keeps the completed responses (true) or the open ones (false); absent keeps both. | 
 **query** | **string** | Query keeps the responses whose answers hold it. | 
 **filed** | **bool** | Filed keeps the responses on the CRM&#39;s timeline (true) or not yet on it (false); absent keeps both. | 
 **sort** | **string** | Sort is submitted_at,desc (the default) or submitted_at,asc. | 
 **includedResponseIds** | **string** | IncludedResponseIDs keeps only these responses, comma separated, at most 90. | 
 **excludedResponseIds** | **string** | ExcludedResponseIDs drops these responses, comma separated, at most 90. | 
 **format** | **string** | Format is json (the default) or csv. | 

### Return type

[**FormExport**](FormExport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileFormResponse

> FormResponse FileFormResponse(ctx, id, response).Execute()

Files a completed response on the CRM's timeline now, and answers the response with whether it is filed and, when it is not, why.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	response := "response_example" // string | Response is the response's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.FileFormResponse(context.Background(), id, response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.FileFormResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FileFormResponse`: FormResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.FileFormResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**response** | **string** | Response is the response&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFileFormResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormResponse**](FormResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForm

> FormForm GetForm(ctx, id).Execute()

Reads one form of the caller's org, by its id or its slug.



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
	id := "id_example" // string | ID is the form's id or its slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.GetForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.GetForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForm`: FormForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.GetForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or its slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormForm**](FormForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormResponse

> FormResponse GetFormResponse(ctx, id, response).Execute()

Reads one response of a form.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	response := "response_example" // string | Response is the response's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.GetFormResponse(context.Background(), id, response).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.GetFormResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormResponse`: FormResponse
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.GetFormResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**response** | **string** | Response is the response&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormResponse**](FormResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormTheme

> FormTheme GetFormTheme(ctx, id).Execute()

Reads one theme of the caller's org.



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
	id := "id_example" // string | ID is the theme's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.GetFormTheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.GetFormTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormTheme`: FormTheme
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.GetFormTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the theme&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormTheme**](FormTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormWebhook

> FormWebhook GetFormWebhook(ctx, id, webhook).Execute()

Reads one webhook of a form.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	webhook := "webhook_example" // string | Webhook is the webhook's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.GetFormWebhook(context.Background(), id, webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.GetFormWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormWebhook`: FormWebhook
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.GetFormWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**webhook** | **string** | Webhook is the webhook&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormWebhook**](FormWebhook.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicForm

> FormPublicForm GetPublicForm(ctx, id).Execute()

Reads a published form as its page draws it: the questions, the screens, the hidden values it takes, how it is answered and its theme.



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
	id := "id_example" // string | ID is the form's public id, \"<org>.<slug or id>\", from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.GetPublicForm(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.GetPublicForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicForm`: FormPublicForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.GetPublicForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s public id, \&quot;&lt;org&gt;.&lt;slug or id&gt;\&quot;, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormPublicForm**](FormPublicForm.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFormResponses

> FormImports ImportFormResponses(ctx, id).FormImportIn(formImportIn).Execute()

Brings responses made elsewhere into a form, keeping when each was made, a batch of at most 500 at a time.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	formImportIn := *openapiclient.NewFormImportIn() // FormImportIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ImportFormResponses(context.Background(), id).FormImportIn(formImportIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ImportFormResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportFormResponses`: FormImports
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ImportFormResponses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportFormResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formImportIn** | [**FormImportIn**](FormImportIn.md) |  | 

### Return type

[**FormImports**](FormImports.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFormDeliveries

> FormDeliveryList ListFormDeliveries(ctx, id, webhook).Page(page).PageSize(pageSize).Execute()

Lists a webhook's deliveries, newest first: each response posted, how many attempts it took, the last status the address answered and why it failed.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	webhook := "webhook_example" // string | Webhook is the webhook's id, from the path.
	page := int64(789) // int64 | Page is the page, from 1; PageSize at most 200, default 50. (optional)
	pageSize := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ListFormDeliveries(context.Background(), id, webhook).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ListFormDeliveries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFormDeliveries`: FormDeliveryList
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ListFormDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**webhook** | **string** | Webhook is the webhook&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFormDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int64** | Page is the page, from 1; PageSize at most 200, default 50. | 
 **pageSize** | **int64** |  | 

### Return type

[**FormDeliveryList**](FormDeliveryList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFormResponses

> FormResponseList ListFormResponses(ctx, id).Page(page).PageSize(pageSize).Since(since).Until(until).Completed(completed).Query(query).Filed(filed).Sort(sort).IncludedResponseIds(includedResponseIds).ExcludedResponseIds(excludedResponseIds).Execute()

Lists a form's responses, newest first, a page at a time: since and until bound when they were made, completed keeps the finished or the open ones, query keeps the ones whose answers hold it, and included_response_ids and excluded_response_ids keep or drop named ones.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	page := int64(789) // int64 | Page is the page, from 1. Default 1. (optional)
	pageSize := int64(789) // int64 | PageSize is how many a page holds: at most 1000, default 25. (optional)
	since := "since_example" // string | Since keeps the responses made at or after it, RFC 3339. (optional)
	until := "until_example" // string | Until keeps the responses made before it, RFC 3339. (optional)
	completed := true // bool | Completed keeps the completed responses (true) or the open ones (false); absent keeps both. (optional)
	query := "query_example" // string | Query keeps the responses whose answers hold it. (optional)
	filed := true // bool | Filed keeps the responses on the CRM's timeline (true) or not yet on it (false); absent keeps both. (optional)
	sort := "sort_example" // string | Sort is submitted_at,desc (the default) or submitted_at,asc. (optional)
	includedResponseIds := "includedResponseIds_example" // string | IncludedResponseIDs keeps only these responses, comma separated, at most 90. (optional)
	excludedResponseIds := "excludedResponseIds_example" // string | ExcludedResponseIDs drops these responses, comma separated, at most 90. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ListFormResponses(context.Background(), id).Page(page).PageSize(pageSize).Since(since).Until(until).Completed(completed).Query(query).Filed(filed).Sort(sort).IncludedResponseIds(includedResponseIds).ExcludedResponseIds(excludedResponseIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ListFormResponses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFormResponses`: FormResponseList
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ListFormResponses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFormResponsesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Page is the page, from 1. Default 1. | 
 **pageSize** | **int64** | PageSize is how many a page holds: at most 1000, default 25. | 
 **since** | **string** | Since keeps the responses made at or after it, RFC 3339. | 
 **until** | **string** | Until keeps the responses made before it, RFC 3339. | 
 **completed** | **bool** | Completed keeps the completed responses (true) or the open ones (false); absent keeps both. | 
 **query** | **string** | Query keeps the responses whose answers hold it. | 
 **filed** | **bool** | Filed keeps the responses on the CRM&#39;s timeline (true) or not yet on it (false); absent keeps both. | 
 **sort** | **string** | Sort is submitted_at,desc (the default) or submitted_at,asc. | 
 **includedResponseIds** | **string** | IncludedResponseIDs keeps only these responses, comma separated, at most 90. | 
 **excludedResponseIds** | **string** | ExcludedResponseIDs drops these responses, comma separated, at most 90. | 

### Return type

[**FormResponseList**](FormResponseList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFormThemes

> FormThemeList ListFormThemes(ctx).Execute()

Lists the caller org's themes by name.



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
	resp, r, err := apiClient.FormAPI.ListFormThemes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ListFormThemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFormThemes`: FormThemeList
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ListFormThemes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFormThemesRequest struct via the builder pattern


### Return type

[**FormThemeList**](FormThemeList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFormWebhooks

> FormWebhookList ListFormWebhooks(ctx, id).Execute()

Lists a form's webhooks.



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
	id := "id_example" // string | ID is the form's id or its slug, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ListFormWebhooks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ListFormWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFormWebhooks`: FormWebhookList
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ListFormWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or its slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFormWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormWebhookList**](FormWebhookList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListForms

> FormFormList ListForms(ctx).Page(page).PageSize(pageSize).Search(search).Tag(tag).Execute()

Lists the caller org's forms, newest first, a page at a time; search keeps the forms whose title holds it and tag the ones carrying it.



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
	page := int64(789) // int64 | Page is the page, from 1. Default 1. (optional)
	pageSize := int64(789) // int64 | PageSize is how many a page holds: at most 200, default 50. (optional)
	search := "search_example" // string | Search keeps the forms whose title holds it. (optional)
	tag := "tag_example" // string | Tag keeps the forms tagged with it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ListForms(context.Background()).Page(page).PageSize(pageSize).Search(search).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ListForms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListForms`: FormFormList
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ListForms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Page is the page, from 1. Default 1. | 
 **pageSize** | **int64** | PageSize is how many a page holds: at most 200, default 50. | 
 **search** | **string** | Search keeps the forms whose title holds it. | 
 **tag** | **string** | Tag keeps the forms tagged with it. | 

### Return type

[**FormFormList**](FormFormList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeliverFormWebhook

> FormDelivery RedeliverFormWebhook(ctx, id, webhook, delivery).Execute()

Posts a delivery's response to its webhook again, now, and answers how it went.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	webhook := "webhook_example" // string | Webhook is the webhook's id, from the path.
	delivery := "delivery_example" // string | Delivery is the delivery's id, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.RedeliverFormWebhook(context.Background(), id, webhook, delivery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.RedeliverFormWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RedeliverFormWebhook`: FormDelivery
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.RedeliverFormWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**webhook** | **string** | Webhook is the webhook&#39;s id, from the path. | 
**delivery** | **string** | Delivery is the delivery&#39;s id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeliverFormWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FormDelivery**](FormDelivery.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceForm

> FormForm ReplaceForm(ctx, id).FormForm(formForm).Execute()

Replaces a form whole and answers it as stored.



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
	id := "id_example" // string | ID is the form's id: 15 lowercase hex characters, derived from the slug when the form has one, so a slug names one form for good.
	formForm := *openapiclient.NewFormForm() // FormForm | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ReplaceForm(context.Background(), id).FormForm(formForm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ReplaceForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceForm`: FormForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ReplaceForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id: 15 lowercase hex characters, derived from the slug when the form has one, so a slug names one form for good. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formForm** | [**FormForm**](FormForm.md) |  | 

### Return type

[**FormForm**](FormForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFormTheme

> FormTheme ReplaceFormTheme(ctx, id).FormTheme(formTheme).Execute()

Replaces a theme whole; every form drawn in it is drawn in the new one.



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
	id := "id_example" // string | ID is the theme's id.
	formTheme := *openapiclient.NewFormTheme() // FormTheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ReplaceFormTheme(context.Background(), id).FormTheme(formTheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ReplaceFormTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFormTheme`: FormTheme
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ReplaceFormTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the theme&#39;s id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFormThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formTheme** | [**FormTheme**](FormTheme.md) |  | 

### Return type

[**FormTheme**](FormTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFormWebhook

> FormWebhook ReplaceFormWebhook(ctx, id, webhook).FormWebhookIn(formWebhookIn).Execute()

Replaces a webhook's tag, address and whether it is enabled.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	webhook := "webhook_example" // string | Webhook is the webhook's id, from the path of a replace.
	formWebhookIn := *openapiclient.NewFormWebhookIn() // FormWebhookIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.ReplaceFormWebhook(context.Background(), id, webhook).FormWebhookIn(formWebhookIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.ReplaceFormWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFormWebhook`: FormWebhook
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.ReplaceFormWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 
**webhook** | **string** | Webhook is the webhook&#39;s id, from the path of a replace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFormWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **formWebhookIn** | [**FormWebhookIn**](FormWebhookIn.md) |  | 

### Return type

[**FormWebhook**](FormWebhook.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFormResponse

> FormReceipt SubmitFormResponse(ctx, id).FormSubmission(formSubmission).Execute()

Answers a published form.



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
	id := "id_example" // string | ID is the form's public id, from the path.
	formSubmission := *openapiclient.NewFormSubmission() // FormSubmission | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.SubmitFormResponse(context.Background(), id).FormSubmission(formSubmission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.SubmitFormResponse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFormResponse`: FormReceipt
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.SubmitFormResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s public id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFormResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formSubmission** | [**FormSubmission**](FormSubmission.md) |  | 

### Return type

[**FormReceipt**](FormReceipt.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummarizeForm

> FormSummary SummarizeForm(ctx, id).Since(since).Until(until).Execute()

Sums up a form's responses: how many were begun and completed, and for each field how many answered it, the distribution of what was said (choices, points, yes and no), the mean, least and most of a number, and how many open responses stopped in front of it.



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
	id := "id_example" // string | ID is the form's id or slug, from the path.
	since := "since_example" // string | Since keeps the responses made at or after it, RFC 3339. (optional)
	until := "until_example" // string | Until keeps the responses made before it, RFC 3339. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.SummarizeForm(context.Background(), id).Since(since).Until(until).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.SummarizeForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeForm`: FormSummary
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.SummarizeForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Since keeps the responses made at or after it, RFC 3339. | 
 **until** | **string** | Until keeps the responses made before it, RFC 3339. | 

### Return type

[**FormSummary**](FormSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateForm

> FormForm UpdateForm(ctx, id).FormFormPatch(formFormPatch).Execute()

Changes the parts of a form the body names and keeps the rest, then answers it as stored.



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
	id := "id_example" // string | ID is the form's id or its slug, from the path.
	formFormPatch := *openapiclient.NewFormFormPatch() // FormFormPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FormAPI.UpdateForm(context.Background(), id).FormFormPatch(formFormPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormAPI.UpdateForm``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateForm`: FormForm
	fmt.Fprintf(os.Stdout, "Response from `FormAPI.UpdateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the form&#39;s id or its slug, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **formFormPatch** | [**FormFormPatch**](FormFormPatch.md) |  | 

### Return type

[**FormForm**](FormForm.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

