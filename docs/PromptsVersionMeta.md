# PromptsVersionMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPromptsVersionMeta

`func NewPromptsVersionMeta() *PromptsVersionMeta`

NewPromptsVersionMeta instantiates a new PromptsVersionMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptsVersionMetaWithDefaults

`func NewPromptsVersionMetaWithDefaults() *PromptsVersionMeta`

NewPromptsVersionMetaWithDefaults instantiates a new PromptsVersionMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PromptsVersionMeta) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptsVersionMeta) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptsVersionMeta) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptsVersionMeta) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *PromptsVersionMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptsVersionMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptsVersionMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptsVersionMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PromptsVersionMeta) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptsVersionMeta) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptsVersionMeta) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptsVersionMeta) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


