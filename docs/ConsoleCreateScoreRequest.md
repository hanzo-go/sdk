# ConsoleCreateScoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**ObservationId** | Pointer to **string** |  | [optional] 
**DatasetRunId** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Value** | **interface{}** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**QueueId** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleCreateScoreRequest

`func NewConsoleCreateScoreRequest(name string, value interface{}, ) *ConsoleCreateScoreRequest`

NewConsoleCreateScoreRequest instantiates a new ConsoleCreateScoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreateScoreRequestWithDefaults

`func NewConsoleCreateScoreRequestWithDefaults() *ConsoleCreateScoreRequest`

NewConsoleCreateScoreRequestWithDefaults instantiates a new ConsoleCreateScoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleCreateScoreRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleCreateScoreRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleCreateScoreRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleCreateScoreRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTraceId

`func (o *ConsoleCreateScoreRequest) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ConsoleCreateScoreRequest) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ConsoleCreateScoreRequest) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ConsoleCreateScoreRequest) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetSessionId

`func (o *ConsoleCreateScoreRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ConsoleCreateScoreRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ConsoleCreateScoreRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ConsoleCreateScoreRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetObservationId

`func (o *ConsoleCreateScoreRequest) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *ConsoleCreateScoreRequest) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *ConsoleCreateScoreRequest) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *ConsoleCreateScoreRequest) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetDatasetRunId

`func (o *ConsoleCreateScoreRequest) GetDatasetRunId() string`

GetDatasetRunId returns the DatasetRunId field if non-nil, zero value otherwise.

### GetDatasetRunIdOk

`func (o *ConsoleCreateScoreRequest) GetDatasetRunIdOk() (*string, bool)`

GetDatasetRunIdOk returns a tuple with the DatasetRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetRunId

`func (o *ConsoleCreateScoreRequest) SetDatasetRunId(v string)`

SetDatasetRunId sets DatasetRunId field to given value.

### HasDatasetRunId

`func (o *ConsoleCreateScoreRequest) HasDatasetRunId() bool`

HasDatasetRunId returns a boolean if a field has been set.

### GetName

`func (o *ConsoleCreateScoreRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleCreateScoreRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleCreateScoreRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ConsoleCreateScoreRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsoleCreateScoreRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsoleCreateScoreRequest) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ConsoleCreateScoreRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConsoleCreateScoreRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetComment

`func (o *ConsoleCreateScoreRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConsoleCreateScoreRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConsoleCreateScoreRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConsoleCreateScoreRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetMetadata

`func (o *ConsoleCreateScoreRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConsoleCreateScoreRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConsoleCreateScoreRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConsoleCreateScoreRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleCreateScoreRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleCreateScoreRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleCreateScoreRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleCreateScoreRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetQueueId

`func (o *ConsoleCreateScoreRequest) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *ConsoleCreateScoreRequest) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *ConsoleCreateScoreRequest) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *ConsoleCreateScoreRequest) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetDataType

`func (o *ConsoleCreateScoreRequest) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ConsoleCreateScoreRequest) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ConsoleCreateScoreRequest) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ConsoleCreateScoreRequest) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetConfigId

`func (o *ConsoleCreateScoreRequest) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ConsoleCreateScoreRequest) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ConsoleCreateScoreRequest) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ConsoleCreateScoreRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


