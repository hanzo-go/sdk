# PubsubStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**PubsubStreamConfig**](PubsubStreamConfig.md) |  | [optional] 
**State** | Pointer to [**PubsubStreamInfoState**](PubsubStreamInfoState.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPubsubStreamInfo

`func NewPubsubStreamInfo() *PubsubStreamInfo`

NewPubsubStreamInfo instantiates a new PubsubStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubStreamInfoWithDefaults

`func NewPubsubStreamInfoWithDefaults() *PubsubStreamInfo`

NewPubsubStreamInfoWithDefaults instantiates a new PubsubStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PubsubStreamInfo) GetConfig() PubsubStreamConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PubsubStreamInfo) GetConfigOk() (*PubsubStreamConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PubsubStreamInfo) SetConfig(v PubsubStreamConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PubsubStreamInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetState

`func (o *PubsubStreamInfo) GetState() PubsubStreamInfoState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PubsubStreamInfo) GetStateOk() (*PubsubStreamInfoState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PubsubStreamInfo) SetState(v PubsubStreamInfoState)`

SetState sets State field to given value.

### HasState

`func (o *PubsubStreamInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreated

`func (o *PubsubStreamInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PubsubStreamInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PubsubStreamInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PubsubStreamInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


