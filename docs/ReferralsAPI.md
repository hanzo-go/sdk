# \ReferralsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReferrals**](ReferralsAPI.md#GetReferrals) | **Get** /v1/referrals | Returns the caller&#39;s referral code, share link and the referrals they have made.
[**PostReferralsClaim**](ReferralsAPI.md#PostReferralsClaim) | **Post** /v1/referrals/claim | Records that the caller&#39;s org signed up through a referral code.



## GetReferrals

> MyReferrals GetReferrals(ctx).Execute()

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
	resp, r, err := apiClient.ReferralsAPI.GetReferrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.GetReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferrals`: MyReferrals
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.GetReferrals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsRequest struct via the builder pattern


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


## PostReferralsClaim

> ClaimView PostReferralsClaim(ctx).ClaimRequest(claimRequest).Execute()

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
	resp, r, err := apiClient.ReferralsAPI.PostReferralsClaim(context.Background()).ClaimRequest(claimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.PostReferralsClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostReferralsClaim`: ClaimView
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.PostReferralsClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReferralsClaimRequest struct via the builder pattern


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

