# AdminBonusDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Referrals** | Pointer to [**[]AdminReferralView**](AdminReferralView.md) | Referrals is every referral in the directory, both orgs exposed. | [optional] 
**Summary** | Pointer to [**AdminSummary**](AdminSummary.md) | Summary is the fleet tally across those referrals. | [optional] 

## Methods

### NewAdminBonusDirectory

`func NewAdminBonusDirectory() *AdminBonusDirectory`

NewAdminBonusDirectory instantiates a new AdminBonusDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminBonusDirectoryWithDefaults

`func NewAdminBonusDirectoryWithDefaults() *AdminBonusDirectory`

NewAdminBonusDirectoryWithDefaults instantiates a new AdminBonusDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferrals

`func (o *AdminBonusDirectory) GetReferrals() []AdminReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *AdminBonusDirectory) GetReferralsOk() (*[]AdminReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *AdminBonusDirectory) SetReferrals(v []AdminReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *AdminBonusDirectory) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetSummary

`func (o *AdminBonusDirectory) GetSummary() AdminSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdminBonusDirectory) GetSummaryOk() (*AdminSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdminBonusDirectory) SetSummary(v AdminSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdminBonusDirectory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


