# CloudCaptableTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullyDilutedShares** | Pointer to **int32** | FullyDilutedShares is outstandingShares plus grantedOptions. | [optional] 
**GrantedOptions** | Pointer to **int32** | GrantedOptions is the shares under non-terminal option grants — grants that are EXERCISED, EXPIRED or CANCELLED are excluded, so nothing double-counts. | [optional] 
**OutstandingShares** | Pointer to **int32** | OutstandingShares is the sum of every issued share certificate. | [optional] 
**ShareClasses** | Pointer to **int32** | ShareClasses is how many share classes the company has authorized. | [optional] 
**Stakeholders** | Pointer to **int32** | Stakeholders is how many stakeholders the company has. | [optional] 

## Methods

### NewCloudCaptableTotals

`func NewCloudCaptableTotals() *CloudCaptableTotals`

NewCloudCaptableTotals instantiates a new CloudCaptableTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCaptableTotalsWithDefaults

`func NewCloudCaptableTotalsWithDefaults() *CloudCaptableTotals`

NewCloudCaptableTotalsWithDefaults instantiates a new CloudCaptableTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullyDilutedShares

`func (o *CloudCaptableTotals) GetFullyDilutedShares() int32`

GetFullyDilutedShares returns the FullyDilutedShares field if non-nil, zero value otherwise.

### GetFullyDilutedSharesOk

`func (o *CloudCaptableTotals) GetFullyDilutedSharesOk() (*int32, bool)`

GetFullyDilutedSharesOk returns a tuple with the FullyDilutedShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyDilutedShares

`func (o *CloudCaptableTotals) SetFullyDilutedShares(v int32)`

SetFullyDilutedShares sets FullyDilutedShares field to given value.

### HasFullyDilutedShares

`func (o *CloudCaptableTotals) HasFullyDilutedShares() bool`

HasFullyDilutedShares returns a boolean if a field has been set.

### GetGrantedOptions

`func (o *CloudCaptableTotals) GetGrantedOptions() int32`

GetGrantedOptions returns the GrantedOptions field if non-nil, zero value otherwise.

### GetGrantedOptionsOk

`func (o *CloudCaptableTotals) GetGrantedOptionsOk() (*int32, bool)`

GetGrantedOptionsOk returns a tuple with the GrantedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantedOptions

`func (o *CloudCaptableTotals) SetGrantedOptions(v int32)`

SetGrantedOptions sets GrantedOptions field to given value.

### HasGrantedOptions

`func (o *CloudCaptableTotals) HasGrantedOptions() bool`

HasGrantedOptions returns a boolean if a field has been set.

### GetOutstandingShares

`func (o *CloudCaptableTotals) GetOutstandingShares() int32`

GetOutstandingShares returns the OutstandingShares field if non-nil, zero value otherwise.

### GetOutstandingSharesOk

`func (o *CloudCaptableTotals) GetOutstandingSharesOk() (*int32, bool)`

GetOutstandingSharesOk returns a tuple with the OutstandingShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingShares

`func (o *CloudCaptableTotals) SetOutstandingShares(v int32)`

SetOutstandingShares sets OutstandingShares field to given value.

### HasOutstandingShares

`func (o *CloudCaptableTotals) HasOutstandingShares() bool`

HasOutstandingShares returns a boolean if a field has been set.

### GetShareClasses

`func (o *CloudCaptableTotals) GetShareClasses() int32`

GetShareClasses returns the ShareClasses field if non-nil, zero value otherwise.

### GetShareClassesOk

`func (o *CloudCaptableTotals) GetShareClassesOk() (*int32, bool)`

GetShareClassesOk returns a tuple with the ShareClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareClasses

`func (o *CloudCaptableTotals) SetShareClasses(v int32)`

SetShareClasses sets ShareClasses field to given value.

### HasShareClasses

`func (o *CloudCaptableTotals) HasShareClasses() bool`

HasShareClasses returns a boolean if a field has been set.

### GetStakeholders

`func (o *CloudCaptableTotals) GetStakeholders() int32`

GetStakeholders returns the Stakeholders field if non-nil, zero value otherwise.

### GetStakeholdersOk

`func (o *CloudCaptableTotals) GetStakeholdersOk() (*int32, bool)`

GetStakeholdersOk returns a tuple with the Stakeholders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeholders

`func (o *CloudCaptableTotals) SetStakeholders(v int32)`

SetStakeholders sets Stakeholders field to given value.

### HasStakeholders

`func (o *CloudCaptableTotals) HasStakeholders() bool`

HasStakeholders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


