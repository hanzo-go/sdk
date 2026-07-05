# MlPipelineStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**LogsUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewMlPipelineStep

`func NewMlPipelineStep() *MlPipelineStep`

NewMlPipelineStep instantiates a new MlPipelineStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlPipelineStepWithDefaults

`func NewMlPipelineStepWithDefaults() *MlPipelineStep`

NewMlPipelineStepWithDefaults instantiates a new MlPipelineStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MlPipelineStep) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlPipelineStep) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlPipelineStep) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlPipelineStep) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *MlPipelineStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlPipelineStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlPipelineStep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlPipelineStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *MlPipelineStep) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MlPipelineStep) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MlPipelineStep) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MlPipelineStep) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *MlPipelineStep) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *MlPipelineStep) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *MlPipelineStep) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *MlPipelineStep) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetLogsUrl

`func (o *MlPipelineStep) GetLogsUrl() string`

GetLogsUrl returns the LogsUrl field if non-nil, zero value otherwise.

### GetLogsUrlOk

`func (o *MlPipelineStep) GetLogsUrlOk() (*string, bool)`

GetLogsUrlOk returns a tuple with the LogsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsUrl

`func (o *MlPipelineStep) SetLogsUrl(v string)`

SetLogsUrl sets LogsUrl field to given value.

### HasLogsUrl

`func (o *MlPipelineStep) HasLogsUrl() bool`

HasLogsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


