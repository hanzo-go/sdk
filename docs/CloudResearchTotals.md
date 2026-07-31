# CloudResearchTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | canonical | [optional] 
**AttemptsRetained** | Pointer to **int32** |  | [optional] 
**Benchmarks** | Pointer to **int32** |  | [optional] 
**ByKind** | Pointer to [**[]CloudKindTotal**](CloudKindTotal.md) |  | [optional] 
**CostUsd** | Pointer to **float32** |  | [optional] 
**Experiments** | Pointer to **int32** | canonical | [optional] 
**ExperimentsRetained** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudResearchTotals

`func NewCloudResearchTotals() *CloudResearchTotals`

NewCloudResearchTotals instantiates a new CloudResearchTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudResearchTotalsWithDefaults

`func NewCloudResearchTotalsWithDefaults() *CloudResearchTotals`

NewCloudResearchTotalsWithDefaults instantiates a new CloudResearchTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *CloudResearchTotals) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *CloudResearchTotals) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *CloudResearchTotals) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *CloudResearchTotals) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *CloudResearchTotals) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *CloudResearchTotals) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *CloudResearchTotals) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *CloudResearchTotals) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetBenchmarks

`func (o *CloudResearchTotals) GetBenchmarks() int32`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *CloudResearchTotals) GetBenchmarksOk() (*int32, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *CloudResearchTotals) SetBenchmarks(v int32)`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *CloudResearchTotals) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.

### GetByKind

`func (o *CloudResearchTotals) GetByKind() []CloudKindTotal`

GetByKind returns the ByKind field if non-nil, zero value otherwise.

### GetByKindOk

`func (o *CloudResearchTotals) GetByKindOk() (*[]CloudKindTotal, bool)`

GetByKindOk returns a tuple with the ByKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByKind

`func (o *CloudResearchTotals) SetByKind(v []CloudKindTotal)`

SetByKind sets ByKind field to given value.

### HasByKind

`func (o *CloudResearchTotals) HasByKind() bool`

HasByKind returns a boolean if a field has been set.

### GetCostUsd

`func (o *CloudResearchTotals) GetCostUsd() float32`

GetCostUsd returns the CostUsd field if non-nil, zero value otherwise.

### GetCostUsdOk

`func (o *CloudResearchTotals) GetCostUsdOk() (*float32, bool)`

GetCostUsdOk returns a tuple with the CostUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostUsd

`func (o *CloudResearchTotals) SetCostUsd(v float32)`

SetCostUsd sets CostUsd field to given value.

### HasCostUsd

`func (o *CloudResearchTotals) HasCostUsd() bool`

HasCostUsd returns a boolean if a field has been set.

### GetExperiments

`func (o *CloudResearchTotals) GetExperiments() int32`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *CloudResearchTotals) GetExperimentsOk() (*int32, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *CloudResearchTotals) SetExperiments(v int32)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *CloudResearchTotals) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *CloudResearchTotals) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *CloudResearchTotals) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *CloudResearchTotals) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *CloudResearchTotals) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetModels

`func (o *CloudResearchTotals) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *CloudResearchTotals) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *CloudResearchTotals) SetModels(v int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *CloudResearchTotals) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetProject

`func (o *CloudResearchTotals) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudResearchTotals) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudResearchTotals) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudResearchTotals) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjects

`func (o *CloudResearchTotals) GetProjects() int32`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CloudResearchTotals) GetProjectsOk() (*int32, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CloudResearchTotals) SetProjects(v int32)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CloudResearchTotals) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


