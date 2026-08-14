# O11yO11yTraceSpanWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[]string** | Columns names the fields each row carries, in row order. | [optional] 
**EndTimestampMillis** | Pointer to **int32** | EndTimestampMillis is when it closes. | [optional] 
**Events** | Pointer to **[][]map[string]interface{}** | Events are the rows, each positionally matching Columns. | [optional] 
**IsSubTree** | Pointer to **bool** | IsSubTree says the window is a subtree of the trace rather than the whole of it. | [optional] 
**StartTimestampMillis** | Pointer to **int32** | StartTimestampMillis is when the window opens. | [optional] 

## Methods

### NewO11yO11yTraceSpanWindow

`func NewO11yO11yTraceSpanWindow() *O11yO11yTraceSpanWindow`

NewO11yO11yTraceSpanWindow instantiates a new O11yO11yTraceSpanWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTraceSpanWindowWithDefaults

`func NewO11yO11yTraceSpanWindowWithDefaults() *O11yO11yTraceSpanWindow`

NewO11yO11yTraceSpanWindowWithDefaults instantiates a new O11yO11yTraceSpanWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *O11yO11yTraceSpanWindow) GetColumns() []string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *O11yO11yTraceSpanWindow) GetColumnsOk() (*[]string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *O11yO11yTraceSpanWindow) SetColumns(v []string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *O11yO11yTraceSpanWindow) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEndTimestampMillis

`func (o *O11yO11yTraceSpanWindow) GetEndTimestampMillis() int32`

GetEndTimestampMillis returns the EndTimestampMillis field if non-nil, zero value otherwise.

### GetEndTimestampMillisOk

`func (o *O11yO11yTraceSpanWindow) GetEndTimestampMillisOk() (*int32, bool)`

GetEndTimestampMillisOk returns a tuple with the EndTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMillis

`func (o *O11yO11yTraceSpanWindow) SetEndTimestampMillis(v int32)`

SetEndTimestampMillis sets EndTimestampMillis field to given value.

### HasEndTimestampMillis

`func (o *O11yO11yTraceSpanWindow) HasEndTimestampMillis() bool`

HasEndTimestampMillis returns a boolean if a field has been set.

### GetEvents

`func (o *O11yO11yTraceSpanWindow) GetEvents() [][]map[string]interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *O11yO11yTraceSpanWindow) GetEventsOk() (*[][]map[string]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *O11yO11yTraceSpanWindow) SetEvents(v [][]map[string]interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *O11yO11yTraceSpanWindow) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIsSubTree

`func (o *O11yO11yTraceSpanWindow) GetIsSubTree() bool`

GetIsSubTree returns the IsSubTree field if non-nil, zero value otherwise.

### GetIsSubTreeOk

`func (o *O11yO11yTraceSpanWindow) GetIsSubTreeOk() (*bool, bool)`

GetIsSubTreeOk returns a tuple with the IsSubTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubTree

`func (o *O11yO11yTraceSpanWindow) SetIsSubTree(v bool)`

SetIsSubTree sets IsSubTree field to given value.

### HasIsSubTree

`func (o *O11yO11yTraceSpanWindow) HasIsSubTree() bool`

HasIsSubTree returns a boolean if a field has been set.

### GetStartTimestampMillis

`func (o *O11yO11yTraceSpanWindow) GetStartTimestampMillis() int32`

GetStartTimestampMillis returns the StartTimestampMillis field if non-nil, zero value otherwise.

### GetStartTimestampMillisOk

`func (o *O11yO11yTraceSpanWindow) GetStartTimestampMillisOk() (*int32, bool)`

GetStartTimestampMillisOk returns a tuple with the StartTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMillis

`func (o *O11yO11yTraceSpanWindow) SetStartTimestampMillis(v int32)`

SetStartTimestampMillis sets StartTimestampMillis field to given value.

### HasStartTimestampMillis

`func (o *O11yO11yTraceSpanWindow) HasStartTimestampMillis() bool`

HasStartTimestampMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


