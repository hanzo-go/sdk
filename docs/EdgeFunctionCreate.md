# EdgeFunctionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Name** | **string** |  | 
**VerifyJwt** | Pointer to **bool** |  | [optional] [default to true]
**ImportMap** | Pointer to **bool** |  | [optional] [default to false]
**Entrypoint** | Pointer to **string** |  | [optional] [default to "index.ts"]

## Methods

### NewEdgeFunctionCreate

`func NewEdgeFunctionCreate(slug string, name string, ) *EdgeFunctionCreate`

NewEdgeFunctionCreate instantiates a new EdgeFunctionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionCreateWithDefaults

`func NewEdgeFunctionCreateWithDefaults() *EdgeFunctionCreate`

NewEdgeFunctionCreateWithDefaults instantiates a new EdgeFunctionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *EdgeFunctionCreate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EdgeFunctionCreate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EdgeFunctionCreate) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *EdgeFunctionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EdgeFunctionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EdgeFunctionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetVerifyJwt

`func (o *EdgeFunctionCreate) GetVerifyJwt() bool`

GetVerifyJwt returns the VerifyJwt field if non-nil, zero value otherwise.

### GetVerifyJwtOk

`func (o *EdgeFunctionCreate) GetVerifyJwtOk() (*bool, bool)`

GetVerifyJwtOk returns a tuple with the VerifyJwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyJwt

`func (o *EdgeFunctionCreate) SetVerifyJwt(v bool)`

SetVerifyJwt sets VerifyJwt field to given value.

### HasVerifyJwt

`func (o *EdgeFunctionCreate) HasVerifyJwt() bool`

HasVerifyJwt returns a boolean if a field has been set.

### GetImportMap

`func (o *EdgeFunctionCreate) GetImportMap() bool`

GetImportMap returns the ImportMap field if non-nil, zero value otherwise.

### GetImportMapOk

`func (o *EdgeFunctionCreate) GetImportMapOk() (*bool, bool)`

GetImportMapOk returns a tuple with the ImportMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMap

`func (o *EdgeFunctionCreate) SetImportMap(v bool)`

SetImportMap sets ImportMap field to given value.

### HasImportMap

`func (o *EdgeFunctionCreate) HasImportMap() bool`

HasImportMap returns a boolean if a field has been set.

### GetEntrypoint

`func (o *EdgeFunctionCreate) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *EdgeFunctionCreate) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *EdgeFunctionCreate) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *EdgeFunctionCreate) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


