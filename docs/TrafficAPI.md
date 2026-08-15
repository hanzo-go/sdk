# \TrafficAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrafficGlobe**](TrafficAPI.md#GetTrafficGlobe) | **Get** /v1/traffic/globe | Returns the PUBLIC live request-geo aggregate for the world.hanzo.ai \&quot;Hanzo mode\&quot; globe: WHERE requests to api.hanzo.ai are coming from, as country/region points with per-service-class counts, plus headline throughput rates.



## GetTrafficGlobe

> GetTrafficGlobe(ctx).Execute()

Returns the PUBLIC live request-geo aggregate for the world.hanzo.ai \"Hanzo mode\" globe: WHERE requests to api.hanzo.ai are coming from, as country/region points with per-service-class counts, plus headline throughput rates.



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
	r, err := apiClient.TrafficAPI.GetTrafficGlobe(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrafficAPI.GetTrafficGlobe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrafficGlobeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

