# \IamOrganizationsAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IamApiControllerAddGroup**](IamOrganizationsAPI.md#IamApiControllerAddGroup) | **Post** /v1/iam/groups | Api Controller Add Group
[**IamApiControllerAddInvitation**](IamOrganizationsAPI.md#IamApiControllerAddInvitation) | **Post** /v1/iam/invitations | Api Controller Add Invitation
[**IamApiControllerAddOrganization**](IamOrganizationsAPI.md#IamApiControllerAddOrganization) | **Post** /v1/iam/organizations | Api Controller Add Organization
[**IamApiControllerDeleteGroup**](IamOrganizationsAPI.md#IamApiControllerDeleteGroup) | **Delete** /v1/iam/groups/{id} | Api Controller Delete Group
[**IamApiControllerDeleteInvitation**](IamOrganizationsAPI.md#IamApiControllerDeleteInvitation) | **Delete** /v1/iam/invitations/{id} | Api Controller Delete Invitation
[**IamApiControllerDeleteOrganization**](IamOrganizationsAPI.md#IamApiControllerDeleteOrganization) | **Delete** /v1/iam/organizations/{id} | Api Controller Delete Organization
[**IamApiControllerGetDefaultApplication**](IamOrganizationsAPI.md#IamApiControllerGetDefaultApplication) | **Get** /v1/iam/applications/default | Api Controller Get Default Application
[**IamApiControllerGetGroup**](IamOrganizationsAPI.md#IamApiControllerGetGroup) | **Get** /v1/iam/groups/{id} | Api Controller Get Group
[**IamApiControllerGetGroups**](IamOrganizationsAPI.md#IamApiControllerGetGroups) | **Get** /v1/iam/groups | Api Controller Get Groups
[**IamApiControllerGetInvitation**](IamOrganizationsAPI.md#IamApiControllerGetInvitation) | **Get** /v1/iam/invitations/{id} | Api Controller Get Invitation
[**IamApiControllerGetInvitationCodeInfo**](IamOrganizationsAPI.md#IamApiControllerGetInvitationCodeInfo) | **Get** /v1/iam/invitation-infos/{id} | Api Controller Get Invitation Code Info
[**IamApiControllerGetInvitations**](IamOrganizationsAPI.md#IamApiControllerGetInvitations) | **Get** /v1/iam/invitations | Api Controller Get Invitations
[**IamApiControllerGetOrganization**](IamOrganizationsAPI.md#IamApiControllerGetOrganization) | **Get** /v1/iam/organizations/{id} | Api Controller Get Organization
[**IamApiControllerGetOrganizationNames**](IamOrganizationsAPI.md#IamApiControllerGetOrganizationNames) | **Get** /v1/iam/organizations/names | Api Controller Get Organization Names
[**IamApiControllerGetOrganizations**](IamOrganizationsAPI.md#IamApiControllerGetOrganizations) | **Get** /v1/iam/organizations | Api Controller Get Organizations
[**IamApiControllerUpdateGroup**](IamOrganizationsAPI.md#IamApiControllerUpdateGroup) | **Put** /v1/iam/groups/{id} | Api Controller Update Group
[**IamApiControllerUpdateInvitation**](IamOrganizationsAPI.md#IamApiControllerUpdateInvitation) | **Put** /v1/iam/invitations/{id} | Api Controller Update Invitation
[**IamApiControllerUpdateOrganization**](IamOrganizationsAPI.md#IamApiControllerUpdateOrganization) | **Put** /v1/iam/organizations/{id} | Api Controller Update Organization
[**IamApiControllerVerifyInvitationGet**](IamOrganizationsAPI.md#IamApiControllerVerifyInvitationGet) | **Get** /v1/iam/invitations/verify | Api Controller Verify Invitation
[**IamApiControllerVerifyInvitationPost**](IamOrganizationsAPI.md#IamApiControllerVerifyInvitationPost) | **Post** /v1/iam/invitations/send | Api Controller Verify Invitation



## IamApiControllerAddGroup

> IamControllersResponse IamApiControllerAddGroup(ctx).IamObjectGroup(iamObjectGroup).Execute()

Api Controller Add Group



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
	iamObjectGroup := *openapiclient.NewIamObjectGroup() // IamObjectGroup | The details of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerAddGroup(context.Background()).IamObjectGroup(iamObjectGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerAddGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddGroup`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerAddGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectGroup** | [**IamObjectGroup**](IamObjectGroup.md) | The details of the group | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddInvitation

> IamControllersResponse IamApiControllerAddInvitation(ctx).IamObjectInvitation(iamObjectInvitation).Execute()

Api Controller Add Invitation



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
	iamObjectInvitation := *openapiclient.NewIamObjectInvitation() // IamObjectInvitation | The details of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerAddInvitation(context.Background()).IamObjectInvitation(iamObjectInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerAddInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddInvitation`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerAddInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectInvitation** | [**IamObjectInvitation**](IamObjectInvitation.md) | The details of the invitation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerAddOrganization

> IamControllersResponse IamApiControllerAddOrganization(ctx).IamObjectOrganization(iamObjectOrganization).Execute()

Api Controller Add Organization



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
	iamObjectOrganization := *openapiclient.NewIamObjectOrganization() // IamObjectOrganization | The details of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerAddOrganization(context.Background()).IamObjectOrganization(iamObjectOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerAddOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerAddOrganization`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerAddOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerAddOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iamObjectOrganization** | [**IamObjectOrganization**](IamObjectOrganization.md) | The details of the organization | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteGroup

> IamControllersResponse IamApiControllerDeleteGroup(ctx, id).IamObjectGroup(iamObjectGroup).Execute()

Api Controller Delete Group



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectGroup := *openapiclient.NewIamObjectGroup() // IamObjectGroup | The details of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerDeleteGroup(context.Background(), id).IamObjectGroup(iamObjectGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerDeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteGroup`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerDeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectGroup** | [**IamObjectGroup**](IamObjectGroup.md) | The details of the group | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteInvitation

> IamControllersResponse IamApiControllerDeleteInvitation(ctx, id).IamObjectInvitation(iamObjectInvitation).Execute()

Api Controller Delete Invitation



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectInvitation := *openapiclient.NewIamObjectInvitation() // IamObjectInvitation | The details of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerDeleteInvitation(context.Background(), id).IamObjectInvitation(iamObjectInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerDeleteInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteInvitation`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerDeleteInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectInvitation** | [**IamObjectInvitation**](IamObjectInvitation.md) | The details of the invitation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerDeleteOrganization

> IamControllersResponse IamApiControllerDeleteOrganization(ctx, id).IamObjectOrganization(iamObjectOrganization).Execute()

Api Controller Delete Organization



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
	id := "id_example" // string | Resource identifier (owner/name)
	iamObjectOrganization := *openapiclient.NewIamObjectOrganization() // IamObjectOrganization | The details of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerDeleteOrganization(context.Background(), id).IamObjectOrganization(iamObjectOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerDeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerDeleteOrganization`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerDeleteOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerDeleteOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectOrganization** | [**IamObjectOrganization**](IamObjectOrganization.md) | The details of the organization | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetDefaultApplication

> IamControllersResponse IamApiControllerGetDefaultApplication(ctx).Id(id).Execute()

Api Controller Get Default Application



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
	id := "id_example" // string | organization id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetDefaultApplication(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetDefaultApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetDefaultApplication`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetDefaultApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetDefaultApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | organization id | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGroup

> IamObjectGroup IamApiControllerGetGroup(ctx, id).Execute()

Api Controller Get Group



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
	id := "id_example" // string | The id ( owner/name ) of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGroup`: IamObjectGroup
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectGroup**](IamObjectGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetGroups

> []IamObjectGroup IamApiControllerGetGroups(ctx).Owner(owner).Execute()

Api Controller Get Groups



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
	owner := "owner_example" // string | The owner of groups

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetGroups(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetGroups`: []IamObjectGroup
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of groups | 

### Return type

[**[]IamObjectGroup**](IamObjectGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetInvitation

> IamObjectInvitation IamApiControllerGetInvitation(ctx, id).Execute()

Api Controller Get Invitation



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
	id := "id_example" // string | The id ( owner/name ) of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetInvitation(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetInvitation`: IamObjectInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectInvitation**](IamObjectInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetInvitationCodeInfo

> IamObjectInvitation IamApiControllerGetInvitationCodeInfo(ctx, id).Code(code).Execute()

Api Controller Get Invitation Code Info



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
	code := "code_example" // string | Invitation code
	id := "id_example" // string | Resource identifier (owner/name)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetInvitationCodeInfo(context.Background(), id).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetInvitationCodeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetInvitationCodeInfo`: IamObjectInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetInvitationCodeInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Resource identifier (owner/name) | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetInvitationCodeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Invitation code | 


### Return type

[**IamObjectInvitation**](IamObjectInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetInvitations

> []IamObjectInvitation IamApiControllerGetInvitations(ctx).Owner(owner).Execute()

Api Controller Get Invitations



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
	owner := "owner_example" // string | The owner of invitations

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetInvitations(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetInvitations`: []IamObjectInvitation
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of invitations | 

### Return type

[**[]IamObjectInvitation**](IamObjectInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrganization

> IamObjectOrganization IamApiControllerGetOrganization(ctx, id).Execute()

Api Controller Get Organization



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
	id := "id_example" // string | organization id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetOrganization(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrganization`: IamObjectOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | organization id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IamObjectOrganization**](IamObjectOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrganizationNames

> []IamObjectOrganization IamApiControllerGetOrganizationNames(ctx).Owner(owner).Execute()

Api Controller Get Organization Names



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
	owner := "owner_example" // string | owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetOrganizationNames(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetOrganizationNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrganizationNames`: []IamObjectOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetOrganizationNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrganizationNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | owner | 

### Return type

[**[]IamObjectOrganization**](IamObjectOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerGetOrganizations

> []IamObjectOrganization IamApiControllerGetOrganizations(ctx).Owner(owner).Execute()

Api Controller Get Organizations



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
	owner := "owner_example" // string | owner

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerGetOrganizations(context.Background()).Owner(owner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerGetOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerGetOrganizations`: []IamObjectOrganization
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerGetOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerGetOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | owner | 

### Return type

[**[]IamObjectOrganization**](IamObjectOrganization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateGroup

> IamControllersResponse IamApiControllerUpdateGroup(ctx, id).IamObjectGroup(iamObjectGroup).Execute()

Api Controller Update Group



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
	id := "id_example" // string | The id ( owner/name ) of the group
	iamObjectGroup := *openapiclient.NewIamObjectGroup() // IamObjectGroup | The details of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerUpdateGroup(context.Background(), id).IamObjectGroup(iamObjectGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerUpdateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateGroup`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerUpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectGroup** | [**IamObjectGroup**](IamObjectGroup.md) | The details of the group | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateInvitation

> IamControllersResponse IamApiControllerUpdateInvitation(ctx, id).IamObjectInvitation(iamObjectInvitation).Execute()

Api Controller Update Invitation



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
	id := "id_example" // string | The id ( owner/name ) of the invitation
	iamObjectInvitation := *openapiclient.NewIamObjectInvitation() // IamObjectInvitation | The details of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerUpdateInvitation(context.Background(), id).IamObjectInvitation(iamObjectInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerUpdateInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateInvitation`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerUpdateInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the invitation | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectInvitation** | [**IamObjectInvitation**](IamObjectInvitation.md) | The details of the invitation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerUpdateOrganization

> IamControllersResponse IamApiControllerUpdateOrganization(ctx, id).IamObjectOrganization(iamObjectOrganization).Execute()

Api Controller Update Organization



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
	id := "id_example" // string | The id ( owner/name ) of the organization
	iamObjectOrganization := *openapiclient.NewIamObjectOrganization() // IamObjectOrganization | The details of the organization

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerUpdateOrganization(context.Background(), id).IamObjectOrganization(iamObjectOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerUpdateOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerUpdateOrganization`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerUpdateOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id ( owner/name ) of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerUpdateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iamObjectOrganization** | [**IamObjectOrganization**](IamObjectOrganization.md) | The details of the organization | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyInvitationGet

> IamControllersResponse IamApiControllerVerifyInvitationGet(ctx).Id(id).Execute()

Api Controller Verify Invitation



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
	id := "id_example" // string | The id ( owner/name ) of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerVerifyInvitationGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerVerifyInvitationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyInvitationGet`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerVerifyInvitationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyInvitationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the invitation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IamApiControllerVerifyInvitationPost

> IamControllersResponse IamApiControllerVerifyInvitationPost(ctx).Id(id).RequestBody(requestBody).Execute()

Api Controller Verify Invitation



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
	id := "id_example" // string | The id ( owner/name ) of the invitation
	requestBody := []string{"Property_example"} // []string | The details of the invitation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamOrganizationsAPI.IamApiControllerVerifyInvitationPost(context.Background()).Id(id).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamOrganizationsAPI.IamApiControllerVerifyInvitationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IamApiControllerVerifyInvitationPost`: IamControllersResponse
	fmt.Fprintf(os.Stdout, "Response from `IamOrganizationsAPI.IamApiControllerVerifyInvitationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIamApiControllerVerifyInvitationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The id ( owner/name ) of the invitation | 
 **requestBody** | **[]string** | The details of the invitation | 

### Return type

[**IamControllersResponse**](IamControllersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

