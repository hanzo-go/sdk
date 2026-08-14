# \ValidatorsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidators**](ValidatorsAPI.md#GetValidators) | **Get** /v1/validators | Returns the validator slots the caller&#39;s org has claimed.
[**GetValidatorsByTokenid**](ValidatorsAPI.md#GetValidatorsByTokenid) | **Get** /v1/validators/{tokenId} | Returns one claimed validator slot, scoped to the caller&#39;s org.
[**GetValidatorsChallenge**](ValidatorsAPI.md#GetValidatorsChallenge) | **Get** /v1/validators/challenge | Issues the single-use nonce and the exact message a wallet must sign to claim a validator slot.
[**PostValidators**](ValidatorsAPI.md#PostValidators) | **Post** /v1/validators | Claims a validator slot and provisions its node, after proving the caller&#39;s wallet owns the slot&#39;s NFT.



## GetValidators

> ValidatorList GetValidators(ctx).Limit(limit).Execute()

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
	resp, r, err := apiClient.ValidatorsAPI.GetValidators(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.GetValidators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidators`: ValidatorList
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.GetValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **string** | Limit is how many slots to return, as a decimal string in the &#x60;?limit&#x3D;&#x60; query. Absent, unparseable or non-positive means 200; over 1000 is clamped to 1000. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**ValidatorList**](ValidatorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidatorsByTokenid

> SlotView GetValidatorsByTokenid(ctx, tokenId).Execute()

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
	resp, r, err := apiClient.ValidatorsAPI.GetValidatorsByTokenid(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.GetValidatorsByTokenid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidatorsByTokenid`: SlotView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.GetValidatorsByTokenid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | TokenID is the slot&#39;s GenesisNFT token id, from the path, as a decimal string. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorsByTokenidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlotView**](SlotView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidatorsChallenge

> ChallengeView GetValidatorsChallenge(ctx).TokenId(tokenId).Execute()

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
	resp, r, err := apiClient.ValidatorsAPI.GetValidatorsChallenge(context.Background()).TokenId(tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.GetValidatorsChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidatorsChallenge`: ChallengeView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.GetValidatorsChallenge`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidatorsChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenId** | **string** | TokenID is the Validator-tier GenesisNFT token id, as a decimal string in the &#x60;?tokenId&#x3D;&#x60; query. A value that is not a positive integer is 400. It is a string rather than a number because the parse that has always served this route trims surrounding whitespace, and one parse rule is better than two. | 

### Return type

[**ChallengeView**](ChallengeView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostValidators

> SlotView PostValidators(ctx).ValidatorClaim(validatorClaim).Execute()

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
	validatorClaim := *openapiclient.NewValidatorClaim() // ValidatorClaim | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatorsAPI.PostValidators(context.Background()).ValidatorClaim(validatorClaim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatorsAPI.PostValidators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostValidators`: SlotView
	fmt.Fprintf(os.Stdout, "Response from `ValidatorsAPI.PostValidators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostValidatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatorClaim** | [**ValidatorClaim**](ValidatorClaim.md) |  | 

### Return type

[**SlotView**](SlotView.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

