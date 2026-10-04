# BenchmarkPresetAccepted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is the verified user who composed it. | [optional] 
**Note** | Pointer to **string** | Note explains what acceptance does and does not promise. | [optional] 
**Preset** | Pointer to [**BenchmarkPreset**](BenchmarkPreset.md) | Preset is the blend with its defaults filled in and Owner set to the caller&#39;s verified org. | [optional] 
**ServedAs** | Pointer to **string** | ServedAs is the model id the serving layer would resolve this blend under. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;accepted\&quot;: the blend is well-formed, not that it is now served. | [optional] 

## Methods

### NewBenchmarkPresetAccepted

`func NewBenchmarkPresetAccepted() *BenchmarkPresetAccepted`

NewBenchmarkPresetAccepted instantiates a new BenchmarkPresetAccepted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPresetAcceptedWithDefaults

`func NewBenchmarkPresetAcceptedWithDefaults() *BenchmarkPresetAccepted`

NewBenchmarkPresetAcceptedWithDefaults instantiates a new BenchmarkPresetAccepted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *BenchmarkPresetAccepted) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *BenchmarkPresetAccepted) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *BenchmarkPresetAccepted) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *BenchmarkPresetAccepted) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetNote

`func (o *BenchmarkPresetAccepted) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BenchmarkPresetAccepted) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BenchmarkPresetAccepted) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BenchmarkPresetAccepted) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPreset

`func (o *BenchmarkPresetAccepted) GetPreset() BenchmarkPreset`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *BenchmarkPresetAccepted) GetPresetOk() (*BenchmarkPreset, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *BenchmarkPresetAccepted) SetPreset(v BenchmarkPreset)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *BenchmarkPresetAccepted) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetServedAs

`func (o *BenchmarkPresetAccepted) GetServedAs() string`

GetServedAs returns the ServedAs field if non-nil, zero value otherwise.

### GetServedAsOk

`func (o *BenchmarkPresetAccepted) GetServedAsOk() (*string, bool)`

GetServedAsOk returns a tuple with the ServedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAs

`func (o *BenchmarkPresetAccepted) SetServedAs(v string)`

SetServedAs sets ServedAs field to given value.

### HasServedAs

`func (o *BenchmarkPresetAccepted) HasServedAs() bool`

HasServedAs returns a boolean if a field has been set.

### GetStatus

`func (o *BenchmarkPresetAccepted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BenchmarkPresetAccepted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BenchmarkPresetAccepted) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BenchmarkPresetAccepted) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


