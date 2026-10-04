# PricingSpeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is what a request asks for. | [optional] 
**Note** | Pointer to **string** | Note is what asking for this window buys. | [optional] 
**Rank** | Pointer to **int64** | Rank orders the windows by price, 1 being the dearest. | [optional] 

## Methods

### NewPricingSpeed

`func NewPricingSpeed() *PricingSpeed`

NewPricingSpeed instantiates a new PricingSpeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingSpeedWithDefaults

`func NewPricingSpeedWithDefaults() *PricingSpeed`

NewPricingSpeedWithDefaults instantiates a new PricingSpeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PricingSpeed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingSpeed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingSpeed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingSpeed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *PricingSpeed) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PricingSpeed) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PricingSpeed) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PricingSpeed) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRank

`func (o *PricingSpeed) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *PricingSpeed) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *PricingSpeed) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *PricingSpeed) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


