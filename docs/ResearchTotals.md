# ResearchTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | canonical | [optional] 
**AttemptsRetained** | Pointer to **int32** |  | [optional] 
**Benchmarks** | Pointer to **int32** |  | [optional] 
**ByKind** | Pointer to [**[]KindTotal**](KindTotal.md) |  | [optional] 
**CostUsd** | Pointer to **float32** |  | [optional] 
**Experiments** | Pointer to **int32** | canonical | [optional] 
**ExperimentsRetained** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to **int32** |  | [optional] 

## Methods

### NewResearchTotals

`func NewResearchTotals() *ResearchTotals`

NewResearchTotals instantiates a new ResearchTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchTotalsWithDefaults

`func NewResearchTotalsWithDefaults() *ResearchTotals`

NewResearchTotalsWithDefaults instantiates a new ResearchTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ResearchTotals) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ResearchTotals) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ResearchTotals) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ResearchTotals) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *ResearchTotals) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *ResearchTotals) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *ResearchTotals) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *ResearchTotals) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetBenchmarks

`func (o *ResearchTotals) GetBenchmarks() int32`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *ResearchTotals) GetBenchmarksOk() (*int32, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *ResearchTotals) SetBenchmarks(v int32)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *ResearchTotals) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetByKind

`func (o *ResearchTotals) GetByKind() []KindTotal`

GetByKind returns the ByKind field if non-nil, zero value otherwise.

### GetByKindOk

`func (o *ResearchTotals) GetByKindOk() (*[]KindTotal, bool)`

GetByKindOk returns a tuple with the ByKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByKind

`func (o *ResearchTotals) SetByKind(v []KindTotal)`

SetByKind sets ByKind field to given value.

### HasByKind

`func (o *ResearchTotals) HasByKind() bool`

HasByKind returns a boolean if a field has been set.

### GetCostUsd

`func (o *ResearchTotals) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ResearchTotals) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ResearchTotals) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ResearchTotals) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetExperiments

`func (o *ResearchTotals) GetExperiments() int32`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *ResearchTotals) GetExperimentsOk() (*int32, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *ResearchTotals) SetExperiments(v int32)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *ResearchTotals) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *ResearchTotals) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *ResearchTotals) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *ResearchTotals) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *ResearchTotals) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetModels

`func (o *ResearchTotals) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ResearchTotals) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ResearchTotals) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *ResearchTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProject

`func (o *ResearchTotals) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchTotals) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchTotals) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchTotals) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *ResearchTotals) GetProjects() int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ResearchTotals) GetProjectsOk() (*int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ResearchTotals) SetProjects(v int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ResearchTotals) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


