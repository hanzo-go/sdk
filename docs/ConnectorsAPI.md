# \ConnectorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnectorsById**](ConnectorsAPI.md#DeleteConnectorsById) | **Delete** /v1/connectors/{id} | Forgets a connector: every custodied secret, then the row.
[**GetConnectors**](ConnectorsAPI.md#GetConnectors) | **Get** /v1/connectors | Lists the caller&#39;s OWN connectors across every provider — the set &#x60;hanzo connector ls&#x60; prints.
[**GetConnectorsByIdToken**](ConnectorsAPI.md#GetConnectorsByIdToken) | **Get** /v1/connectors/{id}/token | Hands the custodied access token to its owner — the ONE place custody exits.
[**GetConnectorsProviders**](ConnectorsAPI.md#GetConnectorsProviders) | **Get** /v1/connectors/providers | Lists the user-scoped provider cards — the catalog of what a user can connect, and how.
[**PostConnectorsByIdRefresh**](ConnectorsAPI.md#PostConnectorsByIdRefresh) | **Post** /v1/connectors/{id}/refresh | Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.
[**PostConnectorsByProviderCredential**](ConnectorsAPI.md#PostConnectorsByProviderCredential) | **Post** /v1/connectors/{provider}/credential | Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI&#39;s local PKCE (Adopt).
[**PostConnectorsByProviderDevice**](ConnectorsAPI.md#PostConnectorsByProviderDevice) | **Post** /v1/connectors/{provider}/device | Begins a device sign-in and returns the code to show the user plus how to poll for completion.
[**PostConnectorsByProviderDeviceByFlowPoll**](ConnectorsAPI.md#PostConnectorsByProviderDeviceByFlowPoll) | **Post** /v1/connectors/{provider}/device/{flow}/poll | Advances a device sign-in.



## DeleteConnectorsById

> DisconnectOut DeleteConnectorsById(ctx, id).Execute()

Forgets a connector: every custodied secret, then the row.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.DeleteConnectorsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.DeleteConnectorsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConnectorsById`: DisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.DeleteConnectorsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DisconnectOut**](DisconnectOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectors

> ConnectorsOut GetConnectors(ctx).Execute()

Lists the caller's OWN connectors across every provider — the set `hanzo connector ls` prints.



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
	resp, r, err := apiClient.ConnectorsAPI.GetConnectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectors`: ConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorsRequest struct via the builder pattern


### Return type

[**ConnectorsOut**](ConnectorsOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorsByIdToken

> ConnectorTokenOut GetConnectorsByIdToken(ctx, id).Execute()

Hands the custodied access token to its owner — the ONE place custody exits.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorsByIdToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorsByIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorsByIdToken`: ConnectorTokenOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorsByIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorsByIdTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorTokenOut**](ConnectorTokenOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorsProviders

> ConnectorProvidersOut GetConnectorsProviders(ctx).Execute()

Lists the user-scoped provider cards — the catalog of what a user can connect, and how.



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
	resp, r, err := apiClient.ConnectorsAPI.GetConnectorsProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.GetConnectorsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorsProviders`: ConnectorProvidersOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.GetConnectorsProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorsProvidersRequest struct via the builder pattern


### Return type

[**ConnectorProvidersOut**](ConnectorProvidersOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectorsByIdRefresh

> RefreshOut PostConnectorsByIdRefresh(ctx, id).Execute()

Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.



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
	id := "openai:work" // string | ID is the connector id, provider + \":\" + label (\"openai:default\") — the auth-profile-id shape. Another user's id is simply no row, so 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PostConnectorsByIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PostConnectorsByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectorsByIdRefresh`: RefreshOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PostConnectorsByIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectorsByIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RefreshOut**](RefreshOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectorsByProviderCredential

> CredentialOut PostConnectorsByProviderCredential(ctx, provider).CredentialIn(credentialIn).Execute()

Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI's local PKCE (Adopt).



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	credentialIn := *openapiclient.NewCredentialIn() // CredentialIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PostConnectorsByProviderCredential(context.Background(), provider).CredentialIn(credentialIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PostConnectorsByProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectorsByProviderCredential`: CredentialOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PostConnectorsByProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectorsByProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialIn** | [**CredentialIn**](CredentialIn.md) |  | 

### Return type

[**CredentialOut**](CredentialOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectorsByProviderDevice

> DeviceStartOut PostConnectorsByProviderDevice(ctx, provider).DeviceStartIn(deviceStartIn).Execute()

Begins a device sign-in and returns the code to show the user plus how to poll for completion.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	deviceStartIn := *openapiclient.NewDeviceStartIn() // DeviceStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PostConnectorsByProviderDevice(context.Background(), provider).DeviceStartIn(deviceStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PostConnectorsByProviderDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectorsByProviderDevice`: DeviceStartOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PostConnectorsByProviderDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectorsByProviderDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceStartIn** | [**DeviceStartIn**](DeviceStartIn.md) |  | 

### Return type

[**DeviceStartOut**](DeviceStartOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostConnectorsByProviderDeviceByFlowPoll

> DevicePollOut PostConnectorsByProviderDeviceByFlowPoll(ctx, provider, flow).Execute()

Advances a device sign-in.



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
	provider := "openai" // string | Provider is the user-scoped provider's registry id, from the path.
	flow := "g_7f2c" // string | Flow is the id deviceStartOut returned. Expired or another user's flow is indistinguishable from an unknown one: 404.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.PostConnectorsByProviderDeviceByFlowPoll(context.Background(), provider, flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.PostConnectorsByProviderDeviceByFlowPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConnectorsByProviderDeviceByFlowPoll`: DevicePollOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.PostConnectorsByProviderDeviceByFlowPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 
**flow** | **string** | Flow is the id deviceStartOut returned. Expired or another user&#39;s flow is indistinguishable from an unknown one: 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostConnectorsByProviderDeviceByFlowPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicePollOut**](DevicePollOut.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

