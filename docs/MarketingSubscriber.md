# MarketingSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clicks** | Pointer to **int64** | Clicks counts the clicks reported. | [optional] 
**ConfirmedAt** | Pointer to **int64** | ConfirmedAt is when the subscription was confirmed or, for one that needed no confirmation, began; 0 while pending. | [optional] 
**Consent** | Pointer to [**MarketingConsent**](MarketingConsent.md) | Consent is the evidence of how and when the subscriber agreed. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the row was created, unix seconds. | [optional] 
**Email** | Pointer to **string** | Email is the address, lower-cased and trimmed; one row per address per list. | [optional] 
**Fields** | Pointer to **map[string]string** | Fields are the values of the org&#39;s custom fields, by key. | [optional] 
**Id** | Pointer to **int64** | ID is the subscriber&#39;s id within the org. Ids grow in the order rows were created, which is the order a list reads in by default. | [optional] 
**Language** | Pointer to **string** | Language is the subscriber&#39;s language, a BCP 47 tag. | [optional] 
**LastClickAt** | Pointer to **int64** | LastClickAt is the last click the email plane reported. | [optional] 
**LastOpenAt** | Pointer to **int64** | LastOpenAt is the last open the email plane reported. | [optional] 
**LastSentAt** | Pointer to **int64** | LastSentAt is the last broadcast sent to them. | [optional] 
**List** | Pointer to **string** | List is the list&#39;s id. | [optional] 
**Name** | Pointer to **string** | Name is the person&#39;s name, as given. | [optional] 
**Opens** | Pointer to **int64** | Opens counts the opens reported. | [optional] 
**Person** | Pointer to **string** | Person is the CRM person the address is: the key the event stream, the CRM and Insights join on. | [optional] 
**Ref** | Pointer to **string** | Ref is that place: a form id, a page, an import batch id. | [optional] 
**Source** | Pointer to **string** | Source is the kind of place the membership came from: form, page, api, import or commerce. | [optional] 
**Status** | Pointer to **string** | Status is pending, subscribed, unsubscribed, bounced, complained or cleaned. | [optional] 
**Tags** | Pointer to **[]string** | Tags label the subscriber; each at most 64 characters. | [optional] 
**UnsubscribedAt** | Pointer to **int64** | UnsubscribedAt is when the subscriber left; 0 while they have not. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last change, unix seconds. | [optional] 

## Methods

### NewMarketingSubscriber

`func NewMarketingSubscriber() *MarketingSubscriber`

NewMarketingSubscriber instantiates a new MarketingSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSubscriberWithDefaults

`func NewMarketingSubscriberWithDefaults() *MarketingSubscriber`

NewMarketingSubscriberWithDefaults instantiates a new MarketingSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClicks

`func (o *MarketingSubscriber) GetClicks() int64`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *MarketingSubscriber) GetClicksOk() (*int64, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *MarketingSubscriber) SetClicks(v int64)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *MarketingSubscriber) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *MarketingSubscriber) GetConfirmedAt() int64`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *MarketingSubscriber) GetConfirmedAtOk() (*int64, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *MarketingSubscriber) SetConfirmedAt(v int64)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *MarketingSubscriber) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetConsent

`func (o *MarketingSubscriber) GetConsent() MarketingConsent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MarketingSubscriber) GetConsentOk() (*MarketingConsent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MarketingSubscriber) SetConsent(v MarketingConsent)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *MarketingSubscriber) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingSubscriber) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingSubscriber) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingSubscriber) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingSubscriber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *MarketingSubscriber) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MarketingSubscriber) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MarketingSubscriber) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MarketingSubscriber) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFields

