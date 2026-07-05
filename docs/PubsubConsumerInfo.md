# PubsubConsumerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**PubsubConsumerConfig**](PubsubConsumerConfig.md) |  | [optional] 
**Delivered** | Pointer to [**PubsubConsumerInfoDelivered**](PubsubConsumerInfoDelivered.md) |  | [optional] 
**AckFloor** | Pointer to [**PubsubConsumerInfoDelivered**](PubsubConsumerInfoDelivered.md) |  | [optional] 
**NumPending** | Pointer to **int32** |  | [optional] 
**NumRedelivered** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubConsumerInfo

`func NewPubsubConsumerInfo() *PubsubConsumerInfo`

NewPubsubConsumerInfo instantiates a new PubsubConsumerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubConsumerInfoWithDefaults

`func NewPubsubConsumerInfoWithDefaults() *PubsubConsumerInfo`

NewPubsubConsumerInfoWithDefaults instantiates a new PubsubConsumerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PubsubConsumerInfo) GetConfig() PubsubConsumerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PubsubConsumerInfo) GetConfigOk() (*PubsubConsumerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PubsubConsumerInfo) SetConfig(v PubsubConsumerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PubsubConsumerInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDelivered

`func (o *PubsubConsumerInfo) GetDelivered() PubsubConsumerInfoDelivered`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *PubsubConsumerInfo) GetDeliveredOk() (*PubsubConsumerInfoDelivered, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *PubsubConsumerInfo) SetDelivered(v PubsubConsumerInfoDelivered)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *PubsubConsumerInfo) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetAckFloor

`func (o *PubsubConsumerInfo) GetAckFloor() PubsubConsumerInfoDelivered`

GetAckFloor returns the AckFloor field if non-nil, zero value otherwise.

### GetAckFloorOk

`func (o *PubsubConsumerInfo) GetAckFloorOk() (*PubsubConsumerInfoDelivered, bool)`

GetAckFloorOk returns a tuple with the AckFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckFloor

`func (o *PubsubConsumerInfo) SetAckFloor(v PubsubConsumerInfoDelivered)`

SetAckFloor sets AckFloor field to given value.

### HasAckFloor

`func (o *PubsubConsumerInfo) HasAckFloor() bool`

HasAckFloor returns a boolean if a field has been set.

### GetNumPending

`func (o *PubsubConsumerInfo) GetNumPending() int32`

GetNumPending returns the NumPending field if non-nil, zero value otherwise.

### GetNumPendingOk

`func (o *PubsubConsumerInfo) GetNumPendingOk() (*int32, bool)`

GetNumPendingOk returns a tuple with the NumPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPending

`func (o *PubsubConsumerInfo) SetNumPending(v int32)`

SetNumPending sets NumPending field to given value.

### HasNumPending

`func (o *PubsubConsumerInfo) HasNumPending() bool`

HasNumPending returns a boolean if a field has been set.

### GetNumRedelivered

`func (o *PubsubConsumerInfo) GetNumRedelivered() int32`

GetNumRedelivered returns the NumRedelivered field if non-nil, zero value otherwise.

### GetNumRedeliveredOk

`func (o *PubsubConsumerInfo) GetNumRedeliveredOk() (*int32, bool)`

GetNumRedeliveredOk returns a tuple with the NumRedelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumRedelivered

`func (o *PubsubConsumerInfo) SetNumRedelivered(v int32)`

SetNumRedelivered sets NumRedelivered field to given value.

### HasNumRedelivered

`func (o *PubsubConsumerInfo) HasNumRedelivered() bool`

HasNumRedelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


