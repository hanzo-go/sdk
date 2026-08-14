# O11yRuleStateTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yRuleStateHistory**](O11yRuleStateHistory.md) |  | [optional] 
**Labels** | Pointer to **map[string][]string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yRuleStateTimeline

`func NewO11yRuleStateTimeline() *O11yRuleStateTimeline`

NewO11yRuleStateTimeline instantiates a new O11yRuleStateTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yRuleStateTimelineWithDefaults

`func NewO11yRuleStateTimelineWithDefaults() *O11yRuleStateTimeline`

NewO11yRuleStateTimelineWithDefaults instantiates a new O11yRuleStateTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yRuleStateTimeline) GetItems() []O11yRuleStateHistory`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yRuleStateTimeline) GetItemsOk() (*[]O11yRuleStateHistory, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yRuleStateTimeline) SetItems(v []O11yRuleStateHistory)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yRuleStateTimeline) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLabels

`func (o *O11yRuleStateTimeline) GetLabels() map[string][]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *O11yRuleStateTimeline) GetLabelsOk() (*map[string][]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *O11yRuleStateTimeline) SetLabels(v map[string][]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *O11yRuleStateTimeline) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTotal

`func (o *O11yRuleStateTimeline) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yRuleStateTimeline) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yRuleStateTimeline) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yRuleStateTimeline) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


