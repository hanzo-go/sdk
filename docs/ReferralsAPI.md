# \ReferralsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Referrals**](ReferralsAPI.md#CloudGetV1Referrals) | **Get** /v1/referrals | Returns the caller&#39;s referral code, share link and the referrals they have made.
[**CloudPostV1ReferralsClaim**](ReferralsAPI.md#CloudPostV1ReferralsClaim) | **Post** /v1/referrals/claim | Records that the caller&#39;s org signed up through a referral code.



## CloudGetV1Referrals

> CloudMyReferrals CloudGetV1Referrals(ctx).Execute()

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
	resp, r, err := apiClient.ReferralsAPI.CloudGetV1Referrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.CloudGetV1Referrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Referrals`: CloudMyReferrals
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.CloudGetV1Referrals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1ReferralsRequest struct via the builder pattern


### Return type

[**CloudMyReferrals**](CloudMyReferrals.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1ReferralsClaim

> CloudClaimView CloudPostV1ReferralsClaim(ctx).CloudClaimRequest(cloudClaimRequest).Execute()

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
	cloudClaimRequest := *openapiclient.NewCloudClaimRequest() // CloudClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferralsAPI.CloudPostV1ReferralsClaim(context.Background()).CloudClaimRequest(cloudClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.CloudPostV1ReferralsClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1ReferralsClaim`: CloudClaimView
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.CloudPostV1ReferralsClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1ReferralsClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudClaimRequest** | [**CloudClaimRequest**](CloudClaimRequest.md) |  | 

### Return type

[**CloudClaimView**](CloudClaimView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

