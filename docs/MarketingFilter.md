# MarketingFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis keeps subscribers by the consent basis recorded for them: explicit, legitimate_interest or unverified. | [optional] 
**ClickedAfter** | Pointer to **int64** | ClickedAfter keeps subscribers who clicked a link at or after this time. | [optional] 
**ConfirmedAfter** | Pointer to **int64** | ConfirmedAfter and ConfirmedBefore bound when the subscription was confirmed, unix seconds, inclusive. | [optional] 
**ConfirmedBefore** | Pointer to **int64** | ConfirmedBefore is the upper bound of ConfirmedAfter&#39;s range. | [optional] 
**CreatedAfter** | Pointer to **int64** | CreatedAfter and CreatedBefore bound when the subscriber joined the list, unix seconds, inclusive. Zero is unbounded. | [optional] 
**CreatedBefore** | Pointer to **int64** | CreatedBefore is the upper bound of CreatedAfter&#39;s range. | [optional] 
**Field** | Pointer to **string** | Field names a custom field, and Value keeps subscribers whose value for it is exactly this. | [optional] 
**NotOpenedSince** | Pointer to **int64** | NotOpenedSince keeps subscribers who have opened nothing since this time, including those who never opened anything. | [optional] 
**OpenedAfter** | Pointer to **int64** | OpenedAfter keeps subscribers who opened a broadcast at or after this time. | [optional] 
**Q** | Pointer to **string** | Q keeps subscribers whose address or name begins with it, case-insensitively. A search is ordered by address. | [optional] 
**Ref** | Pointer to **string** | Ref keeps subscribers from one source instance: a form id, a page, an import batch id. | [optional] 
**Source** | Pointer to **string** | Source keeps subscribers that came from this kind of source: form, page, api, import or commerce. | [optional] 
**Status** | Pointer to **string** | Status is one status or several, comma-separated: pending, subscribed, unsubscribed, bounced, complained, cleaned. | [optional] 
**Tag** | Pointer to **string** | Tag keeps subscribers carrying this tag. | [optional] 
**Value** | Pointer to **string** | Value is the custom field&#39;s value Field keeps. | [optional] 

## Methods

### NewMarketingFilter

`func NewMarketingFilter() *MarketingFilter`

NewMarketingFilter instantiates a new MarketingFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingFilterWithDefaults

`func NewMarketingFilterWithDefaults() *MarketingFilter`

NewMarketingFilterWithDefaults instantiates a new MarketingFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *MarketingFilter) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *MarketingFilter) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *MarketingFilter) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *MarketingFilter) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetClickedAfter

`func (o *MarketingFilter) GetClickedAfter() int64`

GetClickedAfter returns the ClickedAfter field if non-nil, zero value otherwise.

### GetClickedAfterOk

`func (o *MarketingFilter) GetClickedAfterOk() (*int64, bool)`

GetClickedAfterOk returns a tuple with the ClickedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickedAfter

`func (o *MarketingFilter) SetClickedAfter(v int64)`

SetClickedAfter sets ClickedAfter field to given value.

### HasClickedAfter

`func (o *MarketingFilter) HasClickedAfter() bool`

HasClickedAfter returns a boolean if a field has been set.

### GetConfirmedAfter

`func (o *MarketingFilter) GetConfirmedAfter() int64`

GetConfirmedAfter returns the ConfirmedAfter field if non-nil, zero value otherwise.

### GetConfirmedAfterOk

`func (o *MarketingFilter) GetConfirmedAfterOk() (*int64, bool)`

GetConfirmedAfterOk returns a tuple with the ConfirmedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAfter

`func (o *MarketingFilter) SetConfirmedAfter(v int64)`

SetConfirmedAfter sets ConfirmedAfter field to given value.

### HasConfirmedAfter

`func (o *MarketingFilter) HasConfirmedAfter() bool`

HasConfirmedAfter returns a boolean if a field has been set.

### GetConfirmedBefore

`func (o *MarketingFilter) GetConfirmedBefore() int64`

