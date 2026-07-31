# \PiecesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AutomationsListPieces**](PiecesAPI.md#AutomationsListPieces) | **Get** /v1/automations/pieces | List the connector/piece catalogue



## AutomationsListPieces

> AutomationsCatalog AutomationsListPieces(ctx).Execute()

List the connector/piece catalogue

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
	resp, r, err := apiClient.PiecesAPI.AutomationsListPieces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PiecesAPI.AutomationsListPieces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomationsListPieces`: AutomationsCatalog
	fmt.Fprintf(os.Stdout, "Response from `PiecesAPI.AutomationsListPieces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAutomationsListPiecesRequest struct via the builder pattern


### Return type

[**AutomationsCatalog**](AutomationsCatalog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

