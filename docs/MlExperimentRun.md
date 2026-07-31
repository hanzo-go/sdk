# MlExperimentRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExperimentId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**Metrics** | Pointer to **map[string]float32** |  | [optional] 
**Artifacts** | Pointer to [**[]MlExperimentRunArtifactsInner**](MlExperimentRunArtifactsInner.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMlExperimentRun

`func NewMlExperimentRun() *MlExperimentRun`

NewMlExperimentRun instantiates a new MlExperimentRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlExperimentRunWithDefaults

`func NewMlExperimentRunWithDefaults() *MlExperimentRun`

NewMlExperimentRunWithDefaults instantiates a new MlExperimentRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MlExperimentRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlExperimentRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlExperimentRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MlExperimentRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExperimentId

`func (o *MlExperimentRun) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *MlExperimentRun) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *MlExperimentRun) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.

### HasExperimentId

`func (o *MlExperimentRun) HasExperimentId() bool`

HasExperimentId returns a boolean if a field has been set.

### GetStatus

`func (o *MlExperimentRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlExperimentRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlExperimentRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlExperimentRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParams

`func (o *MlExperimentRun) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MlExperimentRun) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MlExperimentRun) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *MlExperimentRun) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetMetrics

`func (o *MlExperimentRun) GetMetrics() map[string]float32`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MlExperimentRun) GetMetricsOk() (*map[string]float32, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MlExperimentRun) SetMetrics(v map[string]float32)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MlExperimentRun) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetArtifacts

`func (o *MlExperimentRun) GetArtifacts() []MlExperimentRunArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *MlExperimentRun) GetArtifactsOk() (*[]MlExperimentRunArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *MlExperimentRun) SetArtifacts(v []MlExperimentRunArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *MlExperimentRun) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetStartedAt

`func (o *MlExperimentRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MlExperimentRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MlExperimentRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MlExperimentRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MlExperimentRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MlExperimentRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MlExperimentRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MlExperimentRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


