# \ReferralsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReferralsClaimReferral**](ReferralsAPI.md#ReferralsClaimReferral) | **Post** /v1/referrals/claim | Claim a referral from a ?ref code
[**ReferralsGetMyReferrals**](ReferralsAPI.md#ReferralsGetMyReferrals) | **Get** /v1/referrals | Get my referral code, link, and referrals



## ReferralsClaimReferral

> ReferralsClaimResponse ReferralsClaimReferral(ctx).ReferralsClaimRequest(referralsClaimRequest).Execute()

Claim a referral from a ?ref code



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
	referralsClaimRequest := *openapiclient.NewReferralsClaimRequest("Code_example") // ReferralsClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferralsAPI.ReferralsClaimReferral(context.Background()).ReferralsClaimRequest(referralsClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.ReferralsClaimReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferralsClaimReferral`: ReferralsClaimResponse
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.ReferralsClaimReferral`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReferralsClaimReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **referralsClaimRequest** | [**ReferralsClaimRequest**](ReferralsClaimRequest.md) |  | 

### Return type

[**ReferralsClaimResponse**](ReferralsClaimResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferralsGetMyReferrals

> ReferralsMyReferralsResponse ReferralsGetMyReferrals(ctx).Execute()

Get my referral code, link, and referrals



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
	resp, r, err := apiClient.ReferralsAPI.ReferralsGetMyReferrals(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferralsAPI.ReferralsGetMyReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferralsGetMyReferrals`: ReferralsMyReferralsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReferralsAPI.ReferralsGetMyReferrals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferralsGetMyReferralsRequest struct via the builder pattern


### Return type

[**ReferralsMyReferralsResponse**](ReferralsMyReferralsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

