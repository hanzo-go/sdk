# \SecurityRulesAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritySecurityHealth**](SecurityRulesAPI.md#SecuritySecurityHealth) | **Get** /v1/security/health | Health check
[**SecuritySecurityListRules**](SecurityRulesAPI.md#SecuritySecurityListRules) | **Get** /v1/security/rules | List the detection ruleset



## SecuritySecurityHealth

> SecuritySecurityHealth200Response SecuritySecurityHealth(ctx).Execute()

Health check

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
	resp, r, err := apiClient.SecurityRulesAPI.SecuritySecurityHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityRulesAPI.SecuritySecurityHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityHealth`: SecuritySecurityHealth200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityRulesAPI.SecuritySecurityHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityHealthRequest struct via the builder pattern


### Return type

[**SecuritySecurityHealth200Response**](SecuritySecurityHealth200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SecuritySecurityListRules

> SecuritySecurityListRules200Response SecuritySecurityListRules(ctx).Execute()

List the detection ruleset

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
	resp, r, err := apiClient.SecurityRulesAPI.SecuritySecurityListRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityRulesAPI.SecuritySecurityListRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SecuritySecurityListRules`: SecuritySecurityListRules200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityRulesAPI.SecuritySecurityListRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritySecurityListRulesRequest struct via the builder pattern


### Return type

[**SecuritySecurityListRules200Response**](SecuritySecurityListRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

