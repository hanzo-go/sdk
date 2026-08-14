# \AppearanceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppearance**](AppearanceAPI.md#GetAppearance) | **Get** /v1/appearance | Returns the signed-in caller&#39;s own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.
[**PostAppearance**](AppearanceAPI.md#PostAppearance) | **Post** /v1/appearance | Stores the caller&#39;s appearance preference on their IAM account, preserving every other field of the row.



## GetAppearance

> Appearance GetAppearance(ctx).Execute()

Returns the signed-in caller's own appearance preference — text size, density and accent — read from their IAM account so it is the same on every device and every Hanzo surface.



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
	resp, r, err := apiClient.AppearanceAPI.GetAppearance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppearanceAPI.GetAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppearance`: Appearance
	fmt.Fprintf(os.Stdout, "Response from `AppearanceAPI.GetAppearance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppearanceRequest struct via the builder pattern


### Return type

[**Appearance**](Appearance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAppearance

> Appearance PostAppearance(ctx).Appearance(appearance).Execute()

Stores the caller's appearance preference on their IAM account, preserving every other field of the row.



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
	appearance := *openapiclient.NewAppearance() // Appearance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppearanceAPI.PostAppearance(context.Background()).Appearance(appearance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppearanceAPI.PostAppearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAppearance`: Appearance
	fmt.Fprintf(os.Stdout, "Response from `AppearanceAPI.PostAppearance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAppearanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appearance** | [**Appearance**](Appearance.md) |  | 

### Return type

[**Appearance**](Appearance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

