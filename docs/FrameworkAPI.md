# \FrameworkAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1FrameworkDoctypeName**](FrameworkAPI.md#CloudDeleteV1FrameworkDoctypeName) | **Delete** /v1/framework/{doctype}/{name} | Removes one document, after its on_trash hooks agree.
[**CloudDeleteV1FrameworkDoctypesName**](FrameworkAPI.md#CloudDeleteV1FrameworkDoctypesName) | **Delete** /v1/framework/doctypes/{name} | Removes a DocType and every document stored under it.
[**CloudDeleteV1FrameworkRolesUserRole**](FrameworkAPI.md#CloudDeleteV1FrameworkRolesUserRole) | **Delete** /v1/framework/roles/{user}/{role} | Removes one (user, role) grant in the caller&#39;s org.
[**CloudGetV1FrameworkDoctype**](FrameworkAPI.md#CloudGetV1FrameworkDoctype) | **Get** /v1/framework/{doctype} | Returns the caller org&#39;s documents of one DocType, filtered, ordered and projected by the query.
[**CloudGetV1FrameworkDoctypeName**](FrameworkAPI.md#CloudGetV1FrameworkDoctypeName) | **Get** /v1/framework/{doctype}/{name} | Returns one document by name, with Password fields redacted.
[**CloudGetV1FrameworkDoctypes**](FrameworkAPI.md#CloudGetV1FrameworkDoctypes) | **Get** /v1/framework/doctypes | Returns every DocType defined in the caller&#39;s org.
[**CloudGetV1FrameworkDoctypesName**](FrameworkAPI.md#CloudGetV1FrameworkDoctypesName) | **Get** /v1/framework/doctypes/{name} | Returns one DocType definition — its fields, naming rule, permissions and lifecycle flags.
[**CloudGetV1FrameworkModules**](FrameworkAPI.md#CloudGetV1FrameworkModules) | **Get** /v1/framework/modules | Returns every app lane compiled into this deployment and the DocTypes each one installs.
[**CloudGetV1FrameworkModulesModule**](FrameworkAPI.md#CloudGetV1FrameworkModulesModule) | **Get** /v1/framework/modules/{module} | Returns one app lane&#39;s install state for the caller&#39;s org: the DocTypes the lane declares, and which of them already exist in the org.
[**CloudGetV1FrameworkRoles**](FrameworkAPI.md#CloudGetV1FrameworkRoles) | **Get** /v1/framework/roles | Returns every (user, role) assignment in the caller&#39;s org.
[**CloudGetV1FrameworkSummary**](FrameworkAPI.md#CloudGetV1FrameworkSummary) | **Get** /v1/framework/summary | Reports how much of the DocType surface the caller&#39;s org uses: how many DocTypes it has defined, and how many documents exist across them.
[**CloudPostV1FrameworkByDoctype**](FrameworkAPI.md#CloudPostV1FrameworkByDoctype) | **Post** /v1/framework/{doctype} | 
[**CloudPostV1FrameworkDoctypeNameCancel**](FrameworkAPI.md#CloudPostV1FrameworkDoctypeNameCancel) | **Post** /v1/framework/{doctype}/{name}/cancel | Moves a submitted document to cancelled (docstatus 1 → 2) after its on_cancel hooks agree.
[**CloudPostV1FrameworkDoctypeNameSubmit**](FrameworkAPI.md#CloudPostV1FrameworkDoctypeNameSubmit) | **Post** /v1/framework/{doctype}/{name}/submit | Moves a draft to submitted (docstatus 0 → 1) after its on_submit hooks agree.
[**CloudPostV1FrameworkDoctypes**](FrameworkAPI.md#CloudPostV1FrameworkDoctypes) | **Post** /v1/framework/doctypes | Defines a DocType in the caller&#39;s org: the metadata that gives a document surface its fields, its naming rule, whether it has a submit/cancel lifecycle, and which role may do what to it.
[**CloudPostV1FrameworkModulesModuleInstall**](FrameworkAPI.md#CloudPostV1FrameworkModulesModuleInstall) | **Post** /v1/framework/modules/{module}/install | Creates an app lane&#39;s DocTypes in the caller&#39;s org.
[**CloudPostV1FrameworkRoles**](FrameworkAPI.md#CloudPostV1FrameworkRoles) | **Post** /v1/framework/roles | Grants one user one role in the caller&#39;s org — how a member gains rights on a DocType, since permissions name roles and never users.
[**CloudPutV1FrameworkByDoctypeByName**](FrameworkAPI.md#CloudPutV1FrameworkByDoctypeByName) | **Put** /v1/framework/{doctype}/{name} | 
[**CloudPutV1FrameworkDoctypesName**](FrameworkAPI.md#CloudPutV1FrameworkDoctypesName) | **Put** /v1/framework/doctypes/{name} | Replaces a DocType definition wholesale (PUT semantics): the stored definition becomes the body.



## CloudDeleteV1FrameworkDoctypeName

> CloudDeleteV1FrameworkDoctypeName(ctx, doctype, name).Execute()

Removes one document, after its on_trash hooks agree.



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
	doctype := "Task" // string | DocType is the document's DocType, from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.CloudDeleteV1FrameworkDoctypeName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudDeleteV1FrameworkDoctypeName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1FrameworkDoctypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CloudDeleteV1FrameworkDoctypesName

> CloudDeleteV1FrameworkDoctypesName(ctx, name).Execute()

Removes a DocType and every document stored under it.



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
	name := "Task" // string | Name is the DocType's name, from the path. A name containing a space (\"Sales Invoice\") arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.CloudDeleteV1FrameworkDoctypesName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudDeleteV1FrameworkDoctypesName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the DocType&#39;s name, from the path. A name containing a space (\&quot;Sales Invoice\&quot;) arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1FrameworkDoctypesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudDeleteV1FrameworkRolesUserRole

> CloudDeleteV1FrameworkRolesUserRole(ctx, user, role).Execute()

Removes one (user, role) grant in the caller's org.



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
	user := "u_alice" // string | User is the assignee whose grant is being revoked, from the path.
	role := "System Manager" // string | Role is the role to revoke, from the path. A role name containing a space (\"System Manager\") arrives percent-encoded and is decoded before it is matched against the stored assignment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.CloudDeleteV1FrameworkRolesUserRole(context.Background(), user, role).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudDeleteV1FrameworkRolesUserRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | User is the assignee whose grant is being revoked, from the path. | 
**role** | **string** | Role is the role to revoke, from the path. A role name containing a space (\&quot;System Manager\&quot;) arrives percent-encoded and is decoded before it is matched against the stored assignment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1FrameworkRolesUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CloudGetV1FrameworkDoctype

> CloudDocumentList CloudGetV1FrameworkDoctype(ctx, doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()

Returns the caller org's documents of one DocType, filtered, ordered and projected by the query.



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
	doctype := "Task" // string | DocType is the DocType to list, from the path.
	filters := "{"priority":"High"}" // string | Filters is a JSON object of equality matches, e.g. {\"priority\":\"High\"}. Every key must be a field the DocType declares (or the managed name / docstatus); an undeclared one is refused rather than silently ignored. (optional)
	fields := "fields_example" // string | Fields projects the response to a subset — a JSON array [\"a\",\"b\"] or a comma list \"a,b\". The envelope keys are always returned. (optional)
	orderBy := "estimate asc" // string | OrderBy is \"<field> [asc|desc]\". Empty means most-recently-updated first. (optional)
	limit := "20" // string | Limit caps the rows returned. Anything that is not a positive integer leaves the engine's default in place. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkDoctype(context.Background(), doctype).Filters(filters).Fields(fields).OrderBy(orderBy).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkDoctype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkDoctype`: CloudDocumentList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkDoctype`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the DocType to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkDoctypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | Filters is a JSON object of equality matches, e.g. {\&quot;priority\&quot;:\&quot;High\&quot;}. Every key must be a field the DocType declares (or the managed name / docstatus); an undeclared one is refused rather than silently ignored. | 
 **fields** | **string** | Fields projects the response to a subset — a JSON array [\&quot;a\&quot;,\&quot;b\&quot;] or a comma list \&quot;a,b\&quot;. The envelope keys are always returned. | 
 **orderBy** | **string** | OrderBy is \&quot;&lt;field&gt; [asc|desc]\&quot;. Empty means most-recently-updated first. | 
 **limit** | **string** | Limit caps the rows returned. Anything that is not a positive integer leaves the engine&#39;s default in place. | 

### Return type

[**CloudDocumentList**](CloudDocumentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkDoctypeName

> map[string]map[string]interface{} CloudGetV1FrameworkDoctypeName(ctx, doctype, name).Execute()

Returns one document by name, with Password fields redacted.



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
	doctype := "Task" // string | DocType is the document's DocType, from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkDoctypeName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkDoctypeName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkDoctypeName`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkDoctypeName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkDoctypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkDoctypes

> CloudDocTypeList CloudGetV1FrameworkDoctypes(ctx).Execute()

Returns every DocType defined in the caller's org.



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
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkDoctypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkDoctypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkDoctypes`: CloudDocTypeList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkDoctypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkDoctypesRequest struct via the builder pattern


### Return type

[**CloudDocTypeList**](CloudDocTypeList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkDoctypesName

> CloudDocType CloudGetV1FrameworkDoctypesName(ctx, name).Execute()

Returns one DocType definition — its fields, naming rule, permissions and lifecycle flags.



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
	name := "Task" // string | Name is the DocType's name, from the path. A name containing a space (\"Sales Invoice\") arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkDoctypesName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkDoctypesName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkDoctypesName`: CloudDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkDoctypesName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the DocType&#39;s name, from the path. A name containing a space (\&quot;Sales Invoice\&quot;) arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkDoctypesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudDocType**](CloudDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkModules

> CloudModuleList CloudGetV1FrameworkModules(ctx).Execute()

Returns every app lane compiled into this deployment and the DocTypes each one installs.



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
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkModules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkModules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkModules`: CloudModuleList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkModulesRequest struct via the builder pattern


### Return type

[**CloudModuleList**](CloudModuleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkModulesModule

> CloudModuleState CloudGetV1FrameworkModulesModule(ctx, module).Execute()

Returns one app lane's install state for the caller's org: the DocTypes the lane declares, and which of them already exist in the org.



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
	module := "cms" // string | Module is the lane's registered name (\"cms\", \"erp\"), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkModulesModule(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkModulesModule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkModulesModule`: CloudModuleState
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkModulesModule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** | Module is the lane&#39;s registered name (\&quot;cms\&quot;, \&quot;erp\&quot;), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkModulesModuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudModuleState**](CloudModuleState.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkRoles

> CloudRoleList CloudGetV1FrameworkRoles(ctx).Execute()

Returns every (user, role) assignment in the caller's org.



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
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkRoles`: CloudRoleList
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkRolesRequest struct via the builder pattern


### Return type

[**CloudRoleList**](CloudRoleList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1FrameworkSummary

> CloudSummaryView CloudGetV1FrameworkSummary(ctx).Execute()

Reports how much of the DocType surface the caller's org uses: how many DocTypes it has defined, and how many documents exist across them.



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
	resp, r, err := apiClient.FrameworkAPI.CloudGetV1FrameworkSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudGetV1FrameworkSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1FrameworkSummary`: CloudSummaryView
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudGetV1FrameworkSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1FrameworkSummaryRequest struct via the builder pattern


### Return type

[**CloudSummaryView**](CloudSummaryView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FrameworkByDoctype

> CloudPostV1FrameworkByDoctype(ctx, doctype).Execute()



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
	doctype := "doctype_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkByDoctype(context.Background(), doctype).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkByDoctype``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkByDoctypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CloudPostV1FrameworkDoctypeNameCancel

> map[string]map[string]interface{} CloudPostV1FrameworkDoctypeNameCancel(ctx, doctype, name).Execute()

Moves a submitted document to cancelled (docstatus 1 → 2) after its on_cancel hooks agree.



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
	doctype := "Task" // string | DocType is the document's DocType, from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkDoctypeNameCancel(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkDoctypeNameCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FrameworkDoctypeNameCancel`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPostV1FrameworkDoctypeNameCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkDoctypeNameCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FrameworkDoctypeNameSubmit

> map[string]map[string]interface{} CloudPostV1FrameworkDoctypeNameSubmit(ctx, doctype, name).Execute()

Moves a draft to submitted (docstatus 0 → 1) after its on_submit hooks agree.



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
	doctype := "Task" // string | DocType is the document's DocType, from the path.
	name := "TASK-00001" // string | Name is the document's name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkDoctypeNameSubmit(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkDoctypeNameSubmit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FrameworkDoctypeNameSubmit`: map[string]map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPostV1FrameworkDoctypeNameSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**doctype** | **string** | DocType is the document&#39;s DocType, from the path. | 
**name** | **string** | Name is the document&#39;s name — its key within the DocType — from the path. A name containing a space arrives percent-encoded and is decoded before it is matched against the stored one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkDoctypeNameSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FrameworkDoctypes

> CloudDocType CloudPostV1FrameworkDoctypes(ctx).CloudDocType(cloudDocType).Execute()

Defines a DocType in the caller's org: the metadata that gives a document surface its fields, its naming rule, whether it has a submit/cancel lifecycle, and which role may do what to it.



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
	cloudDocType := *openapiclient.NewCloudDocType() // CloudDocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkDoctypes(context.Background()).CloudDocType(cloudDocType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkDoctypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FrameworkDoctypes`: CloudDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPostV1FrameworkDoctypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkDoctypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDocType** | [**CloudDocType**](CloudDocType.md) |  | 

### Return type

[**CloudDocType**](CloudDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FrameworkModulesModuleInstall

> CloudInstall CloudPostV1FrameworkModulesModuleInstall(ctx, module).Execute()

Creates an app lane's DocTypes in the caller's org.



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
	module := "cms" // string | Module is the lane's registered name (\"cms\", \"erp\"), from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkModulesModuleInstall(context.Background(), module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkModulesModuleInstall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FrameworkModulesModuleInstall`: CloudInstall
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPostV1FrameworkModulesModuleInstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**module** | **string** | Module is the lane&#39;s registered name (\&quot;cms\&quot;, \&quot;erp\&quot;), from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkModulesModuleInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudInstall**](CloudInstall.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1FrameworkRoles

> CloudRole CloudPostV1FrameworkRoles(ctx).CloudRole(cloudRole).Execute()

Grants one user one role in the caller's org — how a member gains rights on a DocType, since permissions name roles and never users.



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
	cloudRole := *openapiclient.NewCloudRole() // CloudRole | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPostV1FrameworkRoles(context.Background()).CloudRole(cloudRole).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPostV1FrameworkRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1FrameworkRoles`: CloudRole
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPostV1FrameworkRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1FrameworkRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRole** | [**CloudRole**](CloudRole.md) |  | 

### Return type

[**CloudRole**](CloudRole.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1FrameworkByDoctypeByName

> CloudPutV1FrameworkByDoctypeByName(ctx, doctype, name).Execute()



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
	doctype := "doctype_example" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FrameworkAPI.CloudPutV1FrameworkByDoctypeByName(context.Background(), doctype, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPutV1FrameworkByDoctypeByName``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudPutV1FrameworkByDoctypeByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CloudPutV1FrameworkDoctypesName

> CloudDocType CloudPutV1FrameworkDoctypesName(ctx, name).CloudDocType(cloudDocType).Execute()

Replaces a DocType definition wholesale (PUT semantics): the stored definition becomes the body.



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
	name := "Task" // string | 
	cloudDocType := *openapiclient.NewCloudDocType() // CloudDocType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameworkAPI.CloudPutV1FrameworkDoctypesName(context.Background(), name).CloudDocType(cloudDocType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameworkAPI.CloudPutV1FrameworkDoctypesName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1FrameworkDoctypesName`: CloudDocType
	fmt.Fprintf(os.Stdout, "Response from `FrameworkAPI.CloudPutV1FrameworkDoctypesName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1FrameworkDoctypesNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudDocType** | [**CloudDocType**](CloudDocType.md) |  | 

### Return type

[**CloudDocType**](CloudDocType.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