GetConfirmedBefore returns the ConfirmedBefore field if non-nil, zero value otherwise.

### GetConfirmedBeforeOk

`func (o *MarketingFilter) GetConfirmedBeforeOk() (*int64, bool)`

GetConfirmedBeforeOk returns a tuple with the ConfirmedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedBefore

`func (o *MarketingFilter) SetConfirmedBefore(v int64)`

SetConfirmedBefore sets ConfirmedBefore field to given value.

### HasConfirmedBefore

`func (o *MarketingFilter) HasConfirmedBefore() bool`

HasConfirmedBefore returns a boolean if a field has been set.

### GetCreatedAfter

`func (o *MarketingFilter) GetCreatedAfter() int64`

GetCreatedAfter returns the CreatedAfter field if non-nil, zero value otherwise.

### GetCreatedAfterOk

`func (o *MarketingFilter) GetCreatedAfterOk() (*int64, bool)`

GetCreatedAfterOk returns a tuple with the CreatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAfter

`func (o *MarketingFilter) SetCreatedAfter(v int64)`

SetCreatedAfter sets CreatedAfter field to given value.

### HasCreatedAfter

`func (o *MarketingFilter) HasCreatedAfter() bool`

HasCreatedAfter returns a boolean if a field has been set.

### GetCreatedBefore

`func (o *MarketingFilter) GetCreatedBefore() int64`

GetCreatedBefore returns the CreatedBefore field if non-nil, zero value otherwise.

### GetCreatedBeforeOk

`func (o *MarketingFilter) GetCreatedBeforeOk() (*int64, bool)`

GetCreatedBeforeOk returns a tuple with the CreatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBefore

`func (o *MarketingFilter) SetCreatedBefore(v int64)`

SetCreatedBefore sets CreatedBefore field to given value.

### HasCreatedBefore

`func (o *MarketingFilter) HasCreatedBefore() bool`

HasCreatedBefore returns a boolean if a field has been set.

### GetField

`func (o *MarketingFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MarketingFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MarketingFilter) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *MarketingFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetNotOpenedSince

`func (o *MarketingFilter) GetNotOpenedSince() int64`

GetNotOpenedSince returns the NotOpenedSince field if non-nil, zero value otherwise.

### GetNotOpenedSinceOk

`func (o *MarketingFilter) GetNotOpenedSinceOk() (*int64, bool)`

GetNotOpenedSinceOk returns a tuple with the NotOpenedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotOpenedSince

`func (o *MarketingFilter) SetNotOpenedSince(v int64)`

SetNotOpenedSince sets NotOpenedSince field to given value.

### HasNotOpenedSince

`func (o *MarketingFilter) HasNotOpenedSince() bool`

HasNotOpenedSince returns a boolean if a field has been set.

### GetOpenedAfter

`func (o *MarketingFilter) GetOpenedAfter() int64`

GetOpenedAfter returns the OpenedAfter field if non-nil, zero value otherwise.

### GetOpenedAfterOk

`func (o *MarketingFilter) GetOpenedAfterOk() (*int64, bool)`

GetOpenedAfterOk returns a tuple with the OpenedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAfter

`func (o *MarketingFilter) SetOpenedAfter(v int64)`

SetOpenedAfter sets OpenedAfter field to given value.

### HasOpenedAfter

`func (o *MarketingFilter) HasOpenedAfter() bool`

HasOpenedAfter returns a boolean if a field has been set.

### GetQ

`func (o *MarketingFilter) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *MarketingFilter) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *MarketingFilter) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *MarketingFilter) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetRef

`func (o *MarketingFilter) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MarketingFilter) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MarketingFilter) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MarketingFilter) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSource

`func (o *MarketingFilter) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MarketingFilter) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MarketingFilter) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *MarketingFilter) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingFilter) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingFilter) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingFilter) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTag

`func (o *MarketingFilter) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *MarketingFilter) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *MarketingFilter) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *MarketingFilter) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValue

`func (o *MarketingFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MarketingFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MarketingFilter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MarketingFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


