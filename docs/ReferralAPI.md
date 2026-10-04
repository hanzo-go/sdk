# \ReferralAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReferral**](ReferralAPI.md#GetReferral) | **Get** /v1/referral | Returns the caller&#39;s referral code, share link and the referrals they have made.
[**PostReferralClaim**](ReferralAPI.md#PostReferralClaim) | **Post** /v1/referral/claim | Records that the caller&#39;s org signed up through a referral code.



## GetReferral

> ReferralMyReferrals GetReferral(ctx).Execute()

Returns the caller's referral code, share link and the referrals they have made.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferralAPI.GetReferral(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralAPI.GetReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferral`: ReferralMyReferrals
	fmt.Fprintf(os.Stdout, "Response from `ReferralAPI.GetReferral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralRequest struct via the builder pattern


### Return type

[**ReferralMyReferrals**](ReferralMyReferrals.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReferralClaim

> ReferralClaimView PostReferralClaim(ctx).ReferralClaimRequest(referralClaimRequest).Execute()

Records that the caller's org signed up through a referral code.



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
	referralClaimRequest := *openapiclient.NewReferralClaimRequest() // ReferralClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferralAPI.PostReferralClaim(context.Background()).ReferralClaimRequest(referralClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralAPI.PostReferralClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReferralClaim`: ReferralClaimView
	fmt.Fprintf(os.Stdout, "Response from `ReferralAPI.PostReferralClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReferralClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referralClaimRequest** | [**ReferralClaimRequest**](ReferralClaimRequest.md) |  | 

### Return type

[**ReferralClaimView**](ReferralClaimView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

