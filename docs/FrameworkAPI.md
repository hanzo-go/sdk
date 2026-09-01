# \FrameworkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFrameworkByDoctypeByName**](FrameworkAPI.md#DeleteFrameworkByDoctypeByName) | **Delete** /v1/framework/{doctype}/{name} | Removes one document, after its on_trash hooks agree.
[**DeleteFrameworkDoctypesByName**](FrameworkAPI.md#DeleteFrameworkDoctypesByName) | **Delete** /v1/framework/doctypes/{name} | Removes a DocType and every document stored under it.
[**GetFrameworkByDoctype**](FrameworkAPI.md#GetFrameworkByDoctype) | **Get** /v1/framework/{doctype} | Returns the caller org&#39;s documents of one DocType, filtered, ordered and projected by the query.
[**GetFrameworkByDoctypeByName**](FrameworkAPI.md#GetFrameworkByDoctypeByName) | **Get** /v1/framework/{doctype}/{name} | Returns one document by name, with Password fields redacted.
[**GetFrameworkDoctypes**](FrameworkAPI.md#GetFrameworkDoctypes) | **Get** /v1/framework/doctypes | Returns every DocType defined in the caller&#39;s org.
[**GetFrameworkDoctypesByName**](FrameworkAPI.md#GetFrameworkDoctypesByName) | **Get** /v1/framework/doctypes/{name} | Returns one DocType definition — its fields, naming rule, permissions and lifecycle flags.
[**GetFrameworkModules**](FrameworkAPI.md#GetFrameworkModules) | **Get** /v1/framework/modules | Returns every app lane compiled into this deployment and the DocTypes each one installs.
[**GetFrameworkModulesByModule**](FrameworkAPI.md#GetFrameworkModulesByModule) | **Get** /v1/framework/modules/{module} | Returns one app lane&#39;s install state for the caller&#39;s org: the DocTypes the lane declares, and which of them already exist in the org.
[**GetFrameworkSummary**](FrameworkAPI.md#GetFrameworkSummary) | **Get** /v1/framework/summary | Reports how much of the DocType surface the caller&#39;s org uses: how many DocTypes it has defined, and how many documents exist across them.
[**PostFrameworkByDoctype**](FrameworkAPI.md#PostFrameworkByDoctype) | **Post** /v1/framework/{doctype} | Create one document of a DocType, from that DocType&#39;s own fields.
[**PostFrameworkByDoctypeByNameCancel**](FrameworkAPI.md#PostFrameworkByDoctypeByNameCancel) | **Post** /v1/framework/{doctype}/{name}/cancel | Moves a submitted document to cancelled (docstatus 1 → 2) after its on_cancel hooks agree.
[**PostFrameworkByDoctypeByNameSubmit**](FrameworkAPI.md#PostFrameworkByDoctypeByNameSubmit) | **Post** /v1/framework/{doctype}/{name}/submit | Moves a draft to submitted (docstatus 0 → 1) after its on_submit hooks agree.
[**PostFrameworkDoctypes**](FrameworkAPI.md#PostFrameworkDoctypes) | **Post** /v1/framework/doctypes | Defines a DocType in the caller&#39;s org: the metadata that gives a document surface its fields, its naming rule, whether it has a submit/cancel lifecycle, and which role may do what to it.
[**PostFrameworkModulesByModuleInstall**](FrameworkAPI.md#PostFrameworkModulesByModuleInstall) | **Post** /v1/framework/modules/{module}/install | Creates an app lane&#39;s DocTypes in the caller&#39;s org.
[**PutFrameworkByDoctypeByName**](FrameworkAPI.md#PutFrameworkByDoctypeByName) | **Put** /v1/framework/{doctype}/{name} | Replace a draft document&#39;s field data wholesale.
[**PutFrameworkDoctypesByName**](FrameworkAPI.md#PutFrameworkDoctypesByName) | **Put** /v1/framework/doctypes/{name} | Replaces a DocType definition wholesale (PUT semantics): the stored definition becomes the body.



## DeleteFrameworkByDoctypeByName

> DeleteFrameworkByDoctypeByName(ctx, doctype, name).Execute()

Removes one document, after its on_trash hooks agree.



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
	doctype := "Projects.Task" // string | DocType is the document's DocType, by ADDRESS — \"module.name\", from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.DeleteFrameworkByDoctypeByName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.DeleteFrameworkByDoctypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, by ADDRESS — \&quot;module.name\&quot;, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFrameworkByDoctypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFrameworkDoctypesByName

> DeleteFrameworkDoctypesByName(ctx, name).Execute()

Removes a DocType and every document stored under it.



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
	name := "Projects.Task" // string | Name is the DocType's ADDRESS — \"module.name\", e.g. \"kb.page\". A name containing a space (\"erp.Sales Invoice\") arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.DeleteFrameworkDoctypesByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.DeleteFrameworkDoctypesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the DocType&#39;s ADDRESS — \&quot;module.name\&quot;, e.g. \&quot;kb.page\&quot;. A name containing a space (\&quot;erp.Sales Invoice\&quot;) arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFrameworkDoctypesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkByDoctype

> DocumentList GetFrameworkByDoctype(ctx, doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()

Returns the caller org's documents of one DocType, filtered, ordered and projected by the query.



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
	doctype := "Projects.Task" // string | DocType is the DocType to list, by ADDRESS — \"module.name\", from the path.
	filters := "{"priority":"High"}" // string | Filters is a JSON object of equality matches, e.g. {\"priority\":\"High\"}. Every key must be a field the DocType declares (or the managed name / docstatus); an undeclared one is refused rather than silently ignored. (optional)
	fields := "fields_example" // string | Fields projects the response to a subset — a JSON array [\"a\",\"b\"] or a comma list \"a,b\". The envelope keys are always returned. (optional)
	orderBy := "estimate asc" // string | OrderBy is \"<field> [asc|desc]\". Empty means most-recently-updated first. (optional)
	limit := "20" // string | Limit caps the rows returned. Anything that is not a positive integer leaves the engine's default in place. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkByDoctype(context.Background(), doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkByDoctype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkByDoctype`: DocumentList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkByDoctype`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the DocType to list, by ADDRESS — \&quot;module.name\&quot;, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkByDoctypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | Filters is a JSON object of equality matches, e.g. {\&quot;priority\&quot;:\&quot;High\&quot;}. Every key must be a field the DocType declares (or the managed name / docstatus); an undeclared one is refused rather than silently ignored. | 
 **fields** | **string** | Fields projects the response to a subset — a JSON array [\&quot;a\&quot;,\&quot;b\&quot;] or a comma list \&quot;a,b\&quot;. The envelope keys are always returned. | 
 **orderBy** | **string** | OrderBy is \&quot;&lt;field&gt; [asc|desc]\&quot;. Empty means most-recently-updated first. | 
 **limit** | **string** | Limit caps the rows returned. Anything that is not a positive integer leaves the engine&#39;s default in place. | 

### Return type

[**DocumentList**](DocumentList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkByDoctypeByName

> map[string]map[string]interface{} GetFrameworkByDoctypeByName(ctx, doctype, name).Execute()

Returns one document by name, with Password fields redacted.



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
	doctype := "Projects.Task" // string | DocType is the document's DocType, by ADDRESS — \"module.name\", from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkByDoctypeByName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkByDoctypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkByDoctypeByName`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkByDoctypeByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, by ADDRESS — \&quot;module.name\&quot;, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkByDoctypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkDoctypes

> DocTypeList GetFrameworkDoctypes(ctx).Execute()

Returns every DocType defined in the caller's org.



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
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkDoctypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkDoctypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkDoctypes`: DocTypeList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkDoctypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkDoctypesRequest struct via the builder pattern


### Return type

[**DocTypeList**](DocTypeList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkDoctypesByName

> DocType GetFrameworkDoctypesByName(ctx, name).Execute()

Returns one DocType definition — its fields, naming rule, permissions and lifecycle flags.



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
	name := "Projects.Task" // string | Name is the DocType's ADDRESS — \"module.name\", e.g. \"kb.page\". A name containing a space (\"erp.Sales Invoice\") arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkDoctypesByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkDoctypesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkDoctypesByName`: DocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkDoctypesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the DocType&#39;s ADDRESS — \&quot;module.name\&quot;, e.g. \&quot;kb.page\&quot;. A name containing a space (\&quot;erp.Sales Invoice\&quot;) arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkDoctypesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DocType**](DocType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkModules

> ModuleList GetFrameworkModules(ctx).Execute()

Returns every app lane compiled into this deployment and the DocTypes each one installs.



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
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkModules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkModules`: ModuleList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkModulesRequest struct via the builder pattern


### Return type

[**ModuleList**](ModuleList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkModulesByModule

> ModuleState GetFrameworkModulesByModule(ctx, module).Execute()

Returns one app lane's install state for the caller's org: the DocTypes the lane declares, and which of them already exist in the org.



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
	module := "cms" // string | Module is the lane's registered name (\"cms\", \"erp\"), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkModulesByModule(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkModulesByModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkModulesByModule`: ModuleState
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkModulesByModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** | Module is the lane&#39;s registered name (\&quot;cms\&quot;, \&quot;erp\&quot;), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkModulesByModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModuleState**](ModuleState.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFrameworkSummary

> SummaryView GetFrameworkSummary(ctx).Execute()

Reports how much of the DocType surface the caller's org uses: how many DocTypes it has defined, and how many documents exist across them.



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
	resp, r, err := apiClient.FrameworkAPI.GetFrameworkSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.GetFrameworkSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFrameworkSummary`: SummaryView
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.GetFrameworkSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFrameworkSummaryRequest struct via the builder pattern


### Return type

[**SummaryView**](SummaryView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameworkByDoctype

> PostFrameworkByDoctype(ctx, doctype).Execute()

Create one document of a DocType, from that DocType's own fields.



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
	doctype := "doctype_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.PostFrameworkByDoctype(context.Background(), doctype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PostFrameworkByDoctype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameworkByDoctypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameworkByDoctypeByNameCancel

> map[string]map[string]interface{} PostFrameworkByDoctypeByNameCancel(ctx, doctype, name).Execute()

Moves a submitted document to cancelled (docstatus 1 → 2) after its on_cancel hooks agree.



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
	doctype := "Projects.Task" // string | DocType is the document's DocType, by ADDRESS — \"module.name\", from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.PostFrameworkByDoctypeByNameCancel(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PostFrameworkByDoctypeByNameCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameworkByDoctypeByNameCancel`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.PostFrameworkByDoctypeByNameCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, by ADDRESS — \&quot;module.name\&quot;, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameworkByDoctypeByNameCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameworkByDoctypeByNameSubmit

> map[string]map[string]interface{} PostFrameworkByDoctypeByNameSubmit(ctx, doctype, name).Execute()

Moves a draft to submitted (docstatus 0 → 1) after its on_submit hooks agree.



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
	doctype := "Projects.Task" // string | DocType is the document's DocType, by ADDRESS — \"module.name\", from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.PostFrameworkByDoctypeByNameSubmit(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PostFrameworkByDoctypeByNameSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameworkByDoctypeByNameSubmit`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.PostFrameworkByDoctypeByNameSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, by ADDRESS — \&quot;module.name\&quot;, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameworkByDoctypeByNameSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameworkDoctypes

> DocType PostFrameworkDoctypes(ctx).DocType(docType).Execute()

Defines a DocType in the caller's org: the metadata that gives a document surface its fields, its naming rule, whether it has a submit/cancel lifecycle, and which role may do what to it.



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
	docType := *openapiclient.NewDocType() // DocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.PostFrameworkDoctypes(context.Background()).DocType(docType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PostFrameworkDoctypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameworkDoctypes`: DocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.PostFrameworkDoctypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameworkDoctypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **docType** | [**DocType**](DocType.md) |  | 

### Return type

[**DocType**](DocType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFrameworkModulesByModuleInstall

> Install PostFrameworkModulesByModuleInstall(ctx, module).Execute()

Creates an app lane's DocTypes in the caller's org.



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
	module := "cms" // string | Module is the lane's registered name (\"cms\", \"erp\"), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.PostFrameworkModulesByModuleInstall(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PostFrameworkModulesByModuleInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameworkModulesByModuleInstall`: Install
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.PostFrameworkModulesByModuleInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** | Module is the lane&#39;s registered name (\&quot;cms\&quot;, \&quot;erp\&quot;), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameworkModulesByModuleInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Install**](Install.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFrameworkByDoctypeByName

> PutFrameworkByDoctypeByName(ctx, doctype, name).Execute()

Replace a draft document's field data wholesale.



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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.PutFrameworkByDoctypeByName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PutFrameworkByDoctypeByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFrameworkByDoctypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFrameworkDoctypesByName

> DocType PutFrameworkDoctypesByName(ctx, name).DocType(docType).Execute()

Replaces a DocType definition wholesale (PUT semantics): the stored definition becomes the body.



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
	name := "Projects.Task" // string | 
	docType := *openapiclient.NewDocType() // DocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.PutFrameworkDoctypesByName(context.Background(), name).DocType(docType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.PutFrameworkDoctypesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFrameworkDoctypesByName`: DocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.PutFrameworkDoctypesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFrameworkDoctypesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **docType** | [**DocType**](DocType.md) |  | 

### Return type

[**DocType**](DocType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

