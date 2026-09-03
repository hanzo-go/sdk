# O11yO11yLLMTracesPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMTrace**](O11yO11yLLMTrace.md) | Items are the traces, newest first. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int64** | Offset is the row offset this page started at. | [optional] 

## Methods

### NewO11yO11yLLMTracesPage

`func NewO11yO11yLLMTracesPage() *O11yO11yLLMTracesPage`

NewO11yO11yLLMTracesPage instantiates a new O11yO11yLLMTracesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMTracesPageWithDefaults

`func NewO11yO11yLLMTracesPageWithDefaults() *O11yO11yLLMTracesPage`

NewO11yO11yLLMTracesPageWithDefaults instantiates a new O11yO11yLLMTracesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMTracesPage) GetItems() []O11yO11yLLMTrace`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMTracesPage) GetItemsOk() (*[]O11yO11yLLMTrace, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMTracesPage) SetItems(v []O11yO11yLLMTrace)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMTracesPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMTracesPage) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMTracesPage) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMTracesPage) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMTracesPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMTracesPage) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMTracesPage) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMTracesPage) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMTracesPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


