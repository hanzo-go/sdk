# FormFormList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]FormForm**](FormForm.md) | Items are the forms, newest first. | [optional] 
**Page** | Pointer to **int64** | Page is this page, from 1. | [optional] 
**PageSize** | Pointer to **int64** | PageSize is how many a page holds. | [optional] 
**Total** | Pointer to **int64** | Total is how many forms answer the query. | [optional] 

## Methods

### NewFormFormList

`func NewFormFormList() *FormFormList`

NewFormFormList instantiates a new FormFormList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFormListWithDefaults

`func NewFormFormListWithDefaults() *FormFormList`

NewFormFormListWithDefaults instantiates a new FormFormList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FormFormList) GetItems() []FormForm`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FormFormList) GetItemsOk() (*[]FormForm, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FormFormList) SetItems(v []FormForm)`

SetItems sets Items field to given value.

### HasItems

`func (o *FormFormList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPage

`func (o *FormFormList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FormFormList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FormFormList) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *FormFormList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *FormFormList) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FormFormList) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FormFormList) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FormFormList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotal

`func (o *FormFormList) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FormFormList) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FormFormList) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FormFormList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


