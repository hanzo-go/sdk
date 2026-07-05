# ConsoleScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**ObservationId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**ConfigId** | Pointer to **string** |  | [optional] 
**QueueId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewConsoleScore

`func NewConsoleScore() *ConsoleScore`

NewConsoleScore instantiates a new ConsoleScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleScoreWithDefaults

`func NewConsoleScoreWithDefaults() *ConsoleScore`

NewConsoleScoreWithDefaults instantiates a new ConsoleScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleScore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleScore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleScore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConsoleScore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTraceId

`func (o *ConsoleScore) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ConsoleScore) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ConsoleScore) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ConsoleScore) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetObservationId

`func (o *ConsoleScore) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *ConsoleScore) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *ConsoleScore) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *ConsoleScore) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### GetName

`func (o *ConsoleScore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleScore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleScore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsoleScore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ConsoleScore) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConsoleScore) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConsoleScore) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConsoleScore) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ConsoleScore) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ConsoleScore) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetStringValue

`func (o *ConsoleScore) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *ConsoleScore) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *ConsoleScore) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *ConsoleScore) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetDataType

`func (o *ConsoleScore) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *ConsoleScore) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *ConsoleScore) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *ConsoleScore) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetComment

`func (o *ConsoleScore) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConsoleScore) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConsoleScore) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConsoleScore) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSource

`func (o *ConsoleScore) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConsoleScore) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConsoleScore) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConsoleScore) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEnvironment

`func (o *ConsoleScore) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ConsoleScore) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ConsoleScore) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ConsoleScore) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetConfigId

`func (o *ConsoleScore) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ConsoleScore) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ConsoleScore) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ConsoleScore) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetQueueId

`func (o *ConsoleScore) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *ConsoleScore) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *ConsoleScore) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *ConsoleScore) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConsoleScore) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConsoleScore) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConsoleScore) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConsoleScore) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConsoleScore) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConsoleScore) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConsoleScore) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConsoleScore) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


