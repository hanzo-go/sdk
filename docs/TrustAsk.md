# TrustAsk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accept** | Pointer to **bool** | Accept must be true when the centre states an NDA. The text accepted is recorded verbatim on the request, so a later edit to the NDA cannot rewrite what this party agreed to. | [optional] 
**Email** | Pointer to **string** | Email is where the grant will be sent, and the ONLY address the resulting link admits. Required. | [optional] 
**Item** | Pointer to **string** | Item names one published item to ask for. Optional; omitting it asks for everything released on request. | [optional] 
**Party** | Pointer to **string** | Party is the company the asker is from. Optional, and recorded as stated. | [optional] 
**Reason** | Pointer to **string** | Reason is why they want it. Optional, and recorded as stated — it is what the person deciding reads. | [optional] 
**Slug** | Pointer to **string** | Slug is the centre&#39;s public address, taken from the path. | [optional] 

## Methods

### NewTrustAsk

`func NewTrustAsk() *TrustAsk`

NewTrustAsk instantiates a new TrustAsk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustAskWithDefaults

`func NewTrustAskWithDefaults() *TrustAsk`

NewTrustAskWithDefaults instantiates a new TrustAsk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccept

`func (o *TrustAsk) GetAccept() bool`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *TrustAsk) GetAcceptOk() (*bool, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *TrustAsk) SetAccept(v bool)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *TrustAsk) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetEmail

`func (o *TrustAsk) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TrustAsk) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TrustAsk) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TrustAsk) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetItem

`func (o *TrustAsk) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TrustAsk) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TrustAsk) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *TrustAsk) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetParty

`func (o *TrustAsk) GetParty() string`

GetParty returns the Party field if non-nil, zero value otherwise.

### GetPartyOk

`func (o *TrustAsk) GetPartyOk() (*string, bool)`

GetPartyOk returns a tuple with the Party field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParty

`func (o *TrustAsk) SetParty(v string)`

SetParty sets Party field to given value.

### HasParty

`func (o *TrustAsk) HasParty() bool`

HasParty returns a boolean if a field has been set.

### GetReason

`func (o *TrustAsk) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TrustAsk) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TrustAsk) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TrustAsk) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSlug

`func (o *TrustAsk) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TrustAsk) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TrustAsk) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TrustAsk) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


