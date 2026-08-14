# Outcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Control** | Pointer to **bool** | true on the baseline arm; its own lift and stats are zero | [optional] 
**Converted** | Pointer to **int32** | of those, how many fired the metric event | [optional] 
**Exposed** | Pointer to **int32** | subjects the arm enrolled — the denominator | [optional] 
**Lift** | Pointer to **float32** | relative to control: (rate-ctrl)/ctrl | [optional] 
**PValue** | Pointer to **float32** | two-tailed p vs control | [optional] 
**Rate** | Pointer to **float32** | converted over exposed | [optional] 
**Significant** | Pointer to **bool** | pValue &lt; alpha | [optional] 
**Variant** | Pointer to **string** | the arm this row measures | [optional] 
**Z** | Pointer to **float32** | two-proportion z vs control | [optional] 

## Methods

### NewOutcome

`func NewOutcome() *Outcome`

NewOutcome instantiates a new Outcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcomeWithDefaults

`func NewOutcomeWithDefaults() *Outcome`

NewOutcomeWithDefaults instantiates a new Outcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControl

`func (o *Outcome) GetControl() bool`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *Outcome) GetControlOk() (*bool, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *Outcome) SetControl(v bool)`

SetControl sets Control field to given value.

### HasControl

`func (o *Outcome) HasControl() bool`

HasControl returns a boolean if a field has been set.

### GetConverted

`func (o *Outcome) GetConverted() int32`

GetConverted returns the Converted field if non-nil, zero value otherwise.

### GetConvertedOk

`func (o *Outcome) GetConvertedOk() (*int32, bool)`

GetConvertedOk returns a tuple with the Converted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConverted

`func (o *Outcome) SetConverted(v int32)`

SetConverted sets Converted field to given value.

### HasConverted

`func (o *Outcome) HasConverted() bool`

HasConverted returns a boolean if a field has been set.

### GetExposed

`func (o *Outcome) GetExposed() int32`

GetExposed returns the Exposed field if non-nil, zero value otherwise.

### GetExposedOk

`func (o *Outcome) GetExposedOk() (*int32, bool)`

GetExposedOk returns a tuple with the Exposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposed

`func (o *Outcome) SetExposed(v int32)`

SetExposed sets Exposed field to given value.

### HasExposed

`func (o *Outcome) HasExposed() bool`

HasExposed returns a boolean if a field has been set.

### GetLift

`func (o *Outcome) GetLift() float32`

GetLift returns the Lift field if non-nil, zero value otherwise.

### GetLiftOk

`func (o *Outcome) GetLiftOk() (*float32, bool)`

GetLiftOk returns a tuple with the Lift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLift

`func (o *Outcome) SetLift(v float32)`

SetLift sets Lift field to given value.

### HasLift

`func (o *Outcome) HasLift() bool`

HasLift returns a boolean if a field has been set.

### GetPValue

`func (o *Outcome) GetPValue() float32`

GetPValue returns the PValue field if non-nil, zero value otherwise.

### GetPValueOk

`func (o *Outcome) GetPValueOk() (*float32, bool)`

GetPValueOk returns a tuple with the PValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPValue

`func (o *Outcome) SetPValue(v float32)`

SetPValue sets PValue field to given value.

### HasPValue

`func (o *Outcome) HasPValue() bool`

HasPValue returns a boolean if a field has been set.

### GetRate

`func (o *Outcome) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *Outcome) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *Outcome) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *Outcome) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSignificant

`func (o *Outcome) GetSignificant() bool`

GetSignificant returns the Significant field if non-nil, zero value otherwise.

### GetSignificantOk

`func (o *Outcome) GetSignificantOk() (*bool, bool)`

GetSignificantOk returns a tuple with the Significant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificant

`func (o *Outcome) SetSignificant(v bool)`

SetSignificant sets Significant field to given value.

### HasSignificant

`func (o *Outcome) HasSignificant() bool`

HasSignificant returns a boolean if a field has been set.

### GetVariant

`func (o *Outcome) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Outcome) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Outcome) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Outcome) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetZ

`func (o *Outcome) GetZ() float32`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *Outcome) GetZOk() (*float32, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *Outcome) SetZ(v float32)`

SetZ sets Z field to given value.

### HasZ

`func (o *Outcome) HasZ() bool`

HasZ returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


