# IndexList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit is how many rows this page could hold. | [optional] 
**Offset** | Pointer to **int32** | Offset is where this page starts. | [optional] 
**Results** | Pointer to [**[]IndexView**](IndexView.md) | Results are the index definitions themselves. | [optional] 
**Total** | Pointer to **int32** | Total is how many indexes the org holds altogether. | [optional] 

## Methods

### NewIndexList

`func NewIndexList() *IndexList`

NewIndexList instantiates a new IndexList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexListWithDefaults

`func NewIndexListWithDefaults() *IndexList`

NewIndexListWithDefaults instantiates a new IndexList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IndexList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *IndexList) GetResults() []IndexView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IndexList) GetResultsOk() (*[]IndexView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IndexList) SetResults(v []IndexView)`

SetResults sets Results field to given value.

### HasResults

`func (o *IndexList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *IndexList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IndexList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IndexList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IndexList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


