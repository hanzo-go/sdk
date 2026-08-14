# Appearance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accent** | Pointer to **string** | Accent is the one hue — a CSS colour token (a hex, or a bounded functional colour like rgb()/oklch()). Anything else is dropped rather than stored. | [optional] 
**Density** | Pointer to **string** | Density is the spacing step: \&quot;compact\&quot;, \&quot;default\&quot; or \&quot;comfortable\&quot;. | [optional] 
**Type** | Pointer to **float32** | Type is the text-size multiplier, clamped to the ramp window [0.85, 1.4]. Absent (0) leaves the published default. | [optional] 

## Methods

### NewAppearance

`func NewAppearance() *Appearance`

NewAppearance instantiates a new Appearance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppearanceWithDefaults

`func NewAppearanceWithDefaults() *Appearance`

NewAppearanceWithDefaults instantiates a new Appearance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccent

`func (o *Appearance) GetAccent() string`

GetAccent returns the Accent field if non-nil, zero value otherwise.

### GetAccentOk

`func (o *Appearance) GetAccentOk() (*string, bool)`

GetAccentOk returns a tuple with the Accent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccent

`func (o *Appearance) SetAccent(v string)`

SetAccent sets Accent field to given value.

### HasAccent

`func (o *Appearance) HasAccent() bool`

HasAccent returns a boolean if a field has been set.

### GetDensity

`func (o *Appearance) GetDensity() string`

GetDensity returns the Density field if non-nil, zero value otherwise.

### GetDensityOk

`func (o *Appearance) GetDensityOk() (*string, bool)`

GetDensityOk returns a tuple with the Density field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDensity

`func (o *Appearance) SetDensity(v string)`

SetDensity sets Density field to given value.

### HasDensity

`func (o *Appearance) HasDensity() bool`

HasDensity returns a boolean if a field has been set.

### GetType

`func (o *Appearance) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Appearance) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Appearance) SetType(v float32)`

SetType sets Type field to given value.

### HasType

`func (o *Appearance) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


