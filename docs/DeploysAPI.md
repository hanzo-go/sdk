# \DeploysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorsRecordDeploy**](DeploysAPI.md#AuthorsRecordDeploy) | **Post** /v1/authors/deploys/record | Record a deploy



## AuthorsRecordDeploy

> AuthorsRecordDeploy200Response AuthorsRecordDeploy(ctx).AuthorsRecordDeployRequest(authorsRecordDeployRequest).Execute()

Record a deploy



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
	authorsRecordDeployRequest := *openapiclient.NewAuthorsRecordDeployRequest("Project_example") // AuthorsRecordDeployRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploysAPI.AuthorsRecordDeploy(context.Background()).AuthorsRecordDeployRequest(authorsRecordDeployRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploysAPI.AuthorsRecordDeploy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsRecordDeploy`: AuthorsRecordDeploy200Response
	fmt.Fprintf(os.Stdout, "Response from `DeploysAPI.AuthorsRecordDeploy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsRecordDeployRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorsRecordDeployRequest** | [**AuthorsRecordDeployRequest**](AuthorsRecordDeployRequest.md) |  | 

### Return type

[**AuthorsRecordDeploy200Response**](AuthorsRecordDeploy200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

