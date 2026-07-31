# \GeoAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorldWorldAisSnapshot**](GeoAPI.md#WorldWorldAisSnapshot) | **Get** /v1/world/ais-snapshot | AIS vessel snapshot (requires WS_RELAY_URL)
[**WorldWorldClimateAnomalies**](GeoAPI.md#WorldWorldClimateAnomalies) | **Get** /v1/world/climate-anomalies | Climate anomaly feed
[**WorldWorldCloudflareOutages**](GeoAPI.md#WorldWorldCloudflareOutages) | **Get** /v1/world/cloudflare-outages | Cloudflare internet outages
[**WorldWorldEarthquakes**](GeoAPI.md#WorldWorldEarthquakes) | **Get** /v1/world/earthquakes | USGS earthquake feed
[**WorldWorldFaaStatus**](GeoAPI.md#WorldWorldFaaStatus) | **Get** /v1/world/faa-status | FAA airport status
[**WorldWorldFirmsFires**](GeoAPI.md#WorldWorldFirmsFires) | **Get** /v1/world/firms-fires | NASA FIRMS active fires (requires NASA_FIRMS_API_KEY)
[**WorldWorldNgaWarnings**](GeoAPI.md#WorldWorldNgaWarnings) | **Get** /v1/world/nga-warnings | NGA maritime safety warnings
[**WorldWorldOpensky**](GeoAPI.md#WorldWorldOpensky) | **Get** /v1/world/opensky | OpenSky flight states
[**WorldWorldWingbits**](GeoAPI.md#WorldWorldWingbits) | **Get** /v1/world/wingbits | Wingbits ADS-B (requires WINGBITS_API_KEY)
[**WorldWorldWorldpopExposure**](GeoAPI.md#WorldWorldWorldpopExposure) | **Get** /v1/world/worldpop-exposure | Population exposure model



## WorldWorldAisSnapshot

> map[string]interface{} WorldWorldAisSnapshot(ctx).Execute()

AIS vessel snapshot (requires WS_RELAY_URL)

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
	resp, r, err := apiClient.GeoAPI.WorldWorldAisSnapshot(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldAisSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldAisSnapshot`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldAisSnapshot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldAisSnapshotRequest struct via the builder pattern


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


## WorldWorldClimateAnomalies

> map[string]interface{} WorldWorldClimateAnomalies(ctx).Execute()

Climate anomaly feed

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
	resp, r, err := apiClient.GeoAPI.WorldWorldClimateAnomalies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldClimateAnomalies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldClimateAnomalies`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldClimateAnomalies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldClimateAnomaliesRequest struct via the builder pattern


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


## WorldWorldCloudflareOutages

> map[string]interface{} WorldWorldCloudflareOutages(ctx).Execute()

Cloudflare internet outages

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
	resp, r, err := apiClient.GeoAPI.WorldWorldCloudflareOutages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldCloudflareOutages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldCloudflareOutages`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldCloudflareOutages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldCloudflareOutagesRequest struct via the builder pattern


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


## WorldWorldEarthquakes

> map[string]interface{} WorldWorldEarthquakes(ctx).Execute()

USGS earthquake feed

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
	resp, r, err := apiClient.GeoAPI.WorldWorldEarthquakes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldEarthquakes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldEarthquakes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldEarthquakes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldEarthquakesRequest struct via the builder pattern


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


## WorldWorldFaaStatus

> map[string]interface{} WorldWorldFaaStatus(ctx).Execute()

FAA airport status

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
	resp, r, err := apiClient.GeoAPI.WorldWorldFaaStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldFaaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldFaaStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldFaaStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldFaaStatusRequest struct via the builder pattern


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


## WorldWorldFirmsFires

> map[string]interface{} WorldWorldFirmsFires(ctx).Execute()

NASA FIRMS active fires (requires NASA_FIRMS_API_KEY)

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
	resp, r, err := apiClient.GeoAPI.WorldWorldFirmsFires(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldFirmsFires``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldFirmsFires`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldFirmsFires`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldFirmsFiresRequest struct via the builder pattern


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


## WorldWorldNgaWarnings

> map[string]interface{} WorldWorldNgaWarnings(ctx).Execute()

NGA maritime safety warnings

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
	resp, r, err := apiClient.GeoAPI.WorldWorldNgaWarnings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldNgaWarnings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldNgaWarnings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldNgaWarnings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldNgaWarningsRequest struct via the builder pattern


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


## WorldWorldOpensky

> map[string]interface{} WorldWorldOpensky(ctx).Execute()

OpenSky flight states

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
	resp, r, err := apiClient.GeoAPI.WorldWorldOpensky(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldOpensky``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldOpensky`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldOpensky`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldOpenskyRequest struct via the builder pattern


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


## WorldWorldWingbits

> map[string]interface{} WorldWorldWingbits(ctx).Execute()

Wingbits ADS-B (requires WINGBITS_API_KEY)

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
	resp, r, err := apiClient.GeoAPI.WorldWorldWingbits(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldWingbits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldWingbits`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldWingbits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldWingbitsRequest struct via the builder pattern


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


## WorldWorldWorldpopExposure

> map[string]interface{} WorldWorldWorldpopExposure(ctx).Area(area).Execute()

Population exposure model

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
	area := "area_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GeoAPI.WorldWorldWorldpopExposure(context.Background()).Area(area).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GeoAPI.WorldWorldWorldpopExposure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorldWorldWorldpopExposure`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `GeoAPI.WorldWorldWorldpopExposure`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorldWorldWorldpopExposureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **area** | **string** |  | 

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

