# \DbEndpointsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DbCreateEndpoint**](DbEndpointsAPI.md#DbCreateEndpoint) | **Post** /v1/db/projects/{id}/endpoints | Create compute endpoint
[**DbDeleteEndpoint**](DbEndpointsAPI.md#DbDeleteEndpoint) | **Delete** /v1/db/projects/{id}/endpoints/{endpoint_id} | Delete compute endpoint
[**DbGetEndpoint**](DbEndpointsAPI.md#DbGetEndpoint) | **Get** /v1/db/projects/{id}/endpoints/{endpoint_id} | Get compute endpoint
[**DbListEndpoints**](DbEndpointsAPI.md#DbListEndpoints) | **Get** /v1/db/projects/{id}/endpoints | List compute endpoints
[**DbStartEndpoint**](DbEndpointsAPI.md#DbStartEndpoint) | **Post** /v1/db/projects/{id}/endpoints/{endpoint_id}/start | Start compute endpoint
[**DbSuspendEndpoint**](DbEndpointsAPI.md#DbSuspendEndpoint) | **Post** /v1/db/projects/{id}/endpoints/{endpoint_id}/suspend | Suspend compute endpoint
[**DbUpdateEndpoint**](DbEndpointsAPI.md#DbUpdateEndpoint) | **Put** /v1/db/projects/{id}/endpoints/{endpoint_id} | Update compute endpoint



## DbCreateEndpoint

> DbCreateEndpoint201Response DbCreateEndpoint(ctx, id).DbCreateEndpointRequest(dbCreateEndpointRequest).Execute()

Create compute endpoint

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
	id := "id_example" // string | 
	dbCreateEndpointRequest := *openapiclient.NewDbCreateEndpointRequest(*openapiclient.NewDbEndpointCreate("BranchId_example", "Type_example")) // DbCreateEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbCreateEndpoint(context.Background(), id).DbCreateEndpointRequest(dbCreateEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbCreateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbCreateEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbCreateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbCreateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dbCreateEndpointRequest** | [**DbCreateEndpointRequest**](DbCreateEndpointRequest.md) |  | 

### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbDeleteEndpoint

> DbCreateEndpoint201Response DbDeleteEndpoint(ctx, id, endpointId).Execute()

Delete compute endpoint

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
	id := "id_example" // string | 
	endpointId := "endpointId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbDeleteEndpoint(context.Background(), id, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbDeleteEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbDeleteEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbDeleteEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbDeleteEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbGetEndpoint

> DbCreateEndpoint201Response DbGetEndpoint(ctx, id, endpointId).Execute()

Get compute endpoint

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
	id := "id_example" // string | 
	endpointId := "endpointId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbGetEndpoint(context.Background(), id, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbGetEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbGetEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbGetEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbGetEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbListEndpoints

> DbListEndpoints200Response DbListEndpoints(ctx, id).Execute()

List compute endpoints

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbListEndpoints(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbListEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbListEndpoints`: DbListEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbListEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbListEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DbListEndpoints200Response**](DbListEndpoints200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbStartEndpoint

> DbCreateEndpoint201Response DbStartEndpoint(ctx, id, endpointId).Execute()

Start compute endpoint

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
	id := "id_example" // string | 
	endpointId := "endpointId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbStartEndpoint(context.Background(), id, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbStartEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbStartEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbStartEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbStartEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbSuspendEndpoint

> DbCreateEndpoint201Response DbSuspendEndpoint(ctx, id, endpointId).Execute()

Suspend compute endpoint



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
	id := "id_example" // string | 
	endpointId := "endpointId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbSuspendEndpoint(context.Background(), id, endpointId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbSuspendEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbSuspendEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbSuspendEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbSuspendEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DbUpdateEndpoint

> DbCreateEndpoint201Response DbUpdateEndpoint(ctx, id, endpointId).DbUpdateEndpointRequest(dbUpdateEndpointRequest).Execute()

Update compute endpoint

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
	id := "id_example" // string | 
	endpointId := "endpointId_example" // string | 
	dbUpdateEndpointRequest := *openapiclient.NewDbUpdateEndpointRequest(*openapiclient.NewDbUpdateEndpointRequestEndpoint()) // DbUpdateEndpointRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DbEndpointsAPI.DbUpdateEndpoint(context.Background(), id, endpointId).DbUpdateEndpointRequest(dbUpdateEndpointRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DbEndpointsAPI.DbUpdateEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DbUpdateEndpoint`: DbCreateEndpoint201Response
	fmt.Fprintf(os.Stdout, "Response from `DbEndpointsAPI.DbUpdateEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**endpointId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDbUpdateEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dbUpdateEndpointRequest** | [**DbUpdateEndpointRequest**](DbUpdateEndpointRequest.md) |  | 

### Return type

[**DbCreateEndpoint201Response**](DbCreateEndpoint201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

