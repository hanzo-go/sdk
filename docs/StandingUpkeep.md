# StandingUpkeep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtLeast** | Pointer to **bool** | AtLeast reports that some obligation is a minimum, so YearlyCents is a floor rather than a final figure. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code every amount here is denominated in. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state whose obligations these are. | [optional] 
**Obligations** | Pointer to [**[]StandingObligation**](StandingObligation.md) | Obligations are the recurring charges, in the order a reader should see them. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity this prices. | [optional] 
**YearlyCents** | Pointer to **int64** | YearlyCents is what the entity owes every year, all obligations summed. | [optional] 

## Methods

### NewStandingUpkeep

`func NewStandingUpkeep() *StandingUpkeep`

NewStandingUpkeep instantiates a new StandingUpkeep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandingUpkeepWithDefaults

`func NewStandingUpkeepWithDefaults() *StandingUpkeep`

NewStandingUpkeepWithDefaults instantiates a new StandingUpkeep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtLeast

`func (o *StandingUpkeep) GetAtLeast() bool`

GetAtLeast returns the AtLeast field if non-nil, zero value otherwise.

### GetAtLeastOk

`func (o *StandingUpkeep) GetAtLeastOk() (*bool, bool)`

GetAtLeastOk returns a tuple with the AtLeast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtLeast

`func (o *StandingUpkeep) SetAtLeast(v bool)`

SetAtLeast sets AtLeast field to given value.

### HasAtLeast

`func (o *StandingUpkeep) HasAtLeast() bool`

HasAtLeast returns a boolean if a field has been set.

### GetCurrency

`func (o *StandingUpkeep) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StandingUpkeep) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StandingUpkeep) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *StandingUpkeep) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetJurisdiction

`func (o *StandingUpkeep) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *StandingUpkeep) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *StandingUpkeep) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *StandingUpkeep) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetObligations

`func (o *StandingUpkeep) GetObligations() []StandingObligation`

GetObligations returns the Obligations field if non-nil, zero value otherwise.

### GetObligationsOk

`func (o *StandingUpkeep) GetObligationsOk() (*[]StandingObligation, bool)`

GetObligationsOk returns a tuple with the Obligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligations

`func (o *StandingUpkeep) SetObligations(v []StandingObligation)`

SetObligations sets Obligations field to given value.

### HasObligations

`func (o *StandingUpkeep) HasObligations() bool`

HasObligations returns a boolean if a field has been set.

### GetStructure

`func (o *StandingUpkeep) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *StandingUpkeep) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *StandingUpkeep) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *StandingUpkeep) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetYearlyCents

`func (o *StandingUpkeep) GetYearlyCents() int64`

GetYearlyCents returns the YearlyCents field if non-nil, zero value otherwise.

### GetYearlyCentsOk

`func (o *StandingUpkeep) GetYearlyCentsOk() (*int64, bool)`

GetYearlyCentsOk returns a tuple with the YearlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCents

`func (o *StandingUpkeep) SetYearlyCents(v int64)`

SetYearlyCents sets YearlyCents field to given value.

### HasYearlyCents

`func (o *StandingUpkeep) HasYearlyCents() bool`

HasYearlyCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


