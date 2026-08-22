# \ReferralAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReferral**](ReferralAPI.md#GetReferral) | **Get** /v1/referral | Returns the caller&#39;s referral code, share link and the referrals they have made.
[**PostReferralClaim**](ReferralAPI.md#PostReferralClaim) | **Post** /v1/referral/claim | Records that the caller&#39;s org signed up through a referral code.



## GetReferral

> MyReferrals GetReferral(ctx).Execute()

Returns the caller's referral code, share link and the referrals they have made.



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
	resp, r, err := apiClient.ReferralAPI.GetReferral(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralAPI.GetReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferral`: MyReferrals
	fmt.Fprintf(os.Stdout, "Response from `ReferralAPI.GetReferral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralRequest struct via the builder pattern


### Return type

[**MyReferrals**](MyReferrals.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReferralClaim

> ClaimView PostReferralClaim(ctx).ClaimRequest(claimRequest).Execute()

Records that the caller's org signed up through a referral code.



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
	claimRequest := *openapiclient.NewClaimRequest() // ClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferralAPI.PostReferralClaim(context.Background()).ClaimRequest(claimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralAPI.PostReferralClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReferralClaim`: ClaimView
	fmt.Fprintf(os.Stdout, "Response from `ReferralAPI.PostReferralClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReferralClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimRequest** | [**ClaimRequest**](ClaimRequest.md) |  | 

### Return type

[**ClaimView**](ClaimView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

