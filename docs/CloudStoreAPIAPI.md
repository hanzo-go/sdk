# \CloudStoreAPIAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudApiControllerAddStore**](CloudStoreAPIAPI.md#CloudApiControllerAddStore) | **Post** /v1/cloud/add-store | Api Controller Add Store
[**CloudApiControllerDeleteStore**](CloudStoreAPIAPI.md#CloudApiControllerDeleteStore) | **Post** /v1/cloud/delete-store | Api Controller Delete Store
[**CloudApiControllerGetGlobalStores**](CloudStoreAPIAPI.md#CloudApiControllerGetGlobalStores) | **Get** /v1/cloud/get-global-stores | Api Controller Get Global Stores
[**CloudApiControllerGetStore**](CloudStoreAPIAPI.md#CloudApiControllerGetStore) | **Get** /v1/cloud/get-store | Api Controller Get Store
[**CloudApiControllerGetStores**](CloudStoreAPIAPI.md#CloudApiControllerGetStores) | **Get** /v1/cloud/get-stores | Api Controller Get Stores
[**CloudApiControllerRefreshStoreVectors**](CloudStoreAPIAPI.md#CloudApiControllerRefreshStoreVectors) | **Post** /v1/cloud/refresh-store-vectors | Api Controller Refresh Store Vectors
[**CloudApiControllerUpdateStore**](CloudStoreAPIAPI.md#CloudApiControllerUpdateStore) | **Post** /v1/cloud/update-store | Api Controller Update Store



## CloudApiControllerAddStore

> CloudControllersResponse CloudApiControllerAddStore(ctx).CloudObjectStore(cloudObjectStore).Execute()

Api Controller Add Store



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
	cloudObjectStore := *openapiclient.NewCloudObjectStore() // CloudObjectStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerAddStore(context.Background()).CloudObjectStore(cloudObjectStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerAddStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerAddStore`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerAddStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerAddStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectStore** | [**CloudObjectStore**](CloudObjectStore.md) | The details of the store | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerDeleteStore

> CloudControllersResponse CloudApiControllerDeleteStore(ctx).CloudObjectStore(cloudObjectStore).Execute()

Api Controller Delete Store



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
	cloudObjectStore := *openapiclient.NewCloudObjectStore() // CloudObjectStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerDeleteStore(context.Background()).CloudObjectStore(cloudObjectStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerDeleteStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerDeleteStore`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerDeleteStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectStore** | [**CloudObjectStore**](CloudObjectStore.md) | The details of the store | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetGlobalStores

> []CloudObjectStore CloudApiControllerGetGlobalStores(ctx).Execute()

Api Controller Get Global Stores



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
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerGetGlobalStores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerGetGlobalStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetGlobalStores`: []CloudObjectStore
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerGetGlobalStores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetGlobalStoresRequest struct via the builder pattern


### Return type

[**[]CloudObjectStore**](CloudObjectStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetStore

> CloudObjectStore CloudApiControllerGetStore(ctx).Id(id).Execute()

Api Controller Get Store



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
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerGetStore(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerGetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetStore`: CloudObjectStore
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerGetStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the store | 

### Return type

[**CloudObjectStore**](CloudObjectStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerGetStores

> []CloudObjectStore CloudApiControllerGetStores(ctx).Owner(owner).Execute()

Api Controller Get Stores



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
	owner := "owner_example" // string | The owner of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerGetStores(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerGetStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerGetStores`: []CloudObjectStore
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerGetStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerGetStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the store | 

### Return type

[**[]CloudObjectStore**](CloudObjectStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerRefreshStoreVectors

> CloudControllersResponse CloudApiControllerRefreshStoreVectors(ctx).CloudObjectStore(cloudObjectStore).Execute()

Api Controller Refresh Store Vectors



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
	cloudObjectStore := *openapiclient.NewCloudObjectStore() // CloudObjectStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerRefreshStoreVectors(context.Background()).CloudObjectStore(cloudObjectStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerRefreshStoreVectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerRefreshStoreVectors`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerRefreshStoreVectors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerRefreshStoreVectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudObjectStore** | [**CloudObjectStore**](CloudObjectStore.md) | The details of the store | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudApiControllerUpdateStore

> CloudControllersResponse CloudApiControllerUpdateStore(ctx).Id(id).CloudObjectStore(cloudObjectStore).Execute()

Api Controller Update Store



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
	cloudObjectStore := *openapiclient.NewCloudObjectStore() // CloudObjectStore | The details of the store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudStoreAPIAPI.CloudApiControllerUpdateStore(context.Background()).Id(id).CloudObjectStore(cloudObjectStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudStoreAPIAPI.CloudApiControllerUpdateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudApiControllerUpdateStore`: CloudControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `CloudStoreAPIAPI.CloudApiControllerUpdateStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudApiControllerUpdateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id (owner/name) of the store | 
 **cloudObjectStore** | [**CloudObjectStore**](CloudObjectStore.md) | The details of the store | 

### Return type

[**CloudControllersResponse**](CloudControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

