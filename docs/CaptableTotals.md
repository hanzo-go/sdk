# CaptableTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDilutedShares** | Pointer to **int64** | FullyDilutedShares is outstandingShares plus grantedOptions. | [optional] 
**GrantedOptions** | Pointer to **int64** | GrantedOptions is the shares under non-terminal option grants — grants that are EXERCISED, EXPIRED or CANCELLED are excluded, so nothing double-counts. | [optional] 
**OutstandingShares** | Pointer to **int64** | OutstandingShares is the sum of every issued share certificate. | [optional] 
**ShareClasses** | Pointer to **int64** | ShareClasses is how many share classes the company has authorized. | [optional] 
**Stakeholders** | Pointer to **int64** | Stakeholders is how many stakeholders the company has. | [optional] 

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

`func (o *CaptableTotals) GetFullyDilutedShares() int64`

GetFullyDilutedShares returns the FullyDilutedShares field if non-nil, zero value otherwise.

### GetFullyDilutedSharesOk

`func (o *CaptableTotals) GetFullyDilutedSharesOk() (*int64, bool)`

GetFullyDilutedSharesOk returns a tuple with the FullyDilutedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDilutedShares

`func (o *CaptableTotals) SetFullyDilutedShares(v int64)`

SetFullyDilutedShares sets FullyDilutedShares field to given value.

### HasFullyDilutedShares

`func (o *CaptableTotals) HasFullyDilutedShares() bool`

HasFullyDilutedShares returns a boolean if a field has been set.

### GetGrantedOptions

`func (o *CaptableTotals) GetGrantedOptions() int64`

GetGrantedOptions returns the GrantedOptions field if non-nil, zero value otherwise.

### GetGrantedOptionsOk

`func (o *CaptableTotals) GetGrantedOptionsOk() (*int64, bool)`

GetGrantedOptionsOk returns a tuple with the GrantedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedOptions

`func (o *CaptableTotals) SetGrantedOptions(v int64)`

SetGrantedOptions sets GrantedOptions field to given value.

### HasGrantedOptions

`func (o *CaptableTotals) HasGrantedOptions() bool`

HasGrantedOptions returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *CaptableTotals) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *CaptableTotals) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *CaptableTotals) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *CaptableTotals) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetShareClasses

`func (o *CaptableTotals) GetShareClasses() int64`

GetShareClasses returns the ShareClasses field if non-nil, zero value otherwise.

### GetShareClassesOk

`func (o *CaptableTotals) GetShareClassesOk() (*int64, bool)`

GetShareClassesOk returns a tuple with the ShareClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClasses

`func (o *CaptableTotals) SetShareClasses(v int64)`

SetShareClasses sets ShareClasses field to given value.

### HasShareClasses

`func (o *CaptableTotals) HasShareClasses() bool`

HasShareClasses returns a boolean if a field has been set.

### GetStakeholders

`func (o *CaptableTotals) GetStakeholders() int64`

GetStakeholders returns the Stakeholders field if non-nil, zero value otherwise.

### GetStakeholdersOk

`func (o *CaptableTotals) GetStakeholdersOk() (*int64, bool)`

GetStakeholdersOk returns a tuple with the Stakeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholders

`func (o *CaptableTotals) SetStakeholders(v int64)`

SetStakeholders sets Stakeholders field to given value.

### HasStakeholders

`func (o *CaptableTotals) HasStakeholders() bool`

HasStakeholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


