# TrainAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**TrainJob**](TrainJob.md) | Job is the job it belongs to. | [optional] 
**Org** | Pointer to **string** | Org is the job&#39;s org. A platform executor reports for the task acting in it. | [optional] 
**Task** | Pointer to [**TrainClaimed**](TrainClaimed.md) | Task is the unit this executor now holds. | [optional] 

## Methods

### NewTrainAssignment

`func NewTrainAssignment() *TrainAssignment`

NewTrainAssignment instantiates a new TrainAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainAssignmentWithDefaults

`func NewTrainAssignmentWithDefaults() *TrainAssignment`

NewTrainAssignmentWithDefaults instantiates a new TrainAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *TrainAssignment) GetJob() TrainJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *TrainAssignment) GetJobOk() (*TrainJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *TrainAssignment) SetJob(v TrainJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *TrainAssignment) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetOrg

`func (o *TrainAssignment) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *TrainAssignment) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *TrainAssignment) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *TrainAssignment) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetTask

`func (o *TrainAssignment) GetTask() TrainClaimed`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *TrainAssignment) GetTaskOk() (*TrainClaimed, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *TrainAssignment) SetTask(v TrainClaimed)`

SetTask sets Task field to given value.

### HasTask

`func (o *TrainAssignment) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


