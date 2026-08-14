# O11yO11yLLMScoresPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMScore**](O11yO11yLLMScore.md) | Items are the scores, newest first. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int32** | Offset is the row offset this page started at. | [optional] 
**Total** | Pointer to **int32** | Total is how many scores match, across all pages. | [optional] 

## Methods

### NewO11yO11yLLMScoresPage

`func NewO11yO11yLLMScoresPage() *O11yO11yLLMScoresPage`

NewO11yO11yLLMScoresPage instantiates a new O11yO11yLLMScoresPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMScoresPageWithDefaults

`func NewO11yO11yLLMScoresPageWithDefaults() *O11yO11yLLMScoresPage`

NewO11yO11yLLMScoresPageWithDefaults instantiates a new O11yO11yLLMScoresPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMScoresPage) GetItems() []O11yO11yLLMScore`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMScoresPage) GetItemsOk() (*[]O11yO11yLLMScore, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMScoresPage) SetItems(v []O11yO11yLLMScore)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMScoresPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMScoresPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMScoresPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMScoresPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMScoresPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMScoresPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMScoresPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMScoresPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMScoresPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yLLMScoresPage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yLLMScoresPage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yLLMScoresPage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yLLMScoresPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


