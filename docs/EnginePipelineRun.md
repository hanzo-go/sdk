# EnginePipelineRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PipelineId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**Artifacts** | Pointer to [**[]EnginePipelineRunArtifactsInner**](EnginePipelineRunArtifactsInner.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnginePipelineRun

`func NewEnginePipelineRun() *EnginePipelineRun`

NewEnginePipelineRun instantiates a new EnginePipelineRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnginePipelineRunWithDefaults

`func NewEnginePipelineRunWithDefaults() *EnginePipelineRun`

NewEnginePipelineRunWithDefaults instantiates a new EnginePipelineRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnginePipelineRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnginePipelineRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnginePipelineRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnginePipelineRun) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPipelineId

`func (o *EnginePipelineRun) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *EnginePipelineRun) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *EnginePipelineRun) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *EnginePipelineRun) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetName

`func (o *EnginePipelineRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnginePipelineRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnginePipelineRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnginePipelineRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EnginePipelineRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnginePipelineRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnginePipelineRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnginePipelineRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParameters

`func (o *EnginePipelineRun) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EnginePipelineRun) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EnginePipelineRun) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EnginePipelineRun) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetArtifacts

`func (o *EnginePipelineRun) GetArtifacts() []EnginePipelineRunArtifactsInner`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *EnginePipelineRun) GetArtifactsOk() (*[]EnginePipelineRunArtifactsInner, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *EnginePipelineRun) SetArtifacts(v []EnginePipelineRunArtifactsInner)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *EnginePipelineRun) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetStartedAt

`func (o *EnginePipelineRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *EnginePipelineRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *EnginePipelineRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *EnginePipelineRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *EnginePipelineRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EnginePipelineRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EnginePipelineRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *EnginePipelineRun) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnginePipelineRun) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnginePipelineRun) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnginePipelineRun) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnginePipelineRun) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


