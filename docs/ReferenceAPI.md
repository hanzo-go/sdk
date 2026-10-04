# \ReferenceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RiskClearReference**](ReferenceAPI.md#RiskClearReference) | **Delete** /v1/reference/{set} | Removes one of your organisation&#39;s overrides.
[**RiskReference**](ReferenceAPI.md#RiskReference) | **Get** /v1/reference/{set} | Reference describes one set and lists your org&#39;s overrides in it.
[**RiskReferenceSets**](ReferenceAPI.md#RiskReferenceSets) | **Get** /v1/reference | Lists every set this plane publishes, with its version and how fresh it is.
[**RiskResolveReference**](ReferenceAPI.md#RiskResolveReference) | **Post** /v1/reference/resolve | Looks keys up against the reference plane.
[**RiskSetReference**](ReferenceAPI.md#RiskSetReference) | **Put** /v1/reference/{set} | Writes your organisation&#39;s own allow and deny entries over a set.



## RiskClearReference

> ReferenceClearReferenceOut RiskClearReference(ctx, set).Key(key).Execute()

Removes one of your organisation's overrides.



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
	set := "domain" // string | 
	key := "partner.example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskClearReference(context.Background(), set).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskClearReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskClearReference`: ReferenceClearReferenceOut
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

[**ReferenceClearReferenceOut**](ReferenceClearReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReference

> ReferenceReferenceOut RiskReference(ctx, set).After(after).Limit(limit).Execute()

Reference describes one set and lists your org's overrides in it.



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
	set := "domain" // string | 
	after := "after_example" // string | After pages the override listing: the last key of the previous page. (optional)
	limit := int64(50) // int64 | Limit caps the override listing: default 200, maximum 1000. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskReference(context.Background(), set).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReference`: ReferenceReferenceOut
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
 **limit** | **int64** | Limit caps the override listing: default 200, maximum 1000. | 

### Return type

[**ReferenceReferenceOut**](ReferenceReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskReferenceSets

> ReferenceReferenceSetsOut RiskReferenceSets(ctx).Execute()

Lists every set this plane publishes, with its version and how fresh it is.



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
	resp, r, err := apiClient.ReferenceAPI.RiskReferenceSets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskReferenceSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskReferenceSets`: ReferenceReferenceSetsOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskReferenceSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRiskReferenceSetsRequest struct via the builder pattern


### Return type

[**ReferenceReferenceSetsOut**](ReferenceReferenceSetsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskResolveReference

> ReferenceResolveReferenceOut RiskResolveReference(ctx).ReferenceResolveReferenceIn(referenceResolveReferenceIn).Execute()

Looks keys up against the reference plane.



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
	referenceResolveReferenceIn := *openapiclient.NewReferenceResolveReferenceIn() // ReferenceResolveReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskResolveReference(context.Background()).ReferenceResolveReferenceIn(referenceResolveReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskResolveReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskResolveReference`: ReferenceResolveReferenceOut
	fmt.Fprintf(os.Stdout, "Response from `ReferenceAPI.RiskResolveReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRiskResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referenceResolveReferenceIn** | [**ReferenceResolveReferenceIn**](ReferenceResolveReferenceIn.md) |  | 

### Return type

[**ReferenceResolveReferenceOut**](ReferenceResolveReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RiskSetReference

> ReferenceSetReferenceOut RiskSetReference(ctx, set).ReferenceSetReferenceIn(referenceSetReferenceIn).Execute()

Writes your organisation's own allow and deny entries over a set.



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
	set := "domain" // string | 
	referenceSetReferenceIn := *openapiclient.NewReferenceSetReferenceIn() // ReferenceSetReferenceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceAPI.RiskSetReference(context.Background(), set).ReferenceSetReferenceIn(referenceSetReferenceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceAPI.RiskSetReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RiskSetReference`: ReferenceSetReferenceOut
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

 **referenceSetReferenceIn** | [**ReferenceSetReferenceIn**](ReferenceSetReferenceIn.md) |  | 

### Return type

[**ReferenceSetReferenceOut**](ReferenceSetReferenceOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

