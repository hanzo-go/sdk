# IngestOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptsIngested** | Pointer to **int32** | AttemptsIngested is how many attempt versions this call appended. | [optional] 
**AttemptsRetained** | Pointer to **int32** | AttemptsRetained is the full versioned attempt history the store now holds. | [optional] 
**CanonicalAttempts** | Pointer to **int32** | CanonicalAttempts is the deduped attempt count the store now holds. | [optional] 
**CanonicalExperiments** | Pointer to **int32** | CanonicalExperiments is the deduped experiment count the store now holds. | [optional] 
**ExperimentsIngested** | Pointer to **int32** | ExperimentsIngested is how many experiment versions this call appended. | [optional] 
**ExperimentsRetained** | Pointer to **int32** | ExperimentsRetained is the full versioned experiment history the store now holds. | [optional] 
**Project** | Pointer to **string** | Project is the project the batch was filed under — the SERVER&#39;s value, never the body&#39;s. | [optional] 
**RolledUp** | Pointer to **bool** | RolledUp is false when the OLAP roll-up was skipped; the SQLite write still stands. | [optional] 

## Methods

### NewIngestOut

`func NewIngestOut() *IngestOut`

NewIngestOut instantiates a new IngestOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestOutWithDefaults

`func NewIngestOutWithDefaults() *IngestOut`

NewIngestOutWithDefaults instantiates a new IngestOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptsIngested

`func (o *IngestOut) GetAttemptsIngested() int32`

GetAttemptsIngested returns the AttemptsIngested field if non-nil, zero value otherwise.

### GetAttemptsIngestedOk

`func (o *IngestOut) GetAttemptsIngestedOk() (*int32, bool)`

GetAttemptsIngestedOk returns a tuple with the AttemptsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsIngested

`func (o *IngestOut) SetAttemptsIngested(v int32)`

SetAttemptsIngested sets AttemptsIngested field to given value.

### HasAttemptsIngested

`func (o *IngestOut) HasAttemptsIngested() bool`

HasAttemptsIngested returns a boolean if a field has been set.

### GetAttemptsRetained

`func (o *IngestOut) GetAttemptsRetained() int32`

GetAttemptsRetained returns the AttemptsRetained field if non-nil, zero value otherwise.

### GetAttemptsRetainedOk

`func (o *IngestOut) GetAttemptsRetainedOk() (*int32, bool)`

GetAttemptsRetainedOk returns a tuple with the AttemptsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsRetained

`func (o *IngestOut) SetAttemptsRetained(v int32)`

SetAttemptsRetained sets AttemptsRetained field to given value.

### HasAttemptsRetained

`func (o *IngestOut) HasAttemptsRetained() bool`

HasAttemptsRetained returns a boolean if a field has been set.

### GetCanonicalAttempts

`func (o *IngestOut) GetCanonicalAttempts() int32`

GetCanonicalAttempts returns the CanonicalAttempts field if non-nil, zero value otherwise.

### GetCanonicalAttemptsOk

`func (o *IngestOut) GetCanonicalAttemptsOk() (*int32, bool)`

GetCanonicalAttemptsOk returns a tuple with the CanonicalAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalAttempts

`func (o *IngestOut) SetCanonicalAttempts(v int32)`

SetCanonicalAttempts sets CanonicalAttempts field to given value.

### HasCanonicalAttempts

`func (o *IngestOut) HasCanonicalAttempts() bool`

HasCanonicalAttempts returns a boolean if a field has been set.

### GetCanonicalExperiments

`func (o *IngestOut) GetCanonicalExperiments() int32`

GetCanonicalExperiments returns the CanonicalExperiments field if non-nil, zero value otherwise.

### GetCanonicalExperimentsOk

`func (o *IngestOut) GetCanonicalExperimentsOk() (*int32, bool)`

GetCanonicalExperimentsOk returns a tuple with the CanonicalExperiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalExperiments

`func (o *IngestOut) SetCanonicalExperiments(v int32)`

SetCanonicalExperiments sets CanonicalExperiments field to given value.

### HasCanonicalExperiments

`func (o *IngestOut) HasCanonicalExperiments() bool`

HasCanonicalExperiments returns a boolean if a field has been set.

### GetExperimentsIngested

`func (o *IngestOut) GetExperimentsIngested() int32`

GetExperimentsIngested returns the ExperimentsIngested field if non-nil, zero value otherwise.

### GetExperimentsIngestedOk

`func (o *IngestOut) GetExperimentsIngestedOk() (*int32, bool)`

GetExperimentsIngestedOk returns a tuple with the ExperimentsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsIngested

`func (o *IngestOut) SetExperimentsIngested(v int32)`

SetExperimentsIngested sets ExperimentsIngested field to given value.

### HasExperimentsIngested

`func (o *IngestOut) HasExperimentsIngested() bool`

HasExperimentsIngested returns a boolean if a field has been set.

### GetExperimentsRetained

`func (o *IngestOut) GetExperimentsRetained() int32`

GetExperimentsRetained returns the ExperimentsRetained field if non-nil, zero value otherwise.

### GetExperimentsRetainedOk

`func (o *IngestOut) GetExperimentsRetainedOk() (*int32, bool)`

GetExperimentsRetainedOk returns a tuple with the ExperimentsRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentsRetained

`func (o *IngestOut) SetExperimentsRetained(v int32)`

SetExperimentsRetained sets ExperimentsRetained field to given value.

### HasExperimentsRetained

`func (o *IngestOut) HasExperimentsRetained() bool`

HasExperimentsRetained returns a boolean if a field has been set.

### GetProject

`func (o *IngestOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IngestOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IngestOut) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *IngestOut) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRolledUp

`func (o *IngestOut) GetRolledUp() bool`

GetRolledUp returns the RolledUp field if non-nil, zero value otherwise.

### GetRolledUpOk

`func (o *IngestOut) GetRolledUpOk() (*bool, bool)`

GetRolledUpOk returns a tuple with the RolledUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolledUp

`func (o *IngestOut) SetRolledUp(v bool)`

SetRolledUp sets RolledUp field to given value.

### HasRolledUp

`func (o *IngestOut) HasRolledUp() bool`

HasRolledUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


