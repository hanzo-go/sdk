# O11yO11yTraceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]O11yO11yEvent**](O11yO11yEvent.md) | Events are the error events carrying the trace id. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace that was read. | [optional] 

## Methods

### NewO11yO11yTraceDetail

`func NewO11yO11yTraceDetail() *O11yO11yTraceDetail`

NewO11yO11yTraceDetail instantiates a new O11yO11yTraceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTraceDetailWithDefaults

`func NewO11yO11yTraceDetailWithDefaults() *O11yO11yTraceDetail`

NewO11yO11yTraceDetailWithDefaults instantiates a new O11yO11yTraceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *O11yO11yTraceDetail) GetEvents() []O11yO11yEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *O11yO11yTraceDetail) GetEventsOk() (*[]O11yO11yEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *O11yO11yTraceDetail) SetEvents(v []O11yO11yEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *O11yO11yTraceDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yTraceDetail) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yTraceDetail) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yTraceDetail) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yTraceDetail) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


