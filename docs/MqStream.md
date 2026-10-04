# MqStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MqConfig**](MqConfig.md) | Config is the stream&#39;s configuration. | [optional] 
**Created** | Pointer to **time.Time** | Created is when the stream was created. | [optional] 
**Name** | Pointer to **string** | Name is the stream name within the org. | [optional] 
**State** | Pointer to [**MqState**](MqState.md) | State is the stream&#39;s current state. | [optional] 

## Methods

### NewMqStream

`func NewMqStream() *MqStream`

NewMqStream instantiates a new MqStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqStreamWithDefaults

`func NewMqStreamWithDefaults() *MqStream`

NewMqStreamWithDefaults instantiates a new MqStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MqStream) GetConfig() MqConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MqStream) GetConfigOk() (*MqConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MqStream) SetConfig(v MqConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MqStream) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreated

`func (o *MqStream) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MqStream) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MqStream) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MqStream) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetName

`func (o *MqStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqStream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *MqStream) GetState() MqState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MqStream) GetStateOk() (*MqState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MqStream) SetState(v MqState)`

SetState sets State field to given value.

### HasState

`func (o *MqStream) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


