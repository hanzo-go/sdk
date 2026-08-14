# O11yGettableRuleStateTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yGettableRuleStateHistory**](O11yGettableRuleStateHistory.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yGettableRuleStateTimeline

`func NewO11yGettableRuleStateTimeline() *O11yGettableRuleStateTimeline`

NewO11yGettableRuleStateTimeline instantiates a new O11yGettableRuleStateTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableRuleStateTimelineWithDefaults

`func NewO11yGettableRuleStateTimelineWithDefaults() *O11yGettableRuleStateTimeline`

NewO11yGettableRuleStateTimelineWithDefaults instantiates a new O11yGettableRuleStateTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yGettableRuleStateTimeline) GetItems() []O11yGettableRuleStateHistory`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yGettableRuleStateTimeline) GetItemsOk() (*[]O11yGettableRuleStateHistory, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yGettableRuleStateTimeline) SetItems(v []O11yGettableRuleStateHistory)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yGettableRuleStateTimeline) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextCursor

`func (o *O11yGettableRuleStateTimeline) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *O11yGettableRuleStateTimeline) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *O11yGettableRuleStateTimeline) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *O11yGettableRuleStateTimeline) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetTotal

`func (o *O11yGettableRuleStateTimeline) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yGettableRuleStateTimeline) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yGettableRuleStateTimeline) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yGettableRuleStateTimeline) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


