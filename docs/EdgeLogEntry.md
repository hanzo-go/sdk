# EdgeLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 

## Methods

### NewEdgeLogEntry

`func NewEdgeLogEntry() *EdgeLogEntry`

NewEdgeLogEntry instantiates a new EdgeLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeLogEntryWithDefaults

`func NewEdgeLogEntryWithDefaults() *EdgeLogEntry`

NewEdgeLogEntryWithDefaults instantiates a new EdgeLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeLogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeLogEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EdgeLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *EdgeLogEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EdgeLogEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EdgeLogEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EdgeLogEntry) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventType

`func (o *EdgeLogEntry) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EdgeLogEntry) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EdgeLogEntry) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EdgeLogEntry) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetLevel

`func (o *EdgeLogEntry) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *EdgeLogEntry) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *EdgeLogEntry) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *EdgeLogEntry) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *EdgeLogEntry) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EdgeLogEntry) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EdgeLogEntry) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EdgeLogEntry) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExecutionId

`func (o *EdgeLogEntry) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EdgeLogEntry) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EdgeLogEntry) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EdgeLogEntry) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


