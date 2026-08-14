# Blueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Principles** | Pointer to [**[]Principle**](Principle.md) | the 64-principle spine (Zen of Hanzo archetypes) | [optional] 
**Sections** | Pointer to [**[]Section**](Section.md) |  | [optional] 
**Steps** | Pointer to [**[]JourneyStep**](JourneyStep.md) |  | [optional] 
**Strategies** | Pointer to [**[]Strategy**](Strategy.md) |  | [optional] 
**Templates** | Pointer to [**[]Page**](Page.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewBlueprint

`func NewBlueprint() *Blueprint`

NewBlueprint instantiates a new Blueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintWithDefaults

`func NewBlueprintWithDefaults() *Blueprint`

NewBlueprintWithDefaults instantiates a new Blueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *Blueprint) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *Blueprint) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *Blueprint) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *Blueprint) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetEnabled

`func (o *Blueprint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Blueprint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Blueprint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Blueprint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrinciples

`func (o *Blueprint) GetPrinciples() []Principle`

GetPrinciples returns the Principles field if non-nil, zero value otherwise.

### GetPrinciplesOk

`func (o *Blueprint) GetPrinciplesOk() (*[]Principle, bool)`

GetPrinciplesOk returns a tuple with the Principles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciples

`func (o *Blueprint) SetPrinciples(v []Principle)`

SetPrinciples sets Principles field to given value.

### HasPrinciples

`func (o *Blueprint) HasPrinciples() bool`

HasPrinciples returns a boolean if a field has been set.

### GetSections

`func (o *Blueprint) GetSections() []Section`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *Blueprint) GetSectionsOk() (*[]Section, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *Blueprint) SetSections(v []Section)`

SetSections sets Sections field to given value.

### HasSections

`func (o *Blueprint) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSteps

`func (o *Blueprint) GetSteps() []JourneyStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Blueprint) GetStepsOk() (*[]JourneyStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Blueprint) SetSteps(v []JourneyStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Blueprint) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStrategies

`func (o *Blueprint) GetStrategies() []Strategy`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *Blueprint) GetStrategiesOk() (*[]Strategy, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *Blueprint) SetStrategies(v []Strategy)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *Blueprint) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetTemplates

`func (o *Blueprint) GetTemplates() []Page`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Blueprint) GetTemplatesOk() (*[]Page, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *Blueprint) SetTemplates(v []Page)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *Blueprint) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTitle

`func (o *Blueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Blueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Blueprint) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Blueprint) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *Blueprint) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Blueprint) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Blueprint) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Blueprint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


