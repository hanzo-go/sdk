# TraceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TraceView**](TraceView.md) | Data is the caller org&#39;s traces matching the filters, bounded by limit. | [optional] 

## Methods

### NewTraceList

`func NewTraceList() *TraceList`

NewTraceList instantiates a new TraceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceListWithDefaults

`func NewTraceListWithDefaults() *TraceList`

NewTraceListWithDefaults instantiates a new TraceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TraceList) GetData() []TraceView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TraceList) GetDataOk() (*[]TraceView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TraceList) SetData(v []TraceView)`

SetData sets Data field to given value.

### HasData

`func (o *TraceList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


