# CloudIngestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to [**[]CloudAttempt**](CloudAttempt.md) |  | [optional] 
**Experiments** | Pointer to [**[]CloudExperiment**](CloudExperiment.md) |  | [optional] 

## Methods

### NewCloudIngestRequest

`func NewCloudIngestRequest() *CloudIngestRequest`

NewCloudIngestRequest instantiates a new CloudIngestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudIngestRequestWithDefaults

`func NewCloudIngestRequestWithDefaults() *CloudIngestRequest`

NewCloudIngestRequestWithDefaults instantiates a new CloudIngestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *CloudIngestRequest) GetAttempts() []CloudAttempt`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *CloudIngestRequest) GetAttemptsOk() (*[]CloudAttempt, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *CloudIngestRequest) SetAttempts(v []CloudAttempt)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *CloudIngestRequest) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetExperiments

`func (o *CloudIngestRequest) GetExperiments() []CloudExperiment`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *CloudIngestRequest) GetExperimentsOk() (*[]CloudExperiment, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *CloudIngestRequest) SetExperiments(v []CloudExperiment)`

SetExperiments sets Experiments field to given value.

### HasExperiments

`func (o *CloudIngestRequest) HasExperiments() bool`

HasExperiments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


