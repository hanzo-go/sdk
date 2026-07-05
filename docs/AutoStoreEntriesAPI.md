# \AutoStoreEntriesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutoDeleteStoreEntry**](AutoStoreEntriesAPI.md#AutoDeleteStoreEntry) | **Delete** /v1/auto/store-entries/{key} | Delete a store entry
[**AutoGetStoreEntry**](AutoStoreEntriesAPI.md#AutoGetStoreEntry) | **Get** /v1/auto/store-entries/{key} | Get a store entry by key
[**AutoListStoreEntries**](AutoStoreEntriesAPI.md#AutoListStoreEntries) | **Get** /v1/auto/store-entries | List store entries
[**AutoUpsertStoreEntry**](AutoStoreEntriesAPI.md#AutoUpsertStoreEntry) | **Post** /v1/auto/store-entries | Create or update a store entry



## AutoDeleteStoreEntry

> AutoDeleteStoreEntry(ctx, key).Execute()

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
	r, err := apiClient.AutoStoreEntriesAPI.AutoDeleteStoreEntry(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoStoreEntriesAPI.AutoDeleteStoreEntry``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAutoDeleteStoreEntryRequest struct via the builder pattern


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


## AutoGetStoreEntry

> map[string]interface{} AutoGetStoreEntry(ctx, key).Execute()

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
	resp, r, err := apiClient.AutoStoreEntriesAPI.AutoGetStoreEntry(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoStoreEntriesAPI.AutoGetStoreEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoGetStoreEntry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoStoreEntriesAPI.AutoGetStoreEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutoGetStoreEntryRequest struct via the builder pattern


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


## AutoListStoreEntries

> map[string]interface{} AutoListStoreEntries(ctx).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoStoreEntriesAPI.AutoListStoreEntries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoStoreEntriesAPI.AutoListStoreEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoListStoreEntries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoStoreEntriesAPI.AutoListStoreEntries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutoListStoreEntriesRequest struct via the builder pattern


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


## AutoUpsertStoreEntry

> map[string]interface{} AutoUpsertStoreEntry(ctx).AutoUpsertStoreEntryRequest(autoUpsertStoreEntryRequest).Execute()

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
	resp, r, err := apiClient.AutoStoreEntriesAPI.AutoUpsertStoreEntry(context.Background()).AutoUpsertStoreEntryRequest(autoUpsertStoreEntryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoStoreEntriesAPI.AutoUpsertStoreEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutoUpsertStoreEntry`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AutoStoreEntriesAPI.AutoUpsertStoreEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutoUpsertStoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoUpsertStoreEntryRequest** | [**AutoUpsertStoreEntryRequest**](AutoUpsertStoreEntryRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

