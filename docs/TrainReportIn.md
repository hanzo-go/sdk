# TrainReportIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **[]interface{}** | Events are the task&#39;s new events, in order. | [optional] 
**Id** | Pointer to **string** | ID names the job, from the path. | [optional] 
**Lease** | Pointer to **string** | Lease is the task&#39;s lease, as the claim answered it. | [optional] 
**Task** | Pointer to **string** | Task is the reporting task. | [optional] 

## Methods

### NewTrainReportIn

`func NewTrainReportIn() *TrainReportIn`

NewTrainReportIn instantiates a new TrainReportIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainReportInWithDefaults

`func NewTrainReportInWithDefaults() *TrainReportIn`

NewTrainReportInWithDefaults instantiates a new TrainReportIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *TrainReportIn) GetEvents() []interface{}`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TrainReportIn) GetEventsOk() (*[]interface{}, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TrainReportIn) SetEvents(v []interface{})`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TrainReportIn) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *TrainReportIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainReportIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainReportIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainReportIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLease

`func (o *TrainReportIn) GetLease() string`

GetLease returns the Lease field if non-nil, zero value otherwise.

### GetLeaseOk

`func (o *TrainReportIn) GetLeaseOk() (*string, bool)`

GetLeaseOk returns a tuple with the Lease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease

`func (o *TrainReportIn) SetLease(v string)`

SetLease sets Lease field to given value.

### HasLease

`func (o *TrainReportIn) HasLease() bool`

HasLease returns a boolean if a field has been set.

### GetTask

`func (o *TrainReportIn) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TrainReportIn) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TrainReportIn) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *TrainReportIn) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


