# \FlagsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1FlagsDefsKey**](FlagsAPI.md#CloudDeleteV1FlagsDefsKey) | **Delete** /v1/flags/defs/{key} | DeleteFlagDefinition removes one flag definition by key and records the deletion in the change log.
[**CloudGetV1FlagsActivity**](FlagsAPI.md#CloudGetV1FlagsActivity) | **Get** /v1/flags/activity | ListFlagActivity returns the caller&#39;s flag change log newest-first: every create, update and delete, with the actor and the time.
[**CloudGetV1FlagsDefs**](FlagsAPI.md#CloudGetV1FlagsDefs) | **Get** /v1/flags/defs | ListFlagDefinitions returns every flag definition in the caller&#39;s (org, project) store, by key, with its version and who last changed it.
[**CloudGetV1FlagsDefsKey**](FlagsAPI.md#CloudGetV1FlagsDefsKey) | **Get** /v1/flags/defs/{key} | GetFlagDefinition returns one flag definition by key, or 404 when the caller&#39;s store has none under that key.
[**CloudGetV1FlagsHealth**](FlagsAPI.md#CloudGetV1FlagsHealth) | **Get** /v1/flags/health | Health reports that the flag engine is serving.
[**CloudGetV1FlagsWaitlist**](FlagsAPI.md#CloudGetV1FlagsWaitlist) | **Get** /v1/flags/waitlist | WaitlistMode reports whether ONE host is currently gated by the launch waitlist.
[**CloudPostV1Flags**](FlagsAPI.md#CloudPostV1Flags) | **Post** /v1/flags | Evaluate runs the caller&#39;s flag definitions for one identity and returns the PostHog-shaped verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**CloudPostV1FlagsDecide**](FlagsAPI.md#CloudPostV1FlagsDecide) | **Post** /v1/flags/decide | Evaluate runs the caller&#39;s flag definitions for one identity and returns the PostHog-shaped verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.
[**CloudPutV1FlagsDefsKey**](FlagsAPI.md#CloudPutV1FlagsDefsKey) | **Put** /v1/flags/defs/{key} | PutFlagDefinition creates or replaces the flag definition at the path&#39;s key and returns the stored row.



## CloudDeleteV1FlagsDefsKey

> CloudDeletedOut CloudDeleteV1FlagsDefsKey(ctx, key).Execute()

DeleteFlagDefinition removes one flag definition by key and records the deletion in the change log.



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
	resp, r, err := apiClient.FlagsAPI.CloudDeleteV1FlagsDefsKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudDeleteV1FlagsDefsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1FlagsDefsKey`: CloudDeletedOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudDeleteV1FlagsDefsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1FlagsDefsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDeletedOut**](CloudDeletedOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlagsActivity

> CloudActivityOut CloudGetV1FlagsActivity(ctx).Limit(limit).Execute()

ListFlagActivity returns the caller's flag change log newest-first: every create, update and delete, with the actor and the time.



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
	resp, r, err := apiClient.FlagsAPI.CloudGetV1FlagsActivity(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudGetV1FlagsActivity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlagsActivity`: CloudActivityOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudGetV1FlagsActivity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlagsActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit caps the rows returned. 1–500; anything else takes the default 100. | 

### Return type

[**CloudActivityOut**](CloudActivityOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlagsDefs

> CloudDefsOut CloudGetV1FlagsDefs(ctx).Execute()

ListFlagDefinitions returns every flag definition in the caller's (org, project) store, by key, with its version and who last changed it.



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
	resp, r, err := apiClient.FlagsAPI.CloudGetV1FlagsDefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudGetV1FlagsDefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlagsDefs`: CloudDefsOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudGetV1FlagsDefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlagsDefsRequest struct via the builder pattern


### Return type

[**CloudDefsOut**](CloudDefsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlagsDefsKey

> CloudDefRow CloudGetV1FlagsDefsKey(ctx, key).Execute()

GetFlagDefinition returns one flag definition by key, or 404 when the caller's store has none under that key.



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
	resp, r, err := apiClient.FlagsAPI.CloudGetV1FlagsDefsKey(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudGetV1FlagsDefsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlagsDefsKey`: CloudDefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudGetV1FlagsDefsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlagsDefsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDefRow**](CloudDefRow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlagsHealth

> CloudHealthOut CloudGetV1FlagsHealth(ctx).Execute()

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
	resp, r, err := apiClient.FlagsAPI.CloudGetV1FlagsHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudGetV1FlagsHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlagsHealth`: CloudHealthOut
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudGetV1FlagsHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlagsHealthRequest struct via the builder pattern


### Return type

[**CloudHealthOut**](CloudHealthOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FlagsWaitlist

> CloudWaitlistModeView CloudGetV1FlagsWaitlist(ctx).Host(host).Execute()

WaitlistMode reports whether ONE host is currently gated by the launch waitlist.



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
	resp, r, err := apiClient.FlagsAPI.CloudGetV1FlagsWaitlist(context.Background()).Host(host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudGetV1FlagsWaitlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FlagsWaitlist`: CloudWaitlistModeView
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudGetV1FlagsWaitlist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FlagsWaitlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | Host is the host to resolve, e.g. \&quot;chat.hanzo.ai\&quot;. Defaults to the request&#39;s own Host header when omitted, which is what lets a guard running on the governed host ask about itself with no argument. | 

### Return type

[**CloudWaitlistModeView**](CloudWaitlistModeView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Flags

> interface{} CloudPostV1Flags(ctx).CloudEvaluateIn(cloudEvaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the PostHog-shaped verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	cloudEvaluateIn := *openapiclient.NewCloudEvaluateIn() // CloudEvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.CloudPostV1Flags(context.Background()).CloudEvaluateIn(cloudEvaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudPostV1Flags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Flags`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudPostV1Flags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEvaluateIn** | [**CloudEvaluateIn**](CloudEvaluateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FlagsDecide

> interface{} CloudPostV1FlagsDecide(ctx).CloudEvaluateIn(cloudEvaluateIn).Execute()

Evaluate runs the caller's flag definitions for one identity and returns the PostHog-shaped verdict: which flags are on (or which variant), their payloads, and whether any definition failed to compute.



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
	cloudEvaluateIn := *openapiclient.NewCloudEvaluateIn() // CloudEvaluateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlagsAPI.CloudPostV1FlagsDecide(context.Background()).CloudEvaluateIn(cloudEvaluateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudPostV1FlagsDecide``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FlagsDecide`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudPostV1FlagsDecide`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FlagsDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEvaluateIn** | [**CloudEvaluateIn**](CloudEvaluateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1FlagsDefsKey

> CloudDefRow CloudPutV1FlagsDefsKey(ctx, key).Body(body).Execute()

PutFlagDefinition creates or replaces the flag definition at the path's key and returns the stored row.



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
	resp, r, err := apiClient.FlagsAPI.CloudPutV1FlagsDefsKey(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlagsAPI.CloudPutV1FlagsDefsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1FlagsDefsKey`: CloudDefRow
	fmt.Fprintf(os.Stdout, "Response from `FlagsAPI.CloudPutV1FlagsDefsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key is the flag key to write, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1FlagsDefsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**CloudDefRow**](CloudDefRow.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

