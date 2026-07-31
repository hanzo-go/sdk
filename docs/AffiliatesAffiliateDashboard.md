# AffiliatesAffiliateDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAffiliate** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AffiliatesAffiliateStatus**](AffiliatesAffiliateStatus.md) |  | [optional] 
**Code** | Pointer to **string** | The minted affiliate code (empty until approved). | [optional] 
**RequestedCode** | Pointer to **string** | The vanity code requested at apply, pending approval. | [optional] 
**Link** | Pointer to **string** | The &#x60;?aff&#x60; referral link (empty until the affiliate has a code). | [optional] 
**RateBps** | Pointer to **int64** | Commission rate in basis points. | [optional] 
**ReferredCount** | Pointer to **int32** | Number of referred orgs attributed to this affiliate. | [optional] 
**AccruedCents** | Pointer to **int64** | Lifetime commission accrued (USD cents). | [optional] 
**PendingCents** | Pointer to **int64** | Commission accrued but not yet paid (accrued − paid, never negative). | [optional] 
**PaidCents** | Pointer to **int64** | Lifetime commission paid out (USD cents). | [optional] 
**Payouts** | Pointer to [**[]AffiliatesPayout**](AffiliatesPayout.md) |  | [optional] 

## Methods

### NewAffiliatesAffiliateDashboard

`func NewAffiliatesAffiliateDashboard() *AffiliatesAffiliateDashboard`

NewAffiliatesAffiliateDashboard instantiates a new AffiliatesAffiliateDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesAffiliateDashboardWithDefaults

`func NewAffiliatesAffiliateDashboardWithDefaults() *AffiliatesAffiliateDashboard`

NewAffiliatesAffiliateDashboardWithDefaults instantiates a new AffiliatesAffiliateDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAffiliate

`func (o *AffiliatesAffiliateDashboard) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliatesAffiliateDashboard) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliatesAffiliateDashboard) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliatesAffiliateDashboard) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetId

`func (o *AffiliatesAffiliateDashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesAffiliateDashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesAffiliateDashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesAffiliateDashboard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliatesAffiliateDashboard) GetStatus() AffiliatesAffiliateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliatesAffiliateDashboard) GetStatusOk() (*AffiliatesAffiliateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliatesAffiliateDashboard) SetStatus(v AffiliatesAffiliateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliatesAffiliateDashboard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *AffiliatesAffiliateDashboard) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliatesAffiliateDashboard) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliatesAffiliateDashboard) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliatesAffiliateDashboard) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliatesAffiliateDashboard) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliatesAffiliateDashboard) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliatesAffiliateDashboard) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliatesAffiliateDashboard) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetLink

`func (o *AffiliatesAffiliateDashboard) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AffiliatesAffiliateDashboard) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AffiliatesAffiliateDashboard) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AffiliatesAffiliateDashboard) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliatesAffiliateDashboard) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliatesAffiliateDashboard) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliatesAffiliateDashboard) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliatesAffiliateDashboard) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetReferredCount

`func (o *AffiliatesAffiliateDashboard) GetReferredCount() int32`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AffiliatesAffiliateDashboard) GetReferredCountOk() (*int32, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AffiliatesAffiliateDashboard) SetReferredCount(v int32)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AffiliatesAffiliateDashboard) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AffiliatesAffiliateDashboard) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliatesAffiliateDashboard) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliatesAffiliateDashboard) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliatesAffiliateDashboard) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliatesAffiliateDashboard) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliatesAffiliateDashboard) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliatesAffiliateDashboard) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliatesAffiliateDashboard) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliatesAffiliateDashboard) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliatesAffiliateDashboard) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliatesAffiliateDashboard) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliatesAffiliateDashboard) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AffiliatesAffiliateDashboard) GetPayouts() []AffiliatesPayout`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AffiliatesAffiliateDashboard) GetPayoutsOk() (*[]AffiliatesPayout, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AffiliatesAffiliateDashboard) SetPayouts(v []AffiliatesPayout)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AffiliatesAffiliateDashboard) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


