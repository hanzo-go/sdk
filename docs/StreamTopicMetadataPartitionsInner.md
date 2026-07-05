# StreamTopicMetadataPartitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Leader** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **[]int32** |  | [optional] 
**Isr** | Pointer to **[]int32** |  | [optional] 
**EarliestOffset** | Pointer to **int32** |  | [optional] 
**LatestOffset** | Pointer to **int32** |  | [optional] 

## Methods

### NewStreamTopicMetadataPartitionsInner

`func NewStreamTopicMetadataPartitionsInner() *StreamTopicMetadataPartitionsInner`

NewStreamTopicMetadataPartitionsInner instantiates a new StreamTopicMetadataPartitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamTopicMetadataPartitionsInnerWithDefaults

`func NewStreamTopicMetadataPartitionsInnerWithDefaults() *StreamTopicMetadataPartitionsInner`

NewStreamTopicMetadataPartitionsInnerWithDefaults instantiates a new StreamTopicMetadataPartitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamTopicMetadataPartitionsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamTopicMetadataPartitionsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamTopicMetadataPartitionsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StreamTopicMetadataPartitionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeader

`func (o *StreamTopicMetadataPartitionsInner) GetLeader() int32`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *StreamTopicMetadataPartitionsInner) GetLeaderOk() (*int32, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *StreamTopicMetadataPartitionsInner) SetLeader(v int32)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *StreamTopicMetadataPartitionsInner) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *StreamTopicMetadataPartitionsInner) GetReplicas() []int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *StreamTopicMetadataPartitionsInner) GetReplicasOk() (*[]int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *StreamTopicMetadataPartitionsInner) SetReplicas(v []int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *StreamTopicMetadataPartitionsInner) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetIsr

`func (o *StreamTopicMetadataPartitionsInner) GetIsr() []int32`

GetIsr returns the Isr field if non-nil, zero value otherwise.

### GetIsrOk

`func (o *StreamTopicMetadataPartitionsInner) GetIsrOk() (*[]int32, bool)`

GetIsrOk returns a tuple with the Isr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsr

`func (o *StreamTopicMetadataPartitionsInner) SetIsr(v []int32)`

SetIsr sets Isr field to given value.

### HasIsr

`func (o *StreamTopicMetadataPartitionsInner) HasIsr() bool`

HasIsr returns a boolean if a field has been set.

### GetEarliestOffset

`func (o *StreamTopicMetadataPartitionsInner) GetEarliestOffset() int32`

GetEarliestOffset returns the EarliestOffset field if non-nil, zero value otherwise.

### GetEarliestOffsetOk

`func (o *StreamTopicMetadataPartitionsInner) GetEarliestOffsetOk() (*int32, bool)`

GetEarliestOffsetOk returns a tuple with the EarliestOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestOffset

`func (o *StreamTopicMetadataPartitionsInner) SetEarliestOffset(v int32)`

SetEarliestOffset sets EarliestOffset field to given value.

### HasEarliestOffset

`func (o *StreamTopicMetadataPartitionsInner) HasEarliestOffset() bool`

HasEarliestOffset returns a boolean if a field has been set.

### GetLatestOffset

`func (o *StreamTopicMetadataPartitionsInner) GetLatestOffset() int32`

GetLatestOffset returns the LatestOffset field if non-nil, zero value otherwise.

### GetLatestOffsetOk

`func (o *StreamTopicMetadataPartitionsInner) GetLatestOffsetOk() (*int32, bool)`

GetLatestOffsetOk returns a tuple with the LatestOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestOffset

`func (o *StreamTopicMetadataPartitionsInner) SetLatestOffset(v int32)`

SetLatestOffset sets LatestOffset field to given value.

### HasLatestOffset

`func (o *StreamTopicMetadataPartitionsInner) HasLatestOffset() bool`

HasLatestOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


