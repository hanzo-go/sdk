# \CloudAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCloudByProviderAccountsByLabel**](CloudAPI.md#DeleteCloudByProviderAccountsByLabel) | **Delete** /v1/cloud/{provider}/accounts/{label} | Forgets one linked cloud account: it detaches every fleet cluster THIS account folded (its own names, in its own shard — a neighbour&#39;s cluster of the same name is untouched), deletes the sealed credential, and drops the index row.
[**GetCloud**](CloudAPI.md#GetCloud) | **Get** /v1/cloud | Returns the clouds this deployment can link and what linking each one needs — the DigitalOcean token, the AWS role and external id, the GCP credential JSON, the Azure app — plus whether the provider can be linked without storing any long-lived secret.
[**GetCloudAccounts**](CloudAPI.md#GetCloudAccounts) | **Get** /v1/cloud/accounts | Lists the caller org&#39;s linked cloud accounts across every provider: which account each one is at the provider, which fleet clusters it folded, and when it was last discovered.
[**PostCloudByProviderAccounts**](CloudAPI.md#PostCloudByProviderAccounts) | **Post** /v1/cloud/{provider}/accounts | Links one of the caller org&#39;s cloud accounts and folds the Kubernetes clusters it finds there into the ONE Hanzo fleet, so they appear at /v1/clusters and can run work like any managed or bring-your-own cluster.
[**PostCloudByProviderAccountsByLabelSync**](CloudAPI.md#PostCloudByProviderAccountsByLabelSync) | **Post** /v1/cloud/{provider}/accounts/{label}/sync | Re-discovers one already-linked cloud account and reconciles what it folded: kubeconfigs are refreshed, clusters that appeared since the last sync are folded, and clusters this account folded that the provider no longer returns are detached — only this account&#39;s own, in the fleet shard it was linked into.



## DeleteCloudByProviderAccountsByLabel

> UnlinkedView DeleteCloudByProviderAccountsByLabel(ctx, provider, label).Execute()

Forgets one linked cloud account: it detaches every fleet cluster THIS account folded (its own names, in its own shard — a neighbour's cluster of the same name is untouched), deletes the sealed credential, and drops the index row.



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
	provider := "provider_example" // string | Provider is the cloud the account belongs to: digitalocean, aws, gcp or azure. An unknown provider is not found.
	label := "label_example" // string | Label is the org-chosen name of the account within that provider. Empty means \"default\"; anything outside 1–64 of [A-Za-z0-9._-] is refused.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.DeleteCloudByProviderAccountsByLabel(context.Background(), provider, label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.DeleteCloudByProviderAccountsByLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCloudByProviderAccountsByLabel`: UnlinkedView
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.DeleteCloudByProviderAccountsByLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the cloud the account belongs to: digitalocean, aws, gcp or azure. An unknown provider is not found. | 
**label** | **string** | Label is the org-chosen name of the account within that provider. Empty means \&quot;default\&quot;; anything outside 1–64 of [A-Za-z0-9._-] is refused. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudByProviderAccountsByLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UnlinkedView**](UnlinkedView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloud

> ProvidersView GetCloud(ctx).Execute()

Returns the clouds this deployment can link and what linking each one needs — the DigitalOcean token, the AWS role and external id, the GCP credential JSON, the Azure app — plus whether the provider can be linked without storing any long-lived secret.



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
	resp, r, err := apiClient.CloudAPI.GetCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloud`: ProvidersView
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloud`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudRequest struct via the builder pattern


### Return type

[**ProvidersView**](ProvidersView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudAccounts

> CloudAccountsView GetCloudAccounts(ctx).Execute()

Lists the caller org's linked cloud accounts across every provider: which account each one is at the provider, which fleet clusters it folded, and when it was last discovered.



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
	resp, r, err := apiClient.CloudAPI.GetCloudAccounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.GetCloudAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCloudAccounts`: CloudAccountsView
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.GetCloudAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudAccountsRequest struct via the builder pattern


### Return type

[**CloudAccountsView**](CloudAccountsView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudByProviderAccounts

> AccountFoldView PostCloudByProviderAccounts(ctx, provider).VenueLinkRequest(venueLinkRequest).Execute()

Links one of the caller org's cloud accounts and folds the Kubernetes clusters it finds there into the ONE Hanzo fleet, so they appear at /v1/clusters and can run work like any managed or bring-your-own cluster.



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
	provider := "provider_example" // string | Provider is the cloud being linked, from the path: digitalocean, aws, gcp or azure.
	venueLinkRequest := *openapiclient.NewVenueLinkRequest() // VenueLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.PostCloudByProviderAccounts(context.Background(), provider).VenueLinkRequest(venueLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PostCloudByProviderAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudByProviderAccounts`: AccountFoldView
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PostCloudByProviderAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the cloud being linked, from the path: digitalocean, aws, gcp or azure. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudByProviderAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **venueLinkRequest** | [**VenueLinkRequest**](VenueLinkRequest.md) |  | 

### Return type

[**AccountFoldView**](AccountFoldView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloudByProviderAccountsByLabelSync

> AccountFoldView PostCloudByProviderAccountsByLabelSync(ctx, provider, label).Execute()

Re-discovers one already-linked cloud account and reconciles what it folded: kubeconfigs are refreshed, clusters that appeared since the last sync are folded, and clusters this account folded that the provider no longer returns are detached — only this account's own, in the fleet shard it was linked into.



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
	provider := "provider_example" // string | Provider is the cloud the account belongs to: digitalocean, aws, gcp or azure. An unknown provider is not found.
	label := "label_example" // string | Label is the org-chosen name of the account within that provider. Empty means \"default\"; anything outside 1–64 of [A-Za-z0-9._-] is refused.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudAPI.PostCloudByProviderAccountsByLabelSync(context.Background(), provider, label).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudAPI.PostCloudByProviderAccountsByLabelSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloudByProviderAccountsByLabelSync`: AccountFoldView
	fmt.Fprintf(os.Stdout, "Response from `CloudAPI.PostCloudByProviderAccountsByLabelSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider is the cloud the account belongs to: digitalocean, aws, gcp or azure. An unknown provider is not found. | 
**label** | **string** | Label is the org-chosen name of the account within that provider. Empty means \&quot;default\&quot;; anything outside 1–64 of [A-Za-z0-9._-] is refused. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloudByProviderAccountsByLabelSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AccountFoldView**](AccountFoldView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

