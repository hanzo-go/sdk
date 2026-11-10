# MarketingSubscriberIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is the consent basis to record: explicit, legitimate_interest or unverified (the default). | [optional] 
**ConsentAt** | Pointer to **int64** | ConsentAt is when consent was given, unix seconds, when it was given before this call; on add only. | [optional] 
**Email** | Pointer to **string** | Email is the address to add. Required on add; it cannot change. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields set custom field values by key; an empty value clears one. | [optional] 
**Id** | Pointer to **string** | ID is the list, by id or slug, from the path. | [optional] 
**Language** | Pointer to **string** | Language is a BCP 47 tag. | [optional] 
**Name** | Pointer to **string** | Name is the person&#39;s name; at most 200 characters. | [optional] 
**Ref** | Pointer to **string** | Ref is where the subscriber came from (a form id, a page); on add only. | [optional] 
**Sid** | Pointer to **int64** | SID is the subscriber to change, from the path; ignored on add. | [optional] 
**Status** | Pointer to **string** | Status is subscribed (the default on add), pending, which mails the address a link to confirm, or unsubscribed, bounced or cleaned. An address that unsubscribed or complained rejoins only by confirming again itself. | [optional] 
**Tags** | Pointer to **[]string** | Tags replace the subscriber&#39;s tags. | [optional] 

## Methods

### NewMarketingSubscriberIn

`func NewMarketingSubscriberIn() *MarketingSubscriberIn`

NewMarketingSubscriberIn instantiates a new MarketingSubscriberIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriberInWithDefaults

`func NewMarketingSubscriberInWithDefaults() *MarketingSubscriberIn`

NewMarketingSubscriberInWithDefaults instantiates a new MarketingSubscriberIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *MarketingSubscriberIn) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *MarketingSubscriberIn) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *MarketingSubscriberIn) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *MarketingSubscriberIn) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetConsentAt

`func (o *MarketingSubscriberIn) GetConsentAt() int64`

GetConsentAt returns the ConsentAt field if non-nil, zero value otherwise.

### GetConsentAtOk

`func (o *MarketingSubscriberIn) GetConsentAtOk() (*int64, bool)`

GetConsentAtOk returns a tuple with the ConsentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentAt

`func (o *MarketingSubscriberIn) SetConsentAt(v int64)`

SetConsentAt sets ConsentAt field to given value.

### HasConsentAt

`func (o *MarketingSubscriberIn) HasConsentAt() bool`

HasConsentAt returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingSubscriberIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingSubscriberIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingSubscriberIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingSubscriberIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFields

`func (o *MarketingSubscriberIn) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MarketingSubscriberIn) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MarketingSubscriberIn) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MarketingSubscriberIn) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *MarketingSubscriberIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingSubscriberIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingSubscriberIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingSubscriberIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingSubscriberIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingSubscriberIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingSubscriberIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingSubscriberIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetName

`func (o *MarketingSubscriberIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingSubscriberIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingSubscriberIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingSubscriberIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRef

`func (o *MarketingSubscriberIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MarketingSubscriberIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MarketingSubscriberIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MarketingSubscriberIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSid

`func (o *MarketingSubscriberIn) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *MarketingSubscriberIn) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *MarketingSubscriberIn) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *MarketingSubscriberIn) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingSubscriberIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingSubscriberIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingSubscriberIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingSubscriberIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MarketingSubscriberIn) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingSubscriberIn) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingSubscriberIn) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingSubscriberIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


