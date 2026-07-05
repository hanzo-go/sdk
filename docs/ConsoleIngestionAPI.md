# \ConsoleIngestionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConsoleIngestOtelMetrics**](ConsoleIngestionAPI.md#ConsoleIngestOtelMetrics) | **Post** /v1/console/otel/v1/metrics | Ingest OpenTelemetry metrics
[**ConsoleIngestOtelTraces**](ConsoleIngestionAPI.md#ConsoleIngestOtelTraces) | **Post** /v1/console/otel/v1/traces | Ingest OpenTelemetry traces



## ConsoleIngestOtelMetrics

> map[string]interface{} ConsoleIngestOtelMetrics(ctx).Body(body).Execute()

Ingest OpenTelemetry metrics

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleIngestionAPI.ConsoleIngestOtelMetrics(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleIngestionAPI.ConsoleIngestOtelMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleIngestOtelMetrics`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConsoleIngestionAPI.ConsoleIngestOtelMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleIngestOtelMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsoleIngestOtelTraces

> map[string]interface{} ConsoleIngestOtelTraces(ctx).Body(body).Execute()

Ingest OpenTelemetry traces



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConsoleIngestionAPI.ConsoleIngestOtelTraces(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConsoleIngestionAPI.ConsoleIngestOtelTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConsoleIngestOtelTraces`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ConsoleIngestionAPI.ConsoleIngestOtelTraces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConsoleIngestOtelTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

