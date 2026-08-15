# \OrgAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrgSettings**](OrgAPI.md#DeleteOrgSettings) | **Delete** /v1/org/settings | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**DeleteOrgSettingsList**](OrgAPI.md#DeleteOrgSettingsList) | **Delete** /v1/org/settings/list | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**GetOrgSettings**](OrgAPI.md#GetOrgSettings) | **Get** /v1/org/settings | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**GetOrgSettingsList**](OrgAPI.md#GetOrgSettingsList) | **Get** /v1/org/settings/list | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PatchOrgSettings**](OrgAPI.md#PatchOrgSettings) | **Patch** /v1/org/settings | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PatchOrgSettingsList**](OrgAPI.md#PatchOrgSettingsList) | **Patch** /v1/org/settings/list | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PostOrgSettings**](OrgAPI.md#PostOrgSettings) | **Post** /v1/org/settings | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PostOrgSettingsList**](OrgAPI.md#PostOrgSettingsList) | **Post** /v1/org/settings/list | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PutOrgSettings**](OrgAPI.md#PutOrgSettings) | **Put** /v1/org/settings | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).
[**PutOrgSettingsList**](OrgAPI.md#PutOrgSettingsList) | **Put** /v1/org/settings/list | The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



## DeleteOrgSettings

> DeleteOrgSettings(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.DeleteOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.DeleteOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSettingsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgSettingsList

> DeleteOrgSettingsList(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.DeleteOrgSettingsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.DeleteOrgSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgSettingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSettings

> GetOrgSettings(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.GetOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.GetOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgSettingsList

> GetOrgSettingsList(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.GetOrgSettingsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.GetOrgSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrgSettings

> PatchOrgSettings(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PatchOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PatchOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrgSettingsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOrgSettingsList

> PatchOrgSettingsList(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PatchOrgSettingsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PatchOrgSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOrgSettingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgSettings

> PostOrgSettings(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PostOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PostOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgSettingsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgSettingsList

> PostOrgSettingsList(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PostOrgSettingsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PostOrgSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgSettingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgSettings

> PutOrgSettings(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PutOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PutOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgSettingsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOrgSettingsList

> PutOrgSettingsList(ctx).Execute()

The HTTP transport binding for the RESTful router-config nouns (/v1/router/{policy,defaults,ledger,rewards,artifact-meta} and /v1/org/settings[/list]).



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
	r, err := apiClient.OrgAPI.PutOrgSettingsList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgAPI.PutOrgSettingsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutOrgSettingsListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

