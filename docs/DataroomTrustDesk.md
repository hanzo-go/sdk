# DataroomTrustDesk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]DataroomTrustGrantView**](DataroomTrustGrantView.md) | Grants is every grant that has been made, newest first. | [optional] 
**Items** | Pointer to [**[]DataroomTrustItemView**](DataroomTrustItemView.md) | Items is everything the org holds, both tiers, retired included. | [optional] 
**Name** | Pointer to **string** | Name is the centre&#39;s display name. | [optional] 
**Nda** | Pointer to **string** | Nda is the text a party must accept before asking. | [optional] 
**Published** | Pointer to **bool** | Published is whether the centre answers at its public address. | [optional] 
**Requests** | Pointer to [**[]DataroomTrustAskView**](DataroomTrustAskView.md) | Requests is every ask, newest first, open ones included. | [optional] 
**Slug** | Pointer to **string** | Slug is the public address, empty until the centre is published. | [optional] 

## Methods

### NewDataroomTrustDesk

`func NewDataroomTrustDesk() *DataroomTrustDesk`

NewDataroomTrustDesk instantiates a new DataroomTrustDesk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomTrustDeskWithDefaults

`func NewDataroomTrustDeskWithDefaults() *DataroomTrustDesk`

NewDataroomTrustDeskWithDefaults instantiates a new DataroomTrustDesk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *DataroomTrustDesk) GetGrants() []DataroomTrustGrantView`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *DataroomTrustDesk) GetGrantsOk() (*[]DataroomTrustGrantView, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *DataroomTrustDesk) SetGrants(v []DataroomTrustGrantView)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *DataroomTrustDesk) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetItems

`func (o *DataroomTrustDesk) GetItems() []DataroomTrustItemView`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DataroomTrustDesk) GetItemsOk() (*[]DataroomTrustItemView, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DataroomTrustDesk) SetItems(v []DataroomTrustItemView)`

SetItems sets Items field to given value.

### HasItems

`func (o *DataroomTrustDesk) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *DataroomTrustDesk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataroomTrustDesk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataroomTrustDesk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataroomTrustDesk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNda

`func (o *DataroomTrustDesk) GetNda() string`

GetNda returns the Nda field if non-nil, zero value otherwise.

### GetNdaOk

`func (o *DataroomTrustDesk) GetNdaOk() (*string, bool)`

GetNdaOk returns a tuple with the Nda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNda

`func (o *DataroomTrustDesk) SetNda(v string)`

SetNda sets Nda field to given value.

### HasNda

`func (o *DataroomTrustDesk) HasNda() bool`

HasNda returns a boolean if a field has been set.

### GetPublished

`func (o *DataroomTrustDesk) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *DataroomTrustDesk) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *DataroomTrustDesk) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *DataroomTrustDesk) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRequests

`func (o *DataroomTrustDesk) GetRequests() []DataroomTrustAskView`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *DataroomTrustDesk) GetRequestsOk() (*[]DataroomTrustAskView, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *DataroomTrustDesk) SetRequests(v []DataroomTrustAskView)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *DataroomTrustDesk) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSlug

`func (o *DataroomTrustDesk) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DataroomTrustDesk) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DataroomTrustDesk) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *DataroomTrustDesk) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


