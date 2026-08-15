# \AskAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAsk**](AskAPI.md#PostAsk) | **Post** /v1/ask | Ask a grounded question about your own org
[**ResearchWeb**](AskAPI.md#ResearchWeb) | **Post** /v1/ask/web | Research a question on the live web and answer it with sources cited



## PostAsk

> PostAsk(ctx).AskRequest(askRequest).Execute()

Ask a grounded question about your own org



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
	askRequest := *openapiclient.NewAskRequest() // AskRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AskAPI.PostAsk(context.Background()).AskRequest(askRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AskAPI.PostAsk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **askRequest** | [**AskRequest**](AskRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchWeb

> Report ResearchWeb(ctx).WebQuestion(webQuestion).Execute()

Research a question on the live web and answer it with sources cited



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
	webQuestion := *openapiclient.NewWebQuestion() // WebQuestion | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AskAPI.ResearchWeb(context.Background()).WebQuestion(webQuestion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AskAPI.ResearchWeb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchWeb`: Report
	fmt.Fprintf(os.Stdout, "Response from `AskAPI.ResearchWeb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchWebRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webQuestion** | [**WebQuestion**](WebQuestion.md) |  | 

### Return type

[**Report**](Report.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

