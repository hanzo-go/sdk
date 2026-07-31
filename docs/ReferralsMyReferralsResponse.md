# ReferralsMyReferralsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The caller org&#39;s stable referral code. | [optional] 
**Link** | Pointer to **string** | The share link, &#x60;&lt;linkBase&gt;/?ref&#x3D;&lt;code&gt;&#x60;. | [optional] 
**ReferrerBonusCents** | Pointer to **int64** | Bonus granted to the referrer on qualification (1000 &#x3D; $10). | [optional] 
**RefereeBonusCents** | Pointer to **int64** | Bonus granted to the referee on qualification (500 &#x3D; $5). | [optional] 
**CreditsEarnedCents** | Pointer to **int64** | Total credit earned across all of the caller&#39;s referrals. | [optional] 
**Counts** | Pointer to [**ReferralsStatusCounts**](ReferralsStatusCounts.md) |  | [optional] 
**Referrals** | Pointer to [**[]ReferralsMyReferralView**](ReferralsMyReferralView.md) |  | [optional] 

## Methods

### NewReferralsMyReferralsResponse

`func NewReferralsMyReferralsResponse() *ReferralsMyReferralsResponse`

NewReferralsMyReferralsResponse instantiates a new ReferralsMyReferralsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsMyReferralsResponseWithDefaults

`func NewReferralsMyReferralsResponseWithDefaults() *ReferralsMyReferralsResponse`

NewReferralsMyReferralsResponseWithDefaults instantiates a new ReferralsMyReferralsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReferralsMyReferralsResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferralsMyReferralsResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferralsMyReferralsResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferralsMyReferralsResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLink

`func (o *ReferralsMyReferralsResponse) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ReferralsMyReferralsResponse) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ReferralsMyReferralsResponse) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ReferralsMyReferralsResponse) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReferrerBonusCents

`func (o *ReferralsMyReferralsResponse) GetReferrerBonusCents() int64`

GetReferrerBonusCents returns the ReferrerBonusCents field if non-nil, zero value otherwise.

### GetReferrerBonusCentsOk

`func (o *ReferralsMyReferralsResponse) GetReferrerBonusCentsOk() (*int64, bool)`

GetReferrerBonusCentsOk returns a tuple with the ReferrerBonusCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerBonusCents

`func (o *ReferralsMyReferralsResponse) SetReferrerBonusCents(v int64)`

SetReferrerBonusCents sets ReferrerBonusCents field to given value.

### HasReferrerBonusCents

`func (o *ReferralsMyReferralsResponse) HasReferrerBonusCents() bool`

HasReferrerBonusCents returns a boolean if a field has been set.

### GetRefereeBonusCents

`func (o *ReferralsMyReferralsResponse) GetRefereeBonusCents() int64`

GetRefereeBonusCents returns the RefereeBonusCents field if non-nil, zero value otherwise.

### GetRefereeBonusCentsOk

`func (o *ReferralsMyReferralsResponse) GetRefereeBonusCentsOk() (*int64, bool)`

GetRefereeBonusCentsOk returns a tuple with the RefereeBonusCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeBonusCents

`func (o *ReferralsMyReferralsResponse) SetRefereeBonusCents(v int64)`

SetRefereeBonusCents sets RefereeBonusCents field to given value.

### HasRefereeBonusCents

`func (o *ReferralsMyReferralsResponse) HasRefereeBonusCents() bool`

HasRefereeBonusCents returns a boolean if a field has been set.

### GetCreditsEarnedCents

`func (o *ReferralsMyReferralsResponse) GetCreditsEarnedCents() int64`

GetCreditsEarnedCents returns the CreditsEarnedCents field if non-nil, zero value otherwise.

### GetCreditsEarnedCentsOk

`func (o *ReferralsMyReferralsResponse) GetCreditsEarnedCentsOk() (*int64, bool)`

GetCreditsEarnedCentsOk returns a tuple with the CreditsEarnedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsEarnedCents

`func (o *ReferralsMyReferralsResponse) SetCreditsEarnedCents(v int64)`

SetCreditsEarnedCents sets CreditsEarnedCents field to given value.

### HasCreditsEarnedCents

`func (o *ReferralsMyReferralsResponse) HasCreditsEarnedCents() bool`

HasCreditsEarnedCents returns a boolean if a field has been set.

### GetCounts

`func (o *ReferralsMyReferralsResponse) GetCounts() ReferralsStatusCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ReferralsMyReferralsResponse) GetCountsOk() (*ReferralsStatusCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ReferralsMyReferralsResponse) SetCounts(v ReferralsStatusCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *ReferralsMyReferralsResponse) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetReferrals

`func (o *ReferralsMyReferralsResponse) GetReferrals() []ReferralsMyReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *ReferralsMyReferralsResponse) GetReferralsOk() (*[]ReferralsMyReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *ReferralsMyReferralsResponse) SetReferrals(v []ReferralsMyReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *ReferralsMyReferralsResponse) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


