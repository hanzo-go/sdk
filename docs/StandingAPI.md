# \StandingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostStandingUpkeep**](StandingAPI.md#PostStandingUpkeep) | **Post** /v1/standing/upkeep | Reports what keeping this entity costs every year, itemised.



## PostStandingUpkeep

> Upkeep PostStandingUpkeep(ctx).UpkeepIn(upkeepIn).Execute()

Reports what keeping this entity costs every year, itemised.



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
	upkeepIn := *openapiclient.NewUpkeepIn() // UpkeepIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandingAPI.PostStandingUpkeep(context.Background()).UpkeepIn(upkeepIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandingAPI.PostStandingUpkeep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStandingUpkeep`: Upkeep
	fmt.Fprintf(os.Stdout, "Response from `StandingAPI.PostStandingUpkeep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStandingUpkeepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upkeepIn** | [**UpkeepIn**](UpkeepIn.md) |  | 

### Return type

[**Upkeep**](Upkeep.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

