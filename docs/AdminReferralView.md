# AdminReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the referral code the referral was recorded against. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the referral was recorded, as a Unix timestamp. | [optional] 
**Id** | Pointer to **string** | ID is the referral&#39;s handle. | [optional] 
**QualifiedAt** | Pointer to **int32** | QualifiedAt is when the referee first made metered spend, as a Unix timestamp; 0 while still pending. | [optional] 
**RefereeOrg** | Pointer to **string** | RefereeOrg is the org that signed up with it. | [optional] 
**ReferrerOrg** | Pointer to **string** | ReferrerOrg is the org whose code was used. | [optional] 
**Status** | Pointer to **string** | Status is the referral&#39;s lifecycle state: \&quot;signup\&quot; or \&quot;qualified\&quot;. | [optional] 

## Methods

### NewAdminReferralView

`func NewAdminReferralView() *AdminReferralView`

NewAdminReferralView instantiates a new AdminReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminReferralViewWithDefaults

`func NewAdminReferralViewWithDefaults() *AdminReferralView`

NewAdminReferralViewWithDefaults instantiates a new AdminReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AdminReferralView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AdminReferralView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AdminReferralView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AdminReferralView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminReferralView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminReferralView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminReferralView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *AdminReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *AdminReferralView) GetQualifiedAt() int32`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *AdminReferralView) GetQualifiedAtOk() (*int32, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *AdminReferralView) SetQualifiedAt(v int32)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *AdminReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetRefereeOrg

`func (o *AdminReferralView) GetRefereeOrg() string`

GetRefereeOrg returns the RefereeOrg field if non-nil, zero value otherwise.

### GetRefereeOrgOk

`func (o *AdminReferralView) GetRefereeOrgOk() (*string, bool)`

GetRefereeOrgOk returns a tuple with the RefereeOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeOrg

`func (o *AdminReferralView) SetRefereeOrg(v string)`

SetRefereeOrg sets RefereeOrg field to given value.

### HasRefereeOrg

`func (o *AdminReferralView) HasRefereeOrg() bool`

HasRefereeOrg returns a boolean if a field has been set.

### GetReferrerOrg

`func (o *AdminReferralView) GetReferrerOrg() string`

GetReferrerOrg returns the ReferrerOrg field if non-nil, zero value otherwise.

### GetReferrerOrgOk

`func (o *AdminReferralView) GetReferrerOrgOk() (*string, bool)`

GetReferrerOrgOk returns a tuple with the ReferrerOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerOrg

`func (o *AdminReferralView) SetReferrerOrg(v string)`

SetReferrerOrg sets ReferrerOrg field to given value.

### HasReferrerOrg

`func (o *AdminReferralView) HasReferrerOrg() bool`

HasReferrerOrg returns a boolean if a field has been set.

### GetStatus

`func (o *AdminReferralView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminReferralView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminReferralView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


