# CloudListMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Limit is how many rows one page holds. | [optional] 
**Page** | Pointer to **int32** | Page is the 1-based page this response is. | [optional] 
**TotalItems** | Pointer to **int32** | TotalItems is how many rows match in total. | [optional] 
**TotalPages** | Pointer to **int32** | TotalPages is ceil(totalItems/limit), at least 1. | [optional] 

## Methods

### NewCloudListMeta

`func NewCloudListMeta() *CloudListMeta`

NewCloudListMeta instantiates a new CloudListMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudListMetaWithDefaults

`func NewCloudListMetaWithDefaults() *CloudListMeta`

NewCloudListMetaWithDefaults instantiates a new CloudListMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *CloudListMeta) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CloudListMeta) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CloudListMeta) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CloudListMeta) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *CloudListMeta) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CloudListMeta) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CloudListMeta) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *CloudListMeta) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalItems

`func (o *CloudListMeta) GetTotalItems() int32`

GetTotalItems returns the TotalItems field if non-nil, zero value otherwise.

### GetTotalItemsOk

`func (o *CloudListMeta) GetTotalItemsOk() (*int32, bool)`

GetTotalItemsOk returns a tuple with the TotalItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItems

`func (o *CloudListMeta) SetTotalItems(v int32)`

SetTotalItems sets TotalItems field to given value.

### HasTotalItems

`func (o *CloudListMeta) HasTotalItems() bool`

HasTotalItems returns a boolean if a field has been set.

### GetTotalPages

`func (o *CloudListMeta) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *CloudListMeta) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *CloudListMeta) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *CloudListMeta) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


