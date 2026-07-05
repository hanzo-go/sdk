# EnginePipelineRunCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEnginePipelineRunCreate

`func NewEnginePipelineRunCreate() *EnginePipelineRunCreate`

NewEnginePipelineRunCreate instantiates a new EnginePipelineRunCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnginePipelineRunCreateWithDefaults

`func NewEnginePipelineRunCreateWithDefaults() *EnginePipelineRunCreate`

NewEnginePipelineRunCreateWithDefaults instantiates a new EnginePipelineRunCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnginePipelineRunCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnginePipelineRunCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnginePipelineRunCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnginePipelineRunCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *EnginePipelineRunCreate) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EnginePipelineRunCreate) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EnginePipelineRunCreate) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EnginePipelineRunCreate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


