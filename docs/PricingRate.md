# PricingRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to **string** | Component is which of the four this line bills under. | [optional] 
**Name** | Pointer to **string** | Name addresses the rate. | [optional] 
**Note** | Pointer to **string** | Note is what the line covers. | [optional] 
**Per** | Pointer to **string** | Per names one unit, so a reader never has to guess what the number is per. | [optional] 
**RateMicroUsd** | Pointer to **int64** | Micros is the price of ONE unit, in micro-USD. It is the number the ledger books, resolved through the same authority the metering path reads. | [optional] 
**Title** | Pointer to **string** | Title is how the line reads on a price list. | [optional] 

## Methods

### NewPricingRate

`func NewPricingRate() *PricingRate`

NewPricingRate instantiates a new PricingRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingRateWithDefaults

`func NewPricingRateWithDefaults() *PricingRate`

NewPricingRateWithDefaults instantiates a new PricingRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *PricingRate) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PricingRate) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PricingRate) SetComponent(v string)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *PricingRate) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *PricingRate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingRate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingRate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingRate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *PricingRate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PricingRate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PricingRate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PricingRate) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPer

`func (o *PricingRate) GetPer() string`

GetPer returns the Per field if non-nil, zero value otherwise.

### GetPerOk

`func (o *PricingRate) GetPerOk() (*string, bool)`

GetPerOk returns a tuple with the Per field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPer

`func (o *PricingRate) SetPer(v string)`

SetPer sets Per field to given value.

### HasPer

`func (o *PricingRate) HasPer() bool`

HasPer returns a boolean if a field has been set.

### GetRateMicroUsd

`func (o *PricingRate) GetRateMicroUsd() int64`

GetRateMicroUsd returns the RateMicroUsd field if non-nil, zero value otherwise.

### GetRateMicroUsdOk

`func (o *PricingRate) GetRateMicroUsdOk() (*int64, bool)`

GetRateMicroUsdOk returns a tuple with the RateMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMicroUsd

`func (o *PricingRate) SetRateMicroUsd(v int64)`

SetRateMicroUsd sets RateMicroUsd field to given value.

### HasRateMicroUsd

`func (o *PricingRate) HasRateMicroUsd() bool`

HasRateMicroUsd returns a boolean if a field has been set.

### GetTitle

`func (o *PricingRate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PricingRate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PricingRate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PricingRate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


