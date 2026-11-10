# BenchmarkPresetStored

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **string** | Note says what storing does and does not promise. | [optional] 
**Preset** | Pointer to [**BenchmarkPreset**](BenchmarkPreset.md) | Preset is the blend as stored: defaults filled in, Owner the caller&#39;s verified org, At and By the server&#39;s. | [optional] 
**ServedAs** | Pointer to **string** | ServedAs is enso-&lt;name&gt;: the model id a router would resolve this blend under. | [optional] 

## Methods

### NewBenchmarkPresetStored

`func NewBenchmarkPresetStored() *BenchmarkPresetStored`

NewBenchmarkPresetStored instantiates a new BenchmarkPresetStored object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPresetStoredWithDefaults

`func NewBenchmarkPresetStoredWithDefaults() *BenchmarkPresetStored`

NewBenchmarkPresetStoredWithDefaults instantiates a new BenchmarkPresetStored object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *BenchmarkPresetStored) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BenchmarkPresetStored) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BenchmarkPresetStored) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BenchmarkPresetStored) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPreset

`func (o *BenchmarkPresetStored) GetPreset() BenchmarkPreset`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *BenchmarkPresetStored) GetPresetOk() (*BenchmarkPreset, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *BenchmarkPresetStored) SetPreset(v BenchmarkPreset)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *BenchmarkPresetStored) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetServedAs

`func (o *BenchmarkPresetStored) GetServedAs() string`

GetServedAs returns the ServedAs field if non-nil, zero value otherwise.

### GetServedAsOk

`func (o *BenchmarkPresetStored) GetServedAsOk() (*string, bool)`

GetServedAsOk returns a tuple with the ServedAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedAs

`func (o *BenchmarkPresetStored) SetServedAs(v string)`

SetServedAs sets ServedAs field to given value.

### HasServedAs

`func (o *BenchmarkPresetStored) HasServedAs() bool`

HasServedAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


