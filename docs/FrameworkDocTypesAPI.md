# \FrameworkDocTypesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FrameworkCreateDocType**](FrameworkDocTypesAPI.md#FrameworkCreateDocType) | **Post** /v1/framework/doctypes | Define a DocType
[**FrameworkDeleteDocType**](FrameworkDocTypesAPI.md#FrameworkDeleteDocType) | **Delete** /v1/framework/doctypes/{name} | Delete a DocType (and its documents)
[**FrameworkFrameworkSummary**](FrameworkDocTypesAPI.md#FrameworkFrameworkSummary) | **Get** /v1/framework/summary | Org summary (doctype + document counts)
[**FrameworkGetDocType**](FrameworkDocTypesAPI.md#FrameworkGetDocType) | **Get** /v1/framework/doctypes/{name} | Get a DocType definition
[**FrameworkListDocTypes**](FrameworkDocTypesAPI.md#FrameworkListDocTypes) | **Get** /v1/framework/doctypes | List DocType definitions
[**FrameworkReplaceDocType**](FrameworkDocTypesAPI.md#FrameworkReplaceDocType) | **Put** /v1/framework/doctypes/{name} | Replace a DocType definition



## FrameworkCreateDocType

> FrameworkDocType FrameworkCreateDocType(ctx).FrameworkDocType(frameworkDocType).Execute()

Define a DocType

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
	frameworkDocType := *openapiclient.NewFrameworkDocType("Name_example", []openapiclient.FrameworkDocField{*openapiclient.NewFrameworkDocField("Fieldname_example", openapiclient.framework_Fieldtype("Data"))}) // FrameworkDocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocTypesAPI.FrameworkCreateDocType(context.Background()).FrameworkDocType(frameworkDocType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkCreateDocType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkCreateDocType`: FrameworkDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocTypesAPI.FrameworkCreateDocType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkCreateDocTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **frameworkDocType** | [**FrameworkDocType**](FrameworkDocType.md) |  | 

### Return type

[**FrameworkDocType**](FrameworkDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkDeleteDocType

> FrameworkDeleteDocType(ctx, name).Execute()

Delete a DocType (and its documents)

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkDocTypesAPI.FrameworkDeleteDocType(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkDeleteDocType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkDeleteDocTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkFrameworkSummary

> FrameworkFrameworkSummary200Response FrameworkFrameworkSummary(ctx).Execute()

Org summary (doctype + document counts)

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
	resp, r, err := apiClient.FrameworkDocTypesAPI.FrameworkFrameworkSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkFrameworkSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkFrameworkSummary`: FrameworkFrameworkSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocTypesAPI.FrameworkFrameworkSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkFrameworkSummaryRequest struct via the builder pattern


### Return type

[**FrameworkFrameworkSummary200Response**](FrameworkFrameworkSummary200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkGetDocType

> FrameworkDocType FrameworkGetDocType(ctx, name).Execute()

Get a DocType definition

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
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocTypesAPI.FrameworkGetDocType(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkGetDocType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkGetDocType`: FrameworkDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocTypesAPI.FrameworkGetDocType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkGetDocTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FrameworkDocType**](FrameworkDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkListDocTypes

> FrameworkListDocTypes200Response FrameworkListDocTypes(ctx).Execute()

List DocType definitions

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
	resp, r, err := apiClient.FrameworkDocTypesAPI.FrameworkListDocTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkListDocTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkListDocTypes`: FrameworkListDocTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocTypesAPI.FrameworkListDocTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkListDocTypesRequest struct via the builder pattern


### Return type

[**FrameworkListDocTypes200Response**](FrameworkListDocTypes200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FrameworkReplaceDocType

> FrameworkDocType FrameworkReplaceDocType(ctx, name).FrameworkDocType(frameworkDocType).Execute()

Replace a DocType definition

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
	name := "name_example" // string | 
	frameworkDocType := *openapiclient.NewFrameworkDocType("Name_example", []openapiclient.FrameworkDocField{*openapiclient.NewFrameworkDocField("Fieldname_example", openapiclient.framework_Fieldtype("Data"))}) // FrameworkDocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkDocTypesAPI.FrameworkReplaceDocType(context.Background(), name).FrameworkDocType(frameworkDocType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkDocTypesAPI.FrameworkReplaceDocType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameworkReplaceDocType`: FrameworkDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkDocTypesAPI.FrameworkReplaceDocType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFrameworkReplaceDocTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **frameworkDocType** | [**FrameworkDocType**](FrameworkDocType.md) |  | 

### Return type

[**FrameworkDocType**](FrameworkDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

