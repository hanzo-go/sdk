# \PrefsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Prefs**](PrefsAPI.md#CloudGetV1Prefs) | **Get** /v1/prefs | GetPrefs returns the signed-in caller&#39;s OWN preference document — the theme, density and pinned nav that follow them across every Hanzo surface.
[**CloudPatchV1Prefs**](PrefsAPI.md#CloudPatchV1Prefs) | **Patch** /v1/prefs | 



## CloudGetV1Prefs

> CloudPrefsView CloudGetV1Prefs(ctx).Execute()

GetPrefs returns the signed-in caller's OWN preference document — the theme, density and pinned nav that follow them across every Hanzo surface.



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
	resp, r, err := apiClient.PrefsAPI.CloudGetV1Prefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefsAPI.CloudGetV1Prefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Prefs`: CloudPrefsView
	fmt.Fprintf(os.Stdout, "Response from `PrefsAPI.CloudGetV1Prefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1PrefsRequest struct via the builder pattern


### Return type

[**CloudPrefsView**](CloudPrefsView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPatchV1Prefs

> CloudPatchV1Prefs(ctx).Execute()



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
	r, err := apiClient.PrefsAPI.CloudPatchV1Prefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefsAPI.CloudPatchV1Prefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPatchV1PrefsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

