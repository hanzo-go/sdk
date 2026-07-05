# \KbConnectorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KbConnectCallback**](KbConnectorsAPI.md#KbConnectCallback) | **Get** /v1/kb/connectors/{provider}/callback | OAuth callback (provider redirect — org recovered from signed state)
[**KbConnectStart**](KbConnectorsAPI.md#KbConnectStart) | **Get** /v1/kb/connectors/{provider}/connect | Begin an OAuth connection (returns the provider authorize URL)
[**KbDisconnectConnector**](KbConnectorsAPI.md#KbDisconnectConnector) | **Delete** /v1/kb/connectors/{provider} | Disconnect a connector (tombstone token, purge its vector points)
[**KbListCatalog**](KbConnectorsAPI.md#KbListCatalog) | **Get** /v1/kb/connectors/catalog | List every connectable source (native + long-tail pieces)
[**KbListConnectors**](KbConnectorsAPI.md#KbListConnectors) | **Get** /v1/kb/connectors | List this org&#39;s connectors and connection state
[**KbSyncConnector**](KbConnectorsAPI.md#KbSyncConnector) | **Post** /v1/kb/connectors/{provider}/sync | Sync the provider&#39;s documents into this org&#39;s knowledge store



## KbConnectCallback

> KbConnectCallback200Response KbConnectCallback(ctx, provider).Code(code).State(state).Error_(error_).Execute()

OAuth callback (provider redirect — org recovered from signed state)

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
	provider := "provider_example" // string | 
	code := "code_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	error_ := "error__example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbConnectorsAPI.KbConnectCallback(context.Background(), provider).Code(code).State(state).Error_(error_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbConnectCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbConnectCallback`: KbConnectCallback200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbConnectCallback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbConnectCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **code** | **string** |  | 
 **state** | **string** |  | 
 **error_** | **string** |  | 

### Return type

[**KbConnectCallback200Response**](KbConnectCallback200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbConnectStart

> KbConnectStart200Response KbConnectStart(ctx, provider).Execute()

Begin an OAuth connection (returns the provider authorize URL)

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbConnectorsAPI.KbConnectStart(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbConnectStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbConnectStart`: KbConnectStart200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbConnectStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbConnectStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KbConnectStart200Response**](KbConnectStart200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbDisconnectConnector

> KbDisconnectConnector200Response KbDisconnectConnector(ctx, provider).Execute()

Disconnect a connector (tombstone token, purge its vector points)

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbConnectorsAPI.KbDisconnectConnector(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbDisconnectConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbDisconnectConnector`: KbDisconnectConnector200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbDisconnectConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbDisconnectConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KbDisconnectConnector200Response**](KbDisconnectConnector200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbListCatalog

> KbListCatalog200Response KbListCatalog(ctx).Execute()

List every connectable source (native + long-tail pieces)

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
	resp, r, err := apiClient.KbConnectorsAPI.KbListCatalog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbListCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbListCatalog`: KbListCatalog200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbListCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKbListCatalogRequest struct via the builder pattern


### Return type

[**KbListCatalog200Response**](KbListCatalog200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbListConnectors

> KbListConnectors200Response KbListConnectors(ctx).Execute()

List this org's connectors and connection state

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
	resp, r, err := apiClient.KbConnectorsAPI.KbListConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbListConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbListConnectors`: KbListConnectors200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbListConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKbListConnectorsRequest struct via the builder pattern


### Return type

[**KbListConnectors200Response**](KbListConnectors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KbSyncConnector

> KbSyncConnector200Response KbSyncConnector(ctx, provider).Execute()

Sync the provider's documents into this org's knowledge store

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
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KbConnectorsAPI.KbSyncConnector(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KbConnectorsAPI.KbSyncConnector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KbSyncConnector`: KbSyncConnector200Response
	fmt.Fprintf(os.Stdout, "Response from `KbConnectorsAPI.KbSyncConnector`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKbSyncConnectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KbSyncConnector200Response**](KbSyncConnector200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

