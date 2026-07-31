# \SSHKeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GitDeleteGitKey**](SSHKeysAPI.md#GitDeleteGitKey) | **Delete** /v1/git/keys/{id} | Remove an SSH key
[**GitListGitKeys**](SSHKeysAPI.md#GitListGitKeys) | **Get** /v1/git/keys | List the tenant&#39;s SSH keys
[**GitRegisterGitKey**](SSHKeysAPI.md#GitRegisterGitKey) | **Post** /v1/git/keys | Register an SSH public key



## GitDeleteGitKey

> GitDeleteGitKey(ctx, id).Execute()

Remove an SSH key

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
	id := "id_example" // string | Key id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SSHKeysAPI.GitDeleteGitKey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GitDeleteGitKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Key id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGitDeleteGitKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitListGitKeys

> GitListGitKeys200Response GitListGitKeys(ctx).Execute()

List the tenant's SSH keys

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
	resp, r, err := apiClient.SSHKeysAPI.GitListGitKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GitListGitKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitListGitKeys`: GitListGitKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.GitListGitKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGitListGitKeysRequest struct via the builder pattern


### Return type

[**GitListGitKeys200Response**](GitListGitKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GitRegisterGitKey

> GitKey GitRegisterGitKey(ctx).GitRegisterKey(gitRegisterKey).Execute()

Register an SSH public key



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
	gitRegisterKey := *openapiclient.NewGitRegisterKey("ssh-ed25519 AAAAC3Nza... user@host") // GitRegisterKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSHKeysAPI.GitRegisterGitKey(context.Background()).GitRegisterKey(gitRegisterKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GitRegisterGitKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GitRegisterGitKey`: GitKey
	fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.GitRegisterGitKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGitRegisterGitKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitRegisterKey** | [**GitRegisterKey**](GitRegisterKey.md) |  | 

### Return type

[**GitKey**](GitKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

