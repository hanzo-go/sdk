# IndexIndexList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | Limit is how many rows this page could hold. | [optional] 
**Offset** | Pointer to **int64** | Offset is where this page starts. | [optional] 
**Results** | Pointer to [**[]IndexIndexView**](IndexIndexView.md) | Results are the index definitions themselves. | [optional] 
**Total** | Pointer to **int64** | Total is how many indexes the org holds altogether. | [optional] 

## Methods

### NewIndexIndexList

`func NewIndexIndexList() *IndexIndexList`

NewIndexIndexList instantiates a new IndexIndexList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexListWithDefaults

`func NewIndexIndexListWithDefaults() *IndexIndexList`

NewIndexIndexListWithDefaults instantiates a new IndexIndexList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IndexIndexList) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexIndexList) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexIndexList) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexIndexList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexIndexList) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexIndexList) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexIndexList) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexIndexList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *IndexIndexList) GetResults() []IndexIndexView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IndexIndexList) GetResultsOk() (*[]IndexIndexView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IndexIndexList) SetResults(v []IndexIndexView)`

SetResults sets Results field to given value.

### HasResults

`func (o *IndexIndexList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *IndexIndexList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IndexIndexList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IndexIndexList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IndexIndexList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


