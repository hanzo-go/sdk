# \EnablementAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Enablement**](EnablementAPI.md#CloudGetV1Enablement) | **Get** /v1/enablement | GetEnablement returns what the caller&#39;s org can actually use: every managed item with its global state, whether it is effective here, whether this org is already opted into its beta, and whether it may still opt in.
[**CloudPostV1EnablementOptin**](EnablementAPI.md#CloudPostV1EnablementOptin) | **Post** /v1/enablement/optin | OptIntoBeta opts the caller&#39;s OWN org into a beta item.
[**CloudPostV1EnablementOptout**](EnablementAPI.md#CloudPostV1EnablementOptout) | **Post** /v1/enablement/optout | OptOutOfBeta removes the caller&#39;s OWN org from a beta item&#39;s grant list, the reverse of OptIntoBeta and idempotent.



## CloudGetV1Enablement

> CloudEnablementBoard CloudGetV1Enablement(ctx).Execute()

GetEnablement returns what the caller's org can actually use: every managed item with its global state, whether it is effective here, whether this org is already opted into its beta, and whether it may still opt in.



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
	resp, r, err := apiClient.EnablementAPI.CloudGetV1Enablement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.CloudGetV1Enablement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Enablement`: CloudEnablementBoard
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.CloudGetV1Enablement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1EnablementRequest struct via the builder pattern


### Return type

[**CloudEnablementBoard**](CloudEnablementBoard.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1EnablementOptin

> CloudUserEnablementItem CloudPostV1EnablementOptin(ctx).CloudEnablementOptRef(cloudEnablementOptRef).Execute()

OptIntoBeta opts the caller's OWN org into a beta item.



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
	cloudEnablementOptRef := *openapiclient.NewCloudEnablementOptRef() // CloudEnablementOptRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnablementAPI.CloudPostV1EnablementOptin(context.Background()).CloudEnablementOptRef(cloudEnablementOptRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.CloudPostV1EnablementOptin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1EnablementOptin`: CloudUserEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.CloudPostV1EnablementOptin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1EnablementOptinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEnablementOptRef** | [**CloudEnablementOptRef**](CloudEnablementOptRef.md) |  | 

### Return type

[**CloudUserEnablementItem**](CloudUserEnablementItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1EnablementOptout

> CloudUserEnablementItem CloudPostV1EnablementOptout(ctx).CloudEnablementOptRef(cloudEnablementOptRef).Execute()

OptOutOfBeta removes the caller's OWN org from a beta item's grant list, the reverse of OptIntoBeta and idempotent.



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
	cloudEnablementOptRef := *openapiclient.NewCloudEnablementOptRef() // CloudEnablementOptRef | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnablementAPI.CloudPostV1EnablementOptout(context.Background()).CloudEnablementOptRef(cloudEnablementOptRef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnablementAPI.CloudPostV1EnablementOptout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1EnablementOptout`: CloudUserEnablementItem
	fmt.Fprintf(os.Stdout, "Response from `EnablementAPI.CloudPostV1EnablementOptout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1EnablementOptoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEnablementOptRef** | [**CloudEnablementOptRef**](CloudEnablementOptRef.md) |  | 

### Return type

[**CloudUserEnablementItem**](CloudUserEnablementItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

