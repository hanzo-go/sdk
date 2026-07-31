# BaseRecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**TotalItems** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**Items** | Pointer to [**[]BaseRecord**](BaseRecord.md) |  | [optional] 

## Methods

### NewBaseRecordList

`func NewBaseRecordList() *BaseRecordList`

NewBaseRecordList instantiates a new BaseRecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseRecordListWithDefaults

`func NewBaseRecordListWithDefaults() *BaseRecordList`

NewBaseRecordListWithDefaults instantiates a new BaseRecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *BaseRecordList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BaseRecordList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BaseRecordList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *BaseRecordList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *BaseRecordList) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *BaseRecordList) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *BaseRecordList) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *BaseRecordList) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotalItems

`func (o *BaseRecordList) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *BaseRecordList) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *BaseRecordList) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *BaseRecordList) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetTotalPages

`func (o *BaseRecordList) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *BaseRecordList) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *BaseRecordList) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *BaseRecordList) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetItems

`func (o *BaseRecordList) GetItems() []BaseRecord`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BaseRecordList) GetItemsOk() (*[]BaseRecord, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BaseRecordList) SetItems(v []BaseRecord)`

SetItems sets Items field to given value.

### HasItems

`func (o *BaseRecordList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


