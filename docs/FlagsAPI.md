# \FlagsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlagsDefsByKey**](FlagsAPI.md#DeleteFlagsDefsByKey) | **Delete** /v1/flags/defs/{key} | Removes one flag definition by key and records the deletion in the change log.
[**GetFlagsActivity**](FlagsAPI.md#GetFlagsActivity) | **Get** /v1/flags/activity | Returns the caller&#39;s flag change log newest-first: every create, update and delete, with the actor and the time.
[**GetFlagsDefs**](FlagsAPI.md#GetFlagsDefs) | **Get** /v1/flags/defs | Returns every flag definition in the caller&#39;s (org, project) store, by key, with its version and who last changed it.
[**GetFlagsDefsByKey**](FlagsAPI.md#GetFlagsDefsByKey) | **Get** /v1/flags/defs/{key} | Returns one flag definition by key, or 404 when the caller&#39;s store has none under that key.
[**GetFlagsHealth**](FlagsAPI.md#GetFlagsHealth) | **Get** /v1/flags/health | Health reports that the flag engine is serving.
[**GetFlagsWaitlist**](FlagsAPI.md#GetFlagsWaitlist) | **Get** /v1/flags/waitlist | Reports whether ONE host is currently gated by the launch waitlist.
[**PostFlags**](FlagsAPI.md#PostFlags) | **Post** /v1/flags | Evaluate runs the caller&#39;s flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**PostFlagsDecide**](FlagsAPI.md#PostFlagsDecide) | **Post** /v1/flags/decide | Evaluate runs the caller&#39;s flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**PutFlagsDefsByKey**](FlagsAPI.md#PutFlagsDefsByKey) | **Put** /v1/flags/defs/{key} | Creates or replaces the flag definition at the path&#39;s key and returns the stored row.



## DeleteFlagsDefsByKey

> DeletedOut DeleteFlagsDefsByKey(ctx, key).Execute()

Removes one flag definition by key and records the deletion in the change log.



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
	key := "key_example" // string | Key is the flag key to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.DeleteFlagsDefsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.DeleteFlagsDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFlagsDefsByKey`: DeletedOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.DeleteFlagsDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagsDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedOut**](DeletedOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagsActivity

> ActivityOut GetFlagsActivity(ctx).Limit(limit).Execute()

Returns the caller's flag change log newest-first: every create, update and delete, with the actor and the time.



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
	limit := int32(56) // int32 | Limit caps the rows returned. 1–500; anything else takes the default 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.GetFlagsActivity(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.GetFlagsActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagsActivity`: ActivityOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.GetFlagsActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagsActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. 1–500; anything else takes the default 100. | 

### Return type

[**ActivityOut**](ActivityOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagsDefs

> DefsOut GetFlagsDefs(ctx).Execute()

Returns every flag definition in the caller's (org, project) store, by key, with its version and who last changed it.



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
	resp, r, err := apiClient.FlagsAPI.GetFlagsDefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.GetFlagsDefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagsDefs`: DefsOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.GetFlagsDefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagsDefsRequest struct via the builder pattern


### Return type

[**DefsOut**](DefsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagsDefsByKey

> DefRow GetFlagsDefsByKey(ctx, key).Execute()

Returns one flag definition by key, or 404 when the caller's store has none under that key.



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
	key := "key_example" // string | Key is the flag key to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.GetFlagsDefsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.GetFlagsDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagsDefsByKey`: DefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.GetFlagsDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagsDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefRow**](DefRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagsHealth

> HealthOut GetFlagsHealth(ctx).Execute()

Health reports that the flag engine is serving.



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
	resp, r, err := apiClient.FlagsAPI.GetFlagsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.GetFlagsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagsHealth`: HealthOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.GetFlagsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagsHealthRequest struct via the builder pattern


### Return type

[**HealthOut**](HealthOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagsWaitlist

> WaitlistModeView GetFlagsWaitlist(ctx).Host(host).Execute()

Reports whether ONE host is currently gated by the launch waitlist.



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
	host := "host_example" // string | Host is the host to resolve, e.g. \"chat.hanzo.ai\". Defaults to the request's own Host header when omitted, which is what lets a guard running on the governed host ask about itself with no argument. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.GetFlagsWaitlist(context.Background()).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.GetFlagsWaitlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagsWaitlist`: WaitlistModeView
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.GetFlagsWaitlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagsWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | Host is the host to resolve, e.g. \&quot;chat.hanzo.ai\&quot;. Defaults to the request&#39;s own Host header when omitted, which is what lets a guard running on the governed host ask about itself with no argument. | 

### Return type

[**WaitlistModeView**](WaitlistModeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlags

> interface{} PostFlags(ctx).EvaluateIn(evaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	evaluateIn := *openapiclient.NewEvaluateIn() // EvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.PostFlags(context.Background()).EvaluateIn(evaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.PostFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlags`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.PostFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateIn** | [**EvaluateIn**](EvaluateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlagsDecide

> interface{} PostFlagsDecide(ctx).EvaluateIn(evaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	evaluateIn := *openapiclient.NewEvaluateIn() // EvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.PostFlagsDecide(context.Background()).EvaluateIn(evaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.PostFlagsDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlagsDecide`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.PostFlagsDecide`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagsDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateIn** | [**EvaluateIn**](EvaluateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFlagsDefsByKey

> DefRow PutFlagsDefsByKey(ctx, key).Body(body).Execute()

Creates or replaces the flag definition at the path's key and returns the stored row.



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
	key := "new-editor" // string | Key is the flag key to write, from the path.
	body := interface{}({"active":true,"filters":{"groups":[{"rollout_percentage":25}]},"key":"new-editor"}) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.PutFlagsDefsByKey(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.PutFlagsDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFlagsDefsByKey`: DefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.PutFlagsDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to write, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagsDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**DefRow**](DefRow.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

