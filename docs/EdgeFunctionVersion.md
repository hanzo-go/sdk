# EdgeFunctionVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Entrypoint** | Pointer to **string** |  | [optional] 
**ImportMap** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEdgeFunctionVersion

`func NewEdgeFunctionVersion() *EdgeFunctionVersion`

NewEdgeFunctionVersion instantiates a new EdgeFunctionVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionVersionWithDefaults

`func NewEdgeFunctionVersionWithDefaults() *EdgeFunctionVersion`

NewEdgeFunctionVersionWithDefaults instantiates a new EdgeFunctionVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *EdgeFunctionVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EdgeFunctionVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EdgeFunctionVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EdgeFunctionVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *EdgeFunctionVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EdgeFunctionVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EdgeFunctionVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EdgeFunctionVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntrypoint

`func (o *EdgeFunctionVersion) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *EdgeFunctionVersion) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *EdgeFunctionVersion) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *EdgeFunctionVersion) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetImportMap

`func (o *EdgeFunctionVersion) GetImportMap() bool`

GetImportMap returns the ImportMap field if non-nil, zero value otherwise.

### GetImportMapOk

`func (o *EdgeFunctionVersion) GetImportMapOk() (*bool, bool)`

GetImportMapOk returns a tuple with the ImportMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMap

`func (o *EdgeFunctionVersion) SetImportMap(v bool)`

SetImportMap sets ImportMap field to given value.

### HasImportMap

`func (o *EdgeFunctionVersion) HasImportMap() bool`

HasImportMap returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EdgeFunctionVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EdgeFunctionVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EdgeFunctionVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EdgeFunctionVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


