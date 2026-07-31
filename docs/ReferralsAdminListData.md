# ReferralsAdminListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Referrals** | Pointer to [**[]ReferralsAdminReferralView**](ReferralsAdminReferralView.md) |  | [optional] 
**Summary** | Pointer to [**ReferralsAdminSummary**](ReferralsAdminSummary.md) |  | [optional] 

## Methods

### NewReferralsAdminListData

`func NewReferralsAdminListData() *ReferralsAdminListData`

NewReferralsAdminListData instantiates a new ReferralsAdminListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsAdminListDataWithDefaults

`func NewReferralsAdminListDataWithDefaults() *ReferralsAdminListData`

NewReferralsAdminListDataWithDefaults instantiates a new ReferralsAdminListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferrals

`func (o *ReferralsAdminListData) GetReferrals() []ReferralsAdminReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *ReferralsAdminListData) GetReferralsOk() (*[]ReferralsAdminReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *ReferralsAdminListData) SetReferrals(v []ReferralsAdminReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *ReferralsAdminListData) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetSummary

`func (o *ReferralsAdminListData) GetSummary() ReferralsAdminSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReferralsAdminListData) GetSummaryOk() (*ReferralsAdminSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReferralsAdminListData) SetSummary(v ReferralsAdminSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReferralsAdminListData) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


