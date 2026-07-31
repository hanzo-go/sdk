# AffiliatesGetMyAffiliate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAffiliate** | Pointer to **bool** |  | [optional] 
**DefaultRateBps** | Pointer to **int64** | The default commission rate a new affiliate gets, in basis points (2000 &#x3D; 20%). | [optional] 
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

### NewAffiliatesGetMyAffiliate200Response

`func NewAffiliatesGetMyAffiliate200Response() *AffiliatesGetMyAffiliate200Response`

NewAffiliatesGetMyAffiliate200Response instantiates a new AffiliatesGetMyAffiliate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesGetMyAffiliate200ResponseWithDefaults

`func NewAffiliatesGetMyAffiliate200ResponseWithDefaults() *AffiliatesGetMyAffiliate200Response`

NewAffiliatesGetMyAffiliate200ResponseWithDefaults instantiates a new AffiliatesGetMyAffiliate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAffiliate

`func (o *AffiliatesGetMyAffiliate200Response) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliatesGetMyAffiliate200Response) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliatesGetMyAffiliate200Response) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliatesGetMyAffiliate200Response) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetDefaultRateBps

`func (o *AffiliatesGetMyAffiliate200Response) GetDefaultRateBps() int64`

GetDefaultRateBps returns the DefaultRateBps field if non-nil, zero value otherwise.

### GetDefaultRateBpsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetDefaultRateBpsOk() (*int64, bool)`

GetDefaultRateBpsOk returns a tuple with the DefaultRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRateBps

`func (o *AffiliatesGetMyAffiliate200Response) SetDefaultRateBps(v int64)`

SetDefaultRateBps sets DefaultRateBps field to given value.

### HasDefaultRateBps

`func (o *AffiliatesGetMyAffiliate200Response) HasDefaultRateBps() bool`

HasDefaultRateBps returns a boolean if a field has been set.

### GetId

`func (o *AffiliatesGetMyAffiliate200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesGetMyAffiliate200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesGetMyAffiliate200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesGetMyAffiliate200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliatesGetMyAffiliate200Response) GetStatus() AffiliatesAffiliateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliatesGetMyAffiliate200Response) GetStatusOk() (*AffiliatesAffiliateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliatesGetMyAffiliate200Response) SetStatus(v AffiliatesAffiliateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliatesGetMyAffiliate200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCode

`func (o *AffiliatesGetMyAffiliate200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliatesGetMyAffiliate200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliatesGetMyAffiliate200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliatesGetMyAffiliate200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliatesGetMyAffiliate200Response) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliatesGetMyAffiliate200Response) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliatesGetMyAffiliate200Response) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliatesGetMyAffiliate200Response) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetLink

`func (o *AffiliatesGetMyAffiliate200Response) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AffiliatesGetMyAffiliate200Response) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AffiliatesGetMyAffiliate200Response) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AffiliatesGetMyAffiliate200Response) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliatesGetMyAffiliate200Response) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliatesGetMyAffiliate200Response) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliatesGetMyAffiliate200Response) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetReferredCount

`func (o *AffiliatesGetMyAffiliate200Response) GetReferredCount() int32`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AffiliatesGetMyAffiliate200Response) GetReferredCountOk() (*int32, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AffiliatesGetMyAffiliate200Response) SetReferredCount(v int32)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AffiliatesGetMyAffiliate200Response) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AffiliatesGetMyAffiliate200Response) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliatesGetMyAffiliate200Response) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliatesGetMyAffiliate200Response) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliatesGetMyAffiliate200Response) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliatesGetMyAffiliate200Response) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliatesGetMyAffiliate200Response) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliatesGetMyAffiliate200Response) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliatesGetMyAffiliate200Response) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliatesGetMyAffiliate200Response) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AffiliatesGetMyAffiliate200Response) GetPayouts() []AffiliatesPayout`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AffiliatesGetMyAffiliate200Response) GetPayoutsOk() (*[]AffiliatesPayout, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AffiliatesGetMyAffiliate200Response) SetPayouts(v []AffiliatesPayout)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AffiliatesGetMyAffiliate200Response) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


