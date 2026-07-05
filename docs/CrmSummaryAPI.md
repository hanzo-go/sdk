# \CrmSummaryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmCrmSummary**](CrmSummaryAPI.md#CrmCrmSummary) | **Get** /v1/crm/summary | Per-org row counts (companies / contacts / opportunities)



## CrmCrmSummary

> CrmSummary CrmCrmSummary(ctx).Execute()

Per-org row counts (companies / contacts / opportunities)

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
	resp, r, err := apiClient.CrmSummaryAPI.CrmCrmSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CrmSummaryAPI.CrmCrmSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CrmCrmSummary`: CrmSummary
	fmt.Fprintf(os.Stdout, "Response from `CrmSummaryAPI.CrmCrmSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCrmCrmSummaryRequest struct via the builder pattern


### Return type

[**CrmSummary**](CrmSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

