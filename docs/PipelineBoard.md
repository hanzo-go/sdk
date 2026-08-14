# PipelineBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pipelines** | Pointer to [**[]PipelineRow**](PipelineRow.md) | Pipelines are one per application in the caller&#39;s org. | [optional] 

## Methods

### NewPipelineBoard

`func NewPipelineBoard() *PipelineBoard`

NewPipelineBoard instantiates a new PipelineBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineBoardWithDefaults

`func NewPipelineBoardWithDefaults() *PipelineBoard`

NewPipelineBoardWithDefaults instantiates a new PipelineBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelines

`func (o *PipelineBoard) GetPipelines() []PipelineRow`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *PipelineBoard) GetPipelinesOk() (*[]PipelineRow, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *PipelineBoard) SetPipelines(v []PipelineRow)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *PipelineBoard) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


