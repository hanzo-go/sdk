# \FlagAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFlagDefsByKey**](FlagAPI.md#DeleteFlagDefsByKey) | **Delete** /v1/flag/defs/{key} | Removes one flag definition by key and records the deletion in the change log.
[**GetFlagActivity**](FlagAPI.md#GetFlagActivity) | **Get** /v1/flag/activity | Returns the caller&#39;s flag change log newest-first: every create, update and delete, with the actor and the time.
[**GetFlagDefs**](FlagAPI.md#GetFlagDefs) | **Get** /v1/flag/defs | Returns every flag definition in the caller&#39;s (org, project) store, by key, with its version and who last changed it.
[**GetFlagDefsByKey**](FlagAPI.md#GetFlagDefsByKey) | **Get** /v1/flag/defs/{key} | Returns one flag definition by key, or 404 when the caller&#39;s store has none under that key.
[**GetFlagHealth**](FlagAPI.md#GetFlagHealth) | **Get** /v1/flag/health | Health reports that the flag engine is serving.
[**PostFlag**](FlagAPI.md#PostFlag) | **Post** /v1/flag | Evaluate runs the caller&#39;s flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**PostFlagDecide**](FlagAPI.md#PostFlagDecide) | **Post** /v1/flag/decide | Evaluate runs the caller&#39;s flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**PutFlagDefsByKey**](FlagAPI.md#PutFlagDefsByKey) | **Put** /v1/flag/defs/{key} | Creates or replaces the flag definition at the path&#39;s key and returns the stored row.



## DeleteFlagDefsByKey

> DeletedOut DeleteFlagDefsByKey(ctx, key).Execute()

Removes one flag definition by key and records the deletion in the change log.



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
	key := "key_example" // string | Key is the flag key to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.DeleteFlagDefsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.DeleteFlagDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFlagDefsByKey`: DeletedOut
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.DeleteFlagDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlagDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeletedOut**](DeletedOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagActivity

> ActivityOut GetFlagActivity(ctx).Limit(limit).Execute()

Returns the caller's flag change log newest-first: every create, update and delete, with the actor and the time.



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
	limit := int64(789) // int64 | Limit caps the rows returned. 1–500; anything else takes the default 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.GetFlagActivity(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.GetFlagActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagActivity`: ActivityOut
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.GetFlagActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int64** | Limit caps the rows returned. 1–500; anything else takes the default 100. | 

### Return type

[**ActivityOut**](ActivityOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagDefs

> DefsOut GetFlagDefs(ctx).Execute()

Returns every flag definition in the caller's (org, project) store, by key, with its version and who last changed it.



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
	resp, r, err := apiClient.FlagAPI.GetFlagDefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.GetFlagDefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagDefs`: DefsOut
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.GetFlagDefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagDefsRequest struct via the builder pattern


### Return type

[**DefsOut**](DefsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagDefsByKey

> DefRow GetFlagDefsByKey(ctx, key).Execute()

Returns one flag definition by key, or 404 when the caller's store has none under that key.



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
	key := "key_example" // string | Key is the flag key to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.GetFlagDefsByKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.GetFlagDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagDefsByKey`: DefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.GetFlagDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefRow**](DefRow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlagHealth

> HealthOut GetFlagHealth(ctx).Execute()

Health reports that the flag engine is serving.



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
	resp, r, err := apiClient.FlagAPI.GetFlagHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.GetFlagHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlagHealth`: HealthOut
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.GetFlagHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlagHealthRequest struct via the builder pattern


### Return type

[**HealthOut**](HealthOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFlag

> interface{} PostFlag(ctx).EvaluateIn(evaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	evaluateIn := *openapiclient.NewEvaluateIn() // EvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.PostFlag(context.Background()).EvaluateIn(evaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.PostFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlag`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.PostFlag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateIn** | [**EvaluateIn**](EvaluateIn.md) |  | 

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


## PostFlagDecide

> interface{} PostFlagDecide(ctx).EvaluateIn(evaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the flag verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	evaluateIn := *openapiclient.NewEvaluateIn() // EvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.PostFlagDecide(context.Background()).EvaluateIn(evaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.PostFlagDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFlagDecide`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.PostFlagDecide`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFlagDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evaluateIn** | [**EvaluateIn**](EvaluateIn.md) |  | 

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


## PutFlagDefsByKey

> DefRow PutFlagDefsByKey(ctx, key).Body(body).Execute()

Creates or replaces the flag definition at the path's key and returns the stored row.



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
	key := "new-editor" // string | Key is the flag key to write, from the path.
	body := interface{}({"active":true,"filters":{"groups":[{"rollout_percentage":25}]},"key":"new-editor"}) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagAPI.PutFlagDefsByKey(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagAPI.PutFlagDefsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFlagDefsByKey`: DefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagAPI.PutFlagDefsByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to write, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFlagDefsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**DefRow**](DefRow.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

