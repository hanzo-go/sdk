# ResearchIngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**ExperimentsIngested** | Pointer to **int32** | NEW versions added this batch | [optional] 
**AttemptsIngested** | Pointer to **int32** |  | [optional] 
**CanonicalExperiments** | Pointer to **int32** |  | [optional] 
**ExperimentsRetained** | Pointer to **int32** |  | [optional] 
**CanonicalAttempts** | Pointer to **int32** |  | [optional] 
**AttemptsRetained** | Pointer to **int32** |  | [optional] 
**RolledUp** | Pointer to **bool** | false when the datastore was absent (SQLite retains the write) | [optional] 

## Methods

### NewResearchIngestResult

`func NewResearchIngestResult() *ResearchIngestResult`

NewResearchIngestResult instantiates a new ResearchIngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchIngestResultWithDefaults

`func NewResearchIngestResultWithDefaults() *ResearchIngestResult`

NewResearchIngestResultWithDefaults instantiates a new ResearchIngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ResearchIngestResult) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ResearchIngestResult) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ResearchIngestResult) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ResearchIngestResult) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetExperimentsIngested

`func (o *ResearchIngestResult) GetExperimentsIngested() int32`

GetExperimentsIngested returns the ExperimentsIngested field if non-nil, zero value otherwise.

### GetExperimentsIngestedOk

`func (o *ResearchIngestResult) GetExperimentsIngestedOk() (*int32, bool)`

GetExperimentsIngestedOk returns a tuple with the ExperimentsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsIngested

`func (o *ResearchIngestResult) SetExperimentsIngested(v int32)`

SetExperimentsIngested sets ExperimentsIngested field to given value.

### HasExperimentsIngested

`func (o *ResearchIngestResult) HasExperimentsIngested() bool`

HasExperimentsIngested returns a boolean if a field has been set.

### GetAttemptsIngested

`func (o *ResearchIngestResult) GetAttemptsIngested() int32`

GetAttemptsIngested returns the AttemptsIngested field if non-nil, zero value otherwise.

### GetAttemptsIngestedOk

`func (o *ResearchIngestResult) GetAttemptsIngestedOk() (*int32, bool)`

GetAttemptsIngestedOk returns a tuple with the AttemptsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsIngested

`func (o *ResearchIngestResult) SetAttemptsIngested(v int32)`

SetAttemptsIngested sets AttemptsIngested field to given value.

### HasAttemptsIngested

`func (o *ResearchIngestResult) HasAttemptsIngested() bool`

HasAttemptsIngested returns a boolean if a field has been set.

### GetCanonicalExperiments

`func (o *ResearchIngestResult) GetCanonicalExperiments() int32`

GetCanonicalExperiments returns the CanonicalExperiments field if non-nil, zero value otherwise.

### GetCanonicalExperimentsOk

`func (o *ResearchIngestResult) GetCanonicalExperimentsOk() (*int32, bool)`

GetCanonicalExperimentsOk returns a tuple with the CanonicalExperiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalExperiments

`func (o *ResearchIngestResult) SetCanonicalExperiments(v int32)`

SetCanonicalExperiments sets CanonicalExperiments field to given value.

### HasCanonicalExperiments

`func (o *ResearchIngestResult) HasCanonicalExperiments() bool`

HasCanonicalExperiments returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *ResearchIngestResult) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *ResearchIngestResult) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *ResearchIngestResult) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *ResearchIngestResult) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetCanonicalAttempts

`func (o *ResearchIngestResult) GetCanonicalAttempts() int32`

GetCanonicalAttempts returns the CanonicalAttempts field if non-nil, zero value otherwise.

### GetCanonicalAttemptsOk

`func (o *ResearchIngestResult) GetCanonicalAttemptsOk() (*int32, bool)`

GetCanonicalAttemptsOk returns a tuple with the CanonicalAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalAttempts

`func (o *ResearchIngestResult) SetCanonicalAttempts(v int32)`

SetCanonicalAttempts sets CanonicalAttempts field to given value.

### HasCanonicalAttempts

`func (o *ResearchIngestResult) HasCanonicalAttempts() bool`

HasCanonicalAttempts returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *ResearchIngestResult) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *ResearchIngestResult) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *ResearchIngestResult) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *ResearchIngestResult) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetRolledUp

`func (o *ResearchIngestResult) GetRolledUp() bool`

GetRolledUp returns the RolledUp field if non-nil, zero value otherwise.

### GetRolledUpOk

`func (o *ResearchIngestResult) GetRolledUpOk() (*bool, bool)`

GetRolledUpOk returns a tuple with the RolledUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUp

`func (o *ResearchIngestResult) SetRolledUp(v bool)`

SetRolledUp sets RolledUp field to given value.

### HasRolledUp

`func (o *ResearchIngestResult) HasRolledUp() bool`

HasRolledUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


