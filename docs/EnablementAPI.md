# \EnablementAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnablement**](EnablementAPI.md#GetEnablement) | **Get** /v1/enablement | Returns what the caller&#39;s org can actually use: every managed item with its global state, whether it is effective here, whether this org is already opted into its beta, and whether it may still opt in.
[**PostEnablementOptin**](EnablementAPI.md#PostEnablementOptin) | **Post** /v1/enablement/optin | Opts the caller&#39;s OWN org into a beta item.
[**PostEnablementOptout**](EnablementAPI.md#PostEnablementOptout) | **Post** /v1/enablement/optout | Removes the caller&#39;s OWN org from a beta item&#39;s grant list, the reverse of OptIntoBeta and idempotent.



## GetEnablement

> EnablementBoard GetEnablement(ctx).Execute()

Returns what the caller's org can actually use: every managed item with its global state, whether it is effective here, whether this org is already opted into its beta, and whether it may still opt in.



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
	resp, r, err := apiClient.EnablementAPI.GetEnablement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.GetEnablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnablement`: EnablementBoard
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.GetEnablement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnablementRequest struct via the builder pattern


### Return type

[**EnablementBoard**](EnablementBoard.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnablementOptin

> UserEnablementItem PostEnablementOptin(ctx).EnablementOptRef(enablementOptRef).Execute()

Opts the caller's OWN org into a beta item.



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
	enablementOptRef := *openapiclient.NewEnablementOptRef() // EnablementOptRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnablementAPI.PostEnablementOptin(context.Background()).EnablementOptRef(enablementOptRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.PostEnablementOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnablementOptin`: UserEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.PostEnablementOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEnablementOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enablementOptRef** | [**EnablementOptRef**](EnablementOptRef.md) |  | 

### Return type

[**UserEnablementItem**](UserEnablementItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnablementOptout

> UserEnablementItem PostEnablementOptout(ctx).EnablementOptRef(enablementOptRef).Execute()

Removes the caller's OWN org from a beta item's grant list, the reverse of OptIntoBeta and idempotent.



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
	enablementOptRef := *openapiclient.NewEnablementOptRef() // EnablementOptRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnablementAPI.PostEnablementOptout(context.Background()).EnablementOptRef(enablementOptRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.PostEnablementOptout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnablementOptout`: UserEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.PostEnablementOptout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEnablementOptoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enablementOptRef** | [**EnablementOptRef**](EnablementOptRef.md) |  | 

### Return type

[**UserEnablementItem**](UserEnablementItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

