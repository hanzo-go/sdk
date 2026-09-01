# \SpaceAPI

All URIs are relative to *https://api.hanzo.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSpaceBySpaceDrivesByDrive**](SpaceAPI.md#DeleteSpaceBySpaceDrivesByDrive) | **Delete** /v1/space/{space}/drives/{drive} | Removes an EMPTY drive and answers 204.
[**GetSpaceBySpaceDrives**](SpaceAPI.md#GetSpaceBySpaceDrives) | **Get** /v1/space/{space}/drives | Lists a space&#39;s drives.
[**GetSpaceBySpaceDrivesByDriveFiles**](SpaceAPI.md#GetSpaceBySpaceDrivesByDriveFiles) | **Get** /v1/space/{space}/drives/{drive}/files | Lists one folder level of a drive.
[**GetSpaceHealth**](SpaceAPI.md#GetSpaceHealth) | **Get** /v1/space/health | Health reports whether this deployment can serve spaces, drives and files.
[**GetSpaceSpaces**](SpaceAPI.md#GetSpaceSpaces) | **Get** /v1/space/spaces | Lists the caller org&#39;s own spaces.
[**PostSpaceBySpaceDrives**](SpaceAPI.md#PostSpaceBySpaceDrives) | **Post** /v1/space/{space}/drives | Makes a new drive in a space and answers 201 with it.
[**PostSpaceSpaces**](SpaceAPI.md#PostSpaceSpaces) | **Post** /v1/space/spaces | Makes a new space for the caller&#39;s org and answers 201 with it.



## DeleteSpaceBySpaceDrivesByDrive

> DeleteSpaceBySpaceDrivesByDrive(ctx, space, drive).Execute()

Removes an EMPTY drive and answers 204.



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
	space := "space_example" // string | Space is the space's name, from the path.
	drive := "drive_example" // string | Drive is the drive's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpaceAPI.DeleteSpaceBySpaceDrivesByDrive(context.Background(), space, drive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.DeleteSpaceBySpaceDrivesByDrive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Space is the space&#39;s name, from the path. | 
**drive** | **string** | Drive is the drive&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSpaceBySpaceDrivesByDriveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetSpaceBySpaceDrives

> DriveList GetSpaceBySpaceDrives(ctx, space).Execute()

Lists a space's drives.



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
	space := "space_example" // string | Space is the space's name, from the path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.GetSpaceBySpaceDrives(context.Background(), space).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaceBySpaceDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceBySpaceDrives`: DriveList
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaceBySpaceDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Space is the space&#39;s name, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceBySpaceDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DriveList**](DriveList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceBySpaceDrivesByDriveFiles

> FileList GetSpaceBySpaceDrivesByDriveFiles(ctx, space, drive).Folder(folder).Recursive(recursive).Execute()

Lists one folder level of a drive.



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
	space := "space_example" // string | Space is the space to list in, from the path.
	drive := "drive_example" // string | Drive is the drive to list, from the path.
	folder := "folder_example" // string |  (optional)
	recursive := "recursive_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.GetSpaceBySpaceDrivesByDriveFiles(context.Background(), space, drive).Folder(folder).Recursive(recursive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaceBySpaceDrivesByDriveFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceBySpaceDrivesByDriveFiles`: FileList
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaceBySpaceDrivesByDriveFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Space is the space to list in, from the path. | 
**drive** | **string** | Drive is the drive to list, from the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceBySpaceDrivesByDriveFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **folder** | **string** |  | 
 **recursive** | **string** |  | 

### Return type

[**FileList**](FileList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceHealth

> SpaceHealth GetSpaceHealth(ctx).Execute()

Health reports whether this deployment can serve spaces, drives and files.



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
	resp, r, err := apiClient.SpaceAPI.GetSpaceHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaceHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceHealth`: SpaceHealth
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaceHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceHealthRequest struct via the builder pattern


### Return type

[**SpaceHealth**](SpaceHealth.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpaceSpaces

> SpaceList GetSpaceSpaces(ctx).Execute()

Lists the caller org's own spaces.



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
	resp, r, err := apiClient.SpaceAPI.GetSpaceSpaces(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.GetSpaceSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpaceSpaces`: SpaceList
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.GetSpaceSpaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpaceSpacesRequest struct via the builder pattern


### Return type

[**SpaceList**](SpaceList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSpaceBySpaceDrives

> DriveItem PostSpaceBySpaceDrives(ctx, space).DriveIn(driveIn).Execute()

Makes a new drive in a space and answers 201 with it.



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
	space := "space_example" // string | Space is the space to create the drive in, from the path. It carries NO `url:\"-\"`, unlike the field below it, and the difference is the whole reason both tags are written out: zip's binder skips a field tagged \"-\" for EVERY URL source, path params included, so a path-borne value that carried it would arrive empty and the create would refuse a perfectly good address.
	driveIn := *openapiclient.NewDriveIn() // DriveIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.PostSpaceBySpaceDrives(context.Background(), space).DriveIn(driveIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.PostSpaceBySpaceDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSpaceBySpaceDrives`: DriveItem
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.PostSpaceBySpaceDrives`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**space** | **string** | Space is the space to create the drive in, from the path. It carries NO &#x60;url:\&quot;-\&quot;&#x60;, unlike the field below it, and the difference is the whole reason both tags are written out: zip&#39;s binder skips a field tagged \&quot;-\&quot; for EVERY URL source, path params included, so a path-borne value that carried it would arrive empty and the create would refuse a perfectly good address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSpaceBySpaceDrivesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **driveIn** | [**DriveIn**](DriveIn.md) |  | 

### Return type

[**DriveItem**](DriveItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSpaceSpaces

> SpaceItem PostSpaceSpaces(ctx).SpaceIn(spaceIn).Execute()

Makes a new space for the caller's org and answers 201 with it.



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
	spaceIn := *openapiclient.NewSpaceIn() // SpaceIn | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpaceAPI.PostSpaceSpaces(context.Background()).SpaceIn(spaceIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpaceAPI.PostSpaceSpaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSpaceSpaces`: SpaceItem
	fmt.Fprintf(os.Stdout, "Response from `SpaceAPI.PostSpaceSpaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSpaceSpacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spaceIn** | [**SpaceIn**](SpaceIn.md) |  | 

### Return type

[**SpaceItem**](SpaceItem.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

