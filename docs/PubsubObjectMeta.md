# PubsubObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Object name/path | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Chunks** | Pointer to **int32** |  | [optional] 
**Digest** | Pointer to **string** | SHA-256 digest | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPubsubObjectMeta

`func NewPubsubObjectMeta() *PubsubObjectMeta`

NewPubsubObjectMeta instantiates a new PubsubObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubObjectMetaWithDefaults

`func NewPubsubObjectMetaWithDefaults() *PubsubObjectMeta`

NewPubsubObjectMetaWithDefaults instantiates a new PubsubObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PubsubObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PubsubObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PubsubObjectMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PubsubObjectMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PubsubObjectMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PubsubObjectMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PubsubObjectMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PubsubObjectMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *PubsubObjectMeta) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PubsubObjectMeta) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PubsubObjectMeta) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PubsubObjectMeta) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetChunks

`func (o *PubsubObjectMeta) GetChunks() int32`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *PubsubObjectMeta) GetChunksOk() (*int32, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *PubsubObjectMeta) SetChunks(v int32)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *PubsubObjectMeta) HasChunks() bool`

HasChunks returns a boolean if a field has been set.

### GetDigest

`func (o *PubsubObjectMeta) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PubsubObjectMeta) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PubsubObjectMeta) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PubsubObjectMeta) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetHeaders

`func (o *PubsubObjectMeta) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PubsubObjectMeta) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PubsubObjectMeta) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PubsubObjectMeta) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetModified

`func (o *PubsubObjectMeta) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *PubsubObjectMeta) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *PubsubObjectMeta) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *PubsubObjectMeta) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


