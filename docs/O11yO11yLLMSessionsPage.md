# O11yO11yLLMSessionsPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMSession**](O11yO11yLLMSession.md) | Items are the conversations, newest first. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int64** | Offset is the row offset this page started at. | [optional] 

## Methods

### NewO11yO11yLLMSessionsPage

`func NewO11yO11yLLMSessionsPage() *O11yO11yLLMSessionsPage`

NewO11yO11yLLMSessionsPage instantiates a new O11yO11yLLMSessionsPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMSessionsPageWithDefaults

`func NewO11yO11yLLMSessionsPageWithDefaults() *O11yO11yLLMSessionsPage`

NewO11yO11yLLMSessionsPageWithDefaults instantiates a new O11yO11yLLMSessionsPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMSessionsPage) GetItems() []O11yO11yLLMSession`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMSessionsPage) GetItemsOk() (*[]O11yO11yLLMSession, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMSessionsPage) SetItems(v []O11yO11yLLMSession)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMSessionsPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMSessionsPage) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMSessionsPage) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMSessionsPage) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMSessionsPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMSessionsPage) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMSessionsPage) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMSessionsPage) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMSessionsPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


