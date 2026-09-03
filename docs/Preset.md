# Preset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arms** | Pointer to **[]string** | the blend — model ids from the arena | [optional] 
**Name** | Pointer to **string** | served as enso-&lt;name&gt; | [optional] 
**Note** | Pointer to **string** | why this blend (audit) | [optional] 
**Owner** | Pointer to **string** | scoping org (never cross-tenant) | [optional] 
**Panel** | Pointer to **int64** | fan-out width (&gt;&#x3D;1) | [optional] 
**Rank** | Pointer to **[]string** | escalation order over arms | [optional] 

## Methods

### NewPreset

`func NewPreset() *Preset`

NewPreset instantiates a new Preset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresetWithDefaults

`func NewPresetWithDefaults() *Preset`

NewPresetWithDefaults instantiates a new Preset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArms

`func (o *Preset) GetArms() []string`

GetArms returns the Arms field if non-nil, zero value otherwise.

### GetArmsOk

`func (o *Preset) GetArmsOk() (*[]string, bool)`

GetArmsOk returns a tuple with the Arms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArms

`func (o *Preset) SetArms(v []string)`

SetArms sets Arms field to given value.

### HasArms

`func (o *Preset) HasArms() bool`

HasArms returns a boolean if a field has been set.

### GetName

`func (o *Preset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Preset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Preset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Preset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *Preset) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Preset) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Preset) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Preset) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOwner

`func (o *Preset) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Preset) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Preset) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Preset) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPanel

`func (o *Preset) GetPanel() int64`

GetPanel returns the Panel field if non-nil, zero value otherwise.

### GetPanelOk

`func (o *Preset) GetPanelOk() (*int64, bool)`

GetPanelOk returns a tuple with the Panel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanel

`func (o *Preset) SetPanel(v int64)`

SetPanel sets Panel field to given value.

### HasPanel

`func (o *Preset) HasPanel() bool`

HasPanel returns a boolean if a field has been set.

### GetRank

`func (o *Preset) GetRank() []string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Preset) GetRankOk() (*[]string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Preset) SetRank(v []string)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Preset) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


