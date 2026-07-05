# EnginePipelineCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Spec** | **map[string]interface{}** | Pipeline spec (YAML or Argo Workflow) | 
**Parameters** | Pointer to [**[]EnginePipelineCreateParametersInner**](EnginePipelineCreateParametersInner.md) |  | [optional] 

## Methods

### NewEnginePipelineCreate

`func NewEnginePipelineCreate(name string, spec map[string]interface{}, ) *EnginePipelineCreate`

NewEnginePipelineCreate instantiates a new EnginePipelineCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnginePipelineCreateWithDefaults

`func NewEnginePipelineCreateWithDefaults() *EnginePipelineCreate`

NewEnginePipelineCreateWithDefaults instantiates a new EnginePipelineCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnginePipelineCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnginePipelineCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnginePipelineCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EnginePipelineCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnginePipelineCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnginePipelineCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnginePipelineCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpec

`func (o *EnginePipelineCreate) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EnginePipelineCreate) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EnginePipelineCreate) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetParameters

`func (o *EnginePipelineCreate) GetParameters() []EnginePipelineCreateParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *EnginePipelineCreate) GetParametersOk() (*[]EnginePipelineCreateParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *EnginePipelineCreate) SetParameters(v []EnginePipelineCreateParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *EnginePipelineCreate) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


