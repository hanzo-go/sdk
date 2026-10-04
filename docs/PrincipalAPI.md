# \PrincipalAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrincipal**](PrincipalAPI.md#GetPrincipal) | **Get** /v1/principal | Returns the caller&#39;s org as an economic principal, composed from the apps that own each fact: its legal entity and founders&#39; identity verification (company), the W-9 or W-8 it certified (tax), its wallets (wallet), the agents that answer to it — spawned ones with the agent that spawned them — (agents), its screening against the OFAC, UN, EU and UK sanctions lists and the embargoed jurisdictions, run now, and what it still lacks to pay and be paid.
[**GetPrincipalAgentsByRef**](PrincipalAPI.md#GetPrincipalAgentsByRef) | **Get** /v1/principal/agents/{ref} | Attests which principal one of the caller&#39;s agents answers to: the org, the chain of agents that spawned it, and the same signed with the platform&#39;s key as a JWS (EdDSA) a verifier checks offline against GET /v1/principal/keys.
[**GetPrincipalClearance**](PrincipalAPI.md#GetPrincipalClearance) | **Get** /v1/principal/clearance | Lists the caller org&#39;s clearances, newest first, 200 at most; a payee narrows them.
[**GetPrincipalClearanceById**](PrincipalAPI.md#GetPrincipalClearanceById) | **Get** /v1/principal/clearance/{id} | Reads one of the caller org&#39;s clearances, as it was decided.
[**GetPrincipalKeys**](PrincipalAPI.md#GetPrincipalKeys) | **Get** /v1/principal/keys | Publishes the keys attestations are signed with, as a JSON Web Key Set (RFC 7517; Ed25519 keys as RFC 8037 writes them).
[**GetPrincipalStatementsById**](PrincipalAPI.md#GetPrincipalStatementsById) | **Get** /v1/principal/statements/{id} | Reads whether one statement the platform signed about an agent still stands: live, expired, or revoked and when.
[**PostPrincipalClearance**](PrincipalAPI.md#PostPrincipalClearance) | **Post** /v1/principal/clearance | Decides whether the caller&#39;s org may pay another org, BEFORE the payment: what must happen first, what must be withheld and why, what must be reported, by whom and when, and how the payment may settle net of withholding.



## GetPrincipal

> PrincipalView GetPrincipal(ctx).Execute()

Returns the caller's org as an economic principal, composed from the apps that own each fact: its legal entity and founders' identity verification (company), the W-9 or W-8 it certified (tax), its wallets (wallet), the agents that answer to it — spawned ones with the agent that spawned them — (agents), its screening against the OFAC, UN, EU and UK sanctions lists and the embargoed jurisdictions, run now, and what it still lacks to pay and be paid.



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
	resp, r, err := apiClient.PrincipalAPI.GetPrincipal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipal`: PrincipalView
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalRequest struct via the builder pattern


### Return type

[**PrincipalView**](PrincipalView.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalAgentsByRef

> PrincipalAttestation GetPrincipalAgentsByRef(ctx, ref).Execute()

Attests which principal one of the caller's agents answers to: the org, the chain of agents that spawned it, and the same signed with the platform's key as a JWS (EdDSA) a verifier checks offline against GET /v1/principal/keys.



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
	ref := "ref_example" // string | Ref is the agent's id or its name in the caller's org.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAPI.GetPrincipalAgentsByRef(context.Background(), ref).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipalAgentsByRef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalAgentsByRef`: PrincipalAttestation
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipalAgentsByRef`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ref** | **string** | Ref is the agent&#39;s id or its name in the caller&#39;s org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalAgentsByRefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalAttestation**](PrincipalAttestation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalClearance

> PrincipalClearanceList GetPrincipalClearance(ctx).Payee(payee).Execute()

Lists the caller org's clearances, newest first, 200 at most; a payee narrows them.



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
	payee := "payee_example" // string | Payee narrows to one payee. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAPI.GetPrincipalClearance(context.Background()).Payee(payee).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipalClearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalClearance`: PrincipalClearanceList
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipalClearance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalClearanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payee** | **string** | Payee narrows to one payee. | 

### Return type

[**PrincipalClearanceList**](PrincipalClearanceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalClearanceById

> PrincipalClearance GetPrincipalClearanceById(ctx, id).Execute()

Reads one of the caller org's clearances, as it was decided.



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
	id := "id_example" // string | ID is the clearance, \"clr_\"-prefixed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAPI.GetPrincipalClearanceById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipalClearanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalClearanceById`: PrincipalClearance
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipalClearanceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the clearance, \&quot;clr_\&quot;-prefixed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalClearanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalClearance**](PrincipalClearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalKeys

> PrincipalJWKS GetPrincipalKeys(ctx).Execute()

Publishes the keys attestations are signed with, as a JSON Web Key Set (RFC 7517; Ed25519 keys as RFC 8037 writes them).



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
	resp, r, err := apiClient.PrincipalAPI.GetPrincipalKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipalKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalKeys`: PrincipalJWKS
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipalKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalKeysRequest struct via the builder pattern


### Return type

[**PrincipalJWKS**](PrincipalJWKS.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalStatementsById

> PrincipalStatement GetPrincipalStatementsById(ctx, id).Execute()

Reads whether one statement the platform signed about an agent still stands: live, expired, or revoked and when.



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
	id := "id_example" // string | ID is the statement's jti, as its JWS carries it.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAPI.GetPrincipalStatementsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.GetPrincipalStatementsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalStatementsById`: PrincipalStatement
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.GetPrincipalStatementsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the statement&#39;s jti, as its JWS carries it. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalStatementsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalStatement**](PrincipalStatement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPrincipalClearance

> PrincipalClearance PostPrincipalClearance(ctx).PrincipalClearIn(principalClearIn).Execute()

Decides whether the caller's org may pay another org, BEFORE the payment: what must happen first, what must be withheld and why, what must be reported, by whom and when, and how the payment may settle net of withholding.



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
	principalClearIn := *openapiclient.NewPrincipalClearIn() // PrincipalClearIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAPI.PostPrincipalClearance(context.Background()).PrincipalClearIn(principalClearIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAPI.PostPrincipalClearance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPrincipalClearance`: PrincipalClearance
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAPI.PostPrincipalClearance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPrincipalClearanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalClearIn** | [**PrincipalClearIn**](PrincipalClearIn.md) |  | 

### Return type

[**PrincipalClearance**](PrincipalClearance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

