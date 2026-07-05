# MlStartPipelineRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **map[string]string** | Override pipeline parameters for this run | [optional] 

## Methods

### NewMlStartPipelineRunRequest

`func NewMlStartPipelineRunRequest() *MlStartPipelineRunRequest`

NewMlStartPipelineRunRequest instantiates a new MlStartPipelineRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlStartPipelineRunRequestWithDefaults

`func NewMlStartPipelineRunRequestWithDefaults() *MlStartPipelineRunRequest`

NewMlStartPipelineRunRequestWithDefaults instantiates a new MlStartPipelineRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *MlStartPipelineRunRequest) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *MlStartPipelineRunRequest) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *MlStartPipelineRunRequest) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *MlStartPipelineRunRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


