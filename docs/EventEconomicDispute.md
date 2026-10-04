# EventEconomicDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when the dispute was filed, unix seconds. | [optional] 
**By** | Pointer to **string** | By is the org that disputes it: one of the event&#39;s two parties. | [optional] 
**Event** | Pointer to **string** | Event is the id of the event disputed. | [optional] 
**Reason** | Pointer to **string** | Reason is what the party says is wrong. | [optional] 

## Methods

### NewEventEconomicDispute

`func NewEventEconomicDispute() *EventEconomicDispute`

NewEventEconomicDispute instantiates a new EventEconomicDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEconomicDisputeWithDefaults

`func NewEventEconomicDisputeWithDefaults() *EventEconomicDispute`

NewEventEconomicDisputeWithDefaults instantiates a new EventEconomicDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *EventEconomicDispute) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *EventEconomicDispute) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *EventEconomicDispute) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *EventEconomicDispute) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *EventEconomicDispute) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *EventEconomicDispute) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *EventEconomicDispute) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *EventEconomicDispute) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetEvent

`func (o *EventEconomicDispute) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventEconomicDispute) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventEconomicDispute) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventEconomicDispute) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetReason

`func (o *EventEconomicDispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventEconomicDispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventEconomicDispute) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventEconomicDispute) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


