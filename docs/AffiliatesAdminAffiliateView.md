# AffiliatesAdminAffiliateView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**RequestedCode** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AffiliatesAffiliateStatus**](AffiliatesAffiliateStatus.md) |  | [optional] 
**RateBps** | Pointer to **int64** |  | [optional] 
**ReferredCount** | Pointer to **int32** |  | [optional] 
**AccruedCents** | Pointer to **int64** |  | [optional] 
**PendingCents** | Pointer to **int64** |  | [optional] 
**PaidCents** | Pointer to **int64** |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**ApprovedAt** | Pointer to **int64** |  | [optional] 
**SuspendedAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewAffiliatesAdminAffiliateView

`func NewAffiliatesAdminAffiliateView() *AffiliatesAdminAffiliateView`

NewAffiliatesAdminAffiliateView instantiates a new AffiliatesAdminAffiliateView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliatesAdminAffiliateViewWithDefaults

`func NewAffiliatesAdminAffiliateViewWithDefaults() *AffiliatesAdminAffiliateView`

NewAffiliatesAdminAffiliateViewWithDefaults instantiates a new AffiliatesAdminAffiliateView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AffiliatesAdminAffiliateView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliatesAdminAffiliateView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliatesAdminAffiliateView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliatesAdminAffiliateView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *AffiliatesAdminAffiliateView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AffiliatesAdminAffiliateView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AffiliatesAdminAffiliateView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AffiliatesAdminAffiliateView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetCode

`func (o *AffiliatesAdminAffiliateView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliatesAdminAffiliateView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliatesAdminAffiliateView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliatesAdminAffiliateView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliatesAdminAffiliateView) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliatesAdminAffiliateView) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliatesAdminAffiliateView) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliatesAdminAffiliateView) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliatesAdminAffiliateView) GetStatus() AffiliatesAffiliateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliatesAdminAffiliateView) GetStatusOk() (*AffiliatesAffiliateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliatesAdminAffiliateView) SetStatus(v AffiliatesAffiliateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliatesAdminAffiliateView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliatesAdminAffiliateView) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliatesAdminAffiliateView) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliatesAdminAffiliateView) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliatesAdminAffiliateView) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetReferredCount

`func (o *AffiliatesAdminAffiliateView) GetReferredCount() int32`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AffiliatesAdminAffiliateView) GetReferredCountOk() (*int32, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AffiliatesAdminAffiliateView) SetReferredCount(v int32)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AffiliatesAdminAffiliateView) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetAccruedCents

`func (o *AffiliatesAdminAffiliateView) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliatesAdminAffiliateView) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliatesAdminAffiliateView) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliatesAdminAffiliateView) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliatesAdminAffiliateView) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliatesAdminAffiliateView) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliatesAdminAffiliateView) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliatesAdminAffiliateView) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliatesAdminAffiliateView) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliatesAdminAffiliateView) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliatesAdminAffiliateView) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliatesAdminAffiliateView) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AffiliatesAdminAffiliateView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AffiliatesAdminAffiliateView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AffiliatesAdminAffiliateView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AffiliatesAdminAffiliateView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AffiliatesAdminAffiliateView) GetApprovedAt() int64`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AffiliatesAdminAffiliateView) GetApprovedAtOk() (*int64, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AffiliatesAdminAffiliateView) SetApprovedAt(v int64)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AffiliatesAdminAffiliateView) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *AffiliatesAdminAffiliateView) GetSuspendedAt() int64`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *AffiliatesAdminAffiliateView) GetSuspendedAtOk() (*int64, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *AffiliatesAdminAffiliateView) SetSuspendedAt(v int64)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *AffiliatesAdminAffiliateView) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


