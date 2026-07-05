# \MqObjectStoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MqCreateObjectStore**](MqObjectStoreAPI.md#MqCreateObjectStore) | **Post** /v1/mq/objects | Create object store
[**MqDeleteObject**](MqObjectStoreAPI.md#MqDeleteObject) | **Delete** /v1/mq/objects/{store}/{name} | Delete object
[**MqDeleteObjectStore**](MqObjectStoreAPI.md#MqDeleteObjectStore) | **Delete** /v1/mq/objects/{store} | Delete object store
[**MqGetObject**](MqObjectStoreAPI.md#MqGetObject) | **Get** /v1/mq/objects/{store}/{name} | Download object
[**MqGetObjectStore**](MqObjectStoreAPI.md#MqGetObjectStore) | **Get** /v1/mq/objects/{store} | Get store info
[**MqListObjectStores**](MqObjectStoreAPI.md#MqListObjectStores) | **Get** /v1/mq/objects | List object stores
[**MqListObjects**](MqObjectStoreAPI.md#MqListObjects) | **Get** /v1/mq/objects/{store}/list | List objects in store
[**MqPutObject**](MqObjectStoreAPI.md#MqPutObject) | **Put** /v1/mq/objects/{store}/{name} | Upload object



## MqCreateObjectStore

> MqObjectStoreInfo MqCreateObjectStore(ctx).MqObjectStoreConfig(mqObjectStoreConfig).Execute()

Create object store



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
	mqObjectStoreConfig := *openapiclient.NewMqObjectStoreConfig("Name_example") // MqObjectStoreConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqCreateObjectStore(context.Background()).MqObjectStoreConfig(mqObjectStoreConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqCreateObjectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqCreateObjectStore`: MqObjectStoreInfo
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqCreateObjectStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqCreateObjectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mqObjectStoreConfig** | [**MqObjectStoreConfig**](MqObjectStoreConfig.md) |  | 

### Return type

[**MqObjectStoreInfo**](MqObjectStoreInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqDeleteObject

> MqDeleteObject(ctx, store, name).Execute()

Delete object



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
	store := "store_example" // string | Object store name.
	name := "name_example" // string | Object name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqObjectStoreAPI.MqDeleteObject(context.Background(), store, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqDeleteObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 
**name** | **string** | Object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqDeleteObjectStore

> MqDeleteObjectStore(ctx, store).Execute()

Delete object store



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
	store := "store_example" // string | Object store name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MqObjectStoreAPI.MqDeleteObjectStore(context.Background(), store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqDeleteObjectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqDeleteObjectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetObject

> *os.File MqGetObject(ctx, store, name).Execute()

Download object



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
	store := "store_example" // string | Object store name.
	name := "name_example" // string | Object name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqGetObject(context.Background(), store, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqGetObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqGetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 
**name** | **string** | Object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqGetObjectStore

> MqObjectStoreInfo MqGetObjectStore(ctx, store).Execute()

Get store info



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
	store := "store_example" // string | Object store name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqGetObjectStore(context.Background(), store).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqGetObjectStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqGetObjectStore`: MqObjectStoreInfo
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqGetObjectStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqGetObjectStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MqObjectStoreInfo**](MqObjectStoreInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListObjectStores

> MqListObjectStores200Response MqListObjectStores(ctx).Limit(limit).Offset(offset).Execute()

List object stores



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
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqListObjectStores(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqListObjectStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListObjectStores`: MqListObjectStores200Response
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqListObjectStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMqListObjectStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListObjectStores200Response**](MqListObjectStores200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqListObjects

> MqListObjects200Response MqListObjects(ctx, store).Limit(limit).Offset(offset).Execute()

List objects in store



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
	store := "store_example" // string | Object store name.
	limit := int32(56) // int32 | Maximum number of items to return. (optional) (default to 100)
	offset := int32(56) // int32 | Number of items to skip. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqListObjects(context.Background(), store).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqListObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqListObjects`: MqListObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **offset** | **int32** | Number of items to skip. | [default to 0]

### Return type

[**MqListObjects200Response**](MqListObjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MqPutObject

> MqObjectInfo MqPutObject(ctx, store, name).Body(body).XMQObjectDescription(xMQObjectDescription).Execute()

Upload object



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
	store := "store_example" // string | Object store name.
	name := "name_example" // string | Object name.
	body := os.NewFile(1234, "some_file") // *os.File | 
	xMQObjectDescription := "xMQObjectDescription_example" // string | Optional description for the object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MqObjectStoreAPI.MqPutObject(context.Background(), store, name).Body(body).XMQObjectDescription(xMQObjectDescription).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MqObjectStoreAPI.MqPutObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MqPutObject`: MqObjectInfo
	fmt.Fprintf(os.Stdout, "Response from `MqObjectStoreAPI.MqPutObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**store** | **string** | Object store name. | 
**name** | **string** | Object name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMqPutObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **xMQObjectDescription** | **string** | Optional description for the object. | 

### Return type

[**MqObjectInfo**](MqObjectInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

