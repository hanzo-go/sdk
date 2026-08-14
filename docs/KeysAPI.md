# \KeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteKeys**](KeysAPI.md#DeleteKeys) | **Delete** /v1/keys | Revokes the caller&#39;s own API key of the requested class.
[**GetKeys**](KeysAPI.md#GetKeys) | **Get** /v1/keys | Returns the caller&#39;s own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.
[**PostKeys**](KeysAPI.md#PostKeys) | **Post** /v1/keys | Creates — or rotates — the caller&#39;s API key of the requested type and returns it ONCE.



## DeleteKeys

> RevokedKey DeleteKeys(ctx).Type_(type_).Execute()

Revokes the caller's own API key of the requested class.



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
	type_ := "publishable" // string | Type is the key class to act on: \"secret\" (sk-, session-equivalent, belongs on a server) or \"publishable\" (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.DeleteKeys(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.DeleteKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteKeys`: RevokedKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.DeleteKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type is the key class to act on: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. | 

### Return type

[**RevokedKey**](RevokedKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeys

> ApiKeyList GetKeys(ctx).Execute()

Returns the caller's own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.



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
	resp, r, err := apiClient.KeysAPI.GetKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GetKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeys`: ApiKeyList
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GetKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysRequest struct via the builder pattern


### Return type

[**ApiKeyList**](ApiKeyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostKeys

> MintedKey PostKeys(ctx).KeyTypeIn(keyTypeIn).Execute()

Creates — or rotates — the caller's API key of the requested type and returns it ONCE.



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
	keyTypeIn := *openapiclient.NewKeyTypeIn() // KeyTypeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.PostKeys(context.Background()).KeyTypeIn(keyTypeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.PostKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostKeys`: MintedKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.PostKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyTypeIn** | [**KeyTypeIn**](KeyTypeIn.md) |  | 

### Return type

[**MintedKey**](MintedKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

