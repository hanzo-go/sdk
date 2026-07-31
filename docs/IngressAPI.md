# \IngressAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1IngressMiddlewaresId**](IngressAPI.md#CloudDeleteV1IngressMiddlewaresId) | **Delete** /v1/ingress/middlewares/{id} | DeleteMiddleware removes one of the caller org&#39;s edge transforms and hot-applies the change.
[**CloudDeleteV1IngressRoutesId**](IngressAPI.md#CloudDeleteV1IngressRoutesId) | **Delete** /v1/ingress/routes/{id} | DeleteRoute removes one of the caller org&#39;s routing rules and hot-applies the shrunken table, freeing its host for another claim.
[**CloudDeleteV1IngressServicesId**](IngressAPI.md#CloudDeleteV1IngressServicesId) | **Delete** /v1/ingress/services/{id} | DeleteService removes one of the caller org&#39;s backend pools and hot-applies the change.
[**CloudGetV1IngressMiddlewares**](IngressAPI.md#CloudGetV1IngressMiddlewares) | **Get** /v1/ingress/middlewares | ListMiddlewares returns every edge transform the caller&#39;s org has configured, ordered by id.
[**CloudGetV1IngressMiddlewaresId**](IngressAPI.md#CloudGetV1IngressMiddlewaresId) | **Get** /v1/ingress/middlewares/{id} | GetMiddleware returns one of the caller org&#39;s edge transforms by id.
[**CloudGetV1IngressRoutes**](IngressAPI.md#CloudGetV1IngressRoutes) | **Get** /v1/ingress/routes | ListRoutes returns every routing rule the caller&#39;s org has configured, ordered by id.
[**CloudGetV1IngressRoutesId**](IngressAPI.md#CloudGetV1IngressRoutesId) | **Get** /v1/ingress/routes/{id} | GetRoute returns one of the caller org&#39;s routing rules by id.
[**CloudGetV1IngressServices**](IngressAPI.md#CloudGetV1IngressServices) | **Get** /v1/ingress/services | ListServices returns every backend pool the caller&#39;s org has configured, ordered by id.
[**CloudGetV1IngressServicesId**](IngressAPI.md#CloudGetV1IngressServicesId) | **Get** /v1/ingress/services/{id} | GetService returns one of the caller org&#39;s backend pools by id.
[**CloudGetV1IngressStatus**](IngressAPI.md#CloudGetV1IngressStatus) | **Get** /v1/ingress/status | Status reports the ingress edge&#39;s live posture: the role this instance runs in (app or edge), whether its listeners are bound and on which addresses, the ACME posture (staging flag and certificate cache directory), how many hosts the compiled route table currently serves, and how many the ACME HostPolicy will issue a certificate for.
[**CloudGetV1IngressTls**](IngressAPI.md#CloudGetV1IngressTls) | **Get** /v1/ingress/tls | GetTLS returns the caller org&#39;s ACME intent together with the edge-wide TLS facts it lands in: which role this instance runs in, whether its listeners are bound, every host the ACME HostPolicy will issue a certificate for (the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache), and the ACME directory and account email the process was started with.
[**CloudPostV1IngressMiddlewares**](IngressAPI.md#CloudPostV1IngressMiddlewares) | **Post** /v1/ingress/middlewares | PutMiddleware creates or replaces one edge transform and hot-applies it.
[**CloudPostV1IngressRoutes**](IngressAPI.md#CloudPostV1IngressRoutes) | **Post** /v1/ingress/routes | PutRoute creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.
[**CloudPostV1IngressServices**](IngressAPI.md#CloudPostV1IngressServices) | **Post** /v1/ingress/services | PutService creates or replaces one backend pool and hot-applies it.
[**CloudPutV1IngressMiddlewaresId**](IngressAPI.md#CloudPutV1IngressMiddlewaresId) | **Put** /v1/ingress/middlewares/{id} | PutMiddleware creates or replaces one edge transform and hot-applies it.
[**CloudPutV1IngressRoutesId**](IngressAPI.md#CloudPutV1IngressRoutesId) | **Put** /v1/ingress/routes/{id} | PutRoute creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.
[**CloudPutV1IngressServicesId**](IngressAPI.md#CloudPutV1IngressServicesId) | **Put** /v1/ingress/services/{id} | PutService creates or replaces one backend pool and hot-applies it.
[**CloudPutV1IngressTls**](IngressAPI.md#CloudPutV1IngressTls) | **Put** /v1/ingress/tls | PutTLS replaces the caller org&#39;s ACME intent and hot-applies what can be hot-applied.



## CloudDeleteV1IngressMiddlewaresId

> CloudDeleteV1IngressMiddlewaresId(ctx, id).Execute()

DeleteMiddleware removes one of the caller org's edge transforms and hot-applies the change.



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
	id := "strip-api" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.CloudDeleteV1IngressMiddlewaresId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudDeleteV1IngressMiddlewaresId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1IngressMiddlewaresIdRequest struct via the builder pattern


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


## CloudDeleteV1IngressRoutesId

> CloudDeleteV1IngressRoutesId(ctx, id).Execute()

DeleteRoute removes one of the caller org's routing rules and hot-applies the shrunken table, freeing its host for another claim.



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
	id := "web" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.CloudDeleteV1IngressRoutesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudDeleteV1IngressRoutesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1IngressRoutesIdRequest struct via the builder pattern


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


## CloudDeleteV1IngressServicesId

> CloudDeleteV1IngressServicesId(ctx, id).Execute()

DeleteService removes one of the caller org's backend pools and hot-applies the change.



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
	id := "app-pool" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IngressAPI.CloudDeleteV1IngressServicesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudDeleteV1IngressServicesId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCloudDeleteV1IngressServicesIdRequest struct via the builder pattern


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


## CloudGetV1IngressMiddlewares

> CloudIngressMiddlewares CloudGetV1IngressMiddlewares(ctx).Execute()

ListMiddlewares returns every edge transform the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressMiddlewares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressMiddlewares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressMiddlewares`: CloudIngressMiddlewares
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressMiddlewares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressMiddlewaresRequest struct via the builder pattern


### Return type

[**CloudIngressMiddlewares**](CloudIngressMiddlewares.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressMiddlewaresId

> CloudMiddleware CloudGetV1IngressMiddlewaresId(ctx, id).Execute()

GetMiddleware returns one of the caller org's edge transforms by id.



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
	id := "strip-api" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressMiddlewaresId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressMiddlewaresId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressMiddlewaresId`: CloudMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressMiddlewaresId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressMiddlewaresIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudMiddleware**](CloudMiddleware.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressRoutes

> CloudIngressRoutes CloudGetV1IngressRoutes(ctx).Execute()

ListRoutes returns every routing rule the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressRoutes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressRoutes`: CloudIngressRoutes
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressRoutes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressRoutesRequest struct via the builder pattern


### Return type

[**CloudIngressRoutes**](CloudIngressRoutes.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressRoutesId

> CloudRoute CloudGetV1IngressRoutesId(ctx, id).Execute()

GetRoute returns one of the caller org's routing rules by id.



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
	id := "a1b2c3d4e5f60718" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressRoutesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressRoutesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressRoutesId`: CloudRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressRoutesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressRoutesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRoute**](CloudRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressServices

> CloudIngressServices CloudGetV1IngressServices(ctx).Execute()

ListServices returns every backend pool the caller's org has configured, ordered by id.



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
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressServices`: CloudIngressServices
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressServicesRequest struct via the builder pattern


### Return type

[**CloudIngressServices**](CloudIngressServices.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressServicesId

> CloudService CloudGetV1IngressServicesId(ctx, id).Execute()

GetService returns one of the caller org's backend pools by id.



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
	id := "app-pool" // string | ID is the object to act on, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressServicesId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressServicesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressServicesId`: CloudService
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressServicesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the object to act on, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressServicesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudService**](CloudService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressStatus

> CloudIngressStatus CloudGetV1IngressStatus(ctx).Execute()

Status reports the ingress edge's live posture: the role this instance runs in (app or edge), whether its listeners are bound and on which addresses, the ACME posture (staging flag and certificate cache directory), how many hosts the compiled route table currently serves, and how many the ACME HostPolicy will issue a certificate for.



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
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressStatus`: CloudIngressStatus
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressStatusRequest struct via the builder pattern


### Return type

[**CloudIngressStatus**](CloudIngressStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1IngressTls

> CloudIngressTLS CloudGetV1IngressTls(ctx).Execute()

GetTLS returns the caller org's ACME intent together with the edge-wide TLS facts it lands in: which role this instance runs in, whether its listeners are bound, every host the ACME HostPolicy will issue a certificate for (the union across ALL orgs of TLS-marked routes and configured extraHosts, because one process holds one certificate cache), and the ACME directory and account email the process was started with.



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
	resp, r, err := apiClient.IngressAPI.CloudGetV1IngressTls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudGetV1IngressTls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1IngressTls`: CloudIngressTLS
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudGetV1IngressTls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1IngressTlsRequest struct via the builder pattern


### Return type

[**CloudIngressTLS**](CloudIngressTLS.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IngressMiddlewares

> CloudMiddleware CloudPostV1IngressMiddlewares(ctx).CloudMiddleware(cloudMiddleware).Execute()

PutMiddleware creates or replaces one edge transform and hot-applies it.



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
	cloudMiddleware := *openapiclient.NewCloudMiddleware() // CloudMiddleware | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPostV1IngressMiddlewares(context.Background()).CloudMiddleware(cloudMiddleware).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPostV1IngressMiddlewares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IngressMiddlewares`: CloudMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPostV1IngressMiddlewares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IngressMiddlewaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudMiddleware** | [**CloudMiddleware**](CloudMiddleware.md) |  | 

### Return type

[**CloudMiddleware**](CloudMiddleware.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IngressRoutes

> CloudRoute CloudPostV1IngressRoutes(ctx).CloudRoute(cloudRoute).Execute()

PutRoute creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.



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
	cloudRoute := *openapiclient.NewCloudRoute() // CloudRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPostV1IngressRoutes(context.Background()).CloudRoute(cloudRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPostV1IngressRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IngressRoutes`: CloudRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPostV1IngressRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IngressRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudRoute** | [**CloudRoute**](CloudRoute.md) |  | 

### Return type

[**CloudRoute**](CloudRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1IngressServices

> CloudService CloudPostV1IngressServices(ctx).CloudService(cloudService).Execute()

PutService creates or replaces one backend pool and hot-applies it.



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
	cloudService := *openapiclient.NewCloudService() // CloudService | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPostV1IngressServices(context.Background()).CloudService(cloudService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPostV1IngressServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1IngressServices`: CloudService
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPostV1IngressServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1IngressServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudService** | [**CloudService**](CloudService.md) |  | 

### Return type

[**CloudService**](CloudService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1IngressMiddlewaresId

> CloudMiddleware CloudPutV1IngressMiddlewaresId(ctx, id).CloudMiddleware(cloudMiddleware).Execute()

PutMiddleware creates or replaces one edge transform and hot-applies it.



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
	id := "strip-api" // string | ID identifies the transform within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id.
	cloudMiddleware := *openapiclient.NewCloudMiddleware() // CloudMiddleware | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPutV1IngressMiddlewaresId(context.Background(), id).CloudMiddleware(cloudMiddleware).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPutV1IngressMiddlewaresId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1IngressMiddlewaresId`: CloudMiddleware
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPutV1IngressMiddlewaresId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the transform within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1IngressMiddlewaresIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudMiddleware** | [**CloudMiddleware**](CloudMiddleware.md) |  | 

### Return type

[**CloudMiddleware**](CloudMiddleware.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1IngressRoutesId

> CloudRoute CloudPutV1IngressRoutesId(ctx, id).CloudRoute(cloudRoute).Execute()

PutRoute creates or replaces one routing rule and hot-applies the new table — there is no config file and no restart.



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
	id := "web" // string | ID identifies the route within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one.
	cloudRoute := *openapiclient.NewCloudRoute() // CloudRoute | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPutV1IngressRoutesId(context.Background(), id).CloudRoute(cloudRoute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPutV1IngressRoutesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1IngressRoutesId`: CloudRoute
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPutV1IngressRoutesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the route within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1IngressRoutesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudRoute** | [**CloudRoute**](CloudRoute.md) |  | 

### Return type

[**CloudRoute**](CloudRoute.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1IngressServicesId

> CloudService CloudPutV1IngressServicesId(ctx, id).CloudService(cloudService).Execute()

PutService creates or replaces one backend pool and hot-applies it.



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
	id := "app-pool" // string | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id.
	cloudService := *openapiclient.NewCloudService() // CloudService | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPutV1IngressServicesId(context.Background(), id).CloudService(cloudService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPutV1IngressServicesId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1IngressServicesId`: CloudService
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPutV1IngressServicesId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID identifies the pool within the org: [A-Za-z0-9-_.], at most 128 chars. A create that omits it gets a generated one. Routes reference it by this id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1IngressServicesIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudService** | [**CloudService**](CloudService.md) |  | 

### Return type

[**CloudService**](CloudService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPutV1IngressTls

> CloudTLSConfig CloudPutV1IngressTls(ctx).CloudTLSConfig(cloudTLSConfig).Execute()

PutTLS replaces the caller org's ACME intent and hot-applies what can be hot-applied.



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
	cloudTLSConfig := *openapiclient.NewCloudTLSConfig() // CloudTLSConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IngressAPI.CloudPutV1IngressTls(context.Background()).CloudTLSConfig(cloudTLSConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IngressAPI.CloudPutV1IngressTls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPutV1IngressTls`: CloudTLSConfig
	fmt.Fprintf(os.Stdout, "Response from `IngressAPI.CloudPutV1IngressTls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPutV1IngressTlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudTLSConfig** | [**CloudTLSConfig**](CloudTLSConfig.md) |  | 

### Return type

[**CloudTLSConfig**](CloudTLSConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

