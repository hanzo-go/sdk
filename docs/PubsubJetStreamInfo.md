# PubsubJetStreamInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to **int32** |  | [optional] 
**Storage** | Pointer to **int32** |  | [optional] 
**Streams** | Pointer to **int32** |  | [optional] 
**Consumers** | Pointer to **int32** |  | [optional] 
**Api** | Pointer to [**PubsubJetStreamInfoApi**](PubsubJetStreamInfoApi.md) |  | [optional] 

## Methods

### NewPubsubJetStreamInfo

`func NewPubsubJetStreamInfo() *PubsubJetStreamInfo`

NewPubsubJetStreamInfo instantiates a new PubsubJetStreamInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubJetStreamInfoWithDefaults

`func NewPubsubJetStreamInfoWithDefaults() *PubsubJetStreamInfo`

NewPubsubJetStreamInfoWithDefaults instantiates a new PubsubJetStreamInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *PubsubJetStreamInfo) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PubsubJetStreamInfo) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PubsubJetStreamInfo) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PubsubJetStreamInfo) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetStorage

`func (o *PubsubJetStreamInfo) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *PubsubJetStreamInfo) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *PubsubJetStreamInfo) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *PubsubJetStreamInfo) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetStreams

`func (o *PubsubJetStreamInfo) GetStreams() int32`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *PubsubJetStreamInfo) GetStreamsOk() (*int32, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *PubsubJetStreamInfo) SetStreams(v int32)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *PubsubJetStreamInfo) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetConsumers

`func (o *PubsubJetStreamInfo) GetConsumers() int32`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *PubsubJetStreamInfo) GetConsumersOk() (*int32, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *PubsubJetStreamInfo) SetConsumers(v int32)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *PubsubJetStreamInfo) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetApi

`func (o *PubsubJetStreamInfo) GetApi() PubsubJetStreamInfoApi`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *PubsubJetStreamInfo) GetApiOk() (*PubsubJetStreamInfoApi, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *PubsubJetStreamInfo) SetApi(v PubsubJetStreamInfoApi)`

SetApi sets Api field to given value.

### HasApi

`func (o *PubsubJetStreamInfo) HasApi() bool`

HasApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


