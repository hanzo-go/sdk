# TrustPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]TrustItem**](TrustItem.md) | Items is everything the centre publishes: what can be read now, and what exists and is released on request. | [optional] 
**Name** | Pointer to **string** | Name is the org&#39;s display name for its centre. | [optional] 
**Nda** | Pointer to **string** | Nda is the text a party must accept to ask for the gated items, verbatim. Empty when the org asks for none. | [optional] 
**Slug** | Pointer to **string** | Slug is the centre&#39;s public address. | [optional] 

## Methods

### NewTrustPage

`func NewTrustPage() *TrustPage`

NewTrustPage instantiates a new TrustPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustPageWithDefaults

`func NewTrustPageWithDefaults() *TrustPage`

NewTrustPageWithDefaults instantiates a new TrustPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *TrustPage) GetItems() []TrustItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrustPage) GetItemsOk() (*[]TrustItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrustPage) SetItems(v []TrustItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrustPage) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *TrustPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNda

`func (o *TrustPage) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *TrustPage) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *TrustPage) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *TrustPage) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetSlug

`func (o *TrustPage) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TrustPage) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TrustPage) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TrustPage) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


