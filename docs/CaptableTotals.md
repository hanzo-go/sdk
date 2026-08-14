# CaptableTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDilutedShares** | Pointer to **int32** | FullyDilutedShares is outstandingShares plus grantedOptions. | [optional] 
**GrantedOptions** | Pointer to **int32** | GrantedOptions is the shares under non-terminal option grants — grants that are EXERCISED, EXPIRED or CANCELLED are excluded, so nothing double-counts. | [optional] 
**OutstandingShares** | Pointer to **int32** | OutstandingShares is the sum of every issued share certificate. | [optional] 
**ShareClasses** | Pointer to **int32** | ShareClasses is how many share classes the company has authorized. | [optional] 
**Stakeholders** | Pointer to **int32** | Stakeholders is how many stakeholders the company has. | [optional] 

## Methods

### NewCaptableTotals

`func NewCaptableTotals() *CaptableTotals`

NewCaptableTotals instantiates a new CaptableTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableTotalsWithDefaults

`func NewCaptableTotalsWithDefaults() *CaptableTotals`

NewCaptableTotalsWithDefaults instantiates a new CaptableTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyDilutedShares

`func (o *CaptableTotals) GetFullyDilutedShares() int32`

GetFullyDilutedShares returns the FullyDilutedShares field if non-nil, zero value otherwise.

### GetFullyDilutedSharesOk

`func (o *CaptableTotals) GetFullyDilutedSharesOk() (*int32, bool)`

GetFullyDilutedSharesOk returns a tuple with the FullyDilutedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDilutedShares

`func (o *CaptableTotals) SetFullyDilutedShares(v int32)`

SetFullyDilutedShares sets FullyDilutedShares field to given value.

### HasFullyDilutedShares

`func (o *CaptableTotals) HasFullyDilutedShares() bool`

HasFullyDilutedShares returns a boolean if a field has been set.

### GetGrantedOptions

`func (o *CaptableTotals) GetGrantedOptions() int32`

GetGrantedOptions returns the GrantedOptions field if non-nil, zero value otherwise.

### GetGrantedOptionsOk

`func (o *CaptableTotals) GetGrantedOptionsOk() (*int32, bool)`

GetGrantedOptionsOk returns a tuple with the GrantedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedOptions

`func (o *CaptableTotals) SetGrantedOptions(v int32)`

SetGrantedOptions sets GrantedOptions field to given value.

### HasGrantedOptions

`func (o *CaptableTotals) HasGrantedOptions() bool`

HasGrantedOptions returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *CaptableTotals) GetOutstandingShares() int32`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *CaptableTotals) GetOutstandingSharesOk() (*int32, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *CaptableTotals) SetOutstandingShares(v int32)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *CaptableTotals) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetShareClasses

`func (o *CaptableTotals) GetShareClasses() int32`

GetShareClasses returns the ShareClasses field if non-nil, zero value otherwise.

### GetShareClassesOk

`func (o *CaptableTotals) GetShareClassesOk() (*int32, bool)`

GetShareClassesOk returns a tuple with the ShareClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClasses

`func (o *CaptableTotals) SetShareClasses(v int32)`

SetShareClasses sets ShareClasses field to given value.

### HasShareClasses

`func (o *CaptableTotals) HasShareClasses() bool`

HasShareClasses returns a boolean if a field has been set.

### GetStakeholders

`func (o *CaptableTotals) GetStakeholders() int32`

GetStakeholders returns the Stakeholders field if non-nil, zero value otherwise.

### GetStakeholdersOk

`func (o *CaptableTotals) GetStakeholdersOk() (*int32, bool)`

GetStakeholdersOk returns a tuple with the Stakeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholders

`func (o *CaptableTotals) SetStakeholders(v int32)`

SetStakeholders sets Stakeholders field to given value.

### HasStakeholders

`func (o *CaptableTotals) HasStakeholders() bool`

HasStakeholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


