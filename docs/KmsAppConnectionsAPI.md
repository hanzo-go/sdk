# \KmsAppConnectionsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsCreateAppConnection**](KmsAppConnectionsAPI.md#KmsCreateAppConnection) | **Post** /v1/kms/app-connections | Create an app connection
[**KmsDeleteAppConnection**](KmsAppConnectionsAPI.md#KmsDeleteAppConnection) | **Delete** /v1/kms/app-connections/{connectionId} | Delete an app connection
[**KmsGetAppConnection**](KmsAppConnectionsAPI.md#KmsGetAppConnection) | **Get** /v1/kms/app-connections/{connectionId} | Get an app connection by ID
[**KmsListAppConnections**](KmsAppConnectionsAPI.md#KmsListAppConnections) | **Get** /v1/kms/app-connections | List app connections
[**KmsUpdateAppConnection**](KmsAppConnectionsAPI.md#KmsUpdateAppConnection) | **Patch** /v1/kms/app-connections/{connectionId} | Update an app connection



## KmsCreateAppConnection

> KmsCreateAppConnection200Response KmsCreateAppConnection(ctx).KmsCreateAppConnectionRequest(kmsCreateAppConnectionRequest).Execute()

Create an app connection

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
	kmsCreateAppConnectionRequest := *openapiclient.NewKmsCreateAppConnectionRequest("Name_example", "App_example", map[string]interface{}(123)) // KmsCreateAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAppConnectionsAPI.KmsCreateAppConnection(context.Background()).KmsCreateAppConnectionRequest(kmsCreateAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAppConnectionsAPI.KmsCreateAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateAppConnection`: KmsCreateAppConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAppConnectionsAPI.KmsCreateAppConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kmsCreateAppConnectionRequest** | [**KmsCreateAppConnectionRequest**](KmsCreateAppConnectionRequest.md) |  | 

### Return type

[**KmsCreateAppConnection200Response**](KmsCreateAppConnection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteAppConnection

> map[string]interface{} KmsDeleteAppConnection(ctx, connectionId).Execute()

Delete an app connection

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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAppConnectionsAPI.KmsDeleteAppConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAppConnectionsAPI.KmsDeleteAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteAppConnection`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KmsAppConnectionsAPI.KmsDeleteAppConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteAppConnectionRequest struct via the builder pattern


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


## KmsGetAppConnection

> KmsCreateAppConnection200Response KmsGetAppConnection(ctx, connectionId).Execute()

Get an app connection by ID

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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAppConnectionsAPI.KmsGetAppConnection(context.Background(), connectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAppConnectionsAPI.KmsGetAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetAppConnection`: KmsCreateAppConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAppConnectionsAPI.KmsGetAppConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KmsCreateAppConnection200Response**](KmsCreateAppConnection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListAppConnections

> KmsListAppConnections200Response KmsListAppConnections(ctx).OrgId(orgId).Execute()

List app connections

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
	orgId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAppConnectionsAPI.KmsListAppConnections(context.Background()).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAppConnectionsAPI.KmsListAppConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListAppConnections`: KmsListAppConnections200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAppConnectionsAPI.KmsListAppConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListAppConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **string** |  | 

### Return type

[**KmsListAppConnections200Response**](KmsListAppConnections200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateAppConnection

> KmsCreateAppConnection200Response KmsUpdateAppConnection(ctx, connectionId).KmsUpdateAppConnectionRequest(kmsUpdateAppConnectionRequest).Execute()

Update an app connection

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
	connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsUpdateAppConnectionRequest := *openapiclient.NewKmsUpdateAppConnectionRequest() // KmsUpdateAppConnectionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAppConnectionsAPI.KmsUpdateAppConnection(context.Background(), connectionId).KmsUpdateAppConnectionRequest(kmsUpdateAppConnectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAppConnectionsAPI.KmsUpdateAppConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateAppConnection`: KmsCreateAppConnection200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAppConnectionsAPI.KmsUpdateAppConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateAppConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateAppConnectionRequest** | [**KmsUpdateAppConnectionRequest**](KmsUpdateAppConnectionRequest.md) |  | 

### Return type

[**KmsCreateAppConnection200Response**](KmsCreateAppConnection200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

