# EventDisputeIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | Event is the id of the economic event disputed. | 
**Reason** | **string** | Reason is what the caller&#39;s org says is wrong, at most 1024 characters. | 

## Methods

### NewEventDisputeIn

`func NewEventDisputeIn(event string, reason string, ) *EventDisputeIn`

NewEventDisputeIn instantiates a new EventDisputeIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDisputeInWithDefaults

`func NewEventDisputeInWithDefaults() *EventDisputeIn`

NewEventDisputeInWithDefaults instantiates a new EventDisputeIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventDisputeIn) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventDisputeIn) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventDisputeIn) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetReason

`func (o *EventDisputeIn) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventDisputeIn) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventDisputeIn) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


