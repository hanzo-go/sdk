# O11yO11yLLMUsersPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLLMUser**](O11yO11yLLMUser.md) | Items are the end users, newest first. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page cap the read ran with. | [optional] 
**Offset** | Pointer to **int32** | Offset is the row offset this page started at. | [optional] 

## Methods

### NewO11yO11yLLMUsersPage

`func NewO11yO11yLLMUsersPage() *O11yO11yLLMUsersPage`

NewO11yO11yLLMUsersPage instantiates a new O11yO11yLLMUsersPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMUsersPageWithDefaults

`func NewO11yO11yLLMUsersPageWithDefaults() *O11yO11yLLMUsersPage`

NewO11yO11yLLMUsersPageWithDefaults instantiates a new O11yO11yLLMUsersPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLLMUsersPage) GetItems() []O11yO11yLLMUser`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLLMUsersPage) GetItemsOk() (*[]O11yO11yLLMUser, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLLMUsersPage) SetItems(v []O11yO11yLLMUser)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLLMUsersPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yLLMUsersPage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yLLMUsersPage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yLLMUsersPage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yLLMUsersPage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yLLMUsersPage) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yLLMUsersPage) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yLLMUsersPage) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yLLMUsersPage) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


