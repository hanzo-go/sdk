# O11yO11yErrorIssues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yErrorIssue**](O11yO11yErrorIssue.md) | Items are the issues. | [optional] 
**Limit** | Pointer to **int64** | Limit is the page cap that was applied. | [optional] 
**Offset** | Pointer to **int64** | Offset is how many were skipped. | [optional] 
**Total** | Pointer to **int64** | Total is how many matched the filter. | [optional] 

## Methods

### NewO11yO11yErrorIssues

`func NewO11yO11yErrorIssues() *O11yO11yErrorIssues`

NewO11yO11yErrorIssues instantiates a new O11yO11yErrorIssues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorIssuesWithDefaults

`func NewO11yO11yErrorIssuesWithDefaults() *O11yO11yErrorIssues`

NewO11yO11yErrorIssuesWithDefaults instantiates a new O11yO11yErrorIssues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yErrorIssues) GetItems() []O11yO11yErrorIssue`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yErrorIssues) GetItemsOk() (*[]O11yO11yErrorIssue, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yErrorIssues) SetItems(v []O11yO11yErrorIssue)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yErrorIssues) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yErrorIssues) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yErrorIssues) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yErrorIssues) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yErrorIssues) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yErrorIssues) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yErrorIssues) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yErrorIssues) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yErrorIssues) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yErrorIssues) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yErrorIssues) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yErrorIssues) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yErrorIssues) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


