# ProjectSummary

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

### NewProjectSummary

`func NewProjectSummary() *ProjectSummary`

NewProjectSummary instantiates a new ProjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectSummaryWithDefaults

`func NewProjectSummaryWithDefaults() *ProjectSummary`

NewProjectSummaryWithDefaults instantiates a new ProjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *ProjectSummary) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ProjectSummary) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ProjectSummary) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ProjectSummary) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *ProjectSummary) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *ProjectSummary) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *ProjectSummary) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *ProjectSummary) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetBenchmarks

`func (o *ProjectSummary) GetBenchmarks() int32`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *ProjectSummary) GetBenchmarksOk() (*int32, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *ProjectSummary) SetBenchmarks(v int32)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *ProjectSummary) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetCostUsd

`func (o *ProjectSummary) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *ProjectSummary) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *ProjectSummary) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *ProjectSummary) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetExperiments

`func (o *ProjectSummary) GetExperiments() int32`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *ProjectSummary) GetExperimentsOk() (*int32, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *ProjectSummary) SetExperiments(v int32)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *ProjectSummary) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *ProjectSummary) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *ProjectSummary) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *ProjectSummary) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *ProjectSummary) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetKinds

`func (o *ProjectSummary) GetKinds() []string`

GetKinds returns the Kinds field if non-nil, zero value otherwise.

### GetKindsOk

`func (o *ProjectSummary) GetKindsOk() (*[]string, bool)`

GetKindsOk returns a tuple with the Kinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinds

`func (o *ProjectSummary) SetKinds(v []string)`

SetKinds sets Kinds field to given value.

### HasKinds

`func (o *ProjectSummary) HasKinds() bool`

HasKinds returns a boolean if a field has been set.

### GetModels

`func (o *ProjectSummary) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *ProjectSummary) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *ProjectSummary) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *ProjectSummary) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProject

`func (o *ProjectSummary) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectSummary) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectSummary) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectSummary) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


