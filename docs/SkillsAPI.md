# \SkillsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSkillsById**](SkillsAPI.md#DeleteSkillsById) | **Delete** /v1/skills/{id} | Removes one of the caller org&#39;s authored skills.
[**GetSkills**](SkillsAPI.md#GetSkills) | **Get** /v1/skills | Lists the skills the caller&#39;s org can reach — the brand&#39;s embedded catalogue plus the org&#39;s own authored ones — with each one&#39;s activation flag.
[**GetSkillsAuthored**](SkillsAPI.md#GetSkillsAuthored) | **Get** /v1/skills/authored | Lists the caller org&#39;s OWN skills with their SKILL.md bodies.
[**PostSkills**](SkillsAPI.md#PostSkills) | **Post** /v1/skills | Adds or revises one of the caller org&#39;s own skills, and answers 201 with the stored record.



## DeleteSkillsById

> SkillDeleted DeleteSkillsById(ctx, id).Execute()

Removes one of the caller org's authored skills.



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
	id := "id_example" // string | ID is the skill to remove, from the path. It is the skill's name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.DeleteSkillsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.DeleteSkillsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSkillsById`: SkillDeleted
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.DeleteSkillsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID is the skill to remove, from the path. It is the skill&#39;s name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSkillsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkillDeleted**](SkillDeleted.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkills

> SourceToolList GetSkills(ctx).Activated(activated).Execute()

Lists the skills the caller's org can reach — the brand's embedded catalogue plus the org's own authored ones — with each one's activation flag.



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
	activated := "activated_example" // string | Activated keeps only the tools activated for the caller's org and project, and only when it is exactly the string \"true\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.GetSkills(context.Background()).Activated(activated).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.GetSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkills`: SourceToolList
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.GetSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activated** | **string** | Activated keeps only the tools activated for the caller&#39;s org and project, and only when it is exactly the string \&quot;true\&quot;. | 

### Return type

[**SourceToolList**](SourceToolList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSkillsAuthored

> AuthoredSkillList GetSkillsAuthored(ctx).Execute()

Lists the caller org's OWN skills with their SKILL.md bodies.



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
	resp, r, err := apiClient.SkillsAPI.GetSkillsAuthored(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.GetSkillsAuthored``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSkillsAuthored`: AuthoredSkillList
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.GetSkillsAuthored`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkillsAuthoredRequest struct via the builder pattern


### Return type

[**AuthoredSkillList**](AuthoredSkillList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSkills

> SkillWritten PostSkills(ctx).SkillIn(skillIn).Execute()

Adds or revises one of the caller org's own skills, and answers 201 with the stored record.



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
	skillIn := *openapiclient.NewSkillIn() // SkillIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SkillsAPI.PostSkills(context.Background()).SkillIn(skillIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SkillsAPI.PostSkills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSkills`: SkillWritten
	fmt.Fprintf(os.Stdout, "Response from `SkillsAPI.PostSkills`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSkillsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skillIn** | [**SkillIn**](SkillIn.md) |  | 

### Return type

[**SkillWritten**](SkillWritten.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

