# ReferralsAdminReferralView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ReferrerOrg** | Pointer to **string** |  | [optional] 
**RefereeOrg** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ReferralsReferralStatus**](ReferralsReferralStatus.md) |  | [optional] 
**ReferrerGrantCents** | Pointer to **int64** |  | [optional] 
**RefereeGrantCents** | Pointer to **int64** |  | [optional] 
**ReferrerTxn** | Pointer to **string** | Ledger transaction id for the referrer bonus (omitted if empty). | [optional] 
**RefereeTxn** | Pointer to **string** | Ledger transaction id for the referee bonus (omitted if empty). | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**QualifiedAt** | Pointer to **int64** |  | [optional] 
**CreditedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewReferralsAdminReferralView

`func NewReferralsAdminReferralView() *ReferralsAdminReferralView`

NewReferralsAdminReferralView instantiates a new ReferralsAdminReferralView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralsAdminReferralViewWithDefaults

`func NewReferralsAdminReferralViewWithDefaults() *ReferralsAdminReferralView`

NewReferralsAdminReferralViewWithDefaults instantiates a new ReferralsAdminReferralView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferralsAdminReferralView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferralsAdminReferralView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferralsAdminReferralView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReferralsAdminReferralView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferrerOrg

`func (o *ReferralsAdminReferralView) GetReferrerOrg() string`

GetReferrerOrg returns the ReferrerOrg field if non-nil, zero value otherwise.

### GetReferrerOrgOk

`func (o *ReferralsAdminReferralView) GetReferrerOrgOk() (*string, bool)`

GetReferrerOrgOk returns a tuple with the ReferrerOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerOrg

`func (o *ReferralsAdminReferralView) SetReferrerOrg(v string)`

SetReferrerOrg sets ReferrerOrg field to given value.

### HasReferrerOrg

`func (o *ReferralsAdminReferralView) HasReferrerOrg() bool`

HasReferrerOrg returns a boolean if a field has been set.

### GetRefereeOrg

`func (o *ReferralsAdminReferralView) GetRefereeOrg() string`

GetRefereeOrg returns the RefereeOrg field if non-nil, zero value otherwise.

### GetRefereeOrgOk

`func (o *ReferralsAdminReferralView) GetRefereeOrgOk() (*string, bool)`

GetRefereeOrgOk returns a tuple with the RefereeOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeOrg

`func (o *ReferralsAdminReferralView) SetRefereeOrg(v string)`

SetRefereeOrg sets RefereeOrg field to given value.

### HasRefereeOrg

`func (o *ReferralsAdminReferralView) HasRefereeOrg() bool`

HasRefereeOrg returns a boolean if a field has been set.

### GetCode

`func (o *ReferralsAdminReferralView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferralsAdminReferralView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferralsAdminReferralView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferralsAdminReferralView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetStatus

`func (o *ReferralsAdminReferralView) GetStatus() ReferralsReferralStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferralsAdminReferralView) GetStatusOk() (*ReferralsReferralStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferralsAdminReferralView) SetStatus(v ReferralsReferralStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferralsAdminReferralView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReferrerGrantCents

`func (o *ReferralsAdminReferralView) GetReferrerGrantCents() int64`

GetReferrerGrantCents returns the ReferrerGrantCents field if non-nil, zero value otherwise.

### GetReferrerGrantCentsOk

`func (o *ReferralsAdminReferralView) GetReferrerGrantCentsOk() (*int64, bool)`

GetReferrerGrantCentsOk returns a tuple with the ReferrerGrantCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerGrantCents

`func (o *ReferralsAdminReferralView) SetReferrerGrantCents(v int64)`

SetReferrerGrantCents sets ReferrerGrantCents field to given value.

### HasReferrerGrantCents

`func (o *ReferralsAdminReferralView) HasReferrerGrantCents() bool`

HasReferrerGrantCents returns a boolean if a field has been set.

### GetRefereeGrantCents

`func (o *ReferralsAdminReferralView) GetRefereeGrantCents() int64`

GetRefereeGrantCents returns the RefereeGrantCents field if non-nil, zero value otherwise.

### GetRefereeGrantCentsOk

`func (o *ReferralsAdminReferralView) GetRefereeGrantCentsOk() (*int64, bool)`

GetRefereeGrantCentsOk returns a tuple with the RefereeGrantCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeGrantCents

`func (o *ReferralsAdminReferralView) SetRefereeGrantCents(v int64)`

SetRefereeGrantCents sets RefereeGrantCents field to given value.

### HasRefereeGrantCents

`func (o *ReferralsAdminReferralView) HasRefereeGrantCents() bool`

HasRefereeGrantCents returns a boolean if a field has been set.

### GetReferrerTxn

`func (o *ReferralsAdminReferralView) GetReferrerTxn() string`

GetReferrerTxn returns the ReferrerTxn field if non-nil, zero value otherwise.

### GetReferrerTxnOk

`func (o *ReferralsAdminReferralView) GetReferrerTxnOk() (*string, bool)`

GetReferrerTxnOk returns a tuple with the ReferrerTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerTxn

`func (o *ReferralsAdminReferralView) SetReferrerTxn(v string)`

SetReferrerTxn sets ReferrerTxn field to given value.

### HasReferrerTxn

`func (o *ReferralsAdminReferralView) HasReferrerTxn() bool`

HasReferrerTxn returns a boolean if a field has been set.

### GetRefereeTxn

`func (o *ReferralsAdminReferralView) GetRefereeTxn() string`

GetRefereeTxn returns the RefereeTxn field if non-nil, zero value otherwise.

### GetRefereeTxnOk

`func (o *ReferralsAdminReferralView) GetRefereeTxnOk() (*string, bool)`

GetRefereeTxnOk returns a tuple with the RefereeTxn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefereeTxn

`func (o *ReferralsAdminReferralView) SetRefereeTxn(v string)`

SetRefereeTxn sets RefereeTxn field to given value.

### HasRefereeTxn

`func (o *ReferralsAdminReferralView) HasRefereeTxn() bool`

HasRefereeTxn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReferralsAdminReferralView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReferralsAdminReferralView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReferralsAdminReferralView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReferralsAdminReferralView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetQualifiedAt

`func (o *ReferralsAdminReferralView) GetQualifiedAt() int64`

GetQualifiedAt returns the QualifiedAt field if non-nil, zero value otherwise.

### GetQualifiedAtOk

`func (o *ReferralsAdminReferralView) GetQualifiedAtOk() (*int64, bool)`

GetQualifiedAtOk returns a tuple with the QualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedAt

`func (o *ReferralsAdminReferralView) SetQualifiedAt(v int64)`

SetQualifiedAt sets QualifiedAt field to given value.

### HasQualifiedAt

`func (o *ReferralsAdminReferralView) HasQualifiedAt() bool`

HasQualifiedAt returns a boolean if a field has been set.

### GetCreditedAt

`func (o *ReferralsAdminReferralView) GetCreditedAt() int64`

GetCreditedAt returns the CreditedAt field if non-nil, zero value otherwise.

### GetCreditedAtOk

`func (o *ReferralsAdminReferralView) GetCreditedAtOk() (*int64, bool)`

GetCreditedAtOk returns a tuple with the CreditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditedAt

`func (o *ReferralsAdminReferralView) SetCreditedAt(v int64)`

SetCreditedAt sets CreditedAt field to given value.

### HasCreditedAt

`func (o *ReferralsAdminReferralView) HasCreditedAt() bool`

HasCreditedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


