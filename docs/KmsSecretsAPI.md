# \KmsSecretsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KmsBatchSecrets**](KmsSecretsAPI.md#KmsBatchSecrets) | **Post** /v1/kms/v4/secrets/batch | Batch create, update, and delete secrets
[**KmsCreateSecret**](KmsSecretsAPI.md#KmsCreateSecret) | **Post** /v1/kms/v4/secrets/{secretKey} | Create a secret
[**KmsDeleteSecret**](KmsSecretsAPI.md#KmsDeleteSecret) | **Delete** /v1/kms/v4/secrets/{secretKey} | Delete a secret
[**KmsGetSecret**](KmsSecretsAPI.md#KmsGetSecret) | **Get** /v1/kms/v4/secrets/{secretKey} | Get a secret by key
[**KmsKmsDeleteOrgSecret**](KmsSecretsAPI.md#KmsKmsDeleteOrgSecret) | **Delete** /v1/kms/orgs/{org}/secrets/{rest} | Delete a secret
[**KmsKmsGetOrgSecret**](KmsSecretsAPI.md#KmsKmsGetOrgSecret) | **Get** /v1/kms/orgs/{org}/secrets/{rest} | Reveal one secret value
[**KmsKmsListOrgSecrets**](KmsSecretsAPI.md#KmsKmsListOrgSecrets) | **Get** /v1/kms/orgs/{org}/secrets | List an org&#39;s secret metadata (names only, never values)
[**KmsKmsPutOrgSecret**](KmsSecretsAPI.md#KmsKmsPutOrgSecret) | **Post** /v1/kms/orgs/{org}/secrets | Create or upsert a secret (value write-only; version bumps)
[**KmsKmsRotateOrgSecret**](KmsSecretsAPI.md#KmsKmsRotateOrgSecret) | **Patch** /v1/kms/orgs/{org}/secrets/{rest} | Rotate a secret (compare-and-set on version)
[**KmsListSecrets**](KmsSecretsAPI.md#KmsListSecrets) | **Get** /v1/kms/v4/secrets | List secrets
[**KmsUpdateSecret**](KmsSecretsAPI.md#KmsUpdateSecret) | **Patch** /v1/kms/v4/secrets/{secretKey} | Update a secret



## KmsBatchSecrets

> KmsBatchSecrets200Response KmsBatchSecrets(ctx).WorkspaceId(workspaceId).KmsBatchSecretRequest(kmsBatchSecretRequest).Execute()

Batch create, update, and delete secrets

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	kmsBatchSecretRequest := *openapiclient.NewKmsBatchSecretRequest() // KmsBatchSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsBatchSecrets(context.Background()).WorkspaceId(workspaceId).KmsBatchSecretRequest(kmsBatchSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsBatchSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsBatchSecrets`: KmsBatchSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsBatchSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsBatchSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **kmsBatchSecretRequest** | [**KmsBatchSecretRequest**](KmsBatchSecretRequest.md) |  | 

### Return type

[**KmsBatchSecrets200Response**](KmsBatchSecrets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsCreateSecret

> KmsGetSecret200Response KmsCreateSecret(ctx, secretKey).KmsCreateSecretRequest(kmsCreateSecretRequest).Execute()

Create a secret

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
	secretKey := "secretKey_example" // string | 
	kmsCreateSecretRequest := *openapiclient.NewKmsCreateSecretRequest("SecretKey_example", "SecretValue_example", "SecretPath_example", "Environment_example") // KmsCreateSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsCreateSecret(context.Background(), secretKey).KmsCreateSecretRequest(kmsCreateSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsCreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsCreateSecret`: KmsGetSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsCreateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsCreateSecretRequest** | [**KmsCreateSecretRequest**](KmsCreateSecretRequest.md) |  | 

### Return type

[**KmsGetSecret200Response**](KmsGetSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsDeleteSecret

> KmsGetSecret200Response KmsDeleteSecret(ctx, secretKey).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).Execute()

Delete a secret

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
	secretKey := "secretKey_example" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	environment := "environment_example" // string | 
	secretPath := "secretPath_example" // string |  (optional) (default to "/")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsDeleteSecret(context.Background(), secretKey).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsDeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsDeleteSecret`: KmsGetSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsDeleteSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsDeleteSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** |  | 
 **environment** | **string** |  | 
 **secretPath** | **string** |  | [default to &quot;/&quot;]

### Return type

[**KmsGetSecret200Response**](KmsGetSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetSecret

> KmsGetSecret200Response KmsGetSecret(ctx, secretKey).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).ExpandSecretReferences(expandSecretReferences).Version(version).Execute()

Get a secret by key

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
	secretKey := "secretKey_example" // string | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	environment := "environment_example" // string | 
	secretPath := "secretPath_example" // string |  (optional) (default to "/")
	expandSecretReferences := true // bool |  (optional) (default to true)
	version := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsGetSecret(context.Background(), secretKey).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).ExpandSecretReferences(expandSecretReferences).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsGetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsGetSecret`: KmsGetSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsGetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workspaceId** | **string** |  | 
 **environment** | **string** |  | 
 **secretPath** | **string** |  | [default to &quot;/&quot;]
 **expandSecretReferences** | **bool** |  | [default to true]
 **version** | **int32** |  | 

### Return type

[**KmsGetSecret200Response**](KmsGetSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsKmsDeleteOrgSecret

> AnalyticsHeartbeat200Response KmsKmsDeleteOrgSecret(ctx, org, rest).Env(env).Execute()

Delete a secret

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
	org := "org_example" // string | 
	rest := "rest_example" // string | The secret's path + name joined (a/b/c/name).
	env := "env_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsKmsDeleteOrgSecret(context.Background(), org, rest).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsKmsDeleteOrgSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsKmsDeleteOrgSecret`: AnalyticsHeartbeat200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsKmsDeleteOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**rest** | **string** | The secret&#39;s path + name joined (a/b/c/name). | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsKmsDeleteOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** |  | 

### Return type

[**AnalyticsHeartbeat200Response**](AnalyticsHeartbeat200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsKmsGetOrgSecret

> KmsKmsGetOrgSecret200Response KmsKmsGetOrgSecret(ctx, org, rest).Env(env).Execute()

Reveal one secret value

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
	org := "org_example" // string | 
	rest := "rest_example" // string | The secret's path + name joined (a/b/c/name).
	env := "env_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsKmsGetOrgSecret(context.Background(), org, rest).Env(env).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsKmsGetOrgSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsKmsGetOrgSecret`: KmsKmsGetOrgSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsKmsGetOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**rest** | **string** | The secret&#39;s path + name joined (a/b/c/name). | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsKmsGetOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | **string** |  | 

### Return type

[**KmsKmsGetOrgSecret200Response**](KmsKmsGetOrgSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsKmsListOrgSecrets

> KmsKmsListOrgSecrets200Response KmsKmsListOrgSecrets(ctx, org).Env(env).Prefix(prefix).Execute()

List an org's secret metadata (names only, never values)

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
	org := "org_example" // string | 
	env := "env_example" // string | Environment slug (devnet, testnet, mainnet, production, …). Default \"default\". (optional)
	prefix := "prefix_example" // string | Restrict the scan to a subpath within brand/{org}. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsKmsListOrgSecrets(context.Background(), org).Env(env).Prefix(prefix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsKmsListOrgSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsKmsListOrgSecrets`: KmsKmsListOrgSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsKmsListOrgSecrets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsKmsListOrgSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | **string** | Environment slug (devnet, testnet, mainnet, production, …). Default \&quot;default\&quot;. | 
 **prefix** | **string** | Restrict the scan to a subpath within brand/{org}. | 

### Return type

[**KmsKmsListOrgSecrets200Response**](KmsKmsListOrgSecrets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsKmsPutOrgSecret

> KmsKmsPutOrgSecret200Response KmsKmsPutOrgSecret(ctx, org).KmsKmsPutOrgSecretRequest(kmsKmsPutOrgSecretRequest).Execute()

Create or upsert a secret (value write-only; version bumps)

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
	org := "org_example" // string | 
	kmsKmsPutOrgSecretRequest := *openapiclient.NewKmsKmsPutOrgSecretRequest("Name_example", "Value_example") // KmsKmsPutOrgSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsKmsPutOrgSecret(context.Background(), org).KmsKmsPutOrgSecretRequest(kmsKmsPutOrgSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsKmsPutOrgSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsKmsPutOrgSecret`: KmsKmsPutOrgSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsKmsPutOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsKmsPutOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsKmsPutOrgSecretRequest** | [**KmsKmsPutOrgSecretRequest**](KmsKmsPutOrgSecretRequest.md) |  | 

### Return type

[**KmsKmsPutOrgSecret200Response**](KmsKmsPutOrgSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsKmsRotateOrgSecret

> KmsKmsPutOrgSecret200Response KmsKmsRotateOrgSecret(ctx, org, rest).KmsKmsRotateOrgSecretRequest(kmsKmsRotateOrgSecretRequest).IfMatch(ifMatch).Execute()

Rotate a secret (compare-and-set on version)

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
	org := "org_example" // string | 
	rest := "rest_example" // string | The secret's path + name joined (a/b/c/name).
	kmsKmsRotateOrgSecretRequest := *openapiclient.NewKmsKmsRotateOrgSecretRequest("Value_example", int64(123)) // KmsKmsRotateOrgSecretRequest | 
	ifMatch := "ifMatch_example" // string | Current version for the compare-and-set (replay protection). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsKmsRotateOrgSecret(context.Background(), org, rest).KmsKmsRotateOrgSecretRequest(kmsKmsRotateOrgSecretRequest).IfMatch(ifMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsKmsRotateOrgSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsKmsRotateOrgSecret`: KmsKmsPutOrgSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsKmsRotateOrgSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**org** | **string** |  | 
**rest** | **string** | The secret&#39;s path + name joined (a/b/c/name). | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsKmsRotateOrgSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kmsKmsRotateOrgSecretRequest** | [**KmsKmsRotateOrgSecretRequest**](KmsKmsRotateOrgSecretRequest.md) |  | 
 **ifMatch** | **string** | Current version for the compare-and-set (replay protection). | 

### Return type

[**KmsKmsPutOrgSecret200Response**](KmsKmsPutOrgSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsListSecrets

> KmsListSecrets200Response KmsListSecrets(ctx).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).ExpandSecretReferences(expandSecretReferences).Recursive(recursive).IncludeImports(includeImports).TagSlugs(tagSlugs).MetadataFilter(metadataFilter).Execute()

List secrets

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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	environment := "environment_example" // string | 
	secretPath := "secretPath_example" // string |  (optional) (default to "/")
	expandSecretReferences := true // bool |  (optional) (default to true)
	recursive := true // bool |  (optional) (default to false)
	includeImports := true // bool |  (optional) (default to false)
	tagSlugs := "tagSlugs_example" // string |  (optional)
	metadataFilter := "metadataFilter_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsListSecrets(context.Background()).WorkspaceId(workspaceId).Environment(environment).SecretPath(secretPath).ExpandSecretReferences(expandSecretReferences).Recursive(recursive).IncludeImports(includeImports).TagSlugs(tagSlugs).MetadataFilter(metadataFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsListSecrets`: KmsListSecrets200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsListSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKmsListSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** |  | 
 **environment** | **string** |  | 
 **secretPath** | **string** |  | [default to &quot;/&quot;]
 **expandSecretReferences** | **bool** |  | [default to true]
 **recursive** | **bool** |  | [default to false]
 **includeImports** | **bool** |  | [default to false]
 **tagSlugs** | **string** |  | 
 **metadataFilter** | **string** |  | 

### Return type

[**KmsListSecrets200Response**](KmsListSecrets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsUpdateSecret

> KmsGetSecret200Response KmsUpdateSecret(ctx, secretKey).KmsUpdateSecretRequest(kmsUpdateSecretRequest).Execute()

Update a secret

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
	secretKey := "secretKey_example" // string | 
	kmsUpdateSecretRequest := *openapiclient.NewKmsUpdateSecretRequest("SecretKey_example", "SecretPath_example", "Environment_example") // KmsUpdateSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsSecretsAPI.KmsUpdateSecret(context.Background(), secretKey).KmsUpdateSecretRequest(kmsUpdateSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsSecretsAPI.KmsUpdateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KmsUpdateSecret`: KmsGetSecret200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsSecretsAPI.KmsUpdateSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsUpdateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kmsUpdateSecretRequest** | [**KmsUpdateSecretRequest**](KmsUpdateSecretRequest.md) |  | 

### Return type

[**KmsGetSecret200Response**](KmsGetSecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

