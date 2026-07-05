# MqObjectStoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**MqObjectStoreConfig**](MqObjectStoreConfig.md) |  | [optional] 
**Size** | Pointer to **int32** | Total bytes used. | [optional] 
**Objects** | Pointer to **int32** | Number of objects in the store. | [optional] 
**Chunks** | Pointer to **int32** | Total number of chunks across all objects. | [optional] 
**BackingStream** | Pointer to **string** | Name of the underlying JetStream stream. | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMqObjectStoreInfo

`func NewMqObjectStoreInfo() *MqObjectStoreInfo`

NewMqObjectStoreInfo instantiates a new MqObjectStoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqObjectStoreInfoWithDefaults

`func NewMqObjectStoreInfoWithDefaults() *MqObjectStoreInfo`

NewMqObjectStoreInfoWithDefaults instantiates a new MqObjectStoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqObjectStoreInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqObjectStoreInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqObjectStoreInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqObjectStoreInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *MqObjectStoreInfo) GetConfig() MqObjectStoreConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MqObjectStoreInfo) GetConfigOk() (*MqObjectStoreConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MqObjectStoreInfo) SetConfig(v MqObjectStoreConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MqObjectStoreInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSize

`func (o *MqObjectStoreInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MqObjectStoreInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MqObjectStoreInfo) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MqObjectStoreInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetObjects

`func (o *MqObjectStoreInfo) GetObjects() int32`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *MqObjectStoreInfo) GetObjectsOk() (*int32, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *MqObjectStoreInfo) SetObjects(v int32)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *MqObjectStoreInfo) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetChunks

`func (o *MqObjectStoreInfo) GetChunks() int32`

GetChunks returns the Chunks field if non-nil, zero value otherwise.

### GetChunksOk

`func (o *MqObjectStoreInfo) GetChunksOk() (*int32, bool)`

GetChunksOk returns a tuple with the Chunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunks

`func (o *MqObjectStoreInfo) SetChunks(v int32)`

SetChunks sets Chunks field to given value.

### HasChunks

`func (o *MqObjectStoreInfo) HasChunks() bool`

HasChunks returns a boolean if a field has been set.

### GetBackingStream

`func (o *MqObjectStoreInfo) GetBackingStream() string`

GetBackingStream returns the BackingStream field if non-nil, zero value otherwise.

### GetBackingStreamOk

`func (o *MqObjectStoreInfo) GetBackingStreamOk() (*string, bool)`

GetBackingStreamOk returns a tuple with the BackingStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingStream

`func (o *MqObjectStoreInfo) SetBackingStream(v string)`

SetBackingStream sets BackingStream field to given value.

### HasBackingStream

`func (o *MqObjectStoreInfo) HasBackingStream() bool`

HasBackingStream returns a boolean if a field has been set.

### GetCreated

`func (o *MqObjectStoreInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MqObjectStoreInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MqObjectStoreInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MqObjectStoreInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


