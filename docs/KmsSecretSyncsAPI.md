# \KmsSecretSyncsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateSecretSync**](KmsSecretSyncsAPI.md#KmsCreateSecretSync) | **Post** /v1/kms/secret-syncs | Create a secret sync
[**KmsDeleteSecretSync**](KmsSecretSyncsAPI.md#KmsDeleteSecretSync) | **Delete** /v1/kms/secret-syncs/{syncId} | Delete a secret sync
[**KmsGetSecretSync**](KmsSecretSyncsAPI.md#KmsGetSecretSync) | **Get** /v1/kms/secret-syncs/{syncId} | Get a secret sync by ID
[**KmsListSecretSyncs**](KmsSecretSyncsAPI.md#KmsListSecretSyncs) | **Get** /v1/kms/secret-syncs | List secret syncs
[**KmsTriggerSecretSync**](KmsSecretSyncsAPI.md#KmsTriggerSecretSync) | **Post** /v1/kms/secret-syncs/{syncId}/trigger | Manually trigger a secret sync
[**KmsUpdateSecretSync**](KmsSecretSyncsAPI.md#KmsUpdateSecretSync) | **Patch** /v1/kms/secret-syncs/{syncId} | Update a secret sync



## KmsCreateSecretSync

> KmsCreateSecretSync200Response KmsCreateSecretSync(ctx).KmsCreateSecretSyncRequest(kmsCreateSecretSyncRequest).Execute()

Create a secret sync

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
	kmsCreateSecretSyncRequest := *openapiclient.NewKmsCreateSecretSyncRequest("Name_example", "Destination_example", "ProjectId_example", "SourceEnvironment_example", "SourcePath_example", "ConnectionId_example") // KmsCreateSecretSyncRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsCreateSecretSync(context.Background()).KmsCreateSecretSyncRequest(kmsCreateSecretSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsCreateSecretSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSecretSync`: KmsCreateSecretSync200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsCreateSecretSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSecretSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateSecretSyncRequest** | [**KmsCreateSecretSyncRequest**](KmsCreateSecretSyncRequest.md) |  | 

### Return type

[**KmsCreateSecretSync200Response**](KmsCreateSecretSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSecretSync

> map[string]interface{} KmsDeleteSecretSync(ctx, syncId).Execute()

Delete a secret sync

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
	syncId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsDeleteSecretSync(context.Background(), syncId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsDeleteSecretSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSecretSync`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsDeleteSecretSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSecretSyncRequest struct via the builder pattern


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


## KmsGetSecretSync

> KmsCreateSecretSync200Response KmsGetSecretSync(ctx, syncId).Execute()

Get a secret sync by ID

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
	syncId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsGetSecretSync(context.Background(), syncId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsGetSecretSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetSecretSync`: KmsCreateSecretSync200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsGetSecretSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetSecretSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsCreateSecretSync200Response**](KmsCreateSecretSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListSecretSyncs

> KmsListSecretSyncs200Response KmsListSecretSyncs(ctx).ProjectId(projectId).Execute()

List secret syncs

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
	projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsListSecretSyncs(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsListSecretSyncs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSecretSyncs`: KmsListSecretSyncs200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsListSecretSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSecretSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** |  | 

### Return type

[**KmsListSecretSyncs200Response**](KmsListSecretSyncs200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsTriggerSecretSync

> map[string]interface{} KmsTriggerSecretSync(ctx, syncId).Execute()

Manually trigger a secret sync

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
	syncId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsTriggerSecretSync(context.Background(), syncId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsTriggerSecretSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsTriggerSecretSync`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsTriggerSecretSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsTriggerSecretSyncRequest struct via the builder pattern


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


## KmsUpdateSecretSync

> KmsCreateSecretSync200Response KmsUpdateSecretSync(ctx, syncId).KmsUpdateSecretSyncRequest(kmsUpdateSecretSyncRequest).Execute()

Update a secret sync

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
	syncId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateSecretSyncRequest := *openapiclient.NewKmsUpdateSecretSyncRequest() // KmsUpdateSecretSyncRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretSyncsAPI.KmsUpdateSecretSync(context.Background(), syncId).KmsUpdateSecretSyncRequest(kmsUpdateSecretSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretSyncsAPI.KmsUpdateSecretSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateSecretSync`: KmsCreateSecretSync200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretSyncsAPI.KmsUpdateSecretSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateSecretSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateSecretSyncRequest** | [**KmsUpdateSecretSyncRequest**](KmsUpdateSecretSyncRequest.md) |  | 

### Return type

[**KmsCreateSecretSync200Response**](KmsCreateSecretSync200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

