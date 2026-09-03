# O11yO11yLLMAnnotationsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMAnnotation**](O11yO11yLLMAnnotation.md) | Items are the annotations, newest first. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int64** | Offset is the row offset this page started at. | [optional] 
**Total** | Pointer to **int64** | Total is how many annotations match, across all pages. | [optional] 

## Methods

### NewO11yO11yLLMAnnotationsPage

`func NewO11yO11yLLMAnnotationsPage() *O11yO11yLLMAnnotationsPage`

NewO11yO11yLLMAnnotationsPage instantiates a new O11yO11yLLMAnnotationsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMAnnotationsPageWithDefaults

`func NewO11yO11yLLMAnnotationsPageWithDefaults() *O11yO11yLLMAnnotationsPage`

NewO11yO11yLLMAnnotationsPageWithDefaults instantiates a new O11yO11yLLMAnnotationsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMAnnotationsPage) GetItems() []O11yO11yLLMAnnotation`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMAnnotationsPage) GetItemsOk() (*[]O11yO11yLLMAnnotation, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMAnnotationsPage) SetItems(v []O11yO11yLLMAnnotation)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMAnnotationsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMAnnotationsPage) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMAnnotationsPage) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMAnnotationsPage) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMAnnotationsPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMAnnotationsPage) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMAnnotationsPage) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMAnnotationsPage) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMAnnotationsPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yLLMAnnotationsPage) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yLLMAnnotationsPage) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yLLMAnnotationsPage) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yLLMAnnotationsPage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


