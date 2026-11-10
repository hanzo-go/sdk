# StudioRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**ChargedCents** | Pointer to **int64** | ChargedCents is what the finished steps cost; final once the run ends. | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**Error** | Pointer to [**StudioProblem**](StudioProblem.md) |  | [optional] 
**EstimateCents** | Pointer to **int64** | EstimateCents is what was reserved when the run started. | [optional] 
**Finished** | Pointer to **string** |  | [optional] 
**Flow** | Pointer to **string** |  | [optional] 
**Graph** | Pointer to [**StudioGraph**](StudioGraph.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **map[string]interface{}** | Input is what the run was given, as its inputs schema reads it. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to [**[]StudioItem**](StudioItem.md) | Outputs are the Library items the run made, the final steps&#39; first. | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Root** | Pointer to **string** | Root is the first version&#39;s id; equal to id on version 1. | [optional] 
**Shared** | Pointer to **bool** |  | [optional] 
**Started** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]StudioStep**](StudioStep.md) |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** | Text is the last chat step&#39;s text, when the run&#39;s product is words. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewStudioRun

`func NewStudioRun() *StudioRun`

NewStudioRun instantiates a new StudioRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioRunWithDefaults

`func NewStudioRunWithDefaults() *StudioRun`

NewStudioRunWithDefaults instantiates a new StudioRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StudioRun) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StudioRun) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StudioRun) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StudioRun) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetChargedCents

`func (o *StudioRun) GetChargedCents() int64`

GetChargedCents returns the ChargedCents field if non-nil, zero value otherwise.

### GetChargedCentsOk

`func (o *StudioRun) GetChargedCentsOk() (*int64, bool)`

GetChargedCentsOk returns a tuple with the ChargedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargedCents

`func (o *StudioRun) SetChargedCents(v int64)`

SetChargedCents sets ChargedCents field to given value.

### HasChargedCents

`func (o *StudioRun) HasChargedCents() bool`

HasChargedCents returns a boolean if a field has been set.

### GetCreated

`func (o *StudioRun) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StudioRun) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StudioRun) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StudioRun) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreator

`func (o *StudioRun) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *StudioRun) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *StudioRun) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *StudioRun) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetError

`func (o *StudioRun) GetError() StudioProblem`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StudioRun) GetErrorOk() (*StudioProblem, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StudioRun) SetError(v StudioProblem)`

SetError sets Error field to given value.

### HasError

`func (o *StudioRun) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEstimateCents

`func (o *StudioRun) GetEstimateCents() int64`

GetEstimateCents returns the EstimateCents field if non-nil, zero value otherwise.

### GetEstimateCentsOk

`func (o *StudioRun) GetEstimateCentsOk() (*int64, bool)`

GetEstimateCentsOk returns a tuple with the EstimateCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateCents

`func (o *StudioRun) SetEstimateCents(v int64)`

SetEstimateCents sets EstimateCents field to given value.

### HasEstimateCents

`func (o *StudioRun) HasEstimateCents() bool`

HasEstimateCents returns a boolean if a field has been set.

### GetFinished

`func (o *StudioRun) GetFinished() string`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *StudioRun) GetFinishedOk() (*string, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *StudioRun) SetFinished(v string)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *StudioRun) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetFlow

`func (o *StudioRun) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *StudioRun) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *StudioRun) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *StudioRun) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetGraph

`func (o *StudioRun) GetGraph() StudioGraph`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *StudioRun) GetGraphOk() (*StudioGraph, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *StudioRun) SetGraph(v StudioGraph)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *StudioRun) HasGraph() bool`

HasGraph returns a boolean if a field has been set.

### GetId

`func (o *StudioRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *StudioRun) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *StudioRun) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *StudioRun) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *StudioRun) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetKind

`func (o *StudioRun) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StudioRun) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StudioRun) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StudioRun) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrg

`func (o *StudioRun) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StudioRun) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StudioRun) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *StudioRun) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetOutputs

`func (o *StudioRun) GetOutputs() []StudioItem`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *StudioRun) GetOutputsOk() (*[]StudioItem, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *StudioRun) SetOutputs(v []StudioItem)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *StudioRun) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetParent

`func (o *StudioRun) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *StudioRun) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *StudioRun) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *StudioRun) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRoot

`func (o *StudioRun) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *StudioRun) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *StudioRun) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *StudioRun) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetShared

`func (o *StudioRun) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *StudioRun) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *StudioRun) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *StudioRun) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetStarted

`func (o *StudioRun) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *StudioRun) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *StudioRun) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *StudioRun) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *StudioRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *StudioRun) GetSteps() []StudioStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *StudioRun) GetStepsOk() (*[]StudioStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *StudioRun) SetSteps(v []StudioStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *StudioRun) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTemplate

`func (o *StudioRun) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *StudioRun) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *StudioRun) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *StudioRun) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetText

`func (o *StudioRun) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StudioRun) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StudioRun) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StudioRun) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *StudioRun) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioRun) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioRun) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioRun) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVersion

`func (o *StudioRun) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StudioRun) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StudioRun) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StudioRun) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


