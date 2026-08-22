# \ReferenceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskClearReference**](ReferenceAPI.md#RiskClearReference) | **Delete** /v1/reference/{set} | Removes one of your organisation&#39;s overrides.
[**RiskReference**](ReferenceAPI.md#RiskReference) | **Get** /v1/reference/{set} | Reference describes one set and lists your org&#39;s overrides in it.
[**RiskReferenceSets**](ReferenceAPI.md#RiskReferenceSets) | **Get** /v1/reference | Lists every set this plane publishes, with its version and how fresh it is.
[**RiskRefreshReference**](ReferenceAPI.md#RiskRefreshReference) | **Post** /v1/reference/refresh | Takes a new version of one set.
[**RiskResolveReference**](ReferenceAPI.md#RiskResolveReference) | **Post** /v1/reference/resolve | Looks keys up against the reference plane.
[**RiskSetReference**](ReferenceAPI.md#RiskSetReference) | **Put** /v1/reference/{set} | Writes your organisation&#39;s own allow and deny entries over a set.



## RiskClearReference

> ClearReferenceOut RiskClearReference(ctx, set).Key(key).Execute()

Removes one of your organisation's overrides.



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
	set := "domain" // string | 
	key := "partner.example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskClearReference(context.Background(), set).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskClearReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskClearReference`: ClearReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskClearReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskClearReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** |  | 

### Return type

[**ClearReferenceOut**](ClearReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReference

> ReferenceOut RiskReference(ctx, set).After(after).Limit(limit).Execute()

Reference describes one set and lists your org's overrides in it.



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
	set := "domain" // string | 
	after := "after_example" // string | After pages the override listing: the last key of the previous page. (optional)
	limit := int32(50) // int32 | Limit caps the override listing: default 200, maximum 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskReference(context.Background(), set).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReference`: ReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | After pages the override listing: the last key of the previous page. | 
 **limit** | **int32** | Limit caps the override listing: default 200, maximum 1000. | 

### Return type

[**ReferenceOut**](ReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReferenceSets

> ReferenceSetsOut RiskReferenceSets(ctx).Execute()

Lists every set this plane publishes, with its version and how fresh it is.



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
	resp, r, err := apiClient.ReferenceAPI.RiskReferenceSets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskReferenceSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReferenceSets`: ReferenceSetsOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskReferenceSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskReferenceSetsRequest struct via the builder pattern


### Return type

[**ReferenceSetsOut**](ReferenceSetsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskRefreshReference

> RefreshReferenceOut RiskRefreshReference(ctx).RefreshReferenceIn(refreshReferenceIn).Execute()

Takes a new version of one set.



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
	refreshReferenceIn := *openapiclient.NewRefreshReferenceIn() // RefreshReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskRefreshReference(context.Background()).RefreshReferenceIn(refreshReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskRefreshReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskRefreshReference`: RefreshReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskRefreshReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskRefreshReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshReferenceIn** | [**RefreshReferenceIn**](RefreshReferenceIn.md) |  | 

### Return type

[**RefreshReferenceOut**](RefreshReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskResolveReference

> ResolveReferenceOut RiskResolveReference(ctx).ResolveReferenceIn(resolveReferenceIn).Execute()

Looks keys up against the reference plane.



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
	resolveReferenceIn := *openapiclient.NewResolveReferenceIn() // ResolveReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskResolveReference(context.Background()).ResolveReferenceIn(resolveReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskResolveReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskResolveReference`: ResolveReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskResolveReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resolveReferenceIn** | [**ResolveReferenceIn**](ResolveReferenceIn.md) |  | 

### Return type

[**ResolveReferenceOut**](ResolveReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSetReference

> SetReferenceOut RiskSetReference(ctx, set).SetReferenceIn(setReferenceIn).Execute()

Writes your organisation's own allow and deny entries over a set.



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
	set := "domain" // string | 
	setReferenceIn := *openapiclient.NewSetReferenceIn() // SetReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskSetReference(context.Background(), set).SetReferenceIn(setReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskSetReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSetReference`: SetReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskSetReference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRiskSetReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setReferenceIn** | [**SetReferenceIn**](SetReferenceIn.md) |  | 

### Return type

[**SetReferenceOut**](SetReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