`func (o *MarketingSubscriber) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MarketingSubscriber) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MarketingSubscriber) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MarketingSubscriber) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *MarketingSubscriber) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingSubscriber) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingSubscriber) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingSubscriber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *MarketingSubscriber) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MarketingSubscriber) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MarketingSubscriber) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MarketingSubscriber) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastClickAt

`func (o *MarketingSubscriber) GetLastClickAt() int64`

GetLastClickAt returns the LastClickAt field if non-nil, zero value otherwise.

### GetLastClickAtOk

`func (o *MarketingSubscriber) GetLastClickAtOk() (*int64, bool)`

GetLastClickAtOk returns a tuple with the LastClickAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClickAt

`func (o *MarketingSubscriber) SetLastClickAt(v int64)`

SetLastClickAt sets LastClickAt field to given value.

### HasLastClickAt

`func (o *MarketingSubscriber) HasLastClickAt() bool`

HasLastClickAt returns a boolean if a field has been set.

### GetLastOpenAt

`func (o *MarketingSubscriber) GetLastOpenAt() int64`

GetLastOpenAt returns the LastOpenAt field if non-nil, zero value otherwise.

### GetLastOpenAtOk

`func (o *MarketingSubscriber) GetLastOpenAtOk() (*int64, bool)`

GetLastOpenAtOk returns a tuple with the LastOpenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenAt

`func (o *MarketingSubscriber) SetLastOpenAt(v int64)`

SetLastOpenAt sets LastOpenAt field to given value.

### HasLastOpenAt

`func (o *MarketingSubscriber) HasLastOpenAt() bool`

HasLastOpenAt returns a boolean if a field has been set.

### GetLastSentAt

`func (o *MarketingSubscriber) GetLastSentAt() int64`

GetLastSentAt returns the LastSentAt field if non-nil, zero value otherwise.

### GetLastSentAtOk

`func (o *MarketingSubscriber) GetLastSentAtOk() (*int64, bool)`

GetLastSentAtOk returns a tuple with the LastSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentAt

`func (o *MarketingSubscriber) SetLastSentAt(v int64)`

SetLastSentAt sets LastSentAt field to given value.

### HasLastSentAt

`func (o *MarketingSubscriber) HasLastSentAt() bool`

HasLastSentAt returns a boolean if a field has been set.

### GetList

`func (o *MarketingSubscriber) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingSubscriber) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingSubscriber) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingSubscriber) HasList() bool`

HasList returns a boolean if a field has been set.

### GetName

`func (o *MarketingSubscriber) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketingSubscriber) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketingSubscriber) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketingSubscriber) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpens

`func (o *MarketingSubscriber) GetOpens() int64`

GetOpens returns the Opens field if non-nil, zero value otherwise.

### GetOpensOk

`func (o *MarketingSubscriber) GetOpensOk() (*int64, bool)`

GetOpensOk returns a tuple with the Opens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpens

`func (o *MarketingSubscriber) SetOpens(v int64)`

SetOpens sets Opens field to given value.

### HasOpens

`func (o *MarketingSubscriber) HasOpens() bool`

HasOpens returns a boolean if a field has been set.

### GetPerson

`func (o *MarketingSubscriber) GetPerson() string`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *MarketingSubscriber) GetPersonOk() (*string, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *MarketingSubscriber) SetPerson(v string)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *MarketingSubscriber) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetRef

`func (o *MarketingSubscriber) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MarketingSubscriber) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MarketingSubscriber) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MarketingSubscriber) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSource

`func (o *MarketingSubscriber) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingSubscriber) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingSubscriber) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingSubscriber) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingSubscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingSubscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingSubscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingSubscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *MarketingSubscriber) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MarketingSubscriber) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MarketingSubscriber) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MarketingSubscriber) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUnsubscribedAt

`func (o *MarketingSubscriber) GetUnsubscribedAt() int64`

GetUnsubscribedAt returns the UnsubscribedAt field if non-nil, zero value otherwise.

### GetUnsubscribedAtOk

`func (o *MarketingSubscriber) GetUnsubscribedAtOk() (*int64, bool)`

GetUnsubscribedAtOk returns a tuple with the UnsubscribedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribedAt

`func (o *MarketingSubscriber) SetUnsubscribedAt(v int64)`

SetUnsubscribedAt sets UnsubscribedAt field to given value.

### HasUnsubscribedAt

`func (o *MarketingSubscriber) HasUnsubscribedAt() bool`

HasUnsubscribedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingSubscriber) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingSubscriber) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingSubscriber) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingSubscriber) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


