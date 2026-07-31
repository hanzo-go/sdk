# CloudAdminBonusDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Referrals** | Pointer to [**[]CloudAdminReferralView**](CloudAdminReferralView.md) | Referrals is every referral in the ledger, both orgs exposed. | [optional] 
**Summary** | Pointer to [**CloudAdminSummary**](CloudAdminSummary.md) | Summary is the fleet tally across those referrals. | [optional] 

## Methods

### NewCloudAdminBonusDirectory

`func NewCloudAdminBonusDirectory() *CloudAdminBonusDirectory`

NewCloudAdminBonusDirectory instantiates a new CloudAdminBonusDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminBonusDirectoryWithDefaults

`func NewCloudAdminBonusDirectoryWithDefaults() *CloudAdminBonusDirectory`

NewCloudAdminBonusDirectoryWithDefaults instantiates a new CloudAdminBonusDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferrals

`func (o *CloudAdminBonusDirectory) GetReferrals() []CloudAdminReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *CloudAdminBonusDirectory) GetReferralsOk() (*[]CloudAdminReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *CloudAdminBonusDirectory) SetReferrals(v []CloudAdminReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *CloudAdminBonusDirectory) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetSummary

`func (o *CloudAdminBonusDirectory) GetSummary() CloudAdminSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *CloudAdminBonusDirectory) GetSummaryOk() (*CloudAdminSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *CloudAdminBonusDirectory) SetSummary(v CloudAdminSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *CloudAdminBonusDirectory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


