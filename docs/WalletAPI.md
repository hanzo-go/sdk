# \WalletAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWallet**](WalletAPI.md#GetWallet) | **Get** /v1/wallet | Returns the caller org&#39;s wallets, newest first, optionally NARROWED within the org by project, agent or account.
[**GetWalletAccounts**](WalletAPI.md#GetWalletAccounts) | **Get** /v1/wallet/accounts | Returns the caller org&#39;s wallet accounts, newest first.
[**GetWalletById**](WalletAPI.md#GetWalletById) | **Get** /v1/wallet/{id} | Returns one of the caller org&#39;s wallets: its scope, custody kind, tier, chain and on-chain address.
[**PostWallet**](WalletAPI.md#PostWallet) | **Post** /v1/wallet | Provisions a new signing identity under one of the caller org&#39;s accounts and answers the stored wallet including its on-chain address.
[**PostWalletAccounts**](WalletAPI.md#PostWalletAccounts) | **Post** /v1/wallet/accounts | Opens a named wallet account for the caller&#39;s org.
[**PostWalletByIdKeys**](WalletAPI.md#PostWalletByIdKeys) | **Post** /v1/wallet/{id}/keys | Rolls one wallet&#39;s signing material through its own custody backend and answers the wallet with whatever address that produced.
[**PostWalletByIdSign**](WalletAPI.md#PostWalletByIdSign) | **Post** /v1/wallet/{id}/sign | Produces a secp256k1 signature from one of the caller org&#39;s wallets over a 32-byte digest, through whichever custody backend that wallet uses.
[**PostWalletByIdTransactions**](WalletAPI.md#PostWalletByIdTransactions) | **Post** /v1/wallet/{id}/transactions | Composes a Safe transaction on the MPC ring and answers its EIP-712 hash together with the owner approval the ring&#39;s threshold signature produced.



## GetWallet

> WalletList GetWallet(ctx).Project(project).Agent(agent).Account(account).Execute()

Returns the caller org's wallets, newest first, optionally NARROWED within the org by project, agent or account.



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
	project := "project_example" // string | Project narrows to wallets scoped to one project. Must be a url-safe segment. (optional)
	agent := "agent_example" // string | Agent narrows to wallets scoped to one agent. Must be a url-safe segment. (optional)
	account := "acct_9f8c1d" // string | Account narrows to wallets under one account id. Must be a url-safe segment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetWallet(context.Background()).Project(project).Agent(agent).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallet`: WalletList
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to wallets scoped to one project. Must be a url-safe segment. | 
 **agent** | **string** | Agent narrows to wallets scoped to one agent. Must be a url-safe segment. | 
 **account** | **string** | Account narrows to wallets under one account id. Must be a url-safe segment. | 

### Return type

[**WalletList**](WalletList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletAccounts

> AccountList GetWalletAccounts(ctx).Execute()

Returns the caller org's wallet accounts, newest first.



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
	resp, r, err := apiClient.WalletAPI.GetWalletAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetWalletAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletAccounts`: AccountList
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetWalletAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletAccountsRequest struct via the builder pattern


### Return type

[**AccountList**](AccountList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletById

> Wallet GetWalletById(ctx, id).Execute()

Returns one of the caller org's wallets: its scope, custody kind, tier, chain and on-chain address.



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
	id := "wal_4b1e77" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetWalletById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetWalletById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletById`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetWalletById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWallet

> Wallet PostWallet(ctx).CreateWalletIn(createWalletIn).Execute()

Provisions a new signing identity under one of the caller org's accounts and answers the stored wallet including its on-chain address.



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
	createWalletIn := *openapiclient.NewCreateWalletIn() // CreateWalletIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.PostWallet(context.Background()).CreateWalletIn(createWalletIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.PostWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWallet`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.PostWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWalletIn** | [**CreateWalletIn**](CreateWalletIn.md) |  | 

### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletAccounts

> WalletAccount PostWalletAccounts(ctx).CreateAccountIn(createAccountIn).Execute()

Opens a named wallet account for the caller's org.



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
	createAccountIn := *openapiclient.NewCreateAccountIn() // CreateAccountIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.PostWalletAccounts(context.Background()).CreateAccountIn(createAccountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.PostWalletAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletAccounts`: WalletAccount
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.PostWalletAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountIn** | [**CreateAccountIn**](CreateAccountIn.md) |  | 

### Return type

[**WalletAccount**](WalletAccount.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletByIdKeys

> Wallet PostWalletByIdKeys(ctx, id).Execute()

Rolls one wallet's signing material through its own custody backend and answers the wallet with whatever address that produced.



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
	id := "wal_4b1e77" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.PostWalletByIdKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.PostWalletByIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletByIdKeys`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.PostWalletByIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletByIdKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Wallet**](Wallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletByIdSign

> Signature PostWalletByIdSign(ctx, id).SignIn(signIn).Execute()

Produces a secp256k1 signature from one of the caller org's wallets over a 32-byte digest, through whichever custody backend that wallet uses.



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
	id := "wal_4b1e77" // string | 
	signIn := *openapiclient.NewSignIn() // SignIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.PostWalletByIdSign(context.Background(), id).SignIn(signIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.PostWalletByIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletByIdSign`: Signature
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.PostWalletByIdSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletByIdSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signIn** | [**SignIn**](SignIn.md) |  | 

### Return type

[**Signature**](Signature.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWalletByIdTransactions

> SafeProposal PostWalletByIdTransactions(ctx, id).SafeTxIn(safeTxIn).Execute()

Composes a Safe transaction on the MPC ring and answers its EIP-712 hash together with the owner approval the ring's threshold signature produced.



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
	id := "wal_4b1e77" // string | 
	safeTxIn := *openapiclient.NewSafeTxIn() // SafeTxIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.PostWalletByIdTransactions(context.Background(), id).SafeTxIn(safeTxIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.PostWalletByIdTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletByIdTransactions`: SafeProposal
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.PostWalletByIdTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletByIdTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **safeTxIn** | [**SafeTxIn**](SafeTxIn.md) |  | 

### Return type

[**SafeProposal**](SafeProposal.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

