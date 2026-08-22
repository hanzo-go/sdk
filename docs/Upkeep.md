# Upkeep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AtLeast** | Pointer to **bool** | AtLeast reports that some obligation is a minimum, so YearlyCents is a floor rather than a final figure. | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO code every amount here is denominated in. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state whose obligations these are. | [optional] 
**Obligations** | Pointer to [**[]Obligation**](Obligation.md) | Obligations are the recurring charges, in the order a reader should see them. | [optional] 
**Structure** | Pointer to **string** | Structure is the entity this prices. | [optional] 
**YearlyCents** | Pointer to **int32** | YearlyCents is what the entity owes every year, all obligations summed. | [optional] 

## Methods

### NewUpkeep

`func NewUpkeep() *Upkeep`

NewUpkeep instantiates a new Upkeep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpkeepWithDefaults

`func NewUpkeepWithDefaults() *Upkeep`

NewUpkeepWithDefaults instantiates a new Upkeep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAtLeast

`func (o *Upkeep) GetAtLeast() bool`

GetAtLeast returns the AtLeast field if non-nil, zero value otherwise.

### GetAtLeastOk

`func (o *Upkeep) GetAtLeastOk() (*bool, bool)`

GetAtLeastOk returns a tuple with the AtLeast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtLeast

`func (o *Upkeep) SetAtLeast(v bool)`

SetAtLeast sets AtLeast field to given value.

### HasAtLeast

`func (o *Upkeep) HasAtLeast() bool`

HasAtLeast returns a boolean if a field has been set.

### GetCurrency

`func (o *Upkeep) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Upkeep) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Upkeep) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Upkeep) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetJurisdiction

`func (o *Upkeep) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *Upkeep) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *Upkeep) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *Upkeep) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetObligations

`func (o *Upkeep) GetObligations() []Obligation`

GetObligations returns the Obligations field if non-nil, zero value otherwise.

### GetObligationsOk

`func (o *Upkeep) GetObligationsOk() (*[]Obligation, bool)`

GetObligationsOk returns a tuple with the Obligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligations

`func (o *Upkeep) SetObligations(v []Obligation)`

SetObligations sets Obligations field to given value.

### HasObligations

`func (o *Upkeep) HasObligations() bool`

HasObligations returns a boolean if a field has been set.

### GetStructure

`func (o *Upkeep) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *Upkeep) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *Upkeep) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *Upkeep) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetYearlyCents

`func (o *Upkeep) GetYearlyCents() int32`

GetYearlyCents returns the YearlyCents field if non-nil, zero value otherwise.

### GetYearlyCentsOk

`func (o *Upkeep) GetYearlyCentsOk() (*int32, bool)`

GetYearlyCentsOk returns a tuple with the YearlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyCents

`func (o *Upkeep) SetYearlyCents(v int32)`

SetYearlyCents sets YearlyCents field to given value.

### HasYearlyCents

`func (o *Upkeep) HasYearlyCents() bool`

HasYearlyCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


