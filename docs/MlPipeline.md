# MlPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]MlPipelineStep**](MlPipelineStep.md) |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMlPipeline

`func NewMlPipeline() *MlPipeline`

NewMlPipeline instantiates a new MlPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlPipelineWithDefaults

`func NewMlPipelineWithDefaults() *MlPipeline`

NewMlPipelineWithDefaults instantiates a new MlPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MlPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MlPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MlPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MlPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MlPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *MlPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *MlPipeline) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlPipeline) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlPipeline) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlPipeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *MlPipeline) GetSteps() []MlPipelineStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *MlPipeline) GetStepsOk() (*[]MlPipelineStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *MlPipeline) SetSteps(v []MlPipelineStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *MlPipeline) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetParameters

`func (o *MlPipeline) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MlPipeline) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MlPipeline) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MlPipeline) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MlPipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MlPipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MlPipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MlPipeline) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *MlPipeline) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MlPipeline) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MlPipeline) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MlPipeline) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MlPipeline) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MlPipeline) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MlPipeline) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MlPipeline) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


