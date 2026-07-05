# \FlowStoreEntriesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FlowDeleteStoreEntry**](FlowStoreEntriesAPI.md#FlowDeleteStoreEntry) | **Delete** /v1/flow/store-entries/{key} | Delete a store entry
[**FlowGetStoreEntry**](FlowStoreEntriesAPI.md#FlowGetStoreEntry) | **Get** /v1/flow/store-entries/{key} | Get a store entry by key
[**FlowListStoreEntries**](FlowStoreEntriesAPI.md#FlowListStoreEntries) | **Get** /v1/flow/store-entries | List store entries
[**FlowUpsertStoreEntry**](FlowStoreEntriesAPI.md#FlowUpsertStoreEntry) | **Post** /v1/flow/store-entries | Create or update a store entry



## FlowDeleteStoreEntry

> FlowDeleteStoreEntry(ctx, key).Execute()

Delete a store entry

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FlowStoreEntriesAPI.FlowDeleteStoreEntry(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowStoreEntriesAPI.FlowDeleteStoreEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowDeleteStoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlowGetStoreEntry

> map[string]interface{} FlowGetStoreEntry(ctx, key).Execute()

Get a store entry by key

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
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowStoreEntriesAPI.FlowGetStoreEntry(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowStoreEntriesAPI.FlowGetStoreEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowGetStoreEntry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowStoreEntriesAPI.FlowGetStoreEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlowGetStoreEntryRequest struct via the builder pattern


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


## FlowListStoreEntries

> map[string]interface{} FlowListStoreEntries(ctx).Cursor(cursor).Limit(limit).Execute()

List store entries

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
	cursor := "cursor_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowStoreEntriesAPI.FlowListStoreEntries(context.Background()).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowStoreEntriesAPI.FlowListStoreEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowListStoreEntries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FlowStoreEntriesAPI.FlowListStoreEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowListStoreEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

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


## FlowUpsertStoreEntry

> FlowStoreEntry FlowUpsertStoreEntry(ctx).AutoUpsertStoreEntryRequest(autoUpsertStoreEntryRequest).Execute()

Create or update a store entry

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
	autoUpsertStoreEntryRequest := *openapiclient.NewAutoUpsertStoreEntryRequest("Key_example", interface{}(123)) // AutoUpsertStoreEntryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlowStoreEntriesAPI.FlowUpsertStoreEntry(context.Background()).AutoUpsertStoreEntryRequest(autoUpsertStoreEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlowStoreEntriesAPI.FlowUpsertStoreEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FlowUpsertStoreEntry`: FlowStoreEntry
	fmt.Fprintf(os.Stdout, "Response from `FlowStoreEntriesAPI.FlowUpsertStoreEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFlowUpsertStoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoUpsertStoreEntryRequest** | [**AutoUpsertStoreEntryRequest**](AutoUpsertStoreEntryRequest.md) |  | 

### Return type

[**FlowStoreEntry**](FlowStoreEntry.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

