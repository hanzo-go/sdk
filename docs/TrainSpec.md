# TrainSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adaptation** | Pointer to [**TrainAdaptation**](TrainAdaptation.md) | Adaptation is how the update is parameterized. | [optional] 
**BaseModel** | Pointer to **string** | BaseModel is the base the capability is taught to: &#x60;kai&#x60;, or a model the engine loads for a client. | [optional] 
**Dataset** | Pointer to [**TrainDataset**](TrainDataset.md) | Dataset is what the job trains on. | [optional] 
**Evaluation** | Pointer to [**TrainEvaluation**](TrainEvaluation.md) | Evaluation names the target suites the result reports. | [optional] 
**From** | Pointer to **string** | From is an artifact the job starts from, by sha256. | [optional] 
**Objective** | Pointer to [**TrainObjective**](TrainObjective.md) | Objective is the loss and the terms beside it. | [optional] 
**Output** | Pointer to [**TrainOutput**](TrainOutput.md) | Output is what the job produces. | [optional] 
**Protect** | Pointer to [**TrainProtect**](TrainProtect.md) | Protect names what the base already does that must hold, and the budget. | [optional] 
**Resources** | Pointer to [**TrainResources**](TrainResources.md) | Resources places and bounds the run. | [optional] 
**Revision** | Pointer to **string** | Revision pins the base. For kai: a finished run&#39;s name on the executor, or &#x60;hub:&lt;repo&gt;@&lt;rev&gt;&#x60;. | [optional] 

## Methods

### NewTrainSpec

`func NewTrainSpec() *TrainSpec`

NewTrainSpec instantiates a new TrainSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainSpecWithDefaults

`func NewTrainSpecWithDefaults() *TrainSpec`

NewTrainSpecWithDefaults instantiates a new TrainSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdaptation

`func (o *TrainSpec) GetAdaptation() TrainAdaptation`

GetAdaptation returns the Adaptation field if non-nil, zero value otherwise.

### GetAdaptationOk

`func (o *TrainSpec) GetAdaptationOk() (*TrainAdaptation, bool)`

GetAdaptationOk returns a tuple with the Adaptation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptation

`func (o *TrainSpec) SetAdaptation(v TrainAdaptation)`

SetAdaptation sets Adaptation field to given value.

### HasAdaptation

`func (o *TrainSpec) HasAdaptation() bool`

HasAdaptation returns a boolean if a field has been set.

### GetBaseModel

`func (o *TrainSpec) GetBaseModel() string`

GetBaseModel returns the BaseModel field if non-nil, zero value otherwise.

### GetBaseModelOk

`func (o *TrainSpec) GetBaseModelOk() (*string, bool)`

GetBaseModelOk returns a tuple with the BaseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseModel

`func (o *TrainSpec) SetBaseModel(v string)`

SetBaseModel sets BaseModel field to given value.

### HasBaseModel

`func (o *TrainSpec) HasBaseModel() bool`

HasBaseModel returns a boolean if a field has been set.

### GetDataset

`func (o *TrainSpec) GetDataset() TrainDataset`

GetDataset returns the Dataset field if non-nil, zero value otherwise.

### GetDatasetOk

`func (o *TrainSpec) GetDatasetOk() (*TrainDataset, bool)`

GetDatasetOk returns a tuple with the Dataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataset

`func (o *TrainSpec) SetDataset(v TrainDataset)`

SetDataset sets Dataset field to given value.

### HasDataset

`func (o *TrainSpec) HasDataset() bool`

HasDataset returns a boolean if a field has been set.

### GetEvaluation

`func (o *TrainSpec) GetEvaluation() TrainEvaluation`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *TrainSpec) GetEvaluationOk() (*TrainEvaluation, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *TrainSpec) SetEvaluation(v TrainEvaluation)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *TrainSpec) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### GetFrom

`func (o *TrainSpec) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TrainSpec) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TrainSpec) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *TrainSpec) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetObjective

`func (o *TrainSpec) GetObjective() TrainObjective`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *TrainSpec) GetObjectiveOk() (*TrainObjective, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *TrainSpec) SetObjective(v TrainObjective)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *TrainSpec) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetOutput

`func (o *TrainSpec) GetOutput() TrainOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TrainSpec) GetOutputOk() (*TrainOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TrainSpec) SetOutput(v TrainOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TrainSpec) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetProtect

`func (o *TrainSpec) GetProtect() TrainProtect`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *TrainSpec) GetProtectOk() (*TrainProtect, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *TrainSpec) SetProtect(v TrainProtect)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *TrainSpec) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetResources

`func (o *TrainSpec) GetResources() TrainResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *TrainSpec) GetResourcesOk() (*TrainResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *TrainSpec) SetResources(v TrainResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *TrainSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *TrainSpec) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *TrainSpec) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *TrainSpec) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *TrainSpec) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


