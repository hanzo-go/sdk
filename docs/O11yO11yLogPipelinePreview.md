# O11yO11yLogPipelinePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectorLogs** | Pointer to **[]string** | CollectorLogs is what the collector logged while simulating. | [optional] 
**Logs** | Pointer to [**[]O11yO11yLogRecord**](O11yO11yLogRecord.md) | Logs are the sample records after the pipelines ran over them. | [optional] 

## Methods

### NewO11yO11yLogPipelinePreview

`func NewO11yO11yLogPipelinePreview() *O11yO11yLogPipelinePreview`

NewO11yO11yLogPipelinePreview instantiates a new O11yO11yLogPipelinePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogPipelinePreviewWithDefaults

`func NewO11yO11yLogPipelinePreviewWithDefaults() *O11yO11yLogPipelinePreview`

NewO11yO11yLogPipelinePreviewWithDefaults instantiates a new O11yO11yLogPipelinePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectorLogs

`func (o *O11yO11yLogPipelinePreview) GetCollectorLogs() []string`

GetCollectorLogs returns the CollectorLogs field if non-nil, zero value otherwise.

### GetCollectorLogsOk

`func (o *O11yO11yLogPipelinePreview) GetCollectorLogsOk() (*[]string, bool)`

GetCollectorLogsOk returns a tuple with the CollectorLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorLogs

`func (o *O11yO11yLogPipelinePreview) SetCollectorLogs(v []string)`

SetCollectorLogs sets CollectorLogs field to given value.

### HasCollectorLogs

`func (o *O11yO11yLogPipelinePreview) HasCollectorLogs() bool`

HasCollectorLogs returns a boolean if a field has been set.

### GetLogs

`func (o *O11yO11yLogPipelinePreview) GetLogs() []O11yO11yLogRecord`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yO11yLogPipelinePreview) GetLogsOk() (*[]O11yO11yLogRecord, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yO11yLogPipelinePreview) SetLogs(v []O11yO11yLogRecord)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yO11yLogPipelinePreview) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


