# StreamConsumerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **int32** |  | [optional] 
**Offsets** | Pointer to [**[]StreamConsumerGroupOffsetsInner**](StreamConsumerGroupOffsetsInner.md) |  | [optional] 

## Methods

### NewStreamConsumerGroup

`func NewStreamConsumerGroup() *StreamConsumerGroup`

NewStreamConsumerGroup instantiates a new StreamConsumerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConsumerGroupWithDefaults

`func NewStreamConsumerGroupWithDefaults() *StreamConsumerGroup`

NewStreamConsumerGroupWithDefaults instantiates a new StreamConsumerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *StreamConsumerGroup) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamConsumerGroup) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamConsumerGroup) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *StreamConsumerGroup) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetState

`func (o *StreamConsumerGroup) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamConsumerGroup) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamConsumerGroup) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamConsumerGroup) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMembers

`func (o *StreamConsumerGroup) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *StreamConsumerGroup) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *StreamConsumerGroup) SetMembers(v int32)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *StreamConsumerGroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOffsets

`func (o *StreamConsumerGroup) GetOffsets() []StreamConsumerGroupOffsetsInner`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *StreamConsumerGroup) GetOffsetsOk() (*[]StreamConsumerGroupOffsetsInner, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *StreamConsumerGroup) SetOffsets(v []StreamConsumerGroupOffsetsInner)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *StreamConsumerGroup) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


