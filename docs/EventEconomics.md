# EventEconomics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disputes** | Pointer to [**[]EventEconomicDispute**](EventEconomicDispute.md) | Disputes are what either party has disputed about the events in Rows. A dispute changes nothing about the event it names. | [optional] 
**Next** | Pointer to **string** | Next continues the read; empty on the last page. | [optional] 
**Rows** | Pointer to [**[]EventEconomicEvent**](EventEconomicEvent.md) | Rows are the events. | [optional] 
**Superseded** | Pointer to **[]string** | Superseded are the ids in Rows that a correction restates. | [optional] 

## Methods

### NewEventEconomics

`func NewEventEconomics() *EventEconomics`

NewEventEconomics instantiates a new EventEconomics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEconomicsWithDefaults

`func NewEventEconomicsWithDefaults() *EventEconomics`

NewEventEconomicsWithDefaults instantiates a new EventEconomics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisputes

`func (o *EventEconomics) GetDisputes() []EventEconomicDispute`

GetDisputes returns the Disputes field if non-nil, zero value otherwise.

### GetDisputesOk

`func (o *EventEconomics) GetDisputesOk() (*[]EventEconomicDispute, bool)`

GetDisputesOk returns a tuple with the Disputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputes

`func (o *EventEconomics) SetDisputes(v []EventEconomicDispute)`

SetDisputes sets Disputes field to given value.

### HasDisputes

`func (o *EventEconomics) HasDisputes() bool`

HasDisputes returns a boolean if a field has been set.

### GetNext

`func (o *EventEconomics) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *EventEconomics) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *EventEconomics) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *EventEconomics) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRows

`func (o *EventEconomics) GetRows() []EventEconomicEvent`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *EventEconomics) GetRowsOk() (*[]EventEconomicEvent, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *EventEconomics) SetRows(v []EventEconomicEvent)`

SetRows sets Rows field to given value.

### HasRows

`func (o *EventEconomics) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetSuperseded

`func (o *EventEconomics) GetSuperseded() []string`

GetSuperseded returns the Superseded field if non-nil, zero value otherwise.

### GetSupersededOk

`func (o *EventEconomics) GetSupersededOk() (*[]string, bool)`

GetSupersededOk returns a tuple with the Superseded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperseded

`func (o *EventEconomics) SetSuperseded(v []string)`

SetSuperseded sets Superseded field to given value.

### HasSuperseded

`func (o *EventEconomics) HasSuperseded() bool`

HasSuperseded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


