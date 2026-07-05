# \DidIdentityResolutionAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DidLinkIdentity**](DidIdentityResolutionAPI.md#DidLinkIdentity) | **Post** /v1/did/profiles/{profile_id}/identities | Link an external identity
[**DidListLinkedIdentities**](DidIdentityResolutionAPI.md#DidListLinkedIdentities) | **Get** /v1/did/profiles/{profile_id}/identities | List linked identities
[**DidResolveIdentity**](DidIdentityResolutionAPI.md#DidResolveIdentity) | **Get** /v1/did/resolve | Resolve identity across providers
[**DidUnlinkIdentity**](DidIdentityResolutionAPI.md#DidUnlinkIdentity) | **Delete** /v1/did/profiles/{profile_id}/identities/{provider} | Unlink an external identity



## DidLinkIdentity

> DidLinkedIdentity DidLinkIdentity(ctx, profileId).DidLinkIdentityRequest(didLinkIdentityRequest).Execute()

Link an external identity

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
	profileId := "profileId_example" // string | 
	didLinkIdentityRequest := *openapiclient.NewDidLinkIdentityRequest("Provider_example", "ExternalId_example") // DidLinkIdentityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidIdentityResolutionAPI.DidLinkIdentity(context.Background(), profileId).DidLinkIdentityRequest(didLinkIdentityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidIdentityResolutionAPI.DidLinkIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidLinkIdentity`: DidLinkedIdentity
	fmt.Fprintf(os.Stdout, "Response from `DidIdentityResolutionAPI.DidLinkIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidLinkIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **didLinkIdentityRequest** | [**DidLinkIdentityRequest**](DidLinkIdentityRequest.md) |  | 

### Return type

[**DidLinkedIdentity**](DidLinkedIdentity.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidListLinkedIdentities

> DidListLinkedIdentities200Response DidListLinkedIdentities(ctx, profileId).Execute()

List linked identities

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
	profileId := "profileId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidIdentityResolutionAPI.DidListLinkedIdentities(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidIdentityResolutionAPI.DidListLinkedIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidListLinkedIdentities`: DidListLinkedIdentities200Response
	fmt.Fprintf(os.Stdout, "Response from `DidIdentityResolutionAPI.DidListLinkedIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidListLinkedIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DidListLinkedIdentities200Response**](DidListLinkedIdentities200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidResolveIdentity

> DidProfile DidResolveIdentity(ctx).Provider(provider).ExternalId(externalId).Execute()

Resolve identity across providers



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
	provider := "provider_example" // string | 
	externalId := "externalId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DidIdentityResolutionAPI.DidResolveIdentity(context.Background()).Provider(provider).ExternalId(externalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidIdentityResolutionAPI.DidResolveIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DidResolveIdentity`: DidProfile
	fmt.Fprintf(os.Stdout, "Response from `DidIdentityResolutionAPI.DidResolveIdentity`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDidResolveIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** |  | 
 **externalId** | **string** |  | 

### Return type

[**DidProfile**](DidProfile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DidUnlinkIdentity

> DidUnlinkIdentity(ctx, profileId, provider).Execute()

Unlink an external identity

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
	profileId := "profileId_example" // string | 
	provider := "provider_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DidIdentityResolutionAPI.DidUnlinkIdentity(context.Background(), profileId, provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DidIdentityResolutionAPI.DidUnlinkIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** |  | 
**provider** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDidUnlinkIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

