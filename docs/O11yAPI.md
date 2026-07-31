# \O11yAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**O11yO11yDashboards**](O11yAPI.md#O11yO11yDashboards) | **Get** /v1/o11y/dashboards | List dashboards
[**O11yO11yHealth**](O11yAPI.md#O11yO11yHealth) | **Get** /v1/o11y/health | O11y health
[**O11yO11yIngest**](O11yAPI.md#O11yO11yIngest) | **Post** /v1/o11y/ingestion | Ingest LLM-observability events
[**O11yO11yQuery**](O11yAPI.md#O11yO11yQuery) | **Post** /v1/o11y/query | Instant builder query
[**O11yO11yQueryRange**](O11yAPI.md#O11yO11yQueryRange) | **Post** /v1/o11y/query_range | Range builder query
[**O11yO11yRules**](O11yAPI.md#O11yO11yRules) | **Get** /v1/o11y/rules | List alert rules
[**O11yO11yServices**](O11yAPI.md#O11yO11yServices) | **Get** /v1/o11y/services | List traced services
[**O11yO11yVMQuery**](O11yAPI.md#O11yO11yVMQuery) | **Get** /v1/o11y/vm/query | Platform infra health — instant VM query (SuperAdmin)
[**O11yO11yVMQueryRange**](O11yAPI.md#O11yO11yVMQueryRange) | **Get** /v1/o11y/vm/query_range | Platform infra health — range VM query (SuperAdmin)



## O11yO11yDashboards

> []O11yDashboardSummary O11yO11yDashboards(ctx).Execute()

List dashboards



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
	resp, r, err := apiClient.O11yAPI.O11yO11yDashboards(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yDashboards`: []O11yDashboardSummary
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yDashboards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yDashboardsRequest struct via the builder pattern


### Return type

[**[]O11yDashboardSummary**](O11yDashboardSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yHealth

> O11yHealthResponse O11yO11yHealth(ctx).Execute()

O11y health



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
	resp, r, err := apiClient.O11yAPI.O11yO11yHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yHealth`: O11yHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yHealthRequest struct via the builder pattern


### Return type

[**O11yHealthResponse**](O11yHealthResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yIngest

> O11yIngestResult O11yO11yIngest(ctx).O11yIngestBatch(o11yIngestBatch).Execute()

Ingest LLM-observability events



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
	o11yIngestBatch := *openapiclient.NewO11yIngestBatch([]openapiclient.O11yIngestEvent{*openapiclient.NewO11yIngestEvent("Type_example")}) // O11yIngestBatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yIngest(context.Background()).O11yIngestBatch(o11yIngestBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yIngest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yIngest`: O11yIngestResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yIngest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yIngestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yIngestBatch** | [**O11yIngestBatch**](O11yIngestBatch.md) |  | 

### Return type

[**O11yIngestResult**](O11yIngestResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yQuery

> O11yBuilderQueryResult O11yO11yQuery(ctx).O11yBuilderQuery(o11yBuilderQuery).Execute()

Instant builder query



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
	o11yBuilderQuery := *openapiclient.NewO11yBuilderQuery() // O11yBuilderQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yQuery(context.Background()).O11yBuilderQuery(o11yBuilderQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yQuery`: O11yBuilderQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yBuilderQuery** | [**O11yBuilderQuery**](O11yBuilderQuery.md) |  | 

### Return type

[**O11yBuilderQueryResult**](O11yBuilderQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yQueryRange

> O11yBuilderQueryResult O11yO11yQueryRange(ctx).O11yBuilderQuery(o11yBuilderQuery).Execute()

Range builder query



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
	o11yBuilderQuery := *openapiclient.NewO11yBuilderQuery() // O11yBuilderQuery | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yQueryRange(context.Background()).O11yBuilderQuery(o11yBuilderQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yQueryRange`: O11yBuilderQueryResult
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **o11yBuilderQuery** | [**O11yBuilderQuery**](O11yBuilderQuery.md) |  | 

### Return type

[**O11yBuilderQueryResult**](O11yBuilderQueryResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yRules

> []O11yAlertRule O11yO11yRules(ctx).Execute()

List alert rules



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
	resp, r, err := apiClient.O11yAPI.O11yO11yRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yRules`: []O11yAlertRule
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yRulesRequest struct via the builder pattern


### Return type

[**[]O11yAlertRule**](O11yAlertRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yServices

> O11yO11yServices200Response O11yO11yServices(ctx).Execute()

List traced services



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
	resp, r, err := apiClient.O11yAPI.O11yO11yServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yServices`: O11yO11yServices200Response
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yServicesRequest struct via the builder pattern


### Return type

[**O11yO11yServices200Response**](O11yO11yServices200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yVMQuery

> O11yPrometheusResponse O11yO11yVMQuery(ctx).Query(query).Execute()

Platform infra health — instant VM query (SuperAdmin)



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
	query := "query_example" // string | Allowlisted PromQL. Only up, sum(up), and count(up) are permitted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yVMQuery(context.Background()).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yVMQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yVMQuery`: O11yPrometheusResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yVMQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yVMQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Allowlisted PromQL. Only up, sum(up), and count(up) are permitted. | 

### Return type

[**O11yPrometheusResponse**](O11yPrometheusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## O11yO11yVMQueryRange

> O11yPrometheusResponse O11yO11yVMQueryRange(ctx).Query(query).Start(start).End(end).Step(step).Execute()

Platform infra health — range VM query (SuperAdmin)



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
	query := "query_example" // string | Allowlisted PromQL. Only up, sum(up), and count(up) are permitted.
	start := int64(789) // int64 | Range start (Unix seconds, positive integer).
	end := int64(789) // int64 | Range end (Unix seconds, positive integer).
	step := int32(56) // int32 | Step resolution (seconds, positive integer).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.O11yAPI.O11yO11yVMQueryRange(context.Background()).Query(query).Start(start).End(end).Step(step).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `O11yAPI.O11yO11yVMQueryRange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `O11yO11yVMQueryRange`: O11yPrometheusResponse
	fmt.Fprintf(os.Stdout, "Response from `O11yAPI.O11yO11yVMQueryRange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiO11yO11yVMQueryRangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Allowlisted PromQL. Only up, sum(up), and count(up) are permitted. | 
 **start** | **int64** | Range start (Unix seconds, positive integer). | 
 **end** | **int64** | Range end (Unix seconds, positive integer). | 
 **step** | **int32** | Step resolution (seconds, positive integer). | 

### Return type

[**O11yPrometheusResponse**](O11yPrometheusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

