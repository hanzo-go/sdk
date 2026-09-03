# BlueprintCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principles** | Pointer to **int64** | Principles is how many spine archetypes the playbook carries (64 in the shipped corpus). | [optional] 
**Sections** | Pointer to **int64** | Sections is how many phases the journey has. | [optional] 
**Steps** | Pointer to **int64** | Steps is how many checklist items the playbook holds, DISABLED ONES INCLUDED — this counts the authored document, not the journey an org runs, so it is normally larger than the &#x60;total&#x60; on a progress view. | [optional] 
**Strategies** | Pointer to **int64** | Strategies is how many tactics the corpus holds, again counting disabled ones. | [optional] 
**Templates** | Pointer to **int64** | Templates is how many reusable prompts the playbook carries. | [optional] 

## Methods

### NewBlueprintCounts

`func NewBlueprintCounts() *BlueprintCounts`

NewBlueprintCounts instantiates a new BlueprintCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintCountsWithDefaults

`func NewBlueprintCountsWithDefaults() *BlueprintCounts`

NewBlueprintCountsWithDefaults instantiates a new BlueprintCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrinciples

`func (o *BlueprintCounts) GetPrinciples() int64`

GetPrinciples returns the Principles field if non-nil, zero value otherwise.

### GetPrinciplesOk

`func (o *BlueprintCounts) GetPrinciplesOk() (*int64, bool)`

GetPrinciplesOk returns a tuple with the Principles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinciples

`func (o *BlueprintCounts) SetPrinciples(v int64)`

SetPrinciples sets Principles field to given value.

### HasPrinciples

`func (o *BlueprintCounts) HasPrinciples() bool`

HasPrinciples returns a boolean if a field has been set.

### GetSections

`func (o *BlueprintCounts) GetSections() int64`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *BlueprintCounts) GetSectionsOk() (*int64, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *BlueprintCounts) SetSections(v int64)`

SetSections sets Sections field to given value.

### HasSections

`func (o *BlueprintCounts) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetSteps

`func (o *BlueprintCounts) GetSteps() int64`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *BlueprintCounts) GetStepsOk() (*int64, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *BlueprintCounts) SetSteps(v int64)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *BlueprintCounts) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStrategies

`func (o *BlueprintCounts) GetStrategies() int64`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *BlueprintCounts) GetStrategiesOk() (*int64, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *BlueprintCounts) SetStrategies(v int64)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *BlueprintCounts) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.

### GetTemplates

`func (o *BlueprintCounts) GetTemplates() int64`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *BlueprintCounts) GetTemplatesOk() (*int64, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *BlueprintCounts) SetTemplates(v int64)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *BlueprintCounts) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


