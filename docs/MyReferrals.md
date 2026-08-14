# MyReferrals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the org&#39;s STABLE referral code — a deterministic function of the org id, so it never changes and never has to be stored to be reproduced. | [optional] 
**Counts** | Pointer to [**StatusCounts**](StatusCounts.md) | Counts tallies this org&#39;s referrals by status. | [optional] 
**Link** | Pointer to **string** | Link is the shareable signup link carrying the code, on the brand&#39;s own host. | [optional] 
**Referrals** | Pointer to [**[]MyReferralView**](MyReferralView.md) | Referrals is one row per org that signed up with this code. | [optional] 

## Methods

### NewMyReferrals

`func NewMyReferrals() *MyReferrals`

NewMyReferrals instantiates a new MyReferrals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyReferralsWithDefaults

`func NewMyReferralsWithDefaults() *MyReferrals`

NewMyReferralsWithDefaults instantiates a new MyReferrals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MyReferrals) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MyReferrals) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MyReferrals) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MyReferrals) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCounts

`func (o *MyReferrals) GetCounts() StatusCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MyReferrals) GetCountsOk() (*StatusCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MyReferrals) SetCounts(v StatusCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MyReferrals) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetLink

`func (o *MyReferrals) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *MyReferrals) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *MyReferrals) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *MyReferrals) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReferrals

`func (o *MyReferrals) GetReferrals() []MyReferralView`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *MyReferrals) GetReferralsOk() (*[]MyReferralView, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *MyReferrals) SetReferrals(v []MyReferralView)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *MyReferrals) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


