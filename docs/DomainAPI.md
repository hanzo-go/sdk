# \DomainAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDomainAvailability**](DomainAPI.md#GetDomainAvailability) | **Get** /v1/domain/availability | Checks exact names rather than searching for them, and answers the same quote shape search does — purchasable, premium, first-term and renewal price in cents.
[**GetDomainDomains**](DomainAPI.md#GetDomainDomains) | **Get** /v1/domain/domains | Is the domains your org has bought here, newest registration first, each carrying the name, when it was registered, when it expires, what the org paid, the registrar order id and the nameservers it points at.
[**GetDomainHealth**](DomainAPI.md#GetDomainHealth) | **Get** /v1/domain/health | Reports registrar reachability honestly: ok only when the wholesale credentials are present AND name.com accepted them on a live call made while you waited.
[**GetDomainSearch**](DomainAPI.md#GetDomainSearch) | **Get** /v1/domain/search | Finds names built from the keyword q, plus the registrar&#39;s alternate-TLD suggestions, and answers a quote for each: the name, whether it is purchasable, whether it is premium, the first-term and renewal price in cents, and the TLD.
[**PostDomainRegister**](DomainAPI.md#PostDomainRegister) | **Post** /v1/domain/register | Buys a domain for your org and answers the ownership record together with the quote it was bought at.
[**PostDomainRenew**](DomainAPI.md#PostDomainRenew) | **Post** /v1/domain/renew | Extends a domain your org already owns and answers the updated record with its new expiry alongside what was paid.
[**PostDomainTransfer**](DomainAPI.md#PostDomainTransfer) | **Post** /v1/domain/transfer | Moves a domain you own at another registrar onto your org here, using its authCode, and answers the same record-plus-quote a purchase does.



## GetDomainAvailability

> QuoteList GetDomainAvailability(ctx).Domain(domain).Execute()

Checks exact names rather than searching for them, and answers the same quote shape search does — purchasable, premium, first-term and renewal price in cents.



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
	domain := "domain_example" // string | Domain is one name, or several comma-separated, to check in one call. Names are lowercased. It is required.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetDomainAvailability(context.Background()).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainAvailability`: QuoteList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain is one name, or several comma-separated, to check in one call. Names are lowercased. It is required. | 

### Return type

[**QuoteList**](QuoteList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainDomains

> Holdings GetDomainDomains(ctx).Execute()

Is the domains your org has bought here, newest registration first, each carrying the name, when it was registered, when it expires, what the org paid, the registrar order id and the nameservers it points at.



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
	resp, r, err := apiClient.DomainAPI.GetDomainDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainDomains`: Holdings
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainDomainsRequest struct via the builder pattern


### Return type

[**Holdings**](Holdings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainHealth

> Reachability GetDomainHealth(ctx).Execute()

Reports registrar reachability honestly: ok only when the wholesale credentials are present AND name.com accepted them on a live call made while you waited.



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
	resp, r, err := apiClient.DomainAPI.GetDomainHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainHealth`: Reachability
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainHealthRequest struct via the builder pattern


### Return type

[**Reachability**](Reachability.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainSearch

> QuoteList GetDomainSearch(ctx).Q(q).Tld(tld).Execute()

Finds names built from the keyword q, plus the registrar's alternate-TLD suggestions, and answers a quote for each: the name, whether it is purchasable, whether it is premium, the first-term and renewal price in cents, and the TLD.



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
	q := "q_example" // string | Q is the keyword to build names from. It is required.
	tld := "tld_example" // string | TLD narrows the search to a comma-separated set of top-level domains. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.GetDomainSearch(context.Background()).Q(q).Tld(tld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.GetDomainSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainSearch`: QuoteList
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.GetDomainSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Q is the keyword to build names from. It is required. | 
 **tld** | **string** | TLD narrows the search to a comma-separated set of top-level domains. | 

### Return type

[**QuoteList**](QuoteList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDomainRegister

> RegisterResult PostDomainRegister(ctx).Order(order).Execute()

Buys a domain for your org and answers the ownership record together with the quote it was bought at.



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
	order := *openapiclient.NewOrder("Domain_example") // Order | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.PostDomainRegister(context.Background()).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.PostDomainRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDomainRegister`: RegisterResult
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.PostDomainRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDomainRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**Order**](Order.md) |  | 

### Return type

[**RegisterResult**](RegisterResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDomainRenew

> RenewResult PostDomainRenew(ctx).RenewReq(renewReq).Execute()

Extends a domain your org already owns and answers the updated record with its new expiry alongside what was paid.



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
	renewReq := *openapiclient.NewRenewReq("Domain_example") // RenewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.PostDomainRenew(context.Background()).RenewReq(renewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.PostDomainRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDomainRenew`: RenewResult
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.PostDomainRenew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDomainRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renewReq** | [**RenewReq**](RenewReq.md) |  | 

### Return type

[**RenewResult**](RenewResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDomainTransfer

> RegisterResult PostDomainTransfer(ctx).TransferReq(transferReq).Execute()

Moves a domain you own at another registrar onto your org here, using its authCode, and answers the same record-plus-quote a purchase does.



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
	transferReq := *openapiclient.NewTransferReq("AuthCode_example", "Domain_example") // TransferReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DomainAPI.PostDomainTransfer(context.Background()).TransferReq(transferReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainAPI.PostDomainTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDomainTransfer`: RegisterResult
	fmt.Fprintf(os.Stdout, "Response from `DomainAPI.PostDomainTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDomainTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferReq** | [**TransferReq**](TransferReq.md) |  | 

### Return type

[**RegisterResult**](RegisterResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

