# EdgeFunction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | URL-safe function identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** | Current deployed version number | [optional] 
**Runtime** | Pointer to **string** |  | [optional] [default to "deno"]
**Entrypoint** | Pointer to **string** |  | [optional] [default to "index.ts"]
**ImportMap** | Pointer to **bool** |  | [optional] 
**VerifyJwt** | Pointer to **bool** | Whether to verify JWT tokens on invocation | [optional] 
**InvokeUrl** | Pointer to **string** | Public invocation URL | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEdgeFunction

`func NewEdgeFunction() *EdgeFunction`

NewEdgeFunction instantiates a new EdgeFunction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionWithDefaults

`func NewEdgeFunctionWithDefaults() *EdgeFunction`

NewEdgeFunctionWithDefaults instantiates a new EdgeFunction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeFunction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeFunction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeFunction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EdgeFunction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *EdgeFunction) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EdgeFunction) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EdgeFunction) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EdgeFunction) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetName

`func (o *EdgeFunction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EdgeFunction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EdgeFunction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EdgeFunction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EdgeFunction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EdgeFunction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *EdgeFunction) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EdgeFunction) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EdgeFunction) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EdgeFunction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRuntime

`func (o *EdgeFunction) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *EdgeFunction) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *EdgeFunction) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *EdgeFunction) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetEntrypoint

`func (o *EdgeFunction) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *EdgeFunction) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *EdgeFunction) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *EdgeFunction) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetImportMap

`func (o *EdgeFunction) GetImportMap() bool`

GetImportMap returns the ImportMap field if non-nil, zero value otherwise.

### GetImportMapOk

`func (o *EdgeFunction) GetImportMapOk() (*bool, bool)`

GetImportMapOk returns a tuple with the ImportMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMap

`func (o *EdgeFunction) SetImportMap(v bool)`

SetImportMap sets ImportMap field to given value.

### HasImportMap

`func (o *EdgeFunction) HasImportMap() bool`

HasImportMap returns a boolean if a field has been set.

### GetVerifyJwt

`func (o *EdgeFunction) GetVerifyJwt() bool`

GetVerifyJwt returns the VerifyJwt field if non-nil, zero value otherwise.

### GetVerifyJwtOk

`func (o *EdgeFunction) GetVerifyJwtOk() (*bool, bool)`

GetVerifyJwtOk returns a tuple with the VerifyJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyJwt

`func (o *EdgeFunction) SetVerifyJwt(v bool)`

SetVerifyJwt sets VerifyJwt field to given value.

### HasVerifyJwt

`func (o *EdgeFunction) HasVerifyJwt() bool`

HasVerifyJwt returns a boolean if a field has been set.

### GetInvokeUrl

`func (o *EdgeFunction) GetInvokeUrl() string`

GetInvokeUrl returns the InvokeUrl field if non-nil, zero value otherwise.

### GetInvokeUrlOk

`func (o *EdgeFunction) GetInvokeUrlOk() (*string, bool)`

GetInvokeUrlOk returns a tuple with the InvokeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeUrl

`func (o *EdgeFunction) SetInvokeUrl(v string)`

SetInvokeUrl sets InvokeUrl field to given value.

### HasInvokeUrl

`func (o *EdgeFunction) HasInvokeUrl() bool`

HasInvokeUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EdgeFunction) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EdgeFunction) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EdgeFunction) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EdgeFunction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EdgeFunction) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EdgeFunction) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EdgeFunction) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EdgeFunction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


