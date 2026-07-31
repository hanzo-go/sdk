# CloudBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Principles** | Pointer to [**[]CloudPrinciple**](CloudPrinciple.md) | the 64-principle spine (Zen of Hanzo archetypes) | [optional] 
**Sections** | Pointer to [**[]CloudSection**](CloudSection.md) |  | [optional] 
**Steps** | Pointer to [**[]CloudJourneyStep**](CloudJourneyStep.md) |  | [optional] 
**Strategies** | Pointer to [**[]CloudStrategy**](CloudStrategy.md) |  | [optional] 
**Templates** | Pointer to [**[]CloudTemplate**](CloudTemplate.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudBlueprint

`func NewCloudBlueprint() *CloudBlueprint`

NewCloudBlueprint instantiates a new CloudBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBlueprintWithDefaults

`func NewCloudBlueprintWithDefaults() *CloudBlueprint`

NewCloudBlueprintWithDefaults instantiates a new CloudBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *CloudBlueprint) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CloudBlueprint) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CloudBlueprint) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CloudBlueprint) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudBlueprint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudBlueprint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudBlueprint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudBlueprint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrinciples

`func (o *CloudBlueprint) GetPrinciples() []CloudPrinciple`

GetPrinciples returns the Principles field if non-nil, zero value otherwise.

### GetPrinciplesOk

`func (o *CloudBlueprint) GetPrinciplesOk() (*[]CloudPrinciple, bool)`

GetPrinciplesOk returns a tuple with the Principles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciples

`func (o *CloudBlueprint) SetPrinciples(v []CloudPrinciple)`

SetPrinciples sets Principles field to given value.

### HasPrinciples

`func (o *CloudBlueprint) HasPrinciples() bool`

HasPrinciples returns a boolean if a field has been set.

### GetSections

`func (o *CloudBlueprint) GetSections() []CloudSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *CloudBlueprint) GetSectionsOk() (*[]CloudSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *CloudBlueprint) SetSections(v []CloudSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *CloudBlueprint) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSteps

`func (o *CloudBlueprint) GetSteps() []CloudJourneyStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CloudBlueprint) GetStepsOk() (*[]CloudJourneyStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CloudBlueprint) SetSteps(v []CloudJourneyStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CloudBlueprint) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStrategies

`func (o *CloudBlueprint) GetStrategies() []CloudStrategy`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *CloudBlueprint) GetStrategiesOk() (*[]CloudStrategy, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *CloudBlueprint) SetStrategies(v []CloudStrategy)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *CloudBlueprint) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetTemplates

`func (o *CloudBlueprint) GetTemplates() []CloudTemplate`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *CloudBlueprint) GetTemplatesOk() (*[]CloudTemplate, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *CloudBlueprint) SetTemplates(v []CloudTemplate)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *CloudBlueprint) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTitle

`func (o *CloudBlueprint) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudBlueprint) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudBlueprint) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudBlueprint) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *CloudBlueprint) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudBlueprint) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudBlueprint) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudBlueprint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


