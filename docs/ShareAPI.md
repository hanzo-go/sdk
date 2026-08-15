# \ShareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShare**](ShareAPI.md#GetShare) | **Get** /v1/share | Returns the tunnel shares the caller&#39;s org currently has open, across every environment that org has enabled.
[**PostShareEnable**](ShareAPI.md#PostShareEnable) | **Post** /v1/share/enable | Enable provisions the caller org&#39;s tunnel account and returns the credential the &#x60;hanzo share&#x60; CLI needs to run a tunnel.



## GetShare

> SharesOut GetShare(ctx).Execute()

Returns the tunnel shares the caller's org currently has open, across every environment that org has enabled.



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
	resp, r, err := apiClient.ShareAPI.GetShare(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.GetShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShare`: SharesOut
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.GetShare`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareRequest struct via the builder pattern


### Return type

[**SharesOut**](SharesOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostShareEnable

> EnableResp PostShareEnable(ctx).Execute()

Enable provisions the caller org's tunnel account and returns the credential the `hanzo share` CLI needs to run a tunnel.



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
	resp, r, err := apiClient.ShareAPI.PostShareEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAPI.PostShareEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShareEnable`: EnableResp
	fmt.Fprintf(os.Stdout, "Response from `ShareAPI.PostShareEnable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostShareEnableRequest struct via the builder pattern


### Return type

[**EnableResp**](EnableResp.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

