# \IngressAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIngressMiddlewaresById**](IngressAPI.md#DeleteIngressMiddlewaresById) | **Delete** /v1/ingress/middlewares/{id} | Removes one of the caller org&#39;s edge transforms and hot-applies the change.
[**DeleteIngressRoutesById**](IngressAPI.md#DeleteIngressRoutesById) | **Delete** /v1/ingress/routes/{id} | Removes one of the caller org&#39;s routing rules and hot-applies the shrunken table, freeing its host for another claim.
[**DeleteIngressServicesById**](IngressAPI.md#DeleteIngressServicesById) | **Delete** /v1/ingress/services/{id} | Removes one of the caller org&#39;s backend pools and hot-applies the change.
[**GetIngressMiddlewares**](IngressAPI.md#GetIngressMiddlewares) | **Get** /v1/ingress/middlewares | Returns every edge transform the caller&#39;s org has configured, ordered by id.
[**GetIngressMiddlewaresById**](IngressAPI.md#GetIngressMiddlewaresById) | **Get** /v1/ingress/middlewares/{id} | Returns one of the caller org&#39;s edge transforms by id.
[**GetIngressRoutes**](IngressAPI.md#GetIngressRoutes) | **Get** /v1/ingress/routes | Returns every routing rule the caller&#39;s org has configured, ordered by id.
[**GetIngressRoutesById**](IngressAPI.md#GetIngressRoutesById) | **Get** /v1/ingress/routes/{id} | Returns one of the caller org&#39;s routing rules by id.
[**GetIngressServices**](IngressAPI.md#GetIngressServices) | **Get** /v1/ingress/services | Returns every backend pool the caller&#39;s org has configured, ordered by id.
[**GetIngressServicesById**](IngressAPI.md#GetIngressServicesById) | **Get** /v1/ingress/services/{id} | Returns one of the caller org&#39;s backend pools by id.
[**GetIngressStatus**](IngressAPI.md#GetIngressStatus) | **Get** /v1/ingress/status | Reports the ingress edge&#39;s live posture: the role this instance runs in (app or edge), whether its listeners are bound and on which addresses, the ACME posture (staging flag and certificate cache directory), how many hosts the compiled route table currently serves, and how many the ACME HostPolicy will issue a certificate for.
[**GetIngressTls**](IngressAPI.md#GetIngressTls) | **Get** /v1/ingress/tls | Returns the caller org&#39;s ACME intent together with the edge-wide TLS facts it lands in: which role this instance runs in, whether its listeners are bound, every host the ACME HostPolicy will issue a certificate for (the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache), and the ACME directory and account email the process was started with.
[**PostIngressMiddlewares**](IngressAPI.md#PostIngressMiddlewares) | **Post** /v1/ingress/middlewares | Creates or replaces one edge transform and hot-applies it.
[**PostIngressRoutes**](IngressAPI.md#PostIngressRoutes) | **Post** /v1/ingress/routes | Creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.
[**PostIngressServices**](IngressAPI.md#PostIngressServices) | **Post** /v1/ingress/services | Creates or replaces one backend pool and hot-applies it.
[**PutIngressMiddlewaresById**](IngressAPI.md#PutIngressMiddlewaresById) | **Put** /v1/ingress/middlewares/{id} | Creates or replaces one edge transform and hot-applies it.
[**PutIngressRoutesById**](IngressAPI.md#PutIngressRoutesById) | **Put** /v1/ingress/routes/{id} | Creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.
[**PutIngressServicesById**](IngressAPI.md#PutIngressServicesById) | **Put** /v1/ingress/services/{id} | Creates or replaces one backend pool and hot-applies it.
[**PutIngressTls**](IngressAPI.md#PutIngressTls) | **Put** /v1/ingress/tls | Replaces the caller org&#39;s ACME intent and hot-applies what can be hot-applied.



## DeleteIngressMiddlewaresById

> DeleteIngressMiddlewaresById(ctx, id).Execute()

Removes one of the caller org's edge transforms and hot-applies the change.



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
	id := "strip-api" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.DeleteIngressMiddlewaresById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.DeleteIngressMiddlewaresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngressMiddlewaresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngressRoutesById

> DeleteIngressRoutesById(ctx, id).Execute()

Removes one of the caller org's routing rules and hot-applies the shrunken table, freeing its host for another claim.



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
	id := "web" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.DeleteIngressRoutesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.DeleteIngressRoutesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngressRoutesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngressServicesById

> DeleteIngressServicesById(ctx, id).Execute()

Removes one of the caller org's backend pools and hot-applies the change.



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
	id := "app-pool" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.DeleteIngressServicesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.DeleteIngressServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngressServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressMiddlewares

> IngressIngressMiddlewares GetIngressMiddlewares(ctx).Execute()

Returns every edge transform the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.GetIngressMiddlewares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressMiddlewares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressMiddlewares`: IngressIngressMiddlewares
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressMiddlewares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressMiddlewaresRequest struct via the builder pattern


### Return type

[**IngressIngressMiddlewares**](IngressIngressMiddlewares.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressMiddlewaresById

> IngressMiddleware GetIngressMiddlewaresById(ctx, id).Execute()

Returns one of the caller org's edge transforms by id.



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
	id := "strip-api" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.GetIngressMiddlewaresById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressMiddlewaresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressMiddlewaresById`: IngressMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressMiddlewaresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressMiddlewaresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngressMiddleware**](IngressMiddleware.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressRoutes

> IngressIngressRoutes GetIngressRoutes(ctx).Execute()

Returns every routing rule the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.GetIngressRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressRoutes`: IngressIngressRoutes
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressRoutesRequest struct via the builder pattern


### Return type

[**IngressIngressRoutes**](IngressIngressRoutes.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressRoutesById

> IngressRoute GetIngressRoutesById(ctx, id).Execute()

Returns one of the caller org's routing rules by id.



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
	id := "a1b2c3d4e5f60718" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.GetIngressRoutesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressRoutesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressRoutesById`: IngressRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressRoutesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressRoutesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngressRoute**](IngressRoute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressServices

> IngressIngressServices GetIngressServices(ctx).Execute()

Returns every backend pool the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.GetIngressServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressServices`: IngressIngressServices
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressServicesRequest struct via the builder pattern


### Return type

[**IngressIngressServices**](IngressIngressServices.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressServicesById

> IngressUpstream GetIngressServicesById(ctx, id).Execute()

Returns one of the caller org's backend pools by id.



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
	id := "app-pool" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.GetIngressServicesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressServicesById`: IngressUpstream
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngressUpstream**](IngressUpstream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressStatus

> IngressIngressStatus GetIngressStatus(ctx).Execute()

Reports the ingress edge's live posture: the role this instance runs in (app or edge), whether its listeners are bound and on which addresses, the ACME posture (staging flag and certificate cache directory), how many hosts the compiled route table currently serves, and how many the ACME HostPolicy will issue a certificate for.



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
	resp, r, err := apiClient.IngressAPI.GetIngressStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressStatus`: IngressIngressStatus
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressStatusRequest struct via the builder pattern


### Return type

[**IngressIngressStatus**](IngressIngressStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngressTls

> IngressIngressTLS GetIngressTls(ctx).Execute()

Returns the caller org's ACME intent together with the edge-wide TLS facts it lands in: which role this instance runs in, whether its listeners are bound, every host the ACME HostPolicy will issue a certificate for (the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache), and the ACME directory and account email the process was started with.



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
	resp, r, err := apiClient.IngressAPI.GetIngressTls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.GetIngressTls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIngressTls`: IngressIngressTLS
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.GetIngressTls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngressTlsRequest struct via the builder pattern


### Return type

[**IngressIngressTLS**](IngressIngressTLS.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIngressMiddlewares

> IngressMiddleware PostIngressMiddlewares(ctx).IngressMiddleware(ingressMiddleware).Execute()

Creates or replaces one edge transform and hot-applies it.



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
	ingressMiddleware := *openapiclient.NewIngressMiddleware() // IngressMiddleware | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PostIngressMiddlewares(context.Background()).IngressMiddleware(ingressMiddleware).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PostIngressMiddlewares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIngressMiddlewares`: IngressMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PostIngressMiddlewares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIngressMiddlewaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressMiddleware** | [**IngressMiddleware**](IngressMiddleware.md) |  | 

### Return type

[**IngressMiddleware**](IngressMiddleware.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIngressRoutes

> IngressRoute PostIngressRoutes(ctx).IngressRoute(ingressRoute).Execute()

Creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.



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
	ingressRoute := *openapiclient.NewIngressRoute() // IngressRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PostIngressRoutes(context.Background()).IngressRoute(ingressRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PostIngressRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIngressRoutes`: IngressRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PostIngressRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIngressRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressRoute** | [**IngressRoute**](IngressRoute.md) |  | 

### Return type

[**IngressRoute**](IngressRoute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIngressServices

> IngressUpstream PostIngressServices(ctx).IngressUpstream(ingressUpstream).Execute()

Creates or replaces one backend pool and hot-applies it.



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
	ingressUpstream := *openapiclient.NewIngressUpstream() // IngressUpstream | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PostIngressServices(context.Background()).IngressUpstream(ingressUpstream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PostIngressServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIngressServices`: IngressUpstream
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PostIngressServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIngressServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressUpstream** | [**IngressUpstream**](IngressUpstream.md) |  | 

### Return type

[**IngressUpstream**](IngressUpstream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIngressMiddlewaresById

> IngressMiddleware PutIngressMiddlewaresById(ctx, id).IngressMiddleware(ingressMiddleware).Execute()

Creates or replaces one edge transform and hot-applies it.



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
	id := "strip-api" // string | ID identifies the transform within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id.
	ingressMiddleware := *openapiclient.NewIngressMiddleware() // IngressMiddleware | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PutIngressMiddlewaresById(context.Background(), id).IngressMiddleware(ingressMiddleware).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PutIngressMiddlewaresById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIngressMiddlewaresById`: IngressMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PutIngressMiddlewaresById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the transform within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIngressMiddlewaresByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingressMiddleware** | [**IngressMiddleware**](IngressMiddleware.md) |  | 

### Return type

[**IngressMiddleware**](IngressMiddleware.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIngressRoutesById

> IngressRoute PutIngressRoutesById(ctx, id).IngressRoute(ingressRoute).Execute()

Creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.



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
	id := "web" // string | ID identifies the route within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one.
	ingressRoute := *openapiclient.NewIngressRoute() // IngressRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PutIngressRoutesById(context.Background(), id).IngressRoute(ingressRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PutIngressRoutesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIngressRoutesById`: IngressRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PutIngressRoutesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the route within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIngressRoutesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingressRoute** | [**IngressRoute**](IngressRoute.md) |  | 

### Return type

[**IngressRoute**](IngressRoute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIngressServicesById

> IngressUpstream PutIngressServicesById(ctx, id).IngressUpstream(ingressUpstream).Execute()

Creates or replaces one backend pool and hot-applies it.



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
	id := "app-pool" // string | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id.
	ingressUpstream := *openapiclient.NewIngressUpstream() // IngressUpstream | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PutIngressServicesById(context.Background(), id).IngressUpstream(ingressUpstream).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PutIngressServicesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIngressServicesById`: IngressUpstream
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PutIngressServicesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIngressServicesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingressUpstream** | [**IngressUpstream**](IngressUpstream.md) |  | 

### Return type

[**IngressUpstream**](IngressUpstream.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIngressTls

> IngressTLSConfig PutIngressTls(ctx).IngressTLSConfig(ingressTLSConfig).Execute()

Replaces the caller org's ACME intent and hot-applies what can be hot-applied.



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
	ingressTLSConfig := *openapiclient.NewIngressTLSConfig() // IngressTLSConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.PutIngressTls(context.Background()).IngressTLSConfig(ingressTLSConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.PutIngressTls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIngressTls`: IngressTLSConfig
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.PutIngressTls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIngressTlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingressTLSConfig** | [**IngressTLSConfig**](IngressTLSConfig.md) |  | 

### Return type

[**IngressTLSConfig**](IngressTLSConfig.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

