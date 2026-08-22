# \TreasuryAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTreasury**](TreasuryAPI.md#GetTreasury) | **Get** /v1/treasury | Returns the reserve fund&#39;s health and the current revenue-share policy for any validated caller.
[**GetTreasuryAccounts**](TreasuryAPI.md#GetTreasuryAccounts) | **Get** /v1/treasury/accounts | Returns the ledger accounts the caller may see, with their balances.



## GetTreasury

> TreasuryReport GetTreasury(ctx).Execute()

Returns the reserve fund's health and the current revenue-share policy for any validated caller.



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
	resp, r, err := apiClient.TreasuryAPI.GetTreasury(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.GetTreasury``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTreasury`: TreasuryReport
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.GetTreasury`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTreasuryRequest struct via the builder pattern


### Return type

[**TreasuryReport**](TreasuryReport.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTreasuryAccounts

> AccountsOut GetTreasuryAccounts(ctx).Scope(scope).Org(org).Execute()

Returns the ledger accounts the caller may see, with their balances.



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
	scope := "scope_example" // string | Scope is \"house\" to read the reserve/revenue/payout house accounts. SuperAdmin only. (optional)
	org := "org_example" // string | Org names another tenant to read. SuperAdmin only; ignored when scope=house. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TreasuryAPI.GetTreasuryAccounts(context.Background()).Scope(scope).Org(org).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TreasuryAPI.GetTreasuryAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTreasuryAccounts`: AccountsOut
	fmt.Fprintf(os.Stdout, "Response from `TreasuryAPI.GetTreasuryAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTreasuryAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Scope is \&quot;house\&quot; to read the reserve/revenue/payout house accounts. SuperAdmin only. | 
 **org** | **string** | Org names another tenant to read. SuperAdmin only; ignored when scope&#x3D;house. | 

### Return type

[**AccountsOut**](AccountsOut.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

