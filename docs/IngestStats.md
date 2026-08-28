# IngestStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **bool** |  | [optional] 
**DocumentsIndexed** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**FilesIngested** | Pointer to **int32** |  | [optional] 
**FilesSkipped** | Pointer to **int32** |  | [optional] 
**IndexName** | Pointer to **string** |  | [optional] 
**Skipped** | Pointer to **[]string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 

## Methods

### NewIngestStats

`func NewIngestStats() *IngestStats`

NewIngestStats instantiates a new IngestStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestStatsWithDefaults

`func NewIngestStatsWithDefaults() *IngestStats`

NewIngestStatsWithDefaults instantiates a new IngestStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *IngestStats) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *IngestStats) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *IngestStats) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *IngestStats) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetDocumentsIndexed

`func (o *IngestStats) GetDocumentsIndexed() int32`

GetDocumentsIndexed returns the DocumentsIndexed field if non-nil, zero value otherwise.

### GetDocumentsIndexedOk

`func (o *IngestStats) GetDocumentsIndexedOk() (*int32, bool)`

GetDocumentsIndexedOk returns a tuple with the DocumentsIndexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsIndexed

`func (o *IngestStats) SetDocumentsIndexed(v int32)`

SetDocumentsIndexed sets DocumentsIndexed field to given value.

### HasDocumentsIndexed

`func (o *IngestStats) HasDocumentsIndexed() bool`

HasDocumentsIndexed returns a boolean if a field has been set.

### GetErrors

`func (o *IngestStats) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *IngestStats) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *IngestStats) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *IngestStats) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFilesIngested

`func (o *IngestStats) GetFilesIngested() int32`

GetFilesIngested returns the FilesIngested field if non-nil, zero value otherwise.

### GetFilesIngestedOk

`func (o *IngestStats) GetFilesIngestedOk() (*int32, bool)`

GetFilesIngestedOk returns a tuple with the FilesIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesIngested

`func (o *IngestStats) SetFilesIngested(v int32)`

SetFilesIngested sets FilesIngested field to given value.

### HasFilesIngested

`func (o *IngestStats) HasFilesIngested() bool`

HasFilesIngested returns a boolean if a field has been set.

### GetFilesSkipped

`func (o *IngestStats) GetFilesSkipped() int32`

GetFilesSkipped returns the FilesSkipped field if non-nil, zero value otherwise.

### GetFilesSkippedOk

`func (o *IngestStats) GetFilesSkippedOk() (*int32, bool)`

GetFilesSkippedOk returns a tuple with the FilesSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesSkipped

`func (o *IngestStats) SetFilesSkipped(v int32)`

SetFilesSkipped sets FilesSkipped field to given value.

### HasFilesSkipped

`func (o *IngestStats) HasFilesSkipped() bool`

HasFilesSkipped returns a boolean if a field has been set.

### GetIndexName

`func (o *IngestStats) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *IngestStats) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *IngestStats) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *IngestStats) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetSkipped

`func (o *IngestStats) GetSkipped() []string`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *IngestStats) GetSkippedOk() (*[]string, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *IngestStats) SetSkipped(v []string)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *IngestStats) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSource

`func (o *IngestStats) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IngestStats) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IngestStats) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *IngestStats) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStore

`func (o *IngestStats) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *IngestStats) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *IngestStats) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *IngestStats) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetWorkflowId

`func (o *IngestStats) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *IngestStats) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *IngestStats) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *IngestStats) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


