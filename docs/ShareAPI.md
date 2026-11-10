# \ShareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShare**](ShareAPI.md#GetShare) | **Get** /v1/share | Returns the tunnel shares the caller&#39;s account currently has open — their org&#39;s, or in a held org their own — across every environment it has enabled.
[**PostShareEnable**](ShareAPI.md#PostShareEnable) | **Post** /v1/share/enable | Provisions the caller&#39;s tunnel account — their org&#39;s, or in a held org, where every self-serve signup lands, their own — and returns the credential the &#x60;hanzo share&#x60; CLI needs to run a tunnel.
[**PostShareRotate**](ShareAPI.md#PostShareRotate) | **Post** /v1/share/rotate | Replaces the credential of a tunnel account and ends every environment enabled with it, shares and all, so a credential handed out before — to a member who has left, or to anyone who read it — opens nothing from now on.



## GetShare

> ShareSharesOut GetShare(ctx).Execute()

Returns the tunnel shares the caller's account currently has open — their org's, or in a held org their own — across every environment it has enabled.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.GetShare(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.GetShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShare`: ShareSharesOut
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.GetShare`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareRequest struct via the builder pattern


### Return type

[**ShareSharesOut**](ShareSharesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostShareEnable

> ShareEnableResp PostShareEnable(ctx).Execute()

Provisions the caller's tunnel account — their org's, or in a held org, where every self-serve signup lands, their own — and returns the credential the `hanzo share` CLI needs to run a tunnel.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.PostShareEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.PostShareEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShareEnable`: ShareEnableResp
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.PostShareEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostShareEnableRequest struct via the builder pattern


### Return type

[**ShareEnableResp**](ShareEnableResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostShareRotate

> ShareRotated PostShareRotate(ctx).Execute()

Replaces the credential of a tunnel account and ends every environment enabled with it, shares and all, so a credential handed out before — to a member who has left, or to anyone who read it — opens nothing from now on.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/hanzoai/go-sdk/v8"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAPI.PostShareRotate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.PostShareRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShareRotate`: ShareRotated
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.PostShareRotate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostShareRotateRequest struct via the builder pattern


### Return type

[**ShareRotated**](ShareRotated.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

