# \PrefsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrefs**](PrefsAPI.md#GetPrefs) | **Get** /v1/prefs | Returns the signed-in caller&#39;s OWN preference document — the theme, density and pinned nav that follow them across every Hanzo surface.
[**PatchPrefs**](PrefsAPI.md#PatchPrefs) | **Patch** /v1/prefs | Save the preference keys your surface owns, leaving every other key alone



## GetPrefs

> PrefsView GetPrefs(ctx).Execute()

Returns the signed-in caller's OWN preference document — the theme, density and pinned nav that follow them across every Hanzo surface.



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
	resp, r, err := apiClient.PrefsAPI.GetPrefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefsAPI.GetPrefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrefs`: PrefsView
	fmt.Fprintf(os.Stdout, "Response from `PrefsAPI.GetPrefs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrefsRequest struct via the builder pattern


### Return type

[**PrefsView**](PrefsView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPrefs

> PatchPrefs(ctx).Execute()

Save the preference keys your surface owns, leaving every other key alone



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
	r, err := apiClient.PrefsAPI.PatchPrefs(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrefsAPI.PatchPrefs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPrefsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

