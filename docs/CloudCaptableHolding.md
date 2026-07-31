# CloudCaptableHolding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDiluted** | Pointer to **int32** | FullyDiluted is shares plus options. | [optional] 
**Name** | Pointer to **string** | Name is the stakeholder&#39;s name. | [optional] 
**Options** | Pointer to **int32** | Options is the shares under this stakeholder&#39;s non-terminal option grants. | [optional] 
**OwnershipPct** | Pointer to **float32** | OwnershipPct is fullyDiluted as a percentage of the company&#39;s fullyDilutedShares, rounded to two decimals; 0 when nothing is issued. | [optional] 
**Shares** | Pointer to **int32** | Shares is the shares this stakeholder holds by certificate. | [optional] 
**StakeholderId** | Pointer to **string** | StakeholderID is the stakeholder. | [optional] 

## Methods

### NewCloudCaptableHolding

`func NewCloudCaptableHolding() *CloudCaptableHolding`

NewCloudCaptableHolding instantiates a new CloudCaptableHolding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableHoldingWithDefaults

`func NewCloudCaptableHoldingWithDefaults() *CloudCaptableHolding`

NewCloudCaptableHoldingWithDefaults instantiates a new CloudCaptableHolding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyDiluted

`func (o *CloudCaptableHolding) GetFullyDiluted() int32`

GetFullyDiluted returns the FullyDiluted field if non-nil, zero value otherwise.

### GetFullyDilutedOk

`func (o *CloudCaptableHolding) GetFullyDilutedOk() (*int32, bool)`

GetFullyDilutedOk returns a tuple with the FullyDiluted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDiluted

`func (o *CloudCaptableHolding) SetFullyDiluted(v int32)`

SetFullyDiluted sets FullyDiluted field to given value.

### HasFullyDiluted

`func (o *CloudCaptableHolding) HasFullyDiluted() bool`

HasFullyDiluted returns a boolean if a field has been set.

### GetName

`func (o *CloudCaptableHolding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCaptableHolding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCaptableHolding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCaptableHolding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *CloudCaptableHolding) GetOptions() int32`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CloudCaptableHolding) GetOptionsOk() (*int32, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CloudCaptableHolding) SetOptions(v int32)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CloudCaptableHolding) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetOwnershipPct

`func (o *CloudCaptableHolding) GetOwnershipPct() float32`

GetOwnershipPct returns the OwnershipPct field if non-nil, zero value otherwise.

### GetOwnershipPctOk

`func (o *CloudCaptableHolding) GetOwnershipPctOk() (*float32, bool)`

GetOwnershipPctOk returns a tuple with the OwnershipPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipPct

`func (o *CloudCaptableHolding) SetOwnershipPct(v float32)`

SetOwnershipPct sets OwnershipPct field to given value.

### HasOwnershipPct

`func (o *CloudCaptableHolding) HasOwnershipPct() bool`

HasOwnershipPct returns a boolean if a field has been set.

### GetShares

`func (o *CloudCaptableHolding) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *CloudCaptableHolding) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *CloudCaptableHolding) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *CloudCaptableHolding) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetStakeholderId

`func (o *CloudCaptableHolding) GetStakeholderId() string`

GetStakeholderId returns the StakeholderId field if non-nil, zero value otherwise.

### GetStakeholderIdOk

`func (o *CloudCaptableHolding) GetStakeholderIdOk() (*string, bool)`

GetStakeholderIdOk returns a tuple with the StakeholderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholderId

`func (o *CloudCaptableHolding) SetStakeholderId(v string)`

SetStakeholderId sets StakeholderId field to given value.

### HasStakeholderId

`func (o *CloudCaptableHolding) HasStakeholderId() bool`

HasStakeholderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


