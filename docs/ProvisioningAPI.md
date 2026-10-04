# \ProvisioningAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProvisioningDatastoreByName**](ProvisioningAPI.md#DeleteProvisioningDatastoreByName) | **Delete** /v1/provisioning/datastore/{name} | Deprovisions one Hanzo Datastore warehouse.
[**DeleteProvisioningDocdbByName**](ProvisioningAPI.md#DeleteProvisioningDocdbByName) | **Delete** /v1/provisioning/docdb/{name} | Deprovisions one Hanzo DocDB database.
[**DeleteProvisioningKvByName**](ProvisioningAPI.md#DeleteProvisioningKvByName) | **Delete** /v1/provisioning/kv/{name} | Deprovisions one Hanzo KV store.
[**DeleteProvisioningS3ByName**](ProvisioningAPI.md#DeleteProvisioningS3ByName) | **Delete** /v1/provisioning/s3/{name} | Deletes one bucket from the shared object store and removes its metadata row.
[**DeleteProvisioningSearchByName**](ProvisioningAPI.md#DeleteProvisioningSearchByName) | **Delete** /v1/provisioning/search/{name} | Deletes one search index from the shared backend and removes its metadata row.
[**DeleteProvisioningSqlByName**](ProvisioningAPI.md#DeleteProvisioningSqlByName) | **Delete** /v1/provisioning/sql/{name} | Deprovisions one Hanzo SQL database.
[**DeleteProvisioningVectorByName**](ProvisioningAPI.md#DeleteProvisioningVectorByName) | **Delete** /v1/provisioning/vector/{name} | Deletes one vector collection from the shared backend and removes its metadata row.
[**GetProvisioningDatastore**](ProvisioningAPI.md#GetProvisioningDatastore) | **Get** /v1/provisioning/datastore | Lists the caller org&#39;s Hanzo Datastore warehouses.
[**GetProvisioningDatastoreByName**](ProvisioningAPI.md#GetProvisioningDatastoreByName) | **Get** /v1/provisioning/datastore/{name} | Returns one Hanzo Datastore warehouse&#39;s metadata.
[**GetProvisioningDocdb**](ProvisioningAPI.md#GetProvisioningDocdb) | **Get** /v1/provisioning/docdb | Lists the caller org&#39;s Hanzo DocDB document databases.
[**GetProvisioningDocdbByName**](ProvisioningAPI.md#GetProvisioningDocdbByName) | **Get** /v1/provisioning/docdb/{name} | Returns one Hanzo DocDB database&#39;s metadata.
[**GetProvisioningKv**](ProvisioningAPI.md#GetProvisioningKv) | **Get** /v1/provisioning/kv | Lists the caller org&#39;s Hanzo KV stores.
[**GetProvisioningKvByName**](ProvisioningAPI.md#GetProvisioningKvByName) | **Get** /v1/provisioning/kv/{name} | Returns one Hanzo KV store&#39;s metadata.
[**GetProvisioningS3**](ProvisioningAPI.md#GetProvisioningS3) | **Get** /v1/provisioning/s3 | Lists the caller org&#39;s object-storage buckets.
[**GetProvisioningS3ByName**](ProvisioningAPI.md#GetProvisioningS3ByName) | **Get** /v1/provisioning/s3/{name} | Returns one bucket&#39;s metadata.
[**GetProvisioningSearch**](ProvisioningAPI.md#GetProvisioningSearch) | **Get** /v1/provisioning/search | Lists the caller org&#39;s search indexes.
[**GetProvisioningSearchByName**](ProvisioningAPI.md#GetProvisioningSearchByName) | **Get** /v1/provisioning/search/{name} | Returns one search index&#39;s metadata.
[**GetProvisioningSql**](ProvisioningAPI.md#GetProvisioningSql) | **Get** /v1/provisioning/sql | Lists the caller org&#39;s Hanzo SQL databases.
[**GetProvisioningSqlByName**](ProvisioningAPI.md#GetProvisioningSqlByName) | **Get** /v1/provisioning/sql/{name} | Returns one Hanzo SQL database&#39;s metadata.
[**GetProvisioningVector**](ProvisioningAPI.md#GetProvisioningVector) | **Get** /v1/provisioning/vector | Lists the caller org&#39;s vector collections.
[**GetProvisioningVectorByName**](ProvisioningAPI.md#GetProvisioningVectorByName) | **Get** /v1/provisioning/vector/{name} | Returns one vector collection&#39;s metadata.
[**PostProvisioningDatastore**](ProvisioningAPI.md#PostProvisioningDatastore) | **Post** /v1/provisioning/datastore | Launches your org&#39;s OWN Hanzo Datastore instance and answers with its &#x60;datastore://&#x60; connection string.
[**PostProvisioningDocdb**](ProvisioningAPI.md#PostProvisioningDocdb) | **Post** /v1/provisioning/docdb | Launches your org&#39;s OWN document-database instance and answers with its &#x60;mongodb://&#x60; connection string.
[**PostProvisioningKv**](ProvisioningAPI.md#PostProvisioningKv) | **Post** /v1/provisioning/kv | Launches your org&#39;s OWN key-value instance and answers with its &#x60;kv://&#x60; connection string.
[**PostProvisioningS3**](ProvisioningAPI.md#PostProvisioningS3) | **Post** /v1/provisioning/s3 | Creates an S3-compatible bucket inside the already-running shared object store and answers with the endpoint that reaches it.
[**PostProvisioningSearch**](ProvisioningAPI.md#PostProvisioningSearch) | **Post** /v1/provisioning/search | Creates a search index inside the already-running shared search backend and answers with the endpoint that reaches it.
[**PostProvisioningSql**](ProvisioningAPI.md#PostProvisioningSql) | **Post** /v1/provisioning/sql | Launches your org&#39;s OWN PostgreSQL instance and answers with its &#x60;postgres://&#x60; connection string.
[**PostProvisioningVector**](ProvisioningAPI.md#PostProvisioningVector) | **Post** /v1/provisioning/vector | Creates a vector collection inside the already-running shared vector backend and answers with the endpoint that reaches it.



## DeleteProvisioningDatastoreByName

> DeleteProvisioningDatastoreByName(ctx, name).Execute()

Deprovisions one Hanzo Datastore warehouse.



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
	name := "warehouse" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningDatastoreByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningDatastoreByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningDatastoreByNameRequest struct via the builder pattern


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


## DeleteProvisioningDocdbByName

> DeleteProvisioningDocdbByName(ctx, name).Execute()

Deprovisions one Hanzo DocDB database.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningDocdbByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningDocdbByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningDocdbByNameRequest struct via the builder pattern


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


## DeleteProvisioningKvByName

> DeleteProvisioningKvByName(ctx, name).Execute()

Deprovisions one Hanzo KV store.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningKvByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningKvByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningKvByNameRequest struct via the builder pattern


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


## DeleteProvisioningS3ByName

> DeleteProvisioningS3ByName(ctx, name).Execute()

Deletes one bucket from the shared object store and removes its metadata row.



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
	name := "uploads" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningS3ByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningS3ByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningS3ByNameRequest struct via the builder pattern


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


## DeleteProvisioningSearchByName

> DeleteProvisioningSearchByName(ctx, name).Execute()

Deletes one search index from the shared backend and removes its metadata row.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningSearchByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningSearchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningSearchByNameRequest struct via the builder pattern


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


## DeleteProvisioningSqlByName

> DeleteProvisioningSqlByName(ctx, name).Execute()

Deprovisions one Hanzo SQL database.



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
	name := "orders" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningSqlByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningSqlByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningSqlByNameRequest struct via the builder pattern


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


## DeleteProvisioningVectorByName

> DeleteProvisioningVectorByName(ctx, name).Execute()

Deletes one vector collection from the shared backend and removes its metadata row.



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
	name := "embeddings" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProvisioningAPI.DeleteProvisioningVectorByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.DeleteProvisioningVectorByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProvisioningVectorByNameRequest struct via the builder pattern


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


## GetProvisioningDatastore

> []ProvisioningProvisionedSummary GetProvisioningDatastore(ctx).Execute()

Lists the caller org's Hanzo Datastore warehouses.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningDatastore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningDatastore`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningDatastore`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningDatastoreRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningDatastoreByName

> ProvisioningProvisionedResource GetProvisioningDatastoreByName(ctx, name).Execute()

Returns one Hanzo Datastore warehouse's metadata.



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
	name := "warehouse" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningDatastoreByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningDatastoreByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningDatastoreByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningDatastoreByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningDatastoreByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningDocdb

> []ProvisioningProvisionedSummary GetProvisioningDocdb(ctx).Execute()

Lists the caller org's Hanzo DocDB document databases.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningDocdb(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningDocdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningDocdb`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningDocdb`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningDocdbRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningDocdbByName

> ProvisioningProvisionedResource GetProvisioningDocdbByName(ctx, name).Execute()

Returns one Hanzo DocDB database's metadata.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningDocdbByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningDocdbByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningDocdbByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningDocdbByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningDocdbByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningKv

> []ProvisioningProvisionedSummary GetProvisioningKv(ctx).Execute()

Lists the caller org's Hanzo KV stores.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningKv(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningKv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningKv`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningKv`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningKvRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningKvByName

> ProvisioningProvisionedResource GetProvisioningKvByName(ctx, name).Execute()

Returns one Hanzo KV store's metadata.



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
	name := "sessions" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningKvByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningKvByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningKvByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningKvByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningKvByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningS3

> []ProvisioningProvisionedSummary GetProvisioningS3(ctx).Execute()

Lists the caller org's object-storage buckets.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningS3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningS3`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningS3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningS3Request struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningS3ByName

> ProvisioningProvisionedResource GetProvisioningS3ByName(ctx, name).Execute()

Returns one bucket's metadata.



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
	name := "uploads" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningS3ByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningS3ByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningS3ByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningS3ByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningS3ByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningSearch

> []ProvisioningProvisionedSummary GetProvisioningSearch(ctx).Execute()

Lists the caller org's search indexes.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningSearch(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningSearch`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningSearchRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningSearchByName

> ProvisioningProvisionedResource GetProvisioningSearchByName(ctx, name).Execute()

Returns one search index's metadata.



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
	name := "products" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningSearchByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningSearchByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningSearchByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningSearchByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningSearchByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningSql

> []ProvisioningProvisionedSummary GetProvisioningSql(ctx).Execute()

Lists the caller org's Hanzo SQL databases.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningSql(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningSql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningSql`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningSql`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningSqlRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningSqlByName

> ProvisioningProvisionedResource GetProvisioningSqlByName(ctx, name).Execute()

Returns one Hanzo SQL database's metadata.



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
	name := "orders" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningSqlByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningSqlByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningSqlByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningSqlByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningSqlByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningVector

> []ProvisioningProvisionedSummary GetProvisioningVector(ctx).Execute()

Lists the caller org's vector collections.



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
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningVector(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningVector`: []ProvisioningProvisionedSummary
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningVector`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningVectorRequest struct via the builder pattern


### Return type

[**[]ProvisioningProvisionedSummary**](ProvisioningProvisionedSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvisioningVectorByName

> ProvisioningProvisionedResource GetProvisioningVectorByName(ctx, name).Execute()

Returns one vector collection's metadata.



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
	name := "embeddings" // string | Name is the resource's org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.GetProvisioningVectorByName(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.GetProvisioningVectorByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProvisioningVectorByName`: ProvisioningProvisionedResource
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.GetProvisioningVectorByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name is the resource&#39;s org-unique slug, from the path. Lower-cased and trimmed before lookup, exactly as it was at create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningVectorByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningProvisionedResource**](ProvisioningProvisionedResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningDatastore

> ProvisioningProvisionResult PostProvisioningDatastore(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Launches your org's OWN Hanzo Datastore instance and answers with its `datastore://` connection string.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningDatastore(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningDatastore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningDatastore`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningDatastore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningDatastoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningDocdb

> ProvisioningProvisionResult PostProvisioningDocdb(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Launches your org's OWN document-database instance and answers with its `mongodb://` connection string.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningDocdb(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningDocdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningDocdb`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningDocdb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningDocdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningKv

> ProvisioningProvisionResult PostProvisioningKv(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Launches your org's OWN key-value instance and answers with its `kv://` connection string.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningKv(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningKv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningKv`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningKv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningKvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningS3

> ProvisioningProvisionResult PostProvisioningS3(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Creates an S3-compatible bucket inside the already-running shared object store and answers with the endpoint that reaches it.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningS3(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningS3`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningS3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningSearch

> ProvisioningProvisionResult PostProvisioningSearch(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Creates a search index inside the already-running shared search backend and answers with the endpoint that reaches it.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningSearch(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningSearch`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningSql

> ProvisioningProvisionResult PostProvisioningSql(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Launches your org's OWN PostgreSQL instance and answers with its `postgres://` connection string.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningSql(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningSql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningSql`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningSql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningSqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProvisioningVector

> ProvisioningProvisionResult PostProvisioningVector(ctx).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()

Creates a vector collection inside the already-running shared vector backend and answers with the endpoint that reaches it.



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
	provisioningProvisionRequest := *openapiclient.NewProvisioningProvisionRequest() // ProvisioningProvisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProvisioningAPI.PostProvisioningVector(context.Background()).ProvisioningProvisionRequest(provisioningProvisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningAPI.PostProvisioningVector``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostProvisioningVector`: ProvisioningProvisionResult
	fmt.Fprintf(os.Stdout, "Response from `ProvisioningAPI.PostProvisioningVector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProvisioningVectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provisioningProvisionRequest** | [**ProvisioningProvisionRequest**](ProvisioningProvisionRequest.md) |  | 

### Return type

[**ProvisioningProvisionResult**](ProvisioningProvisionResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

