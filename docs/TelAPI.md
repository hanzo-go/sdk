# \TelAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTelCallsById**](TelAPI.md#DeleteTelCallsById) | **Delete** /v1/tel/calls/{id} | Ends a call this org placed.
[**DeleteTelNumbersById**](TelAPI.md#DeleteTelNumbersById) | **Delete** /v1/tel/numbers/{id} | Checks the holding is THIS org&#39;s before it reaches the carrier.
[**GetTelCalls**](TelAPI.md#GetTelCalls) | **Get** /v1/tel/calls | Lists the calls this org has placed or received, newest first.
[**GetTelMessages**](TelAPI.md#GetTelMessages) | **Get** /v1/tel/messages | Lists the messages this org has sent or received, newest first.
[**GetTelNumbers**](TelAPI.md#GetTelNumbers) | **Get** /v1/tel/numbers | Lists the phone numbers this org HOLDS — the ones it has bought and not released.
[**GetTelNumbersAvailable**](TelAPI.md#GetTelNumbersAvailable) | **Get** /v1/tel/numbers/available | Asks the carrier what is available to buy.
[**GetTelSummary**](TelAPI.md#GetTelSummary) | **Get** /v1/tel/summary | Counts what this org holds on the telephony plane: its numbers, its calls and its messages.
[**PostTelCalls**](TelAPI.md#PostTelCalls) | **Post** /v1/tel/calls | Dials.
[**PostTelMessages**](TelAPI.md#PostTelMessages) | **Post** /v1/tel/messages | Sends a message from one of this org&#39;s own numbers.
[**PostTelNumbers**](TelAPI.md#PostTelNumbers) | **Post** /v1/tel/numbers | Provisions with the carrier FIRST and records second.



## DeleteTelCallsById

> DeleteTelCallsById(ctx, id).Execute()

Ends a call this org placed.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TelAPI.DeleteTelCallsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.DeleteTelCallsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelCallsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTelNumbersById

> DeleteTelNumbersById(ctx, id).Execute()

Checks the holding is THIS org's before it reaches the carrier.



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TelAPI.DeleteTelNumbersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.DeleteTelNumbersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTelNumbersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelCalls

> TelCallList GetTelCalls(ctx).Execute()

Lists the calls this org has placed or received, newest first.



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
	resp, r, err := apiClient.TelAPI.GetTelCalls(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.GetTelCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelCalls`: TelCallList
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.GetTelCalls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelCallsRequest struct via the builder pattern


### Return type

[**TelCallList**](TelCallList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelMessages

> TelMessageList GetTelMessages(ctx).Execute()

Lists the messages this org has sent or received, newest first.



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
	resp, r, err := apiClient.TelAPI.GetTelMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.GetTelMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelMessages`: TelMessageList
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.GetTelMessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelMessagesRequest struct via the builder pattern


### Return type

[**TelMessageList**](TelMessageList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelNumbers

> TelNumberList GetTelNumbers(ctx).Execute()

Lists the phone numbers this org HOLDS — the ones it has bought and not released.



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
	resp, r, err := apiClient.TelAPI.GetTelNumbers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.GetTelNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelNumbers`: TelNumberList
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.GetTelNumbers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelNumbersRequest struct via the builder pattern


### Return type

[**TelNumberList**](TelNumberList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelNumbersAvailable

> TelNumberList GetTelNumbersAvailable(ctx).Country(country).Area(area).Type_(type_).Limit(limit).Execute()

Asks the carrier what is available to buy.



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
	country := "country_example" // string |  (optional)
	area := "area_example" // string |  (optional)
	type_ := "type__example" // string |  (optional)
	limit := int64(789) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelAPI.GetTelNumbersAvailable(context.Background()).Country(country).Area(area).Type_(type_).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.GetTelNumbersAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelNumbersAvailable`: TelNumberList
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.GetTelNumbersAvailable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTelNumbersAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 
 **area** | **string** |  | 
 **type_** | **string** |  | 
 **limit** | **int64** |  | 

### Return type

[**TelNumberList**](TelNumberList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTelSummary

> TelSummary GetTelSummary(ctx).Execute()

Counts what this org holds on the telephony plane: its numbers, its calls and its messages.



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
	resp, r, err := apiClient.TelAPI.GetTelSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.GetTelSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTelSummary`: TelSummary
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.GetTelSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTelSummaryRequest struct via the builder pattern


### Return type

[**TelSummary**](TelSummary.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTelCalls

> TelCall PostTelCalls(ctx).TelCallInput(telCallInput).Execute()

Dials.



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
	telCallInput := *openapiclient.NewTelCallInput() // TelCallInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelAPI.PostTelCalls(context.Background()).TelCallInput(telCallInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.PostTelCalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTelCalls`: TelCall
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.PostTelCalls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTelCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telCallInput** | [**TelCallInput**](TelCallInput.md) |  | 

### Return type

[**TelCall**](TelCall.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTelMessages

> TelSMS PostTelMessages(ctx).TelMessageInput(telMessageInput).Execute()

Sends a message from one of this org's own numbers.



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
	telMessageInput := *openapiclient.NewTelMessageInput() // TelMessageInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelAPI.PostTelMessages(context.Background()).TelMessageInput(telMessageInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.PostTelMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTelMessages`: TelSMS
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.PostTelMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTelMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telMessageInput** | [**TelMessageInput**](TelMessageInput.md) |  | 

### Return type

[**TelSMS**](TelSMS.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTelNumbers

> TelNumber PostTelNumbers(ctx).TelBuyInput(telBuyInput).Execute()

Provisions with the carrier FIRST and records second.



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
	telBuyInput := *openapiclient.NewTelBuyInput() // TelBuyInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TelAPI.PostTelNumbers(context.Background()).TelBuyInput(telBuyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TelAPI.PostTelNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTelNumbers`: TelNumber
	fmt.Fprintf(os.Stdout, "Response from `TelAPI.PostTelNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTelNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telBuyInput** | [**TelBuyInput**](TelBuyInput.md) |  | 

### Return type

[**TelNumber**](TelNumber.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

