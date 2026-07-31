# \WalletsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudGetV1Wallets**](WalletsAPI.md#CloudGetV1Wallets) | **Get** /v1/wallets | Returns the caller org&#39;s wallets, newest first, optionally NARROWED within the org by project, agent or account.
[**CloudGetV1WalletsAccounts**](WalletsAPI.md#CloudGetV1WalletsAccounts) | **Get** /v1/wallets/accounts | Returns the caller org&#39;s wallet accounts, newest first.
[**CloudGetV1WalletsId**](WalletsAPI.md#CloudGetV1WalletsId) | **Get** /v1/wallets/{id} | Returns one of the caller org&#39;s wallets: its scope, custody kind, tier, chain and on-chain address.
[**CloudPostV1Wallets**](WalletsAPI.md#CloudPostV1Wallets) | **Post** /v1/wallets | Provisions a new signing identity under one of the caller org&#39;s accounts and answers the stored wallet including its on-chain address.
[**CloudPostV1WalletsAccounts**](WalletsAPI.md#CloudPostV1WalletsAccounts) | **Post** /v1/wallets/accounts | Opens a named wallet account for the caller&#39;s org.
[**CloudPostV1WalletsIdKeys**](WalletsAPI.md#CloudPostV1WalletsIdKeys) | **Post** /v1/wallets/{id}/keys | Rolls one wallet&#39;s signing material through its own custody backend and answers the wallet with whatever address that produced.
[**CloudPostV1WalletsIdSafeTx**](WalletsAPI.md#CloudPostV1WalletsIdSafeTx) | **Post** /v1/wallets/{id}/safe-tx | Composes a Safe transaction on the MPC ring and answers its EIP-712 hash together with the owner approval the ring&#39;s threshold signature produced.
[**CloudPostV1WalletsIdSign**](WalletsAPI.md#CloudPostV1WalletsIdSign) | **Post** /v1/wallets/{id}/sign | Produces a secp256k1 signature from one of the caller org&#39;s wallets over a 32-byte digest, through whichever custody backend that wallet uses.



## CloudGetV1Wallets

> CloudWalletList CloudGetV1Wallets(ctx).Project(project).Agent(agent).Account(account).Execute()

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
	resp, r, err := apiClient.WalletsAPI.CloudGetV1Wallets(context.Background()).Project(project).Agent(agent).Account(account).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudGetV1Wallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Wallets`: CloudWalletList
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudGetV1Wallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Project narrows to wallets scoped to one project. Must be a url-safe segment. | 
 **agent** | **string** | Agent narrows to wallets scoped to one agent. Must be a url-safe segment. | 
 **account** | **string** | Account narrows to wallets under one account id. Must be a url-safe segment. | 

### Return type

[**CloudWalletList**](CloudWalletList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WalletsAccounts

> CloudAccountList CloudGetV1WalletsAccounts(ctx).Execute()

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
	resp, r, err := apiClient.WalletsAPI.CloudGetV1WalletsAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudGetV1WalletsAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WalletsAccounts`: CloudAccountList
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudGetV1WalletsAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WalletsAccountsRequest struct via the builder pattern


### Return type

[**CloudAccountList**](CloudAccountList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1WalletsId

> CloudWallet CloudGetV1WalletsId(ctx, id).Execute()

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
	resp, r, err := apiClient.WalletsAPI.CloudGetV1WalletsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudGetV1WalletsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1WalletsId`: CloudWallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudGetV1WalletsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1WalletsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudWallet**](CloudWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Wallets

> CloudWallet CloudPostV1Wallets(ctx).CloudCreateWalletIn(cloudCreateWalletIn).Execute()

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
	cloudCreateWalletIn := *openapiclient.NewCloudCreateWalletIn() // CloudCreateWalletIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CloudPostV1Wallets(context.Background()).CloudCreateWalletIn(cloudCreateWalletIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudPostV1Wallets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Wallets`: CloudWallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudPostV1Wallets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WalletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateWalletIn** | [**CloudCreateWalletIn**](CloudCreateWalletIn.md) |  | 

### Return type

[**CloudWallet**](CloudWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WalletsAccounts

> CloudWalletAccount CloudPostV1WalletsAccounts(ctx).CloudCreateAccountIn(cloudCreateAccountIn).Execute()

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
	cloudCreateAccountIn := *openapiclient.NewCloudCreateAccountIn() // CloudCreateAccountIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CloudPostV1WalletsAccounts(context.Background()).CloudCreateAccountIn(cloudCreateAccountIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudPostV1WalletsAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WalletsAccounts`: CloudWalletAccount
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudPostV1WalletsAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WalletsAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateAccountIn** | [**CloudCreateAccountIn**](CloudCreateAccountIn.md) |  | 

### Return type

[**CloudWalletAccount**](CloudWalletAccount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WalletsIdKeys

> CloudWallet CloudPostV1WalletsIdKeys(ctx, id).Execute()

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
	resp, r, err := apiClient.WalletsAPI.CloudPostV1WalletsIdKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudPostV1WalletsIdKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WalletsIdKeys`: CloudWallet
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudPostV1WalletsIdKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WalletsIdKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudWallet**](CloudWallet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WalletsIdSafeTx

> CloudSafeProposal CloudPostV1WalletsIdSafeTx(ctx, id).CloudSafeTxIn(cloudSafeTxIn).Execute()

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
	cloudSafeTxIn := *openapiclient.NewCloudSafeTxIn() // CloudSafeTxIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CloudPostV1WalletsIdSafeTx(context.Background(), id).CloudSafeTxIn(cloudSafeTxIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudPostV1WalletsIdSafeTx``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WalletsIdSafeTx`: CloudSafeProposal
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudPostV1WalletsIdSafeTx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WalletsIdSafeTxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSafeTxIn** | [**CloudSafeTxIn**](CloudSafeTxIn.md) |  | 

### Return type

[**CloudSafeProposal**](CloudSafeProposal.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1WalletsIdSign

> CloudSignature CloudPostV1WalletsIdSign(ctx, id).CloudSignIn(cloudSignIn).Execute()

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
	cloudSignIn := *openapiclient.NewCloudSignIn() // CloudSignIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletsAPI.CloudPostV1WalletsIdSign(context.Background(), id).CloudSignIn(cloudSignIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletsAPI.CloudPostV1WalletsIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1WalletsIdSign`: CloudSignature
	fmt.Fprintf(os.Stdout, "Response from `WalletsAPI.CloudPostV1WalletsIdSign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1WalletsIdSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudSignIn** | [**CloudSignIn**](CloudSignIn.md) |  | 

### Return type

[**CloudSignature**](CloudSignature.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

