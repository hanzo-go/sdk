# CloudMyReferrals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the org&#39;s STABLE referral code — a deterministic function of the org id, so it never changes and never has to be stored to be reproduced. | [optional] 
**Counts** | Pointer to [**CloudStatusCounts**](CloudStatusCounts.md) | Counts tallies this org&#39;s referrals by status. | [optional] 
**CreditsEarnedCents** | Pointer to **int32** | CreditsEarnedCents is the total promo credit this org has earned as the REFERRER, in USD cents. | [optional] 
**Link** | Pointer to **string** | Link is the shareable signup link carrying the code, on the brand&#39;s own host. | [optional] 
**RefereeBonusCents** | Pointer to **int32** | RefereeBonusCents is what a referee is granted on qualification, in USD cents. | [optional] 
**Referrals** | Pointer to [**[]CloudMyReferralView**](CloudMyReferralView.md) | Referrals is one row per org that signed up with this code. | [optional] 
**ReferrerBonusCents** | Pointer to **int32** | ReferrerBonusCents is what the referrer is granted when a referee qualifies, in USD cents. | [optional] 

## Methods

### NewCloudMyReferrals

`func NewCloudMyReferrals() *CloudMyReferrals`

NewCloudMyReferrals instantiates a new CloudMyReferrals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudMyReferralsWithDefaults

`func NewCloudMyReferralsWithDefaults() *CloudMyReferrals`

NewCloudMyReferralsWithDefaults instantiates a new CloudMyReferrals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudMyReferrals) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudMyReferrals) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudMyReferrals) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudMyReferrals) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCounts

`func (o *CloudMyReferrals) GetCounts() CloudStatusCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *CloudMyReferrals) GetCountsOk() (*CloudStatusCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *CloudMyReferrals) SetCounts(v CloudStatusCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *CloudMyReferrals) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetCreditsEarnedCents

`func (o *CloudMyReferrals) GetCreditsEarnedCents() int32`

GetCreditsEarnedCents returns the CreditsEarnedCents field if non-nil, zero value otherwise.

### GetCreditsEarnedCentsOk

`func (o *CloudMyReferrals) GetCreditsEarnedCentsOk() (*int32, bool)`

GetCreditsEarnedCentsOk returns a tuple with the CreditsEarnedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsEarnedCents

`func (o *CloudMyReferrals) SetCreditsEarnedCents(v int32)`

SetCreditsEarnedCents sets CreditsEarnedCents field to given value.

### HasCreditsEarnedCents

`func (o *CloudMyReferrals) HasCreditsEarnedCents() bool`

HasCreditsEarnedCents returns a boolean if a field has been set.

### GetLink

`func (o *CloudMyReferrals) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *CloudMyReferrals) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *CloudMyReferrals) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *CloudMyReferrals) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRefereeBonusCents

`func (o *CloudMyReferrals) GetRefereeBonusCents() int32`

GetRefereeBonusCents returns the RefereeBonusCents field if non-nil, zero value otherwise.

### GetRefereeBonusCentsOk

`func (o *CloudMyReferrals) GetRefereeBonusCentsOk() (*int32, bool)`

GetRefereeBonusCentsOk returns a tuple with the RefereeBonusCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeBonusCents

`func (o *CloudMyReferrals) SetRefereeBonusCents(v int32)`

SetRefereeBonusCents sets RefereeBonusCents field to given value.

### HasRefereeBonusCents

`func (o *CloudMyReferrals) HasRefereeBonusCents() bool`

HasRefereeBonusCents returns a boolean if a field has been set.

### GetReferrals

`func (o *CloudMyReferrals) GetReferrals() []CloudMyReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *CloudMyReferrals) GetReferralsOk() (*[]CloudMyReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *CloudMyReferrals) SetReferrals(v []CloudMyReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *CloudMyReferrals) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetReferrerBonusCents

`func (o *CloudMyReferrals) GetReferrerBonusCents() int32`

GetReferrerBonusCents returns the ReferrerBonusCents field if non-nil, zero value otherwise.

### GetReferrerBonusCentsOk

`func (o *CloudMyReferrals) GetReferrerBonusCentsOk() (*int32, bool)`

GetReferrerBonusCentsOk returns a tuple with the ReferrerBonusCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerBonusCents

`func (o *CloudMyReferrals) SetReferrerBonusCents(v int32)`

SetReferrerBonusCents sets ReferrerBonusCents field to given value.

### HasReferrerBonusCents

`func (o *CloudMyReferrals) HasReferrerBonusCents() bool`

HasReferrerBonusCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


