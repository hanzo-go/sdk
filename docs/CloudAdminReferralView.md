# CloudAdminReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the referral code the referral was recorded against. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the referral was recorded, as a Unix timestamp. | [optional] 
**CreditedAt** | Pointer to **int32** | CreditedAt is when the bonuses were latched and granted, as a Unix timestamp; 0 until they are. | [optional] 
**Id** | Pointer to **string** | ID is the referral&#39;s handle. | [optional] 
**QualifiedAt** | Pointer to **int32** | QualifiedAt is when the referee first made metered spend, as a Unix timestamp; 0 while still pending. | [optional] 
**RefereeGrantCents** | Pointer to **int32** | RefereeGrantCents is what the referee was granted, in USD cents; 0 until the referral is credited. | [optional] 
**RefereeOrg** | Pointer to **string** | RefereeOrg is the org that signed up with it. | [optional] 
**RefereeTxn** | Pointer to **string** | RefereeTxn is the commerce ledger transaction that carried the referee&#39;s grant, omitted until one exists. | [optional] 
**ReferrerGrantCents** | Pointer to **int32** | ReferrerGrantCents is what the referrer was granted, in USD cents; 0 until the referral is credited. | [optional] 
**ReferrerOrg** | Pointer to **string** | ReferrerOrg is the org whose code was used. | [optional] 
**ReferrerTxn** | Pointer to **string** | ReferrerTxn is the commerce ledger transaction that carried the referrer&#39;s grant, omitted until one exists. | [optional] 
**Status** | Pointer to **string** | Status is the referral&#39;s lifecycle state: \&quot;signup\&quot;, \&quot;qualified\&quot; or \&quot;credited\&quot;. | [optional] 

## Methods

### NewCloudAdminReferralView

`func NewCloudAdminReferralView() *CloudAdminReferralView`

NewCloudAdminReferralView instantiates a new CloudAdminReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAdminReferralViewWithDefaults

`func NewCloudAdminReferralViewWithDefaults() *CloudAdminReferralView`

NewCloudAdminReferralViewWithDefaults instantiates a new CloudAdminReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudAdminReferralView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudAdminReferralView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudAdminReferralView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudAdminReferralView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAdminReferralView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAdminReferralView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAdminReferralView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAdminReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditedAt

`func (o *CloudAdminReferralView) GetCreditedAt() int32`

GetCreditedAt returns the CreditedAt field if non-nil, zero value otherwise.

### GetCreditedAtOk

`func (o *CloudAdminReferralView) GetCreditedAtOk() (*int32, bool)`

GetCreditedAtOk returns a tuple with the CreditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedAt

`func (o *CloudAdminReferralView) SetCreditedAt(v int32)`

SetCreditedAt sets CreditedAt field to given value.

### HasCreditedAt

`func (o *CloudAdminReferralView) HasCreditedAt() bool`

HasCreditedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudAdminReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAdminReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAdminReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAdminReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *CloudAdminReferralView) GetQualifiedAt() int32`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *CloudAdminReferralView) GetQualifiedAtOk() (*int32, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *CloudAdminReferralView) SetQualifiedAt(v int32)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *CloudAdminReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetRefereeGrantCents

`func (o *CloudAdminReferralView) GetRefereeGrantCents() int32`

GetRefereeGrantCents returns the RefereeGrantCents field if non-nil, zero value otherwise.

### GetRefereeGrantCentsOk

`func (o *CloudAdminReferralView) GetRefereeGrantCentsOk() (*int32, bool)`

GetRefereeGrantCentsOk returns a tuple with the RefereeGrantCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeGrantCents

`func (o *CloudAdminReferralView) SetRefereeGrantCents(v int32)`

SetRefereeGrantCents sets RefereeGrantCents field to given value.

### HasRefereeGrantCents

`func (o *CloudAdminReferralView) HasRefereeGrantCents() bool`

HasRefereeGrantCents returns a boolean if a field has been set.

### GetRefereeOrg

`func (o *CloudAdminReferralView) GetRefereeOrg() string`

GetRefereeOrg returns the RefereeOrg field if non-nil, zero value otherwise.

### GetRefereeOrgOk

`func (o *CloudAdminReferralView) GetRefereeOrgOk() (*string, bool)`

GetRefereeOrgOk returns a tuple with the RefereeOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeOrg

`func (o *CloudAdminReferralView) SetRefereeOrg(v string)`

SetRefereeOrg sets RefereeOrg field to given value.

### HasRefereeOrg

`func (o *CloudAdminReferralView) HasRefereeOrg() bool`

HasRefereeOrg returns a boolean if a field has been set.

### GetRefereeTxn

`func (o *CloudAdminReferralView) GetRefereeTxn() string`

GetRefereeTxn returns the RefereeTxn field if non-nil, zero value otherwise.

### GetRefereeTxnOk

`func (o *CloudAdminReferralView) GetRefereeTxnOk() (*string, bool)`

GetRefereeTxnOk returns a tuple with the RefereeTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeTxn

`func (o *CloudAdminReferralView) SetRefereeTxn(v string)`

SetRefereeTxn sets RefereeTxn field to given value.

### HasRefereeTxn

`func (o *CloudAdminReferralView) HasRefereeTxn() bool`

HasRefereeTxn returns a boolean if a field has been set.

### GetReferrerGrantCents

`func (o *CloudAdminReferralView) GetReferrerGrantCents() int32`

GetReferrerGrantCents returns the ReferrerGrantCents field if non-nil, zero value otherwise.

### GetReferrerGrantCentsOk

`func (o *CloudAdminReferralView) GetReferrerGrantCentsOk() (*int32, bool)`

GetReferrerGrantCentsOk returns a tuple with the ReferrerGrantCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerGrantCents

`func (o *CloudAdminReferralView) SetReferrerGrantCents(v int32)`

SetReferrerGrantCents sets ReferrerGrantCents field to given value.

### HasReferrerGrantCents

`func (o *CloudAdminReferralView) HasReferrerGrantCents() bool`

HasReferrerGrantCents returns a boolean if a field has been set.

### GetReferrerOrg

`func (o *CloudAdminReferralView) GetReferrerOrg() string`

GetReferrerOrg returns the ReferrerOrg field if non-nil, zero value otherwise.

### GetReferrerOrgOk

`func (o *CloudAdminReferralView) GetReferrerOrgOk() (*string, bool)`

GetReferrerOrgOk returns a tuple with the ReferrerOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerOrg

`func (o *CloudAdminReferralView) SetReferrerOrg(v string)`

SetReferrerOrg sets ReferrerOrg field to given value.

### HasReferrerOrg

`func (o *CloudAdminReferralView) HasReferrerOrg() bool`

HasReferrerOrg returns a boolean if a field has been set.

### GetReferrerTxn

`func (o *CloudAdminReferralView) GetReferrerTxn() string`

GetReferrerTxn returns the ReferrerTxn field if non-nil, zero value otherwise.

### GetReferrerTxnOk

`func (o *CloudAdminReferralView) GetReferrerTxnOk() (*string, bool)`

GetReferrerTxnOk returns a tuple with the ReferrerTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerTxn

`func (o *CloudAdminReferralView) SetReferrerTxn(v string)`

SetReferrerTxn sets ReferrerTxn field to given value.

### HasReferrerTxn

`func (o *CloudAdminReferralView) HasReferrerTxn() bool`

HasReferrerTxn returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAdminReferralView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAdminReferralView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAdminReferralView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAdminReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


