# StudioGraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Steps** | Pointer to [**[]StudioSpec**](StudioSpec.md) | Steps are the graph&#39;s steps, 1 to 64 of them. | [optional] 

## Methods

### NewStudioGraph

`func NewStudioGraph() *StudioGraph`

NewStudioGraph instantiates a new StudioGraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioGraphWithDefaults

`func NewStudioGraphWithDefaults() *StudioGraph`

NewStudioGraphWithDefaults instantiates a new StudioGraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteps

`func (o *StudioGraph) GetSteps() []StudioSpec`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *StudioGraph) GetStepsOk() (*[]StudioSpec, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *StudioGraph) SetSteps(v []StudioSpec)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *StudioGraph) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


