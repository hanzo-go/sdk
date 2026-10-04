# PricingCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]PricingComponent**](PricingComponent.md) | Components are the four things a bill is made of. | [optional] 
**Rates** | Pointer to [**[]PricingRate**](PricingRate.md) | Rates are the priced lines, each in micro-USD per its own unit. | [optional] 
**Unit** | Pointer to **string** | Unit is the unit every amount on this card is stated in. | [optional] 
**Windows** | Pointer to [**[]PricingSpeed**](PricingSpeed.md) | Speeds are the completion speeds, dearest first. | [optional] 

## Methods

### NewPricingCard

`func NewPricingCard() *PricingCard`

NewPricingCard instantiates a new PricingCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingCardWithDefaults

`func NewPricingCardWithDefaults() *PricingCard`

NewPricingCardWithDefaults instantiates a new PricingCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *PricingCard) GetComponents() []PricingComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PricingCard) GetComponentsOk() (*[]PricingComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PricingCard) SetComponents(v []PricingComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PricingCard) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetRates

`func (o *PricingCard) GetRates() []PricingRate`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *PricingCard) GetRatesOk() (*[]PricingRate, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *PricingCard) SetRates(v []PricingRate)`

SetRates sets Rates field to given value.

### HasRates

`func (o *PricingCard) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetUnit

`func (o *PricingCard) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PricingCard) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PricingCard) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PricingCard) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetWindows

`func (o *PricingCard) GetWindows() []PricingSpeed`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *PricingCard) GetWindowsOk() (*[]PricingSpeed, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *PricingCard) SetWindows(v []PricingSpeed)`

SetWindows sets Windows field to given value.

### HasWindows

`func (o *PricingCard) HasWindows() bool`

HasWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


