# CaptableHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDiluted** | Pointer to **int64** | FullyDiluted is shares plus options. | [optional] 
**Name** | Pointer to **string** | Name is the stakeholder&#39;s name. | [optional] 
**Options** | Pointer to **int64** | Options is the shares under this stakeholder&#39;s non-terminal option grants. | [optional] 
**OwnershipPct** | Pointer to **float64** | OwnershipPct is fullyDiluted as a percentage of the company&#39;s fullyDilutedShares, rounded to two decimals; 0 when nothing is issued. | [optional] 
**Shares** | Pointer to **int64** | Shares is the shares this stakeholder holds by certificate. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID addresses the stakeholder these totals are for. | [optional] 

## Methods

### NewCaptableHolding

`func NewCaptableHolding() *CaptableHolding`

NewCaptableHolding instantiates a new CaptableHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableHoldingWithDefaults

`func NewCaptableHoldingWithDefaults() *CaptableHolding`

NewCaptableHoldingWithDefaults instantiates a new CaptableHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyDiluted

`func (o *CaptableHolding) GetFullyDiluted() int64`

GetFullyDiluted returns the FullyDiluted field if non-nil, zero value otherwise.

### GetFullyDilutedOk

`func (o *CaptableHolding) GetFullyDilutedOk() (*int64, bool)`

GetFullyDilutedOk returns a tuple with the FullyDiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDiluted

`func (o *CaptableHolding) SetFullyDiluted(v int64)`

SetFullyDiluted sets FullyDiluted field to given value.

### HasFullyDiluted

`func (o *CaptableHolding) HasFullyDiluted() bool`

HasFullyDiluted returns a boolean if a field has been set.

### GetName

`func (o *CaptableHolding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CaptableHolding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CaptableHolding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CaptableHolding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *CaptableHolding) GetOptions() int64`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CaptableHolding) GetOptionsOk() (*int64, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CaptableHolding) SetOptions(v int64)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CaptableHolding) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOwnershipPct

`func (o *CaptableHolding) GetOwnershipPct() float64`

GetOwnershipPct returns the OwnershipPct field if non-nil, zero value otherwise.

### GetOwnershipPctOk

`func (o *CaptableHolding) GetOwnershipPctOk() (*float64, bool)`

GetOwnershipPctOk returns a tuple with the OwnershipPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPct

`func (o *CaptableHolding) SetOwnershipPct(v float64)`

SetOwnershipPct sets OwnershipPct field to given value.

### HasOwnershipPct

`func (o *CaptableHolding) HasOwnershipPct() bool`

HasOwnershipPct returns a boolean if a field has been set.

### GetShares

`func (o *CaptableHolding) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CaptableHolding) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CaptableHolding) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CaptableHolding) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CaptableHolding) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CaptableHolding) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CaptableHolding) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CaptableHolding) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


