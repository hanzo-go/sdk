# \CloudflareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1CloudflareD1DatabasesDatabase**](CloudflareAPI.md#CloudDeleteV1CloudflareD1DatabasesDatabase) | **Delete** /v1/cloudflare/d1/databases/{database} | D1DatabaseDelete deletes a D1 database and everything stored in it.
[**CloudDeleteV1CloudflareKvNamespacesNamespace**](CloudflareAPI.md#CloudDeleteV1CloudflareKvNamespacesNamespace) | **Delete** /v1/cloudflare/kv/namespaces/{namespace} | KVNamespaceDelete deletes a Workers KV namespace and every key in it.
[**CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey**](CloudflareAPI.md#CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey) | **Delete** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | KVValueDelete removes one key from a Workers KV namespace.
[**CloudDeleteV1CloudflarePagesProjectsProject**](CloudflareAPI.md#CloudDeleteV1CloudflarePagesProjectsProject) | **Delete** /v1/cloudflare/pages/projects/{project} | PagesDelete deletes a Cloudflare Pages project, and with it every deployment it has ever made.
[**CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain**](CloudflareAPI.md#CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain) | **Delete** /v1/cloudflare/pages/projects/{project}/domains/{domain} | PagesDomainDelete detaches a custom domain from a Cloudflare Pages project.
[**CloudDeleteV1CloudflareR2BucketsBucket**](CloudflareAPI.md#CloudDeleteV1CloudflareR2BucketsBucket) | **Delete** /v1/cloudflare/r2/buckets/{bucket} | R2BucketDelete deletes an R2 bucket.
[**CloudDeleteV1CloudflareWorkersScriptsScript**](CloudflareAPI.md#CloudDeleteV1CloudflareWorkersScriptsScript) | **Delete** /v1/cloudflare/workers/scripts/{script} | WorkersScriptDelete removes a Worker script from the org&#39;s Cloudflare account.
[**CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute**](CloudflareAPI.md#CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute) | **Delete** /v1/cloudflare/workers/zones/{zone}/routes/{route} | WorkersRouteDelete unbinds a Worker route, so its pattern stops dispatching to a script.
[**CloudGetV1CloudflareD1Databases**](CloudflareAPI.md#CloudGetV1CloudflareD1Databases) | **Get** /v1/cloudflare/d1/databases | D1DatabaseList lists the D1 databases on the org&#39;s Cloudflare account.
[**CloudGetV1CloudflareKvNamespaces**](CloudflareAPI.md#CloudGetV1CloudflareKvNamespaces) | **Get** /v1/cloudflare/kv/namespaces | KVNamespaceList lists the Workers KV namespaces on the org&#39;s Cloudflare account.
[**CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey**](CloudflareAPI.md#CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey) | **Get** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | 
[**CloudGetV1CloudflarePagesProjects**](CloudflareAPI.md#CloudGetV1CloudflarePagesProjects) | **Get** /v1/cloudflare/pages/projects | PagesList lists the org&#39;s Cloudflare Pages projects.
[**CloudGetV1CloudflarePagesProjectsProject**](CloudflareAPI.md#CloudGetV1CloudflarePagesProjectsProject) | **Get** /v1/cloudflare/pages/projects/{project} | PagesGet reads one Cloudflare Pages project — its build config, deployment configs and latest deployment.
[**CloudGetV1CloudflareR2Buckets**](CloudflareAPI.md#CloudGetV1CloudflareR2Buckets) | **Get** /v1/cloudflare/r2/buckets | R2BucketList lists the R2 buckets on the org&#39;s Cloudflare account.
[**CloudGetV1CloudflareWorkersScripts**](CloudflareAPI.md#CloudGetV1CloudflareWorkersScripts) | **Get** /v1/cloudflare/workers/scripts | WorkersScriptList lists the Worker scripts on the org&#39;s Cloudflare account.
[**CloudGetV1CloudflareWorkersSubdomain**](CloudflareAPI.md#CloudGetV1CloudflareWorkersSubdomain) | **Get** /v1/cloudflare/workers/subdomain | WorkersSubdomainGet reads the org account&#39;s workers.dev subdomain — the name under which every subdomain-enabled script is served.
[**CloudGetV1CloudflareWorkersZonesZoneRoutes**](CloudflareAPI.md#CloudGetV1CloudflareWorkersZonesZoneRoutes) | **Get** /v1/cloudflare/workers/zones/{zone}/routes | WorkersRouteList lists the Worker routes bound within one zone — the URL patterns that dispatch to a script.
[**CloudGetV1CloudflareZones**](CloudflareAPI.md#CloudGetV1CloudflareZones) | **Get** /v1/cloudflare/zones | ZonesList lists the Cloudflare zones the org&#39;s connected API token can see, paged and filtered by the query parameters Cloudflare itself accepts.
[**CloudGetV1CloudflareZonesZone**](CloudflareAPI.md#CloudGetV1CloudflareZonesZone) | **Get** /v1/cloudflare/zones/{zone} | ZoneGet reads one Cloudflare zone the org&#39;s token can see.
[**CloudGetV1CloudflareZonesZoneAnalytics**](CloudflareAPI.md#CloudGetV1CloudflareZonesZoneAnalytics) | **Get** /v1/cloudflare/zones/{zone}/analytics | ZoneAnalytics reads a zone&#39;s Cloudflare traffic dashboard — requests, bandwidth, threats and pageviews over the since/until window.
[**CloudPostV1CloudflareAiRunByWildcard1**](CloudflareAPI.md#CloudPostV1CloudflareAiRunByWildcard1) | **Post** /v1/cloudflare/ai/run/{wildcard1} | 
[**CloudPostV1CloudflareD1Databases**](CloudflareAPI.md#CloudPostV1CloudflareD1Databases) | **Post** /v1/cloudflare/d1/databases | D1DatabaseCreate creates a D1 database on the org&#39;s Cloudflare account.
[**CloudPostV1CloudflareD1DatabasesByDatabaseQuery**](CloudflareAPI.md#CloudPostV1CloudflareD1DatabasesByDatabaseQuery) | **Post** /v1/cloudflare/d1/databases/{database}/query | 
[**CloudPostV1CloudflareKvNamespaces**](CloudflareAPI.md#CloudPostV1CloudflareKvNamespaces) | **Post** /v1/cloudflare/kv/namespaces | KVNamespaceCreate creates a Workers KV namespace on the org&#39;s Cloudflare account.
[**CloudPostV1CloudflarePagesProjects**](CloudflareAPI.md#CloudPostV1CloudflarePagesProjects) | **Post** /v1/cloudflare/pages/projects | PagesCreate creates a Cloudflare Pages project on the org&#39;s account.
[**CloudPostV1CloudflarePagesProjectsByProjectDeployments**](CloudflareAPI.md#CloudPostV1CloudflarePagesProjectsByProjectDeployments) | **Post** /v1/cloudflare/pages/projects/{project}/deployments | 
[**CloudPostV1CloudflarePagesProjectsProjectDomains**](CloudflareAPI.md#CloudPostV1CloudflarePagesProjectsProjectDomains) | **Post** /v1/cloudflare/pages/projects/{project}/domains | PagesDomainAdd attaches a custom domain to a Cloudflare Pages project.
[**CloudPostV1CloudflareR2Buckets**](CloudflareAPI.md#CloudPostV1CloudflareR2Buckets) | **Post** /v1/cloudflare/r2/buckets | R2BucketCreate creates an R2 bucket on the org&#39;s Cloudflare account.
[**CloudPostV1CloudflareWorkersScriptsScriptSubdomain**](CloudflareAPI.md#CloudPostV1CloudflareWorkersScriptsScriptSubdomain) | **Post** /v1/cloudflare/workers/scripts/{script}/subdomain | WorkersScriptSubdomainSet publishes or withdraws one Worker script on the account&#39;s workers.dev subdomain.
[**CloudPostV1CloudflareWorkersZonesZoneRoutes**](CloudflareAPI.md#CloudPostV1CloudflareWorkersZonesZoneRoutes) | **Post** /v1/cloudflare/workers/zones/{zone}/routes | WorkersRouteCreate binds a URL pattern in a zone to a Worker script.
[**CloudPostV1CloudflareZonesZonePurge**](CloudflareAPI.md#CloudPostV1CloudflareZonesZonePurge) | **Post** /v1/cloudflare/zones/{zone}/purge | ZonePurge drops a zone&#39;s Cloudflare edge cache — either the whole zone (purge_everything) or exactly the listed file URLs.
[**CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey**](CloudflareAPI.md#CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey) | **Put** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | 
[**CloudPutV1CloudflareWorkersScriptsByScript**](CloudflareAPI.md#CloudPutV1CloudflareWorkersScriptsByScript) | **Put** /v1/cloudflare/workers/scripts/{script} | 



## CloudDeleteV1CloudflareD1DatabasesDatabase

> interface{} CloudDeleteV1CloudflareD1DatabasesDatabase(ctx, database).Execute()

D1DatabaseDelete deletes a D1 database and everything stored in it.



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
	database := "orders" // string | Database is the Cloudflare D1 database id or name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareD1DatabasesDatabase(context.Background(), database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareD1DatabasesDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareD1DatabasesDatabase`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareD1DatabasesDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**database** | **string** | Database is the Cloudflare D1 database id or name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareD1DatabasesDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflareKvNamespacesNamespace

> interface{} CloudDeleteV1CloudflareKvNamespacesNamespace(ctx, namespace).Execute()

KVNamespaceDelete deletes a Workers KV namespace and every key in it.



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
	namespace := "0123456789abcdef0123456789abcdef" // string | Namespace is the Cloudflare KV namespace id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespace(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareKvNamespacesNamespace`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace is the Cloudflare KV namespace id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareKvNamespacesNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey

> interface{} CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey(ctx, namespace, key).Execute()

KVValueDelete removes one key from a Workers KV namespace.



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
	namespace := "0123456789abcdef0123456789abcdef" // string | Namespace is the Cloudflare KV namespace id.
	key := "session/abc" // string | Key is the key within that namespace. KV keys are broad (up to 512 bytes), so this one is escaped rather than charset-restricted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareKvNamespacesNamespaceValuesKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace is the Cloudflare KV namespace id. | 
**key** | **string** | Key is the key within that namespace. KV keys are broad (up to 512 bytes), so this one is escaped rather than charset-restricted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareKvNamespacesNamespaceValuesKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflarePagesProjectsProject

> interface{} CloudDeleteV1CloudflarePagesProjectsProject(ctx, project).Execute()

PagesDelete deletes a Cloudflare Pages project, and with it every deployment it has ever made.



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
	project := "marketing-site" // string | Project is the Pages project name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflarePagesProjectsProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflarePagesProjectsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain

> interface{} CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain(ctx, project, domain).Execute()

PagesDomainDelete detaches a custom domain from a Cloudflare Pages project.



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
	project := "marketing-site" // string | Project is the Pages project name.
	domain := "www.acme.com" // string | Domain is the attached custom domain to detach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain(context.Background(), project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflarePagesProjectsProjectDomainsDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 
**domain** | **string** | Domain is the attached custom domain to detach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflarePagesProjectsProjectDomainsDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflareR2BucketsBucket

> interface{} CloudDeleteV1CloudflareR2BucketsBucket(ctx, bucket).Execute()

R2BucketDelete deletes an R2 bucket.



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
	bucket := "assets" // string | Bucket is the R2 bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareR2BucketsBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareR2BucketsBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareR2BucketsBucket`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareR2BucketsBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the R2 bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareR2BucketsBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflareWorkersScriptsScript

> interface{} CloudDeleteV1CloudflareWorkersScriptsScript(ctx, script).Execute()

WorkersScriptDelete removes a Worker script from the org's Cloudflare account.



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
	script := "edge-router" // string | Script is the Worker script name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareWorkersScriptsScript(context.Background(), script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareWorkersScriptsScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareWorkersScriptsScript`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareWorkersScriptsScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** | Script is the Worker script name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareWorkersScriptsScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute

> interface{} CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute(ctx, zone, route).Execute()

WorkersRouteDelete unbinds a Worker route, so its pattern stops dispatching to a script.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id.
	route := "fedcba9876543210fedcba9876543210" // string | Route is the 32-hex Cloudflare route id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute(context.Background(), zone, route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudDeleteV1CloudflareWorkersZonesZoneRoutesRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 
**route** | **string** | Route is the 32-hex Cloudflare route id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1CloudflareWorkersZonesZoneRoutesRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareD1Databases

> interface{} CloudGetV1CloudflareD1Databases(ctx).Page(page).PerPage(perPage).Name(name).Execute()

D1DatabaseList lists the D1 databases on the org's Cloudflare account.



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
	page := "page_example" // string | Page is the 1-based page of databases to return. (optional)
	perPage := "perPage_example" // string | PerPage is how many databases one page holds. (optional)
	name := "name_example" // string | Name filters to the database with this name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareD1Databases(context.Background()).Page(page).PerPage(perPage).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareD1Databases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareD1Databases`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareD1Databases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareD1DatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of databases to return. | 
 **perPage** | **string** | PerPage is how many databases one page holds. | 
 **name** | **string** | Name filters to the database with this name. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareKvNamespaces

> interface{} CloudGetV1CloudflareKvNamespaces(ctx).Page(page).PerPage(perPage).Order(order).Direction(direction).Execute()

KVNamespaceList lists the Workers KV namespaces on the org's Cloudflare account.



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
	page := "page_example" // string | Page is the 1-based page of namespaces to return. (optional)
	perPage := "perPage_example" // string | PerPage is how many namespaces one page holds. (optional)
	order := "order_example" // string | Order names the field to sort by, and Direction sorts asc or desc. (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareKvNamespaces(context.Background()).Page(page).PerPage(perPage).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareKvNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareKvNamespaces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareKvNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareKvNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of namespaces to return. | 
 **perPage** | **string** | PerPage is how many namespaces one page holds. | 
 **order** | **string** | Order names the field to sort by, and Direction sorts asc or desc. | 
 **direction** | **string** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey

> CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey(ctx, namespace, key).Execute()



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
	namespace := "namespace_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareKvNamespacesByNamespaceValuesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareKvNamespacesByNamespaceValuesByKeyRequest struct via the builder pattern


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


## CloudGetV1CloudflarePagesProjects

> interface{} CloudGetV1CloudflarePagesProjects(ctx).Execute()

PagesList lists the org's Cloudflare Pages projects.



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
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflarePagesProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflarePagesProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflarePagesProjects`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflarePagesProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflarePagesProjectsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflarePagesProjectsProject

> interface{} CloudGetV1CloudflarePagesProjectsProject(ctx, project).Execute()

PagesGet reads one Cloudflare Pages project — its build config, deployment configs and latest deployment.



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
	project := "marketing-site" // string | Project is the Pages project name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflarePagesProjectsProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflarePagesProjectsProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflarePagesProjectsProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflarePagesProjectsProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflarePagesProjectsProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareR2Buckets

> interface{} CloudGetV1CloudflareR2Buckets(ctx).PerPage(perPage).Cursor(cursor).NameContains(nameContains).Order(order).Direction(direction).Execute()

R2BucketList lists the R2 buckets on the org's Cloudflare account.



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
	perPage := "perPage_example" // string | PerPage is how many buckets one page holds. (optional)
	cursor := "cursor_example" // string | Cursor continues from the position a previous page returned. (optional)
	nameContains := "nameContains_example" // string | NameContains filters to buckets whose name contains this substring. (optional)
	order := "order_example" // string | Order names the field to sort by, and Direction sorts asc or desc. (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareR2Buckets(context.Background()).PerPage(perPage).Cursor(cursor).NameContains(nameContains).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareR2Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareR2Buckets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareR2Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareR2BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **string** | PerPage is how many buckets one page holds. | 
 **cursor** | **string** | Cursor continues from the position a previous page returned. | 
 **nameContains** | **string** | NameContains filters to buckets whose name contains this substring. | 
 **order** | **string** | Order names the field to sort by, and Direction sorts asc or desc. | 
 **direction** | **string** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareWorkersScripts

> interface{} CloudGetV1CloudflareWorkersScripts(ctx).Execute()

WorkersScriptList lists the Worker scripts on the org's Cloudflare account.



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
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareWorkersScripts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareWorkersScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareWorkersScripts`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareWorkersScripts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareWorkersScriptsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareWorkersSubdomain

> interface{} CloudGetV1CloudflareWorkersSubdomain(ctx).Execute()

WorkersSubdomainGet reads the org account's workers.dev subdomain — the name under which every subdomain-enabled script is served.



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
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareWorkersSubdomain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareWorkersSubdomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareWorkersSubdomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareWorkersSubdomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareWorkersSubdomainRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareWorkersZonesZoneRoutes

> interface{} CloudGetV1CloudflareWorkersZonesZoneRoutes(ctx, zone).Execute()

WorkersRouteList lists the Worker routes bound within one zone — the URL patterns that dispatch to a script.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareWorkersZonesZoneRoutes(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareWorkersZonesZoneRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareWorkersZonesZoneRoutes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareWorkersZonesZoneRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareWorkersZonesZoneRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareZones

> interface{} CloudGetV1CloudflareZones(ctx).Page(page).PerPage(perPage).Name(name).Status(status).Order(order).Direction(direction).Execute()

ZonesList lists the Cloudflare zones the org's connected API token can see, paged and filtered by the query parameters Cloudflare itself accepts.



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
	page := "page_example" // string | Page is the 1-based page of zones to return. (optional)
	perPage := "perPage_example" // string | PerPage is how many zones one page holds. (optional)
	name := "name_example" // string | Name filters to the zone with this domain name. (optional)
	status := "status_example" // string | Status filters by zone status (active, pending, initializing, …). (optional)
	order := "order_example" // string | Order names the field to sort by, and Direction sorts asc or desc. (optional)
	direction := "direction_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareZones(context.Background()).Page(page).PerPage(perPage).Name(name).Status(status).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareZones`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of zones to return. | 
 **perPage** | **string** | PerPage is how many zones one page holds. | 
 **name** | **string** | Name filters to the zone with this domain name. | 
 **status** | **string** | Status filters by zone status (active, pending, initializing, …). | 
 **order** | **string** | Order names the field to sort by, and Direction sorts asc or desc. | 
 **direction** | **string** |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareZonesZone

> interface{} CloudGetV1CloudflareZonesZone(ctx, zone).Execute()

ZoneGet reads one Cloudflare zone the org's token can see.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareZonesZone(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareZonesZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareZonesZone`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareZonesZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareZonesZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1CloudflareZonesZoneAnalytics

> interface{} CloudGetV1CloudflareZonesZoneAnalytics(ctx, zone).Since(since).Until(until).Continuous(continuous).Execute()

ZoneAnalytics reads a zone's Cloudflare traffic dashboard — requests, bandwidth, threats and pageviews over the since/until window.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id.
	since := "-1440" // string | Since and Until bound the window, in the form Cloudflare accepts — an RFC 3339 time or a negative number of minutes from now (\"-1440\" is the last day). (optional)
	until := "0" // string |  (optional)
	continuous := "continuous_example" // string | Continuous asks Cloudflare for only fully-aggregated buckets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudGetV1CloudflareZonesZoneAnalytics(context.Background(), zone).Since(since).Until(until).Continuous(continuous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudGetV1CloudflareZonesZoneAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1CloudflareZonesZoneAnalytics`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudGetV1CloudflareZonesZoneAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1CloudflareZonesZoneAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Since and Until bound the window, in the form Cloudflare accepts — an RFC 3339 time or a negative number of minutes from now (\&quot;-1440\&quot; is the last day). | 
 **until** | **string** |  | 
 **continuous** | **string** | Continuous asks Cloudflare for only fully-aggregated buckets. | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareAiRunByWildcard1

> CloudPostV1CloudflareAiRunByWildcard1(ctx, wildcard1).Execute()



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
	wildcard1 := "wildcard1_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareAiRunByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareAiRunByWildcard1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wildcard1** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareAiRunByWildcard1Request struct via the builder pattern


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


## CloudPostV1CloudflareD1Databases

> interface{} CloudPostV1CloudflareD1Databases(ctx).CloudDatabaseCreateIn(cloudDatabaseCreateIn).Execute()

D1DatabaseCreate creates a D1 database on the org's Cloudflare account.



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
	cloudDatabaseCreateIn := *openapiclient.NewCloudDatabaseCreateIn() // CloudDatabaseCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareD1Databases(context.Background()).CloudDatabaseCreateIn(cloudDatabaseCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareD1Databases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareD1Databases`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareD1Databases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareD1DatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudDatabaseCreateIn** | [**CloudDatabaseCreateIn**](CloudDatabaseCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareD1DatabasesByDatabaseQuery

> interface{} CloudPostV1CloudflareD1DatabasesByDatabaseQuery(ctx, database).CloudD1Query(cloudD1Query).Execute()



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
	database := "database_example" // string | 
	cloudD1Query := *openapiclient.NewCloudD1Query() // CloudD1Query |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareD1DatabasesByDatabaseQuery(context.Background(), database).CloudD1Query(cloudD1Query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareD1DatabasesByDatabaseQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareD1DatabasesByDatabaseQuery`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareD1DatabasesByDatabaseQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareD1DatabasesByDatabaseQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudD1Query** | [**CloudD1Query**](CloudD1Query.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareKvNamespaces

> interface{} CloudPostV1CloudflareKvNamespaces(ctx).CloudNamespaceCreateIn(cloudNamespaceCreateIn).Execute()

KVNamespaceCreate creates a Workers KV namespace on the org's Cloudflare account.



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
	cloudNamespaceCreateIn := *openapiclient.NewCloudNamespaceCreateIn() // CloudNamespaceCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareKvNamespaces(context.Background()).CloudNamespaceCreateIn(cloudNamespaceCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareKvNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareKvNamespaces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareKvNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareKvNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudNamespaceCreateIn** | [**CloudNamespaceCreateIn**](CloudNamespaceCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflarePagesProjects

> interface{} CloudPostV1CloudflarePagesProjects(ctx).CloudPagesProjectCreate(cloudPagesProjectCreate).Execute()

PagesCreate creates a Cloudflare Pages project on the org's account.



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
	cloudPagesProjectCreate := *openapiclient.NewCloudPagesProjectCreate() // CloudPagesProjectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflarePagesProjects(context.Background()).CloudPagesProjectCreate(cloudPagesProjectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflarePagesProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflarePagesProjects`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflarePagesProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflarePagesProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPagesProjectCreate** | [**CloudPagesProjectCreate**](CloudPagesProjectCreate.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflarePagesProjectsByProjectDeployments

> interface{} CloudPostV1CloudflarePagesProjectsByProjectDeployments(ctx, project).CloudPagesDeploy(cloudPagesDeploy).Execute()



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
	project := "project_example" // string | 
	cloudPagesDeploy := *openapiclient.NewCloudPagesDeploy() // CloudPagesDeploy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflarePagesProjectsByProjectDeployments(context.Background(), project).CloudPagesDeploy(cloudPagesDeploy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflarePagesProjectsByProjectDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflarePagesProjectsByProjectDeployments`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflarePagesProjectsByProjectDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflarePagesProjectsByProjectDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPagesDeploy** | [**CloudPagesDeploy**](CloudPagesDeploy.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflarePagesProjectsProjectDomains

> interface{} CloudPostV1CloudflarePagesProjectsProjectDomains(ctx, project).CloudDomainAddIn(cloudDomainAddIn).Execute()

PagesDomainAdd attaches a custom domain to a Cloudflare Pages project.



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
	project := "marketing-site" // string | Project is the Pages project name, from the path.
	cloudDomainAddIn := *openapiclient.NewCloudDomainAddIn() // CloudDomainAddIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflarePagesProjectsProjectDomains(context.Background(), project).CloudDomainAddIn(cloudDomainAddIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflarePagesProjectsProjectDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflarePagesProjectsProjectDomains`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflarePagesProjectsProjectDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflarePagesProjectsProjectDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudDomainAddIn** | [**CloudDomainAddIn**](CloudDomainAddIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareR2Buckets

> interface{} CloudPostV1CloudflareR2Buckets(ctx).CloudBucketCreateIn(cloudBucketCreateIn).Execute()

R2BucketCreate creates an R2 bucket on the org's Cloudflare account.



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
	cloudBucketCreateIn := *openapiclient.NewCloudBucketCreateIn() // CloudBucketCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareR2Buckets(context.Background()).CloudBucketCreateIn(cloudBucketCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareR2Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareR2Buckets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareR2Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareR2BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudBucketCreateIn** | [**CloudBucketCreateIn**](CloudBucketCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareWorkersScriptsScriptSubdomain

> interface{} CloudPostV1CloudflareWorkersScriptsScriptSubdomain(ctx, script).CloudSubdomainSetIn(cloudSubdomainSetIn).Execute()

WorkersScriptSubdomainSet publishes or withdraws one Worker script on the account's workers.dev subdomain.



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
	script := "edge-router" // string | Script is the Worker script name, from the path.
	cloudSubdomainSetIn := *openapiclient.NewCloudSubdomainSetIn() // CloudSubdomainSetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareWorkersScriptsScriptSubdomain(context.Background(), script).CloudSubdomainSetIn(cloudSubdomainSetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareWorkersScriptsScriptSubdomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareWorkersScriptsScriptSubdomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareWorkersScriptsScriptSubdomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** | Script is the Worker script name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareWorkersScriptsScriptSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSubdomainSetIn** | [**CloudSubdomainSetIn**](CloudSubdomainSetIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareWorkersZonesZoneRoutes

> interface{} CloudPostV1CloudflareWorkersZonesZoneRoutes(ctx, zone).CloudRouteCreateIn(cloudRouteCreateIn).Execute()

WorkersRouteCreate binds a URL pattern in a zone to a Worker script.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id, from the path.
	cloudRouteCreateIn := *openapiclient.NewCloudRouteCreateIn() // CloudRouteCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareWorkersZonesZoneRoutes(context.Background(), zone).CloudRouteCreateIn(cloudRouteCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareWorkersZonesZoneRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareWorkersZonesZoneRoutes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareWorkersZonesZoneRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareWorkersZonesZoneRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRouteCreateIn** | [**CloudRouteCreateIn**](CloudRouteCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1CloudflareZonesZonePurge

> interface{} CloudPostV1CloudflareZonesZonePurge(ctx, zone).CloudPurgeIn(cloudPurgeIn).Execute()

ZonePurge drops a zone's Cloudflare edge cache — either the whole zone (purge_everything) or exactly the listed file URLs.



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
	zone := "0123456789abcdef0123456789abcdef" // string | Zone is the 32-hex Cloudflare zone id, from the path.
	cloudPurgeIn := *openapiclient.NewCloudPurgeIn() // CloudPurgeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPostV1CloudflareZonesZonePurge(context.Background(), zone).CloudPurgeIn(cloudPurgeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPostV1CloudflareZonesZonePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1CloudflareZonesZonePurge`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPostV1CloudflareZonesZonePurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1CloudflareZonesZonePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudPurgeIn** | [**CloudPurgeIn**](CloudPurgeIn.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey

> CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey(ctx, namespace, key).Execute()



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
	namespace := "namespace_example" // string | 
	key := "key_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CloudflareAPI.CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPutV1CloudflareKvNamespacesByNamespaceValuesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CloudflareKvNamespacesByNamespaceValuesByKeyRequest struct via the builder pattern


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


## CloudPutV1CloudflareWorkersScriptsByScript

> interface{} CloudPutV1CloudflareWorkersScriptsByScript(ctx, script).CloudWorkerScriptPut(cloudWorkerScriptPut).Execute()



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
	script := "script_example" // string | 
	cloudWorkerScriptPut := *openapiclient.NewCloudWorkerScriptPut() // CloudWorkerScriptPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.CloudPutV1CloudflareWorkersScriptsByScript(context.Background(), script).CloudWorkerScriptPut(cloudWorkerScriptPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.CloudPutV1CloudflareWorkersScriptsByScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1CloudflareWorkersScriptsByScript`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.CloudPutV1CloudflareWorkersScriptsByScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1CloudflareWorkersScriptsByScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudWorkerScriptPut** | [**CloudWorkerScriptPut**](CloudWorkerScriptPut.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

