# \ConflictAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldAcled**](ConflictAPI.md#WorldWorldAcled) | **Get** /v1/world/acled | ACLED conflict events (requires ACLED_ACCESS_TOKEN)
[**WorldWorldAcledConflict**](ConflictAPI.md#WorldWorldAcledConflict) | **Get** /v1/world/acled-conflict | ACLED conflict summary (requires ACLED_ACCESS_TOKEN)
[**WorldWorldCyberThreats**](ConflictAPI.md#WorldWorldCyberThreats) | **Get** /v1/world/cyber-threats | Cyber threat feed
[**WorldWorldHapi**](ConflictAPI.md#WorldWorldHapi) | **Get** /v1/world/hapi | Humanitarian API (HDX HAPI)
[**WorldWorldUcdp**](ConflictAPI.md#WorldWorldUcdp) | **Get** /v1/world/ucdp | UCDP conflict data
[**WorldWorldUcdpEvents**](ConflictAPI.md#WorldWorldUcdpEvents) | **Get** /v1/world/ucdp-events | UCDP georeferenced events
[**WorldWorldUnhcrPopulation**](ConflictAPI.md#WorldWorldUnhcrPopulation) | **Get** /v1/world/unhcr-population | UNHCR displacement/population



## WorldWorldAcled

> map[string]interface{} WorldWorldAcled(ctx).Execute()

ACLED conflict events (requires ACLED_ACCESS_TOKEN)

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldAcled(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldAcled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldAcled`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldAcled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldAcledRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldAcledConflict

> map[string]interface{} WorldWorldAcledConflict(ctx).Execute()

ACLED conflict summary (requires ACLED_ACCESS_TOKEN)

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldAcledConflict(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldAcledConflict``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldAcledConflict`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldAcledConflict`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldAcledConflictRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldCyberThreats

> map[string]interface{} WorldWorldCyberThreats(ctx).Execute()

Cyber threat feed

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldCyberThreats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldCyberThreats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCyberThreats`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldCyberThreats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCyberThreatsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldHapi

> map[string]interface{} WorldWorldHapi(ctx).Execute()

Humanitarian API (HDX HAPI)

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldHapi(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldHapi``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldHapi`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldHapi`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldHapiRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldUcdp

> map[string]interface{} WorldWorldUcdp(ctx).Execute()

UCDP conflict data

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldUcdp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldUcdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldUcdp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldUcdp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldUcdpRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldUcdpEvents

> map[string]interface{} WorldWorldUcdpEvents(ctx).Execute()

UCDP georeferenced events

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldUcdpEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldUcdpEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldUcdpEvents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldUcdpEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldUcdpEventsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorldWorldUnhcrPopulation

> map[string]interface{} WorldWorldUnhcrPopulation(ctx).Execute()

UNHCR displacement/population

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
	resp, r, err := apiClient.ConflictAPI.WorldWorldUnhcrPopulation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConflictAPI.WorldWorldUnhcrPopulation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldUnhcrPopulation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConflictAPI.WorldWorldUnhcrPopulation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldUnhcrPopulationRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

