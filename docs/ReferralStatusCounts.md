# ReferralStatusCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Qualified** | Pointer to **int64** | Qualified is how many referees have made metered spend. | [optional] 
**Signup** | Pointer to **int64** | Signup is how many referees have signed up but not yet spent. | [optional] 
**Total** | Pointer to **int64** | Total is every referral this org has made. | [optional] 

## Methods

### NewReferralStatusCounts

`func NewReferralStatusCounts() *ReferralStatusCounts`

NewReferralStatusCounts instantiates a new ReferralStatusCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralStatusCountsWithDefaults

`func NewReferralStatusCountsWithDefaults() *ReferralStatusCounts`

NewReferralStatusCountsWithDefaults instantiates a new ReferralStatusCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQualified

`func (o *ReferralStatusCounts) GetQualified() int64`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *ReferralStatusCounts) GetQualifiedOk() (*int64, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *ReferralStatusCounts) SetQualified(v int64)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *ReferralStatusCounts) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetSignup

`func (o *ReferralStatusCounts) GetSignup() int64`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *ReferralStatusCounts) GetSignupOk() (*int64, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *ReferralStatusCounts) SetSignup(v int64)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *ReferralStatusCounts) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetTotal

`func (o *ReferralStatusCounts) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReferralStatusCounts) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReferralStatusCounts) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ReferralStatusCounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


