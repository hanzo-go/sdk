# O11yAnnItemList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]O11yAnnItemView**](O11yAnnItemView.md) | Data is the page of items. | [optional] 
**Meta** | Pointer to [**O11yListMeta**](O11yListMeta.md) | Meta is the paging that produced it. | [optional] 

## Methods

### NewO11yAnnItemList

`func NewO11yAnnItemList() *O11yAnnItemList`

NewO11yAnnItemList instantiates a new O11yAnnItemList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAnnItemListWithDefaults

`func NewO11yAnnItemListWithDefaults() *O11yAnnItemList`

NewO11yAnnItemListWithDefaults instantiates a new O11yAnnItemList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *O11yAnnItemList) GetData() []O11yAnnItemView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *O11yAnnItemList) GetDataOk() (*[]O11yAnnItemView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *O11yAnnItemList) SetData(v []O11yAnnItemView)`

SetData sets Data field to given value.

### HasData

`func (o *O11yAnnItemList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *O11yAnnItemList) GetMeta() O11yListMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yAnnItemList) GetMetaOk() (*O11yListMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yAnnItemList) SetMeta(v O11yListMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yAnnItemList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


