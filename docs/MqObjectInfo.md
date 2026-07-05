# MqObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Object name. | [optional] 
**Description** | Pointer to **string** | Optional description. | [optional] 
**Size** | Pointer to **int64** | Object size in bytes. | [optional] 
**Chunks** | Pointer to **int32** | Number of chunks. | [optional] 
**Digest** | Pointer to **string** | SHA-256 digest of the object. | [optional] 
**Modified** | Pointer to **time.Time** | Last modification timestamp. | [optional] 
**Deleted** | Pointer to **bool** | True if the object has been deleted. | [optional] 
**Headers** | Pointer to **map[string][]string** | Optional metadata headers. | [optional] 

## Methods

### NewMqObjectInfo

`func NewMqObjectInfo() *MqObjectInfo`

NewMqObjectInfo instantiates a new MqObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqObjectInfoWithDefaults

`func NewMqObjectInfoWithDefaults() *MqObjectInfo`

NewMqObjectInfoWithDefaults instantiates a new MqObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MqObjectInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MqObjectInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MqObjectInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MqObjectInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *MqObjectInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MqObjectInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MqObjectInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MqObjectInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetChunks

`func (o *MqObjectInfo) GetChunks() int32`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *MqObjectInfo) GetChunksOk() (*int32, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *MqObjectInfo) SetChunks(v int32)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *MqObjectInfo) HasChunks() bool`

HasChunks returns a boolean if a field has been set.

### GetDigest

`func (o *MqObjectInfo) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *MqObjectInfo) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *MqObjectInfo) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *MqObjectInfo) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetModified

`func (o *MqObjectInfo) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MqObjectInfo) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MqObjectInfo) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *MqObjectInfo) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetDeleted

`func (o *MqObjectInfo) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *MqObjectInfo) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *MqObjectInfo) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *MqObjectInfo) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetHeaders

`func (o *MqObjectInfo) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MqObjectInfo) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MqObjectInfo) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MqObjectInfo) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


