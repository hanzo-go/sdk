# BenchmarkPreset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arms** | Pointer to **[]string** | the blend — model ids from the arena | [optional] 
**At** | Pointer to **time.Time** | At is when this version was stored. Set by the server; a value sent is not read. Zero on the reference blend. | [optional] 
**By** | Pointer to **string** | By is the verified user who stored it. Set by the server; a value sent is not read. | [optional] 
**Name** | Pointer to **string** | the part after enso-: lowercase letters, digits and inner hyphens, at most 64 | [optional] 
**Note** | Pointer to **string** | why this blend (audit) | [optional] 
**Owner** | Pointer to **string** | the org the blend belongs to: the writer&#39;s verified org, never a value sent | [optional] 
**Panel** | Pointer to **int64** | fan-out width (&gt;&#x3D;1) | [optional] 
**Rank** | Pointer to **[]string** | escalation order over arms | [optional] 

## Methods

### NewBenchmarkPreset

`func NewBenchmarkPreset() *BenchmarkPreset`

NewBenchmarkPreset instantiates a new BenchmarkPreset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkPresetWithDefaults

`func NewBenchmarkPresetWithDefaults() *BenchmarkPreset`

NewBenchmarkPresetWithDefaults instantiates a new BenchmarkPreset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArms

`func (o *BenchmarkPreset) GetArms() []string`

GetArms returns the Arms field if non-nil, zero value otherwise.

### GetArmsOk

`func (o *BenchmarkPreset) GetArmsOk() (*[]string, bool)`

GetArmsOk returns a tuple with the Arms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArms

`func (o *BenchmarkPreset) SetArms(v []string)`

SetArms sets Arms field to given value.

### HasArms

`func (o *BenchmarkPreset) HasArms() bool`

HasArms returns a boolean if a field has been set.

### GetAt

`func (o *BenchmarkPreset) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *BenchmarkPreset) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *BenchmarkPreset) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *BenchmarkPreset) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *BenchmarkPreset) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *BenchmarkPreset) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *BenchmarkPreset) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *BenchmarkPreset) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetName

`func (o *BenchmarkPreset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BenchmarkPreset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BenchmarkPreset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BenchmarkPreset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *BenchmarkPreset) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BenchmarkPreset) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BenchmarkPreset) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BenchmarkPreset) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOwner

`func (o *BenchmarkPreset) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BenchmarkPreset) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BenchmarkPreset) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BenchmarkPreset) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPanel

`func (o *BenchmarkPreset) GetPanel() int64`

GetPanel returns the Panel field if non-nil, zero value otherwise.

### GetPanelOk

`func (o *BenchmarkPreset) GetPanelOk() (*int64, bool)`

GetPanelOk returns a tuple with the Panel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanel

`func (o *BenchmarkPreset) SetPanel(v int64)`

SetPanel sets Panel field to given value.

### HasPanel

`func (o *BenchmarkPreset) HasPanel() bool`

HasPanel returns a boolean if a field has been set.

### GetRank

`func (o *BenchmarkPreset) GetRank() []string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BenchmarkPreset) GetRankOk() (*[]string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BenchmarkPreset) SetRank(v []string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BenchmarkPreset) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


