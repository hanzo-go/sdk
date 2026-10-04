# CaptableCaptableTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDilutedShares** | Pointer to **int64** | FullyDilutedShares is outstandingShares plus grantedOptions. | [optional] 
**GrantedOptions** | Pointer to **int64** | GrantedOptions is the shares under non-terminal option grants — grants that are EXERCISED, EXPIRED or CANCELLED are excluded, so nothing double-counts. | [optional] 
**OutstandingShares** | Pointer to **int64** | OutstandingShares is the sum of every issued share certificate. | [optional] 
**ShareClasses** | Pointer to **int64** | ShareClasses is how many share classes the company has authorized. | [optional] 
**Stakeholders** | Pointer to **int64** | Stakeholders is how many stakeholders the company has. | [optional] 

## Methods

### NewCaptableCaptableTotals

`func NewCaptableCaptableTotals() *CaptableCaptableTotals`

NewCaptableCaptableTotals instantiates a new CaptableCaptableTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptableCaptableTotalsWithDefaults

`func NewCaptableCaptableTotalsWithDefaults() *CaptableCaptableTotals`

NewCaptableCaptableTotalsWithDefaults instantiates a new CaptableCaptableTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyDilutedShares

`func (o *CaptableCaptableTotals) GetFullyDilutedShares() int64`

GetFullyDilutedShares returns the FullyDilutedShares field if non-nil, zero value otherwise.

### GetFullyDilutedSharesOk

`func (o *CaptableCaptableTotals) GetFullyDilutedSharesOk() (*int64, bool)`

GetFullyDilutedSharesOk returns a tuple with the FullyDilutedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDilutedShares

`func (o *CaptableCaptableTotals) SetFullyDilutedShares(v int64)`

SetFullyDilutedShares sets FullyDilutedShares field to given value.

### HasFullyDilutedShares

`func (o *CaptableCaptableTotals) HasFullyDilutedShares() bool`

HasFullyDilutedShares returns a boolean if a field has been set.

### GetGrantedOptions

`func (o *CaptableCaptableTotals) GetGrantedOptions() int64`

GetGrantedOptions returns the GrantedOptions field if non-nil, zero value otherwise.

### GetGrantedOptionsOk

`func (o *CaptableCaptableTotals) GetGrantedOptionsOk() (*int64, bool)`

GetGrantedOptionsOk returns a tuple with the GrantedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedOptions

`func (o *CaptableCaptableTotals) SetGrantedOptions(v int64)`

SetGrantedOptions sets GrantedOptions field to given value.

### HasGrantedOptions

`func (o *CaptableCaptableTotals) HasGrantedOptions() bool`

HasGrantedOptions returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *CaptableCaptableTotals) GetOutstandingShares() int64`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *CaptableCaptableTotals) GetOutstandingSharesOk() (*int64, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *CaptableCaptableTotals) SetOutstandingShares(v int64)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *CaptableCaptableTotals) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetShareClasses

`func (o *CaptableCaptableTotals) GetShareClasses() int64`

GetShareClasses returns the ShareClasses field if non-nil, zero value otherwise.

### GetShareClassesOk

`func (o *CaptableCaptableTotals) GetShareClassesOk() (*int64, bool)`

GetShareClassesOk returns a tuple with the ShareClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClasses

`func (o *CaptableCaptableTotals) SetShareClasses(v int64)`

SetShareClasses sets ShareClasses field to given value.

### HasShareClasses

`func (o *CaptableCaptableTotals) HasShareClasses() bool`

HasShareClasses returns a boolean if a field has been set.

### GetStakeholders

`func (o *CaptableCaptableTotals) GetStakeholders() int64`

GetStakeholders returns the Stakeholders field if non-nil, zero value otherwise.

### GetStakeholdersOk

`func (o *CaptableCaptableTotals) GetStakeholdersOk() (*int64, bool)`

GetStakeholdersOk returns a tuple with the Stakeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholders

`func (o *CaptableCaptableTotals) SetStakeholders(v int64)`

SetStakeholders sets Stakeholders field to given value.

### HasStakeholders

`func (o *CaptableCaptableTotals) HasStakeholders() bool`

HasStakeholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


