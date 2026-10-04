# \StandingAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostStandingUpkeep**](StandingAPI.md#PostStandingUpkeep) | **Post** /v1/standing/upkeep | Reports what keeping this entity costs every year, itemised.



## PostStandingUpkeep

> StandingUpkeep PostStandingUpkeep(ctx).StandingUpkeepIn(standingUpkeepIn).Execute()

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
	standingUpkeepIn := *openapiclient.NewStandingUpkeepIn() // StandingUpkeepIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandingAPI.PostStandingUpkeep(context.Background()).StandingUpkeepIn(standingUpkeepIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandingAPI.PostStandingUpkeep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostStandingUpkeep`: StandingUpkeep
	fmt.Fprintf(os.Stdout, "Response from `StandingAPI.PostStandingUpkeep`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStandingUpkeepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standingUpkeepIn** | [**StandingUpkeepIn**](StandingUpkeepIn.md) |  | 

### Return type

[**StandingUpkeep**](StandingUpkeep.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

