# \PubsubObjectStoreAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubDeleteObject**](PubsubObjectStoreAPI.md#PubsubDeleteObject) | **Delete** /v1/pubsub/objects/{bucket}/{name} | Delete an object
[**PubsubGetObject**](PubsubObjectStoreAPI.md#PubsubGetObject) | **Get** /v1/pubsub/objects/{bucket}/{name} | Download an object
[**PubsubListObjects**](PubsubObjectStoreAPI.md#PubsubListObjects) | **Get** /v1/pubsub/objects/{bucket} | List objects in a bucket
[**PubsubPutObject**](PubsubObjectStoreAPI.md#PubsubPutObject) | **Put** /v1/pubsub/objects/{bucket}/{name} | Upload an object



## PubsubDeleteObject

> PubsubDeleteObject(ctx, bucket, name).Execute()

Delete an object

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
	bucket := "bucket_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PubsubObjectStoreAPI.PubsubDeleteObject(context.Background(), bucket, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubObjectStoreAPI.PubsubDeleteObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubDeleteObjectRequest struct via the builder pattern


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


## PubsubGetObject

> *os.File PubsubGetObject(ctx, bucket, name).Execute()

Download an object

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
	bucket := "bucket_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubObjectStoreAPI.PubsubGetObject(context.Background(), bucket, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubObjectStoreAPI.PubsubGetObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubGetObject`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PubsubObjectStoreAPI.PubsubGetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubGetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubListObjects

> PubsubListObjects200Response PubsubListObjects(ctx, bucket).Execute()

List objects in a bucket

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
	bucket := "bucket_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubObjectStoreAPI.PubsubListObjects(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubObjectStoreAPI.PubsubListObjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubListObjects`: PubsubListObjects200Response
	fmt.Fprintf(os.Stdout, "Response from `PubsubObjectStoreAPI.PubsubListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PubsubListObjects200Response**](PubsubListObjects200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PubsubPutObject

> PubsubObjectMeta PubsubPutObject(ctx, bucket, name).Body(body).Description(description).Execute()

Upload an object



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
	bucket := "bucket_example" // string | 
	name := "name_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | 
	description := "description_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PubsubObjectStoreAPI.PubsubPutObject(context.Background(), bucket, name).Body(body).Description(description).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PubsubObjectStoreAPI.PubsubPutObject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PubsubPutObject`: PubsubObjectMeta
	fmt.Fprintf(os.Stdout, "Response from `PubsubObjectStoreAPI.PubsubPutObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPubsubPutObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 
 **description** | **string** |  | 

### Return type

[**PubsubObjectMeta**](PubsubObjectMeta.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

