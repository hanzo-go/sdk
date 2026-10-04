# EventLoss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exhausted** | Pointer to **int64** | Exhausted counts facts the bus abandoned after maxDeliver failed inserts. | [optional] 
**Undecodable** | Pointer to **int64** | Undecodable counts messages acked without landing because they did not parse. | [optional] 

## Methods

### NewEventLoss

`func NewEventLoss() *EventLoss`

NewEventLoss instantiates a new EventLoss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventLossWithDefaults

`func NewEventLossWithDefaults() *EventLoss`

NewEventLossWithDefaults instantiates a new EventLoss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExhausted

`func (o *EventLoss) GetExhausted() int64`

GetExhausted returns the Exhausted field if non-nil, zero value otherwise.

### GetExhaustedOk

`func (o *EventLoss) GetExhaustedOk() (*int64, bool)`

GetExhaustedOk returns a tuple with the Exhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExhausted

`func (o *EventLoss) SetExhausted(v int64)`

SetExhausted sets Exhausted field to given value.

### HasExhausted

`func (o *EventLoss) HasExhausted() bool`

HasExhausted returns a boolean if a field has been set.

### GetUndecodable

`func (o *EventLoss) GetUndecodable() int64`

GetUndecodable returns the Undecodable field if non-nil, zero value otherwise.

### GetUndecodableOk

`func (o *EventLoss) GetUndecodableOk() (*int64, bool)`

GetUndecodableOk returns a tuple with the Undecodable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndecodable

`func (o *EventLoss) SetUndecodable(v int64)`

SetUndecodable sets Undecodable field to given value.

### HasUndecodable

`func (o *EventLoss) HasUndecodable() bool`

HasUndecodable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


