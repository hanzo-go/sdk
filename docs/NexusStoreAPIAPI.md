# \NexusStoreAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NexusAddStore**](NexusStoreAPIAPI.md#NexusAddStore) | **Post** /v1/nexus/add-store | add Store
[**NexusDeleteStore**](NexusStoreAPIAPI.md#NexusDeleteStore) | **Post** /v1/nexus/delete-store | delete Store
[**NexusGetGlobalStores**](NexusStoreAPIAPI.md#NexusGetGlobalStores) | **Get** /v1/nexus/get-global-stores | get Global Stores
[**NexusGetStore**](NexusStoreAPIAPI.md#NexusGetStore) | **Get** /v1/nexus/get-store | get Store
[**NexusGetStores**](NexusStoreAPIAPI.md#NexusGetStores) | **Get** /v1/nexus/get-stores | get Stores
[**NexusRefreshStoreVectors**](NexusStoreAPIAPI.md#NexusRefreshStoreVectors) | **Post** /v1/nexus/refresh-store-vectors | refresh Store Vectors
[**NexusUpdateStore**](NexusStoreAPIAPI.md#NexusUpdateStore) | **Post** /v1/nexus/update-store | update Store



## NexusAddStore

> NexusResponse NexusAddStore(ctx).NexusStore(nexusStore).Execute()

add Store



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
	nexusStore := *openapiclient.NewNexusStore() // NexusStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusAddStore(context.Background()).NexusStore(nexusStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusAddStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusAddStore`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusAddStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusAddStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusStore** | [**NexusStore**](NexusStore.md) | The details of the store | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusDeleteStore

> NexusResponse NexusDeleteStore(ctx).NexusStore(nexusStore).Execute()

delete Store



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
	nexusStore := *openapiclient.NewNexusStore() // NexusStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusDeleteStore(context.Background()).NexusStore(nexusStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusDeleteStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusDeleteStore`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusDeleteStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusStore** | [**NexusStore**](NexusStore.md) | The details of the store | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetGlobalStores

> []NexusStore NexusGetGlobalStores(ctx).Execute()

get Global Stores



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
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusGetGlobalStores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusGetGlobalStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetGlobalStores`: []NexusStore
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusGetGlobalStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetGlobalStoresRequest struct via the builder pattern


### Return type

[**[]NexusStore**](NexusStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetStore

> NexusStore NexusGetStore(ctx).Id(id).Execute()

get Store



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
	id := "id_example" // string | The id (owner/name) of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusGetStore(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusGetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetStore`: NexusStore
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusGetStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the store | 

### Return type

[**NexusStore**](NexusStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusGetStores

> []NexusStore NexusGetStores(ctx).Owner(owner).Execute()

get Stores



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
	owner := "owner_example" // string | The owner of the stores

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusGetStores(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusGetStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusGetStores`: []NexusStore
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusGetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the stores | 

### Return type

[**[]NexusStore**](NexusStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusRefreshStoreVectors

> NexusResponse NexusRefreshStoreVectors(ctx).NexusStore(nexusStore).Execute()

refresh Store Vectors



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
	nexusStore := *openapiclient.NewNexusStore() // NexusStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusRefreshStoreVectors(context.Background()).NexusStore(nexusStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusRefreshStoreVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusRefreshStoreVectors`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusRefreshStoreVectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusRefreshStoreVectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nexusStore** | [**NexusStore**](NexusStore.md) | The details of the store | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NexusUpdateStore

> NexusResponse NexusUpdateStore(ctx).Id(id).NexusStore(nexusStore).Execute()

update Store



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
	id := "id_example" // string | The id (owner/name) of the store
	nexusStore := *openapiclient.NewNexusStore() // NexusStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NexusStoreAPIAPI.NexusUpdateStore(context.Background()).Id(id).NexusStore(nexusStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NexusStoreAPIAPI.NexusUpdateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NexusUpdateStore`: NexusResponse
	fmt.Fprintf(os.Stdout, "Response from `NexusStoreAPIAPI.NexusUpdateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNexusUpdateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the store | 
 **nexusStore** | [**NexusStore**](NexusStore.md) | The details of the store | 

### Return type

[**NexusResponse**](NexusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

