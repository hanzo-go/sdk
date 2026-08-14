# O11yListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit is how many rows one page holds. | [optional] 
**Page** | Pointer to **int32** | Page is the 1-based page this response is. | [optional] 
**TotalItems** | Pointer to **int32** | TotalItems is how many rows match in total. | [optional] 
**TotalPages** | Pointer to **int32** | TotalPages is ceil(totalItems/limit), at least 1. | [optional] 

## Methods

### NewO11yListMeta

`func NewO11yListMeta() *O11yListMeta`

NewO11yListMeta instantiates a new O11yListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yListMetaWithDefaults

`func NewO11yListMetaWithDefaults() *O11yListMeta`

NewO11yListMetaWithDefaults instantiates a new O11yListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *O11yListMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yListMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yListMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yListMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *O11yListMeta) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *O11yListMeta) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *O11yListMeta) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *O11yListMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalItems

`func (o *O11yListMeta) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *O11yListMeta) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *O11yListMeta) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *O11yListMeta) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetTotalPages

`func (o *O11yListMeta) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *O11yListMeta) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *O11yListMeta) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *O11yListMeta) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


