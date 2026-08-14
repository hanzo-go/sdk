# O11yO11yLogPipelinePreviewIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]O11yO11yLogRecord**](O11yO11yLogRecord.md) | Logs are the sample records to transform. | [optional] 
**Pipelines** | Pointer to [**[]O11yO11yLogPipeline**](O11yO11yLogPipeline.md) | Pipelines are the pipelines to simulate, in order. | [optional] 

## Methods

### NewO11yO11yLogPipelinePreviewIn

`func NewO11yO11yLogPipelinePreviewIn() *O11yO11yLogPipelinePreviewIn`

NewO11yO11yLogPipelinePreviewIn instantiates a new O11yO11yLogPipelinePreviewIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPipelinePreviewInWithDefaults

`func NewO11yO11yLogPipelinePreviewInWithDefaults() *O11yO11yLogPipelinePreviewIn`

NewO11yO11yLogPipelinePreviewInWithDefaults instantiates a new O11yO11yLogPipelinePreviewIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yO11yLogPipelinePreviewIn) GetLogs() []O11yO11yLogRecord`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yO11yLogPipelinePreviewIn) GetLogsOk() (*[]O11yO11yLogRecord, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yO11yLogPipelinePreviewIn) SetLogs(v []O11yO11yLogRecord)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yO11yLogPipelinePreviewIn) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetPipelines

`func (o *O11yO11yLogPipelinePreviewIn) GetPipelines() []O11yO11yLogPipeline`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *O11yO11yLogPipelinePreviewIn) GetPipelinesOk() (*[]O11yO11yLogPipeline, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *O11yO11yLogPipelinePreviewIn) SetPipelines(v []O11yO11yLogPipeline)`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *O11yO11yLogPipelinePreviewIn) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


