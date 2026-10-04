# DataroomTrustPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DataroomTrustItem**](DataroomTrustItem.md) | Items is everything the centre publishes: what can be read now, and what exists and is released on request. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s display name for its centre. | [optional] 
**Nda** | Pointer to **string** | Nda is the text a party must accept to ask for the gated items, verbatim. Empty when the org asks for none. | [optional] 
**Slug** | Pointer to **string** | Slug is the centre&#39;s public address. | [optional] 

## Methods

### NewDataroomTrustPage

`func NewDataroomTrustPage() *DataroomTrustPage`

NewDataroomTrustPage instantiates a new DataroomTrustPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustPageWithDefaults

`func NewDataroomTrustPageWithDefaults() *DataroomTrustPage`

NewDataroomTrustPageWithDefaults instantiates a new DataroomTrustPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DataroomTrustPage) GetItems() []DataroomTrustItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DataroomTrustPage) GetItemsOk() (*[]DataroomTrustItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DataroomTrustPage) SetItems(v []DataroomTrustItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DataroomTrustPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *DataroomTrustPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomTrustPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomTrustPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomTrustPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNda

`func (o *DataroomTrustPage) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *DataroomTrustPage) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *DataroomTrustPage) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *DataroomTrustPage) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetSlug

`func (o *DataroomTrustPage) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DataroomTrustPage) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DataroomTrustPage) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DataroomTrustPage) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


