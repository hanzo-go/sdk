# PresetAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** | Note explains what acceptance does and does not promise. | [optional] 
**Preset** | Pointer to [**Preset**](Preset.md) | Preset is the blend with its defaults filled in. | [optional] 
**ServedAs** | Pointer to **string** | ServedAs is the model id the serving layer would resolve this blend under. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;accepted\&quot;: the blend is well-formed, not that it is now served. | [optional] 

## Methods

### NewPresetAccepted

`func NewPresetAccepted() *PresetAccepted`

NewPresetAccepted instantiates a new PresetAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetAcceptedWithDefaults

`func NewPresetAcceptedWithDefaults() *PresetAccepted`

NewPresetAcceptedWithDefaults instantiates a new PresetAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *PresetAccepted) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PresetAccepted) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PresetAccepted) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PresetAccepted) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPreset

`func (o *PresetAccepted) GetPreset() Preset`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *PresetAccepted) GetPresetOk() (*Preset, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *PresetAccepted) SetPreset(v Preset)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *PresetAccepted) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetServedAs

`func (o *PresetAccepted) GetServedAs() string`

GetServedAs returns the ServedAs field if non-nil, zero value otherwise.

### GetServedAsOk

`func (o *PresetAccepted) GetServedAsOk() (*string, bool)`

GetServedAsOk returns a tuple with the ServedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAs

`func (o *PresetAccepted) SetServedAs(v string)`

SetServedAs sets ServedAs field to given value.

### HasServedAs

`func (o *PresetAccepted) HasServedAs() bool`

HasServedAs returns a boolean if a field has been set.

### GetStatus

`func (o *PresetAccepted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PresetAccepted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PresetAccepted) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PresetAccepted) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


