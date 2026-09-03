# OrgEarningView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionCents** | Pointer to **int64** | CommissionCents is what the caller earned from that org across ALL periods, in cents. Deliberately the caller&#39;s own share and nothing else: that org&#39;s spend and the margin on it are not restated here. | [optional] 
**ReferredOrg** | Pointer to **string** | ReferredOrg is the org slug this contribution came from — one the caller referred, directly or up to three levels down. | [optional] 

## Methods

### NewOrgEarningView

`func NewOrgEarningView() *OrgEarningView`

NewOrgEarningView instantiates a new OrgEarningView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgEarningViewWithDefaults

`func NewOrgEarningViewWithDefaults() *OrgEarningView`

NewOrgEarningViewWithDefaults instantiates a new OrgEarningView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionCents

`func (o *OrgEarningView) GetCommissionCents() int64`

GetCommissionCents returns the CommissionCents field if non-nil, zero value otherwise.

### GetCommissionCentsOk

`func (o *OrgEarningView) GetCommissionCentsOk() (*int64, bool)`

GetCommissionCentsOk returns a tuple with the CommissionCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionCents

`func (o *OrgEarningView) SetCommissionCents(v int64)`

SetCommissionCents sets CommissionCents field to given value.

### HasCommissionCents

`func (o *OrgEarningView) HasCommissionCents() bool`

HasCommissionCents returns a boolean if a field has been set.

### GetReferredOrg

`func (o *OrgEarningView) GetReferredOrg() string`

GetReferredOrg returns the ReferredOrg field if non-nil, zero value otherwise.

### GetReferredOrgOk

`func (o *OrgEarningView) GetReferredOrgOk() (*string, bool)`

GetReferredOrgOk returns a tuple with the ReferredOrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredOrg

`func (o *OrgEarningView) SetReferredOrg(v string)`

SetReferredOrg sets ReferredOrg field to given value.

### HasReferredOrg

`func (o *OrgEarningView) HasReferredOrg() bool`

HasReferredOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


