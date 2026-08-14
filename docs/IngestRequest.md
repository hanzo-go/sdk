# IngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to [**[]Attempt**](Attempt.md) |  | [optional] 
**Experiments** | Pointer to [**[]Experiment**](Experiment.md) |  | [optional] 

## Methods

### NewIngestRequest

`func NewIngestRequest() *IngestRequest`

NewIngestRequest instantiates a new IngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestRequestWithDefaults

`func NewIngestRequestWithDefaults() *IngestRequest`

NewIngestRequestWithDefaults instantiates a new IngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *IngestRequest) GetAttempts() []Attempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *IngestRequest) GetAttemptsOk() (*[]Attempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *IngestRequest) SetAttempts(v []Attempt)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *IngestRequest) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetExperiments

`func (o *IngestRequest) GetExperiments() []Experiment`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *IngestRequest) GetExperimentsOk() (*[]Experiment, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *IngestRequest) SetExperiments(v []Experiment)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *IngestRequest) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


