# TrustDesk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]TrustGrantView**](TrustGrantView.md) | Grants is every grant that has been made, newest first. | [optional] 
**Items** | Pointer to [**[]TrustItemView**](TrustItemView.md) | Items is everything the org holds, both tiers, retired included. | [optional] 
**Name** | Pointer to **string** | Name is the centre&#39;s display name. | [optional] 
**Nda** | Pointer to **string** | Nda is the text a party must accept before asking. | [optional] 
**Published** | Pointer to **bool** | Published is whether the centre answers at its public address. | [optional] 
**Requests** | Pointer to [**[]TrustAskView**](TrustAskView.md) | Requests is every ask, newest first, open ones included. | [optional] 
**Slug** | Pointer to **string** | Slug is the public address, empty until the centre is published. | [optional] 

## Methods

### NewTrustDesk

`func NewTrustDesk() *TrustDesk`

NewTrustDesk instantiates a new TrustDesk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustDeskWithDefaults

`func NewTrustDeskWithDefaults() *TrustDesk`

NewTrustDeskWithDefaults instantiates a new TrustDesk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *TrustDesk) GetGrants() []TrustGrantView`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *TrustDesk) GetGrantsOk() (*[]TrustGrantView, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *TrustDesk) SetGrants(v []TrustGrantView)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *TrustDesk) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetItems

`func (o *TrustDesk) GetItems() []TrustItemView`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrustDesk) GetItemsOk() (*[]TrustItemView, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrustDesk) SetItems(v []TrustItemView)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrustDesk) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *TrustDesk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustDesk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustDesk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrustDesk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNda

`func (o *TrustDesk) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *TrustDesk) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *TrustDesk) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *TrustDesk) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetPublished

`func (o *TrustDesk) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TrustDesk) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TrustDesk) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TrustDesk) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRequests

`func (o *TrustDesk) GetRequests() []TrustAskView`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *TrustDesk) GetRequestsOk() (*[]TrustAskView, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *TrustDesk) SetRequests(v []TrustAskView)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *TrustDesk) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSlug

`func (o *TrustDesk) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TrustDesk) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TrustDesk) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TrustDesk) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


