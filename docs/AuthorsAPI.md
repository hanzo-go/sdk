# \AuthorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorsConnectAuthor**](AuthorsAPI.md#AuthorsConnectAuthor) | **Post** /v1/authors/connect | Connect GitHub
[**AuthorsGetMyAuthors**](AuthorsAPI.md#AuthorsGetMyAuthors) | **Get** /v1/authors | Get my author program
[**AuthorsVerifyRepo**](AuthorsAPI.md#AuthorsVerifyRepo) | **Post** /v1/authors/repos/verify | Verify a repo



## AuthorsConnectAuthor

> AuthorsConnectResponse AuthorsConnectAuthor(ctx).AuthorsConnectRequest(authorsConnectRequest).Execute()

Connect GitHub



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
	authorsConnectRequest := *openapiclient.NewAuthorsConnectRequest() // AuthorsConnectRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.AuthorsConnectAuthor(context.Background()).AuthorsConnectRequest(authorsConnectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.AuthorsConnectAuthor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsConnectAuthor`: AuthorsConnectResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.AuthorsConnectAuthor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsConnectAuthorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorsConnectRequest** | [**AuthorsConnectRequest**](AuthorsConnectRequest.md) |  | 

### Return type

[**AuthorsConnectResponse**](AuthorsConnectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsGetMyAuthors

> AuthorsGetMyAuthors200Response AuthorsGetMyAuthors(ctx).Execute()

Get my author program



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
	resp, r, err := apiClient.AuthorsAPI.AuthorsGetMyAuthors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.AuthorsGetMyAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsGetMyAuthors`: AuthorsGetMyAuthors200Response
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.AuthorsGetMyAuthors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsGetMyAuthorsRequest struct via the builder pattern


### Return type

[**AuthorsGetMyAuthors200Response**](AuthorsGetMyAuthors200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthorsVerifyRepo

> AuthorsVerifyRepoResponse AuthorsVerifyRepo(ctx).AuthorsVerifyRepoRequest(authorsVerifyRepoRequest).Execute()

Verify a repo



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
	authorsVerifyRepoRequest := *openapiclient.NewAuthorsVerifyRepoRequest("RepoUrl_example") // AuthorsVerifyRepoRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthorsAPI.AuthorsVerifyRepo(context.Background()).AuthorsVerifyRepoRequest(authorsVerifyRepoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthorsAPI.AuthorsVerifyRepo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorsVerifyRepo`: AuthorsVerifyRepoResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthorsAPI.AuthorsVerifyRepo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorsVerifyRepoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorsVerifyRepoRequest** | [**AuthorsVerifyRepoRequest**](AuthorsVerifyRepoRequest.md) |  | 

### Return type

[**AuthorsVerifyRepoResponse**](AuthorsVerifyRepoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

