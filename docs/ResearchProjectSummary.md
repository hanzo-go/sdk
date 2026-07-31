# ResearchProjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | canonical | [optional] 
**AttemptsRetained** | Pointer to **int32** |  | [optional] 
**Benchmarks** | Pointer to **int32** |  | [optional] 
**CostUsd** | Pointer to **float32** |  | [optional] 
**Experiments** | Pointer to **int32** | canonical | [optional] 
**ExperimentsRetained** | Pointer to **int32** |  | [optional] 
**Kinds** | Pointer to **[]string** |  | [optional] 
**Models** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 

## Methods

### NewResearchProjectSummary

`func NewResearchProjectSummary() *ResearchProjectSummary`

NewResearchProjectSummary instantiates a new ResearchProjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchProjectSummaryWithDefaults

`func NewResearchProjectSummaryWithDefaults() *ResearchProjectSummary`

NewResearchProjectSummaryWithDefaults instantiates a new ResearchProjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ResearchProjectSummary) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ResearchProjectSummary) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ResearchProjectSummary) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ResearchProjectSummary) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *ResearchProjectSummary) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *ResearchProjectSummary) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *ResearchProjectSummary) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *ResearchProjectSummary) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetBenchmarks

`func (o *ResearchProjectSummary) GetBenchmarks() int32`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *ResearchProjectSummary) GetBenchmarksOk() (*int32, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *ResearchProjectSummary) SetBenchmarks(v int32)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *ResearchProjectSummary) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetCostUsd

`func (o *ResearchProjectSummary) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ResearchProjectSummary) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ResearchProjectSummary) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ResearchProjectSummary) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetExperiments

`func (o *ResearchProjectSummary) GetExperiments() int32`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *ResearchProjectSummary) GetExperimentsOk() (*int32, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *ResearchProjectSummary) SetExperiments(v int32)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *ResearchProjectSummary) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *ResearchProjectSummary) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *ResearchProjectSummary) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *ResearchProjectSummary) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *ResearchProjectSummary) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetKinds

`func (o *ResearchProjectSummary) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *ResearchProjectSummary) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *ResearchProjectSummary) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *ResearchProjectSummary) HasKinds() bool`

HasKinds returns a boolean if a field has been set.

### GetModels

`func (o *ResearchProjectSummary) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ResearchProjectSummary) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ResearchProjectSummary) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *ResearchProjectSummary) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProject

`func (o *ResearchProjectSummary) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchProjectSummary) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchProjectSummary) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchProjectSummary) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


