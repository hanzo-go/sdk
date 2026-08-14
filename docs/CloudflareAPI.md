# \CloudflareAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCloudflareD1DatabasesByDatabase**](CloudflareAPI.md#DeleteCloudflareD1DatabasesByDatabase) | **Delete** /v1/cloudflare/d1/databases/{database} | Deletes a D1 database and everything stored in it.
[**DeleteCloudflareKvNamespacesByNamespace**](CloudflareAPI.md#DeleteCloudflareKvNamespacesByNamespace) | **Delete** /v1/cloudflare/kv/namespaces/{namespace} | KVNamespaceDelete deletes a Workers KV namespace and every key in it.
[**DeleteCloudflareKvNamespacesByNamespaceValuesByKey**](CloudflareAPI.md#DeleteCloudflareKvNamespacesByNamespaceValuesByKey) | **Delete** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | KVValueDelete removes one key from a Workers KV namespace.
[**DeleteCloudflarePagesProjectsByProject**](CloudflareAPI.md#DeleteCloudflarePagesProjectsByProject) | **Delete** /v1/cloudflare/pages/projects/{project} | Deletes a Cloudflare Pages project, and with it every deployment it has ever made.
[**DeleteCloudflarePagesProjectsByProjectDomainsByDomain**](CloudflareAPI.md#DeleteCloudflarePagesProjectsByProjectDomainsByDomain) | **Delete** /v1/cloudflare/pages/projects/{project}/domains/{domain} | Detaches a custom domain from a Cloudflare Pages project.
[**DeleteCloudflareR2BucketsByBucket**](CloudflareAPI.md#DeleteCloudflareR2BucketsByBucket) | **Delete** /v1/cloudflare/r2/buckets/{bucket} | Deletes an R2 bucket.
[**DeleteCloudflareWorkersScriptsByScript**](CloudflareAPI.md#DeleteCloudflareWorkersScriptsByScript) | **Delete** /v1/cloudflare/workers/scripts/{script} | Removes a Worker script from the org&#39;s Cloudflare account.
[**DeleteCloudflareWorkersZonesByZoneRoutesByRoute**](CloudflareAPI.md#DeleteCloudflareWorkersZonesByZoneRoutesByRoute) | **Delete** /v1/cloudflare/workers/zones/{zone}/routes/{route} | Unbinds a Worker route, so its pattern stops dispatching to a script.
[**GetCloudflareD1Databases**](CloudflareAPI.md#GetCloudflareD1Databases) | **Get** /v1/cloudflare/d1/databases | Lists the D1 databases on the org&#39;s Cloudflare account.
[**GetCloudflareKvNamespaces**](CloudflareAPI.md#GetCloudflareKvNamespaces) | **Get** /v1/cloudflare/kv/namespaces | KVNamespaceList lists the Workers KV namespaces on the org&#39;s Cloudflare account.
[**GetCloudflareKvNamespacesByNamespaceValuesByKey**](CloudflareAPI.md#GetCloudflareKvNamespacesByNamespaceValuesByKey) | **Get** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | Read a Workers KV value as its stored bytes
[**GetCloudflarePagesProjects**](CloudflareAPI.md#GetCloudflarePagesProjects) | **Get** /v1/cloudflare/pages/projects | Lists the org&#39;s Cloudflare Pages projects.
[**GetCloudflarePagesProjectsByProject**](CloudflareAPI.md#GetCloudflarePagesProjectsByProject) | **Get** /v1/cloudflare/pages/projects/{project} | Reads one Cloudflare Pages project — its build config, deployment configs and latest deployment.
[**GetCloudflareR2Buckets**](CloudflareAPI.md#GetCloudflareR2Buckets) | **Get** /v1/cloudflare/r2/buckets | Lists the R2 buckets on the org&#39;s Cloudflare account.
[**GetCloudflareWorkersScripts**](CloudflareAPI.md#GetCloudflareWorkersScripts) | **Get** /v1/cloudflare/workers/scripts | Lists the Worker scripts on the org&#39;s Cloudflare account.
[**GetCloudflareWorkersSubdomain**](CloudflareAPI.md#GetCloudflareWorkersSubdomain) | **Get** /v1/cloudflare/workers/subdomain | Reads the org account&#39;s workers.dev subdomain — the name under which every subdomain-enabled script is served.
[**GetCloudflareWorkersZonesByZoneRoutes**](CloudflareAPI.md#GetCloudflareWorkersZonesByZoneRoutes) | **Get** /v1/cloudflare/workers/zones/{zone}/routes | Lists the Worker routes bound within one zone — the URL patterns that dispatch to a script.
[**GetCloudflareZones**](CloudflareAPI.md#GetCloudflareZones) | **Get** /v1/cloudflare/zones | Lists the Cloudflare zones the org&#39;s connected API token can see, paged and filtered by the query parameters Cloudflare itself accepts.
[**GetCloudflareZonesByZone**](CloudflareAPI.md#GetCloudflareZonesByZone) | **Get** /v1/cloudflare/zones/{zone} | Reads one Cloudflare zone the org&#39;s token can see.
[**GetCloudflareZonesByZoneAnalytics**](CloudflareAPI.md#GetCloudflareZonesByZoneAnalytics) | **Get** /v1/cloudflare/zones/{zone}/analytics | Reads a zone&#39;s Cloudflare traffic dashboard — requests, bandwidth, threats and pageviews over the since/until window.
[**PostCloudflareAiRunByWildcard1**](CloudflareAPI.md#PostCloudflareAiRunByWildcard1) | **Post** /v1/cloudflare/ai/run/{wildcard1} | Run a Cloudflare Workers AI model and get its output back
[**PostCloudflareD1Databases**](CloudflareAPI.md#PostCloudflareD1Databases) | **Post** /v1/cloudflare/d1/databases | Creates a D1 database on the org&#39;s Cloudflare account.
[**PostCloudflareD1DatabasesByDatabaseQuery**](CloudflareAPI.md#PostCloudflareD1DatabasesByDatabaseQuery) | **Post** /v1/cloudflare/d1/databases/{database}/query | Run a SQL statement against a D1 database
[**PostCloudflareKvNamespaces**](CloudflareAPI.md#PostCloudflareKvNamespaces) | **Post** /v1/cloudflare/kv/namespaces | KVNamespaceCreate creates a Workers KV namespace on the org&#39;s Cloudflare account.
[**PostCloudflarePagesProjects**](CloudflareAPI.md#PostCloudflarePagesProjects) | **Post** /v1/cloudflare/pages/projects | Creates a Cloudflare Pages project on the org&#39;s account.
[**PostCloudflarePagesProjectsByProjectDeployments**](CloudflareAPI.md#PostCloudflarePagesProjectsByProjectDeployments) | **Post** /v1/cloudflare/pages/projects/{project}/deployments | Trigger a new Pages deployment for a project
[**PostCloudflarePagesProjectsByProjectDomains**](CloudflareAPI.md#PostCloudflarePagesProjectsByProjectDomains) | **Post** /v1/cloudflare/pages/projects/{project}/domains | Attaches a custom domain to a Cloudflare Pages project.
[**PostCloudflareR2Buckets**](CloudflareAPI.md#PostCloudflareR2Buckets) | **Post** /v1/cloudflare/r2/buckets | Creates an R2 bucket on the org&#39;s Cloudflare account.
[**PostCloudflareWorkersScriptsByScriptSubdomain**](CloudflareAPI.md#PostCloudflareWorkersScriptsByScriptSubdomain) | **Post** /v1/cloudflare/workers/scripts/{script}/subdomain | Publishes or withdraws one Worker script on the account&#39;s workers.dev subdomain.
[**PostCloudflareWorkersZonesByZoneRoutes**](CloudflareAPI.md#PostCloudflareWorkersZonesByZoneRoutes) | **Post** /v1/cloudflare/workers/zones/{zone}/routes | Binds a URL pattern in a zone to a Worker script.
[**PostCloudflareZonesByZonePurge**](CloudflareAPI.md#PostCloudflareZonesByZonePurge) | **Post** /v1/cloudflare/zones/{zone}/purge | Drops a zone&#39;s Cloudflare edge cache — either the whole zone (purge_everything) or exactly the listed file URLs.
[**PutCloudflareKvNamespacesByNamespaceValuesByKey**](CloudflareAPI.md#PutCloudflareKvNamespacesByNamespaceValuesByKey) | **Put** /v1/cloudflare/kv/namespaces/{namespace}/values/{key} | Write a Workers KV value from the request body
[**PutCloudflareWorkersScriptsByScript**](CloudflareAPI.md#PutCloudflareWorkersScriptsByScript) | **Put** /v1/cloudflare/workers/scripts/{script} | Upload or replace a module Worker script



## DeleteCloudflareD1DatabasesByDatabase

> interface{} DeleteCloudflareD1DatabasesByDatabase(ctx, database).Execute()

Deletes a D1 database and everything stored in it.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareD1DatabasesByDatabase(context.Background(), database).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareD1DatabasesByDatabase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareD1DatabasesByDatabase`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareD1DatabasesByDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**database** | **string** | Database is the Cloudflare D1 database id or name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareD1DatabasesByDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflareKvNamespacesByNamespace

> interface{} DeleteCloudflareKvNamespacesByNamespace(ctx, namespace).Execute()

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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareKvNamespacesByNamespace(context.Background(), namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareKvNamespacesByNamespace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareKvNamespacesByNamespace`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareKvNamespacesByNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace is the Cloudflare KV namespace id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareKvNamespacesByNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflareKvNamespacesByNamespaceValuesByKey

> interface{} DeleteCloudflareKvNamespacesByNamespaceValuesByKey(ctx, namespace, key).Execute()

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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareKvNamespacesByNamespaceValuesByKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareKvNamespacesByNamespaceValuesByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareKvNamespacesByNamespaceValuesByKey`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareKvNamespacesByNamespaceValuesByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Namespace is the Cloudflare KV namespace id. | 
**key** | **string** | Key is the key within that namespace. KV keys are broad (up to 512 bytes), so this one is escaped rather than charset-restricted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareKvNamespacesByNamespaceValuesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflarePagesProjectsByProject

> interface{} DeleteCloudflarePagesProjectsByProject(ctx, project).Execute()

Deletes a Cloudflare Pages project, and with it every deployment it has ever made.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflarePagesProjectsByProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflarePagesProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflarePagesProjectsByProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflarePagesProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflarePagesProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflarePagesProjectsByProjectDomainsByDomain

> interface{} DeleteCloudflarePagesProjectsByProjectDomainsByDomain(ctx, project, domain).Execute()

Detaches a custom domain from a Cloudflare Pages project.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflarePagesProjectsByProjectDomainsByDomain(context.Background(), project, domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflarePagesProjectsByProjectDomainsByDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflarePagesProjectsByProjectDomainsByDomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflarePagesProjectsByProjectDomainsByDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 
**domain** | **string** | Domain is the attached custom domain to detach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflarePagesProjectsByProjectDomainsByDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflareR2BucketsByBucket

> interface{} DeleteCloudflareR2BucketsByBucket(ctx, bucket).Execute()

Deletes an R2 bucket.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareR2BucketsByBucket(context.Background(), bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareR2BucketsByBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareR2BucketsByBucket`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareR2BucketsByBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Bucket is the R2 bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareR2BucketsByBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflareWorkersScriptsByScript

> interface{} DeleteCloudflareWorkersScriptsByScript(ctx, script).Execute()

Removes a Worker script from the org's Cloudflare account.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareWorkersScriptsByScript(context.Background(), script).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareWorkersScriptsByScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareWorkersScriptsByScript`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareWorkersScriptsByScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** | Script is the Worker script name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareWorkersScriptsByScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCloudflareWorkersZonesByZoneRoutesByRoute

> interface{} DeleteCloudflareWorkersZonesByZoneRoutesByRoute(ctx, zone, route).Execute()

Unbinds a Worker route, so its pattern stops dispatching to a script.



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
	resp, r, err := apiClient.CloudflareAPI.DeleteCloudflareWorkersZonesByZoneRoutesByRoute(context.Background(), zone, route).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.DeleteCloudflareWorkersZonesByZoneRoutesByRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudflareWorkersZonesByZoneRoutesByRoute`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.DeleteCloudflareWorkersZonesByZoneRoutesByRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 
**route** | **string** | Route is the 32-hex Cloudflare route id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudflareWorkersZonesByZoneRoutesByRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareD1Databases

> interface{} GetCloudflareD1Databases(ctx).Page(page).PerPage(perPage).Name(name).Execute()

Lists the D1 databases on the org's Cloudflare account.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareD1Databases(context.Background()).Page(page).PerPage(perPage).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareD1Databases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareD1Databases`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareD1Databases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareD1DatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of databases to return. | 
 **perPage** | **string** | PerPage is how many databases one page holds. | 
 **name** | **string** | Name filters to the database with this name. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareKvNamespaces

> interface{} GetCloudflareKvNamespaces(ctx).Page(page).PerPage(perPage).Order(order).Direction(direction).Execute()

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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareKvNamespaces(context.Background()).Page(page).PerPage(perPage).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareKvNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareKvNamespaces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareKvNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareKvNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page is the 1-based page of namespaces to return. | 
 **perPage** | **string** | PerPage is how many namespaces one page holds. | 
 **order** | **string** | Order names the field to sort by, and Direction sorts asc or desc. | 
 **direction** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareKvNamespacesByNamespaceValuesByKey

> GetCloudflareKvNamespacesByNamespaceValuesByKey(ctx, namespace, key).Execute()

Read a Workers KV value as its stored bytes



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
	r, err := apiClient.CloudflareAPI.GetCloudflareKvNamespacesByNamespaceValuesByKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareKvNamespacesByNamespaceValuesByKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetCloudflareKvNamespacesByNamespaceValuesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetCloudflarePagesProjects

> interface{} GetCloudflarePagesProjects(ctx).Execute()

Lists the org's Cloudflare Pages projects.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflarePagesProjects(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflarePagesProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflarePagesProjects`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflarePagesProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflarePagesProjectsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflarePagesProjectsByProject

> interface{} GetCloudflarePagesProjectsByProject(ctx, project).Execute()

Reads one Cloudflare Pages project — its build config, deployment configs and latest deployment.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflarePagesProjectsByProject(context.Background(), project).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflarePagesProjectsByProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflarePagesProjectsByProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflarePagesProjectsByProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflarePagesProjectsByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareR2Buckets

> interface{} GetCloudflareR2Buckets(ctx).PerPage(perPage).Cursor(cursor).NameContains(nameContains).Order(order).Direction(direction).Execute()

Lists the R2 buckets on the org's Cloudflare account.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareR2Buckets(context.Background()).PerPage(perPage).Cursor(cursor).NameContains(nameContains).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareR2Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareR2Buckets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareR2Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareR2BucketsRequest struct via the builder pattern


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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareWorkersScripts

> interface{} GetCloudflareWorkersScripts(ctx).Execute()

Lists the Worker scripts on the org's Cloudflare account.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareWorkersScripts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareWorkersScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareWorkersScripts`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareWorkersScripts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareWorkersScriptsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareWorkersSubdomain

> interface{} GetCloudflareWorkersSubdomain(ctx).Execute()

Reads the org account's workers.dev subdomain — the name under which every subdomain-enabled script is served.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareWorkersSubdomain(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareWorkersSubdomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareWorkersSubdomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareWorkersSubdomain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareWorkersSubdomainRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareWorkersZonesByZoneRoutes

> interface{} GetCloudflareWorkersZonesByZoneRoutes(ctx, zone).Execute()

Lists the Worker routes bound within one zone — the URL patterns that dispatch to a script.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareWorkersZonesByZoneRoutes(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareWorkersZonesByZoneRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareWorkersZonesByZoneRoutes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareWorkersZonesByZoneRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareWorkersZonesByZoneRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareZones

> interface{} GetCloudflareZones(ctx).Page(page).PerPage(perPage).Name(name).Status(status).Order(order).Direction(direction).Execute()

Lists the Cloudflare zones the org's connected API token can see, paged and filtered by the query parameters Cloudflare itself accepts.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareZones(context.Background()).Page(page).PerPage(perPage).Name(name).Status(status).Order(order).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareZones`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareZonesRequest struct via the builder pattern


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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareZonesByZone

> interface{} GetCloudflareZonesByZone(ctx, zone).Execute()

Reads one Cloudflare zone the org's token can see.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareZonesByZone(context.Background(), zone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareZonesByZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareZonesByZone`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareZonesByZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareZonesByZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudflareZonesByZoneAnalytics

> interface{} GetCloudflareZonesByZoneAnalytics(ctx, zone).Since(since).Until(until).Continuous(continuous).Execute()

Reads a zone's Cloudflare traffic dashboard — requests, bandwidth, threats and pageviews over the since/until window.



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
	resp, r, err := apiClient.CloudflareAPI.GetCloudflareZonesByZoneAnalytics(context.Background(), zone).Since(since).Until(until).Continuous(continuous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.GetCloudflareZonesByZoneAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudflareZonesByZoneAnalytics`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.GetCloudflareZonesByZoneAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudflareZonesByZoneAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **since** | **string** | Since and Until bound the window, in the form Cloudflare accepts — an RFC 3339 time or a negative number of minutes from now (\&quot;-1440\&quot; is the last day). | 
 **until** | **string** |  | 
 **continuous** | **string** | Continuous asks Cloudflare for only fully-aggregated buckets. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareAiRunByWildcard1

> PostCloudflareAiRunByWildcard1(ctx, wildcard1).Execute()

Run a Cloudflare Workers AI model and get its output back



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
	r, err := apiClient.CloudflareAPI.PostCloudflareAiRunByWildcard1(context.Background(), wildcard1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareAiRunByWildcard1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostCloudflareAiRunByWildcard1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostCloudflareD1Databases

> interface{} PostCloudflareD1Databases(ctx).DatabaseCreateIn(databaseCreateIn).Execute()

Creates a D1 database on the org's Cloudflare account.



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
	databaseCreateIn := *openapiclient.NewDatabaseCreateIn() // DatabaseCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareD1Databases(context.Background()).DatabaseCreateIn(databaseCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareD1Databases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareD1Databases`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareD1Databases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareD1DatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseCreateIn** | [**DatabaseCreateIn**](DatabaseCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareD1DatabasesByDatabaseQuery

> interface{} PostCloudflareD1DatabasesByDatabaseQuery(ctx, database).D1Query(d1Query).Execute()

Run a SQL statement against a D1 database



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
	d1Query := *openapiclient.NewD1Query() // D1Query |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareD1DatabasesByDatabaseQuery(context.Background(), database).D1Query(d1Query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareD1DatabasesByDatabaseQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareD1DatabasesByDatabaseQuery`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareD1DatabasesByDatabaseQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**database** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareD1DatabasesByDatabaseQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **d1Query** | [**D1Query**](D1Query.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareKvNamespaces

> interface{} PostCloudflareKvNamespaces(ctx).NamespaceCreateIn(namespaceCreateIn).Execute()

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
	namespaceCreateIn := *openapiclient.NewNamespaceCreateIn() // NamespaceCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareKvNamespaces(context.Background()).NamespaceCreateIn(namespaceCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareKvNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareKvNamespaces`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareKvNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareKvNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceCreateIn** | [**NamespaceCreateIn**](NamespaceCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflarePagesProjects

> interface{} PostCloudflarePagesProjects(ctx).PagesProjectCreate(pagesProjectCreate).Execute()

Creates a Cloudflare Pages project on the org's account.



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
	pagesProjectCreate := *openapiclient.NewPagesProjectCreate() // PagesProjectCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflarePagesProjects(context.Background()).PagesProjectCreate(pagesProjectCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflarePagesProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflarePagesProjects`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflarePagesProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflarePagesProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagesProjectCreate** | [**PagesProjectCreate**](PagesProjectCreate.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflarePagesProjectsByProjectDeployments

> interface{} PostCloudflarePagesProjectsByProjectDeployments(ctx, project).PagesDeploy(pagesDeploy).Execute()

Trigger a new Pages deployment for a project



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
	pagesDeploy := *openapiclient.NewPagesDeploy() // PagesDeploy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflarePagesProjectsByProjectDeployments(context.Background(), project).PagesDeploy(pagesDeploy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflarePagesProjectsByProjectDeployments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflarePagesProjectsByProjectDeployments`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflarePagesProjectsByProjectDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflarePagesProjectsByProjectDeploymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pagesDeploy** | [**PagesDeploy**](PagesDeploy.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflarePagesProjectsByProjectDomains

> interface{} PostCloudflarePagesProjectsByProjectDomains(ctx, project).DomainAddIn(domainAddIn).Execute()

Attaches a custom domain to a Cloudflare Pages project.



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
	domainAddIn := *openapiclient.NewDomainAddIn() // DomainAddIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflarePagesProjectsByProjectDomains(context.Background(), project).DomainAddIn(domainAddIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflarePagesProjectsByProjectDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflarePagesProjectsByProjectDomains`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflarePagesProjectsByProjectDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**project** | **string** | Project is the Pages project name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflarePagesProjectsByProjectDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domainAddIn** | [**DomainAddIn**](DomainAddIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareR2Buckets

> interface{} PostCloudflareR2Buckets(ctx).BucketCreateIn(bucketCreateIn).Execute()

Creates an R2 bucket on the org's Cloudflare account.



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
	bucketCreateIn := *openapiclient.NewBucketCreateIn() // BucketCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareR2Buckets(context.Background()).BucketCreateIn(bucketCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareR2Buckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareR2Buckets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareR2Buckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareR2BucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucketCreateIn** | [**BucketCreateIn**](BucketCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareWorkersScriptsByScriptSubdomain

> interface{} PostCloudflareWorkersScriptsByScriptSubdomain(ctx, script).SubdomainSetIn(subdomainSetIn).Execute()

Publishes or withdraws one Worker script on the account's workers.dev subdomain.



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
	subdomainSetIn := *openapiclient.NewSubdomainSetIn() // SubdomainSetIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareWorkersScriptsByScriptSubdomain(context.Background(), script).SubdomainSetIn(subdomainSetIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareWorkersScriptsByScriptSubdomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareWorkersScriptsByScriptSubdomain`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareWorkersScriptsByScriptSubdomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** | Script is the Worker script name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareWorkersScriptsByScriptSubdomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subdomainSetIn** | [**SubdomainSetIn**](SubdomainSetIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareWorkersZonesByZoneRoutes

> interface{} PostCloudflareWorkersZonesByZoneRoutes(ctx, zone).RouteCreateIn(routeCreateIn).Execute()

Binds a URL pattern in a zone to a Worker script.



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
	routeCreateIn := *openapiclient.NewRouteCreateIn() // RouteCreateIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareWorkersZonesByZoneRoutes(context.Background(), zone).RouteCreateIn(routeCreateIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareWorkersZonesByZoneRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareWorkersZonesByZoneRoutes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareWorkersZonesByZoneRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareWorkersZonesByZoneRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routeCreateIn** | [**RouteCreateIn**](RouteCreateIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudflareZonesByZonePurge

> interface{} PostCloudflareZonesByZonePurge(ctx, zone).PurgeIn(purgeIn).Execute()

Drops a zone's Cloudflare edge cache — either the whole zone (purge_everything) or exactly the listed file URLs.



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
	purgeIn := *openapiclient.NewPurgeIn() // PurgeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PostCloudflareZonesByZonePurge(context.Background(), zone).PurgeIn(purgeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PostCloudflareZonesByZonePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudflareZonesByZonePurge`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PostCloudflareZonesByZonePurge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zone** | **string** | Zone is the 32-hex Cloudflare zone id, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudflareZonesByZonePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purgeIn** | [**PurgeIn**](PurgeIn.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCloudflareKvNamespacesByNamespaceValuesByKey

> PutCloudflareKvNamespacesByNamespaceValuesByKey(ctx, namespace, key).Execute()

Write a Workers KV value from the request body



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
	r, err := apiClient.CloudflareAPI.PutCloudflareKvNamespacesByNamespaceValuesByKey(context.Background(), namespace, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PutCloudflareKvNamespacesByNamespaceValuesByKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutCloudflareKvNamespacesByNamespaceValuesByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PutCloudflareWorkersScriptsByScript

> interface{} PutCloudflareWorkersScriptsByScript(ctx, script).WorkerScriptPut(workerScriptPut).Execute()

Upload or replace a module Worker script



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
	workerScriptPut := *openapiclient.NewWorkerScriptPut() // WorkerScriptPut |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudflareAPI.PutCloudflareWorkersScriptsByScript(context.Background(), script).WorkerScriptPut(workerScriptPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudflareAPI.PutCloudflareWorkersScriptsByScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCloudflareWorkersScriptsByScript`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CloudflareAPI.PutCloudflareWorkersScriptsByScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**script** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCloudflareWorkersScriptsByScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workerScriptPut** | [**WorkerScriptPut**](WorkerScriptPut.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

