# ResearchIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experiments** | Pointer to [**[]ResearchExperiment**](ResearchExperiment.md) |  | [optional] 
**Attempts** | Pointer to [**[]ResearchAttempt**](ResearchAttempt.md) |  | [optional] 

## Methods

### NewResearchIngestRequest

`func NewResearchIngestRequest() *ResearchIngestRequest`

NewResearchIngestRequest instantiates a new ResearchIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResearchIngestRequestWithDefaults

`func NewResearchIngestRequestWithDefaults() *ResearchIngestRequest`

NewResearchIngestRequestWithDefaults instantiates a new ResearchIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperiments

`func (o *ResearchIngestRequest) GetExperiments() []ResearchExperiment`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *ResearchIngestRequest) GetExperimentsOk() (*[]ResearchExperiment, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *ResearchIngestRequest) SetExperiments(v []ResearchExperiment)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *ResearchIngestRequest) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.

### GetAttempts

`func (o *ResearchIngestRequest) GetAttempts() []ResearchAttempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *ResearchIngestRequest) GetAttemptsOk() (*[]ResearchAttempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *ResearchIngestRequest) SetAttempts(v []ResearchAttempt)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *ResearchIngestRequest) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


