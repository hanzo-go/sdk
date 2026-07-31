# \KeysAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudDeleteV1Keys**](KeysAPI.md#CloudDeleteV1Keys) | **Delete** /v1/keys | RevokeKey revokes the caller&#39;s own API key of the requested class.
[**CloudGetV1Keys**](KeysAPI.md#CloudGetV1Keys) | **Get** /v1/keys | GetKey returns the caller&#39;s own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.
[**CloudPostV1Keys**](KeysAPI.md#CloudPostV1Keys) | **Post** /v1/keys | MintKey creates — or rotates — the caller&#39;s API key of the requested type and returns it ONCE.
[**GatewayDeleteKey**](KeysAPI.md#GatewayDeleteKey) | **Post** /v1/gateway/key/delete | Delete key
[**GatewayGenerateKey**](KeysAPI.md#GatewayGenerateKey) | **Post** /v1/gateway/key/generate | Generate API key
[**GatewayGetKeyInfo**](KeysAPI.md#GatewayGetKeyInfo) | **Get** /v1/gateway/key/info | Get key info
[**GatewayUpdateKey**](KeysAPI.md#GatewayUpdateKey) | **Post** /v1/gateway/key/update | Update key
[**KmsGetV1KmsKeys**](KeysAPI.md#KmsGetV1KmsKeys) | **Get** /v1/kms/keys | List key sets
[**KmsGetV1KmsKeysId**](KeysAPI.md#KmsGetV1KmsKeysId) | **Get** /v1/kms/keys/{id} | Get a key set
[**KmsGetV1KmsStatus**](KeysAPI.md#KmsGetV1KmsStatus) | **Get** /v1/kms/status | KMS + MPC cluster status
[**KmsPostV1KmsKeysGenerate**](KeysAPI.md#KmsPostV1KmsKeysGenerate) | **Post** /v1/kms/keys/generate | Generate a validator key set via MPC DKG
[**KmsPostV1KmsKeysIdRotate**](KeysAPI.md#KmsPostV1KmsKeysIdRotate) | **Post** /v1/kms/keys/{id}/rotate | Reshare a key set
[**KmsPostV1KmsKeysIdSign**](KeysAPI.md#KmsPostV1KmsKeysIdSign) | **Post** /v1/kms/keys/{id}/sign | Threshold sign
[**KvBatchOperation**](KeysAPI.md#KvBatchOperation) | **Post** /v1/kv/batch | Batch get/set/delete
[**KvDeleteKey**](KeysAPI.md#KvDeleteKey) | **Delete** /v1/kv/keys/{key} | Delete key
[**KvGetKey**](KeysAPI.md#KvGetKey) | **Get** /v1/kv/keys/{key} | Get key value
[**KvIncrKey**](KeysAPI.md#KvIncrKey) | **Post** /v1/kv/keys/{key}/incr | Increment numeric key
[**KvScanKeys**](KeysAPI.md#KvScanKeys) | **Get** /v1/kv/keys | Scan keys
[**KvSetKey**](KeysAPI.md#KvSetKey) | **Put** /v1/kv/keys/{key} | Set key value
[**KvSetKeyTTL**](KeysAPI.md#KvSetKeyTTL) | **Put** /v1/kv/keys/{key}/ttl | Set key TTL
[**SearchCreateKey**](KeysAPI.md#SearchCreateKey) | **Post** /v1/search/keys | Create an API key
[**SearchDeleteKey**](KeysAPI.md#SearchDeleteKey) | **Delete** /v1/search/keys/{keyOrUid} | Delete an API key
[**SearchGetKey**](KeysAPI.md#SearchGetKey) | **Get** /v1/search/keys/{keyOrUid} | Get an API key
[**SearchListKeys**](KeysAPI.md#SearchListKeys) | **Get** /v1/search/keys | List API keys
[**SearchUpdateKey**](KeysAPI.md#SearchUpdateKey) | **Patch** /v1/search/keys/{keyOrUid} | Update an API key



## CloudDeleteV1Keys

> CloudRevokedKey CloudDeleteV1Keys(ctx).Type_(type_).Execute()

RevokeKey revokes the caller's own API key of the requested class.



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
	type_ := "publishable" // string | Type is the key class to act on: \"secret\" (sk-, session-equivalent, belongs on a server) or \"publishable\" (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.CloudDeleteV1Keys(context.Background()).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CloudDeleteV1Keys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudDeleteV1Keys`: CloudRevokedKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CloudDeleteV1Keys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudDeleteV1KeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type is the key class to act on: \&quot;secret\&quot; (sk-, session-equivalent, belongs on a server) or \&quot;publishable\&quot; (pk-, org-identifying, safe in a browser bundle). Omitted means secret, which is what every existing caller means. | 

### Return type

[**CloudRevokedKey**](CloudRevokedKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudGetV1Keys

> CloudApiKeyList CloudGetV1Keys(ctx).Execute()

GetKey returns the caller's own API keys — every type they hold, read AUTHORITATIVELY from IAM rather than from the session claim, which lags a key minted moments ago.



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
	resp, r, err := apiClient.KeysAPI.CloudGetV1Keys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CloudGetV1Keys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudGetV1Keys`: CloudApiKeyList
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CloudGetV1Keys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudGetV1KeysRequest struct via the builder pattern


### Return type

[**CloudApiKeyList**](CloudApiKeyList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudPostV1Keys

> CloudMintedKey CloudPostV1Keys(ctx).CloudKeyTypeIn(cloudKeyTypeIn).Execute()

MintKey creates — or rotates — the caller's API key of the requested type and returns it ONCE.



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
	cloudKeyTypeIn := *openapiclient.NewCloudKeyTypeIn() // CloudKeyTypeIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.CloudPostV1Keys(context.Background()).CloudKeyTypeIn(cloudKeyTypeIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.CloudPostV1Keys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloudPostV1Keys`: CloudMintedKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.CloudPostV1Keys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudPostV1KeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudKeyTypeIn** | [**CloudKeyTypeIn**](CloudKeyTypeIn.md) |  | 

### Return type

[**CloudMintedKey**](CloudMintedKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDeleteKey

> map[string]interface{} GatewayDeleteKey(ctx).GatewayDeleteKeyRequest(gatewayDeleteKeyRequest).Execute()

Delete key

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
	gatewayDeleteKeyRequest := *openapiclient.NewGatewayDeleteKeyRequest([]string{"Keys_example"}) // GatewayDeleteKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GatewayDeleteKey(context.Background()).GatewayDeleteKeyRequest(gatewayDeleteKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GatewayDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayDeleteKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GatewayDeleteKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayDeleteKeyRequest** | [**GatewayDeleteKeyRequest**](GatewayDeleteKeyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGenerateKey

> GatewayKey GatewayGenerateKey(ctx).GatewayGenerateKeyRequest(gatewayGenerateKeyRequest).Execute()

Generate API key

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
	gatewayGenerateKeyRequest := *openapiclient.NewGatewayGenerateKeyRequest() // GatewayGenerateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GatewayGenerateKey(context.Background()).GatewayGenerateKeyRequest(gatewayGenerateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GatewayGenerateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGenerateKey`: GatewayKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GatewayGenerateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGenerateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayGenerateKeyRequest** | [**GatewayGenerateKeyRequest**](GatewayGenerateKeyRequest.md) |  | 

### Return type

[**GatewayKey**](GatewayKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayGetKeyInfo

> GatewayKey GatewayGetKeyInfo(ctx).Key(key).Execute()

Get key info

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
	key := "key_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GatewayGetKeyInfo(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GatewayGetKeyInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayGetKeyInfo`: GatewayKey
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GatewayGetKeyInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayGetKeyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** |  | 

### Return type

[**GatewayKey**](GatewayKey.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayUpdateKey

> map[string]interface{} GatewayUpdateKey(ctx).GatewayUpdateKeyRequest(gatewayUpdateKeyRequest).Execute()

Update key

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
	gatewayUpdateKeyRequest := *openapiclient.NewGatewayUpdateKeyRequest("Key_example") // GatewayUpdateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.GatewayUpdateKey(context.Background()).GatewayUpdateKeyRequest(gatewayUpdateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.GatewayUpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayUpdateKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.GatewayUpdateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayUpdateKeyRequest** | [**GatewayUpdateKeyRequest**](GatewayUpdateKeyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsKeys

> KmsGetV1KmsKeys(ctx).Execute()

List key sets



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
	r, err := apiClient.KeysAPI.KmsGetV1KmsKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsGetV1KmsKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsKeysRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsKeysId

> KmsGetV1KmsKeysId(ctx, id).Execute()

Get a key set



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
	id := "id_example" // string | Key set id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.KmsGetV1KmsKeysId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsGetV1KmsKeysId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Key set id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsKeysIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsGetV1KmsStatus

> KmsGetV1KmsStatus(ctx).Execute()

KMS + MPC cluster status



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
	r, err := apiClient.KeysAPI.KmsGetV1KmsStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsGetV1KmsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsGetV1KmsStatusRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsPostV1KmsKeysGenerate

> KmsPostV1KmsKeysGenerate(ctx).Execute()

Generate a validator key set via MPC DKG



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
	r, err := apiClient.KeysAPI.KmsPostV1KmsKeysGenerate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsPostV1KmsKeysGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiKmsPostV1KmsKeysGenerateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsPostV1KmsKeysIdRotate

> KmsPostV1KmsKeysIdRotate(ctx, id).Execute()

Reshare a key set



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
	id := "id_example" // string | Key set id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.KmsPostV1KmsKeysIdRotate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsPostV1KmsKeysIdRotate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Key set id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsPostV1KmsKeysIdRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KmsPostV1KmsKeysIdSign

> KmsPostV1KmsKeysIdSign(ctx, id).Execute()

Threshold sign



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
	id := "id_example" // string | Key set id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.KmsPostV1KmsKeysIdSign(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KmsPostV1KmsKeysIdSign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Key set id | 

### Other Parameters

Other parameters are passed through a pointer to a apiKmsPostV1KmsKeysIdSignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvBatchOperation

> KvBatchOperation200Response KvBatchOperation(ctx).KvBatchOperationRequest(kvBatchOperationRequest).Execute()

Batch get/set/delete

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
	kvBatchOperationRequest := *openapiclient.NewKvBatchOperationRequest() // KvBatchOperationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvBatchOperation(context.Background()).KvBatchOperationRequest(kvBatchOperationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvBatchOperation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvBatchOperation`: KvBatchOperation200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvBatchOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvBatchOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kvBatchOperationRequest** | [**KvBatchOperationRequest**](KvBatchOperationRequest.md) |  | 

### Return type

[**KvBatchOperation200Response**](KvBatchOperation200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvDeleteKey

> KvDeleteKey200Response KvDeleteKey(ctx, key).Namespace(namespace).Execute()

Delete key

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvDeleteKey(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvDeleteKey`: KvDeleteKey200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvDeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

[**KvDeleteKey200Response**](KvDeleteKey200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvGetKey

> KvKeyValue KvGetKey(ctx, key).Namespace(namespace).Execute()

Get key value

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvGetKey(context.Background(), key).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvGetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvGetKey`: KvKeyValue
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvGetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 

### Return type

[**KvKeyValue**](KvKeyValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvIncrKey

> AnalyticsGetSessionStats200ResponseValue KvIncrKey(ctx, key).Namespace(namespace).KvIncrKeyRequest(kvIncrKeyRequest).Execute()

Increment numeric key

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
	key := "key_example" // string | 
	namespace := "namespace_example" // string |  (optional)
	kvIncrKeyRequest := *openapiclient.NewKvIncrKeyRequest() // KvIncrKeyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvIncrKey(context.Background(), key).Namespace(namespace).KvIncrKeyRequest(kvIncrKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvIncrKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvIncrKey`: AnalyticsGetSessionStats200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvIncrKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvIncrKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **string** |  | 
 **kvIncrKeyRequest** | [**KvIncrKeyRequest**](KvIncrKeyRequest.md) |  | 

### Return type

[**AnalyticsGetSessionStats200ResponseValue**](AnalyticsGetSessionStats200ResponseValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvScanKeys

> KvScanKeys200Response KvScanKeys(ctx).Pattern(pattern).Type_(type_).Cursor(cursor).Count(count).Namespace(namespace).Execute()

Scan keys

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
	pattern := "pattern_example" // string | Glob-style pattern (e.g. user:*, session:*) (optional) (default to "*")
	type_ := "type__example" // string |  (optional)
	cursor := "cursor_example" // string |  (optional) (default to "0")
	count := int32(56) // int32 |  (optional) (default to 100)
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvScanKeys(context.Background()).Pattern(pattern).Type_(type_).Cursor(cursor).Count(count).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvScanKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvScanKeys`: KvScanKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvScanKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiKvScanKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pattern** | **string** | Glob-style pattern (e.g. user:*, session:*) | [default to &quot;*&quot;]
 **type_** | **string** |  | 
 **cursor** | **string** |  | [default to &quot;0&quot;]
 **count** | **int32** |  | [default to 100]
 **namespace** | **string** |  | 

### Return type

[**KvScanKeys200Response**](KvScanKeys200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSetKey

> KvKeyValue KvSetKey(ctx, key).KvSetKeyRequest(kvSetKeyRequest).Namespace(namespace).Execute()

Set key value

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
	key := "key_example" // string | 
	kvSetKeyRequest := *openapiclient.NewKvSetKeyRequest("Value_example") // KvSetKeyRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvSetKey(context.Background(), key).KvSetKeyRequest(kvSetKeyRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvSetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSetKey`: KvKeyValue
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvSetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvSetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvSetKeyRequest** | [**KvSetKeyRequest**](KvSetKeyRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

[**KvKeyValue**](KvKeyValue.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## KvSetKeyTTL

> map[string]interface{} KvSetKeyTTL(ctx, key).KvSetKeyTTLRequest(kvSetKeyTTLRequest).Namespace(namespace).Execute()

Set key TTL

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
	key := "key_example" // string | 
	kvSetKeyTTLRequest := *openapiclient.NewKvSetKeyTTLRequest(int32(123)) // KvSetKeyTTLRequest | 
	namespace := "namespace_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.KvSetKeyTTL(context.Background(), key).KvSetKeyTTLRequest(kvSetKeyTTLRequest).Namespace(namespace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.KvSetKeyTTL``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `KvSetKeyTTL`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.KvSetKeyTTL`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiKvSetKeyTTLRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kvSetKeyTTLRequest** | [**KvSetKeyTTLRequest**](KvSetKeyTTLRequest.md) |  | 
 **namespace** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCreateKey

> SearchKeyView SearchCreateKey(ctx).SearchCreateApiKey(searchCreateApiKey).Execute()

Create an API key

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
	searchCreateApiKey := *openapiclient.NewSearchCreateApiKey([]string{"Actions_example"}, []string{"Indexes_example"}) // SearchCreateApiKey | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SearchCreateKey(context.Background()).SearchCreateApiKey(searchCreateApiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchCreateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCreateKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SearchCreateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchCreateApiKey** | [**SearchCreateApiKey**](SearchCreateApiKey.md) |  | 

### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDeleteKey

> SearchDeleteKey(ctx, keyOrUid).Execute()

Delete an API key

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
	keyOrUid := "keyOrUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeysAPI.SearchDeleteKey(context.Background(), keyOrUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchDeleteKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDeleteKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchGetKey

> SearchKeyView SearchGetKey(ctx, keyOrUid).Execute()

Get an API key

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
	keyOrUid := "keyOrUid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SearchGetKey(context.Background(), keyOrUid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchGetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchGetKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SearchGetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchListKeys

> SearchPaginatedKeys SearchListKeys(ctx).Offset(offset).Limit(limit).Execute()

List API keys

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
	offset := int32(56) // int32 |  (optional) (default to 0)
	limit := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SearchListKeys(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchListKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchListKeys`: SearchPaginatedKeys
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SearchListKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**SearchPaginatedKeys**](SearchPaginatedKeys.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUpdateKey

> SearchKeyView SearchUpdateKey(ctx, keyOrUid).SearchUpdateKeyRequest(searchUpdateKeyRequest).Execute()

Update an API key

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
	keyOrUid := "keyOrUid_example" // string | 
	searchUpdateKeyRequest := *openapiclient.NewSearchUpdateKeyRequest() // SearchUpdateKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeysAPI.SearchUpdateKey(context.Background(), keyOrUid).SearchUpdateKeyRequest(searchUpdateKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeysAPI.SearchUpdateKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUpdateKey`: SearchKeyView
	fmt.Fprintf(os.Stdout, "Response from `KeysAPI.SearchUpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyOrUid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchUpdateKeyRequest** | [**SearchUpdateKeyRequest**](SearchUpdateKeyRequest.md) |  | 

### Return type

[**SearchKeyView**](SearchKeyView.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

