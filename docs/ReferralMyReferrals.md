# ReferralMyReferrals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the org&#39;s STABLE referral code — a deterministic function of the org id, so it never changes and never has to be stored to be reproduced. | [optional] 
**Counts** | Pointer to [**ReferralStatusCounts**](ReferralStatusCounts.md) | Counts tallies this org&#39;s referrals by status. | [optional] 
**Link** | Pointer to **string** | Link is the shareable signup link carrying the code, on the brand&#39;s own host. | [optional] 
**Referrals** | Pointer to [**[]ReferralMyReferralView**](ReferralMyReferralView.md) | Referrals is one row per org that signed up with this code. | [optional] 

## Methods

### NewReferralMyReferrals

`func NewReferralMyReferrals() *ReferralMyReferrals`

NewReferralMyReferrals instantiates a new ReferralMyReferrals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralMyReferralsWithDefaults

`func NewReferralMyReferralsWithDefaults() *ReferralMyReferrals`

NewReferralMyReferralsWithDefaults instantiates a new ReferralMyReferrals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ReferralMyReferrals) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferralMyReferrals) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferralMyReferrals) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferralMyReferrals) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCounts

`func (o *ReferralMyReferrals) GetCounts() ReferralStatusCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ReferralMyReferrals) GetCountsOk() (*ReferralStatusCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ReferralMyReferrals) SetCounts(v ReferralStatusCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *ReferralMyReferrals) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetLink

`func (o *ReferralMyReferrals) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ReferralMyReferrals) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ReferralMyReferrals) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ReferralMyReferrals) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReferrals

`func (o *ReferralMyReferrals) GetReferrals() []ReferralMyReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *ReferralMyReferrals) GetReferralsOk() (*[]ReferralMyReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *ReferralMyReferrals) SetReferrals(v []ReferralMyReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *ReferralMyReferrals) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


