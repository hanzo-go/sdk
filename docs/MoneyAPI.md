# \MoneyAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAdminAnalytics**](MoneyAPI.md#AdminAdminAnalytics) | **Get** /v1/admin/analytics | Native SaaS analytics (retention/growth/churn)
[**AdminAdminFinance**](MoneyAPI.md#AdminAdminFinance) | **Get** /v1/admin/finance | COGS / gross-margin / runway dashboard
[**AdminAdminRevenue**](MoneyAPI.md#AdminAdminRevenue) | **Get** /v1/admin/revenue | Fleet revenue aggregate



## AdminAdminAnalytics

> AdminAdminAnalytics200Response AdminAdminAnalytics(ctx).Range_(range_).Execute()

Native SaaS analytics (retention/growth/churn)

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
	range_ := "range__example" // string |  (optional) (default to "30d")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoneyAPI.AdminAdminAnalytics(context.Background()).Range_(range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoneyAPI.AdminAdminAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminAnalytics`: AdminAdminAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `MoneyAPI.AdminAdminAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** |  | [default to &quot;30d&quot;]

### Return type

[**AdminAdminAnalytics200Response**](AdminAdminAnalytics200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminFinance

> AdminAdminFinance200Response AdminAdminFinance(ctx).Execute()

COGS / gross-margin / runway dashboard

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
	resp, r, err := apiClient.MoneyAPI.AdminAdminFinance(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoneyAPI.AdminAdminFinance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminFinance`: AdminAdminFinance200Response
	fmt.Fprintf(os.Stdout, "Response from `MoneyAPI.AdminAdminFinance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminFinanceRequest struct via the builder pattern


### Return type

[**AdminAdminFinance200Response**](AdminAdminFinance200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminAdminRevenue

> AdminAdminRevenue200Response AdminAdminRevenue(ctx).Execute()

Fleet revenue aggregate

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
	resp, r, err := apiClient.MoneyAPI.AdminAdminRevenue(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoneyAPI.AdminAdminRevenue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminAdminRevenue`: AdminAdminRevenue200Response
	fmt.Fprintf(os.Stdout, "Response from `MoneyAPI.AdminAdminRevenue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminAdminRevenueRequest struct via the builder pattern


### Return type

[**AdminAdminRevenue200Response**](AdminAdminRevenue200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

