# \WalletsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWallets**](WalletsAPI.md#GetWallets) | **Get** /v1/wallets | Returns the caller org&#39;s wallets, newest first, optionally NARROWED within the org by project, agent or account.
[**GetWalletsAccounts**](WalletsAPI.md#GetWalletsAccounts) | **Get** /v1/wallets/accounts | Returns the caller org&#39;s wallet accounts, newest first.
[**GetWalletsById**](WalletsAPI.md#GetWalletsById) | **Get** /v1/wallets/{id} | Returns one of the caller org&#39;s wallets: its scope, custody kind, tier, chain and on-chain address.
[**PostWallets**](WalletsAPI.md#PostWallets) | **Post** /v1/wallets | Provisions a new signing identity under one of the caller org&#39;s accounts and answers the stored wallet including its on-chain address.
[**PostWalletsAccounts**](WalletsAPI.md#PostWalletsAccounts) | **Post** /v1/wallets/accounts | Opens a named wallet account for the caller&#39;s org.
[**PostWalletsByIdKeys**](WalletsAPI.md#PostWalletsByIdKeys) | **Post** /v1/wallets/{id}/keys | Rolls one wallet&#39;s signing material through its own custody backend and answers the wallet with whatever address that produced.
[**PostWalletsByIdSign**](WalletsAPI.md#PostWalletsByIdSign) | **Post** /v1/wallets/{id}/sign | Produces a secp256k1 signature from one of the caller org&#39;s wallets over a 32-byte digest, through whichever custody backend that wallet uses.
[**PostWalletsByIdTransactions**](WalletsAPI.md#PostWalletsByIdTransactions) | **Post** /v1/wallets/{id}/transactions | Composes a Safe transaction on the MPC ring and answers its EIP-712 hash together with the owner approval the ring&#39;s threshold signature produced.



## GetWallets

> WalletList GetWallets(ctx).Project(project).Agent(agent).Account(account).Execute()

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
	resp, r, err := apiClient.WalletsAPI.GetWallets(context.Background()).Project(project).Agent(agent).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWallets`: WalletList
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsRequest struct via the builder pattern


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


## GetWalletsAccounts

> AccountList GetWalletsAccounts(ctx).Execute()

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
	resp, r, err := apiClient.WalletsAPI.GetWalletsAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsAccounts`: AccountList
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsAccountsRequest struct via the builder pattern


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


## GetWalletsById

> Wallet GetWalletsById(ctx, id).Execute()

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
	resp, r, err := apiClient.WalletsAPI.GetWalletsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.GetWalletsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsById`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.GetWalletsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsByIdRequest struct via the builder pattern


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


## PostWallets

> Wallet PostWallets(ctx).CreateWalletIn(createWalletIn).Execute()

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
	resp, r, err := apiClient.WalletsAPI.PostWallets(context.Background()).CreateWalletIn(createWalletIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWallets`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsRequest struct via the builder pattern


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


## PostWalletsAccounts

> WalletAccount PostWalletsAccounts(ctx).CreateAccountIn(createAccountIn).Execute()

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
	resp, r, err := apiClient.WalletsAPI.PostWalletsAccounts(context.Background()).CreateAccountIn(createAccountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsAccounts`: WalletAccount
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsAccountsRequest struct via the builder pattern


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


## PostWalletsByIdKeys

> Wallet PostWalletsByIdKeys(ctx, id).Execute()

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
	resp, r, err := apiClient.WalletsAPI.PostWalletsByIdKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsByIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsByIdKeys`: Wallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsByIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsByIdKeysRequest struct via the builder pattern


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


## PostWalletsByIdSign

> Signature PostWalletsByIdSign(ctx, id).SignIn(signIn).Execute()

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
	resp, r, err := apiClient.WalletsAPI.PostWalletsByIdSign(context.Background(), id).SignIn(signIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsByIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsByIdSign`: Signature
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsByIdSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsByIdSignRequest struct via the builder pattern


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


## PostWalletsByIdTransactions

> SafeProposal PostWalletsByIdTransactions(ctx, id).SafeTxIn(safeTxIn).Execute()

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
	resp, r, err := apiClient.WalletsAPI.PostWalletsByIdTransactions(context.Background(), id).SafeTxIn(safeTxIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.PostWalletsByIdTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostWalletsByIdTransactions`: SafeProposal
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.PostWalletsByIdTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWalletsByIdTransactionsRequest struct via the builder pattern


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

