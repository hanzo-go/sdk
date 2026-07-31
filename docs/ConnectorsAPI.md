# \ConnectorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1ConnectorsId**](ConnectorsAPI.md#CloudDeleteV1ConnectorsId) | **Delete** /v1/connectors/{id} | Forgets a connector: every custodied secret, then the row.
[**CloudGetV1Connectors**](ConnectorsAPI.md#CloudGetV1Connectors) | **Get** /v1/connectors | Lists the caller&#39;s OWN connectors across every provider — the set &#x60;hanzo connector ls&#x60; prints.
[**CloudGetV1ConnectorsIdToken**](ConnectorsAPI.md#CloudGetV1ConnectorsIdToken) | **Get** /v1/connectors/{id}/token | Hands the custodied access token to its owner — the ONE place custody exits.
[**CloudGetV1ConnectorsProviders**](ConnectorsAPI.md#CloudGetV1ConnectorsProviders) | **Get** /v1/connectors/providers | Lists the user-scoped provider cards — the catalog of what a user can connect, and how.
[**CloudPostV1ConnectorsIdRefresh**](ConnectorsAPI.md#CloudPostV1ConnectorsIdRefresh) | **Post** /v1/connectors/{id}/refresh | Forces a token rotation for a connected connector, ahead of the automatic rotation a token read would do inside the expiry window.
[**CloudPostV1ConnectorsProviderCredential**](ConnectorsAPI.md#CloudPostV1ConnectorsProviderCredential) | **Post** /v1/connectors/{provider}/credential | Is the direct intake path: a customer-held token/setup-token (Verify) or an externally obtained OAuth bundle from the CLI&#39;s local PKCE (Adopt).
[**CloudPostV1ConnectorsProviderDevice**](ConnectorsAPI.md#CloudPostV1ConnectorsProviderDevice) | **Post** /v1/connectors/{provider}/device | Begins a device sign-in and returns the code to show the user plus how to poll for completion.
[**CloudPostV1ConnectorsProviderDeviceFlowPoll**](ConnectorsAPI.md#CloudPostV1ConnectorsProviderDeviceFlowPoll) | **Post** /v1/connectors/{provider}/device/{flow}/poll | Advances a device sign-in.



## CloudDeleteV1ConnectorsId

> CloudDisconnectOut CloudDeleteV1ConnectorsId(ctx, id).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudDeleteV1ConnectorsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudDeleteV1ConnectorsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1ConnectorsId`: CloudDisconnectOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudDeleteV1ConnectorsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1ConnectorsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDisconnectOut**](CloudDisconnectOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Connectors

> CloudConnectorsOut CloudGetV1Connectors(ctx).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudGetV1Connectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudGetV1Connectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Connectors`: CloudConnectorsOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudGetV1Connectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ConnectorsRequest struct via the builder pattern


### Return type

[**CloudConnectorsOut**](CloudConnectorsOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ConnectorsIdToken

> CloudConnectorTokenOut CloudGetV1ConnectorsIdToken(ctx, id).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudGetV1ConnectorsIdToken(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudGetV1ConnectorsIdToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ConnectorsIdToken`: CloudConnectorTokenOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudGetV1ConnectorsIdToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ConnectorsIdTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudConnectorTokenOut**](CloudConnectorTokenOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ConnectorsProviders

> CloudConnectorProvidersOut CloudGetV1ConnectorsProviders(ctx).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudGetV1ConnectorsProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudGetV1ConnectorsProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ConnectorsProviders`: CloudConnectorProvidersOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudGetV1ConnectorsProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ConnectorsProvidersRequest struct via the builder pattern


### Return type

[**CloudConnectorProvidersOut**](CloudConnectorProvidersOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ConnectorsIdRefresh

> CloudRefreshOut CloudPostV1ConnectorsIdRefresh(ctx, id).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudPostV1ConnectorsIdRefresh(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudPostV1ConnectorsIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ConnectorsIdRefresh`: CloudRefreshOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudPostV1ConnectorsIdRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the connector id, provider + \&quot;:\&quot; + label (\&quot;openai:default\&quot;) — the auth-profile-id shape. Another user&#39;s id is simply no row, so 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ConnectorsIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRefreshOut**](CloudRefreshOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ConnectorsProviderCredential

> CloudCredentialOut CloudPostV1ConnectorsProviderCredential(ctx, provider).CloudCredentialIn(cloudCredentialIn).Execute()

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
	cloudCredentialIn := *openapiclient.NewCloudCredentialIn() // CloudCredentialIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CloudPostV1ConnectorsProviderCredential(context.Background(), provider).CloudCredentialIn(cloudCredentialIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudPostV1ConnectorsProviderCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ConnectorsProviderCredential`: CloudCredentialOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudPostV1ConnectorsProviderCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ConnectorsProviderCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudCredentialIn** | [**CloudCredentialIn**](CloudCredentialIn.md) |  | 

### Return type

[**CloudCredentialOut**](CloudCredentialOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ConnectorsProviderDevice

> CloudDeviceStartOut CloudPostV1ConnectorsProviderDevice(ctx, provider).CloudDeviceStartIn(cloudDeviceStartIn).Execute()

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
	cloudDeviceStartIn := *openapiclient.NewCloudDeviceStartIn() // CloudDeviceStartIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorsAPI.CloudPostV1ConnectorsProviderDevice(context.Background(), provider).CloudDeviceStartIn(cloudDeviceStartIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudPostV1ConnectorsProviderDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ConnectorsProviderDevice`: CloudDeviceStartOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudPostV1ConnectorsProviderDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ConnectorsProviderDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudDeviceStartIn** | [**CloudDeviceStartIn**](CloudDeviceStartIn.md) |  | 

### Return type

[**CloudDeviceStartOut**](CloudDeviceStartOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ConnectorsProviderDeviceFlowPoll

> CloudDevicePollOut CloudPostV1ConnectorsProviderDeviceFlowPoll(ctx, provider, flow).Execute()

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
	resp, r, err := apiClient.ConnectorsAPI.CloudPostV1ConnectorsProviderDeviceFlowPoll(context.Background(), provider, flow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorsAPI.CloudPostV1ConnectorsProviderDeviceFlowPoll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ConnectorsProviderDeviceFlowPoll`: CloudDevicePollOut
	fmt.Fprintf(os.Stdout, "Response from `ConnectorsAPI.CloudPostV1ConnectorsProviderDeviceFlowPoll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | 
**flow** | **string** | Flow is the id deviceStartOut returned. Expired or another user&#39;s flow is indistinguishable from an unknown one: 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ConnectorsProviderDeviceFlowPollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CloudDevicePollOut**](CloudDevicePollOut.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

