# \ValidatorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Validators**](ValidatorsAPI.md#CloudGetV1Validators) | **Get** /v1/validators | Returns the validator slots the caller&#39;s org has claimed.
[**CloudGetV1ValidatorsChallenge**](ValidatorsAPI.md#CloudGetV1ValidatorsChallenge) | **Get** /v1/validators/challenge | Issues the single-use nonce and the exact message a wallet must sign to claim a validator slot.
[**CloudGetV1ValidatorsTokenId**](ValidatorsAPI.md#CloudGetV1ValidatorsTokenId) | **Get** /v1/validators/{tokenId} | Returns one claimed validator slot, scoped to the caller&#39;s org.
[**CloudPostV1Validators**](ValidatorsAPI.md#CloudPostV1Validators) | **Post** /v1/validators | Claims a validator slot and provisions its node, after proving the caller&#39;s wallet owns the slot&#39;s NFT.



## CloudGetV1Validators

> CloudValidatorList CloudGetV1Validators(ctx).Limit(limit).Execute()

Returns the validator slots the caller's org has claimed.



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
	limit := "limit_example" // string | Limit is how many slots to return, as a decimal string in the `?limit=` query. Absent, unparseable or non-positive means 200; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorsAPI.CloudGetV1Validators(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.CloudGetV1Validators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Validators`: CloudValidatorList
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.CloudGetV1Validators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Limit is how many slots to return, as a decimal string in the &#x60;?limit&#x3D;&#x60; query. Absent, unparseable or non-positive means 200; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**CloudValidatorList**](CloudValidatorList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ValidatorsChallenge

> CloudChallengeView CloudGetV1ValidatorsChallenge(ctx).TokenId(tokenId).Execute()

Issues the single-use nonce and the exact message a wallet must sign to claim a validator slot.



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
	tokenId := "tokenId_example" // string | TokenID is the Validator-tier GenesisNFT token id, as a decimal string in the `?tokenId=` query. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorsAPI.CloudGetV1ValidatorsChallenge(context.Background()).TokenId(tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.CloudGetV1ValidatorsChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ValidatorsChallenge`: CloudChallengeView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.CloudGetV1ValidatorsChallenge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ValidatorsChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | TokenID is the Validator-tier GenesisNFT token id, as a decimal string in the &#x60;?tokenId&#x3D;&#x60; query. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**CloudChallengeView**](CloudChallengeView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1ValidatorsTokenId

> CloudSlotView CloudGetV1ValidatorsTokenId(ctx, tokenId).Execute()

Returns one claimed validator slot, scoped to the caller's org.



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
	tokenId := "tokenId_example" // string | TokenID is the slot's GenesisNFT token id, from the path, as a decimal string. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorsAPI.CloudGetV1ValidatorsTokenId(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.CloudGetV1ValidatorsTokenId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1ValidatorsTokenId`: CloudSlotView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.CloudGetV1ValidatorsTokenId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | TokenID is the slot&#39;s GenesisNFT token id, from the path, as a decimal string. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ValidatorsTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudSlotView**](CloudSlotView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Validators

> CloudSlotView CloudPostV1Validators(ctx).CloudValidatorClaim(cloudValidatorClaim).Execute()

Claims a validator slot and provisions its node, after proving the caller's wallet owns the slot's NFT.



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
	cloudValidatorClaim := *openapiclient.NewCloudValidatorClaim() // CloudValidatorClaim | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorsAPI.CloudPostV1Validators(context.Background()).CloudValidatorClaim(cloudValidatorClaim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.CloudPostV1Validators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Validators`: CloudSlotView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.CloudPostV1Validators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudValidatorClaim** | [**CloudValidatorClaim**](CloudValidatorClaim.md) |  | 

### Return type

[**CloudSlotView**](CloudSlotView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

