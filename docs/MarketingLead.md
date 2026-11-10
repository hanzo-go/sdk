# MarketingLead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lead** | Pointer to **string** | Lead is the reference of what was filed: \&quot;lead_\&quot; and 24 hex characters for a lead, \&quot;cand_\&quot; and 24 for a job application. It is random, not a record&#39;s id, so it says nothing about how many the team has. | [optional] 
**Offer** | Pointer to **string** | Offer says what an application&#39;s code did: \&quot;mailed\&quot;, a new code went to the address; \&quot;held\&quot;, the address already holds one. Absent when no code was given. The code itself is never answered here. | [optional] 
**Sent** | Pointer to **bool** | Sent says whether an investor&#39;s SAFE went out for signature from Hanzo eSign: false when the deployment sends none yet, or eSign did not take it, and the investor is filed either way. Absent on every other form. | [optional] 
**Stage** | Pointer to **string** | Stage is where the lead stands. Always \&quot;new\&quot; on a fresh lead. | [optional] 

## Methods

### NewMarketingLead

`func NewMarketingLead() *MarketingLead`

NewMarketingLead instantiates a new MarketingLead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingLeadWithDefaults

`func NewMarketingLeadWithDefaults() *MarketingLead`

NewMarketingLeadWithDefaults instantiates a new MarketingLead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLead

`func (o *MarketingLead) GetLead() string`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *MarketingLead) GetLeadOk() (*string, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *MarketingLead) SetLead(v string)`

SetLead sets Lead field to given value.

### HasLead

`func (o *MarketingLead) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetOffer

`func (o *MarketingLead) GetOffer() string`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *MarketingLead) GetOfferOk() (*string, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *MarketingLead) SetOffer(v string)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *MarketingLead) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetSent

`func (o *MarketingLead) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MarketingLead) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MarketingLead) SetSent(v bool)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MarketingLead) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetStage

`func (o *MarketingLead) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *MarketingLead) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *MarketingLead) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *MarketingLead) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


