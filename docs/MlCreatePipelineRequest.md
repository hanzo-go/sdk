# MlCreatePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**AutoStart** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewMlCreatePipelineRequest

`func NewMlCreatePipelineRequest(name string, ) *MlCreatePipelineRequest`

NewMlCreatePipelineRequest instantiates a new MlCreatePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlCreatePipelineRequestWithDefaults

`func NewMlCreatePipelineRequestWithDefaults() *MlCreatePipelineRequest`

NewMlCreatePipelineRequestWithDefaults instantiates a new MlCreatePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MlCreatePipelineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlCreatePipelineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlCreatePipelineRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MlCreatePipelineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MlCreatePipelineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MlCreatePipelineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MlCreatePipelineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParameters

`func (o *MlCreatePipelineRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MlCreatePipelineRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MlCreatePipelineRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MlCreatePipelineRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetAutoStart

`func (o *MlCreatePipelineRequest) GetAutoStart() bool`

GetAutoStart returns the AutoStart field if non-nil, zero value otherwise.

### GetAutoStartOk

`func (o *MlCreatePipelineRequest) GetAutoStartOk() (*bool, bool)`

GetAutoStartOk returns a tuple with the AutoStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStart

`func (o *MlCreatePipelineRequest) SetAutoStart(v bool)`

SetAutoStart sets AutoStart field to given value.

### HasAutoStart

`func (o *MlCreatePipelineRequest) HasAutoStart() bool`

HasAutoStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


