# EventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProductEvent**](ProductEvent.md) | Data is the events, newest first. Empty rather than absent when there are none. | [optional] 

## Methods

### NewEventList

`func NewEventList() *EventList`

NewEventList instantiates a new EventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventListWithDefaults

`func NewEventListWithDefaults() *EventList`

NewEventListWithDefaults instantiates a new EventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventList) GetData() []ProductEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventList) GetDataOk() (*[]ProductEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventList) SetData(v []ProductEvent)`

SetData sets Data field to given value.

### HasData

`func (o *EventList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


