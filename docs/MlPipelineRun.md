# MlPipelineRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PipelineId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Metrics** | Pointer to **map[string]float32** |  | [optional] 
**Artifacts** | Pointer to [**[]MlPipelineRunArtifactsInner**](MlPipelineRunArtifactsInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DurationSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewMlPipelineRun

`func NewMlPipelineRun() *MlPipelineRun`

NewMlPipelineRun instantiates a new MlPipelineRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlPipelineRunWithDefaults

`func NewMlPipelineRunWithDefaults() *MlPipelineRun`

NewMlPipelineRunWithDefaults instantiates a new MlPipelineRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MlPipelineRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlPipelineRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlPipelineRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MlPipelineRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPipelineId

`func (o *MlPipelineRun) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *MlPipelineRun) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *MlPipelineRun) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *MlPipelineRun) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetStatus

`func (o *MlPipelineRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlPipelineRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlPipelineRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlPipelineRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParameters

`func (o *MlPipelineRun) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MlPipelineRun) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MlPipelineRun) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MlPipelineRun) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetMetrics

`func (o *MlPipelineRun) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MlPipelineRun) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MlPipelineRun) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MlPipelineRun) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetArtifacts

`func (o *MlPipelineRun) GetArtifacts() []MlPipelineRunArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *MlPipelineRun) GetArtifactsOk() (*[]MlPipelineRunArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *MlPipelineRun) SetArtifacts(v []MlPipelineRunArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *MlPipelineRun) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MlPipelineRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MlPipelineRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MlPipelineRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MlPipelineRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *MlPipelineRun) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *MlPipelineRun) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *MlPipelineRun) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *MlPipelineRun) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


